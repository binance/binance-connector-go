package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	SubAccountList()
}

func SubAccountList() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account List (For Master Account) - /sapi/v1/sub-account/list
	subaccountList, err := client.NewQuerySubAccountListService().Email("test@email.com").
		IsFreeze("").Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountList))
}
