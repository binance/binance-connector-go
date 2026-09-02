/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiCreateSpecialKeyRequest struct {
	ctx            context.Context
	ApiService     *TradeAPIService
	apiName        *string
	symbol         *string
	ip             *string
	publicKey      *string
	permissionMode *models.CreateSpecialKeyPermissionModeParameter
	recvWindow     *int64
}

func (r ApiCreateSpecialKeyRequest) ApiName(apiName string) ApiCreateSpecialKeyRequest {
	r.apiName = &apiName
	return r
}

func (r ApiCreateSpecialKeyRequest) Symbol(symbol string) ApiCreateSpecialKeyRequest {
	r.symbol = &symbol
	return r
}

// Can be added in batches, separated by commas. Max 30 for an API key
func (r ApiCreateSpecialKeyRequest) Ip(ip string) ApiCreateSpecialKeyRequest {
	r.ip = &ip
	return r
}

// 1. If publicKey is inputted it will create an RSA or Ed25519 key.  2. Need to be encoded to URL-encoded format
func (r ApiCreateSpecialKeyRequest) PublicKey(publicKey string) ApiCreateSpecialKeyRequest {
	r.publicKey = &publicKey
	return r
}

// This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE.
func (r ApiCreateSpecialKeyRequest) PermissionMode(permissionMode models.CreateSpecialKeyPermissionModeParameter) ApiCreateSpecialKeyRequest {
	r.permissionMode = &permissionMode
	return r
}

func (r ApiCreateSpecialKeyRequest) RecvWindow(recvWindow int64) ApiCreateSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateSpecialKeyRequest) Execute() (*common.RestApiResponse[models.CreateSpecialKeyResponse], error) {
	return r.ApiService.CreateSpecialKeyExecute(r)
}

/*
CreateSpecialKey Create Special Key(Low-Latency Trading) (TRADE)
Post /sapi/v1/margin/apiKey

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#create-special-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param apiName -
@param symbol -
@param ip -  Can be added in batches, separated by commas. Max 30 for an API key
@param publicKey -  1. If publicKey is inputted it will create an RSA or Ed25519 key.  2. Need to be encoded to URL-encoded format
@param permissionMode -  This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE.
@param recvWindow -
@return ApiCreateSpecialKeyRequest
*/
func (a *TradeAPIService) CreateSpecialKey(ctx context.Context) ApiCreateSpecialKeyRequest {
	return ApiCreateSpecialKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSpecialKeyResponse
func (a *TradeAPIService) CreateSpecialKeyExecute(r ApiCreateSpecialKeyRequest) (*common.RestApiResponse[models.CreateSpecialKeyResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/apiKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.apiName == nil {
		return nil, common.ReportError("apiName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "apiName", r.apiName, "form", "")
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.ip != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "form", "")
	}
	if r.publicKey != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "publicKey", r.publicKey, "form", "")
	}
	if r.permissionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "permissionMode", r.permissionMode, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateSpecialKeyResponse](
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

type ApiDeleteSpecialKeyRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	apiName    *string
	symbol     *string
	recvWindow *int64
}

func (r ApiDeleteSpecialKeyRequest) ApiName(apiName string) ApiDeleteSpecialKeyRequest {
	r.apiName = &apiName
	return r
}

func (r ApiDeleteSpecialKeyRequest) Symbol(symbol string) ApiDeleteSpecialKeyRequest {
	r.symbol = &symbol
	return r
}

func (r ApiDeleteSpecialKeyRequest) RecvWindow(recvWindow int64) ApiDeleteSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteSpecialKeyRequest) Execute() (struct{}, error) {
	return r.ApiService.DeleteSpecialKeyExecute(r)
}

/*
DeleteSpecialKey Delete Special Key(Low-Latency Trading) (TRADE)
Delete /sapi/v1/margin/apiKey

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#delete-special-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param apiName -
@param symbol -
@param recvWindow -
@return ApiDeleteSpecialKeyRequest
*/
func (a *TradeAPIService) DeleteSpecialKey(ctx context.Context) ApiDeleteSpecialKeyRequest {
	return ApiDeleteSpecialKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TradeAPIService) DeleteSpecialKeyExecute(r ApiDeleteSpecialKeyRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/apiKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.apiName != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "apiName", r.apiName, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiEditIpForSpecialKeyRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	ip         *string
	symbol     *string
	recvWindow *int64
}

// Can be added in batches, separated by commas. Max 30 for an API key
func (r ApiEditIpForSpecialKeyRequest) Ip(ip string) ApiEditIpForSpecialKeyRequest {
	r.ip = &ip
	return r
}

// isolated margin pair
func (r ApiEditIpForSpecialKeyRequest) Symbol(symbol string) ApiEditIpForSpecialKeyRequest {
	r.symbol = &symbol
	return r
}

func (r ApiEditIpForSpecialKeyRequest) RecvWindow(recvWindow int64) ApiEditIpForSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEditIpForSpecialKeyRequest) Execute() (struct{}, error) {
	return r.ApiService.EditIpForSpecialKeyExecute(r)
}

/*
EditIpForSpecialKey Edit ip for Special Key(Low-Latency Trading) (TRADE)
Put /sapi/v1/margin/apiKey/ip

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#edit-ip-for-special-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param ip -  Can be added in batches, separated by commas. Max 30 for an API key
@param symbol -  isolated margin pair
@param recvWindow -
@return ApiEditIpForSpecialKeyRequest
*/
func (a *TradeAPIService) EditIpForSpecialKey(ctx context.Context) ApiEditIpForSpecialKeyRequest {
	return ApiEditIpForSpecialKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TradeAPIService) EditIpForSpecialKeyExecute(r ApiEditIpForSpecialKeyRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/apiKey/ip"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.ip == nil {
		return struct{}{}, common.ReportError("ip is required and must be specified")
	}

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiExitSpecialKeyModeRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;
func (r ApiExitSpecialKeyModeRequest) RecvWindow(recvWindow int64) ApiExitSpecialKeyModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiExitSpecialKeyModeRequest) Execute() (*common.RestApiResponse[map[string]interface{}], error) {
	return r.ApiService.ExitSpecialKeyModeExecute(r)
}

/*
ExitSpecialKeyMode Exit Special Key Mode (TRADE)
Post /sapi/v1/margin/exit-special-key-mode

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#exit-special-key-mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`
@return ApiExitSpecialKeyModeRequest
*/
func (a *TradeAPIService) ExitSpecialKeyMode(ctx context.Context) ApiExitSpecialKeyModeRequest {
	return ApiExitSpecialKeyModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *TradeAPIService) ExitSpecialKeyModeExecute(r ApiExitSpecialKeyModeRequest) (*common.RestApiResponse[map[string]interface{}], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/exit-special-key-mode"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[map[string]interface{}](
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

type ApiGetForceLiquidationRecordRequest struct {
	ctx            context.Context
	ApiService     *TradeAPIService
	startTime      *int64
	endTime        *int64
	isolatedSymbol *string
	current        *int64
	size           *int64
	recvWindow     *int64
}

func (r ApiGetForceLiquidationRecordRequest) StartTime(startTime int64) ApiGetForceLiquidationRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetForceLiquidationRecordRequest) EndTime(endTime int64) ApiGetForceLiquidationRecordRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetForceLiquidationRecordRequest) IsolatedSymbol(isolatedSymbol string) ApiGetForceLiquidationRecordRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

func (r ApiGetForceLiquidationRecordRequest) Current(current int64) ApiGetForceLiquidationRecordRequest {
	r.current = &current
	return r
}

func (r ApiGetForceLiquidationRecordRequest) Size(size int64) ApiGetForceLiquidationRecordRequest {
	r.size = &size
	return r
}

func (r ApiGetForceLiquidationRecordRequest) RecvWindow(recvWindow int64) ApiGetForceLiquidationRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetForceLiquidationRecordRequest) Execute() (*common.RestApiResponse[models.GetForceLiquidationRecordResponse], error) {
	return r.ApiService.GetForceLiquidationRecordExecute(r)
}

