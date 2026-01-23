/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package binancesimpleearnrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// FlexibleLockedAPIService FlexibleLockedAPI Service
type FlexibleLockedAPIService Service

type ApiGetCollateralRecordRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetCollateralRecordRequest) ProductId(productId string) ApiGetCollateralRecordRequest {
	r.productId = &productId
	return r
}

func (r ApiGetCollateralRecordRequest) StartTime(startTime int64) ApiGetCollateralRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetCollateralRecordRequest) EndTime(endTime int64) ApiGetCollateralRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetCollateralRecordRequest) Current(current int64) ApiGetCollateralRecordRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetCollateralRecordRequest) Size(size int64) ApiGetCollateralRecordRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetCollateralRecordRequest) RecvWindow(recvWindow int64) ApiGetCollateralRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCollateralRecordRequest) Execute() (*common.RestApiResponse[models.GetCollateralRecordResponse], error) {
	return r.ApiService.GetCollateralRecordExecute(r)
}

/*
GetCollateralRecord Get Collateral Record(USER_DATA)
Get /sapi/v1/simple-earn/flexible/history/collateralRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Collateral-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetCollateralRecordRequest
*/
func (a *FlexibleLockedAPIService) GetCollateralRecord(ctx context.Context) ApiGetCollateralRecordRequest {
	return ApiGetCollateralRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCollateralRecordResponse
func (a *FlexibleLockedAPIService) GetCollateralRecordExecute(r ApiGetCollateralRecordRequest) (*common.RestApiResponse[models.GetCollateralRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/history/collateralRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
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

	resp, err := SendRequest[models.GetCollateralRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexiblePersonalLeftQuotaRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	recvWindow *int64
}

func (r ApiGetFlexiblePersonalLeftQuotaRequest) ProductId(productId string) ApiGetFlexiblePersonalLeftQuotaRequest {
	r.productId = &productId
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexiblePersonalLeftQuotaRequest) RecvWindow(recvWindow int64) ApiGetFlexiblePersonalLeftQuotaRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexiblePersonalLeftQuotaRequest) Execute() (*common.RestApiResponse[models.GetFlexiblePersonalLeftQuotaResponse], error) {
	return r.ApiService.GetFlexiblePersonalLeftQuotaExecute(r)
}

/*
GetFlexiblePersonalLeftQuota Get Flexible Personal Left Quota(USER_DATA)
Get /sapi/v1/simple-earn/flexible/personalLeftQuota

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Flexible-Personal-Left-Quota

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexiblePersonalLeftQuotaRequest
*/
func (a *FlexibleLockedAPIService) GetFlexiblePersonalLeftQuota(ctx context.Context) ApiGetFlexiblePersonalLeftQuotaRequest {
	return ApiGetFlexiblePersonalLeftQuotaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexiblePersonalLeftQuotaResponse
func (a *FlexibleLockedAPIService) GetFlexiblePersonalLeftQuotaExecute(r ApiGetFlexiblePersonalLeftQuotaRequest) (*common.RestApiResponse[models.GetFlexiblePersonalLeftQuotaResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/personalLeftQuota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFlexiblePersonalLeftQuotaResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleProductPositionRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	asset      *string
	productId  *string
	current    *int64
	size       *int64
	recvWindow *int64
}

// USDC or USDT
func (r ApiGetFlexibleProductPositionRequest) Asset(asset string) ApiGetFlexibleProductPositionRequest {
	r.asset = &asset
	return r
}

func (r ApiGetFlexibleProductPositionRequest) ProductId(productId string) ApiGetFlexibleProductPositionRequest {
	r.productId = &productId
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetFlexibleProductPositionRequest) Current(current int64) ApiGetFlexibleProductPositionRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetFlexibleProductPositionRequest) Size(size int64) ApiGetFlexibleProductPositionRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexibleProductPositionRequest) RecvWindow(recvWindow int64) ApiGetFlexibleProductPositionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleProductPositionRequest) Execute() (*common.RestApiResponse[models.GetFlexibleProductPositionResponse], error) {
	return r.ApiService.GetFlexibleProductPositionExecute(r)
}

/*
GetFlexibleProductPosition Get Flexible Product Position(USER_DATA)
Get /sapi/v1/simple-earn/flexible/position

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Flexible-Product-Position

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDC or USDT
@param productId -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexibleProductPositionRequest
*/
func (a *FlexibleLockedAPIService) GetFlexibleProductPosition(ctx context.Context) ApiGetFlexibleProductPositionRequest {
	return ApiGetFlexibleProductPositionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleProductPositionResponse
func (a *FlexibleLockedAPIService) GetFlexibleProductPositionExecute(r ApiGetFlexibleProductPositionRequest) (*common.RestApiResponse[models.GetFlexibleProductPositionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/position"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.productId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
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

	resp, err := SendRequest[models.GetFlexibleProductPositionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleRedemptionRecordRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	redeemId   *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetFlexibleRedemptionRecordRequest) ProductId(productId string) ApiGetFlexibleRedemptionRecordRequest {
	r.productId = &productId
	return r
}

func (r ApiGetFlexibleRedemptionRecordRequest) RedeemId(redeemId string) ApiGetFlexibleRedemptionRecordRequest {
	r.redeemId = &redeemId
	return r
}

// USDC or USDT
func (r ApiGetFlexibleRedemptionRecordRequest) Asset(asset string) ApiGetFlexibleRedemptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetFlexibleRedemptionRecordRequest) StartTime(startTime int64) ApiGetFlexibleRedemptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleRedemptionRecordRequest) EndTime(endTime int64) ApiGetFlexibleRedemptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetFlexibleRedemptionRecordRequest) Current(current int64) ApiGetFlexibleRedemptionRecordRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetFlexibleRedemptionRecordRequest) Size(size int64) ApiGetFlexibleRedemptionRecordRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexibleRedemptionRecordRequest) RecvWindow(recvWindow int64) ApiGetFlexibleRedemptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleRedemptionRecordRequest) Execute() (*common.RestApiResponse[models.GetFlexibleRedemptionRecordResponse], error) {
	return r.ApiService.GetFlexibleRedemptionRecordExecute(r)
}

