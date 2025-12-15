package binanceMarginTrading

import (
	BinanceMarginTradingRestApi "github.com/binance/binance-connector-go/clients/margintrading/src/restapi"
	BinanceMarginTradingWebsocketStreams "github.com/binance/binance-connector-go/clients/margintrading/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceMarginTradingClient struct {
	RestApi          *BinanceMarginTradingRestApi.RestAPIClient
	WebsocketStreams *BinanceMarginTradingWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceMarginTradingClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceMarginTradingClient) {
		c.RestApi = BinanceMarginTradingRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceMarginTradingClient) {
		c.WebsocketStreams = BinanceMarginTradingWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceMarginTradingClient(opts ...Option) *BinanceMarginTradingClient {
	client := &BinanceMarginTradingClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
