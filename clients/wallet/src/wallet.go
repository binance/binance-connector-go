package binanceWallet

import (
	BinanceWalletRestApi "github.com/binance/binance-connector-go/clients/wallet/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceWalletClient struct {
	RestApi *BinanceWalletRestApi.RestAPIClient
}

type Option func(*BinanceWalletClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceWalletClient) {
		c.RestApi = BinanceWalletRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceWalletClient(opts ...Option) *BinanceWalletClient {
	client := &BinanceWalletClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