/*
GetFlexibleRedemptionRecord Get Flexible Redemption Record(USER_DATA)
Get /sapi/v1/simple-earn/flexible/history/redemptionRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Flexible-Redemption-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param redeemId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexibleRedemptionRecordRequest
*/
func (a *FlexibleLockedAPIService) GetFlexibleRedemptionRecord(ctx context.Context) ApiGetFlexibleRedemptionRecordRequest {
	return ApiGetFlexibleRedemptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleRedemptionRecordResponse
func (a *FlexibleLockedAPIService) GetFlexibleRedemptionRecordExecute(r ApiGetFlexibleRedemptionRecordRequest) (*common.RestApiResponse[models.GetFlexibleRedemptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/history/redemptionRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	}
	if r.redeemId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "redeemId", r.redeemId, "form", "")
	}
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

	resp, err := SendRequest[models.GetFlexibleRedemptionRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	type_      *string
	productId  *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// FAST or STANDARD, defaults to STANDARD
func (r ApiGetFlexibleRewardsHistoryRequest) Type(type_ string) ApiGetFlexibleRewardsHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetFlexibleRewardsHistoryRequest) ProductId(productId string) ApiGetFlexibleRewardsHistoryRequest {
	r.productId = &productId
	return r
}

// USDC or USDT
func (r ApiGetFlexibleRewardsHistoryRequest) Asset(asset string) ApiGetFlexibleRewardsHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetFlexibleRewardsHistoryRequest) StartTime(startTime int64) ApiGetFlexibleRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleRewardsHistoryRequest) EndTime(endTime int64) ApiGetFlexibleRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetFlexibleRewardsHistoryRequest) Current(current int64) ApiGetFlexibleRewardsHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetFlexibleRewardsHistoryRequest) Size(size int64) ApiGetFlexibleRewardsHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexibleRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleRewardsHistoryResponse], error) {
	return r.ApiService.GetFlexibleRewardsHistoryExecute(r)
}

/*
GetFlexibleRewardsHistory Get Flexible Rewards History(USER_DATA)
Get /sapi/v1/simple-earn/flexible/history/rewardsRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Flexible-Rewards-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -  FAST or STANDARD, defaults to STANDARD
@param productId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexibleRewardsHistoryRequest
*/
func (a *FlexibleLockedAPIService) GetFlexibleRewardsHistory(ctx context.Context) ApiGetFlexibleRewardsHistoryRequest {
	return ApiGetFlexibleRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleRewardsHistoryResponse
func (a *FlexibleLockedAPIService) GetFlexibleRewardsHistoryExecute(r ApiGetFlexibleRewardsHistoryRequest) (*common.RestApiResponse[models.GetFlexibleRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/history/rewardsRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	if r.productId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	}
	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFlexibleRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleSubscriptionPreviewRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	amount     *float32
	recvWindow *int64
}

func (r ApiGetFlexibleSubscriptionPreviewRequest) ProductId(productId string) ApiGetFlexibleSubscriptionPreviewRequest {
	r.productId = &productId
	return r
}

// Amount
func (r ApiGetFlexibleSubscriptionPreviewRequest) Amount(amount float32) ApiGetFlexibleSubscriptionPreviewRequest {
	r.amount = &amount
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexibleSubscriptionPreviewRequest) RecvWindow(recvWindow int64) ApiGetFlexibleSubscriptionPreviewRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleSubscriptionPreviewRequest) Execute() (*common.RestApiResponse[models.GetFlexibleSubscriptionPreviewResponse], error) {
	return r.ApiService.GetFlexibleSubscriptionPreviewExecute(r)
}

/*
GetFlexibleSubscriptionPreview Get Flexible Subscription Preview(USER_DATA)
Get /sapi/v1/simple-earn/flexible/subscriptionPreview

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Get-Flexible-Subscription-Preview

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param amount -  Amount
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexibleSubscriptionPreviewRequest
*/
func (a *FlexibleLockedAPIService) GetFlexibleSubscriptionPreview(ctx context.Context) ApiGetFlexibleSubscriptionPreviewRequest {
	return ApiGetFlexibleSubscriptionPreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleSubscriptionPreviewResponse
func (a *FlexibleLockedAPIService) GetFlexibleSubscriptionPreviewExecute(r ApiGetFlexibleSubscriptionPreviewRequest) (*common.RestApiResponse[models.GetFlexibleSubscriptionPreviewResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/subscriptionPreview"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFlexibleSubscriptionPreviewResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleSubscriptionRecordRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	purchaseId *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetFlexibleSubscriptionRecordRequest) ProductId(productId string) ApiGetFlexibleSubscriptionRecordRequest {
	r.productId = &productId
	return r
}

func (r ApiGetFlexibleSubscriptionRecordRequest) PurchaseId(purchaseId string) ApiGetFlexibleSubscriptionRecordRequest {
	r.purchaseId = &purchaseId
	return r
}

// USDC or USDT
func (r ApiGetFlexibleSubscriptionRecordRequest) Asset(asset string) ApiGetFlexibleSubscriptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetFlexibleSubscriptionRecordRequest) StartTime(startTime int64) ApiGetFlexibleSubscriptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleSubscriptionRecordRequest) EndTime(endTime int64) ApiGetFlexibleSubscriptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetFlexibleSubscriptionRecordRequest) Current(current int64) ApiGetFlexibleSubscriptionRecordRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetFlexibleSubscriptionRecordRequest) Size(size int64) ApiGetFlexibleSubscriptionRecordRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetFlexibleSubscriptionRecordRequest) RecvWindow(recvWindow int64) ApiGetFlexibleSubscriptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleSubscriptionRecordRequest) Execute() (*common.RestApiResponse[models.GetFlexibleSubscriptionRecordResponse], error) {
	return r.ApiService.GetFlexibleSubscriptionRecordExecute(r)
}

