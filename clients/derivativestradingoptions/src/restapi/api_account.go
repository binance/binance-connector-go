/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountFundingFlowRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	currency   *string
	recordId   *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Asset type, only support USDT  as of now
func (r ApiAccountFundingFlowRequest) Currency(currency string) ApiAccountFundingFlowRequest {
	r.currency = &currency
	return r
}

// Return the recordId and subsequent data, the latest data is returned by default, e.g 100000
func (r ApiAccountFundingFlowRequest) RecordId(recordId int64) ApiAccountFundingFlowRequest {
	r.recordId = &recordId
	return r
}

// Start Time, e.g 1593511200000
func (r ApiAccountFundingFlowRequest) StartTime(startTime int64) ApiAccountFundingFlowRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiAccountFundingFlowRequest) EndTime(endTime int64) ApiAccountFundingFlowRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiAccountFundingFlowRequest) Limit(limit int64) ApiAccountFundingFlowRequest {
	r.limit = &limit
	return r
}

func (r ApiAccountFundingFlowRequest) RecvWindow(recvWindow int64) ApiAccountFundingFlowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountFundingFlowRequest) Execute() (*common.RestApiResponse[models.AccountFundingFlowResponse], error) {
	return r.ApiService.AccountFundingFlowExecute(r)
}

/*
AccountFundingFlow Account Funding Flow (USER_DATA)
Get /eapi/v1/bill

https://developers.binance.com/docs/derivatives/option/account/Account-Funding-Flow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param currency -  Asset type, only support USDT  as of now
@param recordId -  Return the recordId and subsequent data, the latest data is returned by default, e.g 100000
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@param recvWindow -
@return ApiAccountFundingFlowRequest
*/
func (a *AccountAPIService) AccountFundingFlow(ctx context.Context) ApiAccountFundingFlowRequest {
	return ApiAccountFundingFlowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountFundingFlowResponse
func (a *AccountAPIService) AccountFundingFlowExecute(r ApiAccountFundingFlowRequest) (*common.RestApiResponse[models.AccountFundingFlowResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/bill"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.currency == nil {
		return nil, common.ReportError("currency is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "form", "")
	if r.recordId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recordId", r.recordId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountFundingFlowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForOptionTransactionHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

// Timestamp in ms
func (r ApiGetDownloadIdForOptionTransactionHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForOptionTransactionHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms
func (r ApiGetDownloadIdForOptionTransactionHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForOptionTransactionHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForOptionTransactionHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForOptionTransactionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForOptionTransactionHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForOptionTransactionHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForOptionTransactionHistoryExecute(r)
}

/*
GetDownloadIdForOptionTransactionHistory Get Download Id For Option Transaction History (USER_DATA)
Get /eapi/v1/income/asyn

https://developers.binance.com/docs/derivatives/option/account/Get-Download-Id-For-Option-Transaction-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Timestamp in ms
@param endTime -  Timestamp in ms
@param recvWindow -
@return ApiGetDownloadIdForOptionTransactionHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForOptionTransactionHistory(ctx context.Context) ApiGetDownloadIdForOptionTransactionHistoryRequest {
	return ApiGetDownloadIdForOptionTransactionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForOptionTransactionHistoryResponse
func (a *AccountAPIService) GetDownloadIdForOptionTransactionHistoryExecute(r ApiGetDownloadIdForOptionTransactionHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForOptionTransactionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/income/asyn"

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

	resp, err := SendRequest[models.GetDownloadIdForOptionTransactionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOptionTransactionHistoryDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetOptionTransactionHistoryDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetOptionTransactionHistoryDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetOptionTransactionHistoryDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetOptionTransactionHistoryDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOptionTransactionHistoryDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetOptionTransactionHistoryDownloadLinkByIdResponse], error) {
	return r.ApiService.GetOptionTransactionHistoryDownloadLinkByIdExecute(r)
}

/*
GetOptionTransactionHistoryDownloadLinkById Get Option Transaction History Download Link by Id (USER_DATA)
Get /eapi/v1/income/asyn/id

https://developers.binance.com/docs/derivatives/option/account/Get-Option-Transaction-History-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetOptionTransactionHistoryDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetOptionTransactionHistoryDownloadLinkById(ctx context.Context) ApiGetOptionTransactionHistoryDownloadLinkByIdRequest {
	return ApiGetOptionTransactionHistoryDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOptionTransactionHistoryDownloadLinkByIdResponse
func (a *AccountAPIService) GetOptionTransactionHistoryDownloadLinkByIdExecute(r ApiGetOptionTransactionHistoryDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetOptionTransactionHistoryDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/income/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOptionTransactionHistoryDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOptionAccountInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiOptionAccountInformationRequest) RecvWindow(recvWindow int64) ApiOptionAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOptionAccountInformationRequest) Execute() (*common.RestApiResponse[models.OptionAccountInformationResponse], error) {
	return r.ApiService.OptionAccountInformationExecute(r)
}

/*
OptionAccountInformation Option Account Information(TRADE)
Get /eapi/v1/account

https://developers.binance.com/docs/derivatives/option/account/Option-Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiOptionAccountInformationRequest
*/
func (a *AccountAPIService) OptionAccountInformation(ctx context.Context) ApiOptionAccountInformationRequest {
	return ApiOptionAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptionAccountInformationResponse
func (a *AccountAPIService) OptionAccountInformationExecute(r ApiOptionAccountInformationRequest) (*common.RestApiResponse[models.OptionAccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OptionAccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOptionMarginAccountInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiOptionMarginAccountInformationRequest) RecvWindow(recvWindow int64) ApiOptionMarginAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOptionMarginAccountInformationRequest) Execute() (*common.RestApiResponse[models.OptionMarginAccountInformationResponse], error) {
	return r.ApiService.OptionMarginAccountInformationExecute(r)
}

/*
OptionMarginAccountInformation Option Margin Account Information (USER_DATA)
Get /eapi/v1/marginAccount

https://developers.binance.com/docs/derivatives/option/account/Option-Margin-Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiOptionMarginAccountInformationRequest
*/
func (a *AccountAPIService) OptionMarginAccountInformation(ctx context.Context) ApiOptionMarginAccountInformationRequest {
	return ApiOptionMarginAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptionMarginAccountInformationResponse
func (a *AccountAPIService) OptionMarginAccountInformationExecute(r ApiOptionMarginAccountInformationRequest) (*common.RestApiResponse[models.OptionMarginAccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/marginAccount"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OptionMarginAccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
