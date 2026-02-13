/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package binancespotrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountCommissionRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
}

func (r ApiAccountCommissionRequest) Symbol(symbol string) ApiAccountCommissionRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAccountCommissionRequest) Execute() (*common.RestApiResponse[models.AccountCommissionResponse], error) {
	return r.ApiService.AccountCommissionExecute(r)
}

/*
AccountCommission Query Commission Rates
Get /api/v3/account/commission

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-commission-rates-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@return ApiAccountCommissionRequest
*/
func (a *AccountAPIService) AccountCommission(ctx context.Context) ApiAccountCommissionRequest {
	return ApiAccountCommissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountCommissionResponse
func (a *AccountAPIService) AccountCommissionExecute(r ApiAccountCommissionRequest) (*common.RestApiResponse[models.AccountCommissionResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/account/commission"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")

	resp, err := SendRequest[models.AccountCommissionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAllOrderListRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int32
	recvWindow *float32
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiAllOrderListRequest) FromId(fromId int64) ApiAllOrderListRequest {
	r.fromId = &fromId
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiAllOrderListRequest) StartTime(startTime int64) ApiAllOrderListRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiAllOrderListRequest) EndTime(endTime int64) ApiAllOrderListRequest {
	r.endTime = &endTime
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiAllOrderListRequest) Limit(limit int32) ApiAllOrderListRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAllOrderListRequest) RecvWindow(recvWindow float32) ApiAllOrderListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllOrderListRequest) Execute() (*common.RestApiResponse[models.AllOrderListResponse], error) {
	return r.ApiService.AllOrderListExecute(r)
}

/*
AllOrderList Query all Order lists
Get /api/v3/allOrderList

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-all-order-lists-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param limit -  Default: 500; Maximum: 1000.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAllOrderListRequest
*/
func (a *AccountAPIService) AllOrderList(ctx context.Context) ApiAllOrderListRequest {
	return ApiAllOrderListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AllOrderListResponse
func (a *AccountAPIService) AllOrderListExecute(r ApiAllOrderListRequest) (*common.RestApiResponse[models.AllOrderListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/allOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
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

	resp, err := SendRequest[models.AllOrderListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAllOrdersRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int32
	recvWindow *float32
}

func (r ApiAllOrdersRequest) Symbol(symbol string) ApiAllOrdersRequest {
	r.symbol = &symbol
	return r
}

func (r ApiAllOrdersRequest) OrderId(orderId int64) ApiAllOrdersRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiAllOrdersRequest) StartTime(startTime int64) ApiAllOrdersRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiAllOrdersRequest) EndTime(endTime int64) ApiAllOrdersRequest {
	r.endTime = &endTime
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiAllOrdersRequest) Limit(limit int32) ApiAllOrdersRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiAllOrdersRequest) RecvWindow(recvWindow float32) ApiAllOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAllOrdersRequest) Execute() (*common.RestApiResponse[models.AllOrdersResponse], error) {
	return r.ApiService.AllOrdersExecute(r)
}

