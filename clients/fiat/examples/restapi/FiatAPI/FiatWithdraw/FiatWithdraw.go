package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/fiat/src"
	"github.com/binance/binance-connector-go/clients/fiat/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	FiatWithdraw()
}

func FiatWithdraw() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.FiatRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceFiatClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FiatAPI.FiatWithdraw(context.Background()).Currency("currency_example").ApiPaymentMethod("apiPaymentMethod_example").Amount(789).AccountInfo(models.AccountInfo{}).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
