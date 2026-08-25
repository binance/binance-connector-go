/*
Algo Trading REST API

Programmatic access to Binance’s execution algorithms for creating and managing Spot and Futures algo orders.
*/

package binancealgorestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/algo/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// FutureAlgoAPIService FutureAlgoAPI Service
type FutureAlgoAPIService Service

type ApiCancelAlgoOrderFutureAlgoRequest struct {
	ctx          context.Context
	ApiService   *FutureAlgoAPIService
	algoId       *int64
	clientAlgoId *string
	recvWindow   *int64
}

// eg. 14511
func (r ApiCancelAlgoOrderFutureAlgoRequest) AlgoId(algoId int64) ApiCancelAlgoOrderFutureAlgoRequest {
	r.algoId = &algoId
	return r
}

// eg. \&quot;65ce1630101a480b85915d7e11fd5078\&quot;
func (r ApiCancelAlgoOrderFutureAlgoRequest) ClientAlgoId(clientAlgoId string) ApiCancelAlgoOrderFutureAlgoRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// Request validity window in milliseconds
func (r ApiCancelAlgoOrderFutureAlgoRequest) RecvWindow(recvWindow int64) ApiCancelAlgoOrderFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAlgoOrderFutureAlgoRequest) Execute() (*common.RestApiResponse[models.CancelAlgoOrderFutureAlgoResponse], error) {
	return r.ApiService.CancelAlgoOrderFutureAlgoExecute(r)
}

