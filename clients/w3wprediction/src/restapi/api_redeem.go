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

// RedeemAPIService RedeemAPI Service
type RedeemAPIService Service

type ApiBatchRedeemRequest struct {
	ctx           context.Context
	ApiService    *RedeemAPIService
	walletAddress *string
	walletId      *string
	tokenIds      *[]string
	chainId       *string
}

// User&#39;s prediction wallet address
func (r ApiBatchRedeemRequest) WalletAddress(walletAddress string) ApiBatchRedeemRequest {
	r.walletAddress = &walletAddress
	return r
}

// Wallet ID
func (r ApiBatchRedeemRequest) WalletId(walletId string) ApiBatchRedeemRequest {
	r.walletId = &walletId
	return r
}

// List of prediction token IDs to redeem. Not empty. Example: &#x60;tokenIds&#x3D;112233&amp;tokenIds&#x3D;112234&#x60;
func (r ApiBatchRedeemRequest) TokenIds(tokenIds []string) ApiBatchRedeemRequest {
	r.tokenIds = &tokenIds
	return r
}

// Chain ID. Default &#x60;56&#x60; (BSC)
func (r ApiBatchRedeemRequest) ChainId(chainId string) ApiBatchRedeemRequest {
	r.chainId = &chainId
	return r
}

func (r ApiBatchRedeemRequest) Execute() (*common.RestApiResponse[models.BatchRedeemResponse], error) {
	return r.ApiService.BatchRedeemExecute(r)
}

/*
BatchRedeem Batch Redeem (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/batch-redeem

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/redeem#batch-redeem

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param walletId -  Wallet ID
@param tokenIds -  List of prediction token IDs to redeem. Not empty. Example: `tokenIds=112233&tokenIds=112234`
@param chainId -  Chain ID. Default `56` (BSC)
@return ApiBatchRedeemRequest
*/
func (a *RedeemAPIService) BatchRedeem(ctx context.Context) ApiBatchRedeemRequest {
	return ApiBatchRedeemRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BatchRedeemResponse
func (a *RedeemAPIService) BatchRedeemExecute(r ApiBatchRedeemRequest) (*common.RestApiResponse[models.BatchRedeemResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/batch-redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.walletId == nil {
		return nil, common.ReportError("walletId is required and must be specified")
	}

	if r.tokenIds == nil {
		return nil, common.ReportError("tokenIds is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletId", r.walletId, "form", "")
	{
		t := *r.tokenIds
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenIds", t, "form", "multi")
	}
	if r.chainId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "chainId", r.chainId, "form", "")
	}

	resp, err := SendRequest[models.BatchRedeemResponse](
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

type ApiGetRedeemStatusRequest struct {
	ctx           context.Context
	ApiService    *RedeemAPIService
	walletAddress *string
	txHash        *string
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiGetRedeemStatusRequest) WalletAddress(walletAddress string) ApiGetRedeemStatusRequest {
	r.walletAddress = &walletAddress
	return r
}

// Redeem transaction hash
func (r ApiGetRedeemStatusRequest) TxHash(txHash string) ApiGetRedeemStatusRequest {
	r.txHash = &txHash
	return r
}

// Request validity window in milliseconds
func (r ApiGetRedeemStatusRequest) RecvWindow(recvWindow int64) ApiGetRedeemStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRedeemStatusRequest) Execute() (*common.RestApiResponse[models.GetRedeemStatusResponse], error) {
	return r.ApiService.GetRedeemStatusExecute(r)
}

/*
GetRedeemStatus Get Redeem Status (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/redeem/status

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/redeem#get-redeem-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param txHash -  Redeem transaction hash
@param recvWindow -  Request validity window in milliseconds
@return ApiGetRedeemStatusRequest
*/
func (a *RedeemAPIService) GetRedeemStatus(ctx context.Context) ApiGetRedeemStatusRequest {
	return ApiGetRedeemStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRedeemStatusResponse
func (a *RedeemAPIService) GetRedeemStatusExecute(r ApiGetRedeemStatusRequest) (*common.RestApiResponse[models.GetRedeemStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/redeem/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.txHash == nil {
		return nil, common.ReportError("txHash is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txHash", r.txHash, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetRedeemStatusResponse](
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
