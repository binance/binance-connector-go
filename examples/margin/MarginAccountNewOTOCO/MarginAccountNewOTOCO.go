package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountNewOTOCO()
}

func MarginAccountNewOTOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountNewOTOCOService - /sapi/v1/margin/order/otoco
	marginAccountNewOTOCO, err := client.NewMarginAccountNewOTOCOService().
		Symbol("BNBUSDT").WorkingType("LIMIT").WorkingSide("BUY").WorkingPrice(600).WorkingQuantity(1).
		PendingSide("BUY").PendingQuantity(1).PendingAboveType("LIMIT_MAKER").WorkingTimeInForce("GTC").
		PendingAbovePrice(605).PendingBelowType("LIMIT_MAKER").PendingBelowPrice(595).WorkingIcebergQty(0.1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOTOCO))
}
