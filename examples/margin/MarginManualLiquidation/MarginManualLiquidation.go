package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginManualLiquidation()
}

func MarginManualLiquidation() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// MarginManualLiquidationService - /sapi/v1/margin/manual-liquidation
	marginManualLiquidation, err := client.NewMarginManualLiquidationService().
		MarginType("MARGIN").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginManualLiquidation))
}
