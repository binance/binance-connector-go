/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAdjustCrossMarginMaxLeverageRequest struct {
	ctx         context.Context
	ApiService  *AccountAPIService
	maxLeverage *int64
}

// Can only adjust 3 , 5 or 10，Example: maxLeverage &#x3D; 5 or 3 for Cross Margin Classic; maxLeverage&#x3D;10 for Cross Margin Pro 10x leverage or 20x if compliance allows.
func (r ApiAdjustCrossMarginMaxLeverageRequest) MaxLeverage(maxLeverage int64) ApiAdjustCrossMarginMaxLeverageRequest {
	r.maxLeverage = &maxLeverage
	return r
}

func (r ApiAdjustCrossMarginMaxLeverageRequest) Execute() (*common.RestApiResponse[models.AdjustCrossMarginMaxLeverageResponse], error) {
	return r.ApiService.AdjustCrossMarginMaxLeverageExecute(r)
}

/*
AdjustCrossMarginMaxLeverage Adjust cross margin max leverage (USER_DATA)
Post /sapi/v1/margin/max-leverage

https://developers.binance.com/docs/margin_trading/account/Adjust-cross-margin-max-leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param maxLeverage -  Can only adjust 3 , 5 or 10，Example: maxLeverage = 5 or 3 for Cross Margin Classic; maxLeverage=10 for Cross Margin Pro 10x leverage or 20x if compliance allows.
@return ApiAdjustCrossMarginMaxLeverageRequest
*/
func (a *AccountAPIService) AdjustCrossMarginMaxLeverage(ctx context.Context) ApiAdjustCrossMarginMaxLeverageRequest {
	return ApiAdjustCrossMarginMaxLeverageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdjustCrossMarginMaxLeverageResponse
func (a *AccountAPIService) AdjustCrossMarginMaxLeverageExecute(r ApiAdjustCrossMarginMaxLeverageRequest) (*common.RestApiResponse[models.AdjustCrossMarginMaxLeverageResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/max-leverage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.maxLeverage == nil {
		return nil, common.ReportError("maxLeverage is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "maxLeverage", r.maxLeverage, "form", "")

	resp, err := SendRequest[models.AdjustCrossMarginMaxLeverageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDisableIsolatedMarginAccountRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiDisableIsolatedMarginAccountRequest) Symbol(symbol string) ApiDisableIsolatedMarginAccountRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiDisableIsolatedMarginAccountRequest) RecvWindow(recvWindow int64) ApiDisableIsolatedMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDisableIsolatedMarginAccountRequest) Execute() (*common.RestApiResponse[models.DisableIsolatedMarginAccountResponse], error) {
	return r.ApiService.DisableIsolatedMarginAccountExecute(r)
}

/*
DisableIsolatedMarginAccount Disable Isolated Margin Account (TRADE)
Delete /sapi/v1/margin/isolated/account

https://developers.binance.com/docs/margin_trading/account/Disable-Isolated-Margin-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -  No more than 60000
@return ApiDisableIsolatedMarginAccountRequest
*/
func (a *AccountAPIService) DisableIsolatedMarginAccount(ctx context.Context) ApiDisableIsolatedMarginAccountRequest {
	return ApiDisableIsolatedMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DisableIsolatedMarginAccountResponse
func (a *AccountAPIService) DisableIsolatedMarginAccountExecute(r ApiDisableIsolatedMarginAccountRequest) (*common.RestApiResponse[models.DisableIsolatedMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DisableIsolatedMarginAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiEnableIsolatedMarginAccountRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiEnableIsolatedMarginAccountRequest) Symbol(symbol string) ApiEnableIsolatedMarginAccountRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiEnableIsolatedMarginAccountRequest) RecvWindow(recvWindow int64) ApiEnableIsolatedMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEnableIsolatedMarginAccountRequest) Execute() (*common.RestApiResponse[models.EnableIsolatedMarginAccountResponse], error) {
	return r.ApiService.EnableIsolatedMarginAccountExecute(r)
}

/*
EnableIsolatedMarginAccount Enable Isolated Margin Account (TRADE)
Post /sapi/v1/margin/isolated/account

https://developers.binance.com/docs/margin_trading/account/Enable-Isolated-Margin-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -  No more than 60000
@return ApiEnableIsolatedMarginAccountRequest
*/
func (a *AccountAPIService) EnableIsolatedMarginAccount(ctx context.Context) ApiEnableIsolatedMarginAccountRequest {
	return ApiEnableIsolatedMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnableIsolatedMarginAccountResponse
func (a *AccountAPIService) EnableIsolatedMarginAccountExecute(r ApiEnableIsolatedMarginAccountRequest) (*common.RestApiResponse[models.EnableIsolatedMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EnableIsolatedMarginAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetBnbBurnStatusRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiGetBnbBurnStatusRequest) RecvWindow(recvWindow int64) ApiGetBnbBurnStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBnbBurnStatusRequest) Execute() (*common.RestApiResponse[models.GetBnbBurnStatusResponse], error) {
	return r.ApiService.GetBnbBurnStatusExecute(r)
}

/*
GetBnbBurnStatus Get BNB Burn Status (USER_DATA)
Get /sapi/v1/bnbBurn

https://developers.binance.com/docs/margin_trading/account/Get-BNB-Burn-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiGetBnbBurnStatusRequest
*/
func (a *AccountAPIService) GetBnbBurnStatus(ctx context.Context) ApiGetBnbBurnStatusRequest {
	return ApiGetBnbBurnStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBnbBurnStatusResponse
func (a *AccountAPIService) GetBnbBurnStatusExecute(r ApiGetBnbBurnStatusRequest) (*common.RestApiResponse[models.GetBnbBurnStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bnbBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetBnbBurnStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSummaryOfMarginAccountRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiGetSummaryOfMarginAccountRequest) RecvWindow(recvWindow int64) ApiGetSummaryOfMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSummaryOfMarginAccountRequest) Execute() (*common.RestApiResponse[models.GetSummaryOfMarginAccountResponse], error) {
	return r.ApiService.GetSummaryOfMarginAccountExecute(r)
}

/*
GetSummaryOfMarginAccount Get Summary of Margin account (USER_DATA)
Get /sapi/v1/margin/tradeCoeff

https://developers.binance.com/docs/margin_trading/account/Get-Summary-of-Margin-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiGetSummaryOfMarginAccountRequest
*/
func (a *AccountAPIService) GetSummaryOfMarginAccount(ctx context.Context) ApiGetSummaryOfMarginAccountRequest {
	return ApiGetSummaryOfMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSummaryOfMarginAccountResponse
func (a *AccountAPIService) GetSummaryOfMarginAccountExecute(r ApiGetSummaryOfMarginAccountRequest) (*common.RestApiResponse[models.GetSummaryOfMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/tradeCoeff"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSummaryOfMarginAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCrossIsolatedMarginCapitalFlowRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	symbol     *string
	type_      *string
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) Asset(asset string) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.asset = &asset
	return r
}

// isolated margin pair
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) Symbol(symbol string) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.symbol = &symbol
	return r
}

// Transfer Type: ROLL_IN, ROLL_OUT
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) Type(type_ string) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.type_ = &type_
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) StartTime(startTime int64) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) EndTime(endTime int64) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.endTime = &endTime
	return r
}

