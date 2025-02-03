package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	TickerBookTicker()
}

func TickerBookTicker() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "",binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// TickerBookTicker
	tickerBookTicker, err := client.NewTickerBookTickerService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(tickerBookTicker))
}
