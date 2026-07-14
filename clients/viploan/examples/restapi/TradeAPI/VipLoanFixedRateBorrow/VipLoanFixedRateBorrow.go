package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	VipLoanFixedRateBorrow()
}

func VipLoanFixedRateBorrow() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.VipLoanRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceVipLoanClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.VipLoanFixedRateBorrow(context.Background()).SupplyRequest("1212:0.12:100;3434:0.13:50").BorrowCoin("BUSD").LoanTerm(30).BorrowUid(12345678).CollateralCoin("BNB,ETH,BTC").CollateralAccountId("12345,67890,13579").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
