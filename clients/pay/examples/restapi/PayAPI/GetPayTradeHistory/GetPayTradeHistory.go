package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/pay/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetPayTradeHistory()
}

func GetPayTradeHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.PayRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinancePayClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
