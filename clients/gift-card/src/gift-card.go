package binanceGiftCard

import (
	BinanceGiftCardRestApi "github.com/binance/binance-connector-go/clients/giftcard/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceGiftCardClient struct {
	RestApi *BinanceGiftCardRestApi.RestAPIClient
}

type Option func(*BinanceGiftCardClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceGiftCardClient) {
		c.RestApi = BinanceGiftCardRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceGiftCardClient(opts ...Option) *BinanceGiftCardClient {
	client := &BinanceGiftCardClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
