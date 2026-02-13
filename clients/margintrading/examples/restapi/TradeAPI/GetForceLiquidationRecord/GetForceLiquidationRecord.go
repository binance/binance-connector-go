package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetForceLiquidationRecord()
}

func GetForceLiquidationRecord() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.GetForceLiquidationRecord(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
