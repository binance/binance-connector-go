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
	"github.com/binance/binance-connector-go/common/common"
)

// PortfolioMarginEndpointsAPIService PortfolioMarginEndpointsAPI Service
type PortfolioMarginEndpointsAPIService Service

type ApiClassicPortfolioMarginAccountInformationRequest struct {
	ctx        context.Context
	ApiService *PortfolioMarginEndpointsAPIService
	asset      *string
	recvWindow *int64
}

func (r ApiClassicPortfolioMarginAccountInformationRequest) Asset(asset string) ApiClassicPortfolioMarginAccountInformationRequest {
	r.asset = &asset
	return r
}

func (r ApiClassicPortfolioMarginAccountInformationRequest) RecvWindow(recvWindow int64) ApiClassicPortfolioMarginAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiClassicPortfolioMarginAccountInformationRequest) Execute() (*common.RestApiResponse[models.ClassicPortfolioMarginAccountInformationResponse], error) {
	return r.ApiService.ClassicPortfolioMarginAccountInformationExecute(r)
}

/*
ClassicPortfolioMarginAccountInformation Classic Portfolio Margin Account Information (USER_DATA)
Get /fapi/v1/pmAccountInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/portfolio-margin-endpoints/Classic-Portfolio-Margin-Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param recvWindow -
@return ApiClassicPortfolioMarginAccountInformationRequest
*/
func (a *PortfolioMarginEndpointsAPIService) ClassicPortfolioMarginAccountInformation(ctx context.Context) ApiClassicPortfolioMarginAccountInformationRequest {
	return ApiClassicPortfolioMarginAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClassicPortfolioMarginAccountInformationResponse
func (a *PortfolioMarginEndpointsAPIService) ClassicPortfolioMarginAccountInformationExecute(r ApiClassicPortfolioMarginAccountInformationRequest) (*common.RestApiResponse[models.ClassicPortfolioMarginAccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/pmAccountInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ClassicPortfolioMarginAccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
