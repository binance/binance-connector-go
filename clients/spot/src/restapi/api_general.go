/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package binancespotrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// GeneralAPIService GeneralAPI Service
type GeneralAPIService Service

type ApiExchangeInfoRequest struct {
	ctx                context.Context
	ApiService         *GeneralAPIService
	symbol             *string
	symbols            *[]string
	permissions        *[]string
	showPermissionSets *bool
	symbolStatus       *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiExchangeInfoRequest) Symbol(symbol string) ApiExchangeInfoRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiExchangeInfoRequest) Symbols(symbols []string) ApiExchangeInfoRequest {
	r.symbols = &symbols
	return r
}

// List of permissions to query
func (r ApiExchangeInfoRequest) Permissions(permissions []string) ApiExchangeInfoRequest {
	r.permissions = &permissions
	return r
}

// Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not. Defaults to &#x60;true&#x60;
func (r ApiExchangeInfoRequest) ShowPermissionSets(showPermissionSets bool) ApiExchangeInfoRequest {
	r.showPermissionSets = &showPermissionSets
	return r
}

func (r ApiExchangeInfoRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiExchangeInfoRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiExchangeInfoRequest) Execute() (*common.RestApiResponse[models.ExchangeInfoResponse], error) {
	return r.ApiService.ExchangeInfoExecute(r)
}

/*
ExchangeInfo Exchange information
Get /api/v3/exchangeInfo

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-endpoints#exchange-information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param permissions -  List of permissions to query
@param showPermissionSets -  Controls whether the content of the `permissionSets` field is populated or not. Defaults to `true`
@param symbolStatus -
@return ApiExchangeInfoRequest
*/
func (a *GeneralAPIService) ExchangeInfo(ctx context.Context) ApiExchangeInfoRequest {
	return ApiExchangeInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExchangeInfoResponse
func (a *GeneralAPIService) ExchangeInfoExecute(r ApiExchangeInfoRequest) (*common.RestApiResponse[models.ExchangeInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.symbols != nil {
		t := *r.symbols
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", t, "form", "multi")
	}
	if r.permissions != nil {
		t := *r.permissions
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "permissions", t, "form", "multi")
	}
	if r.showPermissionSets != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "showPermissionSets", r.showPermissionSets, "form", "")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.ExchangeInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPingRequest struct {
	ctx        context.Context
	ApiService *GeneralAPIService
}

func (r ApiPingRequest) Execute() (struct{}, error) {
	return r.ApiService.PingExecute(r)
}

/*
Ping Test connectivity
Get /api/v3/ping

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-endpoints#test-connectivity

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiPingRequest
*/
func (a *GeneralAPIService) Ping(ctx context.Context) ApiPingRequest {
	return ApiPingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *GeneralAPIService) PingExecute(r ApiPingRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ping"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiTimeRequest struct {
	ctx        context.Context
	ApiService *GeneralAPIService
}

func (r ApiTimeRequest) Execute() (*common.RestApiResponse[models.TimeResponse], error) {
	return r.ApiService.TimeExecute(r)
}

/*
Time Check server time
Get /api/v3/time

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-endpoints#check-server-time

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiTimeRequest
*/
func (a *GeneralAPIService) Time(ctx context.Context) ApiTimeRequest {
	return ApiTimeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TimeResponse
func (a *GeneralAPIService) TimeExecute(r ApiTimeRequest) (*common.RestApiResponse[models.TimeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/time"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.TimeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
