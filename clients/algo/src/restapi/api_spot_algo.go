/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package binancealgorestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/algo/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// SpotAlgoAPIService SpotAlgoAPI Service
type SpotAlgoAPIService Service

type ApiCancelAlgoOrderSpotAlgoRequest struct {
	ctx        context.Context
	ApiService *SpotAlgoAPIService
	algoId     *int64
	recvWindow *int64
}

// eg. 14511
func (r ApiCancelAlgoOrderSpotAlgoRequest) AlgoId(algoId int64) ApiCancelAlgoOrderSpotAlgoRequest {
	r.algoId = &algoId
	return r
}

func (r ApiCancelAlgoOrderSpotAlgoRequest) RecvWindow(recvWindow int64) ApiCancelAlgoOrderSpotAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAlgoOrderSpotAlgoRequest) Execute() (*common.RestApiResponse[models.CancelAlgoOrderSpotAlgoResponse], error) {
	return r.ApiService.CancelAlgoOrderSpotAlgoExecute(r)
}

/*
CancelAlgoOrderSpotAlgo Cancel Algo Order(TRADE)
Delete /sapi/v1/algo/spot/order

https://developers.binance.com/docs/algo/spot-algo/Cancel-Algo-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param recvWindow -
@return ApiCancelAlgoOrderSpotAlgoRequest
*/
func (a *SpotAlgoAPIService) CancelAlgoOrderSpotAlgo(ctx context.Context) ApiCancelAlgoOrderSpotAlgoRequest {
	return ApiCancelAlgoOrderSpotAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAlgoOrderSpotAlgoResponse
func (a *SpotAlgoAPIService) CancelAlgoOrderSpotAlgoExecute(r ApiCancelAlgoOrderSpotAlgoRequest) (*common.RestApiResponse[models.CancelAlgoOrderSpotAlgoResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/spot/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoId == nil {
		return nil, common.ReportError("algoId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAlgoOrderSpotAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest struct {
	ctx        context.Context
	ApiService *SpotAlgoAPIService
	recvWindow *int64
}

func (r ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest) RecvWindow(recvWindow int64) ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest) Execute() (*common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse], error) {
	return r.ApiService.QueryCurrentAlgoOpenOrdersSpotAlgoExecute(r)
}

/*
QueryCurrentAlgoOpenOrdersSpotAlgo Query Current Algo Open Orders(USER_DATA)
Get /sapi/v1/algo/spot/openOrders

https://developers.binance.com/docs/algo/spot-algo/Query-Current-Algo-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest
*/
func (a *SpotAlgoAPIService) QueryCurrentAlgoOpenOrdersSpotAlgo(ctx context.Context) ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest {
	return ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentAlgoOpenOrdersSpotAlgoResponse
func (a *SpotAlgoAPIService) QueryCurrentAlgoOpenOrdersSpotAlgoExecute(r ApiQueryCurrentAlgoOpenOrdersSpotAlgoRequest) (*common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/spot/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryHistoricalAlgoOrdersSpotAlgoRequest struct {
	ctx        context.Context
	ApiService *SpotAlgoAPIService
	symbol     *string
	side       *string
	startTime  *int64
	endTime    *int64
	page       *int64
	pageSize   *int64
	recvWindow *int64
}

// Trading symbol eg. BTCUSDT
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) Symbol(symbol string) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.symbol = &symbol
	return r
}

// BUY or SELL
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) Side(side string) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.side = &side
	return r
}

// in milliseconds  eg.1641522717552
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) StartTime(startTime int64) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.startTime = &startTime
	return r
}

// in milliseconds  eg.1641522526562
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) EndTime(endTime int64) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.endTime = &endTime
	return r
}

// Default is 1
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) Page(page int64) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.page = &page
	return r
}

// MIN 1, MAX 100; Default 100
func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) PageSize(pageSize int64) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) RecvWindow(recvWindow int64) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) Execute() (*common.RestApiResponse[models.QueryHistoricalAlgoOrdersSpotAlgoResponse], error) {
	return r.ApiService.QueryHistoricalAlgoOrdersSpotAlgoExecute(r)
}

