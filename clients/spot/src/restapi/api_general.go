/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// GeneralAPIService GeneralAPI Service
type GeneralAPIService Service

type ApiExchangeInfoRequest struct {
	ctx                context.Context
	ApiService         *GeneralAPIService
	symbol             *string
	symbols            *[]string
	permissions        *[]models.ExchangeInfoPermissionsParameterInner
	showPermissionSets *bool
	symbolStatus       *models.ExchangeInfoSymbolStatusParameter
}

// Example: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?symbol&#x3D;BNBBTC\&quot;
func (r ApiExchangeInfoRequest) Symbol(symbol string) ApiExchangeInfoRequest {
	r.symbol = &symbol
	return r
}

// Examples: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;%5B%22BNBBTC%22,%22BTCUSDT%22%5D\&quot; or curl -g -X GET &#39;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;[\&quot;BTCUSDT\&quot;,\&quot;BNBBTC\&quot;]&#39;
func (r ApiExchangeInfoRequest) Symbols(symbols []string) ApiExchangeInfoRequest {
	r.symbols = &symbols
	return r
}

// Examples: curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;SPOT\&quot;  curl -X GET \&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\&quot; or curl -g -X GET &#39;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;[\&quot;MARGIN\&quot;,\&quot;LEVERAGED\&quot;]&#39;
func (r ApiExchangeInfoRequest) Permissions(permissions []models.ExchangeInfoPermissionsParameterInner) ApiExchangeInfoRequest {
	r.permissions = &permissions
	return r
}

// Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not.
func (r ApiExchangeInfoRequest) ShowPermissionSets(showPermissionSets bool) ApiExchangeInfoRequest {
	r.showPermissionSets = &showPermissionSets
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. Cannot be used in combination with &#x60;symbols&#x60; or &#x60;symbol&#x60;.
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

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/rest-api/general#exchange-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Example: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?symbol=BNBBTC\"
@param symbols -  Examples: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?symbols=%5B%22BNBBTC%22,%22BTCUSDT%22%5D\" or curl -g -X GET 'https://api.binance.com/api/v3/exchangeInfo?symbols=[\"BTCUSDT\",\"BNBBTC\"]'
@param permissions -  Examples: curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?permissions=SPOT\"  curl -X GET \"https://api.binance.com/api/v3/exchangeInfo?permissions=%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\" or curl -g -X GET 'https://api.binance.com/api/v3/exchangeInfo?permissions=[\"MARGIN\",\"LEVERAGED\"]'
@param showPermissionSets -  Controls whether the content of the `permissionSets` field is populated or not.
@param symbolStatus -  Filters for symbols that have this `tradingStatus`. Cannot be used in combination with `symbols` or `symbol`.
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

	resp, err := SendRequest[models.ExchangeInfoResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiExecutionRulesRequest struct {
	ctx          context.Context
	ApiService   *GeneralAPIService
	symbol       *string
	symbols      *[]string
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Query for specified symbol.
func (r ApiExecutionRulesRequest) Symbol(symbol string) ApiExecutionRulesRequest {
	r.symbol = &symbol
	return r
}

// Query for multiple symbols.
func (r ApiExecutionRulesRequest) Symbols(symbols []string) ApiExecutionRulesRequest {
	r.symbols = &symbols
	return r
}

// Query for all symbols with the specified status.
func (r ApiExecutionRulesRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiExecutionRulesRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiExecutionRulesRequest) Execute() (*common.RestApiResponse[models.ExecutionRulesResponse], error) {
	return r.ApiService.ExecutionRulesExecute(r)
}

/*
ExecutionRules Query Execution Rules
Get /api/v3/executionRules

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/rest-api/general#execution-rules

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Query for specified symbol.
@param symbols -  Query for multiple symbols.
@param symbolStatus -  Query for all symbols with the specified status.
@return ApiExecutionRulesRequest
*/
func (a *GeneralAPIService) ExecutionRules(ctx context.Context) ApiExecutionRulesRequest {
	return ApiExecutionRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExecutionRulesResponse
func (a *GeneralAPIService) ExecutionRulesExecute(r ApiExecutionRulesRequest) (*common.RestApiResponse[models.ExecutionRulesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/executionRules"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.symbols != nil {
		t := *r.symbols
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", t, "form", "multi")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.ExecutionRulesResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
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

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/rest-api/general#ping

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

	_, err := SendRequest[struct{}](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
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

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/rest-api/general#time

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

	resp, err := SendRequest[models.TimeResponse](
		r.ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarQueryParams,
		localVarBodyParameters,
		a.client.cfg,
		false,
	)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
