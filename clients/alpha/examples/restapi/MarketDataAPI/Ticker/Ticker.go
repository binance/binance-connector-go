package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	Ticker()
}

func Ticker() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlphaRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.Ticker(context.Background()).Symbol("symbol_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
