/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiOpenOrdersCancelAllRequest struct {
	ApiService *TradeAPIService
	symbol     *string
	id         *string
	recvWindow *float32
}

func (r ApiOpenOrdersCancelAllRequest) Symbol(symbol string) ApiOpenOrdersCancelAllRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiOpenOrdersCancelAllRequest) Id(id string) ApiOpenOrdersCancelAllRequest {
	r.id = &id
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOpenOrdersCancelAllRequest) RecvWindow(recvWindow float32) ApiOpenOrdersCancelAllRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOpenOrdersCancelAllRequest) Execute() (*common.ResponseOrRaw[models.OpenOrdersCancelAllResponse], error) {
	respChan, errChan, err := r.ApiService.OpenOrdersCancelAllExecute(r)
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

func (r ApiOpenOrdersCancelAllRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OpenOrdersCancelAllResponse], chan error, error) {
	return r.ApiService.OpenOrdersCancelAllExecute(r)
}

/*
OpenOrdersCancelAll WebSocket Cancel open orders
/openOrders.cancelAll

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#cancel-open-orders-trade

@param symbol	@param id Unique WebSocket request ID.	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOpenOrdersCancelAllRequest
*/
func (a *TradeAPIService) OpenOrdersCancelAll() ApiOpenOrdersCancelAllRequest {
	return ApiOpenOrdersCancelAllRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OpenOrdersCancelAllResponse
func (a *TradeAPIService) OpenOrdersCancelAllExecute(r ApiOpenOrdersCancelAllRequest) (chan *common.ResponseOrRaw[models.OpenOrdersCancelAllResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/openOrders.cancelAll"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OpenOrdersCancelAllResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderAmendKeepPriorityRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	newQty            *float32
	id                *string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *float32
}

func (r ApiOrderAmendKeepPriorityRequest) Symbol(symbol string) ApiOrderAmendKeepPriorityRequest {
	r.symbol = &symbol
	return r
}

// &#x60;newQty&#x60; must be greater than 0 and less than the order&#39;s quantity.
func (r ApiOrderAmendKeepPriorityRequest) NewQty(newQty float32) ApiOrderAmendKeepPriorityRequest {
	r.newQty = &newQty
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderAmendKeepPriorityRequest) Id(id string) ApiOrderAmendKeepPriorityRequest {
	r.id = &id
	return r
}

// &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent
func (r ApiOrderAmendKeepPriorityRequest) OrderId(orderId int64) ApiOrderAmendKeepPriorityRequest {
	r.orderId = &orderId
	return r
}

// &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent
func (r ApiOrderAmendKeepPriorityRequest) OrigClientOrderId(origClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderAmendKeepPriorityRequest) NewClientOrderId(newClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderAmendKeepPriorityRequest) RecvWindow(recvWindow float32) ApiOrderAmendKeepPriorityRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderAmendKeepPriorityRequest) Execute() (*common.ResponseOrRaw[models.OrderAmendKeepPriorityResponse], error) {
	respChan, errChan, err := r.ApiService.OrderAmendKeepPriorityExecute(r)
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

func (r ApiOrderAmendKeepPriorityRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderAmendKeepPriorityResponse], chan error, error) {
	return r.ApiService.OrderAmendKeepPriorityExecute(r)
}

/*
OrderAmendKeepPriority WebSocket Order Amend Keep Priority
/order.amend.keepPriority

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#order-amend-keep-priority-trade

@param symbol	@param newQty `newQty` must be greater than 0 and less than the order's quantity. 	@param id Unique WebSocket request ID.	@param orderId `orderId`or`origClientOrderId`mustbesent	@param origClientOrderId `orderId`or`origClientOrderId`mustbesent	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderAmendKeepPriorityRequest
*/
func (a *TradeAPIService) OrderAmendKeepPriority() ApiOrderAmendKeepPriorityRequest {
	return ApiOrderAmendKeepPriorityRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderAmendKeepPriorityResponse
func (a *TradeAPIService) OrderAmendKeepPriorityExecute(r ApiOrderAmendKeepPriorityRequest) (chan *common.ResponseOrRaw[models.OrderAmendKeepPriorityResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.newQty == nil {
		return nil, nil, common.ReportError("newQty is required and must be specified")
	}
	localVarQueryParams["newQty"] = *r.newQty

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderId != nil {
		localVarQueryParams["orderId"] = *r.orderId
	}
	if r.origClientOrderId != nil {
		localVarQueryParams["origClientOrderId"] = *r.origClientOrderId
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.amend.keepPriority"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderAmendKeepPriorityResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderCancelRequest struct {
	ApiService         *TradeAPIService
	symbol             *string
	id                 *string
	orderId            *int64
	origClientOrderId  *string
	newClientOrderId   *string
	cancelRestrictions *models.OrderCancelCancelRestrictionsParameter
	recvWindow         *float32
}

func (r ApiOrderCancelRequest) Symbol(symbol string) ApiOrderCancelRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderCancelRequest) Id(id string) ApiOrderCancelRequest {
	r.id = &id
	return r
}

// &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent
func (r ApiOrderCancelRequest) OrderId(orderId int64) ApiOrderCancelRequest {
	r.orderId = &orderId
	return r
}

// &#x60;orderId&#x60;or&#x60;origClientOrderId&#x60;mustbesent
func (r ApiOrderCancelRequest) OrigClientOrderId(origClientOrderId string) ApiOrderCancelRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderCancelRequest) NewClientOrderId(newClientOrderId string) ApiOrderCancelRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiOrderCancelRequest) CancelRestrictions(cancelRestrictions models.OrderCancelCancelRestrictionsParameter) ApiOrderCancelRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderCancelRequest) RecvWindow(recvWindow float32) ApiOrderCancelRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderCancelRequest) Execute() (*common.ResponseOrRaw[models.OrderCancelResponse], error) {
	respChan, errChan, err := r.ApiService.OrderCancelExecute(r)
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

func (r ApiOrderCancelRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderCancelResponse], chan error, error) {
	return r.ApiService.OrderCancelExecute(r)
}

/*
OrderCancel WebSocket Cancel order
/order.cancel

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#cancel-order-trade

@param symbol	@param id Unique WebSocket request ID.	@param orderId `orderId`or`origClientOrderId`mustbesent	@param origClientOrderId `orderId`or`origClientOrderId`mustbesent	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param cancelRestrictions	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderCancelRequest
*/
func (a *TradeAPIService) OrderCancel() ApiOrderCancelRequest {
	return ApiOrderCancelRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderCancelResponse
func (a *TradeAPIService) OrderCancelExecute(r ApiOrderCancelRequest) (chan *common.ResponseOrRaw[models.OrderCancelResponse], chan error, error) {
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
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.cancelRestrictions != nil {
		localVarQueryParams["cancelRestrictions"] = *r.cancelRestrictions
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

	return SendMessage[models.OrderCancelResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderCancelReplaceRequest struct {
	ApiService                 *TradeAPIService
	symbol                     *string
	cancelReplaceMode          *models.OrderCancelReplaceCancelReplaceModeParameter
	side                       *models.OrderCancelReplaceSideParameter
	type_                      *models.OrderCancelReplaceTypeParameter
	id                         *string
	cancelOrderId              *int64
	cancelOrigClientOrderId    *string
	cancelNewClientOrderId     *string
	timeInForce                *models.OrderCancelReplaceTimeInForceParameter
	price                      *float32
	quantity                   *float32
	quoteOrderQty              *float32
	newClientOrderId           *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice                  *float32
	trailingDelta              *float32
	icebergQty                 *float32
	strategyId                 *int64
	strategyType               *int32
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	cancelRestrictions         *models.OrderCancelCancelRestrictionsParameter
	orderRateLimitExceededMode *models.OrderCancelReplaceOrderRateLimitExceededModeParameter
	pegPriceType               *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue             *int32
	pegOffsetType              *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow                 *float32
}

func (r ApiOrderCancelReplaceRequest) Symbol(symbol string) ApiOrderCancelReplaceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderCancelReplaceRequest) CancelReplaceMode(cancelReplaceMode models.OrderCancelReplaceCancelReplaceModeParameter) ApiOrderCancelReplaceRequest {
	r.cancelReplaceMode = &cancelReplaceMode
	return r
}

func (r ApiOrderCancelReplaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderCancelReplaceRequest {
	r.side = &side
	return r
}

func (r ApiOrderCancelReplaceRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderCancelReplaceRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderCancelReplaceRequest) Id(id string) ApiOrderCancelReplaceRequest {
	r.id = &id
	return r
}

// Cancel order by orderId
func (r ApiOrderCancelReplaceRequest) CancelOrderId(cancelOrderId int64) ApiOrderCancelReplaceRequest {
	r.cancelOrderId = &cancelOrderId
	return r
}

func (r ApiOrderCancelReplaceRequest) CancelOrigClientOrderId(cancelOrigClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return r
}

// New ID for the canceled order. Automatically generated if not sent
func (r ApiOrderCancelReplaceRequest) CancelNewClientOrderId(cancelNewClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelNewClientOrderId = &cancelNewClientOrderId
	return r
}

func (r ApiOrderCancelReplaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderCancelReplaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderCancelReplaceRequest) Price(price float32) ApiOrderCancelReplaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderCancelReplaceRequest) Quantity(quantity float32) ApiOrderCancelReplaceRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderCancelReplaceRequest) QuoteOrderQty(quoteOrderQty float32) ApiOrderCancelReplaceRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderCancelReplaceRequest) NewClientOrderId(newClientOrderId string) ApiOrderCancelReplaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiOrderCancelReplaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderCancelReplaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderCancelReplaceRequest) StopPrice(stopPrice float32) ApiOrderCancelReplaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See Trailing Stop order FAQ
func (r ApiOrderCancelReplaceRequest) TrailingDelta(trailingDelta float32) ApiOrderCancelReplaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

func (r ApiOrderCancelReplaceRequest) IcebergQty(icebergQty float32) ApiOrderCancelReplaceRequest {
	r.icebergQty = &icebergQty
	return r
}

// Arbitrary numeric value identifying the order within an order strategy.
func (r ApiOrderCancelReplaceRequest) StrategyId(strategyId int64) ApiOrderCancelReplaceRequest {
	r.strategyId = &strategyId
	return r
}

// Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderCancelReplaceRequest) StrategyType(strategyType int32) ApiOrderCancelReplaceRequest {
	r.strategyType = &strategyType
	return r
}

func (r ApiOrderCancelReplaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderCancelReplaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiOrderCancelReplaceRequest) CancelRestrictions(cancelRestrictions models.OrderCancelCancelRestrictionsParameter) ApiOrderCancelReplaceRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

