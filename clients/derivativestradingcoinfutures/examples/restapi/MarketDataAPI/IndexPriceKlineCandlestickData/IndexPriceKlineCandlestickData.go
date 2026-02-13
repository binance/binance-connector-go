package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	IndexPriceKlineCandlestickData()
}

func IndexPriceKlineCandlestickData() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingCoinFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
	)
	apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Pair("pair_example").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
