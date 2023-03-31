package main

import (
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsUserData()
}

func WsUserData() {
	wsUserDataHandler := func(event *binance_connector.WsUserDataEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := binance_connector.WsUserDataServe("YourListenKey", wsUserDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
