/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package binancesubaccountrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountManagementAPIService AccountManagementAPI Service
type AccountManagementAPIService Service

type ApiCreateAVirtualSubAccountRequest struct {
	ctx              context.Context
	ApiService       *AccountManagementAPIService
	subAccountString *string
	recvWindow       *int64
}

// Please input a string. We will create a virtual email using that string for you to register
func (r ApiCreateAVirtualSubAccountRequest) SubAccountString(subAccountString string) ApiCreateAVirtualSubAccountRequest {
	r.subAccountString = &subAccountString
	return r
}

func (r ApiCreateAVirtualSubAccountRequest) RecvWindow(recvWindow int64) ApiCreateAVirtualSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateAVirtualSubAccountRequest) Execute() (*common.RestApiResponse[models.CreateAVirtualSubAccountResponse], error) {
	return r.ApiService.CreateAVirtualSubAccountExecute(r)
}

/*
CreateAVirtualSubAccount Create a Virtual Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/virtualSubAccount

https://developers.binance.com/docs/sub_account/account-management/Create-a-Virtual-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param subAccountString -  Please input a string. We will create a virtual email using that string for you to register
@param recvWindow -
@return ApiCreateAVirtualSubAccountRequest
*/
func (a *AccountManagementAPIService) CreateAVirtualSubAccount(ctx context.Context) ApiCreateAVirtualSubAccountRequest {
	return ApiCreateAVirtualSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateAVirtualSubAccountResponse
func (a *AccountManagementAPIService) CreateAVirtualSubAccountExecute(r ApiCreateAVirtualSubAccountRequest) (*common.RestApiResponse[models.CreateAVirtualSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/virtualSubAccount"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.subAccountString == nil {
		return nil, common.ReportError("subAccountString is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountString", r.subAccountString, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateAVirtualSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiEnableFuturesForSubAccountRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiEnableFuturesForSubAccountRequest) Email(email string) ApiEnableFuturesForSubAccountRequest {
	r.email = &email
	return r
}

func (r ApiEnableFuturesForSubAccountRequest) RecvWindow(recvWindow int64) ApiEnableFuturesForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEnableFuturesForSubAccountRequest) Execute() (*common.RestApiResponse[models.EnableFuturesForSubAccountResponse], error) {
	return r.ApiService.EnableFuturesForSubAccountExecute(r)
}

/*
EnableFuturesForSubAccount Enable Futures for Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/futures/enable

https://developers.binance.com/docs/sub_account/account-management/Enable-Futures-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiEnableFuturesForSubAccountRequest
*/
func (a *AccountManagementAPIService) EnableFuturesForSubAccount(ctx context.Context) ApiEnableFuturesForSubAccountRequest {
	return ApiEnableFuturesForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnableFuturesForSubAccountResponse
func (a *AccountManagementAPIService) EnableFuturesForSubAccountExecute(r ApiEnableFuturesForSubAccountRequest) (*common.RestApiResponse[models.EnableFuturesForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/enable"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EnableFuturesForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiEnableOptionsForSubAccountRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiEnableOptionsForSubAccountRequest) Email(email string) ApiEnableOptionsForSubAccountRequest {
	r.email = &email
	return r
}

func (r ApiEnableOptionsForSubAccountRequest) RecvWindow(recvWindow int64) ApiEnableOptionsForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEnableOptionsForSubAccountRequest) Execute() (*common.RestApiResponse[models.EnableOptionsForSubAccountResponse], error) {
	return r.ApiService.EnableOptionsForSubAccountExecute(r)
}

/*
EnableOptionsForSubAccount Enable Options for Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/eoptions/enable

https://developers.binance.com/docs/sub_account/account-management/Enable-Options-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiEnableOptionsForSubAccountRequest
*/
func (a *AccountManagementAPIService) EnableOptionsForSubAccount(ctx context.Context) ApiEnableOptionsForSubAccountRequest {
	return ApiEnableOptionsForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnableOptionsForSubAccountResponse
func (a *AccountManagementAPIService) EnableOptionsForSubAccountExecute(r ApiEnableOptionsForSubAccountRequest) (*common.RestApiResponse[models.EnableOptionsForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/eoptions/enable"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EnableOptionsForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesPositionRiskOfSubAccountRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiGetFuturesPositionRiskOfSubAccountRequest) Email(email string) ApiGetFuturesPositionRiskOfSubAccountRequest {
	r.email = &email
	return r
}

func (r ApiGetFuturesPositionRiskOfSubAccountRequest) RecvWindow(recvWindow int64) ApiGetFuturesPositionRiskOfSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesPositionRiskOfSubAccountRequest) Execute() (*common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountResponse], error) {
	return r.ApiService.GetFuturesPositionRiskOfSubAccountExecute(r)
}

/*
GetFuturesPositionRiskOfSubAccount Get Futures Position-Risk of Sub-account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/futures/positionRisk

https://developers.binance.com/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiGetFuturesPositionRiskOfSubAccountRequest
*/
func (a *AccountManagementAPIService) GetFuturesPositionRiskOfSubAccount(ctx context.Context) ApiGetFuturesPositionRiskOfSubAccountRequest {
	return ApiGetFuturesPositionRiskOfSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesPositionRiskOfSubAccountResponse
func (a *AccountManagementAPIService) GetFuturesPositionRiskOfSubAccountExecute(r ApiGetFuturesPositionRiskOfSubAccountRequest) (*common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesPositionRiskOfSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesPositionRiskOfSubAccountV2Request struct {
	ctx         context.Context
	ApiService  *AccountManagementAPIService
	email       *string
	futuresType *int64
	recvWindow  *int64
}

// [Sub-account email](#email-address)
func (r ApiGetFuturesPositionRiskOfSubAccountV2Request) Email(email string) ApiGetFuturesPositionRiskOfSubAccountV2Request {
	r.email = &email
	return r
}

// 1:USDT-margined Futures，2: Coin-margined Futures
func (r ApiGetFuturesPositionRiskOfSubAccountV2Request) FuturesType(futuresType int64) ApiGetFuturesPositionRiskOfSubAccountV2Request {
	r.futuresType = &futuresType
	return r
}

func (r ApiGetFuturesPositionRiskOfSubAccountV2Request) RecvWindow(recvWindow int64) ApiGetFuturesPositionRiskOfSubAccountV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesPositionRiskOfSubAccountV2Request) Execute() (*common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountV2Response], error) {
	return r.ApiService.GetFuturesPositionRiskOfSubAccountV2Execute(r)
}

/*
GetFuturesPositionRiskOfSubAccountV2 Get Futures Position-Risk of Sub-account V2 (For Master Account) (USER_DATA)
Get /sapi/v2/sub-account/futures/positionRisk

https://developers.binance.com/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param futuresType -  1:USDT-margined Futures，2: Coin-margined Futures
@param recvWindow -
@return ApiGetFuturesPositionRiskOfSubAccountV2Request
*/
func (a *AccountManagementAPIService) GetFuturesPositionRiskOfSubAccountV2(ctx context.Context) ApiGetFuturesPositionRiskOfSubAccountV2Request {
	return ApiGetFuturesPositionRiskOfSubAccountV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesPositionRiskOfSubAccountV2Response
func (a *AccountManagementAPIService) GetFuturesPositionRiskOfSubAccountV2Execute(r ApiGetFuturesPositionRiskOfSubAccountV2Request) (*common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/sub-account/futures/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.futuresType == nil {
		return nil, common.ReportError("futuresType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "futuresType", r.futuresType, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesPositionRiskOfSubAccountV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSubAccountsStatusOnMarginOrFuturesRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	recvWindow *int64
}

// Managed sub-account email
func (r ApiGetSubAccountsStatusOnMarginOrFuturesRequest) Email(email string) ApiGetSubAccountsStatusOnMarginOrFuturesRequest {
	r.email = &email
	return r
}

func (r ApiGetSubAccountsStatusOnMarginOrFuturesRequest) RecvWindow(recvWindow int64) ApiGetSubAccountsStatusOnMarginOrFuturesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSubAccountsStatusOnMarginOrFuturesRequest) Execute() (*common.RestApiResponse[models.GetSubAccountsStatusOnMarginOrFuturesResponse], error) {
	return r.ApiService.GetSubAccountsStatusOnMarginOrFuturesExecute(r)
}

/*
GetSubAccountsStatusOnMarginOrFutures Get Sub-account's Status on Margin Or Futures (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/status

https://developers.binance.com/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Managed sub-account email
@param recvWindow -
@return ApiGetSubAccountsStatusOnMarginOrFuturesRequest
*/
func (a *AccountManagementAPIService) GetSubAccountsStatusOnMarginOrFutures(ctx context.Context) ApiGetSubAccountsStatusOnMarginOrFuturesRequest {
	return ApiGetSubAccountsStatusOnMarginOrFuturesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSubAccountsStatusOnMarginOrFuturesResponse
func (a *AccountManagementAPIService) GetSubAccountsStatusOnMarginOrFuturesExecute(r ApiGetSubAccountsStatusOnMarginOrFuturesRequest) (*common.RestApiResponse[models.GetSubAccountsStatusOnMarginOrFuturesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSubAccountsStatusOnMarginOrFuturesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountListRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	isFreeze   *string
	page       *int64
	limit      *int64
	recvWindow *int64
}

// Managed sub-account email
func (r ApiQuerySubAccountListRequest) Email(email string) ApiQuerySubAccountListRequest {
	r.email = &email
	return r
}

// true or false
func (r ApiQuerySubAccountListRequest) IsFreeze(isFreeze string) ApiQuerySubAccountListRequest {
	r.isFreeze = &isFreeze
	return r
}

// Default value: 1
func (r ApiQuerySubAccountListRequest) Page(page int64) ApiQuerySubAccountListRequest {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiQuerySubAccountListRequest) Limit(limit int64) ApiQuerySubAccountListRequest {
	r.limit = &limit
	return r
}

func (r ApiQuerySubAccountListRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountListRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountListResponse], error) {
	return r.ApiService.QuerySubAccountListExecute(r)
}

/*
QuerySubAccountList Query Sub-account List (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/list

https://developers.binance.com/docs/sub_account/account-management/Query-Sub-account-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Managed sub-account email
@param isFreeze -  true or false
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQuerySubAccountListRequest
*/
func (a *AccountManagementAPIService) QuerySubAccountList(ctx context.Context) ApiQuerySubAccountListRequest {
	return ApiQuerySubAccountListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountListResponse
func (a *AccountManagementAPIService) QuerySubAccountListExecute(r ApiQuerySubAccountListRequest) (*common.RestApiResponse[models.QuerySubAccountListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	}
	if r.isFreeze != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isFreeze", r.isFreeze, "form", "")
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

	resp, err := SendRequest[models.QuerySubAccountListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountTransactionStatisticsRequest struct {
	ctx        context.Context
	ApiService *AccountManagementAPIService
	email      *string
	recvWindow *int64
}

// Managed sub-account email
func (r ApiQuerySubAccountTransactionStatisticsRequest) Email(email string) ApiQuerySubAccountTransactionStatisticsRequest {
	r.email = &email
	return r
}

func (r ApiQuerySubAccountTransactionStatisticsRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountTransactionStatisticsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountTransactionStatisticsRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountTransactionStatisticsResponse], error) {
	return r.ApiService.QuerySubAccountTransactionStatisticsExecute(r)
}

/*
QuerySubAccountTransactionStatistics Query Sub-account Transaction Statistics (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/transaction-statistics

https://developers.binance.com/docs/sub_account/account-management/Query-Sub-account-Transaction-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Managed sub-account email
@param recvWindow -
@return ApiQuerySubAccountTransactionStatisticsRequest
*/
func (a *AccountManagementAPIService) QuerySubAccountTransactionStatistics(ctx context.Context) ApiQuerySubAccountTransactionStatisticsRequest {
	return ApiQuerySubAccountTransactionStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountTransactionStatisticsResponse
func (a *AccountManagementAPIService) QuerySubAccountTransactionStatisticsExecute(r ApiQuerySubAccountTransactionStatisticsRequest) (*common.RestApiResponse[models.QuerySubAccountTransactionStatisticsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/transaction-statistics"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubAccountTransactionStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
