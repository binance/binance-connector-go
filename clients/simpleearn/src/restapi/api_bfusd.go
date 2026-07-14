/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package binancesimpleearnrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// BfusdAPIService BfusdAPI Service
type BfusdAPIService Service

type ApiGetBfusdAccountRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	recvWindow *int64
}

func (r ApiGetBfusdAccountRequest) RecvWindow(recvWindow int64) ApiGetBfusdAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdAccountRequest) Execute() (*common.RestApiResponse[models.GetBfusdAccountResponse], error) {
	return r.ApiService.GetBfusdAccountExecute(r)
}

/*
GetBfusdAccount Get BFUSD Account (USER_DATA)
Get /sapi/v1/bfusd/account

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetBfusdAccountRequest
*/
func (a *BfusdAPIService) GetBfusdAccount(ctx context.Context) ApiGetBfusdAccountRequest {
	return ApiGetBfusdAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdAccountResponse
func (a *BfusdAPIService) GetBfusdAccountExecute(r ApiGetBfusdAccountRequest) (*common.RestApiResponse[models.GetBfusdAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetBfusdAccountResponse](
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

type ApiGetBfusdQuotaDetailsRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	recvWindow *int64
}

func (r ApiGetBfusdQuotaDetailsRequest) RecvWindow(recvWindow int64) ApiGetBfusdQuotaDetailsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdQuotaDetailsRequest) Execute() (*common.RestApiResponse[models.GetBfusdQuotaDetailsResponse], error) {
	return r.ApiService.GetBfusdQuotaDetailsExecute(r)
}

/*
GetBfusdQuotaDetails Get BFUSD Quota Details (USER_DATA)
Get /sapi/v1/bfusd/quota

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-quota-details

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetBfusdQuotaDetailsRequest
*/
func (a *BfusdAPIService) GetBfusdQuotaDetails(ctx context.Context) ApiGetBfusdQuotaDetailsRequest {
	return ApiGetBfusdQuotaDetailsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdQuotaDetailsResponse
func (a *BfusdAPIService) GetBfusdQuotaDetailsExecute(r ApiGetBfusdQuotaDetailsRequest) (*common.RestApiResponse[models.GetBfusdQuotaDetailsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/quota"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetBfusdQuotaDetailsResponse](
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

type ApiGetBfusdRateHistoryRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBfusdRateHistoryRequest) StartTime(startTime int64) ApiGetBfusdRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBfusdRateHistoryRequest) EndTime(endTime int64) ApiGetBfusdRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Starts from 1.
func (r ApiGetBfusdRateHistoryRequest) Current(current int64) ApiGetBfusdRateHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetBfusdRateHistoryRequest) Size(size int64) ApiGetBfusdRateHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBfusdRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetBfusdRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetBfusdRateHistoryResponse], error) {
	return r.ApiService.GetBfusdRateHistoryExecute(r)
}

/*
GetBfusdRateHistory Get BFUSD Rate History (USER_DATA)
Get /sapi/v1/bfusd/history/rateHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-rate-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Starts from 1.
@param size -  Number of results per page.
@param recvWindow -
@return ApiGetBfusdRateHistoryRequest
*/
func (a *BfusdAPIService) GetBfusdRateHistory(ctx context.Context) ApiGetBfusdRateHistoryRequest {
	return ApiGetBfusdRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdRateHistoryResponse
func (a *BfusdAPIService) GetBfusdRateHistoryExecute(r ApiGetBfusdRateHistoryRequest) (*common.RestApiResponse[models.GetBfusdRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/history/rateHistory"

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

	resp, err := SendRequest[models.GetBfusdRateHistoryResponse](
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

type ApiGetBfusdRedemptionHistoryRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBfusdRedemptionHistoryRequest) StartTime(startTime int64) ApiGetBfusdRedemptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBfusdRedemptionHistoryRequest) EndTime(endTime int64) ApiGetBfusdRedemptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page.
func (r ApiGetBfusdRedemptionHistoryRequest) Current(current int64) ApiGetBfusdRedemptionHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetBfusdRedemptionHistoryRequest) Size(size int64) ApiGetBfusdRedemptionHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBfusdRedemptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetBfusdRedemptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdRedemptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetBfusdRedemptionHistoryResponse], error) {
	return r.ApiService.GetBfusdRedemptionHistoryExecute(r)
}