// If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned.
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) FromId(fromId int64) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.fromId = &fromId
	return r
}

// Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) Limit(limit int64) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.limit = &limit
	return r
}

// No more than 60000
func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) RecvWindow(recvWindow int64) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCrossIsolatedMarginCapitalFlowRequest) Execute() (*common.RestApiResponse[models.QueryCrossIsolatedMarginCapitalFlowResponse], error) {
	return r.ApiService.QueryCrossIsolatedMarginCapitalFlowExecute(r)
}

/*
QueryCrossIsolatedMarginCapitalFlow Query Cross Isolated Margin Capital Flow (USER_DATA)
Get /sapi/v1/margin/capital-flow

https://developers.binance.com/docs/margin_trading/account/Query-Cross-Isolated-Margin-Capital-Flow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param symbol -  isolated margin pair
@param type_ -  Transfer Type: ROLL_IN, ROLL_OUT
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param fromId -  If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned.
@param limit -  Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
@param recvWindow -  No more than 60000
@return ApiQueryCrossIsolatedMarginCapitalFlowRequest
*/
func (a *AccountAPIService) QueryCrossIsolatedMarginCapitalFlow(ctx context.Context) ApiQueryCrossIsolatedMarginCapitalFlowRequest {
	return ApiQueryCrossIsolatedMarginCapitalFlowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCrossIsolatedMarginCapitalFlowResponse
func (a *AccountAPIService) QueryCrossIsolatedMarginCapitalFlowExecute(r ApiQueryCrossIsolatedMarginCapitalFlowRequest) (*common.RestApiResponse[models.QueryCrossIsolatedMarginCapitalFlowResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/capital-flow"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCrossIsolatedMarginCapitalFlowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCrossMarginAccountDetailsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiQueryCrossMarginAccountDetailsRequest) RecvWindow(recvWindow int64) ApiQueryCrossMarginAccountDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCrossMarginAccountDetailsRequest) Execute() (*common.RestApiResponse[models.QueryCrossMarginAccountDetailsResponse], error) {
	return r.ApiService.QueryCrossMarginAccountDetailsExecute(r)
}

/*
QueryCrossMarginAccountDetails Query Cross Margin Account Details (USER_DATA)
Get /sapi/v1/margin/account

https://developers.binance.com/docs/margin_trading/account/Query-Cross-Margin-Account-Details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiQueryCrossMarginAccountDetailsRequest
*/
func (a *AccountAPIService) QueryCrossMarginAccountDetails(ctx context.Context) ApiQueryCrossMarginAccountDetailsRequest {
	return ApiQueryCrossMarginAccountDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCrossMarginAccountDetailsResponse
func (a *AccountAPIService) QueryCrossMarginAccountDetailsExecute(r ApiQueryCrossMarginAccountDetailsRequest) (*common.RestApiResponse[models.QueryCrossMarginAccountDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCrossMarginAccountDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCrossMarginFeeDataRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	vipLevel   *int64
	coin       *string
	recvWindow *int64
}

// User&#39;s current specific margin data will be returned if vipLevel is omitted
func (r ApiQueryCrossMarginFeeDataRequest) VipLevel(vipLevel int64) ApiQueryCrossMarginFeeDataRequest {
	r.vipLevel = &vipLevel
	return r
}

func (r ApiQueryCrossMarginFeeDataRequest) Coin(coin string) ApiQueryCrossMarginFeeDataRequest {
	r.coin = &coin
	return r
}

// No more than 60000
func (r ApiQueryCrossMarginFeeDataRequest) RecvWindow(recvWindow int64) ApiQueryCrossMarginFeeDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCrossMarginFeeDataRequest) Execute() (*common.RestApiResponse[models.QueryCrossMarginFeeDataResponse], error) {
	return r.ApiService.QueryCrossMarginFeeDataExecute(r)
}

/*
QueryCrossMarginFeeData Query Cross Margin Fee Data (USER_DATA)
Get /sapi/v1/margin/crossMarginData

https://developers.binance.com/docs/margin_trading/account/Query-Cross-Margin-Fee-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param vipLevel -  User's current specific margin data will be returned if vipLevel is omitted
@param coin -
@param recvWindow -  No more than 60000
@return ApiQueryCrossMarginFeeDataRequest
*/
func (a *AccountAPIService) QueryCrossMarginFeeData(ctx context.Context) ApiQueryCrossMarginFeeDataRequest {
	return ApiQueryCrossMarginFeeDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCrossMarginFeeDataResponse
func (a *AccountAPIService) QueryCrossMarginFeeDataExecute(r ApiQueryCrossMarginFeeDataRequest) (*common.RestApiResponse[models.QueryCrossMarginFeeDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/crossMarginData"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.vipLevel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "vipLevel", r.vipLevel, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCrossMarginFeeDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryEnabledIsolatedMarginAccountLimitRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiQueryEnabledIsolatedMarginAccountLimitRequest) RecvWindow(recvWindow int64) ApiQueryEnabledIsolatedMarginAccountLimitRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryEnabledIsolatedMarginAccountLimitRequest) Execute() (*common.RestApiResponse[models.QueryEnabledIsolatedMarginAccountLimitResponse], error) {
	return r.ApiService.QueryEnabledIsolatedMarginAccountLimitExecute(r)
}

/*
QueryEnabledIsolatedMarginAccountLimit Query Enabled Isolated Margin Account Limit (USER_DATA)
Get /sapi/v1/margin/isolated/accountLimit

https://developers.binance.com/docs/margin_trading/account/Query-Enabled-Isolated-Margin-Account-Limit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiQueryEnabledIsolatedMarginAccountLimitRequest
*/
func (a *AccountAPIService) QueryEnabledIsolatedMarginAccountLimit(ctx context.Context) ApiQueryEnabledIsolatedMarginAccountLimitRequest {
	return ApiQueryEnabledIsolatedMarginAccountLimitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryEnabledIsolatedMarginAccountLimitResponse
func (a *AccountAPIService) QueryEnabledIsolatedMarginAccountLimitExecute(r ApiQueryEnabledIsolatedMarginAccountLimitRequest) (*common.RestApiResponse[models.QueryEnabledIsolatedMarginAccountLimitResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/accountLimit"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryEnabledIsolatedMarginAccountLimitResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryIsolatedMarginAccountInfoRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbols    *string
	recvWindow *int64
}

// Max 5 symbols can be sent; separated by \&quot;,\&quot;. e.g. \&quot;BTCUSDT,BNBUSDT,ADAUSDT\&quot;
func (r ApiQueryIsolatedMarginAccountInfoRequest) Symbols(symbols string) ApiQueryIsolatedMarginAccountInfoRequest {
	r.symbols = &symbols
	return r
}

// No more than 60000
func (r ApiQueryIsolatedMarginAccountInfoRequest) RecvWindow(recvWindow int64) ApiQueryIsolatedMarginAccountInfoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryIsolatedMarginAccountInfoRequest) Execute() (*common.RestApiResponse[models.QueryIsolatedMarginAccountInfoResponse], error) {
	return r.ApiService.QueryIsolatedMarginAccountInfoExecute(r)
}

/*
QueryIsolatedMarginAccountInfo Query Isolated Margin Account Info (USER_DATA)
Get /sapi/v1/margin/isolated/account

https://developers.binance.com/docs/margin_trading/account/Query-Isolated-Margin-Account-Info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbols -  Max 5 symbols can be sent; separated by \",\". e.g. \"BTCUSDT,BNBUSDT,ADAUSDT\"
@param recvWindow -  No more than 60000
@return ApiQueryIsolatedMarginAccountInfoRequest
*/
func (a *AccountAPIService) QueryIsolatedMarginAccountInfo(ctx context.Context) ApiQueryIsolatedMarginAccountInfoRequest {
	return ApiQueryIsolatedMarginAccountInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryIsolatedMarginAccountInfoResponse
func (a *AccountAPIService) QueryIsolatedMarginAccountInfoExecute(r ApiQueryIsolatedMarginAccountInfoRequest) (*common.RestApiResponse[models.QueryIsolatedMarginAccountInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbols != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", r.symbols, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryIsolatedMarginAccountInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryIsolatedMarginFeeDataRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	vipLevel   *int64
	symbol     *string
	recvWindow *int64
}

// User&#39;s current specific margin data will be returned if vipLevel is omitted
func (r ApiQueryIsolatedMarginFeeDataRequest) VipLevel(vipLevel int64) ApiQueryIsolatedMarginFeeDataRequest {
	r.vipLevel = &vipLevel
	return r
}

// isolated margin pair
func (r ApiQueryIsolatedMarginFeeDataRequest) Symbol(symbol string) ApiQueryIsolatedMarginFeeDataRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiQueryIsolatedMarginFeeDataRequest) RecvWindow(recvWindow int64) ApiQueryIsolatedMarginFeeDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryIsolatedMarginFeeDataRequest) Execute() (*common.RestApiResponse[models.QueryIsolatedMarginFeeDataResponse], error) {
	return r.ApiService.QueryIsolatedMarginFeeDataExecute(r)
}

/*
QueryIsolatedMarginFeeData Query Isolated Margin Fee Data (USER_DATA)
Get /sapi/v1/margin/isolatedMarginData

https://developers.binance.com/docs/margin_trading/account/Query-Isolated-Margin-Fee-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param vipLevel -  User's current specific margin data will be returned if vipLevel is omitted
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
@return ApiQueryIsolatedMarginFeeDataRequest
*/
func (a *AccountAPIService) QueryIsolatedMarginFeeData(ctx context.Context) ApiQueryIsolatedMarginFeeDataRequest {
	return ApiQueryIsolatedMarginFeeDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryIsolatedMarginFeeDataResponse
func (a *AccountAPIService) QueryIsolatedMarginFeeDataExecute(r ApiQueryIsolatedMarginFeeDataRequest) (*common.RestApiResponse[models.QueryIsolatedMarginFeeDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolatedMarginData"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.vipLevel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "vipLevel", r.vipLevel, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryIsolatedMarginFeeDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
