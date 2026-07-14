# \AccountAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountCommission**](AccountAPI.md#AccountCommission) | /account.commission | Account Commission Rates (USER_DATA)
[**AccountRateLimitsOrders**](AccountAPI.md#AccountRateLimitsOrders) | /account.rateLimits.orders | Unfilled Order Count (USER_DATA)
[**AccountStatus**](AccountAPI.md#AccountStatus) | /account.status | Account information (USER_DATA)
[**AllOrderLists**](AccountAPI.md#AllOrderLists) | /allOrderLists | Account order list history (USER_DATA)
[**AllOrders**](AccountAPI.md#AllOrders) | /allOrders | Account order history (USER_DATA)
[**MyAllocations**](AccountAPI.md#MyAllocations) | /myAllocations | Account allocations (USER_DATA)
[**MyFilters**](AccountAPI.md#MyFilters) | /myFilters | Query Relevant Filters (USER_DATA)
[**MyPreventedMatches**](AccountAPI.md#MyPreventedMatches) | /myPreventedMatches | Account prevented matches (USER_DATA)
[**MyTrades**](AccountAPI.md#MyTrades) | /myTrades | Account trade history (USER_DATA)
[**OpenOrderListsStatus**](AccountAPI.md#OpenOrderListsStatus) | /openOrderLists.status | Current open Order lists (USER_DATA)
[**OpenOrdersStatus**](AccountAPI.md#OpenOrdersStatus) | /openOrders.status | Current open orders (USER_DATA)
[**OrderAmendments**](AccountAPI.md#OrderAmendments) | /order.amendments | Query Order Amendments (USER_DATA)
[**OrderListStatus**](AccountAPI.md#OrderListStatus) | /orderList.status | Query Order list (USER_DATA)
[**OrderStatus**](AccountAPI.md#OrderStatus) | /order.status | Query order (USER_DATA)


## AccountCommission

> AccountCommissionResponse AccountCommission().Symbol(symbol).Id(id).Execute()

Account Commission Rates (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.AccountCommission().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountCommission``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 

### Return type

[**AccountCommissionResponse**](AccountCommissionResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AccountRateLimitsOrders

> AccountRateLimitsOrdersResponse AccountRateLimitsOrders().Id(id).RecvWindow(recvWindow).Execute()

Unfilled Order Count (USER_DATA)


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


	resp, err := wsClient.WebsocketAPI.AccountAPI.AccountRateLimitsOrders().Id(id).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountRateLimitsOrders``: %v\n", err)
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
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AccountRateLimitsOrdersResponse**](AccountRateLimitsOrdersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AccountStatus

> AccountStatusResponse AccountStatus().Id(id).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

Account information (USER_DATA)


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
	omitZeroBalances := false // bool | When set to `true`, emits only the non-zero balances of an account. Default value: `false`. (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.AccountStatus().Id(id).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountStatus``: %v\n", err)
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
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. Default value: &#x60;false&#x60;. | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AccountStatusResponse**](AccountStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllOrderLists

> AllOrderListsResponse AllOrderLists().Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account order list history (USER_DATA)


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
	fromId := int32(1) // int32 | Order list ID to begin at (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	limit := int32(1) // int32 | Default: 500; Maximum: 1000 (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.AllOrderLists().Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AllOrderLists``: %v\n", err)
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
 **fromId** | **int32** | Order list ID to begin at | 
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **limit** | **int32** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrderListsResponse**](AllOrderListsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account order history (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	orderId := int64(1) // int64 | Order ID to begin at (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	limit := int32(1) // int32 | Default: 500; Maximum: 1000 (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.AllOrders().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AllOrders``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **orderId** | **int64** | Order ID to begin at | 
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **limit** | **int32** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrdersResponse**](AllOrdersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyAllocations

> MyAllocationsResponse MyAllocations().Symbol(symbol).Id(id).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()

Account allocations (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	fromAllocationId := int32(1) // int32 | Allocation ID to begin at (optional)
	limit := int32(1) // int32 | Default: 500; Maximum: 1000 (optional)
	orderId := int64(1) // int64 | Order ID (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.MyAllocations().Symbol(symbol).Id(id).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyAllocations``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **fromAllocationId** | **int32** | Allocation ID to begin at | 
 **limit** | **int32** | Default: 500; Maximum: 1000 | 
 **orderId** | **int64** | Order ID | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyAllocationsResponse**](MyAllocationsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyFilters

> MyFiltersResponse MyFilters().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()

Query Relevant Filters (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.MyFilters().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyFilters``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyFiltersResponse**](MyFiltersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyPreventedMatches

> MyPreventedMatchesResponse MyPreventedMatches().Symbol(symbol).Id(id).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

Account prevented matches (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	preventedMatchId := int64(1) // int64 | Prevented match ID (optional)
	orderId := int64(1) // int64 | Order ID (optional)
	fromPreventedMatchId := int64(1) // int64 | Prevented match ID to begin at (optional)
	limit := int32(1) // int32 | Default: 500; Maximum: 1000 (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.MyPreventedMatches().Symbol(symbol).Id(id).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyPreventedMatches``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **preventedMatchId** | **int64** | Prevented match ID | 
 **orderId** | **int64** | Order ID | 
 **fromPreventedMatchId** | **int64** | Prevented match ID to begin at | 
 **limit** | **int32** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyPreventedMatchesResponse**](MyPreventedMatchesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyTrades

> MyTradesResponse MyTrades().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account trade history (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	orderId := int64(1) // int64 | This can only be used in combination with `symbol`. (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms (optional)
	fromId := int32(1) // int32 | First trade ID to query (optional)
	limit := int32(1) // int32 | Default: 500; Maximum: 1000 (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.MyTrades().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyTrades``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **orderId** | **int64** | This can only be used in combination with &#x60;symbol&#x60;. | 
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **fromId** | **int32** | First trade ID to query | 
 **limit** | **int32** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyTradesResponse**](MyTradesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OpenOrderListsStatus

> OpenOrderListsStatusResponse OpenOrderListsStatus().Id(id).RecvWindow(recvWindow).Execute()

Current open Order lists (USER_DATA)


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


	resp, err := wsClient.WebsocketAPI.AccountAPI.OpenOrderListsStatus().Id(id).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OpenOrderListsStatus``: %v\n", err)
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
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrderListsStatusResponse**](OpenOrderListsStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OpenOrdersStatus

> OpenOrdersStatusResponse OpenOrdersStatus().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current open orders (USER_DATA)


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
	symbol := "BNBUSDT" // string | If omitted, open orders for all symbols are returned (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.OpenOrdersStatus().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OpenOrdersStatus``: %v\n", err)
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
 **symbol** | **string** | If omitted, open orders for all symbols are returned | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrdersStatusResponse**](OpenOrdersStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderAmendments

> OrderAmendmentsResponse OrderAmendments().Symbol(symbol).OrderId(orderId).Id(id).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Order Amendments (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 | Order ID
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	fromExecutionId := int64(1) // int64 | Execution ID to begin at (optional)
	limit := int64(1) // int64 | Default: 500; Maximum: 1000 (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.OrderAmendments().Symbol(symbol).OrderId(orderId).Id(id).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OrderAmendments``: %v\n", err)
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
 **symbol** | **string** |  | 
 **orderId** | **int64** | Order ID | 
 **id** | **string** | Client-generated request identifier. | 
 **fromExecutionId** | **int64** | Execution ID to begin at | 
 **limit** | **int64** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendmentsResponse**](OrderAmendmentsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListStatus

> OrderListStatusResponse OrderListStatus().Id(id).OrigClientOrderId(origClientOrderId).OrderListId(orderListId).RecvWindow(recvWindow).Execute()

Query Order list (USER_DATA)


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
	origClientOrderId := "08985fedd9ea2cf6b28996" // string | Query order list by `listClientOrderId`. `orderListId` or `origClientOrderId` must be provided. (optional)
	orderListId := int32(1) // int32 | Query order list by `orderListId`. `orderListId` or `origClientOrderId` must be provided. (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.OrderListStatus().Id(id).OrigClientOrderId(origClientOrderId).OrderListId(orderListId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OrderListStatus``: %v\n", err)
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
 **origClientOrderId** | **string** | Query order list by &#x60;listClientOrderId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided. | 
 **orderListId** | **int32** | Query order list by &#x60;orderListId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided. | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListStatusResponse**](OrderListStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderStatus

> OrderStatusResponse OrderStatus().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query order (USER_DATA)


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
	symbol := "BNBUSDT" // string | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	orderId := int64(1) // int64 | Lookup order by `orderId` (optional)
	origClientOrderId := "myOrder1" // string | Lookup order by `clientOrderId` (optional)
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


	resp, err := wsClient.WebsocketAPI.AccountAPI.OrderStatus().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OrderStatus``: %v\n", err)
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
 **symbol** | **string** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **orderId** | **int64** | Lookup order by &#x60;orderId&#x60; | 
 **origClientOrderId** | **string** | Lookup order by &#x60;clientOrderId&#x60; | 
 **recvWindow** | **float32** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderStatusResponse**](OrderStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

