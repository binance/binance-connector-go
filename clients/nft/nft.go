package nft

import (
	BinanceNFTRestApi "github.com/binance/binance-connector-go/clients/nft/src/restapi"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceNFTClient struct {
	RestApi *BinanceNFTRestApi.RestAPIClient
}

type Option func(*BinanceNFTClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceNFTClient) {
		c.RestApi = BinanceNFTRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceNFTClient(opts ...Option) *BinanceNFTClient {
	client := &BinanceNFTClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
