package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	IndividualSymbolBookTickerStreams()
}

func IndividualSymbolBookTickerStreams() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.DerivativesTradingUsdsFuturesWebsocketStreamsProdUrl),
	)

	wsClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams().Symbol("btcusdt").Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.IndividualSymbolBookTickerStreamsResponse) {
		b, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received message: %s\n", string(b))
	})

	log.Println("Subscribed. Waiting 10 seconds...")
	time.Sleep(10 * time.Second)

	log.Println("Unsubscribing from stream...")
	handler.Unsubscribe()

	log.Println("Closing WebSocket connection...")
	err = wsClient.WebsocketStreams.CloseWebSocketStreamConnection()
	if err != nil {
		log.Fatalf("Error closing WebSocket connection: %v", err)
	}
}
