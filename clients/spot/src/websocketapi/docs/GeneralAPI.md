# \GeneralAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**ExchangeInfo**](GeneralAPI.md#ExchangeInfo) | /exchangeInfo | Exchange information
[**ExecutionRules**](GeneralAPI.md#ExecutionRules) | /executionRules | Query Execution Rules
[**Ping**](GeneralAPI.md#Ping) | /ping | Test connectivity
[**Time**](GeneralAPI.md#Time) | /time | Check server time


## ExchangeInfo

> ExchangeInfoResponse ExchangeInfo().Id(id).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()

Exchange information


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
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	symbols := []string{"BTCUSDT"} // []string | Describe multiple symbols (optional)
	permissions := []string{"SPOT"} // []string | Filter symbols by permissions (optional)
	showPermissionSets := true // bool | Controls whether the content of the `permissionSets` field is populated or not. Defaults to `true`. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Filters for symbols that have this `tradingStatus`. Valid values: `TRADING`, `HALT`, `BREAK`. Cannot be used in combination with `symbol` or `symbols`. (optional)

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


	resp, err := wsClient.WebsocketAPI.GeneralAPI.ExchangeInfo().Id(id).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.ExchangeInfo``: %v\n", err)
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
 **symbol** | **string** | Describe a single symbol | 
 **symbols** | **[]string** | Describe multiple symbols | 
 **permissions** | **[]string** | Filter symbols by permissions | 
 **showPermissionSets** | **bool** | Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not. Defaults to &#x60;true&#x60;. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Filters for symbols that have this &#x60;tradingStatus&#x60;. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;. Cannot be used in combination with &#x60;symbol&#x60; or &#x60;symbols&#x60;. | 

### Return type

[**ExchangeInfoResponse**](ExchangeInfoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ExecutionRules

> ExecutionRulesResponse ExecutionRules().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()

Query Execution Rules


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
	symbol := "BAZUSD" // string | Query for specified symbol. (optional)
	symbols := []string{"BAZUSD"} // []string | Query for multiple symbols. (optional)
	symbolStatus := models.ExchangeInfoSymbolStatusParameterTrading // ExchangeInfoSymbolStatusParameter | Query for all symbols with the specified status. Supported values: `TRADING`, `HALT`, `BREAK` (optional)

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


	resp, err := wsClient.WebsocketAPI.GeneralAPI.ExecutionRules().Id(id).Symbol(symbol).Symbols(symbols).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.ExecutionRules``: %v\n", err)
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
 **symbol** | **string** | Query for specified symbol. | 
 **symbols** | **[]string** | Query for multiple symbols. | 
 **symbolStatus** | [**ExchangeInfoSymbolStatusParameter**](ExchangeInfoSymbolStatusParameter.md) | Query for all symbols with the specified status. Supported values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; | 

### Return type

[**ExecutionRulesResponse**](ExecutionRulesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ping

> PingResponse Ping().Id(id).Execute()

Test connectivity


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


	resp, err := wsClient.WebsocketAPI.GeneralAPI.Ping().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.Ping``: %v\n", err)
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

[**PingResponse**](PingResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Time

> TimeResponse Time().Id(id).Execute()

Check server time


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


	resp, err := wsClient.WebsocketAPI.GeneralAPI.Time().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `GeneralAPI.Time``: %v\n", err)
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

[**TimeResponse**](TimeResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

