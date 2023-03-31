package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	SubAccountSpotAssetTransferHistory()
}

func SubAccountSpotAssetTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Spot Asset Transfer History (For Master Account) - /sapi/v1/sub-account/sub/transfer/history
	subaccountSpotAssetTransferHistory, err := client.NewQuerySubAccountSpotAssetTransferHistoryService().FromEmail("from@email.com").
		ToEmail("to@email.com").StartTime(1234567891011).EndTime(1).Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(subaccountSpotAssetTransferHistory))
}
