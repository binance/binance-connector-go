# \AuthAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**SessionLogon**](AuthAPI.md#SessionLogon) | /session.logon | Log in with API key (USER_DATA)
[**SessionLogout**](AuthAPI.md#SessionLogout) | /session.logout | Log out of the session
[**SessionStatus**](AuthAPI.md#SessionStatus) | /session.status | Query session status


## SessionLogon

> SessionLogonResponse SessionLogon().Id(id).RecvWindow(recvWindow).Execute()

Log in with API key (USER_DATA)


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
	recvWindow := float32(5000) // float32 | The value cannot be greater than `60000`. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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


	resp, err := wsClient.WebsocketAPI.AuthAPI.SessionLogon().Id(id).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AuthAPI.SessionLogon``: %v\n", err)
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
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SessionLogonResponse**](SessionLogonResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SessionLogout

> SessionLogoutResponse SessionLogout().Id(id).Execute()

Log out of the session


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


	resp, err := wsClient.WebsocketAPI.AuthAPI.SessionLogout().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AuthAPI.SessionLogout``: %v\n", err)
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

[**SessionLogoutResponse**](SessionLogoutResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SessionStatus

> SessionStatusResponse SessionStatus().Id(id).Execute()

Query session status


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


	resp, err := wsClient.WebsocketAPI.AuthAPI.SessionStatus().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AuthAPI.SessionStatus``: %v\n", err)
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

[**SessionStatusResponse**](SessionStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

