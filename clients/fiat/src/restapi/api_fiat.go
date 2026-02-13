/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package binancefiatrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/fiat/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// FiatAPIService FiatAPI Service
type FiatAPIService Service

type ApiDepositRequest struct {
	ctx              context.Context
	ApiService       *FiatAPIService
	currency         *string
	apiPaymentMethod *string
	amount           *int64
	recvWindow       *int64
	ext              *map[string]interface{}
}

func (r ApiDepositRequest) Currency(currency string) ApiDepositRequest {
	r.currency = &currency
	return r
}

func (r ApiDepositRequest) ApiPaymentMethod(apiPaymentMethod string) ApiDepositRequest {
	r.apiPaymentMethod = &apiPaymentMethod
	return r
}

func (r ApiDepositRequest) Amount(amount int64) ApiDepositRequest {
	r.amount = &amount
	return r
}

func (r ApiDepositRequest) RecvWindow(recvWindow int64) ApiDepositRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDepositRequest) Ext(ext map[string]interface{}) ApiDepositRequest {
	r.ext = &ext
	return r
}

func (r ApiDepositRequest) Execute() (*common.RestApiResponse[models.DepositResponse], error) {
	return r.ApiService.DepositExecute(r)
}

/*
Deposit Deposit(TRADE)
Post /sapi/v1/fiat/deposit

https://developers.binance.com/docs/fiat/rest-api/Fiat-Deposit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param currency -
@param apiPaymentMethod -
@param amount -
@param recvWindow -
@param ext -
@return ApiDepositRequest
*/
func (a *FiatAPIService) Deposit(ctx context.Context) ApiDepositRequest {
	return ApiDepositRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositResponse
func (a *FiatAPIService) DepositExecute(r ApiDepositRequest) (*common.RestApiResponse[models.DepositResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/fiat/deposit"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.currency == nil {
		return nil, common.ReportError("currency is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "currency", r.currency, "", "")
	if r.apiPaymentMethod == nil {
		return nil, common.ReportError("apiPaymentMethod is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "apiPaymentMethod", r.apiPaymentMethod, "", "")
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "amount", r.amount, "", "")

	if r.ext != nil {
		localVarBodyParameters["ext"] = *r.ext
	}

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DepositResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFiatWithdrawRequest struct {
	ctx              context.Context
	ApiService       *FiatAPIService
	currency         *string
	apiPaymentMethod *string
	amount           *int64
	accountInfo      *models.AccountInfo
	recvWindow       *int64
	ext              *map[string]interface{}
}

func (r ApiFiatWithdrawRequest) Currency(currency string) ApiFiatWithdrawRequest {
	r.currency = &currency
	return r
}

func (r ApiFiatWithdrawRequest) ApiPaymentMethod(apiPaymentMethod string) ApiFiatWithdrawRequest {
	r.apiPaymentMethod = &apiPaymentMethod
	return r
}

func (r ApiFiatWithdrawRequest) Amount(amount int64) ApiFiatWithdrawRequest {
	r.amount = &amount
	return r
}

func (r ApiFiatWithdrawRequest) AccountInfo(accountInfo models.AccountInfo) ApiFiatWithdrawRequest {
	r.accountInfo = &accountInfo
	return r
}

func (r ApiFiatWithdrawRequest) RecvWindow(recvWindow int64) ApiFiatWithdrawRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFiatWithdrawRequest) Ext(ext map[string]interface{}) ApiFiatWithdrawRequest {
	r.ext = &ext
	return r
}

func (r ApiFiatWithdrawRequest) Execute() (*common.RestApiResponse[models.FiatWithdrawResponse], error) {
	return r.ApiService.FiatWithdrawExecute(r)
}

/*
FiatWithdraw Fiat Withdraw(WITHDRAW)
Post /sapi/v2/fiat/withdraw

https://developers.binance.com/docs/fiat/rest-api/Fiat-Withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param currency -
@param apiPaymentMethod -
@param amount -
@param accountInfo -
@param recvWindow -
@param ext -
@return ApiFiatWithdrawRequest
*/
func (a *FiatAPIService) FiatWithdraw(ctx context.Context) ApiFiatWithdrawRequest {
	return ApiFiatWithdrawRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FiatWithdrawResponse
func (a *FiatAPIService) FiatWithdrawExecute(r ApiFiatWithdrawRequest) (*common.RestApiResponse[models.FiatWithdrawResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/fiat/withdraw"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.currency == nil {
		return nil, common.ReportError("currency is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "currency", r.currency, "", "")
	if r.apiPaymentMethod == nil {
		return nil, common.ReportError("apiPaymentMethod is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "apiPaymentMethod", r.apiPaymentMethod, "", "")
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "amount", r.amount, "", "")
	if r.accountInfo == nil {
		return nil, common.ReportError("accountInfo is required and must be specified")
	}
	common.ParameterAddToHeaderOrQuery(localVarBodyParameters, "accountInfo", r.accountInfo, "", "")

	if r.ext != nil {
		localVarBodyParameters["ext"] = *r.ext
	}

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FiatWithdrawResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFiatDepositWithdrawHistoryRequest struct {
	ctx             context.Context
	ApiService      *FiatAPIService
	transactionType *string
	beginTime       *int64
	endTime         *int64
	page            *int64
	rows            *int64
	recvWindow      *int64
}

// 0-buy,1-sell
func (r ApiGetFiatDepositWithdrawHistoryRequest) TransactionType(transactionType string) ApiGetFiatDepositWithdrawHistoryRequest {
	r.transactionType = &transactionType
	return r
}

func (r ApiGetFiatDepositWithdrawHistoryRequest) BeginTime(beginTime int64) ApiGetFiatDepositWithdrawHistoryRequest {
	r.beginTime = &beginTime
	return r
}

func (r ApiGetFiatDepositWithdrawHistoryRequest) EndTime(endTime int64) ApiGetFiatDepositWithdrawHistoryRequest {
	r.endTime = &endTime
	return r
}

// default 1
func (r ApiGetFiatDepositWithdrawHistoryRequest) Page(page int64) ApiGetFiatDepositWithdrawHistoryRequest {
	r.page = &page
	return r
}

// default 100, max 500
func (r ApiGetFiatDepositWithdrawHistoryRequest) Rows(rows int64) ApiGetFiatDepositWithdrawHistoryRequest {
	r.rows = &rows
	return r
}

func (r ApiGetFiatDepositWithdrawHistoryRequest) RecvWindow(recvWindow int64) ApiGetFiatDepositWithdrawHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFiatDepositWithdrawHistoryRequest) Execute() (*common.RestApiResponse[models.GetFiatDepositWithdrawHistoryResponse], error) {
	return r.ApiService.GetFiatDepositWithdrawHistoryExecute(r)
}

/*
GetFiatDepositWithdrawHistory Get Fiat Deposit/Withdraw History (USER_DATA)
Get /sapi/v1/fiat/orders

https://developers.binance.com/docs/fiat/rest-api/Get-Fiat-Deposit-Withdraw-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param transactionType -  0-buy,1-sell
@param beginTime -
@param endTime -
@param page -  default 1
@param rows -  default 100, max 500
@param recvWindow -
@return ApiGetFiatDepositWithdrawHistoryRequest
*/
func (a *FiatAPIService) GetFiatDepositWithdrawHistory(ctx context.Context) ApiGetFiatDepositWithdrawHistoryRequest {
	return ApiGetFiatDepositWithdrawHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFiatDepositWithdrawHistoryResponse
func (a *FiatAPIService) GetFiatDepositWithdrawHistoryExecute(r ApiGetFiatDepositWithdrawHistoryRequest) (*common.RestApiResponse[models.GetFiatDepositWithdrawHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/fiat/orders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.transactionType == nil {
		return nil, common.ReportError("transactionType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transactionType", r.transactionType, "form", "")
	if r.beginTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "beginTime", r.beginTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.rows != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "rows", r.rows, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFiatDepositWithdrawHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetFiatPaymentsHistoryRequest struct {
	ctx             context.Context
	ApiService      *FiatAPIService
	transactionType *string
	beginTime       *int64
	endTime         *int64
	page            *int64
	rows            *int64
	recvWindow      *int64
}

// 0-buy,1-sell
func (r ApiGetFiatPaymentsHistoryRequest) TransactionType(transactionType string) ApiGetFiatPaymentsHistoryRequest {
	r.transactionType = &transactionType
	return r
}

func (r ApiGetFiatPaymentsHistoryRequest) BeginTime(beginTime int64) ApiGetFiatPaymentsHistoryRequest {
	r.beginTime = &beginTime
	return r
}

func (r ApiGetFiatPaymentsHistoryRequest) EndTime(endTime int64) ApiGetFiatPaymentsHistoryRequest {
	r.endTime = &endTime
	return r
}

// default 1
func (r ApiGetFiatPaymentsHistoryRequest) Page(page int64) ApiGetFiatPaymentsHistoryRequest {
	r.page = &page
	return r
}

// default 100, max 500
func (r ApiGetFiatPaymentsHistoryRequest) Rows(rows int64) ApiGetFiatPaymentsHistoryRequest {
	r.rows = &rows
	return r
}

func (r ApiGetFiatPaymentsHistoryRequest) RecvWindow(recvWindow int64) ApiGetFiatPaymentsHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetFiatPaymentsHistoryRequest) Execute() (*common.RestApiResponse[models.GetFiatPaymentsHistoryResponse], error) {
	return r.ApiService.GetFiatPaymentsHistoryExecute(r)
}

/*
GetFiatPaymentsHistory Get Fiat Payments History (USER_DATA)
Get /sapi/v1/fiat/payments

https://developers.binance.com/docs/fiat/rest-api/Get-Fiat-Payments-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param transactionType -  0-buy,1-sell
@param beginTime -
@param endTime -
@param page -  default 1
@param rows -  default 100, max 500
@param recvWindow -
@return ApiGetFiatPaymentsHistoryRequest
*/
func (a *FiatAPIService) GetFiatPaymentsHistory(ctx context.Context) ApiGetFiatPaymentsHistoryRequest {
	return ApiGetFiatPaymentsHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetFiatPaymentsHistoryResponse
func (a *FiatAPIService) GetFiatPaymentsHistoryExecute(r ApiGetFiatPaymentsHistoryRequest) (*common.RestApiResponse[models.GetFiatPaymentsHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/fiat/payments"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.transactionType == nil {
		return nil, common.ReportError("transactionType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transactionType", r.transactionType, "form", "")
	if r.beginTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "beginTime", r.beginTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.rows != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "rows", r.rows, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetFiatPaymentsHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOrderDetailRequest struct {
	ctx        context.Context
	ApiService *FiatAPIService
	orderNo    *string
	recvWindow *int64
}

// order id retrieved from the api call of withdrawal
func (r ApiGetOrderDetailRequest) OrderNo(orderNo string) ApiGetOrderDetailRequest {
	r.orderNo = &orderNo
	return r
}

func (r ApiGetOrderDetailRequest) RecvWindow(recvWindow int64) ApiGetOrderDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOrderDetailRequest) Execute() (*common.RestApiResponse[models.GetOrderDetailResponse], error) {
	return r.ApiService.GetOrderDetailExecute(r)
}

/*
GetOrderDetail Get Order Detail(USER_DATA)
Get /sapi/v1/fiat/get-order-detail

https://developers.binance.com/docs/fiat/rest-api/Get-Order-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderNo -  order id retrieved from the api call of withdrawal
@param recvWindow -
@return ApiGetOrderDetailRequest
*/
func (a *FiatAPIService) GetOrderDetail(ctx context.Context) ApiGetOrderDetailRequest {
	return ApiGetOrderDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrderDetailResponse
func (a *FiatAPIService) GetOrderDetailExecute(r ApiGetOrderDetailRequest) (*common.RestApiResponse[models.GetOrderDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/fiat/get-order-detail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderNo == nil {
		return nil, common.ReportError("orderNo is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderNo", r.orderNo, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOrderDetailResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
