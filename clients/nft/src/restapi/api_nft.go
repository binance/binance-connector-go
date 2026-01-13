/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package binancenftrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/nft/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// NFTAPIService NFTAPI Service
type NFTAPIService Service

type ApiGetNFTAssetRequest struct {
	ctx        context.Context
	ApiService *NFTAPIService
	limit      *int64
	page       *int64
	recvWindow *int64
}

// Default 50, Max 50
func (r ApiGetNFTAssetRequest) Limit(limit int64) ApiGetNFTAssetRequest {
	r.limit = &limit
	return r
}

// Default 1
func (r ApiGetNFTAssetRequest) Page(page int64) ApiGetNFTAssetRequest {
	r.page = &page
	return r
}

func (r ApiGetNFTAssetRequest) RecvWindow(recvWindow int64) ApiGetNFTAssetRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetNFTAssetRequest) Execute() (*common.RestApiResponse[models.GetNFTAssetResponse], error) {
	return r.ApiService.GetNFTAssetExecute(r)
}

/*
GetNFTAsset Get NFT Asset(USER_DATA)
Get /sapi/v1/nft/user/getAsset

https://developers.binance.com/docs/nft/rest-api/Get-NFT-Asset

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param limit -  Default 50, Max 50
@param page -  Default 1
@param recvWindow -
@return ApiGetNFTAssetRequest
*/
func (a *NFTAPIService) GetNFTAsset(ctx context.Context) ApiGetNFTAssetRequest {
	return ApiGetNFTAssetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNFTAssetResponse
func (a *NFTAPIService) GetNFTAssetExecute(r ApiGetNFTAssetRequest) (*common.RestApiResponse[models.GetNFTAssetResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/nft/user/getAsset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetNFTAssetResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetNFTDepositHistoryRequest struct {
	ctx        context.Context
	ApiService *NFTAPIService
	startTime  *int64
	endTime    *int64
	limit      *int64
	page       *int64
	recvWindow *int64
}

func (r ApiGetNFTDepositHistoryRequest) StartTime(startTime int64) ApiGetNFTDepositHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetNFTDepositHistoryRequest) EndTime(endTime int64) ApiGetNFTDepositHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 50, Max 50
func (r ApiGetNFTDepositHistoryRequest) Limit(limit int64) ApiGetNFTDepositHistoryRequest {
	r.limit = &limit
	return r
}

// Default 1
func (r ApiGetNFTDepositHistoryRequest) Page(page int64) ApiGetNFTDepositHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetNFTDepositHistoryRequest) RecvWindow(recvWindow int64) ApiGetNFTDepositHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetNFTDepositHistoryRequest) Execute() (*common.RestApiResponse[models.GetNFTDepositHistoryResponse], error) {
	return r.ApiService.GetNFTDepositHistoryExecute(r)
}

/*
GetNFTDepositHistory Get NFT Deposit History(USER_DATA)
Get /sapi/v1/nft/history/deposit

https://developers.binance.com/docs/nft/rest-api/Get-NFT-Deposit-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param limit -  Default 50, Max 50
@param page -  Default 1
@param recvWindow -
@return ApiGetNFTDepositHistoryRequest
*/
func (a *NFTAPIService) GetNFTDepositHistory(ctx context.Context) ApiGetNFTDepositHistoryRequest {
	return ApiGetNFTDepositHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNFTDepositHistoryResponse
func (a *NFTAPIService) GetNFTDepositHistoryExecute(r ApiGetNFTDepositHistoryRequest) (*common.RestApiResponse[models.GetNFTDepositHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/nft/history/deposit"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetNFTDepositHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetNFTTransactionHistoryRequest struct {
	ctx        context.Context
	ApiService *NFTAPIService
	orderType  *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	page       *int64
	recvWindow *int64
}

// 0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee
func (r ApiGetNFTTransactionHistoryRequest) OrderType(orderType int64) ApiGetNFTTransactionHistoryRequest {
	r.orderType = &orderType
	return r
}

func (r ApiGetNFTTransactionHistoryRequest) StartTime(startTime int64) ApiGetNFTTransactionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetNFTTransactionHistoryRequest) EndTime(endTime int64) ApiGetNFTTransactionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 50, Max 50
func (r ApiGetNFTTransactionHistoryRequest) Limit(limit int64) ApiGetNFTTransactionHistoryRequest {
	r.limit = &limit
	return r
}

// Default 1
func (r ApiGetNFTTransactionHistoryRequest) Page(page int64) ApiGetNFTTransactionHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetNFTTransactionHistoryRequest) RecvWindow(recvWindow int64) ApiGetNFTTransactionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetNFTTransactionHistoryRequest) Execute() (*common.RestApiResponse[models.GetNFTTransactionHistoryResponse], error) {
	return r.ApiService.GetNFTTransactionHistoryExecute(r)
}

