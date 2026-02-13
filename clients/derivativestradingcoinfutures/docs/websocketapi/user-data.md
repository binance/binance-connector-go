# User Data

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	SessionLogon()
}

func SessionLogon() {
	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.DerivativesTradingCoinFuturesWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsPrivateKey("Your Private Key"),
	)

	wsClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
		client.WithWebsocketAPI(configuration),
	)
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		fmt.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	responseChan, errorChan, err := wsClient.WebsocketAPI.AuthAPI.SessionLogon().Execute()
	if err != nil {
		fmt.Printf("Error executing session logon request: %v\n", err)
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
		fmt.Printf("Error executing user data stream subscribe request: %v\n", err)
		return
	}


	subscribeResult, _ := json.MarshalIndent(resp1.Typed, "", "  ")
	fmt.Printf("Result: %s\n", subscribeResult)

	stream.On("message", func(message map[string]interface{}) {
		eventData, ok := message["event"].(map[string]interface{})
		if !ok {
			log.Println("Failed to determine event type")
			return
		}

		eventType, ok := eventData["e"].(string)
		if !ok {
			log.Println("Failed to determine event type")
			jsonData, _ := json.MarshalIndent(message, "", "  ")
			log.Printf("Full message structure:\n%s\n", string(jsonData))
			return
		}

		var response models.UserDataStreamEventsResponse

		data, _ := json.Marshal(message)
		switch eventType {
		case "outboundAccountPosition":
			response.OutboundAccountPosition = &models.OutboundAccountPosition{}
			json.Unmarshal(data, response.OutboundAccountPosition)
		case "balanceUpdate":
			response.BalanceUpdate = &models.BalanceUpdate{}
			json.Unmarshal(data, response.BalanceUpdate)
		case "executionReport":
			response.ExecutionReport = &models.ExecutionReport{}
			json.Unmarshal(data, response.ExecutionReport)
		case "listStatus":
			response.ListStatus = &models.ListStatus{}
			json.Unmarshal(data, response.ListStatus)
		case "eventStreamTerminated":
			response.EventStreamTerminated = &models.EventStreamTerminated{}
			json.Unmarshal(data, response.EventStreamTerminated)
		case "externalLockUpdate":
			response.ExternalLockUpdate = &models.ExternalLockUpdate{}
			json.Unmarshal(data, response.ExternalLockUpdate)
		default:
			log.Printf("Unhandled event type: %s", eventType)
			return
		}
		jsonData, _ := json.MarshalIndent(response, "", "  ")
		log.Printf("Received event: %s", string(jsonData))
	})

	log.Println("Subscribed. Waiting 10 seconds...")
	time.Sleep(50 * time.Second)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		fmt.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```