/*
CancelAlgoOrderFutureAlgo Cancel Futures Algo Order (TRADE)
Delete /sapi/v1/algo/futures/order

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#cancel-algo-order-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param clientAlgoId -  eg. \"65ce1630101a480b85915d7e11fd5078\"
@param recvWindow -  Request validity window in milliseconds
@return ApiCancelAlgoOrderFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) CancelAlgoOrderFutureAlgo(ctx context.Context) ApiCancelAlgoOrderFutureAlgoRequest {
	return ApiCancelAlgoOrderFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelAlgoOrderFutureAlgoResponse
func (a *FutureAlgoAPIService) CancelAlgoOrderFutureAlgoExecute(r ApiCancelAlgoOrderFutureAlgoRequest) (*common.RestApiResponse[models.CancelAlgoOrderFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/order"

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

	resp, err := SendRequest[models.CancelAlgoOrderFutureAlgoResponse](
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

type ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest struct {
	ctx        context.Context
	ApiService *FutureAlgoAPIService
	recvWindow *int64
}

// Request validity window in milliseconds
func (r ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse], error) {
	return r.ApiService.QueryCurrentAlgoOpenOrdersFutureAlgoExecute(r)
}

/*
QueryCurrentAlgoOpenOrdersFutureAlgo Query Current Futures Algo Open Orders (USER_DATA)
Get /sapi/v1/algo/futures/openOrders

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#query-current-algo-open-orders-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) QueryCurrentAlgoOpenOrdersFutureAlgo(ctx context.Context) ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest {
	return ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentAlgoOpenOrdersFutureAlgoResponse
func (a *FutureAlgoAPIService) QueryCurrentAlgoOpenOrdersFutureAlgoExecute(r ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest) (*common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse](
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

type ApiQueryHistoricalAlgoOrdersFutureAlgoRequest struct {
	ctx        context.Context
	ApiService *FutureAlgoAPIService
	symbol     *string
	side       *models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter
	startTime  *int64
	endTime    *int64
	page       *int64
	pageSize   *int64
	recvWindow *int64
}

// Trading symbol eg. BTCUSDT
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Symbol(symbol string) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.symbol = &symbol
	return r
}

// BUY or SELL
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Side(side models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.side = &side
	return r
}

// in milliseconds  eg.1641522717552
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) StartTime(startTime int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.startTime = &startTime
	return r
}

// in milliseconds  eg.1641522526562
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) EndTime(endTime int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.endTime = &endTime
	return r
}

// Page number
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Page(page int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.page = &page
	return r
}

// Records per page
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) PageSize(pageSize int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.pageSize = &pageSize
	return r
}

// Request validity window in milliseconds
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QueryHistoricalAlgoOrdersFutureAlgoResponse], error) {
	return r.ApiService.QueryHistoricalAlgoOrdersFutureAlgoExecute(r)
}

/*
QueryHistoricalAlgoOrdersFutureAlgo Query Historical Futures Algo Orders (USER_DATA)
Get /sapi/v1/algo/futures/historicalOrders

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#query-historical-algo-orders-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  BUY or SELL
@param startTime -  in milliseconds  eg.1641522717552
@param endTime -  in milliseconds  eg.1641522526562
@param page -  Page number
@param pageSize -  Records per page
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryHistoricalAlgoOrdersFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) QueryHistoricalAlgoOrdersFutureAlgo(ctx context.Context) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	return ApiQueryHistoricalAlgoOrdersFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryHistoricalAlgoOrdersFutureAlgoResponse
func (a *FutureAlgoAPIService) QueryHistoricalAlgoOrdersFutureAlgoExecute(r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) (*common.RestApiResponse[models.QueryHistoricalAlgoOrdersFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/historicalOrders"

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

	resp, err := SendRequest[models.QueryHistoricalAlgoOrdersFutureAlgoResponse](
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

type ApiQuerySubOrdersFutureAlgoRequest struct {
	ctx        context.Context
	ApiService *FutureAlgoAPIService
	algoId     *int64
	page       *int64
	pageSize   *int64
	recvWindow *int64
}

// eg. 14511
func (r ApiQuerySubOrdersFutureAlgoRequest) AlgoId(algoId int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.algoId = &algoId
	return r
}

// Page number
func (r ApiQuerySubOrdersFutureAlgoRequest) Page(page int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.page = &page
	return r
}

// Records per page
func (r ApiQuerySubOrdersFutureAlgoRequest) PageSize(pageSize int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.pageSize = &pageSize
	return r
}

// Request validity window in milliseconds
func (r ApiQuerySubOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QuerySubOrdersFutureAlgoResponse], error) {
	return r.ApiService.QuerySubOrdersFutureAlgoExecute(r)
}

/*
QuerySubOrdersFutureAlgo Query Futures Sub Orders (USER_DATA)
Get /sapi/v1/algo/futures/subOrders

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#query-sub-orders-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param page -  Page number
@param pageSize -  Records per page
@param recvWindow -  Request validity window in milliseconds
@return ApiQuerySubOrdersFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) QuerySubOrdersFutureAlgo(ctx context.Context) ApiQuerySubOrdersFutureAlgoRequest {
	return ApiQuerySubOrdersFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubOrdersFutureAlgoResponse
func (a *FutureAlgoAPIService) QuerySubOrdersFutureAlgoExecute(r ApiQuerySubOrdersFutureAlgoRequest) (*common.RestApiResponse[models.QuerySubOrdersFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/subOrders"

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

	resp, err := SendRequest[models.QuerySubOrdersFutureAlgoResponse](
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

type ApiTimeWeightedAveragePriceFutureAlgoRequest struct {
	ctx          context.Context
	ApiService   *FutureAlgoAPIService
	symbol       *string
	side         *models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter
	quantity     *float32
	duration     *int64
	positionSide *models.TimeWeightedAveragePriceFutureAlgoPositionSideParameter
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *float32
	recvWindow   *int64
}

// Trading symbol eg. BTCUSDT
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Symbol(symbol string) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.symbol = &symbol
	return r
}

// Trading side ( BUY or SELL )
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Side(side models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.side = &side
	return r
}

// Quantity of base asset; The notional (&#x60;quantity&#x60; * &#x60;mark price(base asset)&#x60;) must be more than the equivalent of 1,000 USDT and less than the equivalent of 1,000,000 USDT
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Quantity(quantity float32) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.quantity = &quantity
	return r
}

// Duration for TWAP orders in seconds
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Duration(duration int64) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.duration = &duration
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) PositionSide(positionSide models.TimeWeightedAveragePriceFutureAlgoPositionSideParameter) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.positionSide = &positionSide
	return r
}

// A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) ClientAlgoId(clientAlgoId string) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. Default \&quot;false\&quot;; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) ReduceOnly(reduceOnly bool) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Limit price of the order; If it is not sent, will place order by market price by default
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) LimitPrice(limitPrice float32) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.limitPrice = &limitPrice
	return r
}

// Request validity window in milliseconds
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) RecvWindow(recvWindow int64) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Execute() (*common.RestApiResponse[models.TimeWeightedAveragePriceFutureAlgoResponse], error) {
	return r.ApiService.TimeWeightedAveragePriceFutureAlgoExecute(r)
}

/*
TimeWeightedAveragePriceFutureAlgo Time-Weighted Futures Average Price (Twap) New Order (TRADE)
Post /sapi/v1/algo/futures/newOrderTwap

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#time-weighted-average-price-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  Trading side ( BUY or SELL )
@param quantity -  Quantity of base asset; The notional (`quantity` * `mark price(base asset)`) must be more than the equivalent of 1,000 USDT and less than the equivalent of 1,000,000 USDT
@param duration -  Duration for TWAP orders in seconds
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param clientAlgoId -  A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
@param reduceOnly -  \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
@param limitPrice -  Limit price of the order; If it is not sent, will place order by market price by default
@param recvWindow -  Request validity window in milliseconds
@return ApiTimeWeightedAveragePriceFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) TimeWeightedAveragePriceFutureAlgo(ctx context.Context) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	return ApiTimeWeightedAveragePriceFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TimeWeightedAveragePriceFutureAlgoResponse
func (a *FutureAlgoAPIService) TimeWeightedAveragePriceFutureAlgoExecute(r ApiTimeWeightedAveragePriceFutureAlgoRequest) (*common.RestApiResponse[models.TimeWeightedAveragePriceFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/newOrderTwap"

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
	if *r.duration < 300 {
		return nil, common.ReportError("duration must be greater than 300")
	}
	if *r.duration > 86400 {
		return nil, common.ReportError("duration must be less than 86400")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "form", "")
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.limitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitPrice", r.limitPrice, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TimeWeightedAveragePriceFutureAlgoResponse](
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

type ApiVolumeParticipationFutureAlgoRequest struct {
	ctx          context.Context
	ApiService   *FutureAlgoAPIService
	symbol       *string
	side         *models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter
	quantity     *float32
	urgency      *models.VolumeParticipationFutureAlgoUrgencyParameter
	positionSide *models.TimeWeightedAveragePriceFutureAlgoPositionSideParameter
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *float32
	recvWindow   *int64
}

// Trading symbol eg. BTCUSDT
func (r ApiVolumeParticipationFutureAlgoRequest) Symbol(symbol string) ApiVolumeParticipationFutureAlgoRequest {
	r.symbol = &symbol
	return r
}

// Trading side ( BUY or SELL )
func (r ApiVolumeParticipationFutureAlgoRequest) Side(side models.QueryHistoricalAlgoOrdersFutureAlgoSideParameter) ApiVolumeParticipationFutureAlgoRequest {
	r.side = &side
	return r
}

// Quantity of base asset; The notional (&#x60;quantity&#x60; * &#x60;mark price(base asset)&#x60;) must be more than the equivalent of 10,000 USDT and less than the equivalent of 1,000,000 USDT
func (r ApiVolumeParticipationFutureAlgoRequest) Quantity(quantity float32) ApiVolumeParticipationFutureAlgoRequest {
	r.quantity = &quantity
	return r
}

// Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH
func (r ApiVolumeParticipationFutureAlgoRequest) Urgency(urgency models.VolumeParticipationFutureAlgoUrgencyParameter) ApiVolumeParticipationFutureAlgoRequest {
	r.urgency = &urgency
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiVolumeParticipationFutureAlgoRequest) PositionSide(positionSide models.TimeWeightedAveragePriceFutureAlgoPositionSideParameter) ApiVolumeParticipationFutureAlgoRequest {
	r.positionSide = &positionSide
	return r
}

// A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
func (r ApiVolumeParticipationFutureAlgoRequest) ClientAlgoId(clientAlgoId string) ApiVolumeParticipationFutureAlgoRequest {
	r.clientAlgoId = &clientAlgoId
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;. Default \&quot;false\&quot;; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
func (r ApiVolumeParticipationFutureAlgoRequest) ReduceOnly(reduceOnly bool) ApiVolumeParticipationFutureAlgoRequest {
	r.reduceOnly = &reduceOnly
	return r
}

// Limit price of the order; If it is not sent, will place order by market price by default
func (r ApiVolumeParticipationFutureAlgoRequest) LimitPrice(limitPrice float32) ApiVolumeParticipationFutureAlgoRequest {
	r.limitPrice = &limitPrice
	return r
}

// Request validity window in milliseconds
func (r ApiVolumeParticipationFutureAlgoRequest) RecvWindow(recvWindow int64) ApiVolumeParticipationFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVolumeParticipationFutureAlgoRequest) Execute() (*common.RestApiResponse[models.VolumeParticipationFutureAlgoResponse], error) {
	return r.ApiService.VolumeParticipationFutureAlgoExecute(r)
}

/*
VolumeParticipationFutureAlgo Volume Participation (VP) New Order (TRADE)
Post /sapi/v1/algo/futures/newOrderVp

https://developers.binance.com/en/docs/catalog/advanced-trading-algo-trading/api/rest-api/future-algo#volume-participation-future-algo

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  Trading side ( BUY or SELL )
@param quantity -  Quantity of base asset; The notional (`quantity` * `mark price(base asset)`) must be more than the equivalent of 10,000 USDT and less than the equivalent of 1,000,000 USDT
@param urgency -  Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param clientAlgoId -  A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
@param reduceOnly -  \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
@param limitPrice -  Limit price of the order; If it is not sent, will place order by market price by default
@param recvWindow -  Request validity window in milliseconds
@return ApiVolumeParticipationFutureAlgoRequest
*/
func (a *FutureAlgoAPIService) VolumeParticipationFutureAlgo(ctx context.Context) ApiVolumeParticipationFutureAlgoRequest {
	return ApiVolumeParticipationFutureAlgoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VolumeParticipationFutureAlgoResponse
func (a *FutureAlgoAPIService) VolumeParticipationFutureAlgoExecute(r ApiVolumeParticipationFutureAlgoRequest) (*common.RestApiResponse[models.VolumeParticipationFutureAlgoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/algo/futures/newOrderVp"

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

	if r.urgency == nil {
		return nil, common.ReportError("urgency is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	if r.positionSide != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionSide", r.positionSide, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "urgency", r.urgency, "form", "")
	if r.clientAlgoId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientAlgoId", r.clientAlgoId, "form", "")
	}
	if r.reduceOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "reduceOnly", r.reduceOnly, "form", "")
	}
	if r.limitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitPrice", r.limitPrice, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VolumeParticipationFutureAlgoResponse](
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