/*
GetFlexibleSubscriptionRecord Get Flexible Subscription Record(USER_DATA)
Get /sapi/v1/simple-earn/flexible/history/subscriptionRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Flexible-Subscription-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param purchaseId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetFlexibleSubscriptionRecordRequest
*/
func (a *FlexibleLockedAPIService) GetFlexibleSubscriptionRecord(ctx context.Context) ApiGetFlexibleSubscriptionRecordRequest {
	return ApiGetFlexibleSubscriptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleSubscriptionRecordResponse
func (a *FlexibleLockedAPIService) GetFlexibleSubscriptionRecordExecute(r ApiGetFlexibleSubscriptionRecordRequest) (*common.RestApiResponse[models.GetFlexibleSubscriptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/history/subscriptionRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	}
	if r.purchaseId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "purchaseId", r.purchaseId, "form", "")
	}
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

	resp, err := SendRequest[models.GetFlexibleSubscriptionRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedPersonalLeftQuotaRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	projectId  *string
	recvWindow *int64
}

func (r ApiGetLockedPersonalLeftQuotaRequest) ProjectId(projectId string) ApiGetLockedPersonalLeftQuotaRequest {
	r.projectId = &projectId
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedPersonalLeftQuotaRequest) RecvWindow(recvWindow int64) ApiGetLockedPersonalLeftQuotaRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedPersonalLeftQuotaRequest) Execute() (*common.RestApiResponse[models.GetLockedPersonalLeftQuotaResponse], error) {
	return r.ApiService.GetLockedPersonalLeftQuotaExecute(r)
}

/*
GetLockedPersonalLeftQuota Get Locked Personal Left Quota(USER_DATA)
Get /sapi/v1/simple-earn/locked/personalLeftQuota

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Locked-Personal-Left-Quota

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedPersonalLeftQuotaRequest
*/
func (a *FlexibleLockedAPIService) GetLockedPersonalLeftQuota(ctx context.Context) ApiGetLockedPersonalLeftQuotaRequest {
	return ApiGetLockedPersonalLeftQuotaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedPersonalLeftQuotaResponse
func (a *FlexibleLockedAPIService) GetLockedPersonalLeftQuotaExecute(r ApiGetLockedPersonalLeftQuotaRequest) (*common.RestApiResponse[models.GetLockedPersonalLeftQuotaResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/personalLeftQuota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.projectId == nil {
		return nil, common.ReportError("projectId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetLockedPersonalLeftQuotaResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedProductPositionRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	asset      *string
	positionId *int64
	projectId  *string
	current    *int64
	size       *int64
	recvWindow *int64
}

// USDC or USDT
func (r ApiGetLockedProductPositionRequest) Asset(asset string) ApiGetLockedProductPositionRequest {
	r.asset = &asset
	return r
}

func (r ApiGetLockedProductPositionRequest) PositionId(positionId int64) ApiGetLockedProductPositionRequest {
	r.positionId = &positionId
	return r
}

func (r ApiGetLockedProductPositionRequest) ProjectId(projectId string) ApiGetLockedProductPositionRequest {
	r.projectId = &projectId
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetLockedProductPositionRequest) Current(current int64) ApiGetLockedProductPositionRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetLockedProductPositionRequest) Size(size int64) ApiGetLockedProductPositionRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedProductPositionRequest) RecvWindow(recvWindow int64) ApiGetLockedProductPositionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedProductPositionRequest) Execute() (*common.RestApiResponse[models.GetLockedProductPositionResponse], error) {
	return r.ApiService.GetLockedProductPositionExecute(r)
}

/*
GetLockedProductPosition Get Locked Product Position
Get /sapi/v1/simple-earn/locked/position

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Locked-Product-Position

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDC or USDT
@param positionId -
@param projectId -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedProductPositionRequest
*/
func (a *FlexibleLockedAPIService) GetLockedProductPosition(ctx context.Context) ApiGetLockedProductPositionRequest {
	return ApiGetLockedProductPositionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedProductPositionResponse
func (a *FlexibleLockedAPIService) GetLockedProductPositionExecute(r ApiGetLockedProductPositionRequest) (*common.RestApiResponse[models.GetLockedProductPositionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/position"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.positionId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	}
	if r.projectId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "form", "")
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

	resp, err := SendRequest[models.GetLockedProductPositionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedRedemptionRecordRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	positionId *int64
	redeemId   *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetLockedRedemptionRecordRequest) PositionId(positionId int64) ApiGetLockedRedemptionRecordRequest {
	r.positionId = &positionId
	return r
}

func (r ApiGetLockedRedemptionRecordRequest) RedeemId(redeemId string) ApiGetLockedRedemptionRecordRequest {
	r.redeemId = &redeemId
	return r
}

// USDC or USDT
func (r ApiGetLockedRedemptionRecordRequest) Asset(asset string) ApiGetLockedRedemptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetLockedRedemptionRecordRequest) StartTime(startTime int64) ApiGetLockedRedemptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLockedRedemptionRecordRequest) EndTime(endTime int64) ApiGetLockedRedemptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetLockedRedemptionRecordRequest) Current(current int64) ApiGetLockedRedemptionRecordRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetLockedRedemptionRecordRequest) Size(size int64) ApiGetLockedRedemptionRecordRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedRedemptionRecordRequest) RecvWindow(recvWindow int64) ApiGetLockedRedemptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedRedemptionRecordRequest) Execute() (*common.RestApiResponse[models.GetLockedRedemptionRecordResponse], error) {
	return r.ApiService.GetLockedRedemptionRecordExecute(r)
}

