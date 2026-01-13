/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package binancespotwebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/common/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Spot WebSocket Streams WebSocket Streams v1.1.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
	WebSocketStreamsAPI *WebSocketStreamsAPIService
}

// NewWebsocketStreamsClient creates a new Binance Binance Spot WebSocket Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-spot/1.1.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	wsClient, err := common.NewWebsocketStreams(c.cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.Ws = wsClient

	// API Services
	c.WebSocketStreamsAPI = &WebSocketStreamsAPIService{client: c}

	return c
}

type Service struct {
	client *WebsocketStreamsClient
}

// Connect establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) Connect() error {
	return c.Ws.Connect(c.userAgent)
}

// Subscribe subscribes to the specified streams with optional IDs
//
// @param streams []string - The list of streams to subscribe to
// @param id []string - The optional list of IDs for the subscriptions
// @return error - An error if the subscription fails
func (c *WebsocketStreamsClient) Subscribe(streams []string, id []any) error {
	return c.Ws.Subscribe(streams, id, false)
}

// On registers a callback for the specified stream
//
// @param stream string - The stream name
// @param callback func(map[string]interface{}) - The callback function to handle messages
// @return error - An error if registering the callback fails
func (c *WebsocketStreamsClient) On(stream string, callback func(map[string]interface{})) error {
	return c.Ws.On(stream, callback)
}

// ListSubscriptions lists the current subscriptions for the given ID
//
// @param id string - The subscription ID
// @return map[string]interface{} - The list of current subscriptions
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) ListSubscriptions(id string) (map[string]interface{}, error) {
	return c.Ws.ListSubscriptions(id)
}

// Unsubscribe unsubscribes from the specified streams
//
// @param streams []string - The list of streams to unsubscribe from
// @return error - An error if the unsubscription fails
func (c *WebsocketStreamsClient) Unsubscribe(streams []string) error {
	return c.Ws.Unsubscribe(streams)
}

// CloseWebSocketStreamConnection closes the WebSocket stream connection
//
// @return error - An error if closing the connection fails
func (c *WebsocketStreamsClient) CloseWebSocketStreamConnection() error {
	return c.Ws.CloseWebSocketStreamConnection()
}
