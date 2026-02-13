/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package binancespotrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiDeleteOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *float32
}

func (r ApiDeleteOpenOrdersRequest) Symbol(symbol string) ApiDeleteOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiDeleteOpenOrdersRequest) RecvWindow(recvWindow float32) ApiDeleteOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteOpenOrdersRequest) Execute() (*common.RestApiResponse[models.DeleteOpenOrdersResponse], error) {
	return r.ApiService.DeleteOpenOrdersExecute(r)
}

/*
DeleteOpenOrders Cancel All Open Orders on a Symbol
Delete /api/v3/openOrders

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#cancel-all-open-orders-on-a-symbol-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiDeleteOpenOrdersRequest
*/
func (a *TradeAPIService) DeleteOpenOrders(ctx context.Context) ApiDeleteOpenOrdersRequest {
	return ApiDeleteOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteOpenOrdersResponse
func (a *TradeAPIService) DeleteOpenOrdersExecute(r ApiDeleteOpenOrdersRequest) (*common.RestApiResponse[models.DeleteOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DeleteOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDeleteOrderRequest struct {
	ctx                context.Context
	ApiService         *TradeAPIService
	symbol             *string
	orderId            *int64
	origClientOrderId  *string
	newClientOrderId   *string
	cancelRestrictions *models.DeleteOrderCancelRestrictionsParameter
	recvWindow         *float32
}

func (r ApiDeleteOrderRequest) Symbol(symbol string) ApiDeleteOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiDeleteOrderRequest) OrderId(orderId int64) ApiDeleteOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiDeleteOrderRequest) OrigClientOrderId(origClientOrderId string) ApiDeleteOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiDeleteOrderRequest) NewClientOrderId(newClientOrderId string) ApiDeleteOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiDeleteOrderRequest) CancelRestrictions(cancelRestrictions models.DeleteOrderCancelRestrictionsParameter) ApiDeleteOrderRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiDeleteOrderRequest) RecvWindow(recvWindow float32) ApiDeleteOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteOrderRequest) Execute() (*common.RestApiResponse[models.DeleteOrderResponse], error) {
	return r.ApiService.DeleteOrderExecute(r)
}

/*
DeleteOrder Cancel order
Delete /api/v3/order

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#cancel-order-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param cancelRestrictions -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiDeleteOrderRequest
*/
func (a *TradeAPIService) DeleteOrder(ctx context.Context) ApiDeleteOrderRequest {
	return ApiDeleteOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteOrderResponse
func (a *TradeAPIService) DeleteOrderExecute(r ApiDeleteOrderRequest) (*common.RestApiResponse[models.DeleteOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.cancelRestrictions != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelRestrictions", r.cancelRestrictions, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DeleteOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDeleteOrderListRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderListId       *int64
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *float32
}

func (r ApiDeleteOrderListRequest) Symbol(symbol string) ApiDeleteOrderListRequest {
	r.symbol = &symbol
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiDeleteOrderListRequest) OrderListId(orderListId int64) ApiDeleteOrderListRequest {
	r.orderListId = &orderListId
	return r
}

// A unique Id for the entire orderList
func (r ApiDeleteOrderListRequest) ListClientOrderId(listClientOrderId string) ApiDeleteOrderListRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiDeleteOrderListRequest) NewClientOrderId(newClientOrderId string) ApiDeleteOrderListRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiDeleteOrderListRequest) RecvWindow(recvWindow float32) ApiDeleteOrderListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteOrderListRequest) Execute() (*common.RestApiResponse[models.DeleteOrderListResponse], error) {
	return r.ApiService.DeleteOrderListExecute(r)
}

/*
DeleteOrderList Cancel Order list
Delete /api/v3/orderList

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#cancel-order-list-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param listClientOrderId -  A unique Id for the entire orderList
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiDeleteOrderListRequest
*/
func (a *TradeAPIService) DeleteOrderList(ctx context.Context) ApiDeleteOrderListRequest {
	return ApiDeleteOrderListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteOrderListResponse
func (a *TradeAPIService) DeleteOrderListExecute(r ApiDeleteOrderListRequest) (*common.RestApiResponse[models.DeleteOrderListResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DeleteOrderListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	timeInForce             *models.NewOrderTimeInForceParameter
	quantity                *float32
	quoteOrderQty           *float32
	price                   *float32
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int32
	stopPrice               *float32
	trailingDelta           *int64
	icebergQty              *float32
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	pegPriceType            *models.NewOrderPegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.NewOrderPegOffsetTypeParameter
	recvWindow              *float32
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewOrderRequest) Side(side models.NewOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiNewOrderRequest) QuoteOrderQty(quoteOrderQty float32) ApiNewOrderRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

func (r ApiNewOrderRequest) Price(price float32) ApiNewOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiNewOrderRequest) StrategyId(strategyId int64) ApiNewOrderRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiNewOrderRequest) StrategyType(strategyType int32) ApiNewOrderRequest {
	r.strategyType = &strategyType
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiNewOrderRequest) StopPrice(stopPrice float32) ApiNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiNewOrderRequest) TrailingDelta(trailingDelta int64) ApiNewOrderRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiNewOrderRequest) IcebergQty(icebergQty float32) ApiNewOrderRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiNewOrderRequest) PegPriceType(pegPriceType models.NewOrderPegPriceTypeParameter) ApiNewOrderRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info)
func (r ApiNewOrderRequest) PegOffsetValue(pegOffsetValue int32) ApiNewOrderRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiNewOrderRequest) PegOffsetType(pegOffsetType models.NewOrderPegOffsetTypeParameter) ApiNewOrderRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiNewOrderRequest) RecvWindow(recvWindow float32) ApiNewOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewOrderRequest) Execute() (*common.RestApiResponse[models.NewOrderResponse], error) {
	return r.ApiService.NewOrderExecute(r)
}

/*
NewOrder New order
Post /api/v3/order

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param timeInForce -
@param quantity -
@param quoteOrderQty -
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param strategyId -
@param strategyType -  The value cannot be less than `1000000`.
@param stopPrice -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param trailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -
@param selfTradePreventionMode -
@param pegPriceType -
@param pegOffsetValue -  Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info)
@param pegOffsetType -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiNewOrderRequest
*/
func (a *TradeAPIService) NewOrder(ctx context.Context) ApiNewOrderRequest {
	return ApiNewOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewOrderResponse
func (a *TradeAPIService) NewOrderExecute(r ApiNewOrderRequest) (*common.RestApiResponse[models.NewOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.quoteOrderQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteOrderQty", r.quoteOrderQty, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.strategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.trailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trailingDelta", r.trailingDelta, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.pegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegPriceType", r.pegPriceType, "form", "")
	}
	if r.pegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetValue", r.pegOffsetValue, "form", "")
	}
	if r.pegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetType", r.pegOffsetType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderAmendKeepPriorityRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	newQty            *float32
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

func (r ApiOrderAmendKeepPriorityRequest) OrderId(orderId int64) ApiOrderAmendKeepPriorityRequest {
	r.orderId = &orderId
	return r
}

func (r ApiOrderAmendKeepPriorityRequest) OrigClientOrderId(origClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiOrderAmendKeepPriorityRequest) NewClientOrderId(newClientOrderId string) ApiOrderAmendKeepPriorityRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderAmendKeepPriorityRequest) RecvWindow(recvWindow float32) ApiOrderAmendKeepPriorityRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderAmendKeepPriorityRequest) Execute() (*common.RestApiResponse[models.OrderAmendKeepPriorityResponse], error) {
	return r.ApiService.OrderAmendKeepPriorityExecute(r)
}

