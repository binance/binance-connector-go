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

// OtcAPIService OtcAPI Service
type OtcAPIService Service

type ApiCreateOtcBlocktradeRequest struct {
	ctx           context.Context
	ApiService    *OtcAPIService
	marketId      *string
	tokenId       *string
	side          *models.GetQuoteSideParameter
	makerAmount   *string
	takerAmount   *string
	pricePerShare *string
	expiration    *int64
}

// PredictFun market id
func (r ApiCreateOtcBlocktradeRequest) MarketId(marketId string) ApiCreateOtcBlocktradeRequest {
	r.marketId = &marketId
	return r
}

// ERC-1155 outcome token id
func (r ApiCreateOtcBlocktradeRequest) TokenId(tokenId string) ApiCreateOtcBlocktradeRequest {
	r.tokenId = &tokenId
	return r
}

// Trade side. Enum: &#x60;BUY&#x60; (BID), &#x60;SELL&#x60; (ASK)
func (r ApiCreateOtcBlocktradeRequest) Side(side models.GetQuoteSideParameter) ApiCreateOtcBlocktradeRequest {
	r.side = &side
	return r
}

// Maker amount in wei. BID：USDT; ASK：shares
func (r ApiCreateOtcBlocktradeRequest) MakerAmount(makerAmount string) ApiCreateOtcBlocktradeRequest {
	r.makerAmount = &makerAmount
	return r
}

// Taker amount in wei. BID：shares; ASK：USDT
func (r ApiCreateOtcBlocktradeRequest) TakerAmount(takerAmount string) ApiCreateOtcBlocktradeRequest {
	r.takerAmount = &takerAmount
	return r
}

// Price per share (decimal ether, e.g. &#x60;0.65&#x60;)
func (r ApiCreateOtcBlocktradeRequest) PricePerShare(pricePerShare string) ApiCreateOtcBlocktradeRequest {
	r.pricePerShare = &pricePerShare
	return r
}

// Expiration timestamp in seconds (GTD order)
func (r ApiCreateOtcBlocktradeRequest) Expiration(expiration int64) ApiCreateOtcBlocktradeRequest {
	r.expiration = &expiration
	return r
}

func (r ApiCreateOtcBlocktradeRequest) Execute() (*common.RestApiResponse[models.CreateOtcBlocktradeResponse], error) {
	return r.ApiService.CreateOtcBlocktradeExecute(r)
}

