package fiat

import (
	BinanceFiatRestApi "github.com/binance/binance-connector-go/clients/fiat/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceFiatClient struct {
	RestApi *BinanceFiatRestApi.RestAPIClient
}

type Option func(*BinanceFiatClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceFiatClient) {
		c.RestApi = BinanceFiatRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceFiatClient(opts ...Option) *BinanceFiatClient {
	client := &BinanceFiatClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
