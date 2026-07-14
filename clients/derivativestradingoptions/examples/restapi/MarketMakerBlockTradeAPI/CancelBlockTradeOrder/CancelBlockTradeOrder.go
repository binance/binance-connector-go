package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	CancelBlockTradeOrder()
}

func CancelBlockTradeOrder() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingOptionsRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceDerivativesTradingOptionsClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).BlockOrderMatchingKey("7d046e6e-a429-4335-ab9d-6a681febcde5").Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
