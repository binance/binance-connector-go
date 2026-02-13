package common

import (
	"errors"
	"sync"

	"github.com/binance/binance-connector-go/common/v2/common"
)

type MockWebSocket struct {
	MessagesWritten [][]byte
	HasSentChan     chan int
	ReadChan        chan []byte
	Closed          bool
	ErrorChan       chan error
	mu              sync.Mutex
}

func NewMockWebSocket() *MockWebSocket {
	return &MockWebSocket{
		ReadChan:    make(chan []byte, 100),
		HasSentChan: make(chan int, 1),
	}
}

func (m *MockWebSocket) WriteMessage(_ int, msg []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.MessagesWritten = append(m.MessagesWritten, msg)
	m.HasSentChan <- 1
	return nil
}

func (m *MockWebSocket) ReadMessage() (int, []byte, error) {
	if m.Closed {
		return 0, nil, errors.New("websocket closed")
	}

	msg, ok := <-m.ReadChan
	if !ok {
		return 0, nil, errors.New("no messages queued")
	}

	return 1, msg, nil
}

func (m *MockWebSocket) QueueMessage(msg []byte) {
	select {
	case m.ReadChan <- msg:
	default:
		<-m.ReadChan
		m.ReadChan <- msg
	}
}

func (m *MockWebSocket) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.Closed {
		close(m.ReadChan)
		m.Closed = true
	}
	return nil
}

func (m *MockWebSocket) SetReadLimit(limit int64) {}

func SetupMockClient(id string) (*common.WebSocketConnection, *MockWebSocket, func()) {
	originalUUID := common.GenerateUUID
	common.GenerateUUID = func() string { return id }

	cleanup := func() {
		common.GenerateUUID = originalUUID
	}

	mockWS := NewMockWebSocket()
	conn := &common.WebSocketConnection{
		Id:                id,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockWS,
		SessionLogon:      false,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
		Connected:         common.OPEN,
	}

	return conn, mockWS, cleanup
}
