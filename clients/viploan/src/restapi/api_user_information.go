/*
VIP Loan REST API

Access over-collateralized loan services, manage positions, and monitor collateral via the VIP Loan API.
*/

package binanceviploanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserInformationAPIService UserInformationAPI Service
type UserInformationAPIService Service

type ApiCheckVIPLoanCollateralAccountRequest struct {
	ctx                 context.Context
	ApiService          *UserInformationAPIService
	orderId             *int64
	collateralAccountId *int64
	recvWindow          *int64
}

func (r ApiCheckVIPLoanCollateralAccountRequest) OrderId(orderId int64) ApiCheckVIPLoanCollateralAccountRequest {
	r.orderId = &orderId
	return r
}

func (r ApiCheckVIPLoanCollateralAccountRequest) CollateralAccountId(collateralAccountId int64) ApiCheckVIPLoanCollateralAccountRequest {
	r.collateralAccountId = &collateralAccountId
	return r
}

func (r ApiCheckVIPLoanCollateralAccountRequest) RecvWindow(recvWindow int64) ApiCheckVIPLoanCollateralAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCheckVIPLoanCollateralAccountRequest) Execute() (*common.RestApiResponse[models.CheckVIPLoanCollateralAccountResponse], error) {
	return r.ApiService.CheckVIPLoanCollateralAccountExecute(r)
}

