package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	Ping()
}

func Ping() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// NewPingService
	ping := client.NewPingService().Do(context.Background())
	if ping == nil {
		fmt.Println("Success")
		return
	}
	fmt.Println(binance_connector.PrettyPrint(ping))
}
