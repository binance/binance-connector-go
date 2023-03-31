package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	PingUserStream()
}

func PingUserStream() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	ping := client.NewPingUserStream().ListenKey("your_listen_key").Do(context.Background())
	fmt.Println(ping)
}
