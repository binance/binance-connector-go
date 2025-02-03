package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	Klines()
}

func Klines() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// Klines
	klines, err := client.NewKlinesService().
		Symbol("BTCUSDT").Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(klines))
}
