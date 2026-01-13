package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	MarkPrice()
}

func MarkPrice() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.DerivativesTradingOptionsWebsocketStreamsProdUrl),
	)

	wsClient := client.NewBinanceDerivativesTradingOptionsClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.MarketAPI.MarkPrice().Underlying("btcusdt").Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.MarkPriceResponse) {
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
