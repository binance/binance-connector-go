/*
Binance VIP Loan REST API

OpenAPI Specification for the Binance VIP Loan REST API
*/

package binanceviploanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
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

https://developers.binance.com/docs/vip_loan/user-information/Check-Locked-Value-of-VIP-Collateral-Account

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

	resp, err := SendRequest[models.CheckVIPLoanCollateralAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

func (r ApiGetVIPLoanAccruedInterestRequest) StartTime(startTime int64) ApiGetVIPLoanAccruedInterestRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetVIPLoanAccruedInterestRequest) EndTime(endTime int64) ApiGetVIPLoanAccruedInterestRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetVIPLoanAccruedInterestRequest) Current(current int64) ApiGetVIPLoanAccruedInterestRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
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

https://developers.binance.com/docs/vip_loan/user-information/Get-VIP-Loan-Accrued-Interest

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param loanCoin -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
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

	resp, err := SendRequest[models.GetVIPLoanAccruedInterestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetVIPLoanOngoingOrdersRequest) Current(current int64) ApiGetVIPLoanOngoingOrdersRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
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
GetVIPLoanOngoingOrders Get VIP Loan Ongoing Orders(USER_DATA)
Get /sapi/v1/loan/vip/ongoing/orders

https://developers.binance.com/docs/vip_loan/user-information/Get-VIP-Loan-Ongoing-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param collateralAccountId -
@param loanCoin -
@param collateralCoin -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
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

	resp, err := SendRequest[models.GetVIPLoanOngoingOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiQueryApplicationStatusRequest) Current(current int64) ApiQueryApplicationStatusRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
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
QueryApplicationStatus Query Application Status(USER_DATA)
Get /sapi/v1/loan/vip/request/data

https://developers.binance.com/docs/vip_loan/user-information/Query-Application-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
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

	resp, err := SendRequest[models.QueryApplicationStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
