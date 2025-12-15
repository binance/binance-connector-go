package binanceCryptoLoan

import (
	BinanceCryptoLoanRestApi "github.com/binance/binance-connector-go/clients/cryptoloan/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceCryptoLoanClient struct {
	RestApi *BinanceCryptoLoanRestApi.RestAPIClient
}

type Option func(*BinanceCryptoLoanClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceCryptoLoanClient) {
		c.RestApi = BinanceCryptoLoanRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceCryptoLoanClient(opts ...Option) *BinanceCryptoLoanClient {
	client := &BinanceCryptoLoanClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
