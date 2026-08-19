# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAllEquityOrders**](TradeAPI.md#CancelAllEquityOrders) | **Post** /sapi/v1/equity/order/cancel-all | Cancel All Equity Orders (TRADE)
[**CancelEquityOrder**](TradeAPI.md#CancelEquityOrder) | **Post** /sapi/v1/equity/order/cancel | Cancel Equity Order (TRADE)
[**CurrentOpenOrders**](TradeAPI.md#CurrentOpenOrders) | **Get** /sapi/v1/equity/order/open-orders | Current Open Orders (USER_DATA)
[**EquityOrderDetail**](TradeAPI.md#EquityOrderDetail) | **Get** /sapi/v1/equity/order/detail | Equity Order Detail (USER_DATA)
[**EquityOrderHistory**](TradeAPI.md#EquityOrderHistory) | **Get** /sapi/v1/equity/order/history | Equity Order History (USER_DATA)
[**EquityTradeHistory**](TradeAPI.md#EquityTradeHistory) | **Get** /sapi/v1/equity/trade/history | Equity Trade History (USER_DATA)
[**PlaceEquityOrder**](TradeAPI.md#PlaceEquityOrder) | **Post** /sapi/v1/equity/order/place | Place Equity Order (TRADE)


## CancelAllEquityOrders

> CancelAllEquityOrdersResponse CancelAllEquityOrders(ctx).RecvWindow(recvWindow).Execute()

Cancel All Equity Orders (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllEquityOrders(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllEquityOrders``: %v\n", err)
		return
	}

	// response from `CancelAllEquityOrders`: CancelAllEquityOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**CancelAllEquityOrdersResponse**](CancelAllEquityOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelEquityOrder

> CancelEquityOrderResponse CancelEquityOrder(ctx).OrderId(orderId).RecvWindow(recvWindow).Execute()

Cancel Equity Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := "c3c58f49-7b0d-4b9e-a2db-1a2f9a3b8c71" // string | Equity order id returned by `/order/place` or a query endpoint.
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelEquityOrder(context.Background()).OrderId(orderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelEquityOrder``: %v\n", err)
		return
	}

	// response from `CancelEquityOrder`: CancelEquityOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Equity order id returned by &#x60;/order/place&#x60; or a query endpoint. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**CancelEquityOrderResponse**](CancelEquityOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CurrentOpenOrders

> CurrentOpenOrdersResponse CurrentOpenOrders(ctx).RecvWindow(recvWindow).Execute()

Current Open Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CurrentOpenOrders(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CurrentOpenOrders``: %v\n", err)
		return
	}

	// response from `CurrentOpenOrders`: CurrentOpenOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**CurrentOpenOrdersResponse**](CurrentOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EquityOrderDetail

> EquityOrderDetailResponse EquityOrderDetail(ctx).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Equity Order Detail (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orderId := "c3c58f49-7b0d-4b9e-a2db-1a2f9a3b8c71" // string | Equity order id. Either `orderId` or `clientOrderId` must be provided. (optional)
	clientOrderId := "web_2c9c92b74f1e4a7c8f3b9e1a2d3c4b5a" // string | Client-supplied order id. Either `orderId` or `clientOrderId` must be provided. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.EquityOrderDetail(context.Background()).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.EquityOrderDetail``: %v\n", err)
		return
	}

	// response from `EquityOrderDetail`: EquityOrderDetailResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Equity order id. Either &#x60;orderId&#x60; or &#x60;clientOrderId&#x60; must be provided. | 
 **clientOrderId** | **string** | Client-supplied order id. Either &#x60;orderId&#x60; or &#x60;clientOrderId&#x60; must be provided. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**EquityOrderDetailResponse**](EquityOrderDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EquityOrderHistory

> EquityOrderHistoryResponse EquityOrderHistory(ctx).StartTime(startTime).EndTime(endTime).Symbol(symbol).OrderType(orderType).Side(side).OrderStatus(orderStatus).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Equity Order History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1735800000000) // int64 | Start time (ms epoch).
	endTime := int64(1735900000000) // int64 | End time (ms epoch).
	symbol := "NVDA" // string | US-equity ticker filter, e.g. `NVDA`. (optional)
	orderType := models.PlaceEquityOrderOrderTypeParameterMarket // PlaceEquityOrderOrderTypeParameter | Order type filter: `MARKET` / `LIMIT`. (optional)
	side := models.PlaceEquityOrderSideParameterBuy // PlaceEquityOrderSideParameter | Side filter: `BUY` / `SELL`. (optional)
	orderStatus := "FILLED,CANCELED" // string | Comma-separated status filter. Allowed values: `FILLED`, `PARTIALLY_FILLED`, `CANCELED`, `EXPIRED`, `REJECTED`. (optional)
	current := int32(1) // int32 | Page number, 1-based. Default `1`. (optional)
	size := int32(20) // int32 | Page size. Default `20`, max `100`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.EquityOrderHistory(context.Background()).StartTime(startTime).EndTime(endTime).Symbol(symbol).OrderType(orderType).Side(side).OrderStatus(orderStatus).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.EquityOrderHistory``: %v\n", err)
		return
	}

	// response from `EquityOrderHistory`: EquityOrderHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start time (ms epoch). | 
 **endTime** | **int64** | End time (ms epoch). | 
 **symbol** | **string** | US-equity ticker filter, e.g. &#x60;NVDA&#x60;. | 
 **orderType** | [**PlaceEquityOrderOrderTypeParameter**](PlaceEquityOrderOrderTypeParameter.md) | Order type filter: &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | 
 **side** | [**PlaceEquityOrderSideParameter**](PlaceEquityOrderSideParameter.md) | Side filter: &#x60;BUY&#x60; / &#x60;SELL&#x60;. | 
 **orderStatus** | **string** | Comma-separated status filter. Allowed values: &#x60;FILLED&#x60;, &#x60;PARTIALLY_FILLED&#x60;, &#x60;CANCELED&#x60;, &#x60;EXPIRED&#x60;, &#x60;REJECTED&#x60;. | 
 **current** | **int32** | Page number, 1-based. Default &#x60;1&#x60;. | 
 **size** | **int32** | Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**EquityOrderHistoryResponse**](EquityOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## EquityTradeHistory

> EquityTradeHistoryResponse EquityTradeHistory(ctx).StartTime(startTime).EndTime(endTime).Symbol(symbol).Side(side).OrderId(orderId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Equity Trade History (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	startTime := int64(1735800000000) // int64 | Start time (ms epoch).
	endTime := int64(1735900000000) // int64 | End time (ms epoch).
	symbol := "NVDA" // string | US-equity ticker filter, e.g. `NVDA`. (optional)
	side := models.PlaceEquityOrderSideParameterBuy // PlaceEquityOrderSideParameter | Side filter: `BUY` / `SELL`. (optional)
	orderId := "c3c58f49-7b0d-4b9e-a2db-1a2f9a3b8c71" // string | Narrow the result to executions of a single order. (optional)
	current := int32(1) // int32 | Page number, 1-based. Default `1`. (optional)
	size := int32(20) // int32 | Page size. Default `20`, max `100`. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.EquityTradeHistory(context.Background()).StartTime(startTime).EndTime(endTime).Symbol(symbol).Side(side).OrderId(orderId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.EquityTradeHistory``: %v\n", err)
		return
	}

	// response from `EquityTradeHistory`: EquityTradeHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start time (ms epoch). | 
 **endTime** | **int64** | End time (ms epoch). | 
 **symbol** | **string** | US-equity ticker filter, e.g. &#x60;NVDA&#x60;. | 
 **side** | [**PlaceEquityOrderSideParameter**](PlaceEquityOrderSideParameter.md) | Side filter: &#x60;BUY&#x60; / &#x60;SELL&#x60;. | 
 **orderId** | **string** | Narrow the result to executions of a single order. | 
 **current** | **int32** | Page number, 1-based. Default &#x60;1&#x60;. | 
 **size** | **int32** | Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**EquityTradeHistoryResponse**](EquityTradeHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PlaceEquityOrder

> PlaceEquityOrderResponse PlaceEquityOrder(ctx).Symbol(symbol).Side(side).OrderType(orderType).QuoteAsset(quoteAsset).Price(price).Quantity(quantity).Notional(notional).TimeInForce(timeInForce).TradingSession(tradingSession).WalletType(walletType).ClientOrderId(clientOrderId).Tokenize(tokenize).RecvWindow(recvWindow).Execute()

Place Equity Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US stock ticker, e.g. `AAPL`, `TSLA`. Must be a symbol with tokenization enabled — check via `/market/tokenized-assets`.
	side := models.PlaceEquityOrderSideParameterBuy // PlaceEquityOrderSideParameter | `BUY` / `SELL`.
	orderType := models.PlaceEquityOrderOrderTypeParameterMarket // PlaceEquityOrderOrderTypeParameter | `MARKET` / `LIMIT`.
	quoteAsset := "USDC" // string | Quote asset. Defaults to `USDC`; must be within the server's allowed set. (optional)
	price := "180.50" // string | **Required** for `LIMIT`; **forbidden** for `MARKET`. Maximum 2 decimal places. (optional)
	quantity := "1" // string | **Required** for `LIMIT` (both sides) and `SELL MARKET`; **forbidden** for `BUY MARKET`. (optional)
	notional := "1000.00" // string | **Required** for `BUY MARKET`; **forbidden** for `LIMIT` and `SELL MARKET`. (optional)
	timeInForce := models.PlaceEquityOrderTimeInForceParameterDay // PlaceEquityOrderTimeInForceParameter | `DAY` (default) / `GTC`. `GTC` is only supported for `LIMIT` orders; a fractional-share `GTC` order must be paired with `tradingSession = EXTENDED` or `24H`. (optional)
	tradingSession := models.PlaceEquityOrderTradingSessionParameterRth // PlaceEquityOrderTradingSessionParameter | `RTH` / `EXTENDED` / `24H`. **Required** for `LIMIT`; **forbidden** for `MARKET`. (optional)
	walletType := models.PlaceEquityOrderWalletTypeParameterCard // PlaceEquityOrderWalletTypeParameter | Payment wallet for `BUY` orders: `CARD` (default) / `MAIN`. `SELL` orders always settle to `CARD`. (optional)
	clientOrderId := "web_2c9c92b74f1e4a7c8f3b9e1a2d3c4b5a" // string | Client-supplied order id. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted. (optional)
	tokenize := true // bool | Whether to tokenize the purchased stock asset upon settlement. Default `true`. Set to `false` to receive the underlying equity directly instead of a tokenized asset. (optional)
	recvWindow := int64(5000) // int64 | The value cannot be greater than `60000`. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceStocksClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PlaceEquityOrder(context.Background()).Symbol(symbol).Side(side).OrderType(orderType).QuoteAsset(quoteAsset).Price(price).Quantity(quantity).Notional(notional).TimeInForce(timeInForce).TradingSession(tradingSession).WalletType(walletType).ClientOrderId(clientOrderId).Tokenize(tokenize).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PlaceEquityOrder``: %v\n", err)
		return
	}

	// response from `PlaceEquityOrder`: PlaceEquityOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | US stock ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Must be a symbol with tokenization enabled — check via &#x60;/market/tokenized-assets&#x60;. | 
 **side** | [**PlaceEquityOrderSideParameter**](PlaceEquityOrderSideParameter.md) | &#x60;BUY&#x60; / &#x60;SELL&#x60;. | 
 **orderType** | [**PlaceEquityOrderOrderTypeParameter**](PlaceEquityOrderOrderTypeParameter.md) | &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | 
 **quoteAsset** | **string** | Quote asset. Defaults to &#x60;USDC&#x60;; must be within the server&#39;s allowed set. | 
 **price** | **string** | **Required** for &#x60;LIMIT&#x60;; **forbidden** for &#x60;MARKET&#x60;. Maximum 2 decimal places. | 
 **quantity** | **string** | **Required** for &#x60;LIMIT&#x60; (both sides) and &#x60;SELL MARKET&#x60;; **forbidden** for &#x60;BUY MARKET&#x60;. | 
 **notional** | **string** | **Required** for &#x60;BUY MARKET&#x60;; **forbidden** for &#x60;LIMIT&#x60; and &#x60;SELL MARKET&#x60;. | 
 **timeInForce** | [**PlaceEquityOrderTimeInForceParameter**](PlaceEquityOrderTimeInForceParameter.md) | &#x60;DAY&#x60; (default) / &#x60;GTC&#x60;. &#x60;GTC&#x60; is only supported for &#x60;LIMIT&#x60; orders; a fractional-share &#x60;GTC&#x60; order must be paired with &#x60;tradingSession &#x3D; EXTENDED&#x60; or &#x60;24H&#x60;. | 
 **tradingSession** | [**PlaceEquityOrderTradingSessionParameter**](PlaceEquityOrderTradingSessionParameter.md) | &#x60;RTH&#x60; / &#x60;EXTENDED&#x60; / &#x60;24H&#x60;. **Required** for &#x60;LIMIT&#x60;; **forbidden** for &#x60;MARKET&#x60;. | 
 **walletType** | [**PlaceEquityOrderWalletTypeParameter**](PlaceEquityOrderWalletTypeParameter.md) | Payment wallet for &#x60;BUY&#x60; orders: &#x60;CARD&#x60; (default) / &#x60;MAIN&#x60;. &#x60;SELL&#x60; orders always settle to &#x60;CARD&#x60;. | 
 **clientOrderId** | **string** | Client-supplied order id. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted. | 
 **tokenize** | **bool** | Whether to tokenize the purchased stock asset upon settlement. Default &#x60;true&#x60;. Set to &#x60;false&#x60; to receive the underlying equity directly instead of a tokenized asset. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 

### Return type

[**PlaceEquityOrderResponse**](PlaceEquityOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

