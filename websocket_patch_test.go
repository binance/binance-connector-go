package binance_connector

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var testUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// startEchoServer starts a websocket server that pushes a message every
// millisecond. If closeAfter > 0 the server force-closes the connection
// after that duration (simulating Binance dropping the stream).
func startEchoServer(closeAfter time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := testUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()
		var deadline <-chan time.Time
		if closeAfter > 0 {
			deadline = time.After(closeAfter)
		}
		tick := time.NewTicker(time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				if err := c.WriteMessage(websocket.TextMessage, []byte("x")); err != nil {
					return
				}
			case <-deadline:
				return // force-close from server side
			}
		}
	}))
}

func wsURL(s *httptest.Server) string {
	return "ws" + strings.TrimPrefix(s.URL, "http")
}

// TestProductionCrashSequence reproduces the exact production sequence of the
// 2026-08-20 crashes:
//
//  1. the app stops a worker and close()es stopCh (bnbGetTrade.stop /
//     bnbWebSocketObject.stop);
//  2. the connection dies some time later (Binance disconnect / network);
//  3. the library's read goroutine wakes up with a read error.
//
// Upstream v0.8.0 panics here with "send on closed channel" (websocket.go:88).
// The patched version must survive, close doneCh, and stay silent.
func TestProductionCrashSequence(t *testing.T) {
	srv := startEchoServer(300 * time.Millisecond) // server drops conn after 300ms
	defer srv.Close()

	var handled, errs atomic.Int64
	doneCh, stopCh, err := wsServe(newWsConfig(wsURL(srv)),
		func(message []byte) { handled.Add(1) },
		func(err error) { errs.Add(1) },
	)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}

	// Receive a few messages first.
	time.Sleep(50 * time.Millisecond)
	if handled.Load() == 0 {
		t.Fatalf("expected to receive messages before stop")
	}

	// Step 1: the app stops the worker.
	close(stopCh)

	// Step 2+3: the server force-close fires at t=300ms, long after stop.
	// On upstream v0.8.0 the read goroutine then executes
	// `stopCh <- struct{}{}` on the closed channel and the process panics.
	select {
	case <-doneCh:
		// patched: stop closes the conn immediately, read loop exits.
	case <-time.After(2 * time.Second):
		t.Fatalf("doneCh not closed within 2s after close(stopCh)")
	}

	// Give the doomed goroutine (if any) time to blow up.
	time.Sleep(500 * time.Millisecond)

	if errs.Load() != 0 {
		t.Fatalf("errHandler must not fire for a requested stop, got %d", errs.Load())
	}
}

// TestServerDropThenStop covers the other ordering seen in production: the
// connection errors first (errHandler -> MarkFailed -> restart), and the
// restart's stop() then close()es stopCh while/after the reader is erroring.
func TestServerDropThenStop(t *testing.T) {
	for i := 0; i < 50; i++ {
		srv := startEchoServer(10 * time.Millisecond)
		var errs atomic.Int64
		doneCh, stopCh, err := wsServe(newWsConfig(wsURL(srv)),
			func(message []byte) {},
			func(err error) { errs.Add(1) },
		)
		if err != nil {
			t.Fatalf("dial: %v", err)
		}
		// Race close(stopCh) against the server-side drop at t=10ms.
		time.Sleep(time.Duration(i%20) * time.Millisecond)
		close(stopCh)
		select {
		case <-doneCh:
		case <-time.After(2 * time.Second):
			t.Fatalf("iter %d: doneCh not closed", i)
		}
		srv.Close()
	}
}

// TestStopActuallyClosesConnection verifies the leak fix: after
// close(stopCh), the server must observe the client going away.
func TestStopActuallyClosesConnection(t *testing.T) {
	clientGone := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := testUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()
		defer close(clientGone)
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return // client closed the TCP connection
			}
		}
	}))
	defer srv.Close()

	doneCh, stopCh, err := wsServe(newWsConfig(wsURL(srv)),
		func(message []byte) {}, func(err error) {})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	close(stopCh)
	<-doneCh
	select {
	case <-clientGone:
		// connection really closed -> no fd/goroutine leak
	case <-time.After(2 * time.Second):
		t.Fatalf("server still sees the connection 2s after stop: connection leak")
	}
}

// TestBackwardCompatSendStop keeps the old documented usage working:
// a single send on stopCh must also stop the stream.
func TestBackwardCompatSendStop(t *testing.T) {
	srv := startEchoServer(0)
	defer srv.Close()

	doneCh, stopCh, err := wsServe(newWsConfig(wsURL(srv)),
		func(message []byte) {}, func(err error) {})
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	stopCh <- struct{}{}
	select {
	case <-doneCh:
	case <-time.After(2 * time.Second):
		t.Fatalf("doneCh not closed after send on stopCh")
	}
}
