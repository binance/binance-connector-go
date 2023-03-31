package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountNewOrder()
}

func MarginAccountNewOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountNewOrderService - /sapi/v1/margin/order
	marginAccountNewOrder, err := client.NewMarginAccountNewOrderService().Symbol("BTCUSDT").
		Side("BUY").OrderType("MARKET").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOrder))
}