/*
GetLockedRedemptionRecord Get Locked Redemption Record(USER_DATA)
Get /sapi/v1/simple-earn/locked/history/redemptionRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Locked-Redemption-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param redeemId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedRedemptionRecordRequest
*/
func (a *FlexibleLockedAPIService) GetLockedRedemptionRecord(ctx context.Context) ApiGetLockedRedemptionRecordRequest {
	return ApiGetLockedRedemptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedRedemptionRecordResponse
func (a *FlexibleLockedAPIService) GetLockedRedemptionRecordExecute(r ApiGetLockedRedemptionRecordRequest) (*common.RestApiResponse[models.GetLockedRedemptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/history/redemptionRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	}
	if r.redeemId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "redeemId", r.redeemId, "form", "")
	}
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

	resp, err := SendRequest[models.GetLockedRedemptionRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	positionId *int64
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetLockedRewardsHistoryRequest) PositionId(positionId int64) ApiGetLockedRewardsHistoryRequest {
	r.positionId = &positionId
	return r
}

// USDC or USDT
func (r ApiGetLockedRewardsHistoryRequest) Asset(asset string) ApiGetLockedRewardsHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetLockedRewardsHistoryRequest) StartTime(startTime int64) ApiGetLockedRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLockedRewardsHistoryRequest) EndTime(endTime int64) ApiGetLockedRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetLockedRewardsHistoryRequest) Current(current int64) ApiGetLockedRewardsHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetLockedRewardsHistoryRequest) Size(size int64) ApiGetLockedRewardsHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetLockedRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetLockedRewardsHistoryResponse], error) {
	return r.ApiService.GetLockedRewardsHistoryExecute(r)
}

/*
GetLockedRewardsHistory Get Locked Rewards History(USER_DATA)
Get /sapi/v1/simple-earn/locked/history/rewardsRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Locked-Rewards-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedRewardsHistoryRequest
*/
func (a *FlexibleLockedAPIService) GetLockedRewardsHistory(ctx context.Context) ApiGetLockedRewardsHistoryRequest {
	return ApiGetLockedRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedRewardsHistoryResponse
func (a *FlexibleLockedAPIService) GetLockedRewardsHistoryExecute(r ApiGetLockedRewardsHistoryRequest) (*common.RestApiResponse[models.GetLockedRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/history/rewardsRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	}
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

	resp, err := SendRequest[models.GetLockedRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedSubscriptionPreviewRequest struct {
	ctx           context.Context
	ApiService    *FlexibleLockedAPIService
	projectId     *string
	amount        *float32
	autoSubscribe *bool
	recvWindow    *int64
}

func (r ApiGetLockedSubscriptionPreviewRequest) ProjectId(projectId string) ApiGetLockedSubscriptionPreviewRequest {
	r.projectId = &projectId
	return r
}

// Amount
func (r ApiGetLockedSubscriptionPreviewRequest) Amount(amount float32) ApiGetLockedSubscriptionPreviewRequest {
	r.amount = &amount
	return r
}

// true or false, default true.
func (r ApiGetLockedSubscriptionPreviewRequest) AutoSubscribe(autoSubscribe bool) ApiGetLockedSubscriptionPreviewRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedSubscriptionPreviewRequest) RecvWindow(recvWindow int64) ApiGetLockedSubscriptionPreviewRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedSubscriptionPreviewRequest) Execute() (*common.RestApiResponse[models.GetLockedSubscriptionPreviewResponse], error) {
	return r.ApiService.GetLockedSubscriptionPreviewExecute(r)
}

/*
GetLockedSubscriptionPreview Get Locked Subscription Preview(USER_DATA)
Get /sapi/v1/simple-earn/locked/subscriptionPreview

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Get-Locked-Subscription-Preview

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param amount -  Amount
@param autoSubscribe -  true or false, default true.
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedSubscriptionPreviewRequest
*/
func (a *FlexibleLockedAPIService) GetLockedSubscriptionPreview(ctx context.Context) ApiGetLockedSubscriptionPreviewRequest {
	return ApiGetLockedSubscriptionPreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedSubscriptionPreviewResponse
func (a *FlexibleLockedAPIService) GetLockedSubscriptionPreviewExecute(r ApiGetLockedSubscriptionPreviewRequest) (*common.RestApiResponse[models.GetLockedSubscriptionPreviewResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/subscriptionPreview"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.projectId == nil {
		return nil, common.ReportError("projectId is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.autoSubscribe != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoSubscribe", r.autoSubscribe, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetLockedSubscriptionPreviewResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetLockedSubscriptionRecordRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	purchaseId *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetLockedSubscriptionRecordRequest) PurchaseId(purchaseId string) ApiGetLockedSubscriptionRecordRequest {
	r.purchaseId = &purchaseId
	return r
}

// USDC or USDT
func (r ApiGetLockedSubscriptionRecordRequest) Asset(asset string) ApiGetLockedSubscriptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetLockedSubscriptionRecordRequest) StartTime(startTime int64) ApiGetLockedSubscriptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLockedSubscriptionRecordRequest) EndTime(endTime int64) ApiGetLockedSubscriptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetLockedSubscriptionRecordRequest) Current(current int64) ApiGetLockedSubscriptionRecordRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetLockedSubscriptionRecordRequest) Size(size int64) ApiGetLockedSubscriptionRecordRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetLockedSubscriptionRecordRequest) RecvWindow(recvWindow int64) ApiGetLockedSubscriptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLockedSubscriptionRecordRequest) Execute() (*common.RestApiResponse[models.GetLockedSubscriptionRecordResponse], error) {
	return r.ApiService.GetLockedSubscriptionRecordExecute(r)
}

