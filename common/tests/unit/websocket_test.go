package commontests

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"

	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/binance/binance-connector-go/common/common"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockWebSocketConn struct {
	mu               sync.Mutex
	readMessages     [][]byte
	readMessageTypes []int
	readIndex        int
	errors           []error
	writeMessages    []struct {
		messageType int
		data        []byte
	}
	closeCalled bool
	readLimit   int64
	writeError  error
}

func NewMockWebSocketConn() *MockWebSocketConn {
	return &MockWebSocketConn{
		readMessages:     make([][]byte, 0),
		readMessageTypes: make([]int, 0),
		writeMessages: make([]struct {
			messageType int
			data        []byte
		}, 0),
		readIndex: 0,
	}
}

func (m *MockWebSocketConn) AddReadMessage(messageType int, data []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.readMessages = append(m.readMessages, data)
	m.readMessageTypes = append(m.readMessageTypes, messageType)
}

func (m *MockWebSocketConn) GetWrittenMessages() []struct {
	messageType int
	data        []byte
} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.writeMessages
}

func (m *MockWebSocketConn) WriteMessage(messageType int, data []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.writeMessages = append(m.writeMessages, struct {
		messageType int
		data        []byte
	}{messageType, data})
	return m.writeError
}

func (m *MockWebSocketConn) AddError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errors = append(m.errors, err)
}

func (m *MockWebSocketConn) ReadMessage() (messageType int, p []byte, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.errors) > 0 {
		err = m.errors[0]
		m.errors = m.errors[1:]
		return 0, nil, err
	}

	if m.readIndex >= len(m.readMessages) {
		return 0, nil, io.EOF
	}

	messageType = m.readMessageTypes[m.readIndex]
	p = m.readMessages[m.readIndex]
	m.readIndex++
	return messageType, p, nil
}

func (m *MockWebSocketConn) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.closeCalled = true
	return nil
}

func (m *MockWebSocketConn) SetReadLimit(limit int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.readLimit = limit
}

func (m *MockWebSocketConn) SetWriteError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.writeError = err
}

type MockSigner struct {
	signFunc func([]byte) ([]byte, error)
}

func (m *MockSigner) Sign(data []byte) ([]byte, error) {
	if m.signFunc != nil {
		return m.signFunc(data)
	}
	return []byte("mock-signature"), nil
}

type MockWebSocketCommon struct {
	connections  []*common.WebSocketConnection
	getConnError error
	connectError error
	mu           sync.Mutex
}

func (m *MockWebSocketCommon) GetConnection() (*common.WebSocketConnection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.getConnError != nil {
		return nil, m.getConnError
	}
	if len(m.connections) == 0 {
		return nil, errors.New("no connections available")
	}
	return m.connections[0], nil
}

func (m *MockWebSocketCommon) Connect(cfg *common.ConfigurationWebsocketApi, userAgent string) error {
	return m.connectError
}

func createTestWebsocketAPI(mockConn *MockWebSocketConn) *common.WebsocketAPI {
	cfg := &common.ConfigurationWebsocketApi{
		ApiKey:         "test-api-key",
		ApiSecret:      "test-api-secret",
		Timeout:        5 * time.Second,
		ReconnectDelay: 1 * time.Second,
		PoolSize:       1,
	}

	wsConnection := &common.WebSocketConnection{
		Id:                  "test-conn-1",
		Connected:           common.OPEN,
		PendingMessages:     sync.Map{},
		StreamCallbackMap:   make(map[string][]func(map[string]interface{})),
		Websocket:           mockConn,
		SessionLogon:        false,
		Done:                make(chan struct{}),
		ErrorChan:           make(chan error, 1),
		StreamConnectionMap: make([]string, 0),
	}

	wsCommon := &common.WebSocketCommon{
		Connections:    []*common.WebSocketConnection{wsConnection},
		ReconnectTasks: make(map[string]chan struct{}),
		PoolSize:       1,
	}

	return &common.WebsocketAPI{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: make(map[string][]*common.WebSocketConnection),
		WsCommon:                  wsCommon,
	}
}

func createTestWebsocketStreams(mockConn *MockWebSocketConn) *common.WebsocketStreams {
	cfg := &common.ConfigurationWebsocketStreams{
		BasePath:       "wss://test.example.com",
		ReconnectDelay: 1 * time.Second,
		Compression:    false,
		PoolSize:       1,
	}

	wsConnection := &common.WebSocketConnection{
		Id:                  "test-stream-conn-1",
		Connected:           common.OPEN,
		PendingMessages:     sync.Map{},
		StreamCallbackMap:   make(map[string][]func(map[string]interface{})),
		Websocket:           mockConn,
		StreamConnectionMap: make([]string, 0),
		Done:                make(chan struct{}),
		ErrorChan:           make(chan error, 1),
	}

	wsCommon := &common.WebSocketCommon{
		Connections:    []*common.WebSocketConnection{wsConnection},
		ReconnectTasks: make(map[string]chan struct{}),
		PoolSize:       1,
	}

	return &common.WebsocketStreams{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: make(map[string][]*common.WebSocketConnection),
		WsCommon:                  wsCommon,
	}
}

type TestTradeData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    float64 `json:"qty"`
}

type TestTickerData struct {
	Symbol string `json:"s"`
	Price  string `json:"c"`
	Volume string `json:"v"`
}

func createStreamHandlerWrapperWithStreams(mockConn *MockWebSocketConn) *common.StreamHandlerWrapper {
	ws := createTestWebsocketStreams(mockConn)
	return &common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}
}

func createStreamHandlerWrapperWithAPI(mockConn *MockWebSocketConn) *common.StreamHandlerWrapper {
	api := createTestWebsocketAPI(mockConn)
	return &common.StreamHandlerWrapper{
		WebsocketAPI: api,
	}
}

func triggerStreamCallback(conn *common.WebSocketConnection, stream string, data map[string]interface{}) {
	time.Sleep(10 * time.Millisecond)

	if callbacks, exists := conn.StreamCallbackMap[stream]; exists {
		for _, cb := range callbacks {
			cb(data)
		}
	}
}

type MockWebSocketConfig struct {
	basePath       string
	timeUnit       common.TimeUnit
	httpsAgent     common.HTTPSAgent
	reconnectDelay time.Duration
	compression    bool
	proxy          *common.ProxyConfig
	tlsConfig      *tls.Config
}

type LoggingTransport struct {
	Base http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.Base.RoundTrip(req)
	return resp, err
}

func (m *MockWebSocketConfig) GetAgent() common.HTTPSAgent {
	return m.httpsAgent
}

func (m *MockWebSocketConfig) GetBasePath() string {
	return m.basePath
}

func (m *MockWebSocketConfig) GetReconnectDelay() time.Duration {
	return m.reconnectDelay
}

func (m *MockWebSocketConfig) GetCompression() bool {
	return m.compression
}

func (m *MockWebSocketConfig) GetProxy() *common.ProxyConfig {
	return m.proxy
}

func (m *MockWebSocketConfig) GetTLSConfig() *tls.Config {
	return m.tlsConfig
}

func (m *MockWebSocketConfig) GetTimeUnit() common.TimeUnit {
	return m.timeUnit
}

func NewMockWebSocketConfig() *MockWebSocketConfig {
	return &MockWebSocketConfig{
		basePath:       "/ws",
		httpsAgent:     &LoggingTransport{Base: http.DefaultTransport},
		timeUnit:       "",
		reconnectDelay: 5 * time.Second,
		compression:    false,
		proxy:          nil,
		tlsConfig:      nil,
	}
}

// =============================================================================
// WebSocketConnection Tests
// =============================================================================

func TestWebSocketConnection_Listen(t *testing.T) {
	mockConn := NewMockWebSocketConn()

	testMessage := map[string]interface{}{
		"type": "test",
		"data": "test data",
	}
	messageBytes, _ := json.Marshal(testMessage)

	mockConn.AddReadMessage(websocket.TextMessage, messageBytes)

	wsConn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	wsConn.StreamCallbackMap["test"] = []func(map[string]interface{}){
		func(msg map[string]interface{}) {
			if msg["type"] != "test" {
				t.Errorf("Expected type 'test', got %v", msg["type"])
			}
			if msg["data"] != "test data" {
				t.Errorf("Expected data 'test data', got %v", msg["data"])
			}
		},
	}

	wsConn.Listen()
}

