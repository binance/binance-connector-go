/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountInformationV2Request struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountInformationV2Request) RecvWindow(recvWindow int64) ApiAccountInformationV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInformationV2Request) Execute() (*common.RestApiResponse[models.AccountInformationV2Response], error) {
	return r.ApiService.AccountInformationV2Execute(r)
}

/*
AccountInformationV2 Account Information V2(USER_DATA)
Get /fapi/v2/account

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Information-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountInformationV2Request
*/
func (a *AccountAPIService) AccountInformationV2(ctx context.Context) ApiAccountInformationV2Request {
	return ApiAccountInformationV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountInformationV2Response
func (a *AccountAPIService) AccountInformationV2Execute(r ApiAccountInformationV2Request) (*common.RestApiResponse[models.AccountInformationV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v2/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountInformationV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAccountInformationV3Request struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountInformationV3Request) RecvWindow(recvWindow int64) ApiAccountInformationV3Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInformationV3Request) Execute() (*common.RestApiResponse[models.AccountInformationV3Response], error) {
	return r.ApiService.AccountInformationV3Execute(r)
}

/*
AccountInformationV3 Account Information V3(USER_DATA)
Get /fapi/v3/account

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Information-V3

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountInformationV3Request
*/
func (a *AccountAPIService) AccountInformationV3(ctx context.Context) ApiAccountInformationV3Request {
	return ApiAccountInformationV3Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountInformationV3Response
func (a *AccountAPIService) AccountInformationV3Execute(r ApiAccountInformationV3Request) (*common.RestApiResponse[models.AccountInformationV3Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v3/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountInformationV3Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesAccountBalanceV2Request struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiFuturesAccountBalanceV2Request) RecvWindow(recvWindow int64) ApiFuturesAccountBalanceV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesAccountBalanceV2Request) Execute() (*common.RestApiResponse[models.FuturesAccountBalanceV2Response], error) {
	return r.ApiService.FuturesAccountBalanceV2Execute(r)
}

/*
FuturesAccountBalanceV2 Futures Account Balance V2 (USER_DATA)
Get /fapi/v2/balance

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFuturesAccountBalanceV2Request
*/
func (a *AccountAPIService) FuturesAccountBalanceV2(ctx context.Context) ApiFuturesAccountBalanceV2Request {
	return ApiFuturesAccountBalanceV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesAccountBalanceV2Response
func (a *AccountAPIService) FuturesAccountBalanceV2Execute(r ApiFuturesAccountBalanceV2Request) (*common.RestApiResponse[models.FuturesAccountBalanceV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v2/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesAccountBalanceV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesAccountBalanceV3Request struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiFuturesAccountBalanceV3Request) RecvWindow(recvWindow int64) ApiFuturesAccountBalanceV3Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesAccountBalanceV3Request) Execute() (*common.RestApiResponse[models.FuturesAccountBalanceV3Response], error) {
	return r.ApiService.FuturesAccountBalanceV3Execute(r)
}

/*
FuturesAccountBalanceV3 Futures Account Balance V3 (USER_DATA)
Get /fapi/v3/balance

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Account-Balance-V3

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFuturesAccountBalanceV3Request
*/
func (a *AccountAPIService) FuturesAccountBalanceV3(ctx context.Context) ApiFuturesAccountBalanceV3Request {
	return ApiFuturesAccountBalanceV3Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesAccountBalanceV3Response
func (a *AccountAPIService) FuturesAccountBalanceV3Execute(r ApiFuturesAccountBalanceV3Request) (*common.RestApiResponse[models.FuturesAccountBalanceV3Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v3/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesAccountBalanceV3Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesAccountConfigurationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiFuturesAccountConfigurationRequest) RecvWindow(recvWindow int64) ApiFuturesAccountConfigurationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesAccountConfigurationRequest) Execute() (*common.RestApiResponse[models.FuturesAccountConfigurationResponse], error) {
	return r.ApiService.FuturesAccountConfigurationExecute(r)
}

/*
FuturesAccountConfiguration Futures Account Configuration(USER_DATA)
Get /fapi/v1/accountConfig

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFuturesAccountConfigurationRequest
*/
func (a *AccountAPIService) FuturesAccountConfiguration(ctx context.Context) ApiFuturesAccountConfigurationRequest {
	return ApiFuturesAccountConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesAccountConfigurationResponse
func (a *AccountAPIService) FuturesAccountConfigurationExecute(r ApiFuturesAccountConfigurationRequest) (*common.RestApiResponse[models.FuturesAccountConfigurationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/accountConfig"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesAccountConfigurationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesTradingQuantitativeRulesIndicatorsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiFuturesTradingQuantitativeRulesIndicatorsRequest) Symbol(symbol string) ApiFuturesTradingQuantitativeRulesIndicatorsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiFuturesTradingQuantitativeRulesIndicatorsRequest) RecvWindow(recvWindow int64) ApiFuturesTradingQuantitativeRulesIndicatorsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesTradingQuantitativeRulesIndicatorsRequest) Execute() (*common.RestApiResponse[models.FuturesTradingQuantitativeRulesIndicatorsResponse], error) {
	return r.ApiService.FuturesTradingQuantitativeRulesIndicatorsExecute(r)
}

/*
FuturesTradingQuantitativeRulesIndicators Futures Trading Quantitative Rules Indicators (USER_DATA)
Get /fapi/v1/apiTradingStatus

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Futures-Trading-Quantitative-Rules-Indicators

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiFuturesTradingQuantitativeRulesIndicatorsRequest
*/
func (a *AccountAPIService) FuturesTradingQuantitativeRulesIndicators(ctx context.Context) ApiFuturesTradingQuantitativeRulesIndicatorsRequest {
	return ApiFuturesTradingQuantitativeRulesIndicatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesTradingQuantitativeRulesIndicatorsResponse
func (a *AccountAPIService) FuturesTradingQuantitativeRulesIndicatorsExecute(r ApiFuturesTradingQuantitativeRulesIndicatorsRequest) (*common.RestApiResponse[models.FuturesTradingQuantitativeRulesIndicatorsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/apiTradingStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesTradingQuantitativeRulesIndicatorsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

func (r ApiGetBnbBurnStatusRequest) RecvWindow(recvWindow int64) ApiGetBnbBurnStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBnbBurnStatusRequest) Execute() (*common.RestApiResponse[models.GetBnbBurnStatusResponse], error) {
	return r.ApiService.GetBnbBurnStatusExecute(r)
}

/*
GetBnbBurnStatus Get BNB Burn Status (USER_DATA)
Get /fapi/v1/feeBurn

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-BNB-Burn-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/feeBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetBnbBurnStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCurrentMultiAssetsModeRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetCurrentMultiAssetsModeRequest) RecvWindow(recvWindow int64) ApiGetCurrentMultiAssetsModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCurrentMultiAssetsModeRequest) Execute() (*common.RestApiResponse[models.GetCurrentMultiAssetsModeResponse], error) {
	return r.ApiService.GetCurrentMultiAssetsModeExecute(r)
}

/*
GetCurrentMultiAssetsMode Get Current Multi-Assets Mode (USER_DATA)
Get /fapi/v1/multiAssetsMargin

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Multi-Assets-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetCurrentMultiAssetsModeRequest
*/
func (a *AccountAPIService) GetCurrentMultiAssetsMode(ctx context.Context) ApiGetCurrentMultiAssetsModeRequest {
	return ApiGetCurrentMultiAssetsModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCurrentMultiAssetsModeResponse
func (a *AccountAPIService) GetCurrentMultiAssetsModeExecute(r ApiGetCurrentMultiAssetsModeRequest) (*common.RestApiResponse[models.GetCurrentMultiAssetsModeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/multiAssetsMargin"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCurrentMultiAssetsModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCurrentPositionModeRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetCurrentPositionModeRequest) RecvWindow(recvWindow int64) ApiGetCurrentPositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCurrentPositionModeRequest) Execute() (*common.RestApiResponse[models.GetCurrentPositionModeResponse], error) {
	return r.ApiService.GetCurrentPositionModeExecute(r)
}

/*
GetCurrentPositionMode Get Current Position Mode(USER_DATA)
Get /fapi/v1/positionSide/dual

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Position-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetCurrentPositionModeRequest
*/
func (a *AccountAPIService) GetCurrentPositionMode(ctx context.Context) ApiGetCurrentPositionModeRequest {
	return ApiGetCurrentPositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCurrentPositionModeResponse
func (a *AccountAPIService) GetCurrentPositionModeExecute(r ApiGetCurrentPositionModeRequest) (*common.RestApiResponse[models.GetCurrentPositionModeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCurrentPositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForFuturesOrderHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesOrderHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForFuturesOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesOrderHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForFuturesOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForFuturesOrderHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForFuturesOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForFuturesOrderHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForFuturesOrderHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForFuturesOrderHistoryExecute(r)
}

/*
GetDownloadIdForFuturesOrderHistory Get Download Id For Futures Order History (USER_DATA)
Get /fapi/v1/order/asyn

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Timestamp in ms
@param endTime -  Timestamp in ms
@param recvWindow -
@return ApiGetDownloadIdForFuturesOrderHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForFuturesOrderHistory(ctx context.Context) ApiGetDownloadIdForFuturesOrderHistoryRequest {
	return ApiGetDownloadIdForFuturesOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForFuturesOrderHistoryResponse
func (a *AccountAPIService) GetDownloadIdForFuturesOrderHistoryExecute(r ApiGetDownloadIdForFuturesOrderHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForFuturesOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForFuturesOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForFuturesTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesTradeHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForFuturesTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesTradeHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForFuturesTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForFuturesTradeHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForFuturesTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForFuturesTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForFuturesTradeHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForFuturesTradeHistoryExecute(r)
}

/*
GetDownloadIdForFuturesTradeHistory Get Download Id For Futures Trade History (USER_DATA)
Get /fapi/v1/trade/asyn

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Trade-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Timestamp in ms
@param endTime -  Timestamp in ms
@param recvWindow -
@return ApiGetDownloadIdForFuturesTradeHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForFuturesTradeHistory(ctx context.Context) ApiGetDownloadIdForFuturesTradeHistoryRequest {
	return ApiGetDownloadIdForFuturesTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForFuturesTradeHistoryResponse
func (a *AccountAPIService) GetDownloadIdForFuturesTradeHistoryExecute(r ApiGetDownloadIdForFuturesTradeHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForFuturesTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/trade/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForFuturesTradeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForFuturesTransactionHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesTransactionHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForFuturesTransactionHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiGetDownloadIdForFuturesTransactionHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForFuturesTransactionHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForFuturesTransactionHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForFuturesTransactionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForFuturesTransactionHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForFuturesTransactionHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForFuturesTransactionHistoryExecute(r)
}

/*
GetDownloadIdForFuturesTransactionHistory Get Download Id For Futures Transaction History(USER_DATA)
Get /fapi/v1/income/asyn

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Transaction-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Timestamp in ms
@param endTime -  Timestamp in ms
@param recvWindow -
@return ApiGetDownloadIdForFuturesTransactionHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForFuturesTransactionHistory(ctx context.Context) ApiGetDownloadIdForFuturesTransactionHistoryRequest {
	return ApiGetDownloadIdForFuturesTransactionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForFuturesTransactionHistoryResponse
func (a *AccountAPIService) GetDownloadIdForFuturesTransactionHistoryExecute(r ApiGetDownloadIdForFuturesTransactionHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForFuturesTransactionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/income/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForFuturesTransactionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesOrderHistoryDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetFuturesOrderHistoryDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetFuturesOrderHistoryDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetFuturesOrderHistoryDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetFuturesOrderHistoryDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesOrderHistoryDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetFuturesOrderHistoryDownloadLinkByIdResponse], error) {
	return r.ApiService.GetFuturesOrderHistoryDownloadLinkByIdExecute(r)
}

/*
GetFuturesOrderHistoryDownloadLinkById Get Futures Order History Download Link by Id (USER_DATA)
Get /fapi/v1/order/asyn/id

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Futures-Order-History-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetFuturesOrderHistoryDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetFuturesOrderHistoryDownloadLinkById(ctx context.Context) ApiGetFuturesOrderHistoryDownloadLinkByIdRequest {
	return ApiGetFuturesOrderHistoryDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesOrderHistoryDownloadLinkByIdResponse
func (a *AccountAPIService) GetFuturesOrderHistoryDownloadLinkByIdExecute(r ApiGetFuturesOrderHistoryDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetFuturesOrderHistoryDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/order/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesOrderHistoryDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesTradeDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetFuturesTradeDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetFuturesTradeDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetFuturesTradeDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetFuturesTradeDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesTradeDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetFuturesTradeDownloadLinkByIdResponse], error) {
	return r.ApiService.GetFuturesTradeDownloadLinkByIdExecute(r)
}

/*
GetFuturesTradeDownloadLinkById Get Futures Trade Download Link by Id(USER_DATA)
Get /fapi/v1/trade/asyn/id

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Futures-Trade-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetFuturesTradeDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetFuturesTradeDownloadLinkById(ctx context.Context) ApiGetFuturesTradeDownloadLinkByIdRequest {
	return ApiGetFuturesTradeDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesTradeDownloadLinkByIdResponse
func (a *AccountAPIService) GetFuturesTradeDownloadLinkByIdExecute(r ApiGetFuturesTradeDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetFuturesTradeDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/trade/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesTradeDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetFuturesTransactionHistoryDownloadLinkByIdResponse], error) {
	return r.ApiService.GetFuturesTransactionHistoryDownloadLinkByIdExecute(r)
}

/*
GetFuturesTransactionHistoryDownloadLinkById Get Futures Transaction History Download Link by Id (USER_DATA)
Get /fapi/v1/income/asyn/id

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Futures-Transaction-History-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetFuturesTransactionHistoryDownloadLinkById(ctx context.Context) ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest {
	return ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesTransactionHistoryDownloadLinkByIdResponse
func (a *AccountAPIService) GetFuturesTransactionHistoryDownloadLinkByIdExecute(r ApiGetFuturesTransactionHistoryDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetFuturesTransactionHistoryDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/income/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesTransactionHistoryDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetIncomeHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetIncomeHistoryRequest) Symbol(symbol string) ApiGetIncomeHistoryRequest {
	r.symbol = &symbol
	return r
}

// TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD
func (r ApiGetIncomeHistoryRequest) IncomeType(incomeType string) ApiGetIncomeHistoryRequest {
	r.incomeType = &incomeType
	return r
}

func (r ApiGetIncomeHistoryRequest) StartTime(startTime int64) ApiGetIncomeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetIncomeHistoryRequest) EndTime(endTime int64) ApiGetIncomeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetIncomeHistoryRequest) Page(page int64) ApiGetIncomeHistoryRequest {
	r.page = &page
	return r
}

// Default 100; max 1000
func (r ApiGetIncomeHistoryRequest) Limit(limit int64) ApiGetIncomeHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetIncomeHistoryRequest) RecvWindow(recvWindow int64) ApiGetIncomeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetIncomeHistoryRequest) Execute() (*common.RestApiResponse[models.GetIncomeHistoryResponse], error) {
	return r.ApiService.GetIncomeHistoryExecute(r)
}

/*
GetIncomeHistory Get Income History (USER_DATA)
Get /fapi/v1/income

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Income-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param incomeType -  TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD
@param startTime -
@param endTime -
@param page -
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiGetIncomeHistoryRequest
*/
func (a *AccountAPIService) GetIncomeHistory(ctx context.Context) ApiGetIncomeHistoryRequest {
	return ApiGetIncomeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetIncomeHistoryResponse
func (a *AccountAPIService) GetIncomeHistoryExecute(r ApiGetIncomeHistoryRequest) (*common.RestApiResponse[models.GetIncomeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/income"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.incomeType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "incomeType", r.incomeType, "form", "")
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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetIncomeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNotionalAndLeverageBracketsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiNotionalAndLeverageBracketsRequest) Symbol(symbol string) ApiNotionalAndLeverageBracketsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNotionalAndLeverageBracketsRequest) RecvWindow(recvWindow int64) ApiNotionalAndLeverageBracketsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNotionalAndLeverageBracketsRequest) Execute() (*common.RestApiResponse[models.NotionalAndLeverageBracketsResponse], error) {
	return r.ApiService.NotionalAndLeverageBracketsExecute(r)
}

/*
NotionalAndLeverageBrackets Notional and Leverage Brackets (USER_DATA)
Get /fapi/v1/leverageBracket

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Notional-and-Leverage-Brackets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiNotionalAndLeverageBracketsRequest
*/
func (a *AccountAPIService) NotionalAndLeverageBrackets(ctx context.Context) ApiNotionalAndLeverageBracketsRequest {
	return ApiNotionalAndLeverageBracketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotionalAndLeverageBracketsResponse
func (a *AccountAPIService) NotionalAndLeverageBracketsExecute(r ApiNotionalAndLeverageBracketsRequest) (*common.RestApiResponse[models.NotionalAndLeverageBracketsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NotionalAndLeverageBracketsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserRateLimitRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiQueryUserRateLimitRequest) RecvWindow(recvWindow int64) ApiQueryUserRateLimitRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserRateLimitRequest) Execute() (*common.RestApiResponse[models.QueryUserRateLimitResponse], error) {
	return r.ApiService.QueryUserRateLimitExecute(r)
}

/*
QueryUserRateLimit Query User Rate Limit (USER_DATA)
Get /fapi/v1/rateLimit/order

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Query-Rate-Limit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryUserRateLimitRequest
*/
func (a *AccountAPIService) QueryUserRateLimit(ctx context.Context) ApiQueryUserRateLimitRequest {
	return ApiQueryUserRateLimitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserRateLimitResponse
func (a *AccountAPIService) QueryUserRateLimitExecute(r ApiQueryUserRateLimitRequest) (*common.RestApiResponse[models.QueryUserRateLimitResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/rateLimit/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUserRateLimitResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSymbolConfigurationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiSymbolConfigurationRequest) Symbol(symbol string) ApiSymbolConfigurationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolConfigurationRequest) RecvWindow(recvWindow int64) ApiSymbolConfigurationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSymbolConfigurationRequest) Execute() (*common.RestApiResponse[models.SymbolConfigurationResponse], error) {
	return r.ApiService.SymbolConfigurationExecute(r)
}

/*
SymbolConfiguration Symbol Configuration(USER_DATA)
Get /fapi/v1/symbolConfig

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Symbol-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiSymbolConfigurationRequest
*/
func (a *AccountAPIService) SymbolConfiguration(ctx context.Context) ApiSymbolConfigurationRequest {
	return ApiSymbolConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SymbolConfigurationResponse
func (a *AccountAPIService) SymbolConfigurationExecute(r ApiSymbolConfigurationRequest) (*common.RestApiResponse[models.SymbolConfigurationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/symbolConfig"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SymbolConfigurationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiToggleBnbBurnOnFuturesTradeRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	feeBurn    *string
	recvWindow *int64
}

// \&quot;true\&quot;: Fee Discount On; \&quot;false\&quot;: Fee Discount Off
func (r ApiToggleBnbBurnOnFuturesTradeRequest) FeeBurn(feeBurn string) ApiToggleBnbBurnOnFuturesTradeRequest {
	r.feeBurn = &feeBurn
	return r
}

func (r ApiToggleBnbBurnOnFuturesTradeRequest) RecvWindow(recvWindow int64) ApiToggleBnbBurnOnFuturesTradeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiToggleBnbBurnOnFuturesTradeRequest) Execute() (*common.RestApiResponse[models.ToggleBnbBurnOnFuturesTradeResponse], error) {
	return r.ApiService.ToggleBnbBurnOnFuturesTradeExecute(r)
}

/*
ToggleBnbBurnOnFuturesTrade Toggle BNB Burn On Futures Trade (TRADE)
Post /fapi/v1/feeBurn

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Toggle-BNB-Burn-On-Futures-Trade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param feeBurn -  \"true\": Fee Discount On; \"false\": Fee Discount Off
@param recvWindow -
@return ApiToggleBnbBurnOnFuturesTradeRequest
*/
func (a *AccountAPIService) ToggleBnbBurnOnFuturesTrade(ctx context.Context) ApiToggleBnbBurnOnFuturesTradeRequest {
	return ApiToggleBnbBurnOnFuturesTradeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToggleBnbBurnOnFuturesTradeResponse
func (a *AccountAPIService) ToggleBnbBurnOnFuturesTradeExecute(r ApiToggleBnbBurnOnFuturesTradeRequest) (*common.RestApiResponse[models.ToggleBnbBurnOnFuturesTradeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/feeBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.feeBurn == nil {
		return nil, common.ReportError("feeBurn is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "feeBurn", r.feeBurn, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ToggleBnbBurnOnFuturesTradeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUserCommissionRateRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiUserCommissionRateRequest) Symbol(symbol string) ApiUserCommissionRateRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUserCommissionRateRequest) RecvWindow(recvWindow int64) ApiUserCommissionRateRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUserCommissionRateRequest) Execute() (*common.RestApiResponse[models.UserCommissionRateResponse], error) {
	return r.ApiService.UserCommissionRateExecute(r)
}

/*
UserCommissionRate User Commission Rate (USER_DATA)
Get /fapi/v1/commissionRate

https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/User-Commission-Rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiUserCommissionRateRequest
*/
func (a *AccountAPIService) UserCommissionRate(ctx context.Context) ApiUserCommissionRateRequest {
	return ApiUserCommissionRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserCommissionRateResponse
func (a *AccountAPIService) UserCommissionRateExecute(r ApiUserCommissionRateRequest) (*common.RestApiResponse[models.UserCommissionRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/commissionRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UserCommissionRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
