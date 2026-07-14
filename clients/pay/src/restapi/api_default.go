/*
Binance Pay REST API

Query Binance Pay transaction history.
*/

package binancepayrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/pay/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiGetPayTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *DefaultAPIService
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Start time in milliseconds.
func (r ApiGetPayTradeHistoryRequest) StartTime(startTime int64) ApiGetPayTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

// End time in milliseconds.
func (r ApiGetPayTradeHistoryRequest) EndTime(endTime int64) ApiGetPayTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

// Number of records to return.
func (r ApiGetPayTradeHistoryRequest) Limit(limit int64) ApiGetPayTradeHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds.
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

https://developers.binance.com/en/docs/catalog/investment-and-services-pay/api/rest-api/~#get-pay-trade-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time in milliseconds.
@param endTime -  End time in milliseconds.
@param limit -  Number of records to return.
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetPayTradeHistoryRequest
*/
func (a *DefaultAPIService) GetPayTradeHistory(ctx context.Context) ApiGetPayTradeHistoryRequest {
	return ApiGetPayTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPayTradeHistoryResponse
func (a *DefaultAPIService) GetPayTradeHistoryExecute(r ApiGetPayTradeHistoryRequest) (*common.RestApiResponse[models.GetPayTradeHistoryResponse], error) {
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

	resp, err := SendRequest[models.GetPayTradeHistoryResponse](
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
