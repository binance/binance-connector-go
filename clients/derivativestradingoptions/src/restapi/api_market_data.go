/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiCheckServerTimeRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiCheckServerTimeRequest) Execute() (*common.RestApiResponse[models.CheckServerTimeResponse], error) {
	return r.ApiService.CheckServerTimeExecute(r)
}

/*
CheckServerTime Check Server Time
Get /eapi/v1/time

https://developers.binance.com/docs/derivatives/option/market-data/Check-Server-Time

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCheckServerTimeRequest
*/
func (a *MarketDataAPIService) CheckServerTime(ctx context.Context) ApiCheckServerTimeRequest {
	return ApiCheckServerTimeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CheckServerTimeResponse
func (a *MarketDataAPIService) CheckServerTimeExecute(r ApiCheckServerTimeRequest) (*common.RestApiResponse[models.CheckServerTimeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/time"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.CheckServerTimeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiExchangeInformationRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiExchangeInformationRequest) Execute() (*common.RestApiResponse[models.ExchangeInformationResponse], error) {
	return r.ApiService.ExchangeInformationExecute(r)
}

/*
ExchangeInformation Exchange Information
Get /eapi/v1/exchangeInfo

https://developers.binance.com/docs/derivatives/option/market-data/Exchange-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiExchangeInformationRequest
*/
func (a *MarketDataAPIService) ExchangeInformation(ctx context.Context) ApiExchangeInformationRequest {
	return ApiExchangeInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExchangeInformationResponse
func (a *MarketDataAPIService) ExchangeInformationExecute(r ApiExchangeInformationRequest) (*common.RestApiResponse[models.ExchangeInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.ExchangeInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiHistoricalExerciseRecordsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	underlying *string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// underlying, e.g BTCUSDT
func (r ApiHistoricalExerciseRecordsRequest) Underlying(underlying string) ApiHistoricalExerciseRecordsRequest {
	r.underlying = &underlying
	return r
}

// Start Time, e.g 1593511200000
func (r ApiHistoricalExerciseRecordsRequest) StartTime(startTime int64) ApiHistoricalExerciseRecordsRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiHistoricalExerciseRecordsRequest) EndTime(endTime int64) ApiHistoricalExerciseRecordsRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiHistoricalExerciseRecordsRequest) Limit(limit int64) ApiHistoricalExerciseRecordsRequest {
	r.limit = &limit
	return r
}

func (r ApiHistoricalExerciseRecordsRequest) Execute() (*common.RestApiResponse[models.HistoricalExerciseRecordsResponse], error) {
	return r.ApiService.HistoricalExerciseRecordsExecute(r)
}

/*
HistoricalExerciseRecords Historical Exercise Records
Get /eapi/v1/exerciseHistory

https://developers.binance.com/docs/derivatives/option/market-data/Historical-Exercise-Records

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  underlying, e.g BTCUSDT
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiHistoricalExerciseRecordsRequest
*/
func (a *MarketDataAPIService) HistoricalExerciseRecords(ctx context.Context) ApiHistoricalExerciseRecordsRequest {
	return ApiHistoricalExerciseRecordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HistoricalExerciseRecordsResponse
func (a *MarketDataAPIService) HistoricalExerciseRecordsExecute(r ApiHistoricalExerciseRecordsRequest) (*common.RestApiResponse[models.HistoricalExerciseRecordsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/exerciseHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")
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

	resp, err := SendRequest[models.HistoricalExerciseRecordsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiIndexPriceTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	underlying *string
}

// Option underlying, e.g BTCUSDT
func (r ApiIndexPriceTickerRequest) Underlying(underlying string) ApiIndexPriceTickerRequest {
	r.underlying = &underlying
	return r
}

func (r ApiIndexPriceTickerRequest) Execute() (*common.RestApiResponse[models.IndexPriceTickerResponse], error) {
	return r.ApiService.IndexPriceTickerExecute(r)
}

/*
IndexPriceTicker Index Price Ticker
Get /eapi/v1/index

https://developers.binance.com/docs/derivatives/option/market-data/Index-Price-Ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlying -  Option underlying, e.g BTCUSDT
@return ApiIndexPriceTickerRequest
*/
func (a *MarketDataAPIService) IndexPriceTicker(ctx context.Context) ApiIndexPriceTickerRequest {
	return ApiIndexPriceTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IndexPriceTickerResponse
func (a *MarketDataAPIService) IndexPriceTickerExecute(r ApiIndexPriceTickerRequest) (*common.RestApiResponse[models.IndexPriceTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/index"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlying == nil {
		return nil, common.ReportError("underlying is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlying", r.underlying, "form", "")

	resp, err := SendRequest[models.IndexPriceTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiKlineCandlestickDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	interval   *string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiKlineCandlestickDataRequest) Symbol(symbol string) ApiKlineCandlestickDataRequest {
	r.symbol = &symbol
	return r
}

// Time interval
func (r ApiKlineCandlestickDataRequest) Interval(interval string) ApiKlineCandlestickDataRequest {
	r.interval = &interval
	return r
}

// Start Time, e.g 1593511200000
func (r ApiKlineCandlestickDataRequest) StartTime(startTime int64) ApiKlineCandlestickDataRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiKlineCandlestickDataRequest) EndTime(endTime int64) ApiKlineCandlestickDataRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiKlineCandlestickDataRequest) Limit(limit int64) ApiKlineCandlestickDataRequest {
	r.limit = &limit
	return r
}

func (r ApiKlineCandlestickDataRequest) Execute() (*common.RestApiResponse[models.KlineCandlestickDataResponse], error) {
	return r.ApiService.KlineCandlestickDataExecute(r)
}

/*
KlineCandlestickData Kline/Candlestick Data
Get /eapi/v1/klines

https://developers.binance.com/docs/derivatives/option/market-data/Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param interval -  Time interval
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiKlineCandlestickDataRequest
*/
func (a *MarketDataAPIService) KlineCandlestickData(ctx context.Context) ApiKlineCandlestickDataRequest {
	return ApiKlineCandlestickDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickDataResponse
func (a *MarketDataAPIService) KlineCandlestickDataExecute(r ApiKlineCandlestickDataRequest) (*common.RestApiResponse[models.KlineCandlestickDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/klines"

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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.KlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOldTradesLookupRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	fromId     *int64
	limit      *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiOldTradesLookupRequest) Symbol(symbol string) ApiOldTradesLookupRequest {
	r.symbol = &symbol
	return r
}

// The UniqueId ID from which to return. The latest deal record is returned by default
func (r ApiOldTradesLookupRequest) FromId(fromId int64) ApiOldTradesLookupRequest {
	r.fromId = &fromId
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiOldTradesLookupRequest) Limit(limit int64) ApiOldTradesLookupRequest {
	r.limit = &limit
	return r
}

func (r ApiOldTradesLookupRequest) Execute() (*common.RestApiResponse[models.OldTradesLookupResponse], error) {
	return r.ApiService.OldTradesLookupExecute(r)
}

/*
OldTradesLookup Old Trades Lookup (MARKET_DATA)
Get /eapi/v1/historicalTrades

https://developers.binance.com/docs/derivatives/option/market-data/Old-Trades-Lookup

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param fromId -  The UniqueId ID from which to return. The latest deal record is returned by default
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiOldTradesLookupRequest
*/
func (a *MarketDataAPIService) OldTradesLookup(ctx context.Context) ApiOldTradesLookupRequest {
	return ApiOldTradesLookupRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OldTradesLookupResponse
func (a *MarketDataAPIService) OldTradesLookupExecute(r ApiOldTradesLookupRequest) (*common.RestApiResponse[models.OldTradesLookupResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/historicalTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.OldTradesLookupResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOpenInterestRequest struct {
	ctx             context.Context
	ApiService      *MarketDataAPIService
	underlyingAsset *string
	expiration      *string
}

// underlying asset, e.g ETH/BTC
func (r ApiOpenInterestRequest) UnderlyingAsset(underlyingAsset string) ApiOpenInterestRequest {
	r.underlyingAsset = &underlyingAsset
	return r
}

// expiration date, e.g 221225
func (r ApiOpenInterestRequest) Expiration(expiration string) ApiOpenInterestRequest {
	r.expiration = &expiration
	return r
}

func (r ApiOpenInterestRequest) Execute() (*common.RestApiResponse[models.OpenInterestResponse], error) {
	return r.ApiService.OpenInterestExecute(r)
}

/*
OpenInterest Open Interest
Get /eapi/v1/openInterest

https://developers.binance.com/docs/derivatives/option/market-data/Open-Interest

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param underlyingAsset -  underlying asset, e.g ETH/BTC
@param expiration -  expiration date, e.g 221225
@return ApiOpenInterestRequest
*/
func (a *MarketDataAPIService) OpenInterest(ctx context.Context) ApiOpenInterestRequest {
	return ApiOpenInterestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OpenInterestResponse
func (a *MarketDataAPIService) OpenInterestExecute(r ApiOpenInterestRequest) (*common.RestApiResponse[models.OpenInterestResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/openInterest"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.underlyingAsset == nil {
		return nil, common.ReportError("underlyingAsset is required and must be specified")
	}
	if r.expiration == nil {
		return nil, common.ReportError("expiration is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "underlyingAsset", r.underlyingAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "expiration", r.expiration, "form", "")

	resp, err := SendRequest[models.OpenInterestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOptionMarkPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiOptionMarkPriceRequest) Symbol(symbol string) ApiOptionMarkPriceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOptionMarkPriceRequest) Execute() (*common.RestApiResponse[models.OptionMarkPriceResponse], error) {
	return r.ApiService.OptionMarkPriceExecute(r)
}

/*
OptionMarkPrice Option Mark Price
Get /eapi/v1/mark

https://developers.binance.com/docs/derivatives/option/market-data/Option-Mark-Price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@return ApiOptionMarkPriceRequest
*/
func (a *MarketDataAPIService) OptionMarkPrice(ctx context.Context) ApiOptionMarkPriceRequest {
	return ApiOptionMarkPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptionMarkPriceResponse
func (a *MarketDataAPIService) OptionMarkPriceExecute(r ApiOptionMarkPriceRequest) (*common.RestApiResponse[models.OptionMarkPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/mark"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.OptionMarkPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderBookRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	limit      *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiOrderBookRequest) Symbol(symbol string) ApiOrderBookRequest {
	r.symbol = &symbol
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiOrderBookRequest) Limit(limit int64) ApiOrderBookRequest {
	r.limit = &limit
	return r
}

func (r ApiOrderBookRequest) Execute() (*common.RestApiResponse[models.OrderBookResponse], error) {
	return r.ApiService.OrderBookExecute(r)
}

/*
OrderBook Order Book
Get /eapi/v1/depth

https://developers.binance.com/docs/derivatives/option/market-data/Order-Book

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiOrderBookRequest
*/
func (a *MarketDataAPIService) OrderBook(ctx context.Context) ApiOrderBookRequest {
	return ApiOrderBookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderBookResponse
func (a *MarketDataAPIService) OrderBookExecute(r ApiOrderBookRequest) (*common.RestApiResponse[models.OrderBookResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/depth"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.OrderBookResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRecentBlockTradesListRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	limit      *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiRecentBlockTradesListRequest) Symbol(symbol string) ApiRecentBlockTradesListRequest {
	r.symbol = &symbol
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiRecentBlockTradesListRequest) Limit(limit int64) ApiRecentBlockTradesListRequest {
	r.limit = &limit
	return r
}

func (r ApiRecentBlockTradesListRequest) Execute() (*common.RestApiResponse[models.RecentBlockTradesListResponse], error) {
	return r.ApiService.RecentBlockTradesListExecute(r)
}

/*
RecentBlockTradesList Recent Block Trades List
Get /eapi/v1/blockTrades

https://developers.binance.com/docs/derivatives/option/market-data/Recent-Block-Trade-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiRecentBlockTradesListRequest
*/
func (a *MarketDataAPIService) RecentBlockTradesList(ctx context.Context) ApiRecentBlockTradesListRequest {
	return ApiRecentBlockTradesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RecentBlockTradesListResponse
func (a *MarketDataAPIService) RecentBlockTradesListExecute(r ApiRecentBlockTradesListRequest) (*common.RestApiResponse[models.RecentBlockTradesListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/blockTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.RecentBlockTradesListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRecentTradesListRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	limit      *int64
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiRecentTradesListRequest) Symbol(symbol string) ApiRecentTradesListRequest {
	r.symbol = &symbol
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiRecentTradesListRequest) Limit(limit int64) ApiRecentTradesListRequest {
	r.limit = &limit
	return r
}

func (r ApiRecentTradesListRequest) Execute() (*common.RestApiResponse[models.RecentTradesListResponse], error) {
	return r.ApiService.RecentTradesListExecute(r)
}

/*
RecentTradesList Recent Trades List
Get /eapi/v1/trades

https://developers.binance.com/docs/derivatives/option/market-data/Recent-Trades-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@param limit -  Number of result sets returned Default:100 Max:1000
@return ApiRecentTradesListRequest
*/
func (a *MarketDataAPIService) RecentTradesList(ctx context.Context) ApiRecentTradesListRequest {
	return ApiRecentTradesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RecentTradesListResponse
func (a *MarketDataAPIService) RecentTradesListExecute(r ApiRecentTradesListRequest) (*common.RestApiResponse[models.RecentTradesListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/trades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.RecentTradesListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTestConnectivityRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiTestConnectivityRequest) Execute() (struct{}, error) {
	return r.ApiService.TestConnectivityExecute(r)
}

/*
TestConnectivity Test Connectivity
Get /eapi/v1/ping

https://developers.binance.com/docs/derivatives/option/market-data/Test-Connectivity

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiTestConnectivityRequest
*/
func (a *MarketDataAPIService) TestConnectivity(ctx context.Context) ApiTestConnectivityRequest {
	return ApiTestConnectivityRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MarketDataAPIService) TestConnectivityExecute(r ApiTestConnectivityRequest) (struct{}, error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/ping"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}

type ApiTicker24hrPriceChangeStatisticsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// Option trading pair, e.g BTC-200730-9000-C
func (r ApiTicker24hrPriceChangeStatisticsRequest) Symbol(symbol string) ApiTicker24hrPriceChangeStatisticsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTicker24hrPriceChangeStatisticsRequest) Execute() (*common.RestApiResponse[models.Ticker24hrPriceChangeStatisticsResponse], error) {
	return r.ApiService.Ticker24hrPriceChangeStatisticsExecute(r)
}

/*
Ticker24hrPriceChangeStatistics 24hr Ticker Price Change Statistics
Get /eapi/v1/ticker

https://developers.binance.com/docs/derivatives/option/market-data/24hr-Ticker-Price-Change-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Option trading pair, e.g BTC-200730-9000-C
@return ApiTicker24hrPriceChangeStatisticsRequest
*/
func (a *MarketDataAPIService) Ticker24hrPriceChangeStatistics(ctx context.Context) ApiTicker24hrPriceChangeStatisticsRequest {
	return ApiTicker24hrPriceChangeStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Ticker24hrPriceChangeStatisticsResponse
func (a *MarketDataAPIService) Ticker24hrPriceChangeStatisticsExecute(r ApiTicker24hrPriceChangeStatisticsRequest) (*common.RestApiResponse[models.Ticker24hrPriceChangeStatisticsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/ticker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.Ticker24hrPriceChangeStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
