/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API
*/

package binancederivativestradingportfoliomarginprorestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiGetPortfolioMarginAssetLeverageRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiGetPortfolioMarginAssetLeverageRequest) Execute() (*common.RestApiResponse[models.GetPortfolioMarginAssetLeverageResponse], error) {
	return r.ApiService.GetPortfolioMarginAssetLeverageExecute(r)
}

/*
GetPortfolioMarginAssetLeverage Get Portfolio Margin Asset Leverage(USER_DATA)
Get /sapi/v1/portfolio/margin-asset-leverage

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/market-data/Get-Portfolio-Margin-Asset-Leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiGetPortfolioMarginAssetLeverageRequest
*/
func (a *MarketDataAPIService) GetPortfolioMarginAssetLeverage(ctx context.Context) ApiGetPortfolioMarginAssetLeverageRequest {
	return ApiGetPortfolioMarginAssetLeverageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPortfolioMarginAssetLeverageResponse
func (a *MarketDataAPIService) GetPortfolioMarginAssetLeverageExecute(r ApiGetPortfolioMarginAssetLeverageRequest) (*common.RestApiResponse[models.GetPortfolioMarginAssetLeverageResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/margin-asset-leverage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetPortfolioMarginAssetLeverageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPortfolioMarginCollateralRateRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiPortfolioMarginCollateralRateRequest) Execute() (*common.RestApiResponse[models.PortfolioMarginCollateralRateResponse], error) {
	return r.ApiService.PortfolioMarginCollateralRateExecute(r)
}

/*
PortfolioMarginCollateralRate Portfolio Margin Collateral Rate(MARKET_DATA)
Get /sapi/v1/portfolio/collateralRate

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/market-data/Classic-Portfolio-Margin-Collateral-Rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiPortfolioMarginCollateralRateRequest
*/
func (a *MarketDataAPIService) PortfolioMarginCollateralRate(ctx context.Context) ApiPortfolioMarginCollateralRateRequest {
	return ApiPortfolioMarginCollateralRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PortfolioMarginCollateralRateResponse
func (a *MarketDataAPIService) PortfolioMarginCollateralRateExecute(r ApiPortfolioMarginCollateralRateRequest) (*common.RestApiResponse[models.PortfolioMarginCollateralRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/collateralRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.PortfolioMarginCollateralRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPortfolioMarginProTieredCollateralRateRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	recvWindow *int64
}

func (r ApiPortfolioMarginProTieredCollateralRateRequest) RecvWindow(recvWindow int64) ApiPortfolioMarginProTieredCollateralRateRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPortfolioMarginProTieredCollateralRateRequest) Execute() (*common.RestApiResponse[models.PortfolioMarginProTieredCollateralRateResponse], error) {
	return r.ApiService.PortfolioMarginProTieredCollateralRateExecute(r)
}

/*
PortfolioMarginProTieredCollateralRate Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)
Get /sapi/v2/portfolio/collateralRate

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/market-data/Portfolio-Margin-Pro-Tiered-Collateral-Rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiPortfolioMarginProTieredCollateralRateRequest
*/
func (a *MarketDataAPIService) PortfolioMarginProTieredCollateralRate(ctx context.Context) ApiPortfolioMarginProTieredCollateralRateRequest {
	return ApiPortfolioMarginProTieredCollateralRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PortfolioMarginProTieredCollateralRateResponse
func (a *MarketDataAPIService) PortfolioMarginProTieredCollateralRateExecute(r ApiPortfolioMarginProTieredCollateralRateRequest) (*common.RestApiResponse[models.PortfolioMarginProTieredCollateralRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/portfolio/collateralRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PortfolioMarginProTieredCollateralRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryPortfolioMarginAssetIndexPriceRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	asset      *string
}

func (r ApiQueryPortfolioMarginAssetIndexPriceRequest) Asset(asset string) ApiQueryPortfolioMarginAssetIndexPriceRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryPortfolioMarginAssetIndexPriceRequest) Execute() (*common.RestApiResponse[models.QueryPortfolioMarginAssetIndexPriceResponse], error) {
	return r.ApiService.QueryPortfolioMarginAssetIndexPriceExecute(r)
}

/*
QueryPortfolioMarginAssetIndexPrice Query Portfolio Margin Asset Index Price (MARKET_DATA)
Get /sapi/v1/portfolio/asset-index-price

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/market-data/Query-Portfolio-Margin-Asset-Index-Price

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@return ApiQueryPortfolioMarginAssetIndexPriceRequest
*/
func (a *MarketDataAPIService) QueryPortfolioMarginAssetIndexPrice(ctx context.Context) ApiQueryPortfolioMarginAssetIndexPriceRequest {
	return ApiQueryPortfolioMarginAssetIndexPriceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPortfolioMarginAssetIndexPriceResponse
func (a *MarketDataAPIService) QueryPortfolioMarginAssetIndexPriceExecute(r ApiQueryPortfolioMarginAssetIndexPriceRequest) (*common.RestApiResponse[models.QueryPortfolioMarginAssetIndexPriceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/asset-index-price"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}

	resp, err := SendRequest[models.QueryPortfolioMarginAssetIndexPriceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
