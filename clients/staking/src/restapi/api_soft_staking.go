/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package binancestakingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// SoftStakingAPIService SoftStakingAPI Service
type SoftStakingAPIService Service

type ApiGetSoftStakingProductListRequest struct {
	ctx        context.Context
	ApiService *SoftStakingAPIService
	asset      *string
	current    *int64
	size       *int64
	recvWindow *int64
}

// WBETH or BETH, default to BETH
func (r ApiGetSoftStakingProductListRequest) Asset(asset string) ApiGetSoftStakingProductListRequest {
	r.asset = &asset
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetSoftStakingProductListRequest) Current(current int64) ApiGetSoftStakingProductListRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetSoftStakingProductListRequest) Size(size int64) ApiGetSoftStakingProductListRequest {
	r.size = &size
	return r
}

func (r ApiGetSoftStakingProductListRequest) RecvWindow(recvWindow int64) ApiGetSoftStakingProductListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSoftStakingProductListRequest) Execute() (*common.RestApiResponse[models.GetSoftStakingProductListResponse], error) {
	return r.ApiService.GetSoftStakingProductListExecute(r)
}

/*
GetSoftStakingProductList Get Soft Staking Product List (USER_DATA)
Get /sapi/v1/soft-staking/list

https://developers.binance.com/docs/staking/soft-staking/

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  WBETH or BETH, default to BETH
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetSoftStakingProductListRequest
*/
func (a *SoftStakingAPIService) GetSoftStakingProductList(ctx context.Context) ApiGetSoftStakingProductListRequest {
	return ApiGetSoftStakingProductListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSoftStakingProductListResponse
func (a *SoftStakingAPIService) GetSoftStakingProductListExecute(r ApiGetSoftStakingProductListRequest) (*common.RestApiResponse[models.GetSoftStakingProductListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/soft-staking/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSoftStakingProductListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSoftStakingRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *SoftStakingAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// WBETH or BETH, default to BETH
func (r ApiGetSoftStakingRewardsHistoryRequest) Asset(asset string) ApiGetSoftStakingRewardsHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetSoftStakingRewardsHistoryRequest) StartTime(startTime int64) ApiGetSoftStakingRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSoftStakingRewardsHistoryRequest) EndTime(endTime int64) ApiGetSoftStakingRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetSoftStakingRewardsHistoryRequest) Current(current int64) ApiGetSoftStakingRewardsHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetSoftStakingRewardsHistoryRequest) Size(size int64) ApiGetSoftStakingRewardsHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetSoftStakingRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetSoftStakingRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSoftStakingRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetSoftStakingRewardsHistoryResponse], error) {
	return r.ApiService.GetSoftStakingRewardsHistoryExecute(r)
}

/*
GetSoftStakingRewardsHistory Get Soft Staking Rewards History(USER_DATA)
Get /sapi/v1/soft-staking/history/rewardsRecord

https://developers.binance.com/docs/staking/soft-staking/Get-Soft-Staking-Rewards-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  WBETH or BETH, default to BETH
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetSoftStakingRewardsHistoryRequest
*/
func (a *SoftStakingAPIService) GetSoftStakingRewardsHistory(ctx context.Context) ApiGetSoftStakingRewardsHistoryRequest {
	return ApiGetSoftStakingRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSoftStakingRewardsHistoryResponse
func (a *SoftStakingAPIService) GetSoftStakingRewardsHistoryExecute(r ApiGetSoftStakingRewardsHistoryRequest) (*common.RestApiResponse[models.GetSoftStakingRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/soft-staking/history/rewardsRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSoftStakingRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetSoftStakingRequest struct {
	ctx         context.Context
	ApiService  *SoftStakingAPIService
	softStaking *bool
	recvWindow  *int64
}

// true or false
func (r ApiSetSoftStakingRequest) SoftStaking(softStaking bool) ApiSetSoftStakingRequest {
	r.softStaking = &softStaking
	return r
}

func (r ApiSetSoftStakingRequest) RecvWindow(recvWindow int64) ApiSetSoftStakingRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetSoftStakingRequest) Execute() (*common.RestApiResponse[models.SetSoftStakingResponse], error) {
	return r.ApiService.SetSoftStakingExecute(r)
}

/*
SetSoftStaking Set Soft Staking (USER_DATA)
Get /sapi/v1/soft-staking/set

https://developers.binance.com/docs/staking/soft-staking/Set-Soft-Staking

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param softStaking -  true or false
@param recvWindow -
@return ApiSetSoftStakingRequest
*/
func (a *SoftStakingAPIService) SetSoftStaking(ctx context.Context) ApiSetSoftStakingRequest {
	return ApiSetSoftStakingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetSoftStakingResponse
func (a *SoftStakingAPIService) SetSoftStakingExecute(r ApiSetSoftStakingRequest) (*common.RestApiResponse[models.SetSoftStakingResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/soft-staking/set"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.softStaking == nil {
		return nil, common.ReportError("softStaking is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "softStaking", r.softStaking, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetSoftStakingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
