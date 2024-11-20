package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryLiabilityCoinLeverageBracket()
}

func QueryLiabilityCoinLeverageBracket() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// QueryLiabilityCoinLeverageBracketService - /sapi/v1/margin/leverageBracket
	queryLiabilityCoinLeverageBracket, err := client.NewQueryLiabilityCoinLeverageBracketService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryLiabilityCoinLeverageBracket))
}
