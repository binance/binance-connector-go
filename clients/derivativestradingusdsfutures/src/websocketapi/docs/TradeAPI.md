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
[**PositionInformationV2**](TradeAPI.md#PositionInformationV2) | /v2/account.position | Position Information V2 (USER_DATA)
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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
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
	wsClient := models.NewBinanceDerivativesTradingUsdsFuturesClient(models.WithWebsocketAPI(configuration))

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Symbol.
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	orderId := int64(1) // int64 | Order Id. (optional)
	origClientOrderId := "1" // string | Orig Client Order Id. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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
 **symbol** | **string** | Symbol. | 
 **id** | **string** | Id. | 
 **orderId** | **int64** | Order Id. | 
 **origClientOrderId** | **string** | Orig Client Order Id. | 
 **recvWindow** | **int64** | Recv Window. | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Symbol.
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | `SELL`, `BUY`
	quantity := float64(1.0) // float64 | Order quantity, cannot be sent with `closePosition=true`
	price := float64(1.0) // float64 | Price.
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	orderId := int64(1) // int64 | Order Id. (optional)
	origClientOrderId := "1" // string | Orig Client Order Id. (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price (optional)
	modifyId := int64(1) // int64 | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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
 **symbol** | **string** | Symbol. | 
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | &#x60;SELL&#x60;, &#x60;BUY&#x60; | 
 **quantity** | **float64** | Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60; | 
 **price** | **float64** | Price. | 
 **id** | **string** | Id. | 
 **orderId** | **int64** | Order Id. | 
 **origClientOrderId** | **string** | Orig Client Order Id. | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price | 
 **modifyId** | **int64** | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewAlgoOrder

> NewAlgoOrderResponse NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Algo Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	algoType := models.NewAlgoOrderAlgoTypeParameterConditional // NewAlgoOrderAlgoTypeParameter | Only support `CONDITIONAL`
	symbol := "BTCUSDT" // string | Symbol.
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | Side.
	type_ := models.NewAlgoOrderTypeParameterStopMarket // NewAlgoOrderTypeParameter | For `CONDITIONAL` algoType, `STOP_MARKET`/`TAKE_PROFIT_MARKET`/`STOP`/`TAKE_PROFIT`/`TRAILING_STOP_MARKET` as order type
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewAlgoOrderTimeInForceParameterIoc // NewAlgoOrderTimeInForceParameter | `IOC` or `GTC` or `FOK`, default `GTC` (optional)
	quantity := float64(1.0) // float64 | Cannot be sent with `closePosition`=`true`(Close-All) (optional)
	price := float64(1.0) // float64 | Price. (optional)
	triggerPrice := float64(1.0) // float64 | Trigger Price. (optional)
	workingType := models.NewAlgoOrderWorkingTypeParameterMarkPrice // NewAlgoOrderWorkingTypeParameter | triggerPrice triggered by: `MARK_PRICE`, `CONTRACT_PRICE`. Default `CONTRACT_PRICE` (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price (optional)
	closePosition := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET. (optional)
	priceProtect := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol. (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode; cannot be sent with closePosition=true (optional)
	activatePrice := float64(1.0) // float64 | Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType) (optional)
	callbackRate := float64(1) // float64 | Used with TRAILING_STOP_MARKET orders (optional)
	clientAlgoId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter | \"ACK\", \"RESULT\", default \"ACK\" (optional)
	selfTradePreventionMode := models.NewAlgoOrderSelfTradePreventionModeParameterNone // NewAlgoOrderSelfTradePreventionModeParameter | `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(1770736694138) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType(algoType).Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).Quantity(quantity).Price(price).TriggerPrice(triggerPrice).WorkingType(workingType).PriceMatch(priceMatch).ClosePosition(closePosition).PriceProtect(priceProtect).ReduceOnly(reduceOnly).ActivatePrice(activatePrice).CallbackRate(callbackRate).ClientAlgoId(clientAlgoId).NewOrderRespType(newOrderRespType).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
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
 **timeInForce** | [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) | &#x60;IOC&#x60; or &#x60;GTC&#x60; or &#x60;FOK&#x60;, default &#x60;GTC&#x60; | 
 **quantity** | **float64** | Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All) | 
 **price** | **float64** | Price. | 
 **triggerPrice** | **float64** | Trigger Price. | 
 **workingType** | [**NewAlgoOrderWorkingTypeParameter**](NewAlgoOrderWorkingTypeParameter.md) | triggerPrice triggered by: &#x60;MARK_PRICE&#x60;, &#x60;CONTRACT_PRICE&#x60;. Default &#x60;CONTRACT_PRICE&#x60; | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price | 
 **closePosition** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET. | 
 **priceProtect** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \&quot;MARK_PRICE\&quot; and \&quot;CONTRACT_PRICE\&quot; cannot be larger than the Price Protection Threshold of the symbol. | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with closePosition&#x3D;true | 
 **activatePrice** | **float64** | Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType) | 
 **callbackRate** | **float64** | Used with TRAILING_STOP_MARKET orders | 
 **clientAlgoId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) | \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot; | 
 **selfTradePreventionMode** | [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) | &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**NewAlgoOrderResponse**](NewAlgoOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## NewOrder

> NewOrderResponse NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()

New Order (TRADE)


### Example

```go
package main

import (
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Symbol.
	side := models.ModifyOrderSideParameterBuy // ModifyOrderSideParameter | Side.
	type_ := models.NewOrderTypeParameterLimit // NewOrderTypeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	positionSide := models.NewAlgoOrderPositionSideParameterBoth // NewAlgoOrderPositionSideParameter | Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter | Time In Force. (optional)
	reduceOnly := models.NewAlgoOrderClosePositionParameterTrue // NewAlgoOrderClosePositionParameter | Cannot be sent in Hedge Mode (optional)
	quantity := float64(1.0) // float64 |  (optional)
	price := float64(1.0) // float64 | Price. (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	newOrderRespType := models.NewAlgoOrderNewOrderRespTypeParameterAck // NewAlgoOrderNewOrderRespTypeParameter |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only available for `LIMIT` order; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter | `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE` (optional)
	goodTillDate := int64(1770736694138) // int64 | order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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


	resp, err := wsClient.WebsocketAPI.TradeAPI.NewOrder().Symbol(symbol).Side(side).Type(type_).Id(id).PositionSide(positionSide).TimeInForce(timeInForce).ReduceOnly(reduceOnly).Quantity(quantity).Price(price).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PriceMatch(priceMatch).SelfTradePreventionMode(selfTradePreventionMode).GoodTillDate(goodTillDate).RecvWindow(recvWindow).Execute()
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
 **symbol** | **string** | Symbol. | 
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) | Side. | 
 **type_** | [**NewOrderTypeParameter**](NewOrderTypeParameter.md) |  | 
 **id** | **string** | Id. | 
 **positionSide** | [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) | Time In Force. | 
 **reduceOnly** | [**NewAlgoOrderClosePositionParameter**](NewAlgoOrderClosePositionParameter.md) | Cannot be sent in Hedge Mode | 
 **quantity** | **float64** |  | 
 **price** | **float64** | Price. | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **newOrderRespType** | [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only available for &#x60;LIMIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60; | 
 **goodTillDate** | **int64** | order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000 | 
 **recvWindow** | **int64** | Recv Window. | 

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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	symbol := "BTCUSDT" // string | Symbol. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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
 **id** | **string** | Id. | 
 **symbol** | **string** | Symbol. | 
 **recvWindow** | **int64** | Recv Window. | 

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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	symbol := "BTCUSDT" // string | Symbol. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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
 **id** | **string** | Id. | 
 **symbol** | **string** | Symbol. | 
 **recvWindow** | **int64** | Recv Window. | 

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
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "BTCUSDT" // string | Symbol.
	id := "e9d6b4349871b40611412680b3445fac" // string | Id. (optional)
	orderId := int64(1) // int64 | Order Id. (optional)
	origClientOrderId := "1" // string | Orig Client Order Id. (optional)
	recvWindow := int64(5000) // int64 | Recv Window. (optional)

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
 **symbol** | **string** | Symbol. | 
 **id** | **string** | Id. | 
 **orderId** | **int64** | Order Id. | 
 **origClientOrderId** | **string** | Orig Client Order Id. | 
 **recvWindow** | **int64** | Recv Window. | 

### Return type

[**QueryOrderResponse**](QueryOrderResponse.md)

### Authorization

No authorization required

[[Back to README]](../../../README.md)

