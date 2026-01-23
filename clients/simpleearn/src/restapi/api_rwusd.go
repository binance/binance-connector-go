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

// RwusdAPIService RwusdAPI Service
type RwusdAPIService Service

type ApiGetRwusdAccountRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdAccountRequest) RecvWindow(recvWindow int64) ApiGetRwusdAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdAccountRequest) Execute() (*common.RestApiResponse[models.GetRwusdAccountResponse], error) {
	return r.ApiService.GetRwusdAccountExecute(r)
}

/*
GetRwusdAccount Get RWUSD Account (USER_DATA)
Get /sapi/v1/rwusd/account

https://developers.binance.com/docs/simple_earn/rwusd/account/

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdAccountRequest
*/
func (a *RwusdAPIService) GetRwusdAccount(ctx context.Context) ApiGetRwusdAccountRequest {
	return ApiGetRwusdAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdAccountResponse
func (a *RwusdAPIService) GetRwusdAccountExecute(r ApiGetRwusdAccountRequest) (*common.RestApiResponse[models.GetRwusdAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetRwusdAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRwusdQuotaDetailsRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdQuotaDetailsRequest) RecvWindow(recvWindow int64) ApiGetRwusdQuotaDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdQuotaDetailsRequest) Execute() (*common.RestApiResponse[models.GetRwusdQuotaDetailsResponse], error) {
	return r.ApiService.GetRwusdQuotaDetailsExecute(r)
}

/*
GetRwusdQuotaDetails Get RWUSD Quota Details (USER_DATA)
Get /sapi/v1/rwusd/quota

https://developers.binance.com/docs/simple_earn/rwusd/account/Get-RWUSD-Quota-Details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdQuotaDetailsRequest
*/
func (a *RwusdAPIService) GetRwusdQuotaDetails(ctx context.Context) ApiGetRwusdQuotaDetailsRequest {
	return ApiGetRwusdQuotaDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdQuotaDetailsResponse
func (a *RwusdAPIService) GetRwusdQuotaDetailsExecute(r ApiGetRwusdQuotaDetailsRequest) (*common.RestApiResponse[models.GetRwusdQuotaDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/quota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetRwusdQuotaDetailsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRwusdRateHistoryRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetRwusdRateHistoryRequest) StartTime(startTime int64) ApiGetRwusdRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetRwusdRateHistoryRequest) EndTime(endTime int64) ApiGetRwusdRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetRwusdRateHistoryRequest) Current(current int64) ApiGetRwusdRateHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetRwusdRateHistoryRequest) Size(size int64) ApiGetRwusdRateHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetRwusdRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetRwusdRateHistoryResponse], error) {
	return r.ApiService.GetRwusdRateHistoryExecute(r)
}

/*
GetRwusdRateHistory Get RWUSD Rate History (USER_DATA)
Get /sapi/v1/rwusd/history/rateHistory

https://developers.binance.com/docs/simple_earn/rwusd/history/Get-RWUSD-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdRateHistoryRequest
*/
func (a *RwusdAPIService) GetRwusdRateHistory(ctx context.Context) ApiGetRwusdRateHistoryRequest {
	return ApiGetRwusdRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdRateHistoryResponse
func (a *RwusdAPIService) GetRwusdRateHistoryExecute(r ApiGetRwusdRateHistoryRequest) (*common.RestApiResponse[models.GetRwusdRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/history/rateHistory"

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

	resp, err := SendRequest[models.GetRwusdRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRwusdRedemptionHistoryRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetRwusdRedemptionHistoryRequest) StartTime(startTime int64) ApiGetRwusdRedemptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetRwusdRedemptionHistoryRequest) EndTime(endTime int64) ApiGetRwusdRedemptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetRwusdRedemptionHistoryRequest) Current(current int64) ApiGetRwusdRedemptionHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetRwusdRedemptionHistoryRequest) Size(size int64) ApiGetRwusdRedemptionHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdRedemptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetRwusdRedemptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdRedemptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetRwusdRedemptionHistoryResponse], error) {
	return r.ApiService.GetRwusdRedemptionHistoryExecute(r)
}

/*
GetRwusdRedemptionHistory Get RWUSD Redemption History (USER_DATA)
Get /sapi/v1/rwusd/history/redemptionHistory

https://developers.binance.com/docs/simple_earn/rwusd/history/Get-RWUSD-Redemption-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdRedemptionHistoryRequest
*/
func (a *RwusdAPIService) GetRwusdRedemptionHistory(ctx context.Context) ApiGetRwusdRedemptionHistoryRequest {
	return ApiGetRwusdRedemptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdRedemptionHistoryResponse
func (a *RwusdAPIService) GetRwusdRedemptionHistoryExecute(r ApiGetRwusdRedemptionHistoryRequest) (*common.RestApiResponse[models.GetRwusdRedemptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/history/redemptionHistory"

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

	resp, err := SendRequest[models.GetRwusdRedemptionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRwusdRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetRwusdRewardsHistoryRequest) StartTime(startTime int64) ApiGetRwusdRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetRwusdRewardsHistoryRequest) EndTime(endTime int64) ApiGetRwusdRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetRwusdRewardsHistoryRequest) Current(current int64) ApiGetRwusdRewardsHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetRwusdRewardsHistoryRequest) Size(size int64) ApiGetRwusdRewardsHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetRwusdRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetRwusdRewardsHistoryResponse], error) {
	return r.ApiService.GetRwusdRewardsHistoryExecute(r)
}

