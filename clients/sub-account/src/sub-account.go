package binanceSubAccount

import (
	BinanceSubAccountRestApi "github.com/binance/binance-connector-go/clients/subaccount/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceSubAccountClient struct {
	RestApi *BinanceSubAccountRestApi.RestAPIClient
}

type Option func(*BinanceSubAccountClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceSubAccountClient) {
		c.RestApi = BinanceSubAccountRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceSubAccountClient(opts ...Option) *BinanceSubAccountClient {
	client := &BinanceSubAccountClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
