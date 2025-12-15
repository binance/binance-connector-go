package binancePay

import (
	BinancePayRestApi "github.com/binance/binance-connector-go/clients/pay/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinancePayClient struct {
	RestApi *BinancePayRestApi.RestAPIClient
}

type Option func(*BinancePayClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinancePayClient) {
		c.RestApi = BinancePayRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinancePayClient(opts ...Option) *BinancePayClient {
	client := &BinancePayClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
