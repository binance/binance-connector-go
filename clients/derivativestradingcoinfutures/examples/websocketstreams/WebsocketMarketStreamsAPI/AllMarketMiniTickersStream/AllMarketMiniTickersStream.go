package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	AllMarketMiniTickersStream()
}

func AllMarketMiniTickersStream() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.DerivativesTradingCoinFuturesWebsocketStreamsProdUrl),
	)

	wsClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketMiniTickersStream().Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.AllMarketMiniTickersStreamResponse) {
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
