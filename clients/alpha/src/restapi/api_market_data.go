/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

*/

package binancealpharestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/alpha/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiAggregatedTradesRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List
func (r ApiAggregatedTradesRequest) Symbol(symbol string) ApiAggregatedTradesRequest {
	r.symbol = &symbol
	return r
}

// starting trade ID to fetch from
func (r ApiAggregatedTradesRequest) FromId(fromId int64) ApiAggregatedTradesRequest {
	r.fromId = &fromId
	return r
}

// start timestamp (milliseconds)
func (r ApiAggregatedTradesRequest) StartTime(startTime int64) ApiAggregatedTradesRequest {
	r.startTime = &startTime
	return r
}

// end timestamp (milliseconds)
func (r ApiAggregatedTradesRequest) EndTime(endTime int64) ApiAggregatedTradesRequest {
	r.endTime = &endTime
	return r
}

// number of results to return (default 500, max 1000)
func (r ApiAggregatedTradesRequest) Limit(limit int64) ApiAggregatedTradesRequest {
	r.limit = &limit
	return r
}

func (r ApiAggregatedTradesRequest) Execute() (*common.RestApiResponse[models.AggregatedTradesResponse], error) {
	return r.ApiService.AggregatedTradesExecute(r)
}

/*
AggregatedTrades Aggregated Trades
Get /bapi/defi/v1/public/alpha-trade/agg-trades

https://developers.binance.com/docs/alpha/market-data/rest-api/Aggregated-Trades

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  e.g., \"ALPHA_175USDT\" – use token ID from Token List
@param fromId -  starting trade ID to fetch from
@param startTime -  start timestamp (milliseconds)
@param endTime -  end timestamp (milliseconds)
@param limit -  number of results to return (default 500, max 1000)
@return ApiAggregatedTradesRequest
*/
func (a *MarketDataAPIService) AggregatedTrades(ctx context.Context) ApiAggregatedTradesRequest {
	return ApiAggregatedTradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AggregatedTradesResponse
func (a *MarketDataAPIService) AggregatedTradesExecute(r ApiAggregatedTradesRequest) (*common.RestApiResponse[models.AggregatedTradesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/agg-trades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
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

	resp, err := SendRequest[models.AggregatedTradesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetExchangeInfoRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiGetExchangeInfoRequest) Execute() (*common.RestApiResponse[models.GetExchangeInfoResponse], error) {
	return r.ApiService.GetExchangeInfoExecute(r)
}

/*
GetExchangeInfo Get Exchange Info
Get /bapi/defi/v1/public/alpha-trade/get-exchange-info

https://developers.binance.com/docs/alpha/market-data/rest-api/Get-Exchange-Info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiGetExchangeInfoRequest
*/
func (a *MarketDataAPIService) GetExchangeInfo(ctx context.Context) ApiGetExchangeInfoRequest {
	return ApiGetExchangeInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetExchangeInfoResponse
func (a *MarketDataAPIService) GetExchangeInfoExecute(r ApiGetExchangeInfoRequest) (*common.RestApiResponse[models.GetExchangeInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/get-exchange-info"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetExchangeInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiKlinesRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	interval   *string
	limit      *int64
	startTime  *int64
	endTime    *int64
}

// e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List
func (r ApiKlinesRequest) Symbol(symbol string) ApiKlinesRequest {
	r.symbol = &symbol
	return r
}

// e.g., \&quot;1h\&quot; – supported intervals: 1s, 15s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
func (r ApiKlinesRequest) Interval(interval string) ApiKlinesRequest {
	r.interval = &interval
	return r
}

// number of results to return (default 500, max 1000)
func (r ApiKlinesRequest) Limit(limit int64) ApiKlinesRequest {
	r.limit = &limit
	return r
}

// start timestamp (milliseconds)
func (r ApiKlinesRequest) StartTime(startTime int64) ApiKlinesRequest {
	r.startTime = &startTime
	return r
}

// end timestamp (milliseconds)
func (r ApiKlinesRequest) EndTime(endTime int64) ApiKlinesRequest {
	r.endTime = &endTime
	return r
}

func (r ApiKlinesRequest) Execute() (*common.RestApiResponse[models.KlinesResponse], error) {
	return r.ApiService.KlinesExecute(r)
}

/*
Klines Klines (Candlestick Data)
Get /bapi/defi/v1/public/alpha-trade/klines

https://developers.binance.com/docs/alpha/market-data/rest-api/Klines

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  e.g., \"ALPHA_175USDT\" – use token ID from Token List
@param interval -  e.g., \"1h\" – supported intervals: 1s, 15s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1M
@param limit -  number of results to return (default 500, max 1000)
@param startTime -  start timestamp (milliseconds)
@param endTime -  end timestamp (milliseconds)
@return ApiKlinesRequest
*/
func (a *MarketDataAPIService) Klines(ctx context.Context) ApiKlinesRequest {
	return ApiKlinesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KlinesResponse
func (a *MarketDataAPIService) KlinesExecute(r ApiKlinesRequest) (*common.RestApiResponse[models.KlinesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/klines"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.KlinesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// e.g., \&quot;ALPHA_175USDT\&quot; – use token ID from Token List
func (r ApiTickerRequest) Symbol(symbol string) ApiTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTickerRequest) Execute() (*common.RestApiResponse[models.TickerResponse], error) {
	return r.ApiService.TickerExecute(r)
}

/*
Ticker Ticker (24hr Price Statistics)
Get /bapi/defi/v1/public/alpha-trade/ticker

https://developers.binance.com/docs/alpha/market-data/rest-api/24hr-ticker-price-change

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  e.g., \"ALPHA_175USDT\" – use token ID from Token List
@return ApiTickerRequest
*/
func (a *MarketDataAPIService) Ticker(ctx context.Context) ApiTickerRequest {
	return ApiTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TickerResponse
func (a *MarketDataAPIService) TickerExecute(r ApiTickerRequest) (*common.RestApiResponse[models.TickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/ticker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.TickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTokenListRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiTokenListRequest) Execute() (*common.RestApiResponse[models.TokenListResponse], error) {
	return r.ApiService.TokenListExecute(r)
}

/*
TokenList Token List
Get /bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list

https://developers.binance.com/docs/alpha/market-data/rest-api/Token-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiTokenListRequest
*/
func (a *MarketDataAPIService) TokenList(ctx context.Context) ApiTokenListRequest {
	return ApiTokenListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenListResponse
func (a *MarketDataAPIService) TokenListExecute(r ApiTokenListRequest) (*common.RestApiResponse[models.TokenListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.TokenListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
