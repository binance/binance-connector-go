package dualinvestment

import (
	BinanceDualInvestmentRestApi "github.com/binance/binance-connector-go/clients/dualinvestment/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceDualInvestmentClient struct {
	RestApi *BinanceDualInvestmentRestApi.RestAPIClient
}

type Option func(*BinanceDualInvestmentClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceDualInvestmentClient) {
		c.RestApi = BinanceDualInvestmentRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceDualInvestmentClient(opts ...Option) *BinanceDualInvestmentClient {
	client := &BinanceDualInvestmentClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
