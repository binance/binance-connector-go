package algo

import (
	BinanceAlgoRestApi "github.com/binance/binance-connector-go/clients/algo/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceAlgoClient struct {
	RestApi *BinanceAlgoRestApi.RestAPIClient
}

type Option func(*BinanceAlgoClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceAlgoClient) {
		c.RestApi = BinanceAlgoRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceAlgoClient(opts ...Option) *BinanceAlgoClient {
	client := &BinanceAlgoClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
