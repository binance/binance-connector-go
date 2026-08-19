package stocks

import (
	BinanceStocksRestApi "github.com/binance/binance-connector-go/clients/stocks/src/restapi"
	BinanceStocksWebsocketStreams "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceStocksClient struct {
	RestApi          *BinanceStocksRestApi.RestAPIClient
	WebsocketStreams *BinanceStocksWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceStocksClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceStocksClient) {
		c.RestApi = BinanceStocksRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceStocksClient) {
		c.WebsocketStreams = BinanceStocksWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceStocksClient(opts ...Option) *BinanceStocksClient {
	client := &BinanceStocksClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
