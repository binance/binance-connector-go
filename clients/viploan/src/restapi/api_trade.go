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

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiVipLoanBorrowRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	loanAccountId       *int64
	loanCoin            *string
	loanAmount          *float64
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

func (r ApiVipLoanBorrowRequest) LoanAmount(loanAmount float64) ApiVipLoanBorrowRequest {
	r.loanAmount = &loanAmount
	return r
}

// Collateral account ID(s). Multiple split by &#x60;,&#x60;
func (r ApiVipLoanBorrowRequest) CollateralAccountId(collateralAccountId string) ApiVipLoanBorrowRequest {
	r.collateralAccountId = &collateralAccountId
	return r
}

func (r ApiVipLoanBorrowRequest) CollateralCoin(collateralCoin string) ApiVipLoanBorrowRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// TRUE: flexible rate; FALSE: fixed rate
func (r ApiVipLoanBorrowRequest) IsFlexibleRate(isFlexibleRate bool) ApiVipLoanBorrowRequest {
	r.isFlexibleRate = &isFlexibleRate
	return r
}

// Mandatory for fixed rate. Optional for flexible rate. e.g. 30/60 days
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
VipLoanBorrow VIP Loan Borrow (TRADE)
Post /sapi/v1/loan/vip/borrow

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/trade#vip-loan-borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param loanAccountId -
@param loanCoin -
@param loanAmount -
@param collateralAccountId -  Collateral account ID(s). Multiple split by `,`
@param collateralCoin -
@param isFlexibleRate -  TRUE: flexible rate; FALSE: fixed rate
@param loanTerm -  Mandatory for fixed rate. Optional for flexible rate. e.g. 30/60 days
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

	resp, err := SendRequest[models.VipLoanBorrowResponse](
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

type ApiVipLoanFixedRateBorrowRequest struct {
	ctx                 context.Context
	ApiService          *TradeAPIService
	supplyRequest       *string
	borrowCoin          *string
	loanTerm            *int64
	borrowUid           *int64
	collateralCoin      *string
	collateralAccountId *string
	autoRepay           *bool
	recvWindow          *int64
}

// Supply request string, positional encoding (no key). Multiple entries separated by &#x60;;&#x60;, fields separated by &#x60;:&#x60;, order: &#x60;&lt;requestId&gt;:&lt;interestRate&gt;:&lt;amount&gt;&#x60;. Example: &#x60;1212:0.12:100;3434:0.13:50&#x60;
func (r ApiVipLoanFixedRateBorrowRequest) SupplyRequest(supplyRequest string) ApiVipLoanFixedRateBorrowRequest {
	r.supplyRequest = &supplyRequest
	return r
}

// Borrow coin
func (r ApiVipLoanFixedRateBorrowRequest) BorrowCoin(borrowCoin string) ApiVipLoanFixedRateBorrowRequest {
	r.borrowCoin = &borrowCoin
	return r
}

// Loan term in days
func (r ApiVipLoanFixedRateBorrowRequest) LoanTerm(loanTerm int64) ApiVipLoanFixedRateBorrowRequest {
	r.loanTerm = &loanTerm
	return r
}

// Borrow receiving account UID
func (r ApiVipLoanFixedRateBorrowRequest) BorrowUid(borrowUid int64) ApiVipLoanFixedRateBorrowRequest {
	r.borrowUid = &borrowUid
	return r
}

// Collateral coin(s), multiple separated by &#x60;,&#x60;. Only coin names, no amount (VIP loan collateral amount &#x3D; entire spot account balance)
func (r ApiVipLoanFixedRateBorrowRequest) CollateralCoin(collateralCoin string) ApiVipLoanFixedRateBorrowRequest {
	r.collateralCoin = &collateralCoin
	return r
}

// Collateral account ID(s), multiple separated by &#x60;,&#x60;
func (r ApiVipLoanFixedRateBorrowRequest) CollateralAccountId(collateralAccountId string) ApiVipLoanFixedRateBorrowRequest {
	r.collateralAccountId = &collateralAccountId
	return r
}

// Default: &#x60;true&#x60;. &#x60;true&#x60;: auto repay at expiration; &#x60;false&#x60;: auto-convert to flexible (floating rate) at expiration
func (r ApiVipLoanFixedRateBorrowRequest) AutoRepay(autoRepay bool) ApiVipLoanFixedRateBorrowRequest {
	r.autoRepay = &autoRepay
	return r
}

// The value cannot be greater than &#x60;60000&#x60;
func (r ApiVipLoanFixedRateBorrowRequest) RecvWindow(recvWindow int64) ApiVipLoanFixedRateBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVipLoanFixedRateBorrowRequest) Execute() (*common.RestApiResponse[models.VipLoanFixedRateBorrowResponse], error) {
	return r.ApiService.VipLoanFixedRateBorrowExecute(r)
}

