package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetOpenSymbolList()
}

func GetOpenSymbolList() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.WalletRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceWalletClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AssetAPI.GetOpenSymbolList(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
