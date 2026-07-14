/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// BorrowRepayAPIService BorrowRepayAPI Service
type BorrowRepayAPIService Service

type ApiGetFutureHourlyInterestRateRequest struct {
	ctx        context.Context
	ApiService *BorrowRepayAPIService
	assets     *string
	isIsolated *models.GetFutureHourlyInterestRateIsIsolatedParameter
}

func (r ApiGetFutureHourlyInterestRateRequest) Assets(assets string) ApiGetFutureHourlyInterestRateRequest {
	r.assets = &assets
	return r
}

func (r ApiGetFutureHourlyInterestRateRequest) IsIsolated(isIsolated models.GetFutureHourlyInterestRateIsIsolatedParameter) ApiGetFutureHourlyInterestRateRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiGetFutureHourlyInterestRateRequest) Execute() (*common.RestApiResponse[models.GetFutureHourlyInterestRateResponse], error) {
	return r.ApiService.GetFutureHourlyInterestRateExecute(r)
}

/*
GetFutureHourlyInterestRate Get future hourly interest rate (USER_DATA)
Get /sapi/v1/margin/next-hourly-interest-rate

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#get-future-hourly-interest-rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param assets -
@param isIsolated -
@return ApiGetFutureHourlyInterestRateRequest
*/
func (a *BorrowRepayAPIService) GetFutureHourlyInterestRate(ctx context.Context) ApiGetFutureHourlyInterestRateRequest {
	return ApiGetFutureHourlyInterestRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFutureHourlyInterestRateResponse
func (a *BorrowRepayAPIService) GetFutureHourlyInterestRateExecute(r ApiGetFutureHourlyInterestRateRequest) (*common.RestApiResponse[models.GetFutureHourlyInterestRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/next-hourly-interest-rate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.assets == nil {
		return nil, common.ReportError("assets is required and must be specified")
	}

	if r.isIsolated == nil {
		return nil, common.ReportError("isIsolated is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "assets", r.assets, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")

	resp, err := SendRequest[models.GetFutureHourlyInterestRateResponse](
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

type ApiGetInterestHistoryRequest struct {
	ctx            context.Context
	ApiService     *BorrowRepayAPIService
	asset          *string
	isolatedSymbol *string
	startTime      *int64
	endTime        *int64
	current        *int64
	size           *int64
	recvWindow     *int64
}

func (r ApiGetInterestHistoryRequest) Asset(asset string) ApiGetInterestHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetInterestHistoryRequest) IsolatedSymbol(isolatedSymbol string) ApiGetInterestHistoryRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiGetInterestHistoryRequest) StartTime(startTime int64) ApiGetInterestHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetInterestHistoryRequest) EndTime(endTime int64) ApiGetInterestHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetInterestHistoryRequest) Current(current int64) ApiGetInterestHistoryRequest {
	r.current = &current
	return r
}

func (r ApiGetInterestHistoryRequest) Size(size int64) ApiGetInterestHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetInterestHistoryRequest) RecvWindow(recvWindow int64) ApiGetInterestHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetInterestHistoryRequest) Execute() (*common.RestApiResponse[models.GetInterestHistoryResponse], error) {
	return r.ApiService.GetInterestHistoryExecute(r)
}

