/*
Crypto Loan REST API

Access Binance Crypto Loans to query assets, subscribe to loans, and manage loan positions.
*/

package binancecryptoloanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/cryptoloan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// StableRateAPIService StableRateAPI Service
type StableRateAPIService Service

type ApiGetCryptoLoansIncomeHistoryRequest struct {
	ctx        context.Context
	ApiService *StableRateAPIService
	asset      *string
	type_      *models.GetCryptoLoansIncomeHistoryTypeParameter
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetCryptoLoansIncomeHistoryRequest) Asset(asset string) ApiGetCryptoLoansIncomeHistoryRequest {
	r.asset = &asset
	return r
}

// All types will be returned by default.
func (r ApiGetCryptoLoansIncomeHistoryRequest) Type(type_ models.GetCryptoLoansIncomeHistoryTypeParameter) ApiGetCryptoLoansIncomeHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetCryptoLoansIncomeHistoryRequest) StartTime(startTime int64) ApiGetCryptoLoansIncomeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetCryptoLoansIncomeHistoryRequest) EndTime(endTime int64) ApiGetCryptoLoansIncomeHistoryRequest {
	r.endTime = &endTime
	return r
}

// Number of records to return
func (r ApiGetCryptoLoansIncomeHistoryRequest) Limit(limit int64) ApiGetCryptoLoansIncomeHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiGetCryptoLoansIncomeHistoryRequest) RecvWindow(recvWindow int64) ApiGetCryptoLoansIncomeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCryptoLoansIncomeHistoryRequest) Execute() (*common.RestApiResponse[models.GetCryptoLoansIncomeHistoryResponse], error) {
	return r.ApiService.GetCryptoLoansIncomeHistoryExecute(r)
}

/*
GetCryptoLoansIncomeHistory Get Crypto Loans Income History (USER_DATA)
Get /sapi/v1/loan/income

https://developers.binance.com/en/docs/catalog/investment-and-services-crypto-loan/api/rest-api/stable-rate#get-crypto-loans-income-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param type_ -  All types will be returned by default.
@param startTime -
@param endTime -
@param limit -  Number of records to return
@param recvWindow -  Request validity window in milliseconds
@return ApiGetCryptoLoansIncomeHistoryRequest
*/
func (a *StableRateAPIService) GetCryptoLoansIncomeHistory(ctx context.Context) ApiGetCryptoLoansIncomeHistoryRequest {
	return ApiGetCryptoLoansIncomeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCryptoLoansIncomeHistoryResponse
func (a *StableRateAPIService) GetCryptoLoansIncomeHistoryExecute(r ApiGetCryptoLoansIncomeHistoryRequest) (*common.RestApiResponse[models.GetCryptoLoansIncomeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/income"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCryptoLoansIncomeHistoryResponse](
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

type ApiGetLoanBorrowHistoryRequest struct {
	ctx            context.Context
	ApiService     *StableRateAPIService
	orderId        *int64
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

// orderId in &#x60;POST /sapi/v1/loan/borrow&#x60;
func (r ApiGetLoanBorrowHistoryRequest) OrderId(orderId int64) ApiGetLoanBorrowHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetLoanBorrowHistoryRequest) LoanCoin(loanCoin string) ApiGetLoanBorrowHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetLoanBorrowHistoryRequest) CollateralCoin(collateralCoin string) ApiGetLoanBorrowHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetLoanBorrowHistoryRequest) StartTime(startTime int64) ApiGetLoanBorrowHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLoanBorrowHistoryRequest) EndTime(endTime int64) ApiGetLoanBorrowHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page
func (r ApiGetLoanBorrowHistoryRequest) Current(current int64) ApiGetLoanBorrowHistoryRequest {
	r.current = &current
	return r
}

// Number of records to return
func (r ApiGetLoanBorrowHistoryRequest) Limit(limit int64) ApiGetLoanBorrowHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiGetLoanBorrowHistoryRequest) RecvWindow(recvWindow int64) ApiGetLoanBorrowHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLoanBorrowHistoryRequest) Execute() (*common.RestApiResponse[models.GetLoanBorrowHistoryResponse], error) {
	return r.ApiService.GetLoanBorrowHistoryExecute(r)
}

