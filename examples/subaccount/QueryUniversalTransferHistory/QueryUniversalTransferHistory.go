package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	QueryUniversalTransferHistory()
}

func QueryUniversalTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Universal Transfer History (For Master Account) - /sapi/v1/asset/universalTransfer
	queryUniversalTransferHistory, err := client.NewQueryUniversalTransferHistoryService().FromEmail("from@email.com").
		ToEmail("to@email.com").ClientTranId("123123").StartTime(1234567891011).EndTime(1234567891011).
		Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(queryUniversalTransferHistory))
}
