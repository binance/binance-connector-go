/*
Binance Rebate REST API

OpenAPI Specification for the Binance Rebate REST API

API version: 1.0.0
*/

package binancerebaterestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/rebate/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// RebateAPIService RebateAPI Service
type RebateAPIService Service

type ApiGetSpotRebateHistoryRecordsRequest struct {
	ctx        context.Context
	ApiService *RebateAPIService
	startTime  *int64
	endTime    *int64
	page       *int64
	recvWindow *int64
}

func (r ApiGetSpotRebateHistoryRecordsRequest) StartTime(startTime int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSpotRebateHistoryRecordsRequest) EndTime(endTime int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.endTime = &endTime
	return r
}

// Default 1
func (r ApiGetSpotRebateHistoryRecordsRequest) Page(page int64) ApiGetSpotRebateHistoryRecordsRequest {
	r.page = &page
	return r
}

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

https://developers.binance.com/docs/rebate/rest-api/Get-Spot-Rebate-History-Records

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param page -  Default 1
@param recvWindow -
@return ApiGetSpotRebateHistoryRecordsRequest
*/
func (a *RebateAPIService) GetSpotRebateHistoryRecords(ctx context.Context) ApiGetSpotRebateHistoryRecordsRequest {
	return ApiGetSpotRebateHistoryRecordsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSpotRebateHistoryRecordsResponse
func (a *RebateAPIService) GetSpotRebateHistoryRecordsExecute(r ApiGetSpotRebateHistoryRecordsRequest) (*common.RestApiResponse[models.GetSpotRebateHistoryRecordsResponse], error) {
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

	resp, err := SendRequest[models.GetSpotRebateHistoryRecordsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
