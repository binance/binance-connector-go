package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsAllMarketTickersExample()
}

func WsAllMarketTickersExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsAllMarketTickersHandler := func(event binance_connector.WsAllMarketTickersStatEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAllMarketTickersStatServe(wsAllMarketTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