/*
GetLockedSubscriptionRecord Get Locked Subscription Record(USER_DATA)
Get /sapi/v1/simple-earn/locked/history/subscriptionRecord

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Locked-Subscription-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param purchaseId -
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetLockedSubscriptionRecordRequest
*/
func (a *FlexibleLockedAPIService) GetLockedSubscriptionRecord(ctx context.Context) ApiGetLockedSubscriptionRecordRequest {
	return ApiGetLockedSubscriptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLockedSubscriptionRecordResponse
func (a *FlexibleLockedAPIService) GetLockedSubscriptionRecordExecute(r ApiGetLockedSubscriptionRecordRequest) (*common.RestApiResponse[models.GetLockedSubscriptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/history/subscriptionRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.purchaseId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "purchaseId", r.purchaseId, "form", "")
	}
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

	resp, err := SendRequest[models.GetLockedSubscriptionRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRateHistoryRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	productId  *string
	aprPeriod  *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetRateHistoryRequest) ProductId(productId string) ApiGetRateHistoryRequest {
	r.productId = &productId
	return r
}

// \&quot;DAY\&quot;,\&quot;YEAR\&quot;,default\&quot;DAY\&quot;
func (r ApiGetRateHistoryRequest) AprPeriod(aprPeriod string) ApiGetRateHistoryRequest {
	r.aprPeriod = &aprPeriod
	return r
}

func (r ApiGetRateHistoryRequest) StartTime(startTime int64) ApiGetRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetRateHistoryRequest) EndTime(endTime int64) ApiGetRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetRateHistoryRequest) Current(current int64) ApiGetRateHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetRateHistoryRequest) Size(size int64) ApiGetRateHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetRateHistoryResponse], error) {
	return r.ApiService.GetRateHistoryExecute(r)
}

/*
GetRateHistory Get Rate History(USER_DATA)
Get /sapi/v1/simple-earn/flexible/history/rateHistory

https://developers.binance.com/docs/simple_earn/flexible-locked/history/Get-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param aprPeriod -  \"DAY\",\"YEAR\",default\"DAY\"
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRateHistoryRequest
*/
func (a *FlexibleLockedAPIService) GetRateHistory(ctx context.Context) ApiGetRateHistoryRequest {
	return ApiGetRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRateHistoryResponse
func (a *FlexibleLockedAPIService) GetRateHistoryExecute(r ApiGetRateHistoryRequest) (*common.RestApiResponse[models.GetRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/history/rateHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	if r.aprPeriod != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "aprPeriod", r.aprPeriod, "form", "")
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

	resp, err := SendRequest[models.GetRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSimpleEarnFlexibleProductListRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	asset      *string
	current    *int64
	size       *int64
	recvWindow *int64
}

// USDC or USDT
func (r ApiGetSimpleEarnFlexibleProductListRequest) Asset(asset string) ApiGetSimpleEarnFlexibleProductListRequest {
	r.asset = &asset
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetSimpleEarnFlexibleProductListRequest) Current(current int64) ApiGetSimpleEarnFlexibleProductListRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetSimpleEarnFlexibleProductListRequest) Size(size int64) ApiGetSimpleEarnFlexibleProductListRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetSimpleEarnFlexibleProductListRequest) RecvWindow(recvWindow int64) ApiGetSimpleEarnFlexibleProductListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSimpleEarnFlexibleProductListRequest) Execute() (*common.RestApiResponse[models.GetSimpleEarnFlexibleProductListResponse], error) {
	return r.ApiService.GetSimpleEarnFlexibleProductListExecute(r)
}

/*
GetSimpleEarnFlexibleProductList Get Simple Earn Flexible Product List(USER_DATA)
Get /sapi/v1/simple-earn/flexible/list

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Simple-Earn-Flexible-Product-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDC or USDT
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetSimpleEarnFlexibleProductListRequest
*/
func (a *FlexibleLockedAPIService) GetSimpleEarnFlexibleProductList(ctx context.Context) ApiGetSimpleEarnFlexibleProductListRequest {
	return ApiGetSimpleEarnFlexibleProductListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSimpleEarnFlexibleProductListResponse
func (a *FlexibleLockedAPIService) GetSimpleEarnFlexibleProductListExecute(r ApiGetSimpleEarnFlexibleProductListRequest) (*common.RestApiResponse[models.GetSimpleEarnFlexibleProductListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/list"

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

	resp, err := SendRequest[models.GetSimpleEarnFlexibleProductListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSimpleEarnLockedProductListRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	asset      *string
	current    *int64
	size       *int64
	recvWindow *int64
}

// USDC or USDT
func (r ApiGetSimpleEarnLockedProductListRequest) Asset(asset string) ApiGetSimpleEarnLockedProductListRequest {
	r.asset = &asset
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetSimpleEarnLockedProductListRequest) Current(current int64) ApiGetSimpleEarnLockedProductListRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetSimpleEarnLockedProductListRequest) Size(size int64) ApiGetSimpleEarnLockedProductListRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetSimpleEarnLockedProductListRequest) RecvWindow(recvWindow int64) ApiGetSimpleEarnLockedProductListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSimpleEarnLockedProductListRequest) Execute() (*common.RestApiResponse[models.GetSimpleEarnLockedProductListResponse], error) {
	return r.ApiService.GetSimpleEarnLockedProductListExecute(r)
}

