package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/copytrading/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetFuturesLeadTradingSymbolWhitelist()
}

func GetFuturesLeadTradingSymbolWhitelist() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.CopyTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceCopyTradingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTradingSymbolWhitelist(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
