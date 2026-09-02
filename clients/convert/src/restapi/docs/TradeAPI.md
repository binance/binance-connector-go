# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AcceptQuote**](TradeAPI.md#AcceptQuote) | **Post** /sapi/v1/convert/acceptQuote | Accept Quote (TRADE)
[**CancelLimitOrder**](TradeAPI.md#CancelLimitOrder) | **Post** /sapi/v1/convert/limit/cancelOrder | Cancel limit order (TRADE)
[**GetConvertTradeHistory**](TradeAPI.md#GetConvertTradeHistory) | **Get** /sapi/v1/convert/tradeFlow | Get Convert Trade History (USER_DATA)
[**OrderStatus**](TradeAPI.md#OrderStatus) | **Get** /sapi/v1/convert/orderStatus | Order status (USER_DATA)
[**PlaceLimitOrder**](TradeAPI.md#PlaceLimitOrder) | **Post** /sapi/v1/convert/limit/placeOrder | Place limit order (TRADE)
[**QueryLimitOpenOrders**](TradeAPI.md#QueryLimitOpenOrders) | **Get** /sapi/v1/convert/limit/queryOpenOrders | Query limit open orders (USER_DATA)
[**SendQuoteRequest**](TradeAPI.md#SendQuoteRequest) | **Post** /sapi/v1/convert/getQuote | Send Quote Request (TRADE)


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
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**AcceptQuoteResponse**](AcceptQuoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelLimitOrder

> CancelLimitOrderResponse CancelLimitOrder(ctx).OrderId(orderId).RecvWindow(recvWindow).Execute()

Cancel limit order (TRADE)


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
	orderId := int64(1603680255057330400) // int64 | The orderId from `placeOrder` api
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**CancelLimitOrderResponse**](CancelLimitOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## GetConvertTradeHistory

> GetConvertTradeHistoryResponse GetConvertTradeHistory(ctx).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Convert Trade History (USER_DATA)


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
	limit := int64(100) // int64 | Number of records to return (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **limit** | **int64** | Number of records to return | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**GetConvertTradeHistoryResponse**](GetConvertTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderStatus

> OrderStatusResponse OrderStatus(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status (USER_DATA)


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

Place limit order (TRADE)


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
	baseAsset := "BTC" // string | base asset (use the response `fromIsBase` from `GET /sapi/v1/convert/exchangeInfo` api to check which one is baseAsset )
	quoteAsset := "USDT" // string | quote asset
	limitPrice := float64(1) // float64 | Symbol limit price (from baseAsset to quoteAsset)
	side := models.PlaceLimitOrderSideParameterBuy // PlaceLimitOrderSideParameter | `BUY` or `SELL`
	expiredType := models.PlaceLimitOrderExpiredTypeParameterExpiredType1D // PlaceLimitOrderExpiredTypeParameter | Order expiry duration. 1_D, 3_D, 7_D, 30_D (D means day)
	baseAmount := float64(1) // float64 | Base asset amount. (One of `baseAmount` or `quoteAmount` is required) (optional)
	quoteAmount := float64(1) // float64 | Quote asset amount. (One of `baseAmount` or `quoteAmount` is required) (optional)
	walletType := models.PlaceLimitOrderWalletTypeParameterSpot // PlaceLimitOrderWalletTypeParameter | Wallet to use for payment. Supported values: `SPOT`, `FUNDING`, `EARN`. Combined wallets also supported: `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN`, `SPOT_EARN`. Default is `SPOT`. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **limitPrice** | **float64** | Symbol limit price (from baseAsset to quoteAsset) | 
 **side** | [**PlaceLimitOrderSideParameter**](PlaceLimitOrderSideParameter.md) | &#x60;BUY&#x60; or &#x60;SELL&#x60; | 
 **expiredType** | [**PlaceLimitOrderExpiredTypeParameter**](PlaceLimitOrderExpiredTypeParameter.md) | Order expiry duration. 1_D, 3_D, 7_D, 30_D (D means day) | 
 **baseAmount** | **float64** | Base asset amount. (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required) | 
 **quoteAmount** | **float64** | Quote asset amount. (One of &#x60;baseAmount&#x60; or &#x60;quoteAmount&#x60; is required) | 
 **walletType** | [**PlaceLimitOrderWalletTypeParameter**](PlaceLimitOrderWalletTypeParameter.md) | Wallet to use for payment. Supported values: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;, &#x60;EARN&#x60;. Combined wallets also supported: &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60;, &#x60;SPOT_EARN&#x60;. Default is &#x60;SPOT&#x60;. | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

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
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**QueryLimitOpenOrdersResponse**](QueryLimitOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SendQuoteRequest

> SendQuoteRequestResponse SendQuoteRequest(ctx).FromAsset(fromAsset).ToAsset(toAsset).FromAmount(fromAmount).ToAmount(toAmount).WalletType(walletType).ValidTime(validTime).RecvWindow(recvWindow).Execute()

Send Quote Request (TRADE)


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
	fromAsset := "BTC" // string | 
	toAsset := "USDT" // string | 
	fromAmount := float64(1) // float64 | When specified, it is the amount you will be debited after the conversion (optional)
	toAmount := float64(1) // float64 | When specified, it is the amount you will be credited after the conversion (optional)
	walletType := models.PlaceLimitOrderWalletTypeParameterSpot // PlaceLimitOrderWalletTypeParameter | Wallet to use for payment. Supported values: `SPOT`, `FUNDING`, `EARN`. Combined wallets also supported: `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN`, `SPOT_EARN`. Default is `SPOT`. (optional)
	validTime := models.SendQuoteRequestValidTimeParameterValidTime10s // SendQuoteRequestValidTimeParameter | Quote valid duration. Supported values: 10s, 30s, 1m. Default is 10s. (optional)
	recvWindow := int64(5000) // int64 | Request validity window in milliseconds (optional)

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
 **fromAmount** | **float64** | When specified, it is the amount you will be debited after the conversion | 
 **toAmount** | **float64** | When specified, it is the amount you will be credited after the conversion | 
 **walletType** | [**PlaceLimitOrderWalletTypeParameter**](PlaceLimitOrderWalletTypeParameter.md) | Wallet to use for payment. Supported values: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;, &#x60;EARN&#x60;. Combined wallets also supported: &#x60;SPOT_FUNDING&#x60;, &#x60;FUNDING_EARN&#x60;, &#x60;SPOT_FUNDING_EARN&#x60;, &#x60;SPOT_EARN&#x60;. Default is &#x60;SPOT&#x60;. | 
 **validTime** | [**SendQuoteRequestValidTimeParameter**](SendQuoteRequestValidTimeParameter.md) | Quote valid duration. Supported values: 10s, 30s, 1m. Default is 10s. | 
 **recvWindow** | **int64** | Request validity window in milliseconds | 

### Return type

[**SendQuoteRequestResponse**](SendQuoteRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

