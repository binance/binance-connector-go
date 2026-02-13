package c2c

import (
	BinanceC2CRestApi "github.com/binance/binance-connector-go/clients/c2c/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceC2CClient struct {
	RestApi *BinanceC2CRestApi.RestAPIClient
}

type Option func(*BinanceC2CClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceC2CClient) {
		c.RestApi = BinanceC2CRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceC2CClient(opts ...Option) *BinanceC2CClient {
	client := &BinanceC2CClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
