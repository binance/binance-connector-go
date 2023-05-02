package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CancelReplace()
}

func CancelReplace() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Cancel an Existing Order and Send a New Order (TRADE) - POST /api/v3/order/cancelReplace
	cancelReplace, err := client.NewCancelReplaceService().
		Symbol("BTCUSDT").Side("BUY").OrderType("LIMIT").CancelReplaceMode("STOP_ON_FAILURE").CancelOrderId(13341128).
		TimeInForce("GTC").Quantity(0.001).Price(20000.0).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(cancelReplace))
}