/*
GetForceLiquidationRecord Get Force Liquidation Record (USER_DATA)
Get /sapi/v1/margin/forceLiquidationRec

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#get-force-liquidation-record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param isolatedSymbol -
@param current -
@param size -
@param recvWindow -
@return ApiGetForceLiquidationRecordRequest
*/
func (a *TradeAPIService) GetForceLiquidationRecord(ctx context.Context) ApiGetForceLiquidationRecordRequest {
	return ApiGetForceLiquidationRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetForceLiquidationRecordResponse
func (a *TradeAPIService) GetForceLiquidationRecordExecute(r ApiGetForceLiquidationRecordRequest) (*common.RestApiResponse[models.GetForceLiquidationRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/forceLiquidationRec"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.isolatedSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isolatedSymbol", r.isolatedSymbol, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetForceLiquidationRecordResponse](
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

type ApiGetSmallLiabilityExchangeCoinListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiGetSmallLiabilityExchangeCoinListRequest) RecvWindow(recvWindow int64) ApiGetSmallLiabilityExchangeCoinListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSmallLiabilityExchangeCoinListRequest) Execute() (*common.RestApiResponse[models.GetSmallLiabilityExchangeCoinListResponse], error) {
	return r.ApiService.GetSmallLiabilityExchangeCoinListExecute(r)
}

/*
GetSmallLiabilityExchangeCoinList Get Small Liability Exchange Coin List (USER_DATA)
Get /sapi/v1/margin/exchange-small-liability

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#get-small-liability-exchange-coin-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetSmallLiabilityExchangeCoinListRequest
*/
func (a *TradeAPIService) GetSmallLiabilityExchangeCoinList(ctx context.Context) ApiGetSmallLiabilityExchangeCoinListRequest {
	return ApiGetSmallLiabilityExchangeCoinListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSmallLiabilityExchangeCoinListResponse
func (a *TradeAPIService) GetSmallLiabilityExchangeCoinListExecute(r ApiGetSmallLiabilityExchangeCoinListRequest) (*common.RestApiResponse[models.GetSmallLiabilityExchangeCoinListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/exchange-small-liability"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSmallLiabilityExchangeCoinListResponse](
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

type ApiGetSmallLiabilityExchangeHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	current    *int64
	size       *int64
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) Current(current int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.current = &current
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) Size(size int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) StartTime(startTime int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) EndTime(endTime int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) RecvWindow(recvWindow int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) Execute() (*common.RestApiResponse[models.GetSmallLiabilityExchangeHistoryResponse], error) {
	return r.ApiService.GetSmallLiabilityExchangeHistoryExecute(r)
}

/*
GetSmallLiabilityExchangeHistory Get Small Liability Exchange History (USER_DATA)
Get /sapi/v1/margin/exchange-small-liability-history

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#get-small-liability-exchange-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param current -
@param size -
@param startTime -
@param endTime -
@param recvWindow -
@return ApiGetSmallLiabilityExchangeHistoryRequest
*/
func (a *TradeAPIService) GetSmallLiabilityExchangeHistory(ctx context.Context) ApiGetSmallLiabilityExchangeHistoryRequest {
	return ApiGetSmallLiabilityExchangeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSmallLiabilityExchangeHistoryResponse
func (a *TradeAPIService) GetSmallLiabilityExchangeHistoryExecute(r ApiGetSmallLiabilityExchangeHistoryRequest) (*common.RestApiResponse[models.GetSmallLiabilityExchangeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/exchange-small-liability-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.current == nil {
		return nil, common.ReportError("current is required and must be specified")
	}
	if *r.current < 1 {
		return nil, common.ReportError("current must be greater than 1")
	}

	if r.size == nil {
		return nil, common.ReportError("size is required and must be specified")
	}
	if *r.size > 100 {
		return nil, common.ReportError("size must be less than 100")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSmallLiabilityExchangeHistoryResponse](
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

type ApiLiquidationLoanRepayRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	asset      *string
	amount     *float64
	recvWindow *int64
}

// The asset to repay (e.g. USDT, USDC)
func (r ApiLiquidationLoanRepayRequest) Asset(asset string) ApiLiquidationLoanRepayRequest {
	r.asset = &asset
	return r
}

// Repayment amount, must be greater than 0
func (r ApiLiquidationLoanRepayRequest) Amount(amount float64) ApiLiquidationLoanRepayRequest {
	r.amount = &amount
	return r
}

func (r ApiLiquidationLoanRepayRequest) RecvWindow(recvWindow int64) ApiLiquidationLoanRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiLiquidationLoanRepayRequest) Execute() (*common.RestApiResponse[models.LiquidationLoanRepayResponse], error) {
	return r.ApiService.LiquidationLoanRepayExecute(r)
}

/*
LiquidationLoanRepay Liquidation Loan Repay (MARGIN)
Post /sapi/v1/margin/liquidation-loan/repay

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#liquidation-loan-repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  The asset to repay (e.g. USDT, USDC)
@param amount -  Repayment amount, must be greater than 0
@param recvWindow -
@return ApiLiquidationLoanRepayRequest
*/
func (a *TradeAPIService) LiquidationLoanRepay(ctx context.Context) ApiLiquidationLoanRepayRequest {
	return ApiLiquidationLoanRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LiquidationLoanRepayResponse
func (a *TradeAPIService) LiquidationLoanRepayExecute(r ApiLiquidationLoanRepayRequest) (*common.RestApiResponse[models.LiquidationLoanRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/liquidation-loan/repay"

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

	resp, err := SendRequest[models.LiquidationLoanRepayResponse](
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

type ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	recvWindow *int64
}

func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) Symbol(symbol string) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) RecvWindow(recvWindow int64) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) Execute() (*common.RestApiResponse[models.MarginAccountCancelAllOpenOrdersOnASymbolResponse], error) {
	return r.ApiService.MarginAccountCancelAllOpenOrdersOnASymbolExecute(r)
}

/*
MarginAccountCancelAllOpenOrdersOnASymbol Margin Account Cancel all Open Orders on a Symbol (TRADE)
Delete /sapi/v1/margin/openOrders

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-cancel-all-open-orders-on-asymbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param recvWindow -
@return ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest
*/
func (a *TradeAPIService) MarginAccountCancelAllOpenOrdersOnASymbol(ctx context.Context) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	return ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountCancelAllOpenOrdersOnASymbolResponse
func (a *TradeAPIService) MarginAccountCancelAllOpenOrdersOnASymbolExecute(r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) (*common.RestApiResponse[models.MarginAccountCancelAllOpenOrdersOnASymbolResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountCancelAllOpenOrdersOnASymbolResponse](
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

type ApiMarginAccountCancelOcoRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	orderListId       *int64
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiMarginAccountCancelOcoRequest) Symbol(symbol string) ApiMarginAccountCancelOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountCancelOcoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountCancelOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountCancelOcoRequest) OrderListId(orderListId int64) ApiMarginAccountCancelOcoRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiMarginAccountCancelOcoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountCancelOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiMarginAccountCancelOcoRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountCancelOcoRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiMarginAccountCancelOcoRequest) RecvWindow(recvWindow int64) ApiMarginAccountCancelOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountCancelOcoRequest) Execute() (*common.RestApiResponse[models.MarginAccountCancelOcoResponse], error) {
	return r.ApiService.MarginAccountCancelOcoExecute(r)
}

/*
MarginAccountCancelOco Margin Account Cancel OCO (TRADE)
Delete /sapi/v1/margin/orderList

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-cancel-oco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param orderListId -
@param listClientOrderId -
@param newClientOrderId -
@param recvWindow -
@return ApiMarginAccountCancelOcoRequest
*/
func (a *TradeAPIService) MarginAccountCancelOco(ctx context.Context) ApiMarginAccountCancelOcoRequest {
	return ApiMarginAccountCancelOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountCancelOcoResponse
func (a *TradeAPIService) MarginAccountCancelOcoExecute(r ApiMarginAccountCancelOcoRequest) (*common.RestApiResponse[models.MarginAccountCancelOcoResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountCancelOcoResponse](
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

type ApiMarginAccountCancelOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiMarginAccountCancelOrderRequest) Symbol(symbol string) ApiMarginAccountCancelOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountCancelOrderRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountCancelOrderRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountCancelOrderRequest) OrderId(orderId int64) ApiMarginAccountCancelOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiMarginAccountCancelOrderRequest) OrigClientOrderId(origClientOrderId string) ApiMarginAccountCancelOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiMarginAccountCancelOrderRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountCancelOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

func (r ApiMarginAccountCancelOrderRequest) RecvWindow(recvWindow int64) ApiMarginAccountCancelOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountCancelOrderRequest) Execute() (*common.RestApiResponse[models.MarginAccountCancelOrderResponse], error) {
	return r.ApiService.MarginAccountCancelOrderExecute(r)
}

/*
MarginAccountCancelOrder Margin Account Cancel Order (TRADE)
Delete /sapi/v1/margin/order

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-cancel-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param orderId -
@param origClientOrderId -
@param newClientOrderId -
@param recvWindow -
@return ApiMarginAccountCancelOrderRequest
*/
func (a *TradeAPIService) MarginAccountCancelOrder(ctx context.Context) ApiMarginAccountCancelOrderRequest {
	return ApiMarginAccountCancelOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountCancelOrderResponse
func (a *TradeAPIService) MarginAccountCancelOrderExecute(r ApiMarginAccountCancelOrderRequest) (*common.RestApiResponse[models.MarginAccountCancelOrderResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountCancelOrderResponse](
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

type ApiMarginAccountNewOcoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.MarginAccountNewOrderSideParameter
	quantity                *float64
	price                   *float64
	stopPrice               *float64
	isIsolated              *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	listClientOrderId       *string
	limitClientOrderId      *string
	limitIcebergQty         *float64
	stopClientOrderId       *string
	stopLimitPrice          *float64
	stopIcebergQty          *float64
	stopLimitTimeInForce    *models.MarginAccountNewOcoStopLimitTimeInForceParameter
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *models.MarginAccountNewOrderSideEffectTypeParameter
	selfTradePreventionMode *models.MarginAccountNewOrderSelfTradePreventionModeParameter
	autoRepayAtCancel       *bool
	recvWindow              *int64
}

func (r ApiMarginAccountNewOcoRequest) Symbol(symbol string) ApiMarginAccountNewOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountNewOcoRequest) Side(side models.MarginAccountNewOrderSideParameter) ApiMarginAccountNewOcoRequest {
	r.side = &side
	return r
}

func (r ApiMarginAccountNewOcoRequest) Quantity(quantity float64) ApiMarginAccountNewOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiMarginAccountNewOcoRequest) Price(price float64) ApiMarginAccountNewOcoRequest {
	r.price = &price
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopPrice(stopPrice float64) ApiMarginAccountNewOcoRequest {
	r.stopPrice = &stopPrice
	return r
}

func (r ApiMarginAccountNewOcoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountNewOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// A unique Id for the entire orderList
func (r ApiMarginAccountNewOcoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique Id for the limit order
func (r ApiMarginAccountNewOcoRequest) LimitClientOrderId(limitClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

func (r ApiMarginAccountNewOcoRequest) LimitIcebergQty(limitIcebergQty float64) ApiMarginAccountNewOcoRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

// A unique Id for the stop loss/stop loss limit leg
func (r ApiMarginAccountNewOcoRequest) StopClientOrderId(stopClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

// If provided, &#x60;stopLimitTimeInForce&#x60; is required.
func (r ApiMarginAccountNewOcoRequest) StopLimitPrice(stopLimitPrice float64) ApiMarginAccountNewOcoRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopIcebergQty(stopIcebergQty float64) ApiMarginAccountNewOcoRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopLimitTimeInForce(stopLimitTimeInForce models.MarginAccountNewOcoStopLimitTimeInForceParameter) ApiMarginAccountNewOcoRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

func (r ApiMarginAccountNewOcoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiMarginAccountNewOcoRequest) SideEffectType(sideEffectType models.MarginAccountNewOrderSideEffectTypeParameter) ApiMarginAccountNewOcoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

func (r ApiMarginAccountNewOcoRequest) SelfTradePreventionMode(selfTradePreventionMode models.MarginAccountNewOrderSelfTradePreventionModeParameter) ApiMarginAccountNewOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled.
func (r ApiMarginAccountNewOcoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOcoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

func (r ApiMarginAccountNewOcoRequest) RecvWindow(recvWindow int64) ApiMarginAccountNewOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountNewOcoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOcoResponse], error) {
	return r.ApiService.MarginAccountNewOcoExecute(r)
}

/*
MarginAccountNewOco Margin Account New OCO (TRADE)
Post /sapi/v1/margin/order/oco

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-new-oco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -
@param price -
@param stopPrice -
@param isIsolated -
@param listClientOrderId -  A unique Id for the entire orderList
@param limitClientOrderId -  A unique Id for the limit order
@param limitIcebergQty -
@param stopClientOrderId -  A unique Id for the stop loss/stop loss limit leg
@param stopLimitPrice -  If provided, `stopLimitTimeInForce` is required.
@param stopIcebergQty -
@param stopLimitTimeInForce -
@param newOrderRespType -
@param sideEffectType -
@param selfTradePreventionMode -
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled.
@param recvWindow -
@return ApiMarginAccountNewOcoRequest
*/
func (a *TradeAPIService) MarginAccountNewOco(ctx context.Context) ApiMarginAccountNewOcoRequest {
	return ApiMarginAccountNewOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountNewOcoResponse
func (a *TradeAPIService) MarginAccountNewOcoExecute(r ApiMarginAccountNewOcoRequest) (*common.RestApiResponse[models.MarginAccountNewOcoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order/oco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.quantity == nil {
		return nil, common.ReportError("quantity is required and must be specified")
	}

	if r.price == nil {
		return nil, common.ReportError("price is required and must be specified")
	}

	if r.stopPrice == nil {
		return nil, common.ReportError("stopPrice is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	if r.limitClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitClientOrderId", r.limitClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	if r.limitIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limitIcebergQty", r.limitIcebergQty, "form", "")
	}
	if r.stopClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopClientOrderId", r.stopClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	if r.stopLimitPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitPrice", r.stopLimitPrice, "form", "")
	}
	if r.stopIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopIcebergQty", r.stopIcebergQty, "form", "")
	}
	if r.stopLimitTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopLimitTimeInForce", r.stopLimitTimeInForce, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOcoResponse](
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

type ApiMarginAccountNewOrderRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	side                    *models.MarginAccountNewOrderSideParameter
	type_                   *models.MarginAccountNewOrderTypeParameter
	isIsolated              *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	quantity                *float64
	quoteOrderQty           *float64
	price                   *float64
	stopPrice               *float64
	newClientOrderId        *string
	icebergQty              *float64
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *models.MarginAccountNewOrderSideEffectTypeParameter
	timeInForce             *models.MarginAccountNewOrderTimeInForceParameter
	selfTradePreventionMode *models.MarginAccountNewOrderSelfTradePreventionModeParameter
	trailingDelta           *int64
	autoRepayAtCancel       *bool
	recvWindow              *int64
}

func (r ApiMarginAccountNewOrderRequest) Symbol(symbol string) ApiMarginAccountNewOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountNewOrderRequest) Side(side models.MarginAccountNewOrderSideParameter) ApiMarginAccountNewOrderRequest {
	r.side = &side
	return r
}

func (r ApiMarginAccountNewOrderRequest) Type(type_ models.MarginAccountNewOrderTypeParameter) ApiMarginAccountNewOrderRequest {
	r.type_ = &type_
	return r
}

func (r ApiMarginAccountNewOrderRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountNewOrderRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountNewOrderRequest) Quantity(quantity float64) ApiMarginAccountNewOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiMarginAccountNewOrderRequest) QuoteOrderQty(quoteOrderQty float64) ApiMarginAccountNewOrderRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

func (r ApiMarginAccountNewOrderRequest) Price(price float64) ApiMarginAccountNewOrderRequest {
	r.price = &price
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiMarginAccountNewOrderRequest) StopPrice(stopPrice float64) ApiMarginAccountNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// A unique id among open orders. Automatically generated if not sent.
func (r ApiMarginAccountNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiMarginAccountNewOrderRequest) IcebergQty(icebergQty float64) ApiMarginAccountNewOrderRequest {
	r.icebergQty = &icebergQty
	return r
}

// MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOrderRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiMarginAccountNewOrderRequest) SideEffectType(sideEffectType models.MarginAccountNewOrderSideEffectTypeParameter) ApiMarginAccountNewOrderRequest {
	r.sideEffectType = &sideEffectType
	return r
}

func (r ApiMarginAccountNewOrderRequest) TimeInForce(timeInForce models.MarginAccountNewOrderTimeInForceParameter) ApiMarginAccountNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

func (r ApiMarginAccountNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode models.MarginAccountNewOrderSelfTradePreventionModeParameter) ApiMarginAccountNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiMarginAccountNewOrderRequest) TrailingDelta(trailingDelta int64) ApiMarginAccountNewOrderRequest {
	r.trailingDelta = &trailingDelta
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
func (r ApiMarginAccountNewOrderRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOrderRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

func (r ApiMarginAccountNewOrderRequest) RecvWindow(recvWindow int64) ApiMarginAccountNewOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginAccountNewOrderRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOrderResponse], error) {
	return r.ApiService.MarginAccountNewOrderExecute(r)
}

/*
MarginAccountNewOrder Margin Account New Order (TRADE)
Post /sapi/v1/margin/order

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-new-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -
@param isIsolated -
@param quantity -
@param quoteOrderQty -
@param price -
@param stopPrice -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param newClientOrderId -  A unique id among open orders. Automatically generated if not sent.
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -  MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param sideEffectType -
@param timeInForce -
@param selfTradePreventionMode -
@param trailingDelta -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
@param recvWindow -
@return ApiMarginAccountNewOrderRequest
*/
func (a *TradeAPIService) MarginAccountNewOrder(ctx context.Context) ApiMarginAccountNewOrderRequest {
	return ApiMarginAccountNewOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountNewOrderResponse
func (a *TradeAPIService) MarginAccountNewOrderExecute(r ApiMarginAccountNewOrderRequest) (*common.RestApiResponse[models.MarginAccountNewOrderResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.quantity != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quantity", r.quantity, "form", "")
	}
	if r.quoteOrderQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteOrderQty", r.quoteOrderQty, "form", "")
	}
	if r.price != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "price", r.price, "form", "")
	}
	if r.stopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "stopPrice", r.stopPrice, "form", "")
	}
	if r.newClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newClientOrderId", r.newClientOrderId, "form", "")
	}
	if r.icebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "icebergQty", r.icebergQty, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.timeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeInForce", r.timeInForce, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.trailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trailingDelta", r.trailingDelta, "form", "")
	}
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOrderResponse](
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

type ApiMarginAccountNewOtoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *models.MarginAccountNewOtoWorkingTypeParameter
	workingSide             *models.MarginAccountNewOtoWorkingSideParameter
	workingPrice            *float64
	workingQuantity         *float64
	workingIcebergQty       *float64
	pendingType             *models.MarginAccountNewOrderTypeParameter
	pendingSide             *models.MarginAccountNewOrderSideParameter
	pendingQuantity         *float64
	isIsolated              *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	listClientOrderId       *string
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *models.MarginAccountNewOtoSideEffectTypeParameter
	selfTradePreventionMode *models.MarginAccountNewOrderSelfTradePreventionModeParameter
	autoRepayAtCancel       *bool
	workingClientOrderId    *string
	workingTimeInForce      *models.MarginAccountNewOrderTimeInForceParameter
	pendingClientOrderId    *string
	pendingPrice            *float64
	pendingStopPrice        *float64
	pendingTrailingDelta    *float64
	pendingIcebergQty       *float64
	pendingTimeInForce      *models.MarginAccountNewOrderTimeInForceParameter
}