/*
QueryHistoricalAlgoOrdersSpotAlgo Query Historical Algo Orders(USER_DATA)
Get /sapi/v1/algo/spot/historicalOrders

https://developers.binance.com/docs/algo/spot-algo/Query-Historical-Algo-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  BUY or SELL
@param startTime -  in milliseconds  eg.1641522717552
@param endTime -  in milliseconds  eg.1641522526562
@param page -  Default is 1
@param pageSize -  MIN 1, MAX 100; Default 100
@param recvWindow -
@return ApiQueryHistoricalAlgoOrdersSpotAlgoRequest
*/
func (a *SpotAlgoAPIService) QueryHistoricalAlgoOrdersSpotAlgo(ctx context.Context) ApiQueryHistoricalAlgoOrdersSpotAlgoRequest {
	return ApiQueryHistoricalAlgoOrdersSpotAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryHistoricalAlgoOrdersSpotAlgoResponse
func (a *SpotAlgoAPIService) QueryHistoricalAlgoOrdersSpotAlgoExecute(r ApiQueryHistoricalAlgoOrdersSpotAlgoRequest) (*common.RestApiResponse[models.QueryHistoricalAlgoOrdersSpotAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/spot/historicalOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.side != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
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
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryHistoricalAlgoOrdersSpotAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubOrdersSpotAlgoRequest struct {
	ctx        context.Context
	ApiService *SpotAlgoAPIService
	algoId     *int64
	page       *int64
	pageSize   *int64
	recvWindow *int64
}

// eg. 14511
func (r ApiQuerySubOrdersSpotAlgoRequest) AlgoId(algoId int64) ApiQuerySubOrdersSpotAlgoRequest {
	r.algoId = &algoId
	return r
}

// Default is 1
func (r ApiQuerySubOrdersSpotAlgoRequest) Page(page int64) ApiQuerySubOrdersSpotAlgoRequest {
	r.page = &page
	return r
}

// MIN 1, MAX 100; Default 100
func (r ApiQuerySubOrdersSpotAlgoRequest) PageSize(pageSize int64) ApiQuerySubOrdersSpotAlgoRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiQuerySubOrdersSpotAlgoRequest) RecvWindow(recvWindow int64) ApiQuerySubOrdersSpotAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubOrdersSpotAlgoRequest) Execute() (*common.RestApiResponse[models.QuerySubOrdersSpotAlgoResponse], error) {
	return r.ApiService.QuerySubOrdersSpotAlgoExecute(r)
}

/*
QuerySubOrdersSpotAlgo Query Sub Orders(USER_DATA)
Get /sapi/v1/algo/spot/subOrders

https://developers.binance.com/docs/algo/spot-algo/Query-Sub-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param page -  Default is 1
@param pageSize -  MIN 1, MAX 100; Default 100
@param recvWindow -
@return ApiQuerySubOrdersSpotAlgoRequest
*/
func (a *SpotAlgoAPIService) QuerySubOrdersSpotAlgo(ctx context.Context) ApiQuerySubOrdersSpotAlgoRequest {
	return ApiQuerySubOrdersSpotAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubOrdersSpotAlgoResponse
func (a *SpotAlgoAPIService) QuerySubOrdersSpotAlgoExecute(r ApiQuerySubOrdersSpotAlgoRequest) (*common.RestApiResponse[models.QuerySubOrdersSpotAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/spot/subOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algoId == nil {
		return nil, common.ReportError("algoId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubOrdersSpotAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTimeWeightedAveragePriceSpotAlgoRequest struct {
	ctx          context.Context
	ApiService   *SpotAlgoAPIService
	symbol       *string
	side         *string
	quantity     *float32
	duration     *int64
	clientAlgoId *string
	limitPrice   *float32
}

// Trading symbol eg. BTCUSDT
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) Symbol(symbol string) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.symbol = &symbol
	return r
}

// Trading side ( BUY or SELL )
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) Side(side string) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.side = &side
	return r
}

// Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) Quantity(quantity float32) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.quantity = &quantity
	return r
}

// Duration for TWAP orders in seconds. [300, 86400]
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) Duration(duration int64) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.duration = &duration
	return r
}

// A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) ClientAlgoId(clientAlgoId string) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// Limit price of the order; If it is not sent, will place order by market price by default
func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) LimitPrice(limitPrice float32) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	r.limitPrice = &limitPrice
	return r
}

func (r ApiTimeWeightedAveragePriceSpotAlgoRequest) Execute() (*common.RestApiResponse[models.TimeWeightedAveragePriceSpotAlgoResponse], error) {
	return r.ApiService.TimeWeightedAveragePriceSpotAlgoExecute(r)
}

/*
TimeWeightedAveragePriceSpotAlgo Time-Weighted Average Price(Twap) New Order(TRADE)
Post /sapi/v1/algo/spot/newOrderTwap

https://developers.binance.com/docs/algo/spot-algo/Time-Weighted-Average-Price-New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  Trading side ( BUY or SELL )
@param quantity -  Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
@param duration -  Duration for TWAP orders in seconds. [300, 86400]
@param clientAlgoId -  A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
@param limitPrice -  Limit price of the order; If it is not sent, will place order by market price by default
@return ApiTimeWeightedAveragePriceSpotAlgoRequest
*/
func (a *SpotAlgoAPIService) TimeWeightedAveragePriceSpotAlgo(ctx context.Context) ApiTimeWeightedAveragePriceSpotAlgoRequest {
	return ApiTimeWeightedAveragePriceSpotAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TimeWeightedAveragePriceSpotAlgoResponse
func (a *SpotAlgoAPIService) TimeWeightedAveragePriceSpotAlgoExecute(r ApiTimeWeightedAveragePriceSpotAlgoRequest) (*common.RestApiResponse[models.TimeWeightedAveragePriceSpotAlgoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/spot/newOrderTwap"

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
	if r.duration == nil {
		return nil, common.ReportError("duration is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "form", "")
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
	}
	if r.limitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitPrice", r.limitPrice, "form", "")
	}

	resp, err := SendRequest[models.TimeWeightedAveragePriceSpotAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