/*
OrderAmendKeepPriority Order Amend Keep Priority
Put /api/v3/order/amend/keepPriority

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#order-amend-keep-priority-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param newQty -  `newQty` must be greater than 0 and less than the order's quantity.
@param orderId -
@param origClientOrderId -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderAmendKeepPriorityRequest
*/
func (a *TradeAPIService) OrderAmendKeepPriority(ctx context.Context) ApiOrderAmendKeepPriorityRequest {
	return ApiOrderAmendKeepPriorityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderAmendKeepPriorityResponse
func (a *TradeAPIService) OrderAmendKeepPriorityExecute(r ApiOrderAmendKeepPriorityRequest) (*common.RestApiResponse[models.OrderAmendKeepPriorityResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/amend/keepPriority"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.newQty == nil {
		return nil, common.ReportError("newQty is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newQty", r.newQty, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderAmendKeepPriorityResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderCancelReplaceRequest struct {
	ctx                        context.Context
	ApiService                 *TradeAPIService
	symbol                     *string
	side                       *models.NewOrderSideParameter
	type_                      *models.NewOrderTypeParameter
	cancelReplaceMode          *models.OrderCancelReplaceCancelReplaceModeParameter
	timeInForce                *models.NewOrderTimeInForceParameter
	quantity                   *float32
	quoteOrderQty              *float32
	price                      *float32
	cancelNewClientOrderId     *string
	cancelOrigClientOrderId    *string
	cancelOrderId              *int64
	newClientOrderId           *string
	strategyId                 *int64
	strategyType               *int32
	stopPrice                  *float32
	trailingDelta              *int64
	icebergQty                 *float32
	newOrderRespType           *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode    *models.NewOrderSelfTradePreventionModeParameter
	cancelRestrictions         *models.DeleteOrderCancelRestrictionsParameter
	orderRateLimitExceededMode *models.OrderCancelReplaceOrderRateLimitExceededModeParameter
	pegPriceType               *models.NewOrderPegPriceTypeParameter
	pegOffsetValue             *int32
	pegOffsetType              *models.NewOrderPegOffsetTypeParameter
	recvWindow                 *float32
}

func (r ApiOrderCancelReplaceRequest) Symbol(symbol string) ApiOrderCancelReplaceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderCancelReplaceRequest) Side(side models.NewOrderSideParameter) ApiOrderCancelReplaceRequest {
	r.side = &side
	return r
}

func (r ApiOrderCancelReplaceRequest) Type(type_ models.NewOrderTypeParameter) ApiOrderCancelReplaceRequest {
	r.type_ = &type_
	return r
}

func (r ApiOrderCancelReplaceRequest) CancelReplaceMode(cancelReplaceMode models.OrderCancelReplaceCancelReplaceModeParameter) ApiOrderCancelReplaceRequest {
	r.cancelReplaceMode = &cancelReplaceMode
	return r
}

func (r ApiOrderCancelReplaceRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiOrderCancelReplaceRequest {
	r.timeInForce = &timeInForce
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

func (r ApiOrderCancelReplaceRequest) Price(price float32) ApiOrderCancelReplaceRequest {
	r.price = &price
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default.
func (r ApiOrderCancelReplaceRequest) CancelNewClientOrderId(cancelNewClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelNewClientOrderId = &cancelNewClientOrderId
	return r
}

// Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt; If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt; If both conditions are not met the request will be rejected.
func (r ApiOrderCancelReplaceRequest) CancelOrigClientOrderId(cancelOrigClientOrderId string) ApiOrderCancelReplaceRequest {
	r.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return r
}

// Either &#x60;cancelOrderId&#x60; or &#x60;cancelOrigClientOrderId&#x60; must be sent. &lt;br&gt;&lt;/br&gt;If both &#x60;cancelOrderId&#x60; and &#x60;cancelOrigClientOrderId&#x60; parameters are provided, the &#x60;cancelOrderId&#x60; is searched first, then the &#x60;cancelOrigClientOrderId&#x60; from that result is checked against that order. &lt;br&gt;&lt;/br&gt;If both conditions are not met the request will be rejected.
func (r ApiOrderCancelReplaceRequest) CancelOrderId(cancelOrderId int64) ApiOrderCancelReplaceRequest {
	r.cancelOrderId = &cancelOrderId
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiOrderCancelReplaceRequest) NewClientOrderId(newClientOrderId string) ApiOrderCancelReplaceRequest {
	r.newClientOrderId = &newClientOrderId
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

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiOrderCancelReplaceRequest) StopPrice(stopPrice float32) ApiOrderCancelReplaceRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderCancelReplaceRequest) TrailingDelta(trailingDelta int64) ApiOrderCancelReplaceRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiOrderCancelReplaceRequest) IcebergQty(icebergQty float32) ApiOrderCancelReplaceRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiOrderCancelReplaceRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderCancelReplaceRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderCancelReplaceRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderCancelReplaceRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiOrderCancelReplaceRequest) CancelRestrictions(cancelRestrictions models.DeleteOrderCancelRestrictionsParameter) ApiOrderCancelReplaceRequest {
	r.cancelRestrictions = &cancelRestrictions
	return r
}

func (r ApiOrderCancelReplaceRequest) OrderRateLimitExceededMode(orderRateLimitExceededMode models.OrderCancelReplaceOrderRateLimitExceededModeParameter) ApiOrderCancelReplaceRequest {
	r.orderRateLimitExceededMode = &orderRateLimitExceededMode
	return r
}

func (r ApiOrderCancelReplaceRequest) PegPriceType(pegPriceType models.NewOrderPegPriceTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info)
func (r ApiOrderCancelReplaceRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderCancelReplaceRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiOrderCancelReplaceRequest) PegOffsetType(pegOffsetType models.NewOrderPegOffsetTypeParameter) ApiOrderCancelReplaceRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderCancelReplaceRequest) RecvWindow(recvWindow float32) ApiOrderCancelReplaceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderCancelReplaceRequest) Execute() (*common.RestApiResponse[models.OrderCancelReplaceResponse], error) {
	return r.ApiService.OrderCancelReplaceExecute(r)
}

/*
OrderCancelReplace Cancel an Existing Order and Send a New Order
Post /api/v3/order/cancelReplace

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#cancel-an-existing-order-and-send-a-new-order-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param cancelReplaceMode -
@param timeInForce -
@param quantity -
@param quoteOrderQty -
@param price -
@param cancelNewClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default.
@param cancelOrigClientOrderId -  Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br> If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br> If both conditions are not met the request will be rejected.
@param cancelOrderId -  Either `cancelOrderId` or `cancelOrigClientOrderId` must be sent. <br></br>If both `cancelOrderId` and `cancelOrigClientOrderId` parameters are provided, the `cancelOrderId` is searched first, then the `cancelOrigClientOrderId` from that result is checked against that order. <br></br>If both conditions are not met the request will be rejected.
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param strategyId -
@param strategyType -  The value cannot be less than `1000000`.
@param stopPrice -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param trailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -
@param selfTradePreventionMode -
@param cancelRestrictions -
@param orderRateLimitExceededMode -
@param pegPriceType -
@param pegOffsetValue -  Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info)
@param pegOffsetType -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderCancelReplaceRequest
*/
func (a *TradeAPIService) OrderCancelReplace(ctx context.Context) ApiOrderCancelReplaceRequest {
	return ApiOrderCancelReplaceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderCancelReplaceResponse
func (a *TradeAPIService) OrderCancelReplaceExecute(r ApiOrderCancelReplaceRequest) (*common.RestApiResponse[models.OrderCancelReplaceResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/cancelReplace"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}
	if r.cancelReplaceMode == nil {
		return nil, common.ReportError("cancelReplaceMode is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelReplaceMode", r.cancelReplaceMode, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.quoteOrderQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteOrderQty", r.quoteOrderQty, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.cancelNewClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelNewClientOrderId", r.cancelNewClientOrderId, "form", "")
	}
	if r.cancelOrigClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelOrigClientOrderId", r.cancelOrigClientOrderId, "form", "")
	}
	if r.cancelOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelOrderId", r.cancelOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.strategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.trailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trailingDelta", r.trailingDelta, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.cancelRestrictions != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "cancelRestrictions", r.cancelRestrictions, "form", "")
	}
	if r.orderRateLimitExceededMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderRateLimitExceededMode", r.orderRateLimitExceededMode, "form", "")
	}
	if r.pegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegPriceType", r.pegPriceType, "form", "")
	}
	if r.pegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetValue", r.pegOffsetValue, "form", "")
	}
	if r.pegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetType", r.pegOffsetType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderCancelReplaceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderListOcoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	quantity                *float32
	aboveType               *models.OrderListOcoAboveTypeParameter
	belowType               *models.OrderListOcoBelowTypeParameter
	listClientOrderId       *string
	aboveClientOrderId      *string
	aboveIcebergQty         *int64
	abovePrice              *float32
	aboveStopPrice          *float32
	aboveTrailingDelta      *int64
	aboveTimeInForce        *models.OrderOcoStopLimitTimeInForceParameter
	aboveStrategyId         *int64
	aboveStrategyType       *int32
	abovePegPriceType       *models.OrderListOcoAbovePegPriceTypeParameter
	abovePegOffsetType      *models.OrderListOcoAbovePegOffsetTypeParameter
	abovePegOffsetValue     *int32
	belowClientOrderId      *string
	belowIcebergQty         *int64
	belowPrice              *float32
	belowStopPrice          *float32
	belowTrailingDelta      *int64
	belowTimeInForce        *models.OrderOcoStopLimitTimeInForceParameter
	belowStrategyId         *int64
	belowStrategyType       *int32
	belowPegPriceType       *models.OrderListOcoAbovePegPriceTypeParameter
	belowPegOffsetType      *models.OrderListOcoAbovePegOffsetTypeParameter
	belowPegOffsetValue     *int32
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiOrderListOcoRequest) Symbol(symbol string) ApiOrderListOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListOcoRequest) Side(side models.NewOrderSideParameter) ApiOrderListOcoRequest {
	r.side = &side
	return r
}

func (r ApiOrderListOcoRequest) Quantity(quantity float32) ApiOrderListOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderListOcoRequest) AboveType(aboveType models.OrderListOcoAboveTypeParameter) ApiOrderListOcoRequest {
	r.aboveType = &aboveType
	return r
}

func (r ApiOrderListOcoRequest) BelowType(belowType models.OrderListOcoBelowTypeParameter) ApiOrderListOcoRequest {
	r.belowType = &belowType
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListOcoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
func (r ApiOrderListOcoRequest) AboveClientOrderId(aboveClientOrderId string) ApiOrderListOcoRequest {
	r.aboveClientOrderId = &aboveClientOrderId
	return r
}

// Note that this can only be used if &#x60;aboveTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListOcoRequest) AboveIcebergQty(aboveIcebergQty int64) ApiOrderListOcoRequest {
	r.aboveIcebergQty = &aboveIcebergQty
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListOcoRequest) AbovePrice(abovePrice float32) ApiOrderListOcoRequest {
	r.abovePrice = &abovePrice
	return r
}

// Can be used if &#x60;aboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;. &lt;br&gt;Either &#x60;aboveStopPrice&#x60; or &#x60;aboveTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListOcoRequest) AboveStopPrice(aboveStopPrice float32) ApiOrderListOcoRequest {
	r.aboveStopPrice = &aboveStopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderListOcoRequest) AboveTrailingDelta(aboveTrailingDelta int64) ApiOrderListOcoRequest {
	r.aboveTrailingDelta = &aboveTrailingDelta
	return r
}

