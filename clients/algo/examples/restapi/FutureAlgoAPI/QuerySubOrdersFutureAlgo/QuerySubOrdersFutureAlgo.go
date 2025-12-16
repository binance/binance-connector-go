package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	QuerySubOrdersFutureAlgo()
}

func QuerySubOrdersFutureAlgo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlgoRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceAlgoClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FutureAlgoAPI.QuerySubOrdersFutureAlgo(context.Background()).AlgoId(1).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