/*
CreateOtcBlocktrade Create OTC Blocktrade (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/create

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#create-otc-blocktrade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param marketId -  PredictFun market id
@param tokenId -  ERC-1155 outcome token id
@param side -  Trade side. Enum: `BUY` (BID), `SELL` (ASK)
@param makerAmount -  Maker amount in wei. BID：USDT; ASK：shares
@param takerAmount -  Taker amount in wei. BID：shares; ASK：USDT
@param pricePerShare -  Price per share (decimal ether, e.g. `0.65`)
@param expiration -  Expiration timestamp in seconds (GTD order)
@return ApiCreateOtcBlocktradeRequest
*/
func (a *OtcAPIService) CreateOtcBlocktrade(ctx context.Context) ApiCreateOtcBlocktradeRequest {
	return ApiCreateOtcBlocktradeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateOtcBlocktradeResponse
func (a *OtcAPIService) CreateOtcBlocktradeExecute(r ApiCreateOtcBlocktradeRequest) (*common.RestApiResponse[models.CreateOtcBlocktradeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/create"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.marketId == nil {
		return nil, common.ReportError("marketId is required and must be specified")
	}

	if r.tokenId == nil {
		return nil, common.ReportError("tokenId is required and must be specified")
	}

	if r.side == nil {
		return nil, common.ReportError("side is required and must be specified")
	}

	if r.makerAmount == nil {
		return nil, common.ReportError("makerAmount is required and must be specified")
	}

	if r.takerAmount == nil {
		return nil, common.ReportError("takerAmount is required and must be specified")
	}

	if r.pricePerShare == nil {
		return nil, common.ReportError("pricePerShare is required and must be specified")
	}

	if r.expiration == nil {
		return nil, common.ReportError("expiration is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tokenId", r.tokenId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "side", r.side, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "makerAmount", r.makerAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "takerAmount", r.takerAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pricePerShare", r.pricePerShare, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "expiration", r.expiration, "form", "")

	resp, err := SendRequest[models.CreateOtcBlocktradeResponse](
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

type ApiFulfilOtcBlocktradeRequest struct {
	ctx         context.Context
	ApiService  *OtcAPIService
	orderId     *string
	secretToken *string
}

// Maker order id (returned by maker create)
func (r ApiFulfilOtcBlocktradeRequest) OrderId(orderId string) ApiFulfilOtcBlocktradeRequest {
	r.orderId = &orderId
	return r
}

// One-time fulfilment token the maker shared out-of-band
func (r ApiFulfilOtcBlocktradeRequest) SecretToken(secretToken string) ApiFulfilOtcBlocktradeRequest {
	r.secretToken = &secretToken
	return r
}

func (r ApiFulfilOtcBlocktradeRequest) Execute() (*common.RestApiResponse[models.FulfilOtcBlocktradeResponse], error) {
	return r.ApiService.FulfilOtcBlocktradeExecute(r)
}

/*
FulfilOtcBlocktrade Fulfil OTC Blocktrade (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/fulfil

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#fulfil-otc-blocktrade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Maker order id (returned by maker create)
@param secretToken -  One-time fulfilment token the maker shared out-of-band
@return ApiFulfilOtcBlocktradeRequest
*/
func (a *OtcAPIService) FulfilOtcBlocktrade(ctx context.Context) ApiFulfilOtcBlocktradeRequest {
	return ApiFulfilOtcBlocktradeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FulfilOtcBlocktradeResponse
func (a *OtcAPIService) FulfilOtcBlocktradeExecute(r ApiFulfilOtcBlocktradeRequest) (*common.RestApiResponse[models.FulfilOtcBlocktradeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/fulfil"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}

	if r.secretToken == nil {
		return nil, common.ReportError("secretToken is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "secretToken", r.secretToken, "form", "")

	resp, err := SendRequest[models.FulfilOtcBlocktradeResponse](
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

type ApiGetOtcBlocktradeDetailRequest struct {
	ctx        context.Context
	ApiService *OtcAPIService
	orderId    *string
}

// Maker order id (returned by create)
func (r ApiGetOtcBlocktradeDetailRequest) OrderId(orderId string) ApiGetOtcBlocktradeDetailRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetOtcBlocktradeDetailRequest) Execute() (*common.RestApiResponse[models.GetOtcBlocktradeDetailResponse], error) {
	return r.ApiService.GetOtcBlocktradeDetailExecute(r)
}

/*
GetOtcBlocktradeDetail Get OTC Blocktrade Detail (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/detail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#get-otc-blocktrade-detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Maker order id (returned by create)
@return ApiGetOtcBlocktradeDetailRequest
*/
func (a *OtcAPIService) GetOtcBlocktradeDetail(ctx context.Context) ApiGetOtcBlocktradeDetailRequest {
	return ApiGetOtcBlocktradeDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOtcBlocktradeDetailResponse
func (a *OtcAPIService) GetOtcBlocktradeDetailExecute(r ApiGetOtcBlocktradeDetailRequest) (*common.RestApiResponse[models.GetOtcBlocktradeDetailResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/detail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")

	resp, err := SendRequest[models.GetOtcBlocktradeDetailResponse](
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

type ApiGetOtcBlocktradeEventsRequest struct {
	ctx        context.Context
	ApiService *OtcAPIService
	first      *int32
	after      *string
	eventTypes *[]string
	marketId   *int64
}

// Page size
func (r ApiGetOtcBlocktradeEventsRequest) First(first int32) ApiGetOtcBlocktradeEventsRequest {
	r.first = &first
	return r
}

// Pagination cursor
func (r ApiGetOtcBlocktradeEventsRequest) After(after string) ApiGetOtcBlocktradeEventsRequest {
	r.after = &after
	return r
}

// Filter by event types (e.g. &#x60;[\&quot;MATCH_SUCCESS\&quot;,\&quot;EXPIRE\&quot;]&#x60;)
func (r ApiGetOtcBlocktradeEventsRequest) EventTypes(eventTypes []string) ApiGetOtcBlocktradeEventsRequest {
	r.eventTypes = &eventTypes
	return r
}

// Filter by market id
func (r ApiGetOtcBlocktradeEventsRequest) MarketId(marketId int64) ApiGetOtcBlocktradeEventsRequest {
	r.marketId = &marketId
	return r
}

func (r ApiGetOtcBlocktradeEventsRequest) Execute() (*common.RestApiResponse[models.GetOtcBlocktradeEventsResponse], error) {
	return r.ApiService.GetOtcBlocktradeEventsExecute(r)
}

/*
GetOtcBlocktradeEvents Get OTC Blocktrade Events (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/events

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#get-otc-blocktrade-events

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param first -  Page size
@param after -  Pagination cursor
@param eventTypes -  Filter by event types (e.g. `[\"MATCH_SUCCESS\",\"EXPIRE\"]`)
@param marketId -  Filter by market id
@return ApiGetOtcBlocktradeEventsRequest
*/
func (a *OtcAPIService) GetOtcBlocktradeEvents(ctx context.Context) ApiGetOtcBlocktradeEventsRequest {
	return ApiGetOtcBlocktradeEventsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOtcBlocktradeEventsResponse
func (a *OtcAPIService) GetOtcBlocktradeEventsExecute(r ApiGetOtcBlocktradeEventsRequest) (*common.RestApiResponse[models.GetOtcBlocktradeEventsResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/events"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.first != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "first", r.first, "form", "")
	}
	if r.after != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	}
	if r.eventTypes != nil {
		t := *r.eventTypes
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "eventTypes", t, "form", "multi")
	}
	if r.marketId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marketId", r.marketId, "form", "")
	}

	resp, err := SendRequest[models.GetOtcBlocktradeEventsResponse](
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

type ApiGetOtcReservedBalancesRequest struct {
	ctx        context.Context
	ApiService *OtcAPIService
	assets     *[]models.GetOtcReservedBalancesAssetsParameterInner
}

// Assets to query (max 50)
func (r ApiGetOtcReservedBalancesRequest) Assets(assets []models.GetOtcReservedBalancesAssetsParameterInner) ApiGetOtcReservedBalancesRequest {
	r.assets = &assets
	return r
}

func (r ApiGetOtcReservedBalancesRequest) Execute() (*common.RestApiResponse[models.GetOtcReservedBalancesResponse], error) {
	return r.ApiService.GetOtcReservedBalancesExecute(r)
}

/*
GetOtcReservedBalances Get OTC Reserved Balances (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/reserved-balances

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#get-otc-reserved-balances

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param assets -  Assets to query (max 50)
@return ApiGetOtcReservedBalancesRequest
*/
func (a *OtcAPIService) GetOtcReservedBalances(ctx context.Context) ApiGetOtcReservedBalancesRequest {
	return ApiGetOtcReservedBalancesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOtcReservedBalancesResponse
func (a *OtcAPIService) GetOtcReservedBalancesExecute(r ApiGetOtcReservedBalancesRequest) (*common.RestApiResponse[models.GetOtcReservedBalancesResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/reserved-balances"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.assets == nil {
		return nil, common.ReportError("assets is required and must be specified")
	}
	if len(*r.assets) > 50 {
		return nil, common.ReportError("assets must have less than 50 elements")
	}

	{
		t := *r.assets
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "assets", t, "form", "multi")
	}

	resp, err := SendRequest[models.GetOtcReservedBalancesResponse](
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

type ApiListOtcBlocktradesRequest struct {
	ctx        context.Context
	ApiService *OtcAPIService
	first      *int32
	after      *string
	status     *models.ListOtcBlocktradesStatusParameter
}

// Page size
func (r ApiListOtcBlocktradesRequest) First(first int32) ApiListOtcBlocktradesRequest {
	r.first = &first
	return r
}

// Pagination cursor (from previous response)
func (r ApiListOtcBlocktradesRequest) After(after string) ApiListOtcBlocktradesRequest {
	r.after = &after
	return r
}

// Filter by status. Enum: &#x60;OPEN&#x60;, &#x60;FULFILLED&#x60;, &#x60;MATCHED&#x60;, &#x60;CANCELLED&#x60;, &#x60;EXPIRED&#x60;, &#x60;FAILED&#x60;
func (r ApiListOtcBlocktradesRequest) Status(status models.ListOtcBlocktradesStatusParameter) ApiListOtcBlocktradesRequest {
	r.status = &status
	return r
}

func (r ApiListOtcBlocktradesRequest) Execute() (*common.RestApiResponse[models.ListOtcBlocktradesResponse], error) {
	return r.ApiService.ListOtcBlocktradesExecute(r)
}

/*
ListOtcBlocktrades List OTC Blocktrades (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/list

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#list-otc-blocktrades

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param first -  Page size
@param after -  Pagination cursor (from previous response)
@param status -  Filter by status. Enum: `OPEN`, `FULFILLED`, `MATCHED`, `CANCELLED`, `EXPIRED`, `FAILED`
@return ApiListOtcBlocktradesRequest
*/
func (a *OtcAPIService) ListOtcBlocktrades(ctx context.Context) ApiListOtcBlocktradesRequest {
	return ApiListOtcBlocktradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListOtcBlocktradesResponse
func (a *OtcAPIService) ListOtcBlocktradesExecute(r ApiListOtcBlocktradesRequest) (*common.RestApiResponse[models.ListOtcBlocktradesResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.first != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "first", r.first, "form", "")
	}
	if r.after != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "form", "")
	}
	if r.status != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}

	resp, err := SendRequest[models.ListOtcBlocktradesResponse](
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

type ApiPreviewOtcBlocktradeRequest struct {
	ctx         context.Context
	ApiService  *OtcAPIService
	secretToken *string
}

// One-time fulfilment token the maker shared out-of-band
func (r ApiPreviewOtcBlocktradeRequest) SecretToken(secretToken string) ApiPreviewOtcBlocktradeRequest {
	r.secretToken = &secretToken
	return r
}

func (r ApiPreviewOtcBlocktradeRequest) Execute() (*common.RestApiResponse[models.PreviewOtcBlocktradeResponse], error) {
	return r.ApiService.PreviewOtcBlocktradeExecute(r)
}

/*
PreviewOtcBlocktrade Preview OTC Blocktrade (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/preview

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#preview-otc-blocktrade

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param secretToken -  One-time fulfilment token the maker shared out-of-band
@return ApiPreviewOtcBlocktradeRequest
*/
func (a *OtcAPIService) PreviewOtcBlocktrade(ctx context.Context) ApiPreviewOtcBlocktradeRequest {
	return ApiPreviewOtcBlocktradeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PreviewOtcBlocktradeResponse
func (a *OtcAPIService) PreviewOtcBlocktradeExecute(r ApiPreviewOtcBlocktradeRequest) (*common.RestApiResponse[models.PreviewOtcBlocktradeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/preview"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.secretToken == nil {
		return nil, common.ReportError("secretToken is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "secretToken", r.secretToken, "form", "")

	resp, err := SendRequest[models.PreviewOtcBlocktradeResponse](
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

type ApiRemoveOtcBlocktradesRequest struct {
	ctx        context.Context
	ApiService *OtcAPIService
	orderIds   *[]string
}

// Order ids to remove (max 100). Must be the &#x60;orderId&#x60; returned by Create OTC Blocktrade
func (r ApiRemoveOtcBlocktradesRequest) OrderIds(orderIds []string) ApiRemoveOtcBlocktradesRequest {
	r.orderIds = &orderIds
	return r
}

func (r ApiRemoveOtcBlocktradesRequest) Execute() (*common.RestApiResponse[models.RemoveOtcBlocktradesResponse], error) {
	return r.ApiService.RemoveOtcBlocktradesExecute(r)
}

/*
RemoveOtcBlocktrades Remove OTC Blocktrades (PREDICTION_TRADE)
Post /sapi/v1/w3w/wallet/prediction/otc/blocktrade/remove

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/rest-api/otc#remove-otc-blocktrades

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderIds -  Order ids to remove (max 100). Must be the `orderId` returned by Create OTC Blocktrade
@return ApiRemoveOtcBlocktradesRequest
*/
func (a *OtcAPIService) RemoveOtcBlocktrades(ctx context.Context) ApiRemoveOtcBlocktradesRequest {
	return ApiRemoveOtcBlocktradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RemoveOtcBlocktradesResponse
func (a *OtcAPIService) RemoveOtcBlocktradesExecute(r ApiRemoveOtcBlocktradesRequest) (*common.RestApiResponse[models.RemoveOtcBlocktradesResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/w3w/wallet/prediction/otc/blocktrade/remove"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderIds == nil {
		return nil, common.ReportError("orderIds is required and must be specified")
	}

	{
		t := *r.orderIds
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderIds", t, "form", "multi")
	}

	resp, err := SendRequest[models.RemoveOtcBlocktradesResponse](
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
