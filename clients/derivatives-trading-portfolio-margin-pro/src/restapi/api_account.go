/*
Binance Derivatives Trading Portfolio Margin Pro REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro REST API

API version: 1.0.0
*/

package binancederivativestradingportfoliomarginprorestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomarginpro/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiBnbTransferRequest struct {
	ctx          context.Context
	ApiService   *AccountAPIService
	amount       *float32
	transferSide *string
	recvWindow   *int64
}

func (r ApiBnbTransferRequest) Amount(amount float32) ApiBnbTransferRequest {
	r.amount = &amount
	return r
}

// \&quot;TO_UM\&quot;,\&quot;FROM_UM\&quot;
func (r ApiBnbTransferRequest) TransferSide(transferSide string) ApiBnbTransferRequest {
	r.transferSide = &transferSide
	return r
}

func (r ApiBnbTransferRequest) RecvWindow(recvWindow int64) ApiBnbTransferRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiBnbTransferRequest) Execute() (*common.RestApiResponse[models.BnbTransferResponse], error) {
	return r.ApiService.BnbTransferExecute(r)
}

/*
BnbTransfer BNB transfer(USER_DATA)
Post /sapi/v1/portfolio/bnb-transfer

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/BNB-transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param amount -
@param transferSide -  \"TO_UM\",\"FROM_UM\"
@param recvWindow -
@return ApiBnbTransferRequest
*/
func (a *AccountAPIService) BnbTransfer(ctx context.Context) ApiBnbTransferRequest {
	return ApiBnbTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BnbTransferResponse
func (a *AccountAPIService) BnbTransferExecute(r ApiBnbTransferRequest) (*common.RestApiResponse[models.BnbTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/bnb-transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.transferSide == nil {
		return nil, common.ReportError("transferSide is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferSide", r.transferSide, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.BnbTransferResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiChangeAutoRepayFuturesStatusRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	autoRepay  *string
	recvWindow *int64
}

// Default: &#x60;true&#x60;; &#x60;false&#x60; for turn off the auto-repay futures negative balance function
func (r ApiChangeAutoRepayFuturesStatusRequest) AutoRepay(autoRepay string) ApiChangeAutoRepayFuturesStatusRequest {
	r.autoRepay = &autoRepay
	return r
}

func (r ApiChangeAutoRepayFuturesStatusRequest) RecvWindow(recvWindow int64) ApiChangeAutoRepayFuturesStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeAutoRepayFuturesStatusRequest) Execute() (*common.RestApiResponse[models.ChangeAutoRepayFuturesStatusResponse], error) {
	return r.ApiService.ChangeAutoRepayFuturesStatusExecute(r)
}

/*
ChangeAutoRepayFuturesStatus Change Auto-repay-futures Status(TRADE)
Post /sapi/v1/portfolio/repay-futures-switch

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Change-Auto-repay-futures-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param autoRepay -  Default: `true`; `false` for turn off the auto-repay futures negative balance function
@param recvWindow -
@return ApiChangeAutoRepayFuturesStatusRequest
*/
func (a *AccountAPIService) ChangeAutoRepayFuturesStatus(ctx context.Context) ApiChangeAutoRepayFuturesStatusRequest {
	return ApiChangeAutoRepayFuturesStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeAutoRepayFuturesStatusResponse
func (a *AccountAPIService) ChangeAutoRepayFuturesStatusExecute(r ApiChangeAutoRepayFuturesStatusRequest) (*common.RestApiResponse[models.ChangeAutoRepayFuturesStatusResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/repay-futures-switch"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.autoRepay == nil {
		return nil, common.ReportError("autoRepay is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "autoRepay", r.autoRepay, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeAutoRepayFuturesStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFundAutoCollectionRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiFundAutoCollectionRequest) RecvWindow(recvWindow int64) ApiFundAutoCollectionRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFundAutoCollectionRequest) Execute() (*common.RestApiResponse[models.FundAutoCollectionResponse], error) {
	return r.ApiService.FundAutoCollectionExecute(r)
}

/*
FundAutoCollection Fund Auto-collection(USER_DATA)
Post /sapi/v1/portfolio/auto-collection

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Fund-Auto-collection

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFundAutoCollectionRequest
*/
func (a *AccountAPIService) FundAutoCollection(ctx context.Context) ApiFundAutoCollectionRequest {
	return ApiFundAutoCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FundAutoCollectionResponse
func (a *AccountAPIService) FundAutoCollectionExecute(r ApiFundAutoCollectionRequest) (*common.RestApiResponse[models.FundAutoCollectionResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/auto-collection"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FundAutoCollectionResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFundCollectionByAssetRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	recvWindow *int64
}

// &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60;
func (r ApiFundCollectionByAssetRequest) Asset(asset string) ApiFundCollectionByAssetRequest {
	r.asset = &asset
	return r
}

func (r ApiFundCollectionByAssetRequest) RecvWindow(recvWindow int64) ApiFundCollectionByAssetRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFundCollectionByAssetRequest) Execute() (*common.RestApiResponse[models.FundCollectionByAssetResponse], error) {
	return r.ApiService.FundCollectionByAssetExecute(r)
}

/*
FundCollectionByAsset Fund Collection by Asset(USER_DATA)
Post /sapi/v1/portfolio/asset-collection

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Fund-Collection-by-Asset

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  `LDUSDT` and `RWUSD`
@param recvWindow -
@return ApiFundCollectionByAssetRequest
*/
func (a *AccountAPIService) FundCollectionByAsset(ctx context.Context) ApiFundCollectionByAssetRequest {
	return ApiFundCollectionByAssetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FundCollectionByAssetResponse
func (a *AccountAPIService) FundCollectionByAssetExecute(r ApiFundCollectionByAssetRequest) (*common.RestApiResponse[models.FundCollectionByAssetResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/asset-collection"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FundCollectionByAssetResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAutoRepayFuturesStatusRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetAutoRepayFuturesStatusRequest) RecvWindow(recvWindow int64) ApiGetAutoRepayFuturesStatusRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetAutoRepayFuturesStatusRequest) Execute() (*common.RestApiResponse[models.GetAutoRepayFuturesStatusResponse], error) {
	return r.ApiService.GetAutoRepayFuturesStatusExecute(r)
}

/*
GetAutoRepayFuturesStatus Get Auto-repay-futures Status(USER_DATA)
Get /sapi/v1/portfolio/repay-futures-switch

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Get-Auto-repay-futures-Status

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetAutoRepayFuturesStatusRequest
*/
func (a *AccountAPIService) GetAutoRepayFuturesStatus(ctx context.Context) ApiGetAutoRepayFuturesStatusRequest {
	return ApiGetAutoRepayFuturesStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAutoRepayFuturesStatusResponse
func (a *AccountAPIService) GetAutoRepayFuturesStatusExecute(r ApiGetAutoRepayFuturesStatusRequest) (*common.RestApiResponse[models.GetAutoRepayFuturesStatusResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/repay-futures-switch"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetAutoRepayFuturesStatusResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetPortfolioMarginProAccountBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	recvWindow *int64
}

func (r ApiGetPortfolioMarginProAccountBalanceRequest) Asset(asset string) ApiGetPortfolioMarginProAccountBalanceRequest {
	r.asset = &asset
	return r
}

func (r ApiGetPortfolioMarginProAccountBalanceRequest) RecvWindow(recvWindow int64) ApiGetPortfolioMarginProAccountBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPortfolioMarginProAccountBalanceRequest) Execute() (*common.RestApiResponse[models.GetPortfolioMarginProAccountBalanceResponse], error) {
	return r.ApiService.GetPortfolioMarginProAccountBalanceExecute(r)
}

/*
GetPortfolioMarginProAccountBalance Get Portfolio Margin Pro Account Balance(USER_DATA)
Get /sapi/v1/portfolio/balance

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Get-Classic-Portfolio-Margin-Balance-Info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param recvWindow -
@return ApiGetPortfolioMarginProAccountBalanceRequest
*/
func (a *AccountAPIService) GetPortfolioMarginProAccountBalance(ctx context.Context) ApiGetPortfolioMarginProAccountBalanceRequest {
	return ApiGetPortfolioMarginProAccountBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPortfolioMarginProAccountBalanceResponse
func (a *AccountAPIService) GetPortfolioMarginProAccountBalanceExecute(r ApiGetPortfolioMarginProAccountBalanceRequest) (*common.RestApiResponse[models.GetPortfolioMarginProAccountBalanceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPortfolioMarginProAccountBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetPortfolioMarginProAccountInfoRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetPortfolioMarginProAccountInfoRequest) RecvWindow(recvWindow int64) ApiGetPortfolioMarginProAccountInfoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPortfolioMarginProAccountInfoRequest) Execute() (*common.RestApiResponse[models.GetPortfolioMarginProAccountInfoResponse], error) {
	return r.ApiService.GetPortfolioMarginProAccountInfoExecute(r)
}

/*
GetPortfolioMarginProAccountInfo Get Portfolio Margin Pro Account Info(USER_DATA)
Get /sapi/v1/portfolio/account

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Get-Classic-Portfolio-Margin-Account-Info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetPortfolioMarginProAccountInfoRequest
*/
func (a *AccountAPIService) GetPortfolioMarginProAccountInfo(ctx context.Context) ApiGetPortfolioMarginProAccountInfoRequest {
	return ApiGetPortfolioMarginProAccountInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPortfolioMarginProAccountInfoResponse
func (a *AccountAPIService) GetPortfolioMarginProAccountInfoExecute(r ApiGetPortfolioMarginProAccountInfoRequest) (*common.RestApiResponse[models.GetPortfolioMarginProAccountInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPortfolioMarginProAccountInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetPortfolioMarginProSpanAccountInfoRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetPortfolioMarginProSpanAccountInfoRequest) RecvWindow(recvWindow int64) ApiGetPortfolioMarginProSpanAccountInfoRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetPortfolioMarginProSpanAccountInfoRequest) Execute() (*common.RestApiResponse[models.GetPortfolioMarginProSpanAccountInfoResponse], error) {
	return r.ApiService.GetPortfolioMarginProSpanAccountInfoExecute(r)
}

/*
GetPortfolioMarginProSpanAccountInfo Get Portfolio Margin Pro SPAN Account Info(USER_DATA)
Get /sapi/v2/portfolio/account

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Get-Classic-Portfolio-Margin-Account-Info-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetPortfolioMarginProSpanAccountInfoRequest
*/
func (a *AccountAPIService) GetPortfolioMarginProSpanAccountInfo(ctx context.Context) ApiGetPortfolioMarginProSpanAccountInfoRequest {
	return ApiGetPortfolioMarginProSpanAccountInfoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPortfolioMarginProSpanAccountInfoResponse
func (a *AccountAPIService) GetPortfolioMarginProSpanAccountInfoExecute(r ApiGetPortfolioMarginProSpanAccountInfoRequest) (*common.RestApiResponse[models.GetPortfolioMarginProSpanAccountInfoResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/portfolio/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetPortfolioMarginProSpanAccountInfoResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest struct {
	ctx          context.Context
	ApiService   *AccountAPIService
	asset        *string
	transferType *string
	recvWindow   *int64
}

// &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60;
func (r ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest) Asset(asset string) ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest {
	r.asset = &asset
	return r
}

// &#x60;EARN_TO_FUTURE&#x60; /&#x60;FUTURE_TO_EARN&#x60;
func (r ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest) TransferType(transferType string) ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest {
	r.transferType = &transferType
	return r
}

func (r ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest) RecvWindow(recvWindow int64) ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest) Execute() (*common.RestApiResponse[models.GetTransferableEarnAssetBalanceForPortfolioMarginResponse], error) {
	return r.ApiService.GetTransferableEarnAssetBalanceForPortfolioMarginExecute(r)
}

/*
GetTransferableEarnAssetBalanceForPortfolioMargin Get Transferable Earn Asset Balance for Portfolio Margin (USER_DATA)
Get /sapi/v1/portfolio/earn-asset-balance

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Get-Transferable-Earn-Asset-Balance-for-Portfolio-Margin

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  `LDUSDT` and `RWUSD`
@param transferType -  `EARN_TO_FUTURE` /`FUTURE_TO_EARN`
@param recvWindow -
@return ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest
*/
func (a *AccountAPIService) GetTransferableEarnAssetBalanceForPortfolioMargin(ctx context.Context) ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest {
	return ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTransferableEarnAssetBalanceForPortfolioMarginResponse
func (a *AccountAPIService) GetTransferableEarnAssetBalanceForPortfolioMarginExecute(r ApiGetTransferableEarnAssetBalanceForPortfolioMarginRequest) (*common.RestApiResponse[models.GetTransferableEarnAssetBalanceForPortfolioMarginResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/earn-asset-balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.transferType == nil {
		return nil, common.ReportError("transferType is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferType", r.transferType, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetTransferableEarnAssetBalanceForPortfolioMarginResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPortfolioMarginProBankruptcyLoanRepayRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	from       *string
	recvWindow *int64
}

// SPOT or MARGIN，default SPOT
func (r ApiPortfolioMarginProBankruptcyLoanRepayRequest) From(from string) ApiPortfolioMarginProBankruptcyLoanRepayRequest {
	r.from = &from
	return r
}

func (r ApiPortfolioMarginProBankruptcyLoanRepayRequest) RecvWindow(recvWindow int64) ApiPortfolioMarginProBankruptcyLoanRepayRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPortfolioMarginProBankruptcyLoanRepayRequest) Execute() (*common.RestApiResponse[models.PortfolioMarginProBankruptcyLoanRepayResponse], error) {
	return r.ApiService.PortfolioMarginProBankruptcyLoanRepayExecute(r)
}

/*
PortfolioMarginProBankruptcyLoanRepay Portfolio Margin Pro Bankruptcy Loan Repay
Post /sapi/v1/portfolio/repay

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Classic-Portfolio-Margin-Bankruptcy-Loan-Repay

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param from -  SPOT or MARGIN，default SPOT
@param recvWindow -
@return ApiPortfolioMarginProBankruptcyLoanRepayRequest
*/
func (a *AccountAPIService) PortfolioMarginProBankruptcyLoanRepay(ctx context.Context) ApiPortfolioMarginProBankruptcyLoanRepayRequest {
	return ApiPortfolioMarginProBankruptcyLoanRepayRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PortfolioMarginProBankruptcyLoanRepayResponse
func (a *AccountAPIService) PortfolioMarginProBankruptcyLoanRepayExecute(r ApiPortfolioMarginProBankruptcyLoanRepayRequest) (*common.RestApiResponse[models.PortfolioMarginProBankruptcyLoanRepayResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/repay"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.from != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PortfolioMarginProBankruptcyLoanRepayResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest) RecvWindow(recvWindow int64) ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest) Execute() (*common.RestApiResponse[models.QueryPortfolioMarginProBankruptcyLoanAmountResponse], error) {
	return r.ApiService.QueryPortfolioMarginProBankruptcyLoanAmountExecute(r)
}

/*
QueryPortfolioMarginProBankruptcyLoanAmount Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)
Get /sapi/v1/portfolio/pmLoan

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Query-Classic-Portfolio-Margin-Bankruptcy-Loan-Amount

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest
*/
func (a *AccountAPIService) QueryPortfolioMarginProBankruptcyLoanAmount(ctx context.Context) ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest {
	return ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPortfolioMarginProBankruptcyLoanAmountResponse
func (a *AccountAPIService) QueryPortfolioMarginProBankruptcyLoanAmountExecute(r ApiQueryPortfolioMarginProBankruptcyLoanAmountRequest) (*common.RestApiResponse[models.QueryPortfolioMarginProBankruptcyLoanAmountResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/pmLoan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPortfolioMarginProBankruptcyLoanAmountResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) StartTime(startTime int64) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) EndTime(endTime int64) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) Current(current int64) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) Size(size int64) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) RecvWindow(recvWindow int64) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) Execute() (*common.RestApiResponse[models.QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse], error) {
	return r.ApiService.QueryPortfolioMarginProBankruptcyLoanRepayHistoryExecute(r)
}

/*
QueryPortfolioMarginProBankruptcyLoanRepayHistory Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)
Get /sapi/v1/portfolio/pmloan-history

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Query-Portfolio-Margin-Pro-Bankruptcy-Loan-Repay-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param recvWindow -
@return ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest
*/
func (a *AccountAPIService) QueryPortfolioMarginProBankruptcyLoanRepayHistory(ctx context.Context) ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest {
	return ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse
func (a *AccountAPIService) QueryPortfolioMarginProBankruptcyLoanRepayHistoryExecute(r ApiQueryPortfolioMarginProBankruptcyLoanRepayHistoryRequest) (*common.RestApiResponse[models.QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/pmloan-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPortfolioMarginProBankruptcyLoanRepayHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) Asset(asset string) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) StartTime(startTime int64) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) EndTime(endTime int64) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default:10 Max:100
func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) Size(size int64) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) RecvWindow(recvWindow int64) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) Execute() (*common.RestApiResponse[models.QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse], error) {
	return r.ApiService.QueryPortfolioMarginProNegativeBalanceInterestHistoryExecute(r)
}

/*
QueryPortfolioMarginProNegativeBalanceInterestHistory Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)
Get /sapi/v1/portfolio/interest-history

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Query-Classic-Portfolio-Margin-Negative-Balance-Interest-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param startTime -
@param endTime -
@param size -  Default:10 Max:100
@param recvWindow -
@return ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest
*/
func (a *AccountAPIService) QueryPortfolioMarginProNegativeBalanceInterestHistory(ctx context.Context) ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest {
	return ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse
func (a *AccountAPIService) QueryPortfolioMarginProNegativeBalanceInterestHistoryExecute(r ApiQueryPortfolioMarginProNegativeBalanceInterestHistoryRequest) (*common.RestApiResponse[models.QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/interest-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryPortfolioMarginProNegativeBalanceInterestHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRepayFuturesNegativeBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	from       *string
	recvWindow *int64
}

// SPOT or MARGIN，default SPOT
func (r ApiRepayFuturesNegativeBalanceRequest) From(from string) ApiRepayFuturesNegativeBalanceRequest {
	r.from = &from
	return r
}

func (r ApiRepayFuturesNegativeBalanceRequest) RecvWindow(recvWindow int64) ApiRepayFuturesNegativeBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRepayFuturesNegativeBalanceRequest) Execute() (*common.RestApiResponse[models.RepayFuturesNegativeBalanceResponse], error) {
	return r.ApiService.RepayFuturesNegativeBalanceExecute(r)
}

/*
RepayFuturesNegativeBalance Repay futures Negative Balance(USER_DATA)
Post /sapi/v1/portfolio/repay-futures-negative-balance

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Repay-futures-Negative-Balance

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param from -  SPOT or MARGIN，default SPOT
@param recvWindow -
@return ApiRepayFuturesNegativeBalanceRequest
*/
func (a *AccountAPIService) RepayFuturesNegativeBalance(ctx context.Context) ApiRepayFuturesNegativeBalanceRequest {
	return ApiRepayFuturesNegativeBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RepayFuturesNegativeBalanceResponse
func (a *AccountAPIService) RepayFuturesNegativeBalanceExecute(r ApiRepayFuturesNegativeBalanceRequest) (*common.RestApiResponse[models.RepayFuturesNegativeBalanceResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/repay-futures-negative-balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.from != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RepayFuturesNegativeBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTransferLdusdtRwusdForPortfolioMarginRequest struct {
	ctx          context.Context
	ApiService   *AccountAPIService
	asset        *string
	transferType *string
	amount       *float32
	recvWindow   *int64
}

// &#x60;LDUSDT&#x60; and &#x60;RWUSD&#x60;
func (r ApiTransferLdusdtRwusdForPortfolioMarginRequest) Asset(asset string) ApiTransferLdusdtRwusdForPortfolioMarginRequest {
	r.asset = &asset
	return r
}

// &#x60;EARN_TO_FUTURE&#x60; /&#x60;FUTURE_TO_EARN&#x60;
func (r ApiTransferLdusdtRwusdForPortfolioMarginRequest) TransferType(transferType string) ApiTransferLdusdtRwusdForPortfolioMarginRequest {
	r.transferType = &transferType
	return r
}

func (r ApiTransferLdusdtRwusdForPortfolioMarginRequest) Amount(amount float32) ApiTransferLdusdtRwusdForPortfolioMarginRequest {
	r.amount = &amount
	return r
}

func (r ApiTransferLdusdtRwusdForPortfolioMarginRequest) RecvWindow(recvWindow int64) ApiTransferLdusdtRwusdForPortfolioMarginRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTransferLdusdtRwusdForPortfolioMarginRequest) Execute() (*common.RestApiResponse[models.TransferLdusdtRwusdForPortfolioMarginResponse], error) {
	return r.ApiService.TransferLdusdtRwusdForPortfolioMarginExecute(r)
}

/*
TransferLdusdtRwusdForPortfolioMargin Transfer LDUSDT/RWUSD for Portfolio Margin(TRADE)
Post /sapi/v1/portfolio/earn-asset-transfer

https://developers.binance.com/docs/derivatives/portfolio-margin-pro/account/Transfer-LDUSDT-Portfolio-Margin

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  `LDUSDT` and `RWUSD`
@param transferType -  `EARN_TO_FUTURE` /`FUTURE_TO_EARN`
@param amount -
@param recvWindow -
@return ApiTransferLdusdtRwusdForPortfolioMarginRequest
*/
func (a *AccountAPIService) TransferLdusdtRwusdForPortfolioMargin(ctx context.Context) ApiTransferLdusdtRwusdForPortfolioMarginRequest {
	return ApiTransferLdusdtRwusdForPortfolioMarginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TransferLdusdtRwusdForPortfolioMarginResponse
func (a *AccountAPIService) TransferLdusdtRwusdForPortfolioMarginExecute(r ApiTransferLdusdtRwusdForPortfolioMarginRequest) (*common.RestApiResponse[models.TransferLdusdtRwusdForPortfolioMarginResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/portfolio/earn-asset-transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.transferType == nil {
		return nil, common.ReportError("transferType is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transferType", r.transferType, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TransferLdusdtRwusdForPortfolioMarginResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