func TestWebSocketConnection_Listen_WithError(t *testing.T) {
	mockConn := NewMockWebSocketConn()

	wsConn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	wsConn.Listen()

	select {
	case err := <-wsConn.ErrorChan:
		if err == nil {
			t.Error("Expected an error but got nil")
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("Timeout waiting for error")
	}
}

func TestHandleReadError_NormalClose(t *testing.T) {
	mockConn := NewMockWebSocketConn()

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := &websocket.CloseError{
		Code: websocket.CloseNormalClosure,
		Text: "normal closure",
	}

	conn.HandleReadError(err)

	select {
	case err := <-conn.ErrorChan:
		t.Errorf("Unexpected error sent to ErrorChan: %v", err)
	default:
	}

	if conn.Connected != common.CLOSED {
		t.Errorf("Expected connection to be closed, got %v", conn.Connected)
	}
}

func TestHandleReadError_IntentionalClose(t *testing.T) {
	mockConn := NewMockWebSocketConn()

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := errors.New("use of closed network connection")
	conn.HandleReadError(err)

	select {
	case err := <-conn.ErrorChan:
		t.Errorf("Unexpected error sent to ErrorChan: %v", err)
	default:
	}

	if conn.Connected != common.CLOSED {
		t.Errorf("Expected connection to be closed, got %v", conn.Connected)
	}
}

func TestHandleReadError_OtherErrors(t *testing.T) {
	mockConn := NewMockWebSocketConn()

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	testErr := errors.New("some other error")
	conn.HandleReadError(testErr)

	select {
	case err := <-conn.ErrorChan:
		if err != testErr {
			t.Errorf("Expected error %v, got %v", testErr, err)
		}
	default:
		t.Error("Expected error to be sent to ErrorChan")
	}
}

func TestProcessMessage_SuccessfulResponse(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{
		"status": 200.0,
		"data":   "success",
	}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:        "test-conn",
		Websocket: mockConn,
	}

	err := conn.ProcessMessage()

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func TestProcessMessage_ReadError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	mockConn.AddError(errors.New("read error"))

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil || err.Error() != "read error" {
		t.Errorf("Expected read error, got %v", err)
	}
}

func TestProcessMessage_ErrorParseMessage(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := []byte(`["error", "bad request"`)
	mockConn.AddReadMessage(websocket.TextMessage, testMessage)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "parse_message" {
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Op 'parse_message', got %s", wsErr.Op)
		t.Errorf("Expected parsing error, got %v", wsErr.Err)
		t.Errorf("Expected Message 'failed to parse websocket message', got %s", wsErr.Message)
	}
}

func TestProcessMessage_ErrorResponseCodeFloat64(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := map[string]interface{}{
		"status": 400,
		"error": map[string]interface{}{
			"code": float64(400),
		},
	}
	msgBytes, _ := json.Marshal(testMessage)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "error_response" {
		t.Errorf("Expected Message 0, got %v", wsErr.Code)
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Code '[400]', got %d", wsErr.Err)
		t.Errorf("Expected Err 'received error response from server', got %s", wsErr.Message)
		t.Errorf("Expected Op 'error_response', got %s", wsErr.Op)
	}
}

func TestProcessMessage_ErrorResponseCodeInt(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := map[string]interface{}{
		"status": 400,
		"error": map[string]interface{}{
			"code": int(400),
		},
	}
	msgBytes, _ := json.Marshal(testMessage)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "error_response" {
		t.Errorf("Expected Message 0, got %v", wsErr.Code)
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Code '[400]', got %d", wsErr.Err)
		t.Errorf("Expected Err 'received error response from server', got %s", wsErr.Message)
		t.Errorf("Expected Op 'error_response', got %s", wsErr.Op)
	}
}

func TestProcessMessage_ErrorResponseErrorMsg(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := map[string]interface{}{
		"status": 400,
		"error": map[string]interface{}{
			"code": float64(400),
			"msg":  "bad request",
		},
	}
	msgBytes, _ := json.Marshal(testMessage)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "error_response" {
		t.Errorf("Expected Message 0, got %v", wsErr.Code)
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Code '[400] bad request', got %d", wsErr.Err)
		t.Errorf("Expected Err 'received error response from server', got %s", wsErr.Message)
		t.Errorf("Expected Op 'error_response', got %s", wsErr.Op)
	}
}

func TestProcessMessage_ErrorResponseStatusInt(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := map[string]interface{}{
		"status": int(400),
		"error": map[string]interface{}{
			"code": float64(400),
			"msg":  "bad request",
		},
	}
	msgBytes, _ := json.Marshal(testMessage)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "error_response" {
		t.Errorf("Expected Message 0, got %v", wsErr.Code)
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Code '[400] bad request', got %d", wsErr.Err)
		t.Errorf("Expected Err 'received error response from server', got %s", wsErr.Message)
		t.Errorf("Expected Op 'error_response', got %s", wsErr.Op)
	}
}

func TestProcessMessage_ErrorResponseStatusFloat64(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	testMessage := map[string]interface{}{
		"status": float64(400),
		"error": map[string]interface{}{
			"code": float64(400),
			"msg":  "bad request",
		},
	}
	msgBytes, _ := json.Marshal(testMessage)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: make(map[string][]func(map[string]interface{})),
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Op != "error_response" {
		t.Errorf("Expected Message 0, got %v", wsErr.Code)
		t.Errorf("Expected Id 'test-connection', got %s", wsErr.ConnID)
		t.Errorf("Expected Code '[400] bad request', got %d", wsErr.Err)
		t.Errorf("Expected Err 'received error response from server', got %s", wsErr.Message)
		t.Errorf("Expected Op 'error_response', got %s", wsErr.Op)
	}
}

func TestProcessMessage_ResponseMessageHandled(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{"id": "0", "status": 200, "result": "response data"}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	mockStreamCallbackMap := map[string][]func(map[string]interface{}){"0": {
		func(data map[string]interface{}) {
		},
	}}

	// doneChan := make(chan []byte, 1)
	// mockPendingMsgs := sync.Map{}
	// mockPendingMsgs.Store("0", doneChan)

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		StreamCallbackMap: mockStreamCallbackMap,
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	doneChan := make(chan []byte, 1)
	conn.PendingMessages.Store("0", doneChan)

	go func() {
		err := conn.ProcessMessage()
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	}()

	select {
	case receivedMsg := <-doneChan:
		var receivedData map[string]interface{}
		if err := json.Unmarshal(receivedMsg, &receivedData); err != nil {
			t.Fatalf("Failed to unmarshal received message: %v", err)
		}

		if receivedData["id"] != "0" {
			t.Errorf("Expected id '0', got %v", receivedData["id"])
		}

		if receivedData["result"] != "response data" {
			t.Errorf("Expected result 'response data', got %v", receivedData["result"])
		}

	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for message on channel")
	}

	_, ok := <-doneChan
	if ok {
		t.Error("Expected channel to be closed")
	}

	if _, ok := conn.PendingMessages.Load("0"); ok {
		t.Error("Expected pending message '0' to be deleted")
	}
}

func TestProcessMessage_SubscriptionMessageHandled(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{"subscriptionId": 1.0, "data": "Value"}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	callbackInvoked := false
	var callbackData map[string]interface{}

	mockStreamCallbackMap := map[string][]func(map[string]interface{}){"1": {
		func(data map[string]interface{}) {
			callbackInvoked = true
			callbackData = data
			log.Println("Callback invoked with data:", data)
		},
	}}

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: mockStreamCallbackMap,
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if !callbackInvoked {
		t.Error("Expected callback to be invoked")
	}

	if callbackData["data"] != "Value" {
		t.Errorf("Expected data 'Value', got %v", callbackData["data"])
	}
}

func TestProcessMessage_SubscriptionMessageNotHandled(t *testing.T) {
	var logOutput bytes.Buffer
	originalLogger := log.Writer()
	log.SetOutput(&logOutput)
	defer func() {
		log.SetOutput(originalLogger)
	}()

	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{"subscriptionId": 0, "data": "Value"}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	callbackInvoked := false
	var callbackData map[string]interface{}

	mockStreamCallbackMap := map[string][]func(map[string]interface{}){"1": {
		func(data map[string]interface{}) {
			callbackInvoked = true
			callbackData = data
			log.Println("Callback invoked with data:", data)
		},
	}}

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: mockStreamCallbackMap,
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if callbackInvoked {
		t.Error("Expected callback not to be invoked")
	}

	if callbackData["data"] == "Value" {
		t.Errorf("Expected data 'Value', got %v", callbackData["data"])
	}

	logContent := logOutput.String()
	expectedLog := "No callback registered for subscription ID 0"
	if !strings.Contains(logContent, expectedLog) {
		t.Errorf("Expected log to contain '%s', got '%s'", expectedLog, logContent)
	}
}

func TestProcessMessage_StreamMessageHandled(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{"stream": "stream1", "data": "Value"}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	callbackInvoked := false
	var callbackData map[string]interface{}

	mockStreamCallbackMap := map[string][]func(map[string]interface{}){"stream1": {
		func(data map[string]interface{}) {
			callbackInvoked = true
			callbackData = data
			log.Println("Callback invoked with data:", data)
		},
	}}

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: mockStreamCallbackMap,
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if !callbackInvoked {
		t.Error("Expected callback to be invoked")
	}

	if callbackData["data"] != "Value" {
		t.Errorf("Expected data 'Value', got %v", callbackData["data"])
	}
}

func TestProcessMessage_StreamMessageNotHandled(t *testing.T) {
	var logOutput bytes.Buffer
	originalLogger := log.Writer()
	log.SetOutput(&logOutput)
	defer func() {
		log.SetOutput(originalLogger)
	}()

	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{"stream": "stream1", "data": "Value"}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	callbackInvoked := false
	var callbackData map[string]interface{}

	mockStreamCallbackMap := map[string][]func(map[string]interface{}){"stream2": {
		func(data map[string]interface{}) {
			callbackInvoked = true
			callbackData = data
			log.Println("Callback invoked with data:", data)
		},
	}}

	conn := &common.WebSocketConnection{
		Id:                "test-connection",
		Connected:         common.OPEN,
		PendingMessages:   sync.Map{},
		StreamCallbackMap: mockStreamCallbackMap,
		Websocket:         mockConn,
		Done:              make(chan struct{}),
		ErrorChan:         make(chan error, 1),
	}

	err := conn.ProcessMessage()

	time.Sleep(100 * time.Millisecond)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if callbackInvoked {
		t.Error("Expected callback not to be invoked")
	}

	if callbackData["data"] == "Value" {
		t.Errorf("Expected data 'Value', got %v", callbackData["data"])
	}

	logContent := logOutput.String()
	expectedLog := "No callback registered for stream stream1"
	if !strings.Contains(logContent, expectedLog) {
		t.Errorf("Expected log to contain '%s', got '%s'", expectedLog, logContent)
	}
}

