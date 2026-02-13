package rebate

import (
	BinanceRebateRestApi "github.com/binance/binance-connector-go/clients/rebate/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceRebateClient struct {
	RestApi *BinanceRebateRestApi.RestAPIClient
}

type Option func(*BinanceRebateClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceRebateClient) {
		c.RestApi = BinanceRebateRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceRebateClient(opts ...Option) *BinanceRebateClient {
	client := &BinanceRebateClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
