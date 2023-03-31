package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsTradeExample()
}

func WsTradeExample() {
	wsTradeHandler := func(event *binance_connector.WsTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsTradeServe("LTCBTC", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
