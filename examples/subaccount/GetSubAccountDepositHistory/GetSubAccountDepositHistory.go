package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetSubAccountDepositHistory()
}

func GetSubAccountDepositHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Sub-account Deposit History (For Master Account) - /sapi/v1/capital/deposit/subHisrec
	getSubAccountDepositHistory, err := client.NewGetSubAccountDepositHistoryService().Email("from@email.com").
		Coin("BTC").Status(1).StartTime(1234567891011).EndTime(1234567891011).Limit(10).Offset(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getSubAccountDepositHistory))
}
