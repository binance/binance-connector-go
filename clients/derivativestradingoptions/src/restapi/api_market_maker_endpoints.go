/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketMakerEndpointsAPIService MarketMakerEndpointsAPI Service
type MarketMakerEndpointsAPIService Service

type ApiAutoCancelAllOpenOrdersRequest struct {
	ctx         context.Context
	ApiService  *MarketMakerEndpointsAPIService
	underlyings *string
	recvWindow  *int64
}

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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#auto-cancel-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlyings -
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

	resp, err := SendRequest[models.AutoCancelAllOpenOrdersResponse](
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

type ApiGetAutoCancelAllOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

// Underlying asset.
func (r ApiGetAutoCancelAllOpenOrdersRequest) Underlying(underlying string) ApiGetAutoCancelAllOpenOrdersRequest {
	r.underlying = &underlying
	return r
}

// Recv Window.
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#get-auto-cancel-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Underlying asset.
@param recvWindow -  Recv Window.
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

	resp, err := SendRequest[models.GetAutoCancelAllOpenOrdersResponse](
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

type ApiGetMarketMakerProtectionConfigRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

// Underlying asset.
func (r ApiGetMarketMakerProtectionConfigRequest) Underlying(underlying string) ApiGetMarketMakerProtectionConfigRequest {
	r.underlying = &underlying
	return r
}

// Recv Window.
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#get-market-maker-protection-config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Underlying asset.
@param recvWindow -  Recv Window.
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

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetMarketMakerProtectionConfigResponse](
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

type ApiResetMarketMakerProtectionConfigRequest struct {
	ctx        context.Context
	ApiService *MarketMakerEndpointsAPIService
	underlying *string
	recvWindow *int64
}

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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#reset-market-maker-protection-config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -
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

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ResetMarketMakerProtectionConfigResponse](
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

type ApiSetAutoCancelAllOpenOrdersRequest struct {
	ctx           context.Context
	ApiService    *MarketMakerEndpointsAPIService
	underlying    *string
	countdownTime *int64
	recvWindow    *int64
}

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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#set-auto-cancel-all-open-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -
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

	resp, err := SendRequest[models.SetAutoCancelAllOpenOrdersResponse](
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

func (r ApiSetMarketMakerProtectionConfigRequest) Underlying(underlying string) ApiSetMarketMakerProtectionConfigRequest {
	r.underlying = &underlying
	return r
}

// MMP Interval in milliseconds
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-options/api/rest-api/market-maker-endpoints#set-market-maker-protection-config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -
@param windowTimeInMilliseconds -  MMP Interval in milliseconds
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

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	if r.windowTimeInMilliseconds == nil {
		return nil, common.ReportError("windowTimeInMilliseconds is required and must be specified")
	}
	if *r.windowTimeInMilliseconds < 0 {
		return nil, common.ReportError("windowTimeInMilliseconds must be greater than 0")
	}
	if *r.windowTimeInMilliseconds > 5000 {
		return nil, common.ReportError("windowTimeInMilliseconds must be less than 5000")
	}

	if r.frozenTimeInMilliseconds == nil {
		return nil, common.ReportError("frozenTimeInMilliseconds is required and must be specified")
	}

	if r.qtyLimit == nil {
		return nil, common.ReportError("qtyLimit is required and must be specified")
	}

	if r.deltaLimit == nil {
		return nil, common.ReportError("deltaLimit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "windowTimeInMilliseconds", r.windowTimeInMilliseconds, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "frozenTimeInMilliseconds", r.frozenTimeInMilliseconds, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "qtyLimit", r.qtyLimit, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "deltaLimit", r.deltaLimit, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetMarketMakerProtectionConfigResponse](
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
