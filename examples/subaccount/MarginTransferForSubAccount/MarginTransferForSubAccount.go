package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	MarginTransferForSubAccount()
}

func MarginTransferForSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Margin Transfer for Sub-account (For Master Account) - /sapi/v1/sub-account/margin/transfer
	marginTransferForSubAccount, err := client.NewMarginTransferForSubAccountService().Email("from@email.com").Asset("BTC").
		Amount(0.01).TransferType(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginTransferForSubAccount))
}
