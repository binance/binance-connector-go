package simpleearn

import (
	BinanceSimpleEarnRestApi "github.com/binance/binance-connector-go/clients/simpleearn/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceSimpleEarnClient struct {
	RestApi *BinanceSimpleEarnRestApi.RestAPIClient
}

type Option func(*BinanceSimpleEarnClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceSimpleEarnClient) {
		c.RestApi = BinanceSimpleEarnRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceSimpleEarnClient(opts ...Option) *BinanceSimpleEarnClient {
	client := &BinanceSimpleEarnClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
