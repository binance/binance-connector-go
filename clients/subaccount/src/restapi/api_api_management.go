/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package binancesubaccountrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// ApiManagementAPIService ApiManagementAPI Service
type ApiManagementAPIService Service

type ApiAddIpRestrictionForSubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	status           *int64
	ipAddress        *string
	recvWindow       *int64
}

func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) Email(email string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.email = &email
	return r
}

func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiAddIpRestrictionForSubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

// IP Restriction status. 1 &#x3D; IP Unrestricted. 2 &#x3D; Restrict access to trusted IPs only.
func (r ApiAddIpRestrictionForSubAccountApiKeyRequest) Status(status int64) ApiAddIpRestrictionForSubAccountApiKeyRequest {
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

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#add-ip-restriction-for-sub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -
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

	resp, err := SendRequest[models.AddIpRestrictionForSubAccountApiKeyResponse](
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

type ApiCreateSubAccountApiKeyRequest struct {
	ctx                  context.Context
	ApiService           *ApiManagementAPIService
	email                *string
	apiName              *string
	status               *int64
	canTrade             *bool
	canMarginLoanRepay   *bool
	canFuturesTrade      *bool
	canUniversalTransfer *bool
	canVanillaOptions    *bool
	ipAddress            *string
	thirdPartyName       *string
	publicKey            *string
	recvWindow           *int64
}

// Sub-account email
func (r ApiCreateSubAccountApiKeyRequest) Email(email string) ApiCreateSubAccountApiKeyRequest {
	r.email = &email
	return r
}

// API Key name
func (r ApiCreateSubAccountApiKeyRequest) ApiName(apiName string) ApiCreateSubAccountApiKeyRequest {
	r.apiName = &apiName
	return r
}

// IP restriction status. 1 &#x3D; unrestricted, 2 &#x3D; restricted to trusted IPs, 3 &#x3D; third-party IP restriction
func (r ApiCreateSubAccountApiKeyRequest) Status(status int64) ApiCreateSubAccountApiKeyRequest {
	r.status = &status
	return r
}

// Spot &amp; Margin trading permission, default false
func (r ApiCreateSubAccountApiKeyRequest) CanTrade(canTrade bool) ApiCreateSubAccountApiKeyRequest {
	r.canTrade = &canTrade
	return r
}

// Margin borrow/repay permission, default false
func (r ApiCreateSubAccountApiKeyRequest) CanMarginLoanRepay(canMarginLoanRepay bool) ApiCreateSubAccountApiKeyRequest {
	r.canMarginLoanRepay = &canMarginLoanRepay
	return r
}

// Futures trading permission, default false
func (r ApiCreateSubAccountApiKeyRequest) CanFuturesTrade(canFuturesTrade bool) ApiCreateSubAccountApiKeyRequest {
	r.canFuturesTrade = &canFuturesTrade
	return r
}

// Universal transfer permission, default false
func (r ApiCreateSubAccountApiKeyRequest) CanUniversalTransfer(canUniversalTransfer bool) ApiCreateSubAccountApiKeyRequest {
	r.canUniversalTransfer = &canUniversalTransfer
	return r
}

// Vanilla options permission, default false
func (r ApiCreateSubAccountApiKeyRequest) CanVanillaOptions(canVanillaOptions bool) ApiCreateSubAccountApiKeyRequest {
	r.canVanillaOptions = &canVanillaOptions
	return r
}

// Required when status&#x3D;2. IP address list, max 500 chars
func (r ApiCreateSubAccountApiKeyRequest) IpAddress(ipAddress string) ApiCreateSubAccountApiKeyRequest {
	r.ipAddress = &ipAddress
	return r
}

// Required when status&#x3D;3. Third-party name
func (r ApiCreateSubAccountApiKeyRequest) ThirdPartyName(thirdPartyName string) ApiCreateSubAccountApiKeyRequest {
	r.thirdPartyName = &thirdPartyName
	return r
}

// Ed25519 public key (optional, for Ed25519 type API Key)
func (r ApiCreateSubAccountApiKeyRequest) PublicKey(publicKey string) ApiCreateSubAccountApiKeyRequest {
	r.publicKey = &publicKey
	return r
}

func (r ApiCreateSubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiCreateSubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateSubAccountApiKeyRequest) Execute() (*common.RestApiResponse[models.CreateSubAccountApiKeyResponse], error) {
	return r.ApiService.CreateSubAccountApiKeyExecute(r)
}

/*
CreateSubAccountApiKey Create Sub-account API Key (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/subAccountApi

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#create-sub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Sub-account email
@param apiName -  API Key name
@param status -  IP restriction status. 1 = unrestricted, 2 = restricted to trusted IPs, 3 = third-party IP restriction
@param canTrade -  Spot & Margin trading permission, default false
@param canMarginLoanRepay -  Margin borrow/repay permission, default false
@param canFuturesTrade -  Futures trading permission, default false
@param canUniversalTransfer -  Universal transfer permission, default false
@param canVanillaOptions -  Vanilla options permission, default false
@param ipAddress -  Required when status=2. IP address list, max 500 chars
@param thirdPartyName -  Required when status=3. Third-party name
@param publicKey -  Ed25519 public key (optional, for Ed25519 type API Key)
@param recvWindow -
@return ApiCreateSubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) CreateSubAccountApiKey(ctx context.Context) ApiCreateSubAccountApiKeyRequest {
	return ApiCreateSubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateSubAccountApiKeyResponse
func (a *ApiManagementAPIService) CreateSubAccountApiKeyExecute(r ApiCreateSubAccountApiKeyRequest) (*common.RestApiResponse[models.CreateSubAccountApiKeyResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApi"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	if r.apiName == nil {
		return nil, common.ReportError("apiName is required and must be specified")
	}

	if r.status == nil {
		return nil, common.ReportError("status is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "apiName", r.apiName, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	if r.canTrade != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canTrade", r.canTrade, "form", "")
	}
	if r.canMarginLoanRepay != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canMarginLoanRepay", r.canMarginLoanRepay, "form", "")
	}
	if r.canFuturesTrade != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canFuturesTrade", r.canFuturesTrade, "form", "")
	}
	if r.canUniversalTransfer != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canUniversalTransfer", r.canUniversalTransfer, "form", "")
	}
	if r.canVanillaOptions != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canVanillaOptions", r.canVanillaOptions, "form", "")
	}
	if r.ipAddress != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "form", "")
	}
	if r.thirdPartyName != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "thirdPartyName", r.thirdPartyName, "form", "")
	}
	if r.publicKey != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "publicKey", r.publicKey, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateSubAccountApiKeyResponse](
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

type ApiDeleteIpListForASubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	ipAddress        *string
	recvWindow       *int64
}

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

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#delete-ip-list-for-asub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -
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

	resp, err := SendRequest[models.DeleteIpListForASubAccountApiKeyResponse](
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

type ApiDeleteSubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	recvWindow       *int64
}

// Sub-account email
func (r ApiDeleteSubAccountApiKeyRequest) Email(email string) ApiDeleteSubAccountApiKeyRequest {
	r.email = &email
	return r
}

// The sub-account API Key to be deleted
func (r ApiDeleteSubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiDeleteSubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

func (r ApiDeleteSubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiDeleteSubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDeleteSubAccountApiKeyRequest) Execute() (*common.RestApiResponse[map[string]interface{}], error) {
	return r.ApiService.DeleteSubAccountApiKeyExecute(r)
}

/*
DeleteSubAccountApiKey Delete Sub-account API Key (For Master Account) (USER_DATA)
Delete /sapi/v1/sub-account/subAccountApi

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#delete-sub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Sub-account email
@param subAccountApiKey -  The sub-account API Key to be deleted
@param recvWindow -
@return ApiDeleteSubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) DeleteSubAccountApiKey(ctx context.Context) ApiDeleteSubAccountApiKeyRequest {
	return ApiDeleteSubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *ApiManagementAPIService) DeleteSubAccountApiKeyExecute(r ApiDeleteSubAccountApiKeyRequest) (*common.RestApiResponse[map[string]interface{}], error) {
	localVarHTTPMethod := http.MethodDelete
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApi"

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

	resp, err := SendRequest[map[string]interface{}](
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

type ApiGetIpRestrictionForASubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	recvWindow       *int64
}

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

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#get-ip-restriction-for-asub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -
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

	resp, err := SendRequest[models.GetIpRestrictionForASubAccountApiKeyResponse](
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

type ApiModifySubAccountApiKeyPermissionRequest struct {
	ctx                  context.Context
	ApiService           *ApiManagementAPIService
	email                *string
	subAccountApiKey     *string
	canTrade             *bool
	canMarginLoanRepay   *bool
	canFuturesTrade      *bool
	canUniversalTransfer *bool
	canVanillaOptions    *bool
	recvWindow           *int64
}

// Sub-account email
func (r ApiModifySubAccountApiKeyPermissionRequest) Email(email string) ApiModifySubAccountApiKeyPermissionRequest {
	r.email = &email
	return r
}

// Sub-account API Key
func (r ApiModifySubAccountApiKeyPermissionRequest) SubAccountApiKey(subAccountApiKey string) ApiModifySubAccountApiKeyPermissionRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

// Spot &amp; Margin trading permission
func (r ApiModifySubAccountApiKeyPermissionRequest) CanTrade(canTrade bool) ApiModifySubAccountApiKeyPermissionRequest {
	r.canTrade = &canTrade
	return r
}

// Margin borrow/repay permission
func (r ApiModifySubAccountApiKeyPermissionRequest) CanMarginLoanRepay(canMarginLoanRepay bool) ApiModifySubAccountApiKeyPermissionRequest {
	r.canMarginLoanRepay = &canMarginLoanRepay
	return r
}

// Futures trading permission
func (r ApiModifySubAccountApiKeyPermissionRequest) CanFuturesTrade(canFuturesTrade bool) ApiModifySubAccountApiKeyPermissionRequest {
	r.canFuturesTrade = &canFuturesTrade
	return r
}

// Universal transfer permission
func (r ApiModifySubAccountApiKeyPermissionRequest) CanUniversalTransfer(canUniversalTransfer bool) ApiModifySubAccountApiKeyPermissionRequest {
	r.canUniversalTransfer = &canUniversalTransfer
	return r
}

// Vanilla options permission
func (r ApiModifySubAccountApiKeyPermissionRequest) CanVanillaOptions(canVanillaOptions bool) ApiModifySubAccountApiKeyPermissionRequest {
	r.canVanillaOptions = &canVanillaOptions
	return r
}

func (r ApiModifySubAccountApiKeyPermissionRequest) RecvWindow(recvWindow int64) ApiModifySubAccountApiKeyPermissionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiModifySubAccountApiKeyPermissionRequest) Execute() (*common.RestApiResponse[models.ModifySubAccountApiKeyPermissionResponse], error) {
	return r.ApiService.ModifySubAccountApiKeyPermissionExecute(r)
}

/*
ModifySubAccountApiKeyPermission Modify Sub-account API Key Permission (For Master Account) (USER_DATA)
Post /sapi/v1/sub-account/subAccountApiPermission

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#modify-sub-account-api-key-permission

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Sub-account email
@param subAccountApiKey -  Sub-account API Key
@param canTrade -  Spot & Margin trading permission
@param canMarginLoanRepay -  Margin borrow/repay permission
@param canFuturesTrade -  Futures trading permission
@param canUniversalTransfer -  Universal transfer permission
@param canVanillaOptions -  Vanilla options permission
@param recvWindow -
@return ApiModifySubAccountApiKeyPermissionRequest
*/
func (a *ApiManagementAPIService) ModifySubAccountApiKeyPermission(ctx context.Context) ApiModifySubAccountApiKeyPermissionRequest {
	return ApiModifySubAccountApiKeyPermissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ModifySubAccountApiKeyPermissionResponse
func (a *ApiManagementAPIService) ModifySubAccountApiKeyPermissionExecute(r ApiModifySubAccountApiKeyPermissionRequest) (*common.RestApiResponse[models.ModifySubAccountApiKeyPermissionResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApiPermission"

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
	if r.canTrade != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canTrade", r.canTrade, "form", "")
	}
	if r.canMarginLoanRepay != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canMarginLoanRepay", r.canMarginLoanRepay, "form", "")
	}
	if r.canFuturesTrade != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canFuturesTrade", r.canFuturesTrade, "form", "")
	}
	if r.canUniversalTransfer != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canUniversalTransfer", r.canUniversalTransfer, "form", "")
	}
	if r.canVanillaOptions != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "canVanillaOptions", r.canVanillaOptions, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ModifySubAccountApiKeyPermissionResponse](
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

type ApiQuerySubAccountApiKeyRequest struct {
	ctx              context.Context
	ApiService       *ApiManagementAPIService
	email            *string
	subAccountApiKey *string
	page             *int64
	size             *int64
	recvWindow       *int64
}

// Sub-account email
func (r ApiQuerySubAccountApiKeyRequest) Email(email string) ApiQuerySubAccountApiKeyRequest {
	r.email = &email
	return r
}

// Specify an API Key for exact match
func (r ApiQuerySubAccountApiKeyRequest) SubAccountApiKey(subAccountApiKey string) ApiQuerySubAccountApiKeyRequest {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

// Page number, default 1, minimum 1
func (r ApiQuerySubAccountApiKeyRequest) Page(page int64) ApiQuerySubAccountApiKeyRequest {
	r.page = &page
	return r
}

// Page size, default 30, maximum 100
func (r ApiQuerySubAccountApiKeyRequest) Size(size int64) ApiQuerySubAccountApiKeyRequest {
	r.size = &size
	return r
}

func (r ApiQuerySubAccountApiKeyRequest) RecvWindow(recvWindow int64) ApiQuerySubAccountApiKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQuerySubAccountApiKeyRequest) Execute() (*common.RestApiResponse[models.QuerySubAccountApiKeyResponse], error) {
	return r.ApiService.QuerySubAccountApiKeyExecute(r)
}

/*
QuerySubAccountApiKey Query Sub-account API Key (For Master Account) (USER_DATA)
Get /sapi/v1/sub-account/subAccountApi

https://developers.binance.com/en/docs/catalog/vip-and-institutional-sub-account/api/rest-api/api-management#query-sub-account-api-key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -  Sub-account email
@param subAccountApiKey -  Specify an API Key for exact match
@param page -  Page number, default 1, minimum 1
@param size -  Page size, default 30, maximum 100
@param recvWindow -
@return ApiQuerySubAccountApiKeyRequest
*/
func (a *ApiManagementAPIService) QuerySubAccountApiKey(ctx context.Context) ApiQuerySubAccountApiKeyRequest {
	return ApiQuerySubAccountApiKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QuerySubAccountApiKeyResponse
func (a *ApiManagementAPIService) QuerySubAccountApiKeyExecute(r ApiQuerySubAccountApiKeyRequest) (*common.RestApiResponse[models.QuerySubAccountApiKeyResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/sub-account/subAccountApi"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	if r.subAccountApiKey != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QuerySubAccountApiKeyResponse](
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