func (r ApiOrderCancelReplaceRequest) OrderRateLimitExceededMode(orderRateLimitExceededMode models.OrderCancelReplaceOrderRateLimitExceededModeParameter) ApiOrderCancelReplaceRequest {
	r.orderRateLimitExceededMode = &orderRateLimitExceededMode
	return r
}

func (r ApiOrderCancelReplaceRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level to peg the price to (max: 100)       See Pegged Orders
func (r ApiOrderCancelReplaceRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderCancelReplaceRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiOrderCancelReplaceRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderCancelReplaceRequest) RecvWindow(recvWindow float32) ApiOrderCancelReplaceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderCancelReplaceRequest) Execute() (*common.ResponseOrRaw[models.OrderCancelReplaceResponse], error) {
	respChan, errChan, err := r.ApiService.OrderCancelReplaceExecute(r)
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

func (r ApiOrderCancelReplaceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderCancelReplaceResponse], chan error, error) {
	return r.ApiService.OrderCancelReplaceExecute(r)
}

/*
OrderCancelReplace WebSocket Cancel and replace order
/order.cancelReplace

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#cancel-and-replace-order-trade

@param symbol	@param cancelReplaceMode	@param side	@param type_	@param id Unique WebSocket request ID.	@param cancelOrderId Cancel order by orderId	@param cancelOrigClientOrderId	@param cancelNewClientOrderId New ID for the canceled order. Automatically generated if not sent	@param timeInForce	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param newOrderRespType	@param stopPrice	@param trailingDelta See Trailing Stop order FAQ	@param icebergQty	@param strategyId Arbitrary numeric value identifying the order within an order strategy.	@param strategyType Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.	@param selfTradePreventionMode	@param cancelRestrictions	@param orderRateLimitExceededMode	@param pegPriceType	@param pegOffsetValue Price level to peg the price to (max: 100)       See Pegged Orders	@param pegOffsetType	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderCancelReplaceRequest
*/
func (a *TradeAPIService) OrderCancelReplace() ApiOrderCancelReplaceRequest {
	return ApiOrderCancelReplaceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderCancelReplaceResponse
func (a *TradeAPIService) OrderCancelReplaceExecute(r ApiOrderCancelReplaceRequest) (chan *common.ResponseOrRaw[models.OrderCancelReplaceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.cancelReplaceMode == nil {
		return nil, nil, common.ReportError("cancelReplaceMode is required and must be specified")
	}
	localVarQueryParams["cancelReplaceMode"] = *r.cancelReplaceMode

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
	if r.cancelOrderId != nil {
		localVarQueryParams["cancelOrderId"] = *r.cancelOrderId
	}
	if r.cancelOrigClientOrderId != nil {
		localVarQueryParams["cancelOrigClientOrderId"] = *r.cancelOrigClientOrderId
	}
	if r.cancelNewClientOrderId != nil {
		localVarQueryParams["cancelNewClientOrderId"] = *r.cancelNewClientOrderId
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.quoteOrderQty != nil {
		localVarQueryParams["quoteOrderQty"] = *r.quoteOrderQty
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.trailingDelta != nil {
		localVarQueryParams["trailingDelta"] = *r.trailingDelta
	}
	if r.icebergQty != nil {
		localVarQueryParams["icebergQty"] = *r.icebergQty
	}
	if r.strategyId != nil {
		localVarQueryParams["strategyId"] = *r.strategyId
	}
	if r.strategyType != nil {
		localVarQueryParams["strategyType"] = *r.strategyType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.cancelRestrictions != nil {
		localVarQueryParams["cancelRestrictions"] = *r.cancelRestrictions
	}
	if r.orderRateLimitExceededMode != nil {
		localVarQueryParams["orderRateLimitExceededMode"] = *r.orderRateLimitExceededMode
	}
	if r.pegPriceType != nil {
		localVarQueryParams["pegPriceType"] = *r.pegPriceType
	}
	if r.pegOffsetValue != nil {
		localVarQueryParams["pegOffsetValue"] = *r.pegOffsetValue
	}
	if r.pegOffsetType != nil {
		localVarQueryParams["pegOffsetType"] = *r.pegOffsetType
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.cancelReplace"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderCancelReplaceResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListCancelRequest struct {
	ApiService        *TradeAPIService
	symbol            *string
	id                *string
	orderListId       *int32
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *float32
}

func (r ApiOrderListCancelRequest) Symbol(symbol string) ApiOrderListCancelRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListCancelRequest) Id(id string) ApiOrderListCancelRequest {
	r.id = &id
	return r
}

// Cancel order list by orderListId
func (r ApiOrderListCancelRequest) OrderListId(orderListId int32) ApiOrderListCancelRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiOrderListCancelRequest) ListClientOrderId(listClientOrderId string) ApiOrderListCancelRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderListCancelRequest) NewClientOrderId(newClientOrderId string) ApiOrderListCancelRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListCancelRequest) RecvWindow(recvWindow float32) ApiOrderListCancelRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListCancelRequest) Execute() (*common.ResponseOrRaw[models.OrderListCancelResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListCancelExecute(r)
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

func (r ApiOrderListCancelRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListCancelResponse], chan error, error) {
	return r.ApiService.OrderListCancelExecute(r)
}

/*
OrderListCancel WebSocket Cancel Order list
/orderList.cancel

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#cancel-order-list-trade

@param symbol	@param id Unique WebSocket request ID.	@param orderListId Cancel order list by orderListId	@param listClientOrderId	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListCancelRequest
*/
func (a *TradeAPIService) OrderListCancel() ApiOrderListCancelRequest {
	return ApiOrderListCancelRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListCancelResponse
func (a *TradeAPIService) OrderListCancelExecute(r ApiOrderListCancelRequest) (chan *common.ResponseOrRaw[models.OrderListCancelResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.orderListId != nil {
		localVarQueryParams["orderListId"] = *r.orderListId
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.cancel"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListCancelResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	price                   *float32
	quantity                *float32
	id                      *string
	listClientOrderId       *string
	limitClientOrderId      *string
	limitIcebergQty         *float32
	limitStrategyId         *int64
	limitStrategyType       *int32
	stopPrice               *float32
	trailingDelta           *int32
	stopClientOrderId       *string
	stopLimitPrice          *float32
	stopLimitTimeInForce    *models.OrderListPlaceStopLimitTimeInForceParameter
	stopIcebergQty          *float32
	stopStrategyId          *int64
	stopStrategyType        *int32
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiOrderListPlaceRequest) Symbol(symbol string) ApiOrderListPlaceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderListPlaceRequest {
	r.side = &side
	return r
}

// Price for the limit order
func (r ApiOrderListPlaceRequest) Price(price float32) ApiOrderListPlaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderListPlaceRequest) Quantity(quantity float32) ApiOrderListPlaceRequest {
	r.quantity = &quantity
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceRequest) Id(id string) ApiOrderListPlaceRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent
func (r ApiOrderListPlaceRequest) LimitClientOrderId(limitClientOrderId string) ApiOrderListPlaceRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

func (r ApiOrderListPlaceRequest) LimitIcebergQty(limitIcebergQty float32) ApiOrderListPlaceRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

// Arbitrary numeric value identifying the limit order within an order strategy.
func (r ApiOrderListPlaceRequest) LimitStrategyId(limitStrategyId int64) ApiOrderListPlaceRequest {
	r.limitStrategyId = &limitStrategyId
	return r
}

// &lt;p&gt;Arbitrary numeric value identifying the limit order strategy.&lt;/p&gt;&lt;p&gt;Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.&lt;/p&gt;
func (r ApiOrderListPlaceRequest) LimitStrategyType(limitStrategyType int32) ApiOrderListPlaceRequest {
	r.limitStrategyType = &limitStrategyType
	return r
}

func (r ApiOrderListPlaceRequest) StopPrice(stopPrice float32) ApiOrderListPlaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)
func (r ApiOrderListPlaceRequest) TrailingDelta(trailingDelta int32) ApiOrderListPlaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent
func (r ApiOrderListPlaceRequest) StopClientOrderId(stopClientOrderId string) ApiOrderListPlaceRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

func (r ApiOrderListPlaceRequest) StopLimitPrice(stopLimitPrice float32) ApiOrderListPlaceRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

func (r ApiOrderListPlaceRequest) StopLimitTimeInForce(stopLimitTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

func (r ApiOrderListPlaceRequest) StopIcebergQty(stopIcebergQty float32) ApiOrderListPlaceRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

// Arbitrary numeric value identifying the stop order within an order strategy.
func (r ApiOrderListPlaceRequest) StopStrategyId(stopStrategyId int64) ApiOrderListPlaceRequest {
	r.stopStrategyId = &stopStrategyId
	return r
}

// &lt;p&gt;Arbitrary numeric value identifying the stop order strategy.&lt;/p&gt;&lt;p&gt;Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.&lt;/p&gt;
func (r ApiOrderListPlaceRequest) StopStrategyType(stopStrategyType int32) ApiOrderListPlaceRequest {
	r.stopStrategyType = &stopStrategyType
	return r
}

func (r ApiOrderListPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceExecute(r)
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

func (r ApiOrderListPlaceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceResponse], chan error, error) {
	return r.ApiService.OrderListPlaceExecute(r)
}

/*
	OrderListPlace WebSocket Place new OCO - Deprecated
	/orderList.place

	https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-oco---deprecated-trade

	@param symbol	@param side	@param price Price for the limit order	@param quantity	@param id Unique WebSocket request ID.	@param listClientOrderId	@param limitClientOrderId Arbitrary unique ID among open orders for the limit order. Automatically generated if not sent	@param limitIcebergQty	@param limitStrategyId Arbitrary numeric value identifying the limit order within an order strategy.	@param limitStrategyType <p>Arbitrary numeric value identifying the limit order strategy.</p><p>Values smaller than `1000000` are reserved and cannot be used.</p>	@param stopPrice	@param trailingDelta See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)	@param stopClientOrderId Arbitrary unique ID among open orders for the stop order. Automatically generated if not sent	@param stopLimitPrice	@param stopLimitTimeInForce	@param stopIcebergQty	@param stopStrategyId Arbitrary numeric value identifying the stop order within an order strategy.	@param stopStrategyType <p>Arbitrary numeric value identifying the stop order strategy.</p><p>Values smaller than `1000000` are reserved and cannot be used.</p>	@param newOrderRespType	@param selfTradePreventionMode	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
	@return ApiOrderListPlaceRequest

Deprecated
*/
func (a *TradeAPIService) OrderListPlace() ApiOrderListPlaceRequest {
	return ApiOrderListPlaceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceResponse
//
// Deprecated
func (a *TradeAPIService) OrderListPlaceExecute(r ApiOrderListPlaceRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.side == nil {
		return nil, nil, common.ReportError("side is required and must be specified")
	}
	localVarQueryParams["side"] = *r.side

	if r.price == nil {
		return nil, nil, common.ReportError("price is required and must be specified")
	}
	localVarQueryParams["price"] = *r.price

	if r.quantity == nil {
		return nil, nil, common.ReportError("quantity is required and must be specified")
	}
	localVarQueryParams["quantity"] = *r.quantity

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.limitClientOrderId != nil {
		localVarQueryParams["limitClientOrderId"] = *r.limitClientOrderId
	}
	if r.limitIcebergQty != nil {
		localVarQueryParams["limitIcebergQty"] = *r.limitIcebergQty
	}
	if r.limitStrategyId != nil {
		localVarQueryParams["limitStrategyId"] = *r.limitStrategyId
	}
	if r.limitStrategyType != nil {
		localVarQueryParams["limitStrategyType"] = *r.limitStrategyType
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.trailingDelta != nil {
		localVarQueryParams["trailingDelta"] = *r.trailingDelta
	}
	if r.stopClientOrderId != nil {
		localVarQueryParams["stopClientOrderId"] = *r.stopClientOrderId
	}
	if r.stopLimitPrice != nil {
		localVarQueryParams["stopLimitPrice"] = *r.stopLimitPrice
	}
	if r.stopLimitTimeInForce != nil {
		localVarQueryParams["stopLimitTimeInForce"] = *r.stopLimitTimeInForce
	}
	if r.stopIcebergQty != nil {
		localVarQueryParams["stopIcebergQty"] = *r.stopIcebergQty
	}
	if r.stopStrategyId != nil {
		localVarQueryParams["stopStrategyId"] = *r.stopStrategyId
	}
	if r.stopStrategyType != nil {
		localVarQueryParams["stopStrategyType"] = *r.stopStrategyType
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
		"method": "/orderList.place"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceOcoRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	quantity                *float32
	aboveType               *models.OrderListPlaceOcoAboveTypeParameter
	belowType               *models.OrderListPlaceOcoBelowTypeParameter
	id                      *string
	listClientOrderId       *string
	aboveClientOrderId      *string
	aboveIcebergQty         *int64
	abovePrice              *float32
	aboveStopPrice          *float32
	aboveTrailingDelta      *int64
	aboveTimeInForce        *models.OrderListPlaceStopLimitTimeInForceParameter
	aboveStrategyId         *int64
	aboveStrategyType       *int32
	abovePegPriceType       *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	abovePegOffsetType      *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	abovePegOffsetValue     *int32
	belowClientOrderId      *string
	belowIcebergQty         *int64
	belowPrice              *float32
	belowStopPrice          *float32
	belowTrailingDelta      *int64
	belowTimeInForce        *models.OrderListPlaceStopLimitTimeInForceParameter
	belowStrategyId         *int64
	belowStrategyType       *int32
	belowPegPriceType       *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	belowPegOffsetType      *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	belowPegOffsetValue     *int32
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiOrderListPlaceOcoRequest) Symbol(symbol string) ApiOrderListPlaceOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOcoRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOcoRequest {
	r.side = &side
	return r
}

func (r ApiOrderListPlaceOcoRequest) Quantity(quantity float32) ApiOrderListPlaceOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderListPlaceOcoRequest) AboveType(aboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOcoRequest {
	r.aboveType = &aboveType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowType(belowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowType = &belowType
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceOcoRequest) Id(id string) ApiOrderListPlaceOcoRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceOcoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
func (r ApiOrderListPlaceOcoRequest) AboveClientOrderId(aboveClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.aboveClientOrderId = &aboveClientOrderId
	return r
}

// Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListPlaceOcoRequest) AboveIcebergQty(aboveIcebergQty int64) ApiOrderListPlaceOcoRequest {
	r.aboveIcebergQty = &aboveIcebergQty
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOcoRequest) AbovePrice(abovePrice float32) ApiOrderListPlaceOcoRequest {
	r.abovePrice = &abovePrice
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOcoRequest) AboveStopPrice(aboveStopPrice float32) ApiOrderListPlaceOcoRequest {
	r.aboveStopPrice = &aboveStopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderListPlaceOcoRequest) AboveTrailingDelta(aboveTrailingDelta int64) ApiOrderListPlaceOcoRequest {
	r.aboveTrailingDelta = &aboveTrailingDelta
	return r
}

func (r ApiOrderListPlaceOcoRequest) AboveTimeInForce(aboveTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOcoRequest {
	r.aboveTimeInForce = &aboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the above order within an order strategy.
func (r ApiOrderListPlaceOcoRequest) AboveStrategyId(aboveStrategyId int64) ApiOrderListPlaceOcoRequest {
	r.aboveStrategyId = &aboveStrategyId
	return r
}

// Arbitrary numeric value identifying the above order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOcoRequest) AboveStrategyType(aboveStrategyType int32) ApiOrderListPlaceOcoRequest {
	r.aboveStrategyType = &aboveStrategyType
	return r
}

func (r ApiOrderListPlaceOcoRequest) AbovePegPriceType(abovePegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOcoRequest {
	r.abovePegPriceType = &abovePegPriceType
	return r
}

func (r ApiOrderListPlaceOcoRequest) AbovePegOffsetType(abovePegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOcoRequest {
	r.abovePegOffsetType = &abovePegOffsetType
	return r
}

func (r ApiOrderListPlaceOcoRequest) AbovePegOffsetValue(abovePegOffsetValue int32) ApiOrderListPlaceOcoRequest {
	r.abovePegOffsetValue = &abovePegOffsetValue
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowClientOrderId(belowClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.belowClientOrderId = &belowClientOrderId
	return r
}

// Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListPlaceOcoRequest) BelowIcebergQty(belowIcebergQty int64) ApiOrderListPlaceOcoRequest {
	r.belowIcebergQty = &belowIcebergQty
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOcoRequest) BelowPrice(belowPrice float32) ApiOrderListPlaceOcoRequest {
	r.belowPrice = &belowPrice
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;belowStopPrice&#x60; or &#x60;belowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOcoRequest) BelowStopPrice(belowStopPrice float32) ApiOrderListPlaceOcoRequest {
	r.belowStopPrice = &belowStopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderListPlaceOcoRequest) BelowTrailingDelta(belowTrailingDelta int64) ApiOrderListPlaceOcoRequest {
	r.belowTrailingDelta = &belowTrailingDelta
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowTimeInForce(belowTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOcoRequest {
	r.belowTimeInForce = &belowTimeInForce
	return r
}

// Arbitrary numeric value identifying the below order within an order strategy.
func (r ApiOrderListPlaceOcoRequest) BelowStrategyId(belowStrategyId int64) ApiOrderListPlaceOcoRequest {
	r.belowStrategyId = &belowStrategyId
	return r
}

// Arbitrary numeric value identifying the below order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOcoRequest) BelowStrategyType(belowStrategyType int32) ApiOrderListPlaceOcoRequest {
	r.belowStrategyType = &belowStrategyType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowPegPriceType(belowPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowPegPriceType = &belowPegPriceType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowPegOffsetType(belowPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowPegOffsetType = &belowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowPegOffsetValue(belowPegOffsetValue int32) ApiOrderListPlaceOcoRequest {
	r.belowPegOffsetValue = &belowPegOffsetValue
	return r
}

func (r ApiOrderListPlaceOcoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceOcoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOcoRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceOcoRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceOcoResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceOcoExecute(r)
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

func (r ApiOrderListPlaceOcoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceOcoResponse], chan error, error) {
	return r.ApiService.OrderListPlaceOcoExecute(r)
}

/*
OrderListPlaceOco WebSocket Place new Order list - OCO
/orderList.place.oco

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-order-list---oco-trade

@param symbol	@param side	@param quantity	@param aboveType	@param belowType	@param id Unique WebSocket request ID.	@param listClientOrderId	@param aboveClientOrderId Arbitrary unique ID among open orders for the above order. Automatically generated if not sent	@param aboveIcebergQty Note that this can only be used if `aboveTimeInForce` is `GTC`.	@param abovePrice Can be used if `aboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param aboveStopPrice Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. <br>Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified.	@param aboveTrailingDelta See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).	@param aboveTimeInForce	@param aboveStrategyId Arbitrary numeric value identifying the above order within an order strategy.	@param aboveStrategyType Arbitrary numeric value identifying the above order strategy. <br>Values smaller than 1000000 are reserved and cannot be used.	@param abovePegPriceType	@param abovePegOffsetType	@param abovePegOffsetValue	@param belowClientOrderId	@param belowIcebergQty Note that this can only be used if `belowTimeInForce` is `GTC`.	@param belowPrice Can be used if `belowType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param belowStopPrice Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT` or `TAKE_PROFIT_LIMIT`. <br>Either `belowStopPrice` or `belowTrailingDelta` or both, must be specified.	@param belowTrailingDelta See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).	@param belowTimeInForce	@param belowStrategyId Arbitrary numeric value identifying the below order within an order strategy.	@param belowStrategyType Arbitrary numeric value identifying the below order strategy. <br>Values smaller than 1000000 are reserved and cannot be used.	@param belowPegPriceType	@param belowPegOffsetType	@param belowPegOffsetValue	@param newOrderRespType	@param selfTradePreventionMode	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListPlaceOcoRequest
*/
func (a *TradeAPIService) OrderListPlaceOco() ApiOrderListPlaceOcoRequest {
	return ApiOrderListPlaceOcoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceOcoResponse
func (a *TradeAPIService) OrderListPlaceOcoExecute(r ApiOrderListPlaceOcoRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceOcoResponse], chan error, error) {
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

	if r.aboveType == nil {
		return nil, nil, common.ReportError("aboveType is required and must be specified")
	}
	localVarQueryParams["aboveType"] = *r.aboveType

	if r.belowType == nil {
		return nil, nil, common.ReportError("belowType is required and must be specified")
	}
	localVarQueryParams["belowType"] = *r.belowType

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.aboveClientOrderId != nil {
		localVarQueryParams["aboveClientOrderId"] = *r.aboveClientOrderId
	}
	if r.aboveIcebergQty != nil {
		localVarQueryParams["aboveIcebergQty"] = *r.aboveIcebergQty
	}
	if r.abovePrice != nil {
		localVarQueryParams["abovePrice"] = *r.abovePrice
	}
	if r.aboveStopPrice != nil {
		localVarQueryParams["aboveStopPrice"] = *r.aboveStopPrice
	}
	if r.aboveTrailingDelta != nil {
		localVarQueryParams["aboveTrailingDelta"] = *r.aboveTrailingDelta
	}
	if r.aboveTimeInForce != nil {
		localVarQueryParams["aboveTimeInForce"] = *r.aboveTimeInForce
	}
	if r.aboveStrategyId != nil {
		localVarQueryParams["aboveStrategyId"] = *r.aboveStrategyId
	}
	if r.aboveStrategyType != nil {
		localVarQueryParams["aboveStrategyType"] = *r.aboveStrategyType
	}
	if r.abovePegPriceType != nil {
		localVarQueryParams["abovePegPriceType"] = *r.abovePegPriceType
	}
	if r.abovePegOffsetType != nil {
		localVarQueryParams["abovePegOffsetType"] = *r.abovePegOffsetType
	}
	if r.abovePegOffsetValue != nil {
		localVarQueryParams["abovePegOffsetValue"] = *r.abovePegOffsetValue
	}
	if r.belowClientOrderId != nil {
		localVarQueryParams["belowClientOrderId"] = *r.belowClientOrderId
	}
	if r.belowIcebergQty != nil {
		localVarQueryParams["belowIcebergQty"] = *r.belowIcebergQty
	}
	if r.belowPrice != nil {
		localVarQueryParams["belowPrice"] = *r.belowPrice
	}
	if r.belowStopPrice != nil {
		localVarQueryParams["belowStopPrice"] = *r.belowStopPrice
	}
	if r.belowTrailingDelta != nil {
		localVarQueryParams["belowTrailingDelta"] = *r.belowTrailingDelta
	}
	if r.belowTimeInForce != nil {
		localVarQueryParams["belowTimeInForce"] = *r.belowTimeInForce
	}
	if r.belowStrategyId != nil {
		localVarQueryParams["belowStrategyId"] = *r.belowStrategyId
	}
	if r.belowStrategyType != nil {
		localVarQueryParams["belowStrategyType"] = *r.belowStrategyType
	}
	if r.belowPegPriceType != nil {
		localVarQueryParams["belowPegPriceType"] = *r.belowPegPriceType
	}
	if r.belowPegOffsetType != nil {
		localVarQueryParams["belowPegOffsetType"] = *r.belowPegOffsetType
	}
	if r.belowPegOffsetValue != nil {
		localVarQueryParams["belowPegOffsetValue"] = *r.belowPegOffsetValue
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
		"method": "/orderList.place.oco"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceOcoResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceOpoRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *models.OrderListPlaceOpoWorkingTypeParameter
	workingSide             *models.OrderCancelReplaceSideParameter
	workingPrice            *float32
	workingQuantity         *float32
	pendingType             *models.OrderListPlaceOpoPendingTypeParameter
	pendingSide             *models.OrderCancelReplaceSideParameter
	id                      *string
	listClientOrderId       *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float32
	workingTimeInForce      *models.OrderListPlaceStopLimitTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	workingPegOffsetType    *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float32
	pendingStopPrice        *float32
	pendingTrailingDelta    *float32
	pendingIcebergQty       *float32
	pendingTimeInForce      *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegPriceType     *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingPegOffsetType    *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float32
}

func (r ApiOrderListPlaceOpoRequest) Symbol(symbol string) ApiOrderListPlaceOpoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOpoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingPrice(workingPrice float32) ApiOrderListPlaceOpoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOpoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListPlaceOpoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingType(pendingType models.OrderListPlaceOpoPendingTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingType = &pendingType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceOpoRequest) Id(id string) ApiOrderListPlaceOpoRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceOpoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOpoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListPlaceOpoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOpoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceOpoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOpoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListPlaceOpoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListPlaceOpoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingTimeInForce(workingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOpoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOpoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOpoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOpoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOpoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOpoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOpoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListPlaceOpoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpoRequest) PendingClientOrderId(pendingClientOrderId string) ApiOrderListPlaceOpoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPrice(pendingPrice float32) ApiOrderListPlaceOpoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingStopPrice(pendingStopPrice float32) ApiOrderListPlaceOpoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingTrailingDelta(pendingTrailingDelta float32) ApiOrderListPlaceOpoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpoRequest) PendingIcebergQty(pendingIcebergQty float32) ApiOrderListPlaceOpoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingTimeInForce(pendingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOpoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListPlaceOpoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListPlaceOpoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOpoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListPlaceOpoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPegPriceType(pendingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListPlaceOpoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOpoRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceOpoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceOpoRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceOpoResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceOpoExecute(r)
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

func (r ApiOrderListPlaceOpoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceOpoResponse], chan error, error) {
	return r.ApiService.OrderListPlaceOpoExecute(r)
}

/*
OrderListPlaceOpo WebSocket OPO
/orderList.place.opo

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#opo-trade

@param symbol	@param workingType	@param workingSide	@param workingPrice	@param workingQuantity Sets the quantity for the working order. 	@param pendingType	@param pendingSide	@param id Unique WebSocket request ID.	@param listClientOrderId	@param newOrderRespType	@param selfTradePreventionMode	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. 	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. 	@param workingTimeInForce	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy. 	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param workingPegPriceType	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingClientOrderId Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. 	@param pendingPrice	@param pendingStopPrice	@param pendingTrailingDelta	@param pendingIcebergQty This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`. 	@param pendingTimeInForce	@param pendingStrategyId Arbitrary numeric value identifying the pending order within an order strategy. 	@param pendingStrategyType Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingPegPriceType	@param pendingPegOffsetType	@param pendingPegOffsetValue	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListPlaceOpoRequest
*/
func (a *TradeAPIService) OrderListPlaceOpo() ApiOrderListPlaceOpoRequest {
	return ApiOrderListPlaceOpoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceOpoResponse
func (a *TradeAPIService) OrderListPlaceOpoExecute(r ApiOrderListPlaceOpoRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceOpoResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.workingType == nil {
		return nil, nil, common.ReportError("workingType is required and must be specified")
	}
	localVarQueryParams["workingType"] = *r.workingType

	if r.workingSide == nil {
		return nil, nil, common.ReportError("workingSide is required and must be specified")
	}
	localVarQueryParams["workingSide"] = *r.workingSide

	if r.workingPrice == nil {
		return nil, nil, common.ReportError("workingPrice is required and must be specified")
	}
	localVarQueryParams["workingPrice"] = *r.workingPrice

	if r.workingQuantity == nil {
		return nil, nil, common.ReportError("workingQuantity is required and must be specified")
	}
	localVarQueryParams["workingQuantity"] = *r.workingQuantity

	if r.pendingType == nil {
		return nil, nil, common.ReportError("pendingType is required and must be specified")
	}
	localVarQueryParams["pendingType"] = *r.pendingType

	if r.pendingSide == nil {
		return nil, nil, common.ReportError("pendingSide is required and must be specified")
	}
	localVarQueryParams["pendingSide"] = *r.pendingSide

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.workingClientOrderId != nil {
		localVarQueryParams["workingClientOrderId"] = *r.workingClientOrderId
	}
	if r.workingIcebergQty != nil {
		localVarQueryParams["workingIcebergQty"] = *r.workingIcebergQty
	}
	if r.workingTimeInForce != nil {
		localVarQueryParams["workingTimeInForce"] = *r.workingTimeInForce
	}
	if r.workingStrategyId != nil {
		localVarQueryParams["workingStrategyId"] = *r.workingStrategyId
	}
	if r.workingStrategyType != nil {
		localVarQueryParams["workingStrategyType"] = *r.workingStrategyType
	}
	if r.workingPegPriceType != nil {
		localVarQueryParams["workingPegPriceType"] = *r.workingPegPriceType
	}
	if r.workingPegOffsetType != nil {
		localVarQueryParams["workingPegOffsetType"] = *r.workingPegOffsetType
	}
	if r.workingPegOffsetValue != nil {
		localVarQueryParams["workingPegOffsetValue"] = *r.workingPegOffsetValue
	}
	if r.pendingClientOrderId != nil {
		localVarQueryParams["pendingClientOrderId"] = *r.pendingClientOrderId
	}
	if r.pendingPrice != nil {
		localVarQueryParams["pendingPrice"] = *r.pendingPrice
	}
	if r.pendingStopPrice != nil {
		localVarQueryParams["pendingStopPrice"] = *r.pendingStopPrice
	}
	if r.pendingTrailingDelta != nil {
		localVarQueryParams["pendingTrailingDelta"] = *r.pendingTrailingDelta
	}
	if r.pendingIcebergQty != nil {
		localVarQueryParams["pendingIcebergQty"] = *r.pendingIcebergQty
	}
	if r.pendingTimeInForce != nil {
		localVarQueryParams["pendingTimeInForce"] = *r.pendingTimeInForce
	}
	if r.pendingStrategyId != nil {
		localVarQueryParams["pendingStrategyId"] = *r.pendingStrategyId
	}
	if r.pendingStrategyType != nil {
		localVarQueryParams["pendingStrategyType"] = *r.pendingStrategyType
	}
	if r.pendingPegPriceType != nil {
		localVarQueryParams["pendingPegPriceType"] = *r.pendingPegPriceType
	}
	if r.pendingPegOffsetType != nil {
		localVarQueryParams["pendingPegOffsetType"] = *r.pendingPegOffsetType
	}
	if r.pendingPegOffsetValue != nil {
		localVarQueryParams["pendingPegOffsetValue"] = *r.pendingPegOffsetValue
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.place.opo"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceOpoResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceOpocoRequest struct {
	ApiService                 *TradeAPIService
	symbol                     *string
	workingType                *models.OrderListPlaceOpoWorkingTypeParameter
	workingSide                *models.OrderCancelReplaceSideParameter
	workingPrice               *float32
	workingQuantity            *float32
	pendingSide                *models.OrderCancelReplaceSideParameter
	pendingAboveType           *models.OrderListPlaceOcoAboveTypeParameter
	id                         *string
	listClientOrderId          *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float32
	workingTimeInForce         *models.OrderListPlaceStopLimitTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	workingPegOffsetType       *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float32
	pendingAboveStopPrice      *float32
	pendingAboveTrailingDelta  *float32
	pendingAboveIcebergQty     *float32
	pendingAboveTimeInForce    *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListPlaceOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float32
	pendingBelowStopPrice      *float32
	pendingBelowTrailingDelta  *float32
	pendingBelowIcebergQty     *float32
	pendingBelowTimeInForce    *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float32
}

func (r ApiOrderListPlaceOpocoRequest) Symbol(symbol string) ApiOrderListPlaceOpocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingPrice(workingPrice float32) ApiOrderListPlaceOpocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOpocoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListPlaceOpocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingSide = &pendingSide
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingAboveType(pendingAboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceOpocoRequest) Id(id string) ApiOrderListPlaceOpocoRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceOpocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListPlaceOpocoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOpocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpocoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListPlaceOpocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOpocoRequest) PendingAbovePrice(pendingAbovePrice float32) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float32) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
func (r ApiOrderListPlaceOpocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float32) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float32) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowType(pendingBelowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price
func (r ApiOrderListPlaceOpocoRequest) PendingBelowPrice(pendingBelowPrice float32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOpocoRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceOpocoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceOpocoRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceOpocoResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceOpocoExecute(r)
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

func (r ApiOrderListPlaceOpocoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceOpocoResponse], chan error, error) {
	return r.ApiService.OrderListPlaceOpocoExecute(r)
}

/*
OrderListPlaceOpoco WebSocket OPOCO
/orderList.place.opoco

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#opoco-trade

@param symbol	@param workingType	@param workingSide	@param workingPrice	@param workingQuantity Sets the quantity for the working order. 	@param pendingSide	@param pendingAboveType	@param id Unique WebSocket request ID.	@param listClientOrderId	@param newOrderRespType	@param selfTradePreventionMode	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. 	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. 	@param workingTimeInForce	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy. 	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param workingPegPriceType	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingAboveClientOrderId Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. 	@param pendingAbovePrice Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. 	@param pendingAboveStopPrice Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT` 	@param pendingAboveTrailingDelta See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md) 	@param pendingAboveIcebergQty This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`. 	@param pendingAboveTimeInForce	@param pendingAboveStrategyId Arbitrary numeric value identifying the pending above order within an order strategy. 	@param pendingAboveStrategyType Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingAbovePegPriceType	@param pendingAbovePegOffsetType	@param pendingAbovePegOffsetValue	@param pendingBelowType	@param pendingBelowClientOrderId Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. 	@param pendingBelowPrice Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price 	@param pendingBelowStopPrice Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified. 	@param pendingBelowTrailingDelta	@param pendingBelowIcebergQty This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`. 	@param pendingBelowTimeInForce	@param pendingBelowStrategyId Arbitrary numeric value identifying the pending below order within an order strategy. 	@param pendingBelowStrategyType Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingBelowPegPriceType	@param pendingBelowPegOffsetType	@param pendingBelowPegOffsetValue	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListPlaceOpocoRequest
*/
func (a *TradeAPIService) OrderListPlaceOpoco() ApiOrderListPlaceOpocoRequest {
	return ApiOrderListPlaceOpocoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceOpocoResponse
func (a *TradeAPIService) OrderListPlaceOpocoExecute(r ApiOrderListPlaceOpocoRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceOpocoResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.workingType == nil {
		return nil, nil, common.ReportError("workingType is required and must be specified")
	}
	localVarQueryParams["workingType"] = *r.workingType

	if r.workingSide == nil {
		return nil, nil, common.ReportError("workingSide is required and must be specified")
	}
	localVarQueryParams["workingSide"] = *r.workingSide

	if r.workingPrice == nil {
		return nil, nil, common.ReportError("workingPrice is required and must be specified")
	}
	localVarQueryParams["workingPrice"] = *r.workingPrice

	if r.workingQuantity == nil {
		return nil, nil, common.ReportError("workingQuantity is required and must be specified")
	}
	localVarQueryParams["workingQuantity"] = *r.workingQuantity

	if r.pendingSide == nil {
		return nil, nil, common.ReportError("pendingSide is required and must be specified")
	}
	localVarQueryParams["pendingSide"] = *r.pendingSide

	if r.pendingAboveType == nil {
		return nil, nil, common.ReportError("pendingAboveType is required and must be specified")
	}
	localVarQueryParams["pendingAboveType"] = *r.pendingAboveType

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.workingClientOrderId != nil {
		localVarQueryParams["workingClientOrderId"] = *r.workingClientOrderId
	}
	if r.workingIcebergQty != nil {
		localVarQueryParams["workingIcebergQty"] = *r.workingIcebergQty
	}
	if r.workingTimeInForce != nil {
		localVarQueryParams["workingTimeInForce"] = *r.workingTimeInForce
	}
	if r.workingStrategyId != nil {
		localVarQueryParams["workingStrategyId"] = *r.workingStrategyId
	}
	if r.workingStrategyType != nil {
		localVarQueryParams["workingStrategyType"] = *r.workingStrategyType
	}
	if r.workingPegPriceType != nil {
		localVarQueryParams["workingPegPriceType"] = *r.workingPegPriceType
	}
	if r.workingPegOffsetType != nil {
		localVarQueryParams["workingPegOffsetType"] = *r.workingPegOffsetType
	}
	if r.workingPegOffsetValue != nil {
		localVarQueryParams["workingPegOffsetValue"] = *r.workingPegOffsetValue
	}
	if r.pendingAboveClientOrderId != nil {
		localVarQueryParams["pendingAboveClientOrderId"] = *r.pendingAboveClientOrderId
	}
	if r.pendingAbovePrice != nil {
		localVarQueryParams["pendingAbovePrice"] = *r.pendingAbovePrice
	}
	if r.pendingAboveStopPrice != nil {
		localVarQueryParams["pendingAboveStopPrice"] = *r.pendingAboveStopPrice
	}
	if r.pendingAboveTrailingDelta != nil {
		localVarQueryParams["pendingAboveTrailingDelta"] = *r.pendingAboveTrailingDelta
	}
	if r.pendingAboveIcebergQty != nil {
		localVarQueryParams["pendingAboveIcebergQty"] = *r.pendingAboveIcebergQty
	}
	if r.pendingAboveTimeInForce != nil {
		localVarQueryParams["pendingAboveTimeInForce"] = *r.pendingAboveTimeInForce
	}
	if r.pendingAboveStrategyId != nil {
		localVarQueryParams["pendingAboveStrategyId"] = *r.pendingAboveStrategyId
	}
	if r.pendingAboveStrategyType != nil {
		localVarQueryParams["pendingAboveStrategyType"] = *r.pendingAboveStrategyType
	}
	if r.pendingAbovePegPriceType != nil {
		localVarQueryParams["pendingAbovePegPriceType"] = *r.pendingAbovePegPriceType
	}
	if r.pendingAbovePegOffsetType != nil {
		localVarQueryParams["pendingAbovePegOffsetType"] = *r.pendingAbovePegOffsetType
	}
	if r.pendingAbovePegOffsetValue != nil {
		localVarQueryParams["pendingAbovePegOffsetValue"] = *r.pendingAbovePegOffsetValue
	}
	if r.pendingBelowType != nil {
		localVarQueryParams["pendingBelowType"] = *r.pendingBelowType
	}
	if r.pendingBelowClientOrderId != nil {
		localVarQueryParams["pendingBelowClientOrderId"] = *r.pendingBelowClientOrderId
	}
	if r.pendingBelowPrice != nil {
		localVarQueryParams["pendingBelowPrice"] = *r.pendingBelowPrice
	}
	if r.pendingBelowStopPrice != nil {
		localVarQueryParams["pendingBelowStopPrice"] = *r.pendingBelowStopPrice
	}
	if r.pendingBelowTrailingDelta != nil {
		localVarQueryParams["pendingBelowTrailingDelta"] = *r.pendingBelowTrailingDelta
	}
	if r.pendingBelowIcebergQty != nil {
		localVarQueryParams["pendingBelowIcebergQty"] = *r.pendingBelowIcebergQty
	}
	if r.pendingBelowTimeInForce != nil {
		localVarQueryParams["pendingBelowTimeInForce"] = *r.pendingBelowTimeInForce
	}
	if r.pendingBelowStrategyId != nil {
		localVarQueryParams["pendingBelowStrategyId"] = *r.pendingBelowStrategyId
	}
	if r.pendingBelowStrategyType != nil {
		localVarQueryParams["pendingBelowStrategyType"] = *r.pendingBelowStrategyType
	}
	if r.pendingBelowPegPriceType != nil {
		localVarQueryParams["pendingBelowPegPriceType"] = *r.pendingBelowPegPriceType
	}
	if r.pendingBelowPegOffsetType != nil {
		localVarQueryParams["pendingBelowPegOffsetType"] = *r.pendingBelowPegOffsetType
	}
	if r.pendingBelowPegOffsetValue != nil {
		localVarQueryParams["pendingBelowPegOffsetValue"] = *r.pendingBelowPegOffsetValue
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.place.opoco"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceOpocoResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceOtoRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *models.OrderListPlaceOpoWorkingTypeParameter
	workingSide             *models.OrderCancelReplaceSideParameter
	workingPrice            *float32
	workingQuantity         *float32
	pendingType             *models.OrderListPlaceOpoPendingTypeParameter
	pendingSide             *models.OrderCancelReplaceSideParameter
	pendingQuantity         *float32
	id                      *string
	listClientOrderId       *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float32
	workingTimeInForce      *models.OrderListPlaceStopLimitTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	workingPegOffsetType    *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float32
	pendingStopPrice        *float32
	pendingTrailingDelta    *float32
	pendingIcebergQty       *float32
	pendingTimeInForce      *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegOffsetType    *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingPegPriceType     *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float32
}

func (r ApiOrderListPlaceOtoRequest) Symbol(symbol string) ApiOrderListPlaceOtoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOtoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPrice(workingPrice float32) ApiOrderListPlaceOtoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOtoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListPlaceOtoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingType(pendingType models.OrderListPlaceOpoPendingTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingType = &pendingType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiOrderListPlaceOtoRequest) PendingQuantity(pendingQuantity float32) ApiOrderListPlaceOtoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceOtoRequest) Id(id string) ApiOrderListPlaceOtoRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceOtoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOtoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListPlaceOtoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOtoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceOtoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOtoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListPlaceOtoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListPlaceOtoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingTimeInForce(workingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOtoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOtoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOtoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOtoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOtoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOtoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOtoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListPlaceOtoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtoRequest) PendingClientOrderId(pendingClientOrderId string) ApiOrderListPlaceOtoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPrice(pendingPrice float32) ApiOrderListPlaceOtoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingStopPrice(pendingStopPrice float32) ApiOrderListPlaceOtoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingTrailingDelta(pendingTrailingDelta float32) ApiOrderListPlaceOtoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtoRequest) PendingIcebergQty(pendingIcebergQty float32) ApiOrderListPlaceOtoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingTimeInForce(pendingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOtoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListPlaceOtoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListPlaceOtoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOtoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListPlaceOtoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPegPriceType(pendingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListPlaceOtoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOtoRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceOtoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceOtoRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceOtoResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceOtoExecute(r)
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

func (r ApiOrderListPlaceOtoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceOtoResponse], chan error, error) {
	return r.ApiService.OrderListPlaceOtoExecute(r)
}

/*
OrderListPlaceOto WebSocket Place new Order list - OTO
/orderList.place.oto

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-order-list---oto-trade

@param symbol	@param workingType	@param workingSide	@param workingPrice	@param workingQuantity Sets the quantity for the working order. 	@param pendingType	@param pendingSide	@param pendingQuantity Sets the quantity for the pending order.	@param id Unique WebSocket request ID.	@param listClientOrderId	@param newOrderRespType	@param selfTradePreventionMode	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. 	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. 	@param workingTimeInForce	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy. 	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param workingPegPriceType	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingClientOrderId Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent. 	@param pendingPrice	@param pendingStopPrice	@param pendingTrailingDelta	@param pendingIcebergQty This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`. 	@param pendingTimeInForce	@param pendingStrategyId Arbitrary numeric value identifying the pending order within an order strategy. 	@param pendingStrategyType Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingPegOffsetType	@param pendingPegPriceType	@param pendingPegOffsetValue	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListPlaceOtoRequest
*/
func (a *TradeAPIService) OrderListPlaceOto() ApiOrderListPlaceOtoRequest {
	return ApiOrderListPlaceOtoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceOtoResponse
func (a *TradeAPIService) OrderListPlaceOtoExecute(r ApiOrderListPlaceOtoRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceOtoResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.workingType == nil {
		return nil, nil, common.ReportError("workingType is required and must be specified")
	}
	localVarQueryParams["workingType"] = *r.workingType

	if r.workingSide == nil {
		return nil, nil, common.ReportError("workingSide is required and must be specified")
	}
	localVarQueryParams["workingSide"] = *r.workingSide

	if r.workingPrice == nil {
		return nil, nil, common.ReportError("workingPrice is required and must be specified")
	}
	localVarQueryParams["workingPrice"] = *r.workingPrice

	if r.workingQuantity == nil {
		return nil, nil, common.ReportError("workingQuantity is required and must be specified")
	}
	localVarQueryParams["workingQuantity"] = *r.workingQuantity

	if r.pendingType == nil {
		return nil, nil, common.ReportError("pendingType is required and must be specified")
	}
	localVarQueryParams["pendingType"] = *r.pendingType

	if r.pendingSide == nil {
		return nil, nil, common.ReportError("pendingSide is required and must be specified")
	}
	localVarQueryParams["pendingSide"] = *r.pendingSide

	if r.pendingQuantity == nil {
		return nil, nil, common.ReportError("pendingQuantity is required and must be specified")
	}
	localVarQueryParams["pendingQuantity"] = *r.pendingQuantity

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.workingClientOrderId != nil {
		localVarQueryParams["workingClientOrderId"] = *r.workingClientOrderId
	}
	if r.workingIcebergQty != nil {
		localVarQueryParams["workingIcebergQty"] = *r.workingIcebergQty
	}
	if r.workingTimeInForce != nil {
		localVarQueryParams["workingTimeInForce"] = *r.workingTimeInForce
	}
	if r.workingStrategyId != nil {
		localVarQueryParams["workingStrategyId"] = *r.workingStrategyId
	}
	if r.workingStrategyType != nil {
		localVarQueryParams["workingStrategyType"] = *r.workingStrategyType
	}
	if r.workingPegPriceType != nil {
		localVarQueryParams["workingPegPriceType"] = *r.workingPegPriceType
	}
	if r.workingPegOffsetType != nil {
		localVarQueryParams["workingPegOffsetType"] = *r.workingPegOffsetType
	}
	if r.workingPegOffsetValue != nil {
		localVarQueryParams["workingPegOffsetValue"] = *r.workingPegOffsetValue
	}
	if r.pendingClientOrderId != nil {
		localVarQueryParams["pendingClientOrderId"] = *r.pendingClientOrderId
	}
	if r.pendingPrice != nil {
		localVarQueryParams["pendingPrice"] = *r.pendingPrice
	}
	if r.pendingStopPrice != nil {
		localVarQueryParams["pendingStopPrice"] = *r.pendingStopPrice
	}
	if r.pendingTrailingDelta != nil {
		localVarQueryParams["pendingTrailingDelta"] = *r.pendingTrailingDelta
	}
	if r.pendingIcebergQty != nil {
		localVarQueryParams["pendingIcebergQty"] = *r.pendingIcebergQty
	}
	if r.pendingTimeInForce != nil {
		localVarQueryParams["pendingTimeInForce"] = *r.pendingTimeInForce
	}
	if r.pendingStrategyId != nil {
		localVarQueryParams["pendingStrategyId"] = *r.pendingStrategyId
	}
	if r.pendingStrategyType != nil {
		localVarQueryParams["pendingStrategyType"] = *r.pendingStrategyType
	}
	if r.pendingPegOffsetType != nil {
		localVarQueryParams["pendingPegOffsetType"] = *r.pendingPegOffsetType
	}
	if r.pendingPegPriceType != nil {
		localVarQueryParams["pendingPegPriceType"] = *r.pendingPegPriceType
	}
	if r.pendingPegOffsetValue != nil {
		localVarQueryParams["pendingPegOffsetValue"] = *r.pendingPegOffsetValue
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.place.oto"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceOtoResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderListPlaceOtocoRequest struct {
	ApiService                 *TradeAPIService
	symbol                     *string
	workingType                *models.OrderListPlaceOpoWorkingTypeParameter
	workingSide                *models.OrderCancelReplaceSideParameter
	workingPrice               *float32
	workingQuantity            *float32
	pendingSide                *models.OrderCancelReplaceSideParameter
	pendingQuantity            *float32
	pendingAboveType           *models.OrderListPlaceOcoAboveTypeParameter
	id                         *string
	listClientOrderId          *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float32
	workingTimeInForce         *models.OrderListPlaceStopLimitTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	workingPegOffsetType       *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float32
	pendingAboveStopPrice      *float32
	pendingAboveTrailingDelta  *float32
	pendingAboveIcebergQty     *float32
	pendingAboveTimeInForce    *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListPlaceOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float32
	pendingBelowStopPrice      *float32
	pendingBelowTrailingDelta  *float32
	pendingBelowIcebergQty     *float32
	pendingBelowTimeInForce    *models.OrderListPlaceStopLimitTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderListPlaceOcoAbovePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderListPlaceOcoAbovePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float32
}

func (r ApiOrderListPlaceOtocoRequest) Symbol(symbol string) ApiOrderListPlaceOtocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPrice(workingPrice float32) ApiOrderListPlaceOtocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOtocoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListPlaceOtocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiOrderListPlaceOtocoRequest) PendingQuantity(pendingQuantity float32) ApiOrderListPlaceOtocoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAboveType(pendingAboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderListPlaceOtocoRequest) Id(id string) ApiOrderListPlaceOtocoRequest {
	r.id = &id
	return r
}

func (r ApiOrderListPlaceOtocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListPlaceOtocoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOtocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtocoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListPlaceOtocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListPlaceOtocoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOtocoRequest) PendingAbovePrice(pendingAbovePrice float32) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float32) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
func (r ApiOrderListPlaceOtocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float32) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float32) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowType(pendingBelowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price
func (r ApiOrderListPlaceOtocoRequest) PendingBelowPrice(pendingBelowPrice float32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderListPlaceStopLimitTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderListPlaceOcoAbovePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderListPlaceOcoAbovePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOtocoRequest) RecvWindow(recvWindow float32) ApiOrderListPlaceOtocoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListPlaceOtocoRequest) Execute() (*common.ResponseOrRaw[models.OrderListPlaceOtocoResponse], error) {
	respChan, errChan, err := r.ApiService.OrderListPlaceOtocoExecute(r)
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

func (r ApiOrderListPlaceOtocoRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderListPlaceOtocoResponse], chan error, error) {
	return r.ApiService.OrderListPlaceOtocoExecute(r)
}

/*
OrderListPlaceOtoco WebSocket Place new Order list - OTOCO
/orderList.place.otoco

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-order-list---otoco-trade

@param symbol	@param workingType	@param workingSide	@param workingPrice	@param workingQuantity Sets the quantity for the working order. 	@param pendingSide	@param pendingQuantity Sets the quantity for the pending order.	@param pendingAboveType	@param id Unique WebSocket request ID.	@param listClientOrderId	@param newOrderRespType	@param selfTradePreventionMode	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent. 	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`. 	@param workingTimeInForce	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy. 	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param workingPegPriceType	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingAboveClientOrderId Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent. 	@param pendingAbovePrice Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price. 	@param pendingAboveStopPrice Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT` 	@param pendingAboveTrailingDelta See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md) 	@param pendingAboveIcebergQty This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`. 	@param pendingAboveTimeInForce	@param pendingAboveStrategyId Arbitrary numeric value identifying the pending above order within an order strategy. 	@param pendingAboveStrategyType Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingAbovePegPriceType	@param pendingAbovePegOffsetType	@param pendingAbovePegOffsetValue	@param pendingBelowType	@param pendingBelowClientOrderId Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent. 	@param pendingBelowPrice Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price 	@param pendingBelowStopPrice Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified. 	@param pendingBelowTrailingDelta	@param pendingBelowIcebergQty This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`. 	@param pendingBelowTimeInForce	@param pendingBelowStrategyId Arbitrary numeric value identifying the pending below order within an order strategy. 	@param pendingBelowStrategyType Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used. 	@param pendingBelowPegPriceType	@param pendingBelowPegOffsetType	@param pendingBelowPegOffsetValue	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListPlaceOtocoRequest
*/
func (a *TradeAPIService) OrderListPlaceOtoco() ApiOrderListPlaceOtocoRequest {
	return ApiOrderListPlaceOtocoRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderListPlaceOtocoResponse
func (a *TradeAPIService) OrderListPlaceOtocoExecute(r ApiOrderListPlaceOtocoRequest) (chan *common.ResponseOrRaw[models.OrderListPlaceOtocoResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.workingType == nil {
		return nil, nil, common.ReportError("workingType is required and must be specified")
	}
	localVarQueryParams["workingType"] = *r.workingType

	if r.workingSide == nil {
		return nil, nil, common.ReportError("workingSide is required and must be specified")
	}
	localVarQueryParams["workingSide"] = *r.workingSide

	if r.workingPrice == nil {
		return nil, nil, common.ReportError("workingPrice is required and must be specified")
	}
	localVarQueryParams["workingPrice"] = *r.workingPrice

	if r.workingQuantity == nil {
		return nil, nil, common.ReportError("workingQuantity is required and must be specified")
	}
	localVarQueryParams["workingQuantity"] = *r.workingQuantity

	if r.pendingSide == nil {
		return nil, nil, common.ReportError("pendingSide is required and must be specified")
	}
	localVarQueryParams["pendingSide"] = *r.pendingSide

	if r.pendingQuantity == nil {
		return nil, nil, common.ReportError("pendingQuantity is required and must be specified")
	}
	localVarQueryParams["pendingQuantity"] = *r.pendingQuantity

	if r.pendingAboveType == nil {
		return nil, nil, common.ReportError("pendingAboveType is required and must be specified")
	}
	localVarQueryParams["pendingAboveType"] = *r.pendingAboveType

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.listClientOrderId != nil {
		localVarQueryParams["listClientOrderId"] = *r.listClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.workingClientOrderId != nil {
		localVarQueryParams["workingClientOrderId"] = *r.workingClientOrderId
	}
	if r.workingIcebergQty != nil {
		localVarQueryParams["workingIcebergQty"] = *r.workingIcebergQty
	}
	if r.workingTimeInForce != nil {
		localVarQueryParams["workingTimeInForce"] = *r.workingTimeInForce
	}
	if r.workingStrategyId != nil {
		localVarQueryParams["workingStrategyId"] = *r.workingStrategyId
	}
	if r.workingStrategyType != nil {
		localVarQueryParams["workingStrategyType"] = *r.workingStrategyType
	}
	if r.workingPegPriceType != nil {
		localVarQueryParams["workingPegPriceType"] = *r.workingPegPriceType
	}
	if r.workingPegOffsetType != nil {
		localVarQueryParams["workingPegOffsetType"] = *r.workingPegOffsetType
	}
	if r.workingPegOffsetValue != nil {
		localVarQueryParams["workingPegOffsetValue"] = *r.workingPegOffsetValue
	}
	if r.pendingAboveClientOrderId != nil {
		localVarQueryParams["pendingAboveClientOrderId"] = *r.pendingAboveClientOrderId
	}
	if r.pendingAbovePrice != nil {
		localVarQueryParams["pendingAbovePrice"] = *r.pendingAbovePrice
	}
	if r.pendingAboveStopPrice != nil {
		localVarQueryParams["pendingAboveStopPrice"] = *r.pendingAboveStopPrice
	}
	if r.pendingAboveTrailingDelta != nil {
		localVarQueryParams["pendingAboveTrailingDelta"] = *r.pendingAboveTrailingDelta
	}
	if r.pendingAboveIcebergQty != nil {
		localVarQueryParams["pendingAboveIcebergQty"] = *r.pendingAboveIcebergQty
	}
	if r.pendingAboveTimeInForce != nil {
		localVarQueryParams["pendingAboveTimeInForce"] = *r.pendingAboveTimeInForce
	}
	if r.pendingAboveStrategyId != nil {
		localVarQueryParams["pendingAboveStrategyId"] = *r.pendingAboveStrategyId
	}
	if r.pendingAboveStrategyType != nil {
		localVarQueryParams["pendingAboveStrategyType"] = *r.pendingAboveStrategyType
	}
	if r.pendingAbovePegPriceType != nil {
		localVarQueryParams["pendingAbovePegPriceType"] = *r.pendingAbovePegPriceType
	}
	if r.pendingAbovePegOffsetType != nil {
		localVarQueryParams["pendingAbovePegOffsetType"] = *r.pendingAbovePegOffsetType
	}
	if r.pendingAbovePegOffsetValue != nil {
		localVarQueryParams["pendingAbovePegOffsetValue"] = *r.pendingAbovePegOffsetValue
	}
	if r.pendingBelowType != nil {
		localVarQueryParams["pendingBelowType"] = *r.pendingBelowType
	}
	if r.pendingBelowClientOrderId != nil {
		localVarQueryParams["pendingBelowClientOrderId"] = *r.pendingBelowClientOrderId
	}
	if r.pendingBelowPrice != nil {
		localVarQueryParams["pendingBelowPrice"] = *r.pendingBelowPrice
	}
	if r.pendingBelowStopPrice != nil {
		localVarQueryParams["pendingBelowStopPrice"] = *r.pendingBelowStopPrice
	}
	if r.pendingBelowTrailingDelta != nil {
		localVarQueryParams["pendingBelowTrailingDelta"] = *r.pendingBelowTrailingDelta
	}
	if r.pendingBelowIcebergQty != nil {
		localVarQueryParams["pendingBelowIcebergQty"] = *r.pendingBelowIcebergQty
	}
	if r.pendingBelowTimeInForce != nil {
		localVarQueryParams["pendingBelowTimeInForce"] = *r.pendingBelowTimeInForce
	}
	if r.pendingBelowStrategyId != nil {
		localVarQueryParams["pendingBelowStrategyId"] = *r.pendingBelowStrategyId
	}
	if r.pendingBelowStrategyType != nil {
		localVarQueryParams["pendingBelowStrategyType"] = *r.pendingBelowStrategyType
	}
	if r.pendingBelowPegPriceType != nil {
		localVarQueryParams["pendingBelowPegPriceType"] = *r.pendingBelowPegPriceType
	}
	if r.pendingBelowPegOffsetType != nil {
		localVarQueryParams["pendingBelowPegOffsetType"] = *r.pendingBelowPegOffsetType
	}
	if r.pendingBelowPegOffsetValue != nil {
		localVarQueryParams["pendingBelowPegOffsetValue"] = *r.pendingBelowPegOffsetValue
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/orderList.place.otoco"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderListPlaceOtocoResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderPlaceRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	type_                   *models.OrderCancelReplaceTypeParameter
	id                      *string
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float32
	quantity                *float32
	quoteOrderQty           *float32
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice               *float32
	trailingDelta           *int32
	icebergQty              *float32
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	pegPriceType            *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow              *float32
}

func (r ApiOrderPlaceRequest) Symbol(symbol string) ApiOrderPlaceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderPlaceRequest {
	r.side = &side
	return r
}

func (r ApiOrderPlaceRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderPlaceRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderPlaceRequest) Id(id string) ApiOrderPlaceRequest {
	r.id = &id
	return r
}

func (r ApiOrderPlaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderPlaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderPlaceRequest) Price(price float32) ApiOrderPlaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderPlaceRequest) Quantity(quantity float32) ApiOrderPlaceRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderPlaceRequest) QuoteOrderQty(quoteOrderQty float32) ApiOrderPlaceRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderPlaceRequest) NewClientOrderId(newClientOrderId string) ApiOrderPlaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiOrderPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderPlaceRequest) StopPrice(stopPrice float32) ApiOrderPlaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)
func (r ApiOrderPlaceRequest) TrailingDelta(trailingDelta int32) ApiOrderPlaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

func (r ApiOrderPlaceRequest) IcebergQty(icebergQty float32) ApiOrderPlaceRequest {
	r.icebergQty = &icebergQty
	return r
}

// Arbitrary numeric value identifying the order within an order strategy.
func (r ApiOrderPlaceRequest) StrategyId(strategyId int64) ApiOrderPlaceRequest {
	r.strategyId = &strategyId
	return r
}

// Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderPlaceRequest) StrategyType(strategyType int32) ApiOrderPlaceRequest {
	r.strategyType = &strategyType
	return r
}

func (r ApiOrderPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiOrderPlaceRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderPlaceRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level to peg the price to (max: 100)       See Pegged Orders
func (r ApiOrderPlaceRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderPlaceRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiOrderPlaceRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderPlaceRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderPlaceRequest) RecvWindow(recvWindow float32) ApiOrderPlaceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderPlaceRequest) Execute() (*common.ResponseOrRaw[models.OrderPlaceResponse], error) {
	respChan, errChan, err := r.ApiService.OrderPlaceExecute(r)
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

func (r ApiOrderPlaceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderPlaceResponse], chan error, error) {
	return r.ApiService.OrderPlaceExecute(r)
}

/*
OrderPlace WebSocket Place new order
/order.place

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-order-trade

@param symbol	@param side	@param type_	@param id Unique WebSocket request ID.	@param timeInForce	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param newOrderRespType	@param stopPrice	@param trailingDelta See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)	@param icebergQty	@param strategyId Arbitrary numeric value identifying the order within an order strategy.	@param strategyType Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.	@param selfTradePreventionMode	@param pegPriceType	@param pegOffsetValue Price level to peg the price to (max: 100)       See Pegged Orders	@param pegOffsetType	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderPlaceRequest
*/
func (a *TradeAPIService) OrderPlace() ApiOrderPlaceRequest {
	return ApiOrderPlaceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderPlaceResponse
func (a *TradeAPIService) OrderPlaceExecute(r ApiOrderPlaceRequest) (chan *common.ResponseOrRaw[models.OrderPlaceResponse], chan error, error) {
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
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.quoteOrderQty != nil {
		localVarQueryParams["quoteOrderQty"] = *r.quoteOrderQty
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.trailingDelta != nil {
		localVarQueryParams["trailingDelta"] = *r.trailingDelta
	}
	if r.icebergQty != nil {
		localVarQueryParams["icebergQty"] = *r.icebergQty
	}
	if r.strategyId != nil {
		localVarQueryParams["strategyId"] = *r.strategyId
	}
	if r.strategyType != nil {
		localVarQueryParams["strategyType"] = *r.strategyType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.pegPriceType != nil {
		localVarQueryParams["pegPriceType"] = *r.pegPriceType
	}
	if r.pegOffsetValue != nil {
		localVarQueryParams["pegOffsetValue"] = *r.pegOffsetValue
	}
	if r.pegOffsetType != nil {
		localVarQueryParams["pegOffsetType"] = *r.pegOffsetType
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

	return SendMessage[models.OrderPlaceResponse](a.Ws, localPayload, sendParams)
}

type ApiOrderTestRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	type_                   *models.OrderCancelReplaceTypeParameter
	id                      *string
	computeCommissionRates  *bool
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float32
	quantity                *float32
	quoteOrderQty           *float32
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice               *float32
	trailingDelta           *int32
	icebergQty              *float32
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	pegPriceType            *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow              *float32
}

func (r ApiOrderTestRequest) Symbol(symbol string) ApiOrderTestRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderTestRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderTestRequest {
	r.side = &side
	return r
}

func (r ApiOrderTestRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderTestRequest {
	r.type_ = &type_
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderTestRequest) Id(id string) ApiOrderTestRequest {
	r.id = &id
	return r
}

// Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
func (r ApiOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

func (r ApiOrderTestRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderTestRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderTestRequest) Price(price float32) ApiOrderTestRequest {
	r.price = &price
	return r
}

func (r ApiOrderTestRequest) Quantity(quantity float32) ApiOrderTestRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderTestRequest) QuoteOrderQty(quoteOrderQty float32) ApiOrderTestRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiOrderTestRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderTestRequest) StopPrice(stopPrice float32) ApiOrderTestRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)
func (r ApiOrderTestRequest) TrailingDelta(trailingDelta int32) ApiOrderTestRequest {
	r.trailingDelta = &trailingDelta
	return r
}

func (r ApiOrderTestRequest) IcebergQty(icebergQty float32) ApiOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

// Arbitrary numeric value identifying the order within an order strategy.
func (r ApiOrderTestRequest) StrategyId(strategyId int64) ApiOrderTestRequest {
	r.strategyId = &strategyId
	return r
}

// Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderTestRequest) StrategyType(strategyType int32) ApiOrderTestRequest {
	r.strategyType = &strategyType
	return r
}

func (r ApiOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiOrderTestRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderTestRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level to peg the price to (max: 100)       See Pegged Orders
func (r ApiOrderTestRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderTestRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiOrderTestRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderTestRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderTestRequest) RecvWindow(recvWindow float32) ApiOrderTestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderTestRequest) Execute() (*common.ResponseOrRaw[models.OrderTestResponse], error) {
	respChan, errChan, err := r.ApiService.OrderTestExecute(r)
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

func (r ApiOrderTestRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderTestResponse], chan error, error) {
	return r.ApiService.OrderTestExecute(r)
}

/*
OrderTest WebSocket Test new order
/order.test

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#test-new-order-trade

@param symbol	@param side	@param type_	@param id Unique WebSocket request ID.	@param computeCommissionRates Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.	@param timeInForce	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param newOrderRespType	@param stopPrice	@param trailingDelta See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md)	@param icebergQty	@param strategyId Arbitrary numeric value identifying the order within an order strategy.	@param strategyType Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.	@param selfTradePreventionMode	@param pegPriceType	@param pegOffsetValue Price level to peg the price to (max: 100)       See Pegged Orders	@param pegOffsetType	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderTestRequest
*/
func (a *TradeAPIService) OrderTest() ApiOrderTestRequest {
	return ApiOrderTestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderTestResponse
func (a *TradeAPIService) OrderTestExecute(r ApiOrderTestRequest) (chan *common.ResponseOrRaw[models.OrderTestResponse], chan error, error) {
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
	if r.computeCommissionRates != nil {
		localVarQueryParams["computeCommissionRates"] = *r.computeCommissionRates
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
	}
	if r.price != nil {
		localVarQueryParams["price"] = *r.price
	}
	if r.quantity != nil {
		localVarQueryParams["quantity"] = *r.quantity
	}
	if r.quoteOrderQty != nil {
		localVarQueryParams["quoteOrderQty"] = *r.quoteOrderQty
	}
	if r.newClientOrderId != nil {
		localVarQueryParams["newClientOrderId"] = *r.newClientOrderId
	}
	if r.newOrderRespType != nil {
		localVarQueryParams["newOrderRespType"] = *r.newOrderRespType
	}
	if r.stopPrice != nil {
		localVarQueryParams["stopPrice"] = *r.stopPrice
	}
	if r.trailingDelta != nil {
		localVarQueryParams["trailingDelta"] = *r.trailingDelta
	}
	if r.icebergQty != nil {
		localVarQueryParams["icebergQty"] = *r.icebergQty
	}
	if r.strategyId != nil {
		localVarQueryParams["strategyId"] = *r.strategyId
	}
	if r.strategyType != nil {
		localVarQueryParams["strategyType"] = *r.strategyType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.pegPriceType != nil {
		localVarQueryParams["pegPriceType"] = *r.pegPriceType
	}
	if r.pegOffsetValue != nil {
		localVarQueryParams["pegOffsetValue"] = *r.pegOffsetValue
	}
	if r.pegOffsetType != nil {
		localVarQueryParams["pegOffsetType"] = *r.pegOffsetType
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/order.test"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderTestResponse](a.Ws, localPayload, sendParams)
}

type ApiSorOrderPlaceRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	type_                   *models.OrderCancelReplaceTypeParameter
	quantity                *float32
	id                      *string
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float32
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	icebergQty              *float32
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiSorOrderPlaceRequest) Symbol(symbol string) ApiSorOrderPlaceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSorOrderPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiSorOrderPlaceRequest {
	r.side = &side
	return r
}

func (r ApiSorOrderPlaceRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiSorOrderPlaceRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderPlaceRequest) Quantity(quantity float32) ApiSorOrderPlaceRequest {
	r.quantity = &quantity
	return r
}

// Unique WebSocket request ID.
func (r ApiSorOrderPlaceRequest) Id(id string) ApiSorOrderPlaceRequest {
	r.id = &id
	return r
}

func (r ApiSorOrderPlaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiSorOrderPlaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderPlaceRequest) Price(price float32) ApiSorOrderPlaceRequest {
	r.price = &price
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiSorOrderPlaceRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderPlaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiSorOrderPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiSorOrderPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiSorOrderPlaceRequest) IcebergQty(icebergQty float32) ApiSorOrderPlaceRequest {
	r.icebergQty = &icebergQty
	return r
}

// Arbitrary numeric value identifying the order within an order strategy.
func (r ApiSorOrderPlaceRequest) StrategyId(strategyId int64) ApiSorOrderPlaceRequest {
	r.strategyId = &strategyId
	return r
}

// Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.
func (r ApiSorOrderPlaceRequest) StrategyType(strategyType int32) ApiSorOrderPlaceRequest {
	r.strategyType = &strategyType
	return r
}

func (r ApiSorOrderPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiSorOrderPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderPlaceRequest) RecvWindow(recvWindow float32) ApiSorOrderPlaceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSorOrderPlaceRequest) Execute() (*common.ResponseOrRaw[models.SorOrderPlaceResponse], error) {
	respChan, errChan, err := r.ApiService.SorOrderPlaceExecute(r)
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

func (r ApiSorOrderPlaceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SorOrderPlaceResponse], chan error, error) {
	return r.ApiService.SorOrderPlaceExecute(r)
}

/*
SorOrderPlace WebSocket Place new order using SOR
/sor.order.place

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#place-new-order-using-sor-trade

@param symbol	@param side	@param type_	@param quantity	@param id Unique WebSocket request ID.	@param timeInForce	@param price	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param newOrderRespType	@param icebergQty	@param strategyId Arbitrary numeric value identifying the order within an order strategy.	@param strategyType Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.	@param selfTradePreventionMode	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSorOrderPlaceRequest
*/
func (a *TradeAPIService) SorOrderPlace() ApiSorOrderPlaceRequest {
	return ApiSorOrderPlaceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SorOrderPlaceResponse
func (a *TradeAPIService) SorOrderPlaceExecute(r ApiSorOrderPlaceRequest) (chan *common.ResponseOrRaw[models.SorOrderPlaceResponse], chan error, error) {
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

	if r.quantity == nil {
		return nil, nil, common.ReportError("quantity is required and must be specified")
	}
	localVarQueryParams["quantity"] = *r.quantity

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
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
	if r.icebergQty != nil {
		localVarQueryParams["icebergQty"] = *r.icebergQty
	}
	if r.strategyId != nil {
		localVarQueryParams["strategyId"] = *r.strategyId
	}
	if r.strategyType != nil {
		localVarQueryParams["strategyType"] = *r.strategyType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/sor.order.place"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SorOrderPlaceResponse](a.Ws, localPayload, sendParams)
}

type ApiSorOrderTestRequest struct {
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.OrderCancelReplaceSideParameter
	type_                   *models.OrderCancelReplaceTypeParameter
	quantity                *float32
	id                      *string
	computeCommissionRates  *bool
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float32
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	icebergQty              *float32
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiSorOrderTestRequest) Symbol(symbol string) ApiSorOrderTestRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSorOrderTestRequest) Side(side models.OrderCancelReplaceSideParameter) ApiSorOrderTestRequest {
	r.side = &side
	return r
}

func (r ApiSorOrderTestRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiSorOrderTestRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderTestRequest) Quantity(quantity float32) ApiSorOrderTestRequest {
	r.quantity = &quantity
	return r
}

// Unique WebSocket request ID.
func (r ApiSorOrderTestRequest) Id(id string) ApiSorOrderTestRequest {
	r.id = &id
	return r
}

// Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
func (r ApiSorOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiSorOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

func (r ApiSorOrderTestRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiSorOrderTestRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderTestRequest) Price(price float32) ApiSorOrderTestRequest {
	r.price = &price
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiSorOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiSorOrderTestRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiSorOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiSorOrderTestRequest) IcebergQty(icebergQty float32) ApiSorOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

// Arbitrary numeric value identifying the order within an order strategy.
func (r ApiSorOrderTestRequest) StrategyId(strategyId int64) ApiSorOrderTestRequest {
	r.strategyId = &strategyId
	return r
}

// Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.
func (r ApiSorOrderTestRequest) StrategyType(strategyType int32) ApiSorOrderTestRequest {
	r.strategyType = &strategyType
	return r
}

func (r ApiSorOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiSorOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderTestRequest) RecvWindow(recvWindow float32) ApiSorOrderTestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSorOrderTestRequest) Execute() (*common.ResponseOrRaw[models.SorOrderTestResponse], error) {
	respChan, errChan, err := r.ApiService.SorOrderTestExecute(r)
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

func (r ApiSorOrderTestRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SorOrderTestResponse], chan error, error) {
	return r.ApiService.SorOrderTestExecute(r)
}

/*
SorOrderTest WebSocket Test new order using SOR
/sor.order.test

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/trading-requests#test-new-order-using-sor-trade

@param symbol	@param side	@param type_	@param quantity	@param id Unique WebSocket request ID.	@param computeCommissionRates Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.	@param timeInForce	@param price	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`. 	@param newOrderRespType	@param icebergQty	@param strategyId Arbitrary numeric value identifying the order within an order strategy.	@param strategyType Arbitrary numeric value identifying the order strategy.                 Values smaller than 1000000 are reserved and cannot be used.	@param selfTradePreventionMode	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSorOrderTestRequest
*/
func (a *TradeAPIService) SorOrderTest() ApiSorOrderTestRequest {
	return ApiSorOrderTestRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SorOrderTestResponse
func (a *TradeAPIService) SorOrderTestExecute(r ApiSorOrderTestRequest) (chan *common.ResponseOrRaw[models.SorOrderTestResponse], chan error, error) {
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

	if r.quantity == nil {
		return nil, nil, common.ReportError("quantity is required and must be specified")
	}
	localVarQueryParams["quantity"] = *r.quantity

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.computeCommissionRates != nil {
		localVarQueryParams["computeCommissionRates"] = *r.computeCommissionRates
	}
	if r.timeInForce != nil {
		localVarQueryParams["timeInForce"] = *r.timeInForce
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
	if r.icebergQty != nil {
		localVarQueryParams["icebergQty"] = *r.icebergQty
	}
	if r.strategyId != nil {
		localVarQueryParams["strategyId"] = *r.strategyId
	}
	if r.strategyType != nil {
		localVarQueryParams["strategyType"] = *r.strategyType
	}
	if r.selfTradePreventionMode != nil {
		localVarQueryParams["selfTradePreventionMode"] = *r.selfTradePreventionMode
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/sor.order.test"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SorOrderTestResponse](a.Ws, localPayload, sendParams)
}
