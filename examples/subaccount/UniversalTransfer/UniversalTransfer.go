package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	UniversalTransfer()
}

func UniversalTransfer() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Universal Transfer (For Master Account) - /sapi/v1/asset/universalTransfer
	universalTransfer, err := client.NewUniversalTransferService().FromEmail("from@email.com").ToEmail("to@email.com").
		FromAccountType("SPOT").ToAccountType("SPOT").ClientTranId("123123").Symbol("BTC").Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(universalTransfer))
}
