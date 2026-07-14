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

// PositionAPIService PositionAPI Service
type PositionAPIService Service

type ApiGetPositionByTokenRequest struct {
	ctx           context.Context
	ApiService    *PositionAPIService
	walletAddress *string
	tokenId       *string
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiGetPositionByTokenRequest) WalletAddress(walletAddress string) ApiGetPositionByTokenRequest {
	r.walletAddress = &walletAddress
	return r
}

// Prediction outcome token ID
func (r ApiGetPositionByTokenRequest) TokenId(tokenId string) ApiGetPositionByTokenRequest {
	r.tokenId = &tokenId
	return r
}

// Request validity window in milliseconds
func (r ApiGetPositionByTokenRequest) RecvWindow(recvWindow int64) ApiGetPositionByTokenRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPositionByTokenRequest) Execute() (*common.RestApiResponse[models.GetPositionByTokenResponse], error) {
	return r.ApiService.GetPositionByTokenExecute(r)
}

/*
GetPositionByToken Get Position by Token (USER_DATA)
Get /sapi/v1/w3w/wallet/prediction/position/token

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/position#get-position-by-token

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tokenId -  Prediction outcome token ID
@param recvWindow -  Request validity window in milliseconds
@return ApiGetPositionByTokenRequest
*/
func (a *PositionAPIService) GetPositionByToken(ctx context.Context) ApiGetPositionByTokenRequest {
	return ApiGetPositionByTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPositionByTokenResponse
func (a *PositionAPIService) GetPositionByTokenExecute(r ApiGetPositionByTokenRequest) (*common.RestApiResponse[models.GetPositionByTokenResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/position/token"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	if r.tokenId == nil {
		return nil, common.ReportError("tokenId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPositionByTokenResponse](
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

type ApiQueryPnLRequest struct {
	ctx           context.Context
	ApiService    *PositionAPIService
	walletAddress *string
	tokenId       *string
	marketId      *int64
	marketTopicId *int64
	activeOnly    *bool
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryPnLRequest) WalletAddress(walletAddress string) ApiQueryPnLRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by prediction token ID
func (r ApiQueryPnLRequest) TokenId(tokenId string) ApiQueryPnLRequest {
	r.tokenId = &tokenId
	return r
}

// Filter by market ID. Must be &gt; 0
func (r ApiQueryPnLRequest) MarketId(marketId int64) ApiQueryPnLRequest {
	r.marketId = &marketId
	return r
}

// Filter by market topic ID. Must be &gt; 0
func (r ApiQueryPnLRequest) MarketTopicId(marketTopicId int64) ApiQueryPnLRequest {
	r.marketTopicId = &marketTopicId
	return r
}

// If &#x60;true&#x60;, return only active (unresolved) positions
func (r ApiQueryPnLRequest) ActiveOnly(activeOnly bool) ApiQueryPnLRequest {
	r.activeOnly = &activeOnly
	return r
}

// Request validity window in milliseconds
func (r ApiQueryPnLRequest) RecvWindow(recvWindow int64) ApiQueryPnLRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPnLRequest) Execute() (*common.RestApiResponse[models.QueryPnLResponse], error) {
	return r.ApiService.QueryPnLExecute(r)
}

/*
QueryPnL Query PnL (USER_DATA)
Get /sapi/v1/w3w/wallet/prediction/pnl/query

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/position#query-pn-l

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tokenId -  Filter by prediction token ID
@param marketId -  Filter by market ID. Must be > 0
@param marketTopicId -  Filter by market topic ID. Must be > 0
@param activeOnly -  If `true`, return only active (unresolved) positions
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryPnLRequest
*/
func (a *PositionAPIService) QueryPnL(ctx context.Context) ApiQueryPnLRequest {
	return ApiQueryPnLRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPnLResponse
func (a *PositionAPIService) QueryPnLExecute(r ApiQueryPnLRequest) (*common.RestApiResponse[models.QueryPnLResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/pnl/query"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.tokenId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")
	}
	if r.marketId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	}
	if r.marketTopicId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketTopicId", r.marketTopicId, "form", "")
	}
	if r.activeOnly != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "activeOnly", r.activeOnly, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPnLResponse](
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

type ApiQueryPositionsRequest struct {
	ctx           context.Context
	ApiService    *PositionAPIService
	walletAddress *string
	tab           *string
	offset        *int32
	limit         *int32
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryPositionsRequest) WalletAddress(walletAddress string) ApiQueryPositionsRequest {
	r.walletAddress = &walletAddress
	return r
}

// Position status tab. Values from &#x60;PositionQueryType&#x60;. Default &#x60;ONGOING&#x60;
func (r ApiQueryPositionsRequest) Tab(tab string) ApiQueryPositionsRequest {
	r.tab = &tab
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiQueryPositionsRequest) Offset(offset int32) ApiQueryPositionsRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiQueryPositionsRequest) Limit(limit int32) ApiQueryPositionsRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiQueryPositionsRequest) RecvWindow(recvWindow int64) ApiQueryPositionsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPositionsRequest) Execute() (*common.RestApiResponse[models.QueryPositionsResponse], error) {
	return r.ApiService.QueryPositionsExecute(r)
}

