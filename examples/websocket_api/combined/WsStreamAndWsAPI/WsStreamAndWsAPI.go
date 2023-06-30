package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsStreamAndWsAPI()
}

func WsStreamAndWsAPI() {
	// Websocket Stream
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsTradeHandler := func(event *binance_connector.WsTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsTradeServe("LTCBTC", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Websocket API
	client := binance_connector.NewWebsocketAPIClient("api_key", "secret_key", "wss://testnet.binance.vision/ws-api/v3")
	err2 := client.Connect()
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}

	response, err := client.NewPlaceNewOrderService().Symbol("BTCUSDT").Side("BUY").OrderType("MARKET").Quantity(0.01).
		Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	<-doneCh
	client.WaitForCloseSignal()
}
