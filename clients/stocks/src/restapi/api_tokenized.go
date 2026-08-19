/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package binancestocksrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TokenizedAPIService TokenizedAPI Service
type TokenizedAPIService Service

type ApiTokenizedConvertHistoryRequest struct {
	ctx        context.Context
	ApiService *TokenizedAPIService
	startTime  *int64
	endTime    *int64
	lastId     *int64
	size       *int32
	recvWindow *int64
}

// Start time (ms epoch).
func (r ApiTokenizedConvertHistoryRequest) StartTime(startTime int64) ApiTokenizedConvertHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time (ms epoch).
func (r ApiTokenizedConvertHistoryRequest) EndTime(endTime int64) ApiTokenizedConvertHistoryRequest {
	r.endTime = &endTime
	return r
}

// Last record id from the previous page. Omit (or leave unset) to fetch the first page.
func (r ApiTokenizedConvertHistoryRequest) LastId(lastId int64) ApiTokenizedConvertHistoryRequest {
	r.lastId = &lastId
	return r
}

// Page size. Default &#x60;20&#x60;, max &#x60;100&#x60;.
func (r ApiTokenizedConvertHistoryRequest) Size(size int32) ApiTokenizedConvertHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiTokenizedConvertHistoryRequest) RecvWindow(recvWindow int64) ApiTokenizedConvertHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTokenizedConvertHistoryRequest) Execute() (*common.RestApiResponse[models.TokenizedConvertHistoryResponse], error) {
	return r.ApiService.TokenizedConvertHistoryExecute(r)
}

