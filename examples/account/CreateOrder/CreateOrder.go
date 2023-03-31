package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	NewOrder()
}

func NewOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Binance New Order endpoint - POST /api/v3/order
	newOrder, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side("BUY").Type("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(newOrder))
}
