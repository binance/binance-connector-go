/*
Rebate REST API

Query spot trading rebate history records.
*/

package binancerebaterestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/rebate/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiGetSpotRebateHistoryRecordsRequest struct {
	ctx        context.Context
	ApiService *DefaultAPIService
	startTime  *int64
	endTime    *int64
	page       *int64
	recvWindow *int64
}

// Start time in milliseconds.
func (r ApiGetSpotRebateHistoryRecordsRequest) StartTime(startTime int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.startTime = &startTime
	return r
}

// End time in milliseconds.
func (r ApiGetSpotRebateHistoryRecordsRequest) EndTime(endTime int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.endTime = &endTime
	return r
}

// Page number.
func (r ApiGetSpotRebateHistoryRecordsRequest) Page(page int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.page = &page
	return r
}

// Request validity window in milliseconds.
func (r ApiGetSpotRebateHistoryRecordsRequest) RecvWindow(recvWindow int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSpotRebateHistoryRecordsRequest) Execute() (*common.RestApiResponse[models.GetSpotRebateHistoryRecordsResponse], error) {
	return r.ApiService.GetSpotRebateHistoryRecordsExecute(r)
}

/*
GetSpotRebateHistoryRecords Get Spot Rebate History Records (USER_DATA)
Get /sapi/v1/rebate/taxQuery

https://developers.binance.com/en/docs/catalog/investment-and-services-rebate/api/rest-api/~#get-spot-rebate-history-records

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -  Start time in milliseconds.
@param endTime -  End time in milliseconds.
@param page -  Page number.
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetSpotRebateHistoryRecordsRequest
*/
func (a *DefaultAPIService) GetSpotRebateHistoryRecords(ctx context.Context) ApiGetSpotRebateHistoryRecordsRequest {
	return ApiGetSpotRebateHistoryRecordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSpotRebateHistoryRecordsResponse
func (a *DefaultAPIService) GetSpotRebateHistoryRecordsExecute(r ApiGetSpotRebateHistoryRecordsRequest) (*common.RestApiResponse[models.GetSpotRebateHistoryRecordsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rebate/taxQuery"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSpotRebateHistoryRecordsResponse](
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
