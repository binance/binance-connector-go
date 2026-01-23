package common

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/net/proxy"
)

// WebSocketError represents an error that occurred during WebSocket operations.
func (e *WebSocketError) Error() string {
	var errStr string
	if e.Err != nil {
		errStr = e.Err.Error()
	} else {
		errStr = "no error details"
	}
	var errNbr int
	if e.Code != 0 {
		errNbr = e.Code
	} else {
		errNbr = 0
	}

	return fmt.Sprintf("WebSocket %s error[%v] (conn %s): %v - %s",
		e.Op, errNbr, e.ConnID, errStr, e.Message)
}

// Listen starts listening for incoming messages on the WebSocket connection.
func (c *WebSocketConnection) Listen() {
	go func() {
		for {
			select {
			case <-c.Done:
				log.Println("WebSocket listener shutting down")
				return
			default:
				if err := c.ProcessMessage(); err != nil {
					c.HandleReadError(err)
					return
				}
			}
		}
	}()
}

// ProcessMessage reads and processes a single message from the WebSocket connection.
//
// @return An error if processing fails, otherwise nil.
func (c *WebSocketConnection) ProcessMessage() error {
	_, msg, err := c.Websocket.ReadMessage()
	if err != nil {
		return err
	}

	data, err := c.parseMessage(msg)
	if err != nil {
		return &WebSocketError{
			ConnID:  c.Id,
			Err:     err,
			Message: "failed to parse websocket message",
			Op:      "parse_message",
		}
	}

	if c.isErrorResponse(data) {
		log.Printf("WebSocket received error message: %s\n", string(msg))

		var errorCode int
		var errorMsg string

		if errData, ok := data["error"].(map[string]interface{}); ok {
			if code, ok := errData["code"].(float64); ok {
				errorCode = int(code)
			} else if code, ok := errData["code"].(int); ok {
				errorCode = code
			}

			if msg, ok := errData["msg"].(string); ok {
				errorMsg = msg
			}
		}

		wsErr := fmt.Errorf("[%d] %s", errorCode, errorMsg)

		statusCode := 0
		if status, ok := data["status"].(float64); ok {
			statusCode = int(status)
		} else if status, ok := data["status"].(int); ok {
			statusCode = status
		}

		return &WebSocketError{
			Code:    statusCode,
			ConnID:  c.Id,
			Err:     wsErr,
			Message: "received error response from server",
			Op:      "error_response",
		}
	}

	if c.handleResponseMessage(data, msg) {
		return nil
	}

	if c.handleSubscriptionMessage(data) {
		return nil
	}

	c.handleStreamMessage(data)
	return nil
}

// parseMessage parses the incoming WebSocket message.
//
// @param msg The raw message bytes received from the WebSocket.
// @return A map representing the parsed message data or an error if parsing fails.
func (c *WebSocketConnection) parseMessage(msg []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(msg, &data)
	return data, err
}

// isErrorResponse checks if the message data indicates an error response.
//
// @param data The parsed message data as a map.
// @return True if the message is an error response, otherwise false.
func (c *WebSocketConnection) isErrorResponse(data map[string]interface{}) bool {
	if status, ok := data["status"].(float64); ok {
		return status >= 300
	}
	return false
}

// handleResponseMessage processes incoming response messages and routes them to the appropriate pending message channels.
//
// @param data The parsed message data as a map.
// @param msg The original message bytes.
// @return True if the message was handled as a response message, otherwise false.
func (c *WebSocketConnection) handleResponseMessage(data map[string]interface{}, msg []byte) bool {
	id, ok := data["id"].(string)
	if !ok || id == "" {
		return false
	}

	if chAny, loaded := c.PendingMessages.Load(id); loaded {
		ch := chAny.(chan []byte)
		ch <- msg
		close(ch)
		c.PendingMessages.Delete(id)
	}
	return true
}

// handleSubscriptionMessage processes incoming subscription messages and triggers the appropriate callbacks.
//
// @param data The parsed message data as a map.
// @return True if the message was handled as a subscription message, otherwise false.
func (c *WebSocketConnection) handleSubscriptionMessage(data map[string]interface{}) bool {
	subscriptionId, ok := data["subscriptionId"].(float64)
	if !ok {
		return false
	}

	subscriptionIdStr := strconv.FormatFloat(subscriptionId, 'f', -1, 64)
	callbacks := c.getCallbacks(subscriptionIdStr)

	if len(callbacks) > 0 {
		c.executeCallbacks(callbacks, data)
	} else {
		log.Printf("No callback registered for subscription ID %s", subscriptionIdStr)
	}
	return true
}

// handleStreamMessage processes incoming stream messages and triggers the appropriate callbacks.
//
// @param data The parsed message data as a map.
func (c *WebSocketConnection) handleStreamMessage(data map[string]interface{}) {
	streamName, ok := data["stream"].(string)
	if !ok || streamName == "" {
		return
	}

	callbacks := c.getCallbacks(streamName)
	if len(callbacks) > 0 {
		c.executeCallbacks(callbacks, data)
	} else {
		log.Printf("No callback registered for stream %s", streamName)
	}
}

// getCallbacks retrieves the list of callbacks for the given key.
//
// @param key The key to look up in the StreamCallbackMap.
func (c *WebSocketConnection) getCallbacks(key string) []func(map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.StreamCallbackMap[key]
}

// executeCallbacks executes the provided callbacks with the given data.
//
// @param callbacks The list of callback functions to execute.
// @param data The data to pass to each callback function.
func (c *WebSocketConnection) executeCallbacks(callbacks []func(map[string]interface{}), data map[string]interface{}) {
	for _, cb := range callbacks {
		go func(callback func(map[string]interface{})) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in callback: %v", r)
				}
			}()
			callback(data)
		}(cb)
	}
}

