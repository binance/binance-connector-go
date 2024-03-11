package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	// WsDepthHandlerExample()
	WsDepthLazyHandlerExample()
}

func WsDepthHandlerExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneCh
}

func WsDepthLazyHandlerExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, subCh, err := websocketStreamClient.WsDepthServeSubscribeLazy(wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	stopCh <- struct{}{}
	// }()

	go func() {
		subCh <- binance_connector.Payload{Id: "1234-5678-9101", Method: "SUBSCRIBE", Params: []string{"btcusdt@depth"}}
		fmt.Println("Sent subscribe event btcusdt")

		time.Sleep(2 * time.Second)

		subCh <- binance_connector.Payload{Id: "999-5678-765", Method: "SUBSCRIBE", Params: []string{"ltcbtc@depth"}}
		fmt.Println("Sent subscribe event ltcbtc")
	}()
	// remove this if you do not want to be blocked here
	<-doneCh
}
