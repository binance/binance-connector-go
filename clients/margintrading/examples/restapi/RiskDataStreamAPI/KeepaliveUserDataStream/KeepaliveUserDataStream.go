package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	KeepaliveUserDataStream()
}

func KeepaliveUserDataStream() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.RiskDataStreamAPI.KeepaliveUserDataStream(context.Background()).ListenKey("listenKey_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
