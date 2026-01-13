/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package binancespotwebsocketapi

import (
	"log"
	"runtime"

	"github.com/binance/binance-connector-go/common/common"
)

// WebsocketAPIClient manages communication with the Binance Binance Spot WebSocket API WebSocketAPI v1.1.0
type WebsocketAPIClient struct {
	cfg       *common.ConfigurationWebsocketApi
	userAgent string
	Ws        *common.WebsocketAPI

	// API Services
	AccountAPI        *AccountAPIService
	AuthAPI           *AuthAPIService
	GeneralAPI        *GeneralAPIService
	MarketAPI         *MarketAPIService
	TradeAPI          *TradeAPIService
	UserDataStreamAPI *UserDataStreamAPIService
}

// NewWebsocketAPIClient creates a new Binance Binance Spot WebSocket API WebSocket API client
//
// @param cfg *common.ConfigurationWebsocketApi - The configuration for the WebSocket API client
// @return *WebsocketAPIClient - The newly created WebSocket API client
func NewWebsocketAPIClient(cfg *common.ConfigurationWebsocketApi) *WebsocketAPIClient {
	wsClient, err := common.NewWebsocketAPI(cfg)
	if err != nil {
		log.Fatalf("Error creating WebSocket client: %v", err)
	}

	c := &WebsocketAPIClient{
		cfg:               cfg,
		userAgent:         "binance-spot/1.1.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")",
		Ws:                wsClient,
		AccountAPI:        &AccountAPIService{Ws: wsClient},
		AuthAPI:           &AuthAPIService{Ws: wsClient},
		GeneralAPI:        &GeneralAPIService{Ws: wsClient},
		MarketAPI:         &MarketAPIService{Ws: wsClient},
		TradeAPI:          &TradeAPIService{Ws: wsClient},
		UserDataStreamAPI: &UserDataStreamAPIService{Ws: wsClient},
	}

	return c
}

// SendMessage sends a message through the WebSocket API client
//
// # T is a generic type parameter representing the expected response type
//
// @param c *common.WebsocketAPI - The WebSocket API client
// @param payload map[string]any - The message payload
// @param sendParams common.SendParams - Parameters for sending the message
// @return chan *common.ResponseOrRaw[T] - Channel for receiving responses
// @return chan error - Channel for receiving errors
// @return error - An error if sending the message fails
func SendMessage[T any](c *common.WebsocketAPI, payload map[string]any, sendParams common.SendParams) (chan *common.ResponseOrRaw[T], chan error, error) {
	return common.SendMessage[T](c, payload, sendParams)
}

// Connect establishes the WebSocket connection
//
// @return error - An error if the connection fails
func (c *WebsocketAPIClient) Connect() error {
	return c.Ws.Connect(c.userAgent)
}

// Ping sends a ping to the WebSocket server
//
// @param connection *common.WebSocketConnection - The WebSocket connection
// @return error - An error if the ping fails
func (c *WebsocketAPIClient) Ping(connection *common.WebSocketConnection) error {
	return c.Ws.WsCommon.Ping(connection)
}

// CloseWebSocketConnection closes the WebSocket connection
//
// @return error - An error if closing the connection fails
func (c *WebsocketAPIClient) CloseWebSocketConnection() error {
	return c.Ws.CloseWebSocketConnection()
}
