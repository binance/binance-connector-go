package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	MovePositionForSubAccount()
}

func MovePositionForSubAccount() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SubAccountRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSubAccountClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AssetManagementAPI.MovePositionForSubAccount(context.Background()).FromUserEmail("testFrom@google.com").ToUserEmail("testTo@google.com").ProductType(models.MovePositionForSubAccountProductTypeParameterUm).OrderArgs([]models.MovePositionForSubAccountOrderArgsParameterInner{}).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
