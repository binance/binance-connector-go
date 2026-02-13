package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	EditIpForSpecialKey()
}

func EditIpForSpecialKey() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Ip("ip_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