func TestProcessMessage_MissingErrorFields(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	msg := map[string]interface{}{
		"error": map[string]interface{}{
			"msg": "error without code",
		},
		"status": 500.0,
	}
	msgBytes, _ := json.Marshal(msg)
	mockConn.AddReadMessage(websocket.TextMessage, msgBytes)

	conn := &common.WebSocketConnection{
		Id:        "test-conn",
		Websocket: mockConn,
	}

	err := conn.ProcessMessage()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	wsErr := err.(*common.WebSocketError)
	if wsErr.Err.Error() != "[0] error without code" {
		t.Errorf("Expected Err '[0] error without code', got %v", wsErr.Err)
	}
	if wsErr.Code != 500 {
		t.Errorf("Expected Code 500, got %d", wsErr.Code)
	}
}

func TestWebSocketConnection_IsHealthy(t *testing.T) {
	t.Run("returns true when connected and not done", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Connected: common.OPEN,
			Done:      make(chan struct{}),
		}

		if !conn.IsHealthy() {
			t.Error("Expected IsHealthy() to return true for connected, not done connection")
		}
	})

	t.Run("returns false when disconnected", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Connected: common.CLOSED,
			Done:      make(chan struct{}),
		}

		if conn.IsHealthy() {
			t.Error("Expected IsHealthy() to return false for disconnected connection")
		}
	})

	t.Run("returns false when done channel is closed", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Connected: common.OPEN,
			Done:      make(chan struct{}),
		}

		close(conn.Done)

		if conn.IsHealthy() {
			t.Error("Expected IsHealthy() to return false when done channel is closed")
		}
	})

	t.Run("returns false when connecting", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Connected: common.CONNECTING,
			Done:      make(chan struct{}),
		}

		if conn.IsHealthy() {
			t.Error("Expected IsHealthy() to return false for connecting connection")
		}
	})

	t.Run("returns false when closing", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Connected: common.CLOSING,
			Done:      make(chan struct{}),
		}

		if conn.IsHealthy() {
			t.Error("Expected IsHealthy() to return false for closing connection")
		}
	})
}

func TestConfigurationWrapper_IsAPI(t *testing.T) {
	tests := []struct {
		name     string
		wrapper  common.ConfigurationWrapper
		expected bool
	}{
		{
			name: "API config is nil",
			wrapper: common.ConfigurationWrapper{
				APIConfig: nil,
			},
			expected: false,
		},
		{
			name: "API config is not nil",
			wrapper: common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					ApiKey:    "test-key",
					ApiSecret: "test-secret",
				},
			},
			expected: true,
		},
		{
			name: "Only streams config is set",
			wrapper: common.ConfigurationWrapper{
				APIConfig:     nil,
				StreamsConfig: &common.ConfigurationWebsocketStreams{},
			},
			expected: false,
		},
		{
			name: "Both configs are set",
			wrapper: common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					BasePath: "test-path",
				},
				StreamsConfig: &common.ConfigurationWebsocketStreams{
					BasePath: "test-path",
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wrapper.IsAPI(); got != tt.expected {
				t.Errorf("IsAPI() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConfigurationWrapper_IsStreams(t *testing.T) {
	tests := []struct {
		name     string
		wrapper  common.ConfigurationWrapper
		expected bool
	}{
		{
			name: "Streams config is nil",
			wrapper: common.ConfigurationWrapper{
				StreamsConfig: nil,
			},
			expected: false,
		},
		{
			name: "Streams config is not nil",
			wrapper: common.ConfigurationWrapper{
				StreamsConfig: &common.ConfigurationWebsocketStreams{
					BasePath: "test-path",
				},
			},
			expected: true,
		},
		{
			name: "Only API config is set",
			wrapper: common.ConfigurationWrapper{
				APIConfig:     &common.ConfigurationWebsocketApi{},
				StreamsConfig: nil,
			},
			expected: false,
		},
		{
			name: "Both configs are set",
			wrapper: common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					BasePath: "test-path",
				},
				StreamsConfig: &common.ConfigurationWebsocketStreams{
					BasePath: "test-path",
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wrapper.IsStreams(); got != tt.expected {
				t.Errorf("IsStreams() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// =============================================================================
// WebSocketCommon Tests
// =============================================================================

func TestNewWebSocketCommon(t *testing.T) {
	tests := []struct {
		name          string
		cfg           *common.ConfigurationWrapper
		expectedError bool
		expectedPool  int
		expectedMode  common.WebsocketMode
	}{
		{
			name: "valid API config",
			cfg: &common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					PoolSize: 3,
					Mode:     "api",
				},
			},
			expectedError: false,
			expectedPool:  3,
			expectedMode:  "api",
		},
		{
			name: "valid streams config",
			cfg: &common.ConfigurationWrapper{
				StreamsConfig: &common.ConfigurationWebsocketStreams{
					PoolSize: 5,
					Mode:     "streams",
				},
			},
			expectedError: false,
			expectedPool:  5,
			expectedMode:  "streams",
		},
		{
			name:          "nil config",
			cfg:           nil,
			expectedError: true,
		},
		{
			name: "both configs nil",
			cfg: &common.ConfigurationWrapper{
				APIConfig:     nil,
				StreamsConfig: nil,
			},
			expectedError: true,
		},
		{
			name: "both configs set (should prefer API)",
			cfg: &common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					PoolSize: 2,
					Mode:     "api",
				},
				StreamsConfig: &common.ConfigurationWebsocketStreams{
					PoolSize: 4,
					Mode:     "streams",
				},
			},
			expectedError: false,
			expectedPool:  2,
			expectedMode:  "api",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := common.NewWebSocketCommon(tt.cfg)

			if tt.expectedError {
				if err == nil {
					t.Errorf("NewWebSocketCommon() expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("NewWebSocketCommon() unexpected error: %v", err)
				return
			}

			if got.PoolSize != tt.expectedPool {
				t.Errorf("NewWebSocketCommon() PoolSize = %v, want %v", got.PoolSize, tt.expectedPool)
			}

			if got.Mode != tt.expectedMode {
				t.Errorf("NewWebSocketCommon() Mode = %v, want %v", got.Mode, tt.expectedMode)
			}

			if len(got.Connections) != tt.expectedPool {
				t.Errorf("NewWebSocketCommon() Connections length = %v, want %v", len(got.Connections), tt.expectedPool)
			}

			if got.RoundRobinIndex != 0 {
				t.Errorf("NewWebSocketCommon() RoundRobinIndex = %v, want 0", got.RoundRobinIndex)
			}
		})
	}
}

func TestWebSocketCommon_initializePool(t *testing.T) {
	tests := []struct {
		name          string
		poolSize      int
		expectedError bool
	}{
		{
			name:          "zero pool size",
			poolSize:      0,
			expectedError: false,
		},
		{
			name:          "single connection",
			poolSize:      1,
			expectedError: false,
		},
		{
			name:          "multiple connections",
			poolSize:      5,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := &common.ConfigurationWrapper{
				APIConfig: &common.ConfigurationWebsocketApi{
					PoolSize: tt.poolSize,
					Mode:     "test-mode",
				},
			}

			w, err := common.NewWebSocketCommon(wrapper)

			if tt.expectedError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(w.Connections) != tt.poolSize {
				t.Errorf("created %d connections, want %d", len(w.Connections), tt.poolSize)
			}

			for i, conn := range w.Connections {
				if conn.Connected != common.CONNECTING {
					t.Errorf("connection %d: Connected = %v, want %v", i, conn.Connected, common.CONNECTING)
				}

				if conn.Id == "" {
					t.Errorf("connection %d: Id is empty", i)
				}

				var mapInitialized bool
				conn.PendingMessages.Store("test-key", "test-value")
				if _, ok := conn.PendingMessages.Load("test-key"); ok {
					mapInitialized = true
					conn.PendingMessages.Delete("test-key")
				}
				if !mapInitialized {
					t.Errorf("connection %d: PendingMessages not initialized", i)
				}

				if conn.StreamCallbackMap == nil {
					t.Errorf("connection %d: StreamCallbackMap not initialized", i)
				}

				if conn.StreamConnectionMap == nil {
					t.Errorf("connection %d: StreamConnectionMap not initialized", i)
				}

				if conn.ErrorChan == nil {
					t.Errorf("connection %d: ErrorChan not initialized", i)
				}

				if conn.Done == nil {
					t.Errorf("connection %d: Done channel not initialized", i)
				}
			}
		})
	}
}

func TestWebSocketCommon_CreateWebSocketDialer(t *testing.T) {
	t.Run("creates dialer with default settings", func(t *testing.T) {
		wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
			APIConfig: &common.ConfigurationWebsocketApi{
				PoolSize: 1,
				Mode:     common.SINGLE,
			},
		})

		config := NewMockWebSocketConfig()
		dialer := wsc.CreateWebSocketDialer(config)

		if dialer.ReadBufferSize != 16*1024*1024 {
			t.Errorf("Expected ReadBufferSize 16MB, got %d", dialer.ReadBufferSize)
		}
		if dialer.WriteBufferSize != 16*1024*1024 {
			t.Errorf("Expected WriteBufferSize 16MB, got %d", dialer.WriteBufferSize)
		}
		if dialer.HandshakeTimeout != 45*time.Second {
			t.Errorf("Expected HandshakeTimeout 45s, got %v", dialer.HandshakeTimeout)
		}
	})

	t.Run("respects compression setting", func(t *testing.T) {
		wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
			APIConfig: &common.ConfigurationWebsocketApi{
				PoolSize: 1,
				Mode:     common.SINGLE,
			},
		})

		config := NewMockWebSocketConfig()
		config.compression = true

		dialer := wsc.CreateWebSocketDialer(config)
		if !dialer.EnableCompression {
			t.Error("Expected compression to be enabled")
		}
	})
}

