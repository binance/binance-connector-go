/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Option trading pair.
func (r ApiAccountTradeListRequest) Symbol(symbol string) ApiAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

// Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376
func (r ApiAccountTradeListRequest) FromId(fromId int64) ApiAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Start Time, e.g 1593511200000
func (r ApiAccountTradeListRequest) StartTime(startTime int64) ApiAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiAccountTradeListRequest) EndTime(endTime int64) ApiAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned.
func (r ApiAccountTradeListRequest) Limit(limit int64) ApiAccountTradeListRequest {
	r.limit = &limit
	return r
}

// Recv Window.
func (r ApiAccountTradeListRequest) RecvWindow(recvWindow int64) ApiAccountTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountTradeListRequest) Execute() (*common.RestApiResponse[models.AccountTradeListResponse], error) {
	return r.ApiService.AccountTradeListExecute(r)
}

/*
AccountTradeList Account Trade List (USER_DATA)
Get /eapi/v1/userTrades

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#account-trade-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param fromId -  Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned.
@param recvWindow -  Recv Window.
@return ApiAccountTradeListRequest
*/
func (a *TradeAPIService) AccountTradeList(ctx context.Context) ApiAccountTradeListRequest {
	return ApiAccountTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountTradeListResponse
func (a *TradeAPIService) AccountTradeListExecute(r ApiAccountTradeListRequest) (*common.RestApiResponse[models.AccountTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/userTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountTradeListResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelAllOptionOrdersByUnderlyingRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	underlying *string
	recvWindow *int64
}

// Underlying asset.
func (r ApiCancelAllOptionOrdersByUnderlyingRequest) Underlying(underlying string) ApiCancelAllOptionOrdersByUnderlyingRequest {
	r.underlying = &underlying
	return r
}

// Recv Window.
func (r ApiCancelAllOptionOrdersByUnderlyingRequest) RecvWindow(recvWindow int64) ApiCancelAllOptionOrdersByUnderlyingRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllOptionOrdersByUnderlyingRequest) Execute() (*common.RestApiResponse[models.CancelAllOptionOrdersByUnderlyingResponse], error) {
	return r.ApiService.CancelAllOptionOrdersByUnderlyingExecute(r)
}

/*
CancelAllOptionOrdersByUnderlying Cancel All Option Orders By Underlying (TRADE)
Delete /eapi/v1/allOpenOrdersByUnderlying

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#cancel-all-option-orders-by-underlying

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Underlying asset.
@param recvWindow -  Recv Window.
@return ApiCancelAllOptionOrdersByUnderlyingRequest
*/
func (a *TradeAPIService) CancelAllOptionOrdersByUnderlying(ctx context.Context) ApiCancelAllOptionOrdersByUnderlyingRequest {
	return ApiCancelAllOptionOrdersByUnderlyingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllOptionOrdersByUnderlyingResponse
func (a *TradeAPIService) CancelAllOptionOrdersByUnderlyingExecute(r ApiCancelAllOptionOrdersByUnderlyingRequest) (*common.RestApiResponse[models.CancelAllOptionOrdersByUnderlyingResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/allOpenOrdersByUnderlying"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllOptionOrdersByUnderlyingResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelAllOptionOrdersOnSpecificSymbolRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

// Option trading pair.
func (r ApiCancelAllOptionOrdersOnSpecificSymbolRequest) Symbol(symbol string) ApiCancelAllOptionOrdersOnSpecificSymbolRequest {
	r.symbol = &symbol
	return r
}

// Recv Window.
func (r ApiCancelAllOptionOrdersOnSpecificSymbolRequest) RecvWindow(recvWindow int64) ApiCancelAllOptionOrdersOnSpecificSymbolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllOptionOrdersOnSpecificSymbolRequest) Execute() (*common.RestApiResponse[models.CancelAllOptionOrdersOnSpecificSymbolResponse], error) {
	return r.ApiService.CancelAllOptionOrdersOnSpecificSymbolExecute(r)
}

/*
CancelAllOptionOrdersOnSpecificSymbol Cancel all Option orders on specific symbol (TRADE)
Delete /eapi/v1/allOpenOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#cancel-all-option-orders-on-specific-symbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param recvWindow -  Recv Window.
@return ApiCancelAllOptionOrdersOnSpecificSymbolRequest
*/
func (a *TradeAPIService) CancelAllOptionOrdersOnSpecificSymbol(ctx context.Context) ApiCancelAllOptionOrdersOnSpecificSymbolRequest {
	return ApiCancelAllOptionOrdersOnSpecificSymbolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllOptionOrdersOnSpecificSymbolResponse
func (a *TradeAPIService) CancelAllOptionOrdersOnSpecificSymbolExecute(r ApiCancelAllOptionOrdersOnSpecificSymbolRequest) (*common.RestApiResponse[models.CancelAllOptionOrdersOnSpecificSymbolResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllOptionOrdersOnSpecificSymbolResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelMultipleOptionOrdersRequest struct {
	ctx            context.Context
	ApiService     *TradeAPIService
	symbol         *string
	orderIds       *[]int64
	clientOrderIds *[]string
	recvWindow     *int64
}

// Option trading pair.
func (r ApiCancelMultipleOptionOrdersRequest) Symbol(symbol string) ApiCancelMultipleOptionOrdersRequest {
	r.symbol = &symbol
	return r
}

// Order ID list.
func (r ApiCancelMultipleOptionOrdersRequest) OrderIds(orderIds []int64) ApiCancelMultipleOptionOrdersRequest {
	r.orderIds = &orderIds
	return r
}

// Client order ID list.
func (r ApiCancelMultipleOptionOrdersRequest) ClientOrderIds(clientOrderIds []string) ApiCancelMultipleOptionOrdersRequest {
	r.clientOrderIds = &clientOrderIds
	return r
}

// Recv Window.
func (r ApiCancelMultipleOptionOrdersRequest) RecvWindow(recvWindow int64) ApiCancelMultipleOptionOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelMultipleOptionOrdersRequest) Execute() (*common.RestApiResponse[models.CancelMultipleOptionOrdersResponse], error) {
	return r.ApiService.CancelMultipleOptionOrdersExecute(r)
}

/*
CancelMultipleOptionOrders Cancel Multiple Option Orders (TRADE)
Delete /eapi/v1/batchOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#cancel-multiple-option-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param orderIds -  Order ID list.
@param clientOrderIds -  Client order ID list.
@param recvWindow -  Recv Window.
@return ApiCancelMultipleOptionOrdersRequest
*/
func (a *TradeAPIService) CancelMultipleOptionOrders(ctx context.Context) ApiCancelMultipleOptionOrdersRequest {
	return ApiCancelMultipleOptionOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelMultipleOptionOrdersResponse
func (a *TradeAPIService) CancelMultipleOptionOrdersExecute(r ApiCancelMultipleOptionOrdersRequest) (*common.RestApiResponse[models.CancelMultipleOptionOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/batchOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderIds != nil {
		t := *r.orderIds
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderIds", t, "form", "multi")
	}
	if r.clientOrderIds != nil {
		t := *r.clientOrderIds
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderIds", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelMultipleOptionOrdersResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelOptionOrderRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	orderId       *int64
	clientOrderId *string
	recvWindow    *int64
}

// Option trading pair.
func (r ApiCancelOptionOrderRequest) Symbol(symbol string) ApiCancelOptionOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID.
func (r ApiCancelOptionOrderRequest) OrderId(orderId int64) ApiCancelOptionOrderRequest {
	r.orderId = &orderId
	return r
}

// clientOrderId
func (r ApiCancelOptionOrderRequest) ClientOrderId(clientOrderId string) ApiCancelOptionOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// Recv Window.
func (r ApiCancelOptionOrderRequest) RecvWindow(recvWindow int64) ApiCancelOptionOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelOptionOrderRequest) Execute() (*common.RestApiResponse[models.CancelOptionOrderResponse], error) {
	return r.ApiService.CancelOptionOrderExecute(r)
}

/*
CancelOptionOrder Cancel Option Order (TRADE)
Delete /eapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#cancel-option-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param orderId -  Order ID.
@param clientOrderId -  clientOrderId
@param recvWindow -  Recv Window.
@return ApiCancelOptionOrderRequest
*/
func (a *TradeAPIService) CancelOptionOrder(ctx context.Context) ApiCancelOptionOrderRequest {
	return ApiCancelOptionOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelOptionOrderResponse
func (a *TradeAPIService) CancelOptionOrderExecute(r ApiCancelOptionOrderRequest) (*common.RestApiResponse[models.CancelOptionOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelOptionOrderResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.PlaceMultipleOrdersOrdersParameterInnerSide
	type_                   *models.PlaceMultipleOrdersOrdersParameterInnerType
	quantity                *float32
	price                   *float32
	timeInForce             *models.NewOrderTimeInForceParameter
	reduceOnly              *bool
	postOnly                *bool
	newOrderRespType        *models.NewOrderNewOrderRespTypeParameter
	clientOrderId           *string
	isMmp                   *bool
	selfTradePreventionMode *models.NewOrderSelfTradePreventionModeParameter
	recvWindow              *int64
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewOrderRequest) Side(side models.PlaceMultipleOrdersOrdersParameterInnerSide) ApiNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewOrderRequest) Type(type_ models.PlaceMultipleOrdersOrdersParameterInnerType) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Order Quantity
func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// Order Price
func (r ApiNewOrderRequest) Price(price float32) ApiNewOrderRequest {
	r.price = &price
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewOrderRequest) ReduceOnly(reduceOnly bool) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiNewOrderRequest) PostOnly(postOnly bool) ApiNewOrderRequest {
	r.postOnly = &postOnly
	return r
}

func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// User-defined order ID cannot be repeated in pending orders
func (r ApiNewOrderRequest) ClientOrderId(clientOrderId string) ApiNewOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// is market maker protection order
func (r ApiNewOrderRequest) IsMmp(isMmp bool) ApiNewOrderRequest {
	r.isMmp = &isMmp
	return r
}

// Self-trade prevention mode
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiNewOrderRequest) RecvWindow(recvWindow int64) ApiNewOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewOrderRequest) Execute() (*common.RestApiResponse[models.NewOrderResponse], error) {
	return r.ApiService.NewOrderExecute(r)
}

/*
NewOrder New Order (TRADE)
Post /eapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#new-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param quantity -  Order Quantity
@param price -  Order Price
@param timeInForce -
@param reduceOnly -
@param postOnly -
@param newOrderRespType -
@param clientOrderId -  User-defined order ID cannot be repeated in pending orders
@param isMmp -  is market maker protection order
@param selfTradePreventionMode -  Self-trade prevention mode
@param recvWindow -
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
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/order"

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
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.postOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "postOnly", r.postOnly, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.isMmp != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isMmp", r.isMmp, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewOrderResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOptionPositionInformationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

// Option trading pair.
func (r ApiOptionPositionInformationRequest) Symbol(symbol string) ApiOptionPositionInformationRequest {
	r.symbol = &symbol
	return r
}

// Recv Window.
func (r ApiOptionPositionInformationRequest) RecvWindow(recvWindow int64) ApiOptionPositionInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOptionPositionInformationRequest) Execute() (*common.RestApiResponse[models.OptionPositionInformationResponse], error) {
	return r.ApiService.OptionPositionInformationExecute(r)
}

/*
OptionPositionInformation Option Position Information (USER_DATA)
Get /eapi/v1/position

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#option-position-information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param recvWindow -  Recv Window.
@return ApiOptionPositionInformationRequest
*/
func (a *TradeAPIService) OptionPositionInformation(ctx context.Context) ApiOptionPositionInformationRequest {
	return ApiOptionPositionInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptionPositionInformationResponse
func (a *TradeAPIService) OptionPositionInformationExecute(r ApiOptionPositionInformationRequest) (*common.RestApiResponse[models.OptionPositionInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/position"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OptionPositionInformationResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPlaceMultipleOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orders     *[]models.PlaceMultipleOrdersOrdersParameterInner
	recvWindow *int64
}

// order list. Max 10 orders
func (r ApiPlaceMultipleOrdersRequest) Orders(orders []models.PlaceMultipleOrdersOrdersParameterInner) ApiPlaceMultipleOrdersRequest {
	r.orders = &orders
	return r
}

func (r ApiPlaceMultipleOrdersRequest) RecvWindow(recvWindow int64) ApiPlaceMultipleOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPlaceMultipleOrdersRequest) Execute() (*common.RestApiResponse[models.PlaceMultipleOrdersResponse], error) {
	return r.ApiService.PlaceMultipleOrdersExecute(r)
}

/*
PlaceMultipleOrders Place Multiple Orders (TRADE)
Post /eapi/v1/batchOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#place-multiple-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orders -  order list. Max 10 orders
@param recvWindow -
@return ApiPlaceMultipleOrdersRequest
*/
func (a *TradeAPIService) PlaceMultipleOrders(ctx context.Context) ApiPlaceMultipleOrdersRequest {
	return ApiPlaceMultipleOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PlaceMultipleOrdersResponse
func (a *TradeAPIService) PlaceMultipleOrdersExecute(r ApiPlaceMultipleOrdersRequest) (*common.RestApiResponse[models.PlaceMultipleOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/batchOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orders == nil {
		return nil, common.ReportError("orders is required and must be specified")
	}

	{
		t := *r.orders
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orders", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PlaceMultipleOrdersResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentOpenOptionOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// Option trading pair.
func (r ApiQueryCurrentOpenOptionOrdersRequest) Symbol(symbol string) ApiQueryCurrentOpenOptionOrdersRequest {
	r.symbol = &symbol
	return r
}

// Order ID.
func (r ApiQueryCurrentOpenOptionOrdersRequest) OrderId(orderId int64) ApiQueryCurrentOpenOptionOrdersRequest {
	r.orderId = &orderId
	return r
}

// Start Time, e.g 1593511200000
func (r ApiQueryCurrentOpenOptionOrdersRequest) StartTime(startTime int64) ApiQueryCurrentOpenOptionOrdersRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiQueryCurrentOpenOptionOrdersRequest) EndTime(endTime int64) ApiQueryCurrentOpenOptionOrdersRequest {
	r.endTime = &endTime
	return r
}

// Recv Window.
func (r ApiQueryCurrentOpenOptionOrdersRequest) RecvWindow(recvWindow int64) ApiQueryCurrentOpenOptionOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentOpenOptionOrdersRequest) Execute() (*common.RestApiResponse[models.QueryCurrentOpenOptionOrdersResponse], error) {
	return r.ApiService.QueryCurrentOpenOptionOrdersExecute(r)
}

/*
QueryCurrentOpenOptionOrders Query Current Open Option Orders (USER_DATA)
Get /eapi/v1/openOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#query-current-open-option-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param orderId -  Order ID.
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param recvWindow -  Recv Window.
@return ApiQueryCurrentOpenOptionOrdersRequest
*/
func (a *TradeAPIService) QueryCurrentOpenOptionOrders(ctx context.Context) ApiQueryCurrentOpenOptionOrdersRequest {
	return ApiQueryCurrentOpenOptionOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentOpenOptionOrdersResponse
func (a *TradeAPIService) QueryCurrentOpenOptionOrdersExecute(r ApiQueryCurrentOpenOptionOrdersRequest) (*common.RestApiResponse[models.QueryCurrentOpenOptionOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentOpenOptionOrdersResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryOptionOrderHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Option trading pair.
func (r ApiQueryOptionOrderHistoryRequest) Symbol(symbol string) ApiQueryOptionOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

// Order ID.
func (r ApiQueryOptionOrderHistoryRequest) OrderId(orderId int64) ApiQueryOptionOrderHistoryRequest {
	r.orderId = &orderId
	return r
}

// Start Time, e.g 1593511200000
func (r ApiQueryOptionOrderHistoryRequest) StartTime(startTime int64) ApiQueryOptionOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiQueryOptionOrderHistoryRequest) EndTime(endTime int64) ApiQueryOptionOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned
func (r ApiQueryOptionOrderHistoryRequest) Limit(limit int64) ApiQueryOptionOrderHistoryRequest {
	r.limit = &limit
	return r
}

// Recv Window.
func (r ApiQueryOptionOrderHistoryRequest) RecvWindow(recvWindow int64) ApiQueryOptionOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOptionOrderHistoryRequest) Execute() (*common.RestApiResponse[models.QueryOptionOrderHistoryResponse], error) {
	return r.ApiService.QueryOptionOrderHistoryExecute(r)
}

/*
QueryOptionOrderHistory Query Option Order History (TRADE)
Get /eapi/v1/historyOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#query-option-order-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param orderId -  Order ID.
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned
@param recvWindow -  Recv Window.
@return ApiQueryOptionOrderHistoryRequest
*/
func (a *TradeAPIService) QueryOptionOrderHistory(ctx context.Context) ApiQueryOptionOrderHistoryRequest {
	return ApiQueryOptionOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryOptionOrderHistoryResponse
func (a *TradeAPIService) QueryOptionOrderHistoryExecute(r ApiQueryOptionOrderHistoryRequest) (*common.RestApiResponse[models.QueryOptionOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/historyOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryOptionOrderHistoryResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySingleOrderRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	orderId       *int64
	clientOrderId *string
	recvWindow    *int64
}

// Option trading pair.
func (r ApiQuerySingleOrderRequest) Symbol(symbol string) ApiQuerySingleOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID.
func (r ApiQuerySingleOrderRequest) OrderId(orderId int64) ApiQuerySingleOrderRequest {
	r.orderId = &orderId
	return r
}

// User-defined order ID; cannot be duplicated among open orders.
func (r ApiQuerySingleOrderRequest) ClientOrderId(clientOrderId string) ApiQuerySingleOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// Recv Window.
func (r ApiQuerySingleOrderRequest) RecvWindow(recvWindow int64) ApiQuerySingleOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySingleOrderRequest) Execute() (*common.RestApiResponse[models.QuerySingleOrderResponse], error) {
	return r.ApiService.QuerySingleOrderExecute(r)
}

/*
QuerySingleOrder Query Single Order (TRADE)
Get /eapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#query-single-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param orderId -  Order ID.
@param clientOrderId -  User-defined order ID; cannot be duplicated among open orders.
@param recvWindow -  Recv Window.
@return ApiQuerySingleOrderRequest
*/
func (a *TradeAPIService) QuerySingleOrder(ctx context.Context) ApiQuerySingleOrderRequest {
	return ApiQuerySingleOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySingleOrderResponse
func (a *TradeAPIService) QuerySingleOrderExecute(r ApiQuerySingleOrderRequest) (*common.RestApiResponse[models.QuerySingleOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySingleOrderResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTradfiOptionsContractRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiTradfiOptionsContractRequest) RecvWindow(recvWindow int64) ApiTradfiOptionsContractRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTradfiOptionsContractRequest) Execute() (*common.RestApiResponse[models.TradfiOptionsContractResponse], error) {
	return r.ApiService.TradfiOptionsContractExecute(r)
}

/*
TradfiOptionsContract TradFi Options Contract (USER_DATA)
Post /eapi/v1/stock/contract

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#tradfi-options-contract

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiTradfiOptionsContractRequest
*/
func (a *TradeAPIService) TradfiOptionsContract(ctx context.Context) ApiTradfiOptionsContractRequest {
	return ApiTradfiOptionsContractRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TradfiOptionsContractResponse
func (a *TradeAPIService) TradfiOptionsContractExecute(r ApiTradfiOptionsContractRequest) (*common.RestApiResponse[models.TradfiOptionsContractResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/stock/contract"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TradfiOptionsContractResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUserCommissionRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// Recv Window.
func (r ApiUserCommissionRequest) RecvWindow(recvWindow int64) ApiUserCommissionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUserCommissionRequest) Execute() (*common.RestApiResponse[models.UserCommissionResponse], error) {
	return r.ApiService.UserCommissionExecute(r)
}

/*
UserCommission User Commission (USER_DATA)
Get /eapi/v1/commission

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#user-commission

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Recv Window.
@return ApiUserCommissionRequest
*/
func (a *TradeAPIService) UserCommission(ctx context.Context) ApiUserCommissionRequest {
	return ApiUserCommissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserCommissionResponse
func (a *TradeAPIService) UserCommissionExecute(r ApiUserCommissionRequest) (*common.RestApiResponse[models.UserCommissionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/commission"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UserCommissionResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUserExerciseRecordRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Option trading pair.
func (r ApiUserExerciseRecordRequest) Symbol(symbol string) ApiUserExerciseRecordRequest {
	r.symbol = &symbol
	return r
}

// Start Time, e.g 1593511200000
func (r ApiUserExerciseRecordRequest) StartTime(startTime int64) ApiUserExerciseRecordRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiUserExerciseRecordRequest) EndTime(endTime int64) ApiUserExerciseRecordRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned.
func (r ApiUserExerciseRecordRequest) Limit(limit int64) ApiUserExerciseRecordRequest {
	r.limit = &limit
	return r
}

// Recv Window.
func (r ApiUserExerciseRecordRequest) RecvWindow(recvWindow int64) ApiUserExerciseRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUserExerciseRecordRequest) Execute() (*common.RestApiResponse[models.UserExerciseRecordResponse], error) {
	return r.ApiService.UserExerciseRecordExecute(r)
}

/*
UserExerciseRecord User Exercise Record (USER_DATA)
Get /eapi/v1/exerciseRecord

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/trade#user-exercise-record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair.
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned.
@param recvWindow -  Recv Window.
@return ApiUserExerciseRecordRequest
*/
func (a *TradeAPIService) UserExerciseRecord(ctx context.Context) ApiUserExerciseRecordRequest {
	return ApiUserExerciseRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserExerciseRecordResponse
func (a *TradeAPIService) UserExerciseRecordExecute(r ApiUserExerciseRecordRequest) (*common.RestApiResponse[models.UserExerciseRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/exerciseRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UserExerciseRecordResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