// HandleReadError processes errors encountered during WebSocket read operations.
//
// @param err The error encountered during the read operation.
func (c *WebSocketConnection) HandleReadError(err error) {
	if c.isNormalCloseError(err) {
		log.Printf("WebSocket closed normally for %s", c.Id)
		c.Connected = CLOSED
	} else if strings.Contains(err.Error(), "use of closed network connection") {
		log.Printf("WebSocket %s closed intentionally", c.Id)
		c.Connected = CLOSED
	} else {
		log.Printf("WebSocket read error: %v\n", err)

		select {
		case c.ErrorChan <- err:
		default:
		}

		c.notifyPendingMessagesWithError(err)
	}
}

// isNormalCloseError checks if the error is a normal WebSocket closure error.
//
// @param err The error to check.
// @return True if the error indicates a normal closure, otherwise false.
func (c *WebSocketConnection) isNormalCloseError(err error) bool {
	return websocket.IsCloseError(err,
		websocket.CloseNormalClosure,
		websocket.CloseGoingAway) ||
		errors.Is(err, net.ErrClosed)
}

// notifyPendingMessagesWithError notifies all pending messages with the provided error.
//
// @param err The error to notify pending messages with.
func (c *WebSocketConnection) notifyPendingMessagesWithError(err error) {
	errorMsg := `{"error":"` + strings.ReplaceAll(err.Error(), `"`, `\"`) + `"}`
	c.PendingMessages.Range(func(key, chAny any) bool {
		ch := chAny.(chan []byte)
		ch <- []byte(errorMsg)
		close(ch)
		c.PendingMessages.Delete(key)
		return true
	})
}

// IsHealthy checks if the WebSocket connection is healthy.
//
// @return True if the connection is healthy, otherwise false.
func (c *WebSocketConnection) IsHealthy() bool {
	select {
	case <-c.Done:
		return false
	default:
		return c.Connected == OPEN
	}
}

// IsAPI checks if the configuration is for API.
//
// @return True if the configuration is for API, otherwise false.
func (c ConfigurationWrapper) IsAPI() bool { return c.APIConfig != nil }

// IsStreams checks if the configuration is for Streams.
//
// @return True if the configuration is for Streams, otherwise false.
func (c ConfigurationWrapper) IsStreams() bool { return c.StreamsConfig != nil }

// NewWebSocketCommon creates a new instance of WebSocketCommon based on the provided configuration.
//
// @param cfg The configuration wrapper containing either API or Streams configuration.
// @return A pointer to the WebSocketCommon instance or an error if initialization fails.
func NewWebSocketCommon(cfg *ConfigurationWrapper) (*WebSocketCommon, error) {
	var poolSize int
	var mode WebsocketMode

	if cfg == nil {
		return nil, &WebSocketError{
			Op:      "initialize_websocket_common",
			Err:     errors.New("nil configuration"),
			Message: "Configuration cannot be nil",
		}
	}

	if cfg.IsAPI() {
		poolSize = cfg.APIConfig.PoolSize
		mode = cfg.APIConfig.Mode
	} else if cfg.IsStreams() {
		poolSize = cfg.StreamsConfig.PoolSize
		mode = cfg.StreamsConfig.Mode
	} else {
		return nil, &WebSocketError{
			Op:      "initialize_websocket_common",
			Err:     errors.New("invalid configuration"),
			Message: "Configuration must be for API or Streams",
		}
	}

	wsc := &WebSocketCommon{
		PoolSize:        poolSize,
		Mode:            mode,
		RoundRobinIndex: 0,
	}

	wsc.initializePool()
	return wsc, nil
}

// initializePool initializes the WebSocket connection pool based on the configured pool size.
func (w *WebSocketCommon) initializePool() {
	for i := 0; i < w.PoolSize; i++ {
		log.Printf("Initializing WebSocket connection %d/%d\n", i+1, w.PoolSize)
		id, _ := RandomString()

		w.Connections = append(w.Connections, &WebSocketConnection{
			Connected:           CONNECTING,
			Id:                  id,
			PendingMessages:     sync.Map{},
			StreamCallbackMap:   make(map[string][]func(map[string]interface{})),
			StreamConnectionMap: []string{},
			ErrorChan:           make(chan error, 1),
			Done:                make(chan struct{}),
		})
	}
}

// Connect establishes the WebSocket connection using the provided configuration and user agent.
//
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to be used for the connection.
// @return An error if the connection fails, otherwise nil.
func (w *WebSocketCommon) Connect(config WebSocketConfig, userAgent string) error {
	if err := w.setupProxyDialer(config); err != nil {
		return fmt.Errorf("proxy setup failed: %v", err)
	}

	BasePath := w.prepareBasePath(config)
	headers := w.prepareHeaders(config, userAgent)
	dialer := w.CreateWebSocketDialer(config)

	if w.Mode == SINGLE {
		return w.connectSingleMode(BasePath, headers, dialer, config, userAgent)
	}
	return w.connectPoolMode(BasePath, headers, dialer, config, userAgent)
}

// prepareBasePath constructs the base path for the WebSocket connection.
//
// @param config The WebSocketConfig containing configuration details.
// @return The constructed base path string.
func (w *WebSocketCommon) prepareBasePath(config WebSocketConfig) string {
	BasePath := config.GetBasePath()
	if timeUnit := config.GetTimeUnit(); timeUnit != "" {
		BasePath = BasePath + "?timeUnit=" + string(timeUnit)
	}
	return BasePath
}

