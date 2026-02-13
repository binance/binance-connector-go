package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	EarningsList()
}

func EarningsList() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MiningRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMiningClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MiningAPI.EarningsList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
