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
	"github.com/binance/binance-connector-go/common/v2/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiVipLoanBorrowRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	loanAccountId       *int64
	loanCoin            *string
	loanAmount          *float32
	collateralAccountId *string
	collateralCoin      *string
	isFlexibleRate      *bool
	loanTerm            *int64
	recvWindow          *int64
}

func (r ApiVipLoanBorrowRequest) LoanAccountId(loanAccountId int64) ApiVipLoanBorrowRequest {
	r.loanAccountId = &loanAccountId
	return r
}

func (r ApiVipLoanBorrowRequest) LoanCoin(loanCoin string) ApiVipLoanBorrowRequest {
	r.loanCoin = &loanCoin
	return r
}

func (r ApiVipLoanBorrowRequest) LoanAmount(loanAmount float32) ApiVipLoanBorrowRequest {
	r.loanAmount = &loanAmount
	return r
}

// Multiple split by &#x60;,&#x60;
func (r ApiVipLoanBorrowRequest) CollateralAccountId(collateralAccountId string) ApiVipLoanBorrowRequest {
	r.collateralAccountId = &collateralAccountId
	return r
}

// Multiple split by &#x60;,&#x60;
func (r ApiVipLoanBorrowRequest) CollateralCoin(collateralCoin string) ApiVipLoanBorrowRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// Default: TRUE. TRUE : flexible rate; FALSE: fixed rate
func (r ApiVipLoanBorrowRequest) IsFlexibleRate(isFlexibleRate bool) ApiVipLoanBorrowRequest {
	r.isFlexibleRate = &isFlexibleRate
	return r
}

// Mandatory for fixed rate. Optional for fixed interest rate. Eg: 30/60 days
func (r ApiVipLoanBorrowRequest) LoanTerm(loanTerm int64) ApiVipLoanBorrowRequest {
	r.loanTerm = &loanTerm
	return r
}

func (r ApiVipLoanBorrowRequest) RecvWindow(recvWindow int64) ApiVipLoanBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVipLoanBorrowRequest) Execute() (*common.RestApiResponse[models.VipLoanBorrowResponse], error) {
	return r.ApiService.VipLoanBorrowExecute(r)
}

/*
VipLoanBorrow VIP Loan Borrow(TRADE)
Post /sapi/v1/loan/vip/borrow

https://developers.binance.com/docs/vip_loan/trade/VIP-Loan-Borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanAccountId -
@param loanCoin -
@param loanAmount -
@param collateralAccountId -  Multiple split by `,`
@param collateralCoin -  Multiple split by `,`
@param isFlexibleRate -  Default: TRUE. TRUE : flexible rate; FALSE: fixed rate
@param loanTerm -  Mandatory for fixed rate. Optional for fixed interest rate. Eg: 30/60 days
@param recvWindow -
@return ApiVipLoanBorrowRequest
*/
func (a *TradeAPIService) VipLoanBorrow(ctx context.Context) ApiVipLoanBorrowRequest {
	return ApiVipLoanBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VipLoanBorrowResponse
func (a *TradeAPIService) VipLoanBorrowExecute(r ApiVipLoanBorrowRequest) (*common.RestApiResponse[models.VipLoanBorrowResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/borrow"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.loanAccountId == nil {
		return nil, common.ReportError("loanAccountId is required and must be specified")
	}
	if r.loanCoin == nil {
		return nil, common.ReportError("loanCoin is required and must be specified")
	}
	if r.loanAmount == nil {
		return nil, common.ReportError("loanAmount is required and must be specified")
	}
	if r.collateralAccountId == nil {
		return nil, common.ReportError("collateralAccountId is required and must be specified")
	}
	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}
	if r.isFlexibleRate == nil {
		return nil, common.ReportError("isFlexibleRate is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanAccountId", r.loanAccountId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanCoin", r.loanCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanAmount", r.loanAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralAccountId", r.collateralAccountId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "isFlexibleRate", r.isFlexibleRate, "form", "")
	if r.loanTerm != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanTerm", r.loanTerm, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VipLoanBorrowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiVipLoanRenewRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *int64
	loanTerm   *int64
	recvWindow *int64
}

func (r ApiVipLoanRenewRequest) OrderId(orderId int64) ApiVipLoanRenewRequest {
	r.orderId = &orderId
	return r
}

// 30/60 days
func (r ApiVipLoanRenewRequest) LoanTerm(loanTerm int64) ApiVipLoanRenewRequest {
	r.loanTerm = &loanTerm
	return r
}

func (r ApiVipLoanRenewRequest) RecvWindow(recvWindow int64) ApiVipLoanRenewRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVipLoanRenewRequest) Execute() (*common.RestApiResponse[models.VipLoanRenewResponse], error) {
	return r.ApiService.VipLoanRenewExecute(r)
}

/*
VipLoanRenew VIP Loan Renew(TRADE)
Post /sapi/v1/loan/vip/renew

https://developers.binance.com/docs/vip_loan/trade/VIP-Loan-Renew

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param loanTerm -  30/60 days
@param recvWindow -
@return ApiVipLoanRenewRequest
*/
func (a *TradeAPIService) VipLoanRenew(ctx context.Context) ApiVipLoanRenewRequest {
	return ApiVipLoanRenewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VipLoanRenewResponse
func (a *TradeAPIService) VipLoanRenewExecute(r ApiVipLoanRenewRequest) (*common.RestApiResponse[models.VipLoanRenewResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/renew"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}
	if r.loanTerm == nil {
		return nil, common.ReportError("loanTerm is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanTerm", r.loanTerm, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VipLoanRenewResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiVipLoanRepayRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *int64
	amount     *float32
	recvWindow *int64
}

func (r ApiVipLoanRepayRequest) OrderId(orderId int64) ApiVipLoanRepayRequest {
	r.orderId = &orderId
	return r
}

func (r ApiVipLoanRepayRequest) Amount(amount float32) ApiVipLoanRepayRequest {
	r.amount = &amount
	return r
}

func (r ApiVipLoanRepayRequest) RecvWindow(recvWindow int64) ApiVipLoanRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVipLoanRepayRequest) Execute() (*common.RestApiResponse[models.VipLoanRepayResponse], error) {
	return r.ApiService.VipLoanRepayExecute(r)
}

/*
VipLoanRepay VIP Loan Repay(TRADE)
Post /sapi/v1/loan/vip/repay

https://developers.binance.com/docs/vip_loan/trade/VIP-Loan-Repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -
@param amount -
@param recvWindow -
@return ApiVipLoanRepayRequest
*/
func (a *TradeAPIService) VipLoanRepay(ctx context.Context) ApiVipLoanRepayRequest {
	return ApiVipLoanRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VipLoanRepayResponse
func (a *TradeAPIService) VipLoanRepayExecute(r ApiVipLoanRepayRequest) (*common.RestApiResponse[models.VipLoanRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/repay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VipLoanRepayResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