// prepareHeaders prepares the HTTP headers for the WebSocket connection.
//
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to use for the connection.
// @return The prepared http.Header instance.
func (w *WebSocketCommon) prepareHeaders(config WebSocketConfig, userAgent string) http.Header {
	headers := http.Header{
		"User-Agent": []string{userAgent},
	}
	return headers
}

// CreateWebSocketDialer creates a WebSocket dialer with the specified configuration.
//
// @param config The WebSocketConfig containing configuration details.
// @return A configured websocket.Dialer instance.
func (w *WebSocketCommon) CreateWebSocketDialer(config WebSocketConfig) websocket.Dialer {
	dialer := websocket.Dialer{
		ReadBufferSize:    16 * 1024 * 1024,
		WriteBufferSize:   16 * 1024 * 1024,
		EnableCompression: config.GetCompression(),
		TLSClientConfig:   config.GetTLSConfig(),
		HandshakeTimeout:  45 * time.Second,
	}

	agent := config.GetAgent()

	if agent != nil {
		transport := WsBuildTransport(agent, config.GetTLSConfig())
		if transport != nil {
			if t, ok := transport.(*http.Transport); ok {
				if t.DialContext != nil {
					dialer.NetDialContext = t.DialContext
				}

				if t.Proxy != nil {
					dialer.Proxy = t.Proxy
				}

				if t.TLSClientConfig != nil {
					dialer.TLSClientConfig = t.TLSClientConfig
				}
				if t.TLSClientConfig != nil {
					dialer.TLSClientConfig = t.TLSClientConfig
				} else if config.GetTLSConfig() != nil {
					dialer.TLSClientConfig = config.GetTLSConfig()
				}
			}
		}
	}

	if dialer.NetDialContext == nil {
		dialer.NetDialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := w.ProxyDialer.Dial(network, addr)
			if err != nil {
				return nil, fmt.Errorf("proxy dial failed: %v", err)
			}

			if tlsConfig := config.GetTLSConfig(); tlsConfig != nil && (strings.HasPrefix(addr, "443") || strings.Contains(addr, ":443")) {
				return tls.Client(conn, tlsConfig), nil
			}

			return conn, nil
		}
	}

	if dialer.Proxy == nil {
		dialer.Proxy = http.ProxyFromEnvironment
	}

	return dialer
}

// connectSingleMode establishes a single WebSocket connection.
//
// @param BasePath The base URL for the WebSocket connection.
// @param headers The HTTP headers to include in the connection request.
// @param dialer The WebSocket dialer to use for the connection.
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to use for the connection.
// @return An error if the connection fails, otherwise nil.
func (w *WebSocketCommon) connectSingleMode(BasePath string, headers http.Header, dialer websocket.Dialer, config WebSocketConfig, userAgent string) error {
	conn, _, err := dialer.Dial(BasePath, headers)
	if err != nil {
		return err
	}

	w.configureConnection(conn, w.Connections[0])
	go w.Connections[0].Listen()
	go w.KeepAlive(w.Connections[0], config, userAgent)
	return nil
}

// connectPoolMode establishes WebSocket connections in pool mode.
//
// @param BasePath The base URL for the WebSocket connection.
// @param headers The HTTP headers to include in the connection request.
// @param dialer The WebSocket dialer to use for the connection.
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to use for the connection.
// @return An error if the connection fails, otherwise nil.
func (w *WebSocketCommon) connectPoolMode(BasePath string, headers http.Header, dialer websocket.Dialer, config WebSocketConfig, userAgent string) error {
	var wg sync.WaitGroup
	successChan := make(chan bool, len(w.Connections))
	errChan := make(chan error, len(w.Connections))

	for num, connection := range w.Connections {
		if connection.Websocket == nil {
			wg.Add(1)
			go func(num int, conn *WebSocketConnection) {
				defer wg.Done()

				wsConn, _, err := dialer.Dial(BasePath, headers)
				if err != nil {
					log.Printf("WebSocket connection error: %v", err)
					errChan <- err
					return
				}

				w.configureConnection(wsConn, conn)
				go conn.Listen()
				go w.KeepAlive(conn, config, userAgent)
				successChan <- true
			}(num, connection)
		}
	}

	go func() {
		wg.Wait()
		close(successChan)
		close(errChan)
	}()

	success := false
	for range w.Connections {
		select {
		case <-successChan:
			success = true
		case err := <-errChan:
			log.Printf("Connection error: %v", err)
		}
	}

	if !success {
		return errors.New("failed to establish any WebSocket connections in pool mode")
	}

	return nil
}

// configureConnection sets up the WebSocket connection with necessary configurations.
//
// @param conn The WebSocket connection to configure.
// @param connection The WebSocketConnection wrapper to associate with the connection.
func (w *WebSocketCommon) configureConnection(conn *websocket.Conn, connection *WebSocketConnection) {
	conn.SetReadLimit(64 * 1024 * 1024)
	conn.SetPingHandler(func(appData string) error {
		return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(time.Second))
	})

	connection.Websocket = conn
	connection.Connected = OPEN
}

// KeepAlive periodically checks the health of the WebSocket connection and reconnects if necessary.
//
// @param connection The WebSocketConnection to monitor.
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to use for the connection.
func (w *WebSocketCommon) KeepAlive(connection *WebSocketConnection, config WebSocketConfig, userAgent string) {
	ticker := time.NewTicker(23 * time.Hour)
	defer ticker.Stop()

	reconnectDelay := config.GetReconnectDelay()

	for range ticker.C {
		if !connection.IsHealthy() {
			log.Printf("WebSocket connection ID: %s is not ready", connection.Id)
			return
		}

		if connection.Connected == CLOSING || connection.Connected == CLOSED {
			log.Printf("WebSocket connection ID: %s is closing or closed", connection.Id)
			return
		}

		connection.Connected = CLOSING
		time.Sleep(reconnectDelay)

		w.logoutIfNeeded(connection)

		close(connection.Done)
		connection.Done = make(chan struct{})
		if err := connection.Websocket.Close(); err != nil {
			log.Printf("Failed to close WebSocket: %v", err)
		}
		err := w.reconnect(connection, config, userAgent)
		if err != nil {
			log.Printf("WebSocket reconnection failed for connection ID: %s, error: %v", connection.Id, err)
			return
		}
	}
}