func (r ApiOrderListOcoRequest) AboveTimeInForce(aboveTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOcoRequest {
	r.aboveTimeInForce = &aboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the above order within an order strategy.
func (r ApiOrderListOcoRequest) AboveStrategyId(aboveStrategyId int64) ApiOrderListOcoRequest {
	r.aboveStrategyId = &aboveStrategyId
	return r
}

// Arbitrary numeric value identifying the above order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOcoRequest) AboveStrategyType(aboveStrategyType int32) ApiOrderListOcoRequest {
	r.aboveStrategyType = &aboveStrategyType
	return r
}

func (r ApiOrderListOcoRequest) AbovePegPriceType(abovePegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOcoRequest {
	r.abovePegPriceType = &abovePegPriceType
	return r
}

func (r ApiOrderListOcoRequest) AbovePegOffsetType(abovePegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOcoRequest {
	r.abovePegOffsetType = &abovePegOffsetType
	return r
}

func (r ApiOrderListOcoRequest) AbovePegOffsetValue(abovePegOffsetValue int32) ApiOrderListOcoRequest {
	r.abovePegOffsetValue = &abovePegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the below order. Automatically generated if not sent
func (r ApiOrderListOcoRequest) BelowClientOrderId(belowClientOrderId string) ApiOrderListOcoRequest {
	r.belowClientOrderId = &belowClientOrderId
	return r
}

// Note that this can only be used if &#x60;belowTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiOrderListOcoRequest) BelowIcebergQty(belowIcebergQty int64) ApiOrderListOcoRequest {
	r.belowIcebergQty = &belowIcebergQty
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListOcoRequest) BelowPrice(belowPrice float32) ApiOrderListOcoRequest {
	r.belowPrice = &belowPrice
	return r
}

// Can be used if &#x60;belowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; &lt;br&gt;Either belowStopPrice or belowTrailingDelta or both, must be specified.
func (r ApiOrderListOcoRequest) BelowStopPrice(belowStopPrice float32) ApiOrderListOcoRequest {
	r.belowStopPrice = &belowStopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderListOcoRequest) BelowTrailingDelta(belowTrailingDelta int64) ApiOrderListOcoRequest {
	r.belowTrailingDelta = &belowTrailingDelta
	return r
}

func (r ApiOrderListOcoRequest) BelowTimeInForce(belowTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOcoRequest {
	r.belowTimeInForce = &belowTimeInForce
	return r
}

// Arbitrary numeric value identifying the below order within an order strategy.
func (r ApiOrderListOcoRequest) BelowStrategyId(belowStrategyId int64) ApiOrderListOcoRequest {
	r.belowStrategyId = &belowStrategyId
	return r
}

// Arbitrary numeric value identifying the below order strategy. &lt;br&gt;Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOcoRequest) BelowStrategyType(belowStrategyType int32) ApiOrderListOcoRequest {
	r.belowStrategyType = &belowStrategyType
	return r
}

func (r ApiOrderListOcoRequest) BelowPegPriceType(belowPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOcoRequest {
	r.belowPegPriceType = &belowPegPriceType
	return r
}

func (r ApiOrderListOcoRequest) BelowPegOffsetType(belowPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOcoRequest {
	r.belowPegOffsetType = &belowPegOffsetType
	return r
}

func (r ApiOrderListOcoRequest) BelowPegOffsetValue(belowPegOffsetValue int32) ApiOrderListOcoRequest {
	r.belowPegOffsetValue = &belowPegOffsetValue
	return r
}

func (r ApiOrderListOcoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderListOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListOcoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderListOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListOcoRequest) RecvWindow(recvWindow float32) ApiOrderListOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListOcoRequest) Execute() (*common.RestApiResponse[models.OrderListOcoResponse], error) {
	return r.ApiService.OrderListOcoExecute(r)
}

/*
OrderListOco New Order list - OCO
Post /api/v3/orderList/oco

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-list---oco-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -
@param aboveType -
@param belowType -
@param listClientOrderId -  A unique Id for the entire orderList
@param aboveClientOrderId -  Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
@param aboveIcebergQty -  Note that this can only be used if `aboveTimeInForce` is `GTC`.
@param abovePrice -  Can be used if `aboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.
@param aboveStopPrice -  Can be used if `aboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`. <br>Either `aboveStopPrice` or `aboveTrailingDelta` or both, must be specified.
@param aboveTrailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
@param aboveTimeInForce -
@param aboveStrategyId -  Arbitrary numeric value identifying the above order within an order strategy.
@param aboveStrategyType -  Arbitrary numeric value identifying the above order strategy. <br>Values smaller than 1000000 are reserved and cannot be used.
@param abovePegPriceType -
@param abovePegOffsetType -
@param abovePegOffsetValue -
@param belowClientOrderId -  Arbitrary unique ID among open orders for the below order. Automatically generated if not sent
@param belowIcebergQty -  Note that this can only be used if `belowTimeInForce` is `GTC`.
@param belowPrice -  Can be used if `belowType` is `STOP_LOSS_LIMIT`, `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.
@param belowStopPrice -  Can be used if `belowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT` or `TAKE_PROFIT_LIMIT` <br>Either belowStopPrice or belowTrailingDelta or both, must be specified.
@param belowTrailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
@param belowTimeInForce -
@param belowStrategyId -  Arbitrary numeric value identifying the below order within an order strategy.
@param belowStrategyType -  Arbitrary numeric value identifying the below order strategy. <br>Values smaller than 1000000 are reserved and cannot be used.
@param belowPegPriceType -
@param belowPegOffsetType -
@param belowPegOffsetValue -
@param newOrderRespType -
@param selfTradePreventionMode -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListOcoRequest
*/
func (a *TradeAPIService) OrderListOco(ctx context.Context) ApiOrderListOcoRequest {
	return ApiOrderListOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderListOcoResponse
func (a *TradeAPIService) OrderListOcoExecute(r ApiOrderListOcoRequest) (*common.RestApiResponse[models.OrderListOcoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList/oco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}
	if r.aboveType == nil {
		return nil, common.ReportError("aboveType is required and must be specified")
	}
	if r.belowType == nil {
		return nil, common.ReportError("belowType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveType", r.aboveType, "form", "")
	if r.aboveClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveClientOrderId", r.aboveClientOrderId, "form", "")
	}
	if r.aboveIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveIcebergQty", r.aboveIcebergQty, "form", "")
	}
	if r.abovePrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "abovePrice", r.abovePrice, "form", "")
	}
	if r.aboveStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveStopPrice", r.aboveStopPrice, "form", "")
	}
	if r.aboveTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveTrailingDelta", r.aboveTrailingDelta, "form", "")
	}
	if r.aboveTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveTimeInForce", r.aboveTimeInForce, "form", "")
	}
	if r.aboveStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveStrategyId", r.aboveStrategyId, "form", "")
	}
	if r.aboveStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aboveStrategyType", r.aboveStrategyType, "form", "")
	}
	if r.abovePegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "abovePegPriceType", r.abovePegPriceType, "form", "")
	}
	if r.abovePegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "abovePegOffsetType", r.abovePegOffsetType, "form", "")
	}
	if r.abovePegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "abovePegOffsetValue", r.abovePegOffsetValue, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowType", r.belowType, "form", "")
	if r.belowClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowClientOrderId", r.belowClientOrderId, "form", "")
	}
	if r.belowIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowIcebergQty", r.belowIcebergQty, "form", "")
	}
	if r.belowPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowPrice", r.belowPrice, "form", "")
	}
	if r.belowStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowStopPrice", r.belowStopPrice, "form", "")
	}
	if r.belowTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowTrailingDelta", r.belowTrailingDelta, "form", "")
	}
	if r.belowTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowTimeInForce", r.belowTimeInForce, "form", "")
	}
	if r.belowStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowStrategyId", r.belowStrategyId, "form", "")
	}
	if r.belowStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowStrategyType", r.belowStrategyType, "form", "")
	}
	if r.belowPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowPegPriceType", r.belowPegPriceType, "form", "")
	}
	if r.belowPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowPegOffsetType", r.belowPegOffsetType, "form", "")
	}
	if r.belowPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "belowPegOffsetValue", r.belowPegOffsetValue, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderListOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderListOpoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *models.OrderListOpoWorkingTypeParameter
	workingSide             *models.NewOrderSideParameter
	workingPrice            *float32
	workingQuantity         *float32
	pendingType             *models.OrderListOpoPendingTypeParameter
	pendingSide             *models.NewOrderSideParameter
	listClientOrderId       *string
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float32
	workingTimeInForce      *models.OrderOcoStopLimitTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderListOcoAbovePegPriceTypeParameter
	workingPegOffsetType    *models.OrderListOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float32
	pendingStopPrice        *float32
	pendingTrailingDelta    *float32
	pendingIcebergQty       *float32
	pendingTimeInForce      *models.OrderOcoStopLimitTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegPriceType     *models.OrderListOcoAbovePegPriceTypeParameter
	pendingPegOffsetType    *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float32
}

func (r ApiOrderListOpoRequest) Symbol(symbol string) ApiOrderListOpoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListOpoRequest) WorkingType(workingType models.OrderListOpoWorkingTypeParameter) ApiOrderListOpoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListOpoRequest) WorkingSide(workingSide models.NewOrderSideParameter) ApiOrderListOpoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListOpoRequest) WorkingPrice(workingPrice float32) ApiOrderListOpoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListOpoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListOpoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListOpoRequest) PendingType(pendingType models.OrderListOpoPendingTypeParameter) ApiOrderListOpoRequest {
	r.pendingType = &pendingType
	return r
}

func (r ApiOrderListOpoRequest) PendingSide(pendingSide models.NewOrderSideParameter) ApiOrderListOpoRequest {
	r.pendingSide = &pendingSide
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListOpoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListOpoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListOpoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderListOpoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListOpoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderListOpoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListOpoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListOpoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOpoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListOpoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListOpoRequest) WorkingTimeInForce(workingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOpoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListOpoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListOpoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOpoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListOpoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListOpoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOpoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListOpoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOpoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListOpoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListOpoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiOrderListOpoRequest) PendingClientOrderId(pendingClientOrderId string) ApiOrderListOpoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiOrderListOpoRequest) PendingPrice(pendingPrice float32) ApiOrderListOpoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiOrderListOpoRequest) PendingStopPrice(pendingStopPrice float32) ApiOrderListOpoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiOrderListOpoRequest) PendingTrailingDelta(pendingTrailingDelta float32) ApiOrderListOpoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOpoRequest) PendingIcebergQty(pendingIcebergQty float32) ApiOrderListOpoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

