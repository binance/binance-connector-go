package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSubAccountStatus()
}

func GetSubAccountStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Sub-account's Status on Margin/Futures (For Master Account) - /sapi/v1/sub-account/status
	getSubAccountStatus, err := client.NewGetSubAccountStatusService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSubAccountStatus))
}
