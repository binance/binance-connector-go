# \TradeAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**OpenOrdersCancelAll**](TradeAPI.md#OpenOrdersCancelAll) | /openOrders.cancelAll | WebSocket Cancel open orders
[**OrderCancel**](TradeAPI.md#OrderCancel) | /order.cancel | WebSocket Cancel order
[**OrderCancelReplace**](TradeAPI.md#OrderCancelReplace) | /order.cancelReplace | WebSocket Cancel and replace order
[**OrderListCancel**](TradeAPI.md#OrderListCancel) | /orderList.cancel | WebSocket Cancel Order list
[**OrderListPlace**](TradeAPI.md#OrderListPlace) | /orderList.place | WebSocket Place new OCO - Deprecated
[**OrderListPlaceOco**](TradeAPI.md#OrderListPlaceOco) | /orderList.place.oco | WebSocket Place new Order list - OCO
[**OrderListPlaceOpo**](TradeAPI.md#OrderListPlaceOpo) | /orderList.place.opo | WebSocket OPO
[**OrderListPlaceOpoco**](TradeAPI.md#OrderListPlaceOpoco) | /orderList.place.opoco | WebSocket OPOCO
[**OrderListPlaceOto**](TradeAPI.md#OrderListPlaceOto) | /orderList.place.oto | WebSocket Place new Order list - OTO
[**OrderListPlaceOtoco**](TradeAPI.md#OrderListPlaceOtoco) | /orderList.place.otoco | WebSocket Place new Order list - OTOCO
[**OrderPlace**](TradeAPI.md#OrderPlace) | /order.place | WebSocket Place new order
[**OrderTest**](TradeAPI.md#OrderTest) | /order.test | WebSocket Test new order
[**SorOrderPlace**](TradeAPI.md#SorOrderPlace) | /sor.order.place | WebSocket Place new order using SOR
[**SorOrderTest**](TradeAPI.md#SorOrderTest) | /sor.order.test | WebSocket Test new order using SOR


## OpenOrdersCancelAll

> OpenOrdersCancelAllResponse OpenOrdersCancelAll().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()

WebSocket Cancel open orders


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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OpenOrdersCancelAll().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OpenOrdersCancelAll``: %v\n", err)
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

[**OpenOrdersCancelAllResponse**](OpenOrdersCancelAllResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderCancel

> OrderCancelResponse OrderCancel().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()

WebSocket Cancel order


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
	orderId := int64(1) // int64 | Cancel order by orderId (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	cancelRestrictions := models.OrderCancelCancelRestrictionsParameterOnlyNew // OrderCancelCancelRestrictionsParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderCancel().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderCancel``: %v\n", err)
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
 **orderId** | **int64** | Cancel order by orderId | 
 **origClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **cancelRestrictions** | [**OrderCancelCancelRestrictionsParameter**](OrderCancelCancelRestrictionsParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderCancelResponse**](OrderCancelResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderCancelReplace

> OrderCancelReplaceResponse OrderCancelReplace().Symbol(symbol).CancelReplaceMode(cancelReplaceMode).Side(side).Type(type_).Id(id).CancelOrderId(cancelOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelNewClientOrderId(cancelNewClientOrderId).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).CancelRestrictions(cancelRestrictions).OrderRateLimitExceededMode(orderRateLimitExceededMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

WebSocket Cancel and replace order


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
	cancelReplaceMode := models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure // OrderCancelReplaceCancelReplaceModeParameter | 
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	cancelOrderId := int64(1) // int64 | Cancel order by orderId (optional)
	cancelOrigClientOrderId := "cancelOrigClientOrderId_example" // string |  (optional)
	cancelNewClientOrderId := "cancelNewClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter |  (optional)
	price := float32(1.0) // float32 |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	stopPrice := float32(1.0) // float32 |  (optional)
	trailingDelta := float32(1.0) // float32 | See Trailing Stop order FAQ (optional)
	icebergQty := float32(1.0) // float32 |  (optional)
	strategyId := int64(1) // int64 | Arbitrary numeric value identifying the order within an order strategy. (optional)
	strategyType := int32(1) // int32 | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	cancelRestrictions := models.OrderCancelCancelRestrictionsParameterOnlyNew // OrderCancelCancelRestrictionsParameter |  (optional)
	orderRateLimitExceededMode := models.OrderCancelReplaceOrderRateLimitExceededModeParameterDoNothing // OrderCancelReplaceOrderRateLimitExceededModeParameter |  (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Price level to peg the price to (max: 100)       See Pegged Orders (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderCancelReplace().Symbol(symbol).CancelReplaceMode(cancelReplaceMode).Side(side).Type(type_).Id(id).CancelOrderId(cancelOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelNewClientOrderId(cancelNewClientOrderId).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).CancelRestrictions(cancelRestrictions).OrderRateLimitExceededMode(orderRateLimitExceededMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderCancelReplace``: %v\n", err)
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
 **cancelReplaceMode** | [**OrderCancelReplaceCancelReplaceModeParameter**](OrderCancelReplaceCancelReplaceModeParameter.md) |  | 
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **cancelOrderId** | **int64** | Cancel order by orderId | 
 **cancelOrigClientOrderId** | **string** |  | 
 **cancelNewClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **stopPrice** | **float32** |  | 
 **trailingDelta** | **float32** | See Trailing Stop order FAQ | 
 **icebergQty** | **float32** |  | 
 **strategyId** | **int64** | Arbitrary numeric value identifying the order within an order strategy. | 
 **strategyType** | **int32** | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **cancelRestrictions** | [**OrderCancelCancelRestrictionsParameter**](OrderCancelCancelRestrictionsParameter.md) |  | 
 **orderRateLimitExceededMode** | [**OrderCancelReplaceOrderRateLimitExceededModeParameter**](OrderCancelReplaceOrderRateLimitExceededModeParameter.md) |  | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Price level to peg the price to (max: 100)       See Pegged Orders | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderCancelReplaceResponse**](OrderCancelReplaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListCancel

> OrderListCancelResponse OrderListCancel().Symbol(symbol).Id(id).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

WebSocket Cancel Order list


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
	orderListId := int32(1) // int32 | Cancel order list by orderListId (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListCancel().Symbol(symbol).Id(id).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListCancel``: %v\n", err)
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
 **orderListId** | **int32** | Cancel order list by orderListId | 
 **listClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListCancelResponse**](OrderListCancelResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlace

> OrderListPlaceResponse OrderListPlace().Symbol(symbol).Side(side).Price(price).Quantity(quantity).Id(id).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopIcebergQty(stopIcebergQty).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

WebSocket Place new OCO - Deprecated


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	price := float32(1.0) // float32 | Price for the limit order
	quantity := float32(1.0) // float32 | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	limitClientOrderId := "limitClientOrderId_example" // string | Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent (optional)
	limitIcebergQty := float32(1.0) // float32 |  (optional)
	limitStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the limit order within an order strategy. (optional)
	limitStrategyType := int32(1) // int32 | <p>Arbitrary numeric value identifying the limit order strategy.</p><p>Values smaller than `1000000` are reserved and cannot be used.</p> (optional)
	stopPrice := float32(1.0) // float32 |  (optional)
	trailingDelta := int32(1) // int32 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) (optional)
	stopClientOrderId := "stopClientOrderId_example" // string | Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent (optional)
	stopLimitPrice := float32(1.0) // float32 |  (optional)
	stopLimitTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	stopIcebergQty := float32(1.0) // float32 |  (optional)
	stopStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the stop order within an order strategy. (optional)
	stopStrategyType := int32(1) // int32 | <p>Arbitrary numeric value identifying the stop order strategy.</p><p>Values smaller than `1000000` are reserved and cannot be used.</p> (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlace().Symbol(symbol).Side(side).Price(price).Quantity(quantity).Id(id).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopIcebergQty(stopIcebergQty).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlace``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **price** | **float32** | Price for the limit order | 
 **quantity** | **float32** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **limitClientOrderId** | **string** | Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent | 
 **limitIcebergQty** | **float32** |  | 
 **limitStrategyId** | **int64** | Arbitrary numeric value identifying the limit order within an order strategy. | 
 **limitStrategyType** | **int32** | &lt;p&gt;Arbitrary numeric value identifying the limit order strategy.&lt;/p&gt;&lt;p&gt;Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.&lt;/p&gt; | 
 **stopPrice** | **float32** |  | 
 **trailingDelta** | **int32** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) | 
 **stopClientOrderId** | **string** | Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent | 
 **stopLimitPrice** | **float32** |  | 
 **stopLimitTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **stopIcebergQty** | **float32** |  | 
 **stopStrategyId** | **int64** | Arbitrary numeric value identifying the stop order within an order strategy. | 
 **stopStrategyType** | **int32** | &lt;p&gt;Arbitrary numeric value identifying the stop order strategy.&lt;/p&gt;&lt;p&gt;Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.&lt;/p&gt; | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceResponse**](OrderListPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOco

> OrderListPlaceOcoResponse OrderListPlaceOco().Symbol(symbol).Side(side).Quantity(quantity).AboveType(aboveType).BelowType(belowType).Id(id).ListClientOrderId(listClientOrderId).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveTrailingDelta(aboveTrailingDelta).AboveTimeInForce(aboveTimeInForce).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AbovePegPriceType(abovePegPriceType).AbovePegOffsetType(abovePegOffsetType).AbovePegOffsetValue(abovePegOffsetValue).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowTrailingDelta(belowTrailingDelta).BelowTimeInForce(belowTimeInForce).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowPegPriceType(belowPegPriceType).BelowPegOffsetType(belowPegOffsetType).BelowPegOffsetValue(belowPegOffsetValue).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

WebSocket Place new Order list - OCO


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	quantity := float32(1.0) // float32 | 
	aboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | 
	belowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	aboveClientOrderId := "aboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent (optional)
	aboveIcebergQty := int64(1) // int64 | Note that this can only be used if `aboveTimeInForce` is `GTC`. (optional)
	abovePrice := float32(1.0) // float32 | Can be used if `aboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	aboveStopPrice := float32(1.0) // float32 | Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. <br>Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified. (optional)
	aboveTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	aboveTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	aboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the above order within an order strategy. (optional)
	aboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the above order strategy. <br>Values smaller than 1000000 are reserved and cannot be used. (optional)
	abovePegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	abovePegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	abovePegOffsetValue := int32(1) // int32 |  (optional)
	belowClientOrderId := "belowClientOrderId_example" // string |  (optional)
	belowIcebergQty := int64(1) // int64 | Note that this can only be used if `belowTimeInForce` is `GTC`. (optional)
	belowPrice := float32(1.0) // float32 | Can be used if `belowType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	belowStopPrice := float32(1.0) // float32 | Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT` or `TAKE_PROFIT_LIMIT`. <br>Either `belowStopPrice` or `belowTrailingDelta` or both, must be specified. (optional)
	belowTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	belowTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	belowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the below order within an order strategy. (optional)
	belowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the below order strategy. <br>Values smaller than 1000000 are reserved and cannot be used. (optional)
	belowPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	belowPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	belowPegOffsetValue := int32(1) // int32 |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().Symbol(symbol).Side(side).Quantity(quantity).AboveType(aboveType).BelowType(belowType).Id(id).ListClientOrderId(listClientOrderId).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveTrailingDelta(aboveTrailingDelta).AboveTimeInForce(aboveTimeInForce).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AbovePegPriceType(abovePegPriceType).AbovePegOffsetType(abovePegOffsetType).AbovePegOffsetValue(abovePegOffsetValue).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowTrailingDelta(belowTrailingDelta).BelowTimeInForce(belowTimeInForce).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowPegPriceType(belowPegPriceType).BelowPegOffsetType(belowPegOffsetType).BelowPegOffsetValue(belowPegOffsetValue).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlaceOco``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **quantity** | **float32** |  | 
 **aboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) |  | 
 **belowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **aboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent | 
 **aboveIcebergQty** | **int64** | Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **abovePrice** | **float32** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **aboveStopPrice** | **float32** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified. | 
 **aboveTrailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **aboveTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **aboveStrategyId** | **int64** | Arbitrary numeric value identifying the above order within an order strategy. | 
 **aboveStrategyType** | **int32** | Arbitrary numeric value identifying the above order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used. | 
 **abovePegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **abovePegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **abovePegOffsetValue** | **int32** |  | 
 **belowClientOrderId** | **string** |  | 
 **belowIcebergQty** | **int64** | Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **belowPrice** | **float32** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **belowStopPrice** | **float32** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;belowStopPrice&#x60; or &#x60;belowTrailingDelta&#x60; or both, must be specified. | 
 **belowTrailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **belowTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **belowStrategyId** | **int64** | Arbitrary numeric value identifying the below order within an order strategy. | 
 **belowStrategyType** | **int32** | Arbitrary numeric value identifying the below order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used. | 
 **belowPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **belowPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **belowPegOffsetValue** | **int32** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOcoResponse**](OrderListPlaceOcoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOpo

> OrderListPlaceOpoResponse OrderListPlaceOpo().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

WebSocket OPO


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | 
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingType := models.OrderListPlaceOpoPendingTypeParameterLimit // OrderListPlaceOpoPendingTypeParameter | 
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  (optional)
	pendingPrice := float32(1.0) // float32 |  (optional)
	pendingStopPrice := float32(1.0) // float32 |  (optional)
	pendingTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingIcebergQty := float32(1.0) // float32 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.  (optional)
	pendingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy.  (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlaceOpo``: %v\n", err)
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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingType** | [**OrderListPlaceOpoPendingTypeParameter**](OrderListPlaceOpoPendingTypeParameter.md) |  | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  | 
 **pendingPrice** | **float32** |  | 
 **pendingStopPrice** | **float32** |  | 
 **pendingTrailingDelta** | **float32** |  | 
 **pendingIcebergQty** | **float32** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy.  | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOpoResponse**](OrderListPlaceOpoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOpoco

> OrderListPlaceOpocoResponse OrderListPlaceOpoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

WebSocket OPOCO


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | 
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	pendingAboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  (optional)
	pendingAbovePrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.  (optional)
	pendingAboveStopPrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`  (optional)
	pendingAboveTrailingDelta := float32(1.0) // float32 | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  (optional)
	pendingAboveIcebergQty := float32(1.0) // float32 | This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.  (optional)
	pendingAboveTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy.  (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingAbovePegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingAbovePegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 |  (optional)
	pendingBelowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter |  (optional)
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  (optional)
	pendingBelowPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price  (optional)
	pendingBelowStopPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.  (optional)
	pendingBelowTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingBelowIcebergQty := float32(1.0) // float32 | This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.  (optional)
	pendingBelowTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy.  (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingBelowPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingBelowPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlaceOpoco``: %v\n", err)
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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **pendingAboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  | 
 **pendingAbovePrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.  | 
 **pendingAboveStopPrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;  | 
 **pendingAboveTrailingDelta** | **float32** | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  | 
 **pendingAboveIcebergQty** | **float32** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingAboveTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy.  | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingAbovePegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingAbovePegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingAbovePegOffsetValue** | **int32** |  | 
 **pendingBelowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) |  | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  | 
 **pendingBelowPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price  | 
 **pendingBelowStopPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.  | 
 **pendingBelowTrailingDelta** | **float32** |  | 
 **pendingBelowIcebergQty** | **float32** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingBelowTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy.  | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingBelowPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingBelowPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOpocoResponse**](OrderListPlaceOpocoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOto

> OrderListPlaceOtoResponse OrderListPlaceOto().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegOffsetType(pendingPegOffsetType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

WebSocket Place new Order list - OTO


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | 
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingType := models.OrderListPlaceOpoPendingTypeParameterLimit // OrderListPlaceOpoPendingTypeParameter | 
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	pendingQuantity := float32(1.0) // float32 | Sets the quantity for the pending order.
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  (optional)
	pendingPrice := float32(1.0) // float32 |  (optional)
	pendingStopPrice := float32(1.0) // float32 |  (optional)
	pendingTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingIcebergQty := float32(1.0) // float32 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.  (optional)
	pendingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy.  (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegOffsetType(pendingPegOffsetType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlaceOto``: %v\n", err)
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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingType** | [**OrderListPlaceOpoPendingTypeParameter**](OrderListPlaceOpoPendingTypeParameter.md) |  | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **pendingQuantity** | **float32** | Sets the quantity for the pending order. | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  | 
 **pendingPrice** | **float32** |  | 
 **pendingStopPrice** | **float32** |  | 
 **pendingTrailingDelta** | **float32** |  | 
 **pendingIcebergQty** | **float32** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy.  | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOtoResponse**](OrderListPlaceOtoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOtoco

> OrderListPlaceOtocoResponse OrderListPlaceOtoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

WebSocket Place new Order list - OTOCO


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | 
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	pendingQuantity := float32(1.0) // float32 | Sets the quantity for the pending order.
	pendingAboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	listClientOrderId := "listClientOrderId_example" // string |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  (optional)
	pendingAbovePrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.  (optional)
	pendingAboveStopPrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`  (optional)
	pendingAboveTrailingDelta := float32(1.0) // float32 | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  (optional)
	pendingAboveIcebergQty := float32(1.0) // float32 | This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.  (optional)
	pendingAboveTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy.  (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingAbovePegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingAbovePegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 |  (optional)
	pendingBelowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter |  (optional)
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  (optional)
	pendingBelowPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price  (optional)
	pendingBelowStopPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.  (optional)
	pendingBelowTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingBelowIcebergQty := float32(1.0) // float32 | This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.  (optional)
	pendingBelowTimeInForce := models.OrderListPlaceStopLimitTimeInForceParameterGtc // OrderListPlaceStopLimitTimeInForceParameter |  (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy.  (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingBelowPegPriceType := models.OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListPlaceOcoAbovePegPriceTypeParameter |  (optional)
	pendingBelowPegOffsetType := models.OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel // OrderListPlaceOcoAbovePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListPlaceOtoco``: %v\n", err)
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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **pendingQuantity** | **float32** | Sets the quantity for the pending order. | 
 **pendingAboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **listClientOrderId** | **string** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  | 
 **pendingAbovePrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.  | 
 **pendingAboveStopPrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;  | 
 **pendingAboveTrailingDelta** | **float32** | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  | 
 **pendingAboveIcebergQty** | **float32** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingAboveTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy.  | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingAbovePegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingAbovePegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingAbovePegOffsetValue** | **int32** |  | 
 **pendingBelowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) |  | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  | 
 **pendingBelowPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price  | 
 **pendingBelowStopPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.  | 
 **pendingBelowTrailingDelta** | **float32** |  | 
 **pendingBelowIcebergQty** | **float32** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingBelowTimeInForce** | [**OrderListPlaceStopLimitTimeInForceParameter**](OrderListPlaceStopLimitTimeInForceParameter.md) |  | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy.  | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingBelowPegPriceType** | [**OrderListPlaceOcoAbovePegPriceTypeParameter**](OrderListPlaceOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingBelowPegOffsetType** | [**OrderListPlaceOcoAbovePegOffsetTypeParameter**](OrderListPlaceOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOtocoResponse**](OrderListPlaceOtocoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderPlace

> OrderPlaceResponse OrderPlace().Symbol(symbol).Side(side).Type(type_).Id(id).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

WebSocket Place new order


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter |  (optional)
	price := float32(1.0) // float32 |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	stopPrice := float32(1.0) // float32 |  (optional)
	trailingDelta := int32(1) // int32 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) (optional)
	icebergQty := float32(1.0) // float32 |  (optional)
	strategyId := int64(1) // int64 | Arbitrary numeric value identifying the order within an order strategy. (optional)
	strategyType := int32(1) // int32 | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Price level to peg the price to (max: 100)       See Pegged Orders (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderPlace().Symbol(symbol).Side(side).Type(type_).Id(id).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderPlace``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **stopPrice** | **float32** |  | 
 **trailingDelta** | **int32** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) | 
 **icebergQty** | **float32** |  | 
 **strategyId** | **int64** | Arbitrary numeric value identifying the order within an order strategy. | 
 **strategyType** | **int32** | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Price level to peg the price to (max: 100)       See Pegged Orders | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderPlaceResponse**](OrderPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderTest

> OrderTestResponse OrderTest().Symbol(symbol).Side(side).Type(type_).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

WebSocket Test new order


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	computeCommissionRates := false // bool | Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter |  (optional)
	price := float32(1.0) // float32 |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	stopPrice := float32(1.0) // float32 |  (optional)
	trailingDelta := int32(1) // int32 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) (optional)
	icebergQty := float32(1.0) // float32 |  (optional)
	strategyId := int64(1) // int64 | Arbitrary numeric value identifying the order within an order strategy. (optional)
	strategyType := int32(1) // int32 | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Price level to peg the price to (max: 100)       See Pegged Orders (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderTest().Symbol(symbol).Side(side).Type(type_).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderTest``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **stopPrice** | **float32** |  | 
 **trailingDelta** | **int32** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md) | 
 **icebergQty** | **float32** |  | 
 **strategyId** | **int64** | Arbitrary numeric value identifying the order within an order strategy. | 
 **strategyType** | **int32** | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Price level to peg the price to (max: 100)       See Pegged Orders | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderTestResponse**](OrderTestResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SorOrderPlace

> SorOrderPlaceResponse SorOrderPlace().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

WebSocket Place new order using SOR


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | 
	quantity := float32(1.0) // float32 | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter |  (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	icebergQty := float32(1.0) // float32 |  (optional)
	strategyId := int64(1) // int64 | Arbitrary numeric value identifying the order within an order strategy. (optional)
	strategyType := int32(1) // int32 | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.SorOrderPlace().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SorOrderPlace``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) |  | 
 **quantity** | **float32** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **icebergQty** | **float32** |  | 
 **strategyId** | **int64** | Arbitrary numeric value identifying the order within an order strategy. | 
 **strategyType** | **int32** | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderPlaceResponse**](SorOrderPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SorOrderTest

> SorOrderTestResponse SorOrderTest().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

WebSocket Test new order using SOR


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | 
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | 
	quantity := float32(1.0) // float32 | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	computeCommissionRates := false // bool | Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter |  (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | New ID for the canceled order. Automatically generated if not sent (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter |  (optional)
	icebergQty := float32(1.0) // float32 |  (optional)
	strategyId := int64(1) // int64 | Arbitrary numeric value identifying the order within an order strategy. (optional)
	strategyType := int32(1) // int32 | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter |  (optional)
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


	resp, err := wsClient.WebsocketAPI.TradeAPI.SorOrderTest().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SorOrderTest``: %v\n", err)
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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) |  | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) |  | 
 **quantity** | **float32** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | New ID for the canceled order. Automatically generated if not sent | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) |  | 
 **icebergQty** | **float32** |  | 
 **strategyId** | **int64** | Arbitrary numeric value identifying the order within an order strategy. | 
 **strategyType** | **int32** | Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderTestResponse**](SorOrderTestResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