func (r ApiMarginAccountNewOtoRequest) Symbol(symbol string) ApiMarginAccountNewOtoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingType(workingType models.MarginAccountNewOtoWorkingTypeParameter) ApiMarginAccountNewOtoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingSide(workingSide models.MarginAccountNewOtoWorkingSideParameter) ApiMarginAccountNewOtoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingPrice(workingPrice float64) ApiMarginAccountNewOtoRequest {
	r.workingPrice = &workingPrice
	return r
}

// Sets the quantity for the working order.
func (r ApiMarginAccountNewOtoRequest) WorkingQuantity(workingQuantity float64) ApiMarginAccountNewOtoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiMarginAccountNewOtoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingType(pendingType models.MarginAccountNewOrderTypeParameter) ApiMarginAccountNewOtoRequest {
	r.pendingType = &pendingType
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingSide(pendingSide models.MarginAccountNewOrderSideParameter) ApiMarginAccountNewOtoRequest {
	r.pendingSide = &pendingSide
	return r
}

// Sets the quantity for the pending order.
func (r ApiMarginAccountNewOtoRequest) PendingQuantity(pendingQuantity float64) ApiMarginAccountNewOtoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

func (r ApiMarginAccountNewOtoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountNewOtoRequest {
	r.isIsolated = &isIsolated
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent.&lt;br/&gt;A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.&lt;br/&gt;&#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60; and the &#x60;pendingClientOrderId&#x60;.
func (r ApiMarginAccountNewOtoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOtoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOtoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiMarginAccountNewOtoRequest) SideEffectType(sideEffectType models.MarginAccountNewOtoSideEffectTypeParameter) ApiMarginAccountNewOtoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

func (r ApiMarginAccountNewOtoRequest) SelfTradePreventionMode(selfTradePreventionMode models.MarginAccountNewOrderSelfTradePreventionModeParameter) ApiMarginAccountNewOtoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
func (r ApiMarginAccountNewOtoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOtoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtoRequest) WorkingClientOrderId(workingClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingTimeInForce(workingTimeInForce models.MarginAccountNewOrderTimeInForceParameter) ApiMarginAccountNewOtoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtoRequest) PendingClientOrderId(pendingClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingPrice(pendingPrice float64) ApiMarginAccountNewOtoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingStopPrice(pendingStopPrice float64) ApiMarginAccountNewOtoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingTrailingDelta(pendingTrailingDelta float64) ApiMarginAccountNewOtoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtoRequest) PendingIcebergQty(pendingIcebergQty float64) ApiMarginAccountNewOtoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingTimeInForce(pendingTimeInForce models.MarginAccountNewOrderTimeInForceParameter) ApiMarginAccountNewOtoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

func (r ApiMarginAccountNewOtoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOtoResponse], error) {
	return r.ApiService.MarginAccountNewOtoExecute(r)
}

/*
MarginAccountNewOto Margin Account New OTO (TRADE)
Post /sapi/v1/margin/order/oto

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-new-oto

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -  Sets the quantity for the working order.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`.
@param pendingType -
@param pendingSide -
@param pendingQuantity -  Sets the quantity for the pending order.
@param isIsolated -
@param listClientOrderId -  Arbitrary unique ID among open order lists. Automatically generated if not sent.<br/>A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.<br/>`listClientOrderId` is distinct from the `workingClientOrderId` and the `pendingClientOrderId`.
@param newOrderRespType -  MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param sideEffectType -
@param selfTradePreventionMode -
@param autoRepayAtCancel -  Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingTimeInForce -
@param pendingClientOrderId -  Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
@param pendingPrice -
@param pendingStopPrice -
@param pendingTrailingDelta -
@param pendingIcebergQty -  This can only be used if `pendingTimeInForce` is `GTC`.
@param pendingTimeInForce -
@return ApiMarginAccountNewOtoRequest
*/
func (a *TradeAPIService) MarginAccountNewOto(ctx context.Context) ApiMarginAccountNewOtoRequest {
	return ApiMarginAccountNewOtoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountNewOtoResponse
func (a *TradeAPIService) MarginAccountNewOtoExecute(r ApiMarginAccountNewOtoRequest) (*common.RestApiResponse[models.MarginAccountNewOtoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order/oto"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}

	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}

	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}

	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}

	if r.workingIcebergQty == nil {
		return nil, common.ReportError("workingIcebergQty is required and must be specified")
	}

	if r.pendingType == nil {
		return nil, common.ReportError("pendingType is required and must be specified")
	}

	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}

	if r.pendingQuantity == nil {
		return nil, common.ReportError("pendingQuantity is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingType", r.pendingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	if r.pendingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingClientOrderId", r.pendingClientOrderId, "form", "")
	}
	if r.pendingPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingPrice", r.pendingPrice, "form", "")
	}
	if r.pendingStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingStopPrice", r.pendingStopPrice, "form", "")
	}
	if r.pendingTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTrailingDelta", r.pendingTrailingDelta, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingQuantity", r.pendingQuantity, "form", "")
	if r.pendingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingIcebergQty", r.pendingIcebergQty, "form", "")
	}
	if r.pendingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingTimeInForce", r.pendingTimeInForce, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOtoResponse](
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

