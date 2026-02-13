package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	KeepaliveUserDataStream()
}

func KeepaliveUserDataStream() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingPortfolioMarginRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.UserDataStreamsAPI.KeepaliveUserDataStream(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
