# \UserDataStreamAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**SessionSubscriptions**](UserDataStreamAPI.md#SessionSubscriptions) | /session.subscriptions | Listing all subscriptions
[**UserDataStreamSubscribe**](UserDataStreamAPI.md#UserDataStreamSubscribe) | /userDataStream.subscribe | Subscribe to User Data Stream
[**UserDataStreamSubscribeSignature**](UserDataStreamAPI.md#UserDataStreamSubscribeSignature) | /userDataStream.subscribe.signature | Subscribe to User Data Stream through signature subscription (USER_STREAM)
[**UserDataStreamUnsubscribe**](UserDataStreamAPI.md#UserDataStreamUnsubscribe) | /userDataStream.unsubscribe | WebSocket Unsubscribe from User Data Stream


## SessionSubscriptions

> SessionSubscriptionsResponse SessionSubscriptions().Id(id).Execute()

Listing all subscriptions


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.UserDataStreamAPI.SessionSubscriptions().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.SessionSubscriptions``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 

### Return type

[**SessionSubscriptionsResponse**](SessionSubscriptionsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## UserDataStreamSubscribe

> UserDataStreamSubscribeResponse UserDataStreamSubscribe().Id(id).Execute()

Subscribe to User Data Stream


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	resp, stream, err := wsClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribe().Execute()
	if err != nil {
		log.Printf("Error executing UserDataStreamSubscribe request: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	stream.On("message", func(message map[string]interface{}) {
		formattedMessage, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received event: %s", string(formattedMessage))
	})
	
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
 **id** | **string** | Client-generated request identifier. | 

### Return type

[**UserDataStreamSubscribeResponse**](UserDataStreamSubscribeResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## UserDataStreamSubscribeSignature

> UserDataStreamSubscribeSignatureResponse UserDataStreamSubscribeSignature().Id(id).RecvWindow(recvWindow).Execute()

Subscribe to User Data Stream through signature subscription (USER_STREAM)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	recvWindow := float32(5000) // float32 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	resp, stream, err := wsClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribeSignature().Execute()
	if err != nil {
		log.Printf("Error executing UserDataStreamSubscribeSignature request: %v\n", err)
		return
	}

	result, _ := json.MarshalIndent(resp.Typed, "", "  ")
	log.Printf("Result: %s\n", result)

	stream.On("message", func(message map[string]interface{}) {
		formattedMessage, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received event: %s", string(formattedMessage))
	})
	
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
 **id** | **string** | Client-generated request identifier. | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**UserDataStreamSubscribeSignatureResponse**](UserDataStreamSubscribeSignatureResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## UserDataStreamUnsubscribe

> UserDataStreamUnsubscribeResponse UserDataStreamUnsubscribe().Id(id).SubscriptionId(subscriptionId).Execute()

WebSocket Unsubscribe from User Data Stream


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	subscriptionId := int32(1) // int32 | When called with no parameter, this will close all subscriptions. When called with the subscriptionId parameter, this will attempt to close the subscription with that subscription id, if it exists. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamUnsubscribe().Id(id).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `UserDataStreamAPI.UserDataStreamUnsubscribe``: %v\n", err)
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
 **id** | **string** | Client-generated request identifier. | 
 **subscriptionId** | **int32** | When called with no parameter, this will close all subscriptions. When called with the subscriptionId parameter, this will attempt to close the subscription with that subscription id, if it exists. | 

### Return type

[**UserDataStreamUnsubscribeResponse**](UserDataStreamUnsubscribeResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

