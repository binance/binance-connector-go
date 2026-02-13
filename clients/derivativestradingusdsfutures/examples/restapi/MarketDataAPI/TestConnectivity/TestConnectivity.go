package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	TestConnectivity()
}

func TestConnectivity() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingUsdsFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
