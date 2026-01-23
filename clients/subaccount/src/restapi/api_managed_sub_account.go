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

// ManagedSubAccountAPIService ManagedSubAccountAPI Service
type ManagedSubAccountAPIService Service

type ApiDepositAssetsIntoTheManagedSubAccountRequest struct {
	ctx        context.Context
	ApiService *ManagedSubAccountAPIService
	toEmail    *string
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiDepositAssetsIntoTheManagedSubAccountRequest) ToEmail(toEmail string) ApiDepositAssetsIntoTheManagedSubAccountRequest {
	r.toEmail = &toEmail
	return r
}

func (r ApiDepositAssetsIntoTheManagedSubAccountRequest) Asset(asset string) ApiDepositAssetsIntoTheManagedSubAccountRequest {
	r.asset = &asset
	return r
}

func (r ApiDepositAssetsIntoTheManagedSubAccountRequest) Amount(amount float32) ApiDepositAssetsIntoTheManagedSubAccountRequest {
	r.amount = &amount
	return r
}

func (r ApiDepositAssetsIntoTheManagedSubAccountRequest) RecvWindow(recvWindow int64) ApiDepositAssetsIntoTheManagedSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDepositAssetsIntoTheManagedSubAccountRequest) Execute() (*common.RestApiResponse[models.DepositAssetsIntoTheManagedSubAccountResponse], error) {
	return r.ApiService.DepositAssetsIntoTheManagedSubAccountExecute(r)
}