/*
GetInterestHistory Get Interest History (USER_DATA)
Get /sapi/v1/margin/interestHistory

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#get-interest-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param isolatedSymbol -
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param current -
@param size -
@param recvWindow -
@return ApiGetInterestHistoryRequest
*/
func (a *BorrowRepayAPIService) GetInterestHistory(ctx context.Context) ApiGetInterestHistoryRequest {
	return ApiGetInterestHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetInterestHistoryResponse
func (a *BorrowRepayAPIService) GetInterestHistoryExecute(r ApiGetInterestHistoryRequest) (*common.RestApiResponse[models.GetInterestHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/interestHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetInterestHistoryResponse](
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

type ApiMarginAccountBorrowRepayRequest struct {
	ctx        context.Context
	ApiService *BorrowRepayAPIService
	asset      *string
	isIsolated *models.MarginAccountBorrowRepayIsIsolatedParameter
	amount     *string
	type_      *models.QueryBorrowRepayRecordsInMarginAccountTypeParameter
	symbol     *string
	recvWindow *int64
}

func (r ApiMarginAccountBorrowRepayRequest) Asset(asset string) ApiMarginAccountBorrowRepayRequest {
	r.asset = &asset
	return r
}

// &#x60;TRUE&#x60; for Isolated Margin, &#x60;FALSE&#x60; for Cross Margin
func (r ApiMarginAccountBorrowRepayRequest) IsIsolated(isIsolated models.MarginAccountBorrowRepayIsIsolatedParameter) ApiMarginAccountBorrowRepayRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountBorrowRepayRequest) Amount(amount string) ApiMarginAccountBorrowRepayRequest {
	r.amount = &amount
	return r
}

func (r ApiMarginAccountBorrowRepayRequest) Type(type_ models.QueryBorrowRepayRecordsInMarginAccountTypeParameter) ApiMarginAccountBorrowRepayRequest {
	r.type_ = &type_
	return r
}

// Only for Isolated margin
func (r ApiMarginAccountBorrowRepayRequest) Symbol(symbol string) ApiMarginAccountBorrowRepayRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountBorrowRepayRequest) RecvWindow(recvWindow int64) ApiMarginAccountBorrowRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountBorrowRepayRequest) Execute() (*common.RestApiResponse[models.MarginAccountBorrowRepayResponse], error) {
	return r.ApiService.MarginAccountBorrowRepayExecute(r)
}

/*
MarginAccountBorrowRepay Margin account borrow/repay (USER_DATA)
Post /sapi/v1/margin/borrow-repay

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#margin-account-borrow-repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param isIsolated -  `TRUE` for Isolated Margin, `FALSE` for Cross Margin
@param amount -
@param type_ -
@param symbol -  Only for Isolated margin
@param recvWindow -
@return ApiMarginAccountBorrowRepayRequest
*/
func (a *BorrowRepayAPIService) MarginAccountBorrowRepay(ctx context.Context) ApiMarginAccountBorrowRepayRequest {
	return ApiMarginAccountBorrowRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountBorrowRepayResponse
func (a *BorrowRepayAPIService) MarginAccountBorrowRepayExecute(r ApiMarginAccountBorrowRepayRequest) (*common.RestApiResponse[models.MarginAccountBorrowRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/borrow-repay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	if r.isIsolated == nil {
		return nil, common.ReportError("isIsolated is required and must be specified")
	}

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountBorrowRepayResponse](
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

type ApiQueryBorrowRepayRecordsInMarginAccountRequest struct {
	ctx            context.Context
	ApiService     *BorrowRepayAPIService
	type_          *models.QueryBorrowRepayRecordsInMarginAccountTypeParameter
	asset          *string
	isolatedSymbol *string
	txId           *int64
	startTime      *int64
	endTime        *int64
	current        *int64
	size           *int64
	recvWindow     *int64
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) Type(type_ models.QueryBorrowRepayRecordsInMarginAccountTypeParameter) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.type_ = &type_
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) Asset(asset string) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) IsolatedSymbol(isolatedSymbol string) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) TxId(txId int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.txId = &txId
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) StartTime(startTime int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) EndTime(endTime int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) Current(current int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.current = &current
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) Size(size int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.size = &size
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) RecvWindow(recvWindow int64) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryBorrowRepayRecordsInMarginAccountRequest) Execute() (*common.RestApiResponse[models.QueryBorrowRepayRecordsInMarginAccountResponse], error) {
	return r.ApiService.QueryBorrowRepayRecordsInMarginAccountExecute(r)
}