func (r ApiOrderListOpoRequest) PendingTimeInForce(pendingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOpoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListOpoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListOpoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOpoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListOpoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

func (r ApiOrderListOpoRequest) PendingPegPriceType(pendingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOpoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListOpoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOpoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

func (r ApiOrderListOpoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListOpoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListOpoRequest) RecvWindow(recvWindow float32) ApiOrderListOpoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListOpoRequest) Execute() (*common.RestApiResponse[models.OrderListOpoResponse], error) {
	return r.ApiService.OrderListOpoExecute(r)
}

/*
OrderListOpo New Order List - OPO
Post /api/v3/orderList/opo

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-list---opo-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -  Sets the quantity for the working order.
@param pendingType -
@param pendingSide -
@param listClientOrderId -  A unique Id for the entire orderList
@param newOrderRespType -
@param selfTradePreventionMode -
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.
@param workingTimeInForce -
@param workingStrategyId -  Arbitrary numeric value identifying the working order within an order strategy.
@param workingStrategyType -  Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param workingPegPriceType -
@param workingPegOffsetType -
@param workingPegOffsetValue -
@param pendingClientOrderId -  Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
@param pendingPrice -
@param pendingStopPrice -
@param pendingTrailingDelta -
@param pendingIcebergQty -  This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.
@param pendingTimeInForce -
@param pendingStrategyId -  Arbitrary numeric value identifying the pending order within an order strategy.
@param pendingStrategyType -  Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingPegPriceType -
@param pendingPegOffsetType -
@param pendingPegOffsetValue -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListOpoRequest
*/
func (a *TradeAPIService) OrderListOpo(ctx context.Context) ApiOrderListOpoRequest {
	return ApiOrderListOpoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderListOpoResponse
func (a *TradeAPIService) OrderListOpoExecute(r ApiOrderListOpoRequest) (*common.RestApiResponse[models.OrderListOpoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList/opo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}
	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}
	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}
	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}
	if r.pendingType == nil {
		return nil, common.ReportError("pendingType is required and must be specified")
	}
	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	if r.workingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	}
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	if r.workingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyId", r.workingStrategyId, "form", "")
	}
	if r.workingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyType", r.workingStrategyType, "form", "")
	}
	if r.workingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegPriceType", r.workingPegPriceType, "form", "")
	}
	if r.workingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetType", r.workingPegOffsetType, "form", "")
	}
	if r.workingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetValue", r.workingPegOffsetValue, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingType", r.pendingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	if r.pendingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingClientOrderId", r.pendingClientOrderId, "form", "")
	}
	if r.pendingPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPrice", r.pendingPrice, "form", "")
	}
	if r.pendingStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStopPrice", r.pendingStopPrice, "form", "")
	}
	if r.pendingTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTrailingDelta", r.pendingTrailingDelta, "form", "")
	}
	if r.pendingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingIcebergQty", r.pendingIcebergQty, "form", "")
	}
	if r.pendingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTimeInForce", r.pendingTimeInForce, "form", "")
	}
	if r.pendingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStrategyId", r.pendingStrategyId, "form", "")
	}
	if r.pendingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStrategyType", r.pendingStrategyType, "form", "")
	}
	if r.pendingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegPriceType", r.pendingPegPriceType, "form", "")
	}
	if r.pendingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegOffsetType", r.pendingPegOffsetType, "form", "")
	}
	if r.pendingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegOffsetValue", r.pendingPegOffsetValue, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderListOpoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderListOpocoRequest struct {
	ctx                        context.Context
	ApiService                 *TradeAPIService
	symbol                     *string
	workingType                *models.OrderListOpoWorkingTypeParameter
	workingSide                *models.NewOrderSideParameter
	workingPrice               *float32
	workingQuantity            *float32
	pendingSide                *models.NewOrderSideParameter
	pendingAboveType           *models.OrderListOcoAboveTypeParameter
	listClientOrderId          *string
	newOrderRespType           *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode    *models.NewOrderSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float32
	workingTimeInForce         *models.OrderOcoStopLimitTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderListOcoAbovePegPriceTypeParameter
	workingPegOffsetType       *models.OrderListOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float32
	pendingAboveStopPrice      *float32
	pendingAboveTrailingDelta  *float32
	pendingAboveIcebergQty     *float32
	pendingAboveTimeInForce    *models.OrderOcoStopLimitTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderListOcoAbovePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float32
	pendingBelowStopPrice      *float32
	pendingBelowTrailingDelta  *float32
	pendingBelowIcebergQty     *float32
	pendingBelowTimeInForce    *models.OrderOcoStopLimitTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderListOcoAbovePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float32
}

func (r ApiOrderListOpocoRequest) Symbol(symbol string) ApiOrderListOpocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListOpocoRequest) WorkingType(workingType models.OrderListOpoWorkingTypeParameter) ApiOrderListOpocoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListOpocoRequest) WorkingSide(workingSide models.NewOrderSideParameter) ApiOrderListOpocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListOpocoRequest) WorkingPrice(workingPrice float32) ApiOrderListOpocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListOpocoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListOpocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListOpocoRequest) PendingSide(pendingSide models.NewOrderSideParameter) ApiOrderListOpocoRequest {
	r.pendingSide = &pendingSide
	return r
}

func (r ApiOrderListOpocoRequest) PendingAboveType(pendingAboveType models.OrderListOcoAboveTypeParameter) ApiOrderListOpocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListOpocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListOpocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListOpocoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderListOpocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListOpocoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderListOpocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListOpocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListOpocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOpocoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListOpocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListOpocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOpocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListOpocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListOpocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOpocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListOpocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListOpocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOpocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListOpocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOpocoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListOpocoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListOpocoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiOrderListOpocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiOrderListOpocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListOpocoRequest) PendingAbovePrice(pendingAbovePrice float32) ApiOrderListOpocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListOpocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float32) ApiOrderListOpocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
func (r ApiOrderListOpocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float32) ApiOrderListOpocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOpocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float32) ApiOrderListOpocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiOrderListOpocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOpocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListOpocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListOpocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOpocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListOpocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

func (r ApiOrderListOpocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOpocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

func (r ApiOrderListOpocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOpocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

func (r ApiOrderListOpocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListOpocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowType(pendingBelowType models.OrderListOcoBelowTypeParameter) ApiOrderListOpocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListOpocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListOpocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price
func (r ApiOrderListOpocoRequest) PendingBelowPrice(pendingBelowPrice float32) ApiOrderListOpocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListOpocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float32) ApiOrderListOpocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float32) ApiOrderListOpocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOpocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float32) ApiOrderListOpocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOpocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListOpocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListOpocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOpocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListOpocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOpocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOpocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListOpocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListOpocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListOpocoRequest) RecvWindow(recvWindow float32) ApiOrderListOpocoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListOpocoRequest) Execute() (*common.RestApiResponse[models.OrderListOpocoResponse], error) {
	return r.ApiService.OrderListOpocoExecute(r)
}

