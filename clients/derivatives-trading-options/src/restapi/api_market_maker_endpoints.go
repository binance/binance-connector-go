/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketMakerEndpointsAPIService MarketMakerEndpointsAPI Service
type MarketMakerEndpointsAPIService Service

type ApiAutoCancelAllOpenOrdersRequest struct {
	ctx         context.Context
	ApiService  *MarketMakerEndpointsAPIService
	underlyings *string
	recvWindow  *int64
}

// Option Underlying Symbols, e.g BTCUSDT,ETHUSDT
func (r ApiAutoCancelAllOpenOrdersRequest) Underlyings(underlyings string) ApiAutoCancelAllOpenOrdersRequest {
	r.underlyings = &underlyings
	return r
}

func (r ApiAutoCancelAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiAutoCancelAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAutoCancelAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse], error) {
	return r.ApiService.AutoCancelAllOpenOrdersExecute(r)
}

/*
AutoCancelAllOpenOrders Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)
Post /eapi/v1/countdownCancelAllHeartBeat

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Auto-Cancel-All-Open-Orders-Heartbeat

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlyings -  Option Underlying Symbols, e.g BTCUSDT,ETHUSDT
@param recvWindow -
@return ApiAutoCancelAllOpenOrdersRequest
*/
func (a *MarketMakerEndpointsAPIService) AutoCancelAllOpenOrders(ctx context.Context) ApiAutoCancelAllOpenOrdersRequest {
	return ApiAutoCancelAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AutoCancelAllOpenOrdersResponse
func (a *MarketMakerEndpointsAPIService) AutoCancelAllOpenOrdersExecute(r ApiAutoCancelAllOpenOrdersRequest) (*common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/countdownCancelAllHeartBeat"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlyings == nil {
		return nil, common.ReportError("underlyings is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlyings", r.underlyings, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AutoCancelAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAutoCancelAllOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

// underlying, e.g BTCUSDT
func (r ApiGetAutoCancelAllOpenOrdersRequest) Underlying(underlying string) ApiGetAutoCancelAllOpenOrdersRequest {
	r.underlying = &underlying
	return r
}

func (r ApiGetAutoCancelAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiGetAutoCancelAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetAutoCancelAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.GetAutoCancelAllOpenOrdersResponse], error) {
	return r.ApiService.GetAutoCancelAllOpenOrdersExecute(r)
}

/*
GetAutoCancelAllOpenOrders Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
Get /eapi/v1/countdownCancelAll

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Get-Auto-Cancel-All-Open-Orders-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  underlying, e.g BTCUSDT
@param recvWindow -
@return ApiGetAutoCancelAllOpenOrdersRequest
*/
func (a *MarketMakerEndpointsAPIService) GetAutoCancelAllOpenOrders(ctx context.Context) ApiGetAutoCancelAllOpenOrdersRequest {
	return ApiGetAutoCancelAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAutoCancelAllOpenOrdersResponse
func (a *MarketMakerEndpointsAPIService) GetAutoCancelAllOpenOrdersExecute(r ApiGetAutoCancelAllOpenOrdersRequest) (*common.RestApiResponse[models.GetAutoCancelAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/countdownCancelAll"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetAutoCancelAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetMarketMakerProtectionConfigRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

// underlying, e.g BTCUSDT
func (r ApiGetMarketMakerProtectionConfigRequest) Underlying(underlying string) ApiGetMarketMakerProtectionConfigRequest {
	r.underlying = &underlying
	return r
}

func (r ApiGetMarketMakerProtectionConfigRequest) RecvWindow(recvWindow int64) ApiGetMarketMakerProtectionConfigRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetMarketMakerProtectionConfigRequest) Execute() (*common.RestApiResponse[models.GetMarketMakerProtectionConfigResponse], error) {
	return r.ApiService.GetMarketMakerProtectionConfigExecute(r)
}

/*
GetMarketMakerProtectionConfig Get Market Maker Protection Config (TRADE)
Get /eapi/v1/mmp

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Get-Market-Maker-Protection-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  underlying, e.g BTCUSDT
@param recvWindow -
@return ApiGetMarketMakerProtectionConfigRequest
*/
func (a *MarketMakerEndpointsAPIService) GetMarketMakerProtectionConfig(ctx context.Context) ApiGetMarketMakerProtectionConfigRequest {
	return ApiGetMarketMakerProtectionConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMarketMakerProtectionConfigResponse
func (a *MarketMakerEndpointsAPIService) GetMarketMakerProtectionConfigExecute(r ApiGetMarketMakerProtectionConfigRequest) (*common.RestApiResponse[models.GetMarketMakerProtectionConfigResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/mmp"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetMarketMakerProtectionConfigResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiResetMarketMakerProtectionConfigRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

// underlying, e.g BTCUSDT
func (r ApiResetMarketMakerProtectionConfigRequest) Underlying(underlying string) ApiResetMarketMakerProtectionConfigRequest {
	r.underlying = &underlying
	return r
}

func (r ApiResetMarketMakerProtectionConfigRequest) RecvWindow(recvWindow int64) ApiResetMarketMakerProtectionConfigRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiResetMarketMakerProtectionConfigRequest) Execute() (*common.RestApiResponse[models.ResetMarketMakerProtectionConfigResponse], error) {
	return r.ApiService.ResetMarketMakerProtectionConfigExecute(r)
}

/*
ResetMarketMakerProtectionConfig Reset Market Maker Protection Config (TRADE)
Post /eapi/v1/mmpReset

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Reset-Market-Maker-Protection-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  underlying, e.g BTCUSDT
@param recvWindow -
@return ApiResetMarketMakerProtectionConfigRequest
*/
func (a *MarketMakerEndpointsAPIService) ResetMarketMakerProtectionConfig(ctx context.Context) ApiResetMarketMakerProtectionConfigRequest {
	return ApiResetMarketMakerProtectionConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResetMarketMakerProtectionConfigResponse
func (a *MarketMakerEndpointsAPIService) ResetMarketMakerProtectionConfigExecute(r ApiResetMarketMakerProtectionConfigRequest) (*common.RestApiResponse[models.ResetMarketMakerProtectionConfigResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/mmpReset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ResetMarketMakerProtectionConfigResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetAutoCancelAllOpenOrdersRequest struct {
	ctx           context.Context
	ApiService    *MarketMakerEndpointsAPIService
	underlying    *string
	countdownTime *int64
	recvWindow    *int64
}

// Option underlying, e.g BTCUSDT
func (r ApiSetAutoCancelAllOpenOrdersRequest) Underlying(underlying string) ApiSetAutoCancelAllOpenOrdersRequest {
	r.underlying = &underlying
	return r
}

// Countdown time in milliseconds (ex. 1,000 for 1 second). 0 to disable the timer. Negative values (ex. -10000) are not accepted. Minimum acceptable value is 5,000
func (r ApiSetAutoCancelAllOpenOrdersRequest) CountdownTime(countdownTime int64) ApiSetAutoCancelAllOpenOrdersRequest {
	r.countdownTime = &countdownTime
	return r
}

func (r ApiSetAutoCancelAllOpenOrdersRequest) RecvWindow(recvWindow int64) ApiSetAutoCancelAllOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetAutoCancelAllOpenOrdersRequest) Execute() (*common.RestApiResponse[models.SetAutoCancelAllOpenOrdersResponse], error) {
	return r.ApiService.SetAutoCancelAllOpenOrdersExecute(r)
}

/*
SetAutoCancelAllOpenOrders Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
Post /eapi/v1/countdownCancelAll

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Set-Auto-Cancel-All-Open-Orders-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Option underlying, e.g BTCUSDT
@param countdownTime -  Countdown time in milliseconds (ex. 1,000 for 1 second). 0 to disable the timer. Negative values (ex. -10000) are not accepted. Minimum acceptable value is 5,000
@param recvWindow -
@return ApiSetAutoCancelAllOpenOrdersRequest
*/
func (a *MarketMakerEndpointsAPIService) SetAutoCancelAllOpenOrders(ctx context.Context) ApiSetAutoCancelAllOpenOrdersRequest {
	return ApiSetAutoCancelAllOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetAutoCancelAllOpenOrdersResponse
func (a *MarketMakerEndpointsAPIService) SetAutoCancelAllOpenOrdersExecute(r ApiSetAutoCancelAllOpenOrdersRequest) (*common.RestApiResponse[models.SetAutoCancelAllOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/countdownCancelAll"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}
	if r.countdownTime == nil {
		return nil, common.ReportError("countdownTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "countdownTime", r.countdownTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetAutoCancelAllOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetMarketMakerProtectionConfigRequest struct {
	ctx                      context.Context
	ApiService               *MarketMakerEndpointsAPIService
	underlying               *string
	windowTimeInMilliseconds *int64
	frozenTimeInMilliseconds *int64
	qtyLimit                 *float32
	deltaLimit               *float32
	recvWindow               *int64
}

// underlying, e.g BTCUSDT
func (r ApiSetMarketMakerProtectionConfigRequest) Underlying(underlying string) ApiSetMarketMakerProtectionConfigRequest {
	r.underlying = &underlying
	return r
}

// MMP Interval in milliseconds; Range (0,5000]
func (r ApiSetMarketMakerProtectionConfigRequest) WindowTimeInMilliseconds(windowTimeInMilliseconds int64) ApiSetMarketMakerProtectionConfigRequest {
	r.windowTimeInMilliseconds = &windowTimeInMilliseconds
	return r
}

// MMP frozen time in milliseconds, if set to 0 manual reset is required
func (r ApiSetMarketMakerProtectionConfigRequest) FrozenTimeInMilliseconds(frozenTimeInMilliseconds int64) ApiSetMarketMakerProtectionConfigRequest {
	r.frozenTimeInMilliseconds = &frozenTimeInMilliseconds
	return r
}

// quantity limit
func (r ApiSetMarketMakerProtectionConfigRequest) QtyLimit(qtyLimit float32) ApiSetMarketMakerProtectionConfigRequest {
	r.qtyLimit = &qtyLimit
	return r
}

// net delta limit
func (r ApiSetMarketMakerProtectionConfigRequest) DeltaLimit(deltaLimit float32) ApiSetMarketMakerProtectionConfigRequest {
	r.deltaLimit = &deltaLimit
	return r
}

func (r ApiSetMarketMakerProtectionConfigRequest) RecvWindow(recvWindow int64) ApiSetMarketMakerProtectionConfigRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetMarketMakerProtectionConfigRequest) Execute() (*common.RestApiResponse[models.SetMarketMakerProtectionConfigResponse], error) {
	return r.ApiService.SetMarketMakerProtectionConfigExecute(r)
}

/*
SetMarketMakerProtectionConfig Set Market Maker Protection Config (TRADE)
Post /eapi/v1/mmpSet

https://developers.binance.com/docs/derivatives/option/market-maker-endpoints/Set-Market-Maker-Protection-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  underlying, e.g BTCUSDT
@param windowTimeInMilliseconds -  MMP Interval in milliseconds; Range (0,5000]
@param frozenTimeInMilliseconds -  MMP frozen time in milliseconds, if set to 0 manual reset is required
@param qtyLimit -  quantity limit
@param deltaLimit -  net delta limit
@param recvWindow -
@return ApiSetMarketMakerProtectionConfigRequest
*/
func (a *MarketMakerEndpointsAPIService) SetMarketMakerProtectionConfig(ctx context.Context) ApiSetMarketMakerProtectionConfigRequest {
	return ApiSetMarketMakerProtectionConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetMarketMakerProtectionConfigResponse
func (a *MarketMakerEndpointsAPIService) SetMarketMakerProtectionConfigExecute(r ApiSetMarketMakerProtectionConfigRequest) (*common.RestApiResponse[models.SetMarketMakerProtectionConfigResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/mmpSet"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	}
	if r.windowTimeInMilliseconds != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "windowTimeInMilliseconds", r.windowTimeInMilliseconds, "form", "")
	}
	if r.frozenTimeInMilliseconds != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "frozenTimeInMilliseconds", r.frozenTimeInMilliseconds, "form", "")
	}
	if r.qtyLimit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "qtyLimit", r.qtyLimit, "form", "")
	}
	if r.deltaLimit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "deltaLimit", r.deltaLimit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetMarketMakerProtectionConfigResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
