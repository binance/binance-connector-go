package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsAllMarketTickersExample()
}

func WsAllMarketTickersExample() {
	wsAllMarketTickersHandler := func(event binance_connector.WsAllMarketsStatEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsAllMarketsStatServe(wsAllMarketTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
