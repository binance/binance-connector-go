package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	ServerTime()
}

func ServerTime() {

	client := binance_connector.NewClient("", "")

	// set to debug mode
	client.Debug = true

	// NewServerTimeService
	serverTime, err := client.NewServerTimeService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(serverTime))
}
