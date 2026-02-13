/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
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
GetBorrowInterestRate Get Borrow Interest Rate(USER_DATA)
Get /sapi/v1/loan/vip/request/interestRate

https://developers.binance.com/docs/vip_loan/market-data/Get-Borrow-Interest-Rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
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

	resp, err := SendRequest[models.GetBorrowInterestRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
GetCollateralAssetData Get Collateral Asset Data(USER_DATA)
Get /sapi/v1/loan/vip/collateral/data

https://developers.binance.com/docs/vip_loan/market-data/Get-Collateral-Asset-Data

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

	resp, err := SendRequest[models.GetCollateralAssetDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// default:user&#39;s vip level
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
GetLoanableAssetsData Get Loanable Assets Data(USER_DATA)
Get /sapi/v1/loan/vip/loanable/data

https://developers.binance.com/docs/vip_loan/market-data/Get-Loanable-Assets-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param vipLevel -  default:user's vip level
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

	resp, err := SendRequest[models.GetLoanableAssetsDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

func (r ApiGetVIPLoanInterestRateHistoryRequest) StartTime(startTime int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetVIPLoanInterestRateHistoryRequest) EndTime(endTime int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetVIPLoanInterestRateHistoryRequest) Current(current int64) ApiGetVIPLoanInterestRateHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
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

https://developers.binance.com/docs/vip_loan/market-data/Get-VIP-Loan-Interest-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param recvWindow -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
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

	resp, err := SendRequest[models.GetVIPLoanInterestRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
