/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package binancemargintradingwebsocketstreams

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/clients/margintrading/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

// WebsocketStreamsClient manages communication with the Binance Binance Margin Trading WebSocket Market Streams WebSocket Streams v1.2.0
type WebsocketStreamsClient struct {
	cfg       *common.ConfigurationWebsocketStreams
	userAgent string
	Ws        *common.WebsocketStreams

	// API Services
}

// NewWebsocketStreamsClient creates a new Binance Binance Margin Trading WebSocket Market Streams WebSocket Streams client
//
// @param cfg *common.ConfigurationWebsocketStreams - The configuration for the WebSocket Streams client
// @return *WebsocketStreamsClient - The newly created WebSocket Streams client
func NewWebsocketStreamsClient(cfg *common.ConfigurationWebsocketStreams) *WebsocketStreamsClient {
	c := &WebsocketStreamsClient{cfg: cfg}
	c.userAgent = "binance-margintrading/1.2.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"

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
func (c *WebsocketStreamsClient) Connect() error {

	return c.Ws.Connect(c.userAgent)
}

// RiskData subscribes to risk data stream events with an optional ID
//
// @param listenKey string - The listen key for the risk data stream
// @param id string
// @return *common.StreamHandler[models.RiskDataStreamEventsResponse] - The stream handler for risk data stream events
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) RiskData(listenKey string, id any) (*common.StreamHandler[models.RiskDataStreamEventsResponse], error) {
	if id != "" {
		return common.CreateStreamHandler[models.RiskDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, []any{id}, false,
		)
	} else {
		return common.CreateStreamHandler[models.RiskDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, nil, false)
	}
}

// TradeData subscribes to trade data stream events with an optional ID
//
// @param listenKey string - The listen key for the trade data stream
// @param id string - The optional ID for the subscription
// @return *common.StreamHandler[models.TradeDataStreamEventsResponse] - The stream handler for trade data stream events
// @return error - An error if the operation fails
func (c *WebsocketStreamsClient) TradeData(listenKey string, id any) (*common.StreamHandler[models.TradeDataStreamEventsResponse], error) {
	if id != "" {
		return common.CreateStreamHandler[models.TradeDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketStreams: c.Ws,
		}, listenKey, []any{id}, false,
		)
	} else {
		return common.CreateStreamHandler[models.TradeDataStreamEventsResponse](&common.StreamHandlerWrapper{
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
