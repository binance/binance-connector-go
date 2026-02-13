/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package binancederivativestradingoptionswebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Derivatives Trading Options WebSocket Market Streams WebSocket Streams v1.3.0
type WebsocketStreamsClient struct {
	cfg        *common.ConfigurationWebsocketStreams
	userAgent  string
	WsMarket   *common.WebsocketStreams
	WsPublic   *common.WebsocketStreams
	WsUserData *common.WebsocketStreams

	// API Services
	MarketAPI *MarketAPIService
	PublicAPI *PublicAPIService
}

// NewWebsocketStreamsClient creates a new Binance Binance Derivatives Trading Options WebSocket Market Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-derivativestradingoptions/1.3.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	cfgMarket := *cfg
	cfgMarket.BasePath = cfgMarket.BasePath + "/market/stream"
	wsClientMarket, err := common.NewWebsocketStreams(&cfgMarket)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams common client: %v", err)
	}
	c.WsMarket = wsClientMarket

	cfgPublic := *cfg
	cfgPublic.BasePath = cfgPublic.BasePath + "/public/stream"
	wsClientPublic, err := common.NewWebsocketStreams(&cfgPublic)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams common client: %v", err)
	}
	c.WsPublic = wsClientPublic

	cfgUserData := *cfg
	cfgUserData.BasePath = cfgUserData.BasePath + "/private/stream"
	wsClientUserData, err := common.NewWebsocketStreams(&cfgUserData)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.WsUserData = wsClientUserData

	// API Services
	c.MarketAPI = &MarketAPIService{client: c}
	c.PublicAPI = &PublicAPIService{client: c}

	return c
}

type Service struct {
	client *WebsocketStreamsClient
}

// Connect establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) Connect(streams []string) error {
	var err error
	err = c.ConnectMarket(streams)
	if err != nil {
		return err
	}

	err = c.ConnectPublic(streams)
	if err != nil {
		return err
	}

	return nil
}

// ConnectMarket establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) ConnectMarket(streams []string) error {
	err := c.WsMarket.Connect(c.userAgent, streams)
	if err != nil {
		return err
	}

	return nil
}

// SubscribeMarket subscribes to the specified streams with optional IDs
//
// @param streams []string - The list of streams to subscribe to
// @param id []string - The optional list of IDs for the subscriptions
// @return error - An error if the subscription fails
func (c *WebsocketStreamsClient) SubscribeMarket(streams []string, id []any) error {
	return c.WsMarket.Subscribe(streams, id, true)
}

// OnMarket registers a callback for the specified stream
//
// @param stream string - The stream name
// @param callback func(map[string]interface{}) - The callback function to handle messages
// @return error - An error if registering the callback fails
func (c *WebsocketStreamsClient) OnMarket(stream string, callback func(map[string]interface{})) error {
	return c.WsMarket.On(stream, callback)
}

// ListSubscriptionsMarket lists the current subscriptions for the given ID
//
// @param id string - The subscription ID
// @return map[string]interface{} - The list of current subscriptions
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) ListSubscriptionsMarket(id string) (map[string]interface{}, error) {
	return c.WsMarket.ListSubscriptions(id)
}

// UnsubscribeMarket unsubscribes from the specified streams
//
// @param streams []string - The list of streams to unsubscribe from
// @return error - An error if the unsubscription fails
func (c *WebsocketStreamsClient) UnsubscribeMarket(streams []string) error {
	return c.WsMarket.Unsubscribe(streams)
}

// ConnectPublic establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) ConnectPublic(streams []string) error {
	err := c.WsPublic.Connect(c.userAgent, streams)
	if err != nil {
		return err
	}

	return nil
}

// SubscribePublic subscribes to the specified streams with optional IDs
//
// @param streams []string - The list of streams to subscribe to
// @param id []string - The optional list of IDs for the subscriptions
// @return error - An error if the subscription fails
func (c *WebsocketStreamsClient) SubscribePublic(streams []string, id []any) error {
	return c.WsPublic.Subscribe(streams, id, true)
}

// OnPublic registers a callback for the specified stream
//
// @param stream string - The stream name
// @param callback func(map[string]interface{}) - The callback function to handle messages
// @return error - An error if registering the callback fails
func (c *WebsocketStreamsClient) OnPublic(stream string, callback func(map[string]interface{})) error {
	return c.WsPublic.On(stream, callback)
}

// ListSubscriptionsPublic lists the current subscriptions for the given ID
//
// @param id string - The subscription ID
// @return map[string]interface{} - The list of current subscriptions
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) ListSubscriptionsPublic(id string) (map[string]interface{}, error) {
	return c.WsPublic.ListSubscriptions(id)
}

// UnsubscribePublic unsubscribes from the specified streams
//
// @param streams []string - The list of streams to unsubscribe from
// @return error - An error if the unsubscription fails
func (c *WebsocketStreamsClient) UnsubscribePublic(streams []string) error {
	return c.WsPublic.Unsubscribe(streams)
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
			WebsocketStreams: c.WsUserData,
		}, listenKey, []any{id}, true,
		)
	} else {
		return common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.WsUserData,
		}, listenKey, nil, true)
	}
}

// CloseWebSocketStreamConnection closes the WebSocket stream connection
//
// @return error - An error if closing the connection fails
func (c *WebsocketStreamsClient) CloseWebSocketStreamConnection() error {
	var err error
	err = c.WsMarket.CloseWebSocketStreamConnection()
	if err != nil {
		return err
	}

	err = c.WsPublic.CloseWebSocketStreamConnection()
	if err != nil {
		return err
	}

	return nil
}
