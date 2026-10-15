/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package binancecmswebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// WebsocketStreamsClient manages communication with the Binance Prediction Trading Streams WebSocket Streams v1.0.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
	DefaultAPI *DefaultAPIService
}

// NewWebsocketStreamsClient creates a new Binance Prediction Trading Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}

	// Signed streams sign the connection URL with the configured credentials
	cfg.Signed = true

	c.userAgent = "binance-cms/1.0.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	wsClient, err := common.NewWebsocketStreams(c.cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.Ws = wsClient

	// API Services
	c.DefaultAPI = &DefaultAPIService{client: c}

	return c
}

type Service struct {
	client *WebsocketStreamsClient
}

// Connect establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) Connect(streams []string) error {

	return c.Ws.Connect(c.userAgent, streams)
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
