/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package binancemargintradingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiCrossMarginCollateralRatioRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiCrossMarginCollateralRatioRequest) Execute() (*common.RestApiResponse[models.CrossMarginCollateralRatioResponse], error) {
	return r.ApiService.CrossMarginCollateralRatioExecute(r)
}

/*
CrossMarginCollateralRatio Cross margin collateral ratio (MARKET_DATA)
Get /sapi/v1/margin/crossMarginCollateralRatio

https://developers.binance.com/docs/margin_trading/market-data/Cross-margin-collateral-ratio

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiCrossMarginCollateralRatioRequest
*/
func (a *MarketDataAPIService) CrossMarginCollateralRatio(ctx context.Context) ApiCrossMarginCollateralRatioRequest {
	return ApiCrossMarginCollateralRatioRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CrossMarginCollateralRatioResponse
func (a *MarketDataAPIService) CrossMarginCollateralRatioExecute(r ApiCrossMarginCollateralRatioRequest) (*common.RestApiResponse[models.CrossMarginCollateralRatioResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/crossMarginCollateralRatio"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.CrossMarginCollateralRatioResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAllCrossMarginPairsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

// isolated margin pair
func (r ApiGetAllCrossMarginPairsRequest) Symbol(symbol string) ApiGetAllCrossMarginPairsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetAllCrossMarginPairsRequest) Execute() (*common.RestApiResponse[models.GetAllCrossMarginPairsResponse], error) {
	return r.ApiService.GetAllCrossMarginPairsExecute(r)
}

/*
GetAllCrossMarginPairs Get All Cross Margin Pairs (MARKET_DATA)
Get /sapi/v1/margin/allPairs

https://developers.binance.com/docs/margin_trading/market-data/Get-All-Cross-Margin-Pairs

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@return ApiGetAllCrossMarginPairsRequest
*/
func (a *MarketDataAPIService) GetAllCrossMarginPairs(ctx context.Context) ApiGetAllCrossMarginPairsRequest {
	return ApiGetAllCrossMarginPairsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAllCrossMarginPairsResponse
func (a *MarketDataAPIService) GetAllCrossMarginPairsExecute(r ApiGetAllCrossMarginPairsRequest) (*common.RestApiResponse[models.GetAllCrossMarginPairsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/allPairs"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.GetAllCrossMarginPairsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAllIsolatedMarginSymbolRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	recvWindow *int64
}

// isolated margin pair
func (r ApiGetAllIsolatedMarginSymbolRequest) Symbol(symbol string) ApiGetAllIsolatedMarginSymbolRequest {
	r.symbol = &symbol
	return r
}

// No more than 60000
func (r ApiGetAllIsolatedMarginSymbolRequest) RecvWindow(recvWindow int64) ApiGetAllIsolatedMarginSymbolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetAllIsolatedMarginSymbolRequest) Execute() (*common.RestApiResponse[models.GetAllIsolatedMarginSymbolResponse], error) {
	return r.ApiService.GetAllIsolatedMarginSymbolExecute(r)
}

/*
GetAllIsolatedMarginSymbol Get All Isolated Margin Symbol(MARKET_DATA)
Get /sapi/v1/margin/isolated/allPairs

https://developers.binance.com/docs/margin_trading/market-data/Get-All-Isolated-Margin-Symbol

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  isolated margin pair
@param recvWindow -  No more than 60000
@return ApiGetAllIsolatedMarginSymbolRequest
*/
func (a *MarketDataAPIService) GetAllIsolatedMarginSymbol(ctx context.Context) ApiGetAllIsolatedMarginSymbolRequest {
	return ApiGetAllIsolatedMarginSymbolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAllIsolatedMarginSymbolResponse
func (a *MarketDataAPIService) GetAllIsolatedMarginSymbolExecute(r ApiGetAllIsolatedMarginSymbolRequest) (*common.RestApiResponse[models.GetAllIsolatedMarginSymbolResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolated/allPairs"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetAllIsolatedMarginSymbolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAllMarginAssetsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	asset      *string
}

func (r ApiGetAllMarginAssetsRequest) Asset(asset string) ApiGetAllMarginAssetsRequest {
	r.asset = &asset
	return r
}

func (r ApiGetAllMarginAssetsRequest) Execute() (*common.RestApiResponse[models.GetAllMarginAssetsResponse], error) {
	return r.ApiService.GetAllMarginAssetsExecute(r)
}

/*
GetAllMarginAssets Get All Margin Assets (MARKET_DATA)
Get /sapi/v1/margin/allAssets

https://developers.binance.com/docs/margin_trading/market-data/Get-All-Margin-Assets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@return ApiGetAllMarginAssetsRequest
*/
func (a *MarketDataAPIService) GetAllMarginAssets(ctx context.Context) ApiGetAllMarginAssetsRequest {
	return ApiGetAllMarginAssetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAllMarginAssetsResponse
func (a *MarketDataAPIService) GetAllMarginAssetsExecute(r ApiGetAllMarginAssetsRequest) (*common.RestApiResponse[models.GetAllMarginAssetsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/allAssets"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}

	resp, err := SendRequest[models.GetAllMarginAssetsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDelistScheduleRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiGetDelistScheduleRequest) RecvWindow(recvWindow int64) ApiGetDelistScheduleRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDelistScheduleRequest) Execute() (*common.RestApiResponse[models.GetDelistScheduleResponse], error) {
	return r.ApiService.GetDelistScheduleExecute(r)
}

/*
GetDelistSchedule Get Delist Schedule (MARKET_DATA)
Get /sapi/v1/margin/delist-schedule

https://developers.binance.com/docs/margin_trading/market-data/Get-Delist-Schedule

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiGetDelistScheduleRequest
*/
func (a *MarketDataAPIService) GetDelistSchedule(ctx context.Context) ApiGetDelistScheduleRequest {
	return ApiGetDelistScheduleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDelistScheduleResponse
func (a *MarketDataAPIService) GetDelistScheduleExecute(r ApiGetDelistScheduleRequest) (*common.RestApiResponse[models.GetDelistScheduleResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/delist-schedule"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDelistScheduleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLimitPricePairsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiGetLimitPricePairsRequest) Execute() (*common.RestApiResponse[models.GetLimitPricePairsResponse], error) {
	return r.ApiService.GetLimitPricePairsExecute(r)
}

/*
GetLimitPricePairs Get Limit Price Pairs(MARKET_DATA)
Get /sapi/v1/margin/limit-price-pairs

https://developers.binance.com/docs/margin_trading/market-data/Get-Limit-Price-Pairs

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiGetLimitPricePairsRequest
*/
func (a *MarketDataAPIService) GetLimitPricePairs(ctx context.Context) ApiGetLimitPricePairsRequest {
	return ApiGetLimitPricePairsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLimitPricePairsResponse
func (a *MarketDataAPIService) GetLimitPricePairsExecute(r ApiGetLimitPricePairsRequest) (*common.RestApiResponse[models.GetLimitPricePairsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/limit-price-pairs"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetLimitPricePairsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetListScheduleRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	recvWindow *int64
}

// No more than 60000
func (r ApiGetListScheduleRequest) RecvWindow(recvWindow int64) ApiGetListScheduleRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetListScheduleRequest) Execute() (*common.RestApiResponse[models.GetListScheduleResponse], error) {
	return r.ApiService.GetListScheduleExecute(r)
}

/*
GetListSchedule Get list Schedule (MARKET_DATA)
Get /sapi/v1/margin/list-schedule

https://developers.binance.com/docs/margin_trading/market-data/Get-list-Schedule

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  No more than 60000
@return ApiGetListScheduleRequest
*/
func (a *MarketDataAPIService) GetListSchedule(ctx context.Context) ApiGetListScheduleRequest {
	return ApiGetListScheduleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetListScheduleResponse
func (a *MarketDataAPIService) GetListScheduleExecute(r ApiGetListScheduleRequest) (*common.RestApiResponse[models.GetListScheduleResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/list-schedule"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetListScheduleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryIsolatedMarginTierDataRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
	tier       *int64
	recvWindow *int64
}

func (r ApiQueryIsolatedMarginTierDataRequest) Symbol(symbol string) ApiQueryIsolatedMarginTierDataRequest {
	r.symbol = &symbol
	return r
}

// All margin tier data will be returned if tier is omitted
func (r ApiQueryIsolatedMarginTierDataRequest) Tier(tier int64) ApiQueryIsolatedMarginTierDataRequest {
	r.tier = &tier
	return r
}

// No more than 60000
func (r ApiQueryIsolatedMarginTierDataRequest) RecvWindow(recvWindow int64) ApiQueryIsolatedMarginTierDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryIsolatedMarginTierDataRequest) Execute() (*common.RestApiResponse[models.QueryIsolatedMarginTierDataResponse], error) {
	return r.ApiService.QueryIsolatedMarginTierDataExecute(r)
}

/*
QueryIsolatedMarginTierData Query Isolated Margin Tier Data (USER_DATA)
Get /sapi/v1/margin/isolatedMarginTier

https://developers.binance.com/docs/margin_trading/market-data/Query-Isolated-Margin-Tier-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param tier -  All margin tier data will be returned if tier is omitted
@param recvWindow -  No more than 60000
@return ApiQueryIsolatedMarginTierDataRequest
*/
func (a *MarketDataAPIService) QueryIsolatedMarginTierData(ctx context.Context) ApiQueryIsolatedMarginTierDataRequest {
	return ApiQueryIsolatedMarginTierDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryIsolatedMarginTierDataResponse
func (a *MarketDataAPIService) QueryIsolatedMarginTierDataExecute(r ApiQueryIsolatedMarginTierDataRequest) (*common.RestApiResponse[models.QueryIsolatedMarginTierDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/isolatedMarginTier"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.tier != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tier", r.tier, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryIsolatedMarginTierDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest) Execute() (*common.RestApiResponse[models.QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse], error) {
	return r.ApiService.QueryLiabilityCoinLeverageBracketInCrossMarginProModeExecute(r)
}

/*
QueryLiabilityCoinLeverageBracketInCrossMarginProMode Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
Get /sapi/v1/margin/leverageBracket

https://developers.binance.com/docs/margin_trading/market-data/Query-Liability-Coin-Leverage-Bracket-in-Cross-Margin-Pro-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest
*/
func (a *MarketDataAPIService) QueryLiabilityCoinLeverageBracketInCrossMarginProMode(ctx context.Context) ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest {
	return ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse
func (a *MarketDataAPIService) QueryLiabilityCoinLeverageBracketInCrossMarginProModeExecute(r ApiQueryLiabilityCoinLeverageBracketInCrossMarginProModeRequest) (*common.RestApiResponse[models.QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.QueryLiabilityCoinLeverageBracketInCrossMarginProModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginAvailableInventoryRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	type_      *string
}

// &#x60;MARGIN&#x60;,&#x60;ISOLATED&#x60;
func (r ApiQueryMarginAvailableInventoryRequest) Type(type_ string) ApiQueryMarginAvailableInventoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiQueryMarginAvailableInventoryRequest) Execute() (*common.RestApiResponse[models.QueryMarginAvailableInventoryResponse], error) {
	return r.ApiService.QueryMarginAvailableInventoryExecute(r)
}

/*
QueryMarginAvailableInventory Query Margin Available Inventory(USER_DATA)
Get /sapi/v1/margin/available-inventory

https://developers.binance.com/docs/margin_trading/market-data/Query-margin-avaliable-inventory

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -  `MARGIN`,`ISOLATED`
@return ApiQueryMarginAvailableInventoryRequest
*/
func (a *MarketDataAPIService) QueryMarginAvailableInventory(ctx context.Context) ApiQueryMarginAvailableInventoryRequest {
	return ApiQueryMarginAvailableInventoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginAvailableInventoryResponse
func (a *MarketDataAPIService) QueryMarginAvailableInventoryExecute(r ApiQueryMarginAvailableInventoryRequest) (*common.RestApiResponse[models.QueryMarginAvailableInventoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/available-inventory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")

	resp, err := SendRequest[models.QueryMarginAvailableInventoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginPriceindexRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	symbol     *string
}

func (r ApiQueryMarginPriceindexRequest) Symbol(symbol string) ApiQueryMarginPriceindexRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryMarginPriceindexRequest) Execute() (*common.RestApiResponse[models.QueryMarginPriceindexResponse], error) {
	return r.ApiService.QueryMarginPriceindexExecute(r)
}

/*
QueryMarginPriceindex Query Margin PriceIndex (MARKET_DATA)
Get /sapi/v1/margin/priceIndex

https://developers.binance.com/docs/margin_trading/market-data/Query-Margin-PriceIndex

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiQueryMarginPriceindexRequest
*/
func (a *MarketDataAPIService) QueryMarginPriceindex(ctx context.Context) ApiQueryMarginPriceindexRequest {
	return ApiQueryMarginPriceindexRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginPriceindexResponse
func (a *MarketDataAPIService) QueryMarginPriceindexExecute(r ApiQueryMarginPriceindexRequest) (*common.RestApiResponse[models.QueryMarginPriceindexResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/margin/priceIndex"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.QueryMarginPriceindexResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
