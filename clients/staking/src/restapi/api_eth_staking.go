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
	"github.com/binance/binance-connector-go/common/v2/common"
)

// EthStakingAPIService EthStakingAPI Service
type EthStakingAPIService Service

type ApiEthStakingAccountRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	recvWindow *int64
}

func (r ApiEthStakingAccountRequest) RecvWindow(recvWindow int64) ApiEthStakingAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEthStakingAccountRequest) Execute() (*common.RestApiResponse[models.EthStakingAccountResponse], error) {
	return r.ApiService.EthStakingAccountExecute(r)
}

/*
EthStakingAccount ETH Staking account(USER_DATA)
Get /sapi/v2/eth-staking/account

https://developers.binance.com/docs/staking/eth-staking/account/ETH-Staking-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiEthStakingAccountRequest
*/
func (a *EthStakingAPIService) EthStakingAccount(ctx context.Context) ApiEthStakingAccountRequest {
	return ApiEthStakingAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EthStakingAccountResponse
func (a *EthStakingAPIService) EthStakingAccountExecute(r ApiEthStakingAccountRequest) (*common.RestApiResponse[models.EthStakingAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/eth-staking/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EthStakingAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCurrentEthStakingQuotaRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	recvWindow *int64
}

func (r ApiGetCurrentEthStakingQuotaRequest) RecvWindow(recvWindow int64) ApiGetCurrentEthStakingQuotaRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCurrentEthStakingQuotaRequest) Execute() (*common.RestApiResponse[models.GetCurrentEthStakingQuotaResponse], error) {
	return r.ApiService.GetCurrentEthStakingQuotaExecute(r)
}

/*
GetCurrentEthStakingQuota Get current ETH staking quota(USER_DATA)
Get /sapi/v1/eth-staking/eth/quota

https://developers.binance.com/docs/staking/eth-staking/account/Get-current-ETH-staking-quota

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetCurrentEthStakingQuotaRequest
*/
func (a *EthStakingAPIService) GetCurrentEthStakingQuota(ctx context.Context) ApiGetCurrentEthStakingQuotaRequest {
	return ApiGetCurrentEthStakingQuotaRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCurrentEthStakingQuotaResponse
func (a *EthStakingAPIService) GetCurrentEthStakingQuotaExecute(r ApiGetCurrentEthStakingQuotaRequest) (*common.RestApiResponse[models.GetCurrentEthStakingQuotaResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/quota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCurrentEthStakingQuotaResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetEthRedemptionHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetEthRedemptionHistoryRequest) StartTime(startTime int64) ApiGetEthRedemptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetEthRedemptionHistoryRequest) EndTime(endTime int64) ApiGetEthRedemptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetEthRedemptionHistoryRequest) Current(current int64) ApiGetEthRedemptionHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetEthRedemptionHistoryRequest) Size(size int64) ApiGetEthRedemptionHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetEthRedemptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetEthRedemptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetEthRedemptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetEthRedemptionHistoryResponse], error) {
	return r.ApiService.GetEthRedemptionHistoryExecute(r)
}

/*
GetEthRedemptionHistory Get ETH redemption history(USER_DATA)
Get /sapi/v1/eth-staking/eth/history/redemptionHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-ETH-redemption-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetEthRedemptionHistoryRequest
*/
func (a *EthStakingAPIService) GetEthRedemptionHistory(ctx context.Context) ApiGetEthRedemptionHistoryRequest {
	return ApiGetEthRedemptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEthRedemptionHistoryResponse
func (a *EthStakingAPIService) GetEthRedemptionHistoryExecute(r ApiGetEthRedemptionHistoryRequest) (*common.RestApiResponse[models.GetEthRedemptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/history/redemptionHistory"

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

	resp, err := SendRequest[models.GetEthRedemptionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetEthStakingHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetEthStakingHistoryRequest) StartTime(startTime int64) ApiGetEthStakingHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetEthStakingHistoryRequest) EndTime(endTime int64) ApiGetEthStakingHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetEthStakingHistoryRequest) Current(current int64) ApiGetEthStakingHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetEthStakingHistoryRequest) Size(size int64) ApiGetEthStakingHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetEthStakingHistoryRequest) RecvWindow(recvWindow int64) ApiGetEthStakingHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetEthStakingHistoryRequest) Execute() (*common.RestApiResponse[models.GetEthStakingHistoryResponse], error) {
	return r.ApiService.GetEthStakingHistoryExecute(r)
}

