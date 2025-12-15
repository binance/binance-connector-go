package binanceStaking

import (
	BinanceStakingRestApi "github.com/binance/binance-connector-go/clients/staking/src/restapi"
	"github.com/binance/binance-connector-go/common/common"
)

type BinanceStakingClient struct {
	RestApi *BinanceStakingRestApi.RestAPIClient
}

type Option func(*BinanceStakingClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceStakingClient) {
		c.RestApi = BinanceStakingRestApi.NewRestAPIClient(cfg)
	}
}

func NewBinanceStakingClient(opts ...Option) *BinanceStakingClient {
	client := &BinanceStakingClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
