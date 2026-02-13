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
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AssetManagementAPIService AssetManagementAPI Service
type AssetManagementAPIService Service

type ApiFuturesTransferForSubAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	asset      *string
	amount     *float32
	type_      *int64
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiFuturesTransferForSubAccountRequest) Email(email string) ApiFuturesTransferForSubAccountRequest {
	r.email = &email
	return r
}

func (r ApiFuturesTransferForSubAccountRequest) Asset(asset string) ApiFuturesTransferForSubAccountRequest {
	r.asset = &asset
	return r
}

func (r ApiFuturesTransferForSubAccountRequest) Amount(amount float32) ApiFuturesTransferForSubAccountRequest {
	r.amount = &amount
	return r
}

// 1: transfer from subaccount&#39;s  spot account to margin account 2: transfer from subaccount&#39;s margin account to its spot account
func (r ApiFuturesTransferForSubAccountRequest) Type(type_ int64) ApiFuturesTransferForSubAccountRequest {
	r.type_ = &type_
	return r
}

func (r ApiFuturesTransferForSubAccountRequest) RecvWindow(recvWindow int64) ApiFuturesTransferForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesTransferForSubAccountRequest) Execute() (*common.RestApiResponse[models.FuturesTransferForSubAccountResponse], error) {
	return r.ApiService.FuturesTransferForSubAccountExecute(r)
}

/*
FuturesTransferForSubAccount Futures Transfer for Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/futures/transfer

https://developers.binance.com/docs/sub_account/asset-management/Futures-Transfer-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param asset -
@param amount -
@param type_ -  1: transfer from subaccount's  spot account to margin account 2: transfer from subaccount's margin account to its spot account
@param recvWindow -
@return ApiFuturesTransferForSubAccountRequest
*/
func (a *AssetManagementAPIService) FuturesTransferForSubAccount(ctx context.Context) ApiFuturesTransferForSubAccountRequest {
	return ApiFuturesTransferForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FuturesTransferForSubAccountResponse
func (a *AssetManagementAPIService) FuturesTransferForSubAccountExecute(r ApiFuturesTransferForSubAccountRequest) (*common.RestApiResponse[models.FuturesTransferForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FuturesTransferForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDetailOnSubAccountsFuturesAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiGetDetailOnSubAccountsFuturesAccountRequest) Email(email string) ApiGetDetailOnSubAccountsFuturesAccountRequest {
	r.email = &email
	return r
}

func (r ApiGetDetailOnSubAccountsFuturesAccountRequest) RecvWindow(recvWindow int64) ApiGetDetailOnSubAccountsFuturesAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDetailOnSubAccountsFuturesAccountRequest) Execute() (*common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountResponse], error) {
	return r.ApiService.GetDetailOnSubAccountsFuturesAccountExecute(r)
}

/*
GetDetailOnSubAccountsFuturesAccount Get Detail on Sub-account's Futures Account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/futures/account

https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiGetDetailOnSubAccountsFuturesAccountRequest
*/
func (a *AssetManagementAPIService) GetDetailOnSubAccountsFuturesAccount(ctx context.Context) ApiGetDetailOnSubAccountsFuturesAccountRequest {
	return ApiGetDetailOnSubAccountsFuturesAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDetailOnSubAccountsFuturesAccountResponse
func (a *AssetManagementAPIService) GetDetailOnSubAccountsFuturesAccountExecute(r ApiGetDetailOnSubAccountsFuturesAccountRequest) (*common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDetailOnSubAccountsFuturesAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDetailOnSubAccountsFuturesAccountV2Request struct {
	ctx         context.Context
	ApiService  *AssetManagementAPIService
	email       *string
	futuresType *int64
	recvWindow  *int64
}

// [Sub-account email](#email-address)
func (r ApiGetDetailOnSubAccountsFuturesAccountV2Request) Email(email string) ApiGetDetailOnSubAccountsFuturesAccountV2Request {
	r.email = &email
	return r
}

// 1:USDT-margined Futures，2: Coin-margined Futures
func (r ApiGetDetailOnSubAccountsFuturesAccountV2Request) FuturesType(futuresType int64) ApiGetDetailOnSubAccountsFuturesAccountV2Request {
	r.futuresType = &futuresType
	return r
}

func (r ApiGetDetailOnSubAccountsFuturesAccountV2Request) RecvWindow(recvWindow int64) ApiGetDetailOnSubAccountsFuturesAccountV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDetailOnSubAccountsFuturesAccountV2Request) Execute() (*common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountV2Response], error) {
	return r.ApiService.GetDetailOnSubAccountsFuturesAccountV2Execute(r)
}

