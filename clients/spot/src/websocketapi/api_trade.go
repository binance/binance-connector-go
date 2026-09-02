/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiOpenOrdersCancelAllRequest struct {
	ApiService *TradeAPIService
	symbol     *string
	id         *string
	recvWindow *float64
}

func (r ApiOpenOrdersCancelAllRequest) Symbol(symbol string) ApiOpenOrdersCancelAllRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiOpenOrdersCancelAllRequest) Id(id string) ApiOpenOrdersCancelAllRequest {
	r.id = &id
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOpenOrdersCancelAllRequest) RecvWindow(recvWindow float64) ApiOpenOrdersCancelAllRequest {
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
OpenOrdersCancelAll Cancel open orders (TRADE)
/openOrders.cancelAll

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#open-orders-cancel-all

@param symbol	@param id Client-generated request identifier.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	newQty            *float64
	id                *string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *float64
}

func (r ApiOrderAmendKeepPriorityRequest) Symbol(symbol string) ApiOrderAmendKeepPriorityRequest {
	r.symbol = &symbol
	return r
}

// &#x60;newQty&#x60; must be greater than 0 and less than the order&#39;s quantity.
func (r ApiOrderAmendKeepPriorityRequest) NewQty(newQty float64) ApiOrderAmendKeepPriorityRequest {
	r.newQty = &newQty
	return r
}

// Client-generated request identifier.
func (r ApiOrderAmendKeepPriorityRequest) Id(id string) ApiOrderAmendKeepPriorityRequest {
	r.id = &id
	return r
}

// &#x60;orderId&#x60; or &#x60;origClientOrderId&#x60; must be sent
func (r ApiOrderAmendKeepPriorityRequest) OrderId(orderId int64) ApiOrderAmendKeepPriorityRequest {
	r.orderId = &orderId
	return r
}

// &#x60;orderId&#x60; or &#x60;origClientOrderId&#x60; must be sent
func (r ApiOrderAmendKeepPriorityRequest) OrigClientOrderId(origClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// The new client order ID for the order after being amended. &lt;br&gt; If not sent, one will be randomly generated. &lt;br&gt; It is possible to reuse the current clientOrderId by sending it as the &#x60;newClientOrderId&#x60;.
func (r ApiOrderAmendKeepPriorityRequest) NewClientOrderId(newClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderAmendKeepPriorityRequest) RecvWindow(recvWindow float64) ApiOrderAmendKeepPriorityRequest {
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
OrderAmendKeepPriority Order Amend Keep Priority (TRADE)
/order.amend.keepPriority

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-amend-keep-priority

@param symbol	@param newQty `newQty` must be greater than 0 and less than the order's quantity.	@param id Client-generated request identifier.	@param orderId `orderId` or `origClientOrderId` must be sent	@param origClientOrderId `orderId` or `origClientOrderId` must be sent	@param newClientOrderId The new client order ID for the order after being amended. <br> If not sent, one will be randomly generated. <br> It is possible to reuse the current clientOrderId by sending it as the `newClientOrderId`.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	recvWindow         *float64
}

func (r ApiOrderCancelRequest) Symbol(symbol string) ApiOrderCancelRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiOrderCancelRequest) Id(id string) ApiOrderCancelRequest {
	r.id = &id
	return r
}

func (r ApiOrderCancelRequest) OrderId(orderId int64) ApiOrderCancelRequest {
	r.orderId = &orderId
	return r
}

func (r ApiOrderCancelRequest) OrigClientOrderId(origClientOrderId string) ApiOrderCancelRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default.
func (r ApiOrderCancelRequest) NewClientOrderId(newClientOrderId string) ApiOrderCancelRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Supported values: &lt;br&gt;&#x60;ONLY_NEW&#x60; - Cancel will succeed if the order status is &#x60;NEW&#x60;.&lt;br&gt; &#x60;ONLY_PARTIALLY_FILLED&#x60; - Cancel will succeed if order status is &#x60;PARTIALLY_FILLED&#x60;.
func (r ApiOrderCancelRequest) CancelRestrictions(cancelRestrictions models.OrderCancelCancelRestrictionsParameter) ApiOrderCancelRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderCancelRequest) RecvWindow(recvWindow float64) ApiOrderCancelRequest {
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
OrderCancel Cancel order (TRADE)
/order.cancel

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-cancel

@param symbol	@param id Client-generated request identifier.	@param orderId	@param origClientOrderId	@param newClientOrderId Used to uniquely identify this cancel. Automatically generated by default.	@param cancelRestrictions Supported values: <br>`ONLY_NEW` - Cancel will succeed if the order status is `NEW`.<br> `ONLY_PARTIALLY_FILLED` - Cancel will succeed if order status is `PARTIALLY_FILLED`.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	price                      *float64
	quantity                   *float64
	quoteOrderQty              *float64
	newClientOrderId           *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice                  *float64
	trailingDelta              *float64
	icebergQty                 *float64
	strategyId                 *int64
	strategyType               *int32
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	cancelRestrictions         *models.OrderCancelCancelRestrictionsParameter
	orderRateLimitExceededMode *models.OrderCancelReplaceOrderRateLimitExceededModeParameter
	pegPriceType               *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue             *int32
	pegOffsetType              *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow                 *float64
}

func (r ApiOrderCancelReplaceRequest) Symbol(symbol string) ApiOrderCancelReplaceRequest {
	r.symbol = &symbol
	return r
}

// The allowed values are: &lt;br/&gt; &#x60;STOP_ON_FAILURE&#x60; - If the cancel request fails, the new order placement will not be attempted. &lt;br/&gt; &#x60;ALLOW_FAILURE&#x60; - new order placement will be attempted even if cancel request fails.
func (r ApiOrderCancelReplaceRequest) CancelReplaceMode(cancelReplaceMode models.OrderCancelReplaceCancelReplaceModeParameter) ApiOrderCancelReplaceRequest {
	r.cancelReplaceMode = &cancelReplaceMode
	return r
}

// Please see [Enums](/products/spot/enums#side) for supported values.
func (r ApiOrderCancelReplaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderCancelReplaceRequest {
	r.side = &side
	return r
}

// Please see [Enums](/products/spot/enums#ordertypes) for supported values.
func (r ApiOrderCancelReplaceRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderCancelReplaceRequest {
	r.type_ = &type_
	return r
}

// Client-generated request identifier.
func (r ApiOrderCancelReplaceRequest) Id(id string) ApiOrderCancelReplaceRequest {
	r.id = &id
	return r
}

// Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt;If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt;If both conditions are not met the request will be rejected.
func (r ApiOrderCancelReplaceRequest) CancelOrderId(cancelOrderId int64) ApiOrderCancelReplaceRequest {
	r.cancelOrderId = &cancelOrderId
	return r
}

// Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt; If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt; If both conditions are not met the request will be rejected.
func (r ApiOrderCancelReplaceRequest) CancelOrigClientOrderId(cancelOrigClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default.
func (r ApiOrderCancelReplaceRequest) CancelNewClientOrderId(cancelNewClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelNewClientOrderId = &cancelNewClientOrderId
	return r
}

// Please see [Enums](/products/spot/enums#timeinforce) for supported values.
func (r ApiOrderCancelReplaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderCancelReplaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderCancelReplaceRequest) Price(price float64) ApiOrderCancelReplaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderCancelReplaceRequest) Quantity(quantity float64) ApiOrderCancelReplaceRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderCancelReplaceRequest) QuoteOrderQty(quoteOrderQty float64) ApiOrderCancelReplaceRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// Used to identify the new order.
func (r ApiOrderCancelReplaceRequest) NewClientOrderId(newClientOrderId string) ApiOrderCancelReplaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Allowed values: &lt;br/&gt; &#x60;ACK&#x60;, &#x60;RESULT&#x60;, &#x60;FULL&#x60; &lt;br/&gt; &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; orders types default to &#x60;FULL&#x60;; all other orders default to &#x60;ACK&#x60;
func (r ApiOrderCancelReplaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderCancelReplaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiOrderCancelReplaceRequest) StopPrice(stopPrice float64) ApiOrderCancelReplaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderCancelReplaceRequest) TrailingDelta(trailingDelta float64) ApiOrderCancelReplaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiOrderCancelReplaceRequest) IcebergQty(icebergQty float64) ApiOrderCancelReplaceRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiOrderCancelReplaceRequest) StrategyId(strategyId int64) ApiOrderCancelReplaceRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderCancelReplaceRequest) StrategyType(strategyType int32) ApiOrderCancelReplaceRequest {
	r.strategyType = &strategyType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes).
