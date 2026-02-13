/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API
*/

package binanceconvertrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/convert/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiListAllConvertPairsRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	fromAsset  *string
	toAsset    *string
}

// User spends coin
func (r ApiListAllConvertPairsRequest) FromAsset(fromAsset string) ApiListAllConvertPairsRequest {
	r.fromAsset = &fromAsset
	return r
}

// User receives coin
func (r ApiListAllConvertPairsRequest) ToAsset(toAsset string) ApiListAllConvertPairsRequest {
	r.toAsset = &toAsset
	return r
}

func (r ApiListAllConvertPairsRequest) Execute() (*common.RestApiResponse[models.ListAllConvertPairsResponse], error) {
	return r.ApiService.ListAllConvertPairsExecute(r)
}

/*
ListAllConvertPairs List All Convert Pairs
Get /sapi/v1/convert/exchangeInfo

https://developers.binance.com/docs/convert/market-data/

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAsset -  User spends coin
@param toAsset -  User receives coin
@return ApiListAllConvertPairsRequest
*/
func (a *MarketDataAPIService) ListAllConvertPairs(ctx context.Context) ApiListAllConvertPairsRequest {
	return ApiListAllConvertPairsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListAllConvertPairsResponse
func (a *MarketDataAPIService) ListAllConvertPairsExecute(r ApiListAllConvertPairsRequest) (*common.RestApiResponse[models.ListAllConvertPairsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAsset", r.fromAsset, "form", "")
	}
	if r.toAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAsset", r.toAsset, "form", "")
	}

	resp, err := SendRequest[models.ListAllConvertPairsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryOrderQuantityPrecisionPerAssetRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000
func (r ApiQueryOrderQuantityPrecisionPerAssetRequest) RecvWindow(recvWindow int64) ApiQueryOrderQuantityPrecisionPerAssetRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryOrderQuantityPrecisionPerAssetRequest) Execute() (*common.RestApiResponse[models.QueryOrderQuantityPrecisionPerAssetResponse], error) {
	return r.ApiService.QueryOrderQuantityPrecisionPerAssetExecute(r)
}

/*
QueryOrderQuantityPrecisionPerAsset Query order quantity precision per asset(USER_DATA)
Get /sapi/v1/convert/assetInfo

https://developers.binance.com/docs/convert/market-data/Query-order-quantity-precision-per-asset

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000
@return ApiQueryOrderQuantityPrecisionPerAssetRequest
*/
func (a *MarketDataAPIService) QueryOrderQuantityPrecisionPerAsset(ctx context.Context) ApiQueryOrderQuantityPrecisionPerAssetRequest {
	return ApiQueryOrderQuantityPrecisionPerAssetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryOrderQuantityPrecisionPerAssetResponse
func (a *MarketDataAPIService) QueryOrderQuantityPrecisionPerAssetExecute(r ApiQueryOrderQuantityPrecisionPerAssetRequest) (*common.RestApiResponse[models.QueryOrderQuantityPrecisionPerAssetResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/convert/assetInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryOrderQuantityPrecisionPerAssetResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
