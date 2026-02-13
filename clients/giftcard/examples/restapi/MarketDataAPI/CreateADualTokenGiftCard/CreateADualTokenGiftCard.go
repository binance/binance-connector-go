package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/giftcard"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	CreateADualTokenGiftCard()
}

func CreateADualTokenGiftCard() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.GiftCardRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceGiftCardClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.MarketDataAPI.CreateADualTokenGiftCard(context.Background()).BaseToken("baseToken_example").FaceToken("faceToken_example").BaseTokenAmount(1.0).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
