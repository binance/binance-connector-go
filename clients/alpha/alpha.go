package alpha

import (
	BinanceAlphaRestApi "github.com/binance/binance-connector-go/clients/alpha/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceAlphaClient struct {
	RestApi *BinanceAlphaRestApi.RestAPIClient
}

type Option func(*BinanceAlphaClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceAlphaClient) {
		c.RestApi = BinanceAlphaRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceAlphaClient(opts ...Option) *BinanceAlphaClient {
	client := &BinanceAlphaClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
