/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package binanceviploanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiGetBorrowInterestRateRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	loanCoin   *string
	recvWindow *int64
}

// Max 10 assets, Multiple split by \&quot;,\&quot;
func (r ApiGetBorrowInterestRateRequest) LoanCoin(loanCoin string) ApiGetBorrowInterestRateRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetBorrowInterestRateRequest) RecvWindow(recvWindow int64) ApiGetBorrowInterestRateRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBorrowInterestRateRequest) Execute() (*common.RestApiResponse[models.GetBorrowInterestRateResponse], error) {
	return r.ApiService.GetBorrowInterestRateExecute(r)
}

/*
GetBorrowInterestRate Get Borrow Interest Rate (USER_DATA)
Get /sapi/v1/loan/vip/request/interestRate

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/market-data#get-borrow-interest-rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -  Max 10 assets, Multiple split by \",\"
@param recvWindow -
@return ApiGetBorrowInterestRateRequest
*/
func (a *MarketDataAPIService) GetBorrowInterestRate(ctx context.Context) ApiGetBorrowInterestRateRequest {
	return ApiGetBorrowInterestRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBorrowInterestRateResponse
func (a *MarketDataAPIService) GetBorrowInterestRateExecute(r ApiGetBorrowInterestRateRequest) (*common.RestApiResponse[models.GetBorrowInterestRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/request/interestRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetBorrowInterestRateResponse](
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

type ApiGetCollateralAssetDataRequest struct {
	ctx            context.Context
	ApiService     *MarketDataAPIService
	collateralCoin *string
	recvWindow     *int64
}

func (r ApiGetCollateralAssetDataRequest) CollateralCoin(collateralCoin string) ApiGetCollateralAssetDataRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetCollateralAssetDataRequest) RecvWindow(recvWindow int64) ApiGetCollateralAssetDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCollateralAssetDataRequest) Execute() (*common.RestApiResponse[models.GetCollateralAssetDataResponse], error) {
	return r.ApiService.GetCollateralAssetDataExecute(r)
}

/*
GetCollateralAssetData Get Collateral Asset Data (USER_DATA)
Get /sapi/v1/loan/vip/collateral/data

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/market-data#get-collateral-asset-data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param collateralCoin -
@param recvWindow -
@return ApiGetCollateralAssetDataRequest
*/
func (a *MarketDataAPIService) GetCollateralAssetData(ctx context.Context) ApiGetCollateralAssetDataRequest {
	return ApiGetCollateralAssetDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCollateralAssetDataResponse
func (a *MarketDataAPIService) GetCollateralAssetDataExecute(r ApiGetCollateralAssetDataRequest) (*common.RestApiResponse[models.GetCollateralAssetDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/collateral/data"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCollateralAssetDataResponse](
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

type ApiGetLoanableAssetsDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	loanCoin   *string
	vipLevel   *int64
	recvWindow *int64
}

func (r ApiGetLoanableAssetsDataRequest) LoanCoin(loanCoin string) ApiGetLoanableAssetsDataRequest {
	r.loanCoin = &loanCoin
	return r
}

// Defaults to the user&#39;s VIP level.
func (r ApiGetLoanableAssetsDataRequest) VipLevel(vipLevel int64) ApiGetLoanableAssetsDataRequest {
	r.vipLevel = &vipLevel
	return r
}

func (r ApiGetLoanableAssetsDataRequest) RecvWindow(recvWindow int64) ApiGetLoanableAssetsDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLoanableAssetsDataRequest) Execute() (*common.RestApiResponse[models.GetLoanableAssetsDataResponse], error) {
	return r.ApiService.GetLoanableAssetsDataExecute(r)
}

/*
GetLoanableAssetsData Get Loanable Assets Data (USER_DATA)
Get /sapi/v1/loan/vip/loanable/data

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/market-data#get-loanable-assets-data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param vipLevel -  Defaults to the user's VIP level.
@param recvWindow -
@return ApiGetLoanableAssetsDataRequest
*/
func (a *MarketDataAPIService) GetLoanableAssetsData(ctx context.Context) ApiGetLoanableAssetsDataRequest {
	return ApiGetLoanableAssetsDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLoanableAssetsDataResponse
func (a *MarketDataAPIService) GetLoanableAssetsDataExecute(r ApiGetLoanableAssetsDataRequest) (*common.RestApiResponse[models.GetLoanableAssetsDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/loanable/data"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.vipLevel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "vipLevel", r.vipLevel, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetLoanableAssetsDataResponse](
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

type ApiGetVIPLoanInterestRateHistoryRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	coin       *string
	recvWindow *int64
	startTime  *int64
	endTime    *int64
	current    *int64
	limit      *int64
}

func (r ApiGetVIPLoanInterestRateHistoryRequest) Coin(coin string) ApiGetVIPLoanInterestRateHistoryRequest {
	r.coin = &coin
	return r
}

func (r ApiGetVIPLoanInterestRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

// If both startTime and endTime are omitted, the most recent 90 days are returned.
func (r ApiGetVIPLoanInterestRateHistoryRequest) StartTime(startTime int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.startTime = &startTime
	return r
}

// Maximum interval between startTime and endTime is 180 days. Time is based on UTC+0.
func (r ApiGetVIPLoanInterestRateHistoryRequest) EndTime(endTime int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current page number, starting from 1.
func (r ApiGetVIPLoanInterestRateHistoryRequest) Current(current int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.current = &current
	return r
}

// Number of records per page.
func (r ApiGetVIPLoanInterestRateHistoryRequest) Limit(limit int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetVIPLoanInterestRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetVIPLoanInterestRateHistoryResponse], error) {
	return r.ApiService.GetVIPLoanInterestRateHistoryExecute(r)
}

/*
GetVIPLoanInterestRateHistory Get VIP Loan Interest Rate History (USER_DATA)
Get /sapi/v1/loan/vip/interestRateHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/market-data#get-viploan-interest-rate-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param recvWindow -
@param startTime -  If both startTime and endTime are omitted, the most recent 90 days are returned.
@param endTime -  Maximum interval between startTime and endTime is 180 days. Time is based on UTC+0.
@param current -  Current page number, starting from 1.
@param limit -  Number of records per page.
@return ApiGetVIPLoanInterestRateHistoryRequest
*/
func (a *MarketDataAPIService) GetVIPLoanInterestRateHistory(ctx context.Context) ApiGetVIPLoanInterestRateHistoryRequest {
	return ApiGetVIPLoanInterestRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetVIPLoanInterestRateHistoryResponse
func (a *MarketDataAPIService) GetVIPLoanInterestRateHistoryExecute(r ApiGetVIPLoanInterestRateHistoryRequest) (*common.RestApiResponse[models.GetVIPLoanInterestRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/interestRateHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}

	if r.recvWindow == nil {
		return nil, common.ReportError("recvWindow is required and must be specified")
	}
	if *r.recvWindow > 60000 {
		return nil, common.ReportError("recvWindow must be less than 60000")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")

	resp, err := SendRequest[models.GetVIPLoanInterestRateHistoryResponse](
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

type ApiQueryVIPLoanFixedRateMarketRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	loanCoin   *string
	duration   *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// Loan coin
func (r ApiQueryVIPLoanFixedRateMarketRequest) LoanCoin(loanCoin string) ApiQueryVIPLoanFixedRateMarketRequest {
	r.loanCoin = &loanCoin
	return r
}

// Duration in days, minimum 1
func (r ApiQueryVIPLoanFixedRateMarketRequest) Duration(duration int64) ApiQueryVIPLoanFixedRateMarketRequest {
	r.duration = &duration
	return r
}

// Page number, default 1, minimum 1
func (r ApiQueryVIPLoanFixedRateMarketRequest) Current(current int64) ApiQueryVIPLoanFixedRateMarketRequest {
	r.current = &current
	return r
}

// Page size, default 10, range [1, 100]
func (r ApiQueryVIPLoanFixedRateMarketRequest) Size(size int64) ApiQueryVIPLoanFixedRateMarketRequest {
	r.size = &size
	return r
}

// The value cannot be greater than &#x60;60000&#x60;
func (r ApiQueryVIPLoanFixedRateMarketRequest) RecvWindow(recvWindow int64) ApiQueryVIPLoanFixedRateMarketRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryVIPLoanFixedRateMarketRequest) Execute() (*common.RestApiResponse[models.QueryVIPLoanFixedRateMarketResponse], error) {
	return r.ApiService.QueryVIPLoanFixedRateMarketExecute(r)
}

/*
QueryVIPLoanFixedRateMarket Query VIP Loan Fixed Rate Market (USER_DATA)
Get /sapi/v1/loan/vip/fixed/market

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/market-data#query-viploan-fixed-rate-market

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -  Loan coin
@param duration -  Duration in days, minimum 1
@param current -  Page number, default 1, minimum 1
@param size -  Page size, default 10, range [1, 100]
@param recvWindow -  The value cannot be greater than `60000`
@return ApiQueryVIPLoanFixedRateMarketRequest
*/
func (a *MarketDataAPIService) QueryVIPLoanFixedRateMarket(ctx context.Context) ApiQueryVIPLoanFixedRateMarketRequest {
	return ApiQueryVIPLoanFixedRateMarketRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryVIPLoanFixedRateMarketResponse
func (a *MarketDataAPIService) QueryVIPLoanFixedRateMarketExecute(r ApiQueryVIPLoanFixedRateMarketRequest) (*common.RestApiResponse[models.QueryVIPLoanFixedRateMarketResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/fixed/market"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	if r.duration != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "form", "")
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

	resp, err := SendRequest[models.QueryVIPLoanFixedRateMarketResponse](
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
