package copytrading

import (
	BinanceCopyTradingRestApi "github.com/binance/binance-connector-go/clients/copytrading/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceCopyTradingClient struct {
	RestApi *BinanceCopyTradingRestApi.RestAPIClient
}

type Option func(*BinanceCopyTradingClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceCopyTradingClient) {
		c.RestApi = BinanceCopyTradingRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceCopyTradingClient(opts ...Option) *BinanceCopyTradingClient {
	client := &BinanceCopyTradingClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
