# \TradeAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelAlgoOrder**](TradeAPI.md#CancelAlgoOrder) | /algoOrder.cancel | Cancel Algo Order (TRADE)
[**CancelOrder**](TradeAPI.md#CancelOrder) | /order.cancel | Cancel Order (TRADE)
[**ModifyOrder**](TradeAPI.md#ModifyOrder) | /order.modify | Modify Order (TRADE)
[**NewAlgoOrder**](TradeAPI.md#NewAlgoOrder) | /algoOrder.place | New Algo Order(TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | /order.place | New Order(TRADE)
[**PositionInformation**](TradeAPI.md#PositionInformation) | /account.position | Position Information (USER_DATA)
[**PositionInformationV2**](TradeAPI.md#PositionInformationV2) | /v2/account.position | Position Information V2 (USER_DATA)
[**QueryOrder**](TradeAPI.md#QueryOrder) | /order.status | Query Order (USER_DATA)


## CancelAlgoOrder

> CancelAlgoOrderResponse CancelAlgoOrder().Id(id).Algoid(algoid).Clientalgoid(clientalgoid).RecvWindow(recvWindow).Execute()

Cancel Algo Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	algoid := int64(789) // int64 |  (optional)
	clientalgoid := "clientalgoid_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.CancelAlgoOrder().Id(id).Algoid(algoid).Clientalgoid(clientalgoid).RecvWindow(recvWindow).Execute()
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
 **id** | **string** | Unique WebSocket request ID. | 
 **algoid** | **int64** |  | 
 **clientalgoid** | **string** |  | 
 **recvWindow** | **int64** |  | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

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
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ModifyOrder

> ModifyOrderResponse ModifyOrder().Symbol(symbol).Side(side).Quantity(quantity).Price(price).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | `SELL`, `BUY`
	quantity := float32(1.0) // float32 | Order quantity, cannot be sent with `closePosition=true`
	price := float32(1.0) // float32 | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterNone // ModifyOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol(symbol).Side(side).Quantity(quantity).Price(price).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
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
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **quantity** | **float32** | Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60; | 
 **price** | **float32** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewAlgoOrder

> NewAlgoOrderResponse NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Algo Order(TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	algoType := "algoType_example" // string | Only support `CONDITIONAL`
	symbol := "symbol_example" // string | 
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | `SELL`, `BUY`
	type_ := "type__example" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterGtc // NewAlgoOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 | Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float32(1.0) // float32 |  (optional)
	triggerPrice := float32(1.0) // float32 |  (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterNone // ModifyOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	closePosition := "closePosition_example" // string | `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	priceProtect := "false" // string | \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (optional)
	activatePrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`) (optional)
	callbackRate := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1% (optional)
	clientAlgoId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterExpireTaker // NewAlgoOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(789) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
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
 **algoType** | **string** | Only support &#x60;CONDITIONAL&#x60; | 
 **symbol** | **string** |  | 
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **type_** | **string** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **positionSide** | [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** | Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float32** |  | 
 **triggerPrice** | **float32** |  | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **closePosition** | **string** | &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **priceProtect** | **string** | \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **activatePrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1% | 
 **clientAlgoId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewAlgoOrderResponse**](NewAlgoOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Order(TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | `SELL`, `BUY`
	type_ := "type__example" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterGtc // NewAlgoOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 | Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	reduceOnly := "false" // string | \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	closePosition := "closePosition_example" // string | `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`) (optional)
	callbackRate := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1% (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := "false" // string | \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterNone // ModifyOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterExpireTaker // NewAlgoOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(789) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).ReduceOnly(reduceOnly).Price(price).NewClientOrderId(newClientOrderId).StopPrice(stopPrice).ClosePosition(closePosition).ActivationPrice(activationPrice).CallbackRate(callbackRate).WorkingType(workingType).PriceProtect(priceProtect).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
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
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **type_** | **string** |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **positionSide** | [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** | Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **reduceOnly** | **string** | \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **closePosition** | **string** | &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1% | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | **string** | \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**NewOrderResponse**](NewOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PositionInformation

> PositionInformationResponse PositionInformation().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information (USER_DATA)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.PositionInformation().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()
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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationResponse**](PositionInformationResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PositionInformationV2

> PositionInformationV2Response PositionInformationV2().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V2 (USER_DATA)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	symbol := "symbol_example" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketAPI.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}


	resp, err := wsClient.WebsocketAPI.TradeAPI.PositionInformationV2().Id(id).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TradeAPI.PositionInformationV2``: %v\n", err)
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
 **id** | **string** | Unique WebSocket request ID. | 
 **symbol** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PositionInformationV2Response**](PositionInformationV2Response.md)

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "symbol_example" // string | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationWebsocketApi(
		common.WithWsApiBasePath(common.SpotWebsocketApiProdUrl),
		common.WithWsApiKey("Your API Key"),
		common.WithWsApiSecret("Your API Secret"),
	)
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

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
 **id** | **string** | Unique WebSocket request ID. | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**QueryOrderResponse**](QueryOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

