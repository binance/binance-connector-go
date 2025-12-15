package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	MarginAccountNewOtoco()
}

func MarginAccountNewOtoco() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Symbol("symbol_example").WorkingType("workingType_example").WorkingSide("workingSide_example").WorkingPrice(1.0).WorkingQuantity(1.0).PendingSide("pendingSide_example").PendingQuantity(1.0).PendingAboveType("pendingAboveType_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
