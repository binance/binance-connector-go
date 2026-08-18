package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetOtcReservedBalances()
}

func GetOtcReservedBalances() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.W3wPredictionRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceW3wPredictionClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.OtcAPI.GetOtcReservedBalances(context.Background()).Assets([]models.GetOtcReservedBalancesAssetsParameterInner{}).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
