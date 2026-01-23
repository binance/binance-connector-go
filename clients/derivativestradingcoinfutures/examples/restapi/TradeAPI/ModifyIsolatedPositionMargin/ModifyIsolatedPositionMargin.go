package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	ModifyIsolatedPositionMargin()
}

func ModifyIsolatedPositionMargin() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.DerivativesTradingCoinFuturesRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.ModifyIsolatedPositionMargin(context.Background()).Symbol("symbol_example").Amount(1.0).Type(models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