/*
GetBfusdRedemptionHistory Get BFUSD Redemption History (USER_DATA)
Get /sapi/v1/bfusd/history/redemptionHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-redemption-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page.
@param size -  Number of results per page.
@param recvWindow -
@return ApiGetBfusdRedemptionHistoryRequest
*/
func (a *BfusdAPIService) GetBfusdRedemptionHistory(ctx context.Context) ApiGetBfusdRedemptionHistoryRequest {
	return ApiGetBfusdRedemptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdRedemptionHistoryResponse
func (a *BfusdAPIService) GetBfusdRedemptionHistoryExecute(r ApiGetBfusdRedemptionHistoryRequest) (*common.RestApiResponse[models.GetBfusdRedemptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/history/redemptionHistory"

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

	resp, err := SendRequest[models.GetBfusdRedemptionHistoryResponse](
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

type ApiGetBfusdRewardsHistoryRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBfusdRewardsHistoryRequest) StartTime(startTime int64) ApiGetBfusdRewardsHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBfusdRewardsHistoryRequest) EndTime(endTime int64) ApiGetBfusdRewardsHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page.
func (r ApiGetBfusdRewardsHistoryRequest) Current(current int64) ApiGetBfusdRewardsHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetBfusdRewardsHistoryRequest) Size(size int64) ApiGetBfusdRewardsHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBfusdRewardsHistoryRequest) RecvWindow(recvWindow int64) ApiGetBfusdRewardsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdRewardsHistoryRequest) Execute() (*common.RestApiResponse[models.GetBfusdRewardsHistoryResponse], error) {
	return r.ApiService.GetBfusdRewardsHistoryExecute(r)
}

/*
GetBfusdRewardsHistory Get BFUSD Rewards History (USER_DATA)
Get /sapi/v1/bfusd/history/rewardsHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-rewards-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page.
@param size -  Number of results per page.
@param recvWindow -
@return ApiGetBfusdRewardsHistoryRequest
*/
func (a *BfusdAPIService) GetBfusdRewardsHistory(ctx context.Context) ApiGetBfusdRewardsHistoryRequest {
	return ApiGetBfusdRewardsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdRewardsHistoryResponse
func (a *BfusdAPIService) GetBfusdRewardsHistoryExecute(r ApiGetBfusdRewardsHistoryRequest) (*common.RestApiResponse[models.GetBfusdRewardsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/history/rewardsHistory"

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

	resp, err := SendRequest[models.GetBfusdRewardsHistoryResponse](
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

type ApiGetBfusdSubscriptionHistoryRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	asset      *models.GetBfusdSubscriptionHistoryAssetParameter
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiGetBfusdSubscriptionHistoryRequest) Asset(asset models.GetBfusdSubscriptionHistoryAssetParameter) ApiGetBfusdSubscriptionHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiGetBfusdSubscriptionHistoryRequest) StartTime(startTime int64) ApiGetBfusdSubscriptionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetBfusdSubscriptionHistoryRequest) EndTime(endTime int64) ApiGetBfusdSubscriptionHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page.
func (r ApiGetBfusdSubscriptionHistoryRequest) Current(current int64) ApiGetBfusdSubscriptionHistoryRequest {
	r.current = &current
	return r
}

// Number of results per page.
func (r ApiGetBfusdSubscriptionHistoryRequest) Size(size int64) ApiGetBfusdSubscriptionHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetBfusdSubscriptionHistoryRequest) RecvWindow(recvWindow int64) ApiGetBfusdSubscriptionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetBfusdSubscriptionHistoryRequest) Execute() (*common.RestApiResponse[models.GetBfusdSubscriptionHistoryResponse], error) {
	return r.ApiService.GetBfusdSubscriptionHistoryExecute(r)
}

