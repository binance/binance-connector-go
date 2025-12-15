# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountCommission**](AccountAPI.md#AccountCommission) | **Get** /api/v3/account/commission | Query Commission Rates
[**AllOrderList**](AccountAPI.md#AllOrderList) | **Get** /api/v3/allOrderList | Query all Order lists
[**AllOrders**](AccountAPI.md#AllOrders) | **Get** /api/v3/allOrders | All orders
[**GetAccount**](AccountAPI.md#GetAccount) | **Get** /api/v3/account | Account information
[**GetOpenOrders**](AccountAPI.md#GetOpenOrders) | **Get** /api/v3/openOrders | Current open orders
[**GetOrder**](AccountAPI.md#GetOrder) | **Get** /api/v3/order | Query order
[**GetOrderList**](AccountAPI.md#GetOrderList) | **Get** /api/v3/orderList | Query Order list
[**MyAllocations**](AccountAPI.md#MyAllocations) | **Get** /api/v3/myAllocations | Query Allocations
[**MyFilters**](AccountAPI.md#MyFilters) | **Get** /api/v3/myFilters | Query relevant filters
[**MyPreventedMatches**](AccountAPI.md#MyPreventedMatches) | **Get** /api/v3/myPreventedMatches | Query Prevented Matches
[**MyTrades**](AccountAPI.md#MyTrades) | **Get** /api/v3/myTrades | Account trade list
[**OpenOrderList**](AccountAPI.md#OpenOrderList) | **Get** /api/v3/openOrderList | Query Open Order lists
[**OrderAmendments**](AccountAPI.md#OrderAmendments) | **Get** /api/v3/order/amendments | Query Order Amendments
[**RateLimitOrder**](AccountAPI.md#RateLimitOrder) | **Get** /api/v3/rateLimit/order | Query Unfilled Order Count


## AccountCommission

> AccountCommissionResponse AccountCommission(ctx).Symbol(symbol).Execute()

Query Commission Rates


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AccountCommission(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AccountCommission``: %v\n", err)
		return
	}

	// response from `AccountCommission`: AccountCommissionResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 

### Return type

[**AccountCommissionResponse**](AccountCommissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AllOrderList

> AllOrderListResponse AllOrderList(ctx).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query all Order lists


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AllOrderList(context.Background()).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AllOrderList``: %v\n", err)
		return
	}

	// response from `AllOrderList`: AllOrderListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrderListResponse**](AllOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All orders


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.AllOrders(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.AllOrders``: %v\n", err)
		return
	}

	// response from `AllOrders`: AllOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrdersResponse**](AllOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAccount

> GetAccountResponse GetAccount(ctx).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

Account information


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	omitZeroBalances := false // bool | When set to `true`, emits only the non-zero balances of an account. <br>Default value: `false` (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetAccount(context.Background()).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetAccount``: %v\n", err)
		return
	}

	// response from `GetAccount`: GetAccountResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. &lt;br&gt;Default value: &#x60;false&#x60; | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOpenOrders

> GetOpenOrdersResponse GetOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current open orders


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | Symbol to query (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetOpenOrders``: %v\n", err)
		return
	}

	// response from `GetOpenOrders`: GetOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Symbol to query | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOpenOrdersResponse**](GetOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query order


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetOrder``: %v\n", err)
		return
	}

	// response from `GetOrder`: GetOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrderList

> GetOrderListResponse GetOrderList(ctx).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order list


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.GetOrderList(context.Background()).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.GetOrderList``: %v\n", err)
		return
	}

	// response from `GetOrderList`: GetOrderListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOrderListResponse**](GetOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyAllocations

> MyAllocationsResponse MyAllocations(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()

Query Allocations


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	fromAllocationId := int32(1) // int32 |  (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	orderId := int64(1) // int64 |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.MyAllocations(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyAllocations``: %v\n", err)
		return
	}

	// response from `MyAllocations`: MyAllocationsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **fromAllocationId** | **int32** |  | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **orderId** | **int64** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyAllocationsResponse**](MyAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyFilters

> MyFiltersResponse MyFilters(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query relevant filters


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.MyFilters(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyFilters``: %v\n", err)
		return
	}

	// response from `MyFilters`: MyFiltersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyFiltersResponse**](MyFiltersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyPreventedMatches

> MyPreventedMatchesResponse MyPreventedMatches(ctx).Symbol(symbol).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Prevented Matches


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	preventedMatchId := int64(1) // int64 |  (optional)
	orderId := int64(1) // int64 |  (optional)
	fromPreventedMatchId := int64(1) // int64 |  (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.MyPreventedMatches(context.Background()).Symbol(symbol).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyPreventedMatches``: %v\n", err)
		return
	}

	// response from `MyPreventedMatches`: MyPreventedMatchesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **preventedMatchId** | **int64** |  | 
 **orderId** | **int64** |  | 
 **fromPreventedMatchId** | **int64** |  | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyPreventedMatchesResponse**](MyPreventedMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyTrades

> MyTradesResponse MyTrades(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account trade list


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades from INCLUSIVE. (optional)
	endTime := int64(1735693200000) // int64 | Timestamp in ms to get aggregate trades until INCLUSIVE. (optional)
	fromId := int64(1) // int64 | ID to get aggregate trades from INCLUSIVE. (optional)
	limit := int32(500) // int32 | Default: 500; Maximum: 1000. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.MyTrades(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.MyTrades``: %v\n", err)
		return
	}

	// response from `MyTrades`: MyTradesResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** | Timestamp in ms to get aggregate trades from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get aggregate trades until INCLUSIVE. | 
 **fromId** | **int64** | ID to get aggregate trades from INCLUSIVE. | 
 **limit** | **int32** | Default: 500; Maximum: 1000. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyTradesResponse**](MyTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenOrderList

> OpenOrderListResponse OpenOrderList(ctx).RecvWindow(recvWindow).Execute()

Query Open Order lists


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.OpenOrderList(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OpenOrderList``: %v\n", err)
		return
	}

	// response from `OpenOrderList`: OpenOrderListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrderListResponse**](OpenOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderAmendments

> OrderAmendmentsResponse OrderAmendments(ctx).Symbol(symbol).OrderId(orderId).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Order Amendments


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	orderId := int64(1) // int64 | 
	fromExecutionId := int64(1) // int64 |  (optional)
	limit := int64(500) // int64 | Default:500; Maximum: 1000  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.OrderAmendments(context.Background()).Symbol(symbol).OrderId(orderId).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.OrderAmendments``: %v\n", err)
		return
	}

	// response from `OrderAmendments`: OrderAmendmentsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 
 **orderId** | **int64** |  | 
 **fromExecutionId** | **int64** |  | 
 **limit** | **int64** | Default:500; Maximum: 1000  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendmentsResponse**](OrderAmendmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RateLimitOrder

> RateLimitOrderResponse RateLimitOrder(ctx).RecvWindow(recvWindow).Execute()

Query Unfilled Order Count


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.AccountAPI.RateLimitOrder(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `AccountAPI.RateLimitOrder``: %v\n", err)
		return
	}

	// response from `RateLimitOrder`: RateLimitOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**RateLimitOrderResponse**](RateLimitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

