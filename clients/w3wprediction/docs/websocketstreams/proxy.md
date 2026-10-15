# Proxy

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	PredictionOrderBookByMarketId()
}

func PredictionOrderBookByMarketId() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.W3wPredictionWebsocketStreamsProdUrl),
		common.WithWsStreamsApiKey("your-api-key"),
		common.WithWsStreamsApiSecret("your-api-secret"),
		common.WithWsStreamsProxy(common.ProxyConfig{
			Host:     "127.0.0.1",
			Port:     8080,
			Protocol: "http",
			Auth: struct {
				Username string
				Password string
			}{
				Username: "proxy-user",
				Password: "proxy-password",
			},
		}),
	)

	wsClient := client.NewBinanceW3wPredictionClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.PredictionOrderBookByMarketId().MarketId("8859231").Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.PredictionOrderBookByMarketIdResponse) {
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
```