/*
GetLoanBorrowHistory Get Loan Borrow History (USER_DATA)
Get /sapi/v1/loan/borrow/history

https://developers.binance.com/en/docs/catalog/investment-and-services-crypto-loan/api/rest-api/stable-rate#get-loan-borrow-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  orderId in `POST /sapi/v1/loan/borrow`
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page
@param limit -  Number of records to return
@param recvWindow -  Request validity window in milliseconds
@return ApiGetLoanBorrowHistoryRequest
*/
func (a *StableRateAPIService) GetLoanBorrowHistory(ctx context.Context) ApiGetLoanBorrowHistoryRequest {
	return ApiGetLoanBorrowHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLoanBorrowHistoryResponse
func (a *StableRateAPIService) GetLoanBorrowHistoryExecute(r ApiGetLoanBorrowHistoryRequest) (*common.RestApiResponse[models.GetLoanBorrowHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/borrow/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
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

	resp, err := SendRequest[models.GetLoanBorrowHistoryResponse](
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

type ApiGetLoanLtvAdjustmentHistoryRequest struct {
	ctx            context.Context
	ApiService     *StableRateAPIService
	orderId        *int64
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

// orderId in &#x60;POST /sapi/v1/loan/borrow&#x60;
func (r ApiGetLoanLtvAdjustmentHistoryRequest) OrderId(orderId int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetLoanLtvAdjustmentHistoryRequest) LoanCoin(loanCoin string) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetLoanLtvAdjustmentHistoryRequest) CollateralCoin(collateralCoin string) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetLoanLtvAdjustmentHistoryRequest) StartTime(startTime int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLoanLtvAdjustmentHistoryRequest) EndTime(endTime int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page
func (r ApiGetLoanLtvAdjustmentHistoryRequest) Current(current int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.current = &current
	return r
}

// Number of records to return
func (r ApiGetLoanLtvAdjustmentHistoryRequest) Limit(limit int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiGetLoanLtvAdjustmentHistoryRequest) RecvWindow(recvWindow int64) ApiGetLoanLtvAdjustmentHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLoanLtvAdjustmentHistoryRequest) Execute() (*common.RestApiResponse[models.GetLoanLtvAdjustmentHistoryResponse], error) {
	return r.ApiService.GetLoanLtvAdjustmentHistoryExecute(r)
}

/*
GetLoanLtvAdjustmentHistory Get Loan LTV Adjustment History (USER_DATA)
Get /sapi/v1/loan/ltv/adjustment/history

https://developers.binance.com/en/docs/catalog/investment-and-services-crypto-loan/api/rest-api/stable-rate#get-loan-ltv-adjustment-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  orderId in `POST /sapi/v1/loan/borrow`
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page
@param limit -  Number of records to return
@param recvWindow -  Request validity window in milliseconds
@return ApiGetLoanLtvAdjustmentHistoryRequest
*/
func (a *StableRateAPIService) GetLoanLtvAdjustmentHistory(ctx context.Context) ApiGetLoanLtvAdjustmentHistoryRequest {
	return ApiGetLoanLtvAdjustmentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLoanLtvAdjustmentHistoryResponse
func (a *StableRateAPIService) GetLoanLtvAdjustmentHistoryExecute(r ApiGetLoanLtvAdjustmentHistoryRequest) (*common.RestApiResponse[models.GetLoanLtvAdjustmentHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/ltv/adjustment/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
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

	resp, err := SendRequest[models.GetLoanLtvAdjustmentHistoryResponse](
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

type ApiGetLoanRepaymentHistoryRequest struct {
	ctx            context.Context
	ApiService     *StableRateAPIService
	orderId        *int64
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

// orderId in &#x60;POST /sapi/v1/loan/borrow&#x60;
func (r ApiGetLoanRepaymentHistoryRequest) OrderId(orderId int64) ApiGetLoanRepaymentHistoryRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetLoanRepaymentHistoryRequest) LoanCoin(loanCoin string) ApiGetLoanRepaymentHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetLoanRepaymentHistoryRequest) CollateralCoin(collateralCoin string) ApiGetLoanRepaymentHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetLoanRepaymentHistoryRequest) StartTime(startTime int64) ApiGetLoanRepaymentHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetLoanRepaymentHistoryRequest) EndTime(endTime int64) ApiGetLoanRepaymentHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page
func (r ApiGetLoanRepaymentHistoryRequest) Current(current int64) ApiGetLoanRepaymentHistoryRequest {
	r.current = &current
	return r
}

// Number of records to return
func (r ApiGetLoanRepaymentHistoryRequest) Limit(limit int64) ApiGetLoanRepaymentHistoryRequest {
	r.limit = &limit
	return r
}

// Request validity window in milliseconds
func (r ApiGetLoanRepaymentHistoryRequest) RecvWindow(recvWindow int64) ApiGetLoanRepaymentHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetLoanRepaymentHistoryRequest) Execute() (*common.RestApiResponse[models.GetLoanRepaymentHistoryResponse], error) {
	return r.ApiService.GetLoanRepaymentHistoryExecute(r)
}

/*
GetLoanRepaymentHistory Get Loan Repayment History (USER_DATA)
Get /sapi/v1/loan/repay/history

https://developers.binance.com/en/docs/catalog/investment-and-services-crypto-loan/api/rest-api/stable-rate#get-loan-repayment-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  orderId in `POST /sapi/v1/loan/borrow`
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page
@param limit -  Number of records to return
@param recvWindow -  Request validity window in milliseconds
@return ApiGetLoanRepaymentHistoryRequest
*/
func (a *StableRateAPIService) GetLoanRepaymentHistory(ctx context.Context) ApiGetLoanRepaymentHistoryRequest {
	return ApiGetLoanRepaymentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLoanRepaymentHistoryResponse
func (a *StableRateAPIService) GetLoanRepaymentHistoryExecute(r ApiGetLoanRepaymentHistoryRequest) (*common.RestApiResponse[models.GetLoanRepaymentHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/repay/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
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

	resp, err := SendRequest[models.GetLoanRepaymentHistoryResponse](
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
