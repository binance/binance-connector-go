package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountQueryOpenOCO()
}

func MarginAccountQueryOpenOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryOpenOCOService - /sapi/v1/margin/openOrderList
	marginAccountQueryOpenOCO, err := client.NewMarginAccountQueryOpenOCOService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountQueryOpenOCO))
}
