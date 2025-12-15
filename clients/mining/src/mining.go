package binanceMining

import (
	BinanceMiningRestApi "github.com/binance/binance-connector-go/clients/mining/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceMiningClient struct {
	RestApi *BinanceMiningRestApi.RestAPIClient
}

type Option func(*BinanceMiningClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceMiningClient) {
		c.RestApi = BinanceMiningRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceMiningClient(opts ...Option) *BinanceMiningClient {
	client := &BinanceMiningClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