type ApiMarginAccountNewOtocoRequest struct {
	ctx                       context.Context
	ApiService                *TradeAPIService
	symbol                    *string
	workingType               *models.MarginAccountNewOtoWorkingTypeParameter
	workingSide               *models.MarginAccountNewOtoWorkingSideParameter
	workingPrice              *float64
	workingQuantity           *float64
	pendingSide               *models.MarginAccountNewOrderSideParameter
	pendingQuantity           *float64
	pendingAboveType          *models.MarginAccountNewOtocoPendingAboveTypeParameter
	isIsolated                *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	sideEffectType            *models.MarginAccountNewOtoSideEffectTypeParameter
	autoRepayAtCancel         *bool
	listClientOrderId         *string
	newOrderRespType          *models.MarginAccountNewOrderNewOrderRespTypeParameter
	selfTradePreventionMode   *models.MarginAccountNewOrderSelfTradePreventionModeParameter
	workingClientOrderId      *string
	workingIcebergQty         *float64
	workingTimeInForce        *models.MarginAccountNewOtocoWorkingTimeInForceParameter
	pendingAboveClientOrderId *string
	pendingAbovePrice         *float64
	pendingAboveStopPrice     *float64
	pendingAboveTrailingDelta *float64
	pendingAboveIcebergQty    *float64
	pendingAboveTimeInForce   *models.MarginAccountNewOtocoWorkingTimeInForceParameter
	pendingBelowType          *models.MarginAccountNewOtocoPendingAboveTypeParameter
	pendingBelowClientOrderId *string
	pendingBelowPrice         *float64
	pendingBelowStopPrice     *float64
	pendingBelowTrailingDelta *float64
	pendingBelowIcebergQty    *float64
	pendingBelowTimeInForce   *models.MarginAccountNewOtocoWorkingTimeInForceParameter
}

