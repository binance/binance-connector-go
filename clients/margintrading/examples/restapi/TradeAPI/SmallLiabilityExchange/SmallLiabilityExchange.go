package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	SmallLiabilityExchange()
}

func SmallLiabilityExchange() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).AssetNames([]string{}).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
