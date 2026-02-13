package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetOrderDetail()
}

func GetOrderDetail() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.FiatRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceFiatClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FiatAPI.GetOrderDetail(context.Background()).OrderNo("orderNo_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
