package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/giftcard/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	RedeemABinanceGiftCard()
}

func RedeemABinanceGiftCard() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.GiftCardRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceGiftCardClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.RedeemABinanceGiftCard(context.Background()).Code("code_example").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
