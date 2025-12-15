package binanceDerivativesTradingOptions

import (
	BinanceDerivativesTradingOptionsRestApi "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi"
	BinanceDerivativesTradingOptionsWebsocketStreams "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceDerivativesTradingOptionsClient struct {
	RestApi          *BinanceDerivativesTradingOptionsRestApi.RestAPIClient
	WebsocketStreams *BinanceDerivativesTradingOptionsWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingOptionsClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingOptionsClient) {
		c.RestApi = BinanceDerivativesTradingOptionsRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingOptionsClient) {
		c.WebsocketStreams = BinanceDerivativesTradingOptionsWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingOptionsClient(opts ...Option) *BinanceDerivativesTradingOptionsClient {
	client := &BinanceDerivativesTradingOptionsClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
