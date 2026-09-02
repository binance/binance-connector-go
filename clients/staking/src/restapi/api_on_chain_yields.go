/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package binancestakingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// OnChainYieldsAPIService OnChainYieldsAPI Service
type OnChainYieldsAPIService Service

type ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	projectId  *string
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest) ProjectId(projectId string) ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest {
	r.projectId = &projectId
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedPersonalLeftQuotaResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedPersonalLeftQuotaExecute(r)
}

/*
GetOnChainYieldsLockedPersonalLeftQuota Get On-chain Yields Locked Personal Left Quota (USER_DATA)
Get /sapi/v1/onchain-yields/locked/personalLeftQuota

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-personal-left-quota

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedPersonalLeftQuota(ctx context.Context) ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest {
	return ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedPersonalLeftQuotaResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedPersonalLeftQuotaExecute(r ApiGetOnChainYieldsLockedPersonalLeftQuotaRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedPersonalLeftQuotaResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/personalLeftQuota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.projectId == nil {
		return nil, common.ReportError("projectId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOnChainYieldsLockedPersonalLeftQuotaResponse](
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

type ApiGetOnChainYieldsLockedProductListRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	asset      *string
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedProductListRequest) Asset(asset string) ApiGetOnChainYieldsLockedProductListRequest {
	r.asset = &asset
	return r
}

// Currently querying page
func (r ApiGetOnChainYieldsLockedProductListRequest) Current(current int64) ApiGetOnChainYieldsLockedProductListRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetOnChainYieldsLockedProductListRequest) Size(size int64) ApiGetOnChainYieldsLockedProductListRequest {
	r.size = &size
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedProductListRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedProductListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedProductListRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedProductListResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedProductListExecute(r)
}

/*
GetOnChainYieldsLockedProductList Get On-chain Yields Locked Product List (USER_DATA)
Get /sapi/v1/onchain-yields/locked/list

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-product-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param current -  Currently querying page
@param size -  Number of results per page.
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedProductListRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedProductList(ctx context.Context) ApiGetOnChainYieldsLockedProductListRequest {
	return ApiGetOnChainYieldsLockedProductListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedProductListResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedProductListExecute(r ApiGetOnChainYieldsLockedProductListRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedProductListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/list"

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

	resp, err := SendRequest[models.GetOnChainYieldsLockedProductListResponse](
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

type ApiGetOnChainYieldsLockedProductPositionRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	asset      *string
	positionId *string
	projectId  *string
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedProductPositionRequest) Asset(asset string) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.asset = &asset
	return r
}

func (r ApiGetOnChainYieldsLockedProductPositionRequest) PositionId(positionId string) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.positionId = &positionId
	return r
}

func (r ApiGetOnChainYieldsLockedProductPositionRequest) ProjectId(projectId string) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.projectId = &projectId
	return r
}

// Currently querying page
func (r ApiGetOnChainYieldsLockedProductPositionRequest) Current(current int64) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetOnChainYieldsLockedProductPositionRequest) Size(size int64) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.size = &size
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedProductPositionRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedProductPositionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedProductPositionRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedProductPositionResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedProductPositionExecute(r)
}

/*
GetOnChainYieldsLockedProductPosition Get On-chain Yields Locked Product Position (USER_DATA)
Get /sapi/v1/onchain-yields/locked/position

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-product-position

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param positionId -
@param projectId -
@param current -  Currently querying page
@param size -  Number of results per page.
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedProductPositionRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedProductPosition(ctx context.Context) ApiGetOnChainYieldsLockedProductPositionRequest {
	return ApiGetOnChainYieldsLockedProductPositionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedProductPositionResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedProductPositionExecute(r ApiGetOnChainYieldsLockedProductPositionRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedProductPositionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/position"

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

	resp, err := SendRequest[models.GetOnChainYieldsLockedProductPositionResponse](
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

type ApiGetOnChainYieldsLockedRedemptionRecordRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	positionId *string
	redeemId   *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) PositionId(positionId string) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.positionId = &positionId
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) RedeemId(redeemId string) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.redeemId = &redeemId
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) Asset(asset string) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) StartTime(startTime int64) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) EndTime(endTime int64) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page
func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) Current(current int64) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.current = &current
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) Size(size int64) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.size = &size
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedRedemptionRecordRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedRedemptionRecordResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedRedemptionRecordExecute(r)
}

/*
GetOnChainYieldsLockedRedemptionRecord Get On-chain Yields Locked Redemption Record (USER_DATA)
Get /sapi/v1/onchain-yields/locked/history/redemptionRecord

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-redemption-record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param redeemId -
@param asset -
@param startTime -
@param endTime -
@param current -  Currently querying page
@param size -
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedRedemptionRecordRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedRedemptionRecord(ctx context.Context) ApiGetOnChainYieldsLockedRedemptionRecordRequest {
	return ApiGetOnChainYieldsLockedRedemptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedRedemptionRecordResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedRedemptionRecordExecute(r ApiGetOnChainYieldsLockedRedemptionRecordRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedRedemptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/history/redemptionRecord"

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

	resp, err := SendRequest[models.GetOnChainYieldsLockedRedemptionRecordResponse](
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

type ApiGetOnChainYieldsLockedRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	positionId *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) PositionId(positionId string) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.positionId = &positionId
	return r
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) Asset(asset string) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) StartTime(startTime int64) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) EndTime(endTime int64) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page
func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) Current(current int64) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.current = &current
	return r
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) Size(size int64) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.size = &size
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedRewardsHistoryResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedRewardsHistoryExecute(r)
}

/*
GetOnChainYieldsLockedRewardsHistory Get On-chain Yields Locked Rewards History (USER_DATA)
Get /sapi/v1/onchain-yields/locked/history/rewardsRecord

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-rewards-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param asset -
@param startTime -
@param endTime -
@param current -  Currently querying page
@param size -
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedRewardsHistoryRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedRewardsHistory(ctx context.Context) ApiGetOnChainYieldsLockedRewardsHistoryRequest {
	return ApiGetOnChainYieldsLockedRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedRewardsHistoryResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedRewardsHistoryExecute(r ApiGetOnChainYieldsLockedRewardsHistoryRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/history/rewardsRecord"

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

	resp, err := SendRequest[models.GetOnChainYieldsLockedRewardsHistoryResponse](
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

type ApiGetOnChainYieldsLockedSubscriptionPreviewRequest struct {
	ctx           context.Context
	ApiService    *OnChainYieldsAPIService
	projectId     *string
	amount        *float64
	autoSubscribe *bool
	recvWindow    *int64
}

func (r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) ProjectId(projectId string) ApiGetOnChainYieldsLockedSubscriptionPreviewRequest {
	r.projectId = &projectId
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) Amount(amount float64) ApiGetOnChainYieldsLockedSubscriptionPreviewRequest {
	r.amount = &amount
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) AutoSubscribe(autoSubscribe bool) ApiGetOnChainYieldsLockedSubscriptionPreviewRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedSubscriptionPreviewRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionPreviewResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedSubscriptionPreviewExecute(r)
}

/*
GetOnChainYieldsLockedSubscriptionPreview Get On-chain Yields Locked Subscription Preview (USER_DATA)
Get /sapi/v1/onchain-yields/locked/subscriptionPreview

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-subscription-preview

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param amount -
@param autoSubscribe -
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedSubscriptionPreviewRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedSubscriptionPreview(ctx context.Context) ApiGetOnChainYieldsLockedSubscriptionPreviewRequest {
	return ApiGetOnChainYieldsLockedSubscriptionPreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedSubscriptionPreviewResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedSubscriptionPreviewExecute(r ApiGetOnChainYieldsLockedSubscriptionPreviewRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionPreviewResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/subscriptionPreview"

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

	resp, err := SendRequest[models.GetOnChainYieldsLockedSubscriptionPreviewResponse](
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

type ApiGetOnChainYieldsLockedSubscriptionRecordRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	purchaseId *string
	clientId   *string
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) PurchaseId(purchaseId string) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.purchaseId = &purchaseId
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) ClientId(clientId string) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) Asset(asset string) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) StartTime(startTime int64) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) EndTime(endTime int64) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page
func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) Current(current int64) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.current = &current
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) Size(size int64) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.size = &size
	return r
}

// Request validity window in milliseconds.
func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) RecvWindow(recvWindow int64) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) Execute() (*common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionRecordResponse], error) {
	return r.ApiService.GetOnChainYieldsLockedSubscriptionRecordExecute(r)
}

/*
GetOnChainYieldsLockedSubscriptionRecord Get On-chain Yields Locked Subscription Record (USER_DATA)
Get /sapi/v1/onchain-yields/locked/history/subscriptionRecord

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#get-on-chain-yields-locked-subscription-record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param purchaseId -
@param clientId -
@param asset -
@param startTime -
@param endTime -
@param current -  Currently querying page
@param size -
@param recvWindow -  Request validity window in milliseconds.
@return ApiGetOnChainYieldsLockedSubscriptionRecordRequest
*/
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedSubscriptionRecord(ctx context.Context) ApiGetOnChainYieldsLockedSubscriptionRecordRequest {
	return ApiGetOnChainYieldsLockedSubscriptionRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOnChainYieldsLockedSubscriptionRecordResponse
func (a *OnChainYieldsAPIService) GetOnChainYieldsLockedSubscriptionRecordExecute(r ApiGetOnChainYieldsLockedSubscriptionRecordRequest) (*common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/history/subscriptionRecord"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.purchaseId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "purchaseId", r.purchaseId, "form", "")
	}
	if r.clientId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "form", "")
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

	resp, err := SendRequest[models.GetOnChainYieldsLockedSubscriptionRecordResponse](
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

type ApiOnChainYieldsAccountRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	recvWindow *int64
}

