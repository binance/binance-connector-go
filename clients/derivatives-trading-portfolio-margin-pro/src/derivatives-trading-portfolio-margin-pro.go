package binanceDerivativesTradingPortfolioMarginPro

import (
	BinanceDerivativesTradingPortfolioMarginProRestApi "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro/src/restapi"
	BinanceDerivativesTradingPortfolioMarginProWebsocketStreams "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceDerivativesTradingPortfolioMarginProClient struct {
	RestApi          *BinanceDerivativesTradingPortfolioMarginProRestApi.RestAPIClient
	WebsocketStreams *BinanceDerivativesTradingPortfolioMarginProWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingPortfolioMarginProClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingPortfolioMarginProClient) {
		c.RestApi = BinanceDerivativesTradingPortfolioMarginProRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingPortfolioMarginProClient) {
		c.WebsocketStreams = BinanceDerivativesTradingPortfolioMarginProWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingPortfolioMarginProClient(opts ...Option) *BinanceDerivativesTradingPortfolioMarginProClient {
	client := &BinanceDerivativesTradingPortfolioMarginProClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