/*
GetNFTTransactionHistory Get NFT Transaction History(USER_DATA)
Get /sapi/v1/nft/history/transactions

https://developers.binance.com/docs/nft/rest-api/Get-NFT-Transaction-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderType -  0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee
@param startTime -
@param endTime -
@param limit -  Default 50, Max 50
@param page -  Default 1
@param recvWindow -
@return ApiGetNFTTransactionHistoryRequest
*/
func (a *NFTAPIService) GetNFTTransactionHistory(ctx context.Context) ApiGetNFTTransactionHistoryRequest {
	return ApiGetNFTTransactionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNFTTransactionHistoryResponse
func (a *NFTAPIService) GetNFTTransactionHistoryExecute(r ApiGetNFTTransactionHistoryRequest) (*common.RestApiResponse[models.GetNFTTransactionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/nft/history/transactions"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderType == nil {
		return nil, common.ReportError("orderType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderType", r.orderType, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetNFTTransactionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetNFTWithdrawHistoryRequest struct {
	ctx        context.Context
	ApiService *NFTAPIService
	startTime  *int64
	endTime    *int64
	limit      *int64
	page       *int64
	recvWindow *int64
}

func (r ApiGetNFTWithdrawHistoryRequest) StartTime(startTime int64) ApiGetNFTWithdrawHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetNFTWithdrawHistoryRequest) EndTime(endTime int64) ApiGetNFTWithdrawHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 50, Max 50
func (r ApiGetNFTWithdrawHistoryRequest) Limit(limit int64) ApiGetNFTWithdrawHistoryRequest {
	r.limit = &limit
	return r
}

// Default 1
func (r ApiGetNFTWithdrawHistoryRequest) Page(page int64) ApiGetNFTWithdrawHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetNFTWithdrawHistoryRequest) RecvWindow(recvWindow int64) ApiGetNFTWithdrawHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetNFTWithdrawHistoryRequest) Execute() (*common.RestApiResponse[models.GetNFTWithdrawHistoryResponse], error) {
	return r.ApiService.GetNFTWithdrawHistoryExecute(r)
}

/*
GetNFTWithdrawHistory Get NFT Withdraw History(USER_DATA)
Get /sapi/v1/nft/history/withdraw

https://developers.binance.com/docs/nft/rest-api/Get-NFT-Withdraw-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param limit -  Default 50, Max 50
@param page -  Default 1
@param recvWindow -
@return ApiGetNFTWithdrawHistoryRequest
*/
func (a *NFTAPIService) GetNFTWithdrawHistory(ctx context.Context) ApiGetNFTWithdrawHistoryRequest {
	return ApiGetNFTWithdrawHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNFTWithdrawHistoryResponse
func (a *NFTAPIService) GetNFTWithdrawHistoryExecute(r ApiGetNFTWithdrawHistoryRequest) (*common.RestApiResponse[models.GetNFTWithdrawHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/nft/history/withdraw"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetNFTWithdrawHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
