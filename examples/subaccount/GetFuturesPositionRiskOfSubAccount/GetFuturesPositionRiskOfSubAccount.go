package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	GetFuturesPositionRiskOfSubAccount()
}

func GetFuturesPositionRiskOfSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Get Futures Position-Risk of Sub-account (For Master Account) - /sapi/v1/sub-account/futures/positionRisk
	getFuturesPositionRiskOfSubAccount, err := client.NewGetFuturesPositionRiskOfSubAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFuturesPositionRiskOfSubAccount))
}
