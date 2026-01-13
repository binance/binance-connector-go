# User Data

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	userData()
}

func userData() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsPrivateKey("Your Private Key"),
	)

	wsClient := client.NewBinanceSpotClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		fmt.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.AuthAPI.SessionLogon().ExecuteAsync()
	if err != nil {
		log.Printf("Error executing session logon request: %v\n", err)
		return
	}

	select {
	case resp := <-responseChan:
		result, _ := json.MarshalIndent(resp.Typed, "", "  ")
		log.Printf("Result: %s\n", result)
	case err := <-errorChan:
		log.Printf("Error: %v\n", err)
	}

	resp1, stream, _, err := wsClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribe().Execute()
	if err != nil {
		log.Printf("Error executing user data subscribe request: %v\n", err)
		return
	}

	subscribeResult, _ := json.MarshalIndent(resp1.Typed, "", "  ")
	log.Printf("Result: %s\n", subscribeResult)

	stream.On("message", func(message models.UserDataStreamEventsResponse) {
		b, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received message: %s\n", string(b))
	})

	log.Println("Subscribed. Waiting 10 seconds...")
	time.Sleep(50 * time.Second)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```
