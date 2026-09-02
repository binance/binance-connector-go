# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountCommission**](AccountAPI.md#AccountCommission) | **Get** /api/v3/account/commission | Query Commission Rates (USER_DATA)
[**AllOrderList**](AccountAPI.md#AllOrderList) | **Get** /api/v3/allOrderList | Query all Order lists (USER_DATA)
[**AllOrders**](AccountAPI.md#AllOrders) | **Get** /api/v3/allOrders | All orders (USER_DATA)
[**GetAccount**](AccountAPI.md#GetAccount) | **Get** /api/v3/account | Account information (USER_DATA)
[**GetOpenOrders**](AccountAPI.md#GetOpenOrders) | **Get** /api/v3/openOrders | Current open orders (USER_DATA)
[**GetOrder**](AccountAPI.md#GetOrder) | **Get** /api/v3/order | Query order (USER_DATA)
[**GetOrderList**](AccountAPI.md#GetOrderList) | **Get** /api/v3/orderList | Query Order list (USER_DATA)
[**MyAllocations**](AccountAPI.md#MyAllocations) | **Get** /api/v3/myAllocations | Query Allocations (USER_DATA)
[**MyFilters**](AccountAPI.md#MyFilters) | **Get** /api/v3/myFilters | Query relevant filters (USER_DATA)
[**MyPreventedMatches**](AccountAPI.md#MyPreventedMatches) | **Get** /api/v3/myPreventedMatches | Query Prevented Matches (USER_DATA)
[**MyTrades**](AccountAPI.md#MyTrades) | **Get** /api/v3/myTrades | Account trade list (USER_DATA)
[**OpenOrderList**](AccountAPI.md#OpenOrderList) | **Get** /api/v3/openOrderList | Query Open Order lists (USER_DATA)
[**OrderAmendments**](AccountAPI.md#OrderAmendments) | **Get** /api/v3/order/amendments | Query Order Amendments (USER_DATA)
[**RateLimitOrder**](AccountAPI.md#RateLimitOrder) | **Get** /api/v3/rateLimit/order | Query Unfilled Order Count (USER_DATA)


## AccountCommission

> AccountCommissionResponse AccountCommission(ctx).Symbol(symbol).Execute()

Query Commission Rates (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 

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

Query all Order lists (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromId := int64(1) // int64 | If supplied, neither startTime or endTime can be provided (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	limit := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **fromId** | **int64** | If supplied, neither startTime or endTime can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrderListResponse**](AllOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## AllOrders

> AllOrdersResponse AllOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

All orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "LTCBTC" // string | 
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	limit := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** |  | 
 **recvWindow** | **float64** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**AllOrdersResponse**](AllOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetAccount

> GetAccountResponse GetAccount(ctx).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

Account information (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	omitZeroBalances := false // bool | When set to `true`, emits only the non-zero balances of an account. (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetAccountResponse**](GetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOpenOrders

> GetOpenOrdersResponse GetOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Current open orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "LTCBTC" // string |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **symbol** | **string** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOpenOrdersResponse**](GetOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "LTCBTC" // string | 
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "myOrder1" // string |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetOrderList

> GetOrderListResponse GetOrderList(ctx).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order list (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderListId := int64(27) // int64 | Query order list by `orderListId`. `orderListId` or `origClientOrderId` must be provided. (optional)
	origClientOrderId := "1" // string | Query order list by `listClientOrderId`. `orderListId` or `origClientOrderId` must be provided. (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **orderListId** | **int64** | Query order list by &#x60;orderListId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided. | 
 **origClientOrderId** | **string** | Query order list by &#x60;listClientOrderId&#x60;. &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**GetOrderListResponse**](GetOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyAllocations

> MyAllocationsResponse MyAllocations(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Execute()

Query Allocations (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	fromAllocationId := int32(0) // int32 |  (optional)
	limit := int32(1) // int32 |  (optional)
	orderId := int64(1) // int64 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromAllocationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **orderId** | **int64** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyAllocationsResponse**](MyAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyFilters

> MyFiltersResponse MyFilters(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query relevant filters (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBUSDT" // string | 
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyFiltersResponse**](MyFiltersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyPreventedMatches

> MyPreventedMatchesResponse MyPreventedMatches(ctx).Symbol(symbol).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Prevented Matches (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	preventedMatchId := int64(1) // int64 |  (optional)
	orderId := int64(1) // int64 |  (optional)
	fromPreventedMatchId := int64(1) // int64 |  (optional)
	limit := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **limit** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyPreventedMatchesResponse**](MyPreventedMatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MyTrades

> MyTradesResponse MyTrades(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account trade list (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BNBBTC" // string | 
	orderId := int64(100234) // int64 | This can only be used in combination with `symbol`. (optional)
	startTime := int64(1735693200000) // int64 |  (optional)
	endTime := int64(1735693200000) // int64 |  (optional)
	fromId := int64(1) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **orderId** | **int64** | This can only be used in combination with &#x60;symbol&#x60;. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**MyTradesResponse**](MyTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OpenOrderList

> OpenOrderListResponse OpenOrderList(ctx).RecvWindow(recvWindow).Execute()

Query Open Order lists (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrderListResponse**](OpenOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderAmendments

> OrderAmendmentsResponse OrderAmendments(ctx).Symbol(symbol).OrderId(orderId).FromExecutionId(fromExecutionId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Order Amendments (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | 
	orderId := int64(9) // int64 | 
	fromExecutionId := int64(22) // int64 |  (optional)
	limit := int64(1) // int64 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **limit** | **int64** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendmentsResponse**](OrderAmendmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## RateLimitOrder

> RateLimitOrderResponse RateLimitOrder(ctx).RecvWindow(recvWindow).Execute()

Query Unfilled Order Count (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**RateLimitOrderResponse**](RateLimitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

