package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginIsolatedAccountEnable()
}

func MarginIsolatedAccountEnable() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountEnableService - /sapi/v1/margin/isolated/account
	marginIsolatedAccountEnable, err := client.NewMarginIsolatedAccountEnableService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedAccountEnable))
}
