package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountQueryMaxBorrow()
}

func MarginAccountQueryMaxBorrow() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryMaxBorrowService - /sapi/v1/margin/maxBorrowable
	marginAccountQueryMaxBorrow, err := client.NewMarginAccountQueryMaxBorrowService().Asset("USDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountQueryMaxBorrow))
}
