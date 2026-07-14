/*
Futures (USDⓈ-M) WebSocket API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package binancederivativestradingusdsfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiCancelAlgoOrderRequest struct {
	ApiService   *TradeAPIService
	id           *string
	algoId       *int64
	clientAlgoId *string
	recvWindow   *int64
}

// Id.
func (r ApiCancelAlgoOrderRequest) Id(id string) ApiCancelAlgoOrderRequest {
	r.id = &id
	return r
}

// Algo Id.
func (r ApiCancelAlgoOrderRequest) AlgoId(algoId int64) ApiCancelAlgoOrderRequest {
	r.algoId = &algoId
	return r
}

// Client Algo Id.
func (r ApiCancelAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiCancelAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// Recv Window.
func (r ApiCancelAlgoOrderRequest) RecvWindow(recvWindow int64) ApiCancelAlgoOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAlgoOrderRequest) Execute() (*common.ResponseOrRaw[models.CancelAlgoOrderResponse], error) {
	respChan, errChan, err := r.ApiService.CancelAlgoOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiCancelAlgoOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.CancelAlgoOrderResponse], chan error, error) {
	return r.ApiService.CancelAlgoOrderExecute(r)
}

/*
CancelAlgoOrder Cancel Algo Order (TRADE)
/algoOrder.cancel

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#cancel-algo-order

@param id Id.	@param algoId Algo Id.	@param clientAlgoId Client Algo Id.	@param recvWindow Recv Window.
@return ApiCancelAlgoOrderRequest
*/
func (a *TradeAPIService) CancelAlgoOrder() ApiCancelAlgoOrderRequest {
	return ApiCancelAlgoOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CancelAlgoOrderResponse
func (a *TradeAPIService) CancelAlgoOrderExecute(r ApiCancelAlgoOrderRequest) (chan *common.ResponseOrRaw[models.CancelAlgoOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.algoId != nil {
		localVarQueryParams["algoId"] = *r.algoId
	}
	if r.clientAlgoId != nil {
		localVarQueryParams["clientAlgoId"] = *r.clientAlgoId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/algoOrder.cancel"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.CancelAlgoOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiCancelOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	id                *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Symbol.
func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
	return r
}

// Id.
func (r ApiCancelOrderRequest) Id(id string) ApiCancelOrderRequest {
	r.id = &id
	return r
}

// Order Id.
func (r ApiCancelOrderRequest) OrderId(orderId int64) ApiCancelOrderRequest {
	r.orderId = &orderId
	return r
}

// Orig Client Order Id.
func (r ApiCancelOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Recv Window.
func (r ApiCancelOrderRequest) RecvWindow(recvWindow int64) ApiCancelOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelOrderRequest) Execute() (*common.ResponseOrRaw[models.CancelOrderResponse], error) {
	respChan, errChan, err := r.ApiService.CancelOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiCancelOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.CancelOrderResponse], chan error, error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel Order (TRADE)
/order.cancel

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#cancel-order

@param symbol Symbol.	@param id Id.	@param orderId Order Id.	@param origClientOrderId Orig Client Order Id.	@param recvWindow Recv Window.
@return ApiCancelOrderRequest
*/
func (a *TradeAPIService) CancelOrder() ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CancelOrderResponse
func (a *TradeAPIService) CancelOrderExecute(r ApiCancelOrderRequest) (chan *common.ResponseOrRaw[models.CancelOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.cancel"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.CancelOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiModifyOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	side              *models.ModifyOrderSideParameter
	quantity          *float32
	price             *float32
	id                *string
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.ModifyOrderPriceMatchParameter
	recvWindow        *int64
}

// Symbol.
func (r ApiModifyOrderRequest) Symbol(symbol string) ApiModifyOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiModifyOrderRequest) Side(side models.ModifyOrderSideParameter) ApiModifyOrderRequest {
	r.side = &side
	return r
}

// Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60;
func (r ApiModifyOrderRequest) Quantity(quantity float32) ApiModifyOrderRequest {
	r.quantity = &quantity
	return r
}

// Price.
func (r ApiModifyOrderRequest) Price(price float32) ApiModifyOrderRequest {
	r.price = &price
	return r
}

// Id.
func (r ApiModifyOrderRequest) Id(id string) ApiModifyOrderRequest {
	r.id = &id
	return r
}

// Order Id.
func (r ApiModifyOrderRequest) OrderId(orderId int64) ApiModifyOrderRequest {
	r.orderId = &orderId
	return r
}

// Orig Client Order Id.
func (r ApiModifyOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiModifyOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// Recv Window.
func (r ApiModifyOrderRequest) RecvWindow(recvWindow int64) ApiModifyOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyOrderRequest) Execute() (*common.ResponseOrRaw[models.ModifyOrderResponse], error) {
	respChan, errChan, err := r.ApiService.ModifyOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiModifyOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.ModifyOrderResponse], chan error, error) {
	return r.ApiService.ModifyOrderExecute(r)
}

/*
ModifyOrder Modify Order (TRADE)
/order.modify

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#modify-order

@param symbol Symbol.	@param side `SELL`, `BUY`	@param quantity Order quantity, cannot be sent with `closePosition=true`	@param price Price.	@param id Id.	@param orderId Order Id.	@param origClientOrderId Orig Client Order Id.	@param priceMatch only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price	@param recvWindow Recv Window.
@return ApiModifyOrderRequest
*/
func (a *TradeAPIService) ModifyOrder() ApiModifyOrderRequest {
	return ApiModifyOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ModifyOrderResponse
func (a *TradeAPIService) ModifyOrderExecute(r ApiModifyOrderRequest) (chan *common.ResponseOrRaw[models.ModifyOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.quantity == nil {
		return nil, nil, common.ReportError("quantity is required and must be specified")
	}
	localVarQueryParams["quantity"] = *r.quantity

	if r.price == nil {
		return nil, nil, common.ReportError("price is required and must be specified")
	}
	localVarQueryParams["price"] = *r.price

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.priceMatch != nil {
		localVarQueryParams["priceMatch"] = *r.priceMatch
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.modify"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.ModifyOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiNewAlgoOrderRequest struct {
	ApiService              *TradeAPIService
	algoType                *models.NewAlgoOrderAlgoTypeParameter
	symbol                  *string
	side                    *models.ModifyOrderSideParameter
	type_                   *models.NewAlgoOrderTypeParameter
	id                      *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float32
	price                   *float32
	triggerPrice            *float32
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	closePosition           *models.NewAlgoOrderClosePositionParameter
	priceProtect            *models.NewAlgoOrderClosePositionParameter
	reduceOnly              *models.NewAlgoOrderClosePositionParameter
	activatePrice           *float32
	callbackRate            *float32
	clientAlgoId            *string
	newOrderRespType        *models.NewAlgoOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

// Only support &#x60;CONDITIONAL&#x60;
func (r ApiNewAlgoOrderRequest) AlgoType(algoType models.NewAlgoOrderAlgoTypeParameter) ApiNewAlgoOrderRequest {
	r.algoType = &algoType
	return r
}

// Symbol.
func (r ApiNewAlgoOrderRequest) Symbol(symbol string) ApiNewAlgoOrderRequest {
	r.symbol = &symbol
	return r
}

// Side.
func (r ApiNewAlgoOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewAlgoOrderRequest {
	r.side = &side
	return r
}

// For &#x60;CONDITIONAL&#x60; algoType, &#x60;STOP_MARKET&#x60;/&#x60;TAKE_PROFIT_MARKET&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60;/&#x60;TRAILING_STOP_MARKET&#x60; as order type
func (r ApiNewAlgoOrderRequest) Type(type_ models.NewAlgoOrderTypeParameter) ApiNewAlgoOrderRequest {
	r.type_ = &type_
	return r
}

// Id.
func (r ApiNewAlgoOrderRequest) Id(id string) ApiNewAlgoOrderRequest {
	r.id = &id
	return r
}

// Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewAlgoOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewAlgoOrderRequest {
	r.positionSide = &positionSide
	return r
}

// &#x60;IOC&#x60; or &#x60;GTC&#x60; or &#x60;FOK&#x60;, default &#x60;GTC&#x60;
func (r ApiNewAlgoOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewAlgoOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewAlgoOrderRequest) Quantity(quantity float32) ApiNewAlgoOrderRequest {
	r.quantity = &quantity
	return r
}

// Price.
func (r ApiNewAlgoOrderRequest) Price(price float32) ApiNewAlgoOrderRequest {
	r.price = &price
	return r
}

// Trigger Price.
func (r ApiNewAlgoOrderRequest) TriggerPrice(triggerPrice float32) ApiNewAlgoOrderRequest {
	r.triggerPrice = &triggerPrice
	return r
}

// triggerPrice triggered by: &#x60;MARK_PRICE&#x60;, &#x60;CONTRACT_PRICE&#x60;. Default &#x60;CONTRACT_PRICE&#x60;
func (r ApiNewAlgoOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewAlgoOrderRequest {
	r.workingType = &workingType
	return r
}

// only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price
func (r ApiNewAlgoOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewAlgoOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
func (r ApiNewAlgoOrderRequest) ClosePosition(closePosition models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \&quot;MARK_PRICE\&quot; and \&quot;CONTRACT_PRICE\&quot; cannot be larger than the Price Protection Threshold of the symbol.
func (r ApiNewAlgoOrderRequest) PriceProtect(priceProtect models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// Cannot be sent in Hedge Mode; cannot be sent with closePosition&#x3D;true
func (r ApiNewAlgoOrderRequest) ReduceOnly(reduceOnly models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
func (r ApiNewAlgoOrderRequest) ActivatePrice(activatePrice float32) ApiNewAlgoOrderRequest {
	r.activatePrice = &activatePrice
	return r
}

// Used with TRAILING_STOP_MARKET orders
func (r ApiNewAlgoOrderRequest) CallbackRate(callbackRate float32) ApiNewAlgoOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiNewAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewAlgoOrderRequest) NewOrderRespType(newOrderRespType models.NewAlgoOrderNewOrderRespTypeParameter) ApiNewAlgoOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewAlgoOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiNewAlgoOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
func (r ApiNewAlgoOrderRequest) GoodTillDate(goodTillDate int64) ApiNewAlgoOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

// Recv Window.
func (r ApiNewAlgoOrderRequest) RecvWindow(recvWindow int64) ApiNewAlgoOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewAlgoOrderRequest) Execute() (*common.ResponseOrRaw[models.NewAlgoOrderResponse], error) {
	respChan, errChan, err := r.ApiService.NewAlgoOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiNewAlgoOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.NewAlgoOrderResponse], chan error, error) {
	return r.ApiService.NewAlgoOrderExecute(r)
}

/*
NewAlgoOrder New Algo Order (TRADE)
/algoOrder.place

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#new-algo-order

@param algoType Only support `CONDITIONAL`	@param symbol Symbol.	@param side Side.	@param type_ For `CONDITIONAL` algoType, `STOP_MARKET`/`TAKE_PROFIT_MARKET`/`STOP`/`TAKE_PROFIT`/`TRAILING_STOP_MARKET` as order type	@param id Id.	@param positionSide Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode.	@param timeInForce `IOC` or `GTC` or `FOK`, default `GTC`	@param quantity Cannot be sent with `closePosition`=`true`(Close-All)	@param price Price.	@param triggerPrice Trigger Price.	@param workingType triggerPrice triggered by: `MARK_PRICE`, `CONTRACT_PRICE`. Default `CONTRACT_PRICE`	@param priceMatch only avaliable for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price	@param closePosition Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.	@param priceProtect Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol.	@param reduceOnly Cannot be sent in Hedge Mode; cannot be sent with closePosition=true	@param activatePrice Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)	@param callbackRate Used with TRAILING_STOP_MARKET orders	@param clientAlgoId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param newOrderRespType \"ACK\", \"RESULT\", default \"ACK\"	@param selfTradePreventionMode `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`	@param goodTillDate order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000	@param recvWindow Recv Window.
@return ApiNewAlgoOrderRequest
*/
func (a *TradeAPIService) NewAlgoOrder() ApiNewAlgoOrderRequest {
	return ApiNewAlgoOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return NewAlgoOrderResponse
func (a *TradeAPIService) NewAlgoOrderExecute(r ApiNewAlgoOrderRequest) (chan *common.ResponseOrRaw[models.NewAlgoOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.algoType == nil {
		return nil, nil, common.ReportError("algoType is required and must be specified")
	}
	localVarQueryParams["algoType"] = *r.algoType

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.type_ == nil {
		return nil, nil, common.ReportError("type_ is required and must be specified")
	}
	localVarQueryParams["type_"] = *r.type_

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.positionSide != nil {
		localVarQueryParams["positionSide"] = *r.positionSide
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.triggerPrice != nil {
		localVarQueryParams["triggerPrice"] = *r.triggerPrice
	}
	if r.workingType != nil {
		localVarQueryParams["workingType"] = *r.workingType
	}
	if r.priceMatch != nil {
		localVarQueryParams["priceMatch"] = *r.priceMatch
	}
	if r.closePosition != nil {
		localVarQueryParams["closePosition"] = *r.closePosition
	}
	if r.priceProtect != nil {
		localVarQueryParams["priceProtect"] = *r.priceProtect
	}
	if r.reduceOnly != nil {
		localVarQueryParams["reduceOnly"] = *r.reduceOnly
	}
	if r.activatePrice != nil {
		localVarQueryParams["activatePrice"] = *r.activatePrice
	}
	if r.callbackRate != nil {
		localVarQueryParams["callbackRate"] = *r.callbackRate
	}
	if r.clientAlgoId != nil {
		localVarQueryParams["clientAlgoId"] = *r.clientAlgoId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.goodTillDate != nil {
		localVarQueryParams["goodTillDate"] = *r.goodTillDate
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/algoOrder.place"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.NewAlgoOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiNewOrderRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.ModifyOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	id                      *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewOrderTimeInForceParameter
	reduceOnly              *models.NewAlgoOrderClosePositionParameter
	quantity                *float32
	price                   *float32
	newClientOrderId        *string
	newOrderRespType        *models.NewAlgoOrderNewOrderRespTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

// Symbol.
func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

// Side.
func (r ApiNewOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Id.
func (r ApiNewOrderRequest) Id(id string) ApiNewOrderRequest {
	r.id = &id
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

// Time In Force.
func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Cannot be sent in Hedge Mode
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// Price.
func (r ApiNewOrderRequest) Price(price float32) ApiNewOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewAlgoOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only available for &#x60;LIMIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;NONE&#x60;:No STP / &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
func (r ApiNewOrderRequest) GoodTillDate(goodTillDate int64) ApiNewOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

// Recv Window.
func (r ApiNewOrderRequest) RecvWindow(recvWindow int64) ApiNewOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewOrderRequest) Execute() (*common.ResponseOrRaw[models.NewOrderResponse], error) {
	respChan, errChan, err := r.ApiService.NewOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiNewOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.NewOrderResponse], chan error, error) {
	return r.ApiService.NewOrderExecute(r)
}

/*
NewOrder New Order (TRADE)
/order.place

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#new-order

@param symbol Symbol.	@param side Side.	@param type_	@param id Id.	@param positionSide Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.	@param timeInForce Time In Force.	@param reduceOnly Cannot be sent in Hedge Mode	@param quantity	@param price Price.	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param newOrderRespType	@param priceMatch only available for `LIMIT` order; Can't be passed together with `price`	@param selfTradePreventionMode `NONE`:No STP / `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`	@param goodTillDate order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000	@param recvWindow Recv Window.
@return ApiNewOrderRequest
*/
func (a *TradeAPIService) NewOrder() ApiNewOrderRequest {
	return ApiNewOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return NewOrderResponse
func (a *TradeAPIService) NewOrderExecute(r ApiNewOrderRequest) (chan *common.ResponseOrRaw[models.NewOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.type_ == nil {
		return nil, nil, common.ReportError("type_ is required and must be specified")
	}
	localVarQueryParams["type_"] = *r.type_

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.positionSide != nil {
		localVarQueryParams["positionSide"] = *r.positionSide
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.reduceOnly != nil {
		localVarQueryParams["reduceOnly"] = *r.reduceOnly
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.priceMatch != nil {
		localVarQueryParams["priceMatch"] = *r.priceMatch
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.goodTillDate != nil {
		localVarQueryParams["goodTillDate"] = *r.goodTillDate
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.place"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.NewOrderResponse](a.Ws, localPayload, sendParams)
}

type ApiPositionInformationRequest struct {
	ApiService *TradeAPIService
	id         *string
	symbol     *string
	recvWindow *int64
}

// Id.
func (r ApiPositionInformationRequest) Id(id string) ApiPositionInformationRequest {
	r.id = &id
	return r
}

// Symbol.
func (r ApiPositionInformationRequest) Symbol(symbol string) ApiPositionInformationRequest {
	r.symbol = &symbol
	return r
}

// Recv Window.
func (r ApiPositionInformationRequest) RecvWindow(recvWindow int64) ApiPositionInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionInformationRequest) Execute() (*common.ResponseOrRaw[models.PositionInformationResponse], error) {
	respChan, errChan, err := r.ApiService.PositionInformationExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiPositionInformationRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.PositionInformationResponse], chan error, error) {
	return r.ApiService.PositionInformationExecute(r)
}

/*
PositionInformation Position Information (USER_DATA)
/account.position

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#position-information

@param id Id.	@param symbol Symbol.	@param recvWindow Recv Window.
@return ApiPositionInformationRequest
*/
func (a *TradeAPIService) PositionInformation() ApiPositionInformationRequest {
	return ApiPositionInformationRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PositionInformationResponse
func (a *TradeAPIService) PositionInformationExecute(r ApiPositionInformationRequest) (chan *common.ResponseOrRaw[models.PositionInformationResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.position"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.PositionInformationResponse](a.Ws, localPayload, sendParams)
}

type ApiPositionInformationV2Request struct {
	ApiService *TradeAPIService
	id         *string
	symbol     *string
	recvWindow *int64
}

// Id.
func (r ApiPositionInformationV2Request) Id(id string) ApiPositionInformationV2Request {
	r.id = &id
	return r
}

// Symbol.
func (r ApiPositionInformationV2Request) Symbol(symbol string) ApiPositionInformationV2Request {
	r.symbol = &symbol
	return r
}

// Recv Window.
func (r ApiPositionInformationV2Request) RecvWindow(recvWindow int64) ApiPositionInformationV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionInformationV2Request) Execute() (*common.ResponseOrRaw[models.PositionInformationV2Response], error) {
	respChan, errChan, err := r.ApiService.PositionInformationV2Execute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiPositionInformationV2Request) ExecuteAsync() (chan *common.ResponseOrRaw[models.PositionInformationV2Response], chan error, error) {
	return r.ApiService.PositionInformationV2Execute(r)
}

/*
PositionInformationV2 Position Information V2 (USER_DATA)
/v2/account.position

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#position-information-v2

@param id Id.	@param symbol Symbol.	@param recvWindow Recv Window.
@return ApiPositionInformationV2Request
*/
func (a *TradeAPIService) PositionInformationV2() ApiPositionInformationV2Request {
	return ApiPositionInformationV2Request{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PositionInformationV2Response
func (a *TradeAPIService) PositionInformationV2Execute(r ApiPositionInformationV2Request) (chan *common.ResponseOrRaw[models.PositionInformationV2Response], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/v2/account.position"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.PositionInformationV2Response](a.Ws, localPayload, sendParams)
}

type ApiQueryOrderRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	id                *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Symbol.
func (r ApiQueryOrderRequest) Symbol(symbol string) ApiQueryOrderRequest {
	r.symbol = &symbol
	return r
}

// Id.
func (r ApiQueryOrderRequest) Id(id string) ApiQueryOrderRequest {
	r.id = &id
	return r
}

// Order Id.
func (r ApiQueryOrderRequest) OrderId(orderId int64) ApiQueryOrderRequest {
	r.orderId = &orderId
	return r
}

// Orig Client Order Id.
func (r ApiQueryOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Recv Window.
func (r ApiQueryOrderRequest) RecvWindow(recvWindow int64) ApiQueryOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOrderRequest) Execute() (*common.ResponseOrRaw[models.QueryOrderResponse], error) {
	respChan, errChan, err := r.ApiService.QueryOrderExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiQueryOrderRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.QueryOrderResponse], chan error, error) {
	return r.ApiService.QueryOrderExecute(r)
}

/*
QueryOrder Query Order (USER_DATA)
/order.status

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-api/trade#query-order

@param symbol Symbol.	@param id Id.	@param orderId Order Id.	@param origClientOrderId Orig Client Order Id.	@param recvWindow Recv Window.
@return ApiQueryOrderRequest
*/
func (a *TradeAPIService) QueryOrder() ApiQueryOrderRequest {
	return ApiQueryOrderRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return QueryOrderResponse
func (a *TradeAPIService) QueryOrderExecute(r ApiQueryOrderRequest) (chan *common.ResponseOrRaw[models.QueryOrderResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.QueryOrderResponse](a.Ws, localPayload, sendParams)
}
