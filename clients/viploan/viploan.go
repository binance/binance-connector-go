package viploan

import (
	BinanceVipLoanRestApi "github.com/binance/binance-connector-go/clients/viploan/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceVipLoanClient struct {
	RestApi *BinanceVipLoanRestApi.RestAPIClient
}

type Option func(*BinanceVipLoanClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceVipLoanClient) {
		c.RestApi = BinanceVipLoanRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceVipLoanClient(opts ...Option) *BinanceVipLoanClient {
	client := &BinanceVipLoanClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