/*
QueryBorrowRepayRecordsInMarginAccount Query borrow/repay records in Margin account (USER_DATA)
Get /sapi/v1/margin/borrow-repay

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#query-borrow-repay-records-in-margin-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -
@param asset -
@param isolatedSymbol -
@param txId -
@param startTime -
@param endTime -
@param current -
@param size -
@param recvWindow -
@return ApiQueryBorrowRepayRecordsInMarginAccountRequest
*/
func (a *BorrowRepayAPIService) QueryBorrowRepayRecordsInMarginAccount(ctx context.Context) ApiQueryBorrowRepayRecordsInMarginAccountRequest {
	return ApiQueryBorrowRepayRecordsInMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryBorrowRepayRecordsInMarginAccountResponse
func (a *BorrowRepayAPIService) QueryBorrowRepayRecordsInMarginAccountExecute(r ApiQueryBorrowRepayRecordsInMarginAccountRequest) (*common.RestApiResponse[models.QueryBorrowRepayRecordsInMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/borrow-repay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryBorrowRepayRecordsInMarginAccountResponse](
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

type ApiQueryMarginInterestRateHistoryRequest struct {
	ctx        context.Context
	ApiService *BorrowRepayAPIService
	asset      *string
	vipLevel   *int64
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiQueryMarginInterestRateHistoryRequest) Asset(asset string) ApiQueryMarginInterestRateHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryMarginInterestRateHistoryRequest) VipLevel(vipLevel int64) ApiQueryMarginInterestRateHistoryRequest {
	r.vipLevel = &vipLevel
	return r
}

func (r ApiQueryMarginInterestRateHistoryRequest) StartTime(startTime int64) ApiQueryMarginInterestRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginInterestRateHistoryRequest) EndTime(endTime int64) ApiQueryMarginInterestRateHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryMarginInterestRateHistoryRequest) RecvWindow(recvWindow int64) ApiQueryMarginInterestRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginInterestRateHistoryRequest) Execute() (*common.RestApiResponse[models.QueryMarginInterestRateHistoryResponse], error) {
	return r.ApiService.QueryMarginInterestRateHistoryExecute(r)
}

/*
QueryMarginInterestRateHistory Query Margin Interest Rate History (USER_DATA)
Get /sapi/v1/margin/interestRateHistory

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#query-margin-interest-rate-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param vipLevel -
@param startTime -
@param endTime -
@param recvWindow -
@return ApiQueryMarginInterestRateHistoryRequest
*/
func (a *BorrowRepayAPIService) QueryMarginInterestRateHistory(ctx context.Context) ApiQueryMarginInterestRateHistoryRequest {
	return ApiQueryMarginInterestRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginInterestRateHistoryResponse
func (a *BorrowRepayAPIService) QueryMarginInterestRateHistoryExecute(r ApiQueryMarginInterestRateHistoryRequest) (*common.RestApiResponse[models.QueryMarginInterestRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/interestRateHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.vipLevel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "vipLevel", r.vipLevel, "form", "")
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

	resp, err := SendRequest[models.QueryMarginInterestRateHistoryResponse](
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

type ApiQueryMaxBorrowRequest struct {
	ctx            context.Context
	ApiService     *BorrowRepayAPIService
	asset          *string
	isolatedSymbol *string
	recvWindow     *int64
}

func (r ApiQueryMaxBorrowRequest) Asset(asset string) ApiQueryMaxBorrowRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryMaxBorrowRequest) IsolatedSymbol(isolatedSymbol string) ApiQueryMaxBorrowRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

func (r ApiQueryMaxBorrowRequest) RecvWindow(recvWindow int64) ApiQueryMaxBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMaxBorrowRequest) Execute() (*common.RestApiResponse[models.QueryMaxBorrowResponse], error) {
	return r.ApiService.QueryMaxBorrowExecute(r)
}

/*
QueryMaxBorrow Query Max Borrow (USER_DATA)
Get /sapi/v1/margin/maxBorrowable

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/borrow-repay#query-max-borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param isolatedSymbol -
@param recvWindow -
@return ApiQueryMaxBorrowRequest
*/
func (a *BorrowRepayAPIService) QueryMaxBorrow(ctx context.Context) ApiQueryMaxBorrowRequest {
	return ApiQueryMaxBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMaxBorrowResponse
func (a *BorrowRepayAPIService) QueryMaxBorrowExecute(r ApiQueryMaxBorrowRequest) (*common.RestApiResponse[models.QueryMaxBorrowResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/maxBorrowable"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMaxBorrowResponse](
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
