package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	AddIpRestrictionForSubAccountApiKey()
}

func AddIpRestrictionForSubAccountApiKey() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SubAccountRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSubAccountClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").Status(1).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
