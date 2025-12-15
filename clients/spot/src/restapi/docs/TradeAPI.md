# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**DeleteOpenOrders**](TradeAPI.md#DeleteOpenOrders) | **Delete** /api/v3/openOrders | Cancel All Open Orders on a Symbol
[**DeleteOrder**](TradeAPI.md#DeleteOrder) | **Delete** /api/v3/order | Cancel order
[**DeleteOrderList**](TradeAPI.md#DeleteOrderList) | **Delete** /api/v3/orderList | Cancel Order list
[**NewOrder**](TradeAPI.md#NewOrder) | **Post** /api/v3/order | New order
[**OrderAmendKeepPriority**](TradeAPI.md#OrderAmendKeepPriority) | **Put** /api/v3/order/amend/keepPriority | Order Amend Keep Priority
[**OrderCancelReplace**](TradeAPI.md#OrderCancelReplace) | **Post** /api/v3/order/cancelReplace | Cancel an Existing Order and Send a New Order
[**OrderListOco**](TradeAPI.md#OrderListOco) | **Post** /api/v3/orderList/oco | New Order list - OCO
[**OrderListOpo**](TradeAPI.md#OrderListOpo) | **Post** /api/v3/orderList/opo | New Order List - OPO
[**OrderListOpoco**](TradeAPI.md#OrderListOpoco) | **Post** /api/v3/orderList/opoco | New Order List - OPOCO
[**OrderListOto**](TradeAPI.md#OrderListOto) | **Post** /api/v3/orderList/oto | New Order list - OTO
[**OrderListOtoco**](TradeAPI.md#OrderListOtoco) | **Post** /api/v3/orderList/otoco | New Order list - OTOCO
[**OrderOco**](TradeAPI.md#OrderOco) | **Post** /api/v3/order/oco | New OCO - Deprecated
[**OrderTest**](TradeAPI.md#OrderTest) | **Post** /api/v3/order/test | Test new order
[**SorOrder**](TradeAPI.md#SorOrder) | **Post** /api/v3/sor/order | New order using SOR
[**SorOrderTest**](TradeAPI.md#SorOrderTest) | **Post** /api/v3/sor/order/test | Test new order using SOR


## DeleteOpenOrders

> DeleteOpenOrdersResponse DeleteOpenOrders(ctx).Symbol(symbol).RecvWindow(recvWindow).Execute()

Cancel All Open Orders on a Symbol


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

	resp, err := apiClient.RestApi.TradeAPI.DeleteOpenOrders(context.Background()).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.DeleteOpenOrders``: %v\n", err)
		return
	}

	// response from `DeleteOpenOrders`: DeleteOpenOrdersResponse
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

[**DeleteOpenOrdersResponse**](DeleteOpenOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteOrder

> DeleteOrderResponse DeleteOrder(ctx).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()

Cancel order


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
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	cancelRestrictions := models.DeleteOrderCancelRestrictionsParameterOnlyNew // DeleteOrderCancelRestrictionsParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.DeleteOrder(context.Background()).Symbol(symbol).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).CancelRestrictions(cancelRestrictions).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.DeleteOrder``: %v\n", err)
		return
	}

	// response from `DeleteOrder`: DeleteOrderResponse
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
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **cancelRestrictions** | [**DeleteOrderCancelRestrictionsParameter**](DeleteOrderCancelRestrictionsParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**DeleteOrderResponse**](DeleteOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DeleteOrderList

> DeleteOrderListResponse DeleteOrderList(ctx).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order list


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
	orderListId := int64(1) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.DeleteOrderList(context.Background()).Symbol(symbol).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.DeleteOrderList``: %v\n", err)
		return
	}

	// response from `DeleteOrderList`: DeleteOrderListResponse
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
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**DeleteOrderListResponse**](DeleteOrderListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder(ctx).Symbol(symbol).Side(side).Type(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

New order


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	type_ := models.NewOrderTypeParameterMarket // NewOrderTypeParameter | 
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	price := float32(400.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	pegPriceType := models.NewOrderPegPriceTypeParameterPrimaryPeg // NewOrderPegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info) (optional)
	pegOffsetType := models.NewOrderPegOffsetTypeParameterPriceLevel // NewOrderPegOffsetTypeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **stopPrice** | **float32** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **pegPriceType** | [**NewOrderPegPriceTypeParameter**](NewOrderPegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info) | 
 **pegOffsetType** | [**NewOrderPegOffsetTypeParameter**](NewOrderPegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**NewOrderResponse**](NewOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderAmendKeepPriority

> OrderAmendKeepPriorityResponse OrderAmendKeepPriority(ctx).Symbol(symbol).NewQty(newQty).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Order Amend Keep Priority


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
	newQty := float32(1.0) // float32 | `newQty` must be greater than 0 and less than the order's quantity.
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderAmendKeepPriority(context.Background()).Symbol(symbol).NewQty(newQty).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderAmendKeepPriority``: %v\n", err)
		return
	}

	// response from `OrderAmendKeepPriority`: OrderAmendKeepPriorityResponse
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
 **newQty** | **float32** | &#x60;newQty&#x60; must be greater than 0 and less than the order&#39;s quantity. | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderAmendKeepPriorityResponse**](OrderAmendKeepPriorityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderCancelReplace

> OrderCancelReplaceResponse OrderCancelReplace(ctx).Symbol(symbol).Side(side).Type(type_).CancelReplaceMode(cancelReplaceMode).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).CancelNewClientOrderId(cancelNewClientOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelOrderId(cancelOrderId).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).CancelRestrictions(cancelRestrictions).OrderRateLimitExceededMode(orderRateLimitExceededMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

Cancel an Existing Order and Send a New Order


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	type_ := models.NewOrderTypeParameterMarket // NewOrderTypeParameter | 
	cancelReplaceMode := models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure // OrderCancelReplaceCancelReplaceModeParameter | 
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	price := float32(400.0) // float32 |  (optional)
	cancelNewClientOrderId := "cancelNewClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional)
	cancelOrigClientOrderId := "cancelOrigClientOrderId_example" // string | Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br> If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br> If both conditions are not met the request will be rejected. (optional)
	cancelOrderId := int64(1) // int64 | Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br>If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br>If both conditions are not met the request will be rejected. (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	cancelRestrictions := models.DeleteOrderCancelRestrictionsParameterOnlyNew // DeleteOrderCancelRestrictionsParameter |  (optional)
	orderRateLimitExceededMode := models.OrderCancelReplaceOrderRateLimitExceededModeParameterDoNothing // OrderCancelReplaceOrderRateLimitExceededModeParameter |  (optional)
	pegPriceType := models.NewOrderPegPriceTypeParameterPrimaryPeg // NewOrderPegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info) (optional)
	pegOffsetType := models.NewOrderPegOffsetTypeParameterPriceLevel // NewOrderPegOffsetTypeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderCancelReplace(context.Background()).Symbol(symbol).Side(side).Type(type_).CancelReplaceMode(cancelReplaceMode).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).CancelNewClientOrderId(cancelNewClientOrderId).CancelOrigClientOrderId(cancelOrigClientOrderId).CancelOrderId(cancelOrderId).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).CancelRestrictions(cancelRestrictions).OrderRateLimitExceededMode(orderRateLimitExceededMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderCancelReplace``: %v\n", err)
		return
	}

	// response from `OrderCancelReplace`: OrderCancelReplaceResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **cancelReplaceMode** | [**OrderCancelReplaceCancelReplaceModeParameter**](OrderCancelReplaceCancelReplaceModeParameter.md) |  | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **price** | **float32** |  | 
 **cancelNewClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | 
 **cancelOrigClientOrderId** | **string** | Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt; If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt; If both conditions are not met the request will be rejected. | 
 **cancelOrderId** | **int64** | Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt;If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt;If both conditions are not met the request will be rejected. | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **stopPrice** | **float32** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **cancelRestrictions** | [**DeleteOrderCancelRestrictionsParameter**](DeleteOrderCancelRestrictionsParameter.md) |  | 
 **orderRateLimitExceededMode** | [**OrderCancelReplaceOrderRateLimitExceededModeParameter**](OrderCancelReplaceOrderRateLimitExceededModeParameter.md) |  | 
 **pegPriceType** | [**NewOrderPegPriceTypeParameter**](NewOrderPegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info) | 
 **pegOffsetType** | [**NewOrderPegOffsetTypeParameter**](NewOrderPegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderCancelReplaceResponse**](OrderCancelReplaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderListOco

> OrderListOcoResponse OrderListOco(ctx).Symbol(symbol).Side(side).Quantity(quantity).AboveType(aboveType).BelowType(belowType).ListClientOrderId(listClientOrderId).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveTrailingDelta(aboveTrailingDelta).AboveTimeInForce(aboveTimeInForce).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AbovePegPriceType(abovePegPriceType).AbovePegOffsetType(abovePegOffsetType).AbovePegOffsetValue(abovePegOffsetValue).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowTrailingDelta(belowTrailingDelta).BelowTimeInForce(belowTimeInForce).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowPegPriceType(belowPegPriceType).BelowPegOffsetType(belowPegOffsetType).BelowPegOffsetValue(belowPegOffsetValue).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New Order list - OCO


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	quantity := float32(1.0) // float32 | 
	aboveType := models.OrderListOcoAboveTypeParameterStopLossLimit // OrderListOcoAboveTypeParameter | 
	belowType := models.OrderListOcoBelowTypeParameterStopLoss // OrderListOcoBelowTypeParameter | 
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	aboveClientOrderId := "aboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent (optional)
	aboveIcebergQty := int64(1) // int64 | Note that this can only be used if `aboveTimeInForce` is `GTC`. (optional)
	abovePrice := float32(1.0) // float32 | Can be used if `aboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	aboveStopPrice := float32(1.0) // float32 | Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. <br>Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified. (optional)
	aboveTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	aboveTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	aboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the above order within an order strategy. (optional)
	aboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the above order strategy. <br>Values smaller than 1000000 are reserved and cannot be used. (optional)
	abovePegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	abovePegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	abovePegOffsetValue := int32(1) // int32 |  (optional)
	belowClientOrderId := "belowClientOrderId_example" // string | Arbitrary unique ID among open orders for the below order. Automatically generated if not sent (optional)
	belowIcebergQty := int64(1) // int64 | Note that this can only be used if `belowTimeInForce` is `GTC`. (optional)
	belowPrice := float32(1.0) // float32 | Can be used if `belowType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. (optional)
	belowStopPrice := float32(1.0) // float32 | Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT` or `TAKE_PROFIT_LIMIT` <br>Either belowStopPrice or belowTrailingDelta or both, must be specified. (optional)
	belowTrailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	belowTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	belowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the below order within an order strategy. (optional)
	belowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the below order strategy. <br>Values smaller than 1000000 are reserved and cannot be used. (optional)
	belowPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	belowPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	belowPegOffsetValue := int32(1) // int32 |  (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderListOco(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).AboveType(aboveType).BelowType(belowType).ListClientOrderId(listClientOrderId).AboveClientOrderId(aboveClientOrderId).AboveIcebergQty(aboveIcebergQty).AbovePrice(abovePrice).AboveStopPrice(aboveStopPrice).AboveTrailingDelta(aboveTrailingDelta).AboveTimeInForce(aboveTimeInForce).AboveStrategyId(aboveStrategyId).AboveStrategyType(aboveStrategyType).AbovePegPriceType(abovePegPriceType).AbovePegOffsetType(abovePegOffsetType).AbovePegOffsetValue(abovePegOffsetValue).BelowClientOrderId(belowClientOrderId).BelowIcebergQty(belowIcebergQty).BelowPrice(belowPrice).BelowStopPrice(belowStopPrice).BelowTrailingDelta(belowTrailingDelta).BelowTimeInForce(belowTimeInForce).BelowStrategyId(belowStrategyId).BelowStrategyType(belowStrategyType).BelowPegPriceType(belowPegPriceType).BelowPegOffsetType(belowPegOffsetType).BelowPegOffsetValue(belowPegOffsetValue).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListOco``: %v\n", err)
		return
	}

	// response from `OrderListOco`: OrderListOcoResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **quantity** | **float32** |  | 
 **aboveType** | [**OrderListOcoAboveTypeParameter**](OrderListOcoAboveTypeParameter.md) |  | 
 **belowType** | [**OrderListOcoBelowTypeParameter**](OrderListOcoBelowTypeParameter.md) |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **aboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the above order. Automatically generated if not sent | 
 **aboveIcebergQty** | **int64** | Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **abovePrice** | **float32** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **aboveStopPrice** | **float32** | Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified. | 
 **aboveTrailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **aboveTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **aboveStrategyId** | **int64** | Arbitrary numeric value identifying the above order within an order strategy. | 
 **aboveStrategyType** | **int32** | Arbitrary numeric value identifying the above order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used. | 
 **abovePegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **abovePegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **abovePegOffsetValue** | **int32** |  | 
 **belowClientOrderId** | **string** | Arbitrary unique ID among open orders for the below order. Automatically generated if not sent | 
 **belowIcebergQty** | **int64** | Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;. | 
 **belowPrice** | **float32** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price. | 
 **belowStopPrice** | **float32** | Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; &lt;br&gt;Either belowStopPrice or belowTrailingDelta or both, must be specified. | 
 **belowTrailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **belowTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **belowStrategyId** | **int64** | Arbitrary numeric value identifying the below order within an order strategy. | 
 **belowStrategyType** | **int32** | Arbitrary numeric value identifying the below order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used. | 
 **belowPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **belowPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **belowPegOffsetValue** | **int32** |  | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListOcoResponse**](OrderListOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderListOpo

> OrderListOpoResponse OrderListOpo(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

New Order List - OPO


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
	workingType := models.OrderListOpoWorkingTypeParameterLimit // OrderListOpoWorkingTypeParameter | 
	workingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingType := models.OrderListOpoPendingTypeParameterLimit // OrderListOpoPendingTypeParameter | 
	pendingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  (optional)
	pendingPrice := float32(1.0) // float32 |  (optional)
	pendingStopPrice := float32(1.0) // float32 |  (optional)
	pendingTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingIcebergQty := float32(1.0) // float32 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.  (optional)
	pendingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy.  (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderListOpo(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListOpo``: %v\n", err)
		return
	}

	// response from `OrderListOpo`: OrderListOpoResponse
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
 **workingType** | [**OrderListOpoWorkingTypeParameter**](OrderListOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingType** | [**OrderListOpoPendingTypeParameter**](OrderListOpoPendingTypeParameter.md) |  | 
 **pendingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  | 
 **pendingPrice** | **float32** |  | 
 **pendingStopPrice** | **float32** |  | 
 **pendingTrailingDelta** | **float32** |  | 
 **pendingIcebergQty** | **float32** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy.  | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListOpoResponse**](OrderListOpoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderListOpoco

> OrderListOpocoResponse OrderListOpoco(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingAboveType(pendingAboveType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

New Order List - OPOCO


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
	workingType := models.OrderListOpoWorkingTypeParameterLimit // OrderListOpoWorkingTypeParameter | 
	workingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	pendingAboveType := models.OrderListOcoAboveTypeParameterStopLossLimit // OrderListOcoAboveTypeParameter | 
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  (optional)
	pendingAbovePrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.  (optional)
	pendingAboveStopPrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`  (optional)
	pendingAboveTrailingDelta := float32(1.0) // float32 | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  (optional)
	pendingAboveIcebergQty := float32(1.0) // float32 | This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.  (optional)
	pendingAboveTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy.  (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingAbovePegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingAbovePegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 |  (optional)
	pendingBelowType := models.OrderListOcoBelowTypeParameterStopLoss // OrderListOcoBelowTypeParameter |  (optional)
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  (optional)
	pendingBelowPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price  (optional)
	pendingBelowStopPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.  (optional)
	pendingBelowTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingBelowIcebergQty := float32(1.0) // float32 | This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.  (optional)
	pendingBelowTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy.  (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingBelowPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingBelowPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderListOpoco(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingAboveType(pendingAboveType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListOpoco``: %v\n", err)
		return
	}

	// response from `OrderListOpoco`: OrderListOpocoResponse
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
 **workingType** | [**OrderListOpoWorkingTypeParameter**](OrderListOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **pendingAboveType** | [**OrderListOcoAboveTypeParameter**](OrderListOcoAboveTypeParameter.md) |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  | 
 **pendingAbovePrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.  | 
 **pendingAboveStopPrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;  | 
 **pendingAboveTrailingDelta** | **float32** | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  | 
 **pendingAboveIcebergQty** | **float32** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingAboveTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy.  | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingAbovePegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingAbovePegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingAbovePegOffsetValue** | **int32** |  | 
 **pendingBelowType** | [**OrderListOcoBelowTypeParameter**](OrderListOcoBelowTypeParameter.md) |  | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  | 
 **pendingBelowPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price  | 
 **pendingBelowStopPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.  | 
 **pendingBelowTrailingDelta** | **float32** |  | 
 **pendingBelowIcebergQty** | **float32** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingBelowTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy.  | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingBelowPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingBelowPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListOpocoResponse**](OrderListOpocoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderListOto

> OrderListOtoResponse OrderListOto(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()

New Order list - OTO


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
	workingType := models.OrderListOpoWorkingTypeParameterLimit // OrderListOpoWorkingTypeParameter | 
	workingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingType := models.OrderListOpoPendingTypeParameterLimit // OrderListOpoPendingTypeParameter | 
	pendingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	pendingQuantity := float32(1.0) // float32 | Sets the quantity for the pending order.
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingClientOrderId := "pendingClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  (optional)
	pendingPrice := float32(1.0) // float32 |  (optional)
	pendingStopPrice := float32(1.0) // float32 |  (optional)
	pendingTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingIcebergQty := float32(1.0) // float32 | This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.  (optional)
	pendingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending order within an order strategy.  (optional)
	pendingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderListOto(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingType(pendingType).PendingSide(pendingSide).PendingQuantity(pendingQuantity).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingClientOrderId(pendingClientOrderId).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTrailingDelta(pendingTrailingDelta).PendingIcebergQty(pendingIcebergQty).PendingTimeInForce(pendingTimeInForce).PendingStrategyId(pendingStrategyId).PendingStrategyType(pendingStrategyType).PendingPegPriceType(pendingPegPriceType).PendingPegOffsetType(pendingPegOffsetType).PendingPegOffsetValue(pendingPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListOto``: %v\n", err)
		return
	}

	// response from `OrderListOto`: OrderListOtoResponse
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
 **workingType** | [**OrderListOpoWorkingTypeParameter**](OrderListOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingType** | [**OrderListOpoPendingTypeParameter**](OrderListOpoPendingTypeParameter.md) |  | 
 **pendingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **pendingQuantity** | **float32** | Sets the quantity for the pending order. | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.  | 
 **pendingPrice** | **float32** |  | 
 **pendingStopPrice** | **float32** |  | 
 **pendingTrailingDelta** | **float32** |  | 
 **pendingIcebergQty** | **float32** | This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingStrategyId** | **int64** | Arbitrary numeric value identifying the pending order within an order strategy.  | 
 **pendingStrategyType** | **int32** | Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListOtoResponse**](OrderListOtoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderListOtoco

> OrderListOtocoResponse OrderListOtoco(ctx).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()

New Order list - OTOCO


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
	workingType := models.OrderListOpoWorkingTypeParameterLimit // OrderListOpoWorkingTypeParameter | 
	workingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	workingPrice := float32(1.0) // float32 | 
	workingQuantity := float32(1.0) // float32 | Sets the quantity for the working order. 
	pendingSide := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	pendingQuantity := float32(1.0) // float32 | Sets the quantity for the pending order.
	pendingAboveType := models.OrderListOcoAboveTypeParameterStopLossLimit // OrderListOcoAboveTypeParameter | 
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	workingClientOrderId := "workingClientOrderId_example" // string | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  (optional)
	workingIcebergQty := float32(1.0) // float32 | This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.  (optional)
	workingTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	workingStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the working order within an order strategy.  (optional)
	workingStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	workingPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	workingPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	workingPegOffsetValue := int32(1) // int32 |  (optional)
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  (optional)
	pendingAbovePrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.  (optional)
	pendingAboveStopPrice := float32(1.0) // float32 | Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`  (optional)
	pendingAboveTrailingDelta := float32(1.0) // float32 | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  (optional)
	pendingAboveIcebergQty := float32(1.0) // float32 | This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.  (optional)
	pendingAboveTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingAboveStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending above order within an order strategy.  (optional)
	pendingAboveStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingAbovePegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingAbovePegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingAbovePegOffsetValue := int32(1) // int32 |  (optional)
	pendingBelowType := models.OrderListOcoBelowTypeParameterStopLoss // OrderListOcoBelowTypeParameter |  (optional)
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  (optional)
	pendingBelowPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price  (optional)
	pendingBelowStopPrice := float32(1.0) // float32 | Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.  (optional)
	pendingBelowTrailingDelta := float32(1.0) // float32 |  (optional)
	pendingBelowIcebergQty := float32(1.0) // float32 | This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.  (optional)
	pendingBelowTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	pendingBelowStrategyId := int64(1) // int64 | Arbitrary numeric value identifying the pending below order within an order strategy.  (optional)
	pendingBelowStrategyType := int32(1) // int32 | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  (optional)
	pendingBelowPegPriceType := models.OrderListOcoAbovePegPriceTypeParameterPrimaryPeg // OrderListOcoAbovePegPriceTypeParameter |  (optional)
	pendingBelowPegOffsetType := models.OrderListOcoAbovePegOffsetTypeParameterPriceLevel // OrderListOcoAbovePegOffsetTypeParameter |  (optional)
	pendingBelowPegOffsetValue := int32(1) // int32 |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderListOtoco(context.Background()).Symbol(symbol).WorkingType(workingType).WorkingSide(workingSide).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).PendingSide(pendingSide).PendingQuantity(pendingQuantity).PendingAboveType(pendingAboveType).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).WorkingStrategyId(workingStrategyId).WorkingStrategyType(workingStrategyType).WorkingPegPriceType(workingPegPriceType).WorkingPegOffsetType(workingPegOffsetType).WorkingPegOffsetValue(workingPegOffsetValue).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveStrategyId(pendingAboveStrategyId).PendingAboveStrategyType(pendingAboveStrategyType).PendingAbovePegPriceType(pendingAbovePegPriceType).PendingAbovePegOffsetType(pendingAbovePegOffsetType).PendingAbovePegOffsetValue(pendingAbovePegOffsetValue).PendingBelowType(pendingBelowType).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowStrategyId(pendingBelowStrategyId).PendingBelowStrategyType(pendingBelowStrategyType).PendingBelowPegPriceType(pendingBelowPegPriceType).PendingBelowPegOffsetType(pendingBelowPegOffsetType).PendingBelowPegOffsetValue(pendingBelowPegOffsetValue).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderListOtoco``: %v\n", err)
		return
	}

	// response from `OrderListOtoco`: OrderListOtocoResponse
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
 **workingType** | [**OrderListOpoWorkingTypeParameter**](OrderListOpoWorkingTypeParameter.md) |  | 
 **workingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **workingPrice** | **float32** |  | 
 **workingQuantity** | **float32** | Sets the quantity for the working order.  | 
 **pendingSide** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **pendingQuantity** | **float32** | Sets the quantity for the pending order. | 
 **pendingAboveType** | [**OrderListOcoAboveTypeParameter**](OrderListOcoAboveTypeParameter.md) |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **workingClientOrderId** | **string** | Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.  | 
 **workingIcebergQty** | **float32** | This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **workingTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **workingStrategyId** | **int64** | Arbitrary numeric value identifying the working order within an order strategy.  | 
 **workingStrategyType** | **int32** | Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **workingPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **workingPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **workingPegOffsetValue** | **int32** |  | 
 **pendingAboveClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.  | 
 **pendingAbovePrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.  | 
 **pendingAboveStopPrice** | **float32** | Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;  | 
 **pendingAboveTrailingDelta** | **float32** | See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)  | 
 **pendingAboveIcebergQty** | **float32** | This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingAboveTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingAboveStrategyId** | **int64** | Arbitrary numeric value identifying the pending above order within an order strategy.  | 
 **pendingAboveStrategyType** | **int32** | Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingAbovePegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingAbovePegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingAbovePegOffsetValue** | **int32** |  | 
 **pendingBelowType** | [**OrderListOcoBelowTypeParameter**](OrderListOcoBelowTypeParameter.md) |  | 
 **pendingBelowClientOrderId** | **string** | Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.  | 
 **pendingBelowPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price  | 
 **pendingBelowStopPrice** | **float32** | Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.  | 
 **pendingBelowTrailingDelta** | **float32** |  | 
 **pendingBelowIcebergQty** | **float32** | This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.  | 
 **pendingBelowTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **pendingBelowStrategyId** | **int64** | Arbitrary numeric value identifying the pending below order within an order strategy.  | 
 **pendingBelowStrategyType** | **int32** | Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.  | 
 **pendingBelowPegPriceType** | [**OrderListOcoAbovePegPriceTypeParameter**](OrderListOcoAbovePegPriceTypeParameter.md) |  | 
 **pendingBelowPegOffsetType** | [**OrderListOcoAbovePegOffsetTypeParameter**](OrderListOcoAbovePegOffsetTypeParameter.md) |  | 
 **pendingBelowPegOffsetValue** | **int32** |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderListOtocoResponse**](OrderListOtocoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderOco

> OrderOcoResponse OrderOco(ctx).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).LimitIcebergQty(limitIcebergQty).TrailingDelta(trailingDelta).StopClientOrderId(stopClientOrderId).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New OCO - Deprecated


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	quantity := float32(1.0) // float32 | 
	price := float32(1.0) // float32 | 
	stopPrice := float32(1.0) // float32 | 
	listClientOrderId := "listClientOrderId_example" // string | A unique Id for the entire orderList (optional)
	limitClientOrderId := "limitClientOrderId_example" // string | A unique Id for the limit order (optional)
	limitStrategyId := int64(1) // int64 |  (optional)
	limitStrategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	limitIcebergQty := float32(1.0) // float32 | Used to make the `LIMIT_MAKER` leg an iceberg order. (optional)
	trailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	stopClientOrderId := "stopClientOrderId_example" // string | A unique Id for the stop loss/stop loss limit leg (optional)
	stopStrategyId := int64(1) // int64 |  (optional)
	stopStrategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	stopLimitPrice := float32(1.0) // float32 | If provided, `stopLimitTimeInForce` is required. (optional)
	stopIcebergQty := float32(1.0) // float32 | Used with `STOP_LOSS_LIMIT` leg to make an iceberg order. (optional)
	stopLimitTimeInForce := models.OrderOcoStopLimitTimeInForceParameterGtc // OrderOcoStopLimitTimeInForceParameter |  (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderOco(context.Background()).Symbol(symbol).Side(side).Quantity(quantity).Price(price).StopPrice(stopPrice).ListClientOrderId(listClientOrderId).LimitClientOrderId(limitClientOrderId).LimitStrategyId(limitStrategyId).LimitStrategyType(limitStrategyType).LimitIcebergQty(limitIcebergQty).TrailingDelta(trailingDelta).StopClientOrderId(stopClientOrderId).StopStrategyId(stopStrategyId).StopStrategyType(stopStrategyType).StopLimitPrice(stopLimitPrice).StopIcebergQty(stopIcebergQty).StopLimitTimeInForce(stopLimitTimeInForce).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderOco``: %v\n", err)
		return
	}

	// response from `OrderOco`: OrderOcoResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **quantity** | **float32** |  | 
 **price** | **float32** |  | 
 **stopPrice** | **float32** |  | 
 **listClientOrderId** | **string** | A unique Id for the entire orderList | 
 **limitClientOrderId** | **string** | A unique Id for the limit order | 
 **limitStrategyId** | **int64** |  | 
 **limitStrategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **limitIcebergQty** | **float32** | Used to make the &#x60;LIMIT_MAKER&#x60; leg an iceberg order. | 
 **trailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **stopClientOrderId** | **string** | A unique Id for the stop loss/stop loss limit leg | 
 **stopStrategyId** | **int64** |  | 
 **stopStrategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **stopLimitPrice** | **float32** | If provided, &#x60;stopLimitTimeInForce&#x60; is required. | 
 **stopIcebergQty** | **float32** | Used with &#x60;STOP_LOSS_LIMIT&#x60; leg to make an iceberg order. | 
 **stopLimitTimeInForce** | [**OrderOcoStopLimitTimeInForceParameter**](OrderOcoStopLimitTimeInForceParameter.md) |  | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderOcoResponse**](OrderOcoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## OrderTest

> OrderTestResponse OrderTest(ctx).Symbol(symbol).Side(side).Type(type_).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()

Test new order


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	type_ := models.NewOrderTypeParameterMarket // NewOrderTypeParameter | 
	computeCommissionRates := false // bool | Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 |  (optional)
	quoteOrderQty := float32(1.0) // float32 |  (optional)
	price := float32(400.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders. (optional)
	trailingDelta := int64(1) // int64 | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	pegPriceType := models.NewOrderPegPriceTypeParameterPrimaryPeg // NewOrderPegPriceTypeParameter |  (optional)
	pegOffsetValue := int32(1) // int32 | Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info) (optional)
	pegOffsetType := models.NewOrderPegOffsetTypeParameterPriceLevel // NewOrderPegOffsetTypeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.OrderTest(context.Background()).Symbol(symbol).Side(side).Type(type_).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Quantity(quantity).QuoteOrderQty(quoteOrderQty).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).StopPrice(stopPrice).TrailingDelta(trailingDelta).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).PegPriceType(pegPriceType).PegOffsetValue(pegOffsetValue).PegOffsetType(pegOffsetType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.OrderTest``: %v\n", err)
		return
	}

	// response from `OrderTest`: OrderTestResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** |  | 
 **quoteOrderQty** | **float32** |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **stopPrice** | **float32** | Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders. | 
 **trailingDelta** | **int64** | See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md). | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **pegPriceType** | [**NewOrderPegPriceTypeParameter**](NewOrderPegPriceTypeParameter.md) |  | 
 **pegOffsetValue** | **int32** | Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info) | 
 **pegOffsetType** | [**NewOrderPegOffsetTypeParameter**](NewOrderPegOffsetTypeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**OrderTestResponse**](OrderTestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SorOrder

> SorOrderResponse SorOrder(ctx).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New order using SOR


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	type_ := models.NewOrderTypeParameterMarket // NewOrderTypeParameter | 
	quantity := float32(1.0) // float32 | 
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	price := float32(400.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.SorOrder(context.Background()).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SorOrder``: %v\n", err)
		return
	}

	// response from `SorOrder`: SorOrderResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **quantity** | **float32** |  | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderResponse**](SorOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SorOrderTest

> SorOrderTestResponse SorOrderTest(ctx).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

Test new order using SOR


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
	side := models.NewOrderSideParameterBuy // NewOrderSideParameter | 
	type_ := models.NewOrderTypeParameterMarket // NewOrderTypeParameter | 
	quantity := float32(1.0) // float32 | 
	computeCommissionRates := false // bool | Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	price := float32(400.0) // float32 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string | A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected. (optional)
	strategyId := int64(1) // int64 |  (optional)
	strategyType := int32(1) // int32 | The value cannot be less than `1000000`. (optional)
	icebergQty := float32(1.0) // float32 | Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter |  (optional)
	recvWindow := float32(5000.0) // float32 | The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceSpotClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TradeAPI.SorOrderTest(context.Background()).Symbol(symbol).Side(side).Type(type_).Quantity(quantity).ComputeCommissionRates(computeCommissionRates).TimeInForce(timeInForce).Price(price).NewClientOrderId(newClientOrderId).StrategyId(strategyId).StrategyType(strategyType).IcebergQty(icebergQty).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.SorOrderTest``: %v\n", err)
		return
	}

	// response from `SorOrderTest`: SorOrderTestResponse
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
 **side** | [**NewOrderSideParameter**](NewOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **quantity** | **float32** |  | 
 **computeCommissionRates** | **bool** | Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more. | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected. | 
 **strategyId** | **int64** |  | 
 **strategyType** | **int32** | The value cannot be less than &#x60;1000000&#x60;. | 
 **icebergQty** | **float32** | Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) |  | 
 **recvWindow** | **float32** | The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified. | 

### Return type

[**SorOrderTestResponse**](SorOrderTestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

