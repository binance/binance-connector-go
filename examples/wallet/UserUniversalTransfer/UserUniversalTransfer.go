package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	UserUniversalTransfer()
}

func UserUniversalTransfer() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserUniversalTransferService - /sapi/v1/asset/transfer
	userUniversalTransfer, err := client.NewUserUniversalTransferService().
		TransferType("MAIN_UMFUTURE").Asset("USDT").Amount(20.50).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userUniversalTransfer))
}