/*
GetBfusdSubscriptionHistory Get BFUSD subscription history (USER_DATA)
Get /sapi/v1/bfusd/history/subscriptionHistory

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#get-bfusd-subscription-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param startTime -
@param endTime -
@param current -  Currently querying page.
@param size -  Number of results per page.
@param recvWindow -
@return ApiGetBfusdSubscriptionHistoryRequest
*/
func (a *BfusdAPIService) GetBfusdSubscriptionHistory(ctx context.Context) ApiGetBfusdSubscriptionHistoryRequest {
	return ApiGetBfusdSubscriptionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetBfusdSubscriptionHistoryResponse
func (a *BfusdAPIService) GetBfusdSubscriptionHistoryExecute(r ApiGetBfusdSubscriptionHistoryRequest) (*common.RestApiResponse[models.GetBfusdSubscriptionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/history/subscriptionHistory"

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

	resp, err := SendRequest[models.GetBfusdSubscriptionHistoryResponse](
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

type ApiRedeemBfusdRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	amount     *float32
	type_      *models.RedeemBfusdTypeParameter
	recvWindow *int64
}

// Amount in BFUSD
func (r ApiRedeemBfusdRequest) Amount(amount float32) ApiRedeemBfusdRequest {
	r.amount = &amount
	return r
}

// defaults to STANDARD
func (r ApiRedeemBfusdRequest) Type(type_ models.RedeemBfusdTypeParameter) ApiRedeemBfusdRequest {
	r.type_ = &type_
	return r
}

// Request validity window in milliseconds.
func (r ApiRedeemBfusdRequest) RecvWindow(recvWindow int64) ApiRedeemBfusdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemBfusdRequest) Execute() (*common.RestApiResponse[models.RedeemBfusdResponse], error) {
	return r.ApiService.RedeemBfusdExecute(r)
}

/*
RedeemBfusd Redeem BFUSD (TRADE)
Post /sapi/v1/bfusd/redeem

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#redeem-bfusd

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -  Amount in BFUSD
@param type_ -  defaults to STANDARD
@param recvWindow -  Request validity window in milliseconds.
@return ApiRedeemBfusdRequest
*/
func (a *BfusdAPIService) RedeemBfusd(ctx context.Context) ApiRedeemBfusdRequest {
	return ApiRedeemBfusdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemBfusdResponse
func (a *BfusdAPIService) RedeemBfusdExecute(r ApiRedeemBfusdRequest) (*common.RestApiResponse[models.RedeemBfusdResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/redeem"

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

	resp, err := SendRequest[models.RedeemBfusdResponse](
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

type ApiSubscribeBfusdRequest struct {
	ctx        context.Context
	ApiService *BfusdAPIService
	asset      *string
	amount     *float32
	recvWindow *int64
}

func (r ApiSubscribeBfusdRequest) Asset(asset string) ApiSubscribeBfusdRequest {
	r.asset = &asset
	return r
}

// Amount
func (r ApiSubscribeBfusdRequest) Amount(amount float32) ApiSubscribeBfusdRequest {
	r.amount = &amount
	return r
}

// Request validity window in milliseconds.
func (r ApiSubscribeBfusdRequest) RecvWindow(recvWindow int64) ApiSubscribeBfusdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeBfusdRequest) Execute() (*common.RestApiResponse[models.SubscribeBfusdResponse], error) {
	return r.ApiService.SubscribeBfusdExecute(r)
}

/*
SubscribeBfusd Subscribe BFUSD (TRADE)
Post /sapi/v1/bfusd/subscribe

https://developers.binance.com/en/docs/catalog/investment-and-services-simple-earn/api/rest-api/bfusd#subscribe-bfusd

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param amount -  Amount
@param recvWindow -  Request validity window in milliseconds.
@return ApiSubscribeBfusdRequest
*/
func (a *BfusdAPIService) SubscribeBfusd(ctx context.Context) ApiSubscribeBfusdRequest {
	return ApiSubscribeBfusdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeBfusdResponse
func (a *BfusdAPIService) SubscribeBfusdExecute(r ApiSubscribeBfusdRequest) (*common.RestApiResponse[models.SubscribeBfusdResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bfusd/subscribe"

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

	resp, err := SendRequest[models.SubscribeBfusdResponse](
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
