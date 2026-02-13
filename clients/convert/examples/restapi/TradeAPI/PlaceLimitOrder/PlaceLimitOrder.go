package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	PlaceLimitOrder()
}

func PlaceLimitOrder() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.ConvertRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceConvertClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.PlaceLimitOrder(context.Background()).BaseAsset("baseAsset_example").QuoteAsset("quoteAsset_example").LimitPrice(1.0).Side("BUY").ExpiredType("expiredType_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
