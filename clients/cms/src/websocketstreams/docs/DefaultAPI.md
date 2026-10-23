# \DefaultAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ComAnnouncementEn**](DefaultAPI.md#ComAnnouncementEn) | /com_announcement_en | 


## ComAnnouncementEn




### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/cms"
	responseModels "github.com/binance/binance-connector-go/clients/cms/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceCmsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ComAnnouncementEn().Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.ComAnnouncementEn``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.ComAnnouncementEnResponse) {
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

This endpoint does not need any parameter.

### Authorization

No authorization required

[[Back to README]](../../../README.md)

