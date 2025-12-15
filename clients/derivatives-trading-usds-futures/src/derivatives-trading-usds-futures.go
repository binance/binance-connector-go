package binanceDerivativesTradingUsdsFutures

import (
	BinanceDerivativesTradingUsdsFuturesRestApi "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi"
	BinanceDerivativesTradingUsdsFuturesWebsocketApi "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi"
	BinanceDerivativesTradingUsdsFuturesWebsocketStreams "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceDerivativesTradingUsdsFuturesClient struct {
	RestApi          *BinanceDerivativesTradingUsdsFuturesRestApi.RestAPIClient
	WebsocketAPI     *BinanceDerivativesTradingUsdsFuturesWebsocketApi.WebsocketAPIClient
	WebsocketStreams *BinanceDerivativesTradingUsdsFuturesWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingUsdsFuturesClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.RestApi = BinanceDerivativesTradingUsdsFuturesRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketAPI(cfg *common.ConfigurationWebsocketApi) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.WebsocketAPI = BinanceDerivativesTradingUsdsFuturesWebsocketApi.NewWebsocketAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingUsdsFuturesClient) {
		c.WebsocketStreams = BinanceDerivativesTradingUsdsFuturesWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingUsdsFuturesClient(opts ...Option) *BinanceDerivativesTradingUsdsFuturesClient {
	client := &BinanceDerivativesTradingUsdsFuturesClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
