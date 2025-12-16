package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	Ping()
}

func Ping() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceSpotClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.GeneralAPI.Ping(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