func (r ApiMarginAccountNewOtocoRequest) Symbol(symbol string) ApiMarginAccountNewOtocoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingType(workingType models.MarginAccountNewOtoWorkingTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.workingType = &workingType
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingSide(workingSide models.MarginAccountNewOtoWorkingSideParameter) ApiMarginAccountNewOtocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingPrice(workingPrice float64) ApiMarginAccountNewOtocoRequest {
	r.workingPrice = &workingPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingQuantity(workingQuantity float64) ApiMarginAccountNewOtocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingSide(pendingSide models.MarginAccountNewOrderSideParameter) ApiMarginAccountNewOtocoRequest {
	r.pendingSide = &pendingSide
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingQuantity(pendingQuantity float64) ApiMarginAccountNewOtocoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveType(pendingAboveType models.MarginAccountNewOtocoPendingAboveTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

func (r ApiMarginAccountNewOtocoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiMarginAccountNewOtocoRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountNewOtocoRequest) SideEffectType(sideEffectType models.MarginAccountNewOtoSideEffectTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

// Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
func (r ApiMarginAccountNewOtocoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOtocoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired. &#x60;listClientOrderId&#x60; is distinct from the &#x60;workingClientOrderId&#x60;, &#x60;pendingAboveClientOrderId&#x60;, and the &#x60;pendingBelowClientOrderId&#x60;.
func (r ApiMarginAccountNewOtocoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

func (r ApiMarginAccountNewOtocoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

func (r ApiMarginAccountNewOtocoRequest) SelfTradePreventionMode(selfTradePreventionMode models.MarginAccountNewOrderSelfTradePreventionModeParameter) ApiMarginAccountNewOtocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) WorkingIcebergQty(workingIcebergQty float64) ApiMarginAccountNewOtocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingTimeInForce(workingTimeInForce models.MarginAccountNewOtocoWorkingTimeInForceParameter) ApiMarginAccountNewOtocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAbovePrice(pendingAbovePrice float64) ApiMarginAccountNewOtocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float64) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float64) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float64) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce models.MarginAccountNewOtocoWorkingTimeInForceParameter) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowType(pendingBelowType models.MarginAccountNewOtocoPendingAboveTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowPrice(pendingBelowPrice float64) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float64) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float64) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float64) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce models.MarginAccountNewOtocoWorkingTimeInForceParameter) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

func (r ApiMarginAccountNewOtocoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOtocoResponse], error) {
	return r.ApiService.MarginAccountNewOtocoExecute(r)
}

