/*
Binance Copy Trading REST API

OpenAPI Specification for the Binance Copy Trading REST API

API version: 1.0.0
*/

package binancecopytradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/copytrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// FutureCopyTradingAPIService FutureCopyTradingAPI Service
type FutureCopyTradingAPIService Service

type ApiGetFuturesLeadTraderStatusRequest struct {
	ctx        context.Context
	ApiService *FutureCopyTradingAPIService
	recvWindow *int64
}

func (r ApiGetFuturesLeadTraderStatusRequest) RecvWindow(recvWindow int64) ApiGetFuturesLeadTraderStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesLeadTraderStatusRequest) Execute() (*common.RestApiResponse[models.GetFuturesLeadTraderStatusResponse], error) {
	return r.ApiService.GetFuturesLeadTraderStatusExecute(r)
}

/*
GetFuturesLeadTraderStatus Get Futures Lead Trader Status(TRADE)
Get /sapi/v1/copyTrading/futures/userStatus

https://developers.binance.com/docs/copy_trading/future-copy-trading/Get-Futures-Lead-Trader-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetFuturesLeadTraderStatusRequest
*/
func (a *FutureCopyTradingAPIService) GetFuturesLeadTraderStatus(ctx context.Context) ApiGetFuturesLeadTraderStatusRequest {
	return ApiGetFuturesLeadTraderStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesLeadTraderStatusResponse
func (a *FutureCopyTradingAPIService) GetFuturesLeadTraderStatusExecute(r ApiGetFuturesLeadTraderStatusRequest) (*common.RestApiResponse[models.GetFuturesLeadTraderStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/copyTrading/futures/userStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesLeadTraderStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFuturesLeadTradingSymbolWhitelistRequest struct {
	ctx        context.Context
	ApiService *FutureCopyTradingAPIService
	recvWindow *int64
}

func (r ApiGetFuturesLeadTradingSymbolWhitelistRequest) RecvWindow(recvWindow int64) ApiGetFuturesLeadTradingSymbolWhitelistRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFuturesLeadTradingSymbolWhitelistRequest) Execute() (*common.RestApiResponse[models.GetFuturesLeadTradingSymbolWhitelistResponse], error) {
	return r.ApiService.GetFuturesLeadTradingSymbolWhitelistExecute(r)
}

/*
GetFuturesLeadTradingSymbolWhitelist Get Futures Lead Trading Symbol Whitelist(USER_DATA)
Get /sapi/v1/copyTrading/futures/leadSymbol

https://developers.binance.com/docs/copy_trading/future-copy-trading/Get-Futures-Lead-Trading-Symbol-Whitelist

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetFuturesLeadTradingSymbolWhitelistRequest
*/
func (a *FutureCopyTradingAPIService) GetFuturesLeadTradingSymbolWhitelist(ctx context.Context) ApiGetFuturesLeadTradingSymbolWhitelistRequest {
	return ApiGetFuturesLeadTradingSymbolWhitelistRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFuturesLeadTradingSymbolWhitelistResponse
func (a *FutureCopyTradingAPIService) GetFuturesLeadTradingSymbolWhitelistExecute(r ApiGetFuturesLeadTradingSymbolWhitelistRequest) (*common.RestApiResponse[models.GetFuturesLeadTradingSymbolWhitelistResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/copyTrading/futures/leadSymbol"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFuturesLeadTradingSymbolWhitelistResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
