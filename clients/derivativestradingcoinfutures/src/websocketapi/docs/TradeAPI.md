# \TradeAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAlgoOrder**](TradeAPI.md#CancelAlgoOrder) | /algoOrder.cancel | Cancel Algo Order (TRADE)
[**CancelOrder**](TradeAPI.md#CancelOrder) | /order.cancel | Cancel Order (TRADE)
[**ModifyOrder**](TradeAPI.md#ModifyOrder) | /order.modify | Modify Order (TRADE)
[**NewAlgoOrder**](TradeAPI.md#NewAlgoOrder) | /algoOrder.place | New Algo Order (TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | /order.place | New Order (TRADE)
[**PositionInformation**](TradeAPI.md#PositionInformation) | /account.position | Position Information (USER_DATA)
[**QueryOrder**](TradeAPI.md#QueryOrder) | /order.status | Query Order (USER_DATA)


## CancelAlgoOrder

> CancelAlgoOrderResponse CancelAlgoOrder().Id(id).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()

Cancel Algo Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	algoId := int64(2000000002162519) // int64 | Algo Id. (optional)
	clientAlgoId := "rDMG8WSde6LkyMNtk6s825" // string | Client Algo Id. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.CancelAlgoOrder().Id(id).AlgoId(algoId).ClientAlgoId(clientAlgoId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelAlgoOrder``: %v\n", err)
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
 **id** | **string** | Id. | 
 **algoId** | **int64** | Algo Id. | 
 **clientAlgoId** | **string** | Client Algo Id. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**CancelAlgoOrderResponse**](CancelAlgoOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## CancelOrder

> CancelOrderResponse CancelOrder().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.CancelOrder().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.CancelOrder``: %v\n", err)
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
 **id** | **string** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ModifyOrder

> ModifyOrderResponse ModifyOrder().Symbol(symbol).Side(side).Quantity(quantity).Price(price).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).ModifyId(modifyId).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string | 
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | 
	quantity := float64(1.0) // float64 | Order quantity, cannot be sent with `closePosition=true`
	price := float64(1.0) // float64 | 
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	modifyId := int64(1) // int64 | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol(symbol).Side(side).Quantity(quantity).Price(price).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).ModifyId(modifyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.ModifyOrder``: %v\n", err)
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
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) |  | 
 **quantity** | **float64** | Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60; | 
 **price** | **float64** |  | 
 **id** | **string** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **modifyId** | **int64** | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewAlgoOrder

> NewAlgoOrderResponse NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New Algo Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algoType := models.NewAlgoOrderAlgoTypeParameterConditional // NewAlgoOrderAlgoTypeParameter | Only support `CONDITIONAL`
	symbol := "BTCUSD_PERP" // string | Symbol.
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | Side.
	type_ := models.NewAlgoOrderTypeParameterStopMarket // NewAlgoOrderTypeParameter | For `CONDITIONAL` algoType, `STOP_MARKET`/`TAKE_PROFIT_MARKET`/`STOP`/`TAKE_PROFIT`/`TRAILING_STOP_MARKET` as order type
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterIoc // NewAlgoOrderTimeInForceParameter | `IOC` or `GTC` or `FOK` or `GTX`, default `GTC` (optional)
	quantity := float64(1.0) // float64 | quantity measured by contract number, Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float64(1.0) // float64 | Price. (optional)
	triggerPrice := float64(1.0) // float64 | Trigger Price. (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter | triggerPrice triggered by: `MARK_PRICE`, `CONTRACT_PRICE`. Default `CONTRACT_PRICE` (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only available for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price (optional)
	closePosition := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET. (optional)
	priceProtect := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol. (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode; cannot be sent with closePosition=true (optional)
	activatePrice := float64(1.0) // float64 | Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType) (optional)
	callbackRate := float64(1) // float64 | Used with TRAILING_STOP_MARKET orders (optional)
	clientAlgoId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterNone // NewAlgoOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`: expire taker order when STP triggers/ `EXPIRE_MAKER`: expire maker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers; default `NONE` (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewAlgoOrder``: %v\n", err)
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
 **algoType** | [**NewAlgoOrderAlgoTypeParameter**](NewAlgoOrderAlgoTypeParameter.md) | Only support &#x60;CONDITIONAL&#x60; | 
 **symbol** | **string** | Symbol. | 
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | Side. | 
 **type_** | [**NewAlgoOrderTypeParameter**](NewAlgoOrderTypeParameter.md) | For &#x60;CONDITIONAL&#x60; algoType, &#x60;STOP_MARKET&#x60;/&#x60;TAKE_PROFIT_MARKET&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60;/&#x60;TRAILING_STOP_MARKET&#x60; as order type | 
 **id** | **string** | Id. | 
 **positionSide** | [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) | Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) | &#x60;IOC&#x60; or &#x60;GTC&#x60; or &#x60;FOK&#x60; or &#x60;GTX&#x60;, default &#x60;GTC&#x60; | 
 **quantity** | **float64** | quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float64** | Price. | 
 **triggerPrice** | **float64** | Trigger Price. | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) | triggerPrice triggered by: &#x60;MARK_PRICE&#x60;, &#x60;CONTRACT_PRICE&#x60;. Default &#x60;CONTRACT_PRICE&#x60; | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only available for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price | 
 **closePosition** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET. | 
 **priceProtect** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \&quot;MARK_PRICE\&quot; and \&quot;CONTRACT_PRICE\&quot; cannot be larger than the Price Protection Threshold of the symbol. | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with closePosition&#x3D;true | 
 **activatePrice** | **float64** | Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType) | 
 **callbackRate** | **float64** | Used with TRAILING_STOP_MARKET orders | 
 **clientAlgoId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;: expire maker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**NewAlgoOrderResponse**](NewAlgoOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()

New Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string | 
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | 
	type_ := models.NewOrderTypeParameterLimit // NewOrderTypeParameter | **After CM migration, stop-type values (`STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`) are no longer accepted and will return `-4120`. Use the REST `/dapi/v1/algoOrder` endpoint instead.**
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default `BOTH` for One-way Mode; `LONG` or `SHORT` for Hedge Mode.  It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	quantity := float64(1.0) // float64 | Quantity measured by contract number, Cannot be sent with `closePosition`=`true` (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (Close-All)\" (optional)
	price := float64(1.0) // float64 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	stopPrice := float64(1.0) // float64 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	closePosition := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	activationPrice := float64(1.0) // float64 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different workingType) (optional)
	callbackRate := float64(5000.0) // float64 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1% (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.' (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter | `NONE`: No STP / `EXPIRE_TAKER`: expire taker order when STP triggers/ `EXPIRE_MAKER`: expire taker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.NewOrder``: %v\n", err)
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
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) |  | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) | **After CM migration, stop-type values (&#x60;STOP&#x60;, &#x60;STOP_MARKET&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_MARKET&#x60;, &#x60;TRAILING_STOP_MARKET&#x60;) are no longer accepted and will return &#x60;-4120&#x60;. Use the REST &#x60;/dapi/v1/algoOrder&#x60; endpoint instead.** | 
 **id** | **string** |  | 
 **positionSide** | [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode.  It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **quantity** | **float64** | Quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; (Close-All)\&quot; | 
 **price** | **float64** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **stopPrice** | **float64** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **closePosition** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **activationPrice** | **float64** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different workingType) | 
 **callbackRate** | **float64** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1% | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.&#39; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;: No STP / &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewOrderResponse**](NewOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PositionInformation

> PositionInformationResponse PositionInformation().Id(id).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()

Position Information (USER_DATA)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	marginAsset := "USDT" // string |  (optional)
	pair := "BTCUSD" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.PositionInformation().Id(id).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionInformation``: %v\n", err)
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
 **id** | **string** |  | 
 **marginAsset** | **string** |  | 
 **pair** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationResponse**](PositionInformationResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## QueryOrder

> QueryOrderResponse QueryOrder().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Order (USER_DATA)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSD_PERP" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.QueryOrder().Symbol(symbol).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.QueryOrder``: %v\n", err)
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
 **id** | **string** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryOrderResponse**](QueryOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

