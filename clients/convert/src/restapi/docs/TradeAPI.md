# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AcceptQuote**](TradeAPI.md#AcceptQuote) | **Post** /sapi/v1/convert/acceptQuote | Accept Quote (TRADE)
[**CancelLimitOrder**](TradeAPI.md#CancelLimitOrder) | **Post** /sapi/v1/convert/limit/cancelOrder | Cancel limit order (USER_DATA)
[**GetConvertTradeHistory**](TradeAPI.md#GetConvertTradeHistory) | **Get** /sapi/v1/convert/tradeFlow | Get Convert Trade History(USER_DATA)
[**OrderStatus**](TradeAPI.md#OrderStatus) | **Get** /sapi/v1/convert/orderStatus | Order status(USER_DATA)
[**PlaceLimitOrder**](TradeAPI.md#PlaceLimitOrder) | **Post** /sapi/v1/convert/limit/placeOrder | Place limit order (USER_DATA)
[**QueryLimitOpenOrders**](TradeAPI.md#QueryLimitOpenOrders) | **Get** /sapi/v1/convert/limit/queryOpenOrders | Query limit open orders (USER_DATA)
[**SendQuoteRequest**](TradeAPI.md#SendQuoteRequest) | **Post** /sapi/v1/convert/getQuote | Send Quote Request(USER_DATA)


## AcceptQuote

> AcceptQuoteResponse AcceptQuote(ctx).QuoteId(quoteId).RecvWindow(recvWindow).Execute()

Accept Quote (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	quoteId := "1" // string | 
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AcceptQuote(context.Background()).QuoteId(quoteId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.AcceptQuote``: %v\n", err)
		return
	}

	// response from `AcceptQuote`: AcceptQuoteResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**AcceptQuoteResponse**](AcceptQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelLimitOrder

> CancelLimitOrderResponse CancelLimitOrder(ctx).OrderId(orderId).RecvWindow(recvWindow).Execute()

Cancel limit order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := int64(1) // int64 | The orderId from `placeOrder` api
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelLimitOrder(context.Background()).OrderId(orderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelLimitOrder``: %v\n", err)
		return
	}

	// response from `CancelLimitOrder`: CancelLimitOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **int64** | The orderId from &#x60;placeOrder&#x60; api | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**CancelLimitOrderResponse**](CancelLimitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetConvertTradeHistory

> GetConvertTradeHistoryResponse GetConvertTradeHistory(ctx).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Convert Trade History(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1623319461670) // int64 | 
	endTime := int64(1641782889000) // int64 | 
	limit := int64(100) // int64 | Default 100, Max 1000 (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.GetConvertTradeHistory(context.Background()).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.GetConvertTradeHistory``: %v\n", err)
		return
	}

	// response from `GetConvertTradeHistory`: GetConvertTradeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int64** | Default 100, Max 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetConvertTradeHistoryResponse**](GetConvertTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderStatus

> OrderStatusResponse OrderStatus(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := "1" // string | Either orderId or quoteId is required (optional)
	quoteId := "1" // string | Either orderId or quoteId is required (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderStatus(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderStatus``: %v\n", err)
		return
	}

	// response from `OrderStatus`: OrderStatusResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | 
 **quoteId** | **string** | Either orderId or quoteId is required | 

### Return type

[**OrderStatusResponse**](OrderStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PlaceLimitOrder

> PlaceLimitOrderResponse PlaceLimitOrder(ctx).BaseAsset(baseAsset).QuoteAsset(quoteAsset).LimitPrice(limitPrice).Side(side).ExpiredType(expiredType).BaseAmount(baseAmount).QuoteAmount(quoteAmount).WalletType(walletType).RecvWindow(recvWindow).Execute()

Place limit order (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	baseAsset := "baseAsset_example" // string | base asset (use the response `fromIsBase` from `GET /sapi/v1/convert/exchangeInfo` api to check which one is baseAsset )
	quoteAsset := "quoteAsset_example" // string | quote asset
	limitPrice := float32(1.0) // float32 | Symbol limit price (from baseAsset to quoteAsset)
	side := "BUY" // string | `BUY` or `SELL`
	expiredType := "expiredType_example" // string | 1_D, 3_D, 7_D, 30_D  (D means day)
	baseAmount := float32(1.0) // float32 | Base asset amount.  (One of `baseAmount` or `quoteAmount` is required) (optional)
	quoteAmount := float32(1.0) // float32 | Quote asset amount.  (One of `baseAmount` or `quoteAmount` is required) (optional)
	walletType := "walletType_example" // string | It is to choose which wallet of assets. The wallet selection is `SPOT`, `FUNDING` and `EARN`. Combination of wallet is supported i.e. `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN` or `SPOT_EARN`  Default is `SPOT`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PlaceLimitOrder(context.Background()).BaseAsset(baseAsset).QuoteAsset(quoteAsset).LimitPrice(limitPrice).Side(side).ExpiredType(expiredType).BaseAmount(baseAmount).QuoteAmount(quoteAmount).WalletType(walletType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PlaceLimitOrder``: %v\n", err)
		return
	}

	// response from `PlaceLimitOrder`: PlaceLimitOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **baseAsset** | **string** | base asset (use the response &#x60;fromIsBase&#x60; from &#x60;GET /sapi/v1/convert/exchangeInfo&#x60; api to check which one is baseAsset ) | 
 **quoteAsset** | **string** | quote asset | 
 **limitPrice** | **float32** | Symbol limit price (from baseAsset to quoteAsset) | 
 **side** | **string** | &#x60;BUY&#x60; or &#x60;SELL&#x60; | 
 **expiredType** | **string** | 1_D, 3_D, 7_D, 30_D  (D means day) | 
 **baseAmount** | **float32** | Base asset amount.  (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required) | 
 **quoteAmount** | **float32** | Quote asset amount.  (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required) | 
 **walletType** | **string** | It is to choose which wallet of assets. The wallet selection is &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; and &#x60;EARN&#x60;. Combination of wallet is supported i.e. &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60; or &#x60;SPOT_EARN&#x60;  Default is &#x60;SPOT&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PlaceLimitOrderResponse**](PlaceLimitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryLimitOpenOrders

> QueryLimitOpenOrdersResponse QueryLimitOpenOrders(ctx).RecvWindow(recvWindow).Execute()

Query limit open orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryLimitOpenOrders(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryLimitOpenOrders``: %v\n", err)
		return
	}

	// response from `QueryLimitOpenOrders`: QueryLimitOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**QueryLimitOpenOrdersResponse**](QueryLimitOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SendQuoteRequest

> SendQuoteRequestResponse SendQuoteRequest(ctx).FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).WalletType(walletType).ValidTime(validTime).RecvWindow(recvWindow).Execute()

Send Quote Request(USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/convert"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	fromAsset := "fromAsset_example" // string | 
	toAsset := "toAsset_example" // string | 
	fromAmount := float32(1.0) // float32 | When specified, it is the amount you will be debited after the conversion (optional)
	toAmount := float32(1.0) // float32 | When specified, it is the amount you will be credited after the conversion (optional)
	walletType := "walletType_example" // string | It is to choose which wallet of assets. The wallet selection is `SPOT`, `FUNDING` and `EARN`. Combination of wallet is supported i.e. `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN` or `SPOT_EARN`  Default is `SPOT`. (optional)
	validTime := "10s" // string | 10s, 30s, 1m, default 10s (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than 60000 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceConvertClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.SendQuoteRequest(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).WalletType(walletType).ValidTime(validTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SendQuoteRequest``: %v\n", err)
		return
	}

	// response from `SendQuoteRequest`: SendQuoteRequestResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** |  | 
 **toAsset** | **string** |  | 
 **fromAmount** | **float32** | When specified, it is the amount you will be debited after the conversion | 
 **toAmount** | **float32** | When specified, it is the amount you will be credited after the conversion | 
 **walletType** | **string** | It is to choose which wallet of assets. The wallet selection is &#x60;SPOT&#x60;, &#x60;FUNDING&#x60; and &#x60;EARN&#x60;. Combination of wallet is supported i.e. &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60; or &#x60;SPOT_EARN&#x60;  Default is &#x60;SPOT&#x60;. | 
 **validTime** | **string** | 10s, 30s, 1m, default 10s | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**SendQuoteRequestResponse**](SendQuoteRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

