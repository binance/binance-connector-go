package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginBnbBurnStatus()
}

func MarginBnbBurnStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginBnbBurnStatusService - /sapi/v1/bnbBurn
	marginBnbBurnStatus, err := client.NewMarginBnbBurnStatusService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginBnbBurnStatus))
}