/*
OrderListOpoco New Order List - OPOCO
Post /api/v3/orderList/opoco

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-list---opoco-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -  Sets the quantity for the working order.
@param pendingSide -
@param pendingAboveType -
@param listClientOrderId -  A unique Id for the entire orderList
@param newOrderRespType -
@param selfTradePreventionMode -
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.
@param workingTimeInForce -
@param workingStrategyId -  Arbitrary numeric value identifying the working order within an order strategy.
@param workingStrategyType -  Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param workingPegPriceType -
@param workingPegOffsetType -
@param workingPegOffsetValue -
@param pendingAboveClientOrderId -  Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
@param pendingAbovePrice -  Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.
@param pendingAboveStopPrice -  Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`
@param pendingAboveTrailingDelta -  See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
@param pendingAboveIcebergQty -  This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.
@param pendingAboveTimeInForce -
@param pendingAboveStrategyId -  Arbitrary numeric value identifying the pending above order within an order strategy.
@param pendingAboveStrategyType -  Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingAbovePegPriceType -
@param pendingAbovePegOffsetType -
@param pendingAbovePegOffsetValue -
@param pendingBelowType -
@param pendingBelowClientOrderId -  Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
@param pendingBelowPrice -  Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price
@param pendingBelowStopPrice -  Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.
@param pendingBelowTrailingDelta -
@param pendingBelowIcebergQty -  This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.
@param pendingBelowTimeInForce -
@param pendingBelowStrategyId -  Arbitrary numeric value identifying the pending below order within an order strategy.
@param pendingBelowStrategyType -  Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingBelowPegPriceType -
@param pendingBelowPegOffsetType -
@param pendingBelowPegOffsetValue -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListOpocoRequest
*/
func (a *TradeAPIService) OrderListOpoco(ctx context.Context) ApiOrderListOpocoRequest {
	return ApiOrderListOpocoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderListOpocoResponse
func (a *TradeAPIService) OrderListOpocoExecute(r ApiOrderListOpocoRequest) (*common.RestApiResponse[models.OrderListOpocoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList/opoco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}
	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}
	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}
	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}
	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}
	if r.pendingAboveType == nil {
		return nil, common.ReportError("pendingAboveType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	if r.workingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	}
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	if r.workingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyId", r.workingStrategyId, "form", "")
	}
	if r.workingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyType", r.workingStrategyType, "form", "")
	}
	if r.workingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegPriceType", r.workingPegPriceType, "form", "")
	}
	if r.workingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetType", r.workingPegOffsetType, "form", "")
	}
	if r.workingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetValue", r.workingPegOffsetValue, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveType", r.pendingAboveType, "form", "")
	if r.pendingAboveClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveClientOrderId", r.pendingAboveClientOrderId, "form", "")
	}
	if r.pendingAbovePrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePrice", r.pendingAbovePrice, "form", "")
	}
	if r.pendingAboveStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStopPrice", r.pendingAboveStopPrice, "form", "")
	}
	if r.pendingAboveTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTrailingDelta", r.pendingAboveTrailingDelta, "form", "")
	}
	if r.pendingAboveIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveIcebergQty", r.pendingAboveIcebergQty, "form", "")
	}
	if r.pendingAboveTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTimeInForce", r.pendingAboveTimeInForce, "form", "")
	}
	if r.pendingAboveStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStrategyId", r.pendingAboveStrategyId, "form", "")
	}
	if r.pendingAboveStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStrategyType", r.pendingAboveStrategyType, "form", "")
	}
	if r.pendingAbovePegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegPriceType", r.pendingAbovePegPriceType, "form", "")
	}
	if r.pendingAbovePegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegOffsetType", r.pendingAbovePegOffsetType, "form", "")
	}
	if r.pendingAbovePegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegOffsetValue", r.pendingAbovePegOffsetValue, "form", "")
	}
	if r.pendingBelowType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowType", r.pendingBelowType, "form", "")
	}
	if r.pendingBelowClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowClientOrderId", r.pendingBelowClientOrderId, "form", "")
	}
	if r.pendingBelowPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPrice", r.pendingBelowPrice, "form", "")
	}
	if r.pendingBelowStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStopPrice", r.pendingBelowStopPrice, "form", "")
	}
	if r.pendingBelowTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTrailingDelta", r.pendingBelowTrailingDelta, "form", "")
	}
	if r.pendingBelowIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowIcebergQty", r.pendingBelowIcebergQty, "form", "")
	}
	if r.pendingBelowTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTimeInForce", r.pendingBelowTimeInForce, "form", "")
	}
	if r.pendingBelowStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStrategyId", r.pendingBelowStrategyId, "form", "")
	}
	if r.pendingBelowStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStrategyType", r.pendingBelowStrategyType, "form", "")
	}
	if r.pendingBelowPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegPriceType", r.pendingBelowPegPriceType, "form", "")
	}
	if r.pendingBelowPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegOffsetType", r.pendingBelowPegOffsetType, "form", "")
	}
	if r.pendingBelowPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegOffsetValue", r.pendingBelowPegOffsetValue, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderListOpocoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderListOtoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *models.OrderListOpoWorkingTypeParameter
	workingSide             *models.NewOrderSideParameter
	workingPrice            *float32
	workingQuantity         *float32
	pendingType             *models.OrderListOpoPendingTypeParameter
	pendingSide             *models.NewOrderSideParameter
	pendingQuantity         *float32
	listClientOrderId       *string
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	workingClientOrderId    *string
	workingIcebergQty       *float32
	workingTimeInForce      *models.OrderOcoStopLimitTimeInForceParameter
	workingStrategyId       *int64
	workingStrategyType     *int32
	workingPegPriceType     *models.OrderListOcoAbovePegPriceTypeParameter
	workingPegOffsetType    *models.OrderListOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue   *int32
	pendingClientOrderId    *string
	pendingPrice            *float32
	pendingStopPrice        *float32
	pendingTrailingDelta    *float32
	pendingIcebergQty       *float32
	pendingTimeInForce      *models.OrderOcoStopLimitTimeInForceParameter
	pendingStrategyId       *int64
	pendingStrategyType     *int32
	pendingPegPriceType     *models.OrderListOcoAbovePegPriceTypeParameter
	pendingPegOffsetType    *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingPegOffsetValue   *int32
	recvWindow              *float32
}

func (r ApiOrderListOtoRequest) Symbol(symbol string) ApiOrderListOtoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListOtoRequest) WorkingType(workingType models.OrderListOpoWorkingTypeParameter) ApiOrderListOtoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListOtoRequest) WorkingSide(workingSide models.NewOrderSideParameter) ApiOrderListOtoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListOtoRequest) WorkingPrice(workingPrice float32) ApiOrderListOtoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListOtoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListOtoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListOtoRequest) PendingType(pendingType models.OrderListOpoPendingTypeParameter) ApiOrderListOtoRequest {
	r.pendingType = &pendingType
	return r
}

func (r ApiOrderListOtoRequest) PendingSide(pendingSide models.NewOrderSideParameter) ApiOrderListOtoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiOrderListOtoRequest) PendingQuantity(pendingQuantity float32) ApiOrderListOtoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListOtoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListOtoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListOtoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderListOtoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListOtoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderListOtoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListOtoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListOtoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOtoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListOtoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListOtoRequest) WorkingTimeInForce(workingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOtoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListOtoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListOtoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOtoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListOtoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListOtoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOtoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListOtoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOtoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListOtoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListOtoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiOrderListOtoRequest) PendingClientOrderId(pendingClientOrderId string) ApiOrderListOtoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiOrderListOtoRequest) PendingPrice(pendingPrice float32) ApiOrderListOtoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiOrderListOtoRequest) PendingStopPrice(pendingStopPrice float32) ApiOrderListOtoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiOrderListOtoRequest) PendingTrailingDelta(pendingTrailingDelta float32) ApiOrderListOtoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOtoRequest) PendingIcebergQty(pendingIcebergQty float32) ApiOrderListOtoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

func (r ApiOrderListOtoRequest) PendingTimeInForce(pendingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOtoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending order within an order strategy.
func (r ApiOrderListOtoRequest) PendingStrategyId(pendingStrategyId int64) ApiOrderListOtoRequest {
	r.pendingStrategyId = &pendingStrategyId
	return r
}

// Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOtoRequest) PendingStrategyType(pendingStrategyType int32) ApiOrderListOtoRequest {
	r.pendingStrategyType = &pendingStrategyType
	return r
}

func (r ApiOrderListOtoRequest) PendingPegPriceType(pendingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOtoRequest {
	r.pendingPegPriceType = &pendingPegPriceType
	return r
}

func (r ApiOrderListOtoRequest) PendingPegOffsetType(pendingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOtoRequest {
	r.pendingPegOffsetType = &pendingPegOffsetType
	return r
}

func (r ApiOrderListOtoRequest) PendingPegOffsetValue(pendingPegOffsetValue int32) ApiOrderListOtoRequest {
	r.pendingPegOffsetValue = &pendingPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListOtoRequest) RecvWindow(recvWindow float32) ApiOrderListOtoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListOtoRequest) Execute() (*common.RestApiResponse[models.OrderListOtoResponse], error) {
	return r.ApiService.OrderListOtoExecute(r)
}

/*
OrderListOto New Order list - OTO
Post /api/v3/orderList/oto

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-list---oto-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -  Sets the quantity for the working order.
@param pendingType -
@param pendingSide -
@param pendingQuantity -  Sets the quantity for the pending order.
@param listClientOrderId -  A unique Id for the entire orderList
@param newOrderRespType -
@param selfTradePreventionMode -
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.
@param workingTimeInForce -
@param workingStrategyId -  Arbitrary numeric value identifying the working order within an order strategy.
@param workingStrategyType -  Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param workingPegPriceType -
@param workingPegOffsetType -
@param workingPegOffsetValue -
@param pendingClientOrderId -  Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
@param pendingPrice -
@param pendingStopPrice -
@param pendingTrailingDelta -
@param pendingIcebergQty -  This can only be used if `pendingTimeInForce` is `GTC` or if `pendingType` is `LIMIT_MAKER`.
@param pendingTimeInForce -
@param pendingStrategyId -  Arbitrary numeric value identifying the pending order within an order strategy.
@param pendingStrategyType -  Arbitrary numeric value identifying the pending order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingPegPriceType -
@param pendingPegOffsetType -
@param pendingPegOffsetValue -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListOtoRequest
*/
func (a *TradeAPIService) OrderListOto(ctx context.Context) ApiOrderListOtoRequest {
	return ApiOrderListOtoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderListOtoResponse
func (a *TradeAPIService) OrderListOtoExecute(r ApiOrderListOtoRequest) (*common.RestApiResponse[models.OrderListOtoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList/oto"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}
	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}
	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}
	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}
	if r.pendingType == nil {
		return nil, common.ReportError("pendingType is required and must be specified")
	}
	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}
	if r.pendingQuantity == nil {
		return nil, common.ReportError("pendingQuantity is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	if r.workingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	}
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	if r.workingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyId", r.workingStrategyId, "form", "")
	}
	if r.workingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyType", r.workingStrategyType, "form", "")
	}
	if r.workingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegPriceType", r.workingPegPriceType, "form", "")
	}
	if r.workingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetType", r.workingPegOffsetType, "form", "")
	}
	if r.workingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetValue", r.workingPegOffsetValue, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingType", r.pendingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	if r.pendingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingClientOrderId", r.pendingClientOrderId, "form", "")
	}
	if r.pendingPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPrice", r.pendingPrice, "form", "")
	}
	if r.pendingStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStopPrice", r.pendingStopPrice, "form", "")
	}
	if r.pendingTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTrailingDelta", r.pendingTrailingDelta, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingQuantity", r.pendingQuantity, "form", "")
	if r.pendingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingIcebergQty", r.pendingIcebergQty, "form", "")
	}
	if r.pendingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTimeInForce", r.pendingTimeInForce, "form", "")
	}
	if r.pendingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStrategyId", r.pendingStrategyId, "form", "")
	}
	if r.pendingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStrategyType", r.pendingStrategyType, "form", "")
	}
	if r.pendingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegPriceType", r.pendingPegPriceType, "form", "")
	}
	if r.pendingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegOffsetType", r.pendingPegOffsetType, "form", "")
	}
	if r.pendingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPegOffsetValue", r.pendingPegOffsetValue, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderListOtoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderListOtocoRequest struct {
	ctx                        context.Context
	ApiService                 *TradeAPIService
	symbol                     *string
	workingType                *models.OrderListOpoWorkingTypeParameter
	workingSide                *models.NewOrderSideParameter
	workingPrice               *float32
	workingQuantity            *float32
	pendingSide                *models.NewOrderSideParameter
	pendingQuantity            *float32
	pendingAboveType           *models.OrderListOcoAboveTypeParameter
	listClientOrderId          *string
	newOrderRespType           *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode    *models.NewOrderSelfTradePreventionModeParameter
	workingClientOrderId       *string
	workingIcebergQty          *float32
	workingTimeInForce         *models.OrderOcoStopLimitTimeInForceParameter
	workingStrategyId          *int64
	workingStrategyType        *int32
	workingPegPriceType        *models.OrderListOcoAbovePegPriceTypeParameter
	workingPegOffsetType       *models.OrderListOcoAbovePegOffsetTypeParameter
	workingPegOffsetValue      *int32
	pendingAboveClientOrderId  *string
	pendingAbovePrice          *float32
	pendingAboveStopPrice      *float32
	pendingAboveTrailingDelta  *float32
	pendingAboveIcebergQty     *float32
	pendingAboveTimeInForce    *models.OrderOcoStopLimitTimeInForceParameter
	pendingAboveStrategyId     *int64
	pendingAboveStrategyType   *int32
	pendingAbovePegPriceType   *models.OrderListOcoAbovePegPriceTypeParameter
	pendingAbovePegOffsetType  *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingAbovePegOffsetValue *int32
	pendingBelowType           *models.OrderListOcoBelowTypeParameter
	pendingBelowClientOrderId  *string
	pendingBelowPrice          *float32
	pendingBelowStopPrice      *float32
	pendingBelowTrailingDelta  *float32
	pendingBelowIcebergQty     *float32
	pendingBelowTimeInForce    *models.OrderOcoStopLimitTimeInForceParameter
	pendingBelowStrategyId     *int64
	pendingBelowStrategyType   *int32
	pendingBelowPegPriceType   *models.OrderListOcoAbovePegPriceTypeParameter
	pendingBelowPegOffsetType  *models.OrderListOcoAbovePegOffsetTypeParameter
	pendingBelowPegOffsetValue *int32
	recvWindow                 *float32
}

