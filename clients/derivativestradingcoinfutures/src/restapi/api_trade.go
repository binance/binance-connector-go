/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	orderId    *string
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

// Symbol
func (r ApiAccountTradeListRequest) Symbol(symbol string) ApiAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

// pair
func (r ApiAccountTradeListRequest) Pair(pair string) ApiAccountTradeListRequest {
	r.pair = &pair
	return r
}

// Order ID
func (r ApiAccountTradeListRequest) OrderId(orderId string) ApiAccountTradeListRequest {
	r.orderId = &orderId
	return r
}

// Start time
func (r ApiAccountTradeListRequest) StartTime(startTime int64) ApiAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

// End time
func (r ApiAccountTradeListRequest) EndTime(endTime int64) ApiAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// Trade id to fetch from. Default gets most recent trades.
func (r ApiAccountTradeListRequest) FromId(fromId int64) ApiAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Maximum number of records to return.
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
Get /dapi/v1/userTrades

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#account-trade-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param pair -  pair
@param orderId -  Order ID
@param startTime -  Start time
@param endTime -  End time
@param fromId -  Trade id to fetch from. Default gets most recent trades.
@param limit -  Maximum number of records to return.
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/userTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
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
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
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

type ApiAllOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Symbol
func (r ApiAllOrdersRequest) Symbol(symbol string) ApiAllOrdersRequest {
	r.symbol = &symbol
	return r
}

// Pair
func (r ApiAllOrdersRequest) Pair(pair string) ApiAllOrdersRequest {
	r.pair = &pair
	return r
}

func (r ApiAllOrdersRequest) OrderId(orderId int64) ApiAllOrdersRequest {
	r.orderId = &orderId
	return r
}

// Start time
func (r ApiAllOrdersRequest) StartTime(startTime int64) ApiAllOrdersRequest {
	r.startTime = &startTime
	return r
}

// End time
func (r ApiAllOrdersRequest) EndTime(endTime int64) ApiAllOrdersRequest {
	r.endTime = &endTime
	return r
}

// Maximum number of records to return.
func (r ApiAllOrdersRequest) Limit(limit int64) ApiAllOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiAllOrdersRequest) RecvWindow(recvWindow int64) ApiAllOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllOrdersRequest) Execute() (*common.RestApiResponse[models.AllOrdersResponse], error) {
	return r.ApiService.AllOrdersExecute(r)
}

