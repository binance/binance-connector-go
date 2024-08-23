package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	WsMarketMiniTickers()
}

func WsMarketMiniTickers() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsMarketMiniTickersHandler := func(event binance_connector.WsMarketMiniTickerStatEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsMarketMiniTickersStatServe("BNBBTC", wsMarketMiniTickersHandler, errHandler)
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
