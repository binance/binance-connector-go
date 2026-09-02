# \TradeAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**OpenOrdersCancelAll**](TradeAPI.md#OpenOrdersCancelAll) | /openOrders.cancelAll | Cancel open orders (TRADE)
[**OrderAmendKeepPriority**](TradeAPI.md#OrderAmendKeepPriority) | /order.amend.keepPriority | Order Amend Keep Priority (TRADE)
[**OrderCancel**](TradeAPI.md#OrderCancel) | /order.cancel | Cancel order (TRADE)
[**OrderCancelReplace**](TradeAPI.md#OrderCancelReplace) | /order.cancelReplace | Cancel and replace order (TRADE)
[**OrderListCancel**](TradeAPI.md#OrderListCancel) | /orderList.cancel | Cancel Order list (TRADE)
[**OrderListPlace**](TradeAPI.md#OrderListPlace) | /orderList.place | Place new OCO - Deprecated (TRADE)
[**OrderListPlaceOco**](TradeAPI.md#OrderListPlaceOco) | /orderList.place.oco | Place new Order list - OCO (TRADE)
[**OrderListPlaceOpo**](TradeAPI.md#OrderListPlaceOpo) | /orderList.place.opo | OPO (TRADE)
[**OrderListPlaceOpoco**](TradeAPI.md#OrderListPlaceOpoco) | /orderList.place.opoco | OPOCO (TRADE)
[**OrderListPlaceOto**](TradeAPI.md#OrderListPlaceOto) | /orderList.place.oto | Place new Order list - OTO (TRADE)
[**OrderListPlaceOtoco**](TradeAPI.md#OrderListPlaceOtoco) | /orderList.place.otoco | Place new Order list - OTOCO (TRADE)
[**OrderPlace**](TradeAPI.md#OrderPlace) | /order.place | Place new order (TRADE)
[**OrderTest**](TradeAPI.md#OrderTest) | /order.test | Test new order (TRADE)
[**SorOrderPlace**](TradeAPI.md#SorOrderPlace) | /sor.order.place | Place new order using SOR (TRADE)
[**SorOrderTest**](TradeAPI.md#SorOrderTest) | /sor.order.test | Test new order using SOR (TRADE)


## OpenOrdersCancelAll

> OpenOrdersCancelAllResponse OpenOrdersCancelAll().Symbol(symbol).Id(id).RecvWindow(recvWindow).Execute()

Cancel open orders (TRADE)


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
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Client-generated request identifier. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OpenOrdersCancelAllResponse**](OpenOrdersCancelAllResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderAmendKeepPriority

> OrderAmendKeepPriorityResponse OrderAmendKeepPriority().Symbol(symbol).NewQty(newQty).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Order Amend Keep Priority (TRADE)


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
	newQty := float64(1) // float64 | `newQty` must be greater than 0 and less than the order's quantity.
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	orderId := int64(1) // int64 | `orderId` or `origClientOrderId` must be sent (optional)
	origClientOrderId := "myOrder1" // string | `orderId` or `origClientOrderId` must be sent (optional)
	newClientOrderId := "myOrder2" // string | The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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


	resp, err := wsClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().Symbol(symbol).NewQty(newQty).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderAmendKeepPriority``: %v\n", err)
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
 **newQty** | **float64** | &#x60;newQty&#x60; must be greater than 0 and less than the order&#39;s quantity. | 
 **id** | **string** | Client-generated request identifier. | 
 **orderId** | **int64** | &#x60;orderId&#x60; or &#x60;origClientOrderId&#x60; must be sent | 
 **origClientOrderId** | **string** | &#x60;orderId&#x60; or &#x60;origClientOrderId&#x60; must be sent | 
 **newClientOrderId** | **string** | The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendKeepPriorityResponse**](OrderAmendKeepPriorityResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderCancel

> OrderCancelResponse OrderCancel().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()

Cancel order (TRADE)


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
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "myOrder1" // string |  (optional)
	newClientOrderId := "cancelMyOrder1" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional)
	cancelRestrictions := models.OrderCancelCancelRestrictionsParameterOnlyNew // OrderCancelCancelRestrictionsParameter | Supported values: <br>`ONLY_NEW` - Cancel will succeed if the order status is `NEW`.<br> `ONLY_PARTIALLY_FILLED` - Cancel will succeed if order status is `PARTIALLY_FILLED`. (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Client-generated request identifier. | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | 
 **cancelRestrictions** | [**OrderCancelCancelRestrictionsParameter**](OrderCancelCancelRestrictionsParameter.md) | Supported values: &lt;br&gt;&#x60;ONLY_NEW&#x60; - Cancel will succeed if the order status is &#x60;NEW&#x60;.&lt;br&gt; &#x60;ONLY_PARTIALLY_FILLED&#x60; - Cancel will succeed if order status is &#x60;PARTIALLY_FILLED&#x60;. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderCancelResponse**](OrderCancelResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderCancelReplace

> OrderCancelReplaceResponse OrderCancelReplace().Symbol(symbol).CancelReplaceMode(cancelReplaceMode).Side(side).Type(type_).Id(id).CancelOrderId(cancelOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelNewClientOrderId(cancelNewClientOrderId).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).CancelRestrictions(cancelRestrictions).OrderRateLimitExceededMode(orderRateLimitExceededMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

Cancel and replace order (TRADE)


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
	cancelReplaceMode := models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure // OrderCancelReplaceCancelReplaceModeParameter | The allowed values are: <br/> `STOP_ON_FAILURE` - If the cancel request fails, the new order placement will not be attempted. <br/> `ALLOW_FAILURE` - new order placement will be attempted even if cancel request fails.
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Please see [Enums](/products/spot/enums#side) for supported values.
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | Please see [Enums](/products/spot/enums#ordertypes) for supported values.
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	cancelOrderId := int64(1) // int64 | Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br>If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br>If both conditions are not met the request will be rejected. (optional)
	cancelOrigClientOrderId := "myOrder1" // string | Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br> If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br> If both conditions are not met the request will be rejected. (optional)
	cancelNewClientOrderId := "cancelMyOrder1" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Please see [Enums](/products/spot/enums#timeinforce) for supported values. (optional)
	price := float64(1) // float64 |  (optional)
	quantity := float64(1) // float64 |  (optional)
	quoteOrderQty := float64(1) // float64 |  (optional)
	newClientOrderId := "myOrder2" // string | Used to identify the new order. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Allowed values: <br/> `ACK`, `RESULT`, `FULL` <br/> `MARKET` and `LIMIT` orders types default to `FULL`; all other orders default to `ACK` (optional)
	stopPrice := float64(1) // float64 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := float64(1) // float64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	icebergQty := float64(1) // float64 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes). (optional)
	cancelRestrictions := models.OrderCancelCancelRestrictionsParameterOnlyNew // OrderCancelCancelRestrictionsParameter | Supported values: <br>`ONLY_NEW` - Cancel will succeed if the order status is `NEW`.<br> `ONLY_PARTIALLY_FILLED` - Cancel will succeed if order status is `PARTIALLY_FILLED`. (optional)
	orderRateLimitExceededMode := models.OrderCancelReplaceOrderRateLimitExceededModeParameterDoNothing // OrderCancelReplaceOrderRateLimitExceededModeParameter | Supported values: <br> `DO_NOTHING` (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit<br> `CANCEL_ONLY` - will always cancel the order (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | `PRIMARY_PEG` or `MARKET_PEG` <br> See Pegged Orders (optional)
	pegOffsetValue := int32(1) // int32 | Price level to peg the price to (max: 100) <br> See Pegged Orders (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter | Only `PRICE_LEVEL` is supported <br> See Pegged Orders (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **cancelReplaceMode** | [**OrderCancelReplaceCancelReplaceModeParameter**](OrderCancelReplaceCancelReplaceModeParameter.md) | The allowed values are: &lt;br/&gt; &#x60;STOP_ON_FAILURE&#x60; - If the cancel request fails, the new order placement will not be attempted. &lt;br/&gt; &#x60;ALLOW_FAILURE&#x60; - new order placement will be attempted even if cancel request fails. | 
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Please see [Enums](/products/spot/enums#side) for supported values. | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) | Please see [Enums](/products/spot/enums#ordertypes) for supported values. | 
 **id** | **string** | Client-generated request identifier. | 
 **cancelOrderId** | **int64** | Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt;If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt;If both conditions are not met the request will be rejected. | 
 **cancelOrigClientOrderId** | **string** | Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt; If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt; If both conditions are not met the request will be rejected. | 
 **cancelNewClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Please see [Enums](/products/spot/enums#timeinforce) for supported values. | 
 **price** | **float64** |  | 
 **quantity** | **float64** |  | 
 **quoteOrderQty** | **float64** |  | 
 **newClientOrderId** | **string** | Used to identify the new order. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Allowed values: &lt;br/&gt; &#x60;ACK&#x60;, &#x60;RESULT&#x60;, &#x60;FULL&#x60; &lt;br/&gt; &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; orders types default to &#x60;FULL&#x60;; all other orders default to &#x60;ACK&#x60; | 
 **stopPrice** | **float64** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **float64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes). | 
 **cancelRestrictions** | [**OrderCancelCancelRestrictionsParameter**](OrderCancelCancelRestrictionsParameter.md) | Supported values: &lt;br&gt;&#x60;ONLY_NEW&#x60; - Cancel will succeed if the order status is &#x60;NEW&#x60;.&lt;br&gt; &#x60;ONLY_PARTIALLY_FILLED&#x60; - Cancel will succeed if order status is &#x60;PARTIALLY_FILLED&#x60;. | 
 **orderRateLimitExceededMode** | [**OrderCancelReplaceOrderRateLimitExceededModeParameter**](OrderCancelReplaceOrderRateLimitExceededModeParameter.md) | Supported values: &lt;br&gt; &#x60;DO_NOTHING&#x60; (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit&lt;br&gt; &#x60;CANCEL_ONLY&#x60; - will always cancel the order | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60; &lt;br&gt; See Pegged Orders | 
 **pegOffsetValue** | **int32** | Price level to peg the price to (max: 100) &lt;br&gt; See Pegged Orders | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) | Only &#x60;PRICE_LEVEL&#x60; is supported &lt;br&gt; See Pegged Orders | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderCancelReplaceResponse**](OrderCancelReplaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListCancel

> OrderListCancelResponse OrderListCancel().Symbol(symbol).Id(id).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order list (TRADE)


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
	orderListId := int32(1) // int32 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "C3wyj4WVEktd7u9aVBRXcN" // string | Either `orderListId` or `listClientOrderId` must be provided (optional)
	newClientOrderId := "cancelMyOrder1" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **id** | **string** | Client-generated request identifier. | 
 **orderListId** | **int32** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListCancelResponse**](OrderListCancelResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlace

> OrderListPlaceResponse OrderListPlace().Symbol(symbol).Side(side).Price(price).Quantity(quantity).Id(id).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).StopClientOrderId(stopClientOrderId).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).StopIcebergQty(stopIcebergQty).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

Place new OCO - Deprecated (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Please see [Enums](/products/spot/enums#side) for supported values.
	price := float64(1) // float64 | 
	quantity := float64(1) // float64 | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "08985fedd9ea2cf6b28996" // string | A unique Id for the entire orderList (optional)
	limitClientOrderId := "limitOrder1" // string | A unique Id for the limit order (optional)
	limitIcebergQty := float64(1) // float64 | Used to make the `LIMIT_MAKER` leg an iceberg order. (optional)
	limitStrategyId := int64(1) // int64 |  (optional)
	limitStrategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	stopPrice := float64(1) // float64 |  (optional)
	trailingDelta := int32(1) // int32 |  (optional)
	stopClientOrderId := "stopOrder1" // string | A unique Id for the stop loss/stop loss limit leg (optional)
	stopLimitPrice := float64(1) // float64 | If provided, `stopLimitTimeInForce` is required. (optional)
	stopLimitTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Valid values are `GTC`/`FOK`/`IOC` (optional)
	stopIcebergQty := float64(1) // float64 | Used with `STOP_LOSS_LIMIT` leg to make an iceberg order. (optional)
	stopStrategyId := int64(1) // int64 |  (optional)
	stopStrategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed values are dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Please see [Enums](/products/spot/enums#side) for supported values. | 
 **price** | **float64** |  | 
 **quantity** | **float64** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitIcebergQty** | **float64** | Used to make the &#x60;LIMIT_MAKER&#x60; leg an iceberg order. | 
 **limitStrategyId** | **int64** |  | 
 **limitStrategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **stopPrice** | **float64** |  | 
 **trailingDelta** | **int32** |  | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopLimitPrice** | **float64** | If provided, &#x60;stopLimitTimeInForce&#x60; is required. | 
 **stopLimitTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Valid values are &#x60;GTC&#x60;/&#x60;FOK&#x60;/&#x60;IOC&#x60; | 
 **stopIcebergQty** | **float64** | Used with &#x60;STOP_LOSS_LIMIT&#x60; leg to make an iceberg order. | 
 **stopStrategyId** | **int64** |  | 
 **stopStrategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed values are dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceResponse**](OrderListPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOco

> OrderListPlaceOcoResponse OrderListPlaceOco().Symbol(symbol).Side(side).Quantity(quantity).AboveType(aboveType).BelowType(belowType).Id(id).ListClientOrderId(listClientOrderId).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveTrailingDelta(aboveTrailingDelta).AboveTimeInForce(aboveTimeInForce).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AbovePegPriceType(abovePegPriceType).AbovePegOffsetType(abovePegOffsetType).AbovePegOffsetValue(abovePegOffsetValue).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowTrailingDelta(belowTrailingDelta).BelowTimeInForce(belowTimeInForce).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowPegPriceType(belowPegPriceType).BelowPegOffsetType(belowPegOffsetType).BelowPegOffsetValue(belowPegOffsetValue).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

Place new Order list - OCO (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | `BUY` or `SELL`
	quantity := float64(1) // float64 | Quantity for both orders of the order list.
	aboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | 
	belowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter | Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "cKPMnDCbcLQILtDYM4f4fX" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `aboveClientOrderId` and the `belowClientOrderId`. (optional)
	aboveClientOrderId := "aboveOrder1" // string | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent. (optional)
	aboveIcebergQty := int64(1) // int64 | Note that this can only be used if `aboveTimeInForce` is `GTC`. (optional)
	abovePrice := float64(1) // float64 | Can be used if `aboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	aboveStopPrice := float64(1) // float64 | Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified. (optional)
	aboveTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	aboveTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Required if `aboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`. (optional)
	aboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the above order within an order strategy. (optional)
	aboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the above order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	abovePegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | `PRIMARY_PEG` or `MARKET_PEG`. See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	abovePegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	abovePegOffsetValue := int32(1) // int32 |  (optional)
	belowClientOrderId := "belowOrder1" // string | Arbitrary unique ID among open orders for the below order. Automatically generated if not sent. (optional)
	belowIcebergQty := int64(1) // int64 | Note that this can only be used if `belowTimeInForce` is `GTC`. (optional)
	belowPrice := float64(1) // float64 | Can be used if `belowType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	belowStopPrice := float64(1) // float64 | Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `belowStopPrice` or `belowTrailingDelta` or both, must be specified. (optional)
	belowTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	belowTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Required if `belowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`. (optional)
	belowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the below order within an order strategy. (optional)
	belowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the below order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	belowPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	belowPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	belowPegOffsetValue := int32(1) // int32 |  (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Select response format: `ACK`, `RESULT`, `FULL`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | &#x60;BUY&#x60; or &#x60;SELL&#x60; | 
 **quantity** | **float64** | Quantity for both orders of the order list. | 
 **aboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) |  | 
 **belowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) | Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60; | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;aboveClientOrderId&#x60; and the &#x60;belowClientOrderId&#x60;. | 
 **aboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent. | 
 **aboveIcebergQty** | **int64** | Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **abovePrice** | **float64** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **aboveStopPrice** | **float64** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified. | 
 **aboveTrailingDelta** | **int64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **aboveTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Required if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **aboveStrategyId** | **int64** | Arbitrary numeric value identifying the above order within an order strategy. | 
 **aboveStrategyType** | **int32** | Arbitrary numeric value identifying the above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **abovePegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60;. See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **abovePegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **abovePegOffsetValue** | **int32** |  | 
 **belowClientOrderId** | **string** | Arbitrary unique ID among open orders for the below order. Automatically generated if not sent. | 
 **belowIcebergQty** | **int64** | Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **belowPrice** | **float64** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **belowStopPrice** | **float64** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;belowStopPrice&#x60; or &#x60;belowTrailingDelta&#x60; or both, must be specified. | 
 **belowTrailingDelta** | **int64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **belowTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Required if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **belowStrategyId** | **int64** | Arbitrary numeric value identifying the below order within an order strategy. | 
 **belowStrategyType** | **int32** | Arbitrary numeric value identifying the below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **belowPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **belowPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **belowPegOffsetValue** | **int32** |  | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Select response format: &#x60;ACK&#x60;, &#x60;RESULT&#x60;, &#x60;FULL&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOcoResponse**](OrderListPlaceOcoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOpo

> OrderListPlaceOpoResponse OrderListPlaceOpo().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

OPO (TRADE)


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | Supported values: `LIMIT`, `LIMIT_MAKER`
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	workingPrice := float64(1) // float64 | Price for the working order.
	workingQuantity := float64(1) // float64 | Sets the quantity for the working order.
	pendingType := models.OrderListPlaceOpoPendingTypeParameterLimit // OrderListPlaceOpoPendingTypeParameter | Supported values: [Order Types](/products/spot/enums#ordertypes). Note that `MARKET` orders using `quoteOrderQty` are not supported.
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "OiOgqvRagBefpzdM5gjYX3" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	workingClientOrderId := "workingOrder1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float64(1) // float64 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. (optional)
	workingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy. (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	workingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	workingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingOrder1" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. (optional)
	pendingPrice := float64(1) // float64 | Price for the pending order. (optional)
	pendingStopPrice := float64(1) // float64 | Stop price for the pending order. (optional)
	pendingTrailingDelta := float64(1) // float64 | Trailing delta for the pending order. (optional)
	pendingIcebergQty := float64(1) // float64 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`. (optional)
	pendingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy. (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) | Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60; | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **workingPrice** | **float64** | Price for the working order. | 
 **workingQuantity** | **float64** | Sets the quantity for the working order. | 
 **pendingType** | [**OrderListPlaceOpoPendingTypeParameter**](OrderListPlaceOpoPendingTypeParameter.md) | Supported values: [Order Types](/products/spot/enums#ordertypes). Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported. | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **workingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy. | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **workingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **workingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. | 
 **pendingPrice** | **float64** | Price for the pending order. | 
 **pendingStopPrice** | **float64** | Stop price for the pending order. | 
 **pendingTrailingDelta** | **float64** | Trailing delta for the pending order. | 
 **pendingIcebergQty** | **float64** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy. | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOpoResponse**](OrderListPlaceOpoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOpoco

> OrderListPlaceOpocoResponse OrderListPlaceOpoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

OPOCO (TRADE)


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | 
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	workingPrice := float64(1) // float64 | Price for the working order.
	workingQuantity := float64(1) // float64 | Sets the quantity for the working order.
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	pendingAboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | Supported values: `STOP_LOSS_LIMIT`, `STOP_LOSS`, `LIMIT_MAKER`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "TVbG6ymkYMXTj7tczbOsBf" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	workingClientOrderId := "workingOrder1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float64(1) // float64 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. (optional)
	workingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy. (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	workingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	workingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	workingPegOffsetValue := int32(1) // int32 | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingAboveClientOrderId := "pendingAboveOrder1" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. (optional)
	pendingAbovePrice := float64(1) // float64 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	pendingAboveStopPrice := float64(1) // float64 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. (optional)
	pendingAboveTrailingDelta := float64(1) // float64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	pendingAboveIcebergQty := float64(1) // float64 | This can only be used if `pendingAboveTimeInForce` is `GTC` or `pendingAboveType` is `LIMIT_MAKER`. (optional)
	pendingAboveTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Required if `pendingAboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`. (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy. (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingAbovePegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingAbovePegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingBelowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter | Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT` (optional)
	pendingBelowClientOrderId := "pendingBelowOrder1" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. (optional)
	pendingBelowPrice := float64(1) // float64 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	pendingBelowStopPrice := float64(1) // float64 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified. (optional)
	pendingBelowTrailingDelta := float64(1) // float64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	pendingBelowIcebergQty := float64(1) // float64 | This can only be used if `pendingBelowTimeInForce` is `GTC` or `pendingBelowType` is `LIMIT_MAKER`. (optional)
	pendingBelowTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy. (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingBelowPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingBelowPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **workingPrice** | **float64** | Price for the working order. | 
 **workingQuantity** | **float64** | Sets the quantity for the working order. | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **pendingAboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) | Supported values: &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;STOP_LOSS&#x60;, &#x60;LIMIT_MAKER&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60; | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **workingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy. | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **workingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **workingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **workingPegOffsetValue** | **int32** | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. | 
 **pendingAbovePrice** | **float64** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **pendingAboveStopPrice** | **float64** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **pendingAboveTrailingDelta** | **float64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **pendingAboveIcebergQty** | **float64** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingAboveTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Required if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy. | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingAbovePegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingAbovePegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingAbovePegOffsetValue** | **int32** | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingBelowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) | Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60; | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. | 
 **pendingBelowPrice** | **float64** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **pendingBelowStopPrice** | **float64** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified. | 
 **pendingBelowTrailingDelta** | **float64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **pendingBelowIcebergQty** | **float64** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60; or &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingBelowTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy. | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingBelowPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingBelowPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOpocoResponse**](OrderListPlaceOpocoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOto

> OrderListPlaceOtoResponse OrderListPlaceOto().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegOffsetType(pendingPegOffsetType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

Place new Order list - OTO (TRADE)


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | Supported values: `LIMIT`, `LIMIT_MAKER`
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	workingPrice := float64(1) // float64 | 
	workingQuantity := float64(1) // float64 | Sets the quantity for the working order.
	pendingType := models.OrderListPlaceOpoPendingTypeParameterLimit // OrderListPlaceOpoPendingTypeParameter | Supported values: [Order Types](/products/spot/enums#ordertypes). Note that `MARKET` orders using `quoteOrderQty` are not supported.
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	pendingQuantity := float64(1) // float64 | Sets the quantity for the pending order.
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "KA4EBjGnzvSwSCQsDdTrlf" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	workingClientOrderId := "workingOrder1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float64(1) // float64 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. (optional)
	workingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy. (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	workingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	workingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingOrder1" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. (optional)
	pendingPrice := float64(1) // float64 |  (optional)
	pendingStopPrice := float64(1) // float64 |  (optional)
	pendingTrailingDelta := float64(1) // float64 |  (optional)
	pendingIcebergQty := float64(1) // float64 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`. (optional)
	pendingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy. (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	pendingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) | Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60; | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **workingPrice** | **float64** |  | 
 **workingQuantity** | **float64** | Sets the quantity for the working order. | 
 **pendingType** | [**OrderListPlaceOpoPendingTypeParameter**](OrderListPlaceOpoPendingTypeParameter.md) | Supported values: [Order Types](/products/spot/enums#ordertypes). Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported. | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **pendingQuantity** | **float64** | Sets the quantity for the pending order. | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **workingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy. | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **workingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **workingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. | 
 **pendingPrice** | **float64** |  | 
 **pendingStopPrice** | **float64** |  | 
 **pendingTrailingDelta** | **float64** |  | 
 **pendingIcebergQty** | **float64** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy. | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **pendingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOtoResponse**](OrderListPlaceOtoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderListPlaceOtoco

> OrderListPlaceOtocoResponse OrderListPlaceOtoco().Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).Id(id).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

Place new Order list - OTOCO (TRADE)


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
	workingType := models.OrderListPlaceOpoWorkingTypeParameterLimit // OrderListPlaceOpoWorkingTypeParameter | Supported values: `LIMIT`, `LIMIT_MAKER`
	workingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	workingPrice := float64(1) // float64 | 
	workingQuantity := float64(1) // float64 | Sets the quantity for the working order.
	pendingSide := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Supported values: [Order Side](/products/spot/enums#side)
	pendingQuantity := float64(1) // float64 | Sets the quantity for the pending orders.
	pendingAboveType := models.OrderListPlaceOcoAboveTypeParameterStopLossLimit // OrderListPlaceOcoAboveTypeParameter | Supported values: `STOP_LOSS_LIMIT`, `STOP_LOSS`, `LIMIT_MAKER`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	listClientOrderId := "GaeJHjZPasPItFj4x7Mqm6" // string | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	workingClientOrderId := "workingOrder1" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. (optional)
	workingIcebergQty := float64(1) // float64 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. (optional)
	workingTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Supported values: [Time In Force](/products/spot/enums#timeinforce) (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy. (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	workingPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	workingPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingAboveClientOrderId := "pendingAboveOrder1" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. (optional)
	pendingAbovePrice := float64(1) // float64 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	pendingAboveStopPrice := float64(1) // float64 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. (optional)
	pendingAboveTrailingDelta := float64(1) // float64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	pendingAboveIcebergQty := float64(1) // float64 | This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`. (optional)
	pendingAboveTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Required if `pendingAboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`. (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy. (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingAbovePegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingAbovePegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 |  (optional)
	pendingBelowType := models.OrderListPlaceOcoBelowTypeParameterStopLoss // OrderListPlaceOcoBelowTypeParameter | Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT` (optional)
	pendingBelowClientOrderId := "pendingBelowOrder1" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. (optional)
	pendingBelowPrice := float64(1) // float64 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	pendingBelowStopPrice := float64(1) // float64 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified. (optional)
	pendingBelowTrailingDelta := float64(1) // float64 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	pendingBelowIcebergQty := float64(1) // float64 | This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`. (optional)
	pendingBelowTimeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Required if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`. (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy. (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than `1000000` are reserved and cannot be used. (optional)
	pendingBelowPegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pendingBelowPegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **workingType** | [**OrderListPlaceOpoWorkingTypeParameter**](OrderListPlaceOpoWorkingTypeParameter.md) | Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60; | 
 **workingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **workingPrice** | **float64** |  | 
 **workingQuantity** | **float64** | Sets the quantity for the working order. | 
 **pendingSide** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Supported values: [Order Side](/products/spot/enums#side) | 
 **pendingQuantity** | **float64** | Sets the quantity for the pending orders. | 
 **pendingAboveType** | [**OrderListPlaceOcoAboveTypeParameter**](OrderListPlaceOcoAboveTypeParameter.md) | Supported values: &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;STOP_LOSS&#x60;, &#x60;LIMIT_MAKER&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60; | 
 **id** | **string** | Client-generated request identifier. | 
 **listClientOrderId** | **string** | Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype) | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. | 
 **workingIcebergQty** | **float64** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **workingTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Supported values: [Time In Force](/products/spot/enums#timeinforce) | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy. | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **workingPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **workingPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. | 
 **pendingAbovePrice** | **float64** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **pendingAboveStopPrice** | **float64** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **pendingAboveTrailingDelta** | **float64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **pendingAboveIcebergQty** | **float64** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingAboveTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Required if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy. | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingAbovePegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingAbovePegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **pendingAbovePegOffsetValue** | **int32** |  | 
 **pendingBelowType** | [**OrderListPlaceOcoBelowTypeParameter**](OrderListPlaceOcoBelowTypeParameter.md) | Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60; | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. | 
 **pendingBelowPrice** | **float64** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **pendingBelowStopPrice** | **float64** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified. | 
 **pendingBelowTrailingDelta** | **float64** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **pendingBelowIcebergQty** | **float64** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;. | 
 **pendingBelowTimeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Required if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy. | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used. | 
 **pendingBelowPegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pendingBelowPegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListPlaceOtocoResponse**](OrderListPlaceOtocoResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderPlace

> OrderPlaceResponse OrderPlace().Symbol(symbol).Side(side).Type(type_).Id(id).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

Place new order (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Please see [Enums](/products/spot/enums#side) for supported values.
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | Please see [Enums](/products/spot/enums#ordertypes) for supported values.
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Please see [Enums](/products/spot/enums#timeinforce) for supported values. (optional)
	price := float64(1) // float64 |  (optional)
	quantity := float64(1) // float64 |  (optional)
	quoteOrderQty := float64(1) // float64 |  (optional)
	newClientOrderId := "myOrder1" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | `MARKET` and `LIMIT` order types default to `FULL`, all other orders default to `ACK`. (optional)
	stopPrice := float64(1) // float64 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := int32(1) // int32 | See Trailing Stop order FAQ (optional)
	icebergQty := float64(1) // float64 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | See Pegged Orders Info (optional)
	pegOffsetValue := int32(1) // int32 | Price level to peg the price to (max: 100). See Pegged Orders Info (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter | Only `PRICE_LEVEL` is supported. See Pegged Orders Info (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Please see [Enums](/products/spot/enums#side) for supported values. | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) | Please see [Enums](/products/spot/enums#ordertypes) for supported values. | 
 **id** | **string** | Client-generated request identifier. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Please see [Enums](/products/spot/enums#timeinforce) for supported values. | 
 **price** | **float64** |  | 
 **quantity** | **float64** |  | 
 **quoteOrderQty** | **float64** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; order types default to &#x60;FULL&#x60;, all other orders default to &#x60;ACK&#x60;. | 
 **stopPrice** | **float64** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **int32** | See Trailing Stop order FAQ | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | See Pegged Orders Info | 
 **pegOffsetValue** | **int32** | Price level to peg the price to (max: 100). See Pegged Orders Info | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) | Only &#x60;PRICE_LEVEL&#x60; is supported. See Pegged Orders Info | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderPlaceResponse**](OrderPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## OrderTest

> OrderTestResponse OrderTest().Symbol(symbol).Side(side).Type(type_).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

Test new order (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Please see [Enums](/products/spot/enums#side) for supported values.
	type_ := models.OrderCancelReplaceTypeParameterMarket // OrderCancelReplaceTypeParameter | Please see [Enums](/products/spot/enums#ordertypes) for supported values.
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	computeCommissionRates := false // bool | Default: `false` <br> See [Commissions FAQ](/products/spot/faqs/commission_faq#test-order-diferences) to learn more. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Please see [Enums](/products/spot/enums#timeinforce) for supported values. (optional)
	price := float64(1) // float64 |  (optional)
	quantity := float64(1) // float64 |  (optional)
	quoteOrderQty := float64(1) // float64 |  (optional)
	newClientOrderId := "myOrder1" // string | A unique id among open orders. Automatically generated if not sent. Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Set the response JSON. `ACK`, `RESULT`, or `FULL`; `MARKET` and `LIMIT` order types default to `FULL`, all other orders default to `ACK`. (optional)
	stopPrice := float64(1) // float64 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := int32(1) // int32 | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) (optional)
	icebergQty := float64(1) // float64 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	pegPriceType := models.OrderCancelReplacePegPriceTypeParameterPrimaryPeg // OrderCancelReplacePegPriceTypeParameter | `PRIMARY_PEG` or `MARKET_PEG`. See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pegOffsetValue := int32(1) // int32 | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	pegOffsetType := models.OrderCancelReplacePegOffsetTypeParameterPriceLevel // OrderCancelReplacePegOffsetTypeParameter | Only `PRICE_LEVEL` is supported. See [Pegged Orders](/products/spot/faqs/pegged_orders) (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Please see [Enums](/products/spot/enums#side) for supported values. | 
 **type_** | [**OrderCancelReplaceTypeParameter**](OrderCancelReplaceTypeParameter.md) | Please see [Enums](/products/spot/enums#ordertypes) for supported values. | 
 **id** | **string** | Client-generated request identifier. | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](/products/spot/faqs/commission_faq#test-order-diferences) to learn more. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Please see [Enums](/products/spot/enums#timeinforce) for supported values. | 
 **price** | **float64** |  | 
 **quantity** | **float64** |  | 
 **quoteOrderQty** | **float64** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;; &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; order types default to &#x60;FULL&#x60;, all other orders default to &#x60;ACK&#x60;. | 
 **stopPrice** | **float64** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **int32** | See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq) | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **pegPriceType** | [**OrderCancelReplacePegPriceTypeParameter**](OrderCancelReplacePegPriceTypeParameter.md) | &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60;. See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pegOffsetValue** | **int32** | Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **pegOffsetType** | [**OrderCancelReplacePegOffsetTypeParameter**](OrderCancelReplacePegOffsetTypeParameter.md) | Only &#x60;PRICE_LEVEL&#x60; is supported. See [Pegged Orders](/products/spot/faqs/pegged_orders) | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderTestResponse**](OrderTestResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SorOrderPlace

> SorOrderPlaceResponse SorOrderPlace().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

Place new order using SOR (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | `BUY` or `SELL`
	type_ := models.SorOrderPlaceTypeParameterMarket // SorOrderPlaceTypeParameter | Only `LIMIT` and `MARKET` orders are supported.
	quantity := float64(1) // float64 | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Applicable only to `LIMIT` order type. (optional)
	price := float64(1) // float64 |  (optional)
	newClientOrderId := "myOrder1" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Set the response JSON. `ACK`, `RESULT`, or `FULL`. Default to `FULL` (optional)
	icebergQty := float64(1) // float64 | Used with `LIMIT` to create an iceberg order. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes). (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | &#x60;BUY&#x60; or &#x60;SELL&#x60; | 
 **type_** | [**SorOrderPlaceTypeParameter**](SorOrderPlaceTypeParameter.md) | Only &#x60;LIMIT&#x60; and &#x60;MARKET&#x60; orders are supported. | 
 **quantity** | **float64** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Applicable only to &#x60;LIMIT&#x60; order type. | 
 **price** | **float64** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;. Default to &#x60;FULL&#x60; | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60; to create an iceberg order. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes). | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderPlaceResponse**](SorOrderPlaceResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## SorOrderTest

> SorOrderTestResponse SorOrderTest().Symbol(symbol).Side(side).Type(type_).Quantity(quantity).Id(id).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).IcebergQty(icebergQty).StrategyId(strategyId).StrategyType(strategyType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

Test new order using SOR (TRADE)


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
	side := models.OrderCancelReplaceSideParameterBuy // OrderCancelReplaceSideParameter | Please see [Enums](/products/spot/enums#side) for supported values.
	type_ := models.SorOrderPlaceTypeParameterMarket // SorOrderPlaceTypeParameter | Please see [Enums](/products/spot/enums#ordertypes) for supported values.
	quantity := float64(1) // float64 | 
	id := "7d3b9b46-5f4f-4c8b-9a2d-0a8f9a4a0f5b" // string | Client-generated request identifier. (optional)
	computeCommissionRates := false // bool | Default: `false` (optional)
	timeInForce := models.OrderCancelReplaceTimeInForceParameterGtc // OrderCancelReplaceTimeInForceParameter | Please see [Enums](/products/spot/enums#timeinforce) for supported values. (optional)
	price := float64(1) // float64 |  (optional)
	newClientOrderId := "myOrder1" // string | A unique id among open orders. Automatically generated if not sent. Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	newOrderRespType := models.OrderCancelReplaceNewOrderRespTypeParameterAck // OrderCancelReplaceNewOrderRespTypeParameter | Set the response JSON. `ACK`, `RESULT`, or `FULL`. Default to `FULL`. (optional)
	icebergQty := float64(1) // float64 | Used with `LIMIT` to create an iceberg order. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	selfTradePreventionMode := models.OrderCancelReplaceSelfTradePreventionModeParameterNone // OrderCancelReplaceSelfTradePreventionModeParameter | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) (optional)
	recvWindow := float64(5000) // float64 | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

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
 **side** | [**OrderCancelReplaceSideParameter**](OrderCancelReplaceSideParameter.md) | Please see [Enums](/products/spot/enums#side) for supported values. | 
 **type_** | [**SorOrderPlaceTypeParameter**](SorOrderPlaceTypeParameter.md) | Please see [Enums](/products/spot/enums#ordertypes) for supported values. | 
 **quantity** | **float64** |  | 
 **id** | **string** | Client-generated request identifier. | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; | 
 **timeInForce** | [**OrderCancelReplaceTimeInForceParameter**](OrderCancelReplaceTimeInForceParameter.md) | Please see [Enums](/products/spot/enums#timeinforce) for supported values. | 
 **price** | **float64** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **newOrderRespType** | [**OrderCancelReplaceNewOrderRespTypeParameter**](OrderCancelReplaceNewOrderRespTypeParameter.md) | Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;. Default to &#x60;FULL&#x60;. | 
 **icebergQty** | **float64** | Used with &#x60;LIMIT&#x60; to create an iceberg order. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **selfTradePreventionMode** | [**OrderCancelReplaceSelfTradePreventionModeParameter**](OrderCancelReplaceSelfTradePreventionModeParameter.md) | The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes) | 
 **recvWindow** | **float64** | Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderTestResponse**](SorOrderTestResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

