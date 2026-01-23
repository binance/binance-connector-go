/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
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
	permissionMode *string
	recvWindow     *int64
}

func (r ApiCreateSpecialKeyRequest) ApiName(apiName string) ApiCreateSpecialKeyRequest {
	r.apiName = &apiName
	return r
}

// isolated margin pair
func (r ApiCreateSpecialKeyRequest) Symbol(symbol string) ApiCreateSpecialKeyRequest {
	r.symbol = &symbol
	return r
}

// Can be added in batches, separated by commas. Max 30 for an API key
func (r ApiCreateSpecialKeyRequest) Ip(ip string) ApiCreateSpecialKeyRequest {
	r.ip = &ip
	return r
}

// 1. If publicKey is inputted it will create an RSA or Ed25519 key. &lt;br /&gt;2. Need to be encoded to URL-encoded format
func (r ApiCreateSpecialKeyRequest) PublicKey(publicKey string) ApiCreateSpecialKeyRequest {
	r.publicKey = &publicKey
	return r
}

// This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE.
func (r ApiCreateSpecialKeyRequest) PermissionMode(permissionMode string) ApiCreateSpecialKeyRequest {
	r.permissionMode = &permissionMode
	return r
}

// No more than 60000
func (r ApiCreateSpecialKeyRequest) RecvWindow(recvWindow int64) ApiCreateSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateSpecialKeyRequest) Execute() (*common.RestApiResponse[models.CreateSpecialKeyResponse], error) {
	return r.ApiService.CreateSpecialKeyExecute(r)
}

