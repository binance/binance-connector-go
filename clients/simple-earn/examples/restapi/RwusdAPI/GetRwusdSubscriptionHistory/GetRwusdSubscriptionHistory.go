package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetRwusdSubscriptionHistory()
}

func GetRwusdSubscriptionHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SimpleEarnRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSimpleEarnClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.RwusdAPI.GetRwusdSubscriptionHistory(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
