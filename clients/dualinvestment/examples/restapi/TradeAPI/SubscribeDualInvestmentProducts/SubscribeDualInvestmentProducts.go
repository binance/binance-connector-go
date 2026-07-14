package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/dualinvestment"
	"github.com/binance/binance-connector-go/clients/dualinvestment/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	SubscribeDualInvestmentProducts()
}

func SubscribeDualInvestmentProducts() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DualInvestmentRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceDualInvestmentClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.SubscribeDualInvestmentProducts(context.Background()).Id("741590").OrderId("8257205859").DepositAmount(1).AutoCompoundPlan(models.ChangeAutoCompoundStatusAutoCompoundPlanParameterNone).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
