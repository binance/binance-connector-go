package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "",binance_connector.SIGNATURE_HMAC_SHA256, baseURL)

	// ExchangeInfo
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(exchangeInfo))
}