/*
MarginAccountNewOtoco Margin Account New OTOCO (TRADE)
Post /sapi/v1/margin/order/otoco

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-account-new-otoco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -
@param workingSide -
@param workingPrice -
@param workingQuantity -
@param pendingSide -
@param pendingQuantity -
@param pendingAboveType -
@param isIsolated -
@param sideEffectType -
@param autoRepayAtCancel -  Only when MARGIN_BUY order takes effect, true means that the debt generated by the order needs to be repaid after the order is cancelled.
@param listClientOrderId -  Arbitrary unique ID among open order lists. Automatically generated if not sent. A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired. `listClientOrderId` is distinct from the `workingClientOrderId`, `pendingAboveClientOrderId`, and the `pendingBelowClientOrderId`.
@param newOrderRespType -
@param selfTradePreventionMode -
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`.
@param workingTimeInForce -
@param pendingAboveClientOrderId -  Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
@param pendingAbovePrice -
@param pendingAboveStopPrice -
@param pendingAboveTrailingDelta -
@param pendingAboveIcebergQty -  This can only be used if `pendingAboveTimeInForce` is `GTC`.
@param pendingAboveTimeInForce -
@param pendingBelowType -
@param pendingBelowClientOrderId -  Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
@param pendingBelowPrice -
@param pendingBelowStopPrice -
@param pendingBelowTrailingDelta -
@param pendingBelowIcebergQty -  This can only be used if `pendingBelowTimeInForce` is `GTC`.
@param pendingBelowTimeInForce -
@return ApiMarginAccountNewOtocoRequest
*/
func (a *TradeAPIService) MarginAccountNewOtoco(ctx context.Context) ApiMarginAccountNewOtocoRequest {
	return ApiMarginAccountNewOtocoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginAccountNewOtocoResponse
func (a *TradeAPIService) MarginAccountNewOtocoExecute(r ApiMarginAccountNewOtocoRequest) (*common.RestApiResponse[models.MarginAccountNewOtocoResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order/otoco"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	if r.workingType == nil {
		return nil, common.ReportError("workingType is required and must be specified")
	}

	if r.workingSide == nil {
		return nil, common.ReportError("workingSide is required and must be specified")
	}

	if r.workingPrice == nil {
		return nil, common.ReportError("workingPrice is required and must be specified")
	}

	if r.workingQuantity == nil {
		return nil, common.ReportError("workingQuantity is required and must be specified")
	}

	if r.pendingSide == nil {
		return nil, common.ReportError("pendingSide is required and must be specified")
	}

	if r.pendingQuantity == nil {
		return nil, common.ReportError("pendingQuantity is required and must be specified")
	}

	if r.pendingAboveType == nil {
		return nil, common.ReportError("pendingAboveType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.sideEffectType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sideEffectType", r.sideEffectType, "form", "")
	}
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	if r.listClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "listClientOrderId", r.listClientOrderId, "form", "")
	}
	if r.newOrderRespType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "newOrderRespType", r.newOrderRespType, "form", "")
	}
	if r.selfTradePreventionMode != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "selfTradePreventionMode", r.selfTradePreventionMode, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingType", r.workingType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingSide", r.workingSide, "form", "")
	if r.workingClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingClientOrderId", r.workingClientOrderId, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingPrice", r.workingPrice, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingQuantity", r.workingQuantity, "form", "")
	if r.workingIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingIcebergQty", r.workingIcebergQty, "form", "")
	}
	if r.workingTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workingTimeInForce", r.workingTimeInForce, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingSide", r.pendingSide, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingQuantity", r.pendingQuantity, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveType", r.pendingAboveType, "form", "")
	if r.pendingAboveClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveClientOrderId", r.pendingAboveClientOrderId, "form", "")
	}
	if r.pendingAbovePrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAbovePrice", r.pendingAbovePrice, "form", "")
	}
	if r.pendingAboveStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveStopPrice", r.pendingAboveStopPrice, "form", "")
	}
	if r.pendingAboveTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTrailingDelta", r.pendingAboveTrailingDelta, "form", "")
	}
	if r.pendingAboveIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveIcebergQty", r.pendingAboveIcebergQty, "form", "")
	}
	if r.pendingAboveTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingAboveTimeInForce", r.pendingAboveTimeInForce, "form", "")
	}
	if r.pendingBelowType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowType", r.pendingBelowType, "form", "")
	}
	if r.pendingBelowClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowClientOrderId", r.pendingBelowClientOrderId, "form", "")
	}
	if r.pendingBelowPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowPrice", r.pendingBelowPrice, "form", "")
	}
	if r.pendingBelowStopPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowStopPrice", r.pendingBelowStopPrice, "form", "")
	}
	if r.pendingBelowTrailingDelta != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTrailingDelta", r.pendingBelowTrailingDelta, "form", "")
	}
	if r.pendingBelowIcebergQty != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowIcebergQty", r.pendingBelowIcebergQty, "form", "")
	}
	if r.pendingBelowTimeInForce != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingBelowTimeInForce", r.pendingBelowTimeInForce, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOtocoResponse](
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

type ApiMarginManualLiquidationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	type_      *models.QueryMarginAvailableInventoryTypeParameter
	symbol     *string
	recvWindow *int64
}

func (r ApiMarginManualLiquidationRequest) Type(type_ models.QueryMarginAvailableInventoryTypeParameter) ApiMarginManualLiquidationRequest {
	r.type_ = &type_
	return r
}

// When type selects &#x60;ISOLATED&#x60;, &#x60;symbol&#x60; must be filled in
func (r ApiMarginManualLiquidationRequest) Symbol(symbol string) ApiMarginManualLiquidationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarginManualLiquidationRequest) RecvWindow(recvWindow int64) ApiMarginManualLiquidationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginManualLiquidationRequest) Execute() (*common.RestApiResponse[models.MarginManualLiquidationResponse], error) {
	return r.ApiService.MarginManualLiquidationExecute(r)
}

/*
MarginManualLiquidation Margin Manual Liquidation (TRADE)
Post /sapi/v1/margin/manual-liquidation

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#margin-manual-liquidation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -
@param symbol -  When type selects `ISOLATED`, `symbol` must be filled in
@param recvWindow -
@return ApiMarginManualLiquidationRequest
*/
func (a *TradeAPIService) MarginManualLiquidation(ctx context.Context) ApiMarginManualLiquidationRequest {
	return ApiMarginManualLiquidationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginManualLiquidationResponse
func (a *TradeAPIService) MarginManualLiquidationExecute(r ApiMarginManualLiquidationRequest) (*common.RestApiResponse[models.MarginManualLiquidationResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/manual-liquidation"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginManualLiquidationResponse](
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

type ApiQueryCurrentMarginOrderCountUsageRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryCurrentMarginOrderCountUsageRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryCurrentMarginOrderCountUsageRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryCurrentMarginOrderCountUsageRequest) Symbol(symbol string) ApiQueryCurrentMarginOrderCountUsageRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryCurrentMarginOrderCountUsageRequest) RecvWindow(recvWindow int64) ApiQueryCurrentMarginOrderCountUsageRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCurrentMarginOrderCountUsageRequest) Execute() (*common.RestApiResponse[models.QueryCurrentMarginOrderCountUsageResponse], error) {
	return r.ApiService.QueryCurrentMarginOrderCountUsageExecute(r)
}

