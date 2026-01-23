/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountFundingFlowRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	currency   *string
	recordId   *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// Asset type, only support USDT  as of now
func (r ApiAccountFundingFlowRequest) Currency(currency string) ApiAccountFundingFlowRequest {
	r.currency = &currency
	return r
}

// Return the recordId and subsequent data, the latest data is returned by default, e.g 100000
func (r ApiAccountFundingFlowRequest) RecordId(recordId int64) ApiAccountFundingFlowRequest {
	r.recordId = &recordId
	return r
}

// Start Time, e.g 1593511200000
func (r ApiAccountFundingFlowRequest) StartTime(startTime int64) ApiAccountFundingFlowRequest {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r ApiAccountFundingFlowRequest) EndTime(endTime int64) ApiAccountFundingFlowRequest {
	r.endTime = &endTime
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r ApiAccountFundingFlowRequest) Limit(limit int64) ApiAccountFundingFlowRequest {
	r.limit = &limit
	return r
}

func (r ApiAccountFundingFlowRequest) RecvWindow(recvWindow int64) ApiAccountFundingFlowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountFundingFlowRequest) Execute() (*common.RestApiResponse[models.AccountFundingFlowResponse], error) {
	return r.ApiService.AccountFundingFlowExecute(r)
}

/*
AccountFundingFlow Account Funding Flow (USER_DATA)
Get /eapi/v1/bill

https://developers.binance.com/docs/derivatives/options-trading/account/Account-Funding-Flow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param currency -  Asset type, only support USDT  as of now
@param recordId -  Return the recordId and subsequent data, the latest data is returned by default, e.g 100000
@param startTime -  Start Time, e.g 1593511200000
@param endTime -  End Time, e.g 1593512200000
@param limit -  Number of result sets returned Default:100 Max:1000
@param recvWindow -
@return ApiAccountFundingFlowRequest
*/
func (a *AccountAPIService) AccountFundingFlow(ctx context.Context) ApiAccountFundingFlowRequest {
	return ApiAccountFundingFlowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountFundingFlowResponse
func (a *AccountAPIService) AccountFundingFlowExecute(r ApiAccountFundingFlowRequest) (*common.RestApiResponse[models.AccountFundingFlowResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/bill"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.currency == nil {
		return nil, common.ReportError("currency is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "form", "")
	if r.recordId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recordId", r.recordId, "form", "")
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

	resp, err := SendRequest[models.AccountFundingFlowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiOptionMarginAccountInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiOptionMarginAccountInformationRequest) RecvWindow(recvWindow int64) ApiOptionMarginAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiOptionMarginAccountInformationRequest) Execute() (*common.RestApiResponse[models.OptionMarginAccountInformationResponse], error) {
	return r.ApiService.OptionMarginAccountInformationExecute(r)
}

/*
OptionMarginAccountInformation Option Margin Account Information (USER_DATA)
Get /eapi/v1/marginAccount

https://developers.binance.com/docs/derivatives/options-trading/account/Option-Margin-Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiOptionMarginAccountInformationRequest
*/
func (a *AccountAPIService) OptionMarginAccountInformation(ctx context.Context) ApiOptionMarginAccountInformationRequest {
	return ApiOptionMarginAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return OptionMarginAccountInformationResponse
func (a *AccountAPIService) OptionMarginAccountInformationExecute(r ApiOptionMarginAccountInformationRequest) (*common.RestApiResponse[models.OptionMarginAccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/eapi/v1/marginAccount"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.OptionMarginAccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
