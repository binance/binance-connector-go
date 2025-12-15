/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiAccountTradeListRequest) Symbol(symbol string) ApiAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

// The UniqueId ID from which to return. The latest deal record is returned by default
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

// Number of result sets returned Default:100 Max:1000
func (r ApiAccountTradeListRequest) Limit(limit int64) ApiAccountTradeListRequest {
	r.limit = &limit
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param fromId -  The UniqueId ID from which to return. The latest deal record is returned by default
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@param recvWindow -
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

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
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

	resp, err := SendRequest[models.AccountTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option underlying, e.g BTCUSDT
func (r ApiCancelAllOptionOrdersByUnderlyingRequest) Underlying(underlying string) ApiCancelAllOptionOrdersByUnderlyingRequest {
	r.underlying = &underlying
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Cancel-All-Option-Orders-By-Underlying

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Option underlying, e.g BTCUSDT
@param recvWindow -
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

	resp, err := SendRequest[models.CancelAllOptionOrdersByUnderlyingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiCancelAllOptionOrdersOnSpecificSymbolRequest) Symbol(symbol string) ApiCancelAllOptionOrdersOnSpecificSymbolRequest {
	r.symbol = &symbol
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Cancel-all-Option-orders-on-specific-symbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param recvWindow -
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

	resp, err := SendRequest[models.CancelAllOptionOrdersOnSpecificSymbolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiCancelMultipleOptionOrdersRequest) Symbol(symbol string) ApiCancelMultipleOptionOrdersRequest {
	r.symbol = &symbol
	return r
}

// Order ID, e.g [4611875134427365377,4611875134427365378]
func (r ApiCancelMultipleOptionOrdersRequest) OrderIds(orderIds []int64) ApiCancelMultipleOptionOrdersRequest {
	r.orderIds = &orderIds
	return r
}

// User-defined order ID, e.g [\&quot;my_id_1\&quot;,\&quot;my_id_2\&quot;]
func (r ApiCancelMultipleOptionOrdersRequest) ClientOrderIds(clientOrderIds []string) ApiCancelMultipleOptionOrdersRequest {
	r.clientOrderIds = &clientOrderIds
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Cancel-Multiple-Option-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param orderIds -  Order ID, e.g [4611875134427365377,4611875134427365378]
@param clientOrderIds -  User-defined order ID, e.g [\"my_id_1\",\"my_id_2\"]
@param recvWindow -
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

	resp, err := SendRequest[models.CancelMultipleOptionOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiCancelOptionOrderRequest) Symbol(symbol string) ApiCancelOptionOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID, e.g 4611875134427365377
func (r ApiCancelOptionOrderRequest) OrderId(orderId int64) ApiCancelOptionOrderRequest {
	r.orderId = &orderId
	return r
}

// User-defined order ID, e.g 10000
func (r ApiCancelOptionOrderRequest) ClientOrderId(clientOrderId string) ApiCancelOptionOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Cancel-Option-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param orderId -  Order ID, e.g 4611875134427365377
@param clientOrderId -  User-defined order ID, e.g 10000
@param recvWindow -
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

	resp, err := SendRequest[models.CancelOptionOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewOrderRequest struct {
	ctx              context.Context
	ApiService       *TradeAPIService
	symbol           *string
	side             *models.PlaceMultipleOrdersOrdersParameterInnerSide
	type_            *models.PlaceMultipleOrdersOrdersParameterInnerType
	quantity         *float32
	price            *float32
	timeInForce      *models.PlaceMultipleOrdersOrdersParameterInnerTimeInForce
	reduceOnly       *bool
	postOnly         *bool
	newOrderRespType *models.PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType
	clientOrderId    *string
	isMmp            *bool
	recvWindow       *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

// Buy/sell direction: SELL, BUY
func (r ApiNewOrderRequest) Side(side models.PlaceMultipleOrdersOrdersParameterInnerSide) ApiNewOrderRequest {
	r.side = &side
	return r
}

// Order Type: LIMIT(only support limit)
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

// Time in force method（Default GTC）
func (r ApiNewOrderRequest) TimeInForce(timeInForce models.PlaceMultipleOrdersOrdersParameterInnerTimeInForce) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// Reduce Only（Default false）
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly bool) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Post Only（Default false）
func (r ApiNewOrderRequest) PostOnly(postOnly bool) ApiNewOrderRequest {
	r.postOnly = &postOnly
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, Default \&quot;ACK\&quot;
func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// User-defined order ID, e.g 10000
func (r ApiNewOrderRequest) ClientOrderId(clientOrderId string) ApiNewOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// is market maker protection order, true/false
func (r ApiNewOrderRequest) IsMmp(isMmp bool) ApiNewOrderRequest {
	r.isMmp = &isMmp
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

https://developers.binance.com/docs/derivatives/option/trade/New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param side -  Buy/sell direction: SELL, BUY
@param type_ -  Order Type: LIMIT(only support limit)
@param quantity -  Order Quantity
@param price -  Order Price
@param timeInForce -  Time in force method（Default GTC）
@param reduceOnly -  Reduce Only（Default false）
@param postOnly -  Post Only（Default false）
@param newOrderRespType -  \"ACK\", \"RESULT\", Default \"ACK\"
@param clientOrderId -  User-defined order ID, e.g 10000
@param isMmp -  is market maker protection order, true/false
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
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiOptionPositionInformationRequest) Symbol(symbol string) ApiOptionPositionInformationRequest {
	r.symbol = &symbol
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Option-Position-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param recvWindow -
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

	resp, err := SendRequest[models.OptionPositionInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
PlaceMultipleOrders Place Multiple Orders(TRADE)
Post /eapi/v1/batchOrders

https://developers.binance.com/docs/derivatives/option/trade/Place-Multiple-Orders

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

	resp, err := SendRequest[models.PlaceMultipleOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiQueryCurrentOpenOptionOrdersRequest) Symbol(symbol string) ApiQueryCurrentOpenOptionOrdersRequest {
	r.symbol = &symbol
	return r
}

// Order ID, e.g 4611875134427365377
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

https://developers.binance.com/docs/derivatives/option/trade/Query-Current-Open-Option-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param orderId -  Order ID, e.g 4611875134427365377
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param recvWindow -
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

	resp, err := SendRequest[models.QueryCurrentOpenOptionOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiQueryOptionOrderHistoryRequest) Symbol(symbol string) ApiQueryOptionOrderHistoryRequest {
	r.symbol = &symbol
	return r
}

// Order ID, e.g 4611875134427365377
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

// Number of result sets returned Default:100 Max:1000
func (r ApiQueryOptionOrderHistoryRequest) Limit(limit int64) ApiQueryOptionOrderHistoryRequest {
	r.limit = &limit
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Query-Option-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param orderId -  Order ID, e.g 4611875134427365377
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@param recvWindow -
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

	resp, err := SendRequest[models.QueryOptionOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiQuerySingleOrderRequest) Symbol(symbol string) ApiQuerySingleOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID, e.g 4611875134427365377
func (r ApiQuerySingleOrderRequest) OrderId(orderId int64) ApiQuerySingleOrderRequest {
	r.orderId = &orderId
	return r
}

// User-defined order ID, e.g 10000
func (r ApiQuerySingleOrderRequest) ClientOrderId(clientOrderId string) ApiQuerySingleOrderRequest {
	r.clientOrderId = &clientOrderId
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/Query-Single-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param orderId -  Order ID, e.g 4611875134427365377
@param clientOrderId -  User-defined order ID, e.g 10000
@param recvWindow -
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

	resp, err := SendRequest[models.QuerySingleOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Option trading pair, e.g BTC-200730-9000-C
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

// Number of result sets returned Default:100 Max:1000
func (r ApiUserExerciseRecordRequest) Limit(limit int64) ApiUserExerciseRecordRequest {
	r.limit = &limit
	return r
}

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

https://developers.binance.com/docs/derivatives/option/trade/User-Exercise-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@param recvWindow -
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

	resp, err := SendRequest[models.UserExerciseRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
