package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	FlexibleLoanAdjustLtv()
}

func FlexibleLoanAdjustLtv() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.CryptoLoanRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceCryptoLoanClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanAdjustLtv(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").AdjustmentAmount(1.0).Direction("direction_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