/*
GetRwusdRewardsHistory Get RWUSD Rewards History (USER_DATA)
Get /sapi/v1/rwusd/history/rewardsHistory

https://developers.binance.com/docs/simple_earn/rwusd/history/Get-RWUSD-Rewards-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdRewardsHistoryRequest
*/
func (a *RwusdAPIService) GetRwusdRewardsHistory(ctx context.Context) ApiGetRwusdRewardsHistoryRequest {
	return ApiGetRwusdRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdRewardsHistoryResponse
func (a *RwusdAPIService) GetRwusdRewardsHistoryExecute(r ApiGetRwusdRewardsHistoryRequest) (*common.RestApiResponse[models.GetRwusdRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/history/rewardsHistory"

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

	resp, err := SendRequest[models.GetRwusdRewardsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetRwusdSubscriptionHistoryRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

// USDC or USDT
func (r ApiGetRwusdSubscriptionHistoryRequest) Asset(asset string) ApiGetRwusdSubscriptionHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetRwusdSubscriptionHistoryRequest) StartTime(startTime int64) ApiGetRwusdSubscriptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetRwusdSubscriptionHistoryRequest) EndTime(endTime int64) ApiGetRwusdSubscriptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1. Default: 1
func (r ApiGetRwusdSubscriptionHistoryRequest) Current(current int64) ApiGetRwusdSubscriptionHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page. Default: 10, Max: 100
func (r ApiGetRwusdSubscriptionHistoryRequest) Size(size int64) ApiGetRwusdSubscriptionHistoryRequest {
	r.size = &size
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiGetRwusdSubscriptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetRwusdSubscriptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRwusdSubscriptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetRwusdSubscriptionHistoryResponse], error) {
	return r.ApiService.GetRwusdSubscriptionHistoryExecute(r)
}

/*
GetRwusdSubscriptionHistory Get RWUSD subscription history(USER_DATA)
Get /sapi/v1/rwusd/history/subscriptionHistory

https://developers.binance.com/docs/simple_earn/rwusd/history/Get-RWUSD-subscription-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDC or USDT
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1. Default: 1
@param size -  Number of results per page. Default: 10, Max: 100
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiGetRwusdSubscriptionHistoryRequest
*/
func (a *RwusdAPIService) GetRwusdSubscriptionHistory(ctx context.Context) ApiGetRwusdSubscriptionHistoryRequest {
	return ApiGetRwusdSubscriptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetRwusdSubscriptionHistoryResponse
func (a *RwusdAPIService) GetRwusdSubscriptionHistoryExecute(r ApiGetRwusdSubscriptionHistoryRequest) (*common.RestApiResponse[models.GetRwusdSubscriptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/history/subscriptionHistory"

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

	resp, err := SendRequest[models.GetRwusdSubscriptionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemRwusdRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	amount     *float32
	type_      *string
	recvWindow *int64
}

// Amount
func (r ApiRedeemRwusdRequest) Amount(amount float32) ApiRedeemRwusdRequest {
	r.amount = &amount
	return r
}

// FAST or STANDARD, defaults to STANDARD
func (r ApiRedeemRwusdRequest) Type(type_ string) ApiRedeemRwusdRequest {
	r.type_ = &type_
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiRedeemRwusdRequest) RecvWindow(recvWindow int64) ApiRedeemRwusdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemRwusdRequest) Execute() (*common.RestApiResponse[models.RedeemRwusdResponse], error) {
	return r.ApiService.RedeemRwusdExecute(r)
}

/*
RedeemRwusd Redeem RWUSD(TRADE)
Post /sapi/v1/rwusd/redeem

https://developers.binance.com/docs/simple_earn/rwusd/earn/Redeem-RWUSD

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount
@param type_ -  FAST or STANDARD, defaults to STANDARD
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiRedeemRwusdRequest
*/
func (a *RwusdAPIService) RedeemRwusd(ctx context.Context) ApiRedeemRwusdRequest {
	return ApiRedeemRwusdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemRwusdResponse
func (a *RwusdAPIService) RedeemRwusdExecute(r ApiRedeemRwusdRequest) (*common.RestApiResponse[models.RedeemRwusdResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/redeem"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemRwusdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeRwusdRequest struct {
	ctx        context.Context
	ApiService *RwusdAPIService
	asset      *string
	amount     *float32
	recvWindow *int64
}

// USDT or USDC (whichever is eligible)
func (r ApiSubscribeRwusdRequest) Asset(asset string) ApiSubscribeRwusdRequest {
	r.asset = &asset
	return r
}

// Amount
func (r ApiSubscribeRwusdRequest) Amount(amount float32) ApiSubscribeRwusdRequest {
	r.amount = &amount
	return r
}

// The value cannot be greater than 60000 (ms)
func (r ApiSubscribeRwusdRequest) RecvWindow(recvWindow int64) ApiSubscribeRwusdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeRwusdRequest) Execute() (*common.RestApiResponse[models.SubscribeRwusdResponse], error) {
	return r.ApiService.SubscribeRwusdExecute(r)
}

/*
SubscribeRwusd Subscribe RWUSD(TRADE)
Post /sapi/v1/rwusd/subscribe

https://developers.binance.com/docs/simple_earn/rwusd/earn/Subscribe-RWUSD

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  USDT or USDC (whichever is eligible)
@param amount -  Amount
@param recvWindow -  The value cannot be greater than 60000 (ms)
@return ApiSubscribeRwusdRequest
*/
func (a *RwusdAPIService) SubscribeRwusd(ctx context.Context) ApiSubscribeRwusdRequest {
	return ApiSubscribeRwusdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeRwusdResponse
func (a *RwusdAPIService) SubscribeRwusdExecute(r ApiSubscribeRwusdRequest) (*common.RestApiResponse[models.SubscribeRwusdResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/rwusd/subscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeRwusdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
