package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsAllMarketMiniTickers()
}

func WsAllMarketMiniTickers() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsAllMarketMiniTickersHandler := func(event binance_connector.WsAllMarketMiniTickersStatEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsAllMarketMiniTickersStatServe(wsAllMarketMiniTickersHandler, errHandler)
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
