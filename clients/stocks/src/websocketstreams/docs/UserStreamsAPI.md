# \UserStreamsAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**OrderReportStream**](UserStreamsAPI.md#OrderReportStream) | /&lt;listenKey&gt;@orderReport | Order Report Stream


## OrderReportStream

Order Report Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	listenKey := "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s6af0dkfj2a97b8a91d" // string | User data listen key obtained from the Listen Key endpoint.

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.UserStreamsAPI.OrderReportStream().ListenKey(listenKey).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserStreamsAPI.OrderReportStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.OrderReportStreamResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** | User data listen key obtained from the Listen Key endpoint. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

