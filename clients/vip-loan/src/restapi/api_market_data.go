/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API

API version: 1.0.0
*/

package binanceviploanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
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

	resp, err := SendRequest[models.GetBorrowInterestRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

	resp, err := SendRequest[models.GetCollateralAssetDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

	resp, err := SendRequest[models.GetLoanableAssetsDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
