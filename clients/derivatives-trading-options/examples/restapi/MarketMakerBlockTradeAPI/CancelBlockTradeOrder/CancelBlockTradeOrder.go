package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src"
	"github.com/binance/binance-connector-go/common/common"
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
	_, err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).BlockOrderMatchingKey("blockOrderMatchingKey_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