/*
TokenizedConvertHistory Tokenized Convert History (USER_DATA)
Get /sapi/v1/equity/tokenized/history

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/tokenized#tokenized-convert-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time (ms epoch).
@param endTime -  End time (ms epoch).
@param lastId -  Last record id from the previous page. Omit (or leave unset) to fetch the first page.
@param size -  Page size. Default `20`, max `100`.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiTokenizedConvertHistoryRequest
*/
func (a *TokenizedAPIService) TokenizedConvertHistory(ctx context.Context) ApiTokenizedConvertHistoryRequest {
	return ApiTokenizedConvertHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenizedConvertHistoryResponse
func (a *TokenizedAPIService) TokenizedConvertHistoryExecute(r ApiTokenizedConvertHistoryRequest) (*common.RestApiResponse[models.TokenizedConvertHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/tokenized/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.lastId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "lastId", r.lastId, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TokenizedConvertHistoryResponse](
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

type ApiTokenizedConvertStatusRequest struct {
	ctx             context.Context
	ApiService      *TokenizedAPIService
	issuerRequestId *string
	convertType     *models.TokenizedConvertStatusConvertTypeParameter
	recvWindow      *int64
}

// Convert request id returned by &#x60;/tokenized/mint&#x60; or &#x60;/redeem&#x60;.
func (r ApiTokenizedConvertStatusRequest) IssuerRequestId(issuerRequestId string) ApiTokenizedConvertStatusRequest {
	r.issuerRequestId = &issuerRequestId
	return r
}

// &#x60;MINT&#x60; or &#x60;REDEEM&#x60;.
func (r ApiTokenizedConvertStatusRequest) ConvertType(convertType models.TokenizedConvertStatusConvertTypeParameter) ApiTokenizedConvertStatusRequest {
	r.convertType = &convertType
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiTokenizedConvertStatusRequest) RecvWindow(recvWindow int64) ApiTokenizedConvertStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTokenizedConvertStatusRequest) Execute() (*common.RestApiResponse[models.TokenizedConvertStatusResponse], error) {
	return r.ApiService.TokenizedConvertStatusExecute(r)
}

/*
TokenizedConvertStatus Tokenized Convert Status (USER_DATA)
Get /sapi/v1/equity/tokenized/convert-status

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/tokenized#tokenized-convert-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param issuerRequestId -  Convert request id returned by `/tokenized/mint` or `/redeem`.
@param convertType -  `MINT` or `REDEEM`.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiTokenizedConvertStatusRequest
*/
func (a *TokenizedAPIService) TokenizedConvertStatus(ctx context.Context) ApiTokenizedConvertStatusRequest {
	return ApiTokenizedConvertStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenizedConvertStatusResponse
func (a *TokenizedAPIService) TokenizedConvertStatusExecute(r ApiTokenizedConvertStatusRequest) (*common.RestApiResponse[models.TokenizedConvertStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/tokenized/convert-status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.issuerRequestId == nil {
		return nil, common.ReportError("issuerRequestId is required and must be specified")
	}

	if r.convertType == nil {
		return nil, common.ReportError("convertType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "issuerRequestId", r.issuerRequestId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "convertType", r.convertType, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TokenizedConvertStatusResponse](
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

type ApiTokenizedMintRequest struct {
	ctx                   context.Context
	ApiService            *TokenizedAPIService
	underlyingAsset       *string
	underlyingAssetAmount *string
	clientOrderId         *string
	recvWindow            *int64
}

// Underlying US-equity ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Resolved against the active-symbol list; unknown tickers return &#x60;-26004&#x60;. The target tokenized asset is looked up from this field via &#x60;/market/tokenized-assets&#x60;.
func (r ApiTokenizedMintRequest) UnderlyingAsset(underlyingAsset string) ApiTokenizedMintRequest {
	r.underlyingAsset = &underlyingAsset
	return r
}

// Quantity of the underlying asset to mint from. Must be &gt; 0.
func (r ApiTokenizedMintRequest) UnderlyingAssetAmount(underlyingAssetAmount string) ApiTokenizedMintRequest {
	r.underlyingAssetAmount = &underlyingAssetAmount
	return r
}

// Client order id for idempotency. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted.
func (r ApiTokenizedMintRequest) ClientOrderId(clientOrderId string) ApiTokenizedMintRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiTokenizedMintRequest) RecvWindow(recvWindow int64) ApiTokenizedMintRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTokenizedMintRequest) Execute() (*common.RestApiResponse[models.TokenizedMintResponse], error) {
	return r.ApiService.TokenizedMintExecute(r)
}

/*
TokenizedMint Tokenized Mint (TRADE)
Post /sapi/v1/equity/tokenized/mint

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/tokenized#tokenized-mint

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlyingAsset -  Underlying US-equity ticker, e.g. `AAPL`, `TSLA`. Resolved against the active-symbol list; unknown tickers return `-26004`. The target tokenized asset is looked up from this field via `/market/tokenized-assets`.
@param underlyingAssetAmount -  Quantity of the underlying asset to mint from. Must be > 0.
@param clientOrderId -  Client order id for idempotency. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiTokenizedMintRequest
*/
func (a *TokenizedAPIService) TokenizedMint(ctx context.Context) ApiTokenizedMintRequest {
	return ApiTokenizedMintRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenizedMintResponse
func (a *TokenizedAPIService) TokenizedMintExecute(r ApiTokenizedMintRequest) (*common.RestApiResponse[models.TokenizedMintResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/tokenized/mint"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlyingAsset == nil {
		return nil, common.ReportError("underlyingAsset is required and must be specified")
	}

	if r.underlyingAssetAmount == nil {
		return nil, common.ReportError("underlyingAssetAmount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlyingAsset", r.underlyingAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlyingAssetAmount", r.underlyingAssetAmount, "form", "")
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TokenizedMintResponse](
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

type ApiTokenizedRedeemRequest struct {
	ctx                  context.Context
	ApiService           *TokenizedAPIService
	tokenizedAsset       *string
	tokenizedAssetAmount *string
	clientOrderId        *string
	recvWindow           *int64
}

// Tokenized asset to redeem, e.g. &#x60;AAPLB&#x60;. Not a US-equity ticker — this is the on-chain tokenized asset identifier. Unknown asset returns &#x60;-1102&#x60; (the message currently says the parameter was empty/malformed, but it was in fact sent — it is simply unknown). The target underlying ticker is looked up from this field via &#x60;/market/tokenized-assets&#x60;.
func (r ApiTokenizedRedeemRequest) TokenizedAsset(tokenizedAsset string) ApiTokenizedRedeemRequest {
	r.tokenizedAsset = &tokenizedAsset
	return r
}

// Quantity of the tokenized asset to redeem. Must be &gt; 0.
func (r ApiTokenizedRedeemRequest) TokenizedAssetAmount(tokenizedAssetAmount string) ApiTokenizedRedeemRequest {
	r.tokenizedAssetAmount = &tokenizedAssetAmount
	return r
}

// Client order id for idempotency. Format &#x60;^[a-zA-Z0-9-_]{32,36}$&#x60;. Auto-generated when omitted.
func (r ApiTokenizedRedeemRequest) ClientOrderId(clientOrderId string) ApiTokenizedRedeemRequest {
	r.clientOrderId = &clientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;.
func (r ApiTokenizedRedeemRequest) RecvWindow(recvWindow int64) ApiTokenizedRedeemRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTokenizedRedeemRequest) Execute() (*common.RestApiResponse[models.TokenizedRedeemResponse], error) {
	return r.ApiService.TokenizedRedeemExecute(r)
}

/*
TokenizedRedeem Tokenized Redeem (TRADE)
Post /sapi/v1/equity/tokenized/redeem

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/tokenized#tokenized-redeem

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param tokenizedAsset -  Tokenized asset to redeem, e.g. `AAPLB`. Not a US-equity ticker — this is the on-chain tokenized asset identifier. Unknown asset returns `-1102` (the message currently says the parameter was empty/malformed, but it was in fact sent — it is simply unknown). The target underlying ticker is looked up from this field via `/market/tokenized-assets`.
@param tokenizedAssetAmount -  Quantity of the tokenized asset to redeem. Must be > 0.
@param clientOrderId -  Client order id for idempotency. Format `^[a-zA-Z0-9-_]{32,36}$`. Auto-generated when omitted.
@param recvWindow -  The value cannot be greater than `60000`.
@return ApiTokenizedRedeemRequest
*/
func (a *TokenizedAPIService) TokenizedRedeem(ctx context.Context) ApiTokenizedRedeemRequest {
	return ApiTokenizedRedeemRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenizedRedeemResponse
func (a *TokenizedAPIService) TokenizedRedeemExecute(r ApiTokenizedRedeemRequest) (*common.RestApiResponse[models.TokenizedRedeemResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/tokenized/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.tokenizedAsset == nil {
		return nil, common.ReportError("tokenizedAsset is required and must be specified")
	}

	if r.tokenizedAssetAmount == nil {
		return nil, common.ReportError("tokenizedAssetAmount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenizedAsset", r.tokenizedAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenizedAssetAmount", r.tokenizedAssetAmount, "form", "")
	if r.clientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientOrderId", r.clientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TokenizedRedeemResponse](
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
