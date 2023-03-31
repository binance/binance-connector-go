package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryMarginPriceIndex()
}

func QueryMarginPriceIndex() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// QueryMarginPriceIndexService - /sapi/v1/margin/priceIndex
	queryMarginPriceIndex, err := client.NewQueryMarginPriceIndexService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryMarginPriceIndex))
}
