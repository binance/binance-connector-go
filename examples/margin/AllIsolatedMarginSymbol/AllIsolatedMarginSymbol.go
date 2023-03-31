package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AllIsolatedMarginSymbol()
}

func AllIsolatedMarginSymbol() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// AllIsolatedMarginSymbolService - /sapi/v1/margin/isolated/allPairs
	allIsolatedMarginSymbol, err := client.NewAllIsolatedMarginSymbolService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(allIsolatedMarginSymbol))
}
