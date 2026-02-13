package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	QueryManagedSubAccountTransferLogMasterAccountInvestor()
}

func QueryManagedSubAccountTransferLogMasterAccountInvestor() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SubAccountRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSubAccountClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor(context.Background()).Email("sub-account-email@email.com").StartTime(1623319461670).EndTime(1641782889000).Page(789).Limit(789).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
