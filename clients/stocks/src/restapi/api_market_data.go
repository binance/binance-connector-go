/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package binancestocksrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiExchangeInfoRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// Filter to a single US-equity ticker, e.g. &#x60;AAPL&#x60;. When omitted, returns all active symbols. An unknown ticker returns an empty &#x60;symbols&#x60; array (HTTP 200), not an error.
func (r ApiExchangeInfoRequest) Symbol(symbol string) ApiExchangeInfoRequest {
	r.symbol = &symbol
	return r
}

func (r ApiExchangeInfoRequest) Execute() (*common.RestApiResponse[models.ExchangeInfoResponse], error) {
	return r.ApiService.ExchangeInfoExecute(r)
}

/*
ExchangeInfo Exchange Info (MARKET_DATA)
Get /sapi/v1/equity/market/exchangeInfo

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/market-data#exchange-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Filter to a single US-equity ticker, e.g. `AAPL`. When omitted, returns all active symbols. An unknown ticker returns an empty `symbols` array (HTTP 200), not an error.
@return ApiExchangeInfoRequest
*/
func (a *MarketDataAPIService) ExchangeInfo(ctx context.Context) ApiExchangeInfoRequest {
	return ApiExchangeInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExchangeInfoResponse
func (a *MarketDataAPIService) ExchangeInfoExecute(r ApiExchangeInfoRequest) (*common.RestApiResponse[models.ExchangeInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/market/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
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

type ApiLatestQuoteRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// US-equity ticker, e.g. &#x60;AAPL&#x60;, &#x60;TSLA&#x60;. Case-insensitive; uppercased server-side.
func (r ApiLatestQuoteRequest) Symbol(symbol string) ApiLatestQuoteRequest {
	r.symbol = &symbol
	return r
}

func (r ApiLatestQuoteRequest) Execute() (*common.RestApiResponse[models.LatestQuoteResponse], error) {
	return r.ApiService.LatestQuoteExecute(r)
}

/*
LatestQuote Latest Quote (MARKET_DATA)
Get /sapi/v1/equity/market/quote

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/market-data#latest-quote

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  US-equity ticker, e.g. `AAPL`, `TSLA`. Case-insensitive; uppercased server-side.
@return ApiLatestQuoteRequest
*/
func (a *MarketDataAPIService) LatestQuote(ctx context.Context) ApiLatestQuoteRequest {
	return ApiLatestQuoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LatestQuoteResponse
func (a *MarketDataAPIService) LatestQuoteExecute(r ApiLatestQuoteRequest) (*common.RestApiResponse[models.LatestQuoteResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/market/quote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.LatestQuoteResponse](
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

type ApiTokenizedAssetsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiTokenizedAssetsRequest) Execute() (*common.RestApiResponse[models.TokenizedAssetsResponse], error) {
	return r.ApiService.TokenizedAssetsExecute(r)
}

/*
TokenizedAssets Tokenized Assets (MARKET_DATA)
Get /sapi/v1/equity/market/tokenized-assets

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/rest-api/market-data#tokenized-assets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiTokenizedAssetsRequest
*/
func (a *MarketDataAPIService) TokenizedAssets(ctx context.Context) ApiTokenizedAssetsRequest {
	return ApiTokenizedAssetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenizedAssetsResponse
func (a *MarketDataAPIService) TokenizedAssetsExecute(r ApiTokenizedAssetsRequest) (*common.RestApiResponse[models.TokenizedAssetsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/equity/market/tokenized-assets"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.TokenizedAssetsResponse](
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
