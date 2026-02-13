/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package binancederivativestradingcoinfutureswebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Derivatives Trading COIN Futures WebSocket Market Streams WebSocket Streams v1.4.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
	WebsocketMarketStreamsAPI *WebsocketMarketStreamsAPIService
}

// NewWebsocketStreamsClient creates a new Binance Binance Derivatives Trading COIN Futures WebSocket Market Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-derivativestradingcoinfutures/1.4.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	wsClient, err := common.NewWebsocketStreams(c.cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.Ws = wsClient

	// API Services
	c.WebsocketMarketStreamsAPI = &WebsocketMarketStreamsAPIService{client: c}

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

// UserData subscribes to user data stream events with an optional ID
//
// @param listenKey string - The listen key for the user data stream
// @param id string - The optional ID for the subscription
// @return *common.StreamHandler[models.UserDataStreamEventsResponse] - The stream handler for user data stream events
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) UserData(listenKey string, id any) (*common.StreamHandler[models.UserDataStreamEventsResponse], error) {
	if id != "" {
		return common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, []any{id}, false,
		)
	} else {
		return common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, nil, false)
	}
}

// CloseWebSocketStreamConnection closes the WebSocket stream connection
//
// @return error - An error if closing the connection fails
func (c *WebsocketStreamsClient) CloseWebSocketStreamConnection() error {

	return c.Ws.CloseWebSocketStreamConnection()
}