/*
GetDetailOnSubAccountsFuturesAccountV2 Get Detail on Sub-account's Futures Account V2 (For Master Account) (USER_DATA)
Get /sapi/v2/sub-account/futures/account

https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param futuresType -  1:USDT-margined Futures，2: Coin-margined Futures
@param recvWindow -
@return ApiGetDetailOnSubAccountsFuturesAccountV2Request
*/
func (a *AssetManagementAPIService) GetDetailOnSubAccountsFuturesAccountV2(ctx context.Context) ApiGetDetailOnSubAccountsFuturesAccountV2Request {
	return ApiGetDetailOnSubAccountsFuturesAccountV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDetailOnSubAccountsFuturesAccountV2Response
func (a *AssetManagementAPIService) GetDetailOnSubAccountsFuturesAccountV2Execute(r ApiGetDetailOnSubAccountsFuturesAccountV2Request) (*common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/sub-account/futures/account"

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

	resp, err := SendRequest[models.GetDetailOnSubAccountsFuturesAccountV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDetailOnSubAccountsMarginAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiGetDetailOnSubAccountsMarginAccountRequest) Email(email string) ApiGetDetailOnSubAccountsMarginAccountRequest {
	r.email = &email
	return r
}

func (r ApiGetDetailOnSubAccountsMarginAccountRequest) RecvWindow(recvWindow int64) ApiGetDetailOnSubAccountsMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDetailOnSubAccountsMarginAccountRequest) Execute() (*common.RestApiResponse[models.GetDetailOnSubAccountsMarginAccountResponse], error) {
	return r.ApiService.GetDetailOnSubAccountsMarginAccountExecute(r)
}

/*
GetDetailOnSubAccountsMarginAccount Get Detail on Sub-account's Margin Account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/margin/account

https://developers.binance.com/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiGetDetailOnSubAccountsMarginAccountRequest
*/
func (a *AssetManagementAPIService) GetDetailOnSubAccountsMarginAccount(ctx context.Context) ApiGetDetailOnSubAccountsMarginAccountRequest {
	return ApiGetDetailOnSubAccountsMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDetailOnSubAccountsMarginAccountResponse
func (a *AssetManagementAPIService) GetDetailOnSubAccountsMarginAccountExecute(r ApiGetDetailOnSubAccountsMarginAccountRequest) (*common.RestApiResponse[models.GetDetailOnSubAccountsMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/margin/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDetailOnSubAccountsMarginAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetMovePositionHistoryForSubAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	symbol     *string
	page       *int64
	row        *int64
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) Symbol(symbol string) ApiGetMovePositionHistoryForSubAccountRequest {
	r.symbol = &symbol
	return r
}

// Page
func (r ApiGetMovePositionHistoryForSubAccountRequest) Page(page int64) ApiGetMovePositionHistoryForSubAccountRequest {
	r.page = &page
	return r
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) Row(row int64) ApiGetMovePositionHistoryForSubAccountRequest {
	r.row = &row
	return r
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) StartTime(startTime int64) ApiGetMovePositionHistoryForSubAccountRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) EndTime(endTime int64) ApiGetMovePositionHistoryForSubAccountRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) RecvWindow(recvWindow int64) ApiGetMovePositionHistoryForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetMovePositionHistoryForSubAccountRequest) Execute() (*common.RestApiResponse[models.GetMovePositionHistoryForSubAccountResponse], error) {
	return r.ApiService.GetMovePositionHistoryForSubAccountExecute(r)
}

