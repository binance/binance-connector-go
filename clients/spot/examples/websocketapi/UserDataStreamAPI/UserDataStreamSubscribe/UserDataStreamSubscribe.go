package main

import (
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	UserDataStreamSubscribe()
}

func UserDataStreamSubscribe() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
	)

	wsClient := client.NewBinanceSpotClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	resp, stream, errorChan, err := wsClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribe().Execute()
	if err != nil || errorChan != nil {
		log.Printf("Error executing UserDataStreamSubscribe request: %v, %v\n", err, errorChan)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	stream.On("message", func(message models.UserDataStreamEventsResponse) {
		formattedMessage, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received event: %s", string(formattedMessage))
	})

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