/*
AllOrders All orders
Get /api/v3/allOrders

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#all-orders-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param limit -  Default: 500; Maximum: 1000.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiAllOrdersRequest
*/
func (a *AccountAPIService) AllOrders(ctx context.Context) ApiAllOrdersRequest {
	return ApiAllOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AllOrdersResponse
func (a *AccountAPIService) AllOrdersExecute(r ApiAllOrdersRequest) (*common.RestApiResponse[models.AllOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/allOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
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

	resp, err := SendRequest[models.AllOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAccountRequest struct {
	ctx              context.Context
	ApiService       *AccountAPIService
	omitZeroBalances *bool
	recvWindow       *float32
}

// When set to &#x60;true&#x60;, emits only the non-zero balances of an account. &lt;br&gt;Default value: &#x60;false&#x60;
func (r ApiGetAccountRequest) OmitZeroBalances(omitZeroBalances bool) ApiGetAccountRequest {
	r.omitZeroBalances = &omitZeroBalances
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiGetAccountRequest) RecvWindow(recvWindow float32) ApiGetAccountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetAccountRequest) Execute() (*common.RestApiResponse[models.GetAccountResponse], error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Account information
Get /api/v3/account

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#account-information-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param omitZeroBalances -  When set to `true`, emits only the non-zero balances of an account. <br>Default value: `false`
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiGetAccountRequest
*/
func (a *AccountAPIService) GetAccount(ctx context.Context) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccountResponse
func (a *AccountAPIService) GetAccountExecute(r ApiGetAccountRequest) (*common.RestApiResponse[models.GetAccountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.omitZeroBalances != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "omitZeroBalances", r.omitZeroBalances, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetAccountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOpenOrdersRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *float32
}

// Symbol to query
func (r ApiGetOpenOrdersRequest) Symbol(symbol string) ApiGetOpenOrdersRequest {
	r.symbol = &symbol
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiGetOpenOrdersRequest) RecvWindow(recvWindow float32) ApiGetOpenOrdersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOpenOrdersRequest) Execute() (*common.RestApiResponse[models.GetOpenOrdersResponse], error) {
	return r.ApiService.GetOpenOrdersExecute(r)
}

/*
GetOpenOrders Current open orders
Get /api/v3/openOrders

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#current-open-orders-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -  Symbol to query
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiGetOpenOrdersRequest
*/
func (a *AccountAPIService) GetOpenOrders(ctx context.Context) ApiGetOpenOrdersRequest {
	return ApiGetOpenOrdersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOpenOrdersResponse
func (a *AccountAPIService) GetOpenOrdersExecute(r ApiGetOpenOrdersRequest) (*common.RestApiResponse[models.GetOpenOrdersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrders"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOpenOrdersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOrderRequest struct {
	ctx               context.Context
	ApiService        *AccountAPIService
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *float32
}

func (r ApiGetOrderRequest) Symbol(symbol string) ApiGetOrderRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetOrderRequest) OrderId(orderId int64) ApiGetOrderRequest {
	r.orderId = &orderId
	return r
}

func (r ApiGetOrderRequest) OrigClientOrderId(origClientOrderId string) ApiGetOrderRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiGetOrderRequest) RecvWindow(recvWindow float32) ApiGetOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOrderRequest) Execute() (*common.RestApiResponse[models.GetOrderResponse], error) {
	return r.ApiService.GetOrderExecute(r)
}

/*
GetOrder Query order
Get /api/v3/order

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-order-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param origClientOrderId -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiGetOrderRequest
*/
func (a *AccountAPIService) GetOrder(ctx context.Context) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrderResponse
func (a *AccountAPIService) GetOrderExecute(r ApiGetOrderRequest) (*common.RestApiResponse[models.GetOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOrderListRequest struct {
	ctx               context.Context
	ApiService        *AccountAPIService
	orderListId       *int64
	origClientOrderId *string
	recvWindow        *float32
}

// Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided
func (r ApiGetOrderListRequest) OrderListId(orderListId int64) ApiGetOrderListRequest {
	r.orderListId = &orderListId
	return r
}

func (r ApiGetOrderListRequest) OrigClientOrderId(origClientOrderId string) ApiGetOrderListRequest {
	r.origClientOrderId = &origClientOrderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiGetOrderListRequest) RecvWindow(recvWindow float32) ApiGetOrderListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetOrderListRequest) Execute() (*common.RestApiResponse[models.GetOrderListResponse], error) {
	return r.ApiService.GetOrderListExecute(r)
}

/*
GetOrderList Query Order list
Get /api/v3/orderList

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-order-list-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param orderListId -  Either `orderListId` or `listClientOrderId` must be provided
@param origClientOrderId -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiGetOrderListRequest
*/
func (a *AccountAPIService) GetOrderList(ctx context.Context) ApiGetOrderListRequest {
	return ApiGetOrderListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOrderListResponse
func (a *AccountAPIService) GetOrderListExecute(r ApiGetOrderListRequest) (*common.RestApiResponse[models.GetOrderListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/orderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.orderListId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderListId", r.orderListId, "form", "")
	}
	if r.origClientOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "origClientOrderId", r.origClientOrderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetOrderListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMyAllocationsRequest struct {
	ctx              context.Context
	ApiService       *AccountAPIService
	symbol           *string
	startTime        *int64
	endTime          *int64
	fromAllocationId *int32
	limit            *int32
	orderId          *int64
	recvWindow       *float32
}

func (r ApiMyAllocationsRequest) Symbol(symbol string) ApiMyAllocationsRequest {
	r.symbol = &symbol
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiMyAllocationsRequest) StartTime(startTime int64) ApiMyAllocationsRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiMyAllocationsRequest) EndTime(endTime int64) ApiMyAllocationsRequest {
	r.endTime = &endTime
	return r
}

func (r ApiMyAllocationsRequest) FromAllocationId(fromAllocationId int32) ApiMyAllocationsRequest {
	r.fromAllocationId = &fromAllocationId
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiMyAllocationsRequest) Limit(limit int32) ApiMyAllocationsRequest {
	r.limit = &limit
	return r
}

func (r ApiMyAllocationsRequest) OrderId(orderId int64) ApiMyAllocationsRequest {
	r.orderId = &orderId
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyAllocationsRequest) RecvWindow(recvWindow float32) ApiMyAllocationsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyAllocationsRequest) Execute() (*common.RestApiResponse[models.MyAllocationsResponse], error) {
	return r.ApiService.MyAllocationsExecute(r)
}

/*
MyAllocations Query Allocations
Get /api/v3/myAllocations

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-allocations-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param fromAllocationId -
@param limit -  Default: 500; Maximum: 1000.
@param orderId -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyAllocationsRequest
*/
func (a *AccountAPIService) MyAllocations(ctx context.Context) ApiMyAllocationsRequest {
	return ApiMyAllocationsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MyAllocationsResponse
func (a *AccountAPIService) MyAllocationsExecute(r ApiMyAllocationsRequest) (*common.RestApiResponse[models.MyAllocationsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/myAllocations"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromAllocationId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromAllocationId", r.fromAllocationId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MyAllocationsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMyFiltersRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *float32
}

func (r ApiMyFiltersRequest) Symbol(symbol string) ApiMyFiltersRequest {
	r.symbol = &symbol
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyFiltersRequest) RecvWindow(recvWindow float32) ApiMyFiltersRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyFiltersRequest) Execute() (*common.RestApiResponse[models.MyFiltersResponse], error) {
	return r.ApiService.MyFiltersExecute(r)
}

/*
MyFilters Query relevant filters
Get /api/v3/myFilters

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-relevant-filters-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyFiltersRequest
*/
func (a *AccountAPIService) MyFilters(ctx context.Context) ApiMyFiltersRequest {
	return ApiMyFiltersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MyFiltersResponse
func (a *AccountAPIService) MyFiltersExecute(r ApiMyFiltersRequest) (*common.RestApiResponse[models.MyFiltersResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/myFilters"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MyFiltersResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMyPreventedMatchesRequest struct {
	ctx                  context.Context
	ApiService           *AccountAPIService
	symbol               *string
	preventedMatchId     *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                *int32
	recvWindow           *float32
}

func (r ApiMyPreventedMatchesRequest) Symbol(symbol string) ApiMyPreventedMatchesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMyPreventedMatchesRequest) PreventedMatchId(preventedMatchId int64) ApiMyPreventedMatchesRequest {
	r.preventedMatchId = &preventedMatchId
	return r
}

func (r ApiMyPreventedMatchesRequest) OrderId(orderId int64) ApiMyPreventedMatchesRequest {
	r.orderId = &orderId
	return r
}

func (r ApiMyPreventedMatchesRequest) FromPreventedMatchId(fromPreventedMatchId int64) ApiMyPreventedMatchesRequest {
	r.fromPreventedMatchId = &fromPreventedMatchId
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiMyPreventedMatchesRequest) Limit(limit int32) ApiMyPreventedMatchesRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyPreventedMatchesRequest) RecvWindow(recvWindow float32) ApiMyPreventedMatchesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyPreventedMatchesRequest) Execute() (*common.RestApiResponse[models.MyPreventedMatchesResponse], error) {
	return r.ApiService.MyPreventedMatchesExecute(r)
}

/*
MyPreventedMatches Query Prevented Matches
Get /api/v3/myPreventedMatches

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-prevented-matches-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param preventedMatchId -
@param orderId -
@param fromPreventedMatchId -
@param limit -  Default: 500; Maximum: 1000.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyPreventedMatchesRequest
*/
func (a *AccountAPIService) MyPreventedMatches(ctx context.Context) ApiMyPreventedMatchesRequest {
	return ApiMyPreventedMatchesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MyPreventedMatchesResponse
func (a *AccountAPIService) MyPreventedMatchesExecute(r ApiMyPreventedMatchesRequest) (*common.RestApiResponse[models.MyPreventedMatchesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/myPreventedMatches"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.preventedMatchId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "preventedMatchId", r.preventedMatchId, "form", "")
	}
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.fromPreventedMatchId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromPreventedMatchId", r.fromPreventedMatchId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MyPreventedMatchesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMyTradesRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int32
	recvWindow *float32
}

func (r ApiMyTradesRequest) Symbol(symbol string) ApiMyTradesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiMyTradesRequest) OrderId(orderId int64) ApiMyTradesRequest {
	r.orderId = &orderId
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiMyTradesRequest) StartTime(startTime int64) ApiMyTradesRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiMyTradesRequest) EndTime(endTime int64) ApiMyTradesRequest {
	r.endTime = &endTime
	return r
}

// ID to get aggregate trades from INCLUSIVE.
func (r ApiMyTradesRequest) FromId(fromId int64) ApiMyTradesRequest {
	r.fromId = &fromId
	return r
}

// Default: 500; Maximum: 1000.
func (r ApiMyTradesRequest) Limit(limit int32) ApiMyTradesRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiMyTradesRequest) RecvWindow(recvWindow float32) ApiMyTradesRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMyTradesRequest) Execute() (*common.RestApiResponse[models.MyTradesResponse], error) {
	return r.ApiService.MyTradesExecute(r)
}

/*
MyTrades Account trade list
Get /api/v3/myTrades

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#account-trade-list-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param startTime -  Timestamp in ms to get aggregate trades from INCLUSIVE.
@param endTime -  Timestamp in ms to get aggregate trades until INCLUSIVE.
@param fromId -  ID to get aggregate trades from INCLUSIVE.
@param limit -  Default: 500; Maximum: 1000.
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiMyTradesRequest
*/
func (a *AccountAPIService) MyTrades(ctx context.Context) ApiMyTradesRequest {
	return ApiMyTradesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MyTradesResponse
func (a *AccountAPIService) MyTradesExecute(r ApiMyTradesRequest) (*common.RestApiResponse[models.MyTradesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/myTrades"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.orderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.fromId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromId", r.fromId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MyTradesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOpenOrderListRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *float32
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOpenOrderListRequest) RecvWindow(recvWindow float32) ApiOpenOrderListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOpenOrderListRequest) Execute() (*common.RestApiResponse[models.OpenOrderListResponse], error) {
	return r.ApiService.OpenOrderListExecute(r)
}

/*
OpenOrderList Query Open Order lists
Get /api/v3/openOrderList

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-open-order-lists-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOpenOrderListRequest
*/
func (a *AccountAPIService) OpenOrderList(ctx context.Context) ApiOpenOrderListRequest {
	return ApiOpenOrderListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OpenOrderListResponse
func (a *AccountAPIService) OpenOrderListExecute(r ApiOpenOrderListRequest) (*common.RestApiResponse[models.OpenOrderListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/openOrderList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OpenOrderListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOrderAmendmentsRequest struct {
	ctx             context.Context
	ApiService      *AccountAPIService
	symbol          *string
	orderId         *int64
	fromExecutionId *int64
	limit           *int64
	recvWindow      *float32
}

func (r ApiOrderAmendmentsRequest) Symbol(symbol string) ApiOrderAmendmentsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiOrderAmendmentsRequest) OrderId(orderId int64) ApiOrderAmendmentsRequest {
	r.orderId = &orderId
	return r
}

func (r ApiOrderAmendmentsRequest) FromExecutionId(fromExecutionId int64) ApiOrderAmendmentsRequest {
	r.fromExecutionId = &fromExecutionId
	return r
}

// Default:500; Maximum: 1000
func (r ApiOrderAmendmentsRequest) Limit(limit int64) ApiOrderAmendmentsRequest {
	r.limit = &limit
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiOrderAmendmentsRequest) RecvWindow(recvWindow float32) ApiOrderAmendmentsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOrderAmendmentsRequest) Execute() (*common.RestApiResponse[models.OrderAmendmentsResponse], error) {
	return r.ApiService.OrderAmendmentsExecute(r)
}

/*
OrderAmendments Query Order Amendments
Get /api/v3/order/amendments

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-order-amendments-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param orderId -
@param fromExecutionId -
@param limit -  Default:500; Maximum: 1000
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiOrderAmendmentsRequest
*/
func (a *AccountAPIService) OrderAmendments(ctx context.Context) ApiOrderAmendmentsRequest {
	return ApiOrderAmendmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OrderAmendmentsResponse
func (a *AccountAPIService) OrderAmendmentsExecute(r ApiOrderAmendmentsRequest) (*common.RestApiResponse[models.OrderAmendmentsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/order/amendments"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	if r.fromExecutionId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromExecutionId", r.fromExecutionId, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OrderAmendmentsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRateLimitOrderRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *float32
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiRateLimitOrderRequest) RecvWindow(recvWindow float32) ApiRateLimitOrderRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRateLimitOrderRequest) Execute() (*common.RestApiResponse[models.RateLimitOrderResponse], error) {
	return r.ApiService.RateLimitOrderExecute(r)
}

/*
RateLimitOrder Query Unfilled Order Count
Get /api/v3/rateLimit/order

https://developers.binance.com/docs/binance-spot-api-docs/rest-api/account-endpoints#query-unfilled-order-count-user_data

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiRateLimitOrderRequest
*/
func (a *AccountAPIService) RateLimitOrder(ctx context.Context) ApiRateLimitOrderRequest {
	return ApiRateLimitOrderRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RateLimitOrderResponse
func (a *AccountAPIService) RateLimitOrderExecute(r ApiRateLimitOrderRequest) (*common.RestApiResponse[models.RateLimitOrderResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/api/v3/rateLimit/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RateLimitOrderResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
