# \UserDataStreamsAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CloseUserDataStream**](UserDataStreamsAPI.md#CloseUserDataStream) | /userDataStream.stop | Close User Data Stream (USER_STREAM)
[**KeepaliveUserDataStream**](UserDataStreamsAPI.md#KeepaliveUserDataStream) | /userDataStream.ping | Keepalive User Data Stream (USER_STREAM)
[**StartUserDataStream**](UserDataStreamsAPI.md#StartUserDataStream) | /userDataStream.start | Start User Data Stream (USER_STREAM)


## CloseUserDataStream

> CloseUserDataStreamResponse CloseUserDataStream().Id(id).Execute()

Close User Data Stream (USER_STREAM)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.UserDataStreamsAPI.CloseUserDataStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.CloseUserDataStream``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Unique WebSocket request ID. | 

### Return type

[**CloseUserDataStreamResponse**](CloseUserDataStreamResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## KeepaliveUserDataStream

> KeepaliveUserDataStreamResponse KeepaliveUserDataStream().Id(id).Execute()

Keepalive User Data Stream (USER_STREAM)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.UserDataStreamsAPI.KeepaliveUserDataStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.KeepaliveUserDataStream``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Unique WebSocket request ID. | 

### Return type

[**KeepaliveUserDataStreamResponse**](KeepaliveUserDataStreamResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## StartUserDataStream

> StartUserDataStreamResponse StartUserDataStream().Id(id).Execute()

Start User Data Stream (USER_STREAM)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.UserDataStreamsAPI.StartUserDataStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamsAPI.StartUserDataStream``: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	err = wsClient.WebsocketAPI.CloseWebSocketConnection()
	if err != nil {
		log.Printf("Error closing WebSocket connection: %v\n", err)
		return
	}
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Unique WebSocket request ID. | 

### Return type

[**StartUserDataStreamResponse**](StartUserDataStreamResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