/*
AllOrders All Orders (USER_DATA)
Get /dapi/v1/allOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#all-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param pair -  Pair
@param orderId -
@param startTime -  Start time
@param endTime -  End time
@param limit -  Maximum number of records to return.
@param recvWindow -
@return ApiAllOrdersRequest
*/
func (a *TradeAPIService) AllOrders(ctx context.Context) ApiAllOrdersRequest {
	return ApiAllOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AllOrdersResponse
func (a *TradeAPIService) AllOrdersExecute(r ApiAllOrdersRequest) (*common.RestApiResponse[models.AllOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AllOrdersResponse](
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

type ApiAutoCancelAllOpenOrdersRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	countdownTime *int64
	recvWindow    *int64
}

func (r ApiAutoCancelAllOpenOrdersRequest) Symbol(symbol string) ApiAutoCancelAllOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

// countdown time, 1000 for 1 second. 0 to cancel the timer
func (r ApiAutoCancelAllOpenOrdersRequest) CountdownTime(countdownTime int64) ApiAutoCancelAllOpenOrdersRequest {
	r.countdownTime = &countdownTime
	return r
}

func (r ApiAutoCancelAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiAutoCancelAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAutoCancelAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse], error) {
	return r.ApiService.AutoCancelAllOpenOrdersExecute(r)
}

/*
AutoCancelAllOpenOrders Auto-Cancel All Open Orders (TRADE)
Post /dapi/v1/countdownCancelAll

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#auto-cancel-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param countdownTime -  countdown time, 1000 for 1 second. 0 to cancel the timer
@param recvWindow -
@return ApiAutoCancelAllOpenOrdersRequest
*/
func (a *TradeAPIService) AutoCancelAllOpenOrders(ctx context.Context) ApiAutoCancelAllOpenOrdersRequest {
	return ApiAutoCancelAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AutoCancelAllOpenOrdersResponse
func (a *TradeAPIService) AutoCancelAllOpenOrdersExecute(r ApiAutoCancelAllOpenOrdersRequest) (*common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/countdownCancelAll"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.countdownTime == nil {
		return nil, common.ReportError("countdownTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "countdownTime", r.countdownTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AutoCancelAllOpenOrdersResponse](
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

type ApiCancelAlgoOrderRequest struct {
	ctx          context.Context
	ApiService   *TradeAPIService
	algoId       *int64
	clientAlgoId *string
	recvWindow   *int64
}

func (r ApiCancelAlgoOrderRequest) AlgoId(algoId int64) ApiCancelAlgoOrderRequest {
	r.algoId = &algoId
	return r
}

func (r ApiCancelAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiCancelAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

func (r ApiCancelAlgoOrderRequest) RecvWindow(recvWindow int64) ApiCancelAlgoOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAlgoOrderRequest) Execute() (*common.RestApiResponse[models.CancelAlgoOrderResponse], error) {
	return r.ApiService.CancelAlgoOrderExecute(r)
}

/*
CancelAlgoOrder Cancel Algo Order (TRADE)
Delete /dapi/v1/algoOrder

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#cancel-algo-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -
@param clientAlgoId -
@param recvWindow -
@return ApiCancelAlgoOrderRequest
*/
func (a *TradeAPIService) CancelAlgoOrder(ctx context.Context) ApiCancelAlgoOrderRequest {
	return ApiCancelAlgoOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAlgoOrderResponse
func (a *TradeAPIService) CancelAlgoOrderExecute(r ApiCancelAlgoOrderRequest) (*common.RestApiResponse[models.CancelAlgoOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/algoOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	}
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAlgoOrderResponse](
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

type ApiCancelAllAlgoOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCancelAllAlgoOpenOrdersRequest) Symbol(symbol string) ApiCancelAllAlgoOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllAlgoOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllAlgoOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllAlgoOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllAlgoOpenOrdersResponse], error) {
	return r.ApiService.CancelAllAlgoOpenOrdersExecute(r)
}

/*
CancelAllAlgoOpenOrders Cancel All Algo Open Orders (TRADE)
Delete /dapi/v1/algoOpenOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#cancel-all-algo-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCancelAllAlgoOpenOrdersRequest
*/
func (a *TradeAPIService) CancelAllAlgoOpenOrders(ctx context.Context) ApiCancelAllAlgoOpenOrdersRequest {
	return ApiCancelAllAlgoOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllAlgoOpenOrdersResponse
func (a *TradeAPIService) CancelAllAlgoOpenOrdersExecute(r ApiCancelAllAlgoOpenOrdersRequest) (*common.RestApiResponse[models.CancelAllAlgoOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/algoOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllAlgoOpenOrdersResponse](
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

type ApiCancelAllOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

// Symbol
func (r ApiCancelAllOpenOrdersRequest) Symbol(symbol string) ApiCancelAllOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCancelAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCancelAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CancelAllOpenOrdersResponse], error) {
	return r.ApiService.CancelAllOpenOrdersExecute(r)
}

/*
CancelAllOpenOrders Cancel All Open Orders (TRADE)
Delete /dapi/v1/allOpenOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#cancel-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param recvWindow -
@return ApiCancelAllOpenOrdersRequest
*/
func (a *TradeAPIService) CancelAllOpenOrders(ctx context.Context) ApiCancelAllOpenOrdersRequest {
	return ApiCancelAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAllOpenOrdersResponse
func (a *TradeAPIService) CancelAllOpenOrdersExecute(r ApiCancelAllOpenOrdersRequest) (*common.RestApiResponse[models.CancelAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllOpenOrdersResponse](
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

type ApiCancelMultipleOrdersRequest struct {
	ctx                   context.Context
	ApiService            *TradeAPIService
	symbol                *string
	orderIdList           *[]int64
	origClientOrderIdList *[]string
	recvWindow            *int64
}

// Symbol
func (r ApiCancelMultipleOrdersRequest) Symbol(symbol string) ApiCancelMultipleOrdersRequest {
	r.symbol = &symbol
	return r
}

// Order IDs to cancel.
func (r ApiCancelMultipleOrdersRequest) OrderIdList(orderIdList []int64) ApiCancelMultipleOrdersRequest {
	r.orderIdList = &orderIdList
	return r
}

// Original client order IDs to cancel.
func (r ApiCancelMultipleOrdersRequest) OrigClientOrderIdList(origClientOrderIdList []string) ApiCancelMultipleOrdersRequest {
	r.origClientOrderIdList = &origClientOrderIdList
	return r
}

func (r ApiCancelMultipleOrdersRequest) RecvWindow(recvWindow int64) ApiCancelMultipleOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelMultipleOrdersRequest) Execute() (*common.RestApiResponse[models.CancelMultipleOrdersResponse], error) {
	return r.ApiService.CancelMultipleOrdersExecute(r)
}

/*
CancelMultipleOrders Cancel Multiple Orders (TRADE)
Delete /dapi/v1/batchOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#cancel-multiple-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param orderIdList -  Order IDs to cancel.
@param origClientOrderIdList -  Original client order IDs to cancel.
@param recvWindow -
@return ApiCancelMultipleOrdersRequest
*/
func (a *TradeAPIService) CancelMultipleOrders(ctx context.Context) ApiCancelMultipleOrdersRequest {
	return ApiCancelMultipleOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelMultipleOrdersResponse
func (a *TradeAPIService) CancelMultipleOrdersExecute(r ApiCancelMultipleOrdersRequest) (*common.RestApiResponse[models.CancelMultipleOrdersResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/batchOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderIdList != nil {
		t := *r.orderIdList
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderIdList", t, "form", "multi")
	}
	if r.origClientOrderIdList != nil {
		t := *r.origClientOrderIdList
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderIdList", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelMultipleOrdersResponse](
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

type ApiCancelOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Symbol
func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID
func (r ApiCancelOrderRequest) OrderId(orderId int64) ApiCancelOrderRequest {
	r.orderId = &orderId
	return r
}

// Client order ID
func (r ApiCancelOrderRequest) OrigClientOrderId(origClientOrderId string) ApiCancelOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiCancelOrderRequest) RecvWindow(recvWindow int64) ApiCancelOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelOrderRequest) Execute() (*common.RestApiResponse[models.CancelOrderResponse], error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel Order (TRADE)
Delete /dapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#cancel-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param orderId -  Order ID
@param origClientOrderId -  Client order ID
@param recvWindow -
@return ApiCancelOrderRequest
*/
func (a *TradeAPIService) CancelOrder(ctx context.Context) ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelOrderResponse
func (a *TradeAPIService) CancelOrderExecute(r ApiCancelOrderRequest) (*common.RestApiResponse[models.CancelOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order"

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
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelOrderResponse](
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

type ApiChangeInitialLeverageRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	leverage   *int64
	recvWindow *int64
}

// Symbol
func (r ApiChangeInitialLeverageRequest) Symbol(symbol string) ApiChangeInitialLeverageRequest {
	r.symbol = &symbol
	return r
}

// target initial leverage: int from 1 to 125
func (r ApiChangeInitialLeverageRequest) Leverage(leverage int64) ApiChangeInitialLeverageRequest {
	r.leverage = &leverage
	return r
}

func (r ApiChangeInitialLeverageRequest) RecvWindow(recvWindow int64) ApiChangeInitialLeverageRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeInitialLeverageRequest) Execute() (*common.RestApiResponse[models.ChangeInitialLeverageResponse], error) {
	return r.ApiService.ChangeInitialLeverageExecute(r)
}

/*
ChangeInitialLeverage Change Initial Leverage (TRADE)
Post /dapi/v1/leverage

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#change-initial-leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param leverage -  target initial leverage: int from 1 to 125
@param recvWindow -
@return ApiChangeInitialLeverageRequest
*/
func (a *TradeAPIService) ChangeInitialLeverage(ctx context.Context) ApiChangeInitialLeverageRequest {
	return ApiChangeInitialLeverageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeInitialLeverageResponse
func (a *TradeAPIService) ChangeInitialLeverageExecute(r ApiChangeInitialLeverageRequest) (*common.RestApiResponse[models.ChangeInitialLeverageResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/leverage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.leverage == nil {
		return nil, common.ReportError("leverage is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "leverage", r.leverage, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeInitialLeverageResponse](
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

type ApiChangeMarginTypeRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	marginType *models.ChangeMarginTypeMarginTypeParameter
	recvWindow *int64
}

// Symbol
func (r ApiChangeMarginTypeRequest) Symbol(symbol string) ApiChangeMarginTypeRequest {
	r.symbol = &symbol
	return r
}

func (r ApiChangeMarginTypeRequest) MarginType(marginType models.ChangeMarginTypeMarginTypeParameter) ApiChangeMarginTypeRequest {
	r.marginType = &marginType
	return r
}

func (r ApiChangeMarginTypeRequest) RecvWindow(recvWindow int64) ApiChangeMarginTypeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeMarginTypeRequest) Execute() (*common.RestApiResponse[models.ChangeMarginTypeResponse], error) {
	return r.ApiService.ChangeMarginTypeExecute(r)
}

/*
ChangeMarginType Change Margin Type (TRADE)
Post /dapi/v1/marginType

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#change-margin-type

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param marginType -
@param recvWindow -
@return ApiChangeMarginTypeRequest
*/
func (a *TradeAPIService) ChangeMarginType(ctx context.Context) ApiChangeMarginTypeRequest {
	return ApiChangeMarginTypeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeMarginTypeResponse
func (a *TradeAPIService) ChangeMarginTypeExecute(r ApiChangeMarginTypeRequest) (*common.RestApiResponse[models.ChangeMarginTypeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/marginType"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.marginType == nil {
		return nil, common.ReportError("marginType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marginType", r.marginType, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeMarginTypeResponse](
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

type ApiChangePositionModeRequest struct {
	ctx              context.Context
	ApiService       *TradeAPIService
	dualSidePosition *string
	recvWindow       *int64
}

// \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode
func (r ApiChangePositionModeRequest) DualSidePosition(dualSidePosition string) ApiChangePositionModeRequest {
	r.dualSidePosition = &dualSidePosition
	return r
}

func (r ApiChangePositionModeRequest) RecvWindow(recvWindow int64) ApiChangePositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangePositionModeRequest) Execute() (*common.RestApiResponse[models.ChangePositionModeResponse], error) {
	return r.ApiService.ChangePositionModeExecute(r)
}

/*
ChangePositionMode Change Position Mode (TRADE)
Post /dapi/v1/positionSide/dual

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#change-position-mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param dualSidePosition -  \"true\": Hedge Mode; \"false\": One-way Mode
@param recvWindow -
@return ApiChangePositionModeRequest
*/
func (a *TradeAPIService) ChangePositionMode(ctx context.Context) ApiChangePositionModeRequest {
	return ApiChangePositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangePositionModeResponse
func (a *TradeAPIService) ChangePositionModeExecute(r ApiChangePositionModeRequest) (*common.RestApiResponse[models.ChangePositionModeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.dualSidePosition == nil {
		return nil, common.ReportError("dualSidePosition is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dualSidePosition", r.dualSidePosition, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangePositionModeResponse](
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

type ApiCurrentAllAlgoOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	algoType   *string
	symbol     *string
	algoId     *int64
	recvWindow *int64
}

func (r ApiCurrentAllAlgoOpenOrdersRequest) AlgoType(algoType string) ApiCurrentAllAlgoOpenOrdersRequest {
	r.algoType = &algoType
	return r
}

func (r ApiCurrentAllAlgoOpenOrdersRequest) Symbol(symbol string) ApiCurrentAllAlgoOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCurrentAllAlgoOpenOrdersRequest) AlgoId(algoId int64) ApiCurrentAllAlgoOpenOrdersRequest {
	r.algoId = &algoId
	return r
}

func (r ApiCurrentAllAlgoOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCurrentAllAlgoOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCurrentAllAlgoOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CurrentAllAlgoOpenOrdersResponse], error) {
	return r.ApiService.CurrentAllAlgoOpenOrdersExecute(r)
}

/*
CurrentAllAlgoOpenOrders Current All Algo Open Orders (USER_DATA)
Get /dapi/v1/openAlgoOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#current-all-algo-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoType -
@param symbol -
@param algoId -
@param recvWindow -
@return ApiCurrentAllAlgoOpenOrdersRequest
*/
func (a *TradeAPIService) CurrentAllAlgoOpenOrders(ctx context.Context) ApiCurrentAllAlgoOpenOrdersRequest {
	return ApiCurrentAllAlgoOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CurrentAllAlgoOpenOrdersResponse
func (a *TradeAPIService) CurrentAllAlgoOpenOrdersExecute(r ApiCurrentAllAlgoOpenOrdersRequest) (*common.RestApiResponse[models.CurrentAllAlgoOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/openAlgoOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoType", r.algoType, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.algoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CurrentAllAlgoOpenOrdersResponse](
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

type ApiCurrentAllOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	pair       *string
	recvWindow *int64
}

// Symbol. **After CM migration, an invalid &#x60;symbol&#x60; returns &#x60;-1121&#x60; (previously a silent &#x60;200&#x60;).**
func (r ApiCurrentAllOpenOrdersRequest) Symbol(symbol string) ApiCurrentAllOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

// Pair
func (r ApiCurrentAllOpenOrdersRequest) Pair(pair string) ApiCurrentAllOpenOrdersRequest {
	r.pair = &pair
	return r
}

func (r ApiCurrentAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiCurrentAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCurrentAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.CurrentAllOpenOrdersResponse], error) {
	return r.ApiService.CurrentAllOpenOrdersExecute(r)
}

/*
CurrentAllOpenOrders Current All Open Orders (USER_DATA)
Get /dapi/v1/openOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#current-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol. **After CM migration, an invalid `symbol` returns `-1121` (previously a silent `200`).**
@param pair -  Pair
@param recvWindow -
@return ApiCurrentAllOpenOrdersRequest
*/
func (a *TradeAPIService) CurrentAllOpenOrders(ctx context.Context) ApiCurrentAllOpenOrdersRequest {
	return ApiCurrentAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CurrentAllOpenOrdersResponse
func (a *TradeAPIService) CurrentAllOpenOrdersExecute(r ApiCurrentAllOpenOrdersRequest) (*common.RestApiResponse[models.CurrentAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CurrentAllOpenOrdersResponse](
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

type ApiGetOrderModifyHistoryRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	startTime         *int64
	endTime           *int64
	limit             *int64
	recvWindow        *int64
}

// Symbol
func (r ApiGetOrderModifyHistoryRequest) Symbol(symbol string) ApiGetOrderModifyHistoryRequest {
	r.symbol = &symbol
	return r
}

// Order ID
func (r ApiGetOrderModifyHistoryRequest) OrderId(orderId int64) ApiGetOrderModifyHistoryRequest {
	r.orderId = &orderId
	return r
}

// Client order ID
func (r ApiGetOrderModifyHistoryRequest) OrigClientOrderId(origClientOrderId string) ApiGetOrderModifyHistoryRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Timestamp in ms to get modification history from INCLUSIVE
func (r ApiGetOrderModifyHistoryRequest) StartTime(startTime int64) ApiGetOrderModifyHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get modification history until INCLUSIVE
func (r ApiGetOrderModifyHistoryRequest) EndTime(endTime int64) ApiGetOrderModifyHistoryRequest {
	r.endTime = &endTime
	return r
}

// Maximum number of records to return.
func (r ApiGetOrderModifyHistoryRequest) Limit(limit int64) ApiGetOrderModifyHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetOrderModifyHistoryRequest) RecvWindow(recvWindow int64) ApiGetOrderModifyHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOrderModifyHistoryRequest) Execute() (*common.RestApiResponse[models.GetOrderModifyHistoryResponse], error) {
	return r.ApiService.GetOrderModifyHistoryExecute(r)
}

/*
GetOrderModifyHistory Get Order Modify History (USER_DATA)
Get /dapi/v1/orderAmendment

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#get-order-modify-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param orderId -  Order ID
@param origClientOrderId -  Client order ID
@param startTime -  Timestamp in ms to get modification history from INCLUSIVE
@param endTime -  Timestamp in ms to get modification history until INCLUSIVE
@param limit -  Maximum number of records to return.
@param recvWindow -
@return ApiGetOrderModifyHistoryRequest
*/
func (a *TradeAPIService) GetOrderModifyHistory(ctx context.Context) ApiGetOrderModifyHistoryRequest {
	return ApiGetOrderModifyHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrderModifyHistoryResponse
func (a *TradeAPIService) GetOrderModifyHistoryExecute(r ApiGetOrderModifyHistoryRequest) (*common.RestApiResponse[models.GetOrderModifyHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/orderAmendment"

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

	resp, err := SendRequest[models.GetOrderModifyHistoryResponse](
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

type ApiGetPositionMarginChangeHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	type_      *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetPositionMarginChangeHistoryRequest) Symbol(symbol string) ApiGetPositionMarginChangeHistoryRequest {
	r.symbol = &symbol
	return r
}

// 1: Add position margin,2: Reduce position margin
func (r ApiGetPositionMarginChangeHistoryRequest) Type(type_ int64) ApiGetPositionMarginChangeHistoryRequest {
	r.type_ = &type_
	return r
}

// Start time
func (r ApiGetPositionMarginChangeHistoryRequest) StartTime(startTime int64) ApiGetPositionMarginChangeHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time
func (r ApiGetPositionMarginChangeHistoryRequest) EndTime(endTime int64) ApiGetPositionMarginChangeHistoryRequest {
	r.endTime = &endTime
	return r
}

// Maximum number of records to return.
func (r ApiGetPositionMarginChangeHistoryRequest) Limit(limit int64) ApiGetPositionMarginChangeHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPositionMarginChangeHistoryRequest) RecvWindow(recvWindow int64) ApiGetPositionMarginChangeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPositionMarginChangeHistoryRequest) Execute() (*common.RestApiResponse[models.GetPositionMarginChangeHistoryResponse], error) {
	return r.ApiService.GetPositionMarginChangeHistoryExecute(r)
}

/*
GetPositionMarginChangeHistory Get Position Margin Change History (TRADE)
Get /dapi/v1/positionMargin/history

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#get-position-margin-change-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param type_ -  1: Add position margin,2: Reduce position margin
@param startTime -  Start time
@param endTime -  End time
@param limit -  Maximum number of records to return.
@param recvWindow -
@return ApiGetPositionMarginChangeHistoryRequest
*/
func (a *TradeAPIService) GetPositionMarginChangeHistory(ctx context.Context) ApiGetPositionMarginChangeHistoryRequest {
	return ApiGetPositionMarginChangeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPositionMarginChangeHistoryResponse
func (a *TradeAPIService) GetPositionMarginChangeHistoryExecute(r ApiGetPositionMarginChangeHistoryRequest) (*common.RestApiResponse[models.GetPositionMarginChangeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/positionMargin/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

	resp, err := SendRequest[models.GetPositionMarginChangeHistoryResponse](
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

type ApiModifyIsolatedPositionMarginRequest struct {
	ctx          context.Context
	ApiService   *TradeAPIService
	symbol       *string
	amount       *float64
	type_        *int64
	positionSide *models.NewOrderPositionSideParameter
	recvWindow   *int64
}

// Symbol
func (r ApiModifyIsolatedPositionMarginRequest) Symbol(symbol string) ApiModifyIsolatedPositionMarginRequest {
	r.symbol = &symbol
	return r
}

// Margin asset
func (r ApiModifyIsolatedPositionMarginRequest) Amount(amount float64) ApiModifyIsolatedPositionMarginRequest {
	r.amount = &amount
	return r
}

// 1: Add position margin,2: Reduce position margin
func (r ApiModifyIsolatedPositionMarginRequest) Type(type_ int64) ApiModifyIsolatedPositionMarginRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode.
func (r ApiModifyIsolatedPositionMarginRequest) PositionSide(positionSide models.NewOrderPositionSideParameter) ApiModifyIsolatedPositionMarginRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiModifyIsolatedPositionMarginRequest) RecvWindow(recvWindow int64) ApiModifyIsolatedPositionMarginRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyIsolatedPositionMarginRequest) Execute() (*common.RestApiResponse[models.ModifyIsolatedPositionMarginResponse], error) {
	return r.ApiService.ModifyIsolatedPositionMarginExecute(r)
}

/*
ModifyIsolatedPositionMargin Modify Isolated Position Margin (TRADE)
Post /dapi/v1/positionMargin

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#modify-isolated-position-margin

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param amount -  Margin asset
@param type_ -  1: Add position margin,2: Reduce position margin
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode.
@param recvWindow -
@return ApiModifyIsolatedPositionMarginRequest
*/
func (a *TradeAPIService) ModifyIsolatedPositionMargin(ctx context.Context) ApiModifyIsolatedPositionMarginRequest {
	return ApiModifyIsolatedPositionMarginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifyIsolatedPositionMarginResponse
func (a *TradeAPIService) ModifyIsolatedPositionMarginExecute(r ApiModifyIsolatedPositionMarginRequest) (*common.RestApiResponse[models.ModifyIsolatedPositionMarginResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/positionMargin"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyIsolatedPositionMarginResponse](
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

type ApiModifyMultipleOrdersRequest struct {
	ctx         context.Context
	ApiService  *TradeAPIService
	batchOrders *[]models.ModifyMultipleOrdersBatchOrdersParameterInner
	recvWindow  *int64
}

// order list. Max 5 orders
func (r ApiModifyMultipleOrdersRequest) BatchOrders(batchOrders []models.ModifyMultipleOrdersBatchOrdersParameterInner) ApiModifyMultipleOrdersRequest {
	r.batchOrders = &batchOrders
	return r
}

func (r ApiModifyMultipleOrdersRequest) RecvWindow(recvWindow int64) ApiModifyMultipleOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifyMultipleOrdersRequest) Execute() (*common.RestApiResponse[models.ModifyMultipleOrdersResponse], error) {
	return r.ApiService.ModifyMultipleOrdersExecute(r)
}

/*
ModifyMultipleOrders Modify Multiple Orders (TRADE)
Put /dapi/v1/batchOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#modify-multiple-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param batchOrders -  order list. Max 5 orders
@param recvWindow -
@return ApiModifyMultipleOrdersRequest
*/
func (a *TradeAPIService) ModifyMultipleOrders(ctx context.Context) ApiModifyMultipleOrdersRequest {
	return ApiModifyMultipleOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifyMultipleOrdersResponse
func (a *TradeAPIService) ModifyMultipleOrdersExecute(r ApiModifyMultipleOrdersRequest) (*common.RestApiResponse[models.ModifyMultipleOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/batchOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.batchOrders == nil {
		return nil, common.ReportError("batchOrders is required and must be specified")
	}

	{
		t := *r.batchOrders
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "batchOrders", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyMultipleOrdersResponse](
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

type ApiModifyOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	side              *models.NewAlgoOrderSideParameter
	orderId           *int64
	origClientOrderId *string
	quantity          *float64
	price             *float64
	priceMatch        *models.NewAlgoOrderPriceMatchParameter
	modifyId          *int64
	recvWindow        *int64
}

// Symbol
func (r ApiModifyOrderRequest) Symbol(symbol string) ApiModifyOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiModifyOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiModifyOrderRequest {
	r.side = &side
	return r
}

// Order ID
func (r ApiModifyOrderRequest) OrderId(orderId int64) ApiModifyOrderRequest {
	r.orderId = &orderId
	return r
}

// Client order ID
func (r ApiModifyOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// Order quantity, cannot be sent with &#x60;closePosition&#x3D;true&#x60;. **After CM migration, this parameter becomes mandatory** (must be sent together with &#x60;price&#x60;).
func (r ApiModifyOrderRequest) Quantity(quantity float64) ApiModifyOrderRequest {
	r.quantity = &quantity
	return r
}

// Order price. **After CM migration, this parameter becomes mandatory** (must be sent together with &#x60;quantity&#x60;).
func (r ApiModifyOrderRequest) Price(price float64) ApiModifyOrderRequest {
	r.price = &price
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiModifyOrderRequest {
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

func (r ApiModifyOrderRequest) Execute() (*common.RestApiResponse[models.ModifyOrderResponse], error) {
	return r.ApiService.ModifyOrderExecute(r)
}

/*
ModifyOrder Modify Order (TRADE)
Put /dapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#modify-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param side -
@param orderId -  Order ID
@param origClientOrderId -  Client order ID
@param quantity -  Order quantity, cannot be sent with `closePosition=true`. **After CM migration, this parameter becomes mandatory** (must be sent together with `price`).
@param price -  Order price. **After CM migration, this parameter becomes mandatory** (must be sent together with `quantity`).
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price`
@param modifyId -  User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness.
@param recvWindow -
@return ApiModifyOrderRequest
*/
func (a *TradeAPIService) ModifyOrder(ctx context.Context) ApiModifyOrderRequest {
	return ApiModifyOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifyOrderResponse
func (a *TradeAPIService) ModifyOrderExecute(r ApiModifyOrderRequest) (*common.RestApiResponse[models.ModifyOrderResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.modifyId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "modifyId", r.modifyId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyOrderResponse](
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

type ApiNewAlgoOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	algoType                *models.NewAlgoOrderAlgoTypeParameter
	symbol                  *string
	side                    *models.NewAlgoOrderSideParameter
	type_                   *models.NewAlgoOrderTypeParameter
	positionSide            *string
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float64
	price                   *float64
	triggerPrice            *float64
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceMatch              *models.NewAlgoOrderPriceMatchParameter
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

func (r ApiNewAlgoOrderRequest) AlgoType(algoType models.NewAlgoOrderAlgoTypeParameter) ApiNewAlgoOrderRequest {
	r.algoType = &algoType
	return r
}

func (r ApiNewAlgoOrderRequest) Symbol(symbol string) ApiNewAlgoOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewAlgoOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiNewAlgoOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewAlgoOrderRequest) Type(type_ models.NewAlgoOrderTypeParameter) ApiNewAlgoOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewAlgoOrderRequest) PositionSide(positionSide string) ApiNewAlgoOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewAlgoOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewAlgoOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewAlgoOrderRequest) Quantity(quantity float64) ApiNewAlgoOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiNewAlgoOrderRequest) Price(price float64) ApiNewAlgoOrderRequest {
	r.price = &price
	return r
}

// Trigger price
func (r ApiNewAlgoOrderRequest) TriggerPrice(triggerPrice float64) ApiNewAlgoOrderRequest {
	r.triggerPrice = &triggerPrice
	return r
}

func (r ApiNewAlgoOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewAlgoOrderRequest {
	r.workingType = &workingType
	return r
}

// only available for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewAlgoOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiNewAlgoOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewAlgoOrderRequest) ClosePosition(closePosition models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60; order. when price reaches the triggerPrice ，the difference rate between \&quot;MARK_PRICE\&quot; and \&quot;CONTRACT_PRICE\&quot; cannot be larger than the Price Protection Threshold of the symbol.
func (r ApiNewAlgoOrderRequest) PriceProtect(priceProtect models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewAlgoOrderRequest) ReduceOnly(reduceOnly models.NewAlgoOrderClosePositionParameter) ApiNewAlgoOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiNewAlgoOrderRequest) ActivatePrice(activatePrice float64) ApiNewAlgoOrderRequest {
	r.activatePrice = &activatePrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders
func (r ApiNewAlgoOrderRequest) CallbackRate(callbackRate float64) ApiNewAlgoOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiNewAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiNewAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

func (r ApiNewAlgoOrderRequest) NewOrderRespType(newOrderRespType models.NewAlgoOrderNewOrderRespTypeParameter) ApiNewAlgoOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers / &#x60;EXPIRE_MAKER&#x60;: expire maker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiNewAlgoOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiNewAlgoOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

func (r ApiNewAlgoOrderRequest) RecvWindow(recvWindow int64) ApiNewAlgoOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNewAlgoOrderRequest) Execute() (*common.RestApiResponse[models.NewAlgoOrderResponse], error) {
	return r.ApiService.NewAlgoOrderExecute(r)
}

/*
NewAlgoOrder New Algo Order (TRADE)
Post /dapi/v1/algoOrder

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#new-algo-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoType -
@param symbol -
@param side -
@param type_ -
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param timeInForce -
@param quantity -  quantity measured by contract number, Cannot be sent with `closePosition`=`true`(Close-All)
@param price -
@param triggerPrice -  Trigger price
@param workingType -
@param priceMatch -  only available for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price`
@param closePosition -  Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.
@param priceProtect -  Used with `STOP_MARKET` or `TAKE_PROFIT_MARKET` order. when price reaches the triggerPrice ，the difference rate between \"MARK_PRICE\" and \"CONTRACT_PRICE\" cannot be larger than the Price Protection Threshold of the symbol.
@param reduceOnly -  Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`
@param activatePrice -  Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders
@param clientAlgoId -  A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`
@param newOrderRespType -
@param selfTradePreventionMode -  `EXPIRE_TAKER`: expire taker order when STP triggers / `EXPIRE_MAKER`: expire maker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers; default `NONE`
@param recvWindow -
@return ApiNewAlgoOrderRequest
*/
func (a *TradeAPIService) NewAlgoOrder(ctx context.Context) ApiNewAlgoOrderRequest {
	return ApiNewAlgoOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NewAlgoOrderResponse
func (a *TradeAPIService) NewAlgoOrderExecute(r ApiNewAlgoOrderRequest) (*common.RestApiResponse[models.NewAlgoOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/algoOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoType == nil {
		return nil, common.ReportError("algoType is required and must be specified")
	}

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoType", r.algoType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.triggerPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "triggerPrice", r.triggerPrice, "form", "")
	}
	if r.workingType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.closePosition != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "closePosition", r.closePosition, "form", "")
	}
	if r.priceProtect != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceProtect", r.priceProtect, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.activatePrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activatePrice", r.activatePrice, "form", "")
	}
	if r.callbackRate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "callbackRate", r.callbackRate, "form", "")
	}
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
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

	resp, err := SendRequest[models.NewAlgoOrderResponse](
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
	side                    *models.NewAlgoOrderSideParameter
	type_                   *models.NewOrderTypeParameter
	positionSide            *models.NewOrderPositionSideParameter
	reduceOnly              *models.NewAlgoOrderClosePositionParameter
	quantity                *float64
	price                   *float64
	newClientOrderId        *string
	stopPrice               *float64
	closePosition           *string
	activationPrice         *float64
	callbackRate            *float64
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceProtect            *models.NewAlgoOrderClosePositionParameter
	newOrderRespType        *models.NewAlgoOrderNewOrderRespTypeParameter
	priceMatch              *models.NewAlgoOrderPriceMatchParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	recvWindow              *int64
}

// Symbol
func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNewOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

// **After CM migration, stop-type values (&#x60;STOP&#x60;, &#x60;STOP_MARKET&#x60;, &#x60;TAKE_PROFIT&#x60;, &#x60;TAKE_PROFIT_MARKET&#x60;, &#x60;TRAILING_STOP_MARKET&#x60;) are no longer accepted by this endpoint and will return &#x60;-4120&#x60;. Use the new &#x60;/dapi/v1/algoOrder&#x60; endpoint instead.**
func (r ApiNewOrderRequest) Type(type_ models.NewOrderTypeParameter) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. Cannot be sent in Hedge Mode; cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;(Close-All)
func (r ApiNewOrderRequest) ReduceOnly(reduceOnly models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// quantity measured by contract number, Cannot be sent with &#x60;closePosition&#x60;&#x3D;&#x60;true&#x60;
func (r ApiNewOrderRequest) Quantity(quantity float64) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// Order price
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

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All,used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiNewOrderRequest) ClosePosition(closePosition string) ApiNewOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiNewOrderRequest) ActivationPrice(activationPrice float64) ApiNewOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 10 where 1 for 1%
func (r ApiNewOrderRequest) CallbackRate(callbackRate float64) ApiNewOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// &#39;stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;.
func (r ApiNewOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiNewOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiNewOrderRequest) PriceProtect(priceProtect models.NewAlgoOrderClosePositionParameter) ApiNewOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.NewAlgoOrderNewOrderRespTypeParameter) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiNewOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;EXPIRE_TAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;: expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;: expire both orders when STP triggers
func (r ApiNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiNewOrderRequest {
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
Post /dapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#new-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param side -
@param type_ -  **After CM migration, stop-type values (`STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`) are no longer accepted by this endpoint and will return `-4120`. Use the new `/dapi/v1/algoOrder` endpoint instead.**
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param reduceOnly -  \"true\" or \"false\". Cannot be sent in Hedge Mode; cannot be sent with `closePosition`=`true`(Close-All)
@param quantity -  quantity measured by contract number, Cannot be sent with `closePosition`=`true`
@param price -  Order price
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`
@param stopPrice -  Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param closePosition -  `true`, `false`；Close-All,used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.
@param activationPrice -  Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 10 where 1 for 1%
@param timeInForce -
@param workingType -  'stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\".
@param priceProtect -  \"true\" or \"false\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param newOrderRespType -
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; Can't be passed together with `price`
@param selfTradePreventionMode -  `EXPIRE_TAKER`: expire taker order when STP triggers/ `EXPIRE_MAKER`: expire taker order when STP triggers/ `EXPIRE_BOTH`: expire both orders when STP triggers
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order"

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
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.closePosition != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "closePosition", r.closePosition, "form", "")
	}
	if r.activationPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activationPrice", r.activationPrice, "form", "")
	}
	if r.callbackRate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "callbackRate", r.callbackRate, "form", "")
	}
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.workingType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	}
	if r.priceProtect != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceProtect", r.priceProtect, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
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

type ApiPlaceMultipleOrdersRequest struct {
	ctx         context.Context
	ApiService  *TradeAPIService
	batchOrders *[]models.PlaceMultipleOrdersBatchOrdersParameterInner
	recvWindow  *int64
}

// order list. Max 5 orders
func (r ApiPlaceMultipleOrdersRequest) BatchOrders(batchOrders []models.PlaceMultipleOrdersBatchOrdersParameterInner) ApiPlaceMultipleOrdersRequest {
	r.batchOrders = &batchOrders
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
Post /dapi/v1/batchOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#place-multiple-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param batchOrders -  order list. Max 5 orders
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/batchOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.batchOrders == nil {
		return nil, common.ReportError("batchOrders is required and must be specified")
	}

	{
		t := *r.batchOrders
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "batchOrders", t, "form", "multi")
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

type ApiPositionAdlQuantileEstimationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiPositionAdlQuantileEstimationRequest) Symbol(symbol string) ApiPositionAdlQuantileEstimationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiPositionAdlQuantileEstimationRequest) RecvWindow(recvWindow int64) ApiPositionAdlQuantileEstimationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionAdlQuantileEstimationRequest) Execute() (*common.RestApiResponse[models.PositionAdlQuantileEstimationResponse], error) {
	return r.ApiService.PositionAdlQuantileEstimationExecute(r)
}

/*
PositionAdlQuantileEstimation Position ADL Quantile Estimation (USER_DATA)
Get /dapi/v1/adlQuantile

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#position-adl-quantile-estimation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiPositionAdlQuantileEstimationRequest
*/
func (a *TradeAPIService) PositionAdlQuantileEstimation(ctx context.Context) ApiPositionAdlQuantileEstimationRequest {
	return ApiPositionAdlQuantileEstimationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PositionAdlQuantileEstimationResponse
func (a *TradeAPIService) PositionAdlQuantileEstimationExecute(r ApiPositionAdlQuantileEstimationRequest) (*common.RestApiResponse[models.PositionAdlQuantileEstimationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/adlQuantile"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PositionAdlQuantileEstimationResponse](
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

type ApiPositionInformationRequest struct {
	ctx         context.Context
	ApiService  *TradeAPIService
	marginAsset *string
	pair        *string
	recvWindow  *int64
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

func (r ApiPositionInformationRequest) Execute() (*common.RestApiResponse[models.PositionInformationResponse], error) {
	return r.ApiService.PositionInformationExecute(r)
}

/*
PositionInformation Position Information (USER_DATA)
Get /dapi/v1/positionRisk

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#position-information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param marginAsset -
@param pair -
@param recvWindow -
@return ApiPositionInformationRequest
*/
func (a *TradeAPIService) PositionInformation(ctx context.Context) ApiPositionInformationRequest {
	return ApiPositionInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PositionInformationResponse
func (a *TradeAPIService) PositionInformationExecute(r ApiPositionInformationRequest) (*common.RestApiResponse[models.PositionInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.marginAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marginAsset", r.marginAsset, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PositionInformationResponse](
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

type ApiQueryAlgoOrderRequest struct {
	ctx          context.Context
	ApiService   *TradeAPIService
	algoId       *int64
	clientAlgoId *string
	recvWindow   *int64
}

// Order ID
func (r ApiQueryAlgoOrderRequest) AlgoId(algoId int64) ApiQueryAlgoOrderRequest {
	r.algoId = &algoId
	return r
}

// Client order ID
func (r ApiQueryAlgoOrderRequest) ClientAlgoId(clientAlgoId string) ApiQueryAlgoOrderRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

func (r ApiQueryAlgoOrderRequest) RecvWindow(recvWindow int64) ApiQueryAlgoOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAlgoOrderRequest) Execute() (*common.RestApiResponse[models.QueryAlgoOrderResponse], error) {
	return r.ApiService.QueryAlgoOrderExecute(r)
}

/*
QueryAlgoOrder Query Algo Order (USER_DATA)
Get /dapi/v1/algoOrder

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#query-algo-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  Order ID
@param clientAlgoId -  Client order ID
@param recvWindow -
@return ApiQueryAlgoOrderRequest
*/
func (a *TradeAPIService) QueryAlgoOrder(ctx context.Context) ApiQueryAlgoOrderRequest {
	return ApiQueryAlgoOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAlgoOrderResponse
func (a *TradeAPIService) QueryAlgoOrderExecute(r ApiQueryAlgoOrderRequest) (*common.RestApiResponse[models.QueryAlgoOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/algoOrder"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	}
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAlgoOrderResponse](
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

type ApiQueryAllAlgoOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	algoId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Symbol
func (r ApiQueryAllAlgoOrdersRequest) Symbol(symbol string) ApiQueryAllAlgoOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) AlgoId(algoId int64) ApiQueryAllAlgoOrdersRequest {
	r.algoId = &algoId
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) StartTime(startTime int64) ApiQueryAllAlgoOrdersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) EndTime(endTime int64) ApiQueryAllAlgoOrdersRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) Limit(limit int64) ApiQueryAllAlgoOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) RecvWindow(recvWindow int64) ApiQueryAllAlgoOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryAllAlgoOrdersRequest) Execute() (*common.RestApiResponse[models.QueryAllAlgoOrdersResponse], error) {
	return r.ApiService.QueryAllAlgoOrdersExecute(r)
}

/*
QueryAllAlgoOrders Query All Algo Orders (USER_DATA)
Get /dapi/v1/allAlgoOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#query-all-algo-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param algoId -
@param startTime -
@param endTime -
@param limit -
@param recvWindow -
@return ApiQueryAllAlgoOrdersRequest
*/
func (a *TradeAPIService) QueryAllAlgoOrders(ctx context.Context) ApiQueryAllAlgoOrdersRequest {
	return ApiQueryAllAlgoOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryAllAlgoOrdersResponse
func (a *TradeAPIService) QueryAllAlgoOrdersExecute(r ApiQueryAllAlgoOrdersRequest) (*common.RestApiResponse[models.QueryAllAlgoOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/allAlgoOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.algoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
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

	resp, err := SendRequest[models.QueryAllAlgoOrdersResponse](
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

type ApiQueryCurrentOpenOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Symbol
func (r ApiQueryCurrentOpenOrderRequest) Symbol(symbol string) ApiQueryCurrentOpenOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID
func (r ApiQueryCurrentOpenOrderRequest) OrderId(orderId int64) ApiQueryCurrentOpenOrderRequest {
	r.orderId = &orderId
	return r
}

// Client order ID
func (r ApiQueryCurrentOpenOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryCurrentOpenOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryCurrentOpenOrderRequest) RecvWindow(recvWindow int64) ApiQueryCurrentOpenOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentOpenOrderRequest) Execute() (*common.RestApiResponse[models.QueryCurrentOpenOrderResponse], error) {
	return r.ApiService.QueryCurrentOpenOrderExecute(r)
}

/*
QueryCurrentOpenOrder Query Current Open Order (USER_DATA)
Get /dapi/v1/openOrder

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#query-current-open-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param orderId -  Order ID
@param origClientOrderId -  Client order ID
@param recvWindow -
@return ApiQueryCurrentOpenOrderRequest
*/
func (a *TradeAPIService) QueryCurrentOpenOrder(ctx context.Context) ApiQueryCurrentOpenOrderRequest {
	return ApiQueryCurrentOpenOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentOpenOrderResponse
func (a *TradeAPIService) QueryCurrentOpenOrderExecute(r ApiQueryCurrentOpenOrderRequest) (*common.RestApiResponse[models.QueryCurrentOpenOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/openOrder"

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
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentOpenOrderResponse](
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

type ApiQueryOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

// Symbol
func (r ApiQueryOrderRequest) Symbol(symbol string) ApiQueryOrderRequest {
	r.symbol = &symbol
	return r
}

// Order ID
func (r ApiQueryOrderRequest) OrderId(orderId int64) ApiQueryOrderRequest {
	r.orderId = &orderId
	return r
}

// Client order ID
func (r ApiQueryOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryOrderRequest) RecvWindow(recvWindow int64) ApiQueryOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOrderRequest) Execute() (*common.RestApiResponse[models.QueryOrderResponse], error) {
	return r.ApiService.QueryOrderExecute(r)
}

/*
QueryOrder Query Order (USER_DATA)
Get /dapi/v1/order

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#query-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol
@param orderId -  Order ID
@param origClientOrderId -  Client order ID
@param recvWindow -
@return ApiQueryOrderRequest
*/
func (a *TradeAPIService) QueryOrder(ctx context.Context) ApiQueryOrderRequest {
	return ApiQueryOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryOrderResponse
func (a *TradeAPIService) QueryOrderExecute(r ApiQueryOrderRequest) (*common.RestApiResponse[models.QueryOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order"

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
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryOrderResponse](
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

type ApiUsersForceOrdersRequest struct {
	ctx           context.Context
	ApiService    *TradeAPIService
	symbol        *string
	autoCloseType *models.UsersForceOrdersAutoCloseTypeParameter
	startTime     *int64
	endTime       *int64
	limit         *int64
	recvWindow    *int64
}

func (r ApiUsersForceOrdersRequest) Symbol(symbol string) ApiUsersForceOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUsersForceOrdersRequest) AutoCloseType(autoCloseType models.UsersForceOrdersAutoCloseTypeParameter) ApiUsersForceOrdersRequest {
	r.autoCloseType = &autoCloseType
	return r
}

func (r ApiUsersForceOrdersRequest) StartTime(startTime int64) ApiUsersForceOrdersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiUsersForceOrdersRequest) EndTime(endTime int64) ApiUsersForceOrdersRequest {
	r.endTime = &endTime
	return r
}

// Maximum number of records to return.
func (r ApiUsersForceOrdersRequest) Limit(limit int64) ApiUsersForceOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiUsersForceOrdersRequest) RecvWindow(recvWindow int64) ApiUsersForceOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUsersForceOrdersRequest) Execute() (*common.RestApiResponse[models.UsersForceOrdersResponse], error) {
	return r.ApiService.UsersForceOrdersExecute(r)
}

/*
UsersForceOrders User's Force Orders (USER_DATA)
Get /dapi/v1/forceOrders

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/rest-api/trade#users-force-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param autoCloseType -
@param startTime -
@param endTime -
@param limit -  Maximum number of records to return.
@param recvWindow -
@return ApiUsersForceOrdersRequest
*/
func (a *TradeAPIService) UsersForceOrders(ctx context.Context) ApiUsersForceOrdersRequest {
	return ApiUsersForceOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UsersForceOrdersResponse
func (a *TradeAPIService) UsersForceOrdersExecute(r ApiUsersForceOrdersRequest) (*common.RestApiResponse[models.UsersForceOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/forceOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.autoCloseType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoCloseType", r.autoCloseType, "form", "")
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

	resp, err := SendRequest[models.UsersForceOrdersResponse](
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
