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
	CreateOtcBlocktrade()
}

func CreateOtcBlocktrade() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.W3wPredictionRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceW3wPredictionClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.OtcAPI.CreateOtcBlocktrade(context.Background()).MarketId("123").TokenId("71321045679252212594626385532706912750332728571942532289631379312455583992563").Side(models.GetQuoteSideParameterBuy).MakerAmount("600000000000000000000").TakerAmount("1000000000000000000000").PricePerShare("0.65").Expiration(1790000000).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