func TestWebSocketCommon_Connect(t *testing.T) {
	t.Run("single mode connection", func(t *testing.T) {
		upgrader := websocket.Upgrader{}
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return
			}
			defer func() { _ = conn.Close() }()
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					return
				}
			}
		}))
		defer s.Close()

		u, _ := url.Parse(s.URL)
		u.Scheme = "ws"

		config := NewMockWebSocketConfig()
		config.basePath = u.String()

		wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
			APIConfig: &common.ConfigurationWebsocketApi{
				PoolSize: 1,
				Mode:     common.SINGLE,
			},
		})

		err := wsc.Connect(config, "test-agent")

		if err != nil {
			t.Fatalf("Connect failed: %v", err)
		}

		if len(wsc.Connections) != 1 {
			t.Errorf("Expected 1 connection, got %d", len(wsc.Connections))
		}

		if wsc.Connections[0].Websocket == nil {
			t.Error("Websocket connection should not be nil")
		}
	})

	t.Run("pool mode connection", func(t *testing.T) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}

		var connectionCount int
		var mu sync.Mutex

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				t.Logf("WebSocket upgrade failed: %v", err)
				return
			}
			defer func() { _ = conn.Close() }()

			mu.Lock()
			connectionCount++
			mu.Unlock()

			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					return
				}
			}
		}))
		defer s.Close()

		u, _ := url.Parse(s.URL)
		u.Scheme = "ws"

		wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
			APIConfig: &common.ConfigurationWebsocketApi{
				PoolSize: 3,
				Mode:     common.POOL,
			},
		})

		if len(wsc.Connections) == 0 {
			wsc.Connections = make([]*common.WebSocketConnection, 3)
			for i := 0; i < 3; i++ {
				wsc.Connections[i] = &common.WebSocketConnection{}
			}
		}

		config := NewMockWebSocketConfig()
		config.basePath = u.String()

		err := wsc.Connect(config, "test-agent")

		if err != nil {
			t.Fatalf("Connect failed: %v", err)
		}

		if len(wsc.Connections) != 3 {
			t.Errorf("Expected 3 connections, got %d", len(wsc.Connections))
		}

		for i, conn := range wsc.Connections {
			if conn.Websocket == nil {
				t.Errorf("Connection %d has nil WebSocket", i)
			}
			if conn.Connected != common.OPEN {
				t.Errorf("Connection %d is not OPEN", i)
			}
		}

		mu.Lock()
		if connectionCount != 3 {
			t.Errorf("Expected 3 server connections, got %d", connectionCount)
		}
		mu.Unlock()
	})
}

func TestWebSocketCommon_KeepAlive(t *testing.T) {
	t.Run("should not reconnect when connection is healthy", func(t *testing.T) {
		wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
			APIConfig: &common.ConfigurationWebsocketApi{
				PoolSize: 1,
				Mode:     common.SINGLE,
			},
		})

		connection := &common.WebSocketConnection{
			Id:        "test-connection",
			Connected: common.OPEN,
			Done:      make(chan struct{}),
		}

		go wsc.KeepAlive(connection, NewMockWebSocketConfig(), "test-agent")
		time.Sleep(100 * time.Millisecond)
		assert.Equal(t, common.OPEN, connection.Connected)
	})

	t.Run("should not reconnect when connection is closing or closed", func(t *testing.T) {
		testCases := []struct {
			name  string
			state common.WebsocketStatus
		}{
			{"closing", common.CLOSING},
			{"closed", common.CLOSED},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
					APIConfig: &common.ConfigurationWebsocketApi{
						PoolSize: 1,
						Mode:     common.SINGLE,
					},
				})

				connection := &common.WebSocketConnection{
					Id:        "test-connection",
					Connected: tc.state,
					Done:      make(chan struct{}),
				}

				go wsc.KeepAlive(connection, NewMockWebSocketConfig(), "test-agent")

				time.Sleep(100 * time.Millisecond)

				assert.Equal(t, tc.state, connection.Connected)
			})
		}
	})
}

func TestWebsocketCommon_Ping(t *testing.T) {
	wsc, _ := common.NewWebSocketCommon(&common.ConfigurationWrapper{
		APIConfig: &common.ConfigurationWebsocketApi{
			PoolSize: 1,
			Mode:     common.SINGLE,
		},
	})

	t.Run("successful ping", func(t *testing.T) {
		mockConn := NewMockWebSocketConn()
		conn := &common.WebSocketConnection{
			Id:        "test-connection",
			Connected: common.OPEN,
			Websocket: mockConn,
			Done:      make(chan struct{}),
		}

		err := wsc.Ping(conn)
		assert.NoError(t, err)

		writtenMsgs := mockConn.GetWrittenMessages()
		assert.Len(t, writtenMsgs, 1)
		assert.Equal(t, websocket.TextMessage, writtenMsgs[0].messageType)
		assert.Equal(t, []byte("ping"), writtenMsgs[0].data)
	})

	t.Run("nil websocket", func(t *testing.T) {
		conn := &common.WebSocketConnection{
			Id:        "test-connection",
			Connected: common.OPEN,
			Websocket: nil,
			Done:      make(chan struct{}),
		}

		err := wsc.Ping(conn)
		assert.EqualError(t, err, "WebSocket connection is not established")
	})

	t.Run("connection not open", func(t *testing.T) {
		mockConn := NewMockWebSocketConn()
		conn := &common.WebSocketConnection{
			Id:        "test-connection",
			Connected: common.CLOSED,
			Websocket: mockConn,
			Done:      make(chan struct{}),
		}

		err := wsc.Ping(conn)
		assert.EqualError(t, err, "WebSocket connection is not open")
	})

	t.Run("write error", func(t *testing.T) {
		mockConn := NewMockWebSocketConn()
		mockConn.SetWriteError(errors.New("write failed"))
		conn := &common.WebSocketConnection{
			Id:        "test-connection",
			Connected: common.OPEN,
			Websocket: mockConn,
			Done:      make(chan struct{}),
		}

		err := wsc.Ping(conn)
		assert.EqualError(t, err, "write failed")
	})
}

// =============================================================================
// WebsocketAPI Tests
// =============================================================================

func TestNewWebsocketAPI(t *testing.T) {
	cfg := &common.ConfigurationWebsocketApi{
		ApiKey:    "test-key",
		ApiSecret: "test-secret",
		PoolSize:  2,
	}

	api, err := common.NewWebsocketAPI(cfg)

	require.NoError(t, err)
	assert.NotNil(t, api)
	assert.Equal(t, cfg, api.Cfg)
	assert.NotNil(t, api.GlobalStreamConnectionMap)
	assert.NotNil(t, api.WsCommon)
}

func TestWebsocketAPI_Connect(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	err := api.Connect("test-user-agent")

	assert.NoError(t, err)
}

func TestSendMessage_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	responseMsg := map[string]interface{}{
		"id":     "test-id-123",
		"result": map[string]interface{}{"status": "ok"},
	}
	responseMsgBytes, _ := json.Marshal(responseMsg)

	payload := map[string]any{
		"method": "test.method",
		"params": map[string]any{
			"id":    "test-id-123",
			"param": "value",
		},
	}

	sendParams := common.SendParams{
		WithAPIKey:       true,
		WithSessionLogon: false,
		Signed:           false,
	}

	respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)

	require.NoError(t, err)
	require.NotNil(t, respChan)
	require.NotNil(t, errChan)

	go func() {
		time.Sleep(50 * time.Millisecond)
		conn := api.WsCommon.Connections[0]
		respChanInterface, ok := conn.PendingMessages.Load("test-id-123")
		if !ok {
			t.Logf("Warning: pending message not found for test-id-123")
			return
		}
		respChanTyped := respChanInterface.(chan []byte)
		respChanTyped <- responseMsgBytes
	}()

	select {
	case resp := <-respChan:
		require.NotNil(t, resp)
		if resp.Raw != nil {
			assert.NotNil(t, resp.Raw)
		}
	case err := <-errChan:
		t.Fatalf("Unexpected error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for response")
	}

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.GreaterOrEqual(t, len(writtenMsgs), 1)
}

