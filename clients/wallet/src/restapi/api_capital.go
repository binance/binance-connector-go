/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package binancewalletrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// CapitalAPIService CapitalAPI Service
type CapitalAPIService Service

type ApiAllCoinsInformationRequest struct {
	ctx        context.Context
	ApiService *CapitalAPIService
	recvWindow *int64
}

func (r ApiAllCoinsInformationRequest) RecvWindow(recvWindow int64) ApiAllCoinsInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllCoinsInformationRequest) Execute() (*common.RestApiResponse[models.AllCoinsInformationResponse], error) {
	return r.ApiService.AllCoinsInformationExecute(r)
}

/*
AllCoinsInformation All Coins' Information (USER_DATA)
Get /sapi/v1/capital/config/getall

https://developers.binance.com/docs/wallet/capital/all-coins-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAllCoinsInformationRequest
*/
func (a *CapitalAPIService) AllCoinsInformation(ctx context.Context) ApiAllCoinsInformationRequest {
	return ApiAllCoinsInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AllCoinsInformationResponse
func (a *CapitalAPIService) AllCoinsInformationExecute(r ApiAllCoinsInformationRequest) (*common.RestApiResponse[models.AllCoinsInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/config/getall"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AllCoinsInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDepositAddressRequest struct {
	ctx        context.Context
	ApiService *CapitalAPIService
	coin       *string
	network    *string
	amount     *float32
	recvWindow *int64
}

func (r ApiDepositAddressRequest) Coin(coin string) ApiDepositAddressRequest {
	r.coin = &coin
	return r
}

func (r ApiDepositAddressRequest) Network(network string) ApiDepositAddressRequest {
	r.network = &network
	return r
}

func (r ApiDepositAddressRequest) Amount(amount float32) ApiDepositAddressRequest {
	r.amount = &amount
	return r
}

func (r ApiDepositAddressRequest) RecvWindow(recvWindow int64) ApiDepositAddressRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDepositAddressRequest) Execute() (*common.RestApiResponse[models.DepositAddressResponse], error) {
	return r.ApiService.DepositAddressExecute(r)
}

/*
DepositAddress Deposit Address(supporting network) (USER_DATA)
Get /sapi/v1/capital/deposit/address

https://developers.binance.com/docs/wallet/capital/deposite-address

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param network -
@param amount -
@param recvWindow -
@return ApiDepositAddressRequest
*/
func (a *CapitalAPIService) DepositAddress(ctx context.Context) ApiDepositAddressRequest {
	return ApiDepositAddressRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositAddressResponse
func (a *CapitalAPIService) DepositAddressExecute(r ApiDepositAddressRequest) (*common.RestApiResponse[models.DepositAddressResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/address"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}

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

	resp, err := SendRequest[models.DepositAddressResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDepositHistoryRequest struct {
	ctx           context.Context
	ApiService    *CapitalAPIService
	includeSource *bool
	coin          *string
	status        *int64
	startTime     *int64
	endTime       *int64
	offset        *int64
	limit         *int64
	recvWindow    *int64
	txId          *string
}

// Default: &#x60;false&#x60;, return &#x60;sourceAddress&#x60;field when set to &#x60;true&#x60;
func (r ApiDepositHistoryRequest) IncludeSource(includeSource bool) ApiDepositHistoryRequest {
	r.includeSource = &includeSource
	return r
}

func (r ApiDepositHistoryRequest) Coin(coin string) ApiDepositHistoryRequest {
	r.coin = &coin
	return r
}

// 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed)
func (r ApiDepositHistoryRequest) Status(status int64) ApiDepositHistoryRequest {
	r.status = &status
	return r
}

func (r ApiDepositHistoryRequest) StartTime(startTime int64) ApiDepositHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiDepositHistoryRequest) EndTime(endTime int64) ApiDepositHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default: 0
func (r ApiDepositHistoryRequest) Offset(offset int64) ApiDepositHistoryRequest {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiDepositHistoryRequest) Limit(limit int64) ApiDepositHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiDepositHistoryRequest) RecvWindow(recvWindow int64) ApiDepositHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDepositHistoryRequest) TxId(txId string) ApiDepositHistoryRequest {
	r.txId = &txId
	return r
}

func (r ApiDepositHistoryRequest) Execute() (*common.RestApiResponse[models.DepositHistoryResponse], error) {
	return r.ApiService.DepositHistoryExecute(r)
}

/*
DepositHistory Deposit History (supporting network) (USER_DATA)
Get /sapi/v1/capital/deposit/hisrec

https://developers.binance.com/docs/wallet/capital/deposite-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param includeSource -  Default: `false`, return `sourceAddress`field when set to `true`
@param coin -
@param status -  0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed)
@param startTime -
@param endTime -
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@param recvWindow -
@param txId -
@return ApiDepositHistoryRequest
*/
func (a *CapitalAPIService) DepositHistory(ctx context.Context) ApiDepositHistoryRequest {
	return ApiDepositHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositHistoryResponse
func (a *CapitalAPIService) DepositHistoryExecute(r ApiDepositHistoryRequest) (*common.RestApiResponse[models.DepositHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/hisrec"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.includeSource != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "includeSource", r.includeSource, "form", "")
	}
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
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}

	resp, err := SendRequest[models.DepositHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchDepositAddressListWithNetworkRequest struct {
	ctx        context.Context
	ApiService *CapitalAPIService
	coin       *string
	network    *string
}

func (r ApiFetchDepositAddressListWithNetworkRequest) Coin(coin string) ApiFetchDepositAddressListWithNetworkRequest {
	r.coin = &coin
	return r
}

func (r ApiFetchDepositAddressListWithNetworkRequest) Network(network string) ApiFetchDepositAddressListWithNetworkRequest {
	r.network = &network
	return r
}

func (r ApiFetchDepositAddressListWithNetworkRequest) Execute() (*common.RestApiResponse[models.FetchDepositAddressListWithNetworkResponse], error) {
	return r.ApiService.FetchDepositAddressListWithNetworkExecute(r)
}

/*
FetchDepositAddressListWithNetwork Fetch deposit address list with network(USER_DATA)
Get /sapi/v1/capital/deposit/address/list

https://developers.binance.com/docs/wallet/capital/Fetch-deposit-address-list-with-network

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param network -
@return ApiFetchDepositAddressListWithNetworkRequest
*/
func (a *CapitalAPIService) FetchDepositAddressListWithNetwork(ctx context.Context) ApiFetchDepositAddressListWithNetworkRequest {
	return ApiFetchDepositAddressListWithNetworkRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchDepositAddressListWithNetworkResponse
func (a *CapitalAPIService) FetchDepositAddressListWithNetworkExecute(r ApiFetchDepositAddressListWithNetworkRequest) (*common.RestApiResponse[models.FetchDepositAddressListWithNetworkResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/address/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}

	resp, err := SendRequest[models.FetchDepositAddressListWithNetworkResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchWithdrawAddressListRequest struct {
	ctx        context.Context
	ApiService *CapitalAPIService
}

func (r ApiFetchWithdrawAddressListRequest) Execute() (*common.RestApiResponse[models.FetchWithdrawAddressListResponse], error) {
	return r.ApiService.FetchWithdrawAddressListExecute(r)
}

/*
FetchWithdrawAddressList Fetch withdraw address list (USER_DATA)
Get /sapi/v1/capital/withdraw/address/list

https://developers.binance.com/docs/wallet/capital/fetch-withdraw-address

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFetchWithdrawAddressListRequest
*/
func (a *CapitalAPIService) FetchWithdrawAddressList(ctx context.Context) ApiFetchWithdrawAddressListRequest {
	return ApiFetchWithdrawAddressListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchWithdrawAddressListResponse
func (a *CapitalAPIService) FetchWithdrawAddressListExecute(r ApiFetchWithdrawAddressListRequest) (*common.RestApiResponse[models.FetchWithdrawAddressListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/withdraw/address/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.FetchWithdrawAddressListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchWithdrawQuotaRequest struct {
	ctx        context.Context
	ApiService *CapitalAPIService
}

func (r ApiFetchWithdrawQuotaRequest) Execute() (*common.RestApiResponse[models.FetchWithdrawQuotaResponse], error) {
	return r.ApiService.FetchWithdrawQuotaExecute(r)
}

/*
FetchWithdrawQuota Fetch withdraw quota (USER_DATA)
Get /sapi/v1/capital/withdraw/quota

https://developers.binance.com/docs/wallet/capital/Fetch-withdraw-quota

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiFetchWithdrawQuotaRequest
*/
func (a *CapitalAPIService) FetchWithdrawQuota(ctx context.Context) ApiFetchWithdrawQuotaRequest {
	return ApiFetchWithdrawQuotaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchWithdrawQuotaResponse
func (a *CapitalAPIService) FetchWithdrawQuotaExecute(r ApiFetchWithdrawQuotaRequest) (*common.RestApiResponse[models.FetchWithdrawQuotaResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/withdraw/quota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.FetchWithdrawQuotaResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOneClickArrivalDepositApplyRequest struct {
	ctx          context.Context
	ApiService   *CapitalAPIService
	depositId    *int64
	txId         *string
	subAccountId *int64
	subUserId    *int64
}

// Deposit record Id, priority use
func (r ApiOneClickArrivalDepositApplyRequest) DepositId(depositId int64) ApiOneClickArrivalDepositApplyRequest {
	r.depositId = &depositId
	return r
}

func (r ApiOneClickArrivalDepositApplyRequest) TxId(txId string) ApiOneClickArrivalDepositApplyRequest {
	r.txId = &txId
	return r
}

// Sub-accountId of Cloud user
func (r ApiOneClickArrivalDepositApplyRequest) SubAccountId(subAccountId int64) ApiOneClickArrivalDepositApplyRequest {
	r.subAccountId = &subAccountId
	return r
}

// Sub-userId of parent user
func (r ApiOneClickArrivalDepositApplyRequest) SubUserId(subUserId int64) ApiOneClickArrivalDepositApplyRequest {
	r.subUserId = &subUserId
	return r
}

func (r ApiOneClickArrivalDepositApplyRequest) Execute() (*common.RestApiResponse[models.OneClickArrivalDepositApplyResponse], error) {
	return r.ApiService.OneClickArrivalDepositApplyExecute(r)
}

/*
OneClickArrivalDepositApply One click arrival deposit apply (for expired address deposit) (USER_DATA)
Post /sapi/v1/capital/deposit/credit-apply

https://developers.binance.com/docs/wallet/capital/one-click-arrival-deposite-apply

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param depositId -  Deposit record Id, priority use
@param txId -
@param subAccountId -  Sub-accountId of Cloud user
@param subUserId -  Sub-userId of parent user
@return ApiOneClickArrivalDepositApplyRequest
*/
func (a *CapitalAPIService) OneClickArrivalDepositApply(ctx context.Context) ApiOneClickArrivalDepositApplyRequest {
	return ApiOneClickArrivalDepositApplyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OneClickArrivalDepositApplyResponse
func (a *CapitalAPIService) OneClickArrivalDepositApplyExecute(r ApiOneClickArrivalDepositApplyRequest) (*common.RestApiResponse[models.OneClickArrivalDepositApplyResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/deposit/credit-apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.depositId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "depositId", r.depositId, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.subAccountId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountId", r.subAccountId, "form", "")
	}
	if r.subUserId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subUserId", r.subUserId, "form", "")
	}

	resp, err := SendRequest[models.OneClickArrivalDepositApplyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawRequest struct {
	ctx                context.Context
	ApiService         *CapitalAPIService
	coin               *string
	address            *string
	amount             *float32
	withdrawOrderId    *string
	network            *string
	addressTag         *string
	transactionFeeFlag *bool
	name               *string
	walletType         *int64
	recvWindow         *int64
}

func (r ApiWithdrawRequest) Coin(coin string) ApiWithdrawRequest {
	r.coin = &coin
	return r
}

func (r ApiWithdrawRequest) Address(address string) ApiWithdrawRequest {
	r.address = &address
	return r
}

func (r ApiWithdrawRequest) Amount(amount float32) ApiWithdrawRequest {
	r.amount = &amount
	return r
}

// client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query.
func (r ApiWithdrawRequest) WithdrawOrderId(withdrawOrderId string) ApiWithdrawRequest {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

func (r ApiWithdrawRequest) Network(network string) ApiWithdrawRequest {
	r.network = &network
	return r
}

// Secondary address identifier for coins like XRP,XMR etc.
func (r ApiWithdrawRequest) AddressTag(addressTag string) ApiWithdrawRequest {
	r.addressTag = &addressTag
	return r
}

// When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;.
func (r ApiWithdrawRequest) TransactionFeeFlag(transactionFeeFlag bool) ApiWithdrawRequest {
	r.transactionFeeFlag = &transactionFeeFlag
	return r
}

// Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60;
func (r ApiWithdrawRequest) Name(name string) ApiWithdrawRequest {
	r.name = &name
	return r
}

// The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit
func (r ApiWithdrawRequest) WalletType(walletType int64) ApiWithdrawRequest {
	r.walletType = &walletType
	return r
}

func (r ApiWithdrawRequest) RecvWindow(recvWindow int64) ApiWithdrawRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawRequest) Execute() (*common.RestApiResponse[models.WithdrawResponse], error) {
	return r.ApiService.WithdrawExecute(r)
}

/*
Withdraw Withdraw(USER_DATA)
Post /sapi/v1/capital/withdraw/apply

https://developers.binance.com/docs/wallet/capital/Withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param address -
@param amount -
@param withdrawOrderId -  client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query.
@param network -
@param addressTag -  Secondary address identifier for coins like XRP,XMR etc.
@param transactionFeeFlag -  When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`.
@param name -  Description of the address. Address book cap is 200, space in name should be encoded into `%20`
@param walletType -  The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit
@param recvWindow -
@return ApiWithdrawRequest
*/
func (a *CapitalAPIService) Withdraw(ctx context.Context) ApiWithdrawRequest {
	return ApiWithdrawRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawResponse
func (a *CapitalAPIService) WithdrawExecute(r ApiWithdrawRequest) (*common.RestApiResponse[models.WithdrawResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/withdraw/apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}
	if r.address == nil {
		return nil, common.ReportError("address is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "form", "")
	if r.addressTag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "addressTag", r.addressTag, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.transactionFeeFlag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transactionFeeFlag", r.transactionFeeFlag, "form", "")
	}
	if r.name != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WithdrawResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawHistoryRequest struct {
	ctx             context.Context
	ApiService      *CapitalAPIService
	coin            *string
	withdrawOrderId *string
	status          *int64
	offset          *int64
	limit           *int64
	idList          *string
	startTime       *int64
	endTime         *int64
	recvWindow      *int64
}

func (r ApiWithdrawHistoryRequest) Coin(coin string) ApiWithdrawHistoryRequest {
	r.coin = &coin
	return r
}

// client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query.
func (r ApiWithdrawHistoryRequest) WithdrawOrderId(withdrawOrderId string) ApiWithdrawHistoryRequest {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

// 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed)
func (r ApiWithdrawHistoryRequest) Status(status int64) ApiWithdrawHistoryRequest {
	r.status = &status
	return r
}

// Default: 0
func (r ApiWithdrawHistoryRequest) Offset(offset int64) ApiWithdrawHistoryRequest {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiWithdrawHistoryRequest) Limit(limit int64) ApiWithdrawHistoryRequest {
	r.limit = &limit
	return r
}

// id list returned in the response of POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, separated by &#x60;,&#x60;
func (r ApiWithdrawHistoryRequest) IdList(idList string) ApiWithdrawHistoryRequest {
	r.idList = &idList
	return r
}

func (r ApiWithdrawHistoryRequest) StartTime(startTime int64) ApiWithdrawHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiWithdrawHistoryRequest) EndTime(endTime int64) ApiWithdrawHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiWithdrawHistoryRequest) RecvWindow(recvWindow int64) ApiWithdrawHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawHistoryRequest) Execute() (*common.RestApiResponse[models.WithdrawHistoryResponse], error) {
	return r.ApiService.WithdrawHistoryExecute(r)
}

/*
WithdrawHistory Withdraw History (supporting network) (USER_DATA)
Get /sapi/v1/capital/withdraw/history

https://developers.binance.com/docs/wallet/capital/Withdraw-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param withdrawOrderId -  client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query.
@param status -  0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed)
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@param idList -  id list returned in the response of POST `/sapi/v1/capital/withdraw/apply`, separated by `,`
@param startTime -
@param endTime -
@param recvWindow -
@return ApiWithdrawHistoryRequest
*/
func (a *CapitalAPIService) WithdrawHistory(ctx context.Context) ApiWithdrawHistoryRequest {
	return ApiWithdrawHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawHistoryResponse
func (a *CapitalAPIService) WithdrawHistoryExecute(r ApiWithdrawHistoryRequest) (*common.RestApiResponse[models.WithdrawHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/capital/withdraw/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.status != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.idList != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "idList", r.idList, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WithdrawHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
