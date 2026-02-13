package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	QueryPortfolioMarginAssetIndexPrice()
}

func QueryPortfolioMarginAssetIndexPrice() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingPortfolioMarginProRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingPortfolioMarginProClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.QueryPortfolioMarginAssetIndexPrice(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
