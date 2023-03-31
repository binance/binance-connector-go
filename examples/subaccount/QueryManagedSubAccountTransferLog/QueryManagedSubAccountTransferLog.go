package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryManagedSubAccountTransferLog()
}

func QueryManagedSubAccountTransferLog() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Managed Sub Account Transfer Log (Investor) (USER_DATA)
	queryManagedSubAccountTransferLog, err := client.NewQueryManagedSubAccountTransferLogService().Email("email@email.com").
		StartTime(123123).EndTime(123123).Page(1).Limit(10).Transfers("").TransferFunctionAccountType("").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryManagedSubAccountTransferLog))
}