func (r ApiOrderListOtocoRequest) Symbol(symbol string) ApiOrderListOtocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderListOtocoRequest) WorkingType(workingType models.OrderListOpoWorkingTypeParameter) ApiOrderListOtocoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiOrderListOtocoRequest) WorkingSide(workingSide models.NewOrderSideParameter) ApiOrderListOtocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiOrderListOtocoRequest) WorkingPrice(workingPrice float32) ApiOrderListOtocoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiOrderListOtocoRequest) WorkingQuantity(workingQuantity float32) ApiOrderListOtocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiOrderListOtocoRequest) PendingSide(pendingSide models.NewOrderSideParameter) ApiOrderListOtocoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiOrderListOtocoRequest) PendingQuantity(pendingQuantity float32) ApiOrderListOtocoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

func (r ApiOrderListOtocoRequest) PendingAboveType(pendingAboveType models.OrderListOcoAboveTypeParameter) ApiOrderListOtocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderListOtocoRequest) ListClientOrderId(listClientOrderId string) ApiOrderListOtocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiOrderListOtocoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderListOtocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderListOtocoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderListOtocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiOrderListOtocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiOrderListOtocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;workingType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOtocoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiOrderListOtocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiOrderListOtocoRequest) WorkingTimeInForce(workingTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOtocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary numeric value identifying the working order within an order strategy.
func (r ApiOrderListOtocoRequest) WorkingStrategyId(workingStrategyId int64) ApiOrderListOtocoRequest {
	r.workingStrategyId = &workingStrategyId
	return r
}

// Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOtocoRequest) WorkingStrategyType(workingStrategyType int32) ApiOrderListOtocoRequest {
	r.workingStrategyType = &workingStrategyType
	return r
}

func (r ApiOrderListOtocoRequest) WorkingPegPriceType(workingPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOtocoRequest {
	r.workingPegPriceType = &workingPegPriceType
	return r
}

func (r ApiOrderListOtocoRequest) WorkingPegOffsetType(workingPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOtocoRequest {
	r.workingPegOffsetType = &workingPegOffsetType
	return r
}

func (r ApiOrderListOtocoRequest) WorkingPegOffsetValue(workingPegOffsetValue int32) ApiOrderListOtocoRequest {
	r.workingPegOffsetValue = &workingPegOffsetValue
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiOrderListOtocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiOrderListOtocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; , &#x60;LIMIT_MAKER&#x60;, or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify the limit price.
func (r ApiOrderListOtocoRequest) PendingAbovePrice(pendingAbovePrice float32) ApiOrderListOtocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

// Can be used if &#x60;pendingAboveType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_LIMIT&#x60;
func (r ApiOrderListOtocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float32) ApiOrderListOtocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

// See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
func (r ApiOrderListOtocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float32) ApiOrderListOtocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60; or if &#x60;pendingAboveType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOtocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float32) ApiOrderListOtocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiOrderListOtocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOtocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending above order within an order strategy.
func (r ApiOrderListOtocoRequest) PendingAboveStrategyId(pendingAboveStrategyId int64) ApiOrderListOtocoRequest {
	r.pendingAboveStrategyId = &pendingAboveStrategyId
	return r
}

// Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOtocoRequest) PendingAboveStrategyType(pendingAboveStrategyType int32) ApiOrderListOtocoRequest {
	r.pendingAboveStrategyType = &pendingAboveStrategyType
	return r
}

func (r ApiOrderListOtocoRequest) PendingAbovePegPriceType(pendingAbovePegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOtocoRequest {
	r.pendingAbovePegPriceType = &pendingAbovePegPriceType
	return r
}

func (r ApiOrderListOtocoRequest) PendingAbovePegOffsetType(pendingAbovePegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOtocoRequest {
	r.pendingAbovePegOffsetType = &pendingAbovePegOffsetType
	return r
}

func (r ApiOrderListOtocoRequest) PendingAbovePegOffsetValue(pendingAbovePegOffsetValue int32) ApiOrderListOtocoRequest {
	r.pendingAbovePegOffsetValue = &pendingAbovePegOffsetValue
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowType(pendingBelowType models.OrderListOcoBelowTypeParameter) ApiOrderListOtocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiOrderListOtocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiOrderListOtocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS_LIMIT&#x60; or &#x60;TAKE_PROFIT_LIMIT&#x60; to specify limit price
func (r ApiOrderListOtocoRequest) PendingBelowPrice(pendingBelowPrice float32) ApiOrderListOtocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

// Can be used if &#x60;pendingBelowType&#x60; is &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT&#x60;. Either &#x60;pendingBelowStopPrice&#x60; or &#x60;pendingBelowTrailingDelta&#x60; or both, must be specified.
func (r ApiOrderListOtocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float32) ApiOrderListOtocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float32) ApiOrderListOtocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;, or if &#x60;pendingBelowType&#x60; is &#x60;LIMIT_MAKER&#x60;.
func (r ApiOrderListOtocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float32) ApiOrderListOtocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderListOtocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

// Arbitrary numeric value identifying the pending below order within an order strategy.
func (r ApiOrderListOtocoRequest) PendingBelowStrategyId(pendingBelowStrategyId int64) ApiOrderListOtocoRequest {
	r.pendingBelowStrategyId = &pendingBelowStrategyId
	return r
}

// Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
func (r ApiOrderListOtocoRequest) PendingBelowStrategyType(pendingBelowStrategyType int32) ApiOrderListOtocoRequest {
	r.pendingBelowStrategyType = &pendingBelowStrategyType
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowPegPriceType(pendingBelowPegPriceType models.OrderListOcoAbovePegPriceTypeParameter) ApiOrderListOtocoRequest {
	r.pendingBelowPegPriceType = &pendingBelowPegPriceType
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowPegOffsetType(pendingBelowPegOffsetType models.OrderListOcoAbovePegOffsetTypeParameter) ApiOrderListOtocoRequest {
	r.pendingBelowPegOffsetType = &pendingBelowPegOffsetType
	return r
}

func (r ApiOrderListOtocoRequest) PendingBelowPegOffsetValue(pendingBelowPegOffsetValue int32) ApiOrderListOtocoRequest {
	r.pendingBelowPegOffsetValue = &pendingBelowPegOffsetValue
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderListOtocoRequest) RecvWindow(recvWindow float32) ApiOrderListOtocoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderListOtocoRequest) Execute() (*common.RestApiResponse[models.OrderListOtocoResponse], error) {
	return r.ApiService.OrderListOtocoExecute(r)
}

/*
OrderListOtoco New Order list - OTOCO
Post /api/v3/orderList/otoco

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-list---otoco-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -  Sets the quantity for the working order.
@param pendingSide -
@param pendingQuantity -  Sets the quantity for the pending order.
@param pendingAboveType -
@param listClientOrderId -  A unique Id for the entire orderList
@param newOrderRespType -
@param selfTradePreventionMode -
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`, or if `workingType` is `LIMIT_MAKER`.
@param workingTimeInForce -
@param workingStrategyId -  Arbitrary numeric value identifying the working order within an order strategy.
@param workingStrategyType -  Arbitrary numeric value identifying the working order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param workingPegPriceType -
@param workingPegOffsetType -
@param workingPegOffsetValue -
@param pendingAboveClientOrderId -  Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
@param pendingAbovePrice -  Can be used if `pendingAboveType` is `STOP_LOSS_LIMIT` , `LIMIT_MAKER`, or `TAKE_PROFIT_LIMIT` to specify the limit price.
@param pendingAboveStopPrice -  Can be used if `pendingAboveType` is `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, `TAKE_PROFIT_LIMIT`
@param pendingAboveTrailingDelta -  See [Trailing Stop FAQ](./faqs/trailing-stop-faq.md)
@param pendingAboveIcebergQty -  This can only be used if `pendingAboveTimeInForce` is `GTC` or if `pendingAboveType` is `LIMIT_MAKER`.
@param pendingAboveTimeInForce -
@param pendingAboveStrategyId -  Arbitrary numeric value identifying the pending above order within an order strategy.
@param pendingAboveStrategyType -  Arbitrary numeric value identifying the pending above order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingAbovePegPriceType -
@param pendingAbovePegOffsetType -
@param pendingAbovePegOffsetValue -
@param pendingBelowType -
@param pendingBelowClientOrderId -  Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
@param pendingBelowPrice -  Can be used if `pendingBelowType` is `STOP_LOSS_LIMIT` or `TAKE_PROFIT_LIMIT` to specify limit price
@param pendingBelowStopPrice -  Can be used if `pendingBelowType` is `STOP_LOSS`, `STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT`. Either `pendingBelowStopPrice` or `pendingBelowTrailingDelta` or both, must be specified.
@param pendingBelowTrailingDelta -
@param pendingBelowIcebergQty -  This can only be used if `pendingBelowTimeInForce` is `GTC`, or if `pendingBelowType` is `LIMIT_MAKER`.
@param pendingBelowTimeInForce -
@param pendingBelowStrategyId -  Arbitrary numeric value identifying the pending below order within an order strategy.
@param pendingBelowStrategyType -  Arbitrary numeric value identifying the pending below order strategy. Values smaller than 1000000 are reserved and cannot be used.
@param pendingBelowPegPriceType -
@param pendingBelowPegOffsetType -
@param pendingBelowPegOffsetValue -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderListOtocoRequest
*/
func (a *TradeAPIService) OrderListOtoco(ctx context.Context) ApiOrderListOtocoRequest {
	return ApiOrderListOtocoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderListOtocoResponse
func (a *TradeAPIService) OrderListOtocoExecute(r ApiOrderListOtocoRequest) (*common.RestApiResponse[models.OrderListOtocoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList/otoco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}
	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}
	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}
	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}
	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}
	if r.pendingQuantity == nil {
		return nil, common.ReportError("pendingQuantity is required and must be specified")
	}
	if r.pendingAboveType == nil {
		return nil, common.ReportError("pendingAboveType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	if r.workingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	}
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	if r.workingStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyId", r.workingStrategyId, "form", "")
	}
	if r.workingStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingStrategyType", r.workingStrategyType, "form", "")
	}
	if r.workingPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegPriceType", r.workingPegPriceType, "form", "")
	}
	if r.workingPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetType", r.workingPegOffsetType, "form", "")
	}
	if r.workingPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPegOffsetValue", r.workingPegOffsetValue, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingQuantity", r.pendingQuantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveType", r.pendingAboveType, "form", "")
	if r.pendingAboveClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveClientOrderId", r.pendingAboveClientOrderId, "form", "")
	}
	if r.pendingAbovePrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePrice", r.pendingAbovePrice, "form", "")
	}
	if r.pendingAboveStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStopPrice", r.pendingAboveStopPrice, "form", "")
	}
	if r.pendingAboveTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTrailingDelta", r.pendingAboveTrailingDelta, "form", "")
	}
	if r.pendingAboveIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveIcebergQty", r.pendingAboveIcebergQty, "form", "")
	}
	if r.pendingAboveTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTimeInForce", r.pendingAboveTimeInForce, "form", "")
	}
	if r.pendingAboveStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStrategyId", r.pendingAboveStrategyId, "form", "")
	}
	if r.pendingAboveStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStrategyType", r.pendingAboveStrategyType, "form", "")
	}
	if r.pendingAbovePegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegPriceType", r.pendingAbovePegPriceType, "form", "")
	}
	if r.pendingAbovePegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegOffsetType", r.pendingAbovePegOffsetType, "form", "")
	}
	if r.pendingAbovePegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePegOffsetValue", r.pendingAbovePegOffsetValue, "form", "")
	}
	if r.pendingBelowType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowType", r.pendingBelowType, "form", "")
	}
	if r.pendingBelowClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowClientOrderId", r.pendingBelowClientOrderId, "form", "")
	}
	if r.pendingBelowPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPrice", r.pendingBelowPrice, "form", "")
	}
	if r.pendingBelowStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStopPrice", r.pendingBelowStopPrice, "form", "")
	}
	if r.pendingBelowTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTrailingDelta", r.pendingBelowTrailingDelta, "form", "")
	}
	if r.pendingBelowIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowIcebergQty", r.pendingBelowIcebergQty, "form", "")
	}
	if r.pendingBelowTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTimeInForce", r.pendingBelowTimeInForce, "form", "")
	}
	if r.pendingBelowStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStrategyId", r.pendingBelowStrategyId, "form", "")
	}
	if r.pendingBelowStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStrategyType", r.pendingBelowStrategyType, "form", "")
	}
	if r.pendingBelowPegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegPriceType", r.pendingBelowPegPriceType, "form", "")
	}
	if r.pendingBelowPegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegOffsetType", r.pendingBelowPegOffsetType, "form", "")
	}
	if r.pendingBelowPegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPegOffsetValue", r.pendingBelowPegOffsetValue, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderListOtocoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderOcoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	quantity                *float32
	price                   *float32
	stopPrice               *float32
	listClientOrderId       *string
	limitClientOrderId      *string
	limitStrategyId         *int64
	limitStrategyType       *int32
	limitIcebergQty         *float32
	trailingDelta           *int64
	stopClientOrderId       *string
	stopStrategyId          *int64
	stopStrategyType        *int32
	stopLimitPrice          *float32
	stopIcebergQty          *float32
	stopLimitTimeInForce    *models.OrderOcoStopLimitTimeInForceParameter
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiOrderOcoRequest) Symbol(symbol string) ApiOrderOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderOcoRequest) Side(side models.NewOrderSideParameter) ApiOrderOcoRequest {
	r.side = &side
	return r
}

