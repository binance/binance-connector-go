package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	CloseUserDataStream()
}

func CloseUserDataStream() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingOptionsRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingOptionsClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.UserDataStreamsAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
