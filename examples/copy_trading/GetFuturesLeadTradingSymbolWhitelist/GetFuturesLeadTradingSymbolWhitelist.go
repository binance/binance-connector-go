package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetFuturesLeadTradingSymbolWhitelist()
}

func GetFuturesLeadTradingSymbolWhitelist() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Futures Lead Trading Symbol Whitelist - /sapi/v1/copyTrading/futures/leadSymbol
	getFuturesLeadTradingSymbolWhitelist, err := client.NewGetFuturesLeadTradingSymbolWhitelistService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFuturesLeadTradingSymbolWhitelist))
}
