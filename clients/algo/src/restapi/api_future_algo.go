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

// FutureAlgoAPIService FutureAlgoAPI Service
type FutureAlgoAPIService Service

type ApiCancelAlgoOrderFutureAlgoRequest struct {
	ctx        context.Context
	ApiService *FutureAlgoAPIService
	algoId     *int64
	recvWindow *int64
}

// eg. 14511
func (r ApiCancelAlgoOrderFutureAlgoRequest) AlgoId(algoId int64) ApiCancelAlgoOrderFutureAlgoRequest {
	r.algoId = &algoId
	return r
}

func (r ApiCancelAlgoOrderFutureAlgoRequest) RecvWindow(recvWindow int64) ApiCancelAlgoOrderFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelAlgoOrderFutureAlgoRequest) Execute() (*common.RestApiResponse[models.CancelAlgoOrderFutureAlgoResponse], error) {
	return r.ApiService.CancelAlgoOrderFutureAlgoExecute(r)
}

/*
CancelAlgoOrderFutureAlgo Cancel Algo Order(TRADE)
Delete /sapi/v1/algo/futures/order

https://developers.binance.com/docs/algo/future-algo/Cancel-Algo-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param recvWindow -
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

	if r.algoId == nil {
		return nil, common.ReportError("algoId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algoId", r.algoId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelAlgoOrderFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentAlgoOpenOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse], error) {
	return r.ApiService.QueryCurrentAlgoOpenOrdersFutureAlgoExecute(r)
}

/*
QueryCurrentAlgoOpenOrdersFutureAlgo Query Current Algo Open Orders(USER_DATA)
Get /sapi/v1/algo/futures/openOrders

https://developers.binance.com/docs/algo/future-algo/Query-Current-Algo-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
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

	resp, err := SendRequest[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryHistoricalAlgoOrdersFutureAlgoRequest struct {
	ctx        context.Context
	ApiService *FutureAlgoAPIService
	symbol     *string
	side       *string
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
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Side(side string) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
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

// Default is 1
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Page(page int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.page = &page
	return r
}

// MIN 1, MAX 100; Default 100
func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) PageSize(pageSize int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQueryHistoricalAlgoOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryHistoricalAlgoOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QueryHistoricalAlgoOrdersFutureAlgoResponse], error) {
	return r.ApiService.QueryHistoricalAlgoOrdersFutureAlgoExecute(r)
}

/*
QueryHistoricalAlgoOrdersFutureAlgo Query Historical Algo Orders(USER_DATA)
Get /sapi/v1/algo/futures/historicalOrders

https://developers.binance.com/docs/algo/future-algo/Query-Historical-Algo-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  BUY or SELL
@param startTime -  in milliseconds  eg.1641522717552
@param endTime -  in milliseconds  eg.1641522526562
@param page -  Default is 1
@param pageSize -  MIN 1, MAX 100; Default 100
@param recvWindow -
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

	resp, err := SendRequest[models.QueryHistoricalAlgoOrdersFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// Default is 1
func (r ApiQuerySubOrdersFutureAlgoRequest) Page(page int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.page = &page
	return r
}

// MIN 1, MAX 100; Default 100
func (r ApiQuerySubOrdersFutureAlgoRequest) PageSize(pageSize int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiQuerySubOrdersFutureAlgoRequest) RecvWindow(recvWindow int64) ApiQuerySubOrdersFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubOrdersFutureAlgoRequest) Execute() (*common.RestApiResponse[models.QuerySubOrdersFutureAlgoResponse], error) {
	return r.ApiService.QuerySubOrdersFutureAlgoExecute(r)
}

/*
QuerySubOrdersFutureAlgo Query Sub Orders(USER_DATA)
Get /sapi/v1/algo/futures/subOrders

https://developers.binance.com/docs/algo/future-algo/Query-Sub-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algoId -  eg. 14511
@param page -  Default is 1
@param pageSize -  MIN 1, MAX 100; Default 100
@param recvWindow -
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

	resp, err := SendRequest[models.QuerySubOrdersFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTimeWeightedAveragePriceFutureAlgoRequest struct {
	ctx          context.Context
	ApiService   *FutureAlgoAPIService
	symbol       *string
	side         *string
	quantity     *float32
	duration     *int64
	positionSide *string
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
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Side(side string) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.side = &side
	return r
}

// Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Quantity(quantity float32) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.quantity = &quantity
	return r
}

// Duration for TWAP orders in seconds. [300, 86400]
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Duration(duration int64) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.duration = &duration
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) PositionSide(positionSide string) ApiTimeWeightedAveragePriceFutureAlgoRequest {
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

func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) RecvWindow(recvWindow int64) ApiTimeWeightedAveragePriceFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTimeWeightedAveragePriceFutureAlgoRequest) Execute() (*common.RestApiResponse[models.TimeWeightedAveragePriceFutureAlgoResponse], error) {
	return r.ApiService.TimeWeightedAveragePriceFutureAlgoExecute(r)
}

/*
TimeWeightedAveragePriceFutureAlgo Time-Weighted Average Price(Twap) New Order(TRADE)
Post /sapi/v1/algo/futures/newOrderTwap

https://developers.binance.com/docs/algo/future-algo/Time-Weighted-Average-Price-New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  Trading side ( BUY or SELL )
@param quantity -  Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
@param duration -  Duration for TWAP orders in seconds. [300, 86400]
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param clientAlgoId -  A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
@param reduceOnly -  \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
@param limitPrice -  Limit price of the order; If it is not sent, will place order by market price by default
@param recvWindow -
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

	resp, err := SendRequest[models.TimeWeightedAveragePriceFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiVolumeParticipationFutureAlgoRequest struct {
	ctx          context.Context
	ApiService   *FutureAlgoAPIService
	symbol       *string
	side         *string
	quantity     *float32
	urgency      *string
	positionSide *string
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
func (r ApiVolumeParticipationFutureAlgoRequest) Side(side string) ApiVolumeParticipationFutureAlgoRequest {
	r.side = &side
	return r
}

// Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
func (r ApiVolumeParticipationFutureAlgoRequest) Quantity(quantity float32) ApiVolumeParticipationFutureAlgoRequest {
	r.quantity = &quantity
	return r
}

// Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH
func (r ApiVolumeParticipationFutureAlgoRequest) Urgency(urgency string) ApiVolumeParticipationFutureAlgoRequest {
	r.urgency = &urgency
	return r
}

// Default &#x60;BOTH&#x60; for One-way Mode ; &#x60;LONG&#x60; or &#x60;SHORT&#x60; for Hedge Mode. It must be sent in Hedge Mode.
func (r ApiVolumeParticipationFutureAlgoRequest) PositionSide(positionSide string) ApiVolumeParticipationFutureAlgoRequest {
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

func (r ApiVolumeParticipationFutureAlgoRequest) RecvWindow(recvWindow int64) ApiVolumeParticipationFutureAlgoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVolumeParticipationFutureAlgoRequest) Execute() (*common.RestApiResponse[models.VolumeParticipationFutureAlgoResponse], error) {
	return r.ApiService.VolumeParticipationFutureAlgoExecute(r)
}

/*
VolumeParticipationFutureAlgo Volume Participation(VP) New Order (TRADE)
Post /sapi/v1/algo/futures/newOrderVp

https://developers.binance.com/docs/algo/future-algo/Volume-Participation-New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Trading symbol eg. BTCUSDT
@param side -  Trading side ( BUY or SELL )
@param quantity -  Quantity of base asset; Maximum notional per order is 200k, 2mm or 10mm, depending on symbol. Please reduce your size if you order is above the maximum notional per order.
@param urgency -  Represent the relative speed of the current execution; ENUM: LOW, MEDIUM, HIGH
@param positionSide -  Default `BOTH` for One-way Mode ; `LONG` or `SHORT` for Hedge Mode. It must be sent in Hedge Mode.
@param clientAlgoId -  A unique id among Algo orders (length should be 32 characters)， If it is not sent, we will give default value
@param reduceOnly -  \"true\" or \"false\". Default \"false\"; Cannot be sent in Hedge Mode; Cannot be sent when you open a position
@param limitPrice -  Limit price of the order; If it is not sent, will place order by market price by default
@param recvWindow -
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

	resp, err := SendRequest[models.VolumeParticipationFutureAlgoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