/*
DepositAssetsIntoTheManagedSubAccount Deposit Assets Into The Managed Sub-account (For Investor Master Account) (USER_DATA)
Post /sapi/v1/managed-subaccount/deposit

https://developers.binance.com/docs/sub_account/managed-sub-account/Deposit-Assets-Into-The-Managed-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param toEmail -
@param asset -
@param amount -
@param recvWindow -
@return ApiDepositAssetsIntoTheManagedSubAccountRequest
*/
func (a *ManagedSubAccountAPIService) DepositAssetsIntoTheManagedSubAccount(ctx context.Context) ApiDepositAssetsIntoTheManagedSubAccountRequest {
	return ApiDepositAssetsIntoTheManagedSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositAssetsIntoTheManagedSubAccountResponse
func (a *ManagedSubAccountAPIService) DepositAssetsIntoTheManagedSubAccountExecute(r ApiDepositAssetsIntoTheManagedSubAccountRequest) (*common.RestApiResponse[models.DepositAssetsIntoTheManagedSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/deposit"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.toEmail == nil {
		return nil, common.ReportError("toEmail is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toEmail", r.toEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DepositAssetsIntoTheManagedSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetManagedSubAccountDepositAddressRequest struct {
	ctx        context.Context
	ApiService *ManagedSubAccountAPIService
	email      *string
	coin       *string
	network    *string
	amount     *float32
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiGetManagedSubAccountDepositAddressRequest) Email(email string) ApiGetManagedSubAccountDepositAddressRequest {
	r.email = &email
	return r
}

func (r ApiGetManagedSubAccountDepositAddressRequest) Coin(coin string) ApiGetManagedSubAccountDepositAddressRequest {
	r.coin = &coin
	return r
}

// networks can be found in &#x60;GET /sapi/v1/capital/deposit/address&#x60;
func (r ApiGetManagedSubAccountDepositAddressRequest) Network(network string) ApiGetManagedSubAccountDepositAddressRequest {
	r.network = &network
	return r
}

func (r ApiGetManagedSubAccountDepositAddressRequest) Amount(amount float32) ApiGetManagedSubAccountDepositAddressRequest {
	r.amount = &amount
	return r
}

func (r ApiGetManagedSubAccountDepositAddressRequest) RecvWindow(recvWindow int64) ApiGetManagedSubAccountDepositAddressRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetManagedSubAccountDepositAddressRequest) Execute() (*common.RestApiResponse[models.GetManagedSubAccountDepositAddressResponse], error) {
	return r.ApiService.GetManagedSubAccountDepositAddressExecute(r)
}

/*
GetManagedSubAccountDepositAddress Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/deposit/address

https://developers.binance.com/docs/sub_account/managed-sub-account/Get-Managed-Sub-account-Deposit-Address

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param coin -
@param network -  networks can be found in `GET /sapi/v1/capital/deposit/address`
@param amount -
@param recvWindow -
@return ApiGetManagedSubAccountDepositAddressRequest
*/
func (a *ManagedSubAccountAPIService) GetManagedSubAccountDepositAddress(ctx context.Context) ApiGetManagedSubAccountDepositAddressRequest {
	return ApiGetManagedSubAccountDepositAddressRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetManagedSubAccountDepositAddressResponse
func (a *ManagedSubAccountAPIService) GetManagedSubAccountDepositAddressExecute(r ApiGetManagedSubAccountDepositAddressRequest) (*common.RestApiResponse[models.GetManagedSubAccountDepositAddressResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/deposit/address"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.amount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetManagedSubAccountDepositAddressResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountAssetDetailsRequest struct {
	ctx        context.Context
	ApiService *ManagedSubAccountAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountAssetDetailsRequest) Email(email string) ApiQueryManagedSubAccountAssetDetailsRequest {
	r.email = &email
	return r
}

func (r ApiQueryManagedSubAccountAssetDetailsRequest) RecvWindow(recvWindow int64) ApiQueryManagedSubAccountAssetDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryManagedSubAccountAssetDetailsRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountAssetDetailsResponse], error) {
	return r.ApiService.QueryManagedSubAccountAssetDetailsExecute(r)
}

/*
QueryManagedSubAccountAssetDetails Query Managed Sub-account Asset Details (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/asset

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Asset-Details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiQueryManagedSubAccountAssetDetailsRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountAssetDetails(ctx context.Context) ApiQueryManagedSubAccountAssetDetailsRequest {
	return ApiQueryManagedSubAccountAssetDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountAssetDetailsResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountAssetDetailsExecute(r ApiQueryManagedSubAccountAssetDetailsRequest) (*common.RestApiResponse[models.QueryManagedSubAccountAssetDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/asset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountAssetDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountFuturesAssetDetailsRequest struct {
	ctx         context.Context
	ApiService  *ManagedSubAccountAPIService
	email       *string
	accountType *string
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountFuturesAssetDetailsRequest) Email(email string) ApiQueryManagedSubAccountFuturesAssetDetailsRequest {
	r.email = &email
	return r
}

// No input or input \&quot;MARGIN\&quot; to get Cross Margin account details. Input \&quot;ISOLATED_MARGIN\&quot; to get Isolated Margin account details.
func (r ApiQueryManagedSubAccountFuturesAssetDetailsRequest) AccountType(accountType string) ApiQueryManagedSubAccountFuturesAssetDetailsRequest {
	r.accountType = &accountType
	return r
}

func (r ApiQueryManagedSubAccountFuturesAssetDetailsRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountFuturesAssetDetailsResponse], error) {
	return r.ApiService.QueryManagedSubAccountFuturesAssetDetailsExecute(r)
}

/*
QueryManagedSubAccountFuturesAssetDetails Query Managed Sub-account Futures Asset Details (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/fetch-future-asset

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Futures-Asset-Details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param accountType -  No input or input \"MARGIN\" to get Cross Margin account details. Input \"ISOLATED_MARGIN\" to get Isolated Margin account details.
@return ApiQueryManagedSubAccountFuturesAssetDetailsRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountFuturesAssetDetails(ctx context.Context) ApiQueryManagedSubAccountFuturesAssetDetailsRequest {
	return ApiQueryManagedSubAccountFuturesAssetDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountFuturesAssetDetailsResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountFuturesAssetDetailsExecute(r ApiQueryManagedSubAccountFuturesAssetDetailsRequest) (*common.RestApiResponse[models.QueryManagedSubAccountFuturesAssetDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/fetch-future-asset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.accountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountFuturesAssetDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountListRequest struct {
	ctx        context.Context
	ApiService *ManagedSubAccountAPIService
	email      *string
	page       *int64
	limit      *int64
	recvWindow *int64
}

// Managed sub-account email
func (r ApiQueryManagedSubAccountListRequest) Email(email string) ApiQueryManagedSubAccountListRequest {
	r.email = &email
	return r
}

// Default value: 1
func (r ApiQueryManagedSubAccountListRequest) Page(page int64) ApiQueryManagedSubAccountListRequest {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiQueryManagedSubAccountListRequest) Limit(limit int64) ApiQueryManagedSubAccountListRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryManagedSubAccountListRequest) RecvWindow(recvWindow int64) ApiQueryManagedSubAccountListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryManagedSubAccountListRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountListResponse], error) {
	return r.ApiService.QueryManagedSubAccountListExecute(r)
}

/*
QueryManagedSubAccountList Query Managed Sub-account List (For Investor) (USER_DATA)
Get /sapi/v1/managed-subaccount/info

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Managed sub-account email
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQueryManagedSubAccountListRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountList(ctx context.Context) ApiQueryManagedSubAccountListRequest {
	return ApiQueryManagedSubAccountListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountListResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountListExecute(r ApiQueryManagedSubAccountListRequest) (*common.RestApiResponse[models.QueryManagedSubAccountListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/info"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
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

	resp, err := SendRequest[models.QueryManagedSubAccountListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountMarginAssetDetailsRequest struct {
	ctx         context.Context
	ApiService  *ManagedSubAccountAPIService
	email       *string
	accountType *string
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountMarginAssetDetailsRequest) Email(email string) ApiQueryManagedSubAccountMarginAssetDetailsRequest {
	r.email = &email
	return r
}

// No input or input \&quot;MARGIN\&quot; to get Cross Margin account details. Input \&quot;ISOLATED_MARGIN\&quot; to get Isolated Margin account details.
func (r ApiQueryManagedSubAccountMarginAssetDetailsRequest) AccountType(accountType string) ApiQueryManagedSubAccountMarginAssetDetailsRequest {
	r.accountType = &accountType
	return r
}

func (r ApiQueryManagedSubAccountMarginAssetDetailsRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountMarginAssetDetailsResponse], error) {
	return r.ApiService.QueryManagedSubAccountMarginAssetDetailsExecute(r)
}

/*
QueryManagedSubAccountMarginAssetDetails Query Managed Sub-account Margin Asset Details (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/marginAsset

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Margin-Asset-Details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param accountType -  No input or input \"MARGIN\" to get Cross Margin account details. Input \"ISOLATED_MARGIN\" to get Isolated Margin account details.
@return ApiQueryManagedSubAccountMarginAssetDetailsRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountMarginAssetDetails(ctx context.Context) ApiQueryManagedSubAccountMarginAssetDetailsRequest {
	return ApiQueryManagedSubAccountMarginAssetDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountMarginAssetDetailsResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountMarginAssetDetailsExecute(r ApiQueryManagedSubAccountMarginAssetDetailsRequest) (*common.RestApiResponse[models.QueryManagedSubAccountMarginAssetDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/marginAsset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.accountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountMarginAssetDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountSnapshotRequest struct {
	ctx        context.Context
	ApiService *ManagedSubAccountAPIService
	email      *string
	type_      *string
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountSnapshotRequest) Email(email string) ApiQueryManagedSubAccountSnapshotRequest {
	r.email = &email
	return r
}

// \&quot;SPOT\&quot;, \&quot;MARGIN\&quot;（cross）, \&quot;FUTURES\&quot;（UM）
func (r ApiQueryManagedSubAccountSnapshotRequest) Type(type_ string) ApiQueryManagedSubAccountSnapshotRequest {
	r.type_ = &type_
	return r
}

func (r ApiQueryManagedSubAccountSnapshotRequest) StartTime(startTime int64) ApiQueryManagedSubAccountSnapshotRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryManagedSubAccountSnapshotRequest) EndTime(endTime int64) ApiQueryManagedSubAccountSnapshotRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1, Max value: 200
func (r ApiQueryManagedSubAccountSnapshotRequest) Limit(limit int64) ApiQueryManagedSubAccountSnapshotRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryManagedSubAccountSnapshotRequest) RecvWindow(recvWindow int64) ApiQueryManagedSubAccountSnapshotRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryManagedSubAccountSnapshotRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountSnapshotResponse], error) {
	return r.ApiService.QueryManagedSubAccountSnapshotExecute(r)
}

/*
QueryManagedSubAccountSnapshot Query Managed Sub-account Snapshot (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/accountSnapshot

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-account-Snapshot

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param type_ -  \"SPOT\", \"MARGIN\"（cross）, \"FUTURES\"（UM）
@param startTime -
@param endTime -
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQueryManagedSubAccountSnapshotRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountSnapshot(ctx context.Context) ApiQueryManagedSubAccountSnapshotRequest {
	return ApiQueryManagedSubAccountSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountSnapshotResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountSnapshotExecute(r ApiQueryManagedSubAccountSnapshotRequest) (*common.RestApiResponse[models.QueryManagedSubAccountSnapshotResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/accountSnapshot"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
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

	resp, err := SendRequest[models.QueryManagedSubAccountSnapshotResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest struct {
	ctx                         context.Context
	ApiService                  *ManagedSubAccountAPIService
	email                       *string
	startTime                   *int64
	endTime                     *int64
	page                        *int64
	limit                       *int64
	transfers                   *string
	transferFunctionAccountType *string
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) Email(email string) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.email = &email
	return r
}

// Start Time
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) StartTime(startTime int64) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.startTime = &startTime
	return r
}

// End Time (The start time and end time interval cannot exceed half a year)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) EndTime(endTime int64) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.endTime = &endTime
	return r
}

// Page
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) Page(page int64) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.page = &page
	return r
}

// Limit (Max: 500)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) Limit(limit int64) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.limit = &limit
	return r
}

// Transfer Direction (FROM/TO)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) Transfers(transfers string) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.transfers = &transfers
	return r
}

// Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) TransferFunctionAccountType(transferFunctionAccountType string) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	r.transferFunctionAccountType = &transferFunctionAccountType
	return r
}

func (r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse], error) {
	return r.ApiService.QueryManagedSubAccountTransferLogMasterAccountInvestorExecute(r)
}

/*
QueryManagedSubAccountTransferLogMasterAccountInvestor Query Managed Sub Account Transfer Log (For Investor Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/queryTransLogForInvestor

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-Account-Transfer-Log-Investor

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param startTime -  Start Time
@param endTime -  End Time (The start time and end time interval cannot exceed half a year)
@param page -  Page
@param limit -  Limit (Max: 500)
@param transfers -  Transfer Direction (FROM/TO)
@param transferFunctionAccountType -  Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
@return ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogMasterAccountInvestor(ctx context.Context) ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest {
	return ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountTransferLogMasterAccountInvestorResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogMasterAccountInvestorExecute(r ApiQueryManagedSubAccountTransferLogMasterAccountInvestorRequest) (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/queryTransLogForInvestor"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}
	if r.page == nil {
		return nil, common.ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return nil, common.ReportError("limit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.transfers != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transfers", r.transfers, "form", "")
	}
	if r.transferFunctionAccountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferFunctionAccountType", r.transferFunctionAccountType, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest struct {
	ctx                         context.Context
	ApiService                  *ManagedSubAccountAPIService
	email                       *string
	startTime                   *int64
	endTime                     *int64
	page                        *int64
	limit                       *int64
	transfers                   *string
	transferFunctionAccountType *string
}

// [Sub-account email](#email-address)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) Email(email string) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.email = &email
	return r
}

// Start Time
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) StartTime(startTime int64) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.startTime = &startTime
	return r
}

// End Time (The start time and end time interval cannot exceed half a year)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) EndTime(endTime int64) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.endTime = &endTime
	return r
}

// Page
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) Page(page int64) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.page = &page
	return r
}

// Limit (Max: 500)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) Limit(limit int64) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.limit = &limit
	return r
}

// Transfer Direction (FROM/TO)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) Transfers(transfers string) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.transfers = &transfers
	return r
}

// Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) TransferFunctionAccountType(transferFunctionAccountType string) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	r.transferFunctionAccountType = &transferFunctionAccountType
	return r
}

func (r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse], error) {
	return r.ApiService.QueryManagedSubAccountTransferLogMasterAccountTradingExecute(r)
}

/*
QueryManagedSubAccountTransferLogMasterAccountTrading Query Managed Sub Account Transfer Log (For Trading Team Master Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/queryTransLogForTradeParent

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-Account-Transfer-Log-Trading-Team-Master

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param startTime -  Start Time
@param endTime -  End Time (The start time and end time interval cannot exceed half a year)
@param page -  Page
@param limit -  Limit (Max: 500)
@param transfers -  Transfer Direction (FROM/TO)
@param transferFunctionAccountType -  Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
@return ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogMasterAccountTrading(ctx context.Context) ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest {
	return ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountTransferLogMasterAccountTradingResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogMasterAccountTradingExecute(r ApiQueryManagedSubAccountTransferLogMasterAccountTradingRequest) (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/queryTransLogForTradeParent"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}
	if r.page == nil {
		return nil, common.ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return nil, common.ReportError("limit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.transfers != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transfers", r.transfers, "form", "")
	}
	if r.transferFunctionAccountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferFunctionAccountType", r.transferFunctionAccountType, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest struct {
	ctx                         context.Context
	ApiService                  *ManagedSubAccountAPIService
	startTime                   *int64
	endTime                     *int64
	page                        *int64
	limit                       *int64
	transfers                   *string
	transferFunctionAccountType *string
	recvWindow                  *int64
}

// Start Time
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) StartTime(startTime int64) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.startTime = &startTime
	return r
}

// End Time (The start time and end time interval cannot exceed half a year)
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) EndTime(endTime int64) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.endTime = &endTime
	return r
}

// Page
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) Page(page int64) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.page = &page
	return r
}

// Limit (Max: 500)
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) Limit(limit int64) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.limit = &limit
	return r
}

// Transfer Direction (FROM/TO)
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) Transfers(transfers string) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.transfers = &transfers
	return r
}

// Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) TransferFunctionAccountType(transferFunctionAccountType string) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.transferFunctionAccountType = &transferFunctionAccountType
	return r
}

func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) RecvWindow(recvWindow int64) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) Execute() (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogSubAccountTradingResponse], error) {
	return r.ApiService.QueryManagedSubAccountTransferLogSubAccountTradingExecute(r)
}

/*
QueryManagedSubAccountTransferLogSubAccountTrading Query Managed Sub Account Transfer Log (For Trading Team Sub Account) (USER_DATA)
Get /sapi/v1/managed-subaccount/query-trans-log

https://developers.binance.com/docs/sub_account/managed-sub-account/Query-Managed-Sub-Account-Transfer-Log-Trading-Team-Sub

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start Time
@param endTime -  End Time (The start time and end time interval cannot exceed half a year)
@param page -  Page
@param limit -  Limit (Max: 500)
@param transfers -  Transfer Direction (FROM/TO)
@param transferFunctionAccountType -  Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE)
@param recvWindow -
@return ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest
*/
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogSubAccountTrading(ctx context.Context) ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest {
	return ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryManagedSubAccountTransferLogSubAccountTradingResponse
func (a *ManagedSubAccountAPIService) QueryManagedSubAccountTransferLogSubAccountTradingExecute(r ApiQueryManagedSubAccountTransferLogSubAccountTradingRequest) (*common.RestApiResponse[models.QueryManagedSubAccountTransferLogSubAccountTradingResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/query-trans-log"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}
	if r.page == nil {
		return nil, common.ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return nil, common.ReportError("limit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.transfers != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transfers", r.transfers, "form", "")
	}
	if r.transferFunctionAccountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferFunctionAccountType", r.transferFunctionAccountType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryManagedSubAccountTransferLogSubAccountTradingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawlAssetsFromTheManagedSubAccountRequest struct {
	ctx          context.Context
	ApiService   *ManagedSubAccountAPIService
	fromEmail    *string
	asset        *string
	amount       *float32
	transferDate *int64
	recvWindow   *int64
}

func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) FromEmail(fromEmail string) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	r.fromEmail = &fromEmail
	return r
}

func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) Asset(asset string) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	r.asset = &asset
	return r
}

func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) Amount(amount float32) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	r.amount = &amount
	return r
}

// Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now
func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) TransferDate(transferDate int64) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	r.transferDate = &transferDate
	return r
}

func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) RecvWindow(recvWindow int64) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) Execute() (*common.RestApiResponse[models.WithdrawlAssetsFromTheManagedSubAccountResponse], error) {
	return r.ApiService.WithdrawlAssetsFromTheManagedSubAccountExecute(r)
}

/*
WithdrawlAssetsFromTheManagedSubAccount Withdrawl Assets From The Managed Sub-account (For Investor Master Account) (USER_DATA)
Post /sapi/v1/managed-subaccount/withdraw

https://developers.binance.com/docs/sub_account/managed-sub-account/Withdrawl-Assets-From-The-Managed-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromEmail -
@param asset -
@param amount -
@param transferDate -  Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now
@param recvWindow -
@return ApiWithdrawlAssetsFromTheManagedSubAccountRequest
*/
func (a *ManagedSubAccountAPIService) WithdrawlAssetsFromTheManagedSubAccount(ctx context.Context) ApiWithdrawlAssetsFromTheManagedSubAccountRequest {
	return ApiWithdrawlAssetsFromTheManagedSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawlAssetsFromTheManagedSubAccountResponse
func (a *ManagedSubAccountAPIService) WithdrawlAssetsFromTheManagedSubAccountExecute(r ApiWithdrawlAssetsFromTheManagedSubAccountRequest) (*common.RestApiResponse[models.WithdrawlAssetsFromTheManagedSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/managed-subaccount/withdraw"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromEmail == nil {
		return nil, common.ReportError("fromEmail is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromEmail", r.fromEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.transferDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferDate", r.transferDate, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WithdrawlAssetsFromTheManagedSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
