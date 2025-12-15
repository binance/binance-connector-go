package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	CompressedAggregateTradesList()
}

func CompressedAggregateTradesList() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingCoinFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Symbol("symbol_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
