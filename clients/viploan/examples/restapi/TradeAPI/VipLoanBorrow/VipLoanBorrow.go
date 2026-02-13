package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	VipLoanBorrow()
}

func VipLoanBorrow() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.VipLoanRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceVipLoanClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).LoanAccountId(1).LoanCoin("loanCoin_example").LoanAmount(1.0).CollateralAccountId("1").CollateralCoin("collateralCoin_example").IsFlexibleRate(true).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