/*
QueryPositions Query Positions (USER_DATA)
Get /sapi/v1/w3w/wallet/prediction/position/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/position#query-positions

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param tab -  Position status tab. Values from `PositionQueryType`. Default `ONGOING`
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryPositionsRequest
*/
func (a *PositionAPIService) QueryPositions(ctx context.Context) ApiQueryPositionsRequest {
	return ApiQueryPositionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPositionsResponse
func (a *PositionAPIService) QueryPositionsExecute(r ApiQueryPositionsRequest) (*common.RestApiResponse[models.QueryPositionsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/position/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.tab != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tab", r.tab, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPositionsResponse](
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

type ApiQueryPositionsByFilterRequest struct {
	ctx           context.Context
	ApiService    *PositionAPIService
	walletAddress *string
	marketTopicId *int64
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQueryPositionsByFilterRequest) WalletAddress(walletAddress string) ApiQueryPositionsByFilterRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by market topic ID
func (r ApiQueryPositionsByFilterRequest) MarketTopicId(marketTopicId int64) ApiQueryPositionsByFilterRequest {
	r.marketTopicId = &marketTopicId
	return r
}

// Request validity window in milliseconds
func (r ApiQueryPositionsByFilterRequest) RecvWindow(recvWindow int64) ApiQueryPositionsByFilterRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPositionsByFilterRequest) Execute() (*common.RestApiResponse[models.QueryPositionsByFilterResponse], error) {
	return r.ApiService.QueryPositionsByFilterExecute(r)
}

/*
QueryPositionsByFilter Query Positions by Filter (USER_DATA)
Get /sapi/v1/w3w/wallet/prediction/position/filter

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/position#query-positions-by-filter

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param marketTopicId -  Filter by market topic ID
@param recvWindow -  Request validity window in milliseconds
@return ApiQueryPositionsByFilterRequest
*/
func (a *PositionAPIService) QueryPositionsByFilter(ctx context.Context) ApiQueryPositionsByFilterRequest {
	return ApiQueryPositionsByFilterRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPositionsByFilterResponse
func (a *PositionAPIService) QueryPositionsByFilterExecute(r ApiQueryPositionsByFilterRequest) (*common.RestApiResponse[models.QueryPositionsByFilterResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/position/filter"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	}
	if r.marketTopicId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketTopicId", r.marketTopicId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPositionsByFilterResponse](
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

type ApiQuerySettledPositionHistoryRequest struct {
	ctx           context.Context
	ApiService    *PositionAPIService
	walletAddress *string
	l1Category    *string
	result        *int32
	startDate     *string
	endDate       *string
	offset        *int32
	limit         *int32
	recvWindow    *int64
}

// User&#39;s prediction wallet address
func (r ApiQuerySettledPositionHistoryRequest) WalletAddress(walletAddress string) ApiQuerySettledPositionHistoryRequest {
	r.walletAddress = &walletAddress
	return r
}

// Filter by level-1 category
func (r ApiQuerySettledPositionHistoryRequest) L1Category(l1Category string) ApiQuerySettledPositionHistoryRequest {
	r.l1Category = &l1Category
	return r
}

// Settlement result filter
func (r ApiQuerySettledPositionHistoryRequest) Result(result int32) ApiQuerySettledPositionHistoryRequest {
	r.result = &result
	return r
}

// Start date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≤ &#x60;endDate&#x60;
func (r ApiQuerySettledPositionHistoryRequest) StartDate(startDate string) ApiQuerySettledPositionHistoryRequest {
	r.startDate = &startDate
	return r
}

// End date. Format: &#x60;yyyy-MM-dd&#x60;. Must be ≥ &#x60;startDate&#x60;
func (r ApiQuerySettledPositionHistoryRequest) EndDate(endDate string) ApiQuerySettledPositionHistoryRequest {
	r.endDate = &endDate
	return r
}

// Pagination offset. Default &#x60;0&#x60;
func (r ApiQuerySettledPositionHistoryRequest) Offset(offset int32) ApiQuerySettledPositionHistoryRequest {
	r.offset = &offset
	return r
}

// Page size. Default &#x60;20&#x60;, range 1–100
func (r ApiQuerySettledPositionHistoryRequest) Limit(limit int32) ApiQuerySettledPositionHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiQuerySettledPositionHistoryRequest) RecvWindow(recvWindow int64) ApiQuerySettledPositionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySettledPositionHistoryRequest) Execute() (*common.RestApiResponse[models.QuerySettledPositionHistoryResponse], error) {
	return r.ApiService.QuerySettledPositionHistoryExecute(r)
}

/*
QuerySettledPositionHistory Query Settled Position History (USER_DATA)
Get /sapi/v1/w3w/wallet/prediction/position/settled-history

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/position#query-settled-position-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param walletAddress -  User's prediction wallet address
@param l1Category -  Filter by level-1 category
@param result -  Settlement result filter
@param startDate -  Start date. Format: `yyyy-MM-dd`. Must be ≤ `endDate`
@param endDate -  End date. Format: `yyyy-MM-dd`. Must be ≥ `startDate`
@param offset -  Pagination offset. Default `0`
@param limit -  Page size. Default `20`, range 1–100
@param recvWindow -  Request validity window in milliseconds
@return ApiQuerySettledPositionHistoryRequest
*/
func (a *PositionAPIService) QuerySettledPositionHistory(ctx context.Context) ApiQuerySettledPositionHistoryRequest {
	return ApiQuerySettledPositionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySettledPositionHistoryResponse
func (a *PositionAPIService) QuerySettledPositionHistoryExecute(r ApiQuerySettledPositionHistoryRequest) (*common.RestApiResponse[models.QuerySettledPositionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/position/settled-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.walletAddress == nil {
		return nil, common.ReportError("walletAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletAddress", r.walletAddress, "form", "")
	if r.l1Category != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "l1Category", r.l1Category, "form", "")
	}
	if r.result != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "result", r.result, "form", "")
	}
	if r.startDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySettledPositionHistoryResponse](
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