/*
GetMovePositionHistoryForSubAccount Get Move Position History for Sub-account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/futures/move-position

https://developers.binance.com/docs/sub_account/asset-management/Get-Move-Position-History-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param page -  Page
@param row -
@param startTime -
@param endTime -
@param recvWindow -
@return ApiGetMovePositionHistoryForSubAccountRequest
*/
func (a *AssetManagementAPIService) GetMovePositionHistoryForSubAccount(ctx context.Context) ApiGetMovePositionHistoryForSubAccountRequest {
	return ApiGetMovePositionHistoryForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMovePositionHistoryForSubAccountResponse
func (a *AssetManagementAPIService) GetMovePositionHistoryForSubAccountExecute(r ApiGetMovePositionHistoryForSubAccountRequest) (*common.RestApiResponse[models.GetMovePositionHistoryForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/move-position"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.page == nil {
		return nil, common.ReportError("page is required and must be specified")
	}
	if r.row == nil {
		return nil, common.ReportError("row is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "row", r.row, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetMovePositionHistoryForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSubAccountDepositAddressRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	coin       *string
	network    *string
	amount     *float32
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiGetSubAccountDepositAddressRequest) Email(email string) ApiGetSubAccountDepositAddressRequest {
	r.email = &email
	return r
}

func (r ApiGetSubAccountDepositAddressRequest) Coin(coin string) ApiGetSubAccountDepositAddressRequest {
	r.coin = &coin
	return r
}

// networks can be found in &#x60;GET /sapi/v1/capital/deposit/address&#x60;
func (r ApiGetSubAccountDepositAddressRequest) Network(network string) ApiGetSubAccountDepositAddressRequest {
	r.network = &network
	return r
}

func (r ApiGetSubAccountDepositAddressRequest) Amount(amount float32) ApiGetSubAccountDepositAddressRequest {
	r.amount = &amount
	return r
}

func (r ApiGetSubAccountDepositAddressRequest) RecvWindow(recvWindow int64) ApiGetSubAccountDepositAddressRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSubAccountDepositAddressRequest) Execute() (*common.RestApiResponse[models.GetSubAccountDepositAddressResponse], error) {
	return r.ApiService.GetSubAccountDepositAddressExecute(r)
}

/*
GetSubAccountDepositAddress Get Sub-account Deposit Address (For Master Account) (USER_DATA)
Get /sapi/v1/capital/deposit/subAddress

https://developers.binance.com/docs/sub_account/asset-management/Get-Sub-account-Deposit-Address

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param coin -
@param network -  networks can be found in `GET /sapi/v1/capital/deposit/address`
@param amount -
@param recvWindow -
@return ApiGetSubAccountDepositAddressRequest
*/
func (a *AssetManagementAPIService) GetSubAccountDepositAddress(ctx context.Context) ApiGetSubAccountDepositAddressRequest {
	return ApiGetSubAccountDepositAddressRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSubAccountDepositAddressResponse
func (a *AssetManagementAPIService) GetSubAccountDepositAddressExecute(r ApiGetSubAccountDepositAddressRequest) (*common.RestApiResponse[models.GetSubAccountDepositAddressResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/subAddress"

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

	resp, err := SendRequest[models.GetSubAccountDepositAddressResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSubAccountDepositHistoryRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	coin       *string
	status     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	offset     *int64
	recvWindow *int64
	txId       *string
}

// [Sub-account email](#email-address)
func (r ApiGetSubAccountDepositHistoryRequest) Email(email string) ApiGetSubAccountDepositHistoryRequest {
	r.email = &email
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) Coin(coin string) ApiGetSubAccountDepositHistoryRequest {
	r.coin = &coin
	return r
}

// 0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
func (r ApiGetSubAccountDepositHistoryRequest) Status(status int64) ApiGetSubAccountDepositHistoryRequest {
	r.status = &status
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) StartTime(startTime int64) ApiGetSubAccountDepositHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) EndTime(endTime int64) ApiGetSubAccountDepositHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1, Max value: 200
func (r ApiGetSubAccountDepositHistoryRequest) Limit(limit int64) ApiGetSubAccountDepositHistoryRequest {
	r.limit = &limit
	return r
}

// default:0
func (r ApiGetSubAccountDepositHistoryRequest) Offset(offset int64) ApiGetSubAccountDepositHistoryRequest {
	r.offset = &offset
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) RecvWindow(recvWindow int64) ApiGetSubAccountDepositHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) TxId(txId string) ApiGetSubAccountDepositHistoryRequest {
	r.txId = &txId
	return r
}

func (r ApiGetSubAccountDepositHistoryRequest) Execute() (*common.RestApiResponse[models.GetSubAccountDepositHistoryResponse], error) {
	return r.ApiService.GetSubAccountDepositHistoryExecute(r)
}

/*
GetSubAccountDepositHistory Get Sub-account Deposit History (For Master Account) (USER_DATA)
Get /sapi/v1/capital/deposit/subHisrec

https://developers.binance.com/docs/sub_account/asset-management/Get-Sub-account-Deposit-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param coin -
@param status -  0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success)
@param startTime -
@param endTime -
@param limit -  Default value: 1, Max value: 200
@param offset -  default:0
@param recvWindow -
@param txId -
@return ApiGetSubAccountDepositHistoryRequest
*/
func (a *AssetManagementAPIService) GetSubAccountDepositHistory(ctx context.Context) ApiGetSubAccountDepositHistoryRequest {
	return ApiGetSubAccountDepositHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSubAccountDepositHistoryResponse
func (a *AssetManagementAPIService) GetSubAccountDepositHistoryExecute(r ApiGetSubAccountDepositHistoryRequest) (*common.RestApiResponse[models.GetSubAccountDepositHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/subHisrec"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.status != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
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
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}

	resp, err := SendRequest[models.GetSubAccountDepositHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSummaryOfSubAccountsFuturesAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	page       *int64
	limit      *int64
	recvWindow *int64
}

// Page
func (r ApiGetSummaryOfSubAccountsFuturesAccountRequest) Page(page int64) ApiGetSummaryOfSubAccountsFuturesAccountRequest {
	r.page = &page
	return r
}

// Limit (Max: 500)
func (r ApiGetSummaryOfSubAccountsFuturesAccountRequest) Limit(limit int64) ApiGetSummaryOfSubAccountsFuturesAccountRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSummaryOfSubAccountsFuturesAccountRequest) RecvWindow(recvWindow int64) ApiGetSummaryOfSubAccountsFuturesAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSummaryOfSubAccountsFuturesAccountRequest) Execute() (*common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountResponse], error) {
	return r.ApiService.GetSummaryOfSubAccountsFuturesAccountExecute(r)
}

/*
GetSummaryOfSubAccountsFuturesAccount Get Summary of Sub-account's Futures Account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/futures/accountSummary

https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Futures-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param page -  Page
@param limit -  Limit (Max: 500)
@param recvWindow -
@return ApiGetSummaryOfSubAccountsFuturesAccountRequest
*/
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsFuturesAccount(ctx context.Context) ApiGetSummaryOfSubAccountsFuturesAccountRequest {
	return ApiGetSummaryOfSubAccountsFuturesAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSummaryOfSubAccountsFuturesAccountResponse
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsFuturesAccountExecute(r ApiGetSummaryOfSubAccountsFuturesAccountRequest) (*common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/accountSummary"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.page == nil {
		return nil, common.ReportError("page is required and must be specified")
	}
	if r.limit == nil {
		return nil, common.ReportError("limit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSummaryOfSubAccountsFuturesAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSummaryOfSubAccountsFuturesAccountV2Request struct {
	ctx         context.Context
	ApiService  *AssetManagementAPIService
	futuresType *int64
	page        *int64
	limit       *int64
	recvWindow  *int64
}

// 1:USDT-margined Futures，2: Coin-margined Futures
func (r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) FuturesType(futuresType int64) ApiGetSummaryOfSubAccountsFuturesAccountV2Request {
	r.futuresType = &futuresType
	return r
}

// Default value: 1
func (r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) Page(page int64) ApiGetSummaryOfSubAccountsFuturesAccountV2Request {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) Limit(limit int64) ApiGetSummaryOfSubAccountsFuturesAccountV2Request {
	r.limit = &limit
	return r
}

func (r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) RecvWindow(recvWindow int64) ApiGetSummaryOfSubAccountsFuturesAccountV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) Execute() (*common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountV2Response], error) {
	return r.ApiService.GetSummaryOfSubAccountsFuturesAccountV2Execute(r)
}

/*
GetSummaryOfSubAccountsFuturesAccountV2 Get Summary of Sub-account's Futures Account V2 (For Master Account) (USER_DATA)
Get /sapi/v2/sub-account/futures/accountSummary

https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Futures-Account-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param futuresType -  1:USDT-margined Futures，2: Coin-margined Futures
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiGetSummaryOfSubAccountsFuturesAccountV2Request
*/
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsFuturesAccountV2(ctx context.Context) ApiGetSummaryOfSubAccountsFuturesAccountV2Request {
	return ApiGetSummaryOfSubAccountsFuturesAccountV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSummaryOfSubAccountsFuturesAccountV2Response
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsFuturesAccountV2Execute(r ApiGetSummaryOfSubAccountsFuturesAccountV2Request) (*common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/sub-account/futures/accountSummary"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.futuresType == nil {
		return nil, common.ReportError("futuresType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "futuresType", r.futuresType, "form", "")
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSummaryOfSubAccountsFuturesAccountV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSummaryOfSubAccountsMarginAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	recvWindow *int64
}

func (r ApiGetSummaryOfSubAccountsMarginAccountRequest) RecvWindow(recvWindow int64) ApiGetSummaryOfSubAccountsMarginAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSummaryOfSubAccountsMarginAccountRequest) Execute() (*common.RestApiResponse[models.GetSummaryOfSubAccountsMarginAccountResponse], error) {
	return r.ApiService.GetSummaryOfSubAccountsMarginAccountExecute(r)
}

/*
GetSummaryOfSubAccountsMarginAccount Get Summary of Sub-account's Margin Account (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/margin/accountSummary

https://developers.binance.com/docs/sub_account/asset-management/Get-Summary-of-Sub-accounts-Margin-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetSummaryOfSubAccountsMarginAccountRequest
*/
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsMarginAccount(ctx context.Context) ApiGetSummaryOfSubAccountsMarginAccountRequest {
	return ApiGetSummaryOfSubAccountsMarginAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSummaryOfSubAccountsMarginAccountResponse
func (a *AssetManagementAPIService) GetSummaryOfSubAccountsMarginAccountExecute(r ApiGetSummaryOfSubAccountsMarginAccountRequest) (*common.RestApiResponse[models.GetSummaryOfSubAccountsMarginAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/margin/accountSummary"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSummaryOfSubAccountsMarginAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginTransferForSubAccountRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	asset      *string
	amount     *float32
	type_      *int64
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiMarginTransferForSubAccountRequest) Email(email string) ApiMarginTransferForSubAccountRequest {
	r.email = &email
	return r
}

func (r ApiMarginTransferForSubAccountRequest) Asset(asset string) ApiMarginTransferForSubAccountRequest {
	r.asset = &asset
	return r
}

func (r ApiMarginTransferForSubAccountRequest) Amount(amount float32) ApiMarginTransferForSubAccountRequest {
	r.amount = &amount
	return r
}

// 1: transfer from subaccount&#39;s  spot account to margin account 2: transfer from subaccount&#39;s margin account to its spot account
func (r ApiMarginTransferForSubAccountRequest) Type(type_ int64) ApiMarginTransferForSubAccountRequest {
	r.type_ = &type_
	return r
}

func (r ApiMarginTransferForSubAccountRequest) RecvWindow(recvWindow int64) ApiMarginTransferForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginTransferForSubAccountRequest) Execute() (*common.RestApiResponse[models.MarginTransferForSubAccountResponse], error) {
	return r.ApiService.MarginTransferForSubAccountExecute(r)
}

/*
MarginTransferForSubAccount Margin Transfer for Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/margin/transfer

https://developers.binance.com/docs/sub_account/asset-management/Margin-Transfer-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param asset -
@param amount -
@param type_ -  1: transfer from subaccount's  spot account to margin account 2: transfer from subaccount's margin account to its spot account
@param recvWindow -
@return ApiMarginTransferForSubAccountRequest
*/
func (a *AssetManagementAPIService) MarginTransferForSubAccount(ctx context.Context) ApiMarginTransferForSubAccountRequest {
	return ApiMarginTransferForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginTransferForSubAccountResponse
func (a *AssetManagementAPIService) MarginTransferForSubAccountExecute(r ApiMarginTransferForSubAccountRequest) (*common.RestApiResponse[models.MarginTransferForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/margin/transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginTransferForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMovePositionForSubAccountRequest struct {
	ctx           context.Context
	ApiService    *AssetManagementAPIService
	fromUserEmail *string
	toUserEmail   *string
	productType   *string
	orderArgs     *[]models.MovePositionForSubAccountOrderArgsParameterInner
	recvWindow    *int64
}

func (r ApiMovePositionForSubAccountRequest) FromUserEmail(fromUserEmail string) ApiMovePositionForSubAccountRequest {
	r.fromUserEmail = &fromUserEmail
	return r
}

func (r ApiMovePositionForSubAccountRequest) ToUserEmail(toUserEmail string) ApiMovePositionForSubAccountRequest {
	r.toUserEmail = &toUserEmail
	return r
}

// Only support UM
func (r ApiMovePositionForSubAccountRequest) ProductType(productType string) ApiMovePositionForSubAccountRequest {
	r.productType = &productType
	return r
}

// Max 10 positions supported. When input request parameter,orderArgs.symbol should be STRING, orderArgs.quantity should be BIGDECIMAL, and orderArgs.positionSide should be STRING, positionSide support BOTH,LONG and SHORT. Each entry should be like orderArgs[0].symbol&#x3D;BTCUSDT,orderArgs[0].quantity&#x3D;0.001,orderArgs[0].positionSide&#x3D;BOTH. Example of the request parameter array: orderArgs[0].symbol&#x3D;BTCUSDT orderArgs[0].quantity&#x3D;0.001 orderArgs[0].positionSide&#x3D;BOTH orderArgs[1].symbol&#x3D;ETHUSDT orderArgs[1].quantity&#x3D;0.01 orderArgs[1].positionSide&#x3D;BOTH
func (r ApiMovePositionForSubAccountRequest) OrderArgs(orderArgs []models.MovePositionForSubAccountOrderArgsParameterInner) ApiMovePositionForSubAccountRequest {
	r.orderArgs = &orderArgs
	return r
}

func (r ApiMovePositionForSubAccountRequest) RecvWindow(recvWindow int64) ApiMovePositionForSubAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMovePositionForSubAccountRequest) Execute() (*common.RestApiResponse[models.MovePositionForSubAccountResponse], error) {
	return r.ApiService.MovePositionForSubAccountExecute(r)
}

/*
MovePositionForSubAccount Move Position for Sub-account (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/futures/move-position

https://developers.binance.com/docs/sub_account/asset-management/Move-Position-for-Sub-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromUserEmail -
@param toUserEmail -
@param productType -  Only support UM
@param orderArgs -  Max 10 positions supported. When input request parameter,orderArgs.symbol should be STRING, orderArgs.quantity should be BIGDECIMAL, and orderArgs.positionSide should be STRING, positionSide support BOTH,LONG and SHORT. Each entry should be like orderArgs[0].symbol=BTCUSDT,orderArgs[0].quantity=0.001,orderArgs[0].positionSide=BOTH. Example of the request parameter array: orderArgs[0].symbol=BTCUSDT orderArgs[0].quantity=0.001 orderArgs[0].positionSide=BOTH orderArgs[1].symbol=ETHUSDT orderArgs[1].quantity=0.01 orderArgs[1].positionSide=BOTH
@param recvWindow -
@return ApiMovePositionForSubAccountRequest
*/
func (a *AssetManagementAPIService) MovePositionForSubAccount(ctx context.Context) ApiMovePositionForSubAccountRequest {
	return ApiMovePositionForSubAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MovePositionForSubAccountResponse
func (a *AssetManagementAPIService) MovePositionForSubAccountExecute(r ApiMovePositionForSubAccountRequest) (*common.RestApiResponse[models.MovePositionForSubAccountResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/move-position"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromUserEmail == nil {
		return nil, common.ReportError("fromUserEmail is required and must be specified")
	}
	if r.toUserEmail == nil {
		return nil, common.ReportError("toUserEmail is required and must be specified")
	}
	if r.productType == nil {
		return nil, common.ReportError("productType is required and must be specified")
	}
	if r.orderArgs == nil {
		return nil, common.ReportError("orderArgs is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromUserEmail", r.fromUserEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toUserEmail", r.toUserEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productType", r.productType, "form", "")
	{
		t := *r.orderArgs
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderArgs", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MovePositionForSubAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountAssetsRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiQuerySubAccountAssetsRequest) Email(email string) ApiQuerySubAccountAssetsRequest {
	r.email = &email
	return r
}

func (r ApiQuerySubAccountAssetsRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountAssetsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountAssetsRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountAssetsResponse], error) {
	return r.ApiService.QuerySubAccountAssetsExecute(r)
}

/*
QuerySubAccountAssets Query Sub-account Assets (For Master Account) (USER_DATA)
Get /sapi/v3/sub-account/assets

https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Assets-V4

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiQuerySubAccountAssetsRequest
*/
func (a *AssetManagementAPIService) QuerySubAccountAssets(ctx context.Context) ApiQuerySubAccountAssetsRequest {
	return ApiQuerySubAccountAssetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountAssetsResponse
func (a *AssetManagementAPIService) QuerySubAccountAssetsExecute(r ApiQuerySubAccountAssetsRequest) (*common.RestApiResponse[models.QuerySubAccountAssetsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v3/sub-account/assets"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubAccountAssetsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountAssetsAssetManagementRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	recvWindow *int64
}

// [Sub-account email](#email-address)
func (r ApiQuerySubAccountAssetsAssetManagementRequest) Email(email string) ApiQuerySubAccountAssetsAssetManagementRequest {
	r.email = &email
	return r
}

func (r ApiQuerySubAccountAssetsAssetManagementRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountAssetsAssetManagementRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountAssetsAssetManagementRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountAssetsAssetManagementResponse], error) {
	return r.ApiService.QuerySubAccountAssetsAssetManagementExecute(r)
}

/*
QuerySubAccountAssetsAssetManagement Query Sub-account Assets (For Master Account) (USER_DATA)
Get /sapi/v4/sub-account/assets

https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Assets-V4

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param recvWindow -
@return ApiQuerySubAccountAssetsAssetManagementRequest
*/
func (a *AssetManagementAPIService) QuerySubAccountAssetsAssetManagement(ctx context.Context) ApiQuerySubAccountAssetsAssetManagementRequest {
	return ApiQuerySubAccountAssetsAssetManagementRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountAssetsAssetManagementResponse
func (a *AssetManagementAPIService) QuerySubAccountAssetsAssetManagementExecute(r ApiQuerySubAccountAssetsAssetManagementRequest) (*common.RestApiResponse[models.QuerySubAccountAssetsAssetManagementResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v4/sub-account/assets"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubAccountAssetsAssetManagementResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountFuturesAssetTransferHistoryRequest struct {
	ctx         context.Context
	ApiService  *AssetManagementAPIService
	email       *string
	futuresType *int64
	startTime   *int64
	endTime     *int64
	page        *int64
	limit       *int64
	recvWindow  *int64
}

// [Sub-account email](#email-address)
func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) Email(email string) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.email = &email
	return r
}

// 1:USDT-margined Futures，2: Coin-margined Futures
func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) FuturesType(futuresType int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.futuresType = &futuresType
	return r
}

func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) StartTime(startTime int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) EndTime(endTime int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1
func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) Page(page int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) Limit(limit int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountFuturesAssetTransferHistoryResponse], error) {
	return r.ApiService.QuerySubAccountFuturesAssetTransferHistoryExecute(r)
}

/*
QuerySubAccountFuturesAssetTransferHistory Query Sub-account Futures Asset Transfer History (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/futures/internalTransfer

https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Futures-Asset-Transfer-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param futuresType -  1:USDT-margined Futures，2: Coin-margined Futures
@param startTime -
@param endTime -
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQuerySubAccountFuturesAssetTransferHistoryRequest
*/
func (a *AssetManagementAPIService) QuerySubAccountFuturesAssetTransferHistory(ctx context.Context) ApiQuerySubAccountFuturesAssetTransferHistoryRequest {
	return ApiQuerySubAccountFuturesAssetTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountFuturesAssetTransferHistoryResponse
func (a *AssetManagementAPIService) QuerySubAccountFuturesAssetTransferHistoryExecute(r ApiQuerySubAccountFuturesAssetTransferHistoryRequest) (*common.RestApiResponse[models.QuerySubAccountFuturesAssetTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/internalTransfer"

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

	resp, err := SendRequest[models.QuerySubAccountFuturesAssetTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountSpotAssetTransferHistoryRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	fromEmail  *string
	toEmail    *string
	startTime  *int64
	endTime    *int64
	page       *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) FromEmail(fromEmail string) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.fromEmail = &fromEmail
	return r
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) ToEmail(toEmail string) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.toEmail = &toEmail
	return r
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) StartTime(startTime int64) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) EndTime(endTime int64) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1
func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) Page(page int64) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) Limit(limit int64) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountSpotAssetTransferHistoryRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountSpotAssetTransferHistoryResponse], error) {
	return r.ApiService.QuerySubAccountSpotAssetTransferHistoryExecute(r)
}

/*
QuerySubAccountSpotAssetTransferHistory Query Sub-account Spot Asset Transfer History (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/sub/transfer/history

https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Spot-Asset-Transfer-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromEmail -
@param toEmail -
@param startTime -
@param endTime -
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQuerySubAccountSpotAssetTransferHistoryRequest
*/
func (a *AssetManagementAPIService) QuerySubAccountSpotAssetTransferHistory(ctx context.Context) ApiQuerySubAccountSpotAssetTransferHistoryRequest {
	return ApiQuerySubAccountSpotAssetTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountSpotAssetTransferHistoryResponse
func (a *AssetManagementAPIService) QuerySubAccountSpotAssetTransferHistoryExecute(r ApiQuerySubAccountSpotAssetTransferHistoryRequest) (*common.RestApiResponse[models.QuerySubAccountSpotAssetTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/sub/transfer/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromEmail", r.fromEmail, "form", "")
	}
	if r.toEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toEmail", r.toEmail, "form", "")
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

	resp, err := SendRequest[models.QuerySubAccountSpotAssetTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuerySubAccountSpotAssetsSummaryRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	email      *string
	page       *int64
	size       *int64
	recvWindow *int64
}

// Managed sub-account email
func (r ApiQuerySubAccountSpotAssetsSummaryRequest) Email(email string) ApiQuerySubAccountSpotAssetsSummaryRequest {
	r.email = &email
	return r
}

// Default value: 1
func (r ApiQuerySubAccountSpotAssetsSummaryRequest) Page(page int64) ApiQuerySubAccountSpotAssetsSummaryRequest {
	r.page = &page
	return r
}

// default 10, max 20
func (r ApiQuerySubAccountSpotAssetsSummaryRequest) Size(size int64) ApiQuerySubAccountSpotAssetsSummaryRequest {
	r.size = &size
	return r
}

func (r ApiQuerySubAccountSpotAssetsSummaryRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountSpotAssetsSummaryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountSpotAssetsSummaryRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountSpotAssetsSummaryResponse], error) {
	return r.ApiService.QuerySubAccountSpotAssetsSummaryExecute(r)
}

/*
QuerySubAccountSpotAssetsSummary Query Sub-account Spot Assets Summary (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/spotSummary

https://developers.binance.com/docs/sub_account/asset-management/Query-Sub-account-Spot-Assets-Summary

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Managed sub-account email
@param page -  Default value: 1
@param size -  default 10, max 20
@param recvWindow -
@return ApiQuerySubAccountSpotAssetsSummaryRequest
*/
func (a *AssetManagementAPIService) QuerySubAccountSpotAssetsSummary(ctx context.Context) ApiQuerySubAccountSpotAssetsSummaryRequest {
	return ApiQuerySubAccountSpotAssetsSummaryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountSpotAssetsSummaryResponse
func (a *AssetManagementAPIService) QuerySubAccountSpotAssetsSummaryExecute(r ApiQuerySubAccountSpotAssetsSummaryRequest) (*common.RestApiResponse[models.QuerySubAccountSpotAssetsSummaryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/spotSummary"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubAccountSpotAssetsSummaryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUniversalTransferHistoryRequest struct {
	ctx          context.Context
	ApiService   *AssetManagementAPIService
	fromEmail    *string
	toEmail      *string
	clientTranId *string
	startTime    *int64
	endTime      *int64
	page         *int64
	limit        *int64
	recvWindow   *int64
}

func (r ApiQueryUniversalTransferHistoryRequest) FromEmail(fromEmail string) ApiQueryUniversalTransferHistoryRequest {
	r.fromEmail = &fromEmail
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) ToEmail(toEmail string) ApiQueryUniversalTransferHistoryRequest {
	r.toEmail = &toEmail
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) ClientTranId(clientTranId string) ApiQueryUniversalTransferHistoryRequest {
	r.clientTranId = &clientTranId
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) StartTime(startTime int64) ApiQueryUniversalTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) EndTime(endTime int64) ApiQueryUniversalTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1
func (r ApiQueryUniversalTransferHistoryRequest) Page(page int64) ApiQueryUniversalTransferHistoryRequest {
	r.page = &page
	return r
}

// Default value: 1, Max value: 200
func (r ApiQueryUniversalTransferHistoryRequest) Limit(limit int64) ApiQueryUniversalTransferHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) RecvWindow(recvWindow int64) ApiQueryUniversalTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUniversalTransferHistoryRequest) Execute() (*common.RestApiResponse[models.QueryUniversalTransferHistoryResponse], error) {
	return r.ApiService.QueryUniversalTransferHistoryExecute(r)
}

/*
QueryUniversalTransferHistory Query Universal Transfer History (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/universalTransfer

https://developers.binance.com/docs/sub_account/asset-management/Query-Universal-Transfer-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromEmail -
@param toEmail -
@param clientTranId -
@param startTime -
@param endTime -
@param page -  Default value: 1
@param limit -  Default value: 1, Max value: 200
@param recvWindow -
@return ApiQueryUniversalTransferHistoryRequest
*/
func (a *AssetManagementAPIService) QueryUniversalTransferHistory(ctx context.Context) ApiQueryUniversalTransferHistoryRequest {
	return ApiQueryUniversalTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUniversalTransferHistoryResponse
func (a *AssetManagementAPIService) QueryUniversalTransferHistoryExecute(r ApiQueryUniversalTransferHistoryRequest) (*common.RestApiResponse[models.QueryUniversalTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/universalTransfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromEmail", r.fromEmail, "form", "")
	}
	if r.toEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toEmail", r.toEmail, "form", "")
	}
	if r.clientTranId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientTranId", r.clientTranId, "form", "")
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

	resp, err := SendRequest[models.QueryUniversalTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubAccountFuturesAssetTransferRequest struct {
	ctx         context.Context
	ApiService  *AssetManagementAPIService
	fromEmail   *string
	toEmail     *string
	futuresType *int64
	asset       *string
	amount      *float32
	recvWindow  *int64
}

func (r ApiSubAccountFuturesAssetTransferRequest) FromEmail(fromEmail string) ApiSubAccountFuturesAssetTransferRequest {
	r.fromEmail = &fromEmail
	return r
}

func (r ApiSubAccountFuturesAssetTransferRequest) ToEmail(toEmail string) ApiSubAccountFuturesAssetTransferRequest {
	r.toEmail = &toEmail
	return r
}

// 1:USDT-margined Futures，2: Coin-margined Futures
func (r ApiSubAccountFuturesAssetTransferRequest) FuturesType(futuresType int64) ApiSubAccountFuturesAssetTransferRequest {
	r.futuresType = &futuresType
	return r
}

func (r ApiSubAccountFuturesAssetTransferRequest) Asset(asset string) ApiSubAccountFuturesAssetTransferRequest {
	r.asset = &asset
	return r
}

func (r ApiSubAccountFuturesAssetTransferRequest) Amount(amount float32) ApiSubAccountFuturesAssetTransferRequest {
	r.amount = &amount
	return r
}

func (r ApiSubAccountFuturesAssetTransferRequest) RecvWindow(recvWindow int64) ApiSubAccountFuturesAssetTransferRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubAccountFuturesAssetTransferRequest) Execute() (*common.RestApiResponse[models.SubAccountFuturesAssetTransferResponse], error) {
	return r.ApiService.SubAccountFuturesAssetTransferExecute(r)
}

/*
SubAccountFuturesAssetTransfer Sub-account Futures Asset Transfer (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/futures/internalTransfer

https://developers.binance.com/docs/sub_account/asset-management/Sub-account-Futures-Asset-Transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromEmail -
@param toEmail -
@param futuresType -  1:USDT-margined Futures，2: Coin-margined Futures
@param asset -
@param amount -
@param recvWindow -
@return ApiSubAccountFuturesAssetTransferRequest
*/
func (a *AssetManagementAPIService) SubAccountFuturesAssetTransfer(ctx context.Context) ApiSubAccountFuturesAssetTransferRequest {
	return ApiSubAccountFuturesAssetTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubAccountFuturesAssetTransferResponse
func (a *AssetManagementAPIService) SubAccountFuturesAssetTransferExecute(r ApiSubAccountFuturesAssetTransferRequest) (*common.RestApiResponse[models.SubAccountFuturesAssetTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/futures/internalTransfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromEmail == nil {
		return nil, common.ReportError("fromEmail is required and must be specified")
	}
	if r.toEmail == nil {
		return nil, common.ReportError("toEmail is required and must be specified")
	}
	if r.futuresType == nil {
		return nil, common.ReportError("futuresType is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromEmail", r.fromEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toEmail", r.toEmail, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "futuresType", r.futuresType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubAccountFuturesAssetTransferResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubAccountTransferHistoryRequest struct {
	ctx               context.Context
	ApiService        *AssetManagementAPIService
	asset             *string
	type_             *int64
	startTime         *int64
	endTime           *int64
	limit             *int64
	returnFailHistory *bool
	recvWindow        *int64
}

// If not sent, result of all assets will be returned
func (r ApiSubAccountTransferHistoryRequest) Asset(asset string) ApiSubAccountTransferHistoryRequest {
	r.asset = &asset
	return r
}

// 1: transfer in, 2: transfer out
func (r ApiSubAccountTransferHistoryRequest) Type(type_ int64) ApiSubAccountTransferHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiSubAccountTransferHistoryRequest) StartTime(startTime int64) ApiSubAccountTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiSubAccountTransferHistoryRequest) EndTime(endTime int64) ApiSubAccountTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default value: 1, Max value: 200
func (r ApiSubAccountTransferHistoryRequest) Limit(limit int64) ApiSubAccountTransferHistoryRequest {
	r.limit = &limit
	return r
}

// Default &#x60;False&#x60;, return PROCESS and SUCCESS status history; If &#x60;True&#x60;,return PROCESS and SUCCESS and FAILURE status history
func (r ApiSubAccountTransferHistoryRequest) ReturnFailHistory(returnFailHistory bool) ApiSubAccountTransferHistoryRequest {
	r.returnFailHistory = &returnFailHistory
	return r
}

func (r ApiSubAccountTransferHistoryRequest) RecvWindow(recvWindow int64) ApiSubAccountTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubAccountTransferHistoryRequest) Execute() (*common.RestApiResponse[models.SubAccountTransferHistoryResponse], error) {
	return r.ApiService.SubAccountTransferHistoryExecute(r)
}

/*
SubAccountTransferHistory Sub-account Transfer History (For Sub-account) (USER_DATA)
Get /sapi/v1/sub-account/transfer/subUserHistory

https://developers.binance.com/docs/sub_account/asset-management/Sub-account-Transfer-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  If not sent, result of all assets will be returned
@param type_ -  1: transfer in, 2: transfer out
@param startTime -
@param endTime -
@param limit -  Default value: 1, Max value: 200
@param returnFailHistory -  Default `False`, return PROCESS and SUCCESS status history; If `True`,return PROCESS and SUCCESS and FAILURE status history
@param recvWindow -
@return ApiSubAccountTransferHistoryRequest
*/
func (a *AssetManagementAPIService) SubAccountTransferHistory(ctx context.Context) ApiSubAccountTransferHistoryRequest {
	return ApiSubAccountTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubAccountTransferHistoryResponse
func (a *AssetManagementAPIService) SubAccountTransferHistoryExecute(r ApiSubAccountTransferHistoryRequest) (*common.RestApiResponse[models.SubAccountTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/transfer/subUserHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.returnFailHistory != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "returnFailHistory", r.returnFailHistory, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubAccountTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTransferToMasterRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiTransferToMasterRequest) Asset(asset string) ApiTransferToMasterRequest {
	r.asset = &asset
	return r
}

func (r ApiTransferToMasterRequest) Amount(amount float32) ApiTransferToMasterRequest {
	r.amount = &amount
	return r
}

func (r ApiTransferToMasterRequest) RecvWindow(recvWindow int64) ApiTransferToMasterRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTransferToMasterRequest) Execute() (*common.RestApiResponse[models.TransferToMasterResponse], error) {
	return r.ApiService.TransferToMasterExecute(r)
}

/*
TransferToMaster Transfer to Master (For Sub-account) (USER_DATA)
Post /sapi/v1/sub-account/transfer/subToMaster

https://developers.binance.com/docs/sub_account/asset-management/Transfer-to-Master

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param amount -
@param recvWindow -
@return ApiTransferToMasterRequest
*/
func (a *AssetManagementAPIService) TransferToMaster(ctx context.Context) ApiTransferToMasterRequest {
	return ApiTransferToMasterRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransferToMasterResponse
func (a *AssetManagementAPIService) TransferToMasterExecute(r ApiTransferToMasterRequest) (*common.RestApiResponse[models.TransferToMasterResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/transfer/subToMaster"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TransferToMasterResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTransferToSubAccountOfSameMasterRequest struct {
	ctx        context.Context
	ApiService *AssetManagementAPIService
	toEmail    *string
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiTransferToSubAccountOfSameMasterRequest) ToEmail(toEmail string) ApiTransferToSubAccountOfSameMasterRequest {
	r.toEmail = &toEmail
	return r
}

func (r ApiTransferToSubAccountOfSameMasterRequest) Asset(asset string) ApiTransferToSubAccountOfSameMasterRequest {
	r.asset = &asset
	return r
}

func (r ApiTransferToSubAccountOfSameMasterRequest) Amount(amount float32) ApiTransferToSubAccountOfSameMasterRequest {
	r.amount = &amount
	return r
}

func (r ApiTransferToSubAccountOfSameMasterRequest) RecvWindow(recvWindow int64) ApiTransferToSubAccountOfSameMasterRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTransferToSubAccountOfSameMasterRequest) Execute() (*common.RestApiResponse[models.TransferToSubAccountOfSameMasterResponse], error) {
	return r.ApiService.TransferToSubAccountOfSameMasterExecute(r)
}

/*
TransferToSubAccountOfSameMaster Transfer to Sub-account of Same Master (For Sub-account) (USER_DATA)
Post /sapi/v1/sub-account/transfer/subToSub

https://developers.binance.com/docs/sub_account/asset-management/Transfer-to-Sub-account-of-Same-Master

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param toEmail -
@param asset -
@param amount -
@param recvWindow -
@return ApiTransferToSubAccountOfSameMasterRequest
*/
func (a *AssetManagementAPIService) TransferToSubAccountOfSameMaster(ctx context.Context) ApiTransferToSubAccountOfSameMasterRequest {
	return ApiTransferToSubAccountOfSameMasterRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransferToSubAccountOfSameMasterResponse
func (a *AssetManagementAPIService) TransferToSubAccountOfSameMasterExecute(r ApiTransferToSubAccountOfSameMasterRequest) (*common.RestApiResponse[models.TransferToSubAccountOfSameMasterResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/transfer/subToSub"

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

	resp, err := SendRequest[models.TransferToSubAccountOfSameMasterResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUniversalTransferRequest struct {
	ctx             context.Context
	ApiService      *AssetManagementAPIService
	fromAccountType *string
	toAccountType   *string
	asset           *string
	amount          *float32
	fromEmail       *string
	toEmail         *string
	clientTranId    *string
	symbol          *string
	recvWindow      *int64
}

// \&quot;SPOT\&quot;,\&quot;USDT_FUTURE\&quot;,\&quot;COIN_FUTURE\&quot;,\&quot;MARGIN\&quot;(Cross),\&quot;ISOLATED_MARGIN\&quot;
func (r ApiUniversalTransferRequest) FromAccountType(fromAccountType string) ApiUniversalTransferRequest {
	r.fromAccountType = &fromAccountType
	return r
}

// \&quot;SPOT\&quot;,\&quot;USDT_FUTURE\&quot;,\&quot;COIN_FUTURE\&quot;,\&quot;MARGIN\&quot;(Cross),\&quot;ISOLATED_MARGIN\&quot;
func (r ApiUniversalTransferRequest) ToAccountType(toAccountType string) ApiUniversalTransferRequest {
	r.toAccountType = &toAccountType
	return r
}

func (r ApiUniversalTransferRequest) Asset(asset string) ApiUniversalTransferRequest {
	r.asset = &asset
	return r
}

func (r ApiUniversalTransferRequest) Amount(amount float32) ApiUniversalTransferRequest {
	r.amount = &amount
	return r
}

func (r ApiUniversalTransferRequest) FromEmail(fromEmail string) ApiUniversalTransferRequest {
	r.fromEmail = &fromEmail
	return r
}

func (r ApiUniversalTransferRequest) ToEmail(toEmail string) ApiUniversalTransferRequest {
	r.toEmail = &toEmail
	return r
}

func (r ApiUniversalTransferRequest) ClientTranId(clientTranId string) ApiUniversalTransferRequest {
	r.clientTranId = &clientTranId
	return r
}

// Only supported under ISOLATED_MARGIN type
func (r ApiUniversalTransferRequest) Symbol(symbol string) ApiUniversalTransferRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUniversalTransferRequest) RecvWindow(recvWindow int64) ApiUniversalTransferRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUniversalTransferRequest) Execute() (*common.RestApiResponse[models.UniversalTransferResponse], error) {
	return r.ApiService.UniversalTransferExecute(r)
}

/*
UniversalTransfer Universal Transfer (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/universalTransfer

https://developers.binance.com/docs/sub_account/asset-management/Universal-Transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAccountType -  \"SPOT\",\"USDT_FUTURE\",\"COIN_FUTURE\",\"MARGIN\"(Cross),\"ISOLATED_MARGIN\"
@param toAccountType -  \"SPOT\",\"USDT_FUTURE\",\"COIN_FUTURE\",\"MARGIN\"(Cross),\"ISOLATED_MARGIN\"
@param asset -
@param amount -
@param fromEmail -
@param toEmail -
@param clientTranId -
@param symbol -  Only supported under ISOLATED_MARGIN type
@param recvWindow -
@return ApiUniversalTransferRequest
*/
func (a *AssetManagementAPIService) UniversalTransfer(ctx context.Context) ApiUniversalTransferRequest {
	return ApiUniversalTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UniversalTransferResponse
func (a *AssetManagementAPIService) UniversalTransferExecute(r ApiUniversalTransferRequest) (*common.RestApiResponse[models.UniversalTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/universalTransfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromAccountType == nil {
		return nil, common.ReportError("fromAccountType is required and must be specified")
	}
	if r.toAccountType == nil {
		return nil, common.ReportError("toAccountType is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	if r.fromEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromEmail", r.fromEmail, "form", "")
	}
	if r.toEmail != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toEmail", r.toEmail, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAccountType", r.fromAccountType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAccountType", r.toAccountType, "form", "")
	if r.clientTranId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientTranId", r.clientTranId, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UniversalTransferResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
