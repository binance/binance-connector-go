# \TradeAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CancelOrder**](TradeAPI.md#CancelOrder) | /order.cancel | Cancel Order (TRADE)
[**ModifyOrder**](TradeAPI.md#ModifyOrder) | /order.modify | Modify Order (TRADE)
[**NewOrder**](TradeAPI.md#NewOrder) | /order.place | New Order (TRADE)
[**PositionInformation**](TradeAPI.md#PositionInformation) | /account.position | Position Information (USER_DATA)
[**QueryOrder**](TradeAPI.md#QueryOrder) | /order.status | Query Order (USER_DATA)


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

> ModifyOrderResponse ModifyOrder().Symbol(symbol).Side(side).Quantity(quantity).Price(price).Id(id).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

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
	quantity := float32(1.0) // float32 | Order quantity, cannot be sent with `closePosition=true`
	price := float32(1.0) // float32 | 
	id := "e9d6b4349871b40611412680b3445fac" // string |  (optional)
	orderId := int64(1) // int64 |  (optional)
	origClientOrderId := "1" // string |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
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
 **side** | [**ModifyOrderSideParameter**](ModifyOrderSideParameter.md) |  | 
 **quantity** | **float32** | Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60; | 
 **price** | **float32** |  | 
 **id** | **string** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **recvWindow** | **int64** |  | 

### Return type

[**ModifyOrderResponse**](ModifyOrderResponse.md)

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
	positionSide := models.NewOrderPositionSideParameterBoth // NewOrderPositionSideParameter | Default `BOTH` for One-way Mode; `LONG` or `SHORT` for Hedge Mode.  It must be sent in Hedge Mode. (optional)
	timeInForce := models.NewOrderTimeInForceParameterGtc // NewOrderTimeInForceParameter |  (optional)
	quantity := float32(1.0) // float32 | Quantity measured by contract number, Cannot be sent with `closePosition`=`true` (optional)
	reduceOnly := models.NewOrderReduceOnlyParameterTrue // NewOrderReduceOnlyParameter | Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (Close-All)\" (optional)
	price := float32(1.0) // float32 |  (optional)
	newClientOrderId := "1" // string | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$` (optional)
	stopPrice := float32(1.0) // float32 | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders. (optional)
	closePosition := models.NewOrderReduceOnlyParameterTrue // NewOrderReduceOnlyParameter | `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`. (optional)
	activationPrice := float32(1.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different workingType) (optional)
	callbackRate := float32(5000.0) // float32 | Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1% (optional)
	workingType := models.NewOrderWorkingTypeParameterMarkPrice // NewOrderWorkingTypeParameter | stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\" (optional)
	priceProtect := models.NewOrderReduceOnlyParameterTrue // NewOrderReduceOnlyParameter | Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.' (optional)
	newOrderRespType := models.NewOrderNewOrderRespTypeParameterAck // NewOrderNewOrderRespTypeParameter |  (optional)
	priceMatch := models.ModifyOrderPriceMatchParameterOpponent // ModifyOrderPriceMatchParameter | only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price` (optional)
	selfTradePreventionMode := models.NewOrderSelfTradePreventionModeParameterNone // NewOrderSelfTradePreventionModeParameter | `NONE`: No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers (optional)
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
 **positionSide** | [**NewOrderPositionSideParameter**](NewOrderPositionSideParameter.md) | Default &#x60;BOTH&#x60; for One-way Mode; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode.  It must be sent in Hedge Mode. | 
 **timeInForce** | [**NewOrderTimeInForceParameter**](NewOrderTimeInForceParameter.md) |  | 
 **quantity** | **float32** | Quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; | 
 **reduceOnly** | [**NewOrderReduceOnlyParameter**](NewOrderReduceOnlyParameter.md) | Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; (Close-All)\&quot; | 
 **price** | **float32** |  | 
 **newClientOrderId** | **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | 
 **stopPrice** | **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | 
 **closePosition** | [**NewOrderReduceOnlyParameter**](NewOrderReduceOnlyParameter.md) | &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;. | 
 **activationPrice** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different workingType) | 
 **callbackRate** | **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1% | 
 **workingType** | [**NewOrderWorkingTypeParameter**](NewOrderWorkingTypeParameter.md) | stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot; | 
 **priceProtect** | [**NewOrderReduceOnlyParameter**](NewOrderReduceOnlyParameter.md) | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.&#39; | 
 **newOrderRespType** | [**NewOrderNewOrderRespTypeParameter**](NewOrderNewOrderRespTypeParameter.md) |  | 
 **priceMatch** | [**ModifyOrderPriceMatchParameter**](ModifyOrderPriceMatchParameter.md) | only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60; | 
 **selfTradePreventionMode** | [**NewOrderSelfTradePreventionModeParameter**](NewOrderSelfTradePreventionModeParameter.md) | &#x60;NONE&#x60;: No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers | 
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