// reconnect handles the reconnection logic for a WebSocket connection.
//
// @param conn The WebSocketConnection to reconnect.
// @param config The WebSocketConfig containing configuration details.
// @param userAgent The user agent string to use for the connection.
// @return An error if the reconnection fails, otherwise nil.
func (w *WebSocketCommon) reconnect(conn *WebSocketConnection, config WebSocketConfig, userAgent string) error {
	BasePath := w.prepareBasePath(config)
	headers := w.prepareHeaders(config, userAgent)
	dialer := w.CreateWebSocketDialer(config)

	newConn, _, err := dialer.Dial(BasePath, headers)
	if err != nil {
		return err
	}

	w.configureConnection(newConn, conn)
	go conn.Listen()

	if conn.SessionLogonRequest != nil && !conn.SessionLogon {
		w.restoreSessionIfNeeded(conn)
		time.Sleep(1 * time.Second)
		w.resubscribeUserDataStreams(conn)
	} else if len(conn.StreamConnectionMap) > 0 {
		w.resubscribeRegularStreams(conn)
	}

	return nil
}

// logoutIfNeeded logs out of the session if a session logon exists.
//
// @param connection The WebSocketConnection to log out from.
func (w *WebSocketCommon) logoutIfNeeded(connection *WebSocketConnection) {
	if !connection.SessionLogon {
		return
	}

	disconnect := map[string]any{
		"method": "session.logout",
		"params": map[string]any{},
		"id":     GenerateUUID(),
	}
	message, err := json.Marshal(disconnect)
	if err != nil {
		log.Printf("Error during session logout: %v", err)
	} else {
		connection.mu.Lock()
		err = connection.Websocket.WriteMessage(websocket.TextMessage, message)
		connection.mu.Unlock()
		if err != nil {
			log.Printf("Error sending session logout message: %v", err)
		}
	}
	connection.SessionLogon = false
}

// restoreSessionIfNeeded restores the session after reconnection if a session logon request exists.
//
// @param connection The WebSocketConnection to restore the session on.
func (w *WebSocketCommon) restoreSessionIfNeeded(connection *WebSocketConnection) {
	connection.SessionLogonRequest.Params["timestamp"] = GetTimestamp()
	delete(connection.SessionLogonRequest.Params, "signature")

	signature, err := SignMessage(connection.SessionLogonRequest.Priv, []byte(Urlencode(connection.SessionLogonRequest.Params)))
	if err != nil {
		panic(err)
	}
	connection.SessionLogonRequest.Params["signature"] = signature

	data := map[string]interface{}{
		"method": connection.SessionLogonRequest.Method,
		"params": connection.SessionLogonRequest.Params,
		"id":     connection.SessionLogonRequest.ID,
	}
	message, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error during session logon on reconnection: %v", err)
		return
	}
	connection.mu.Lock()
	err = connection.Websocket.WriteMessage(websocket.TextMessage, message)
	connection.mu.Unlock()
	if err != nil {
		log.Printf("Error sending session logon message on reconnection: %v", err)
		return
	}
	log.Printf("Session logon sent on reconnection for connection ID: %s", connection.Id)
	connection.SessionLogon = true
}

// resubscribeUserDataStreams resubscribes to all user data streams after reconnection.
//
// @param connection The WebSocketConnection to resubscribe streams on.
func (w *WebSocketCommon) resubscribeUserDataStreams(connection *WebSocketConnection) {
	for _, streamID := range connection.StreamConnectionMap {
		log.Printf("Resubscribing to stream ID: %s on reconnection", streamID)
		subscribePayload := map[string]interface{}{
			"method": "userDataStream.subscribe",
			"params": make(map[string]any),
			"id":     GenerateUUID(),
		}

		log.Printf("Resubscribing with payload: %v\n", subscribePayload)
		message, err := json.Marshal(subscribePayload)
		if err != nil {
			log.Printf("Error during resubscription to stream %s: %v", streamID, err)
			continue
		}
		connection.mu.Lock()
		err = connection.Websocket.WriteMessage(websocket.TextMessage, message)
		connection.mu.Unlock()
		if err != nil {
			log.Printf("Error sending resubscription message for stream %s: %v", streamID, err)
			continue
		}
	}
}

// resubscribeRegularStreams resubscribes to all regular streams after reconnection.
//
// @param connection The WebSocketConnection to resubscribe streams on.
func (w *WebSocketCommon) resubscribeRegularStreams(connection *WebSocketConnection) {
	for _, stream := range connection.StreamConnectionMap {
		subscribePayload := map[string]interface{}{
			"method": "SUBSCRIBE",
			"params": []string{stream},
			"id":     GenerateUUID(),
		}

		message, err := json.Marshal(subscribePayload)
		if err != nil {
			log.Printf("Error during resubscription to stream %s: %v", stream, err)
			continue
		}

		connection.mu.Lock()
		err = connection.Websocket.WriteMessage(websocket.TextMessage, message)
		connection.mu.Unlock()
		if err != nil {
			log.Printf("Error sending resubscription message for stream %s: %v", stream, err)
			continue
		}
		log.Printf("Resubscribed to stream %s on reconnection", stream)
	}
}

