package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsBookTickerExample()
}

func WsBookTickerExample() {
	wsBookTickerHandler := func(event *binance_connector.WsBookTickerEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsBookTickerServe("LTCBTC", wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
