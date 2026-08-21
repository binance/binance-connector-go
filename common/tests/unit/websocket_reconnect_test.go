package commontests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

// TestHandleReadError_TriggersReconnect verifies that an unexpected read error
// (e.g. a temporary network interruption) now triggers a forced reconnection,
// instead of silently killing the listener. The signal is sent as
// non-graceful (false) so the read loop is restarted. Regression test for #96.
func TestHandleReadError_TriggersReconnect(t *testing.T) {
	conn := &common.WebSocketConnection{
		Id:            "test-connection",
		Connected:     common.OPEN,
		ReconnectChan: make(chan bool, 1),
		ErrorChan:     make(chan error, 1),
		Done:          make(chan struct{}),
	}

	conn.HandleReadError(errors.New("connection reset by peer"))

	select {
	case graceful := <-conn.ReconnectChan:
		if graceful {
			t.Error("expected a forced (non-graceful) reconnect on network drop")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("expected reconnect signal on unexpected read error")
	}
}

// TestReconnect_OnNetworkDrop verifies the client re-establishes the websocket
// connection after the underlying connection is dropped (simulated by closing
// the raw TCP connection without a clean websocket close frame). Regression
// test for #96.
func TestReconnect_OnNetworkDrop(t *testing.T) {
	var (
		connectionCount int
		mu              sync.Mutex
	)

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		mu.Lock()
		first := connectionCount == 0
		connectionCount++
		mu.Unlock()

		if first {
			// Simulate a sudden network drop: close the raw TCP connection
			// without sending a websocket close frame.
			_ = wsConn.UnderlyingConn().Close()
			return
		}

		// Keep subsequent (reconnect) connections alive.
		go func() {
			for {
				if _, _, e := wsConn.ReadMessage(); e != nil {
					return
				}
			}
		}()
	}))
	defer server.Close()

	u, _ := url.Parse(server.URL)
	u.Scheme = "ws"

	config := NewMockWebSocketConfig()
	config.basePath = u.String()
	config.reconnectDelay = 0

	wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
		APIConfig: &common.ConfigurationWebsocketApi{
			PoolSize: 1,
			Mode:     common.SINGLE,
		},
	})

	if err := wsc.Connect(config, "test-agent", []string{}); err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	require.Eventually(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		return connectionCount >= 2
	}, 3*time.Second, 50*time.Millisecond, "expected client to reconnect after a network drop")
}
