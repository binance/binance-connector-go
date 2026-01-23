/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// ConvertAPIService ConvertAPI Service
type ConvertAPIService Service

type ApiAcceptTheOfferedQuoteRequest struct {
	ctx        context.Context
	ApiService *ConvertAPIService
	quoteId    *string
	recvWindow *int64
}

func (r ApiAcceptTheOfferedQuoteRequest) QuoteId(quoteId string) ApiAcceptTheOfferedQuoteRequest {
	r.quoteId = &quoteId
	return r
}

func (r ApiAcceptTheOfferedQuoteRequest) RecvWindow(recvWindow int64) ApiAcceptTheOfferedQuoteRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAcceptTheOfferedQuoteRequest) Execute() (*common.RestApiResponse[models.AcceptTheOfferedQuoteResponse], error) {
	return r.ApiService.AcceptTheOfferedQuoteExecute(r)
}

/*
AcceptTheOfferedQuote Accept the offered quote (USER_DATA)
Post /fapi/v1/convert/acceptQuote

https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Accept-Quote

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param quoteId -
@param recvWindow -
@return ApiAcceptTheOfferedQuoteRequest
*/
func (a *ConvertAPIService) AcceptTheOfferedQuote(ctx context.Context) ApiAcceptTheOfferedQuoteRequest {
	return ApiAcceptTheOfferedQuoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AcceptTheOfferedQuoteResponse
func (a *ConvertAPIService) AcceptTheOfferedQuoteExecute(r ApiAcceptTheOfferedQuoteRequest) (*common.RestApiResponse[models.AcceptTheOfferedQuoteResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/convert/acceptQuote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.quoteId == nil {
		return nil, common.ReportError("quoteId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteId", r.quoteId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AcceptTheOfferedQuoteResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiListAllConvertPairsRequest struct {
	ctx        context.Context
	ApiService *ConvertAPIService
	fromAsset  *string
	toAsset    *string
}

// User spends coin
func (r ApiListAllConvertPairsRequest) FromAsset(fromAsset string) ApiListAllConvertPairsRequest {
	r.fromAsset = &fromAsset
	return r
}

// User receives coin
func (r ApiListAllConvertPairsRequest) ToAsset(toAsset string) ApiListAllConvertPairsRequest {
	r.toAsset = &toAsset
	return r
}

func (r ApiListAllConvertPairsRequest) Execute() (*common.RestApiResponse[models.ListAllConvertPairsResponse], error) {
	return r.ApiService.ListAllConvertPairsExecute(r)
}

/*
ListAllConvertPairs List All Convert Pairs
Get /fapi/v1/convert/exchangeInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAsset -  User spends coin
@param toAsset -  User receives coin
@return ApiListAllConvertPairsRequest
*/
func (a *ConvertAPIService) ListAllConvertPairs(ctx context.Context) ApiListAllConvertPairsRequest {
	return ApiListAllConvertPairsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ListAllConvertPairsResponse
func (a *ConvertAPIService) ListAllConvertPairsExecute(r ApiListAllConvertPairsRequest) (*common.RestApiResponse[models.ListAllConvertPairsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/convert/exchangeInfo"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAsset", r.fromAsset, "form", "")
	}
	if r.toAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAsset", r.toAsset, "form", "")
	}

	resp, err := SendRequest[models.ListAllConvertPairsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderStatusRequest struct {
	ctx        context.Context
	ApiService *ConvertAPIService
	orderId    *string
	quoteId    *string
}

// Either orderId or quoteId is required
func (r ApiOrderStatusRequest) OrderId(orderId string) ApiOrderStatusRequest {
	r.orderId = &orderId
	return r
}

// Either orderId or quoteId is required
func (r ApiOrderStatusRequest) QuoteId(quoteId string) ApiOrderStatusRequest {
	r.quoteId = &quoteId
	return r
}

func (r ApiOrderStatusRequest) Execute() (*common.RestApiResponse[models.OrderStatusResponse], error) {
	return r.ApiService.OrderStatusExecute(r)
}

/*
OrderStatus Order status(USER_DATA)
Get /fapi/v1/convert/orderStatus

https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Order-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderId -  Either orderId or quoteId is required
@param quoteId -  Either orderId or quoteId is required
@return ApiOrderStatusRequest
*/
func (a *ConvertAPIService) OrderStatus(ctx context.Context) ApiOrderStatusRequest {
	return ApiOrderStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderStatusResponse
func (a *ConvertAPIService) OrderStatusExecute(r ApiOrderStatusRequest) (*common.RestApiResponse[models.OrderStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/convert/orderStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.quoteId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteId", r.quoteId, "form", "")
	}

	resp, err := SendRequest[models.OrderStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSendQuoteRequestRequest struct {
	ctx        context.Context
	ApiService *ConvertAPIService
	fromAsset  *string
	toAsset    *string
	fromAmount *float32
	toAmount   *float32
	validTime  *string
	recvWindow *int64
}

func (r ApiSendQuoteRequestRequest) FromAsset(fromAsset string) ApiSendQuoteRequestRequest {
	r.fromAsset = &fromAsset
	return r
}

func (r ApiSendQuoteRequestRequest) ToAsset(toAsset string) ApiSendQuoteRequestRequest {
	r.toAsset = &toAsset
	return r
}

// When specified, it is the amount you will be debited after the conversion
func (r ApiSendQuoteRequestRequest) FromAmount(fromAmount float32) ApiSendQuoteRequestRequest {
	r.fromAmount = &fromAmount
	return r
}

// When specified, it is the amount you will be credited after the conversion
func (r ApiSendQuoteRequestRequest) ToAmount(toAmount float32) ApiSendQuoteRequestRequest {
	r.toAmount = &toAmount
	return r
}

// 10s, default 10s
func (r ApiSendQuoteRequestRequest) ValidTime(validTime string) ApiSendQuoteRequestRequest {
	r.validTime = &validTime
	return r
}

func (r ApiSendQuoteRequestRequest) RecvWindow(recvWindow int64) ApiSendQuoteRequestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSendQuoteRequestRequest) Execute() (*common.RestApiResponse[models.SendQuoteRequestResponse], error) {
	return r.ApiService.SendQuoteRequestExecute(r)
}

/*
SendQuoteRequest Send Quote Request(USER_DATA)
Post /fapi/v1/convert/getQuote

https://developers.binance.com/docs/derivatives/usds-margined-futures/convert/Send-quote-request

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromAsset -
@param toAsset -
@param fromAmount -  When specified, it is the amount you will be debited after the conversion
@param toAmount -  When specified, it is the amount you will be credited after the conversion
@param validTime -  10s, default 10s
@param recvWindow -
@return ApiSendQuoteRequestRequest
*/
func (a *ConvertAPIService) SendQuoteRequest(ctx context.Context) ApiSendQuoteRequestRequest {
	return ApiSendQuoteRequestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SendQuoteRequestResponse
func (a *ConvertAPIService) SendQuoteRequestExecute(r ApiSendQuoteRequestRequest) (*common.RestApiResponse[models.SendQuoteRequestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/fapi/v1/convert/getQuote"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromAsset == nil {
		return nil, common.ReportError("fromAsset is required and must be specified")
	}
	if r.toAsset == nil {
		return nil, common.ReportError("toAsset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAsset", r.fromAsset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAsset", r.toAsset, "form", "")
	if r.fromAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAmount", r.fromAmount, "form", "")
	}
	if r.toAmount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toAmount", r.toAmount, "form", "")
	}
	if r.validTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "validTime", r.validTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SendQuoteRequestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