// isConnectionReady checks if the WebSocket connection is open.
//
// @param connection The WebSocketConnection to check.
// @return True if the connection is open, otherwise false.
func (w *WebSocketCommon) isConnectionReady(connection *WebSocketConnection) bool {
	return connection.Connected == OPEN
}

// getAvailableConnections retrieves the list of available WebSocket connections based on the mode.
//
// @return A slice of pointers to WebSocketConnection that are available.
func (w *WebSocketCommon) getAvailableConnections() []*WebSocketConnection {
	if w.Mode == SINGLE {
		return []*WebSocketConnection{w.Connections[0]}
	}

	availableConnections := FilterArrays(w.Connections, w.isConnectionReady)
	return availableConnections
}

// GetConnection retrieves an available WebSocket connection using a round-robin strategy.
//
// @return A pointer to the WebSocketConnection or an error if no connections are available.
func (w *WebSocketCommon) GetConnection() (*WebSocketConnection, error) {
	availableConnections := w.getAvailableConnections()

	if len(availableConnections) == 0 {
		return nil, errors.New("no available websocket connection")
	}

	selectConnection := availableConnections[w.RoundRobinIndex%len(availableConnections)]
	w.RoundRobinIndex = (w.RoundRobinIndex + 1) % len(availableConnections)
	return selectConnection, nil
}

// setupProxyDialer sets up the proxy dialer based on the provided WebSocket configuration.
//
// @param config The WebSocket configuration containing proxy settings.
// @return An error if the proxy dialer setup fails, otherwise nil.
func (w *WebSocketCommon) setupProxyDialer(config WebSocketConfig) error {
	proxyConfig := config.GetProxy()

	if proxyConfig == nil || proxyConfig.IsEmpty() {
		w.ProxyDialer = proxy.Direct
		return nil
	}

	proxyURL, err := proxyConfig.URL()
	if err != nil {
		return fmt.Errorf("invalid proxy configuration: %v", err)
	}

	switch proxyURL.Scheme {
	case "http", "https":
		dialer, err := createHTTPProxyDialer(proxyURL)
		if err != nil {
			return fmt.Errorf("failed to create HTTP proxy dialer: %v", err)
		}
		w.ProxyDialer = dialer

	case "socks5":
		auth := &proxy.Auth{}
		if proxyURL.User != nil {
			auth.User = proxyURL.User.Username()
			if password, ok := proxyURL.User.Password(); ok {
				auth.Password = password
			}
		}

		dialer, err := proxy.SOCKS5("tcp", proxyURL.Host, auth, proxy.Direct)
		if err != nil {
			return fmt.Errorf("failed to create SOCKS5 proxy dialer: %v", err)
		}
		w.ProxyDialer = dialer

	default:
		dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
		if err != nil {
			return fmt.Errorf("unsupported proxy type: %s", proxyURL.Scheme)
		}
		w.ProxyDialer = dialer
	}

	return nil
}

// Ping sends a ping message over the WebSocket connection.
//
// @param conn The WebSocket connection to send the ping on.
// @return An error if the ping fails, otherwise nil.
func (w *WebSocketCommon) Ping(conn *WebSocketConnection) error {
	if conn.Websocket == nil {
		return errors.New("WebSocket connection is not established")
	}

	if conn.Connected != OPEN {
		return errors.New("WebSocket connection is not open")
	}

	log.Printf("Sending ping on WebSocket connection ID: %s", conn.Id)
	return conn.Websocket.WriteMessage(websocket.TextMessage, []byte("ping"))
}

// NewWebsocketAPI creates a new instance of WebsocketAPI.
//
// @param cfg The configuration for the WebSocket API.
// @return A pointer to the WebsocketAPI instance or an error if initialization fails.
func NewWebsocketAPI(cfg *ConfigurationWebsocketApi) (*WebsocketAPI, error) {
	wsCommon, err := NewWebSocketCommon(&ConfigurationWrapper{APIConfig: cfg})
	if err != nil {
		return nil, err
	}
	globalStreamConnectionMap := make(map[string][]*WebSocketConnection)

	return &WebsocketAPI{
		Cfg:                       cfg,
		GlobalStreamConnectionMap: globalStreamConnectionMap,
		WsCommon:                  wsCommon,
	}, nil
}

// Connect establishes the WebSocket connection using the provided user agent.
//
// @param userAgent The user agent string to be used for the connection.
// @return An error if the connection fails, otherwise nil.
func (w *WebsocketAPI) Connect(userAgent string) error {
	return w.WsCommon.Connect(w.Cfg, userAgent)
}

