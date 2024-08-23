package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	AggTradesExample()
}

func AggTradesExample() {
	// Initialise Websocket Client with Testnet base URL and false for "isCombined" parameter
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")

	wsAggTradeHandler := func(event *binance_connector.WsAggTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsAggTradeServe("BTCUSDT", wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(10 * time.Second)
		stopCh <- struct{}{}
	}()
	<-doneCh
}
