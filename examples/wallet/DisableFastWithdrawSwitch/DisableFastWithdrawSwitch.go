package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	DisableFastWithdrawSwitch()
}

func DisableFastWithdrawSwitch() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// DisableFastWithdrawSwitchService  - /sapi/v1/account/disableFastWithdrawSwitch
	disableFastWithdrawSwitch, err := client.NewDisableFastWithdrawSwitchService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(disableFastWithdrawSwitch))
}