// SendMessage sends a message over the WebSocket connection and returns channels for the response and error.
//
// @param T The type of the expected response.
// @param w The WebsocketAPI instance.
// @param payload The message payload to be sent.
// @param sendParams Parameters for sending the message.
// @return A channel for the typed response, a channel for errors, and an error if sending fails.
func SendMessage[T any](w *WebsocketAPI, payload map[string]any, sendParams SendParams) (chan *ResponseOrRaw[T], chan error, error) {
	conn, err := w.WsCommon.GetConnection()
	if err != nil {
		return nil, nil, err
	}

	params, ok := payload["params"].(map[string]any)
	if !ok {
		return nil, nil, errors.New("payload[params] is not a map")
	}

	skipAuth := !sendParams.WithSessionLogon && conn.SessionLogon

	newParams, err := processParams(params, sendParams.WithAPIKey, skipAuth, w.Cfg.ApiKey)
	if err != nil {
		return nil, nil, err
	}

	var priv crypto.Signer
	if sendParams.Signed {
		priv, err = w.handleSignature(conn, newParams, skipAuth)
		if err != nil {
			return nil, nil, err
		}
	}

	reqID := GenerateUUID()
	if params["id"] != nil {
		reqID = params["id"].(string)
	}
	finalPayload := map[string]interface{}{
		"method": payload["method"],
		"params": newParams,
		"id":     reqID,
	}

	message, err := json.Marshal(finalPayload)
	if err != nil {
		return nil, nil, err
	}

	respChan := make(chan []byte, 1)
	conn.PendingMessages.Store(reqID, respChan)

	conn.mu.Lock()
	err = conn.Websocket.WriteMessage(websocket.TextMessage, message)
	conn.mu.Unlock()
	if err != nil {
		return nil, nil, err
	}

	responseChan := make(chan *ResponseOrRaw[T], 1)
	errorChan := make(chan error, 1)

	go func() {
		select {
		case msg := <-respChan:
			conn.PendingMessages.Delete(reqID)

			if msg == nil {
				errorChan <- errors.New("received nil message")
				return
			}

			if sendParams.WithSessionLogon {
				conn.SessionLogon = true
				conn.SessionLogonRequest = &SessionLogonRequest{
					Method: payload["method"].(string),
					Params: newParams,
					ID:     reqID,
					Priv:   priv,
				}
			}

			var raw map[string]interface{}
			if err := json.Unmarshal(msg, &raw); err != nil {
				errorChan <- err
				return
			}

			if e, ok := raw["error"]; ok {
				errorMap, ok := e.(map[string]interface{})
				if ok {
					errorMsg := fmt.Sprintf("[-%v] %v", errorMap["code"], errorMap["msg"])
					errorChan <- &WebSocketError{
						ConnID:  conn.Id,
						Op:      "error_response",
						Err:     fmt.Errorf("%v", errorMap),
						Message: errorMsg,
					}
					return
				}
				errorChan <- &WebSocketError{
					ConnID:  conn.Id,
					Op:      "error_response",
					Err:     fmt.Errorf("server error: %v", e),
					Message: "received error response from server",
				}
				return
			}
			if status, ok := raw["status"].(float64); ok && status >= 400 {
				errorChan <- &WebSocketError{
					ConnID:  conn.Id,
					Op:      "error_response",
					Err:     fmt.Errorf("server error: status %v", status),
					Message: "received error status from server",
				}
				return
			}

			var data T
			if err := json.Unmarshal(msg, &data); err != nil {
				responseChan <- &ResponseOrRaw[T]{Raw: raw}
				errorChan <- nil
				return
			}

			responseChan <- &ResponseOrRaw[T]{Typed: &data}
			errorChan <- nil
		case <-time.After(w.Cfg.Timeout):
			errorChan <- errors.New("timeout waiting for response")
		}
	}()

	return responseChan, errorChan, nil
}

// handleSignature handles the signing of the request parameters.
//
// @param conn The WebSocket connection.
// @param newParams The parameters to be signed.
// @param skipAuth Whether to skip authentication.
// @return The crypto.Signer used for signing and an error if signing fails.
func (w *WebsocketAPI) handleSignature(conn *WebSocketConnection, newParams map[string]any, skipAuth bool) (crypto.Signer, error) {
	priv := crypto.Signer(nil)
	newParams["timestamp"] = GetTimestamp()
	if !skipAuth {
		newParams["apiKey"] = w.Cfg.ApiKey

		if w.Cfg.PrivateKey != "" {
			if w.Cfg.Signer == nil {
				var err error
				priv, err = LoadPrivateKey(w.Cfg.PrivateKey, w.Cfg.PrivateKeyPassphrase)
				if err != nil {
					return nil, &WebSocketError{
						ConnID:  conn.Id,
						Op:      "load_private_key",
						Err:     err,
						Message: "failed to load private key for signing",
					}
				}
				w.Cfg.Signer = &cryptoSigner{s: priv}
			}

			signature, err := w.Cfg.Signer.Sign([]byte(Urlencode(newParams)))
			if err != nil {
				return nil, &WebSocketError{
					ConnID:  conn.Id,
					Op:      "sign_message",
					Err:     err,
					Message: "failed to sign request",
				}
			}
			newParams["signature"] = signature
		} else {
			signer := hmac.New(sha256.New, []byte(w.Cfg.ApiSecret))
			signer.Write([]byte(Urlencode(newParams)))
			signatureBytes := signer.Sum(nil)
			newParams["signature"] = hex.EncodeToString(signatureBytes)
		}
	}
	return priv, nil
}

// Subscribe subscribes to the specified stream ID.
//
// @param id The stream ID to subscribe to.
// @return An error if the subscription fails, otherwise nil.
func (w *WebsocketAPI) Subscribe(id string) error {
	conn, err := w.WsCommon.GetConnection()
	if err != nil {
		return err
	}

	w.GlobalStreamConnectionMap[id] = append(w.GlobalStreamConnectionMap[id], conn)
	conn.StreamConnectionMap = append(conn.StreamConnectionMap, id)
	return nil
}

// Unsubscribe unsubscribes from the specified stream ID.
//
// @param id The stream ID to unsubscribe from.
// @return An error if the unsubscription fails, otherwise nil.
func (w *WebsocketAPI) Unsubscribe(id string) error {
	if _, ok := w.GlobalStreamConnectionMap[id]; !ok {
		log.Printf("Stream %s not associated with an active connection.", id)
		return nil
	}

	conn := w.GlobalStreamConnectionMap[id][0]

	conn.mu.Lock()
	defer conn.mu.Unlock()

	delete(w.GlobalStreamConnectionMap, id)
	return nil
}

