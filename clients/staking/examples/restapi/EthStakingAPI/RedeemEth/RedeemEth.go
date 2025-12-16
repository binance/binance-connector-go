package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	RedeemEth()
}

func RedeemEth() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.StakingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceStakingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.EthStakingAPI.RedeemEth(context.Background()).Amount(1.0).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
