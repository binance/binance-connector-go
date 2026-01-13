/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiBasisRequest struct {
	ctx          context.Context
	ApiService   *MarketDataAPIService
	pair         *string
	contractType *models.BasisContractTypeParameter
	period       *models.BasisPeriodParameter
	limit        *int64
	startTime    *int64
	endTime      *int64
}

// BTCUSD
func (r ApiBasisRequest) Pair(pair string) ApiBasisRequest {
	r.pair = &pair
	return r
}

// ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (r ApiBasisRequest) ContractType(contractType models.BasisContractTypeParameter) ApiBasisRequest {
	r.contractType = &contractType
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiBasisRequest) Period(period models.BasisPeriodParameter) ApiBasisRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiBasisRequest) Limit(limit int64) ApiBasisRequest {
	r.limit = &limit
	return r
}

func (r ApiBasisRequest) StartTime(startTime int64) ApiBasisRequest {
	r.startTime = &startTime
	return r
}

func (r ApiBasisRequest) EndTime(endTime int64) ApiBasisRequest {
	r.endTime = &endTime
	return r
}

func (r ApiBasisRequest) Execute() (*common.RestApiResponse[models.BasisResponse], error) {
	return r.ApiService.BasisExecute(r)
}

/*
Basis Basis
Get /futures/data/basis

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Basis

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param contractType -  ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiBasisRequest
*/
func (a *MarketDataAPIService) Basis(ctx context.Context) ApiBasisRequest {
	return ApiBasisRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BasisResponse
func (a *MarketDataAPIService) BasisExecute(r ApiBasisRequest) (*common.RestApiResponse[models.BasisResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/basis"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.contractType == nil {
		return nil, common.ReportError("contractType is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "contractType", r.contractType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.BasisResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCheckServerTimeRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiCheckServerTimeRequest) Execute() (*common.RestApiResponse[models.CheckServerTimeResponse], error) {
	return r.ApiService.CheckServerTimeExecute(r)
}

/*
CheckServerTime Check Server time
Get /dapi/v1/time

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Check-Server-time

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/time"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.CheckServerTimeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCompressedAggregateTradesListRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiCompressedAggregateTradesListRequest) Symbol(symbol string) ApiCompressedAggregateTradesListRequest {
	r.symbol = &symbol
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiCompressedAggregateTradesListRequest) FromId(fromId int64) ApiCompressedAggregateTradesListRequest {
	r.fromId = &fromId
	return r
}

func (r ApiCompressedAggregateTradesListRequest) StartTime(startTime int64) ApiCompressedAggregateTradesListRequest {
	r.startTime = &startTime
	return r
}

func (r ApiCompressedAggregateTradesListRequest) EndTime(endTime int64) ApiCompressedAggregateTradesListRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiCompressedAggregateTradesListRequest) Limit(limit int64) ApiCompressedAggregateTradesListRequest {
	r.limit = &limit
	return r
}

func (r ApiCompressedAggregateTradesListRequest) Execute() (*common.RestApiResponse[models.CompressedAggregateTradesListResponse], error) {
	return r.ApiService.CompressedAggregateTradesListExecute(r)
}

/*
CompressedAggregateTradesList Compressed/Aggregate Trades List
Get /dapi/v1/aggTrades

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Compressed-Aggregate-Trades-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiCompressedAggregateTradesListRequest
*/
func (a *MarketDataAPIService) CompressedAggregateTradesList(ctx context.Context) ApiCompressedAggregateTradesListRequest {
	return ApiCompressedAggregateTradesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CompressedAggregateTradesListResponse
func (a *MarketDataAPIService) CompressedAggregateTradesListExecute(r ApiCompressedAggregateTradesListRequest) (*common.RestApiResponse[models.CompressedAggregateTradesListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/aggTrades"

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

	resp, err := SendRequest[models.CompressedAggregateTradesListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiContinuousContractKlineCandlestickDataRequest struct {
	ctx          context.Context
	ApiService   *MarketDataAPIService
	pair         *string
	contractType *models.BasisContractTypeParameter
	interval     *models.ContinuousContractKlineCandlestickDataIntervalParameter
	startTime    *int64
	endTime      *int64
	limit        *int64
}

// BTCUSD
func (r ApiContinuousContractKlineCandlestickDataRequest) Pair(pair string) ApiContinuousContractKlineCandlestickDataRequest {
	r.pair = &pair
	return r
}

// ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (r ApiContinuousContractKlineCandlestickDataRequest) ContractType(contractType models.BasisContractTypeParameter) ApiContinuousContractKlineCandlestickDataRequest {
	r.contractType = &contractType
	return r
}

func (r ApiContinuousContractKlineCandlestickDataRequest) Interval(interval models.ContinuousContractKlineCandlestickDataIntervalParameter) ApiContinuousContractKlineCandlestickDataRequest {
	r.interval = &interval
	return r
}

func (r ApiContinuousContractKlineCandlestickDataRequest) StartTime(startTime int64) ApiContinuousContractKlineCandlestickDataRequest {
	r.startTime = &startTime
	return r
}

func (r ApiContinuousContractKlineCandlestickDataRequest) EndTime(endTime int64) ApiContinuousContractKlineCandlestickDataRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiContinuousContractKlineCandlestickDataRequest) Limit(limit int64) ApiContinuousContractKlineCandlestickDataRequest {
	r.limit = &limit
	return r
}

func (r ApiContinuousContractKlineCandlestickDataRequest) Execute() (*common.RestApiResponse[models.ContinuousContractKlineCandlestickDataResponse], error) {
	return r.ApiService.ContinuousContractKlineCandlestickDataExecute(r)
}

/*
ContinuousContractKlineCandlestickData Continuous Contract Kline/Candlestick Data
Get /dapi/v1/continuousKlines

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Continuous-Contract-Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param contractType -  ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
@param interval -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiContinuousContractKlineCandlestickDataRequest
*/
func (a *MarketDataAPIService) ContinuousContractKlineCandlestickData(ctx context.Context) ApiContinuousContractKlineCandlestickDataRequest {
	return ApiContinuousContractKlineCandlestickDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ContinuousContractKlineCandlestickDataResponse
func (a *MarketDataAPIService) ContinuousContractKlineCandlestickDataExecute(r ApiContinuousContractKlineCandlestickDataRequest) (*common.RestApiResponse[models.ContinuousContractKlineCandlestickDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/continuousKlines"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.contractType == nil {
		return nil, common.ReportError("contractType is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "contractType", r.contractType, "form", "")
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

	resp, err := SendRequest[models.ContinuousContractKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
Get /dapi/v1/exchangeInfo

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Exchange-Information

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.ExchangeInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFundingRateHistoryOfPerpetualFuturesRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) Symbol(symbol string) ApiGetFundingRateHistoryOfPerpetualFuturesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) StartTime(startTime int64) ApiGetFundingRateHistoryOfPerpetualFuturesRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) EndTime(endTime int64) ApiGetFundingRateHistoryOfPerpetualFuturesRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) Limit(limit int64) ApiGetFundingRateHistoryOfPerpetualFuturesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) Execute() (*common.RestApiResponse[models.GetFundingRateHistoryOfPerpetualFuturesResponse], error) {
	return r.ApiService.GetFundingRateHistoryOfPerpetualFuturesExecute(r)
}

/*
GetFundingRateHistoryOfPerpetualFutures Get Funding Rate History of Perpetual Futures
Get /dapi/v1/fundingRate

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Get-Funding-Rate-History-of-Perpetual-Futures

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiGetFundingRateHistoryOfPerpetualFuturesRequest
*/
func (a *MarketDataAPIService) GetFundingRateHistoryOfPerpetualFutures(ctx context.Context) ApiGetFundingRateHistoryOfPerpetualFuturesRequest {
	return ApiGetFundingRateHistoryOfPerpetualFuturesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFundingRateHistoryOfPerpetualFuturesResponse
func (a *MarketDataAPIService) GetFundingRateHistoryOfPerpetualFuturesExecute(r ApiGetFundingRateHistoryOfPerpetualFuturesRequest) (*common.RestApiResponse[models.GetFundingRateHistoryOfPerpetualFuturesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/fundingRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.GetFundingRateHistoryOfPerpetualFuturesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFundingRateInfoRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiGetFundingRateInfoRequest) Execute() (*common.RestApiResponse[models.GetFundingRateInfoResponse], error) {
	return r.ApiService.GetFundingRateInfoExecute(r)
}

/*
GetFundingRateInfo Get Funding Rate Info
Get /dapi/v1/fundingInfo

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Get-Funding-Info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiGetFundingRateInfoRequest
*/
func (a *MarketDataAPIService) GetFundingRateInfo(ctx context.Context) ApiGetFundingRateInfoRequest {
	return ApiGetFundingRateInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFundingRateInfoResponse
func (a *MarketDataAPIService) GetFundingRateInfoExecute(r ApiGetFundingRateInfoRequest) (*common.RestApiResponse[models.GetFundingRateInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/fundingInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetFundingRateInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiIndexPriceAndMarkPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	pair       *string
}

func (r ApiIndexPriceAndMarkPriceRequest) Symbol(symbol string) ApiIndexPriceAndMarkPriceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiIndexPriceAndMarkPriceRequest) Pair(pair string) ApiIndexPriceAndMarkPriceRequest {
	r.pair = &pair
	return r
}

func (r ApiIndexPriceAndMarkPriceRequest) Execute() (*common.RestApiResponse[models.IndexPriceAndMarkPriceResponse], error) {
	return r.ApiService.IndexPriceAndMarkPriceExecute(r)
}

/*
IndexPriceAndMarkPrice Index Price and Mark Price
Get /dapi/v1/premiumIndex

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Index-Price-and-Mark-Price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@return ApiIndexPriceAndMarkPriceRequest
*/
func (a *MarketDataAPIService) IndexPriceAndMarkPrice(ctx context.Context) ApiIndexPriceAndMarkPriceRequest {
	return ApiIndexPriceAndMarkPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IndexPriceAndMarkPriceResponse
func (a *MarketDataAPIService) IndexPriceAndMarkPriceExecute(r ApiIndexPriceAndMarkPriceRequest) (*common.RestApiResponse[models.IndexPriceAndMarkPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/premiumIndex"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}

	resp, err := SendRequest[models.IndexPriceAndMarkPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiIndexPriceKlineCandlestickDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	pair       *string
	interval   *models.ContinuousContractKlineCandlestickDataIntervalParameter
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// BTCUSD
func (r ApiIndexPriceKlineCandlestickDataRequest) Pair(pair string) ApiIndexPriceKlineCandlestickDataRequest {
	r.pair = &pair
	return r
}

func (r ApiIndexPriceKlineCandlestickDataRequest) Interval(interval models.ContinuousContractKlineCandlestickDataIntervalParameter) ApiIndexPriceKlineCandlestickDataRequest {
	r.interval = &interval
	return r
}

func (r ApiIndexPriceKlineCandlestickDataRequest) StartTime(startTime int64) ApiIndexPriceKlineCandlestickDataRequest {
	r.startTime = &startTime
	return r
}

func (r ApiIndexPriceKlineCandlestickDataRequest) EndTime(endTime int64) ApiIndexPriceKlineCandlestickDataRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiIndexPriceKlineCandlestickDataRequest) Limit(limit int64) ApiIndexPriceKlineCandlestickDataRequest {
	r.limit = &limit
	return r
}

func (r ApiIndexPriceKlineCandlestickDataRequest) Execute() (*common.RestApiResponse[models.IndexPriceKlineCandlestickDataResponse], error) {
	return r.ApiService.IndexPriceKlineCandlestickDataExecute(r)
}

/*
IndexPriceKlineCandlestickData Index Price Kline/Candlestick Data
Get /dapi/v1/indexPriceKlines

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Index-Price-Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param interval -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiIndexPriceKlineCandlestickDataRequest
*/
func (a *MarketDataAPIService) IndexPriceKlineCandlestickData(ctx context.Context) ApiIndexPriceKlineCandlestickDataRequest {
	return ApiIndexPriceKlineCandlestickDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IndexPriceKlineCandlestickDataResponse
func (a *MarketDataAPIService) IndexPriceKlineCandlestickDataExecute(r ApiIndexPriceKlineCandlestickDataRequest) (*common.RestApiResponse[models.IndexPriceKlineCandlestickDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/indexPriceKlines"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
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

	resp, err := SendRequest[models.IndexPriceKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiKlineCandlestickDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	interval   *models.ContinuousContractKlineCandlestickDataIntervalParameter
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiKlineCandlestickDataRequest) Symbol(symbol string) ApiKlineCandlestickDataRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlineCandlestickDataRequest) Interval(interval models.ContinuousContractKlineCandlestickDataIntervalParameter) ApiKlineCandlestickDataRequest {
	r.interval = &interval
	return r
}

func (r ApiKlineCandlestickDataRequest) StartTime(startTime int64) ApiKlineCandlestickDataRequest {
	r.startTime = &startTime
	return r
}

func (r ApiKlineCandlestickDataRequest) EndTime(endTime int64) ApiKlineCandlestickDataRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiKlineCandlestickDataRequest) Limit(limit int64) ApiKlineCandlestickDataRequest {
	r.limit = &limit
	return r
}

func (r ApiKlineCandlestickDataRequest) Execute() (*common.RestApiResponse[models.KlineCandlestickDataResponse], error) {
	return r.ApiService.KlineCandlestickDataExecute(r)
}

/*
KlineCandlestickData Kline/Candlestick Data
Get /dapi/v1/klines

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param interval -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/klines"

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

type ApiLongShortRatioRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	pair       *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

// BTCUSD
func (r ApiLongShortRatioRequest) Pair(pair string) ApiLongShortRatioRequest {
	r.pair = &pair
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiLongShortRatioRequest) Period(period models.BasisPeriodParameter) ApiLongShortRatioRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiLongShortRatioRequest) Limit(limit int64) ApiLongShortRatioRequest {
	r.limit = &limit
	return r
}

func (r ApiLongShortRatioRequest) StartTime(startTime int64) ApiLongShortRatioRequest {
	r.startTime = &startTime
	return r
}

func (r ApiLongShortRatioRequest) EndTime(endTime int64) ApiLongShortRatioRequest {
	r.endTime = &endTime
	return r
}

func (r ApiLongShortRatioRequest) Execute() (*common.RestApiResponse[models.LongShortRatioResponse], error) {
	return r.ApiService.LongShortRatioExecute(r)
}

/*
LongShortRatio Long/Short Ratio
Get /futures/data/globalLongShortAccountRatio

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Long-Short-Ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiLongShortRatioRequest
*/
func (a *MarketDataAPIService) LongShortRatio(ctx context.Context) ApiLongShortRatioRequest {
	return ApiLongShortRatioRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LongShortRatioResponse
func (a *MarketDataAPIService) LongShortRatioExecute(r ApiLongShortRatioRequest) (*common.RestApiResponse[models.LongShortRatioResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/globalLongShortAccountRatio"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.LongShortRatioResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarkPriceKlineCandlestickDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	interval   *models.ContinuousContractKlineCandlestickDataIntervalParameter
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiMarkPriceKlineCandlestickDataRequest) Symbol(symbol string) ApiMarkPriceKlineCandlestickDataRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarkPriceKlineCandlestickDataRequest) Interval(interval models.ContinuousContractKlineCandlestickDataIntervalParameter) ApiMarkPriceKlineCandlestickDataRequest {
	r.interval = &interval
	return r
}

func (r ApiMarkPriceKlineCandlestickDataRequest) StartTime(startTime int64) ApiMarkPriceKlineCandlestickDataRequest {
	r.startTime = &startTime
	return r
}

func (r ApiMarkPriceKlineCandlestickDataRequest) EndTime(endTime int64) ApiMarkPriceKlineCandlestickDataRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiMarkPriceKlineCandlestickDataRequest) Limit(limit int64) ApiMarkPriceKlineCandlestickDataRequest {
	r.limit = &limit
	return r
}

func (r ApiMarkPriceKlineCandlestickDataRequest) Execute() (*common.RestApiResponse[models.MarkPriceKlineCandlestickDataResponse], error) {
	return r.ApiService.MarkPriceKlineCandlestickDataExecute(r)
}

/*
MarkPriceKlineCandlestickData Mark Price Kline/Candlestick Data
Get /dapi/v1/markPriceKlines

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Mark-Price-Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param interval -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiMarkPriceKlineCandlestickDataRequest
*/
func (a *MarketDataAPIService) MarkPriceKlineCandlestickData(ctx context.Context) ApiMarkPriceKlineCandlestickDataRequest {
	return ApiMarkPriceKlineCandlestickDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarkPriceKlineCandlestickDataResponse
func (a *MarketDataAPIService) MarkPriceKlineCandlestickDataExecute(r ApiMarkPriceKlineCandlestickDataRequest) (*common.RestApiResponse[models.MarkPriceKlineCandlestickDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/markPriceKlines"

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

	resp, err := SendRequest[models.MarkPriceKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOldTradesLookupRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	limit      *int64
	fromId     *int64
}

func (r ApiOldTradesLookupRequest) Symbol(symbol string) ApiOldTradesLookupRequest {
	r.symbol = &symbol
	return r
}

// Default 100; max 1000
func (r ApiOldTradesLookupRequest) Limit(limit int64) ApiOldTradesLookupRequest {
	r.limit = &limit
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiOldTradesLookupRequest) FromId(fromId int64) ApiOldTradesLookupRequest {
	r.fromId = &fromId
	return r
}

func (r ApiOldTradesLookupRequest) Execute() (*common.RestApiResponse[models.OldTradesLookupResponse], error) {
	return r.ApiService.OldTradesLookupExecute(r)
}

/*
OldTradesLookup Old Trades Lookup(MARKET_DATA)
Get /dapi/v1/historicalTrades

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Old-Trades-Lookup

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default 100; max 1000
@param fromId -  ID to get aggregate trades from INCLUSIVE.
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/historicalTrades"

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

	resp, err := SendRequest[models.OldTradesLookupResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOpenInterestRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiOpenInterestRequest) Symbol(symbol string) ApiOpenInterestRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOpenInterestRequest) Execute() (*common.RestApiResponse[models.OpenInterestResponse], error) {
	return r.ApiService.OpenInterestExecute(r)
}

/*
OpenInterest Open Interest
Get /dapi/v1/openInterest

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Open-Interest

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/openInterest"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.OpenInterestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOpenInterestStatisticsRequest struct {
	ctx          context.Context
	ApiService   *MarketDataAPIService
	pair         *string
	contractType *models.BasisContractTypeParameter
	period       *models.BasisPeriodParameter
	limit        *int64
	startTime    *int64
	endTime      *int64
}

// BTCUSD
func (r ApiOpenInterestStatisticsRequest) Pair(pair string) ApiOpenInterestStatisticsRequest {
	r.pair = &pair
	return r
}

// ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (r ApiOpenInterestStatisticsRequest) ContractType(contractType models.BasisContractTypeParameter) ApiOpenInterestStatisticsRequest {
	r.contractType = &contractType
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiOpenInterestStatisticsRequest) Period(period models.BasisPeriodParameter) ApiOpenInterestStatisticsRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiOpenInterestStatisticsRequest) Limit(limit int64) ApiOpenInterestStatisticsRequest {
	r.limit = &limit
	return r
}

func (r ApiOpenInterestStatisticsRequest) StartTime(startTime int64) ApiOpenInterestStatisticsRequest {
	r.startTime = &startTime
	return r
}

func (r ApiOpenInterestStatisticsRequest) EndTime(endTime int64) ApiOpenInterestStatisticsRequest {
	r.endTime = &endTime
	return r
}

func (r ApiOpenInterestStatisticsRequest) Execute() (*common.RestApiResponse[models.OpenInterestStatisticsResponse], error) {
	return r.ApiService.OpenInterestStatisticsExecute(r)
}

/*
OpenInterestStatistics Open Interest Statistics
Get /futures/data/openInterestHist

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Open-Interest-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param contractType -  ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiOpenInterestStatisticsRequest
*/
func (a *MarketDataAPIService) OpenInterestStatistics(ctx context.Context) ApiOpenInterestStatisticsRequest {
	return ApiOpenInterestStatisticsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OpenInterestStatisticsResponse
func (a *MarketDataAPIService) OpenInterestStatisticsExecute(r ApiOpenInterestStatisticsRequest) (*common.RestApiResponse[models.OpenInterestStatisticsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/openInterestHist"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.contractType == nil {
		return nil, common.ReportError("contractType is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "contractType", r.contractType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.OpenInterestStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

func (r ApiOrderBookRequest) Symbol(symbol string) ApiOrderBookRequest {
	r.symbol = &symbol
	return r
}

// Default 100; max 1000
func (r ApiOrderBookRequest) Limit(limit int64) ApiOrderBookRequest {
	r.limit = &limit
	return r
}

func (r ApiOrderBookRequest) Execute() (*common.RestApiResponse[models.OrderBookResponse], error) {
	return r.ApiService.OrderBookExecute(r)
}

/*
OrderBook Order Book
Get /dapi/v1/depth

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Order-Book

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/depth"

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

type ApiPremiumIndexKlineDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	interval   *models.ContinuousContractKlineCandlestickDataIntervalParameter
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiPremiumIndexKlineDataRequest) Symbol(symbol string) ApiPremiumIndexKlineDataRequest {
	r.symbol = &symbol
	return r
}

func (r ApiPremiumIndexKlineDataRequest) Interval(interval models.ContinuousContractKlineCandlestickDataIntervalParameter) ApiPremiumIndexKlineDataRequest {
	r.interval = &interval
	return r
}

func (r ApiPremiumIndexKlineDataRequest) StartTime(startTime int64) ApiPremiumIndexKlineDataRequest {
	r.startTime = &startTime
	return r
}

func (r ApiPremiumIndexKlineDataRequest) EndTime(endTime int64) ApiPremiumIndexKlineDataRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiPremiumIndexKlineDataRequest) Limit(limit int64) ApiPremiumIndexKlineDataRequest {
	r.limit = &limit
	return r
}

func (r ApiPremiumIndexKlineDataRequest) Execute() (*common.RestApiResponse[models.PremiumIndexKlineDataResponse], error) {
	return r.ApiService.PremiumIndexKlineDataExecute(r)
}

/*
PremiumIndexKlineData Premium index Kline Data
Get /dapi/v1/premiumIndexKlines

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Premium-index-Kline-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param interval -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiPremiumIndexKlineDataRequest
*/
func (a *MarketDataAPIService) PremiumIndexKlineData(ctx context.Context) ApiPremiumIndexKlineDataRequest {
	return ApiPremiumIndexKlineDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PremiumIndexKlineDataResponse
func (a *MarketDataAPIService) PremiumIndexKlineDataExecute(r ApiPremiumIndexKlineDataRequest) (*common.RestApiResponse[models.PremiumIndexKlineDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/premiumIndexKlines"

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

	resp, err := SendRequest[models.PremiumIndexKlineDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryIndexPriceConstituentsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiQueryIndexPriceConstituentsRequest) Symbol(symbol string) ApiQueryIndexPriceConstituentsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryIndexPriceConstituentsRequest) Execute() (*common.RestApiResponse[models.QueryIndexPriceConstituentsResponse], error) {
	return r.ApiService.QueryIndexPriceConstituentsExecute(r)
}

/*
QueryIndexPriceConstituents Query Index Price Constituents
Get /dapi/v1/constituents

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Index-Constituents

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiQueryIndexPriceConstituentsRequest
*/
func (a *MarketDataAPIService) QueryIndexPriceConstituents(ctx context.Context) ApiQueryIndexPriceConstituentsRequest {
	return ApiQueryIndexPriceConstituentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryIndexPriceConstituentsResponse
func (a *MarketDataAPIService) QueryIndexPriceConstituentsExecute(r ApiQueryIndexPriceConstituentsRequest) (*common.RestApiResponse[models.QueryIndexPriceConstituentsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/constituents"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.QueryIndexPriceConstituentsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

func (r ApiRecentTradesListRequest) Symbol(symbol string) ApiRecentTradesListRequest {
	r.symbol = &symbol
	return r
}

// Default 100; max 1000
func (r ApiRecentTradesListRequest) Limit(limit int64) ApiRecentTradesListRequest {
	r.limit = &limit
	return r
}

func (r ApiRecentTradesListRequest) Execute() (*common.RestApiResponse[models.RecentTradesListResponse], error) {
	return r.ApiService.RecentTradesListExecute(r)
}

/*
RecentTradesList Recent Trades List
Get /dapi/v1/trades

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Recent-Trades-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default 100; max 1000
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/trades"

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

type ApiSymbolOrderBookTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	pair       *string
}

func (r ApiSymbolOrderBookTickerRequest) Symbol(symbol string) ApiSymbolOrderBookTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolOrderBookTickerRequest) Pair(pair string) ApiSymbolOrderBookTickerRequest {
	r.pair = &pair
	return r
}

func (r ApiSymbolOrderBookTickerRequest) Execute() (*common.RestApiResponse[models.SymbolOrderBookTickerResponse], error) {
	return r.ApiService.SymbolOrderBookTickerExecute(r)
}

/*
SymbolOrderBookTicker Symbol Order Book Ticker
Get /dapi/v1/ticker/bookTicker

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Symbol-Order-Book-Ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@return ApiSymbolOrderBookTickerRequest
*/
func (a *MarketDataAPIService) SymbolOrderBookTicker(ctx context.Context) ApiSymbolOrderBookTickerRequest {
	return ApiSymbolOrderBookTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SymbolOrderBookTickerResponse
func (a *MarketDataAPIService) SymbolOrderBookTickerExecute(r ApiSymbolOrderBookTickerRequest) (*common.RestApiResponse[models.SymbolOrderBookTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/ticker/bookTicker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}

	resp, err := SendRequest[models.SymbolOrderBookTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSymbolPriceTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	pair       *string
}

func (r ApiSymbolPriceTickerRequest) Symbol(symbol string) ApiSymbolPriceTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolPriceTickerRequest) Pair(pair string) ApiSymbolPriceTickerRequest {
	r.pair = &pair
	return r
}

func (r ApiSymbolPriceTickerRequest) Execute() (*common.RestApiResponse[models.SymbolPriceTickerResponse], error) {
	return r.ApiService.SymbolPriceTickerExecute(r)
}

/*
SymbolPriceTicker Symbol Price Ticker
Get /dapi/v1/ticker/price

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Symbol-Price-Ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
@return ApiSymbolPriceTickerRequest
*/
func (a *MarketDataAPIService) SymbolPriceTicker(ctx context.Context) ApiSymbolPriceTickerRequest {
	return ApiSymbolPriceTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SymbolPriceTickerResponse
func (a *MarketDataAPIService) SymbolPriceTickerExecute(r ApiSymbolPriceTickerRequest) (*common.RestApiResponse[models.SymbolPriceTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/ticker/price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}

	resp, err := SendRequest[models.SymbolPriceTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTakerBuySellVolumeRequest struct {
	ctx          context.Context
	ApiService   *MarketDataAPIService
	pair         *string
	contractType *models.BasisContractTypeParameter
	period       *models.BasisPeriodParameter
	limit        *int64
	startTime    *int64
	endTime      *int64
}

// BTCUSD
func (r ApiTakerBuySellVolumeRequest) Pair(pair string) ApiTakerBuySellVolumeRequest {
	r.pair = &pair
	return r
}

// ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (r ApiTakerBuySellVolumeRequest) ContractType(contractType models.BasisContractTypeParameter) ApiTakerBuySellVolumeRequest {
	r.contractType = &contractType
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiTakerBuySellVolumeRequest) Period(period models.BasisPeriodParameter) ApiTakerBuySellVolumeRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiTakerBuySellVolumeRequest) Limit(limit int64) ApiTakerBuySellVolumeRequest {
	r.limit = &limit
	return r
}

func (r ApiTakerBuySellVolumeRequest) StartTime(startTime int64) ApiTakerBuySellVolumeRequest {
	r.startTime = &startTime
	return r
}

func (r ApiTakerBuySellVolumeRequest) EndTime(endTime int64) ApiTakerBuySellVolumeRequest {
	r.endTime = &endTime
	return r
}

func (r ApiTakerBuySellVolumeRequest) Execute() (*common.RestApiResponse[models.TakerBuySellVolumeResponse], error) {
	return r.ApiService.TakerBuySellVolumeExecute(r)
}

/*
TakerBuySellVolume Taker Buy/Sell Volume
Get /futures/data/takerBuySellVol

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Taker-Buy-Sell-Volume

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param contractType -  ALL, CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiTakerBuySellVolumeRequest
*/
func (a *MarketDataAPIService) TakerBuySellVolume(ctx context.Context) ApiTakerBuySellVolumeRequest {
	return ApiTakerBuySellVolumeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TakerBuySellVolumeResponse
func (a *MarketDataAPIService) TakerBuySellVolumeExecute(r ApiTakerBuySellVolumeRequest) (*common.RestApiResponse[models.TakerBuySellVolumeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/takerBuySellVol"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.contractType == nil {
		return nil, common.ReportError("contractType is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "contractType", r.contractType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.TakerBuySellVolumeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
Get /dapi/v1/ping

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Test-Connectivity

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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/ping"

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
	pair       *string
}

func (r ApiTicker24hrPriceChangeStatisticsRequest) Symbol(symbol string) ApiTicker24hrPriceChangeStatisticsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTicker24hrPriceChangeStatisticsRequest) Pair(pair string) ApiTicker24hrPriceChangeStatisticsRequest {
	r.pair = &pair
	return r
}

func (r ApiTicker24hrPriceChangeStatisticsRequest) Execute() (*common.RestApiResponse[models.Ticker24hrPriceChangeStatisticsResponse], error) {
	return r.ApiService.Ticker24hrPriceChangeStatisticsExecute(r)
}

/*
Ticker24hrPriceChangeStatistics 24hr Ticker Price Change Statistics
Get /dapi/v1/ticker/24hr

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/24hr-Ticker-Price-Change-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param pair -
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
	localVarPath := a.client.cfg.BasePath + "/dapi/v1/ticker/24hr"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}

	resp, err := SendRequest[models.Ticker24hrPriceChangeStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTopTraderLongShortRatioAccountsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

func (r ApiTopTraderLongShortRatioAccountsRequest) Symbol(symbol string) ApiTopTraderLongShortRatioAccountsRequest {
	r.symbol = &symbol
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiTopTraderLongShortRatioAccountsRequest) Period(period models.BasisPeriodParameter) ApiTopTraderLongShortRatioAccountsRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiTopTraderLongShortRatioAccountsRequest) Limit(limit int64) ApiTopTraderLongShortRatioAccountsRequest {
	r.limit = &limit
	return r
}

func (r ApiTopTraderLongShortRatioAccountsRequest) StartTime(startTime int64) ApiTopTraderLongShortRatioAccountsRequest {
	r.startTime = &startTime
	return r
}

func (r ApiTopTraderLongShortRatioAccountsRequest) EndTime(endTime int64) ApiTopTraderLongShortRatioAccountsRequest {
	r.endTime = &endTime
	return r
}

func (r ApiTopTraderLongShortRatioAccountsRequest) Execute() (*common.RestApiResponse[models.TopTraderLongShortRatioAccountsResponse], error) {
	return r.ApiService.TopTraderLongShortRatioAccountsExecute(r)
}

/*
TopTraderLongShortRatioAccounts Top Trader Long/Short Ratio (Accounts)
Get /futures/data/topLongShortAccountRatio

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Top-Long-Short-Account-Ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiTopTraderLongShortRatioAccountsRequest
*/
func (a *MarketDataAPIService) TopTraderLongShortRatioAccounts(ctx context.Context) ApiTopTraderLongShortRatioAccountsRequest {
	return ApiTopTraderLongShortRatioAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TopTraderLongShortRatioAccountsResponse
func (a *MarketDataAPIService) TopTraderLongShortRatioAccountsExecute(r ApiTopTraderLongShortRatioAccountsRequest) (*common.RestApiResponse[models.TopTraderLongShortRatioAccountsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/topLongShortAccountRatio"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.TopTraderLongShortRatioAccountsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTopTraderLongShortRatioPositionsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	pair       *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

// BTCUSD
func (r ApiTopTraderLongShortRatioPositionsRequest) Pair(pair string) ApiTopTraderLongShortRatioPositionsRequest {
	r.pair = &pair
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiTopTraderLongShortRatioPositionsRequest) Period(period models.BasisPeriodParameter) ApiTopTraderLongShortRatioPositionsRequest {
	r.period = &period
	return r
}

// Default 100; max 1000
func (r ApiTopTraderLongShortRatioPositionsRequest) Limit(limit int64) ApiTopTraderLongShortRatioPositionsRequest {
	r.limit = &limit
	return r
}

func (r ApiTopTraderLongShortRatioPositionsRequest) StartTime(startTime int64) ApiTopTraderLongShortRatioPositionsRequest {
	r.startTime = &startTime
	return r
}

func (r ApiTopTraderLongShortRatioPositionsRequest) EndTime(endTime int64) ApiTopTraderLongShortRatioPositionsRequest {
	r.endTime = &endTime
	return r
}

func (r ApiTopTraderLongShortRatioPositionsRequest) Execute() (*common.RestApiResponse[models.TopTraderLongShortRatioPositionsResponse], error) {
	return r.ApiService.TopTraderLongShortRatioPositionsExecute(r)
}

/*
TopTraderLongShortRatioPositions Top Trader Long/Short Ratio (Positions)
Get /futures/data/topLongShortPositionRatio

https://developers.binance.com/docs/derivatives/coin-margined-futures/market-data/rest-api/Top-Trader-Long-Short-Ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -  BTCUSD
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 100; max 1000
@param startTime -
@param endTime -
@return ApiTopTraderLongShortRatioPositionsRequest
*/
func (a *MarketDataAPIService) TopTraderLongShortRatioPositions(ctx context.Context) ApiTopTraderLongShortRatioPositionsRequest {
	return ApiTopTraderLongShortRatioPositionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TopTraderLongShortRatioPositionsResponse
func (a *MarketDataAPIService) TopTraderLongShortRatioPositionsExecute(r ApiTopTraderLongShortRatioPositionsRequest) (*common.RestApiResponse[models.TopTraderLongShortRatioPositionsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/topLongShortPositionRatio"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.period == nil {
		return nil, common.ReportError("period is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.TopTraderLongShortRatioPositionsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
