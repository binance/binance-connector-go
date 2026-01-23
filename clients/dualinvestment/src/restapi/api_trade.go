/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package binancedualinvestmentrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/dualinvestment/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TradeAPIService TradeAPI Service
type TradeAPIService Service

type ApiChangeAutoCompoundStatusRequest struct {
	ctx              context.Context
	ApiService       *TradeAPIService
	positionId       *string
	autoCompoundPlan *string
	recvWindow       *int64
}

// Get positionId from &#x60;/sapi/v1/dci/product/positions&#x60;
func (r ApiChangeAutoCompoundStatusRequest) PositionId(positionId string) ApiChangeAutoCompoundStatusRequest {
	r.positionId = &positionId
	return r
}

func (r ApiChangeAutoCompoundStatusRequest) AutoCompoundPlan(autoCompoundPlan string) ApiChangeAutoCompoundStatusRequest {
	r.autoCompoundPlan = &autoCompoundPlan
	return r
}

// The value cannot be greater than 60000
func (r ApiChangeAutoCompoundStatusRequest) RecvWindow(recvWindow int64) ApiChangeAutoCompoundStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeAutoCompoundStatusRequest) Execute() (*common.RestApiResponse[models.ChangeAutoCompoundStatusResponse], error) {
	return r.ApiService.ChangeAutoCompoundStatusExecute(r)
}

