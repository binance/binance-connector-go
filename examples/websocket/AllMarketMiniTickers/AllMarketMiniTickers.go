package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsAllMarketMiniTickers()
}

func WsAllMarketMiniTickers() {
	wsAllMarketMiniTickersHandler := func(event binance_connector.WsAllMiniMarketsStatEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsAllMiniMarketsStatServe(wsAllMarketMiniTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
