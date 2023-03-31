package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	SubAccountFuturesAssetTransfer()
}

func SubAccountFuturesAssetTransfer() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Sub-account Futures Asset Transfer (For Master Account) - /sapi/v1/sub-account/futures/internalTransfer
	subaccountFuturesAssetTransfer, err := client.NewSubAccountFuturesAssetTransferService().FromEmail("from@email.com").
		ToEmail("to@email.com").FuturesType(1).Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountFuturesAssetTransfer))
}
