/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API
*/

package binancestakingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// SolStakingAPIService SolStakingAPI Service
type SolStakingAPIService Service

type ApiClaimBoostRewardsRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	recvWindow *int64
}

func (r ApiClaimBoostRewardsRequest) RecvWindow(recvWindow int64) ApiClaimBoostRewardsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiClaimBoostRewardsRequest) Execute() (*common.RestApiResponse[models.ClaimBoostRewardsResponse], error) {
	return r.ApiService.ClaimBoostRewardsExecute(r)
}

/*
ClaimBoostRewards Claim Boost Rewards(TRADE)
Post /sapi/v1/sol-staking/sol/claim

https://developers.binance.com/docs/staking/sol-staking/staking/Claim-Boost-Rewards

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiClaimBoostRewardsRequest
*/
func (a *SolStakingAPIService) ClaimBoostRewards(ctx context.Context) ApiClaimBoostRewardsRequest {
	return ApiClaimBoostRewardsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClaimBoostRewardsResponse
func (a *SolStakingAPIService) ClaimBoostRewardsExecute(r ApiClaimBoostRewardsRequest) (*common.RestApiResponse[models.ClaimBoostRewardsResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/claim"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ClaimBoostRewardsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetBnsolRateHistoryRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBnsolRateHistoryRequest) StartTime(startTime int64) ApiGetBnsolRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBnsolRateHistoryRequest) EndTime(endTime int64) ApiGetBnsolRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetBnsolRateHistoryRequest) Current(current int64) ApiGetBnsolRateHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetBnsolRateHistoryRequest) Size(size int64) ApiGetBnsolRateHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBnsolRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetBnsolRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBnsolRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetBnsolRateHistoryResponse], error) {
	return r.ApiService.GetBnsolRateHistoryExecute(r)
}

/*
GetBnsolRateHistory Get BNSOL Rate History(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/rateHistory

https://developers.binance.com/docs/staking/sol-staking/history/Get-BNSOL-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetBnsolRateHistoryRequest
*/
func (a *SolStakingAPIService) GetBnsolRateHistory(ctx context.Context) ApiGetBnsolRateHistoryRequest {
	return ApiGetBnsolRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBnsolRateHistoryResponse
func (a *SolStakingAPIService) GetBnsolRateHistoryExecute(r ApiGetBnsolRateHistoryRequest) (*common.RestApiResponse[models.GetBnsolRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/rateHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetBnsolRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetBnsolRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBnsolRewardsHistoryRequest) StartTime(startTime int64) ApiGetBnsolRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBnsolRewardsHistoryRequest) EndTime(endTime int64) ApiGetBnsolRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetBnsolRewardsHistoryRequest) Current(current int64) ApiGetBnsolRewardsHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetBnsolRewardsHistoryRequest) Size(size int64) ApiGetBnsolRewardsHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBnsolRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetBnsolRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBnsolRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetBnsolRewardsHistoryResponse], error) {
	return r.ApiService.GetBnsolRewardsHistoryExecute(r)
}

/*
GetBnsolRewardsHistory Get BNSOL rewards history(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/bnsolRewardsHistory

https://developers.binance.com/docs/staking/sol-staking/history/Get-BNSOL-rewards-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetBnsolRewardsHistoryRequest
*/
func (a *SolStakingAPIService) GetBnsolRewardsHistory(ctx context.Context) ApiGetBnsolRewardsHistoryRequest {
	return ApiGetBnsolRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBnsolRewardsHistoryResponse
func (a *SolStakingAPIService) GetBnsolRewardsHistoryExecute(r ApiGetBnsolRewardsHistoryRequest) (*common.RestApiResponse[models.GetBnsolRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/bnsolRewardsHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetBnsolRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetBoostRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	type_      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// \&quot;CLAIM\&quot;, \&quot;DISTRIBUTE\&quot;, default \&quot;CLAIM\&quot;
func (r ApiGetBoostRewardsHistoryRequest) Type(type_ string) ApiGetBoostRewardsHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetBoostRewardsHistoryRequest) StartTime(startTime int64) ApiGetBoostRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBoostRewardsHistoryRequest) EndTime(endTime int64) ApiGetBoostRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetBoostRewardsHistoryRequest) Current(current int64) ApiGetBoostRewardsHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetBoostRewardsHistoryRequest) Size(size int64) ApiGetBoostRewardsHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBoostRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetBoostRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBoostRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetBoostRewardsHistoryResponse], error) {
	return r.ApiService.GetBoostRewardsHistoryExecute(r)
}

