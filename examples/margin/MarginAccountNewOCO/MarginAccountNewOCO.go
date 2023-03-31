package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginAccountNewOCO()
}

func MarginAccountNewOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountNewOCOService - /sapi/v1/margin/order/oco
	marginAccountNewOCO, err := client.NewMarginAccountNewOCOService().Symbol("BTCUSDT").
		Side("BUY").Quantity(0.01).Price(20000).StopPrice(18000).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountNewOCO))
}