/*
CreateSpecialKey Create Special Key(Low-Latency Trading)(TRADE)
Post /sapi/v1/margin/apiKey

https://developers.binance.com/docs/margin_trading/trade/Create-Special-Key-of-Low-Latency-Trading

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param apiName -
@param symbol -  isolated margin pair
@param ip -  Can be added in batches, separated by commas. Max 30 for an API key
@param publicKey -  1. If publicKey is inputted it will create an RSA or Ed25519 key. <br />2. Need to be encoded to URL-encoded format
@param permissionMode -  This parameter is only for the Ed25519 API key, and does not effact for other encryption methods. The value can be TRADE (TRADE for all permissions) or READ (READ for USER_DATA, FIX_API_READ_ONLY). The default value is TRADE.
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.CreateSpecialKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// isolated margin pair
func (r ApiDeleteSpecialKeyRequest) Symbol(symbol string) ApiDeleteSpecialKeyRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiDeleteSpecialKeyRequest) RecvWindow(recvWindow int64) ApiDeleteSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteSpecialKeyRequest) Execute() (struct{}, error) {
	return r.ApiService.DeleteSpecialKeyExecute(r)
}

/*
DeleteSpecialKey Delete Special Key(Low-Latency Trading)(TRADE)
Delete /sapi/v1/margin/apiKey

https://developers.binance.com/docs/margin_trading/trade/Delete-Special-Key-of-Low-Latency-Trading

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param apiName -
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// No more than 60000
func (r ApiEditIpForSpecialKeyRequest) RecvWindow(recvWindow int64) ApiEditIpForSpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEditIpForSpecialKeyRequest) Execute() (struct{}, error) {
	return r.ApiService.EditIpForSpecialKeyExecute(r)
}

/*
EditIpForSpecialKey Edit ip for Special Key(Low-Latency Trading)(TRADE)
Put /sapi/v1/margin/apiKey/ip

https://developers.binance.com/docs/margin_trading/trade/Edit-ip-for-Special-Key-of-Low-Latency-Trading

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param ip -  Can be added in batches, separated by commas. Max 30 for an API key
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
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

// Only supports querying data from the past 90 days.
func (r ApiGetForceLiquidationRecordRequest) StartTime(startTime int64) ApiGetForceLiquidationRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetForceLiquidationRecordRequest) EndTime(endTime int64) ApiGetForceLiquidationRecordRequest {
	r.endTime = &endTime
	return r
}

// isolated symbol
func (r ApiGetForceLiquidationRecordRequest) IsolatedSymbol(isolatedSymbol string) ApiGetForceLiquidationRecordRequest {
	r.isolatedSymbol = &isolatedSymbol
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetForceLiquidationRecordRequest) Current(current int64) ApiGetForceLiquidationRecordRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiGetForceLiquidationRecordRequest) Size(size int64) ApiGetForceLiquidationRecordRequest {
	r.size = &size
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Get-Force-Liquidation-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param isolatedSymbol -  isolated symbol
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.GetForceLiquidationRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Get-Small-Liability-Exchange-Coin-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.GetSmallLiabilityExchangeCoinListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// Currently querying page. Start from 1. Default:1
func (r ApiGetSmallLiabilityExchangeHistoryRequest) Current(current int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetSmallLiabilityExchangeHistoryRequest) Size(size int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.size = &size
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiGetSmallLiabilityExchangeHistoryRequest) StartTime(startTime int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSmallLiabilityExchangeHistoryRequest) EndTime(endTime int64) ApiGetSmallLiabilityExchangeHistoryRequest {
	r.endTime = &endTime
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Get-Small-Liability-Exchange-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param recvWindow -  No more than 60000
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
	if r.size == nil {
		return nil, common.ReportError("size is required and must be specified")
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

	resp, err := SendRequest[models.GetSmallLiabilityExchangeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *string
	recvWindow *int64
}

func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) Symbol(symbol string) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	r.symbol = &symbol
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest) IsIsolated(isIsolated string) ApiMarginAccountCancelAllOpenOrdersOnASymbolRequest {
	r.isIsolated = &isIsolated
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-Cancel-All-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.MarginAccountCancelAllOpenOrdersOnASymbolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountCancelOcoRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *string
	orderListId       *int64
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiMarginAccountCancelOcoRequest) Symbol(symbol string) ApiMarginAccountCancelOcoRequest {
	r.symbol = &symbol
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountCancelOcoRequest) IsIsolated(isIsolated string) ApiMarginAccountCancelOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountCancelOcoRequest) OrderListId(orderListId int64) ApiMarginAccountCancelOcoRequest {
	r.orderListId = &orderListId
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountCancelOcoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountCancelOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiMarginAccountCancelOcoRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountCancelOcoRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-Cancel-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.MarginAccountCancelOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountCancelOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
}

func (r ApiMarginAccountCancelOrderRequest) Symbol(symbol string) ApiMarginAccountCancelOrderRequest {
	r.symbol = &symbol
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountCancelOrderRequest) IsIsolated(isIsolated string) ApiMarginAccountCancelOrderRequest {
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

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiMarginAccountCancelOrderRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountCancelOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-Cancel-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param orderId -
@param origClientOrderId -
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.MarginAccountCancelOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	quantity                *float32
	price                   *float32
	stopPrice               *float32
	isIsolated              *string
	listClientOrderId       *string
	limitClientOrderId      *string
	limitIcebergQty         *float32
	stopClientOrderId       *string
	stopLimitPrice          *float32
	stopIcebergQty          *float32
	stopLimitTimeInForce    *string
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *string
	selfTradePreventionMode *string
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

func (r ApiMarginAccountNewOcoRequest) Quantity(quantity float32) ApiMarginAccountNewOcoRequest {
	r.quantity = &quantity
	return r
}

func (r ApiMarginAccountNewOcoRequest) Price(price float32) ApiMarginAccountNewOcoRequest {
	r.price = &price
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopPrice(stopPrice float32) ApiMarginAccountNewOcoRequest {
	r.stopPrice = &stopPrice
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountNewOcoRequest) IsIsolated(isIsolated string) ApiMarginAccountNewOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountNewOcoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// A unique Id for the limit order
func (r ApiMarginAccountNewOcoRequest) LimitClientOrderId(limitClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.limitClientOrderId = &limitClientOrderId
	return r
}

func (r ApiMarginAccountNewOcoRequest) LimitIcebergQty(limitIcebergQty float32) ApiMarginAccountNewOcoRequest {
	r.limitIcebergQty = &limitIcebergQty
	return r
}

// A unique Id for the stop loss/stop loss limit leg
func (r ApiMarginAccountNewOcoRequest) StopClientOrderId(stopClientOrderId string) ApiMarginAccountNewOcoRequest {
	r.stopClientOrderId = &stopClientOrderId
	return r
}

// If provided, &#x60;stopLimitTimeInForce&#x60; is required.
func (r ApiMarginAccountNewOcoRequest) StopLimitPrice(stopLimitPrice float32) ApiMarginAccountNewOcoRequest {
	r.stopLimitPrice = &stopLimitPrice
	return r
}

func (r ApiMarginAccountNewOcoRequest) StopIcebergQty(stopIcebergQty float32) ApiMarginAccountNewOcoRequest {
	r.stopIcebergQty = &stopIcebergQty
	return r
}

// Valid values are &#x60;GTC&#x60;/&#x60;FOK&#x60;/&#x60;IOC&#x60;
func (r ApiMarginAccountNewOcoRequest) StopLimitTimeInForce(stopLimitTimeInForce string) ApiMarginAccountNewOcoRequest {
	r.stopLimitTimeInForce = &stopLimitTimeInForce
	return r
}

// Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOcoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOcoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
func (r ApiMarginAccountNewOcoRequest) SideEffectType(sideEffectType string) ApiMarginAccountNewOcoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
func (r ApiMarginAccountNewOcoRequest) SelfTradePreventionMode(selfTradePreventionMode string) ApiMarginAccountNewOcoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
func (r ApiMarginAccountNewOcoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOcoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-New-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param quantity -
@param price -
@param stopPrice -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param limitClientOrderId -  A unique Id for the limit order
@param limitIcebergQty -
@param stopClientOrderId -  A unique Id for the stop loss/stop loss limit leg
@param stopLimitPrice -  If provided, `stopLimitTimeInForce` is required.
@param stopIcebergQty -
@param stopLimitTimeInForce -  Valid values are `GTC`/`FOK`/`IOC`
@param newOrderRespType -  Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
@param selfTradePreventionMode -  The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.MarginAccountNewOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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
	type_                   *string
	isIsolated              *string
	quantity                *float32
	quoteOrderQty           *float32
	price                   *float32
	stopPrice               *float32
	newClientOrderId        *string
	icebergQty              *float32
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *string
	timeInForce             *models.MarginAccountNewOrderTimeInForceParameter
	selfTradePreventionMode *string
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

// &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60;
func (r ApiMarginAccountNewOrderRequest) Type(type_ string) ApiMarginAccountNewOrderRequest {
	r.type_ = &type_
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountNewOrderRequest) IsIsolated(isIsolated string) ApiMarginAccountNewOrderRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiMarginAccountNewOrderRequest) Quantity(quantity float32) ApiMarginAccountNewOrderRequest {
	r.quantity = &quantity
	return r
}

func (r ApiMarginAccountNewOrderRequest) QuoteOrderQty(quoteOrderQty float32) ApiMarginAccountNewOrderRequest {
	r.quoteOrderQty = &quoteOrderQty
	return r
}

func (r ApiMarginAccountNewOrderRequest) Price(price float32) ApiMarginAccountNewOrderRequest {
	r.price = &price
	return r
}

// Used with &#x60;STOP_LOSS&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, &#x60;TAKE_PROFIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; orders.
func (r ApiMarginAccountNewOrderRequest) StopPrice(stopPrice float32) ApiMarginAccountNewOrderRequest {
	r.stopPrice = &stopPrice
	return r
}

// Used to uniquely identify this cancel. Automatically generated by default
func (r ApiMarginAccountNewOrderRequest) NewClientOrderId(newClientOrderId string) ApiMarginAccountNewOrderRequest {
	r.newClientOrderId = &newClientOrderId
	return r
}

// Used with &#x60;LIMIT&#x60;, &#x60;STOP_LOSS_LIMIT&#x60;, and &#x60;TAKE_PROFIT_LIMIT&#x60; to create an iceberg order.
func (r ApiMarginAccountNewOrderRequest) IcebergQty(icebergQty float32) ApiMarginAccountNewOrderRequest {
	r.icebergQty = &icebergQty
	return r
}

// Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOrderRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOrderRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
func (r ApiMarginAccountNewOrderRequest) SideEffectType(sideEffectType string) ApiMarginAccountNewOrderRequest {
	r.sideEffectType = &sideEffectType
	return r
}

// GTC,IOC,FOK
func (r ApiMarginAccountNewOrderRequest) TimeInForce(timeInForce models.MarginAccountNewOrderTimeInForceParameter) ApiMarginAccountNewOrderRequest {
	r.timeInForce = &timeInForce
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
func (r ApiMarginAccountNewOrderRequest) SelfTradePreventionMode(selfTradePreventionMode string) ApiMarginAccountNewOrderRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
func (r ApiMarginAccountNewOrderRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOrderRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-New-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param side -
@param type_ -  `MARGIN`,`ISOLATED`
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param quantity -
@param quoteOrderQty -
@param price -
@param stopPrice -  Used with `STOP_LOSS`, `STOP_LOSS_LIMIT`, `TAKE_PROFIT`, and `TAKE_PROFIT_LIMIT` orders.
@param newClientOrderId -  Used to uniquely identify this cancel. Automatically generated by default
@param icebergQty -  Used with `LIMIT`, `STOP_LOSS_LIMIT`, and `TAKE_PROFIT_LIMIT` to create an iceberg order.
@param newOrderRespType -  Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
@param timeInForce -  GTC,IOC,FOK
@param selfTradePreventionMode -  The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
@param recvWindow -  No more than 60000
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
	if r.autoRepayAtCancel != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepayAtCancel", r.autoRepayAtCancel, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginAccountNewOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountNewOtoRequest struct {
	ctx                     context.Context
	ApiService              *TradeAPIService
	symbol                  *string
	workingType             *string
	workingSide             *string
	workingPrice            *float32
	workingQuantity         *float32
	workingIcebergQty       *float32
	pendingType             *string
	pendingSide             *string
	pendingQuantity         *float32
	isIsolated              *string
	listClientOrderId       *string
	newOrderRespType        *models.MarginAccountNewOrderNewOrderRespTypeParameter
	sideEffectType          *string
	selfTradePreventionMode *string
	autoRepayAtCancel       *bool
	workingClientOrderId    *string
	workingTimeInForce      *string
	pendingClientOrderId    *string
	pendingPrice            *float32
	pendingStopPrice        *float32
	pendingTrailingDelta    *float32
	pendingIcebergQty       *float32
	pendingTimeInForce      *string
}

func (r ApiMarginAccountNewOtoRequest) Symbol(symbol string) ApiMarginAccountNewOtoRequest {
	r.symbol = &symbol
	return r
}

// Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;
func (r ApiMarginAccountNewOtoRequest) WorkingType(workingType string) ApiMarginAccountNewOtoRequest {
	r.workingType = &workingType
	return r
}

// BUY, SELL
func (r ApiMarginAccountNewOtoRequest) WorkingSide(workingSide string) ApiMarginAccountNewOtoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingPrice(workingPrice float32) ApiMarginAccountNewOtoRequest {
	r.workingPrice = &workingPrice
	return r
}

func (r ApiMarginAccountNewOtoRequest) WorkingQuantity(workingQuantity float32) ApiMarginAccountNewOtoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiMarginAccountNewOtoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// Supported values: [Order Types](https://developers.binance.com/docs/binance-spot-api-docs/enums#order-types-ordertypes-type) Note that &#x60;MARKET&#x60; orders using &#x60;quoteOrderQty&#x60; are not supported.
func (r ApiMarginAccountNewOtoRequest) PendingType(pendingType string) ApiMarginAccountNewOtoRequest {
	r.pendingType = &pendingType
	return r
}

// BUY, SELL
func (r ApiMarginAccountNewOtoRequest) PendingSide(pendingSide string) ApiMarginAccountNewOtoRequest {
	r.pendingSide = &pendingSide
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingQuantity(pendingQuantity float32) ApiMarginAccountNewOtoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountNewOtoRequest) IsIsolated(isIsolated string) ApiMarginAccountNewOtoRequest {
	r.isIsolated = &isIsolated
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountNewOtoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOtoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOtoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
func (r ApiMarginAccountNewOtoRequest) SideEffectType(sideEffectType string) ApiMarginAccountNewOtoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
func (r ApiMarginAccountNewOtoRequest) SelfTradePreventionMode(selfTradePreventionMode string) ApiMarginAccountNewOtoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
func (r ApiMarginAccountNewOtoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOtoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtoRequest) WorkingClientOrderId(workingClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// GTC,IOC,FOK
func (r ApiMarginAccountNewOtoRequest) WorkingTimeInForce(workingTimeInForce string) ApiMarginAccountNewOtoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtoRequest) PendingClientOrderId(pendingClientOrderId string) ApiMarginAccountNewOtoRequest {
	r.pendingClientOrderId = &pendingClientOrderId
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingPrice(pendingPrice float32) ApiMarginAccountNewOtoRequest {
	r.pendingPrice = &pendingPrice
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingStopPrice(pendingStopPrice float32) ApiMarginAccountNewOtoRequest {
	r.pendingStopPrice = &pendingStopPrice
	return r
}

func (r ApiMarginAccountNewOtoRequest) PendingTrailingDelta(pendingTrailingDelta float32) ApiMarginAccountNewOtoRequest {
	r.pendingTrailingDelta = &pendingTrailingDelta
	return r
}

// This can only be used if &#x60;pendingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtoRequest) PendingIcebergQty(pendingIcebergQty float32) ApiMarginAccountNewOtoRequest {
	r.pendingIcebergQty = &pendingIcebergQty
	return r
}

// GTC,IOC,FOK
func (r ApiMarginAccountNewOtoRequest) PendingTimeInForce(pendingTimeInForce string) ApiMarginAccountNewOtoRequest {
	r.pendingTimeInForce = &pendingTimeInForce
	return r
}

func (r ApiMarginAccountNewOtoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOtoResponse], error) {
	return r.ApiService.MarginAccountNewOtoExecute(r)
}

/*
MarginAccountNewOto Margin Account New OTO (TRADE)
Post /sapi/v1/margin/order/oto

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-New-OTO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -  Supported values: `LIMIT`, `LIMIT_MAKER`
@param workingSide -  BUY, SELL
@param workingPrice -
@param workingQuantity -
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`.
@param pendingType -  Supported values: [Order Types](https://developers.binance.com/docs/binance-spot-api-docs/enums#order-types-ordertypes-type) Note that `MARKET` orders using `quoteOrderQty` are not supported.
@param pendingSide -  BUY, SELL
@param pendingQuantity -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param newOrderRespType -  Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
@param selfTradePreventionMode -  The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingTimeInForce -  GTC,IOC,FOK
@param pendingClientOrderId -  Arbitrary unique ID among open orders for the pending order. Automatically generated if not sent.
@param pendingPrice -
@param pendingStopPrice -
@param pendingTrailingDelta -
@param pendingIcebergQty -  This can only be used if `pendingTimeInForce` is `GTC`.
@param pendingTimeInForce -  GTC,IOC,FOK
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

	resp, err := SendRequest[models.MarginAccountNewOtoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginAccountNewOtocoRequest struct {
	ctx                       context.Context
	ApiService                *TradeAPIService
	symbol                    *string
	workingType               *string
	workingSide               *string
	workingPrice              *float32
	workingQuantity           *float32
	pendingSide               *string
	pendingQuantity           *float32
	pendingAboveType          *string
	isIsolated                *string
	sideEffectType            *string
	autoRepayAtCancel         *bool
	listClientOrderId         *string
	newOrderRespType          *models.MarginAccountNewOrderNewOrderRespTypeParameter
	selfTradePreventionMode   *string
	workingClientOrderId      *string
	workingIcebergQty         *float32
	workingTimeInForce        *string
	pendingAboveClientOrderId *string
	pendingAbovePrice         *float32
	pendingAboveStopPrice     *float32
	pendingAboveTrailingDelta *float32
	pendingAboveIcebergQty    *float32
	pendingAboveTimeInForce   *string
	pendingBelowType          *string
	pendingBelowClientOrderId *string
	pendingBelowPrice         *float32
	pendingBelowStopPrice     *float32
	pendingBelowTrailingDelta *float32
	pendingBelowIcebergQty    *float32
	pendingBelowTimeInForce   *string
}

func (r ApiMarginAccountNewOtocoRequest) Symbol(symbol string) ApiMarginAccountNewOtocoRequest {
	r.symbol = &symbol
	return r
}

// Supported values: &#x60;LIMIT&#x60;, &#x60;LIMIT_MAKER&#x60;
func (r ApiMarginAccountNewOtocoRequest) WorkingType(workingType string) ApiMarginAccountNewOtocoRequest {
	r.workingType = &workingType
	return r
}

// BUY, SELL
func (r ApiMarginAccountNewOtocoRequest) WorkingSide(workingSide string) ApiMarginAccountNewOtocoRequest {
	r.workingSide = &workingSide
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingPrice(workingPrice float32) ApiMarginAccountNewOtocoRequest {
	r.workingPrice = &workingPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) WorkingQuantity(workingQuantity float32) ApiMarginAccountNewOtocoRequest {
	r.workingQuantity = &workingQuantity
	return r
}

// BUY, SELL
func (r ApiMarginAccountNewOtocoRequest) PendingSide(pendingSide string) ApiMarginAccountNewOtocoRequest {
	r.pendingSide = &pendingSide
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingQuantity(pendingQuantity float32) ApiMarginAccountNewOtocoRequest {
	r.pendingQuantity = &pendingQuantity
	return r
}

// Supported values: &#x60;LIMIT_MAKER&#x60;, &#x60;STOP_LOSS&#x60;, and &#x60;STOP_LOSS_LIMIT&#x60;
func (r ApiMarginAccountNewOtocoRequest) PendingAboveType(pendingAboveType string) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveType = &pendingAboveType
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiMarginAccountNewOtocoRequest) IsIsolated(isIsolated string) ApiMarginAccountNewOtocoRequest {
	r.isIsolated = &isIsolated
	return r
}

// NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
func (r ApiMarginAccountNewOtocoRequest) SideEffectType(sideEffectType string) ApiMarginAccountNewOtocoRequest {
	r.sideEffectType = &sideEffectType
	return r
}

// Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
func (r ApiMarginAccountNewOtocoRequest) AutoRepayAtCancel(autoRepayAtCancel bool) ApiMarginAccountNewOtocoRequest {
	r.autoRepayAtCancel = &autoRepayAtCancel
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiMarginAccountNewOtocoRequest) ListClientOrderId(listClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.listClientOrderId = &listClientOrderId
	return r
}

// Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (r ApiMarginAccountNewOtocoRequest) NewOrderRespType(newOrderRespType models.MarginAccountNewOrderNewOrderRespTypeParameter) ApiMarginAccountNewOtocoRequest {
	r.newOrderRespType = &newOrderRespType
	return r
}

// The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
func (r ApiMarginAccountNewOtocoRequest) SelfTradePreventionMode(selfTradePreventionMode string) ApiMarginAccountNewOtocoRequest {
	r.selfTradePreventionMode = &selfTradePreventionMode
	return r
}

// Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) WorkingClientOrderId(workingClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.workingClientOrderId = &workingClientOrderId
	return r
}

// This can only be used if &#x60;workingTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) WorkingIcebergQty(workingIcebergQty float32) ApiMarginAccountNewOtocoRequest {
	r.workingIcebergQty = &workingIcebergQty
	return r
}

// GTC,IOC,FOK
func (r ApiMarginAccountNewOtocoRequest) WorkingTimeInForce(workingTimeInForce string) ApiMarginAccountNewOtocoRequest {
	r.workingTimeInForce = &workingTimeInForce
	return r
}

// Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) PendingAboveClientOrderId(pendingAboveClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveClientOrderId = &pendingAboveClientOrderId
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAbovePrice(pendingAbovePrice float32) ApiMarginAccountNewOtocoRequest {
	r.pendingAbovePrice = &pendingAbovePrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveStopPrice(pendingAboveStopPrice float32) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveStopPrice = &pendingAboveStopPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveTrailingDelta(pendingAboveTrailingDelta float32) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveTrailingDelta = &pendingAboveTrailingDelta
	return r
}

// This can only be used if &#x60;pendingAboveTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) PendingAboveIcebergQty(pendingAboveIcebergQty float32) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveIcebergQty = &pendingAboveIcebergQty
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingAboveTimeInForce(pendingAboveTimeInForce string) ApiMarginAccountNewOtocoRequest {
	r.pendingAboveTimeInForce = &pendingAboveTimeInForce
	return r
}

// Supported values: &#x60;LIMIT_MAKER&#x60;, &#x60;STOP_LOSS&#x60;, and &#x60;STOP_LOSS_LIMIT&#x60;
func (r ApiMarginAccountNewOtocoRequest) PendingBelowType(pendingBelowType string) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowType = &pendingBelowType
	return r
}

// Arbitrary unique ID among open orders for the pending below order. Automatically generated if not sent.
func (r ApiMarginAccountNewOtocoRequest) PendingBelowClientOrderId(pendingBelowClientOrderId string) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowClientOrderId = &pendingBelowClientOrderId
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowPrice(pendingBelowPrice float32) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowPrice = &pendingBelowPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowStopPrice(pendingBelowStopPrice float32) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowStopPrice = &pendingBelowStopPrice
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowTrailingDelta(pendingBelowTrailingDelta float32) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowTrailingDelta = &pendingBelowTrailingDelta
	return r
}

// This can only be used if &#x60;pendingBelowTimeInForce&#x60; is &#x60;GTC&#x60;.
func (r ApiMarginAccountNewOtocoRequest) PendingBelowIcebergQty(pendingBelowIcebergQty float32) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowIcebergQty = &pendingBelowIcebergQty
	return r
}

func (r ApiMarginAccountNewOtocoRequest) PendingBelowTimeInForce(pendingBelowTimeInForce string) ApiMarginAccountNewOtocoRequest {
	r.pendingBelowTimeInForce = &pendingBelowTimeInForce
	return r
}

func (r ApiMarginAccountNewOtocoRequest) Execute() (*common.RestApiResponse[models.MarginAccountNewOtocoResponse], error) {
	return r.ApiService.MarginAccountNewOtocoExecute(r)
}

/*
MarginAccountNewOtoco Margin Account New OTOCO (TRADE)
Post /sapi/v1/margin/order/otoco

https://developers.binance.com/docs/margin_trading/trade/Margin-Account-New-OTOCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param workingType -  Supported values: `LIMIT`, `LIMIT_MAKER`
@param workingSide -  BUY, SELL
@param workingPrice -
@param workingQuantity -
@param pendingSide -  BUY, SELL
@param pendingQuantity -
@param pendingAboveType -  Supported values: `LIMIT_MAKER`, `STOP_LOSS`, and `STOP_LOSS_LIMIT`
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param sideEffectType -  NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY,AUTO_BORROW_REPAY; default NO_SIDE_EFFECT. More info in [FAQ](https://www.binance.com/en/support/faq/how-to-use-the-sideeffecttype-parameter-with-the-margin-order-endpoints-f9fc51cda1984bf08b95e0d96c4570bc)
@param autoRepayAtCancel -  Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
@param listClientOrderId -  Either `orderListId` or `listClientOrderId` must be provided
@param newOrderRespType -  Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
@param selfTradePreventionMode -  The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE
@param workingClientOrderId -  Arbitrary unique ID among open orders for the working order. Automatically generated if not sent.
@param workingIcebergQty -  This can only be used if `workingTimeInForce` is `GTC`.
@param workingTimeInForce -  GTC,IOC,FOK
@param pendingAboveClientOrderId -  Arbitrary unique ID among open orders for the pending above order. Automatically generated if not sent.
@param pendingAbovePrice -
@param pendingAboveStopPrice -
@param pendingAboveTrailingDelta -
@param pendingAboveIcebergQty -  This can only be used if `pendingAboveTimeInForce` is `GTC`.
@param pendingAboveTimeInForce -
@param pendingBelowType -  Supported values: `LIMIT_MAKER`, `STOP_LOSS`, and `STOP_LOSS_LIMIT`
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

	resp, err := SendRequest[models.MarginAccountNewOtocoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginManualLiquidationRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	type_      *string
	symbol     *string
	recvWindow *int64
}

// &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60;
func (r ApiMarginManualLiquidationRequest) Type(type_ string) ApiMarginManualLiquidationRequest {
	r.type_ = &type_
	return r
}

// isolated margin pair
func (r ApiMarginManualLiquidationRequest) Symbol(symbol string) ApiMarginManualLiquidationRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiMarginManualLiquidationRequest) RecvWindow(recvWindow int64) ApiMarginManualLiquidationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginManualLiquidationRequest) Execute() (*common.RestApiResponse[models.MarginManualLiquidationResponse], error) {
	return r.ApiService.MarginManualLiquidationExecute(r)
}

/*
MarginManualLiquidation Margin Manual Liquidation(MARGIN)
Post /sapi/v1/margin/manual-liquidation

https://developers.binance.com/docs/margin_trading/trade/Margin-Manual-Liquidation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -  `MARGIN`,`ISOLATED`
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.MarginManualLiquidationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCurrentMarginOrderCountUsageRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *string
	symbol     *string
	recvWindow *int64
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryCurrentMarginOrderCountUsageRequest) IsIsolated(isIsolated string) ApiQueryCurrentMarginOrderCountUsageRequest {
	r.isIsolated = &isIsolated
	return r
}

// isolated margin pair
func (r ApiQueryCurrentMarginOrderCountUsageRequest) Symbol(symbol string) ApiQueryCurrentMarginOrderCountUsageRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Current-Margin-Order-Count-Usage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryCurrentMarginOrderCountUsageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsAllOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *string
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsAllOcoRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsAllOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// isolated margin pair
func (r ApiQueryMarginAccountsAllOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsAllOcoRequest {
	r.symbol = &symbol
	return r
}

// If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned.
func (r ApiQueryMarginAccountsAllOcoRequest) FromId(fromId int64) ApiQueryMarginAccountsAllOcoRequest {
	r.fromId = &fromId
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiQueryMarginAccountsAllOcoRequest) StartTime(startTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsAllOcoRequest) EndTime(endTime int64) ApiQueryMarginAccountsAllOcoRequest {
	r.endTime = &endTime
	return r
}

// Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
func (r ApiQueryMarginAccountsAllOcoRequest) Limit(limit int64) ApiQueryMarginAccountsAllOcoRequest {
	r.limit = &limit
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-all-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param symbol -  isolated margin pair
@param fromId -  If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned.
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param limit -  Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsAllOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsAllOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *string
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

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsAllOrdersRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsAllOrdersRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) OrderId(orderId int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.orderId = &orderId
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiQueryMarginAccountsAllOrdersRequest) StartTime(startTime int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsAllOrdersRequest) EndTime(endTime int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.endTime = &endTime
	return r
}

// Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
func (r ApiQueryMarginAccountsAllOrdersRequest) Limit(limit int64) ApiQueryMarginAccountsAllOrdersRequest {
	r.limit = &limit
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-All-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param orderId -
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param limit -  Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsAllOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOcoRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	isIsolated        *string
	symbol            *string
	orderListId       *int64
	origClientOrderId *string
	recvWindow        *int64
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsOcoRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// isolated margin pair
func (r ApiQueryMarginAccountsOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsOcoRequest {
	r.symbol = &symbol
	return r
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiQueryMarginAccountsOcoRequest) OrderListId(orderListId int64) ApiQueryMarginAccountsOcoRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiQueryMarginAccountsOcoRequest) OrigClientOrderId(origClientOrderId string) ApiQueryMarginAccountsOcoRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param symbol -  isolated margin pair
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param origClientOrderId -
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOpenOcoRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	isIsolated *string
	symbol     *string
	recvWindow *int64
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsOpenOcoRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsOpenOcoRequest {
	r.isIsolated = &isIsolated
	return r
}

// isolated margin pair
func (r ApiQueryMarginAccountsOpenOcoRequest) Symbol(symbol string) ApiQueryMarginAccountsOpenOcoRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-Open-OCO

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsOpenOcoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *string
	recvWindow *int64
}

// isolated margin pair
func (r ApiQueryMarginAccountsOpenOrdersRequest) Symbol(symbol string) ApiQueryMarginAccountsOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsOpenOrdersRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsOpenOrdersRequest {
	r.isIsolated = &isIsolated
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-Open-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsOrderRequest struct {
	ctx               context.Context
	ApiService        *TradeAPIService
	symbol            *string
	isIsolated        *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (r ApiQueryMarginAccountsOrderRequest) Symbol(symbol string) ApiQueryMarginAccountsOrderRequest {
	r.symbol = &symbol
	return r
}

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsOrderRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsOrderRequest {
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

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-Order

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param orderId -
@param origClientOrderId -
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAccountsTradeListRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	symbol     *string
	isIsolated *string
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

// for isolated margin or not, \&quot;TRUE\&quot;, \&quot;FALSE\&quot;，default \&quot;FALSE\&quot;
func (r ApiQueryMarginAccountsTradeListRequest) IsIsolated(isIsolated string) ApiQueryMarginAccountsTradeListRequest {
	r.isIsolated = &isIsolated
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) OrderId(orderId int64) ApiQueryMarginAccountsTradeListRequest {
	r.orderId = &orderId
	return r
}

// Only supports querying data from the past 90 days.
func (r ApiQueryMarginAccountsTradeListRequest) StartTime(startTime int64) ApiQueryMarginAccountsTradeListRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryMarginAccountsTradeListRequest) EndTime(endTime int64) ApiQueryMarginAccountsTradeListRequest {
	r.endTime = &endTime
	return r
}

// If &#x60;fromId&#x60; is set, data with &#x60;id&#x60; greater than &#x60;fromId&#x60; will be returned. Otherwise, the latest data will be returned.
func (r ApiQueryMarginAccountsTradeListRequest) FromId(fromId int64) ApiQueryMarginAccountsTradeListRequest {
	r.fromId = &fromId
	return r
}

// Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
func (r ApiQueryMarginAccountsTradeListRequest) Limit(limit int64) ApiQueryMarginAccountsTradeListRequest {
	r.limit = &limit
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Query-Margin-Account-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param isIsolated -  for isolated margin or not, \"TRUE\", \"FALSE\"，default \"FALSE\"
@param orderId -
@param startTime -  Only supports querying data from the past 90 days.
@param endTime -
@param fromId -  If `fromId` is set, data with `id` greater than `fromId` will be returned. Otherwise, the latest data will be returned.
@param limit -  Limit on the number of data records returned per request. Default: 500; Maximum: 1000.
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QueryMarginAccountsTradeListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// isolated margin pair
func (r ApiQuerySpecialKeyRequest) Symbol(symbol string) ApiQuerySpecialKeyRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiQuerySpecialKeyRequest) RecvWindow(recvWindow int64) ApiQuerySpecialKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySpecialKeyRequest) Execute() (*common.RestApiResponse[models.QuerySpecialKeyResponse], error) {
	return r.ApiService.QuerySpecialKeyExecute(r)
}

/*
QuerySpecialKey Query Special key(Low Latency Trading)(TRADE)
Get /sapi/v1/margin/apiKey

https://developers.binance.com/docs/margin_trading/trade/Query-Special-Key-of-Low-Latency-Trading

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QuerySpecialKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
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

// isolated margin pair
func (r ApiQuerySpecialKeyListRequest) Symbol(symbol string) ApiQuerySpecialKeyListRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiQuerySpecialKeyListRequest) RecvWindow(recvWindow int64) ApiQuerySpecialKeyListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySpecialKeyListRequest) Execute() (*common.RestApiResponse[models.QuerySpecialKeyListResponse], error) {
	return r.ApiService.QuerySpecialKeyListExecute(r)
}

/*
QuerySpecialKeyList Query Special key List(Low Latency Trading)(TRADE)
Get /sapi/v1/margin/api-key-list

https://developers.binance.com/docs/margin_trading/trade/Query-Special-Key-List-of-Low-Latency-Trading

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
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

	resp, err := SendRequest[models.QuerySpecialKeyListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSmallLiabilityExchangeRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	assetNames *[]string
	recvWindow *int64
}

// The assets list of small liability exchange， Example: assetNames &#x3D; BTC,ETH
func (r ApiSmallLiabilityExchangeRequest) AssetNames(assetNames []string) ApiSmallLiabilityExchangeRequest {
	r.assetNames = &assetNames
	return r
}

// No more than 60000
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

https://developers.binance.com/docs/margin_trading/trade/Small-Liability-Exchange

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param assetNames -  The assets list of small liability exchange， Example: assetNames = BTC,ETH
@param recvWindow -  No more than 60000
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

	{
		t := *r.assetNames
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "assetNames", t, "form", "multi")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}
