/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API
*/

package binancepayrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/pay/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// PayAPIService PayAPI Service
type PayAPIService Service

type ApiGetPayTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *PayAPIService
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetPayTradeHistoryRequest) StartTime(startTime int64) ApiGetPayTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetPayTradeHistoryRequest) EndTime(endTime int64) ApiGetPayTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

// default 100, max 100
func (r ApiGetPayTradeHistoryRequest) Limit(limit int64) ApiGetPayTradeHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetPayTradeHistoryRequest) RecvWindow(recvWindow int64) ApiGetPayTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPayTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetPayTradeHistoryResponse], error) {
	return r.ApiService.GetPayTradeHistoryExecute(r)
}

/*
GetPayTradeHistory Get Pay Trade History
Get /sapi/v1/pay/transactions

https://developers.binance.com/docs/pay/rest-api/Get-Pay-Trade-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param limit -  default 100, max 100
@param recvWindow -
@return ApiGetPayTradeHistoryRequest
*/
func (a *PayAPIService) GetPayTradeHistory(ctx context.Context) ApiGetPayTradeHistoryRequest {
	return ApiGetPayTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPayTradeHistoryResponse
func (a *PayAPIService) GetPayTradeHistoryExecute(r ApiGetPayTradeHistoryRequest) (*common.RestApiResponse[models.GetPayTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/pay/transactions"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPayTradeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
