package convert

import (
	BinanceConvertRestApi "github.com/binance/binance-connector-go/clients/convert/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceConvertClient struct {
	RestApi *BinanceConvertRestApi.RestAPIClient
}

type Option func(*BinanceConvertClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceConvertClient) {
		c.RestApi = BinanceConvertRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceConvertClient(opts ...Option) *BinanceConvertClient {
	client := &BinanceConvertClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