/*
GetBoostRewardsHistory Get Boost Rewards History(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/boostRewardsHistory

https://developers.binance.com/docs/staking/sol-staking/history/Get-Boost-Rewards-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -  \"CLAIM\", \"DISTRIBUTE\", default \"CLAIM\"
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetBoostRewardsHistoryRequest
*/
func (a *SolStakingAPIService) GetBoostRewardsHistory(ctx context.Context) ApiGetBoostRewardsHistoryRequest {
	return ApiGetBoostRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBoostRewardsHistoryResponse
func (a *SolStakingAPIService) GetBoostRewardsHistoryExecute(r ApiGetBoostRewardsHistoryRequest) (*common.RestApiResponse[models.GetBoostRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/boostRewardsHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

	resp, err := SendRequest[models.GetBoostRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSolRedemptionHistoryRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetSolRedemptionHistoryRequest) StartTime(startTime int64) ApiGetSolRedemptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSolRedemptionHistoryRequest) EndTime(endTime int64) ApiGetSolRedemptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetSolRedemptionHistoryRequest) Current(current int64) ApiGetSolRedemptionHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetSolRedemptionHistoryRequest) Size(size int64) ApiGetSolRedemptionHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetSolRedemptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetSolRedemptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSolRedemptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetSolRedemptionHistoryResponse], error) {
	return r.ApiService.GetSolRedemptionHistoryExecute(r)
}

/*
GetSolRedemptionHistory Get SOL redemption history(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/redemptionHistory

https://developers.binance.com/docs/staking/sol-staking/history/Get-SOL-redemption-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetSolRedemptionHistoryRequest
*/
func (a *SolStakingAPIService) GetSolRedemptionHistory(ctx context.Context) ApiGetSolRedemptionHistoryRequest {
	return ApiGetSolRedemptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSolRedemptionHistoryResponse
func (a *SolStakingAPIService) GetSolRedemptionHistoryExecute(r ApiGetSolRedemptionHistoryRequest) (*common.RestApiResponse[models.GetSolRedemptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/redemptionHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetSolRedemptionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSolStakingHistoryRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetSolStakingHistoryRequest) StartTime(startTime int64) ApiGetSolStakingHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetSolStakingHistoryRequest) EndTime(endTime int64) ApiGetSolStakingHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetSolStakingHistoryRequest) Current(current int64) ApiGetSolStakingHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetSolStakingHistoryRequest) Size(size int64) ApiGetSolStakingHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetSolStakingHistoryRequest) RecvWindow(recvWindow int64) ApiGetSolStakingHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSolStakingHistoryRequest) Execute() (*common.RestApiResponse[models.GetSolStakingHistoryResponse], error) {
	return r.ApiService.GetSolStakingHistoryExecute(r)
}

