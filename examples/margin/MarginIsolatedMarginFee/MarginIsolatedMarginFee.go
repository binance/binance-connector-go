package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginIsolatedMarginFee()
}

func MarginIsolatedMarginFee() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedMarginFeeService - /sapi/v1/margin/isolatedMarginData
	marginIsolatedMarginFee, err := client.NewMarginIsolatedMarginFeeService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedMarginFee))
}