func (r ApiOrderOcoRequest) Quantity(quantity float32) ApiOrderOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiOrderOcoRequest) Price(price float32) ApiOrderOcoRequest {
	r.price = &price
	return r
}

func (r ApiOrderOcoRequest) StopPrice(stopPrice float32) ApiOrderOcoRequest {
	r.stopPrice = &stopPrice
	return r
}

// A unique Id for the entire orderList
func (r ApiOrderOcoRequest) ListClientOrderId(listClientOrderId string) ApiOrderOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique Id for the limit order
func (r ApiOrderOcoRequest) LimitClientOrderId(limitClientOrderId string) ApiOrderOcoRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

func (r ApiOrderOcoRequest) LimitStrategyId(limitStrategyId int64) ApiOrderOcoRequest {
	r.limitStrategyId = &limitStrategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderOcoRequest) LimitStrategyType(limitStrategyType int32) ApiOrderOcoRequest {
	r.limitStrategyType = &limitStrategyType
	return r
}

// Used to make the &#x60;LIMIT_MAKER&#x60; leg an iceberg order.
func (r ApiOrderOcoRequest) LimitIcebergQty(limitIcebergQty float32) ApiOrderOcoRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderOcoRequest) TrailingDelta(trailingDelta int64) ApiOrderOcoRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// A unique Id for the stop loss/stop loss limit leg
func (r ApiOrderOcoRequest) StopClientOrderId(stopClientOrderId string) ApiOrderOcoRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

func (r ApiOrderOcoRequest) StopStrategyId(stopStrategyId int64) ApiOrderOcoRequest {
	r.stopStrategyId = &stopStrategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiOrderOcoRequest) StopStrategyType(stopStrategyType int32) ApiOrderOcoRequest {
	r.stopStrategyType = &stopStrategyType
	return r
}

// If provided, &#x60;stopLimitTimeInForce&#x60; is required.
func (r ApiOrderOcoRequest) StopLimitPrice(stopLimitPrice float32) ApiOrderOcoRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

// Used with &#x60;STOP_LOSS_LIMIT&#x60; leg to make an iceberg order.
func (r ApiOrderOcoRequest) StopIcebergQty(stopIcebergQty float32) ApiOrderOcoRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

func (r ApiOrderOcoRequest) StopLimitTimeInForce(stopLimitTimeInForce models.OrderOcoStopLimitTimeInForceParameter) ApiOrderOcoRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

func (r ApiOrderOcoRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderOcoRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderOcoRequest) RecvWindow(recvWindow float32) ApiOrderOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderOcoRequest) Execute() (*common.RestApiResponse[models.OrderOcoResponse], error) {
	return r.ApiService.OrderOcoExecute(r)
}