func TestSendMessage_WithSignature_HMAC(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	api.Cfg.Signer = &MockSigner{
		signFunc: func(data []byte) ([]byte, error) {
			return []byte("hmac-signature-bytes"), nil
		},
	}

	responseMsg := map[string]interface{}{
		"id":     "test-id-456",
		"result": map[string]interface{}{"status": "ok"},
	}
	responseMsgBytes, _ := json.Marshal(responseMsg)

	payload := map[string]any{
		"method": "test.signed.method",
		"params": map[string]any{
			"id": "test-id-456",
		},
	}

	sendParams := common.SendParams{
		WithAPIKey: true,
		Signed:     true,
	}

	_, _, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)

	require.NoError(t, err)

	conn, _ := api.WsCommon.GetConnection()
	respChanInterface, _ := conn.PendingMessages.Load("test-id-456")
	respChanTyped := respChanInterface.(chan []byte)
	respChanTyped <- responseMsgBytes

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 1, len(writtenMsgs))

	var sentMsg map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &sentMsg)
	require.NoError(t, err)

	params := sentMsg["params"].(map[string]interface{})
	assert.Contains(t, params, "signature")
	assert.Contains(t, params, "timestamp")
	assert.Contains(t, params, "apiKey")
}

func TestSendMessage_WithSessionLogon(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	responseMsg := map[string]interface{}{
		"id":     "logon-id",
		"result": map[string]interface{}{"session": "active"},
	}
	responseMsgBytes, _ := json.Marshal(responseMsg)

	payload := map[string]any{
		"method": "session.logon",
		"params": map[string]any{
			"id": "logon-id",
		},
	}

	sendParams := common.SendParams{
		WithAPIKey:       true,
		WithSessionLogon: true,
		Signed:           false,
	}

	respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)
	require.NoError(t, err)

	conn, _ := api.WsCommon.GetConnection()
	assert.False(t, conn.SessionLogon)

	respChanInterface, _ := conn.PendingMessages.Load("logon-id")
	respChanTyped := respChanInterface.(chan []byte)
	respChanTyped <- responseMsgBytes

	select {
	case <-respChan:
		assert.True(t, conn.SessionLogon)
		assert.NotNil(t, conn.SessionLogonRequest)
	case <-errChan:
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}
}

func TestSendMessage_ErrorResponse(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	errorMsg := map[string]interface{}{
		"id": "error-id",
		"error": map[string]interface{}{
			"code": -1000,
			"msg":  "Invalid request",
		},
	}
	errorMsgBytes, _ := json.Marshal(errorMsg)

	payload := map[string]any{
		"method": "test.method",
		"params": map[string]any{
			"id": "error-id",
		},
	}

	sendParams := common.SendParams{}

	respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)
	require.NoError(t, err)

	go func() {
		time.Sleep(50 * time.Millisecond)
		conn := api.WsCommon.Connections[0]
		respChanInterface, ok := conn.PendingMessages.Load("error-id")
		if ok {
			respChanTyped := respChanInterface.(chan []byte)
			respChanTyped <- errorMsgBytes
		}
	}()

	select {
	case <-respChan:
		t.Fatal("should not receive response")
	case err := <-errChan:
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "1000")
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}
}

func TestSendMessage_StatusError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	statusMsg := map[string]interface{}{
		"id":     "status-id",
		"status": float64(404),
	}
	statusMsgBytes, _ := json.Marshal(statusMsg)

	payload := map[string]any{
		"method": "test.method",
		"params": map[string]any{
			"id": "status-id",
		},
	}

	sendParams := common.SendParams{}

	respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)
	require.NoError(t, err)

	conn, _ := api.WsCommon.GetConnection()
	respChanInterface, _ := conn.PendingMessages.Load("status-id")
	respChanTyped := respChanInterface.(chan []byte)
	respChanTyped <- statusMsgBytes

	select {
	case <-respChan:
		t.Fatal("should not receive response")
	case err := <-errChan:
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status 404")
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}
}

func TestSendMessage_Timeout(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)
	api.Cfg.Timeout = 100 * time.Millisecond

	payload := map[string]any{
		"method": "test.method",
		"params": map[string]any{
			"id": "timeout-id",
		},
	}

	sendParams := common.SendParams{}

	respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)
	require.NoError(t, err)

	select {
	case <-respChan:
		t.Fatal("should not receive response")
	case err := <-errChan:
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "timeout")
	case <-time.After(1 * time.Second):
		t.Fatal("test timeout")
	}
}

func TestSendMessage_InvalidParams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	payload := map[string]any{
		"method": "test.method",
		"params": "invalid-params",
	}

	sendParams := common.SendParams{}

	_, _, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not a map")
}

func TestSendMessage_WriteError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	mockConn.SetWriteError(errors.New("write failed"))
	api := createTestWebsocketAPI(mockConn)

	payload := map[string]any{
		"method": "test.method",
		"params": map[string]any{
			"id": "write-error-id",
		},
	}

	sendParams := common.SendParams{}

	_, _, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write failed")
}

func TestWebsocketAPI_Subscribe(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	streamID := "test-stream-1"
	err := api.Subscribe(streamID)

	assert.NoError(t, err)
	assert.Contains(t, api.GlobalStreamConnectionMap, streamID)
	assert.Equal(t, 1, len(api.GlobalStreamConnectionMap[streamID]))

	conn, _ := api.WsCommon.GetConnection()
	assert.Contains(t, conn.StreamConnectionMap, streamID)
}

func TestWebsocketAPI_Subscribe_MultipleStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	streams := []string{"stream-1", "stream-2", "stream-3"}

	for _, streamID := range streams {
		err := api.Subscribe(streamID)
		assert.NoError(t, err)
	}

	conn, _ := api.WsCommon.GetConnection()
	assert.Equal(t, len(streams), len(conn.StreamConnectionMap))

	for _, streamID := range streams {
		assert.Contains(t, api.GlobalStreamConnectionMap, streamID)
	}
}

func TestWebsocketAPI_Unsubscribe(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	streamID := "test-stream-1"

	err := api.Subscribe(streamID)
	require.NoError(t, err)
	assert.Contains(t, api.GlobalStreamConnectionMap, streamID)

	err = api.Unsubscribe(streamID)
	assert.NoError(t, err)
	assert.NotContains(t, api.GlobalStreamConnectionMap, streamID)
}

func TestWebsocketAPI_Unsubscribe_NonExistent(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	streamID := "non-existent-stream"

	err := api.Unsubscribe(streamID)
	assert.NoError(t, err)
}

func TestWebsocketAPI_CloseWebSocketConnection(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	conn := api.WsCommon.Connections[0]
	require.NotNil(t, conn)
	require.NotNil(t, conn.Websocket)

	err := api.CloseWebSocketConnection()
	assert.NoError(t, err)

	assert.Equal(t, common.CLOSING, conn.Connected)

	mockConn.mu.Lock()
	closeCalled := mockConn.closeCalled
	mockConn.mu.Unlock()
	assert.True(t, closeCalled)
}

func TestWebsocketAPI_CloseWebSocketConnection_MultipleConnections(t *testing.T) {
	mockConn1 := NewMockWebSocketConn()
	mockConn2 := NewMockWebSocketConn()

	cfg := &common.ConfigurationWebsocketApi{
		ApiKey:    "test-key",
		ApiSecret: "test-secret",
		Timeout:   5 * time.Second,
		PoolSize:  2,
	}

	conn1 := &common.WebSocketConnection{
		Id:        "conn-1",
		Connected: common.OPEN,
		Websocket: mockConn1,
	}

	conn2 := &common.WebSocketConnection{
		Id:        "conn-2",
		Connected: common.OPEN,
		Websocket: mockConn2,
	}

	wsCommon := &common.WebSocketCommon{
		Connections: []*common.WebSocketConnection{conn1, conn2},
	}

	api := &common.WebsocketAPI{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: make(map[string][]*common.WebSocketConnection),
		WsCommon:                  wsCommon,
	}

	err := api.CloseWebSocketConnection()
	assert.NoError(t, err)
	assert.True(t, mockConn1.closeCalled)
	assert.True(t, mockConn2.closeCalled)
	assert.Equal(t, common.CLOSING, conn1.Connected)
	assert.Equal(t, common.CLOSING, conn2.Connected)
}

func TestSendMessage_ConcurrentRequests(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	api := createTestWebsocketAPI(mockConn)

	numRequests := 10
	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func(index int) {
			defer wg.Done()

			responseMsg := map[string]interface{}{
				"id":     "test-id-" + string(rune(index)),
				"result": map[string]interface{}{"index": index},
			}
			responseMsgBytes, _ := json.Marshal(responseMsg)

			payload := map[string]any{
				"method": "test.method",
				"params": map[string]any{
					"index": index,
				},
			}

			sendParams := common.SendParams{}

			respChan, errChan, err := common.SendMessage[map[string]interface{}](api, payload, sendParams)
			if err != nil {
				t.Errorf("SendMessage error: %v", err)
				return
			}

			conn, _ := api.WsCommon.GetConnection()
			conn.PendingMessages.Range(func(key, value interface{}) bool {
				respChanTyped := value.(chan []byte)
				respChanTyped <- responseMsgBytes
				return false
			})

			select {
			case <-respChan:
			case <-errChan:
			case <-time.After(1 * time.Second):
			}
		}(i)
	}

	wg.Wait()
}

// =============================================================================
// WebsocketStreams Tests
// =============================================================================

