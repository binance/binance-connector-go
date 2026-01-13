package derivativestradingportfoliomargin

import (
	BinanceDerivativesTradingPortfolioMarginRestApi "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/restapi"
	BinanceDerivativesTradingPortfolioMarginWebsocketStreams "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceDerivativesTradingPortfolioMarginClient struct {
	RestApi          *BinanceDerivativesTradingPortfolioMarginRestApi.RestAPIClient
	WebsocketStreams *BinanceDerivativesTradingPortfolioMarginWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceDerivativesTradingPortfolioMarginClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDerivativesTradingPortfolioMarginClient) {
		c.RestApi = BinanceDerivativesTradingPortfolioMarginRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceDerivativesTradingPortfolioMarginClient) {
		c.WebsocketStreams = BinanceDerivativesTradingPortfolioMarginWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceDerivativesTradingPortfolioMarginClient(opts ...Option) *BinanceDerivativesTradingPortfolioMarginClient {
	client := &BinanceDerivativesTradingPortfolioMarginClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