/*
QueryCurrentMarginOrderCountUsage Query Current Margin Order Count Usage (TRADE)
Get /sapi/v1/margin/rateLimit/order

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-current-margin-order-count-usage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -
@param symbol -
@param recvWindow -
@return ApiQueryCurrentMarginOrderCountUsageRequest
*/
func (a *TradeAPIService) QueryCurrentMarginOrderCountUsage(ctx context.Context) ApiQueryCurrentMarginOrderCountUsageRequest {
	return ApiQueryCurrentMarginOrderCountUsageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCurrentMarginOrderCountUsageResponse
func (a *TradeAPIService) QueryCurrentMarginOrderCountUsageExecute(r ApiQueryCurrentMarginOrderCountUsageRequest) (*common.RestApiResponse[models.QueryCurrentMarginOrderCountUsageResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/rateLimit/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCurrentMarginOrderCountUsageResponse](
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

type ApiQueryLiquidationLoanRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

func (r ApiQueryLiquidationLoanRequest) RecvWindow(recvWindow int64) ApiQueryLiquidationLoanRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryLiquidationLoanRequest) Execute() (*common.RestApiResponse[models.QueryLiquidationLoanResponse], error) {
	return r.ApiService.QueryLiquidationLoanExecute(r)
}

/*
QueryLiquidationLoan Query Liquidation Loan (USER_DATA)
Get /sapi/v1/margin/liquidation-loan

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-liquidation-loan

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryLiquidationLoanRequest
*/
func (a *TradeAPIService) QueryLiquidationLoan(ctx context.Context) ApiQueryLiquidationLoanRequest {
	return ApiQueryLiquidationLoanRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryLiquidationLoanResponse
func (a *TradeAPIService) QueryLiquidationLoanExecute(r ApiQueryLiquidationLoanRequest) (*common.RestApiResponse[models.QueryLiquidationLoanResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/liquidation-loan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryLiquidationLoanResponse](
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

type ApiQueryLiquidationLoanRepayHistoryRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// Start time in Unix timestamp (milliseconds). Defaults to 7 days ago if not specified
func (r ApiQueryLiquidationLoanRepayHistoryRequest) StartTime(startTime int64) ApiQueryLiquidationLoanRepayHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time in Unix timestamp (milliseconds). Defaults to now if not specified
func (r ApiQueryLiquidationLoanRepayHistoryRequest) EndTime(endTime int64) ApiQueryLiquidationLoanRepayHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current page number, default &#x60;1&#x60;
func (r ApiQueryLiquidationLoanRepayHistoryRequest) Current(current int64) ApiQueryLiquidationLoanRepayHistoryRequest {
	r.current = &current
	return r
}

// Page size, default &#x60;50&#x60;
func (r ApiQueryLiquidationLoanRepayHistoryRequest) Size(size int64) ApiQueryLiquidationLoanRepayHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryLiquidationLoanRepayHistoryRequest) RecvWindow(recvWindow int64) ApiQueryLiquidationLoanRepayHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryLiquidationLoanRepayHistoryRequest) Execute() (*common.RestApiResponse[models.QueryLiquidationLoanRepayHistoryResponse], error) {
	return r.ApiService.QueryLiquidationLoanRepayHistoryExecute(r)
}

/*
QueryLiquidationLoanRepayHistory Query Liquidation Loan Repay History (USER_DATA)
Get /sapi/v1/margin/liquidation-loan/repay-history

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-liquidation-loan-repay-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time in Unix timestamp (milliseconds). Defaults to 7 days ago if not specified
@param endTime -  End time in Unix timestamp (milliseconds). Defaults to now if not specified
@param current -  Current page number, default `1`
@param size -  Page size, default `50`
@param recvWindow -
@return ApiQueryLiquidationLoanRepayHistoryRequest
*/
func (a *TradeAPIService) QueryLiquidationLoanRepayHistory(ctx context.Context) ApiQueryLiquidationLoanRepayHistoryRequest {
	return ApiQueryLiquidationLoanRepayHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryLiquidationLoanRepayHistoryResponse
func (a *TradeAPIService) QueryLiquidationLoanRepayHistoryExecute(r ApiQueryLiquidationLoanRepayHistoryRequest) (*common.RestApiResponse[models.QueryLiquidationLoanRepayHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/liquidation-loan/repay-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryLiquidationLoanRepayHistoryResponse](
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

type ApiQueryMarginAccountsAllOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryMarginAccountsAllOcoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsAllOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsAllOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) FromId(fromId int64) ApiQueryMarginAccountsAllOcoRequest {
	r.fromId = &fromId
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) StartTime(startTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) EndTime(endTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) Limit(limit int64) ApiQueryMarginAccountsAllOcoRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsAllOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsAllOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsAllOcoExecute(r)
}

/*
QueryMarginAccountsAllOco Query Margin Account's all OCO (USER_DATA)
Get /sapi/v1/margin/allOrderList

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-all-oco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -
@param symbol -
@param fromId -
@param startTime -
@param endTime -
@param limit -
@param recvWindow -
@return ApiQueryMarginAccountsAllOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsAllOco(ctx context.Context) ApiQueryMarginAccountsAllOcoRequest {
	return ApiQueryMarginAccountsAllOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsAllOcoResponse
func (a *TradeAPIService) QueryMarginAccountsAllOcoExecute(r ApiQueryMarginAccountsAllOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsAllOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/allOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
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

	resp, err := SendRequest[models.QueryMarginAccountsAllOcoResponse](
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

type ApiQueryMarginAccountsAllOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryMarginAccountsAllOrdersRequest) Symbol(symbol string) ApiQueryMarginAccountsAllOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsAllOrdersRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) OrderId(orderId int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) StartTime(startTime int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) EndTime(endTime int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) Limit(limit int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsAllOrdersResponse], error) {
	return r.ApiService.QueryMarginAccountsAllOrdersExecute(r)
}

/*
QueryMarginAccountsAllOrders Query Margin Account's All Orders (USER_DATA)
Get /sapi/v1/margin/allOrders

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-all-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param orderId -
@param startTime -
@param endTime -
@param limit -
@param recvWindow -
@return ApiQueryMarginAccountsAllOrdersRequest
*/
func (a *TradeAPIService) QueryMarginAccountsAllOrders(ctx context.Context) ApiQueryMarginAccountsAllOrdersRequest {
	return ApiQueryMarginAccountsAllOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsAllOrdersResponse
func (a *TradeAPIService) QueryMarginAccountsAllOrdersExecute(r ApiQueryMarginAccountsAllOrdersRequest) (*common.RestApiResponse[models.QueryMarginAccountsAllOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
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

	resp, err := SendRequest[models.QueryMarginAccountsAllOrdersResponse](
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

type ApiQueryMarginAccountsOcoRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	isIsolated        *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	symbol            *string
	orderListId       *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryMarginAccountsOcoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) OrderListId(orderListId int64) ApiQueryMarginAccountsOcoRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) OrigClientOrderId(origClientOrderId string) ApiQueryMarginAccountsOcoRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsOcoExecute(r)
}

/*
QueryMarginAccountsOco Query Margin Account's OCO (USER_DATA)
Get /sapi/v1/margin/orderList

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-oco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -
@param symbol -
@param orderListId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryMarginAccountsOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOco(ctx context.Context) ApiQueryMarginAccountsOcoRequest {
	return ApiQueryMarginAccountsOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOcoResponse
func (a *TradeAPIService) QueryMarginAccountsOcoExecute(r ApiQueryMarginAccountsOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOcoResponse](
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

type ApiQueryMarginAccountsOpenOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryMarginAccountsOpenOcoRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsOpenOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsOpenOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsOpenOcoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsOpenOcoRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOpenOcoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOpenOcoRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOpenOcoResponse], error) {
	return r.ApiService.QueryMarginAccountsOpenOcoExecute(r)
}

/*
QueryMarginAccountsOpenOco Query Margin Account's Open OCO (USER_DATA)
Get /sapi/v1/margin/openOrderList

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-open-oco

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -
@param symbol -
@param recvWindow -
@return ApiQueryMarginAccountsOpenOcoRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOpenOco(ctx context.Context) ApiQueryMarginAccountsOpenOcoRequest {
	return ApiQueryMarginAccountsOpenOcoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOpenOcoResponse
func (a *TradeAPIService) QueryMarginAccountsOpenOcoExecute(r ApiQueryMarginAccountsOpenOcoRequest) (*common.RestApiResponse[models.QueryMarginAccountsOpenOcoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/openOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOpenOcoResponse](
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

type ApiQueryMarginAccountsOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	recvWindow *int64
}

// isolated margin pair
func (r ApiQueryMarginAccountsOpenOrdersRequest) Symbol(symbol string) ApiQueryMarginAccountsOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsOpenOrdersRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsOpenOrdersRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsOpenOrdersRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOpenOrdersRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOpenOrdersResponse], error) {
	return r.ApiService.QueryMarginAccountsOpenOrdersExecute(r)
}

/*
QueryMarginAccountsOpenOrders Query Margin Account's Open Orders (USER_DATA)
Get /sapi/v1/margin/openOrders

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@param isIsolated -
@param recvWindow -
@return ApiQueryMarginAccountsOpenOrdersRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOpenOrders(ctx context.Context) ApiQueryMarginAccountsOpenOrdersRequest {
	return ApiQueryMarginAccountsOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOpenOrdersResponse
func (a *TradeAPIService) QueryMarginAccountsOpenOrdersExecute(r ApiQueryMarginAccountsOpenOrdersRequest) (*common.RestApiResponse[models.QueryMarginAccountsOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOpenOrdersResponse](
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

type ApiQueryMarginAccountsOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryMarginAccountsOrderRequest) Symbol(symbol string) ApiQueryMarginAccountsOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsOrderRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsOrderRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsOrderRequest) OrderId(orderId int64) ApiQueryMarginAccountsOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryMarginAccountsOrderRequest) OrigClientOrderId(origClientOrderId string) ApiQueryMarginAccountsOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

func (r ApiQueryMarginAccountsOrderRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsOrderRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsOrderResponse], error) {
	return r.ApiService.QueryMarginAccountsOrderExecute(r)
}

/*
QueryMarginAccountsOrder Query Margin Account's Order (USER_DATA)
Get /sapi/v1/margin/order

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param orderId -
@param origClientOrderId -
@param recvWindow -
@return ApiQueryMarginAccountsOrderRequest
*/
func (a *TradeAPIService) QueryMarginAccountsOrder(ctx context.Context) ApiQueryMarginAccountsOrderRequest {
	return ApiQueryMarginAccountsOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsOrderResponse
func (a *TradeAPIService) QueryMarginAccountsOrderExecute(r ApiQueryMarginAccountsOrderRequest) (*common.RestApiResponse[models.QueryMarginAccountsOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginAccountsOrderResponse](
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

type ApiQueryMarginAccountsTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiQueryMarginAccountsTradeListRequest) Symbol(symbol string) ApiQueryMarginAccountsTradeListRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryMarginAccountsTradeListRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) OrderId(orderId int64) ApiQueryMarginAccountsTradeListRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) StartTime(startTime int64) ApiQueryMarginAccountsTradeListRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) EndTime(endTime int64) ApiQueryMarginAccountsTradeListRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) FromId(fromId int64) ApiQueryMarginAccountsTradeListRequest {
	r.fromId = &fromId
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) Limit(limit int64) ApiQueryMarginAccountsTradeListRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) RecvWindow(recvWindow int64) ApiQueryMarginAccountsTradeListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) Execute() (*common.RestApiResponse[models.QueryMarginAccountsTradeListResponse], error) {
	return r.ApiService.QueryMarginAccountsTradeListExecute(r)
}

/*
QueryMarginAccountsTradeList Query Margin Account's Trade List (USER_DATA)
Get /sapi/v1/margin/myTrades

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-margin-accounts-trade-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -
@param orderId -
@param startTime -
@param endTime -
@param fromId -
@param limit -
@param recvWindow -
@return ApiQueryMarginAccountsTradeListRequest
*/
func (a *TradeAPIService) QueryMarginAccountsTradeList(ctx context.Context) ApiQueryMarginAccountsTradeListRequest {
	return ApiQueryMarginAccountsTradeListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAccountsTradeListResponse
func (a *TradeAPIService) QueryMarginAccountsTradeListExecute(r ApiQueryMarginAccountsTradeListRequest) (*common.RestApiResponse[models.QueryMarginAccountsTradeListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/myTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
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

	resp, err := SendRequest[models.QueryMarginAccountsTradeListResponse](
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

type ApiQueryPreventedMatchesRequest struct {
	ctx                  context.Context
	ApiService           *TradeAPIService
	symbol               *string
	preventedMatchId     *int64
	orderId              *int64
	fromPreventedMatchId *int64
	isIsolated           *models.QueryMarginAccountsOpenOrdersIsIsolatedParameter
	recvWindow           *int64
}

func (r ApiQueryPreventedMatchesRequest) Symbol(symbol string) ApiQueryPreventedMatchesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryPreventedMatchesRequest) PreventedMatchId(preventedMatchId int64) ApiQueryPreventedMatchesRequest {
	r.preventedMatchId = &preventedMatchId
	return r
}

func (r ApiQueryPreventedMatchesRequest) OrderId(orderId int64) ApiQueryPreventedMatchesRequest {
	r.orderId = &orderId
	return r
}

func (r ApiQueryPreventedMatchesRequest) FromPreventedMatchId(fromPreventedMatchId int64) ApiQueryPreventedMatchesRequest {
	r.fromPreventedMatchId = &fromPreventedMatchId
	return r
}

func (r ApiQueryPreventedMatchesRequest) IsIsolated(isIsolated models.QueryMarginAccountsOpenOrdersIsIsolatedParameter) ApiQueryPreventedMatchesRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryPreventedMatchesRequest) RecvWindow(recvWindow int64) ApiQueryPreventedMatchesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPreventedMatchesRequest) Execute() (*common.RestApiResponse[models.QueryPreventedMatchesResponse], error) {
	return r.ApiService.QueryPreventedMatchesExecute(r)
}

/*
QueryPreventedMatches Query Prevented Matches (USER_DATA)
Get /sapi/v1/margin/myPreventedMatches

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-prevented-matches

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param preventedMatchId -
@param orderId -
@param fromPreventedMatchId -
@param isIsolated -
@param recvWindow -
@return ApiQueryPreventedMatchesRequest
*/
func (a *TradeAPIService) QueryPreventedMatches(ctx context.Context) ApiQueryPreventedMatchesRequest {
	return ApiQueryPreventedMatchesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPreventedMatchesResponse
func (a *TradeAPIService) QueryPreventedMatchesExecute(r ApiQueryPreventedMatchesRequest) (*common.RestApiResponse[models.QueryPreventedMatchesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/myPreventedMatches"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.preventedMatchId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "preventedMatchId", r.preventedMatchId, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.fromPreventedMatchId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromPreventedMatchId", r.fromPreventedMatchId, "form", "")
	}
	if r.isIsolated != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isIsolated", r.isIsolated, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPreventedMatchesResponse](
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

type ApiQuerySpecialKeyRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQuerySpecialKeyRequest) Symbol(symbol string) ApiQuerySpecialKeyRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQuerySpecialKeyRequest) RecvWindow(recvWindow int64) ApiQuerySpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySpecialKeyRequest) Execute() (*common.RestApiResponse[models.QuerySpecialKeyResponse], error) {
	return r.ApiService.QuerySpecialKeyExecute(r)
}