// The value cannot be greater than &#x60;60000&#x60;
func (r ApiOnChainYieldsAccountRequest) RecvWindow(recvWindow int64) ApiOnChainYieldsAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOnChainYieldsAccountRequest) Execute() (*common.RestApiResponse[models.OnChainYieldsAccountResponse], error) {
	return r.ApiService.OnChainYieldsAccountExecute(r)
}

/*
OnChainYieldsAccount On-chain Yields Account (USER_DATA)
Get /sapi/v1/onchain-yields/account

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#on-chain-yields-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`
@return ApiOnChainYieldsAccountRequest
*/
func (a *OnChainYieldsAPIService) OnChainYieldsAccount(ctx context.Context) ApiOnChainYieldsAccountRequest {
	return ApiOnChainYieldsAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OnChainYieldsAccountResponse
func (a *OnChainYieldsAPIService) OnChainYieldsAccountExecute(r ApiOnChainYieldsAccountRequest) (*common.RestApiResponse[models.OnChainYieldsAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OnChainYieldsAccountResponse](
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

type ApiRedeemOnChainYieldsLockedProductRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	positionId *string
	channelId  *string
	recvWindow *int64
}

// Locked product position ID
func (r ApiRedeemOnChainYieldsLockedProductRequest) PositionId(positionId string) ApiRedeemOnChainYieldsLockedProductRequest {
	r.positionId = &positionId
	return r
}

func (r ApiRedeemOnChainYieldsLockedProductRequest) ChannelId(channelId string) ApiRedeemOnChainYieldsLockedProductRequest {
	r.channelId = &channelId
	return r
}

// Request validity window in milliseconds.
func (r ApiRedeemOnChainYieldsLockedProductRequest) RecvWindow(recvWindow int64) ApiRedeemOnChainYieldsLockedProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemOnChainYieldsLockedProductRequest) Execute() (*common.RestApiResponse[models.RedeemOnChainYieldsLockedProductResponse], error) {
	return r.ApiService.RedeemOnChainYieldsLockedProductExecute(r)
}

/*
RedeemOnChainYieldsLockedProduct Redeem On-chain Yields Locked Product (TRADE)
Post /sapi/v1/onchain-yields/locked/redeem

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#redeem-on-chain-yields-locked-product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -  Locked product position ID
@param channelId -
@param recvWindow -  Request validity window in milliseconds.
@return ApiRedeemOnChainYieldsLockedProductRequest
*/
func (a *OnChainYieldsAPIService) RedeemOnChainYieldsLockedProduct(ctx context.Context) ApiRedeemOnChainYieldsLockedProductRequest {
	return ApiRedeemOnChainYieldsLockedProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemOnChainYieldsLockedProductResponse
func (a *OnChainYieldsAPIService) RedeemOnChainYieldsLockedProductExecute(r ApiRedeemOnChainYieldsLockedProductRequest) (*common.RestApiResponse[models.RedeemOnChainYieldsLockedProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId == nil {
		return nil, common.ReportError("positionId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	if r.channelId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "channelId", r.channelId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemOnChainYieldsLockedProductResponse](
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

type ApiSetOnChainYieldsLockedAutoSubscribeRequest struct {
	ctx           context.Context
	ApiService    *OnChainYieldsAPIService
	positionId    *string
	autoSubscribe *bool
	recvWindow    *int64
}

func (r ApiSetOnChainYieldsLockedAutoSubscribeRequest) PositionId(positionId string) ApiSetOnChainYieldsLockedAutoSubscribeRequest {
	r.positionId = &positionId
	return r
}

func (r ApiSetOnChainYieldsLockedAutoSubscribeRequest) AutoSubscribe(autoSubscribe bool) ApiSetOnChainYieldsLockedAutoSubscribeRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

// Request validity window in milliseconds.
func (r ApiSetOnChainYieldsLockedAutoSubscribeRequest) RecvWindow(recvWindow int64) ApiSetOnChainYieldsLockedAutoSubscribeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetOnChainYieldsLockedAutoSubscribeRequest) Execute() (*common.RestApiResponse[models.SetOnChainYieldsLockedAutoSubscribeResponse], error) {
	return r.ApiService.SetOnChainYieldsLockedAutoSubscribeExecute(r)
}

/*
SetOnChainYieldsLockedAutoSubscribe Set On-chain Yields Locked Auto Subscribe (USER_DATA)
Post /sapi/v1/onchain-yields/locked/setAutoSubscribe

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#set-on-chain-yields-locked-auto-subscribe

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param autoSubscribe -
@param recvWindow -  Request validity window in milliseconds.
@return ApiSetOnChainYieldsLockedAutoSubscribeRequest
*/
func (a *OnChainYieldsAPIService) SetOnChainYieldsLockedAutoSubscribe(ctx context.Context) ApiSetOnChainYieldsLockedAutoSubscribeRequest {
	return ApiSetOnChainYieldsLockedAutoSubscribeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetOnChainYieldsLockedAutoSubscribeResponse
func (a *OnChainYieldsAPIService) SetOnChainYieldsLockedAutoSubscribeExecute(r ApiSetOnChainYieldsLockedAutoSubscribeRequest) (*common.RestApiResponse[models.SetOnChainYieldsLockedAutoSubscribeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/setAutoSubscribe"

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

	resp, err := SendRequest[models.SetOnChainYieldsLockedAutoSubscribeResponse](
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

type ApiSetOnChainYieldsLockedProductRedeemOptionRequest struct {
	ctx        context.Context
	ApiService *OnChainYieldsAPIService
	positionId *string
	redeemTo   *models.SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter
	recvWindow *int64
}

func (r ApiSetOnChainYieldsLockedProductRedeemOptionRequest) PositionId(positionId string) ApiSetOnChainYieldsLockedProductRedeemOptionRequest {
	r.positionId = &positionId
	return r
}

func (r ApiSetOnChainYieldsLockedProductRedeemOptionRequest) RedeemTo(redeemTo models.SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) ApiSetOnChainYieldsLockedProductRedeemOptionRequest {
	r.redeemTo = &redeemTo
	return r
}

// Request validity window in milliseconds.
func (r ApiSetOnChainYieldsLockedProductRedeemOptionRequest) RecvWindow(recvWindow int64) ApiSetOnChainYieldsLockedProductRedeemOptionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSetOnChainYieldsLockedProductRedeemOptionRequest) Execute() (*common.RestApiResponse[models.SetOnChainYieldsLockedProductRedeemOptionResponse], error) {
	return r.ApiService.SetOnChainYieldsLockedProductRedeemOptionExecute(r)
}

/*
SetOnChainYieldsLockedProductRedeemOption Set On-chain Yields Locked Product Redeem Option (USER_DATA)
Post /sapi/v1/onchain-yields/locked/setRedeemOption

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#set-on-chain-yields-locked-product-redeem-option

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -
@param redeemTo -
@param recvWindow -  Request validity window in milliseconds.
@return ApiSetOnChainYieldsLockedProductRedeemOptionRequest
*/
func (a *OnChainYieldsAPIService) SetOnChainYieldsLockedProductRedeemOption(ctx context.Context) ApiSetOnChainYieldsLockedProductRedeemOptionRequest {
	return ApiSetOnChainYieldsLockedProductRedeemOptionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SetOnChainYieldsLockedProductRedeemOptionResponse
func (a *OnChainYieldsAPIService) SetOnChainYieldsLockedProductRedeemOptionExecute(r ApiSetOnChainYieldsLockedProductRedeemOptionRequest) (*common.RestApiResponse[models.SetOnChainYieldsLockedProductRedeemOptionResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/setRedeemOption"

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

	resp, err := SendRequest[models.SetOnChainYieldsLockedProductRedeemOptionResponse](
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

type ApiSubscribeOnChainYieldsLockedProductRequest struct {
	ctx           context.Context
	ApiService    *OnChainYieldsAPIService
	projectId     *string
	amount        *float64
	autoSubscribe *bool
	sourceAccount *models.SubscribeOnChainYieldsLockedProductSourceAccountParameter
	redeemTo      *models.SubscribeOnChainYieldsLockedProductRedeemToParameter
	channelId     *string
	clientId      *string
	recvWindow    *int64
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) ProjectId(projectId string) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.projectId = &projectId
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) Amount(amount float64) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.amount = &amount
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) AutoSubscribe(autoSubscribe bool) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.autoSubscribe = &autoSubscribe
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) SourceAccount(sourceAccount models.SubscribeOnChainYieldsLockedProductSourceAccountParameter) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.sourceAccount = &sourceAccount
	return r
}

// Takes effect when Auto Subscribe is false
func (r ApiSubscribeOnChainYieldsLockedProductRequest) RedeemTo(redeemTo models.SubscribeOnChainYieldsLockedProductRedeemToParameter) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.redeemTo = &redeemTo
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) ChannelId(channelId string) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.channelId = &channelId
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) ClientId(clientId string) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.clientId = &clientId
	return r
}

// Request validity window in milliseconds.
func (r ApiSubscribeOnChainYieldsLockedProductRequest) RecvWindow(recvWindow int64) ApiSubscribeOnChainYieldsLockedProductRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeOnChainYieldsLockedProductRequest) Execute() (*common.RestApiResponse[models.SubscribeOnChainYieldsLockedProductResponse], error) {
	return r.ApiService.SubscribeOnChainYieldsLockedProductExecute(r)
}

/*
SubscribeOnChainYieldsLockedProduct Subscribe On-chain Yields Locked Product (TRADE)
Post /sapi/v1/onchain-yields/locked/subscribe

https://developers.binance.com/en/docs/catalog/investment-and-services-staking/api/rest-api/on-chain-yields#subscribe-on-chain-yields-locked-product

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param projectId -
@param amount -
@param autoSubscribe -
@param sourceAccount -
@param redeemTo -  Takes effect when Auto Subscribe is false
@param channelId -
@param clientId -
@param recvWindow -  Request validity window in milliseconds.
@return ApiSubscribeOnChainYieldsLockedProductRequest
*/
func (a *OnChainYieldsAPIService) SubscribeOnChainYieldsLockedProduct(ctx context.Context) ApiSubscribeOnChainYieldsLockedProductRequest {
	return ApiSubscribeOnChainYieldsLockedProductRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeOnChainYieldsLockedProductResponse
func (a *OnChainYieldsAPIService) SubscribeOnChainYieldsLockedProductExecute(r ApiSubscribeOnChainYieldsLockedProductRequest) (*common.RestApiResponse[models.SubscribeOnChainYieldsLockedProductResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/onchain-yields/locked/subscribe"

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
	if r.channelId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "channelId", r.channelId, "form", "")
	}
	if r.clientId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeOnChainYieldsLockedProductResponse](
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