func TestNewWebsocketStreams(t *testing.T) {
	cfg := &common.ConfigurationWebsocketStreams{
		BasePath:       "wss://test.example.com",
		ReconnectDelay: 1 * time.Second,
		PoolSize:       2,
	}

	ws, err := common.NewWebsocketStreams(cfg)

	require.NoError(t, err)
	assert.NotNil(t, ws)
	assert.Equal(t, cfg, ws.Cfg)
	assert.NotNil(t, ws.GlobalStreamConnectionMap)
	assert.NotNil(t, ws.WsCommon)
	assert.Empty(t, ws.GlobalStreamConnectionMap)
}

func TestWebsocketStreams_Connect(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	err := ws.Connect("test-user-agent")

	assert.NoError(t, err)
}

func TestWebsocketStreams_Subscribe_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade", "ethusdt@trade"}
	ids := []any{"id-1", "id-2"}

	err := ws.Subscribe(streams, ids, false)

	require.NoError(t, err)

	assert.Contains(t, ws.GlobalStreamConnectionMap, "btcusdt@trade")
	assert.Contains(t, ws.GlobalStreamConnectionMap, "ethusdt@trade")

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 2, len(writtenMsgs))

	var msg1 map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &msg1)
	require.NoError(t, err)
	assert.Equal(t, "SUBSCRIBE", msg1["method"])
	assert.Equal(t, "id-1", msg1["id"])
	params1 := msg1["params"].([]interface{})
	assert.Equal(t, "btcusdt@trade", params1[0])

	var msg2 map[string]interface{}
	err = json.Unmarshal(writtenMsgs[1].data, &msg2)
	require.NoError(t, err)
	assert.Equal(t, "SUBSCRIBE", msg2["method"])
	assert.Equal(t, "id-2", msg2["id"])
}

func TestWebsocketStreams_Subscribe_NoStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{}
	ids := []any{}

	err := ws.Subscribe(streams, ids, false)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no streams to subscribe")
}

func TestWebsocketStreams_Subscribe_AutoGenerateID(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	ids := []any{}

	err := ws.Subscribe(streams, ids, false)

	require.NoError(t, err)

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 1, len(writtenMsgs))

	var msg map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &msg)
	require.NoError(t, err)
	assert.NotEmpty(t, msg["id"])
}

func TestWebsocketStreams_Subscribe_AutoGenerateNumberID(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	ids := []any{}

	err := ws.Subscribe(streams, ids, true)

	require.NoError(t, err)

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 1, len(writtenMsgs))

	var msg map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &msg)
	require.NoError(t, err)
	assert.NotEmpty(t, msg["id"])
	assert.IsType(t, float64(0), msg["id"])
}

func TestWebsocketStreams_Subscribe_NumberID(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	ids := []any{12345}

	err := ws.Subscribe(streams, ids, true)

	require.NoError(t, err)

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 1, len(writtenMsgs))

	var msg map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &msg)
	require.NoError(t, err)
	assert.NotEmpty(t, msg["id"])
	assert.IsType(t, float64(0), msg["id"])
	assert.Equal(t, float64(12345), msg["id"])
}

func TestWebsocketStreams_Subscribe_AlreadySubscribed(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	ids := []any{"id-1"}

	err := ws.Subscribe(streams, ids, false)
	require.NoError(t, err)

	err = ws.Subscribe(streams, ids, false)
	require.NoError(t, err)

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.Equal(t, 1, len(writtenMsgs))
}

func TestWebsocketStreams_Subscribe_WriteError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	mockConn.SetWriteError(errors.New("write failed"))
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	ids := []any{"id-1"}

	err := ws.Subscribe(streams, ids, false)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write failed")
}

func TestWebsocketStreams_On_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1"}, false)
	require.NoError(t, err)

	callbackCalled := false
	callback := func(data map[string]interface{}) {
		callbackCalled = true
	}

	err = ws.On("btcusdt@trade", callback)

	require.NoError(t, err)

	conn := ws.WsCommon.Connections[0]
	assert.Len(t, conn.StreamCallbackMap["btcusdt@trade"], 1)

	conn.StreamCallbackMap["btcusdt@trade"][0](map[string]interface{}{"test": "data"})
	assert.True(t, callbackCalled)
}

func TestWebsocketStreams_On_MultipleCallbacks(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1"}, false)
	require.NoError(t, err)

	callback1Called := false
	callback1 := func(data map[string]interface{}) {
		callback1Called = true
	}

	callback2Called := false
	callback2 := func(data map[string]interface{}) {
		callback2Called = true
	}

	err = ws.On("btcusdt@trade", callback1)
	require.NoError(t, err)

	err = ws.On("btcusdt@trade", callback2)
	require.NoError(t, err)

	conn := ws.WsCommon.Connections[0]
	assert.Len(t, conn.StreamCallbackMap["btcusdt@trade"], 2)

	conn.StreamCallbackMap["btcusdt@trade"][0](map[string]interface{}{"test": "data"})
	conn.StreamCallbackMap["btcusdt@trade"][1](map[string]interface{}{"test": "data"})

	assert.True(t, callback1Called)
	assert.True(t, callback2Called)
}

func TestWebsocketStreams_On_NotSubscribed(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	callback := func(data map[string]interface{}) {}

	err := ws.On("unsubscribed-stream", callback)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not subscribed")
}

func TestWebsocketStreams_ListSubscriptions_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	responseMsg := map[string]interface{}{
		"id":     "list-id-1",
		"result": []interface{}{"btcusdt@trade", "ethusdt@trade"},
	}
	responseMsgBytes, _ := json.Marshal(responseMsg)

	resultChan := make(chan map[string]interface{}, 1)
	errorChan := make(chan error, 1)

	go func() {
		result, err := ws.ListSubscriptions("list-id-1")
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- result
		}
	}()

	time.Sleep(100 * time.Millisecond)

	conn := ws.WsCommon.Connections[0]
	respChanInterface, ok := conn.PendingMessages.Load("list-id-1")
	require.True(t, ok)

	respChan := respChanInterface.(chan []byte)
	respChan <- responseMsgBytes

	select {
	case result := <-resultChan:
		assert.NotNil(t, result)
		assert.Equal(t, "list-id-1", result["id"])
		assert.NotNil(t, result["result"])
	case err := <-errorChan:
		t.Fatalf("Unexpected error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("timeout waiting for result")
	}

	writtenMsgs := mockConn.GetWrittenMessages()
	assert.GreaterOrEqual(t, len(writtenMsgs), 1)

	var msg map[string]interface{}
	err := json.Unmarshal(writtenMsgs[len(writtenMsgs)-1].data, &msg)
	require.NoError(t, err)
	assert.Equal(t, "LIST_SUBSCRIPTIONS", msg["method"])
}

func TestWebsocketStreams_ListSubscriptions_AutoGenerateID(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	go func() {
		time.Sleep(100 * time.Millisecond)
		conn := ws.WsCommon.Connections[0]

		var foundID string
		conn.PendingMessages.Range(func(key, value interface{}) bool {
			foundID = key.(string)
			return false
		})

		if foundID != "" {
			responseMsg := map[string]interface{}{
				"id":     foundID,
				"result": []interface{}{"btcusdt@trade"},
			}
			responseMsgBytes, _ := json.Marshal(responseMsg)

			respChanInterface, ok := conn.PendingMessages.Load(foundID)
			if ok {
				respChan := respChanInterface.(chan []byte)
				respChan <- responseMsgBytes
			}
		}
	}()

	result, err := ws.ListSubscriptions("")

	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestWebsocketStreams_ListSubscriptions_Timeout(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	result, err := ws.ListSubscriptions("timeout-id")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "timeout")
}

func TestWebsocketStreams_ListSubscriptions_NilMessage(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	go func() {
		time.Sleep(100 * time.Millisecond)
		conn := ws.WsCommon.Connections[0]
		respChanInterface, ok := conn.PendingMessages.Load("nil-id")
		if ok {
			respChan := respChanInterface.(chan []byte)
			respChan <- nil
		}
	}()

	result, err := ws.ListSubscriptions("nil-id")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "nil message")
}

func TestWebsocketStreams_ListSubscriptions_InvalidJSON(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	go func() {
		time.Sleep(100 * time.Millisecond)
		conn := ws.WsCommon.Connections[0]
		respChanInterface, ok := conn.PendingMessages.Load("invalid-json-id")
		if ok {
			respChan := respChanInterface.(chan []byte)
			respChan <- []byte("{invalid json}")
		}
	}()

	result, err := ws.ListSubscriptions("invalid-json-id")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestWebsocketStreams_Unsubscribe_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade", "ethusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1", "id-2"}, false)
	require.NoError(t, err)

	_ = mockConn.GetWrittenMessages()

	err = ws.Unsubscribe([]string{"btcusdt@trade"})

	require.NoError(t, err)

	assert.NotContains(t, ws.GlobalStreamConnectionMap, "btcusdt@trade")
	assert.Contains(t, ws.GlobalStreamConnectionMap, "ethusdt@trade")

	writtenMsgs := mockConn.GetWrittenMessages()
	require.Greater(t, len(writtenMsgs), 0, "Expected at least one unsubscribe message")

	lastMsg := writtenMsgs[len(writtenMsgs)-1]
	var msg map[string]interface{}
	err = json.Unmarshal(lastMsg.data, &msg)
	require.NoError(t, err)
	assert.Equal(t, "UNSUBSCRIBE", msg["method"])
	params := msg["params"].([]interface{})
	assert.Equal(t, "btcusdt@trade", params[0])
}

