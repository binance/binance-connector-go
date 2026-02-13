package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	OrderListOpo()
}

func OrderListOpo() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSpotClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.OrderListOpo(context.Background()).Symbol("BNBUSDT").WorkingType(models.OrderListOpoWorkingTypeParameterLimit).WorkingSide(models.NewOrderSideParameterBuy).WorkingPrice(1.0).WorkingQuantity(1.0).PendingType(models.OrderListOpoPendingTypeParameterLimit).PendingSide(models.NewOrderSideParameterBuy).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
