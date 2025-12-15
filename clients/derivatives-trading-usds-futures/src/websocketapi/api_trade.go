/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package binancederivativestradingusdsfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiCancelAlgoOrderRequest struct {
	ApiService   *TradeAPIService
	id           *string
	algoid       *int64
	clientalgoid *string
	recvWindow   *int64
}

// Unique WebSocket request ID.
func (r ApiCancelAlgoOrderRequest) Id(id string) ApiCancelAlgoOrderRequest {
	r.id = &id
	return r
}

func (r ApiCancelAlgoOrderRequest) Algoid(algoid int64) ApiCancelAlgoOrderRequest {
	r.algoid = &algoid
	return r
}

func (r ApiCancelAlgoOrderRequest) Clientalgoid(clientalgoid string) ApiCancelAlgoOrderRequest {
	r.clientalgoid = &clientalgoid
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Algo-Order

@param id Unique WebSocket request ID.	@param algoid	@param clientalgoid	@param recvWindow
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
	if r.algoid != nil {
		localVarQueryParams["algoid"] = *r.algoid
	}
	if r.clientalgoid != nil {
		localVarQueryParams["clientalgoid"] = *r.clientalgoid
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

func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiCancelOrderRequest) Id(id string) ApiCancelOrderRequest {
	r.id = &id
	return r
}

func (r ApiCancelOrderRequest) OrderId(orderId int64) ApiCancelOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCancelOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Cancel-Order

@param symbol	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param recvWindow
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

func (r ApiModifyOrderRequest) Price(price float32) ApiModifyOrderRequest {
	r.price = &price
	return r
}

// Unique WebSocket request ID.
func (r ApiModifyOrderRequest) Id(id string) ApiModifyOrderRequest {
	r.id = &id
	return r
}

func (r ApiModifyOrderRequest) OrderId(orderId int64) ApiModifyOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiModifyOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiModifyOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Modify-Order

@param symbol	@param side `SELL`, `BUY`	@param quantity Order quantity, cannot be sent with `closePosition=true`	@param price	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param priceMatch only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`	@param recvWindow
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
	algoType                *string
	symbol                  *string
	side                    *models.ModifyOrderSideParameter
	type_                   *string
	id                      *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float32
	price                   *float32
	triggerPrice            *float32
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	closePosition           *string
	priceProtect            *string
	reduceOnly              *string
	activationPrice         *float32
	callbackRate            *float32
	clientAlgoId            *string
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

// Only support &#x60;CONDITIONAL&#x60;
func (r ApiNewAlgoOrderRequest) AlgoType(algoType string) ApiNewAlgoOrderRequest {
	r.algoType = &algoType
	return r
}

func (r ApiNewAlgoOrderRequest) Symbol(symbol string) ApiNewAlgoOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiNewAlgoOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewAlgoOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewAlgoOrderRequest) Type(type_ string) ApiNewAlgoOrderRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiNewAlgoOrderRequest) Id(id string) ApiNewAlgoOrderRequest {
	r.id = &id
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewAlgoOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewAlgoOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewAlgoOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewAlgoOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewAlgoOrderRequest) Quantity(quantity float32) ApiNewAlgoOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiNewAlgoOrderRequest) Price(price float32) ApiNewAlgoOrderRequest {
	r.price = &price
	return r
}

func (r ApiNewAlgoOrderRequest) TriggerPrice(triggerPrice float32) ApiNewAlgoOrderRequest {
	r.triggerPrice = &triggerPrice
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewAlgoOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewAlgoOrderRequest {
	r.workingType = &workingType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewAlgoOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewAlgoOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewAlgoOrderRequest) ClosePosition(closePosition string) ApiNewAlgoOrderRequest {
	r.closePosition = &closePosition
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewAlgoOrderRequest) PriceProtect(priceProtect string) ApiNewAlgoOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewAlgoOrderRequest) ReduceOnly(reduceOnly string) ApiNewAlgoOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiNewAlgoOrderRequest) ActivationPrice(activationPrice float32) ApiNewAlgoOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1%
func (r ApiNewAlgoOrderRequest) CallbackRate(callbackRate float32) ApiNewAlgoOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiNewAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
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
NewAlgoOrder New Algo Order(TRADE)
/algoOrder.place

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/New-Algo-Order

@param algoType Only support `CONDITIONAL`	@param symbol	@param side `SELL`, `BUY`	@param type_	@param id Unique WebSocket request ID.	@param positionSide Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.	@param timeInForce	@param quantity Cannot be sent with `closePosition`=`true`(Close-All)	@param price	@param triggerPrice	@param workingType stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"	@param priceMatch only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`	@param closePosition `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.	@param priceProtect \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param reduceOnly \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`	@param activationPrice Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)	@param callbackRate Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1%	@param clientAlgoId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param selfTradePreventionMode `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`	@param goodTillDate order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000	@param recvWindow
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
	if r.activationPrice != nil {
		localVarQueryParams["activationPrice"] = *r.activationPrice
	}
	if r.callbackRate != nil {
		localVarQueryParams["callbackRate"] = *r.callbackRate
	}
	if r.clientAlgoId != nil {
		localVarQueryParams["clientAlgoId"] = *r.clientAlgoId
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
	type_                   *string
	id                      *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float32
	reduceOnly              *string
	price                   *float32
	newClientOrderId        *string
	stopPrice               *float32
	closePosition           *string
	activationPrice         *float32
	callbackRate            *float32
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceProtect            *string
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiNewOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewOrderRequest) Type(type_ string) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiNewOrderRequest) Id(id string) ApiNewOrderRequest {
	r.id = &id
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly string) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewOrderRequest) Price(price float32) ApiNewOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) StopPrice(stopPrice float32) ApiNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewOrderRequest) ClosePosition(closePosition string) ApiNewOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiNewOrderRequest) ActivationPrice(activationPrice float32) ApiNewOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1%
func (r ApiNewOrderRequest) CallbackRate(callbackRate float32) ApiNewOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) PriceProtect(priceProtect string) ApiNewOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
func (r ApiNewOrderRequest) GoodTillDate(goodTillDate int64) ApiNewOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

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
NewOrder New Order(TRADE)
/order.place

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/New-Order

