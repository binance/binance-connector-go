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
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketAPIService MarketAPI Service
type MarketAPIService Service

type ApiAggTradesRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int32
}

func (r ApiAggTradesRequest) Symbol(symbol string) ApiAggTradesRequest {
	r.symbol = &symbol
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiAggTradesRequest) FromId(fromId int64) ApiAggTradesRequest {
	r.fromId = &fromId
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiAggTradesRequest) StartTime(startTime int64) ApiAggTradesRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiAggTradesRequest) EndTime(endTime int64) ApiAggTradesRequest {
	r.endTime = &endTime
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiAggTradesRequest) Limit(limit int32) ApiAggTradesRequest {
	r.limit = &limit
	return r
}

func (r ApiAggTradesRequest) Execute() (*common.RestApiResponse[models.AggTradesResponse], error) {
	return r.ApiService.AggTradesExecute(r)
}

/*
AggTrades Compressed/Aggregate trades list
Get /api/v3/aggTrades

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#compressedaggregate-trades-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param limit -  Default: 500; Maximum: 1000.
@return ApiAggTradesRequest
*/
func (a *MarketAPIService) AggTrades(ctx context.Context) ApiAggTradesRequest {
	return ApiAggTradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AggTradesResponse
func (a *MarketAPIService) AggTradesExecute(r ApiAggTradesRequest) (*common.RestApiResponse[models.AggTradesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/aggTrades"

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

	resp, err := SendRequest[models.AggTradesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAvgPriceRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
}

func (r ApiAvgPriceRequest) Symbol(symbol string) ApiAvgPriceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAvgPriceRequest) Execute() (*common.RestApiResponse[models.AvgPriceResponse], error) {
	return r.ApiService.AvgPriceExecute(r)
}

/*
AvgPrice Current average price
Get /api/v3/avgPrice

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#current-average-price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiAvgPriceRequest
*/
func (a *MarketAPIService) AvgPrice(ctx context.Context) ApiAvgPriceRequest {
	return ApiAvgPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AvgPriceResponse
func (a *MarketAPIService) AvgPriceExecute(r ApiAvgPriceRequest) (*common.RestApiResponse[models.AvgPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/avgPrice"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.AvgPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDepthRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	limit        *int32
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

func (r ApiDepthRequest) Symbol(symbol string) ApiDepthRequest {
	r.symbol = &symbol
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiDepthRequest) Limit(limit int32) ApiDepthRequest {
	r.limit = &limit
	return r
}

func (r ApiDepthRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiDepthRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiDepthRequest) Execute() (*common.RestApiResponse[models.DepthResponse], error) {
	return r.ApiService.DepthExecute(r)
}

/*
Depth Order book
Get /api/v3/depth

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#order-book

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default: 500; Maximum: 1000.
@param symbolStatus -
@return ApiDepthRequest
*/
func (a *MarketAPIService) Depth(ctx context.Context) ApiDepthRequest {
	return ApiDepthRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepthResponse
func (a *MarketAPIService) DepthExecute(r ApiDepthRequest) (*common.RestApiResponse[models.DepthResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/depth"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.DepthResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetTradesRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
	limit      *int32
}

func (r ApiGetTradesRequest) Symbol(symbol string) ApiGetTradesRequest {
	r.symbol = &symbol
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiGetTradesRequest) Limit(limit int32) ApiGetTradesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetTradesRequest) Execute() (*common.RestApiResponse[models.GetTradesResponse], error) {
	return r.ApiService.GetTradesExecute(r)
}

/*
GetTrades Recent trades list
Get /api/v3/trades

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#recent-trades-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default: 500; Maximum: 1000.
@return ApiGetTradesRequest
*/
func (a *MarketAPIService) GetTrades(ctx context.Context) ApiGetTradesRequest {
	return ApiGetTradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTradesResponse
func (a *MarketAPIService) GetTradesExecute(r ApiGetTradesRequest) (*common.RestApiResponse[models.GetTradesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/trades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.GetTradesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiHistoricalTradesRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
	limit      *int32
	fromId     *int64
}

func (r ApiHistoricalTradesRequest) Symbol(symbol string) ApiHistoricalTradesRequest {
	r.symbol = &symbol
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiHistoricalTradesRequest) Limit(limit int32) ApiHistoricalTradesRequest {
	r.limit = &limit
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiHistoricalTradesRequest) FromId(fromId int64) ApiHistoricalTradesRequest {
	r.fromId = &fromId
	return r
}

func (r ApiHistoricalTradesRequest) Execute() (*common.RestApiResponse[models.HistoricalTradesResponse], error) {
	return r.ApiService.HistoricalTradesExecute(r)
}

/*
HistoricalTrades Old trade lookup
Get /api/v3/historicalTrades

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#old-trade-lookup

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default: 500; Maximum: 1000.
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@return ApiHistoricalTradesRequest
*/
func (a *MarketAPIService) HistoricalTrades(ctx context.Context) ApiHistoricalTradesRequest {
	return ApiHistoricalTradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HistoricalTradesResponse
func (a *MarketAPIService) HistoricalTradesExecute(r ApiHistoricalTradesRequest) (*common.RestApiResponse[models.HistoricalTradesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/historicalTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}

	resp, err := SendRequest[models.HistoricalTradesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiKlinesRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
	interval   *models.KlinesIntervalParameter
	startTime  *int64
	endTime    *int64
	timeZone   *string
	limit      *int32
}

func (r ApiKlinesRequest) Symbol(symbol string) ApiKlinesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlinesRequest) Interval(interval models.KlinesIntervalParameter) ApiKlinesRequest {
	r.interval = &interval
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiKlinesRequest) StartTime(startTime int64) ApiKlinesRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiKlinesRequest) EndTime(endTime int64) ApiKlinesRequest {
	r.endTime = &endTime
	return r
}

// Default: 0 (UTC)
func (r ApiKlinesRequest) TimeZone(timeZone string) ApiKlinesRequest {
	r.timeZone = &timeZone
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiKlinesRequest) Limit(limit int32) ApiKlinesRequest {
	r.limit = &limit
	return r
}

func (r ApiKlinesRequest) Execute() (*common.RestApiResponse[models.KlinesResponse], error) {
	return r.ApiService.KlinesExecute(r)
}

/*
Klines Kline/Candlestick data
Get /api/v3/klines

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#klinecandlestick-data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param interval -
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param timeZone -  Default: 0 (UTC)
@param limit -  Default: 500; Maximum: 1000.
@return ApiKlinesRequest
*/
func (a *MarketAPIService) Klines(ctx context.Context) ApiKlinesRequest {
	return ApiKlinesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KlinesResponse
func (a *MarketAPIService) KlinesExecute(r ApiKlinesRequest) (*common.RestApiResponse[models.KlinesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/klines"

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
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.timeZone != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeZone", r.timeZone, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.KlinesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTickerRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	symbols      *[]string
	windowSize   *models.TickerWindowSizeParameter
	type_        *models.TickerTypeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiTickerRequest) Symbol(symbol string) ApiTickerRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiTickerRequest) Symbols(symbols []string) ApiTickerRequest {
	r.symbols = &symbols
	return r
}

func (r ApiTickerRequest) WindowSize(windowSize models.TickerWindowSizeParameter) ApiTickerRequest {
	r.windowSize = &windowSize
	return r
}

func (r ApiTickerRequest) Type(type_ models.TickerTypeParameter) ApiTickerRequest {
	r.type_ = &type_
	return r
}

func (r ApiTickerRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerRequest) Execute() (*common.RestApiResponse[models.TickerResponse], error) {
	return r.ApiService.TickerExecute(r)
}

/*
Ticker Rolling window price change statistics
Get /api/v3/ticker

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#rolling-window-price-change-statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param windowSize -
@param type_ -
@param symbolStatus -
@return ApiTickerRequest
*/
func (a *MarketAPIService) Ticker(ctx context.Context) ApiTickerRequest {
	return ApiTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TickerResponse
func (a *MarketAPIService) TickerExecute(r ApiTickerRequest) (*common.RestApiResponse[models.TickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ticker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.symbols != nil {
		t := *r.symbols
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", t, "form", "multi")
	}
	if r.windowSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "windowSize", r.windowSize, "form", "")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.TickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTicker24hrRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	symbols      *[]string
	type_        *models.TickerTypeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiTicker24hrRequest) Symbol(symbol string) ApiTicker24hrRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiTicker24hrRequest) Symbols(symbols []string) ApiTicker24hrRequest {
	r.symbols = &symbols
	return r
}

func (r ApiTicker24hrRequest) Type(type_ models.TickerTypeParameter) ApiTicker24hrRequest {
	r.type_ = &type_
	return r
}

func (r ApiTicker24hrRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTicker24hrRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTicker24hrRequest) Execute() (*common.RestApiResponse[models.Ticker24hrResponse], error) {
	return r.ApiService.Ticker24hrExecute(r)
}

/*
Ticker24hr 24hr ticker price change statistics
Get /api/v3/ticker/24hr

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#24hr-ticker-price-change-statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param type_ -
@param symbolStatus -
@return ApiTicker24hrRequest
*/
func (a *MarketAPIService) Ticker24hr(ctx context.Context) ApiTicker24hrRequest {
	return ApiTicker24hrRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Ticker24hrResponse
func (a *MarketAPIService) Ticker24hrExecute(r ApiTicker24hrRequest) (*common.RestApiResponse[models.Ticker24hrResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ticker/24hr"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.symbols != nil {
		t := *r.symbols
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", t, "form", "multi")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.Ticker24hrResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTickerBookTickerRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	symbols      *[]string
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiTickerBookTickerRequest) Symbol(symbol string) ApiTickerBookTickerRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiTickerBookTickerRequest) Symbols(symbols []string) ApiTickerBookTickerRequest {
	r.symbols = &symbols
	return r
}

func (r ApiTickerBookTickerRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerBookTickerRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerBookTickerRequest) Execute() (*common.RestApiResponse[models.TickerBookTickerResponse], error) {
	return r.ApiService.TickerBookTickerExecute(r)
}

/*
TickerBookTicker Symbol order book ticker
Get /api/v3/ticker/bookTicker

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#symbol-order-book-ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param symbolStatus -
@return ApiTickerBookTickerRequest
*/
func (a *MarketAPIService) TickerBookTicker(ctx context.Context) ApiTickerBookTickerRequest {
	return ApiTickerBookTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TickerBookTickerResponse
func (a *MarketAPIService) TickerBookTickerExecute(r ApiTickerBookTickerRequest) (*common.RestApiResponse[models.TickerBookTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ticker/bookTicker"

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

	resp, err := SendRequest[models.TickerBookTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTickerPriceRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	symbols      *[]string
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiTickerPriceRequest) Symbol(symbol string) ApiTickerPriceRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiTickerPriceRequest) Symbols(symbols []string) ApiTickerPriceRequest {
	r.symbols = &symbols
	return r
}

func (r ApiTickerPriceRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerPriceRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerPriceRequest) Execute() (*common.RestApiResponse[models.TickerPriceResponse], error) {
	return r.ApiService.TickerPriceExecute(r)
}

/*
TickerPrice Symbol price ticker
Get /api/v3/ticker/price

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#symbol-price-ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param symbolStatus -
@return ApiTickerPriceRequest
*/
func (a *MarketAPIService) TickerPrice(ctx context.Context) ApiTickerPriceRequest {
	return ApiTickerPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TickerPriceResponse
func (a *MarketAPIService) TickerPriceExecute(r ApiTickerPriceRequest) (*common.RestApiResponse[models.TickerPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ticker/price"

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

	resp, err := SendRequest[models.TickerPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTickerTradingDayRequest struct {
	ctx          context.Context
	ApiService   *MarketAPIService
	symbol       *string
	symbols      *[]string
	timeZone     *string
	type_        *models.TickerTypeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Symbol to query
func (r ApiTickerTradingDayRequest) Symbol(symbol string) ApiTickerTradingDayRequest {
	r.symbol = &symbol
	return r
}

// List of symbols to query
func (r ApiTickerTradingDayRequest) Symbols(symbols []string) ApiTickerTradingDayRequest {
	r.symbols = &symbols
	return r
}

// Default: 0 (UTC)
func (r ApiTickerTradingDayRequest) TimeZone(timeZone string) ApiTickerTradingDayRequest {
	r.timeZone = &timeZone
	return r
}

func (r ApiTickerTradingDayRequest) Type(type_ models.TickerTypeParameter) ApiTickerTradingDayRequest {
	r.type_ = &type_
	return r
}

func (r ApiTickerTradingDayRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerTradingDayRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerTradingDayRequest) Execute() (*common.RestApiResponse[models.TickerTradingDayResponse], error) {
	return r.ApiService.TickerTradingDayExecute(r)
}

/*
TickerTradingDay Trading Day Ticker
Get /api/v3/ticker/tradingDay

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#trading-day-ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param symbols -  List of symbols to query
@param timeZone -  Default: 0 (UTC)
@param type_ -
@param symbolStatus -
@return ApiTickerTradingDayRequest
*/
func (a *MarketAPIService) TickerTradingDay(ctx context.Context) ApiTickerTradingDayRequest {
	return ApiTickerTradingDayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TickerTradingDayResponse
func (a *MarketAPIService) TickerTradingDayExecute(r ApiTickerTradingDayRequest) (*common.RestApiResponse[models.TickerTradingDayResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/ticker/tradingDay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.symbols != nil {
		t := *r.symbols
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbols", t, "form", "multi")
	}
	if r.timeZone != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeZone", r.timeZone, "form", "")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.symbolStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbolStatus", r.symbolStatus, "form", "")
	}

	resp, err := SendRequest[models.TickerTradingDayResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUiKlinesRequest struct {
	ctx        context.Context
	ApiService *MarketAPIService
	symbol     *string
	interval   *models.KlinesIntervalParameter
	startTime  *int64
	endTime    *int64
	timeZone   *string
	limit      *int32
}

func (r ApiUiKlinesRequest) Symbol(symbol string) ApiUiKlinesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUiKlinesRequest) Interval(interval models.KlinesIntervalParameter) ApiUiKlinesRequest {
	r.interval = &interval
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiUiKlinesRequest) StartTime(startTime int64) ApiUiKlinesRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiUiKlinesRequest) EndTime(endTime int64) ApiUiKlinesRequest {
	r.endTime = &endTime
	return r
}

// Default: 0 (UTC)
func (r ApiUiKlinesRequest) TimeZone(timeZone string) ApiUiKlinesRequest {
	r.timeZone = &timeZone
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiUiKlinesRequest) Limit(limit int32) ApiUiKlinesRequest {
	r.limit = &limit
	return r
}

func (r ApiUiKlinesRequest) Execute() (*common.RestApiResponse[models.UiKlinesResponse], error) {
	return r.ApiService.UiKlinesExecute(r)
}

/*
UiKlines UIKlines
Get /api/v3/uiKlines

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/market-data-endpoints#uiklines

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param interval -
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param timeZone -  Default: 0 (UTC)
@param limit -  Default: 500; Maximum: 1000.
@return ApiUiKlinesRequest
*/
func (a *MarketAPIService) UiKlines(ctx context.Context) ApiUiKlinesRequest {
	return ApiUiKlinesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UiKlinesResponse
func (a *MarketAPIService) UiKlinesExecute(r ApiUiKlinesRequest) (*common.RestApiResponse[models.UiKlinesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/uiKlines"

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
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.timeZone != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "timeZone", r.timeZone, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.UiKlinesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
