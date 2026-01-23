/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package binancewalletrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountApiTradingStatusRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountApiTradingStatusRequest) RecvWindow(recvWindow int64) ApiAccountApiTradingStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountApiTradingStatusRequest) Execute() (*common.RestApiResponse[models.AccountApiTradingStatusResponse], error) {
	return r.ApiService.AccountApiTradingStatusExecute(r)
}

/*
AccountApiTradingStatus Account API Trading Status (USER_DATA)
Get /sapi/v1/account/apiTradingStatus

https://developers.binance.com/docs/wallet/account/Account-API-Trading-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountApiTradingStatusRequest
*/
func (a *AccountAPIService) AccountApiTradingStatus(ctx context.Context) ApiAccountApiTradingStatusRequest {
	return ApiAccountApiTradingStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountApiTradingStatusResponse
func (a *AccountAPIService) AccountApiTradingStatusExecute(r ApiAccountApiTradingStatusRequest) (*common.RestApiResponse[models.AccountApiTradingStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/apiTradingStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountApiTradingStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAccountInfoRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountInfoRequest) RecvWindow(recvWindow int64) ApiAccountInfoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInfoRequest) Execute() (*common.RestApiResponse[models.AccountInfoResponse], error) {
	return r.ApiService.AccountInfoExecute(r)
}

/*
AccountInfo Account info (USER_DATA)
Get /sapi/v1/account/info

https://developers.binance.com/docs/wallet/account/Account-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountInfoRequest
*/
func (a *AccountAPIService) AccountInfo(ctx context.Context) ApiAccountInfoRequest {
	return ApiAccountInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountInfoResponse
func (a *AccountAPIService) AccountInfoExecute(r ApiAccountInfoRequest) (*common.RestApiResponse[models.AccountInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/info"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAccountStatusRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountStatusRequest) RecvWindow(recvWindow int64) ApiAccountStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountStatusRequest) Execute() (*common.RestApiResponse[models.AccountStatusResponse], error) {
	return r.ApiService.AccountStatusExecute(r)
}

/*
AccountStatus Account Status (USER_DATA)
Get /sapi/v1/account/status

https://developers.binance.com/docs/wallet/account/Account-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountStatusRequest
*/
func (a *AccountAPIService) AccountStatus(ctx context.Context) ApiAccountStatusRequest {
	return ApiAccountStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountStatusResponse
func (a *AccountAPIService) AccountStatusExecute(r ApiAccountStatusRequest) (*common.RestApiResponse[models.AccountStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDailyAccountSnapshotRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	type_      *string
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiDailyAccountSnapshotRequest) Type(type_ string) ApiDailyAccountSnapshotRequest {
	r.type_ = &type_
	return r
}

func (r ApiDailyAccountSnapshotRequest) StartTime(startTime int64) ApiDailyAccountSnapshotRequest {
	r.startTime = &startTime
	return r
}

func (r ApiDailyAccountSnapshotRequest) EndTime(endTime int64) ApiDailyAccountSnapshotRequest {
	r.endTime = &endTime
	return r
}

// min 7, max 30, default 7
func (r ApiDailyAccountSnapshotRequest) Limit(limit int64) ApiDailyAccountSnapshotRequest {
	r.limit = &limit
	return r
}

func (r ApiDailyAccountSnapshotRequest) RecvWindow(recvWindow int64) ApiDailyAccountSnapshotRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDailyAccountSnapshotRequest) Execute() (*common.RestApiResponse[models.DailyAccountSnapshotResponse], error) {
	return r.ApiService.DailyAccountSnapshotExecute(r)
}

/*
DailyAccountSnapshot Daily Account Snapshot (USER_DATA)
Get /sapi/v1/accountSnapshot

https://developers.binance.com/docs/wallet/account/daily-account-snapshoot

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -
@param startTime -
@param endTime -
@param limit -  min 7, max 30, default 7
@param recvWindow -
@return ApiDailyAccountSnapshotRequest
*/
func (a *AccountAPIService) DailyAccountSnapshot(ctx context.Context) ApiDailyAccountSnapshotRequest {
	return ApiDailyAccountSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DailyAccountSnapshotResponse
func (a *AccountAPIService) DailyAccountSnapshotExecute(r ApiDailyAccountSnapshotRequest) (*common.RestApiResponse[models.DailyAccountSnapshotResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/accountSnapshot"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

	resp, err := SendRequest[models.DailyAccountSnapshotResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDisableFastWithdrawSwitchRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiDisableFastWithdrawSwitchRequest) RecvWindow(recvWindow int64) ApiDisableFastWithdrawSwitchRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDisableFastWithdrawSwitchRequest) Execute() (struct{}, error) {
	return r.ApiService.DisableFastWithdrawSwitchExecute(r)
}

/*
DisableFastWithdrawSwitch Disable Fast Withdraw Switch (USER_DATA)
Post /sapi/v1/account/disableFastWithdrawSwitch

https://developers.binance.com/docs/wallet/account/Disable-Fast-Withdraw-Switch

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiDisableFastWithdrawSwitchRequest
*/
func (a *AccountAPIService) DisableFastWithdrawSwitch(ctx context.Context) ApiDisableFastWithdrawSwitchRequest {
	return ApiDisableFastWithdrawSwitchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AccountAPIService) DisableFastWithdrawSwitchExecute(r ApiDisableFastWithdrawSwitchRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/disableFastWithdrawSwitch"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiEnableFastWithdrawSwitchRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiEnableFastWithdrawSwitchRequest) RecvWindow(recvWindow int64) ApiEnableFastWithdrawSwitchRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEnableFastWithdrawSwitchRequest) Execute() (struct{}, error) {
	return r.ApiService.EnableFastWithdrawSwitchExecute(r)
}

/*
EnableFastWithdrawSwitch Enable Fast Withdraw Switch (USER_DATA)
Post /sapi/v1/account/enableFastWithdrawSwitch

https://developers.binance.com/docs/wallet/account/Enable-Fast-Withdraw-Switch

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiEnableFastWithdrawSwitchRequest
*/
func (a *AccountAPIService) EnableFastWithdrawSwitch(ctx context.Context) ApiEnableFastWithdrawSwitchRequest {
	return ApiEnableFastWithdrawSwitchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AccountAPIService) EnableFastWithdrawSwitchExecute(r ApiEnableFastWithdrawSwitchRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/enableFastWithdrawSwitch"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiGetApiKeyPermissionRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetApiKeyPermissionRequest) RecvWindow(recvWindow int64) ApiGetApiKeyPermissionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetApiKeyPermissionRequest) Execute() (*common.RestApiResponse[models.GetApiKeyPermissionResponse], error) {
	return r.ApiService.GetApiKeyPermissionExecute(r)
}

/*
GetApiKeyPermission Get API Key Permission (USER_DATA)
Get /sapi/v1/account/apiRestrictions

https://developers.binance.com/docs/wallet/account/api-key-permission

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetApiKeyPermissionRequest
*/
func (a *AccountAPIService) GetApiKeyPermission(ctx context.Context) ApiGetApiKeyPermissionRequest {
	return ApiGetApiKeyPermissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetApiKeyPermissionResponse
func (a *AccountAPIService) GetApiKeyPermissionExecute(r ApiGetApiKeyPermissionRequest) (*common.RestApiResponse[models.GetApiKeyPermissionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/account/apiRestrictions"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetApiKeyPermissionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
