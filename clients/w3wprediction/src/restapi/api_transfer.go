/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package binancew3wpredictionrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TransferAPIService TransferAPI Service
type TransferAPIService Service

type ApiApplyMmDepositRequest struct {
	ctx             context.Context
	ApiService      *TransferAPIService
	fromToken       *string
	fromTokenAmount *string
	toToken         *string
	accountType     *models.PlaceOrderAccountTypeParameter
	chainId         *string
}

// Source token symbol (e.g. &#x60;USDT&#x60;)
func (r ApiApplyMmDepositRequest) FromToken(fromToken string) ApiApplyMmDepositRequest {
	r.fromToken = &fromToken
	return r
}

// Source token amount in WEI (18 decimals). Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT
func (r ApiApplyMmDepositRequest) FromTokenAmount(fromTokenAmount string) ApiApplyMmDepositRequest {
	r.fromTokenAmount = &fromTokenAmount
	return r
}

// Target token symbol (e.g. &#x60;USDT&#x60;)
func (r ApiApplyMmDepositRequest) ToToken(toToken string) ApiApplyMmDepositRequest {
	r.toToken = &toToken
	return r
}

// Target CEX account type. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;
func (r ApiApplyMmDepositRequest) AccountType(accountType models.PlaceOrderAccountTypeParameter) ApiApplyMmDepositRequest {
	r.accountType = &accountType
	return r
}

// Chain ID. Default &#x60;56&#x60; (BSC)
func (r ApiApplyMmDepositRequest) ChainId(chainId string) ApiApplyMmDepositRequest {
	r.chainId = &chainId
	return r
}

func (r ApiApplyMmDepositRequest) Execute() (*common.RestApiResponse[models.ApplyMmDepositResponse], error) {
	return r.ApiService.ApplyMmDepositExecute(r)
}

