package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	CloseUserDataStream()
}

func CloseUserDataStream() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingCoinFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.UserDataStreamsAPI.CloseUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