@param symbol	@param side `SELL`, `BUY`	@param type_	@param id Unique WebSocket request ID.	@param positionSide Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.	@param timeInForce	@param quantity Cannot be sent with `closePosition`=`true`(Close-All)	@param reduceOnly \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`	@param price	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param stopPrice Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param closePosition `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.	@param activationPrice Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)	@param callbackRate Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1%	@param workingType stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"	@param priceProtect \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param newOrderRespType \"ACK\", \"RESULT\", default \"ACK\"	@param priceMatch only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`	@param selfTradePreventionMode `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`	@param goodTillDate order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000	@param recvWindow
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
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.reduceOnly != nil {
		localVarQueryParams["reduceOnly"] = *r.reduceOnly
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.closePosition != nil {
		localVarQueryParams["closePosition"] = *r.closePosition
	}
	if r.activationPrice != nil {
		localVarQueryParams["activationPrice"] = *r.activationPrice
	}
	if r.callbackRate != nil {
		localVarQueryParams["callbackRate"] = *r.callbackRate
	}
	if r.workingType != nil {
		localVarQueryParams["workingType"] = *r.workingType
	}
	if r.priceProtect != nil {
		localVarQueryParams["priceProtect"] = *r.priceProtect
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

// Unique WebSocket request ID.
func (r ApiPositionInformationRequest) Id(id string) ApiPositionInformationRequest {
	r.id = &id
	return r
}

func (r ApiPositionInformationRequest) Symbol(symbol string) ApiPositionInformationRequest {
	r.symbol = &symbol
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Position-Information

@param id Unique WebSocket request ID.	@param symbol	@param recvWindow
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

// Unique WebSocket request ID.
func (r ApiPositionInformationV2Request) Id(id string) ApiPositionInformationV2Request {
	r.id = &id
	return r
}

func (r ApiPositionInformationV2Request) Symbol(symbol string) ApiPositionInformationV2Request {
	r.symbol = &symbol
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Position-Info-V2

@param id Unique WebSocket request ID.	@param symbol	@param recvWindow
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

func (r ApiQueryOrderRequest) Symbol(symbol string) ApiQueryOrderRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiQueryOrderRequest) Id(id string) ApiQueryOrderRequest {
	r.id = &id
	return r
}

func (r ApiQueryOrderRequest) OrderId(orderId int64) ApiQueryOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

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

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/websocket-api/Query-Order

@param symbol	@param id Unique WebSocket request ID.	@param orderId	@param origClientOrderId	@param recvWindow
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