/*
GetSimpleEarnLockedProductList Get Simple Earn Locked Product List(USER_DATA)
Get /sapi/v1/simple-earn/locked/list

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Get-Simple-Earn-Locked-Product-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDC or USDT
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetSimpleEarnLockedProductListRequest
*/
func (a *FlexibleLockedAPIService) GetSimpleEarnLockedProductList(ctx context.Context) ApiGetSimpleEarnLockedProductListRequest {
	return ApiGetSimpleEarnLockedProductListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSimpleEarnLockedProductListResponse
func (a *FlexibleLockedAPIService) GetSimpleEarnLockedProductListExecute(r ApiGetSimpleEarnLockedProductListRequest) (*common.RestApiResponse[models.GetSimpleEarnLockedProductListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/list"

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

	resp, err := SendRequest[models.GetSimpleEarnLockedProductListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemFlexibleProductRequest struct {
	ctx         context.Context
	ApiService  *FlexibleLockedAPIService
	productId   *string
	redeemAll   *bool
	amount      *float32
	destAccount *string
	recvWindow  *int64
}

func (r ApiRedeemFlexibleProductRequest) ProductId(productId string) ApiRedeemFlexibleProductRequest {
	r.productId = &productId
	return r
}

// true or false, default to false
func (r ApiRedeemFlexibleProductRequest) RedeemAll(redeemAll bool) ApiRedeemFlexibleProductRequest {
	r.redeemAll = &redeemAll
	return r
}

// if redeemAll is false, amount is mandatory
func (r ApiRedeemFlexibleProductRequest) Amount(amount float32) ApiRedeemFlexibleProductRequest {
	r.amount = &amount
	return r
}

// &#x60;SPOT&#x60;,&#x60;FUND&#x60;, default &#x60;SPOT&#x60;
func (r ApiRedeemFlexibleProductRequest) DestAccount(destAccount string) ApiRedeemFlexibleProductRequest {
	r.destAccount = &destAccount
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiRedeemFlexibleProductRequest) RecvWindow(recvWindow int64) ApiRedeemFlexibleProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemFlexibleProductRequest) Execute() (*common.RestApiResponse[models.RedeemFlexibleProductResponse], error) {
	return r.ApiService.RedeemFlexibleProductExecute(r)
}

/*
RedeemFlexibleProduct Redeem Flexible Product(TRADE)
Post /sapi/v1/simple-earn/flexible/redeem

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Redeem-Flexible-Product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param redeemAll -  true or false, default to false
@param amount -  if redeemAll is false, amount is mandatory
@param destAccount -  `SPOT`,`FUND`, default `SPOT`
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiRedeemFlexibleProductRequest
*/
func (a *FlexibleLockedAPIService) RedeemFlexibleProduct(ctx context.Context) ApiRedeemFlexibleProductRequest {
	return ApiRedeemFlexibleProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemFlexibleProductResponse
func (a *FlexibleLockedAPIService) RedeemFlexibleProductExecute(r ApiRedeemFlexibleProductRequest) (*common.RestApiResponse[models.RedeemFlexibleProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	if r.redeemAll != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "redeemAll", r.redeemAll, "form", "")
	}
	if r.amount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	}
	if r.destAccount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "destAccount", r.destAccount, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemFlexibleProductResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemLockedProductRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	positionId *string
	recvWindow *int64
}

func (r ApiRedeemLockedProductRequest) PositionId(positionId string) ApiRedeemLockedProductRequest {
	r.positionId = &positionId
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiRedeemLockedProductRequest) RecvWindow(recvWindow int64) ApiRedeemLockedProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemLockedProductRequest) Execute() (*common.RestApiResponse[models.RedeemLockedProductResponse], error) {
	return r.ApiService.RedeemLockedProductExecute(r)
}

/*
RedeemLockedProduct Redeem Locked Product(TRADE)
Post /sapi/v1/simple-earn/locked/redeem

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Redeem-Locked-Product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiRedeemLockedProductRequest
*/
func (a *FlexibleLockedAPIService) RedeemLockedProduct(ctx context.Context) ApiRedeemLockedProductRequest {
	return ApiRedeemLockedProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemLockedProductResponse
func (a *FlexibleLockedAPIService) RedeemLockedProductExecute(r ApiRedeemLockedProductRequest) (*common.RestApiResponse[models.RedeemLockedProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId == nil {
		return nil, common.ReportError("positionId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemLockedProductResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetFlexibleAutoSubscribeRequest struct {
	ctx           context.Context
	ApiService    *FlexibleLockedAPIService
	productId     *string
	autoSubscribe *bool
	recvWindow    *int64
}

func (r ApiSetFlexibleAutoSubscribeRequest) ProductId(productId string) ApiSetFlexibleAutoSubscribeRequest {
	r.productId = &productId
	return r
}

// true or false
func (r ApiSetFlexibleAutoSubscribeRequest) AutoSubscribe(autoSubscribe bool) ApiSetFlexibleAutoSubscribeRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSetFlexibleAutoSubscribeRequest) RecvWindow(recvWindow int64) ApiSetFlexibleAutoSubscribeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetFlexibleAutoSubscribeRequest) Execute() (*common.RestApiResponse[models.SetFlexibleAutoSubscribeResponse], error) {
	return r.ApiService.SetFlexibleAutoSubscribeExecute(r)
}

/*
SetFlexibleAutoSubscribe Set Flexible Auto Subscribe(USER_DATA)
Post /sapi/v1/simple-earn/flexible/setAutoSubscribe

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Set-Flexible-Auto-Subscribe

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param autoSubscribe -  true or false
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSetFlexibleAutoSubscribeRequest
*/
func (a *FlexibleLockedAPIService) SetFlexibleAutoSubscribe(ctx context.Context) ApiSetFlexibleAutoSubscribeRequest {
	return ApiSetFlexibleAutoSubscribeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetFlexibleAutoSubscribeResponse
func (a *FlexibleLockedAPIService) SetFlexibleAutoSubscribeExecute(r ApiSetFlexibleAutoSubscribeRequest) (*common.RestApiResponse[models.SetFlexibleAutoSubscribeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/setAutoSubscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}
	if r.autoSubscribe == nil {
		return nil, common.ReportError("autoSubscribe is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoSubscribe", r.autoSubscribe, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetFlexibleAutoSubscribeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetLockedAutoSubscribeRequest struct {
	ctx           context.Context
	ApiService    *FlexibleLockedAPIService
	positionId    *string
	autoSubscribe *bool
	recvWindow    *int64
}

func (r ApiSetLockedAutoSubscribeRequest) PositionId(positionId string) ApiSetLockedAutoSubscribeRequest {
	r.positionId = &positionId
	return r
}

// true or false
func (r ApiSetLockedAutoSubscribeRequest) AutoSubscribe(autoSubscribe bool) ApiSetLockedAutoSubscribeRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSetLockedAutoSubscribeRequest) RecvWindow(recvWindow int64) ApiSetLockedAutoSubscribeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetLockedAutoSubscribeRequest) Execute() (*common.RestApiResponse[models.SetLockedAutoSubscribeResponse], error) {
	return r.ApiService.SetLockedAutoSubscribeExecute(r)
}

/*
SetLockedAutoSubscribe Set Locked Auto Subscribe(USER_DATA)
Post /sapi/v1/simple-earn/locked/setAutoSubscribe

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Set-Locked-Auto-Subscribe

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param autoSubscribe -  true or false
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSetLockedAutoSubscribeRequest
*/
func (a *FlexibleLockedAPIService) SetLockedAutoSubscribe(ctx context.Context) ApiSetLockedAutoSubscribeRequest {
	return ApiSetLockedAutoSubscribeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetLockedAutoSubscribeResponse
func (a *FlexibleLockedAPIService) SetLockedAutoSubscribeExecute(r ApiSetLockedAutoSubscribeRequest) (*common.RestApiResponse[models.SetLockedAutoSubscribeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/setAutoSubscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId == nil {
		return nil, common.ReportError("positionId is required and must be specified")
	}
	if r.autoSubscribe == nil {
		return nil, common.ReportError("autoSubscribe is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoSubscribe", r.autoSubscribe, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetLockedAutoSubscribeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSetLockedProductRedeemOptionRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	positionId *string
	redeemTo   *string
	recvWindow *int64
}

func (r ApiSetLockedProductRedeemOptionRequest) PositionId(positionId string) ApiSetLockedProductRedeemOptionRequest {
	r.positionId = &positionId
	return r
}

// &#x60;SPOT&#x60;,&#39;FLEXIBLE&#39;
func (r ApiSetLockedProductRedeemOptionRequest) RedeemTo(redeemTo string) ApiSetLockedProductRedeemOptionRequest {
	r.redeemTo = &redeemTo
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSetLockedProductRedeemOptionRequest) RecvWindow(recvWindow int64) ApiSetLockedProductRedeemOptionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetLockedProductRedeemOptionRequest) Execute() (*common.RestApiResponse[models.SetLockedProductRedeemOptionResponse], error) {
	return r.ApiService.SetLockedProductRedeemOptionExecute(r)
}

/*
SetLockedProductRedeemOption Set Locked Product Redeem Option(USER_DATA)
Post /sapi/v1/simple-earn/locked/setRedeemOption

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Set-Locked-Redeem-Option

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param redeemTo -  `SPOT`,'FLEXIBLE'
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSetLockedProductRedeemOptionRequest
*/
func (a *FlexibleLockedAPIService) SetLockedProductRedeemOption(ctx context.Context) ApiSetLockedProductRedeemOptionRequest {
	return ApiSetLockedProductRedeemOptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetLockedProductRedeemOptionResponse
func (a *FlexibleLockedAPIService) SetLockedProductRedeemOptionExecute(r ApiSetLockedProductRedeemOptionRequest) (*common.RestApiResponse[models.SetLockedProductRedeemOptionResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/setRedeemOption"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId == nil {
		return nil, common.ReportError("positionId is required and must be specified")
	}
	if r.redeemTo == nil {
		return nil, common.ReportError("redeemTo is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "redeemTo", r.redeemTo, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SetLockedProductRedeemOptionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSimpleAccountRequest struct {
	ctx        context.Context
	ApiService *FlexibleLockedAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000 (ms)
func (r ApiSimpleAccountRequest) RecvWindow(recvWindow int64) ApiSimpleAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSimpleAccountRequest) Execute() (*common.RestApiResponse[models.SimpleAccountResponse], error) {
	return r.ApiService.SimpleAccountExecute(r)
}

/*
SimpleAccount Simple Account(USER_DATA)
Get /sapi/v1/simple-earn/account

https://developers.binance.com/docs/simple_earn/flexible-locked/account/Simple-Account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSimpleAccountRequest
*/
func (a *FlexibleLockedAPIService) SimpleAccount(ctx context.Context) ApiSimpleAccountRequest {
	return ApiSimpleAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SimpleAccountResponse
func (a *FlexibleLockedAPIService) SimpleAccountExecute(r ApiSimpleAccountRequest) (*common.RestApiResponse[models.SimpleAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SimpleAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeFlexibleProductRequest struct {
	ctx           context.Context
	ApiService    *FlexibleLockedAPIService
	productId     *string
	amount        *float32
	autoSubscribe *bool
	sourceAccount *string
	recvWindow    *int64
}

func (r ApiSubscribeFlexibleProductRequest) ProductId(productId string) ApiSubscribeFlexibleProductRequest {
	r.productId = &productId
	return r
}

// Amount
func (r ApiSubscribeFlexibleProductRequest) Amount(amount float32) ApiSubscribeFlexibleProductRequest {
	r.amount = &amount
	return r
}

// true or false, default true.
func (r ApiSubscribeFlexibleProductRequest) AutoSubscribe(autoSubscribe bool) ApiSubscribeFlexibleProductRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// &#x60;SPOT&#x60;,&#x60;FUND&#x60;,&#x60;ALL&#x60;, default &#x60;SPOT&#x60;
func (r ApiSubscribeFlexibleProductRequest) SourceAccount(sourceAccount string) ApiSubscribeFlexibleProductRequest {
	r.sourceAccount = &sourceAccount
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSubscribeFlexibleProductRequest) RecvWindow(recvWindow int64) ApiSubscribeFlexibleProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeFlexibleProductRequest) Execute() (*common.RestApiResponse[models.SubscribeFlexibleProductResponse], error) {
	return r.ApiService.SubscribeFlexibleProductExecute(r)
}

/*
SubscribeFlexibleProduct Subscribe Flexible Product(TRADE)
Post /sapi/v1/simple-earn/flexible/subscribe

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Subscribe-Flexible-Product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param productId -
@param amount -  Amount
@param autoSubscribe -  true or false, default true.
@param sourceAccount -  `SPOT`,`FUND`,`ALL`, default `SPOT`
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSubscribeFlexibleProductRequest
*/
func (a *FlexibleLockedAPIService) SubscribeFlexibleProduct(ctx context.Context) ApiSubscribeFlexibleProductRequest {
	return ApiSubscribeFlexibleProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeFlexibleProductResponse
func (a *FlexibleLockedAPIService) SubscribeFlexibleProductExecute(r ApiSubscribeFlexibleProductRequest) (*common.RestApiResponse[models.SubscribeFlexibleProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/flexible/subscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.productId == nil {
		return nil, common.ReportError("productId is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.autoSubscribe != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoSubscribe", r.autoSubscribe, "form", "")
	}
	if r.sourceAccount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sourceAccount", r.sourceAccount, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeFlexibleProductResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeLockedProductRequest struct {
	ctx           context.Context
	ApiService    *FlexibleLockedAPIService
	projectId     *string
	amount        *float32
	autoSubscribe *bool
	sourceAccount *string
	redeemTo      *string
	recvWindow    *int64
}

func (r ApiSubscribeLockedProductRequest) ProjectId(projectId string) ApiSubscribeLockedProductRequest {
	r.projectId = &projectId
	return r
}

// Amount
func (r ApiSubscribeLockedProductRequest) Amount(amount float32) ApiSubscribeLockedProductRequest {
	r.amount = &amount
	return r
}

// true or false, default true.
func (r ApiSubscribeLockedProductRequest) AutoSubscribe(autoSubscribe bool) ApiSubscribeLockedProductRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// &#x60;SPOT&#x60;,&#x60;FUND&#x60;,&#x60;ALL&#x60;, default &#x60;SPOT&#x60;
func (r ApiSubscribeLockedProductRequest) SourceAccount(sourceAccount string) ApiSubscribeLockedProductRequest {
	r.sourceAccount = &sourceAccount
	return r
}

// &#x60;SPOT&#x60;,&#x60;FLEXIBLE&#x60;, default &#x60;SPOT&#x60;
func (r ApiSubscribeLockedProductRequest) RedeemTo(redeemTo string) ApiSubscribeLockedProductRequest {
	r.redeemTo = &redeemTo
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSubscribeLockedProductRequest) RecvWindow(recvWindow int64) ApiSubscribeLockedProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeLockedProductRequest) Execute() (*common.RestApiResponse[models.SubscribeLockedProductResponse], error) {
	return r.ApiService.SubscribeLockedProductExecute(r)
}

/*
SubscribeLockedProduct Subscribe Locked Product(TRADE)
Post /sapi/v1/simple-earn/locked/subscribe

https://developers.binance.com/docs/simple_earn/flexible-locked/earn/Subscribe-Locked-Product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param amount -  Amount
@param autoSubscribe -  true or false, default true.
@param sourceAccount -  `SPOT`,`FUND`,`ALL`, default `SPOT`
@param redeemTo -  `SPOT`,`FLEXIBLE`, default `SPOT`
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSubscribeLockedProductRequest
*/
func (a *FlexibleLockedAPIService) SubscribeLockedProduct(ctx context.Context) ApiSubscribeLockedProductRequest {
	return ApiSubscribeLockedProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeLockedProductResponse
func (a *FlexibleLockedAPIService) SubscribeLockedProductExecute(r ApiSubscribeLockedProductRequest) (*common.RestApiResponse[models.SubscribeLockedProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/simple-earn/locked/subscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.projectId == nil {
		return nil, common.ReportError("projectId is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.autoSubscribe != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoSubscribe", r.autoSubscribe, "form", "")
	}
	if r.sourceAccount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sourceAccount", r.sourceAccount, "form", "")
	}
	if r.redeemTo != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "redeemTo", r.redeemTo, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeLockedProductResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
