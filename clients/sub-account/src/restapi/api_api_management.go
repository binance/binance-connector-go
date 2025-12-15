/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package binancesubaccountrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// ApiManagementAPIService ApiManagementAPI Service
type ApiManagementAPIService Service

type ApiAddIpRestrictionForSubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	status           *string
	ipAddress        *string
	recvWindow       *int64
}

// [Sub-account email](#email-address)
func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) Email(email string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.email = &email
	return r
}

func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

// IP Restriction status. 1 &#x3D; IP Unrestricted. 2 &#x3D; Restrict access to trusted IPs only.
func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) Status(status string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.status = &status
	return r
}

// Insert static IP in batch, separated by commas.
func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) IpAddress(ipAddress string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.ipAddress = &ipAddress
	return r
}

func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) Execute() (*common.RestApiResponse[models.AddIpRestrictionForSubAccountApiKeyResponse], error) {
	return r.ApiService.AddIpRestrictionForSubAccountApiKeyExecute(r)
}

/*
AddIpRestrictionForSubAccountApiKey Add IP Restriction for Sub-Account API key (For Master Account) (USER_DATA)
Post /sapi/v2/sub-account/subAccountApi/ipRestriction

https://developers.binance.com/docs/sub_account/api-management/Add-IP-Restriction-for-Sub-Account-API-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param subAccountApiKey -
@param status -  IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
@param ipAddress -  Insert static IP in batch, separated by commas.
@param recvWindow -
@return ApiAddIpRestrictionForSubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) AddIpRestrictionForSubAccountApiKey(ctx context.Context) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	return ApiAddIpRestrictionForSubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddIpRestrictionForSubAccountApiKeyResponse
func (a *ApiManagementAPIService) AddIpRestrictionForSubAccountApiKeyExecute(r ApiAddIpRestrictionForSubAccountApiKeyRequest) (*common.RestApiResponse[models.AddIpRestrictionForSubAccountApiKeyResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/sub-account/subAccountApi/ipRestriction"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return nil, common.ReportError("subAccountApiKey is required and must be specified")
	}
	if r.status == nil {
		return nil, common.ReportError("status is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	if r.ipAddress != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AddIpRestrictionForSubAccountApiKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDeleteIpListForASubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	ipAddress        *string
	recvWindow       *int64
}

// [Sub-account email](#email-address)
func (r ApiDeleteIpListForASubAccountApiKeyRequest) Email(email string) ApiDeleteIpListForASubAccountApiKeyRequest {
	r.email = &email
	return r
}

func (r ApiDeleteIpListForASubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiDeleteIpListForASubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

// IPs to be deleted. Can be added in batches, separated by commas
func (r ApiDeleteIpListForASubAccountApiKeyRequest) IpAddress(ipAddress string) ApiDeleteIpListForASubAccountApiKeyRequest {
	r.ipAddress = &ipAddress
	return r
}

func (r ApiDeleteIpListForASubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiDeleteIpListForASubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteIpListForASubAccountApiKeyRequest) Execute() (*common.RestApiResponse[models.DeleteIpListForASubAccountApiKeyResponse], error) {
	return r.ApiService.DeleteIpListForASubAccountApiKeyExecute(r)
}

/*
DeleteIpListForASubAccountApiKey Delete IP List For a Sub-account API Key (For Master Account) (USER_DATA)
Delete /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList

https://developers.binance.com/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param subAccountApiKey -
@param ipAddress -  IPs to be deleted. Can be added in batches, separated by commas
@param recvWindow -
@return ApiDeleteIpListForASubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) DeleteIpListForASubAccountApiKey(ctx context.Context) ApiDeleteIpListForASubAccountApiKeyRequest {
	return ApiDeleteIpListForASubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DeleteIpListForASubAccountApiKeyResponse
func (a *ApiManagementAPIService) DeleteIpListForASubAccountApiKeyExecute(r ApiDeleteIpListForASubAccountApiKeyRequest) (*common.RestApiResponse[models.DeleteIpListForASubAccountApiKeyResponse], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return nil, common.ReportError("subAccountApiKey is required and must be specified")
	}
	if r.ipAddress == nil {
		return nil, common.ReportError("ipAddress is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DeleteIpListForASubAccountApiKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetIpRestrictionForASubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	recvWindow       *int64
}

// [Sub-account email](#email-address)
func (r ApiGetIpRestrictionForASubAccountApiKeyRequest) Email(email string) ApiGetIpRestrictionForASubAccountApiKeyRequest {
	r.email = &email
	return r
}

func (r ApiGetIpRestrictionForASubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiGetIpRestrictionForASubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

func (r ApiGetIpRestrictionForASubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiGetIpRestrictionForASubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetIpRestrictionForASubAccountApiKeyRequest) Execute() (*common.RestApiResponse[models.GetIpRestrictionForASubAccountApiKeyResponse], error) {
	return r.ApiService.GetIpRestrictionForASubAccountApiKeyExecute(r)
}

/*
GetIpRestrictionForASubAccountApiKey Get IP Restriction for a Sub-account API Key (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/subAccountApi/ipRestriction

https://developers.binance.com/docs/sub_account/api-management/Get-IP-Restriction-for-a-Sub-account-API-Key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  [Sub-account email](#email-address)
@param subAccountApiKey -
@param recvWindow -
@return ApiGetIpRestrictionForASubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) GetIpRestrictionForASubAccountApiKey(ctx context.Context) ApiGetIpRestrictionForASubAccountApiKeyRequest {
	return ApiGetIpRestrictionForASubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetIpRestrictionForASubAccountApiKeyResponse
func (a *ApiManagementAPIService) GetIpRestrictionForASubAccountApiKeyExecute(r ApiGetIpRestrictionForASubAccountApiKeyRequest) (*common.RestApiResponse[models.GetIpRestrictionForASubAccountApiKeyResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApi/ipRestriction"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return nil, common.ReportError("subAccountApiKey is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetIpRestrictionForASubAccountApiKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
