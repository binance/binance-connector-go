package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/clients/alpha/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	Klines()
}

func Klines() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlphaRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.Klines(context.Background()).Symbol("ALPHA_175USDT").Interval(models.KlinesIntervalParameterInterval1s).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
