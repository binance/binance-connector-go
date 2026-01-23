/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API
*/

package binancecryptoloanrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/cryptoloan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// FlexibleRateAPIService FlexibleRateAPI Service
type FlexibleRateAPIService Service

type ApiCheckCollateralRepayRateRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	recvWindow     *int64
}

func (r ApiCheckCollateralRepayRateRequest) LoanCoin(loanCoin string) ApiCheckCollateralRepayRateRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiCheckCollateralRepayRateRequest) CollateralCoin(collateralCoin string) ApiCheckCollateralRepayRateRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiCheckCollateralRepayRateRequest) RecvWindow(recvWindow int64) ApiCheckCollateralRepayRateRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCheckCollateralRepayRateRequest) Execute() (*common.RestApiResponse[models.CheckCollateralRepayRateResponse], error) {
	return r.ApiService.CheckCollateralRepayRateExecute(r)
}

/*
CheckCollateralRepayRate Check Collateral Repay Rate (USER_DATA)
Get /sapi/v2/loan/flexible/repay/rate

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Check-Collateral-Repay-Rate

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param recvWindow -
@return ApiCheckCollateralRepayRateRequest
*/
func (a *FlexibleRateAPIService) CheckCollateralRepayRate(ctx context.Context) ApiCheckCollateralRepayRateRequest {
	return ApiCheckCollateralRepayRateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CheckCollateralRepayRateResponse
func (a *FlexibleRateAPIService) CheckCollateralRepayRateExecute(r ApiCheckCollateralRepayRateRequest) (*common.RestApiResponse[models.CheckCollateralRepayRateResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/repay/rate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}
	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CheckCollateralRepayRateResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFlexibleLoanAdjustLtvRequest struct {
	ctx              context.Context
	ApiService       *FlexibleRateAPIService
	loanCoin         *string
	collateralCoin   *string
	adjustmentAmount *float32
	direction        *string
	recvWindow       *int64
}

func (r ApiFlexibleLoanAdjustLtvRequest) LoanCoin(loanCoin string) ApiFlexibleLoanAdjustLtvRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiFlexibleLoanAdjustLtvRequest) CollateralCoin(collateralCoin string) ApiFlexibleLoanAdjustLtvRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiFlexibleLoanAdjustLtvRequest) AdjustmentAmount(adjustmentAmount float32) ApiFlexibleLoanAdjustLtvRequest {
	r.adjustmentAmount = &adjustmentAmount
	return r
}

// \&quot;ADDITIONAL\&quot;, \&quot;REDUCED\&quot;
func (r ApiFlexibleLoanAdjustLtvRequest) Direction(direction string) ApiFlexibleLoanAdjustLtvRequest {
	r.direction = &direction
	return r
}

func (r ApiFlexibleLoanAdjustLtvRequest) RecvWindow(recvWindow int64) ApiFlexibleLoanAdjustLtvRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFlexibleLoanAdjustLtvRequest) Execute() (*common.RestApiResponse[models.FlexibleLoanAdjustLtvResponse], error) {
	return r.ApiService.FlexibleLoanAdjustLtvExecute(r)
}

