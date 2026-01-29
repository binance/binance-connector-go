# \AccountAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountCommission**](AccountAPI.md#AccountCommission) | /account.commission | WebSocket Account Commission Rates
[**AccountRateLimitsOrders**](AccountAPI.md#AccountRateLimitsOrders) | /account.rateLimits.orders | WebSocket Unfilled Order Count
[**AccountStatus**](AccountAPI.md#AccountStatus) | /account.status | WebSocket Account information
[**AllOrderLists**](AccountAPI.md#AllOrderLists) | /allOrderLists | WebSocket Account order list history
[**AllOrders**](AccountAPI.md#AllOrders) | /allOrders | WebSocket Account order history
[**MyAllocations**](AccountAPI.md#MyAllocations) | /myAllocations | WebSocket Account allocations
[**MyFilters**](AccountAPI.md#MyFilters) | /myFilters | WebSocket Query Relevant Filters
[**MyPreventedMatches**](AccountAPI.md#MyPreventedMatches) | /myPreventedMatches | WebSocket Account prevented matches
[**MyTrades**](AccountAPI.md#MyTrades) | /myTrades | WebSocket Account trade history
[**OpenOrderListsStatus**](AccountAPI.md#OpenOrderListsStatus) | /openOrderLists.status | WebSocket Current open Order lists
[**OpenOrdersStatus**](AccountAPI.md#OpenOrdersStatus) | /openOrders.status | WebSocket Current open orders
[**OrderAmendments**](AccountAPI.md#OrderAmendments) | /order.amendments | WebSocket Query Order Amendments
[**OrderListStatus**](AccountAPI.md#OrderListStatus) | /orderList.status | WebSocket Query Order list
[**OrderStatus**](AccountAPI.md#OrderStatus) | /order.status | WebSocket Query order


## AccountCommission

> AccountCommissionResponse AccountCommission().Symbol(symbol).Id(id).Execute()

WebSocket Account Commission Rates


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 

### Return type

[**AccountCommissionResponse**](AccountCommissionResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AccountRateLimitsOrders

> AccountRateLimitsOrdersResponse AccountRateLimitsOrders().Id(id).RecvWindow(recvWindow).Execute()

WebSocket Unfilled Order Count


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AccountRateLimitsOrdersResponse**](AccountRateLimitsOrdersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AccountStatus

> AccountStatusResponse AccountStatus().Id(id).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

WebSocket Account information


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	omitZeroBalances := false // bool | When set to `true`, emits only the non-zero balances of an account. <br>Default value: false (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. &lt;br&gt;Default value: false | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AccountStatusResponse**](AccountStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllOrderLists

> AllOrderListsResponse AllOrderLists().Id(id).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

WebSocket Account order list history


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	fromId := int32(1) // int32 | Trade ID to begin at (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **fromId** | **int32** | Trade ID to begin at | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrderListsResponse**](AllOrderListsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

WebSocket Account order history


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 | `orderId`or`origClientOrderId`mustbesent (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrdersResponse**](AllOrdersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyAllocations

> MyAllocationsResponse MyAllocations().Symbol(symbol).Id(id).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()

WebSocket Account allocations


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	fromAllocationId := int32(1) // int32 |  (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	orderId := int64(1) // int64 | `orderId`or`origClientOrderId`mustbesent (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromAllocationId** | **int32** |  | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **orderId** | **int64** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyAllocationsResponse**](MyAllocationsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyFilters

> MyFiltersResponse MyFilters().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()

WebSocket Query Relevant Filters


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyFiltersResponse**](MyFiltersResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyPreventedMatches

> MyPreventedMatchesResponse MyPreventedMatches().Symbol(symbol).Id(id).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

WebSocket Account prevented matches


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	preventedMatchId := int64(1) // int64 |  (optional)
	orderId := int64(1) // int64 | `orderId`or`origClientOrderId`mustbesent (optional)
	fromPreventedMatchId := int64(1) // int64 |  (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **preventedMatchId** | **int64** |  | 
 **orderId** | **int64** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **fromPreventedMatchId** | **int64** |  | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyPreventedMatchesResponse**](MyPreventedMatchesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MyTrades

> MyTradesResponse MyTrades().Symbol(symbol).Id(id).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

WebSocket Account trade history


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 | `orderId`or`origClientOrderId`mustbesent (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	fromId := int32(1) // int32 | Trade ID to begin at (optional)
	limit := int32(100) // int32 | Default: 100; Maximum: 5000 (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int32** | Trade ID to begin at | 
 **limit** | **int32** | Default: 100; Maximum: 5000 | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyTradesResponse**](MyTradesResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OpenOrderListsStatus

> OpenOrderListsStatusResponse OpenOrderListsStatus().Id(id).RecvWindow(recvWindow).Execute()

WebSocket Current open Order lists


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrderListsStatusResponse**](OpenOrderListsStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OpenOrdersStatus

> OpenOrdersStatusResponse OpenOrdersStatus().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()

WebSocket Current open orders


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "BNBUSDT" // string | Describe a single symbol (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** | Describe a single symbol | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrdersStatusResponse**](OpenOrdersStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderAmendments

> OrderAmendmentsResponse OrderAmendments().Symbol(symbol).OrderId(orderId).Id(id).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()

WebSocket Query Order Amendments


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	fromExecutionId := int64(1) // int64 |  (optional)
	limit := int64(500) // int64 | Default: 500; Maximum: 1000 (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **orderId** | **int64** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **fromExecutionId** | **int64** |  | 
 **limit** | **int64** | Default: 500; Maximum: 1000 | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendmentsResponse**](OrderAmendmentsResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListStatus

> OrderListStatusResponse OrderListStatus().Id(id).OrigClientOrderId(origClientOrderId).OrderListId(orderListId).RecvWindow(recvWindow).Execute()

WebSocket Query Order list


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	origClientOrderId := "origClientOrderId_example" // string | `orderId`or`origClientOrderId`mustbesent (optional)
	orderListId := int32(1) // int32 | Cancel order list by orderListId (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **origClientOrderId** | **string** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **orderListId** | **int32** | Cancel order list by orderListId | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListStatusResponse**](OrderListStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderStatus

> OrderStatusResponse OrderStatus().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

WebSocket Query order


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 | `orderId`or`origClientOrderId`mustbesent (optional)
	origClientOrderId := "origClientOrderId_example" // string | `orderId`or`origClientOrderId`mustbesent (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **origClientOrderId** | **string** | &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderStatusResponse**](OrderStatusResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