func TestWebsocketStreams_Unsubscribe_NoStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	err := ws.Unsubscribe([]string{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no streams to unsubscribe")
}

func TestWebsocketStreams_Unsubscribe_NotSubscribed(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	err := ws.Unsubscribe([]string{"non-existent-stream"})

	assert.NoError(t, err)
}

func TestWebsocketStreams_Unsubscribe_WriteError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1"}, false)
	require.NoError(t, err)

	mockConn.SetWriteError(errors.New("write failed"))

	err = ws.Unsubscribe([]string{"btcusdt@trade"})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "write failed")
}

func TestWebsocketStreams_Unsubscribe_RemovesCallback(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1"}, false)
	require.NoError(t, err)

	callback := func(data map[string]interface{}) {}
	err = ws.On("btcusdt@trade", callback)
	require.NoError(t, err)

	conn := ws.WsCommon.Connections[0]
	assert.Len(t, conn.StreamCallbackMap["btcusdt@trade"], 1)

	err = ws.Unsubscribe([]string{"btcusdt@trade"})
	require.NoError(t, err)

	assert.NotContains(t, conn.StreamCallbackMap, "btcusdt@trade")
}

func TestWebsocketStreams_IsSubscribed_True(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1"}, false)
	require.NoError(t, err)

	result := ws.IsSubscribed("btcusdt@trade")

	assert.True(t, result)
}

func TestWebsocketStreams_IsSubscribed_False(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	result := ws.IsSubscribed("non-existent-stream")

	assert.False(t, result)
}

func TestWebsocketStreams_CloseWebSocketStreamConnection_Success(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	err := ws.CloseWebSocketStreamConnection()

	assert.NoError(t, err)

	conn := ws.WsCommon.Connections[0]
	assert.Equal(t, common.CLOSING, conn.Connected)

	mockConn.mu.Lock()
	closeCalled := mockConn.closeCalled
	mockConn.mu.Unlock()
	assert.True(t, closeCalled)
}

func TestWebsocketStreams_CloseWebSocketStreamConnection_MultipleConnections(t *testing.T) {
	mockConn1 := NewMockWebSocketConn()
	mockConn2 := NewMockWebSocketConn()

	cfg := &common.ConfigurationWebsocketStreams{
		BasePath: "wss://test.example.com",
		PoolSize: 2,
	}

	conn1 := &common.WebSocketConnection{
		Id:        "stream-conn-1",
		Connected: common.OPEN,
		Websocket: mockConn1,
	}

	conn2 := &common.WebSocketConnection{
		Id:        "stream-conn-2",
		Connected: common.OPEN,
		Websocket: mockConn2,
	}

	wsCommon := &common.WebSocketCommon{
		Connections: []*common.WebSocketConnection{conn1, conn2},
	}

	ws := &common.WebsocketStreams{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: make(map[string][]*common.WebSocketConnection),
		WsCommon:                  wsCommon,
	}

	err := ws.CloseWebSocketStreamConnection()

	assert.NoError(t, err)

	mockConn1.mu.Lock()
	close1 := mockConn1.closeCalled
	mockConn1.mu.Unlock()

	mockConn2.mu.Lock()
	close2 := mockConn2.closeCalled
	mockConn2.mu.Unlock()

	assert.True(t, close1)
	assert.True(t, close2)
	assert.Equal(t, common.CLOSING, conn1.Connected)
	assert.Equal(t, common.CLOSING, conn2.Connected)
}

func TestWebsocketStreams_CloseWebSocketStreamConnection_NilWebsocket(t *testing.T) {
	cfg := &common.ConfigurationWebsocketStreams{
		BasePath: "wss://test.example.com",
		PoolSize: 1,
	}

	conn := &common.WebSocketConnection{
		Id:        "stream-conn-1",
		Connected: common.OPEN,
		Websocket: nil,
	}

	wsCommon := &common.WebSocketCommon{
		Connections: []*common.WebSocketConnection{conn},
	}

	ws := &common.WebsocketStreams{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: make(map[string][]*common.WebSocketConnection),
		WsCommon:                  wsCommon,
	}

	err := ws.CloseWebSocketStreamConnection()

	assert.NoError(t, err)
}

func TestWebsocketStreams_Subscribe_ConcurrentSubscriptions(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	var wg sync.WaitGroup
	numGoroutines := 5
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			stream := []string{"stream-" + string(rune('A'+index))}
			err := ws.Subscribe(stream, []any{}, false)
			assert.NoError(t, err)
		}(i)
	}

	wg.Wait()

	assert.Equal(t, numGoroutines, len(ws.GlobalStreamConnectionMap))
}

func TestWebsocketStreams_CompleteWorkflow(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	ws := createTestWebsocketStreams(mockConn)

	streams := []string{"btcusdt@trade", "ethusdt@trade"}
	err := ws.Subscribe(streams, []any{"id-1", "id-2"}, false)
	require.NoError(t, err)

	assert.True(t, ws.IsSubscribed("btcusdt@trade"))
	assert.True(t, ws.IsSubscribed("ethusdt@trade"))

	btcCallbackCalled := false
	btcCallback := func(data map[string]interface{}) {
		btcCallbackCalled = true
	}

	err = ws.On("btcusdt@trade", btcCallback)
	require.NoError(t, err)

	conn := ws.WsCommon.Connections[0]
	assert.Len(t, conn.StreamCallbackMap["btcusdt@trade"], 1)
	assert.Equal(t, btcCallbackCalled, false)

	err = ws.Unsubscribe([]string{"ethusdt@trade"})
	require.NoError(t, err)

	assert.True(t, ws.IsSubscribed("btcusdt@trade"))
	assert.False(t, ws.IsSubscribed("ethusdt@trade"))

	err = ws.CloseWebSocketStreamConnection()
	assert.NoError(t, err)

	assert.Equal(t, common.CLOSING, conn.Connected)
}

// =============================================================================
// CreateStreamHandler Tests
// =============================================================================

func TestCreateStreamHandler_WithWebsocketStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)

	require.NoError(t, err)
	assert.NotNil(t, handler)
	assert.Equal(t, stream, handler.Stream)
	assert.Equal(t, wrapper, handler.Wrapper)

	assert.Contains(t, wrapper.WebsocketStreams.GlobalStreamConnectionMap, stream)
}

func TestCreateStreamHandler_WithWebsocketAPI(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithAPI(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{}, false)

	require.NoError(t, err)
	assert.NotNil(t, handler)
	assert.Equal(t, stream, handler.Stream)
	assert.Equal(t, wrapper, handler.Wrapper)

	assert.Contains(t, wrapper.WebsocketAPI.GlobalStreamConnectionMap, stream)
}

func TestCreateStreamHandler_InvalidWrapper(t *testing.T) {
	wrapper := &common.StreamHandlerWrapper{}

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{}, false)

	assert.Error(t, err)
	assert.Nil(t, handler)

	wsErr, ok := err.(*common.WebSocketError)
	require.True(t, ok, "Expected WebSocketError")
	assert.Equal(t, "create_stream_handler", wsErr.Op)
	assert.Contains(t, wsErr.Message, "invalid StreamHandlerWrapper")
}

func TestCreateStreamHandler_WithCustomIDs(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	customIDs := []any{"custom-id-123"}

	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, customIDs, false)

	require.NoError(t, err)
	assert.NotNil(t, handler)

	writtenMsgs := mockConn.GetWrittenMessages()
	require.Greater(t, len(writtenMsgs), 0)

	var msg map[string]interface{}
	err = json.Unmarshal(writtenMsgs[0].data, &msg)
	require.NoError(t, err)
	assert.Equal(t, "custom-id-123", msg["id"])
}

