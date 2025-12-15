/*
Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams

API version: 1.0.0
*/

package binancederivativestradingportfoliomarginprowebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams WebSocket Streams v1.0.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
}

// NewWebsocketStreamsClient creates a new Binance Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-derivativestradingportfoliomarginpro/1.0.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

	wsClient, err := common.NewWebsocketStreams(c.cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket Streams client: %v", err)
	}
	c.Ws = wsClient

	// API Services

	return c
}

// UserData subscribes to user data stream events with an optional ID
//
// @param listenKey string - The listen key for the user data stream
// @param id string - The optional ID for the subscription
// @return *common.StreamHandler[models.UserDataStreamEventsResponse] - The stream handler for user data stream events
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) UserData(listenKey string, id string) (*common.StreamHandler[models.UserDataStreamEventsResponse], error) {
	if id != "" {
		return common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, []string{id},
		)
	} else {
		return common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, nil)
	}
}