/*
GetSolStakingHistory Get SOL staking history(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/stakingHistory

https://developers.binance.com/docs/staking/sol-staking/history/Get-SOL-staking-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetSolStakingHistoryRequest
*/
func (a *SolStakingAPIService) GetSolStakingHistory(ctx context.Context) ApiGetSolStakingHistoryRequest {
	return ApiGetSolStakingHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSolStakingHistoryResponse
func (a *SolStakingAPIService) GetSolStakingHistoryExecute(r ApiGetSolStakingHistoryRequest) (*common.RestApiResponse[models.GetSolStakingHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/stakingHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetSolStakingHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetSolStakingQuotaDetailsRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	recvWindow *int64
}

func (r ApiGetSolStakingQuotaDetailsRequest) RecvWindow(recvWindow int64) ApiGetSolStakingQuotaDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetSolStakingQuotaDetailsRequest) Execute() (*common.RestApiResponse[models.GetSolStakingQuotaDetailsResponse], error) {
	return r.ApiService.GetSolStakingQuotaDetailsExecute(r)
}

/*
GetSolStakingQuotaDetails Get SOL staking quota details(USER_DATA)
Get /sapi/v1/sol-staking/sol/quota

https://developers.binance.com/docs/staking/sol-staking/account/Get-SOL-staking-quota-details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetSolStakingQuotaDetailsRequest
*/
func (a *SolStakingAPIService) GetSolStakingQuotaDetails(ctx context.Context) ApiGetSolStakingQuotaDetailsRequest {
	return ApiGetSolStakingQuotaDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetSolStakingQuotaDetailsResponse
func (a *SolStakingAPIService) GetSolStakingQuotaDetailsExecute(r ApiGetSolStakingQuotaDetailsRequest) (*common.RestApiResponse[models.GetSolStakingQuotaDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/quota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetSolStakingQuotaDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUnclaimedRewardsRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	recvWindow *int64
}

func (r ApiGetUnclaimedRewardsRequest) RecvWindow(recvWindow int64) ApiGetUnclaimedRewardsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUnclaimedRewardsRequest) Execute() (*common.RestApiResponse[models.GetUnclaimedRewardsResponse], error) {
	return r.ApiService.GetUnclaimedRewardsExecute(r)
}

/*
GetUnclaimedRewards Get Unclaimed Rewards(USER_DATA)
Get /sapi/v1/sol-staking/sol/history/unclaimedRewards

https://developers.binance.com/docs/staking/sol-staking/history/Get-Unclaimed-Rewards

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetUnclaimedRewardsRequest
*/
func (a *SolStakingAPIService) GetUnclaimedRewards(ctx context.Context) ApiGetUnclaimedRewardsRequest {
	return ApiGetUnclaimedRewardsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUnclaimedRewardsResponse
func (a *SolStakingAPIService) GetUnclaimedRewardsExecute(r ApiGetUnclaimedRewardsRequest) (*common.RestApiResponse[models.GetUnclaimedRewardsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/history/unclaimedRewards"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUnclaimedRewardsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemSolRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	amount     *float32
	recvWindow *int64
}

// Amount in SOL.
func (r ApiRedeemSolRequest) Amount(amount float32) ApiRedeemSolRequest {
	r.amount = &amount
	return r
}

func (r ApiRedeemSolRequest) RecvWindow(recvWindow int64) ApiRedeemSolRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemSolRequest) Execute() (*common.RestApiResponse[models.RedeemSolResponse], error) {
	return r.ApiService.RedeemSolExecute(r)
}

/*
RedeemSol Redeem SOL(TRADE)
Post /sapi/v1/sol-staking/sol/redeem

https://developers.binance.com/docs/staking/sol-staking/staking/Redeem-SOL

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in SOL.
@param recvWindow -
@return ApiRedeemSolRequest
*/
func (a *SolStakingAPIService) RedeemSol(ctx context.Context) ApiRedeemSolRequest {
	return ApiRedeemSolRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemSolResponse
func (a *SolStakingAPIService) RedeemSolExecute(r ApiRedeemSolRequest) (*common.RestApiResponse[models.RedeemSolResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemSolResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSolStakingAccountRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	recvWindow *int64
}

func (r ApiSolStakingAccountRequest) RecvWindow(recvWindow int64) ApiSolStakingAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSolStakingAccountRequest) Execute() (*common.RestApiResponse[models.SolStakingAccountResponse], error) {
	return r.ApiService.SolStakingAccountExecute(r)
}

/*
SolStakingAccount SOL Staking account(USER_DATA)
Get /sapi/v1/sol-staking/account

https://developers.binance.com/docs/staking/sol-staking/account/SOL-Staking-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiSolStakingAccountRequest
*/
func (a *SolStakingAPIService) SolStakingAccount(ctx context.Context) ApiSolStakingAccountRequest {
	return ApiSolStakingAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SolStakingAccountResponse
func (a *SolStakingAPIService) SolStakingAccountExecute(r ApiSolStakingAccountRequest) (*common.RestApiResponse[models.SolStakingAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SolStakingAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeSolStakingRequest struct {
	ctx        context.Context
	ApiService *SolStakingAPIService
	amount     *float32
	recvWindow *int64
}

// Amount in SOL.
func (r ApiSubscribeSolStakingRequest) Amount(amount float32) ApiSubscribeSolStakingRequest {
	r.amount = &amount
	return r
}

func (r ApiSubscribeSolStakingRequest) RecvWindow(recvWindow int64) ApiSubscribeSolStakingRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeSolStakingRequest) Execute() (*common.RestApiResponse[models.SubscribeSolStakingResponse], error) {
	return r.ApiService.SubscribeSolStakingExecute(r)
}

/*
SubscribeSolStaking Subscribe SOL Staking(TRADE)
Post /sapi/v1/sol-staking/sol/stake

https://developers.binance.com/docs/staking/sol-staking/staking/Subscribe-SOL-Staking

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in SOL.
@param recvWindow -
@return ApiSubscribeSolStakingRequest
*/
func (a *SolStakingAPIService) SubscribeSolStaking(ctx context.Context) ApiSubscribeSolStakingRequest {
	return ApiSubscribeSolStakingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeSolStakingResponse
func (a *SolStakingAPIService) SubscribeSolStakingExecute(r ApiSubscribeSolStakingRequest) (*common.RestApiResponse[models.SubscribeSolStakingResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sol-staking/sol/stake"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeSolStakingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
