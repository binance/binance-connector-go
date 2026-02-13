package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	QueryUserNegativeBalanceAutoExchangeRecord()
}

func QueryUserNegativeBalanceAutoExchangeRecord() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingPortfolioMarginRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.QueryUserNegativeBalanceAutoExchangeRecord(context.Background()).StartTime(1623319461670).EndTime(1641782889000).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
