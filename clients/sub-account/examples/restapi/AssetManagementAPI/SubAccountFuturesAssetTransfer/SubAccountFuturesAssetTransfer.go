package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	SubAccountFuturesAssetTransfer()
}

func SubAccountFuturesAssetTransfer() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SubAccountRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSubAccountClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountFuturesAssetTransfer(context.Background()).FromEmail("fromEmail_example").ToEmail("toEmail_example").FuturesType(789).Asset("asset_example").Amount(1.0).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