/*
GetEthStakingHistory Get ETH staking history(USER_DATA)
Get /sapi/v1/eth-staking/eth/history/stakingHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-ETH-staking-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetEthStakingHistoryRequest
*/
func (a *EthStakingAPIService) GetEthStakingHistory(ctx context.Context) ApiGetEthStakingHistoryRequest {
	return ApiGetEthStakingHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetEthStakingHistoryResponse
func (a *EthStakingAPIService) GetEthStakingHistoryExecute(r ApiGetEthStakingHistoryRequest) (*common.RestApiResponse[models.GetEthStakingHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/history/stakingHistory"

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

	resp, err := SendRequest[models.GetEthStakingHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetWbethRateHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetWbethRateHistoryRequest) StartTime(startTime int64) ApiGetWbethRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetWbethRateHistoryRequest) EndTime(endTime int64) ApiGetWbethRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetWbethRateHistoryRequest) Current(current int64) ApiGetWbethRateHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetWbethRateHistoryRequest) Size(size int64) ApiGetWbethRateHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetWbethRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetWbethRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetWbethRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetWbethRateHistoryResponse], error) {
	return r.ApiService.GetWbethRateHistoryExecute(r)
}

/*
GetWbethRateHistory Get WBETH Rate History(USER_DATA)
Get /sapi/v1/eth-staking/eth/history/rateHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-BETH-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetWbethRateHistoryRequest
*/
func (a *EthStakingAPIService) GetWbethRateHistory(ctx context.Context) ApiGetWbethRateHistoryRequest {
	return ApiGetWbethRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetWbethRateHistoryResponse
func (a *EthStakingAPIService) GetWbethRateHistoryExecute(r ApiGetWbethRateHistoryRequest) (*common.RestApiResponse[models.GetWbethRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/history/rateHistory"

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

	resp, err := SendRequest[models.GetWbethRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetWbethRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetWbethRewardsHistoryRequest) StartTime(startTime int64) ApiGetWbethRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetWbethRewardsHistoryRequest) EndTime(endTime int64) ApiGetWbethRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetWbethRewardsHistoryRequest) Current(current int64) ApiGetWbethRewardsHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetWbethRewardsHistoryRequest) Size(size int64) ApiGetWbethRewardsHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetWbethRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetWbethRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetWbethRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetWbethRewardsHistoryResponse], error) {
	return r.ApiService.GetWbethRewardsHistoryExecute(r)
}

/*
GetWbethRewardsHistory Get WBETH rewards history(USER_DATA)
Get /sapi/v1/eth-staking/eth/history/wbethRewardsHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-WBETH-rewards-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetWbethRewardsHistoryRequest
*/
func (a *EthStakingAPIService) GetWbethRewardsHistory(ctx context.Context) ApiGetWbethRewardsHistoryRequest {
	return ApiGetWbethRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetWbethRewardsHistoryResponse
func (a *EthStakingAPIService) GetWbethRewardsHistoryExecute(r ApiGetWbethRewardsHistoryRequest) (*common.RestApiResponse[models.GetWbethRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/history/wbethRewardsHistory"

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

	resp, err := SendRequest[models.GetWbethRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetWbethUnwrapHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetWbethUnwrapHistoryRequest) StartTime(startTime int64) ApiGetWbethUnwrapHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetWbethUnwrapHistoryRequest) EndTime(endTime int64) ApiGetWbethUnwrapHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetWbethUnwrapHistoryRequest) Current(current int64) ApiGetWbethUnwrapHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetWbethUnwrapHistoryRequest) Size(size int64) ApiGetWbethUnwrapHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetWbethUnwrapHistoryRequest) RecvWindow(recvWindow int64) ApiGetWbethUnwrapHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetWbethUnwrapHistoryRequest) Execute() (*common.RestApiResponse[models.GetWbethUnwrapHistoryResponse], error) {
	return r.ApiService.GetWbethUnwrapHistoryExecute(r)
}

/*
GetWbethUnwrapHistory Get WBETH unwrap history(USER_DATA)
Get /sapi/v1/eth-staking/wbeth/history/unwrapHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-WBETH-unwrap-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetWbethUnwrapHistoryRequest
*/
func (a *EthStakingAPIService) GetWbethUnwrapHistory(ctx context.Context) ApiGetWbethUnwrapHistoryRequest {
	return ApiGetWbethUnwrapHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetWbethUnwrapHistoryResponse
func (a *EthStakingAPIService) GetWbethUnwrapHistoryExecute(r ApiGetWbethUnwrapHistoryRequest) (*common.RestApiResponse[models.GetWbethUnwrapHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/wbeth/history/unwrapHistory"

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

	resp, err := SendRequest[models.GetWbethUnwrapHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetWbethWrapHistoryRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetWbethWrapHistoryRequest) StartTime(startTime int64) ApiGetWbethWrapHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetWbethWrapHistoryRequest) EndTime(endTime int64) ApiGetWbethWrapHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetWbethWrapHistoryRequest) Current(current int64) ApiGetWbethWrapHistoryRequest {
	r.current = &current
	return r
}

// Default:10, Max:100
func (r ApiGetWbethWrapHistoryRequest) Size(size int64) ApiGetWbethWrapHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetWbethWrapHistoryRequest) RecvWindow(recvWindow int64) ApiGetWbethWrapHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetWbethWrapHistoryRequest) Execute() (*common.RestApiResponse[models.GetWbethWrapHistoryResponse], error) {
	return r.ApiService.GetWbethWrapHistoryExecute(r)
}

