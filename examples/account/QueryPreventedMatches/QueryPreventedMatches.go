package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryPreventedMatches()
}

func QueryPreventedMatches() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Prevented Matches (USER_DATA) - GET /api/v3/preventedMatches
	getQueryPreventedMatchesService, err := client.NewGetQueryPreventedMatchesService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getQueryPreventedMatchesService))
}
