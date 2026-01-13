/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package binancederivativestradingportfoliomarginrestapi

import (
	"context"
	"net/http"
	"net/url"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiTestConnectivityRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
}

func (r ApiTestConnectivityRequest) Execute() (struct{}, error) {
	return r.ApiService.TestConnectivityExecute(r)
}

/*
TestConnectivity Test Connectivity
Get /papi/v1/ping

https://developers.binance.com/docs/derivatives/portfolio-margin/market-data/Test-Connectivity

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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/ping"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	_, err := SendRequest[struct{}](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil {
		return struct{}{}, err
	}

	return struct{}{}, nil
}
