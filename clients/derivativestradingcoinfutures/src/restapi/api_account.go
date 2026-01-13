/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountInformationRequest) RecvWindow(recvWindow int64) ApiAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInformationRequest) Execute() (*common.RestApiResponse[models.AccountInformationResponse], error) {
	return r.ApiService.AccountInformationExecute(r)
}

/*
AccountInformation Account Information (USER_DATA)
Get /dapi/v1/account

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountInformationRequest
*/
func (a *AccountAPIService) AccountInformation(ctx context.Context) ApiAccountInformationRequest {
	return ApiAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountInformationResponse
func (a *AccountAPIService) AccountInformationExecute(r ApiAccountInformationRequest) (*common.RestApiResponse[models.AccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFuturesAccountBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiFuturesAccountBalanceRequest) RecvWindow(recvWindow int64) ApiFuturesAccountBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesAccountBalanceRequest) Execute() (*common.RestApiResponse[models.FuturesAccountBalanceResponse], error) {
	return r.ApiService.FuturesAccountBalanceExecute(r)
}

/*
FuturesAccountBalance Futures Account Balance (USER_DATA)
Get /dapi/v1/balance

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Futures-Account-Balance

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFuturesAccountBalanceRequest
*/
func (a *AccountAPIService) FuturesAccountBalance(ctx context.Context) ApiFuturesAccountBalanceRequest {
	return ApiFuturesAccountBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesAccountBalanceResponse
func (a *AccountAPIService) FuturesAccountBalanceExecute(r ApiFuturesAccountBalanceRequest) (*common.RestApiResponse[models.FuturesAccountBalanceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesAccountBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
Get /dapi/v1/positionSide/dual

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Current-Position-Mode

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/positionSide/dual"

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
Get /dapi/v1/order/asyn

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Order-History

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order/asyn"

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
Get /dapi/v1/trade/asyn

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Trade-History

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/trade/asyn"

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
Get /dapi/v1/income/asyn

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Download-Id-For-Futures-Transaction-History

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/income/asyn"

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
Get /dapi/v1/order/asyn/id

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Futures-Order-History-Download-Link-by-Id

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/order/asyn/id"

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
Get /dapi/v1/trade/asyn/id

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Futures-Trade-Download-Link-by-Id

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/trade/asyn/id"

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
Get /dapi/v1/income/asyn/id

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Futures-Transaction-History-Download-Link-by-Id

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/income/asyn/id"

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

// \&quot;TRANSFER\&quot;,\&quot;WELCOME_BONUS\&quot;, \&quot;FUNDING_FEE\&quot;, \&quot;REALIZED_PNL\&quot;, \&quot;COMMISSION\&quot;, \&quot;INSURANCE_CLEAR\&quot;, and \&quot;DELIVERED_SETTELMENT\&quot;
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
GetIncomeHistory Get Income History(USER_DATA)
Get /dapi/v1/income

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Get-Income-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param incomeType -  \"TRANSFER\",\"WELCOME_BONUS\", \"FUNDING_FEE\", \"REALIZED_PNL\", \"COMMISSION\", \"INSURANCE_CLEAR\", and \"DELIVERED_SETTELMENT\"
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/income"

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

type ApiNotionalBracketForPairRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	pair       *string
	recvWindow *int64
}

func (r ApiNotionalBracketForPairRequest) Pair(pair string) ApiNotionalBracketForPairRequest {
	r.pair = &pair
	return r
}

func (r ApiNotionalBracketForPairRequest) RecvWindow(recvWindow int64) ApiNotionalBracketForPairRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNotionalBracketForPairRequest) Execute() (*common.RestApiResponse[models.NotionalBracketForPairResponse], error) {
	return r.ApiService.NotionalBracketForPairExecute(r)
}

/*
NotionalBracketForPair Notional Bracket for Pair(USER_DATA)
Get /dapi/v1/leverageBracket

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Notional-Bracket-for-Pair

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -
@param recvWindow -
@return ApiNotionalBracketForPairRequest
*/
func (a *AccountAPIService) NotionalBracketForPair(ctx context.Context) ApiNotionalBracketForPairRequest {
	return ApiNotionalBracketForPairRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotionalBracketForPairResponse
func (a *AccountAPIService) NotionalBracketForPairExecute(r ApiNotionalBracketForPairRequest) (*common.RestApiResponse[models.NotionalBracketForPairResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NotionalBracketForPairResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiNotionalBracketForSymbolRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiNotionalBracketForSymbolRequest) Symbol(symbol string) ApiNotionalBracketForSymbolRequest {
	r.symbol = &symbol
	return r
}

func (r ApiNotionalBracketForSymbolRequest) RecvWindow(recvWindow int64) ApiNotionalBracketForSymbolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiNotionalBracketForSymbolRequest) Execute() (*common.RestApiResponse[models.NotionalBracketForSymbolResponse], error) {
	return r.ApiService.NotionalBracketForSymbolExecute(r)
}

/*
NotionalBracketForSymbol Notional Bracket for Symbol(USER_DATA)
Get /dapi/v2/leverageBracket

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/Notional-Bracket-for-Symbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiNotionalBracketForSymbolRequest
*/
func (a *AccountAPIService) NotionalBracketForSymbol(ctx context.Context) ApiNotionalBracketForSymbolRequest {
	return ApiNotionalBracketForSymbolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotionalBracketForSymbolResponse
func (a *AccountAPIService) NotionalBracketForSymbolExecute(r ApiNotionalBracketForSymbolRequest) (*common.RestApiResponse[models.NotionalBracketForSymbolResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v2/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.NotionalBracketForSymbolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
Get /dapi/v1/commissionRate

https://developers.binance.com/docs/derivatives/coin-margined-futures/account/rest-api/User-Commission-Rate

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/commissionRate"

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