/*
GetWbethWrapHistory Get WBETH wrap history(USER_DATA)
Get /sapi/v1/eth-staking/wbeth/history/wrapHistory

https://developers.binance.com/docs/staking/eth-staking/history/Get-WBETH-wrap-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10, Max:100
@param recvWindow -
@return ApiGetWbethWrapHistoryRequest
*/
func (a *EthStakingAPIService) GetWbethWrapHistory(ctx context.Context) ApiGetWbethWrapHistoryRequest {
	return ApiGetWbethWrapHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetWbethWrapHistoryResponse
func (a *EthStakingAPIService) GetWbethWrapHistoryExecute(r ApiGetWbethWrapHistoryRequest) (*common.RestApiResponse[models.GetWbethWrapHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/wbeth/history/wrapHistory"

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

	resp, err := SendRequest[models.GetWbethWrapHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemEthRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	amount     *float32
	asset      *string
	recvWindow *int64
}

// Amount in SOL.
func (r ApiRedeemEthRequest) Amount(amount float32) ApiRedeemEthRequest {
	r.amount = &amount
	return r
}

// WBETH or BETH, default to BETH
func (r ApiRedeemEthRequest) Asset(asset string) ApiRedeemEthRequest {
	r.asset = &asset
	return r
}

func (r ApiRedeemEthRequest) RecvWindow(recvWindow int64) ApiRedeemEthRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemEthRequest) Execute() (*common.RestApiResponse[models.RedeemEthResponse], error) {
	return r.ApiService.RedeemEthExecute(r)
}

/*
RedeemEth Redeem ETH(TRADE)
Post /sapi/v1/eth-staking/eth/redeem

https://developers.binance.com/docs/staking/eth-staking/staking/Redeem-ETH

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in SOL.
@param asset -  WBETH or BETH, default to BETH
@param recvWindow -
@return ApiRedeemEthRequest
*/
func (a *EthStakingAPIService) RedeemEth(ctx context.Context) ApiRedeemEthRequest {
	return ApiRedeemEthRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemEthResponse
func (a *EthStakingAPIService) RedeemEthExecute(r ApiRedeemEthRequest) (*common.RestApiResponse[models.RedeemEthResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/eth/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemEthResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeEthStakingRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	amount     *float32
	recvWindow *int64
}

// Amount in SOL.
func (r ApiSubscribeEthStakingRequest) Amount(amount float32) ApiSubscribeEthStakingRequest {
	r.amount = &amount
	return r
}

func (r ApiSubscribeEthStakingRequest) RecvWindow(recvWindow int64) ApiSubscribeEthStakingRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeEthStakingRequest) Execute() (*common.RestApiResponse[models.SubscribeEthStakingResponse], error) {
	return r.ApiService.SubscribeEthStakingExecute(r)
}

/*
SubscribeEthStaking Subscribe ETH Staking(TRADE)
Post /sapi/v2/eth-staking/eth/stake

https://developers.binance.com/docs/staking/eth-staking/staking/Subscribe-ETH-Staking

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in SOL.
@param recvWindow -
@return ApiSubscribeEthStakingRequest
*/
func (a *EthStakingAPIService) SubscribeEthStaking(ctx context.Context) ApiSubscribeEthStakingRequest {
	return ApiSubscribeEthStakingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeEthStakingResponse
func (a *EthStakingAPIService) SubscribeEthStakingExecute(r ApiSubscribeEthStakingRequest) (*common.RestApiResponse[models.SubscribeEthStakingResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/eth-staking/eth/stake"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeEthStakingResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWrapBethRequest struct {
	ctx        context.Context
	ApiService *EthStakingAPIService
	amount     *float32
	recvWindow *int64
}

// Amount in SOL.
func (r ApiWrapBethRequest) Amount(amount float32) ApiWrapBethRequest {
	r.amount = &amount
	return r
}

func (r ApiWrapBethRequest) RecvWindow(recvWindow int64) ApiWrapBethRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWrapBethRequest) Execute() (*common.RestApiResponse[models.WrapBethResponse], error) {
	return r.ApiService.WrapBethExecute(r)
}

/*
WrapBeth Wrap BETH(TRADE)
Post /sapi/v1/eth-staking/wbeth/wrap

https://developers.binance.com/docs/staking/eth-staking/staking/Wrap-BETH

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in SOL.
@param recvWindow -
@return ApiWrapBethRequest
*/
func (a *EthStakingAPIService) WrapBeth(ctx context.Context) ApiWrapBethRequest {
	return ApiWrapBethRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WrapBethResponse
func (a *EthStakingAPIService) WrapBethExecute(r ApiWrapBethRequest) (*common.RestApiResponse[models.WrapBethResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/eth-staking/wbeth/wrap"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WrapBethResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