func (r ApiOrderCancelReplaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderCancelReplaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Supported values: &lt;br&gt;&#x60;ONLY_NEW&#x60; - Cancel will succeed if the order status is &#x60;NEW&#x60;.&lt;br&gt; &#x60;ONLY_PARTIALLY_FILLED&#x60; - Cancel will succeed if order status is &#x60;PARTIALLY_FILLED&#x60;.
func (r ApiOrderCancelReplaceRequest) CancelRestrictions(cancelRestrictions models.OrderCancelCancelRestrictionsParameter) ApiOrderCancelReplaceRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

// Supported values: &lt;br&gt; &#x60;DO_NOTHING&#x60; (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit&lt;br&gt; &#x60;CANCEL_ONLY&#x60; - will always cancel the order
func (r ApiOrderCancelReplaceRequest) OrderRateLimitExceededMode(orderRateLimitExceededMode models.OrderCancelReplaceOrderRateLimitExceededModeParameter) ApiOrderCancelReplaceRequest {
	r.orderRateLimitExceededMode = &orderRateLimitExceededMode
	return r
}

// &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60; &lt;br&gt; See Pegged Orders
func (r ApiOrderCancelReplaceRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level to peg the price to (max: 100) &lt;br&gt; See Pegged Orders
func (r ApiOrderCancelReplaceRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderCancelReplaceRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

// Only &#x60;PRICE_LEVEL&#x60; is supported &lt;br&gt; See Pegged Orders
func (r ApiOrderCancelReplaceRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderCancelReplaceRequest) RecvWindow(recvWindow float64) ApiOrderCancelReplaceRequest {
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
OrderCancelReplace Cancel and replace order (TRADE)
/order.cancelReplace

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-cancel-replace

@param symbol	@param cancelReplaceMode The allowed values are: <br/> `STOP_ON_FAILURE` - If the cancel request fails, the new order placement will not be attempted. <br/> `ALLOW_FAILURE` - new order placement will be attempted even if cancel request fails.	@param side Please see [Enums](/products/spot/enums#side) for supported values.	@param type_ Please see [Enums](/products/spot/enums#ordertypes) for supported values.	@param id Client-generated request identifier.	@param cancelOrderId Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br>If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br>If both conditions are not met the request will be rejected.	@param cancelOrigClientOrderId Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br> If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br> If both conditions are not met the request will be rejected.	@param cancelNewClientOrderId Used to uniquely identify this cancel. Automatically generated by default.	@param timeInForce Please see [Enums](/products/spot/enums#timeinforce) for supported values.	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId Used to identify the new order.	@param newOrderRespType Allowed values: <br/> `ACK`, `RESULT`, `FULL` <br/> `MARKET` and `LIMIT` orders types default to `FULL`; all other orders default to `ACK`	@param stopPrice Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.	@param trailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param icebergQty Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.	@param strategyId	@param strategyType The value cannot be less than `1000000`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes).	@param cancelRestrictions Supported values: <br>`ONLY_NEW` - Cancel will succeed if the order status is `NEW`.<br> `ONLY_PARTIALLY_FILLED` - Cancel will succeed if order status is `PARTIALLY_FILLED`.	@param orderRateLimitExceededMode Supported values: <br> `DO_NOTHING` (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit<br> `CANCEL_ONLY` - will always cancel the order	@param pegPriceType `PRIMARY_PEG` or `MARKET_PEG` <br> See Pegged Orders	@param pegOffsetValue Price level to peg the price to (max: 100) <br> See Pegged Orders	@param pegOffsetType Only `PRICE_LEVEL` is supported <br> See Pegged Orders	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	recvWindow        *float64
}

func (r ApiOrderListCancelRequest) Symbol(symbol string) ApiOrderListCancelRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiOrderListCancelRequest) Id(id string) ApiOrderListCancelRequest {
	r.id = &id
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiOrderListCancelRequest) OrderListId(orderListId int32) ApiOrderListCancelRequest {
	r.orderListId = &orderListId
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiOrderListCancelRequest) ListClientOrderId(listClientOrderId string) ApiOrderListCancelRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default.
func (r ApiOrderListCancelRequest) NewClientOrderId(newClientOrderId string) ApiOrderListCancelRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListCancelRequest) RecvWindow(recvWindow float64) ApiOrderListCancelRequest {
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
OrderListCancel Cancel Order list (TRADE)
/orderList.cancel

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-cancel

@param symbol	@param id Client-generated request identifier.	@param orderListId Either `orderListId` or `listClientOrderId` must be provided	@param listClientOrderId Either `orderListId` or `listClientOrderId` must be provided	@param newClientOrderId Used to uniquely identify this cancel. Automatically generated by default.	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	price                   *float64
	quantity                *float64
	id                      *string
	listClientOrderId       *string
	limitClientOrderId      *string
	limitIcebergQty         *float64
	limitStrategyId         *int64
	limitStrategyType       *int32
	stopPrice               *float64
	trailingDelta           *int32
	stopClientOrderId       *string
	stopLimitPrice          *float64
	stopLimitTimeInForce    *models.OrderCancelReplaceTimeInForceParameter
	stopIcebergQty          *float64
	stopStrategyId          *int64
	stopStrategyType        *int32
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float64
}

func (r ApiOrderListPlaceRequest) Symbol(symbol string) ApiOrderListPlaceRequest {
	r.symbol = &symbol
	return r
}

// Please see [Enums](/products/spot/enums#side) for supported values.
func (r ApiOrderListPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderListPlaceRequest {
	r.side = &side
	return r
}

func (r ApiOrderListPlaceRequest) Price(price float64) ApiOrderListPlaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderListPlaceRequest) Quantity(quantity float64) ApiOrderListPlaceRequest {
	r.quantity = &quantity
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceRequest) Id(id string) ApiOrderListPlaceRequest {
	r.id = &id
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListPlaceRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique Id for the limit order
func (r ApiOrderListPlaceRequest) LimitClientOrderId(limitClientOrderId string) ApiOrderListPlaceRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

// Used to make the &#x60;LIMIT_MAKER&#x60; leg an iceberg order.
func (r ApiOrderListPlaceRequest) LimitIcebergQty(limitIcebergQty float64) ApiOrderListPlaceRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

func (r ApiOrderListPlaceRequest) LimitStrategyId(limitStrategyId int64) ApiOrderListPlaceRequest {
	r.limitStrategyId = &limitStrategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderListPlaceRequest) LimitStrategyType(limitStrategyType int32) ApiOrderListPlaceRequest {
	r.limitStrategyType = &limitStrategyType
	return r
}

func (r ApiOrderListPlaceRequest) StopPrice(stopPrice float64) ApiOrderListPlaceRequest {
	r.stopPrice = &stopPrice
	return r
}

func (r ApiOrderListPlaceRequest) TrailingDelta(trailingDelta int32) ApiOrderListPlaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// A unique Id for the stop loss/stop loss limit leg
func (r ApiOrderListPlaceRequest) StopClientOrderId(stopClientOrderId string) ApiOrderListPlaceRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

// If provided, &#x60;stopLimitTimeInForce&#x60; is required.
func (r ApiOrderListPlaceRequest) StopLimitPrice(stopLimitPrice float64) ApiOrderListPlaceRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

// Valid values are &#x60;GTC&#x60;/&#x60;FOK&#x60;/&#x60;IOC&#x60;
func (r ApiOrderListPlaceRequest) StopLimitTimeInForce(stopLimitTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

// Used with &#x60;STOP_LOSS_LIMIT&#x60; leg to make an iceberg order.
func (r ApiOrderListPlaceRequest) StopIcebergQty(stopIcebergQty float64) ApiOrderListPlaceRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

func (r ApiOrderListPlaceRequest) StopStrategyId(stopStrategyId int64) ApiOrderListPlaceRequest {
	r.stopStrategyId = &stopStrategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderListPlaceRequest) StopStrategyType(stopStrategyType int32) ApiOrderListPlaceRequest {
	r.stopStrategyType = &stopStrategyType
	return r
}

// Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)
func (r ApiOrderListPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed values are dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
func (r ApiOrderListPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceRequest {
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
	OrderListPlace Place new OCO - Deprecated (TRADE)
	/orderList.place

	https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place

	@param symbol	@param side Please see [Enums](/products/spot/enums#side) for supported values.	@param price	@param quantity	@param id Client-generated request identifier.	@param listClientOrderId A unique Id for the entire orderList	@param limitClientOrderId A unique Id for the limit order	@param limitIcebergQty Used to make the `LIMIT_MAKER` leg an iceberg order.	@param limitStrategyId	@param limitStrategyType The value cannot be less than `1000000`.	@param stopPrice	@param trailingDelta	@param stopClientOrderId A unique Id for the stop loss/stop loss limit leg	@param stopLimitPrice If provided, `stopLimitTimeInForce` is required.	@param stopLimitTimeInForce Valid values are `GTC`/`FOK`/`IOC`	@param stopIcebergQty Used with `STOP_LOSS_LIMIT` leg to make an iceberg order.	@param stopStrategyId	@param stopStrategyType The value cannot be less than `1000000`.	@param newOrderRespType Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)	@param selfTradePreventionMode The allowed values are dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	quantity                *float64
	aboveType               *models.OrderListPlaceOcoAboveTypeParameter
	belowType               *models.OrderListPlaceOcoBelowTypeParameter
	id                      *string
	listClientOrderId       *string
	aboveClientOrderId      *string
	aboveIcebergQty         *int64
	abovePrice              *float64
	aboveStopPrice          *float64
	aboveTrailingDelta      *int64
	aboveTimeInForce        *models.OrderCancelReplaceTimeInForceParameter
	aboveStrategyId         *int64
	aboveStrategyType       *int32
	abovePegPriceType       *models.OrderCancelReplacePegPriceTypeParameter
	abovePegOffsetType      *models.OrderCancelReplacePegOffsetTypeParameter
	abovePegOffsetValue     *int32
	belowClientOrderId      *string
	belowIcebergQty         *int64
	belowPrice              *float64
	belowStopPrice          *float64
	belowTrailingDelta      *int64
	belowTimeInForce        *models.OrderCancelReplaceTimeInForceParameter
	belowStrategyId         *int64
	belowStrategyType       *int32
	belowPegPriceType       *models.OrderCancelReplacePegPriceTypeParameter
	belowPegOffsetType      *models.OrderCancelReplacePegOffsetTypeParameter
	belowPegOffsetValue     *int32
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float64
}

func (r ApiOrderListPlaceOcoRequest) Symbol(symbol string) ApiOrderListPlaceOcoRequest {
	r.symbol = &symbol
	return r
}

// &#x60;BUY&#x60; or &#x60;SELL&#x60;
func (r ApiOrderListPlaceOcoRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOcoRequest {
	r.side = &side
	return r
}

// Quantity for both orders of the order list.
func (r ApiOrderListPlaceOcoRequest) Quantity(quantity float64) ApiOrderListPlaceOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderListPlaceOcoRequest) AboveType(aboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOcoRequest {
	r.aboveType = &aboveType
	return r
}

// Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOcoRequest) BelowType(belowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowType = &belowType
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceOcoRequest) Id(id string) ApiOrderListPlaceOcoRequest {
	r.id = &id
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;aboveClientOrderId&#x60; and the &#x60;belowClientOrderId&#x60;.
func (r ApiOrderListPlaceOcoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Arbitrary unique ID among open orders for the above order. Automatically generated if not sent.
func (r ApiOrderListPlaceOcoRequest) AboveClientOrderId(aboveClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.aboveClientOrderId = &aboveClientOrderId
	return r
}

// Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListPlaceOcoRequest) AboveIcebergQty(aboveIcebergQty int64) ApiOrderListPlaceOcoRequest {
	r.aboveIcebergQty = &aboveIcebergQty
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOcoRequest) AbovePrice(abovePrice float64) ApiOrderListPlaceOcoRequest {
	r.abovePrice = &abovePrice
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOcoRequest) AboveStopPrice(aboveStopPrice float64) ApiOrderListPlaceOcoRequest {
	r.aboveStopPrice = &aboveStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOcoRequest) AboveTrailingDelta(aboveTrailingDelta int64) ApiOrderListPlaceOcoRequest {
	r.aboveTrailingDelta = &aboveTrailingDelta
	return r
}

// Required if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOcoRequest) AboveTimeInForce(aboveTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOcoRequest {
	r.aboveTimeInForce = &aboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the above order within an order strategy.
func (r ApiOrderListPlaceOcoRequest) AboveStrategyId(aboveStrategyId int64) ApiOrderListPlaceOcoRequest {
	r.aboveStrategyId = &aboveStrategyId
	return r
}

// Arbitrary numeric value identifying the above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOcoRequest) AboveStrategyType(aboveStrategyType int32) ApiOrderListPlaceOcoRequest {
	r.aboveStrategyType = &aboveStrategyType
	return r
}

// &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60;. See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOcoRequest) AbovePegPriceType(abovePegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOcoRequest {
	r.abovePegPriceType = &abovePegPriceType
	return r
}

func (r ApiOrderListPlaceOcoRequest) AbovePegOffsetType(abovePegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOcoRequest {
	r.abovePegOffsetType = &abovePegOffsetType
	return r
}

func (r ApiOrderListPlaceOcoRequest) AbovePegOffsetValue(abovePegOffsetValue int32) ApiOrderListPlaceOcoRequest {
	r.abovePegOffsetValue = &abovePegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the below order. Automatically generated if not sent.
func (r ApiOrderListPlaceOcoRequest) BelowClientOrderId(belowClientOrderId string) ApiOrderListPlaceOcoRequest {
	r.belowClientOrderId = &belowClientOrderId
	return r
}

// Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListPlaceOcoRequest) BelowIcebergQty(belowIcebergQty int64) ApiOrderListPlaceOcoRequest {
	r.belowIcebergQty = &belowIcebergQty
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOcoRequest) BelowPrice(belowPrice float64) ApiOrderListPlaceOcoRequest {
	r.belowPrice = &belowPrice
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;belowStopPrice&#x60; or &#x60;belowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOcoRequest) BelowStopPrice(belowStopPrice float64) ApiOrderListPlaceOcoRequest {
	r.belowStopPrice = &belowStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOcoRequest) BelowTrailingDelta(belowTrailingDelta int64) ApiOrderListPlaceOcoRequest {
	r.belowTrailingDelta = &belowTrailingDelta
	return r
}

// Required if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOcoRequest) BelowTimeInForce(belowTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOcoRequest {
	r.belowTimeInForce = &belowTimeInForce
	return r
}

// Arbitrary numeric value identifying the below order within an order strategy.
func (r ApiOrderListPlaceOcoRequest) BelowStrategyId(belowStrategyId int64) ApiOrderListPlaceOcoRequest {
	r.belowStrategyId = &belowStrategyId
	return r
}

// Arbitrary numeric value identifying the below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOcoRequest) BelowStrategyType(belowStrategyType int32) ApiOrderListPlaceOcoRequest {
	r.belowStrategyType = &belowStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOcoRequest) BelowPegPriceType(belowPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowPegPriceType = &belowPegPriceType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowPegOffsetType(belowPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOcoRequest {
	r.belowPegOffsetType = &belowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOcoRequest) BelowPegOffsetValue(belowPegOffsetValue int32) ApiOrderListPlaceOcoRequest {
	r.belowPegOffsetValue = &belowPegOffsetValue
	return r
}

// Select response format: &#x60;ACK&#x60;, &#x60;RESULT&#x60;, &#x60;FULL&#x60;.
func (r ApiOrderListPlaceOcoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
func (r ApiOrderListPlaceOcoRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderListPlaceOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOcoRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceOcoRequest {
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
OrderListPlaceOco Place new Order list - OCO (TRADE)
/orderList.place.oco

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place-oco

@param symbol	@param side `BUY` or `SELL`	@param quantity Quantity for both orders of the order list.	@param aboveType	@param belowType Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`	@param id Client-generated request identifier.	@param listClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `aboveClientOrderId` and the `belowClientOrderId`.	@param aboveClientOrderId Arbitrary unique ID among open orders for the above order. Automatically generated if not sent.	@param aboveIcebergQty Note that this can only be used if `aboveTimeInForce` is `GTC`.	@param abovePrice Can be used if `aboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param aboveStopPrice Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified.	@param aboveTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param aboveTimeInForce Required if `aboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`.	@param aboveStrategyId Arbitrary numeric value identifying the above order within an order strategy.	@param aboveStrategyType Arbitrary numeric value identifying the above order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param abovePegPriceType `PRIMARY_PEG` or `MARKET_PEG`. See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param abovePegOffsetType	@param abovePegOffsetValue	@param belowClientOrderId Arbitrary unique ID among open orders for the below order. Automatically generated if not sent.	@param belowIcebergQty Note that this can only be used if `belowTimeInForce` is `GTC`.	@param belowPrice Can be used if `belowType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param belowStopPrice Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `belowStopPrice` or `belowTrailingDelta` or both, must be specified.	@param belowTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param belowTimeInForce Required if `belowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`.	@param belowStrategyId Arbitrary numeric value identifying the below order within an order strategy.	@param belowStrategyType Arbitrary numeric value identifying the below order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param belowPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param belowPegOffsetType	@param belowPegOffsetValue	@param newOrderRespType Select response format: `ACK`, `RESULT`, `FULL`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	workingPrice            *float64
	workingQuantity         *float64
	pendingType             *models.OrderListPlaceOpoPendingTypeParameter
	pendingSide             *models.OrderCancelReplaceSideParameter
	id                      *string
	listClientOrderId       *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float64
	workingTimeInForce      *models.OrderCancelReplaceTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderCancelReplacePegPriceTypeParameter
	workingPegOffsetType    *models.OrderCancelReplacePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float64
	pendingStopPrice        *float64
	pendingTrailingDelta    *float64
	pendingIcebergQty       *float64
	pendingTimeInForce      *models.OrderCancelReplaceTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegPriceType     *models.OrderCancelReplacePegPriceTypeParameter
	pendingPegOffsetType    *models.OrderCancelReplacePegOffsetTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float64
}

func (r ApiOrderListPlaceOpoRequest) Symbol(symbol string) ApiOrderListPlaceOpoRequest {
	r.symbol = &symbol
	return r
}

// Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;
func (r ApiOrderListPlaceOpoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOpoRequest {
	r.workingType = &workingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOpoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpoRequest {
	r.workingSide = &workingSide
	return r
}

// Price for the working order.
func (r ApiOrderListPlaceOpoRequest) WorkingPrice(workingPrice float64) ApiOrderListPlaceOpoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOpoRequest) WorkingQuantity(workingQuantity float64) ApiOrderListPlaceOpoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// Supported values: [Order Types](/products/spot/enums#ordertypes). Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported.
func (r ApiOrderListPlaceOpoRequest) PendingType(pendingType models.OrderListPlaceOpoPendingTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingType = &pendingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOpoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceOpoRequest) Id(id string) ApiOrderListPlaceOpoRequest {
	r.id = &id
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;.
func (r ApiOrderListPlaceOpoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOpoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)
func (r ApiOrderListPlaceOpoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOpoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
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
func (r ApiOrderListPlaceOpoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiOrderListPlaceOpoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOpoRequest) WorkingTimeInForce(workingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOpoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOpoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOpoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOpoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOpoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpoRequest) WorkingPegPriceType(workingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOpoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOpoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOpoRequest {
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

// Price for the pending order.
func (r ApiOrderListPlaceOpoRequest) PendingPrice(pendingPrice float64) ApiOrderListPlaceOpoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

// Stop price for the pending order.
func (r ApiOrderListPlaceOpoRequest) PendingStopPrice(pendingStopPrice float64) ApiOrderListPlaceOpoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

// Trailing delta for the pending order.
func (r ApiOrderListPlaceOpoRequest) PendingTrailingDelta(pendingTrailingDelta float64) ApiOrderListPlaceOpoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpoRequest) PendingIcebergQty(pendingIcebergQty float64) ApiOrderListPlaceOpoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOpoRequest) PendingTimeInForce(pendingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOpoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListPlaceOpoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListPlaceOpoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOpoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListPlaceOpoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpoRequest) PendingPegPriceType(pendingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOpoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListPlaceOpoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOpoRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceOpoRequest {
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
OrderListPlaceOpo OPO (TRADE)
/orderList.place.opo

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place-opo

@param symbol	@param workingType Supported values: `LIMIT`, `LIMIT_MAKER`	@param workingSide Supported values: [Order Side](/products/spot/enums#side)	@param workingPrice Price for the working order.	@param workingQuantity Sets the quantity for the working order.	@param pendingType Supported values: [Order Types](/products/spot/enums#ordertypes). Note that `MARKET` orders using `quoteOrderQty` are not supported.	@param pendingSide Supported values: [Order Side](/products/spot/enums#side)	@param id Client-generated request identifier.	@param listClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`.	@param newOrderRespType Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.	@param workingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy.	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param workingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingClientOrderId Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.	@param pendingPrice Price for the pending order.	@param pendingStopPrice Stop price for the pending order.	@param pendingTrailingDelta Trailing delta for the pending order.	@param pendingIcebergQty This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.	@param pendingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param pendingStrategyId Arbitrary numeric value identifying the pending order within an order strategy.	@param pendingStrategyType Arbitrary numeric value identifying the pending order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingPegOffsetType	@param pendingPegOffsetValue	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	workingPrice               *float64
	workingQuantity            *float64
	pendingSide                *models.OrderCancelReplaceSideParameter
	pendingAboveType           *models.OrderListPlaceOcoAboveTypeParameter
	id                         *string
	listClientOrderId          *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float64
	workingTimeInForce         *models.OrderCancelReplaceTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderCancelReplacePegPriceTypeParameter
	workingPegOffsetType       *models.OrderCancelReplacePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float64
	pendingAboveStopPrice      *float64
	pendingAboveTrailingDelta  *float64
	pendingAboveIcebergQty     *float64
	pendingAboveTimeInForce    *models.OrderCancelReplaceTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderCancelReplacePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderCancelReplacePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListPlaceOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float64
	pendingBelowStopPrice      *float64
	pendingBelowTrailingDelta  *float64
	pendingBelowIcebergQty     *float64
	pendingBelowTimeInForce    *models.OrderCancelReplaceTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderCancelReplacePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderCancelReplacePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float64
}

func (r ApiOrderListPlaceOpocoRequest) Symbol(symbol string) ApiOrderListPlaceOpocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListPlaceOpocoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingType = &workingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOpocoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpocoRequest {
	r.workingSide = &workingSide
	return r
}

// Price for the working order.
func (r ApiOrderListPlaceOpocoRequest) WorkingPrice(workingPrice float64) ApiOrderListPlaceOpocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOpocoRequest) WorkingQuantity(workingQuantity float64) ApiOrderListPlaceOpocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOpocoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Supported values: &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;STOP_LOSS&#x60;, &#x60;LIMIT_MAKER&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOpocoRequest) PendingAboveType(pendingAboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceOpocoRequest) Id(id string) ApiOrderListPlaceOpocoRequest {
	r.id = &id
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;.
func (r ApiOrderListPlaceOpocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)
func (r ApiOrderListPlaceOpocoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
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
func (r ApiOrderListPlaceOpocoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiOrderListPlaceOpocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOpocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

// Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOpocoRequest) PendingAbovePrice(pendingAbovePrice float64) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float64) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOpocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float64) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float64) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

// Required if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

// Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

// Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOpocoRequest) PendingBelowType(pendingBelowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowPrice(pendingBelowPrice float64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOpocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60; or &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOpocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOpocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOpocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListPlaceOpocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOpocoRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceOpocoRequest {
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
OrderListPlaceOpoco OPOCO (TRADE)
/orderList.place.opoco

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place-opoco

@param symbol	@param workingType	@param workingSide Supported values: [Order Side](/products/spot/enums#side)	@param workingPrice Price for the working order.	@param workingQuantity Sets the quantity for the working order.	@param pendingSide Supported values: [Order Side](/products/spot/enums#side)	@param pendingAboveType Supported values: `STOP_LOSS_LIMIT`, `STOP_LOSS`, `LIMIT_MAKER`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`	@param id Client-generated request identifier.	@param listClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`.	@param newOrderRespType Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.	@param workingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy.	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param workingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param workingPegOffsetType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param workingPegOffsetValue Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingAboveClientOrderId Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.	@param pendingAbovePrice Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param pendingAboveStopPrice Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`.	@param pendingAboveTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param pendingAboveIcebergQty This can only be used if `pendingAboveTimeInForce` is `GTC` or `pendingAboveType` is `LIMIT_MAKER`.	@param pendingAboveTimeInForce Required if `pendingAboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`.	@param pendingAboveStrategyId Arbitrary numeric value identifying the pending above order within an order strategy.	@param pendingAboveStrategyType Arbitrary numeric value identifying the pending above order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingAbovePegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingAbovePegOffsetType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingAbovePegOffsetValue Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingBelowType Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`	@param pendingBelowClientOrderId Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.	@param pendingBelowPrice Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param pendingBelowStopPrice Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.	@param pendingBelowTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param pendingBelowIcebergQty This can only be used if `pendingBelowTimeInForce` is `GTC` or `pendingBelowType` is `LIMIT_MAKER`.	@param pendingBelowTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param pendingBelowStrategyId Arbitrary numeric value identifying the pending below order within an order strategy.	@param pendingBelowStrategyType Arbitrary numeric value identifying the pending below order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingBelowPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingBelowPegOffsetType	@param pendingBelowPegOffsetValue	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	workingPrice            *float64
	workingQuantity         *float64
	pendingType             *models.OrderListPlaceOpoPendingTypeParameter
	pendingSide             *models.OrderCancelReplaceSideParameter
	pendingQuantity         *float64
	id                      *string
	listClientOrderId       *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float64
	workingTimeInForce      *models.OrderCancelReplaceTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderCancelReplacePegPriceTypeParameter
	workingPegOffsetType    *models.OrderCancelReplacePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float64
	pendingStopPrice        *float64
	pendingTrailingDelta    *float64
	pendingIcebergQty       *float64
	pendingTimeInForce      *models.OrderCancelReplaceTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegOffsetType    *models.OrderCancelReplacePegOffsetTypeParameter
	pendingPegPriceType     *models.OrderCancelReplacePegPriceTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float64
}

func (r ApiOrderListPlaceOtoRequest) Symbol(symbol string) ApiOrderListPlaceOtoRequest {
	r.symbol = &symbol
	return r
}

// Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;
func (r ApiOrderListPlaceOtoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOtoRequest {
	r.workingType = &workingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOtoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPrice(workingPrice float64) ApiOrderListPlaceOtoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOtoRequest) WorkingQuantity(workingQuantity float64) ApiOrderListPlaceOtoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// Supported values: [Order Types](/products/spot/enums#ordertypes). Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported.
func (r ApiOrderListPlaceOtoRequest) PendingType(pendingType models.OrderListPlaceOpoPendingTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingType = &pendingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOtoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiOrderListPlaceOtoRequest) PendingQuantity(pendingQuantity float64) ApiOrderListPlaceOtoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceOtoRequest) Id(id string) ApiOrderListPlaceOtoRequest {
	r.id = &id
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;.
func (r ApiOrderListPlaceOtoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOtoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)
func (r ApiOrderListPlaceOtoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOtoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
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
func (r ApiOrderListPlaceOtoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiOrderListPlaceOtoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOtoRequest) WorkingTimeInForce(workingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOtoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOtoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOtoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOtoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOtoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOtoRequest) WorkingPegPriceType(workingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOtoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOtoRequest {
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

func (r ApiOrderListPlaceOtoRequest) PendingPrice(pendingPrice float64) ApiOrderListPlaceOtoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingStopPrice(pendingStopPrice float64) ApiOrderListPlaceOtoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingTrailingDelta(pendingTrailingDelta float64) ApiOrderListPlaceOtoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtoRequest) PendingIcebergQty(pendingIcebergQty float64) ApiOrderListPlaceOtoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOtoRequest) PendingTimeInForce(pendingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOtoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListPlaceOtoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListPlaceOtoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOtoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListPlaceOtoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOtoRequest) PendingPegPriceType(pendingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOtoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListPlaceOtoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOtoRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceOtoRequest {
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
OrderListPlaceOto Place new Order list - OTO (TRADE)
/orderList.place.oto

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place-oto

@param symbol	@param workingType Supported values: `LIMIT`, `LIMIT_MAKER`	@param workingSide Supported values: [Order Side](/products/spot/enums#side)	@param workingPrice	@param workingQuantity Sets the quantity for the working order.	@param pendingType Supported values: [Order Types](/products/spot/enums#ordertypes). Note that `MARKET` orders using `quoteOrderQty` are not supported.	@param pendingSide Supported values: [Order Side](/products/spot/enums#side)	@param pendingQuantity Sets the quantity for the pending order.	@param id Client-generated request identifier.	@param listClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`.	@param newOrderRespType Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.	@param workingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy.	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param workingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingClientOrderId Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.	@param pendingPrice	@param pendingStopPrice	@param pendingTrailingDelta	@param pendingIcebergQty This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.	@param pendingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param pendingStrategyId Arbitrary numeric value identifying the pending order within an order strategy.	@param pendingStrategyType Arbitrary numeric value identifying the pending order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingPegOffsetType	@param pendingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingPegOffsetValue	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	workingPrice               *float64
	workingQuantity            *float64
	pendingSide                *models.OrderCancelReplaceSideParameter
	pendingQuantity            *float64
	pendingAboveType           *models.OrderListPlaceOcoAboveTypeParameter
	id                         *string
	listClientOrderId          *string
	newOrderRespType           *models.OrderCancelReplaceNewOrderRespTypeParameter
	selfTradePreventionMode    *models.OrderCancelReplaceSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float64
	workingTimeInForce         *models.OrderCancelReplaceTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderCancelReplacePegPriceTypeParameter
	workingPegOffsetType       *models.OrderCancelReplacePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float64
	pendingAboveStopPrice      *float64
	pendingAboveTrailingDelta  *float64
	pendingAboveIcebergQty     *float64
	pendingAboveTimeInForce    *models.OrderCancelReplaceTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderCancelReplacePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderCancelReplacePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListPlaceOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float64
	pendingBelowStopPrice      *float64
	pendingBelowTrailingDelta  *float64
	pendingBelowIcebergQty     *float64
	pendingBelowTimeInForce    *models.OrderCancelReplaceTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderCancelReplacePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderCancelReplacePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float64
}

func (r ApiOrderListPlaceOtocoRequest) Symbol(symbol string) ApiOrderListPlaceOtocoRequest {
	r.symbol = &symbol
	return r
}

// Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;
func (r ApiOrderListPlaceOtocoRequest) WorkingType(workingType models.OrderListPlaceOpoWorkingTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.workingType = &workingType
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOtocoRequest) WorkingSide(workingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPrice(workingPrice float64) ApiOrderListPlaceOtocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListPlaceOtocoRequest) WorkingQuantity(workingQuantity float64) ApiOrderListPlaceOtocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// Supported values: [Order Side](/products/spot/enums#side)
func (r ApiOrderListPlaceOtocoRequest) PendingSide(pendingSide models.OrderCancelReplaceSideParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending orders.
func (r ApiOrderListPlaceOtocoRequest) PendingQuantity(pendingQuantity float64) ApiOrderListPlaceOtocoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// Supported values: &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;STOP_LOSS&#x60;, &#x60;LIMIT_MAKER&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOtocoRequest) PendingAboveType(pendingAboveType models.OrderListPlaceOcoAboveTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// Client-generated request identifier.
func (r ApiOrderListPlaceOtocoRequest) Id(id string) ApiOrderListPlaceOtocoRequest {
	r.id = &id
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same &#x60;listClientOrderId&#x60; is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;.
func (r ApiOrderListPlaceOtocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)
func (r ApiOrderListPlaceOtocoRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
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
func (r ApiOrderListPlaceOtocoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiOrderListPlaceOtocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// Supported values: [Time In Force](/products/spot/enums#timeinforce)
func (r ApiOrderListPlaceOtocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOtocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
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

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOtocoRequest) PendingAbovePrice(pendingAbovePrice float64) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float64) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOtocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float64) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float64) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

// Required if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListPlaceOtocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

// Supported values: &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListPlaceOtocoRequest) PendingBelowType(pendingBelowType models.OrderListPlaceOcoBelowTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowPrice(pendingBelowPrice float64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderListPlaceOtocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

// Required if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60;.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than &#x60;1000000&#x60; are reserved and cannot be used.
func (r ApiOrderListPlaceOtocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

// See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListPlaceOtocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListPlaceOtocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListPlaceOtocoRequest) RecvWindow(recvWindow float64) ApiOrderListPlaceOtocoRequest {
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
OrderListPlaceOtoco Place new Order list - OTOCO (TRADE)
/orderList.place.otoco

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-list-place-otoco

@param symbol	@param workingType Supported values: `LIMIT`, `LIMIT_MAKER`	@param workingSide Supported values: [Order Side](/products/spot/enums#side)	@param workingPrice	@param workingQuantity Sets the quantity for the working order.	@param pendingSide Supported values: [Order Side](/products/spot/enums#side)	@param pendingQuantity Sets the quantity for the pending orders.	@param pendingAboveType Supported values: `STOP_LOSS_LIMIT`, `STOP_LOSS`, `LIMIT_MAKER`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`	@param id Client-generated request identifier.	@param listClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same `listClientOrderId` is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`.	@param newOrderRespType Format of the JSON response. Supported values: [Order Response Type](/products/spot/enums#orderresponsetype)	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param workingClientOrderId Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.	@param workingIcebergQty This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.	@param workingTimeInForce Supported values: [Time In Force](/products/spot/enums#timeinforce)	@param workingStrategyId Arbitrary numeric value identifying the working order within an order strategy.	@param workingStrategyType Arbitrary numeric value identifying the working order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param workingPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param workingPegOffsetType	@param workingPegOffsetValue	@param pendingAboveClientOrderId Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.	@param pendingAbovePrice Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param pendingAboveStopPrice Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`.	@param pendingAboveTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param pendingAboveIcebergQty This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.	@param pendingAboveTimeInForce Required if `pendingAboveType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`.	@param pendingAboveStrategyId Arbitrary numeric value identifying the pending above order within an order strategy.	@param pendingAboveStrategyType Arbitrary numeric value identifying the pending above order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingAbovePegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingAbovePegOffsetType	@param pendingAbovePegOffsetValue	@param pendingBelowType Supported values: `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`	@param pendingBelowClientOrderId Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.	@param pendingBelowPrice Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify the limit price.	@param pendingBelowStopPrice Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.	@param pendingBelowTrailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param pendingBelowIcebergQty This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.	@param pendingBelowTimeInForce Required if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT`.	@param pendingBelowStrategyId Arbitrary numeric value identifying the pending below order within an order strategy.	@param pendingBelowStrategyType Arbitrary numeric value identifying the pending below order strategy. Values smaller than `1000000` are reserved and cannot be used.	@param pendingBelowPegPriceType See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pendingBelowPegOffsetType	@param pendingBelowPegOffsetValue	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	price                   *float64
	quantity                *float64
	quoteOrderQty           *float64
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice               *float64
	trailingDelta           *int32
	icebergQty              *float64
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	pegPriceType            *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow              *float64
}

func (r ApiOrderPlaceRequest) Symbol(symbol string) ApiOrderPlaceRequest {
	r.symbol = &symbol
	return r
}

// Please see [Enums](/products/spot/enums#side) for supported values.
func (r ApiOrderPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderPlaceRequest {
	r.side = &side
	return r
}

// Please see [Enums](/products/spot/enums#ordertypes) for supported values.
func (r ApiOrderPlaceRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderPlaceRequest {
	r.type_ = &type_
	return r
}

// Client-generated request identifier.
func (r ApiOrderPlaceRequest) Id(id string) ApiOrderPlaceRequest {
	r.id = &id
	return r
}

// Please see [Enums](/products/spot/enums#timeinforce) for supported values.
func (r ApiOrderPlaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderPlaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderPlaceRequest) Price(price float64) ApiOrderPlaceRequest {
	r.price = &price
	return r
}

func (r ApiOrderPlaceRequest) Quantity(quantity float64) ApiOrderPlaceRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderPlaceRequest) QuoteOrderQty(quoteOrderQty float64) ApiOrderPlaceRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiOrderPlaceRequest) NewClientOrderId(newClientOrderId string) ApiOrderPlaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; order types default to &#x60;FULL&#x60;, all other orders default to &#x60;ACK&#x60;.
func (r ApiOrderPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiOrderPlaceRequest) StopPrice(stopPrice float64) ApiOrderPlaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See Trailing Stop order FAQ
func (r ApiOrderPlaceRequest) TrailingDelta(trailingDelta int32) ApiOrderPlaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiOrderPlaceRequest) IcebergQty(icebergQty float64) ApiOrderPlaceRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiOrderPlaceRequest) StrategyId(strategyId int64) ApiOrderPlaceRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderPlaceRequest) StrategyType(strategyType int32) ApiOrderPlaceRequest {
	r.strategyType = &strategyType
	return r
}

// The allowed enums is dependent on what is configured on the symbol.
func (r ApiOrderPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// See Pegged Orders Info
func (r ApiOrderPlaceRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderPlaceRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level to peg the price to (max: 100). See Pegged Orders Info
func (r ApiOrderPlaceRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderPlaceRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

// Only &#x60;PRICE_LEVEL&#x60; is supported. See Pegged Orders Info
func (r ApiOrderPlaceRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderPlaceRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderPlaceRequest) RecvWindow(recvWindow float64) ApiOrderPlaceRequest {
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
OrderPlace Place new order (TRADE)
/order.place

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-place

@param symbol	@param side Please see [Enums](/products/spot/enums#side) for supported values.	@param type_ Please see [Enums](/products/spot/enums#ordertypes) for supported values.	@param id Client-generated request identifier.	@param timeInForce Please see [Enums](/products/spot/enums#timeinforce) for supported values.	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.	@param newOrderRespType `MARKET` and `LIMIT` order types default to `FULL`, all other orders default to `ACK`.	@param stopPrice Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.	@param trailingDelta See Trailing Stop order FAQ	@param icebergQty Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.	@param strategyId	@param strategyType The value cannot be less than `1000000`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol.	@param pegPriceType See Pegged Orders Info	@param pegOffsetValue Price level to peg the price to (max: 100). See Pegged Orders Info	@param pegOffsetType Only `PRICE_LEVEL` is supported. See Pegged Orders Info	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	price                   *float64
	quantity                *float64
	quoteOrderQty           *float64
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	stopPrice               *float64
	trailingDelta           *int32
	icebergQty              *float64
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	pegPriceType            *models.OrderCancelReplacePegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.OrderCancelReplacePegOffsetTypeParameter
	recvWindow              *float64
}

func (r ApiOrderTestRequest) Symbol(symbol string) ApiOrderTestRequest {
	r.symbol = &symbol
	return r
}

// Please see [Enums](/products/spot/enums#side) for supported values.
func (r ApiOrderTestRequest) Side(side models.OrderCancelReplaceSideParameter) ApiOrderTestRequest {
	r.side = &side
	return r
}

// Please see [Enums](/products/spot/enums#ordertypes) for supported values.
func (r ApiOrderTestRequest) Type(type_ models.OrderCancelReplaceTypeParameter) ApiOrderTestRequest {
	r.type_ = &type_
	return r
}

// Client-generated request identifier.
func (r ApiOrderTestRequest) Id(id string) ApiOrderTestRequest {
	r.id = &id
	return r
}

// Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](/products/spot/faqs/commission_faq#test-order-diferences) to learn more.
func (r ApiOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

// Please see [Enums](/products/spot/enums#timeinforce) for supported values.
func (r ApiOrderTestRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiOrderTestRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiOrderTestRequest) Price(price float64) ApiOrderTestRequest {
	r.price = &price
	return r
}

func (r ApiOrderTestRequest) Quantity(quantity float64) ApiOrderTestRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderTestRequest) QuoteOrderQty(quoteOrderQty float64) ApiOrderTestRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

// A unique id among open orders. Automatically generated if not sent. Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;; &#x60;MARKET&#x60; and &#x60;LIMIT&#x60; order types default to &#x60;FULL&#x60;, all other orders default to &#x60;ACK&#x60;.
func (r ApiOrderTestRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiOrderTestRequest) StopPrice(stopPrice float64) ApiOrderTestRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)
func (r ApiOrderTestRequest) TrailingDelta(trailingDelta int32) ApiOrderTestRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiOrderTestRequest) IcebergQty(icebergQty float64) ApiOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiOrderTestRequest) StrategyId(strategyId int64) ApiOrderTestRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderTestRequest) StrategyType(strategyType int32) ApiOrderTestRequest {
	r.strategyType = &strategyType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
func (r ApiOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// &#x60;PRIMARY_PEG&#x60; or &#x60;MARKET_PEG&#x60;. See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderTestRequest) PegPriceType(pegPriceType models.OrderCancelReplacePegPriceTypeParameter) ApiOrderTestRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderTestRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderTestRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

// Only &#x60;PRICE_LEVEL&#x60; is supported. See [Pegged Orders](/products/spot/faqs/pegged_orders)
func (r ApiOrderTestRequest) PegOffsetType(pegOffsetType models.OrderCancelReplacePegOffsetTypeParameter) ApiOrderTestRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderTestRequest) RecvWindow(recvWindow float64) ApiOrderTestRequest {
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
OrderTest Test new order (TRADE)
/order.test

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#order-test

@param symbol	@param side Please see [Enums](/products/spot/enums#side) for supported values.	@param type_ Please see [Enums](/products/spot/enums#ordertypes) for supported values.	@param id Client-generated request identifier.	@param computeCommissionRates Default: `false` <br> See [Commissions FAQ](/products/spot/faqs/commission_faq#test-order-diferences) to learn more.	@param timeInForce Please see [Enums](/products/spot/enums#timeinforce) for supported values.	@param price	@param quantity	@param quoteOrderQty	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.	@param newOrderRespType Set the response JSON. `ACK`, `RESULT`, or `FULL`; `MARKET` and `LIMIT` order types default to `FULL`, all other orders default to `ACK`.	@param stopPrice Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.	@param trailingDelta See [Trailing Stop order FAQ](/products/spot/faqs/trailing-stop-faq)	@param icebergQty Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.	@param strategyId	@param strategyType The value cannot be less than `1000000`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param pegPriceType `PRIMARY_PEG` or `MARKET_PEG`. See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pegOffsetValue Price level for pegging (max: 100). See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param pegOffsetType Only `PRICE_LEVEL` is supported. See [Pegged Orders](/products/spot/faqs/pegged_orders)	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	type_                   *models.SorOrderPlaceTypeParameter
	quantity                *float64
	id                      *string
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float64
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	icebergQty              *float64
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float64
}

func (r ApiSorOrderPlaceRequest) Symbol(symbol string) ApiSorOrderPlaceRequest {
	r.symbol = &symbol
	return r
}

// &#x60;BUY&#x60; or &#x60;SELL&#x60;
func (r ApiSorOrderPlaceRequest) Side(side models.OrderCancelReplaceSideParameter) ApiSorOrderPlaceRequest {
	r.side = &side
	return r
}

// Only &#x60;LIMIT&#x60; and &#x60;MARKET&#x60; orders are supported.
func (r ApiSorOrderPlaceRequest) Type(type_ models.SorOrderPlaceTypeParameter) ApiSorOrderPlaceRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderPlaceRequest) Quantity(quantity float64) ApiSorOrderPlaceRequest {
	r.quantity = &quantity
	return r
}

// Client-generated request identifier.
func (r ApiSorOrderPlaceRequest) Id(id string) ApiSorOrderPlaceRequest {
	r.id = &id
	return r
}

// Applicable only to &#x60;LIMIT&#x60; order type.
func (r ApiSorOrderPlaceRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiSorOrderPlaceRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderPlaceRequest) Price(price float64) ApiSorOrderPlaceRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiSorOrderPlaceRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderPlaceRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;. Default to &#x60;FULL&#x60;
func (r ApiSorOrderPlaceRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiSorOrderPlaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;LIMIT&#x60; to create an iceberg order.
func (r ApiSorOrderPlaceRequest) IcebergQty(icebergQty float64) ApiSorOrderPlaceRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiSorOrderPlaceRequest) StrategyId(strategyId int64) ApiSorOrderPlaceRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiSorOrderPlaceRequest) StrategyType(strategyType int32) ApiSorOrderPlaceRequest {
	r.strategyType = &strategyType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes).
func (r ApiSorOrderPlaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiSorOrderPlaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderPlaceRequest) RecvWindow(recvWindow float64) ApiSorOrderPlaceRequest {
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
SorOrderPlace Place new order using SOR (TRADE)
/sor.order.place

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#sor-order-place

@param symbol	@param side `BUY` or `SELL`	@param type_ Only `LIMIT` and `MARKET` orders are supported.	@param quantity	@param id Client-generated request identifier.	@param timeInForce Applicable only to `LIMIT` order type.	@param price	@param newClientOrderId A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.	@param newOrderRespType Set the response JSON. `ACK`, `RESULT`, or `FULL`. Default to `FULL`	@param icebergQty Used with `LIMIT` to create an iceberg order.	@param strategyId	@param strategyType The value cannot be less than `1000000`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: [STP Modes](/products/spot/enums#stpmodes).	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
	type_                   *models.SorOrderPlaceTypeParameter
	quantity                *float64
	id                      *string
	computeCommissionRates  *bool
	timeInForce             *models.OrderCancelReplaceTimeInForceParameter
	price                   *float64
	newClientOrderId        *string
	newOrderRespType        *models.OrderCancelReplaceNewOrderRespTypeParameter
	icebergQty              *float64
	strategyId              *int64
	strategyType            *int32
	selfTradePreventionMode *models.OrderCancelReplaceSelfTradePreventionModeParameter
	recvWindow              *float64
}

func (r ApiSorOrderTestRequest) Symbol(symbol string) ApiSorOrderTestRequest {
	r.symbol = &symbol
	return r
}

// Please see [Enums](/products/spot/enums#side) for supported values.
func (r ApiSorOrderTestRequest) Side(side models.OrderCancelReplaceSideParameter) ApiSorOrderTestRequest {
	r.side = &side
	return r
}

// Please see [Enums](/products/spot/enums#ordertypes) for supported values.
func (r ApiSorOrderTestRequest) Type(type_ models.SorOrderPlaceTypeParameter) ApiSorOrderTestRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderTestRequest) Quantity(quantity float64) ApiSorOrderTestRequest {
	r.quantity = &quantity
	return r
}

// Client-generated request identifier.
func (r ApiSorOrderTestRequest) Id(id string) ApiSorOrderTestRequest {
	r.id = &id
	return r
}

// Default: &#x60;false&#x60;
func (r ApiSorOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiSorOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

// Please see [Enums](/products/spot/enums#timeinforce) for supported values.
func (r ApiSorOrderTestRequest) TimeInForce(timeInForce models.OrderCancelReplaceTimeInForceParameter) ApiSorOrderTestRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderTestRequest) Price(price float64) ApiSorOrderTestRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiSorOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Set the response JSON. &#x60;ACK&#x60;, &#x60;RESULT&#x60;, or &#x60;FULL&#x60;. Default to &#x60;FULL&#x60;.
func (r ApiSorOrderTestRequest) NewOrderRespType(newOrderRespType models.OrderCancelReplaceNewOrderRespTypeParameter) ApiSorOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// Used with &#x60;LIMIT&#x60; to create an iceberg order.
func (r ApiSorOrderTestRequest) IcebergQty(icebergQty float64) ApiSorOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiSorOrderTestRequest) StrategyId(strategyId int64) ApiSorOrderTestRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiSorOrderTestRequest) StrategyType(strategyType int32) ApiSorOrderTestRequest {
	r.strategyType = &strategyType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)
func (r ApiSorOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.OrderCancelReplaceSelfTradePreventionModeParameter) ApiSorOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderTestRequest) RecvWindow(recvWindow float64) ApiSorOrderTestRequest {
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
SorOrderTest Test new order using SOR (TRADE)
/sor.order.test

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/trade#sor-order-test

@param symbol	@param side Please see [Enums](/products/spot/enums#side) for supported values.	@param type_ Please see [Enums](/products/spot/enums#ordertypes) for supported values.	@param quantity	@param id Client-generated request identifier.	@param computeCommissionRates Default: `false`	@param timeInForce Please see [Enums](/products/spot/enums#timeinforce) for supported values.	@param price	@param newClientOrderId A unique id among open orders. Automatically generated if not sent. Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.	@param newOrderRespType Set the response JSON. `ACK`, `RESULT`, or `FULL`. Default to `FULL`.	@param icebergQty Used with `LIMIT` to create an iceberg order.	@param strategyId	@param strategyType The value cannot be less than `1000000`.	@param selfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. Supported values: [STP Modes](/products/spot/enums#stpmodes)	@param recvWindow Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
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