/*
	OrderOco New OCO - Deprecated
	Post /api/v3/order/oco

	https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-oco---deprecated-trade

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param symbol -
	@param side -
	@param quantity -
	@param price -
	@param stopPrice -
	@param listClientOrderId -  A unique Id for the entire orderList
	@param limitClientOrderId -  A unique Id for the limit order
	@param limitStrategyId -
	@param limitStrategyType -  The value cannot be less than `1000000`.
	@param limitIcebergQty -  Used to make the `LIMIT_MAKER` leg an iceberg order.
	@param trailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
	@param stopClientOrderId -  A unique Id for the stop loss/stop loss limit leg
	@param stopStrategyId -
	@param stopStrategyType -  The value cannot be less than `1000000`.
	@param stopLimitPrice -  If provided, `stopLimitTimeInForce` is required.
	@param stopIcebergQty -  Used with `STOP_LOSS_LIMIT` leg to make an iceberg order.
	@param stopLimitTimeInForce -
	@param newOrderRespType -
	@param selfTradePreventionMode -
	@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
	@return ApiOrderOcoRequest

Deprecated
*/
func (a *TradeAPIService) OrderOco(ctx context.Context) ApiOrderOcoRequest {
	return ApiOrderOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderOcoResponse
//
// Deprecated
func (a *TradeAPIService) OrderOcoExecute(r ApiOrderOcoRequest) (*common.RestApiResponse[models.OrderOcoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/oco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}
	if r.price == nil {
		return nil, common.ReportError("price is required and must be specified")
	}
	if r.stopPrice == nil {
		return nil, common.ReportError("stopPrice is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.limitClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitClientOrderId", r.limitClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.limitStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitStrategyId", r.limitStrategyId, "form", "")
	}
	if r.limitStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitStrategyType", r.limitStrategyType, "form", "")
	}
	if r.limitIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitIcebergQty", r.limitIcebergQty, "form", "")
	}
	if r.trailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trailingDelta", r.trailingDelta, "form", "")
	}
	if r.stopClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopClientOrderId", r.stopClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	if r.stopStrategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopStrategyId", r.stopStrategyId, "form", "")
	}
	if r.stopStrategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopStrategyType", r.stopStrategyType, "form", "")
	}
	if r.stopLimitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitPrice", r.stopLimitPrice, "form", "")
	}
	if r.stopIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopIcebergQty", r.stopIcebergQty, "form", "")
	}
	if r.stopLimitTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitTimeInForce", r.stopLimitTimeInForce, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderTestRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	computeCommissionRates  *bool
	timeInForce             *models.NewOrderTimeInForceParameter
	quantity                *float32
	quoteOrderQty           *float32
	price                   *float32
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int32
	stopPrice               *float32
	trailingDelta           *int64
	icebergQty              *float32
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	pegPriceType            *models.NewOrderPegPriceTypeParameter
	pegOffsetValue          *int32
	pegOffsetType           *models.NewOrderPegOffsetTypeParameter
	recvWindow              *float32
}

func (r ApiOrderTestRequest) Symbol(symbol string) ApiOrderTestRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderTestRequest) Side(side models.NewOrderSideParameter) ApiOrderTestRequest {
	r.side = &side
	return r
}

func (r ApiOrderTestRequest) Type(type_ models.NewOrderTypeParameter) ApiOrderTestRequest {
	r.type_ = &type_
	return r
}

// Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
func (r ApiOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

func (r ApiOrderTestRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiOrderTestRequest {
	r.timeInForce = &timeInForce
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

func (r ApiOrderTestRequest) Price(price float32) ApiOrderTestRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
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

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiOrderTestRequest) StopPrice(stopPrice float32) ApiOrderTestRequest {
	r.stopPrice = &stopPrice
	return r
}

// See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
func (r ApiOrderTestRequest) TrailingDelta(trailingDelta int64) ApiOrderTestRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiOrderTestRequest) IcebergQty(icebergQty float32) ApiOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiOrderTestRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiOrderTestRequest) PegPriceType(pegPriceType models.NewOrderPegPriceTypeParameter) ApiOrderTestRequest {
	r.pegPriceType = &pegPriceType
	return r
}

// Priceleveltopegthepriceto(max:100).&lt;br&gt;See[PeggedOrdersInfo](#pegged-orders-info)
func (r ApiOrderTestRequest) PegOffsetValue(pegOffsetValue int32) ApiOrderTestRequest {
	r.pegOffsetValue = &pegOffsetValue
	return r
}

func (r ApiOrderTestRequest) PegOffsetType(pegOffsetType models.NewOrderPegOffsetTypeParameter) ApiOrderTestRequest {
	r.pegOffsetType = &pegOffsetType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderTestRequest) RecvWindow(recvWindow float32) ApiOrderTestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderTestRequest) Execute() (*common.RestApiResponse[models.OrderTestResponse], error) {
	return r.ApiService.OrderTestExecute(r)
}

/*
OrderTest Test new order
Post /api/v3/order/test

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#test-new-order-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param computeCommissionRates -  Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
@param timeInForce -
@param quantity -
@param quoteOrderQty -
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param strategyId -
@param strategyType -  The value cannot be less than `1000000`.
@param stopPrice -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param trailingDelta -  See [Trailing Stop order FAQ](faqs/trailing-stop-faq.md).
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -
@param selfTradePreventionMode -
@param pegPriceType -
@param pegOffsetValue -  Priceleveltopegthepriceto(max:100).<br>See[PeggedOrdersInfo](#pegged-orders-info)
@param pegOffsetType -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderTestRequest
*/
func (a *TradeAPIService) OrderTest(ctx context.Context) ApiOrderTestRequest {
	return ApiOrderTestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderTestResponse
func (a *TradeAPIService) OrderTestExecute(r ApiOrderTestRequest) (*common.RestApiResponse[models.OrderTestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/test"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	if r.computeCommissionRates != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "computeCommissionRates", r.computeCommissionRates, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.quoteOrderQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteOrderQty", r.quoteOrderQty, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.strategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.trailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trailingDelta", r.trailingDelta, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.pegPriceType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegPriceType", r.pegPriceType, "form", "")
	}
	if r.pegOffsetValue != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetValue", r.pegOffsetValue, "form", "")
	}
	if r.pegOffsetType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pegOffsetType", r.pegOffsetType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderTestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSorOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	quantity                *float32
	timeInForce             *models.NewOrderTimeInForceParameter
	price                   *float32
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int32
	icebergQty              *float32
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiSorOrderRequest) Symbol(symbol string) ApiSorOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSorOrderRequest) Side(side models.NewOrderSideParameter) ApiSorOrderRequest {
	r.side = &side
	return r
}

func (r ApiSorOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiSorOrderRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderRequest) Quantity(quantity float32) ApiSorOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiSorOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiSorOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderRequest) Price(price float32) ApiSorOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiSorOrderRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiSorOrderRequest) StrategyId(strategyId int64) ApiSorOrderRequest {
	r.strategyId = &strategyId
	return r
}

// The value cannot be less than &#x60;1000000&#x60;.
func (r ApiSorOrderRequest) StrategyType(strategyType int32) ApiSorOrderRequest {
	r.strategyType = &strategyType
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiSorOrderRequest) IcebergQty(icebergQty float32) ApiSorOrderRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiSorOrderRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiSorOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiSorOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiSorOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderRequest) RecvWindow(recvWindow float32) ApiSorOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSorOrderRequest) Execute() (*common.RestApiResponse[models.SorOrderResponse], error) {
	return r.ApiService.SorOrderExecute(r)
}

/*
SorOrder New order using SOR
Post /api/v3/sor/order

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-using-sor-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param quantity -
@param timeInForce -
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param strategyId -
@param strategyType -  The value cannot be less than `1000000`.
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -
@param selfTradePreventionMode -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSorOrderRequest
*/
func (a *TradeAPIService) SorOrder(ctx context.Context) ApiSorOrderRequest {
	return ApiSorOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SorOrderResponse
func (a *TradeAPIService) SorOrderExecute(r ApiSorOrderRequest) (*common.RestApiResponse[models.SorOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/sor/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.strategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SorOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSorOrderTestRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	quantity                *float32
	computeCommissionRates  *bool
	timeInForce             *models.NewOrderTimeInForceParameter
	price                   *float32
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int32
	icebergQty              *float32
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *float32
}

func (r ApiSorOrderTestRequest) Symbol(symbol string) ApiSorOrderTestRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSorOrderTestRequest) Side(side models.NewOrderSideParameter) ApiSorOrderTestRequest {
	r.side = &side
	return r
}

func (r ApiSorOrderTestRequest) Type(type_ models.NewOrderTypeParameter) ApiSorOrderTestRequest {
	r.type_ = &type_
	return r
}

func (r ApiSorOrderTestRequest) Quantity(quantity float32) ApiSorOrderTestRequest {
	r.quantity = &quantity
	return r
}

// Default: &#x60;false&#x60; &lt;br&gt; See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
func (r ApiSorOrderTestRequest) ComputeCommissionRates(computeCommissionRates bool) ApiSorOrderTestRequest {
	r.computeCommissionRates = &computeCommissionRates
	return r
}

func (r ApiSorOrderTestRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiSorOrderTestRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiSorOrderTestRequest) Price(price float32) ApiSorOrderTestRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent.&lt;br/&gt; Orders with the same &#x60;newClientOrderID&#x60; can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (r ApiSorOrderTestRequest) NewClientOrderId(newClientOrderId string) ApiSorOrderTestRequest {
	r.newClientOrderId = &newClientOrderId
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

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiSorOrderTestRequest) IcebergQty(icebergQty float32) ApiSorOrderTestRequest {
	r.icebergQty = &icebergQty
	return r
}

func (r ApiSorOrderTestRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiSorOrderTestRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiSorOrderTestRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiSorOrderTestRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSorOrderTestRequest) RecvWindow(recvWindow float32) ApiSorOrderTestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSorOrderTestRequest) Execute() (*common.RestApiResponse[models.SorOrderTestResponse], error) {
	return r.ApiService.SorOrderTestExecute(r)
}

/*
SorOrderTest Test new order using SOR
Post /api/v3/sor/order/test

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#test-new-order-using-sor-trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param quantity -
@param computeCommissionRates -  Default: `false` <br> See [Commissions FAQ](faqs/commission_faq.md#test-order-diferences) to learn more.
@param timeInForce -
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.<br/> Orders with the same `newClientOrderID` can be accepted only when the previous one is filled, otherwise the order will be rejected.
@param strategyId -
@param strategyType -  The value cannot be less than `1000000`.
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -
@param selfTradePreventionMode -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSorOrderTestRequest
*/
func (a *TradeAPIService) SorOrderTest(ctx context.Context) ApiSorOrderTestRequest {
	return ApiSorOrderTestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SorOrderTestResponse
func (a *TradeAPIService) SorOrderTestExecute(r ApiSorOrderTestRequest) (*common.RestApiResponse[models.SorOrderTestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/api/v3/sor/order/test"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}
	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}

	if r.computeCommissionRates != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "computeCommissionRates", r.computeCommissionRates, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.strategyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyId", r.strategyId, "form", "")
	}
	if r.strategyType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "strategyType", r.strategyType, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SorOrderTestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
