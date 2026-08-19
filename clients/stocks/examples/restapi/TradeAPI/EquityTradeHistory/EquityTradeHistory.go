package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	EquityTradeHistory()
}

func EquityTradeHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.StocksRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceStocksClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.EquityTradeHistory(context.Background()).StartTime(1735800000000).EndTime(1735900000000).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
