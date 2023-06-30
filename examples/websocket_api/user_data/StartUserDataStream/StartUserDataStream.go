package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	StartUserDataStream()
}

func StartUserDataStream() {
	client := binance_connector.NewWebsocketAPIClient("API_KEY", "")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	response, err := client.NewStartUserDataStreamService().Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response.Result.ListenKey))
	client.WaitForCloseSignal()
}
