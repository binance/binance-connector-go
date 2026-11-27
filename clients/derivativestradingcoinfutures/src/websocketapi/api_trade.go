/*
Futures (COIN-M) WebSocket API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package binancederivativestradingcoinfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#cancel-algo-order

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

func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
	return r
}

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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#cancel-order

@param symbol	@param id	@param orderId	@param origClientOrderId	@param recvWindow
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
	quantity          *float64
	price             *float64
	id                *string
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.ModifyOrderPriceMatchParameter
	modifyId          *int64
	recvWindow        *int64
}

func (r ApiModifyOrderRequest) Symbol(symbol string) ApiModifyOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiModifyOrderRequest) Side(side models.ModifyOrderSideParameter) ApiModifyOrderRequest {
	r.side = &side
	return r
}

// Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60;
func (r ApiModifyOrderRequest) Quantity(quantity float64) ApiModifyOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiModifyOrderRequest) Price(price float64) ApiModifyOrderRequest {
	r.price = &price
	return r
}

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

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiModifyOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness.
func (r ApiModifyOrderRequest) ModifyId(modifyId int64) ApiModifyOrderRequest {
	r.modifyId = &modifyId
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#modify-order

@param symbol	@param side	@param quantity Order quantity, cannot be sent with `closePosition=true`	@param price	@param id	@param orderId	@param origClientOrderId	@param priceMatch only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price`	@param modifyId User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness.	@param recvWindow
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
	if r.modifyId != nil {
		localVarQueryParams["modifyId"] = *r.modifyId
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
	quantity                *float64
	price                   *float64
	triggerPrice            *float64
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	closePosition           *models.NewAlgoOrderClosePositionParameter
	priceProtect            *models.NewAlgoOrderClosePositionParameter
	reduceOnly              *models.NewAlgoOrderClosePositionParameter
	activatePrice           *float64
	callbackRate            *float64
	clientAlgoId            *string
	newOrderRespType        *models.NewAlgoOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
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

// &#x60;IOC&#x60; or &#x60;GTC&#x60; or &#x60;FOK&#x60; or &#x60;GTX&#x60;, default &#x60;GTC&#x60;
func (r ApiNewAlgoOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewAlgoOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewAlgoOrderRequest) Quantity(quantity float64) ApiNewAlgoOrderRequest {
	r.quantity = &quantity
	return r
}

// Price.
func (r ApiNewAlgoOrderRequest) Price(price float64) ApiNewAlgoOrderRequest {
	r.price = &price
	return r
}

// Trigger Price.
func (r ApiNewAlgoOrderRequest) TriggerPrice(triggerPrice float64) ApiNewAlgoOrderRequest {
	r.triggerPrice = &triggerPrice
	return r
}

// triggerPrice triggered by: &#x60;MARK_PRICE&#x60;, &#x60;CONTRACT_PRICE&#x60;. Default &#x60;CONTRACT_PRICE&#x60;
func (r ApiNewAlgoOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewAlgoOrderRequest {
	r.workingType = &workingType
	return r
}

// only available for LIMIT/STOP/TAKE_PROFIT order; Can&#39;t be passed together with price
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
func (r ApiNewAlgoOrderRequest) ActivatePrice(activatePrice float64) ApiNewAlgoOrderRequest {
	r.activatePrice = &activatePrice
	return r
}

// Used with TRAILING_STOP_MARKET orders
func (r ApiNewAlgoOrderRequest) CallbackRate(callbackRate float64) ApiNewAlgoOrderRequest {
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

// &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;: expire maker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewAlgoOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiNewAlgoOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#new-algo-order

@param algoType Only support `CONDITIONAL`	@param symbol Symbol.	@param side Side.	@param type_ For `CONDITIONAL` algoType, `STOP_MARKET`/`TAKE_PROFIT_MARKET`/`STOP`/`TAKE_PROFIT`/`TRAILING_STOP_MARKET` as order type	@param id Id.	@param positionSide Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode.	@param timeInForce `IOC` or `GTC` or `FOK` or `GTX`, default `GTC`	@param quantity quantity measured by contract number, Cannot be sent with `closePosition`=`true`(Close-All)	@param price Price.	@param triggerPrice Trigger Price.	@param workingType triggerPrice triggered by: `MARK_PRICE`, `CONTRACT_PRICE`. Default `CONTRACT_PRICE`	@param priceMatch only available for LIMIT/STOP/TAKE_PROFIT order; Can't be passed together with price	@param closePosition Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.	@param priceProtect Used with STOP_MARKET or TAKE_PROFIT_MARKET order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol.	@param reduceOnly Cannot be sent in Hedge Mode; cannot be sent with closePosition=true	@param activatePrice Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)	@param callbackRate Used with TRAILING_STOP_MARKET orders	@param clientAlgoId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param newOrderRespType \"ACK\", \"RESULT\", default \"ACK\"	@param selfTradePreventionMode `EXPIRE_TAKER`: expire taker order when STP triggers/ `EXPIRE_MAKER`: expire maker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers; default `NONE`	@param recvWindow Recv Window.
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
	quantity                *float64
	reduceOnly              *models.NewAlgoOrderClosePositionParameter
	price                   *float64
	newClientOrderId        *string
	stopPrice               *float64
	closePosition           *models.NewAlgoOrderClosePositionParameter
	activationPrice         *float64
	callbackRate            *float64
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceProtect            *models.NewAlgoOrderClosePositionParameter
	newOrderRespType        *models.NewAlgoOrderNewOrderRespTypeParameter
	priceMatch              *models.ModifyOrderPriceMatchParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *int64
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewOrderRequest) Side(side models.ModifyOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

// **After CM migration, stop-type values (&#x60;STOP&#x60;, &#x60;STOP_MARKET&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_MARKET&#x60;, &#x60;TRAILING_STOP_MARKET&#x60;) are no longer accepted and will return &#x60;-4120&#x60;. Use the REST &#x60;/dapi/v1/algoOrder&#x60; endpoint instead.**
func (r ApiNewOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

func (r ApiNewOrderRequest) Id(id string) ApiNewOrderRequest {
	r.id = &id
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode.  It must be sent in Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewOrderRequest) Quantity(quantity float64) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60; (Close-All)\&quot;
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewOrderRequest) Price(price float64) ApiNewOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) StopPrice(stopPrice float64) ApiNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewOrderRequest) ClosePosition(closePosition models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different workingType)
func (r ApiNewOrderRequest) ActivationPrice(activationPrice float64) ApiNewOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1%
func (r ApiNewOrderRequest) CallbackRate(callbackRate float64) ApiNewOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiNewOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewOrderRequest {
	r.workingType = &workingType
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.&#39;
func (r ApiNewOrderRequest) PriceProtect(priceProtect models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewAlgoOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.ModifyOrderPriceMatchParameter) ApiNewOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;NONE&#x60;: No STP / &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
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
NewOrder New Order (TRADE)
/order.place

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#new-order

@param symbol	@param side	@param type_ **After CM migration, stop-type values (`STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`) are no longer accepted and will return `-4120`. Use the REST `/dapi/v1/algoOrder` endpoint instead.**	@param id	@param positionSide Default `BOTH` for One-way Mode; `LONG` or `SHORT` for Hedge Mode.  It must be sent in Hedge Mode.	@param timeInForce	@param quantity Quantity measured by contract number, Cannot be sent with `closePosition`=`true`	@param reduceOnly Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true` (Close-All)\"	@param price	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`	@param stopPrice Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.	@param closePosition `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.	@param activationPrice Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different workingType)	@param callbackRate Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1%	@param workingType stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"	@param priceProtect Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.'	@param newOrderRespType	@param priceMatch only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price`	@param selfTradePreventionMode `NONE`: No STP / `EXPIRE_TAKER`: expire taker order when STP triggers/ `EXPIRE_MAKER`: expire taker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers	@param recvWindow
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
	ApiService  *TradeAPIService
	id          *string
	marginAsset *string
	pair        *string
	recvWindow  *int64
}

func (r ApiPositionInformationRequest) Id(id string) ApiPositionInformationRequest {
	r.id = &id
	return r
}

func (r ApiPositionInformationRequest) MarginAsset(marginAsset string) ApiPositionInformationRequest {
	r.marginAsset = &marginAsset
	return r
}

func (r ApiPositionInformationRequest) Pair(pair string) ApiPositionInformationRequest {
	r.pair = &pair
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#position-information

@param id	@param marginAsset	@param pair	@param recvWindow
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
	if r.marginAsset != nil {
		localVarQueryParams["marginAsset"] = *r.marginAsset
	}
	if r.pair != nil {
		localVarQueryParams["pair"] = *r.pair
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/trade#query-order

@param symbol	@param id	@param orderId	@param origClientOrderId	@param recvWindow
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