/*
ChangeAutoCompoundStatus Change Auto-Compound status(USER_DATA)
Post /sapi/v1/dci/product/auto_compound/edit-status

https://developers.binance.com/docs/advanced_earn/dual-investment/trade/Change-Auto-Compound-status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param positionId -  Get positionId from `/sapi/v1/dci/product/positions`
@param autoCompoundPlan -
@param recvWindow -  The value cannot be greater than 60000
@return ApiChangeAutoCompoundStatusRequest
*/
func (a *TradeAPIService) ChangeAutoCompoundStatus(ctx context.Context) ApiChangeAutoCompoundStatusRequest {
	return ApiChangeAutoCompoundStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeAutoCompoundStatusResponse
func (a *TradeAPIService) ChangeAutoCompoundStatusExecute(r ApiChangeAutoCompoundStatusRequest) (*common.RestApiResponse[models.ChangeAutoCompoundStatusResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/dci/product/auto_compound/edit-status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.positionId == nil {
		return nil, common.ReportError("positionId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "positionId", r.positionId, "form", "")
	if r.autoCompoundPlan != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "AutoCompoundPlan", r.autoCompoundPlan, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeAutoCompoundStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCheckDualInvestmentAccountsRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	recvWindow *int64
}

// The value cannot be greater than 60000
func (r ApiCheckDualInvestmentAccountsRequest) RecvWindow(recvWindow int64) ApiCheckDualInvestmentAccountsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCheckDualInvestmentAccountsRequest) Execute() (*common.RestApiResponse[models.CheckDualInvestmentAccountsResponse], error) {
	return r.ApiService.CheckDualInvestmentAccountsExecute(r)
}

/*
CheckDualInvestmentAccounts Check Dual Investment accounts(USER_DATA)
Get /sapi/v1/dci/product/accounts

https://developers.binance.com/docs/advanced_earn/dual-investment/trade/Check-Dual-Investment-accounts

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -  The value cannot be greater than 60000
@return ApiCheckDualInvestmentAccountsRequest
*/
func (a *TradeAPIService) CheckDualInvestmentAccounts(ctx context.Context) ApiCheckDualInvestmentAccountsRequest {
	return ApiCheckDualInvestmentAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CheckDualInvestmentAccountsResponse
func (a *TradeAPIService) CheckDualInvestmentAccountsExecute(r ApiCheckDualInvestmentAccountsRequest) (*common.RestApiResponse[models.CheckDualInvestmentAccountsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/dci/product/accounts"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CheckDualInvestmentAccountsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDualInvestmentPositionsRequest struct {
	ctx        context.Context
	ApiService *TradeAPIService
	status     *string
	pageSize   *int64
	pageIndex  *int64
	recvWindow *int64
}

// &#x60;PENDING&#x60;:Products are purchasing, will give results later;&#x60;PURCHASE_SUCCESS&#x60;:purchase successfully;&#x60;SETTLED&#x60;: Products are finish settling;&#x60;PURCHASE_FAIL&#x60;:fail to purchase;&#x60;REFUNDING&#x60;:refund ongoing;&#x60;REFUND_SUCCESS&#x60;:refund to spot account successfully; &#x60;SETTLING&#x60;:Products are settling. If don&#39;t fill this field, will response all the position status.
func (r ApiGetDualInvestmentPositionsRequest) Status(status string) ApiGetDualInvestmentPositionsRequest {
	r.status = &status
	return r
}

// Default: 10, Maximum: 100
func (r ApiGetDualInvestmentPositionsRequest) PageSize(pageSize int64) ApiGetDualInvestmentPositionsRequest {
	r.pageSize = &pageSize
	return r
}

// Default: 1
func (r ApiGetDualInvestmentPositionsRequest) PageIndex(pageIndex int64) ApiGetDualInvestmentPositionsRequest {
	r.pageIndex = &pageIndex
	return r
}

// The value cannot be greater than 60000
func (r ApiGetDualInvestmentPositionsRequest) RecvWindow(recvWindow int64) ApiGetDualInvestmentPositionsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDualInvestmentPositionsRequest) Execute() (*common.RestApiResponse[models.GetDualInvestmentPositionsResponse], error) {
	return r.ApiService.GetDualInvestmentPositionsExecute(r)
}

/*
GetDualInvestmentPositions Get Dual Investment positions(USER_DATA)
Get /sapi/v1/dci/product/positions

https://developers.binance.com/docs/advanced_earn/dual-investment/trade/Get-Dual-Investment-positions

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param status -  `PENDING`:Products are purchasing, will give results later;`PURCHASE_SUCCESS`:purchase successfully;`SETTLED`: Products are finish settling;`PURCHASE_FAIL`:fail to purchase;`REFUNDING`:refund ongoing;`REFUND_SUCCESS`:refund to spot account successfully; `SETTLING`:Products are settling. If don't fill this field, will response all the position status.
@param pageSize -  Default: 10, Maximum: 100
@param pageIndex -  Default: 1
@param recvWindow -  The value cannot be greater than 60000
@return ApiGetDualInvestmentPositionsRequest
*/
func (a *TradeAPIService) GetDualInvestmentPositions(ctx context.Context) ApiGetDualInvestmentPositionsRequest {
	return ApiGetDualInvestmentPositionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDualInvestmentPositionsResponse
func (a *TradeAPIService) GetDualInvestmentPositionsExecute(r ApiGetDualInvestmentPositionsRequest) (*common.RestApiResponse[models.GetDualInvestmentPositionsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/dci/product/positions"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.status != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDualInvestmentPositionsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubscribeDualInvestmentProductsRequest struct {
	ctx              context.Context
	ApiService       *TradeAPIService
	id               *string
	orderId          *string
	depositAmount    *float32
	autoCompoundPlan *string
	recvWindow       *int64
}

// get id from &#x60;/sapi/v1/dci/product/list&#x60;
func (r ApiSubscribeDualInvestmentProductsRequest) Id(id string) ApiSubscribeDualInvestmentProductsRequest {
	r.id = &id
	return r
}

// get orderId from &#x60;/sapi/v1/dci/product/list&#x60;
func (r ApiSubscribeDualInvestmentProductsRequest) OrderId(orderId string) ApiSubscribeDualInvestmentProductsRequest {
	r.orderId = &orderId
	return r
}

// the amount for subscribing
func (r ApiSubscribeDualInvestmentProductsRequest) DepositAmount(depositAmount float32) ApiSubscribeDualInvestmentProductsRequest {
	r.depositAmount = &depositAmount
	return r
}

// &#x60;NONE&#x60;: switch off the plan, &#x60;STANDARD&#x60;:standard plan,&#x60;ADVANCED&#x60;:advanced plan
func (r ApiSubscribeDualInvestmentProductsRequest) AutoCompoundPlan(autoCompoundPlan string) ApiSubscribeDualInvestmentProductsRequest {
	r.autoCompoundPlan = &autoCompoundPlan
	return r
}

// The value cannot be greater than 60000
func (r ApiSubscribeDualInvestmentProductsRequest) RecvWindow(recvWindow int64) ApiSubscribeDualInvestmentProductsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubscribeDualInvestmentProductsRequest) Execute() (*common.RestApiResponse[models.SubscribeDualInvestmentProductsResponse], error) {
	return r.ApiService.SubscribeDualInvestmentProductsExecute(r)
}

/*
SubscribeDualInvestmentProducts Subscribe Dual Investment products(USER_DATA)
Post /sapi/v1/dci/product/subscribe

https://developers.binance.com/docs/advanced_earn/dual-investment/trade/Subscribe-Dual-Investment-products

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param id -  get id from `/sapi/v1/dci/product/list`
@param orderId -  get orderId from `/sapi/v1/dci/product/list`
@param depositAmount -  the amount for subscribing
@param autoCompoundPlan -  `NONE`: switch off the plan, `STANDARD`:standard plan,`ADVANCED`:advanced plan
@param recvWindow -  The value cannot be greater than 60000
@return ApiSubscribeDualInvestmentProductsRequest
*/
func (a *TradeAPIService) SubscribeDualInvestmentProducts(ctx context.Context) ApiSubscribeDualInvestmentProductsRequest {
	return ApiSubscribeDualInvestmentProductsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscribeDualInvestmentProductsResponse
func (a *TradeAPIService) SubscribeDualInvestmentProductsExecute(r ApiSubscribeDualInvestmentProductsRequest) (*common.RestApiResponse[models.SubscribeDualInvestmentProductsResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/dci/product/subscribe"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.id == nil {
		return nil, common.ReportError("id is required and must be specified")
	}
	if r.orderId == nil {
		return nil, common.ReportError("orderId is required and must be specified")
	}
	if r.depositAmount == nil {
		return nil, common.ReportError("depositAmount is required and must be specified")
	}
	if r.autoCompoundPlan == nil {
		return nil, common.ReportError("autoCompoundPlan is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "orderId", r.orderId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "depositAmount", r.depositAmount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoCompoundPlan", r.autoCompoundPlan, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.SubscribeDualInvestmentProductsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
