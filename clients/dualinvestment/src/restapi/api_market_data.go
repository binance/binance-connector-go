/*
Dual Investment REST API

Query products, request quotes, and subscribe to Advanced Earn Dual Investment strategies.
*/

package binancedualinvestmentrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/dualinvestment/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiGetDualInvestmentProductListRequest struct {
	ctx           context.Context
	ApiService    *MarketDataAPIService
	optionType    *models.GetDualInvestmentProductListOptionTypeParameter
	exercisedCoin *string
	investCoin    *string
	pageSize      *int64
	pageIndex     *int64
	recvWindow    *int64
}

// Input CALL or PUT
func (r ApiGetDualInvestmentProductListRequest) OptionType(optionType models.GetDualInvestmentProductListOptionTypeParameter) ApiGetDualInvestmentProductListRequest {
	r.optionType = &optionType
	return r
}

// Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType: CALL&#x60;, &#x60;exercisedCoin: USDT&#x60;, &#x60;investCoin: BNB&#x60;; if you subscribe to a low buy product (put option), you should input: &#x60;optionType: PUT&#x60;, &#x60;exercisedCoin: BNB&#x60;, &#x60;investCoin: USDT&#x60;
func (r ApiGetDualInvestmentProductListRequest) ExercisedCoin(exercisedCoin string) ApiGetDualInvestmentProductListRequest {
	r.exercisedCoin = &exercisedCoin
	return r
}

// Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType: CALL&#x60;, &#x60;exercisedCoin: USDT&#x60;, &#x60;investCoin: BNB&#x60;; if you subscribe to a low buy product (put option), you should input: &#x60;optionType: PUT&#x60;, &#x60;exercisedCoin: BNB&#x60;, &#x60;investCoin: USDT&#x60;
func (r ApiGetDualInvestmentProductListRequest) InvestCoin(investCoin string) ApiGetDualInvestmentProductListRequest {
	r.investCoin = &investCoin
	return r
}

// Number of records per page
func (r ApiGetDualInvestmentProductListRequest) PageSize(pageSize int64) ApiGetDualInvestmentProductListRequest {
	r.pageSize = &pageSize
	return r
}

// Page index
func (r ApiGetDualInvestmentProductListRequest) PageIndex(pageIndex int64) ApiGetDualInvestmentProductListRequest {
	r.pageIndex = &pageIndex
	return r
}

// Request validity window in milliseconds
func (r ApiGetDualInvestmentProductListRequest) RecvWindow(recvWindow int64) ApiGetDualInvestmentProductListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDualInvestmentProductListRequest) Execute() (*common.RestApiResponse[models.GetDualInvestmentProductListResponse], error) {
	return r.ApiService.GetDualInvestmentProductListExecute(r)
}

/*
GetDualInvestmentProductList Get Dual Investment product list
Get /sapi/v1/dci/product/list

https://developers.binance.com/en/docs/catalog/investment-and-services-dual-investment/api/rest-api/market-data#get-dual-investment-product-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param optionType -  Input CALL or PUT
@param exercisedCoin -  Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType: CALL`, `exercisedCoin: USDT`, `investCoin: BNB`; if you subscribe to a low buy product (put option), you should input: `optionType: PUT`, `exercisedCoin: BNB`, `investCoin: USDT`
@param investCoin -  Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType: CALL`, `exercisedCoin: USDT`, `investCoin: BNB`; if you subscribe to a low buy product (put option), you should input: `optionType: PUT`, `exercisedCoin: BNB`, `investCoin: USDT`
@param pageSize -  Number of records per page
@param pageIndex -  Page index
@param recvWindow -  Request validity window in milliseconds
@return ApiGetDualInvestmentProductListRequest
*/
func (a *MarketDataAPIService) GetDualInvestmentProductList(ctx context.Context) ApiGetDualInvestmentProductListRequest {
	return ApiGetDualInvestmentProductListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDualInvestmentProductListResponse
func (a *MarketDataAPIService) GetDualInvestmentProductListExecute(r ApiGetDualInvestmentProductListRequest) (*common.RestApiResponse[models.GetDualInvestmentProductListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/dci/product/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.optionType == nil {
		return nil, common.ReportError("optionType is required and must be specified")
	}

	if r.exercisedCoin == nil {
		return nil, common.ReportError("exercisedCoin is required and must be specified")
	}

	if r.investCoin == nil {
		return nil, common.ReportError("investCoin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "optionType", r.optionType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "exercisedCoin", r.exercisedCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "investCoin", r.investCoin, "form", "")
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDualInvestmentProductListResponse](
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
