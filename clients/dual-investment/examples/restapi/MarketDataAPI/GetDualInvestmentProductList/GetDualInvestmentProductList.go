package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/dualinvestment/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetDualInvestmentProductList()
}

func GetDualInvestmentProductList() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DualInvestmentRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceDualInvestmentClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.GetDualInvestmentProductList(context.Background()).OptionType("optionType_example").ExercisedCoin("exercisedCoin_example").InvestCoin("investCoin_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