func TestStreamHandler_On_SingleObject_WithStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var receivedData TestTradeData
	var wg sync.WaitGroup
	wg.Add(1)

	handler.On("message", func(data TestTradeData) {
		receivedData = data
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	conn := wrapper.WebsocketStreams.WsCommon.Connections[0]
	callbacks := conn.StreamCallbackMap[stream]
	assert.Len(t, callbacks, 1)

	testData := map[string]interface{}{
		"data": map[string]interface{}{
			"symbol": "BTCUSDT",
			"price":  50000.0,
			"qty":    1.5,
		},
	}

	triggerStreamCallback(conn, stream, testData)

	wg.Wait()

	assert.Equal(t, "BTCUSDT", receivedData.Symbol)
	assert.Equal(t, 50000.0, receivedData.Price)
	assert.Equal(t, 1.5, receivedData.Qty)
}

func TestStreamHandler_On_SingleObject_WithAPI(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithAPI(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{}, false)
	require.NoError(t, err)

	var receivedData TestTradeData
	var wg sync.WaitGroup
	wg.Add(1)

	handler.On("message", func(data TestTradeData) {
		receivedData = data
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	conn := wrapper.WebsocketAPI.WsCommon.Connections[0]
	callbacks := conn.StreamCallbackMap[stream]
	assert.Len(t, callbacks, 1)

	testData := map[string]interface{}{
		"symbol": "BTCUSDT",
		"price":  50000.0,
		"qty":    1.5,
	}

	triggerStreamCallback(conn, stream, testData)

	wg.Wait()

	assert.Equal(t, "BTCUSDT", receivedData.Symbol)
	assert.Equal(t, 50000.0, receivedData.Price)
	assert.Equal(t, 1.5, receivedData.Qty)
}

func TestStreamHandler_On_ArrayOfObjects(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "!ticker@arr"
	handler, err := common.CreateStreamHandler[TestTickerData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var receivedData []TestTickerData
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(3)

	handler.On("message", func(data TestTickerData) {
		mu.Lock()
		receivedData = append(receivedData, data)
		mu.Unlock()
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	testData := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"s": "BTCUSDT",
				"c": "50000.00",
				"v": "1000.5",
			},
			map[string]interface{}{
				"s": "ETHUSDT",
				"c": "3000.00",
				"v": "500.25",
			},
			map[string]interface{}{
				"s": "BNBUSDT",
				"c": "400.00",
				"v": "200.10",
			},
		},
	}

	conn := wrapper.WebsocketStreams.WsCommon.Connections[0]
	triggerStreamCallback(conn, stream, testData)

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	assert.Len(t, receivedData, 3)
	assert.Equal(t, "BTCUSDT", receivedData[0].Symbol)
	assert.Equal(t, "ETHUSDT", receivedData[1].Symbol)
	assert.Equal(t, "BNBUSDT", receivedData[2].Symbol)
}

func TestStreamHandler_On_NonMessageEvent(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	callbackCalled := false
	handler.On("connect", func(data TestTradeData) {
		callbackCalled = true
	})

	time.Sleep(50 * time.Millisecond)

	conn := wrapper.WebsocketStreams.WsCommon.Connections[0]
	callbacks := conn.StreamCallbackMap[stream]

	assert.Len(t, callbacks, 0, "Callback should not be registered for non-message events")
	assert.False(t, callbackCalled)
}

func TestStreamHandler_On_StreamNotSubscribed(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler := &common.StreamHandler[TestTradeData]{
		Stream:  stream,
		Wrapper: wrapper,
	}

	callbackCalled := false
	handler.On("message", func(data TestTradeData) {
		callbackCalled = true
	})

	time.Sleep(50 * time.Millisecond)

	assert.False(t, callbackCalled)
}

func TestStreamHandler_On_MultipleCallbacks(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var callback1Called, callback2Called bool
	var wg sync.WaitGroup
	wg.Add(2)

	handler.On("message", func(data TestTradeData) {
		callback1Called = true
		wg.Done()
	})

	handler.On("message", func(data TestTradeData) {
		callback2Called = true
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	conn := wrapper.WebsocketStreams.WsCommon.Connections[0]
	callbacks := conn.StreamCallbackMap[stream]
	assert.Len(t, callbacks, 2)

	testData := map[string]interface{}{
		"symbol": "BTCUSDT",
		"price":  50000.0,
		"qty":    1.5,
	}

	triggerStreamCallback(conn, stream, testData)

	wg.Wait()

	assert.True(t, callback1Called)
	assert.True(t, callback2Called)
}

func TestStreamHandler_On_InvalidJSON(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	callbackCalled := false
	handler.On("message", func(data TestTradeData) {
		callbackCalled = true
	})

	time.Sleep(50 * time.Millisecond)

	testData := map[string]interface{}{
		"data": "invalid-data-type",
	}

	conn := wrapper.WebsocketStreams.WsCommon.Connections[0]
	triggerStreamCallback(conn, stream, testData)

	time.Sleep(100 * time.Millisecond)
	assert.False(t, callbackCalled)
}

func TestStreamHandler_OnError(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var receivedError error
	var wg sync.WaitGroup
	wg.Add(1)

	handler.OnError(func(err error) {
		receivedError = err
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	testError := errors.New("connection error")
	conn := wrapper.WebsocketStreams.GlobalStreamConnectionMap[stream][0]

	conn.ErrorChan <- testError

	wg.Wait()

	assert.Equal(t, testError, receivedError)
}

func TestStreamHandler_OnError_MultipleErrors(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var receivedErrors []error
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(3)

	handler.OnError(func(err error) {
		mu.Lock()
		receivedErrors = append(receivedErrors, err)
		mu.Unlock()
		wg.Done()
	})

	time.Sleep(50 * time.Millisecond)

	conn := wrapper.WebsocketStreams.GlobalStreamConnectionMap[stream][0]

	testErrors := []error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}

	for _, e := range testErrors {
		conn.ErrorChan <- e
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	assert.Len(t, receivedErrors, 3)
}

func TestStreamHandler_OnError_StreamNotSubscribed(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler := &common.StreamHandler[TestTradeData]{
		Stream:  stream,
		Wrapper: wrapper,
	}

	errorCallbackCalled := false
	handler.OnError(func(err error) {
		errorCallbackCalled = true
	})

	time.Sleep(100 * time.Millisecond)
	assert.False(t, errorCallbackCalled)
}

func TestStreamHandler_Unsubscribe_WithStreams(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	assert.Contains(t, wrapper.WebsocketStreams.GlobalStreamConnectionMap, stream)

	messagesBefore := len(mockConn.GetWrittenMessages())

	handler.Unsubscribe()

	time.Sleep(50 * time.Millisecond)

	assert.NotContains(t, wrapper.WebsocketStreams.GlobalStreamConnectionMap, stream)

	allMsgs := mockConn.GetWrittenMessages()
	require.Greater(t, len(allMsgs), messagesBefore, "Expected new message after unsubscribe")

	var unsubMsg map[string]interface{}
	found := false
	for i := len(allMsgs) - 1; i >= messagesBefore; i-- {
		var msg map[string]interface{}
		if err := json.Unmarshal(allMsgs[i].data, &msg); err == nil {
			if method, ok := msg["method"].(string); ok && method == "UNSUBSCRIBE" {
				unsubMsg = msg
				found = true
				break
			}
		}
	}

	require.True(t, found, "Could not find UNSUBSCRIBE message")
	assert.Equal(t, "UNSUBSCRIBE", unsubMsg["method"])

	params, ok := unsubMsg["params"].([]interface{})
	require.True(t, ok)
	require.Greater(t, len(params), 0)
	assert.Equal(t, stream, params[0])
}

func TestStreamHandler_Unsubscribe_WithAPI(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithAPI(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{}, false)
	require.NoError(t, err)

	handler.Unsubscribe()

	assert.NotNil(t, handler)
}

func TestStreamHandler_CompleteWorkflow(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var receivedTrades []TestTradeData
	var mu sync.Mutex
	var messageWg sync.WaitGroup
	messageWg.Add(2)

	handler.On("message", func(data TestTradeData) {
		mu.Lock()
		receivedTrades = append(receivedTrades, data)
		mu.Unlock()
		messageWg.Done()
	})

	var receivedErrors []error
	var errorWg sync.WaitGroup
	errorWg.Add(1)

	handler.OnError(func(err error) {
		mu.Lock()
		receivedErrors = append(receivedErrors, err)
		mu.Unlock()
		errorWg.Done()
	})

	time.Sleep(100 * time.Millisecond)

	conn := wrapper.WebsocketStreams.GlobalStreamConnectionMap[stream][0]

	testTrades := []map[string]interface{}{
		{
			"symbol": "BTCUSDT",
			"price":  50000.0,
			"qty":    1.5,
		},
		{
			"symbol": "BTCUSDT",
			"price":  50100.0,
			"qty":    2.0,
		},
	}

	for _, trade := range testTrades {
		triggerStreamCallback(conn, stream, trade)
		time.Sleep(10 * time.Millisecond)
	}

	messageWg.Wait()

	conn.ErrorChan <- errors.New("test error")

	errorWg.Wait()

	mu.Lock()
	assert.Len(t, receivedTrades, 2)
	assert.Equal(t, 50000.0, receivedTrades[0].Price)
	assert.Equal(t, 50100.0, receivedTrades[1].Price)
	assert.Len(t, receivedErrors, 1)
	assert.Equal(t, "test error", receivedErrors[0].Error())
	mu.Unlock()

	handler.Unsubscribe()
	assert.NotContains(t, wrapper.WebsocketStreams.GlobalStreamConnectionMap, stream)
}

func TestStreamHandler_ConcurrentCallbacks(t *testing.T) {
	mockConn := NewMockWebSocketConn()
	wrapper := createStreamHandlerWrapperWithStreams(mockConn)

	stream := "btcusdt@trade"
	handler, err := common.CreateStreamHandler[TestTradeData](wrapper, stream, []any{"id-1"}, false)
	require.NoError(t, err)

	var callbackCount int
	var mu sync.Mutex
	var wg sync.WaitGroup
	numCallbacks := 100
	wg.Add(numCallbacks)

	handler.On("message", func(data TestTradeData) {
		mu.Lock()
		callbackCount++
		mu.Unlock()
		wg.Done()
	})

	time.Sleep(100 * time.Millisecond)
	conn := wrapper.WebsocketStreams.GlobalStreamConnectionMap[stream][0]
	testData := map[string]interface{}{
		"symbol": "BTCUSDT",
		"price":  50000.0,
		"qty":    1.5,
	}

	for i := 0; i < numCallbacks; i++ {
		go func() {
			triggerStreamCallback(conn, stream, testData)
		}()
	}

	wg.Wait()

	mu.Lock()
	defer mu.Unlock()
	assert.Equal(t, numCallbacks, callbackCount)
}
