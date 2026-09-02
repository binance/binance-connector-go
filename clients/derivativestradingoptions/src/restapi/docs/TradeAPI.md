# \TradeAPI

All URIs are relative to *https://eapi.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AccountTradeList**](TradeAPI.md#AccountTradeList) | **Get** /eapi/v1/userTrades | Account Trade List (USER_DATA)
[**CancelAllOptionOrdersByUnderlying**](TradeAPI.md#CancelAllOptionOrdersByUnderlying) | **Delete** /eapi/v1/allOpenOrdersByUnderlying | Cancel All Option Orders By Underlying (TRADE)
[**CancelAllOptionOrdersOnSpecificSymbol**](TradeAPI.md#CancelAllOptionOrdersOnSpecificSymbol) | **Delete** /eapi/v1/allOpenOrders | Cancel all Option orders on specific symbol (TRADE)
[**CancelMultipleOptionOrders**](TradeAPI.md#CancelMultipleOptionOrders) | **Delete** /eapi/v1/batchOrders | Cancel Multiple Option Orders (TRADE)
[**CancelOptionOrder**](TradeAPI.md#CancelOptionOrder) | **Delete** /eapi/v1/order | Cancel Option Order (TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | **Post** /eapi/v1/order | New Order (TRADE)
[**OptionPositionInformation**](TradeAPI.md#OptionPositionInformation) | **Get** /eapi/v1/position | Option Position Information (USER_DATA)
[**PlaceMultipleOrders**](TradeAPI.md#PlaceMultipleOrders) | **Post** /eapi/v1/batchOrders | Place Multiple Orders (TRADE)
[**QueryCurrentOpenOptionOrders**](TradeAPI.md#QueryCurrentOpenOptionOrders) | **Get** /eapi/v1/openOrders | Query Current Open Option Orders (USER_DATA)
[**QueryOptionOrderHistory**](TradeAPI.md#QueryOptionOrderHistory) | **Get** /eapi/v1/historyOrders | Query Option Order History (TRADE)
[**QuerySingleOrder**](TradeAPI.md#QuerySingleOrder) | **Get** /eapi/v1/order | Query Single Order (TRADE)
[**TradfiOptionsContract**](TradeAPI.md#TradfiOptionsContract) | **Post** /eapi/v1/stock/contract | TradFi Options Contract (USER_DATA)
[**UserCommission**](TradeAPI.md#UserCommission) | **Get** /eapi/v1/commission | User Commission (USER_DATA)
[**UserExerciseRecord**](TradeAPI.md#UserExerciseRecord) | **Get** /eapi/v1/exerciseRecord | User Exercise Record (USER_DATA)


## AccountTradeList

> AccountTradeListResponse AccountTradeList(ctx).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	fromId := int64(1) // int64 | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(20) // int64 | Number of result sets returned. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.AccountTradeList(context.Background()).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.AccountTradeList``: %v\n", err)
		return
	}

	// response from `AccountTradeList`: AccountTradeListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**AccountTradeListResponse**](AccountTradeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllOptionOrdersByUnderlying

> CancelAllOptionOrdersByUnderlyingResponse CancelAllOptionOrdersByUnderlying(ctx).Underlying(underlying).RecvWindow(recvWindow).Execute()

Cancel All Option Orders By Underlying (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	underlying := "BTCUSDT" // string | Underlying asset.
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllOptionOrdersByUnderlying(context.Background()).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllOptionOrdersByUnderlying``: %v\n", err)
		return
	}

	// response from `CancelAllOptionOrdersByUnderlying`: CancelAllOptionOrdersByUnderlyingResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Underlying asset. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**CancelAllOptionOrdersByUnderlyingResponse**](CancelAllOptionOrdersByUnderlyingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelAllOptionOrdersOnSpecificSymbol

> CancelAllOptionOrdersOnSpecificSymbolResponse CancelAllOptionOrdersOnSpecificSymbol(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel all Option orders on specific symbol (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelAllOptionOrdersOnSpecificSymbol(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAllOptionOrdersOnSpecificSymbol``: %v\n", err)
		return
	}

	// response from `CancelAllOptionOrdersOnSpecificSymbol`: CancelAllOptionOrdersOnSpecificSymbolResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**CancelAllOptionOrdersOnSpecificSymbolResponse**](CancelAllOptionOrdersOnSpecificSymbolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelMultipleOptionOrders

> CancelMultipleOptionOrdersResponse CancelMultipleOptionOrders(ctx).Symbol(symbol).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()

Cancel Multiple Option Orders (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	orderIds := []int64{int64(4611875134427365000)} // []int64 | Order ID list. (optional)
	clientOrderIds := []string{"my_id_1"} // []string | Client order ID list. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelMultipleOptionOrders(context.Background()).Symbol(symbol).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelMultipleOptionOrders``: %v\n", err)
		return
	}

	// response from `CancelMultipleOptionOrders`: CancelMultipleOptionOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **orderIds** | **[]int64** | Order ID list. | 
 **clientOrderIds** | **[]string** | Client order ID list. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**CancelMultipleOptionOrdersResponse**](CancelMultipleOptionOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CancelOptionOrder

> CancelOptionOrderResponse CancelOptionOrder(ctx).Symbol(symbol).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Cancel Option Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	orderId := int64(4611875134427365000) // int64 | Order ID. (optional)
	clientOrderId := "10000" // string | clientOrderId (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.CancelOptionOrder(context.Background()).Symbol(symbol).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelOptionOrder``: %v\n", err)
		return
	}

	// response from `CancelOptionOrder`: CancelOptionOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **orderId** | **int64** | Order ID. | 
 **clientOrderId** | **string** | clientOrderId | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**CancelOptionOrderResponse**](CancelOptionOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder(ctx).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Price(price).TimeInForce(timeInForce).ReduceOnly(reduceOnly).PostOnly(postOnly).NewOrderRespType(newOrderRespType).ClientOrderId(clientOrderId).IsMmp(isMmp).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | 
	side := models.PlaceMultipleOrdersOrdersParameterInnerSideBuy // PlaceMultipleOrdersOrdersParameterInnerSide | 
	type_ := models.PlaceMultipleOrdersOrdersParameterInnerTypeLimit // PlaceMultipleOrdersOrdersParameterInnerType | 
	quantity := float64(1.0) // float64 | Order Quantity
	price := float64(1.0) // float64 | Order Price (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	reduceOnly := false // bool |  (optional)
	postOnly := false // bool |  (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	clientOrderId := "1" // string | User-defined order ID cannot be repeated in pending orders (optional)
	isMmp := true // bool | is market maker protection order (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter | Self-trade prevention mode (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Price(price).TimeInForce(timeInForce).ReduceOnly(reduceOnly).PostOnly(postOnly).NewOrderRespType(newOrderRespType).ClientOrderId(clientOrderId).IsMmp(isMmp).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewOrder``: %v\n", err)
		return
	}

	// response from `NewOrder`: NewOrderResponse
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
 **side** | [**PlaceMultipleOrdersOrdersParameterInnerSide**](PlaceMultipleOrdersOrdersParameterInnerSide.md) |  | 
 **type_** | [**PlaceMultipleOrdersOrdersParameterInnerType**](PlaceMultipleOrdersOrdersParameterInnerType.md) |  | 
 **quantity** | **float64** | Order Quantity | 
 **price** | **float64** | Order Price | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **reduceOnly** | **bool** |  | 
 **postOnly** | **bool** |  | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **clientOrderId** | **string** | User-defined order ID cannot be repeated in pending orders | 
 **isMmp** | **bool** | is market maker protection order | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) | Self-trade prevention mode | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewOrderResponse**](NewOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OptionPositionInformation

> OptionPositionInformationResponse OptionPositionInformation(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Option Position Information (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OptionPositionInformation(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OptionPositionInformation``: %v\n", err)
		return
	}

	// response from `OptionPositionInformation`: OptionPositionInformationResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**OptionPositionInformationResponse**](OptionPositionInformationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## PlaceMultipleOrders

> PlaceMultipleOrdersResponse PlaceMultipleOrders(ctx).Orders(orders).RecvWindow(recvWindow).Execute()

Place Multiple Orders (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	orders := []models.PlaceMultipleOrdersOrdersParameterInner{*models.NewPlaceMultipleOrdersOrdersParameterInner("BTC-200730-9000-C", models.placeMultipleOrders_orders_parameter_inner_side("BUY"), models.placeMultipleOrders_orders_parameter_inner_type("LIMIT"), float64(1.0))} // []PlaceMultipleOrdersOrdersParameterInner | order list. Max 10 orders
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.PlaceMultipleOrders(context.Background()).Orders(orders).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PlaceMultipleOrders``: %v\n", err)
		return
	}

	// response from `PlaceMultipleOrders`: PlaceMultipleOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **orders** | [**[]PlaceMultipleOrdersOrdersParameterInner**](PlaceMultipleOrdersOrdersParameterInner.md) | order list. Max 10 orders | 
 **recvWindow** | **int64** |  | 

### Return type

[**PlaceMultipleOrdersResponse**](PlaceMultipleOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryCurrentOpenOptionOrders

> QueryCurrentOpenOptionOrdersResponse QueryCurrentOpenOptionOrders(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Query Current Open Option Orders (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair. (optional)
	orderId := int64(4611875134427365000) // int64 | Order ID. (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryCurrentOpenOptionOrders(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryCurrentOpenOptionOrders``: %v\n", err)
		return
	}

	// response from `QueryCurrentOpenOptionOrders`: QueryCurrentOpenOptionOrdersResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **orderId** | **int64** | Order ID. | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**QueryCurrentOpenOptionOrdersResponse**](QueryCurrentOpenOptionOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QueryOptionOrderHistory

> QueryOptionOrderHistoryResponse QueryOptionOrderHistory(ctx).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Option Order History (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	orderId := int64(4611875134427365000) // int64 | Order ID. (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(20) // int64 | Number of result sets returned (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QueryOptionOrderHistory(context.Background()).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryOptionOrderHistory``: %v\n", err)
		return
	}

	// response from `QueryOptionOrderHistory`: QueryOptionOrderHistoryResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **orderId** | **int64** | Order ID. | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**QueryOptionOrderHistoryResponse**](QueryOptionOrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## QuerySingleOrder

> QuerySingleOrderResponse QuerySingleOrder(ctx).Symbol(symbol).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Query Single Order (TRADE)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair.
	orderId := int64(4611875134427365000) // int64 | Order ID. (optional)
	clientOrderId := "abc123" // string | User-defined order ID; cannot be duplicated among open orders. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.QuerySingleOrder(context.Background()).Symbol(symbol).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QuerySingleOrder``: %v\n", err)
		return
	}

	// response from `QuerySingleOrder`: QuerySingleOrderResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **orderId** | **int64** | Order ID. | 
 **clientOrderId** | **string** | User-defined order ID; cannot be duplicated among open orders. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**QuerySingleOrderResponse**](QuerySingleOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## TradfiOptionsContract

> TradfiOptionsContractResponse TradfiOptionsContract(ctx).RecvWindow(recvWindow).Execute()

TradFi Options Contract (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.TradfiOptionsContract(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.TradfiOptionsContract``: %v\n", err)
		return
	}

	// response from `TradfiOptionsContract`: TradfiOptionsContractResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**TradfiOptionsContractResponse**](TradfiOptionsContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UserCommission

> UserCommissionResponse UserCommission(ctx).RecvWindow(recvWindow).Execute()

User Commission (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.UserCommission(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.UserCommission``: %v\n", err)
		return
	}

	// response from `UserCommission`: UserCommissionResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**UserCommissionResponse**](UserCommissionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## UserExerciseRecord

> UserExerciseRecordResponse UserExerciseRecord(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User Exercise Record (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTC-200730-9000-C" // string | Option trading pair. (optional)
	startTime := int64(1623319461670) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(1641782889000) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int64(20) // int64 | Number of result sets returned. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.UserExerciseRecord(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.UserExerciseRecord``: %v\n", err)
		return
	}

	// response from `UserExerciseRecord`: UserExerciseRecordResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair. | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int64** | Number of result sets returned. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**UserExerciseRecordResponse**](UserExerciseRecordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