/*
QuerySpecialKey Query Special key(Low Latency Trading) (TRADE)
Get /sapi/v1/margin/apiKey

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-special-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQuerySpecialKeyRequest
*/
func (a *TradeAPIService) QuerySpecialKey(ctx context.Context) ApiQuerySpecialKeyRequest {
	return ApiQuerySpecialKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySpecialKeyResponse
func (a *TradeAPIService) QuerySpecialKeyExecute(r ApiQuerySpecialKeyRequest) (*common.RestApiResponse[models.QuerySpecialKeyResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/apiKey"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySpecialKeyResponse](
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

type ApiQuerySpecialKeyListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQuerySpecialKeyListRequest) Symbol(symbol string) ApiQuerySpecialKeyListRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQuerySpecialKeyListRequest) RecvWindow(recvWindow int64) ApiQuerySpecialKeyListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySpecialKeyListRequest) Execute() (*common.RestApiResponse[models.QuerySpecialKeyListResponse], error) {
	return r.ApiService.QuerySpecialKeyListExecute(r)
}

/*
QuerySpecialKeyList Query Special key List(Low Latency Trading) (TRADE)
Get /sapi/v1/margin/api-key-list

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#query-special-key-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQuerySpecialKeyListRequest
*/
func (a *TradeAPIService) QuerySpecialKeyList(ctx context.Context) ApiQuerySpecialKeyListRequest {
	return ApiQuerySpecialKeyListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySpecialKeyListResponse
func (a *TradeAPIService) QuerySpecialKeyListExecute(r ApiQuerySpecialKeyListRequest) (*common.RestApiResponse[models.QuerySpecialKeyListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/api-key-list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySpecialKeyListResponse](
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

type ApiSmallLiabilityExchangeRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	assetNames *string
	recvWindow *int64
}

// The assets list of small liability exchange
func (r ApiSmallLiabilityExchangeRequest) AssetNames(assetNames string) ApiSmallLiabilityExchangeRequest {
	r.assetNames = &assetNames
	return r
}

func (r ApiSmallLiabilityExchangeRequest) RecvWindow(recvWindow int64) ApiSmallLiabilityExchangeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSmallLiabilityExchangeRequest) Execute() (struct{}, error) {
	return r.ApiService.SmallLiabilityExchangeExecute(r)
}

/*
SmallLiabilityExchange Small Liability Exchange (MARGIN)
Post /sapi/v1/margin/exchange-small-liability

https://developers.binance.com/en/docs/catalog/core-trading-margin-trading/api/rest-api/trade#small-liability-exchange

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param assetNames -  The assets list of small liability exchange
@param recvWindow -
@return ApiSmallLiabilityExchangeRequest
*/
func (a *TradeAPIService) SmallLiabilityExchange(ctx context.Context) ApiSmallLiabilityExchangeRequest {
	return ApiSmallLiabilityExchangeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TradeAPIService) SmallLiabilityExchangeExecute(r ApiSmallLiabilityExchangeRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/exchange-small-liability"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.assetNames == nil {
		return struct{}{}, common.ReportError("assetNames is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "assetNames", r.assetNames, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		true,
	)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}
