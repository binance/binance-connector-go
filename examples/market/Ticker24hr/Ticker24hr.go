package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	Ticker24hr()
}

func Ticker24hr() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// Ticker24hr
	ticker24hr, err := client.NewTicker24hrService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(ticker24hr))
}
