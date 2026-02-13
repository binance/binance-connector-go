/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiAdlRiskRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiAdlRiskRequest) Symbol(symbol string) ApiAdlRiskRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAdlRiskRequest) Execute() (*common.RestApiResponse[models.AdlRiskResponse], error) {
	return r.ApiService.AdlRiskExecute(r)
}

/*
AdlRisk ADL Risk
Get /fapi/v1/symbolAdlRisk

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/ADL-Risk

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiAdlRiskRequest
*/
func (a *MarketDataAPIService) AdlRisk(ctx context.Context) ApiAdlRiskRequest {
	return ApiAdlRiskRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdlRiskResponse
func (a *MarketDataAPIService) AdlRiskExecute(r ApiAdlRiskRequest) (*common.RestApiResponse[models.AdlRiskResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/symbolAdlRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.AdlRiskResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

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

func (r ApiBasisRequest) Pair(pair string) ApiBasisRequest {
	r.pair = &pair
	return r
}

func (r ApiBasisRequest) ContractType(contractType models.BasisContractTypeParameter) ApiBasisRequest {
	r.contractType = &contractType
	return r
}

// \&quot;5m\&quot;,\&quot;15m\&quot;,\&quot;30m\&quot;,\&quot;1h\&quot;,\&quot;2h\&quot;,\&quot;4h\&quot;,\&quot;6h\&quot;,\&quot;12h\&quot;,\&quot;1d\&quot;
func (r ApiBasisRequest) Period(period models.BasisPeriodParameter) ApiBasisRequest {
	r.period = &period
	return r
}

// Default 30,Max 500
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

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Basis

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -
@param contractType -
@param period -  \"5m\",\"15m\",\"30m\",\"1h\",\"2h\",\"4h\",\"6h\",\"12h\",\"1d\"
@param limit -  Default 30,Max 500
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
	if r.limit == nil {
		return nil, common.ReportError("limit is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "contractType", r.contractType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "period", r.period, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}

	resp, err := SendRequest[models.BasisResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
CheckServerTime Check Server Time
Get /fapi/v1/time

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Check-Server-Time

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/time"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.CheckServerTimeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCompositeIndexSymbolInformationRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiCompositeIndexSymbolInformationRequest) Symbol(symbol string) ApiCompositeIndexSymbolInformationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCompositeIndexSymbolInformationRequest) Execute() (*common.RestApiResponse[models.CompositeIndexSymbolInformationResponse], error) {
	return r.ApiService.CompositeIndexSymbolInformationExecute(r)
}

/*
CompositeIndexSymbolInformation Composite Index Symbol Information
Get /fapi/v1/indexInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Composite-Index-Symbol-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiCompositeIndexSymbolInformationRequest
*/
func (a *MarketDataAPIService) CompositeIndexSymbolInformation(ctx context.Context) ApiCompositeIndexSymbolInformationRequest {
	return ApiCompositeIndexSymbolInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CompositeIndexSymbolInformationResponse
func (a *MarketDataAPIService) CompositeIndexSymbolInformationExecute(r ApiCompositeIndexSymbolInformationRequest) (*common.RestApiResponse[models.CompositeIndexSymbolInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/indexInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.CompositeIndexSymbolInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/aggTrades

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Compressed-Aggregate-Trades-List

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/aggTrades"

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

	resp, err := SendRequest[models.CompressedAggregateTradesListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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

func (r ApiContinuousContractKlineCandlestickDataRequest) Pair(pair string) ApiContinuousContractKlineCandlestickDataRequest {
	r.pair = &pair
	return r
}

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
Get /fapi/v1/continuousKlines

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Continuous-Contract-Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -
@param contractType -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/continuousKlines"

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

	resp, err := SendRequest[models.ContinuousContractKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/exchangeInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Exchange-Information

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.ExchangeInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFundingRateHistoryRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

func (r ApiGetFundingRateHistoryRequest) Symbol(symbol string) ApiGetFundingRateHistoryRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetFundingRateHistoryRequest) StartTime(startTime int64) ApiGetFundingRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFundingRateHistoryRequest) EndTime(endTime int64) ApiGetFundingRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default 100; max 1000
func (r ApiGetFundingRateHistoryRequest) Limit(limit int64) ApiGetFundingRateHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFundingRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetFundingRateHistoryResponse], error) {
	return r.ApiService.GetFundingRateHistoryExecute(r)
}

/*
GetFundingRateHistory Get Funding Rate History
Get /fapi/v1/fundingRate

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Get-Funding-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param startTime -
@param endTime -
@param limit -  Default 100; max 1000
@return ApiGetFundingRateHistoryRequest
*/
func (a *MarketDataAPIService) GetFundingRateHistory(ctx context.Context) ApiGetFundingRateHistoryRequest {
	return ApiGetFundingRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFundingRateHistoryResponse
func (a *MarketDataAPIService) GetFundingRateHistoryExecute(r ApiGetFundingRateHistoryRequest) (*common.RestApiResponse[models.GetFundingRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/fundingRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
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

	resp, err := SendRequest[models.GetFundingRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/fundingInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Get-Funding-Rate-Info

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/fundingInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetFundingRateInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/indexPriceKlines

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Index-Price-Kline-Candlestick-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/indexPriceKlines"

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

	resp, err := SendRequest[models.IndexPriceKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/klines

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Kline-Candlestick-Data

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/klines"

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

	resp, err := SendRequest[models.KlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiLongShortRatioRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

func (r ApiLongShortRatioRequest) Symbol(symbol string) ApiLongShortRatioRequest {
	r.symbol = &symbol
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

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Long-Short-Ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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

	resp, err := SendRequest[models.LongShortRatioResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarkPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiMarkPriceRequest) Symbol(symbol string) ApiMarkPriceRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMarkPriceRequest) Execute() (*common.RestApiResponse[models.MarkPriceResponse], error) {
	return r.ApiService.MarkPriceExecute(r)
}

/*
MarkPrice Mark Price
Get /fapi/v1/premiumIndex

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Mark-Price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiMarkPriceRequest
*/
func (a *MarketDataAPIService) MarkPrice(ctx context.Context) ApiMarkPriceRequest {
	return ApiMarkPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarkPriceResponse
func (a *MarketDataAPIService) MarkPriceExecute(r ApiMarkPriceRequest) (*common.RestApiResponse[models.MarkPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/premiumIndex"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.MarkPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/markPriceKlines

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Mark-Price-Kline-Candlestick-Data

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/markPriceKlines"

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

	resp, err := SendRequest[models.MarkPriceKlineCandlestickDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMultiAssetsModeAssetIndexRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiMultiAssetsModeAssetIndexRequest) Symbol(symbol string) ApiMultiAssetsModeAssetIndexRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMultiAssetsModeAssetIndexRequest) Execute() (*common.RestApiResponse[models.MultiAssetsModeAssetIndexResponse], error) {
	return r.ApiService.MultiAssetsModeAssetIndexExecute(r)
}

/*
MultiAssetsModeAssetIndex Multi-Assets Mode Asset Index
Get /fapi/v1/assetIndex

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Multi-Assets-Mode-Asset-Index

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiMultiAssetsModeAssetIndexRequest
*/
func (a *MarketDataAPIService) MultiAssetsModeAssetIndex(ctx context.Context) ApiMultiAssetsModeAssetIndexRequest {
	return ApiMultiAssetsModeAssetIndexRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MultiAssetsModeAssetIndexResponse
func (a *MarketDataAPIService) MultiAssetsModeAssetIndexExecute(r ApiMultiAssetsModeAssetIndexRequest) (*common.RestApiResponse[models.MultiAssetsModeAssetIndexResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/assetIndex"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.MultiAssetsModeAssetIndexResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
OldTradesLookup Old Trades Lookup (MARKET_DATA)
Get /fapi/v1/historicalTrades

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Old-Trades-Lookup

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/historicalTrades"

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

	resp, err := SendRequest[models.OldTradesLookupResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/openInterest

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Open-Interest

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/openInterest"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.OpenInterestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOpenInterestStatisticsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

func (r ApiOpenInterestStatisticsRequest) Symbol(symbol string) ApiOpenInterestStatisticsRequest {
	r.symbol = &symbol
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

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Open-Interest-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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

	resp, err := SendRequest[models.OpenInterestStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/depth

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Order-Book

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/depth"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.OrderBookResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/premiumIndexKlines

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Premium-index-Kline-Data

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/premiumIndexKlines"

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

	resp, err := SendRequest[models.PremiumIndexKlineDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQuarterlyContractSettlementPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	pair       *string
}

func (r ApiQuarterlyContractSettlementPriceRequest) Pair(pair string) ApiQuarterlyContractSettlementPriceRequest {
	r.pair = &pair
	return r
}

func (r ApiQuarterlyContractSettlementPriceRequest) Execute() (*common.RestApiResponse[models.QuarterlyContractSettlementPriceResponse], error) {
	return r.ApiService.QuarterlyContractSettlementPriceExecute(r)
}

/*
QuarterlyContractSettlementPrice Quarterly Contract Settlement Price
Get /futures/data/delivery-price

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Delivery-Price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pair -
@return ApiQuarterlyContractSettlementPriceRequest
*/
func (a *MarketDataAPIService) QuarterlyContractSettlementPrice(ctx context.Context) ApiQuarterlyContractSettlementPriceRequest {
	return ApiQuarterlyContractSettlementPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuarterlyContractSettlementPriceResponse
func (a *MarketDataAPIService) QuarterlyContractSettlementPriceExecute(r ApiQuarterlyContractSettlementPriceRequest) (*common.RestApiResponse[models.QuarterlyContractSettlementPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/futures/data/delivery-price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")

	resp, err := SendRequest[models.QuarterlyContractSettlementPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/constituents

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Index-Constituents

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/constituents"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.QueryIndexPriceConstituentsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryInsuranceFundBalanceSnapshotRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiQueryInsuranceFundBalanceSnapshotRequest) Symbol(symbol string) ApiQueryInsuranceFundBalanceSnapshotRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryInsuranceFundBalanceSnapshotRequest) Execute() (*common.RestApiResponse[models.QueryInsuranceFundBalanceSnapshotResponse], error) {
	return r.ApiService.QueryInsuranceFundBalanceSnapshotExecute(r)
}

/*
QueryInsuranceFundBalanceSnapshot Query Insurance Fund Balance Snapshot
Get /fapi/v1/insuranceBalance

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Query-Insurance-Fund-Balance-Snapshot

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiQueryInsuranceFundBalanceSnapshotRequest
*/
func (a *MarketDataAPIService) QueryInsuranceFundBalanceSnapshot(ctx context.Context) ApiQueryInsuranceFundBalanceSnapshotRequest {
	return ApiQueryInsuranceFundBalanceSnapshotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryInsuranceFundBalanceSnapshotResponse
func (a *MarketDataAPIService) QueryInsuranceFundBalanceSnapshotExecute(r ApiQueryInsuranceFundBalanceSnapshotRequest) (*common.RestApiResponse[models.QueryInsuranceFundBalanceSnapshotResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/insuranceBalance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.QueryInsuranceFundBalanceSnapshotResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/trades

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Recent-Trades-List

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/trades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.RecentTradesListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRpiOrderBookRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	limit      *int64
}

func (r ApiRpiOrderBookRequest) Symbol(symbol string) ApiRpiOrderBookRequest {
	r.symbol = &symbol
	return r
}

// Default 100; max 1000
func (r ApiRpiOrderBookRequest) Limit(limit int64) ApiRpiOrderBookRequest {
	r.limit = &limit
	return r
}

func (r ApiRpiOrderBookRequest) Execute() (*common.RestApiResponse[models.RpiOrderBookResponse], error) {
	return r.ApiService.RpiOrderBookExecute(r)
}

/*
RpiOrderBook RPI Order Book
Get /fapi/v1/rpiDepth

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Order-Book-RPI

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param limit -  Default 100; max 1000
@return ApiRpiOrderBookRequest
*/
func (a *MarketDataAPIService) RpiOrderBook(ctx context.Context) ApiRpiOrderBookRequest {
	return ApiRpiOrderBookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RpiOrderBookResponse
func (a *MarketDataAPIService) RpiOrderBookExecute(r ApiRpiOrderBookRequest) (*common.RestApiResponse[models.RpiOrderBookResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/rpiDepth"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.RpiOrderBookResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSymbolOrderBookTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiSymbolOrderBookTickerRequest) Symbol(symbol string) ApiSymbolOrderBookTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolOrderBookTickerRequest) Execute() (*common.RestApiResponse[models.SymbolOrderBookTickerResponse], error) {
	return r.ApiService.SymbolOrderBookTickerExecute(r)
}

/*
SymbolOrderBookTicker Symbol Order Book Ticker
Get /fapi/v1/ticker/bookTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Symbol-Order-Book-Ticker

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/ticker/bookTicker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.SymbolOrderBookTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSymbolPriceTickerRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiSymbolPriceTickerRequest) Symbol(symbol string) ApiSymbolPriceTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolPriceTickerRequest) Execute() (*common.RestApiResponse[models.SymbolPriceTickerResponse], error) {
	return r.ApiService.SymbolPriceTickerExecute(r)
}

/*
	SymbolPriceTicker Symbol Price Ticker
	Get /fapi/v1/ticker/price

	https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Symbol-Price-Ticker

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param symbol -
	@return ApiSymbolPriceTickerRequest

Deprecated
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
//
// Deprecated
func (a *MarketDataAPIService) SymbolPriceTickerExecute(r ApiSymbolPriceTickerRequest) (*common.RestApiResponse[models.SymbolPriceTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/ticker/price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.SymbolPriceTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSymbolPriceTickerV2Request struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiSymbolPriceTickerV2Request) Symbol(symbol string) ApiSymbolPriceTickerV2Request {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolPriceTickerV2Request) Execute() (*common.RestApiResponse[models.SymbolPriceTickerV2Response], error) {
	return r.ApiService.SymbolPriceTickerV2Execute(r)
}

/*
SymbolPriceTickerV2 Symbol Price Ticker V2
Get /fapi/v2/ticker/price

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Symbol-Price-Ticker-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiSymbolPriceTickerV2Request
*/
func (a *MarketDataAPIService) SymbolPriceTickerV2(ctx context.Context) ApiSymbolPriceTickerV2Request {
	return ApiSymbolPriceTickerV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SymbolPriceTickerV2Response
func (a *MarketDataAPIService) SymbolPriceTickerV2Execute(r ApiSymbolPriceTickerV2Request) (*common.RestApiResponse[models.SymbolPriceTickerV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v2/ticker/price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.SymbolPriceTickerV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTakerBuySellVolumeRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

func (r ApiTakerBuySellVolumeRequest) Symbol(symbol string) ApiTakerBuySellVolumeRequest {
	r.symbol = &symbol
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
Get /futures/data/takerlongshortRatio

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Taker-BuySell-Volume

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/futures/data/takerlongshortRatio"

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

	resp, err := SendRequest[models.TakerBuySellVolumeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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
Get /fapi/v1/ping

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Test-Connectivity

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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/ping"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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

func (r ApiTicker24hrPriceChangeStatisticsRequest) Symbol(symbol string) ApiTicker24hrPriceChangeStatisticsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTicker24hrPriceChangeStatisticsRequest) Execute() (*common.RestApiResponse[models.Ticker24hrPriceChangeStatisticsResponse], error) {
	return r.ApiService.Ticker24hrPriceChangeStatisticsExecute(r)
}

/*
Ticker24hrPriceChangeStatistics 24hr Ticker Price Change Statistics
Get /fapi/v1/ticker/24hr

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/24hr-Ticker-Price-Change-Statistics

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/ticker/24hr"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.Ticker24hrPriceChangeStatisticsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
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

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Top-Long-Short-Account-Ratio

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

	resp, err := SendRequest[models.TopTraderLongShortRatioAccountsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTopTraderLongShortRatioPositionsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	period     *models.BasisPeriodParameter
	limit      *int64
	startTime  *int64
	endTime    *int64
}

func (r ApiTopTraderLongShortRatioPositionsRequest) Symbol(symbol string) ApiTopTraderLongShortRatioPositionsRequest {
	r.symbol = &symbol
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

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Top-Trader-Long-Short-Ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
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

	resp, err := SendRequest[models.TopTraderLongShortRatioPositionsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTradingScheduleRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiTradingScheduleRequest) Execute() (*common.RestApiResponse[models.TradingScheduleResponse], error) {
	return r.ApiService.TradingScheduleExecute(r)
}

/*
TradingSchedule Trading Schedule
Get /fapi/v1/tradingSchedule

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/rest-api/Trading-Schedule

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiTradingScheduleRequest
*/
func (a *MarketDataAPIService) TradingSchedule(ctx context.Context) ApiTradingScheduleRequest {
	return ApiTradingScheduleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TradingScheduleResponse
func (a *MarketDataAPIService) TradingScheduleExecute(r ApiTradingScheduleRequest) (*common.RestApiResponse[models.TradingScheduleResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/tradingSchedule"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.TradingScheduleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
