package binance_connector

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

type WebsocketStreamClient struct {
	Endpoint   string
	IsCombined bool
}

func NewWebsocketStreamClient(isCombined bool, baseURL ...string) *WebsocketStreamClient {
	// Set default base URL to production WS URL
	url := "wss://stream.binance.com:9443"

	if len(baseURL) > 0 {
		url = baseURL[0]
	}

	// Append to baseURL based on whether the client is for combined streams or not
	if isCombined {
		url += "/stream?streams="
	} else {
		url += "/ws"
	}

	return &WebsocketStreamClient{
		Endpoint:   url,
		IsCombined: isCombined,
	}
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

// PATCHED (2026-08-20, market_data_service fork):
//
// The upstream v0.8.0 implementation had three defects that could crash or
// leak the whole process:
//
//  1. The internal read goroutine performed `stopCh <- struct{}{}` on read
//     error. If the caller had close()d stopCh (the natural way to stop),
//     this panicked with "send on closed channel" inside a library-owned
//     goroutine, which no application-side recover() can catch.
//  2. The connection was never Close()d on stop, leaking the TCP connection
//     and the read goroutine until the server dropped the connection.
//  3. The `silent` flag was read/written from two goroutines without
//     synchronization (data race).
//
// The patched contract (mirrors the battle-tested go-binance pattern):
//   - To stop: the caller close()es stopCh. (A single send also still works
//     for backward compatibility.)
//   - The library then closes the connection itself, which unblocks
//     ReadMessage; the read loop exits and doneCh is closed.
//   - errHandler is not invoked for errors caused by a requested stop.
//   - The read goroutine never sends on stopCh.
var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	headers := http.Header{}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))
	c, _, err := Dialer.Dial(cfg.Endpoint, headers)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This goroutine exits when ReadMessage returns an error, either
		// because the connection failed or because the watcher goroutine
		// below closed the connection in response to a stop request.
		defer close(doneCh)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// silent suppresses errHandler once the caller has requested a stop:
		// the subsequent "use of closed network connection" read error is
		// expected and must not be reported.
		var silent atomic.Bool
		// Watcher: waits for a stop request (close or send on stopCh) or for
		// the read loop to finish (doneCh), then closes the connection. This
		// is the only place the library closes the connection, and it
		// guarantees no connection/goroutine leak on stop.
		go func() {
			select {
			case <-stopCh:
				silent.Store(true)
			case <-doneCh:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent.Load() {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				// PATCHED: close the connection on pong timeout so the read
				// loop unblocks and the error surfaces to errHandler instead
				// of the connection silently going stale.
				c.Close()
				return
			}
		}
	}()
}
