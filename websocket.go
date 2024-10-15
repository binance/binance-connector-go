package binance_connector

import (
	"fmt"
	"net/http"
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
		defer close(doneCh)
		defer c.Close()
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		go readMessages(c, handler, errHandler, doneCh, stopCh)
		<-stopCh
	}()
	return
}

func readMessages(c *websocket.Conn, handler WsHandler, errHandler ErrHandler, doneCh, stopCh chan struct{}) {
	for {
		select {
		case <-doneCh:
			return
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				select {
				case <-doneCh:
					return
				case stopCh <- struct{}{}:
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}
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
				return
			}
		}
	}()
}
