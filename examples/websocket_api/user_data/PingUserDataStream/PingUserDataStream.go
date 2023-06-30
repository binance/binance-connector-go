package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	PingUserDataStream()
}

func PingUserDataStream() {
	client := binance_connector.NewWebsocketAPIClient("API_KEY", "")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewPingUserDataStreamService().ListenKey("LISTEN_KEY").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response.Response))

	client.WaitForCloseSignal()
}
