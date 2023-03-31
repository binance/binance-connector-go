package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsKlineExample()
}

func WsKlineExample() {
	wsKlineHandler := func(event *binance_connector.WsKlineEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsKlineServe("LTCBTC", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
