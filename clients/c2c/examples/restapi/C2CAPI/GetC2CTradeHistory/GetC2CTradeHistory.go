package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/c2c/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetC2CTradeHistory()
}

func GetC2CTradeHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.C2CRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceC2CClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.C2CAPI.GetC2CTradeHistory(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
