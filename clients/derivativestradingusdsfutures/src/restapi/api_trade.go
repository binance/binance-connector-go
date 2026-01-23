/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiAccountTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiAccountTradeListRequest) Symbol(symbol string) ApiAccountTradeListRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAccountTradeListRequest) OrderId(orderId int64) ApiAccountTradeListRequest {
	r.orderId = &orderId
	return r
}

func (r ApiAccountTradeListRequest) StartTime(startTime int64) ApiAccountTradeListRequest {
	r.startTime = &startTime
	return r
}

func (r ApiAccountTradeListRequest) EndTime(endTime int64) ApiAccountTradeListRequest {
	r.endTime = &endTime
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiAccountTradeListRequest) FromId(fromId int64) ApiAccountTradeListRequest {
	r.fromId = &fromId
	return r
}

// Default 100; max 1000
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
Get /fapi/v1/userTrades

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -
@param endTime -
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/userTrades"

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
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAllOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiAllOrdersRequest) Symbol(symbol string) ApiAllOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAllOrdersRequest) OrderId(orderId int64) ApiAllOrdersRequest {
	r.orderId = &orderId
	return r
}

func (r ApiAllOrdersRequest) StartTime(startTime int64) ApiAllOrdersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiAllOrdersRequest) EndTime(endTime int64) ApiAllOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
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
Get /fapi/v1/allOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/All-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/allOrders"

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

	resp, err := SendRequest[models.AllOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
Post /fapi/v1/countdownCancelAll

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Auto-Cancel-All-Open-Orders

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/countdownCancelAll"

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

	resp, err := SendRequest[models.AutoCancelAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
Delete /fapi/v1/algoOrder

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-Algo-Order

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/algoOrder"

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

	resp, err := SendRequest[models.CancelAlgoOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
Delete /fapi/v1/algoOpenOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-All-Algo-Open-Orders

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/algoOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllAlgoOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
Delete /fapi/v1/allOpenOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-All-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/allOpenOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiCancelMultipleOrdersRequest) Symbol(symbol string) ApiCancelMultipleOrdersRequest {
	r.symbol = &symbol
	return r
}

// max length 10 &lt;br /&gt; e.g. [1234567,2345678]
func (r ApiCancelMultipleOrdersRequest) OrderIdList(orderIdList []int64) ApiCancelMultipleOrdersRequest {
	r.orderIdList = &orderIdList
	return r
}

// max length 10&lt;br /&gt; e.g. [\&quot;my_id_1\&quot;,\&quot;my_id_2\&quot;], encode the double quotes. No space after comma.
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
Delete /fapi/v1/batchOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-Multiple-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderIdList -  max length 10 <br /> e.g. [1234567,2345678]
@param origClientOrderIdList -  max length 10<br /> e.g. [\"my_id_1\",\"my_id_2\"], encode the double quotes. No space after comma.
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/batchOrders"

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

	resp, err := SendRequest[models.CancelMultipleOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiCancelOrderRequest) Symbol(symbol string) ApiCancelOrderRequest {
	r.symbol = &symbol
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

func (r ApiCancelOrderRequest) Execute() (*common.RestApiResponse[models.CancelOrderResponse], error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel Order (TRADE)
Delete /fapi/v1/order

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order"

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

	resp, err := SendRequest[models.CancelOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
ChangeInitialLeverage Change Initial Leverage(TRADE)
Post /fapi/v1/leverage

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Initial-Leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/leverage"

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

	resp, err := SendRequest[models.ChangeInitialLeverageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiChangeMarginTypeRequest) Symbol(symbol string) ApiChangeMarginTypeRequest {
	r.symbol = &symbol
	return r
}

// ISOLATED, CROSSED
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
ChangeMarginType Change Margin Type(TRADE)
Post /fapi/v1/marginType

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Margin-Type

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param marginType -  ISOLATED, CROSSED
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/marginType"

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

	resp, err := SendRequest[models.ChangeMarginTypeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiChangeMultiAssetsModeRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	multiAssetsMargin *string
	recvWindow        *int64
}

// \&quot;true\&quot;: Multi-Assets Mode; \&quot;false\&quot;: Single-Asset Mode
func (r ApiChangeMultiAssetsModeRequest) MultiAssetsMargin(multiAssetsMargin string) ApiChangeMultiAssetsModeRequest {
	r.multiAssetsMargin = &multiAssetsMargin
	return r
}

func (r ApiChangeMultiAssetsModeRequest) RecvWindow(recvWindow int64) ApiChangeMultiAssetsModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeMultiAssetsModeRequest) Execute() (*common.RestApiResponse[models.ChangeMultiAssetsModeResponse], error) {
	return r.ApiService.ChangeMultiAssetsModeExecute(r)
}

/*
ChangeMultiAssetsMode Change Multi-Assets Mode (TRADE)
Post /fapi/v1/multiAssetsMargin

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Multi-Assets-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param multiAssetsMargin -  \"true\": Multi-Assets Mode; \"false\": Single-Asset Mode
@param recvWindow -
@return ApiChangeMultiAssetsModeRequest
*/
func (a *TradeAPIService) ChangeMultiAssetsMode(ctx context.Context) ApiChangeMultiAssetsModeRequest {
	return ApiChangeMultiAssetsModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeMultiAssetsModeResponse
func (a *TradeAPIService) ChangeMultiAssetsModeExecute(r ApiChangeMultiAssetsModeRequest) (*common.RestApiResponse[models.ChangeMultiAssetsModeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/multiAssetsMargin"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.multiAssetsMargin == nil {
		return nil, common.ReportError("multiAssetsMargin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "multiAssetsMargin", r.multiAssetsMargin, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeMultiAssetsModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
ChangePositionMode Change Position Mode(TRADE)
Post /fapi/v1/positionSide/dual

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Position-Mode

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.dualSidePosition == nil {
		return nil, common.ReportError("dualSidePosition is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dualSidePosition", r.dualSidePosition, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangePositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
Get /fapi/v1/openAlgoOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Current-All-Algo-Open-Orders

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/openAlgoOrders"

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

	resp, err := SendRequest[models.CurrentAllAlgoOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCurrentAllOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCurrentAllOpenOrdersRequest) Symbol(symbol string) ApiCurrentAllOpenOrdersRequest {
	r.symbol = &symbol
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
Get /fapi/v1/openOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Current-All-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CurrentAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesTradfiPerpsContractRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiFuturesTradfiPerpsContractRequest) RecvWindow(recvWindow int64) ApiFuturesTradfiPerpsContractRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesTradfiPerpsContractRequest) Execute() (*common.RestApiResponse[models.FuturesTradfiPerpsContractResponse], error) {
	return r.ApiService.FuturesTradfiPerpsContractExecute(r)
}

/*
FuturesTradfiPerpsContract Futures TradFi Perps Contract(USER_DATA)
Post /fapi/v1/stock/contract

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Futures-TradFi-Perps-Contract

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFuturesTradfiPerpsContractRequest
*/
func (a *TradeAPIService) FuturesTradfiPerpsContract(ctx context.Context) ApiFuturesTradfiPerpsContractRequest {
	return ApiFuturesTradfiPerpsContractRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesTradfiPerpsContractResponse
func (a *TradeAPIService) FuturesTradfiPerpsContractExecute(r ApiFuturesTradfiPerpsContractRequest) (*common.RestApiResponse[models.FuturesTradfiPerpsContractResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/stock/contract"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesTradfiPerpsContractResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiGetOrderModifyHistoryRequest) Symbol(symbol string) ApiGetOrderModifyHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetOrderModifyHistoryRequest) OrderId(orderId int64) ApiGetOrderModifyHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetOrderModifyHistoryRequest) OrigClientOrderId(origClientOrderId string) ApiGetOrderModifyHistoryRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiGetOrderModifyHistoryRequest) StartTime(startTime int64) ApiGetOrderModifyHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetOrderModifyHistoryRequest) EndTime(endTime int64) ApiGetOrderModifyHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
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
Get /fapi/v1/orderAmendment

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Get-Order-Modify-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/orderAmendment"

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

	resp, err := SendRequest[models.GetOrderModifyHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// 1: Add position margin，2: Reduce position margin
func (r ApiGetPositionMarginChangeHistoryRequest) Type(type_ int64) ApiGetPositionMarginChangeHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetPositionMarginChangeHistoryRequest) StartTime(startTime int64) ApiGetPositionMarginChangeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetPositionMarginChangeHistoryRequest) EndTime(endTime int64) ApiGetPositionMarginChangeHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
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
Get /fapi/v1/positionMargin/history

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Get-Position-Margin-Change-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param type_ -  1: Add position margin，2: Reduce position margin
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/positionMargin/history"

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

	resp, err := SendRequest[models.GetPositionMarginChangeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiModifyIsolatedPositionMarginRequest struct {
	ctx          context.Context
	ApiService   *TradeAPIService
	symbol       *string
	amount       *float32
	type_        *string
	positionSide *models.NewAlgoOrderPositionSideParameter
	recvWindow   *int64
}

func (r ApiModifyIsolatedPositionMarginRequest) Symbol(symbol string) ApiModifyIsolatedPositionMarginRequest {
	r.symbol = &symbol
	return r
}

func (r ApiModifyIsolatedPositionMarginRequest) Amount(amount float32) ApiModifyIsolatedPositionMarginRequest {
	r.amount = &amount
	return r
}

func (r ApiModifyIsolatedPositionMarginRequest) Type(type_ string) ApiModifyIsolatedPositionMarginRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode.
func (r ApiModifyIsolatedPositionMarginRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiModifyIsolatedPositionMarginRequest {
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
ModifyIsolatedPositionMargin Modify Isolated Position Margin(TRADE)
Post /fapi/v1/positionMargin

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Modify-Isolated-Position-Margin

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param amount -
@param type_ -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/positionMargin"

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

	resp, err := SendRequest[models.ModifyIsolatedPositionMarginResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
ModifyMultipleOrders Modify Multiple Orders(TRADE)
Put /fapi/v1/batchOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Modify-Multiple-Orders

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/batchOrders"

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

	resp, err := SendRequest[models.ModifyMultipleOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	quantity          *float32
	price             *float32
	orderId           *int64
	origClientOrderId *string
	priceMatch        *models.NewAlgoOrderPriceMatchParameter
	recvWindow        *int64
}

func (r ApiModifyOrderRequest) Symbol(symbol string) ApiModifyOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiModifyOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiModifyOrderRequest {
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

func (r ApiModifyOrderRequest) OrderId(orderId int64) ApiModifyOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiModifyOrderRequest) OrigClientOrderId(origClientOrderId string) ApiModifyOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiModifyOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiModifyOrderRequest {
	r.priceMatch = &priceMatch
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
Put /fapi/v1/order

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Modify-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -  `SELL`, `BUY`
@param quantity -  Order quantity, cannot be sent with `closePosition=true`
@param price -
@param orderId -
@param origClientOrderId -
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order"

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

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.priceMatch != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "priceMatch", r.priceMatch, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifyOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNewAlgoOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	algoType                *string
	symbol                  *string
	side                    *models.NewAlgoOrderSideParameter
	type_                   *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float32
	price                   *float32
	triggerPrice            *float32
	workingType             *models.NewAlgoOrderWorkingTypeParameter
	priceMatch              *models.NewAlgoOrderPriceMatchParameter
	closePosition           *string
	priceProtect            *string
	reduceOnly              *string
	activatePrice           *float32
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
func (r ApiNewAlgoOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiNewAlgoOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewAlgoOrderRequest) Type(type_ string) ApiNewAlgoOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode.
func (r ApiNewAlgoOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewAlgoOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewAlgoOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewAlgoOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

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
func (r ApiNewAlgoOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiNewAlgoOrderRequest {
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

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode
func (r ApiNewAlgoOrderRequest) ReduceOnly(reduceOnly string) ApiNewAlgoOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiNewAlgoOrderRequest) ActivatePrice(activatePrice float32) ApiNewAlgoOrderRequest {
	r.activatePrice = &activatePrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1%
func (r ApiNewAlgoOrderRequest) CallbackRate(callbackRate float32) ApiNewAlgoOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

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

func (r ApiNewAlgoOrderRequest) Execute() (*common.RestApiResponse[models.NewAlgoOrderResponse], error) {
	return r.ApiService.NewAlgoOrderExecute(r)
}

/*
NewAlgoOrder New Algo Order(TRADE)
Post /fapi/v1/algoOrder

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/New-Algo-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoType -  Only support `CONDITIONAL`
@param symbol -
@param side -  `SELL`, `BUY`
@param type_ -
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode.
@param timeInForce -
@param quantity -
@param price -
@param triggerPrice -
@param workingType -  stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param closePosition -  `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.
@param priceProtect -  \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode
@param activatePrice -  Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1%
@param clientAlgoId -
@param selfTradePreventionMode -  `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`
@param goodTillDate -  order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/algoOrder"

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
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.goodTillDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "goodTillDate", r.goodTillDate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NewAlgoOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	type_                   *string
	positionSide            *models.NewAlgoOrderPositionSideParameter
	timeInForce             *models.NewAlgoOrderTimeInForceParameter
	quantity                *float32
	reduceOnly              *string
	price                   *float32
	newClientOrderId        *string
	newOrderRespType        *models.PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType
	priceMatch              *models.NewAlgoOrderPriceMatchParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

func (r ApiNewOrderRequest) Symbol(symbol string) ApiNewOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiNewOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiNewOrderRequest) Type(type_ string) ApiNewOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode.
func (r ApiNewOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiNewOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiNewOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiNewOrderRequest) Quantity(quantity float32) ApiNewOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode
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

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiNewOrderRequest) NewOrderRespType(newOrderRespType models.PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) ApiNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiNewOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiNewOrderRequest {
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

func (r ApiNewOrderRequest) Execute() (*common.RestApiResponse[models.NewOrderResponse], error) {
	return r.ApiService.NewOrderExecute(r)
}

/*
NewOrder New Order(TRADE)
Post /fapi/v1/order

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -  `SELL`, `BUY`
@param type_ -
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param selfTradePreventionMode -  `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`
@param goodTillDate -  order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order"

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
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
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
	if r.goodTillDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "goodTillDate", r.goodTillDate, "form", "")
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
PlaceMultipleOrders Place Multiple Orders(TRADE)
Post /fapi/v1/batchOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Place-Multiple-Orders

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/batchOrders"

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

	resp, err := SendRequest[models.PlaceMultipleOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
PositionAdlQuantileEstimation Position ADL Quantile Estimation(USER_DATA)
Get /fapi/v1/adlQuantile

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Position-ADL-Quantile-Estimation

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/adlQuantile"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PositionAdlQuantileEstimationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPositionInformationV2Request struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiPositionInformationV2Request) Symbol(symbol string) ApiPositionInformationV2Request {
	r.symbol = &symbol
	return r
}

func (r ApiPositionInformationV2Request) RecvWindow(recvWindow int64) ApiPositionInformationV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionInformationV2Request) Execute() (*common.RestApiResponse[models.PositionInformationV2Response], error) {
	return r.ApiService.PositionInformationV2Execute(r)
}

/*
PositionInformationV2 Position Information V2 (USER_DATA)
Get /fapi/v2/positionRisk

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Position-Information-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiPositionInformationV2Request
*/
func (a *TradeAPIService) PositionInformationV2(ctx context.Context) ApiPositionInformationV2Request {
	return ApiPositionInformationV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PositionInformationV2Response
func (a *TradeAPIService) PositionInformationV2Execute(r ApiPositionInformationV2Request) (*common.RestApiResponse[models.PositionInformationV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v2/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PositionInformationV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPositionInformationV3Request struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiPositionInformationV3Request) Symbol(symbol string) ApiPositionInformationV3Request {
	r.symbol = &symbol
	return r
}

func (r ApiPositionInformationV3Request) RecvWindow(recvWindow int64) ApiPositionInformationV3Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPositionInformationV3Request) Execute() (*common.RestApiResponse[models.PositionInformationV3Response], error) {
	return r.ApiService.PositionInformationV3Execute(r)
}

/*
PositionInformationV3 Position Information V3 (USER_DATA)
Get /fapi/v3/positionRisk

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Position-Information-V3

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiPositionInformationV3Request
*/
func (a *TradeAPIService) PositionInformationV3(ctx context.Context) ApiPositionInformationV3Request {
	return ApiPositionInformationV3Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PositionInformationV3Response
func (a *TradeAPIService) PositionInformationV3Execute(r ApiPositionInformationV3Request) (*common.RestApiResponse[models.PositionInformationV3Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v3/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PositionInformationV3Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiQueryAlgoOrderRequest) AlgoId(algoId int64) ApiQueryAlgoOrderRequest {
	r.algoId = &algoId
	return r
}

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
Get /fapi/v1/algoOrder

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Query-Algo-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -
@param clientAlgoId -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/algoOrder"

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

	resp, err := SendRequest[models.QueryAlgoOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	page       *int64
	limit      *int64
	recvWindow *int64
}

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

func (r ApiQueryAllAlgoOrdersRequest) Page(page int64) ApiQueryAllAlgoOrdersRequest {
	r.page = &page
	return r
}

// Default 100; max 1000
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
Get /fapi/v1/allAlgoOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Query-All-Algo-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param algoId -
@param startTime -
@param endTime -
@param page -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/allAlgoOrders"

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
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryAllAlgoOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiQueryCurrentOpenOrderRequest) Symbol(symbol string) ApiQueryCurrentOpenOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentOpenOrderRequest) OrderId(orderId int64) ApiQueryCurrentOpenOrderRequest {
	r.orderId = &orderId
	return r
}

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
Get /fapi/v1/openOrder

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Query-Current-Open-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/openOrder"

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

	resp, err := SendRequest[models.QueryCurrentOpenOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiQueryOrderRequest) Symbol(symbol string) ApiQueryOrderRequest {
	r.symbol = &symbol
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

func (r ApiQueryOrderRequest) Execute() (*common.RestApiResponse[models.QueryOrderResponse], error) {
	return r.ApiService.QueryOrderExecute(r)
}

/*
QueryOrder Query Order (USER_DATA)
Get /fapi/v1/order

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Query-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order"

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

	resp, err := SendRequest[models.QueryOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTestOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.NewAlgoOrderSideParameter
	type_                   *string
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
	newOrderRespType        *models.PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType
	priceMatch              *models.NewAlgoOrderPriceMatchParameter
	selfTradePreventionMode *models.NewAlgoOrderSelfTradePreventionModeParameter
	goodTillDate            *int64
	recvWindow              *int64
}

func (r ApiTestOrderRequest) Symbol(symbol string) ApiTestOrderRequest {
	r.symbol = &symbol
	return r
}

// &#x60;SELL&#x60;, &#x60;BUY&#x60;
func (r ApiTestOrderRequest) Side(side models.NewAlgoOrderSideParameter) ApiTestOrderRequest {
	r.side = &side
	return r
}

func (r ApiTestOrderRequest) Type(type_ string) ApiTestOrderRequest {
	r.type_ = &type_
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent with Hedge Mode.
func (r ApiTestOrderRequest) PositionSide(positionSide models.NewAlgoOrderPositionSideParameter) ApiTestOrderRequest {
	r.positionSide = &positionSide
	return r
}

func (r ApiTestOrderRequest) TimeInForce(timeInForce models.NewAlgoOrderTimeInForceParameter) ApiTestOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiTestOrderRequest) Quantity(quantity float32) ApiTestOrderRequest {
	r.quantity = &quantity
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. default \&quot;false\&quot;. Cannot be sent in Hedge Mode
func (r ApiTestOrderRequest) ReduceOnly(reduceOnly string) ApiTestOrderRequest {
	r.reduceOnly = &reduceOnly
	return r
}

func (r ApiTestOrderRequest) Price(price float32) ApiTestOrderRequest {
	r.price = &price
	return r
}

// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60;
func (r ApiTestOrderRequest) NewClientOrderId(newClientOrderId string) ApiTestOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiTestOrderRequest) StopPrice(stopPrice float32) ApiTestOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// &#x60;true&#x60;, &#x60;false&#x60;；Close-All，used with &#x60;STOP_MARKET&#x60; or &#x60;TAKE_PROFIT_MARKET&#x60;.
func (r ApiTestOrderRequest) ClosePosition(closePosition string) ApiTestOrderRequest {
	r.closePosition = &closePosition
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;)
func (r ApiTestOrderRequest) ActivationPrice(activationPrice float32) ApiTestOrderRequest {
	r.activationPrice = &activationPrice
	return r
}

// Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 5 where 1 for 1%
func (r ApiTestOrderRequest) CallbackRate(callbackRate float32) ApiTestOrderRequest {
	r.callbackRate = &callbackRate
	return r
}

// stopPrice triggered by: \&quot;MARK_PRICE\&quot;, \&quot;CONTRACT_PRICE\&quot;. Default \&quot;CONTRACT_PRICE\&quot;
func (r ApiTestOrderRequest) WorkingType(workingType models.NewAlgoOrderWorkingTypeParameter) ApiTestOrderRequest {
	r.workingType = &workingType
	return r
}

// \&quot;TRUE\&quot; or \&quot;FALSE\&quot;, default \&quot;FALSE\&quot;. Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders.
func (r ApiTestOrderRequest) PriceProtect(priceProtect string) ApiTestOrderRequest {
	r.priceProtect = &priceProtect
	return r
}

// \&quot;ACK\&quot;, \&quot;RESULT\&quot;, default \&quot;ACK\&quot;
func (r ApiTestOrderRequest) NewOrderRespType(newOrderRespType models.PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) ApiTestOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// only avaliable for &#x60;LIMIT&#x60;/&#x60;STOP&#x60;/&#x60;TAKE_PROFIT&#x60; order; can be set to &#x60;OPPONENT&#x60;/ &#x60;OPPONENT_5&#x60;/ &#x60;OPPONENT_10&#x60;/ &#x60;OPPONENT_20&#x60;: /&#x60;QUEUE&#x60;/ &#x60;QUEUE_5&#x60;/ &#x60;QUEUE_10&#x60;/ &#x60;QUEUE_20&#x60;; Can&#39;t be passed together with &#x60;price&#x60;
func (r ApiTestOrderRequest) PriceMatch(priceMatch models.NewAlgoOrderPriceMatchParameter) ApiTestOrderRequest {
	r.priceMatch = &priceMatch
	return r
}

// &#x60;EXPIRE_TAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_MAKER&#x60;:expire taker order when STP triggers/ &#x60;EXPIRE_BOTH&#x60;:expire both orders when STP triggers; default &#x60;NONE&#x60;
func (r ApiTestOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.NewAlgoOrderSelfTradePreventionModeParameter) ApiTestOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// order cancel time for timeInForce &#x60;GTD&#x60;, mandatory when &#x60;timeInforce&#x60; set to &#x60;GTD&#x60;; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
func (r ApiTestOrderRequest) GoodTillDate(goodTillDate int64) ApiTestOrderRequest {
	r.goodTillDate = &goodTillDate
	return r
}

func (r ApiTestOrderRequest) RecvWindow(recvWindow int64) ApiTestOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTestOrderRequest) Execute() (*common.RestApiResponse[models.TestOrderResponse], error) {
	return r.ApiService.TestOrderExecute(r)
}

/*
TestOrder Test Order(TRADE)
Post /fapi/v1/order/test

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/New-Order-Test

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -  `SELL`, `BUY`
@param type_ -
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent with Hedge Mode.
@param timeInForce -
@param quantity -
@param reduceOnly -  \"true\" or \"false\". default \"false\". Cannot be sent in Hedge Mode
@param price -
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`
@param stopPrice -  Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param closePosition -  `true`, `false`；Close-All，used with `STOP_MARKET` or `TAKE_PROFIT_MARKET`.
@param activationPrice -  Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)
@param callbackRate -  Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 5 where 1 for 1%
@param workingType -  stopPrice triggered by: \"MARK_PRICE\", \"CONTRACT_PRICE\". Default \"CONTRACT_PRICE\"
@param priceProtect -  \"TRUE\" or \"FALSE\", default \"FALSE\". Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
@param newOrderRespType -  \"ACK\", \"RESULT\", default \"ACK\"
@param priceMatch -  only avaliable for `LIMIT`/`STOP`/`TAKE_PROFIT` order; can be set to `OPPONENT`/ `OPPONENT_5`/ `OPPONENT_10`/ `OPPONENT_20`: /`QUEUE`/ `QUEUE_5`/ `QUEUE_10`/ `QUEUE_20`; Can't be passed together with `price`
@param selfTradePreventionMode -  `EXPIRE_TAKER`:expire taker order when STP triggers/ `EXPIRE_MAKER`:expire taker order when STP triggers/ `EXPIRE_BOTH`:expire both orders when STP triggers; default `NONE`
@param goodTillDate -  order cancel time for timeInForce `GTD`, mandatory when `timeInforce` set to `GTD`; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
@param recvWindow -
@return ApiTestOrderRequest
*/
func (a *TradeAPIService) TestOrder(ctx context.Context) ApiTestOrderRequest {
	return ApiTestOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TestOrderResponse
func (a *TradeAPIService) TestOrderExecute(r ApiTestOrderRequest) (*common.RestApiResponse[models.TestOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order/test"

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
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
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
	if r.goodTillDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "goodTillDate", r.goodTillDate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TestOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// \&quot;LIQUIDATION\&quot; for liquidation orders, \&quot;ADL\&quot; for ADL orders.
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

// Default 100; max 1000
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
Get /fapi/v1/forceOrders

https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Users-Force-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param autoCloseType -  \"LIQUIDATION\" for liquidation orders, \"ADL\" for ADL orders.
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/forceOrders"

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

	resp, err := SendRequest[models.UsersForceOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
