/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package binancew3wpredictionrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// WalletAPIService WalletAPI Service
type WalletAPIService Service

type ApiGetPortfolioRequest struct {
	ctx           context.Context
	ApiService    *WalletAPIService
	walletAddress *string
	tokenId       *string
	marketId      *int64
	marketTopicId *int64
	activeOnly    *bool
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiGetPortfolioRequest) WalletAddress(walletAddress string) ApiGetPortfolioRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by prediction token ID
func (r ApiGetPortfolioRequest) TokenId(tokenId string) ApiGetPortfolioRequest {
	r.tokenId = &tokenId
	return r
}

// Filter by market ID. Must be &gt; 0
func (r ApiGetPortfolioRequest) MarketId(marketId int64) ApiGetPortfolioRequest {
	r.marketId = &marketId
	return r
}

// Filter by market topic ID. Must be &gt; 0
func (r ApiGetPortfolioRequest) MarketTopicId(marketTopicId int64) ApiGetPortfolioRequest {
	r.marketTopicId = &marketTopicId
	return r
}

// If &#x60;true&#x60;, return only active (unresolved) positions
func (r ApiGetPortfolioRequest) ActiveOnly(activeOnly bool) ApiGetPortfolioRequest {
	r.activeOnly = &activeOnly
	return r
}

// Request validity window in milliseconds
func (r ApiGetPortfolioRequest) RecvWindow(recvWindow int64) ApiGetPortfolioRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPortfolioRequest) Execute() (*common.RestApiResponse[models.GetPortfolioResponse], error) {
	return r.ApiService.GetPortfolioExecute(r)
}

/*
GetPortfolio Get Portfolio (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/pnl/portfolio

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/wallet#get-portfolio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tokenId -  Filter by prediction token ID
@param marketId -  Filter by market ID. Must be > 0
@param marketTopicId -  Filter by market topic ID. Must be > 0
@param activeOnly -  If `true`, return only active (unresolved) positions
@param recvWindow -  Request validity window in milliseconds
@return ApiGetPortfolioRequest
*/
func (a *WalletAPIService) GetPortfolio(ctx context.Context) ApiGetPortfolioRequest {
	return ApiGetPortfolioRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPortfolioResponse
func (a *WalletAPIService) GetPortfolioExecute(r ApiGetPortfolioRequest) (*common.RestApiResponse[models.GetPortfolioResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/pnl/portfolio"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.tokenId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")
	}
	if r.marketId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	}
	if r.marketTopicId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketTopicId", r.marketTopicId, "form", "")
	}
	if r.activeOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activeOnly", r.activeOnly, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPortfolioResponse](
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

type ApiGetQuotaStatusRequest struct {
	ctx        context.Context
	ApiService *WalletAPIService
	recvWindow *int64
}

// Request validity window in milliseconds
func (r ApiGetQuotaStatusRequest) RecvWindow(recvWindow int64) ApiGetQuotaStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetQuotaStatusRequest) Execute() (*common.RestApiResponse[models.GetQuotaStatusResponse], error) {
	return r.ApiService.GetQuotaStatusExecute(r)
}

/*
GetQuotaStatus Get Quota Status (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/quota/limit/status

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/wallet#get-quota-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Request validity window in milliseconds
@return ApiGetQuotaStatusRequest
*/
func (a *WalletAPIService) GetQuotaStatus(ctx context.Context) ApiGetQuotaStatusRequest {
	return ApiGetQuotaStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetQuotaStatusResponse
func (a *WalletAPIService) GetQuotaStatusExecute(r ApiGetQuotaStatusRequest) (*common.RestApiResponse[models.GetQuotaStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/quota/limit/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetQuotaStatusResponse](
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

type ApiListPredictionWalletsRequest struct {
	ctx        context.Context
	ApiService *WalletAPIService
	recvWindow *int64
}

// Request validity window in milliseconds
func (r ApiListPredictionWalletsRequest) RecvWindow(recvWindow int64) ApiListPredictionWalletsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiListPredictionWalletsRequest) Execute() (*common.RestApiResponse[models.ListPredictionWalletsResponse], error) {
	return r.ApiService.ListPredictionWalletsExecute(r)
}

/*
ListPredictionWallets List Prediction Wallets (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/wallet/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/wallet#list-prediction-wallets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Request validity window in milliseconds
@return ApiListPredictionWalletsRequest
*/
func (a *WalletAPIService) ListPredictionWallets(ctx context.Context) ApiListPredictionWalletsRequest {
	return ApiListPredictionWalletsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListPredictionWalletsResponse
func (a *WalletAPIService) ListPredictionWalletsExecute(r ApiListPredictionWalletsRequest) (*common.RestApiResponse[models.ListPredictionWalletsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/wallet/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ListPredictionWalletsResponse](
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

type ApiQueryPaymentOptionBalancesRequest struct {
	ctx        context.Context
	ApiService *WalletAPIService
	recvWindow *int64
}

// Request validity window in milliseconds
func (r ApiQueryPaymentOptionBalancesRequest) RecvWindow(recvWindow int64) ApiQueryPaymentOptionBalancesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPaymentOptionBalancesRequest) Execute() (*common.RestApiResponse[models.QueryPaymentOptionBalancesResponse], error) {
	return r.ApiService.QueryPaymentOptionBalancesExecute(r)
}

/*
QueryPaymentOptionBalances Query Payment Option Balances (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/balance/payment-options

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/wallet#query-payment-option-balances

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryPaymentOptionBalancesRequest
*/
func (a *WalletAPIService) QueryPaymentOptionBalances(ctx context.Context) ApiQueryPaymentOptionBalancesRequest {
	return ApiQueryPaymentOptionBalancesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPaymentOptionBalancesResponse
func (a *WalletAPIService) QueryPaymentOptionBalancesExecute(r ApiQueryPaymentOptionBalancesRequest) (*common.RestApiResponse[models.QueryPaymentOptionBalancesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/balance/payment-options"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPaymentOptionBalancesResponse](
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