/*
VipLoanFixedRateBorrow VIP Loan Fixed Rate Borrow (TRADE)
Post /sapi/v1/loan/vip/fixed/borrow

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/trade#vip-loan-fixed-rate-borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param supplyRequest -  Supply request string, positional encoding (no key). Multiple entries separated by `;`, fields separated by `:`, order: `<requestId>:<interestRate>:<amount>`. Example: `1212:0.12:100;3434:0.13:50`
@param borrowCoin -  Borrow coin
@param loanTerm -  Loan term in days
@param borrowUid -  Borrow receiving account UID
@param collateralCoin -  Collateral coin(s), multiple separated by `,`. Only coin names, no amount (VIP loan collateral amount = entire spot account balance)
@param collateralAccountId -  Collateral account ID(s), multiple separated by `,`
@param autoRepay -  Default: `true`. `true`: auto repay at expiration; `false`: auto-convert to flexible (floating rate) at expiration
@param recvWindow -  The value cannot be greater than `60000`
@return ApiVipLoanFixedRateBorrowRequest
*/
func (a *TradeAPIService) VipLoanFixedRateBorrow(ctx context.Context) ApiVipLoanFixedRateBorrowRequest {
	return ApiVipLoanFixedRateBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VipLoanFixedRateBorrowResponse
func (a *TradeAPIService) VipLoanFixedRateBorrowExecute(r ApiVipLoanFixedRateBorrowRequest) (*common.RestApiResponse[models.VipLoanFixedRateBorrowResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/loan/vip/fixed/borrow"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.supplyRequest == nil {
		return nil, common.ReportError("supplyRequest is required and must be specified")
	}

	if r.borrowCoin == nil {
		return nil, common.ReportError("borrowCoin is required and must be specified")
	}

	if r.loanTerm == nil {
		return nil, common.ReportError("loanTerm is required and must be specified")
	}

	if r.borrowUid == nil {
		return nil, common.ReportError("borrowUid is required and must be specified")
	}

	if r.collateralCoin == nil {
		return nil, common.ReportError("collateralCoin is required and must be specified")
	}

	if r.collateralAccountId == nil {
		return nil, common.ReportError("collateralAccountId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "supplyRequest", r.supplyRequest, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "borrowCoin", r.borrowCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "loanTerm", r.loanTerm, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "borrowUid", r.borrowUid, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralCoin", r.collateralCoin, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "collateralAccountId", r.collateralAccountId, "form", "")
	if r.autoRepay != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepay", r.autoRepay, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VipLoanFixedRateBorrowResponse](
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
VipLoanRenew VIP Loan Renew (TRADE)
Post /sapi/v1/loan/vip/renew

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/trade#vip-loan-renew

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

	resp, err := SendRequest[models.VipLoanRenewResponse](
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

type ApiVipLoanRepayRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	orderId    *int64
	amount     *float64
	recvWindow *int64
}

func (r ApiVipLoanRepayRequest) OrderId(orderId int64) ApiVipLoanRepayRequest {
	r.orderId = &orderId
	return r
}

func (r ApiVipLoanRepayRequest) Amount(amount float64) ApiVipLoanRepayRequest {
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
VipLoanRepay VIP Loan Repay (TRADE)
Post /sapi/v1/loan/vip/repay

https://developers.binance.com/en/docs/catalog/investment-and-services-vip-loan/api/rest-api/trade#vip-loan-repay

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

	resp, err := SendRequest[models.VipLoanRepayResponse](
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