/*
CheckVIPLoanCollateralAccount Check VIP Loan Collateral Account (USER_DATA)
Get /sapi/v1/loan/vip/collateral/account

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/user-information#check-viploan-collateral-account

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param collateralAccountId -
@param recvWindow -
@return ApiCheckVIPLoanCollateralAccountRequest
*/
func (a *UserInformationAPIService) CheckVIPLoanCollateralAccount(ctx context.Context) ApiCheckVIPLoanCollateralAccountRequest {
	return ApiCheckVIPLoanCollateralAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CheckVIPLoanCollateralAccountResponse
func (a *UserInformationAPIService) CheckVIPLoanCollateralAccountExecute(r ApiCheckVIPLoanCollateralAccountRequest) (*common.RestApiResponse[models.CheckVIPLoanCollateralAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/collateral/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.collateralAccountId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralAccountId", r.collateralAccountId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CheckVIPLoanCollateralAccountResponse](
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

type ApiGetVIPLoanAccruedInterestRequest struct {
	ctx        context.Context
	ApiService *UserInformationAPIService
	orderId    *int64
	loanCoin   *string
	startTime  *int64
	endTime    *int64
	current    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetVIPLoanAccruedInterestRequest) OrderId(orderId int64) ApiGetVIPLoanAccruedInterestRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetVIPLoanAccruedInterestRequest) LoanCoin(loanCoin string) ApiGetVIPLoanAccruedInterestRequest {
	r.loanCoin = &loanCoin
	return r
}

// If both startTime and endTime are omitted, the most recent 90 days are returned.
func (r ApiGetVIPLoanAccruedInterestRequest) StartTime(startTime int64) ApiGetVIPLoanAccruedInterestRequest {
	r.startTime = &startTime
	return r
}

// Maximum interval between startTime and endTime is 90 days.
func (r ApiGetVIPLoanAccruedInterestRequest) EndTime(endTime int64) ApiGetVIPLoanAccruedInterestRequest {
	r.endTime = &endTime
	return r
}

// Current page number, starting from 1.
func (r ApiGetVIPLoanAccruedInterestRequest) Current(current int64) ApiGetVIPLoanAccruedInterestRequest {
	r.current = &current
	return r
}

// Number of records per page.
func (r ApiGetVIPLoanAccruedInterestRequest) Limit(limit int64) ApiGetVIPLoanAccruedInterestRequest {
	r.limit = &limit
	return r
}

func (r ApiGetVIPLoanAccruedInterestRequest) RecvWindow(recvWindow int64) ApiGetVIPLoanAccruedInterestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetVIPLoanAccruedInterestRequest) Execute() (*common.RestApiResponse[models.GetVIPLoanAccruedInterestResponse], error) {
	return r.ApiService.GetVIPLoanAccruedInterestExecute(r)
}

/*
GetVIPLoanAccruedInterest Get VIP Loan Accrued Interest (USER_DATA)
Get /sapi/v1/loan/vip/accruedInterest

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/user-information#get-viploan-accrued-interest

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param loanCoin -
@param startTime -  If both startTime and endTime are omitted, the most recent 90 days are returned.
@param endTime -  Maximum interval between startTime and endTime is 90 days.
@param current -  Current page number, starting from 1.
@param limit -  Number of records per page.
@param recvWindow -
@return ApiGetVIPLoanAccruedInterestRequest
*/
func (a *UserInformationAPIService) GetVIPLoanAccruedInterest(ctx context.Context) ApiGetVIPLoanAccruedInterestRequest {
	return ApiGetVIPLoanAccruedInterestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetVIPLoanAccruedInterestResponse
func (a *UserInformationAPIService) GetVIPLoanAccruedInterestExecute(r ApiGetVIPLoanAccruedInterestRequest) (*common.RestApiResponse[models.GetVIPLoanAccruedInterestResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/accruedInterest"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetVIPLoanAccruedInterestResponse](
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

type ApiGetVIPLoanOngoingOrdersRequest struct {
	ctx                 context.Context
	ApiService          *UserInformationAPIService
	orderId             *int64
	collateralAccountId *int64
	loanCoin            *string
	collateralCoin      *string
	current             *int64
	limit               *int64
	recvWindow          *int64
}

func (r ApiGetVIPLoanOngoingOrdersRequest) OrderId(orderId int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) CollateralAccountId(collateralAccountId int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.collateralAccountId = &collateralAccountId
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) LoanCoin(loanCoin string) ApiGetVIPLoanOngoingOrdersRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) CollateralCoin(collateralCoin string) ApiGetVIPLoanOngoingOrdersRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) Current(current int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.current = &current
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) Limit(limit int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) RecvWindow(recvWindow int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetVIPLoanOngoingOrdersRequest) Execute() (*common.RestApiResponse[models.GetVIPLoanOngoingOrdersResponse], error) {
	return r.ApiService.GetVIPLoanOngoingOrdersExecute(r)
}

/*
GetVIPLoanOngoingOrders Get VIP Loan Ongoing Orders (USER_DATA)
Get /sapi/v1/loan/vip/ongoing/orders

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/user-information#get-viploan-ongoing-orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param collateralAccountId -
@param loanCoin -
@param collateralCoin -
@param current -
@param limit -
@param recvWindow -
@return ApiGetVIPLoanOngoingOrdersRequest
*/
func (a *UserInformationAPIService) GetVIPLoanOngoingOrders(ctx context.Context) ApiGetVIPLoanOngoingOrdersRequest {
	return ApiGetVIPLoanOngoingOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetVIPLoanOngoingOrdersResponse
func (a *UserInformationAPIService) GetVIPLoanOngoingOrdersExecute(r ApiGetVIPLoanOngoingOrdersRequest) (*common.RestApiResponse[models.GetVIPLoanOngoingOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/ongoing/orders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.collateralAccountId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralAccountId", r.collateralAccountId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetVIPLoanOngoingOrdersResponse](
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

type ApiGetVIPLoanRepaymentHistoryRequest struct {
	ctx        context.Context
	ApiService *UserInformationAPIService
	orderId    *int64
	loanCoin   *string
	startTime  *int64
	endTime    *int64
	current    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetVIPLoanRepaymentHistoryRequest) OrderId(orderId int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetVIPLoanRepaymentHistoryRequest) LoanCoin(loanCoin string) ApiGetVIPLoanRepaymentHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

// If both startTime and endTime are omitted, the most recent 90 days are returned.
func (r ApiGetVIPLoanRepaymentHistoryRequest) StartTime(startTime int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.startTime = &startTime
	return r
}

// Maximum interval between startTime and endTime is 180 days.
func (r ApiGetVIPLoanRepaymentHistoryRequest) EndTime(endTime int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current page number, starting from 1.
func (r ApiGetVIPLoanRepaymentHistoryRequest) Current(current int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.current = &current
	return r
}

// Number of records per page.
func (r ApiGetVIPLoanRepaymentHistoryRequest) Limit(limit int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetVIPLoanRepaymentHistoryRequest) RecvWindow(recvWindow int64) ApiGetVIPLoanRepaymentHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetVIPLoanRepaymentHistoryRequest) Execute() (*common.RestApiResponse[models.GetVIPLoanRepaymentHistoryResponse], error) {
	return r.ApiService.GetVIPLoanRepaymentHistoryExecute(r)
}

/*
GetVIPLoanRepaymentHistory Get VIP Loan Repayment History (USER_DATA)
Get /sapi/v1/loan/vip/repay/history

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/user-information#get-viploan-repayment-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param loanCoin -
@param startTime -  If both startTime and endTime are omitted, the most recent 90 days are returned.
@param endTime -  Maximum interval between startTime and endTime is 180 days.
@param current -  Current page number, starting from 1.
@param limit -  Number of records per page.
@param recvWindow -
@return ApiGetVIPLoanRepaymentHistoryRequest
*/
func (a *UserInformationAPIService) GetVIPLoanRepaymentHistory(ctx context.Context) ApiGetVIPLoanRepaymentHistoryRequest {
	return ApiGetVIPLoanRepaymentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetVIPLoanRepaymentHistoryResponse
func (a *UserInformationAPIService) GetVIPLoanRepaymentHistoryExecute(r ApiGetVIPLoanRepaymentHistoryRequest) (*common.RestApiResponse[models.GetVIPLoanRepaymentHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/repay/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetVIPLoanRepaymentHistoryResponse](
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

type ApiQueryApplicationStatusRequest struct {
	ctx        context.Context
	ApiService *UserInformationAPIService
	current    *int64
	limit      *int64
	recvWindow *int64
}

// Current page number, starting from 1.
func (r ApiQueryApplicationStatusRequest) Current(current int64) ApiQueryApplicationStatusRequest {
	r.current = &current
	return r
}

func (r ApiQueryApplicationStatusRequest) Limit(limit int64) ApiQueryApplicationStatusRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryApplicationStatusRequest) RecvWindow(recvWindow int64) ApiQueryApplicationStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryApplicationStatusRequest) Execute() (*common.RestApiResponse[models.QueryApplicationStatusResponse], error) {
	return r.ApiService.QueryApplicationStatusExecute(r)
}

/*
QueryApplicationStatus Query Application Status (USER_DATA)
Get /sapi/v1/loan/vip/request/data

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/user-information#query-application-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param current -  Current page number, starting from 1.
@param limit -
@param recvWindow -
@return ApiQueryApplicationStatusRequest
*/
func (a *UserInformationAPIService) QueryApplicationStatus(ctx context.Context) ApiQueryApplicationStatusRequest {
	return ApiQueryApplicationStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryApplicationStatusResponse
func (a *UserInformationAPIService) QueryApplicationStatusExecute(r ApiQueryApplicationStatusRequest) (*common.RestApiResponse[models.QueryApplicationStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/request/data"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryApplicationStatusResponse](
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
