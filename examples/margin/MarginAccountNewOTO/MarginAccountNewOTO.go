package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountNewOTO()
}

func MarginAccountNewOTO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// MarginAccountNewOTOService - /sapi/v1/margin/order/oto
	marginAccountNewOTO, err := client.NewMarginAccountNewOTOService().
		Symbol("BNBUSDT").WorkingType("LIMIT").WorkingSide("BUY").WorkingPrice(600).WorkingQuantity(1).
		PendingType("LIMIT").PendingSide("BUY").PendingQuantity(1).WorkingTimeInForce("GTC").
		PendingPrice(595).PendingTimeInForce("GTC").WorkingIcebergQty(0.1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOTO))
}