/*
ApplyMmDeposit Apply MM Deposit (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/deposit/apply

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#apply-mm-deposit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromToken -  Source token symbol (e.g. `USDT`)
@param fromTokenAmount -  Source token amount in WEI (18 decimals). Example: `1000000000000000000` = 1 USDT
@param toToken -  Target token symbol (e.g. `USDT`)
@param accountType -  Target CEX account type. Enum: `SPOT`, `FUNDING`
@param chainId -  Chain ID. Default `56` (BSC)
@return ApiApplyMmDepositRequest
*/
func (a *TransferAPIService) ApplyMmDeposit(ctx context.Context) ApiApplyMmDepositRequest {
	return ApiApplyMmDepositRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApplyMmDepositResponse
func (a *TransferAPIService) ApplyMmDepositExecute(r ApiApplyMmDepositRequest) (*common.RestApiResponse[models.ApplyMmDepositResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/deposit/apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromToken == nil {
		return nil, common.ReportError("fromToken is required and must be specified")
	}

	if r.fromTokenAmount == nil {
		return nil, common.ReportError("fromTokenAmount is required and must be specified")
	}

	if r.toToken == nil {
		return nil, common.ReportError("toToken is required and must be specified")
	}

	if r.accountType == nil {
		return nil, common.ReportError("accountType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromToken", r.fromToken, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromTokenAmount", r.fromTokenAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toToken", r.toToken, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	if r.chainId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "chainId", r.chainId, "form", "")
	}

	resp, err := SendRequest[models.ApplyMmDepositResponse](
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

type ApiApplyMmWithdrawRequest struct {
	ctx             context.Context
	ApiService      *TransferAPIService
	coin            *string
	network         *string
	amount          *string
	withdrawOrderId *string
	walletType      *models.ApplyMmWithdrawWalletTypeParameter
	name            *string
}

// Coin to withdraw (e.g. &#x60;USDT&#x60;)
func (r ApiApplyMmWithdrawRequest) Coin(coin string) ApiApplyMmWithdrawRequest {
	r.coin = &coin
	return r
}

// Network (e.g. &#x60;BEP20&#x60;)
func (r ApiApplyMmWithdrawRequest) Network(network string) ApiApplyMmWithdrawRequest {
	r.network = &network
	return r
}

// Amount to withdraw (must be &gt; 0)
func (r ApiApplyMmWithdrawRequest) Amount(amount string) ApiApplyMmWithdrawRequest {
	r.amount = &amount
	return r
}

// Client withdraw order id (idempotency key)
func (r ApiApplyMmWithdrawRequest) WithdrawOrderId(withdrawOrderId string) ApiApplyMmWithdrawRequest {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

// Source CEX account type. Enum: &#x60;0&#x60; (SPOT), &#x60;1&#x60; (FUNDING). Default &#x60;0&#x60;. Must be &#x60;0&#x60; or &#x60;1&#x60;; any other value is rejected
func (r ApiApplyMmWithdrawRequest) WalletType(walletType models.ApplyMmWithdrawWalletTypeParameter) ApiApplyMmWithdrawRequest {
	r.walletType = &walletType
	return r
}

// Remark for the withdraw
func (r ApiApplyMmWithdrawRequest) Name(name string) ApiApplyMmWithdrawRequest {
	r.name = &name
	return r
}

func (r ApiApplyMmWithdrawRequest) Execute() (*common.RestApiResponse[models.ApplyMmWithdrawResponse], error) {
	return r.ApiService.ApplyMmWithdrawExecute(r)
}

/*
ApplyMmWithdraw Apply MM Withdraw (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/withdraw/apply

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#apply-mm-withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -  Coin to withdraw (e.g. `USDT`)
@param network -  Network (e.g. `BEP20`)
@param amount -  Amount to withdraw (must be > 0)
@param withdrawOrderId -  Client withdraw order id (idempotency key)
@param walletType -  Source CEX account type. Enum: `0` (SPOT), `1` (FUNDING). Default `0`. Must be `0` or `1`; any other value is rejected
@param name -  Remark for the withdraw
@return ApiApplyMmWithdrawRequest
*/
func (a *TransferAPIService) ApplyMmWithdraw(ctx context.Context) ApiApplyMmWithdrawRequest {
	return ApiApplyMmWithdrawRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApplyMmWithdrawResponse
func (a *TransferAPIService) ApplyMmWithdrawExecute(r ApiApplyMmWithdrawRequest) (*common.RestApiResponse[models.ApplyMmWithdrawResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/withdraw/apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}

	if r.network == nil {
		return nil, common.ReportError("network is required and must be specified")
	}

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	if r.name != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}

	resp, err := SendRequest[models.ApplyMmWithdrawResponse](
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

type ApiCreateInboundTransferRequest struct {
	ctx             context.Context
	ApiService      *TransferAPIService
	walletId        *string
	walletAddress   *string
	fromTokenAmount *string
	accountType     *models.PlaceOrderAccountTypeParameter
	fromToken       *string
	toToken         *string
	chainId         *string
}

// Wallet ID
func (r ApiCreateInboundTransferRequest) WalletId(walletId string) ApiCreateInboundTransferRequest {
	r.walletId = &walletId
	return r
}

// User&#39;s prediction wallet address
func (r ApiCreateInboundTransferRequest) WalletAddress(walletAddress string) ApiCreateInboundTransferRequest {
	r.walletAddress = &walletAddress
	return r
}

// Transfer amount in wei (18 decimals). Must be &gt; 0. Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT
func (r ApiCreateInboundTransferRequest) FromTokenAmount(fromTokenAmount string) ApiCreateInboundTransferRequest {
	r.fromTokenAmount = &fromTokenAmount
	return r
}

// Destination CEX account. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;
func (r ApiCreateInboundTransferRequest) AccountType(accountType models.PlaceOrderAccountTypeParameter) ApiCreateInboundTransferRequest {
	r.accountType = &accountType
	return r
}

// Source token symbol. Default &#x60;USDT&#x60;
func (r ApiCreateInboundTransferRequest) FromToken(fromToken string) ApiCreateInboundTransferRequest {
	r.fromToken = &fromToken
	return r
}

// Destination token symbol. Default &#x60;USDT&#x60;
func (r ApiCreateInboundTransferRequest) ToToken(toToken string) ApiCreateInboundTransferRequest {
	r.toToken = &toToken
	return r
}

// Chain ID. Default &#x60;56&#x60; (BSC)
func (r ApiCreateInboundTransferRequest) ChainId(chainId string) ApiCreateInboundTransferRequest {
	r.chainId = &chainId
	return r
}

func (r ApiCreateInboundTransferRequest) Execute() (*common.RestApiResponse[models.CreateInboundTransferResponse], error) {
	return r.ApiService.CreateInboundTransferExecute(r)
}

/*
CreateInboundTransfer Create Inbound Transfer (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/transfer/inbound

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#create-inbound-transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletId -  Wallet ID
@param walletAddress -  User's prediction wallet address
@param fromTokenAmount -  Transfer amount in wei (18 decimals). Must be > 0. Example: `1000000000000000000` = 1 USDT
@param accountType -  Destination CEX account. Enum: `SPOT`, `FUNDING`
@param fromToken -  Source token symbol. Default `USDT`
@param toToken -  Destination token symbol. Default `USDT`
@param chainId -  Chain ID. Default `56` (BSC)
@return ApiCreateInboundTransferRequest
*/
func (a *TransferAPIService) CreateInboundTransfer(ctx context.Context) ApiCreateInboundTransferRequest {
	return ApiCreateInboundTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateInboundTransferResponse
func (a *TransferAPIService) CreateInboundTransferExecute(r ApiCreateInboundTransferRequest) (*common.RestApiResponse[models.CreateInboundTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/transfer/inbound"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletId == nil {
		return nil, common.ReportError("walletId is required and must be specified")
	}

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.fromTokenAmount == nil {
		return nil, common.ReportError("fromTokenAmount is required and must be specified")
	}

	if r.accountType == nil {
		return nil, common.ReportError("accountType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletId", r.walletId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromTokenAmount", r.fromTokenAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	if r.fromToken != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromToken", r.fromToken, "form", "")
	}
	if r.toToken != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toToken", r.toToken, "form", "")
	}
	if r.chainId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "chainId", r.chainId, "form", "")
	}

	resp, err := SendRequest[models.CreateInboundTransferResponse](
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

type ApiCreateOutboundTransferRequest struct {
	ctx             context.Context
	ApiService      *TransferAPIService
	walletId        *string
	walletAddress   *string
	fromTokenAmount *string
	accountType     *models.PlaceOrderAccountTypeParameter
	sourceBiz       *models.CreateOutboundTransferSourceBizParameter
	fromToken       *string
	toToken         *string
	chainId         *string
}

// Wallet ID
func (r ApiCreateOutboundTransferRequest) WalletId(walletId string) ApiCreateOutboundTransferRequest {
	r.walletId = &walletId
	return r
}

// User&#39;s prediction wallet address
func (r ApiCreateOutboundTransferRequest) WalletAddress(walletAddress string) ApiCreateOutboundTransferRequest {
	r.walletAddress = &walletAddress
	return r
}

// Transfer amount in wei (18 decimals). Must be &gt; 0. Example: &#x60;1000000000000000000&#x60; &#x3D; 1 USDT
func (r ApiCreateOutboundTransferRequest) FromTokenAmount(fromTokenAmount string) ApiCreateOutboundTransferRequest {
	r.fromTokenAmount = &fromTokenAmount
	return r
}

// Source CEX account. Enum: &#x60;SPOT&#x60;, &#x60;FUNDING&#x60;
func (r ApiCreateOutboundTransferRequest) AccountType(accountType models.PlaceOrderAccountTypeParameter) ApiCreateOutboundTransferRequest {
	r.accountType = &accountType
	return r
}

// Business source. Enum: &#x60;USER_TRANSFER&#x60;, &#x60;PREDICTION_BUY&#x60;
func (r ApiCreateOutboundTransferRequest) SourceBiz(sourceBiz models.CreateOutboundTransferSourceBizParameter) ApiCreateOutboundTransferRequest {
	r.sourceBiz = &sourceBiz
	return r
}

// Source token symbol. Default &#x60;USDT&#x60;
func (r ApiCreateOutboundTransferRequest) FromToken(fromToken string) ApiCreateOutboundTransferRequest {
	r.fromToken = &fromToken
	return r
}

// Destination token symbol. Default &#x60;USDT&#x60;
func (r ApiCreateOutboundTransferRequest) ToToken(toToken string) ApiCreateOutboundTransferRequest {
	r.toToken = &toToken
	return r
}

// Chain ID. Default &#x60;56&#x60; (BSC)
func (r ApiCreateOutboundTransferRequest) ChainId(chainId string) ApiCreateOutboundTransferRequest {
	r.chainId = &chainId
	return r
}

func (r ApiCreateOutboundTransferRequest) Execute() (*common.RestApiResponse[models.CreateOutboundTransferResponse], error) {
	return r.ApiService.CreateOutboundTransferExecute(r)
}

/*
CreateOutboundTransfer Create Outbound Transfer (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/transfer/outbound

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#create-outbound-transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletId -  Wallet ID
@param walletAddress -  User's prediction wallet address
@param fromTokenAmount -  Transfer amount in wei (18 decimals). Must be > 0. Example: `1000000000000000000` = 1 USDT
@param accountType -  Source CEX account. Enum: `SPOT`, `FUNDING`
@param sourceBiz -  Business source. Enum: `USER_TRANSFER`, `PREDICTION_BUY`
@param fromToken -  Source token symbol. Default `USDT`
@param toToken -  Destination token symbol. Default `USDT`
@param chainId -  Chain ID. Default `56` (BSC)
@return ApiCreateOutboundTransferRequest
*/
func (a *TransferAPIService) CreateOutboundTransfer(ctx context.Context) ApiCreateOutboundTransferRequest {
	return ApiCreateOutboundTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateOutboundTransferResponse
func (a *TransferAPIService) CreateOutboundTransferExecute(r ApiCreateOutboundTransferRequest) (*common.RestApiResponse[models.CreateOutboundTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/transfer/outbound"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletId == nil {
		return nil, common.ReportError("walletId is required and must be specified")
	}

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.fromTokenAmount == nil {
		return nil, common.ReportError("fromTokenAmount is required and must be specified")
	}

	if r.accountType == nil {
		return nil, common.ReportError("accountType is required and must be specified")
	}

	if r.sourceBiz == nil {
		return nil, common.ReportError("sourceBiz is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletId", r.walletId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromTokenAmount", r.fromTokenAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sourceBiz", r.sourceBiz, "form", "")
	if r.fromToken != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromToken", r.fromToken, "form", "")
	}
	if r.toToken != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toToken", r.toToken, "form", "")
	}
	if r.chainId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "chainId", r.chainId, "form", "")
	}

	resp, err := SendRequest[models.CreateOutboundTransferResponse](
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

type ApiQueryTransferListRequest struct {
	ctx           context.Context
	ApiService    *TransferAPIService
	walletAddress *string
	startDate     *string
	endDate       *string
	tokenSymbol   *string
	direction     *models.QueryTransferListDirectionParameter
	offset        *int32
	limit         *int32
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryTransferListRequest) WalletAddress(walletAddress string) ApiQueryTransferListRequest {
	r.walletAddress = &walletAddress
	return r
}

// Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60;
func (r ApiQueryTransferListRequest) StartDate(startDate string) ApiQueryTransferListRequest {
	r.startDate = &startDate
	return r
}

// End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60;
func (r ApiQueryTransferListRequest) EndDate(endDate string) ApiQueryTransferListRequest {
	r.endDate = &endDate
	return r
}

// Filter by token symbol (e.g. &#x60;USDT&#x60;)
func (r ApiQueryTransferListRequest) TokenSymbol(tokenSymbol string) ApiQueryTransferListRequest {
	r.tokenSymbol = &tokenSymbol
	return r
}

// Filter by direction. Enum: &#x60;INBOUND&#x60;, &#x60;OUTBOUND&#x60;
func (r ApiQueryTransferListRequest) Direction(direction models.QueryTransferListDirectionParameter) ApiQueryTransferListRequest {
	r.direction = &direction
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiQueryTransferListRequest) Offset(offset int32) ApiQueryTransferListRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiQueryTransferListRequest) Limit(limit int32) ApiQueryTransferListRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiQueryTransferListRequest) RecvWindow(recvWindow int64) ApiQueryTransferListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryTransferListRequest) Execute() (*common.RestApiResponse[models.QueryTransferListResponse], error) {
	return r.ApiService.QueryTransferListExecute(r)
}

/*
QueryTransferList Query Transfer List (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/transfer/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#query-transfer-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param startDate -  Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate`
@param endDate -  End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate`
@param tokenSymbol -  Filter by token symbol (e.g. `USDT`)
@param direction -  Filter by direction. Enum: `INBOUND`, `OUTBOUND`
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryTransferListRequest
*/
func (a *TransferAPIService) QueryTransferList(ctx context.Context) ApiQueryTransferListRequest {
	return ApiQueryTransferListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryTransferListResponse
func (a *TransferAPIService) QueryTransferListExecute(r ApiQueryTransferListRequest) (*common.RestApiResponse[models.QueryTransferListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/transfer/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.startDate == nil {
		return nil, common.ReportError("startDate is required and must be specified")
	}

	if r.endDate == nil {
		return nil, common.ReportError("endDate is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	if r.tokenSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenSymbol", r.tokenSymbol, "form", "")
	}
	if r.direction != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
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

	resp, err := SendRequest[models.QueryTransferListResponse](
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

type ApiQueryTransferStatusRequest struct {
	ctx        context.Context
	ApiService *TransferAPIService
	transferId *string
	recvWindow *int64
}

// Transfer ID returned from outbound/inbound transfer
func (r ApiQueryTransferStatusRequest) TransferId(transferId string) ApiQueryTransferStatusRequest {
	r.transferId = &transferId
	return r
}

// Request validity window in milliseconds
func (r ApiQueryTransferStatusRequest) RecvWindow(recvWindow int64) ApiQueryTransferStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryTransferStatusRequest) Execute() (*common.RestApiResponse[models.QueryTransferStatusResponse], error) {
	return r.ApiService.QueryTransferStatusExecute(r)
}

/*
QueryTransferStatus Query Transfer Status (PREDICTION_TRADE)
Get /sapi/v1/w3w/wallet/prediction/transfer/status

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/transfer#query-transfer-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param transferId -  Transfer ID returned from outbound/inbound transfer
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryTransferStatusRequest
*/
func (a *TransferAPIService) QueryTransferStatus(ctx context.Context) ApiQueryTransferStatusRequest {
	return ApiQueryTransferStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryTransferStatusResponse
func (a *TransferAPIService) QueryTransferStatusExecute(r ApiQueryTransferStatusRequest) (*common.RestApiResponse[models.QueryTransferStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/transfer/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.transferId == nil {
		return nil, common.ReportError("transferId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferId", r.transferId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryTransferStatusResponse](
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
