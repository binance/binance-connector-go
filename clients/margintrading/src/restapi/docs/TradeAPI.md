# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CreateSpecialKey**](TradeAPI.md#CreateSpecialKey) | **Post** /sapi/v1/margin/apiKey | Create Special Key(Low-Latency Trading)(TRADE)
[**DeleteSpecialKey**](TradeAPI.md#DeleteSpecialKey) | **Delete** /sapi/v1/margin/apiKey | Delete Special Key(Low-Latency Trading)(TRADE)
[**EditIpForSpecialKey**](TradeAPI.md#EditIpForSpecialKey) | **Put** /sapi/v1/margin/apiKey/ip | Edit ip for Special Key(Low-Latency Trading)(TRADE)
[**GetForceLiquidationRecord**](TradeAPI.md#GetForceLiquidationRecord) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**GetSmallLiabilityExchangeCoinList**](TradeAPI.md#GetSmallLiabilityExchangeCoinList) | **Get** /sapi/v1/margin/exchange-small-liability | Get Small Liability Exchange Coin List (USER_DATA)
[**GetSmallLiabilityExchangeHistory**](TradeAPI.md#GetSmallLiabilityExchangeHistory) | **Get** /sapi/v1/margin/exchange-small-liability-history | Get Small Liability Exchange History (USER_DATA)
[**MarginAccountCancelAllOpenOrdersOnASymbol**](TradeAPI.md#MarginAccountCancelAllOpenOrdersOnASymbol) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**MarginAccountCancelOco**](TradeAPI.md#MarginAccountCancelOco) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**MarginAccountCancelOrder**](TradeAPI.md#MarginAccountCancelOrder) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginAccountNewOco**](TradeAPI.md#MarginAccountNewOco) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**MarginAccountNewOrder**](TradeAPI.md#MarginAccountNewOrder) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginAccountNewOto**](TradeAPI.md#MarginAccountNewOto) | **Post** /sapi/v1/margin/order/oto | Margin Account New OTO (TRADE)
[**MarginAccountNewOtoco**](TradeAPI.md#MarginAccountNewOtoco) | **Post** /sapi/v1/margin/order/otoco | Margin Account New OTOCO (TRADE)
[**MarginManualLiquidation**](TradeAPI.md#MarginManualLiquidation) | **Post** /sapi/v1/margin/manual-liquidation | Margin Manual Liquidation(MARGIN)
[**QueryCurrentMarginOrderCountUsage**](TradeAPI.md#QueryCurrentMarginOrderCountUsage) | **Get** /sapi/v1/margin/rateLimit/order | Query Current Margin Order Count Usage (TRADE)
[**QueryMarginAccountsAllOco**](TradeAPI.md#QueryMarginAccountsAllOco) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**QueryMarginAccountsAllOrders**](TradeAPI.md#QueryMarginAccountsAllOrders) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**QueryMarginAccountsOco**](TradeAPI.md#QueryMarginAccountsOco) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**QueryMarginAccountsOpenOco**](TradeAPI.md#QueryMarginAccountsOpenOco) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**QueryMarginAccountsOpenOrders**](TradeAPI.md#QueryMarginAccountsOpenOrders) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**QueryMarginAccountsOrder**](TradeAPI.md#QueryMarginAccountsOrder) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**QueryMarginAccountsTradeList**](TradeAPI.md#QueryMarginAccountsTradeList) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**QuerySpecialKey**](TradeAPI.md#QuerySpecialKey) | **Get** /sapi/v1/margin/apiKey | Query Special key(Low Latency Trading)(TRADE)
[**QuerySpecialKeyList**](TradeAPI.md#QuerySpecialKeyList) | **Get** /sapi/v1/margin/api-key-list | Query Special key List(Low Latency Trading)(TRADE)
[**SmallLiabilityExchange**](TradeAPI.md#SmallLiabilityExchange) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)


## CreateSpecialKey

> CreateSpecialKeyResponse CreateSpecialKey(ctx).ApiName(apiName).Symbol(symbol).Ip(ip).PublicKey(publicKey).PermissionMode(permissionMode).RecvWindow(recvWindow).Execute()

Create Special Key(Low-Latency Trading)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	apiName := "apiName_example" // string | 
	symbol := "symbol_example" // string | isolated margin pair (optional)
	ip := "ip_example" // string | Can be added in batches, separated by commas. Max 30 for an API key (optional)
	publicKey := "publicKey_example" // string | 1. If publicKey is inputted it will create an RSA or Ed25519 key. <br />2. Need to be encoded to URL-encoded format (optional)
	permissionMode := "value" // string | This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE. (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CreateSpecialKey(context.Background()).ApiName(apiName).Symbol(symbol).Ip(ip).PublicKey(publicKey).PermissionMode(permissionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CreateSpecialKey``: %v\n", err)
		return
	}

	// response from `CreateSpecialKey`: CreateSpecialKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | 
 **symbol** | **string** | isolated margin pair | 
 **ip** | **string** | Can be added in batches, separated by commas. Max 30 for an API key | 
 **publicKey** | **string** | 1. If publicKey is inputted it will create an RSA or Ed25519 key. &lt;br /&gt;2. Need to be encoded to URL-encoded format | 
 **permissionMode** | **string** | This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**CreateSpecialKeyResponse**](CreateSpecialKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteSpecialKey

> DeleteSpecialKey(ctx).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()

Delete Special Key(Low-Latency Trading)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	apiName := "apiName_example" // string |  (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.DeleteSpecialKey(context.Background()).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.DeleteSpecialKey``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## EditIpForSpecialKey

> EditIpForSpecialKey(ctx).Ip(ip).Symbol(symbol).RecvWindow(recvWindow).Execute()

Edit ip for Special Key(Low-Latency Trading)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ip := "ip_example" // string | Can be added in batches, separated by commas. Max 30 for an API key
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Ip(ip).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.EditIpForSpecialKey``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | Can be added in batches, separated by commas. Max 30 for an API key | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)


## GetForceLiquidationRecord

> GetForceLiquidationRecordResponse GetForceLiquidationRecord(ctx).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional)
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(10) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetForceLiquidationRecord(context.Background()).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetForceLiquidationRecord``: %v\n", err)
		return
	}

	// response from `GetForceLiquidationRecord`: GetForceLiquidationRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetForceLiquidationRecordResponse**](GetForceLiquidationRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSmallLiabilityExchangeCoinList

> GetSmallLiabilityExchangeCoinListResponse GetSmallLiabilityExchangeCoinList(ctx).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange Coin List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeCoinList(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetSmallLiabilityExchangeCoinList``: %v\n", err)
		return
	}

	// response from `GetSmallLiabilityExchangeCoinList`: GetSmallLiabilityExchangeCoinListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetSmallLiabilityExchangeCoinListResponse**](GetSmallLiabilityExchangeCoinListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetSmallLiabilityExchangeHistory

> GetSmallLiabilityExchangeHistoryResponse GetSmallLiabilityExchangeHistory(ctx).Current(current).Size(size).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	current := int64(1) // int64 | Currently querying page. Start from 1. Default:1
	size := int64(10) // int64 | Default:10, Max:100
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeHistory(context.Background()).Current(current).Size(size).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetSmallLiabilityExchangeHistory``: %v\n", err)
		return
	}

	// response from `GetSmallLiabilityExchangeHistory`: GetSmallLiabilityExchangeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetSmallLiabilityExchangeHistoryResponse**](GetSmallLiabilityExchangeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelAllOpenOrdersOnASymbol

> MarginAccountCancelAllOpenOrdersOnASymbolResponse MarginAccountCancelAllOpenOrdersOnASymbol(ctx).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol(context.Background()).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelAllOpenOrdersOnASymbol`: MarginAccountCancelAllOpenOrdersOnASymbolResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountCancelAllOpenOrdersOnASymbolResponse**](MarginAccountCancelAllOpenOrdersOnASymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelOco

> MarginAccountCancelOcoResponse MarginAccountCancelOco(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOco(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelOco``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelOco`: MarginAccountCancelOcoResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountCancelOcoResponse**](MarginAccountCancelOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountCancelOrder

> MarginAccountCancelOrderResponse MarginAccountCancelOrder(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOrder(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountCancelOrder``: %v\n", err)
		return
	}

	// response from `MarginAccountCancelOrder`: MarginAccountCancelOrderResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountCancelOrderResponse**](MarginAccountCancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOco

> MarginAccountNewOcoResponse MarginAccountNewOco(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()

Margin Account New OCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	quantity := float32(1.0) // float32 | 
	price := float32(1.0) // float32 | 
	stopPrice := float32(1.0) // float32 | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	limitClientOrderId := "1" // string | A unique Id for the limit order (optional)
	limitIcebergQty := float32(1.0) // float32 |  (optional)
	stopClientOrderId := "1" // string | A unique Id for the stop loss/stop loss limit leg (optional)
	stopLimitPrice := float32(1.0) // float32 | If provided, `stopLimitTimeInForce` is required. (optional)
	stopIcebergQty := float32(1.0) // float32 |  (optional)
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string | Valid values are `GTC`/`FOK`/`IOC` (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	sideEffectType := "NO_SIDE_EFFECT" // string | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) (optional)
	selfTradePreventionMode := "NONE" // string | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOco``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOco`: MarginAccountNewOcoResponse
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
 **side** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **quantity** | **float32** |  | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** |  | 
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitIcebergQty** | **float32** |  | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopLimitPrice** | **float32** | If provided, &#x60;stopLimitTimeInForce&#x60; is required. | 
 **stopIcebergQty** | **float32** |  | 
 **stopLimitTimeInForce** | **string** | Valid values are &#x60;GTC&#x60;/&#x60;FOK&#x60;/&#x60;IOC&#x60; | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | **string** | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) | 
 **selfTradePreventionMode** | **string** | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountNewOcoResponse**](MarginAccountNewOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOrder

> MarginAccountNewOrderResponse MarginAccountNewOrder(ctx).Symbol(symbol).Side(side).Type(type_).IsIsolated(isIsolated).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()

Margin Account New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.MarginAccountNewOrderSideParameterBuy // MarginAccountNewOrderSideParameter | 
	type_ := "type__example" // string | `MARGIN`,`ISOLATED`
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	price := float32(1.0) // float32 |  (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	newClientOrderId := "1" // string | Used to uniquely identify this cancel. Automatically generated by default (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	sideEffectType := "NO_SIDE_EFFECT" // string | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) (optional)
	timeInForce := models.MarginAccountNewOrderTimeInForceParameterGtc // MarginAccountNewOrderTimeInForceParameter | GTC,IOC,FOK (optional)
	selfTradePreventionMode := "NONE" // string | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).IsIsolated(isIsolated).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).StopPrice(stopPrice).NewClientOrderId(newClientOrderId).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).TimeInForce(timeInForce).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOrder``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOrder`: MarginAccountNewOrderResponse
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
 **side** | [**MarginAccountNewOrderSideParameter**](MarginAccountNewOrderSideParameter.md) |  | 
 **type_** | **string** | &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60; | 
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | **string** | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) | 
 **timeInForce** | [**MarginAccountNewOrderTimeInForceParameter**](MarginAccountNewOrderTimeInForceParameter.md) | GTC,IOC,FOK | 
 **selfTradePreventionMode** | **string** | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginAccountNewOrderResponse**](MarginAccountNewOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOto

> MarginAccountNewOtoResponse MarginAccountNewOto(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingIcebergQty(workingIcebergQty).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).Execute()

Margin Account New OTO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	workingType := "workingType_example" // string | Supported values: `LIMIT`, `LIMIT_MAKER`
	workingSide := "workingSide_example" // string | BUY, SELL
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | 
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`.
	pendingType := "Order Types" // string | Supported values: [Order Types](https://developers.binance.com/docs/binance-spot-api-docs/enums#order-types-ordertypes-type) Note that `MARKET` orders using `quoteOrderQty` are not supported.
	pendingSide := "pendingSide_example" // string | BUY, SELL
	pendingQuantity := float32(1.0) // float32 | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	sideEffectType := "NO_SIDE_EFFECT" // string | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) (optional)
	selfTradePreventionMode := "NONE" // string | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true (optional)
	workingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingTimeInForce := "workingTimeInForce_example" // string | GTC,IOC,FOK (optional)
	pendingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. (optional)
	pendingPrice := float32(1.0) // float32 |  (optional)
	pendingStopPrice := float32(1.0) // float32 |  (optional)
	pendingTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingIcebergQty := float32(1.0) // float32 | This can only be used if `pendingTimeInForce` is `GTC`. (optional)
	pendingTimeInForce := "pendingTimeInForce_example" // string | GTC,IOC,FOK (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOto(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingIcebergQty(workingIcebergQty).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SideEffectType(sideEffectType).SelfTradePreventionMode(selfTradePreventionMode).AutoRepayAtCancel(autoRepayAtCancel).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOto``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOto`: MarginAccountNewOtoResponse
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
 **workingType** | **string** | Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60; | 
 **workingSide** | **string** | BUY, SELL | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** |  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingType** | **string** | Supported values: [Order Types](https://developers.binance.com/docs/binance-spot-api-docs/enums#order-types-ordertypes-type) Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported. | 
 **pendingSide** | **string** | BUY, SELL | 
 **pendingQuantity** | **float32** |  | 
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **sideEffectType** | **string** | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) | 
 **selfTradePreventionMode** | **string** | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingTimeInForce** | **string** | GTC,IOC,FOK | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. | 
 **pendingPrice** | **float32** |  | 
 **pendingStopPrice** | **float32** |  | 
 **pendingTrailingDelta** | **float32** |  | 
 **pendingIcebergQty** | **float32** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingTimeInForce** | **string** | GTC,IOC,FOK | 

### Return type

[**MarginAccountNewOtoResponse**](MarginAccountNewOtoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginAccountNewOtoco

> MarginAccountNewOtocoResponse MarginAccountNewOtoco(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).IsIsolated(isIsolated).SideEffectType(sideEffectType).AutoRepayAtCancel(autoRepayAtCancel).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).Execute()

Margin Account New OTOCO (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	workingType := "workingType_example" // string | Supported values: `LIMIT`, `LIMIT_MAKER`
	workingSide := "workingSide_example" // string | BUY, SELL
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | 
	pendingSide := "pendingSide_example" // string | BUY, SELL
	pendingQuantity := float32(1.0) // float32 | 
	pendingAboveType := "pendingAboveType_example" // string | Supported values: `LIMIT_MAKER`, `STOP_LOSS`, and `STOP_LOSS_LIMIT`
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	sideEffectType := "NO_SIDE_EFFECT" // string | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) (optional)
	autoRepayAtCancel := true // bool | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true (optional)
	listClientOrderId := "1" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	newOrderRespType := models.MarginAccountNewOrderNewOrderRespTypeParameterAck // MarginAccountNewOrderNewOrderRespTypeParameter | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. (optional)
	selfTradePreventionMode := "NONE" // string | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE (optional)
	workingClientOrderId := "1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`. (optional)
	workingTimeInForce := "workingTimeInForce_example" // string | GTC,IOC,FOK (optional)
	pendingAboveClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. (optional)
	pendingAbovePrice := float32(1.0) // float32 |  (optional)
	pendingAboveStopPrice := float32(1.0) // float32 |  (optional)
	pendingAboveTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingAboveIcebergQty := float32(1.0) // float32 | This can only be used if `pendingAboveTimeInForce` is `GTC`. (optional)
	pendingAboveTimeInForce := "pendingAboveTimeInForce_example" // string |  (optional)
	pendingBelowType := "pendingBelowType_example" // string | Supported values: `LIMIT_MAKER`, `STOP_LOSS`, and `STOP_LOSS_LIMIT` (optional)
	pendingBelowClientOrderId := "1" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. (optional)
	pendingBelowPrice := float32(1.0) // float32 |  (optional)
	pendingBelowStopPrice := float32(1.0) // float32 |  (optional)
	pendingBelowTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingBelowIcebergQty := float32(1.0) // float32 | This can only be used if `pendingBelowTimeInForce` is `GTC`. (optional)
	pendingBelowTimeInForce := "pendingBelowTimeInForce_example" // string |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).IsIsolated(isIsolated).SideEffectType(sideEffectType).AutoRepayAtCancel(autoRepayAtCancel).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginAccountNewOtoco``: %v\n", err)
		return
	}

	// response from `MarginAccountNewOtoco`: MarginAccountNewOtocoResponse
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
 **workingType** | **string** | Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60; | 
 **workingSide** | **string** | BUY, SELL | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** |  | 
 **pendingSide** | **string** | BUY, SELL | 
 **pendingQuantity** | **float32** |  | 
 **pendingAboveType** | **string** | Supported values: &#x60;LIMIT_MAKER&#x60;, &#x60;STOP_LOSS&#x60;, and &#x60;STOP_LOSS_LIMIT&#x60; | 
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **sideEffectType** | **string** | NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc) | 
 **autoRepayAtCancel** | **bool** | Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **newOrderRespType** | [**MarginAccountNewOrderNewOrderRespTypeParameter**](MarginAccountNewOrderNewOrderRespTypeParameter.md) | Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK. | 
 **selfTradePreventionMode** | **string** | The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **workingTimeInForce** | **string** | GTC,IOC,FOK | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. | 
 **pendingAbovePrice** | **float32** |  | 
 **pendingAboveStopPrice** | **float32** |  | 
 **pendingAboveTrailingDelta** | **float32** |  | 
 **pendingAboveIcebergQty** | **float32** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingAboveTimeInForce** | **string** |  | 
 **pendingBelowType** | **string** | Supported values: &#x60;LIMIT_MAKER&#x60;, &#x60;STOP_LOSS&#x60;, and &#x60;STOP_LOSS_LIMIT&#x60; | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. | 
 **pendingBelowPrice** | **float32** |  | 
 **pendingBelowStopPrice** | **float32** |  | 
 **pendingBelowTrailingDelta** | **float32** |  | 
 **pendingBelowIcebergQty** | **float32** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **pendingBelowTimeInForce** | **string** |  | 

### Return type

[**MarginAccountNewOtocoResponse**](MarginAccountNewOtocoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## MarginManualLiquidation

> MarginManualLiquidationResponse MarginManualLiquidation(ctx).Type(type_).Symbol(symbol).RecvWindow(recvWindow).Execute()

Margin Manual Liquidation(MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	type_ := "type__example" // string | `MARGIN`,`ISOLATED`
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.MarginManualLiquidation(context.Background()).Type(type_).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.MarginManualLiquidation``: %v\n", err)
		return
	}

	// response from `MarginManualLiquidation`: MarginManualLiquidationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60; | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginManualLiquidationResponse**](MarginManualLiquidationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentMarginOrderCountUsage

> QueryCurrentMarginOrderCountUsageResponse QueryCurrentMarginOrderCountUsage(ctx).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Order Count Usage (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentMarginOrderCountUsage(context.Background()).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentMarginOrderCountUsage``: %v\n", err)
		return
	}

	// response from `QueryCurrentMarginOrderCountUsage`: QueryCurrentMarginOrderCountUsageResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryCurrentMarginOrderCountUsageResponse**](QueryCurrentMarginOrderCountUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsAllOco

> QueryMarginAccountsAllOcoResponse QueryMarginAccountsAllOco(ctx).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	fromId := int64(1) // int64 | If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned. (optional)
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(500) // int64 | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsAllOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsAllOco`: QueryMarginAccountsAllOcoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **symbol** | **string** | isolated margin pair | 
 **fromId** | **int64** | If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned. | 
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsAllOcoResponse**](QueryMarginAccountsAllOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsAllOrders

> QueryMarginAccountsAllOrdersResponse QueryMarginAccountsAllOrders(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	limit := int64(500) // int64 | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOrders(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsAllOrders``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsAllOrders`: QueryMarginAccountsAllOrdersResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsAllOrdersResponse**](QueryMarginAccountsAllOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOco

> QueryMarginAccountsOcoResponse QueryMarginAccountsOco(ctx).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOco`: QueryMarginAccountsOcoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **symbol** | **string** | isolated margin pair | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsOcoResponse**](QueryMarginAccountsOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOpenOco

> QueryMarginAccountsOpenOcoResponse QueryMarginAccountsOpenOco(ctx).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOco(context.Background()).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOpenOco``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOpenOco`: QueryMarginAccountsOpenOcoResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsOpenOcoResponse**](QueryMarginAccountsOpenOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOpenOrders

> QueryMarginAccountsOpenOrdersResponse QueryMarginAccountsOpenOrders(ctx).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | isolated margin pair (optional)
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOrders(context.Background()).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOpenOrders``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOpenOrders`: QueryMarginAccountsOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsOpenOrdersResponse**](QueryMarginAccountsOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsOrder

> QueryMarginAccountsOrderResponse QueryMarginAccountsOrder(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOrder(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsOrder``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsOrder`: QueryMarginAccountsOrderResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsOrderResponse**](QueryMarginAccountsOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryMarginAccountsTradeList

> QueryMarginAccountsTradeListResponse QueryMarginAccountsTradeList(ctx).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | 
	isIsolated := "false" // string | for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\" (optional)
	orderId := int64(1) // int64 |  (optional)
	startTime := int64(1623319461670) // int64 | Only supports querying data from the past 90 days. (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	fromId := int64(1) // int64 | If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned. (optional)
	limit := int64(500) // int64 | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsTradeList(context.Background()).Symbol(symbol).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryMarginAccountsTradeList``: %v\n", err)
		return
	}

	// response from `QueryMarginAccountsTradeList`: QueryMarginAccountsTradeListResponse
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
 **isIsolated** | **string** | for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot; | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** | Only supports querying data from the past 90 days. | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned. | 
 **limit** | **int64** | Limit on the number of data records returned per request. Default: 500; Maximum: 1000. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QueryMarginAccountsTradeListResponse**](QueryMarginAccountsTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySpecialKey

> QuerySpecialKeyResponse QuerySpecialKey(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key(Low Latency Trading)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKey(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QuerySpecialKey``: %v\n", err)
		return
	}

	// response from `QuerySpecialKey`: QuerySpecialKeyResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QuerySpecialKeyResponse**](QuerySpecialKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySpecialKeyList

> QuerySpecialKeyListResponse QuerySpecialKeyList(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key List(Low Latency Trading)(TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "symbol_example" // string | isolated margin pair (optional)
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKeyList(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QuerySpecialKeyList``: %v\n", err)
		return
	}

	// response from `QuerySpecialKeyList`: QuerySpecialKeyListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | isolated margin pair | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**QuerySpecialKeyListResponse**](QuerySpecialKeyListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SmallLiabilityExchange

> SmallLiabilityExchange(ctx).AssetNames(assetNames).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	assetNames := []string{"BTC"} // []string | The assets list of small liability exchange， Example: assetNames = BTC,ETH
	recvWindow := int64(5000) // int64 | No more than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceMarginTradingClient(models.WithRestAPI(configuration))

	 err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).AssetNames(assetNames).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SmallLiabilityExchange``: %v\n", err)
		return
	}

}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **assetNames** | **[]string** | The assets list of small liability exchange， Example: assetNames &#x3D; BTC,ETH | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: Not defined

[[Back to README]](../../../README.md)

