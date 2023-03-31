package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginToggleBnbBurn()
}

func MarginToggleBnbBurn() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginToggleBnbBurnService - /sapi/v1/bnbBurn
	marginToggleBnbBurn, err := client.NewMarginToggleBnbBurnService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginToggleBnbBurn))
}