// CloseWebSocketConnection closes all active WebSocket connections.
//
// @return An error if any connection fails to close, otherwise nil.
func (w *WebsocketAPI) CloseWebSocketConnection() error {
	for _, conn := range w.WsCommon.Connections {
		log.Println("Closing WebSocket connection:", conn.Id)

		if conn.Websocket != nil {
			conn.Connected = CLOSING
			err := conn.Websocket.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewWebsocketStreams creates a new instance of WebsocketStreams.
//
// @param cfg The configuration for the WebSocket streams.
// @return A pointer to the WebsocketStreams instance or an error if initialization fails.
func NewWebsocketStreams(cfg *ConfigurationWebsocketStreams) (*WebsocketStreams, error) {
	wsCommon, err := NewWebSocketCommon(&ConfigurationWrapper{StreamsConfig: cfg})
	if err != nil {
		return nil, err
	}
	globalStreamConnectionMap := make(map[string][]*WebSocketConnection)

	return &WebsocketStreams{
		Cfg:                       cfg,
		WsCommon:                  wsCommon,
		GlobalStreamConnectionMap: globalStreamConnectionMap,
	}, nil
}

func (w *WebsocketStreams) Connect(userAgent string) error {
	return w.WsCommon.Connect(w.Cfg, userAgent)
}

// Subscribe subscribes to the specified streams.
//
// @param streams A slice of stream names to subscribe to.
// @param id A slice of IDs corresponding to each stream. If empty, new UUIDs will be generated.
// @return An error if any subscription fails, otherwise nil.
func (w *WebsocketStreams) Subscribe(streams []string, id []any, strictInt bool) error {
	if len(streams) == 0 {
		return errors.New("no streams to subscribe")
	}

	for num, stream := range streams {
		if _, ok := w.GlobalStreamConnectionMap[stream]; ok {
			continue
		}

		conn, err := w.WsCommon.GetConnection()
		if err != nil {
			return err
		}

		var streamID interface{}
		if strictInt {
			streamID = GenerateIntUUID()
		} else {
			streamID = GenerateUUID()
		}

		if len(id) > num {
			streamID = id[num]
		}
		subscribePayload := map[string]interface{}{
			"method": "SUBSCRIBE",
			"params": []string{stream},
			"id":     streamID,
		}

		message, err := json.Marshal(subscribePayload)
		if err != nil {
			return err
		}

		log.Printf("Sending subscription1: %s", string(message))
		conn.mu.Lock()
		err = conn.Websocket.WriteMessage(websocket.TextMessage, message)
		conn.mu.Unlock()
		if err != nil {
			return err
		}

		w.GlobalStreamConnectionMap[stream] = append(w.GlobalStreamConnectionMap[stream], conn)
		conn.StreamConnectionMap = append(conn.StreamConnectionMap, stream)
	}
	return nil
}

// On registers a callback function for the specified stream.
//
// @param stream The name of the stream to register the callback for.
// @param callback The callback function to be invoked when a message is received for the specified stream.
// @return An error if the stream is not subscribed, otherwise nil.
func (w *WebsocketStreams) On(stream string, callback func(map[string]interface{})) error {
	if _, ok := w.GlobalStreamConnectionMap[stream]; !ok {
		return fmt.Errorf("stream %s not subscribed", stream)
	}

	connections := w.GlobalStreamConnectionMap[stream]

	for _, conn := range connections {
		conn.mu.Lock()
		conn.StreamCallbackMap[stream] = append(conn.StreamCallbackMap[stream], callback)
		conn.mu.Unlock()
	}
	return nil
}

// ListSubscriptions lists the current subscriptions for the specified stream ID.
//
// @param streamId The stream ID to list subscriptions for. If empty, a new UUID will be generated.
// @return A map containing the subscription details or an error if the operation fails.
func (w *WebsocketStreams) ListSubscriptions(streamId string) (map[string]interface{}, error) {
	conn, err := w.WsCommon.GetConnection()
	if err != nil {
		return nil, err
	}

	id := GenerateUUID()
	if streamId != "" {
		id = streamId
	}
	subscribePayload := map[string]interface{}{
		"method": "LIST_SUBSCRIPTIONS",
		"id":     id,
	}

	message, err := json.Marshal(subscribePayload)
	if err != nil {
		return nil, err
	}

	respChan := make(chan []byte, 1)
	conn.PendingMessages.Store(id, respChan)

	conn.mu.Lock()
	err = conn.Websocket.WriteMessage(websocket.TextMessage, message)
	conn.mu.Unlock()
	if err != nil {
		return nil, err
	}

	select {
	case msg := <-respChan:
		conn.PendingMessages.Delete(id)

		if msg == nil {
			return nil, errors.New("received nil message")
		}

		var raw map[string]interface{}
		if err := json.Unmarshal(msg, &raw); err != nil {
			return nil, err
		}

		return raw, nil
	case <-time.After(10 * time.Second):
		return nil, errors.New("timeout waiting for response")
	}
}

// Unsubscribe unsubscribes from the specified streams.
//
// @param streams A slice of stream names to unsubscribe from.
// @return An error if any unsubscription fails, otherwise nil.
func (w *WebsocketStreams) Unsubscribe(streams []string) error {
	if len(streams) == 0 {
		return errors.New("no streams to unsubscribe")
	}

	for _, stream := range streams {
		if _, ok := w.GlobalStreamConnectionMap[stream]; !ok {
			log.Printf("Stream %s not associated with an active connection.", stream)
			continue
		}

		conn := w.GlobalStreamConnectionMap[stream][0]

		unsubscribePayload := map[string]interface{}{
			"method": "UNSUBSCRIBE",
			"params": []string{stream},
			"id":     GenerateUUID(),
		}

		message, err := json.Marshal(unsubscribePayload)
		if err != nil {
			return err
		}

		conn.mu.Lock()
		err = conn.Websocket.WriteMessage(websocket.TextMessage, message)
		conn.mu.Unlock()
		if err != nil {
			return err
		}

		delete(w.GlobalStreamConnectionMap, stream)
		conn.mu.Lock()
		var updatedStreams []string
		for _, s := range conn.StreamConnectionMap {
			if s != stream {
				updatedStreams = append(updatedStreams, s)
			}
		}
		conn.StreamConnectionMap = updatedStreams
		conn.mu.Unlock()

		conn.mu.Lock()
		delete(conn.StreamCallbackMap, stream)
		conn.mu.Unlock()

		log.Printf("Unsubscribed from stream %s", stream)
	}

	return nil
}

// IsSubscribed checks if the specified stream is currently subscribed.
//
// @param stream The name of the stream to check.
// @return True if the stream is subscribed, otherwise false.
func (w *WebsocketStreams) IsSubscribed(stream string) bool {
	if _, ok := w.GlobalStreamConnectionMap[stream]; ok {
		return true
	} else {
		return false
	}
}

// CloseWebSocketStreamConnection closes all active WebSocket connections in the WebsocketStreams instance.
//
// @return An error if any connection fails to close, otherwise nil.
func (w *WebsocketStreams) CloseWebSocketStreamConnection() error {
	for _, conn := range w.WsCommon.Connections {
		log.Println("Closing WebSocket connection:", conn.Id)

		if conn.Websocket != nil {
			conn.Connected = CLOSING
			err := conn.Websocket.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateStreamHandler creates a new StreamHandler for the specified stream and registers it with the provided StreamHandlerWrapper.
//
// @param wrapper The StreamHandlerWrapper containing either a WebsocketStreams or WebsocketAPI instance.
// @param stream The name of the stream to handle.
// @param id An optional slice of IDs associated with the stream.
// @return A pointer to the created StreamHandler or an error if subscription fails.
func CreateStreamHandler[T any](wrapper *StreamHandlerWrapper, stream string, id []any, strictInt bool) (*StreamHandler[T], error) {
	if wrapper.WebsocketStreams != nil {
		if err := wrapper.WebsocketStreams.Subscribe([]string{stream}, id, strictInt); err != nil {
			return nil, err
		}
	} else if wrapper.WebsocketAPI != nil {
		if err := wrapper.WebsocketAPI.Subscribe(stream); err != nil {
			return nil, err
		}
	} else {
		return nil, &WebSocketError{
			Op:      "create_stream_handler",
			Err:     errors.New("invalid StreamHandlerWrapper"),
			Message: "invalid StreamHandlerWrapper: neither WebsocketStreams nor WebsocketAPI is set",
		}
	}
	return &StreamHandler[T]{Stream: stream, Wrapper: wrapper}, nil
}

// On registers a callback for the specified event on the stream associated with the StreamHandler.
//
// @param event The event type to listen for (e.g., "message").
// @param callback The callback function to handle incoming messages of type T.
func (h *StreamHandler[T]) On(event string, callback func(T)) {
	h.Callback = callback

	connections := []*WebSocketConnection{}
	exists := false
	if h.Wrapper.WebsocketStreams != nil {
		connections, exists = h.Wrapper.WebsocketStreams.GlobalStreamConnectionMap[h.Stream]
		if !exists {
			return
		}
	} else if h.Wrapper.WebsocketAPI != nil {
		connections, exists = h.Wrapper.WebsocketAPI.GlobalStreamConnectionMap[h.Stream]
		if !exists {
			return
		}
	}

	if event != "message" {
		return
	}

	for _, conn := range connections {
		conn.mu.Lock()

		handler := func(msg map[string]interface{}) {
			var dataToProcess interface{}
			if data, exists := msg["data"]; exists {
				dataToProcess = data
			} else {
				dataToProcess = msg
			}

			jsonData, err := json.Marshal(dataToProcess)
			if err != nil {
				log.Printf("Failed to marshal data for stream %s: %v", h.Stream, err)
				return
			}

			var singleItem T
			if err := json.Unmarshal(jsonData, &singleItem); err == nil {
				callback(singleItem)
				return
			}

			var items []T
			if err := json.Unmarshal(jsonData, &items); err != nil {
				log.Printf("Failed to unmarshal data for stream %s: %v", h.Stream, err)
				return
			}

			for _, item := range items {
				callback(item)
			}
		}

		conn.StreamCallbackMap[h.Stream] = append(conn.StreamCallbackMap[h.Stream], handler)
		conn.mu.Unlock()
	}
}

// OnError registers an error callback for the stream associated with the StreamHandler.
//
// @param cb The callback function to handle errors.
func (h *StreamHandler[T]) OnError(cb func(error)) {
	h.ErrorCb = cb
	connections, exists := h.Wrapper.WebsocketStreams.GlobalStreamConnectionMap[h.Stream]
	if !exists {
		return
	}

	log.Println("\n\nRegistering OnError handlers for stream:", h.Stream)
	for _, conn := range connections {
		go func(conn *WebSocketConnection) {
			for err := range conn.ErrorChan {
				cb(err)
			}
		}(conn)
	}
}

// Unsubscribe unsubscribes from the stream associated with the StreamHandler.
func (h *StreamHandler[T]) Unsubscribe() {
	if h.Wrapper.WebsocketStreams != nil {
		if err := h.Wrapper.WebsocketStreams.Unsubscribe([]string{h.Stream}); err != nil {
			log.Printf("Error unsubscribing from stream %s: %v", h.Stream, err)
		}
	}
}