/*
FlexibleLoanAdjustLtv Flexible Loan Adjust LTV(TRADE)
Post /sapi/v2/loan/flexible/adjust/ltv

https://developers.binance.com/docs/crypto_loan/flexible-rate/trade/Flexible-Loan-Adjust-LTV

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param adjustmentAmount -
@param direction -  \"ADDITIONAL\", \"REDUCED\"
@param recvWindow -
@return ApiFlexibleLoanAdjustLtvRequest
*/
func (a *FlexibleRateAPIService) FlexibleLoanAdjustLtv(ctx context.Context) ApiFlexibleLoanAdjustLtvRequest {
	return ApiFlexibleLoanAdjustLtvRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlexibleLoanAdjustLtvResponse
func (a *FlexibleRateAPIService) FlexibleLoanAdjustLtvExecute(r ApiFlexibleLoanAdjustLtvRequest) (*common.RestApiResponse[models.FlexibleLoanAdjustLtvResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/adjust/ltv"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}
	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}
	if r.adjustmentAmount == nil {
		return nil, common.ReportError("adjustmentAmount is required and must be specified")
	}
	if r.direction == nil {
		return nil, common.ReportError("direction is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "adjustmentAmount", r.adjustmentAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FlexibleLoanAdjustLtvResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFlexibleLoanBorrowRequest struct {
	ctx              context.Context
	ApiService       *FlexibleRateAPIService
	loanCoin         *string
	collateralCoin   *string
	loanAmount       *float32
	collateralAmount *float32
	recvWindow       *int64
}

func (r ApiFlexibleLoanBorrowRequest) LoanCoin(loanCoin string) ApiFlexibleLoanBorrowRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiFlexibleLoanBorrowRequest) CollateralCoin(collateralCoin string) ApiFlexibleLoanBorrowRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// Mandatory when collateralAmount is empty
func (r ApiFlexibleLoanBorrowRequest) LoanAmount(loanAmount float32) ApiFlexibleLoanBorrowRequest {
	r.loanAmount = &loanAmount
	return r
}

// Mandatory when loanAmount is empty
func (r ApiFlexibleLoanBorrowRequest) CollateralAmount(collateralAmount float32) ApiFlexibleLoanBorrowRequest {
	r.collateralAmount = &collateralAmount
	return r
}

func (r ApiFlexibleLoanBorrowRequest) RecvWindow(recvWindow int64) ApiFlexibleLoanBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFlexibleLoanBorrowRequest) Execute() (*common.RestApiResponse[models.FlexibleLoanBorrowResponse], error) {
	return r.ApiService.FlexibleLoanBorrowExecute(r)
}

/*
FlexibleLoanBorrow Flexible Loan Borrow(TRADE)
Post /sapi/v2/loan/flexible/borrow

https://developers.binance.com/docs/crypto_loan/flexible-rate/trade/Flexible-Loan-Borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param loanAmount -  Mandatory when collateralAmount is empty
@param collateralAmount -  Mandatory when loanAmount is empty
@param recvWindow -
@return ApiFlexibleLoanBorrowRequest
*/
func (a *FlexibleRateAPIService) FlexibleLoanBorrow(ctx context.Context) ApiFlexibleLoanBorrowRequest {
	return ApiFlexibleLoanBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlexibleLoanBorrowResponse
func (a *FlexibleRateAPIService) FlexibleLoanBorrowExecute(r ApiFlexibleLoanBorrowRequest) (*common.RestApiResponse[models.FlexibleLoanBorrowResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/borrow"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}
	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	if r.loanAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanAmount", r.loanAmount, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	if r.collateralAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralAmount", r.collateralAmount, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FlexibleLoanBorrowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFlexibleLoanRepayRequest struct {
	ctx              context.Context
	ApiService       *FlexibleRateAPIService
	loanCoin         *string
	collateralCoin   *string
	repayAmount      *float32
	collateralReturn *bool
	fullRepayment    *bool
	repaymentType    *int64
	recvWindow       *int64
}

func (r ApiFlexibleLoanRepayRequest) LoanCoin(loanCoin string) ApiFlexibleLoanRepayRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiFlexibleLoanRepayRequest) CollateralCoin(collateralCoin string) ApiFlexibleLoanRepayRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// repay amount of loanCoin
func (r ApiFlexibleLoanRepayRequest) RepayAmount(repayAmount float32) ApiFlexibleLoanRepayRequest {
	r.repayAmount = &repayAmount
	return r
}

// Default: TRUE. TRUE: Return extra collateral to spot account; FALSE: Keep extra collateral in the order, and lower LTV.
func (r ApiFlexibleLoanRepayRequest) CollateralReturn(collateralReturn bool) ApiFlexibleLoanRepayRequest {
	r.collateralReturn = &collateralReturn
	return r
}

// Default: FALSE. TRUE: Full repayment; FALSE: Partial repayment, based on loanAmount
func (r ApiFlexibleLoanRepayRequest) FullRepayment(fullRepayment bool) ApiFlexibleLoanRepayRequest {
	r.fullRepayment = &fullRepayment
	return r
}

// Default: 1. 1: Repayment with loan asset; 2: Repayment with collateral
func (r ApiFlexibleLoanRepayRequest) RepaymentType(repaymentType int64) ApiFlexibleLoanRepayRequest {
	r.repaymentType = &repaymentType
	return r
}

func (r ApiFlexibleLoanRepayRequest) RecvWindow(recvWindow int64) ApiFlexibleLoanRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFlexibleLoanRepayRequest) Execute() (*common.RestApiResponse[models.FlexibleLoanRepayResponse], error) {
	return r.ApiService.FlexibleLoanRepayExecute(r)
}

/*
FlexibleLoanRepay Flexible Loan Repay(TRADE)
Post /sapi/v2/loan/flexible/repay

https://developers.binance.com/docs/crypto_loan/flexible-rate/trade/Flexible-Loan-Repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param repayAmount -  repay amount of loanCoin
@param collateralReturn -  Default: TRUE. TRUE: Return extra collateral to spot account; FALSE: Keep extra collateral in the order, and lower LTV.
@param fullRepayment -  Default: FALSE. TRUE: Full repayment; FALSE: Partial repayment, based on loanAmount
@param repaymentType -  Default: 1. 1: Repayment with loan asset; 2: Repayment with collateral
@param recvWindow -
@return ApiFlexibleLoanRepayRequest
*/
func (a *FlexibleRateAPIService) FlexibleLoanRepay(ctx context.Context) ApiFlexibleLoanRepayRequest {
	return ApiFlexibleLoanRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FlexibleLoanRepayResponse
func (a *FlexibleRateAPIService) FlexibleLoanRepayExecute(r ApiFlexibleLoanRepayRequest) (*common.RestApiResponse[models.FlexibleLoanRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/repay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}
	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}
	if r.repayAmount == nil {
		return nil, common.ReportError("repayAmount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "repayAmount", r.repayAmount, "form", "")
	if r.collateralReturn != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralReturn", r.collateralReturn, "form", "")
	}
	if r.fullRepayment != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fullRepayment", r.fullRepayment, "form", "")
	}
	if r.repaymentType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "repaymentType", r.repaymentType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FlexibleLoanRepayResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanAssetsDataRequest struct {
	ctx        context.Context
	ApiService *FlexibleRateAPIService
	loanCoin   *string
	recvWindow *int64
}

func (r ApiGetFlexibleLoanAssetsDataRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanAssetsDataRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanAssetsDataRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanAssetsDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanAssetsDataRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanAssetsDataResponse], error) {
	return r.ApiService.GetFlexibleLoanAssetsDataExecute(r)
}

/*
GetFlexibleLoanAssetsData Get Flexible Loan Assets Data(USER_DATA)
Get /sapi/v2/loan/flexible/loanable/data

https://developers.binance.com/docs/crypto_loan/flexible-rate/market-data/Get-Flexible-Loan-Assets-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param recvWindow -
@return ApiGetFlexibleLoanAssetsDataRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanAssetsData(ctx context.Context) ApiGetFlexibleLoanAssetsDataRequest {
	return ApiGetFlexibleLoanAssetsDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanAssetsDataResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanAssetsDataExecute(r ApiGetFlexibleLoanAssetsDataRequest) (*common.RestApiResponse[models.GetFlexibleLoanAssetsDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/loanable/data"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFlexibleLoanAssetsDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanBorrowHistoryRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) StartTime(startTime int64) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) EndTime(endTime int64) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanBorrowHistoryRequest) Current(current int64) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanBorrowHistoryRequest) Limit(limit int64) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanBorrowHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanBorrowHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanBorrowHistoryResponse], error) {
	return r.ApiService.GetFlexibleLoanBorrowHistoryExecute(r)
}

/*
GetFlexibleLoanBorrowHistory Get Flexible Loan Borrow History(USER_DATA)
Get /sapi/v2/loan/flexible/borrow/history

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Get-Flexible-Loan-Borrow-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@param recvWindow -
@return ApiGetFlexibleLoanBorrowHistoryRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanBorrowHistory(ctx context.Context) ApiGetFlexibleLoanBorrowHistoryRequest {
	return ApiGetFlexibleLoanBorrowHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanBorrowHistoryResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanBorrowHistoryExecute(r ApiGetFlexibleLoanBorrowHistoryRequest) (*common.RestApiResponse[models.GetFlexibleLoanBorrowHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/borrow/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetFlexibleLoanBorrowHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanCollateralAssetsDataRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	collateralCoin *string
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanCollateralAssetsDataRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanCollateralAssetsDataRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetFlexibleLoanCollateralAssetsDataRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanCollateralAssetsDataRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanCollateralAssetsDataRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanCollateralAssetsDataResponse], error) {
	return r.ApiService.GetFlexibleLoanCollateralAssetsDataExecute(r)
}

/*
GetFlexibleLoanCollateralAssetsData Get Flexible Loan Collateral Assets Data(USER_DATA)
Get /sapi/v2/loan/flexible/collateral/data

https://developers.binance.com/docs/crypto_loan/flexible-rate/market-data/Get-Flexible-Loan-Collateral-Assets-Data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param collateralCoin -
@param recvWindow -
@return ApiGetFlexibleLoanCollateralAssetsDataRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanCollateralAssetsData(ctx context.Context) ApiGetFlexibleLoanCollateralAssetsDataRequest {
	return ApiGetFlexibleLoanCollateralAssetsDataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanCollateralAssetsDataResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanCollateralAssetsDataExecute(r ApiGetFlexibleLoanCollateralAssetsDataRequest) (*common.RestApiResponse[models.GetFlexibleLoanCollateralAssetsDataResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/collateral/data"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.collateralCoin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFlexibleLoanCollateralAssetsDataResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanInterestRateHistoryRequest struct {
	ctx        context.Context
	ApiService *FlexibleRateAPIService
	coin       *string
	recvWindow *int64
	startTime  *int64
	endTime    *int64
	current    *int64
	limit      *int64
}

func (r ApiGetFlexibleLoanInterestRateHistoryRequest) Coin(coin string) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.coin = &coin
	return r
}

func (r ApiGetFlexibleLoanInterestRateHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanInterestRateHistoryRequest) StartTime(startTime int64) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleLoanInterestRateHistoryRequest) EndTime(endTime int64) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanInterestRateHistoryRequest) Current(current int64) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanInterestRateHistoryRequest) Limit(limit int64) ApiGetFlexibleLoanInterestRateHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanInterestRateHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanInterestRateHistoryResponse], error) {
	return r.ApiService.GetFlexibleLoanInterestRateHistoryExecute(r)
}

/*
GetFlexibleLoanInterestRateHistory Get Flexible Loan Interest Rate History (USER_DATA)
Get /sapi/v2/loan/interestRateHistory

https://developers.binance.com/docs/crypto_loan/flexible-rate/market-data/Get-Flexible-Loan-Interest-Rate-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param recvWindow -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@return ApiGetFlexibleLoanInterestRateHistoryRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanInterestRateHistory(ctx context.Context) ApiGetFlexibleLoanInterestRateHistoryRequest {
	return ApiGetFlexibleLoanInterestRateHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanInterestRateHistoryResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanInterestRateHistoryExecute(r ApiGetFlexibleLoanInterestRateHistoryRequest) (*common.RestApiResponse[models.GetFlexibleLoanInterestRateHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/interestRateHistory"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}
	if r.recvWindow == nil {
		return nil, common.ReportError("recvWindow is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
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
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")

	resp, err := SendRequest[models.GetFlexibleLoanInterestRateHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanLiquidationHistoryRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) StartTime(startTime int64) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) EndTime(endTime int64) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanLiquidationHistoryRequest) Current(current int64) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanLiquidationHistoryRequest) Limit(limit int64) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanLiquidationHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanLiquidationHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanLiquidationHistoryResponse], error) {
	return r.ApiService.GetFlexibleLoanLiquidationHistoryExecute(r)
}

/*
GetFlexibleLoanLiquidationHistory Get Flexible Loan Liquidation History (USER_DATA)
Get /sapi/v2/loan/flexible/liquidation/history

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Get-Flexible-Loan-Liquidation-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@param recvWindow -
@return ApiGetFlexibleLoanLiquidationHistoryRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanLiquidationHistory(ctx context.Context) ApiGetFlexibleLoanLiquidationHistoryRequest {
	return ApiGetFlexibleLoanLiquidationHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanLiquidationHistoryResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanLiquidationHistoryExecute(r ApiGetFlexibleLoanLiquidationHistoryRequest) (*common.RestApiResponse[models.GetFlexibleLoanLiquidationHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/liquidation/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetFlexibleLoanLiquidationHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanLtvAdjustmentHistoryRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) StartTime(startTime int64) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) EndTime(endTime int64) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) Current(current int64) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) Limit(limit int64) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanLtvAdjustmentHistoryResponse], error) {
	return r.ApiService.GetFlexibleLoanLtvAdjustmentHistoryExecute(r)
}

/*
GetFlexibleLoanLtvAdjustmentHistory Get Flexible Loan LTV Adjustment History(USER_DATA)
Get /sapi/v2/loan/flexible/ltv/adjustment/history

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Get-Flexible-Loan-LTV-Adjustment-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@param recvWindow -
@return ApiGetFlexibleLoanLtvAdjustmentHistoryRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanLtvAdjustmentHistory(ctx context.Context) ApiGetFlexibleLoanLtvAdjustmentHistoryRequest {
	return ApiGetFlexibleLoanLtvAdjustmentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanLtvAdjustmentHistoryResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanLtvAdjustmentHistoryExecute(r ApiGetFlexibleLoanLtvAdjustmentHistoryRequest) (*common.RestApiResponse[models.GetFlexibleLoanLtvAdjustmentHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/ltv/adjustment/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetFlexibleLoanLtvAdjustmentHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanOngoingOrdersRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	current        *int64
	limit          *int64
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanOngoingOrdersRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanOngoingOrdersRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanOngoingOrdersRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanOngoingOrdersRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanOngoingOrdersRequest) Current(current int64) ApiGetFlexibleLoanOngoingOrdersRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanOngoingOrdersRequest) Limit(limit int64) ApiGetFlexibleLoanOngoingOrdersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanOngoingOrdersRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanOngoingOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanOngoingOrdersRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanOngoingOrdersResponse], error) {
	return r.ApiService.GetFlexibleLoanOngoingOrdersExecute(r)
}

/*
GetFlexibleLoanOngoingOrders Get Flexible Loan Ongoing Orders(USER_DATA)
Get /sapi/v2/loan/flexible/ongoing/orders

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Get-Flexible-Loan-Ongoing-Orders

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@param recvWindow -
@return ApiGetFlexibleLoanOngoingOrdersRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanOngoingOrders(ctx context.Context) ApiGetFlexibleLoanOngoingOrdersRequest {
	return ApiGetFlexibleLoanOngoingOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanOngoingOrdersResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanOngoingOrdersExecute(r ApiGetFlexibleLoanOngoingOrdersRequest) (*common.RestApiResponse[models.GetFlexibleLoanOngoingOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/ongoing/orders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetFlexibleLoanOngoingOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFlexibleLoanRepaymentHistoryRequest struct {
	ctx            context.Context
	ApiService     *FlexibleRateAPIService
	loanCoin       *string
	collateralCoin *string
	startTime      *int64
	endTime        *int64
	current        *int64
	limit          *int64
	recvWindow     *int64
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) LoanCoin(loanCoin string) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) CollateralCoin(collateralCoin string) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.collateralCoin = &collateralCoin
	return r
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) StartTime(startTime int64) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) EndTime(endTime int64) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.endTime = &endTime
	return r
}

// Current querying page. Start from 1; default: 1; max: 1000
func (r ApiGetFlexibleLoanRepaymentHistoryRequest) Current(current int64) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.current = &current
	return r
}

// Default: 10; max: 100
func (r ApiGetFlexibleLoanRepaymentHistoryRequest) Limit(limit int64) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) RecvWindow(recvWindow int64) ApiGetFlexibleLoanRepaymentHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFlexibleLoanRepaymentHistoryRequest) Execute() (*common.RestApiResponse[models.GetFlexibleLoanRepaymentHistoryResponse], error) {
	return r.ApiService.GetFlexibleLoanRepaymentHistoryExecute(r)
}

/*
GetFlexibleLoanRepaymentHistory Get Flexible Loan Repayment History(USER_DATA)
Get /sapi/v2/loan/flexible/repay/history

https://developers.binance.com/docs/crypto_loan/flexible-rate/user-information/Get-Flexible-Loan-Repayment-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanCoin -
@param collateralCoin -
@param startTime -
@param endTime -
@param current -  Current querying page. Start from 1; default: 1; max: 1000
@param limit -  Default: 10; max: 100
@param recvWindow -
@return ApiGetFlexibleLoanRepaymentHistoryRequest
*/
func (a *FlexibleRateAPIService) GetFlexibleLoanRepaymentHistory(ctx context.Context) ApiGetFlexibleLoanRepaymentHistoryRequest {
	return ApiGetFlexibleLoanRepaymentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFlexibleLoanRepaymentHistoryResponse
func (a *FlexibleRateAPIService) GetFlexibleLoanRepaymentHistoryExecute(r ApiGetFlexibleLoanRepaymentHistoryRequest) (*common.RestApiResponse[models.GetFlexibleLoanRepaymentHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/loan/flexible/repay/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

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

	resp, err := SendRequest[models.GetFlexibleLoanRepaymentHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
