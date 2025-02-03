package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountAdjustCrossMaxLeverage()
}

func MarginAccountAdjustCrossMaxLeverage() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// MarginAccountAdjustCrossMaxLeverageService - /sapi/v1/margin/max-leverage
	marginAccountAdjustCrossMaxLeverage, err := client.NewMarginAccountAdjustCrossMaxLeverageService().
		MaxLeverage(5).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountAdjustCrossMaxLeverage))
}
