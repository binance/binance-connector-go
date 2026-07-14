/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package binancew3wpredictionrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiGetMarketDetailRequest struct {
	ctx           context.Context
	ApiService    *MarketDataAPIService
	marketTopicId *int64
}

// Market topic ID. Must be &gt; 0
func (r ApiGetMarketDetailRequest) MarketTopicId(marketTopicId int64) ApiGetMarketDetailRequest {
	r.marketTopicId = &marketTopicId
	return r
}

func (r ApiGetMarketDetailRequest) Execute() (*common.RestApiResponse[models.GetMarketDetailResponse], error) {
	return r.ApiService.GetMarketDetailExecute(r)
}

/*
GetMarketDetail Get Market Detail
Get /sapi/v1/w3w/wallet/prediction/market/detail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#get-market-detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param marketTopicId -  Market topic ID. Must be > 0
@return ApiGetMarketDetailRequest
*/
func (a *MarketDataAPIService) GetMarketDetail(ctx context.Context) ApiGetMarketDetailRequest {
	return ApiGetMarketDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMarketDetailResponse
func (a *MarketDataAPIService) GetMarketDetailExecute(r ApiGetMarketDetailRequest) (*common.RestApiResponse[models.GetMarketDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/market/detail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.marketTopicId == nil {
		return nil, common.ReportError("marketTopicId is required and must be specified")
	}
	if *r.marketTopicId < 1 {
		return nil, common.ReportError("marketTopicId must be greater than 1")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketTopicId", r.marketTopicId, "form", "")

	resp, err := SendRequest[models.GetMarketDetailResponse](
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

type ApiListPredictionCategoriesRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiListPredictionCategoriesRequest) Execute() (*common.RestApiResponse[models.ListPredictionCategoriesResponse], error) {
	return r.ApiService.ListPredictionCategoriesExecute(r)
}

/*
ListPredictionCategories List Prediction Categories
Get /sapi/v1/w3w/wallet/prediction/category/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#list-prediction-categories

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiListPredictionCategoriesRequest
*/
func (a *MarketDataAPIService) ListPredictionCategories(ctx context.Context) ApiListPredictionCategoriesRequest {
	return ApiListPredictionCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListPredictionCategoriesResponse
func (a *MarketDataAPIService) ListPredictionCategoriesExecute(r ApiListPredictionCategoriesRequest) (*common.RestApiResponse[models.ListPredictionCategoriesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/category/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.ListPredictionCategoriesResponse](
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

type ApiListPredictionMarketsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	l1Category *string
	l2Category *string
	sortBy     *models.ListPredictionMarketsSortByParameter
	orderBy    *models.ListPredictionMarketsOrderByParameter
	offset     *int32
	limit      *int32
}

// Level-1 category filter
func (r ApiListPredictionMarketsRequest) L1Category(l1Category string) ApiListPredictionMarketsRequest {
	r.l1Category = &l1Category
	return r
}

// Level-2 category filter
func (r ApiListPredictionMarketsRequest) L2Category(l2Category string) ApiListPredictionMarketsRequest {
	r.l2Category = &l2Category
	return r
}

// Sort field. Enum: &#x60;RECOMMENDED&#x60;, &#x60;VOLUME&#x60;, &#x60;PARTICIPANTS&#x60;, &#x60;CREATED_TIME&#x60;, &#x60;END_DATE&#x60;
func (r ApiListPredictionMarketsRequest) SortBy(sortBy models.ListPredictionMarketsSortByParameter) ApiListPredictionMarketsRequest {
	r.sortBy = &sortBy
	return r
}

// Sort direction. Enum: &#x60;ASC&#x60;, &#x60;DESC&#x60;
func (r ApiListPredictionMarketsRequest) OrderBy(orderBy models.ListPredictionMarketsOrderByParameter) ApiListPredictionMarketsRequest {
	r.orderBy = &orderBy
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiListPredictionMarketsRequest) Offset(offset int32) ApiListPredictionMarketsRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiListPredictionMarketsRequest) Limit(limit int32) ApiListPredictionMarketsRequest {
	r.limit = &limit
	return r
}

func (r ApiListPredictionMarketsRequest) Execute() (*common.RestApiResponse[models.ListPredictionMarketsResponse], error) {
	return r.ApiService.ListPredictionMarketsExecute(r)
}

/*
ListPredictionMarkets List Prediction Markets
Get /sapi/v1/w3w/wallet/prediction/market/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#list-prediction-markets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param l1Category -  Level-1 category filter
@param l2Category -  Level-2 category filter
@param sortBy -  Sort field. Enum: `RECOMMENDED`, `VOLUME`, `PARTICIPANTS`, `CREATED_TIME`, `END_DATE`
@param orderBy -  Sort direction. Enum: `ASC`, `DESC`
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@return ApiListPredictionMarketsRequest
*/
func (a *MarketDataAPIService) ListPredictionMarkets(ctx context.Context) ApiListPredictionMarketsRequest {
	return ApiListPredictionMarketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListPredictionMarketsResponse
func (a *MarketDataAPIService) ListPredictionMarketsExecute(r ApiListPredictionMarketsRequest) (*common.RestApiResponse[models.ListPredictionMarketsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/market/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.l1Category != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "l1Category", r.l1Category, "form", "")
	}
	if r.l2Category != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "l2Category", r.l2Category, "form", "")
	}
	if r.sortBy != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
	}
	if r.orderBy != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.ListPredictionMarketsResponse](
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

type ApiMarketSearchRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	query      *string
	topK       *int32
}

// Search keyword. Not blank
func (r ApiMarketSearchRequest) Query(query string) ApiMarketSearchRequest {
	r.query = &query
	return r
}

// Max number of results to return. Default &#x60;20&#x60;, range 1–50
func (r ApiMarketSearchRequest) TopK(topK int32) ApiMarketSearchRequest {
	r.topK = &topK
	return r
}

func (r ApiMarketSearchRequest) Execute() (*common.RestApiResponse[models.MarketSearchResponse], error) {
	return r.ApiService.MarketSearchExecute(r)
}

/*
MarketSearch Market Search
Get /sapi/v1/w3w/wallet/prediction/market/search

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#market-search

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param query -  Search keyword. Not blank
@param topK -  Max number of results to return. Default `20`, range 1–50
@return ApiMarketSearchRequest
*/
func (a *MarketDataAPIService) MarketSearch(ctx context.Context) ApiMarketSearchRequest {
	return ApiMarketSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarketSearchResponse
func (a *MarketDataAPIService) MarketSearchExecute(r ApiMarketSearchRequest) (*common.RestApiResponse[models.MarketSearchResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/market/search"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.query == nil {
		return nil, common.ReportError("query is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	if r.topK != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "topK", r.topK, "form", "")
	}

	resp, err := SendRequest[models.MarketSearchResponse](
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

type ApiQueryLastTradePriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	marketId   *int64
}

// Market ID. Must be &gt; 0
func (r ApiQueryLastTradePriceRequest) MarketId(marketId int64) ApiQueryLastTradePriceRequest {
	r.marketId = &marketId
	return r
}

func (r ApiQueryLastTradePriceRequest) Execute() (*common.RestApiResponse[models.QueryLastTradePriceResponse], error) {
	return r.ApiService.QueryLastTradePriceExecute(r)
}

/*
QueryLastTradePrice Query Last Trade Price
Get /sapi/v1/w3w/wallet/prediction/order-book/last-trade-price

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#query-last-trade-price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param marketId -  Market ID. Must be > 0
@return ApiQueryLastTradePriceRequest
*/
func (a *MarketDataAPIService) QueryLastTradePrice(ctx context.Context) ApiQueryLastTradePriceRequest {
	return ApiQueryLastTradePriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryLastTradePriceResponse
func (a *MarketDataAPIService) QueryLastTradePriceExecute(r ApiQueryLastTradePriceRequest) (*common.RestApiResponse[models.QueryLastTradePriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/order-book/last-trade-price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.marketId == nil {
		return nil, common.ReportError("marketId is required and must be specified")
	}
	if *r.marketId < 1 {
		return nil, common.ReportError("marketId must be greater than 1")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")

	resp, err := SendRequest[models.QueryLastTradePriceResponse](
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

type ApiQueryOrderBookRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	vendor     *string
	marketId   *int64
	tokenId    *string
}

// Vendor identifier (e.g. &#x60;predict_fun&#x60;)
func (r ApiQueryOrderBookRequest) Vendor(vendor string) ApiQueryOrderBookRequest {
	r.vendor = &vendor
	return r
}

// Market ID. Must be &gt; 0
func (r ApiQueryOrderBookRequest) MarketId(marketId int64) ApiQueryOrderBookRequest {
	r.marketId = &marketId
	return r
}

// Prediction outcome token ID
func (r ApiQueryOrderBookRequest) TokenId(tokenId string) ApiQueryOrderBookRequest {
	r.tokenId = &tokenId
	return r
}

func (r ApiQueryOrderBookRequest) Execute() (*common.RestApiResponse[models.QueryOrderBookResponse], error) {
	return r.ApiService.QueryOrderBookExecute(r)
}

/*
QueryOrderBook Query Order Book
Get /sapi/v1/w3w/wallet/prediction/order-book

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/market-data#query-order-book

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param vendor -  Vendor identifier (e.g. `predict_fun`)
@param marketId -  Market ID. Must be > 0
@param tokenId -  Prediction outcome token ID
@return ApiQueryOrderBookRequest
*/
func (a *MarketDataAPIService) QueryOrderBook(ctx context.Context) ApiQueryOrderBookRequest {
	return ApiQueryOrderBookRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryOrderBookResponse
func (a *MarketDataAPIService) QueryOrderBookExecute(r ApiQueryOrderBookRequest) (*common.RestApiResponse[models.QueryOrderBookResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/order-book"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.vendor == nil {
		return nil, common.ReportError("vendor is required and must be specified")
	}

	if r.marketId == nil {
		return nil, common.ReportError("marketId is required and must be specified")
	}
	if *r.marketId < 1 {
		return nil, common.ReportError("marketId must be greater than 1")
	}

	if r.tokenId == nil {
		return nil, common.ReportError("tokenId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "vendor", r.vendor, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")

	resp, err := SendRequest[models.QueryOrderBookResponse](
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
