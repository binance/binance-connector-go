/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package binancederivativestradingportfoliomarginwebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Derivatives Trading Portfolio Margin WebSocket Market Streams WebSocket Streams v1.4.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
}

// NewWebsocketStreamsClient creates a new Binance Binance Derivatives Trading Portfolio Margin WebSocket Market Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-derivativestradingportfoliomargin/1.4.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	wsClient, err := common.NewWebsocketStreams(c.cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.Ws = wsClient

	// API Services

	return c
}

// Connect establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketStreamsClient) Connect(streams []string) error {
	return c.Ws.Connect(c.userAgent, streams)
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
