package main

import (
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	OrderBook()
}

func OrderBook() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.DerivativesTradingUsdsFuturesWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
	)

	wsClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.MarketDataAPI.OrderBook().Symbol("symbol_example").ExecuteAsync()
	if err != nil {
		log.Printf("Error executing OrderBook request: %v\n", err)
		return
	}

	select {
	case resp := <-responseChan:
		result, _ := json.MarshalIndent(resp.Typed, "", "  ")
		log.Printf("Result: %s\n", result)
	case err := <-errorChan:
		log.Printf("Error: %v\n", err)
	}

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
