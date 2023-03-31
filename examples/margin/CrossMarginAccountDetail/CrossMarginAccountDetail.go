package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	CrossMarginAccountDetail()
}

func CrossMarginAccountDetail() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// CrossMarginAccountDetailService - /sapi/v1/margin/account
	crossMarginAccountDetail, err := client.NewCrossMarginAccountDetailService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(crossMarginAccountDetail))
}
