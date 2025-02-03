package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	TickerPrice()
}

func TickerPrice() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "",binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// TickerPrice
	tickerPrice, err := client.NewTickerPriceService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(tickerPrice))
}
