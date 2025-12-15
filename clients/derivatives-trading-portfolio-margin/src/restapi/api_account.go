/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package binancederivativestradingportfoliomarginrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService Service

type ApiAccountBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	recvWindow *int64
}

func (r ApiAccountBalanceRequest) Asset(asset string) ApiAccountBalanceRequest {
	r.asset = &asset
	return r
}

func (r ApiAccountBalanceRequest) RecvWindow(recvWindow int64) ApiAccountBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountBalanceRequest) Execute() (*common.RestApiResponse[models.AccountBalanceResponse], error) {
	return r.ApiService.AccountBalanceExecute(r)
}

/*
AccountBalance Account Balance(USER_DATA)
Get /papi/v1/balance

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Account-Balance

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param recvWindow -
@return ApiAccountBalanceRequest
*/
func (a *AccountAPIService) AccountBalance(ctx context.Context) ApiAccountBalanceRequest {
	return ApiAccountBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountBalanceResponse
func (a *AccountAPIService) AccountBalanceExecute(r ApiAccountBalanceRequest) (*common.RestApiResponse[models.AccountBalanceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAccountInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiAccountInformationRequest) RecvWindow(recvWindow int64) ApiAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInformationRequest) Execute() (*common.RestApiResponse[models.AccountInformationResponse], error) {
	return r.ApiService.AccountInformationExecute(r)
}

/*
AccountInformation Account Information(USER_DATA)
Get /papi/v1/account

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Account-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAccountInformationRequest
*/
func (a *AccountAPIService) AccountInformation(ctx context.Context) ApiAccountInformationRequest {
	return ApiAccountInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountInformationResponse
func (a *AccountAPIService) AccountInformationExecute(r ApiAccountInformationRequest) (*common.RestApiResponse[models.AccountInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

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
BnbTransfer BNB transfer (TRADE)
Post /papi/v1/bnb-transfer

https://developers.binance.com/docs/derivatives/portfolio-margin/account/BNB-transfer

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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/bnb-transfer"

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
Post /papi/v1/repay-futures-switch

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-Auto-repay-futures-Status

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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/repay-futures-switch"

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

type ApiChangeCmInitialLeverageRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	leverage   *int64
	recvWindow *int64
}

func (r ApiChangeCmInitialLeverageRequest) Symbol(symbol string) ApiChangeCmInitialLeverageRequest {
	r.symbol = &symbol
	return r
}

// target initial leverage: int from 1 to 125
func (r ApiChangeCmInitialLeverageRequest) Leverage(leverage int64) ApiChangeCmInitialLeverageRequest {
	r.leverage = &leverage
	return r
}

func (r ApiChangeCmInitialLeverageRequest) RecvWindow(recvWindow int64) ApiChangeCmInitialLeverageRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeCmInitialLeverageRequest) Execute() (*common.RestApiResponse[models.ChangeCmInitialLeverageResponse], error) {
	return r.ApiService.ChangeCmInitialLeverageExecute(r)
}

/*
ChangeCmInitialLeverage Change CM Initial Leverage (TRADE)
Post /papi/v1/cm/leverage

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-CM-Initial-Leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param leverage -  target initial leverage: int from 1 to 125
@param recvWindow -
@return ApiChangeCmInitialLeverageRequest
*/
func (a *AccountAPIService) ChangeCmInitialLeverage(ctx context.Context) ApiChangeCmInitialLeverageRequest {
	return ApiChangeCmInitialLeverageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeCmInitialLeverageResponse
func (a *AccountAPIService) ChangeCmInitialLeverageExecute(r ApiChangeCmInitialLeverageRequest) (*common.RestApiResponse[models.ChangeCmInitialLeverageResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/leverage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.leverage == nil {
		return nil, common.ReportError("leverage is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "leverage", r.leverage, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeCmInitialLeverageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiChangeCmPositionModeRequest struct {
	ctx              context.Context
	ApiService       *AccountAPIService
	dualSidePosition *string
	recvWindow       *int64
}

// \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode
func (r ApiChangeCmPositionModeRequest) DualSidePosition(dualSidePosition string) ApiChangeCmPositionModeRequest {
	r.dualSidePosition = &dualSidePosition
	return r
}

func (r ApiChangeCmPositionModeRequest) RecvWindow(recvWindow int64) ApiChangeCmPositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeCmPositionModeRequest) Execute() (*common.RestApiResponse[models.ChangeCmPositionModeResponse], error) {
	return r.ApiService.ChangeCmPositionModeExecute(r)
}

/*
ChangeCmPositionMode Change CM Position Mode(TRADE)
Post /papi/v1/cm/positionSide/dual

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-CM-Position-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param dualSidePosition -  \"true\": Hedge Mode; \"false\": One-way Mode
@param recvWindow -
@return ApiChangeCmPositionModeRequest
*/
func (a *AccountAPIService) ChangeCmPositionMode(ctx context.Context) ApiChangeCmPositionModeRequest {
	return ApiChangeCmPositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeCmPositionModeResponse
func (a *AccountAPIService) ChangeCmPositionModeExecute(r ApiChangeCmPositionModeRequest) (*common.RestApiResponse[models.ChangeCmPositionModeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.dualSidePosition == nil {
		return nil, common.ReportError("dualSidePosition is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dualSidePosition", r.dualSidePosition, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeCmPositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiChangeUmInitialLeverageRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	leverage   *int64
	recvWindow *int64
}

func (r ApiChangeUmInitialLeverageRequest) Symbol(symbol string) ApiChangeUmInitialLeverageRequest {
	r.symbol = &symbol
	return r
}

// target initial leverage: int from 1 to 125
func (r ApiChangeUmInitialLeverageRequest) Leverage(leverage int64) ApiChangeUmInitialLeverageRequest {
	r.leverage = &leverage
	return r
}

func (r ApiChangeUmInitialLeverageRequest) RecvWindow(recvWindow int64) ApiChangeUmInitialLeverageRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeUmInitialLeverageRequest) Execute() (*common.RestApiResponse[models.ChangeUmInitialLeverageResponse], error) {
	return r.ApiService.ChangeUmInitialLeverageExecute(r)
}

/*
ChangeUmInitialLeverage Change UM Initial Leverage(TRADE)
Post /papi/v1/um/leverage

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-UM-Initial-Leverage

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param leverage -  target initial leverage: int from 1 to 125
@param recvWindow -
@return ApiChangeUmInitialLeverageRequest
*/
func (a *AccountAPIService) ChangeUmInitialLeverage(ctx context.Context) ApiChangeUmInitialLeverageRequest {
	return ApiChangeUmInitialLeverageRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeUmInitialLeverageResponse
func (a *AccountAPIService) ChangeUmInitialLeverageExecute(r ApiChangeUmInitialLeverageRequest) (*common.RestApiResponse[models.ChangeUmInitialLeverageResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/leverage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.leverage == nil {
		return nil, common.ReportError("leverage is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "leverage", r.leverage, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeUmInitialLeverageResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiChangeUmPositionModeRequest struct {
	ctx              context.Context
	ApiService       *AccountAPIService
	dualSidePosition *string
	recvWindow       *int64
}

// \&quot;true\&quot;: Hedge Mode; \&quot;false\&quot;: One-way Mode
func (r ApiChangeUmPositionModeRequest) DualSidePosition(dualSidePosition string) ApiChangeUmPositionModeRequest {
	r.dualSidePosition = &dualSidePosition
	return r
}

func (r ApiChangeUmPositionModeRequest) RecvWindow(recvWindow int64) ApiChangeUmPositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiChangeUmPositionModeRequest) Execute() (*common.RestApiResponse[models.ChangeUmPositionModeResponse], error) {
	return r.ApiService.ChangeUmPositionModeExecute(r)
}

/*
ChangeUmPositionMode Change UM Position Mode(TRADE)
Post /papi/v1/um/positionSide/dual

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-UM-Position-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param dualSidePosition -  \"true\": Hedge Mode; \"false\": One-way Mode
@param recvWindow -
@return ApiChangeUmPositionModeRequest
*/
func (a *AccountAPIService) ChangeUmPositionMode(ctx context.Context) ApiChangeUmPositionModeRequest {
	return ApiChangeUmPositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ChangeUmPositionModeResponse
func (a *AccountAPIService) ChangeUmPositionModeExecute(r ApiChangeUmPositionModeRequest) (*common.RestApiResponse[models.ChangeUmPositionModeResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.dualSidePosition == nil {
		return nil, common.ReportError("dualSidePosition is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dualSidePosition", r.dualSidePosition, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ChangeUmPositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCmNotionalAndLeverageBracketsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiCmNotionalAndLeverageBracketsRequest) Symbol(symbol string) ApiCmNotionalAndLeverageBracketsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiCmNotionalAndLeverageBracketsRequest) RecvWindow(recvWindow int64) ApiCmNotionalAndLeverageBracketsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCmNotionalAndLeverageBracketsRequest) Execute() (*common.RestApiResponse[models.CmNotionalAndLeverageBracketsResponse], error) {
	return r.ApiService.CmNotionalAndLeverageBracketsExecute(r)
}

/*
CmNotionalAndLeverageBrackets CM Notional and Leverage Brackets(USER_DATA)
Get /papi/v1/cm/leverageBracket

https://developers.binance.com/docs/derivatives/portfolio-margin/account/CM-Notional-and-Leverage-Brackets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiCmNotionalAndLeverageBracketsRequest
*/
func (a *AccountAPIService) CmNotionalAndLeverageBrackets(ctx context.Context) ApiCmNotionalAndLeverageBracketsRequest {
	return ApiCmNotionalAndLeverageBracketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CmNotionalAndLeverageBracketsResponse
func (a *AccountAPIService) CmNotionalAndLeverageBracketsExecute(r ApiCmNotionalAndLeverageBracketsRequest) (*common.RestApiResponse[models.CmNotionalAndLeverageBracketsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CmNotionalAndLeverageBracketsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
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
FundAutoCollection Fund Auto-collection(TRADE)
Post /papi/v1/auto-collection

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Fund-Auto-collection

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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/auto-collection"

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
FundCollectionByAsset Fund Collection by Asset(TRADE)
Post /papi/v1/asset-collection

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Fund-Collection-by-Asset

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/asset-collection"

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
Get /papi/v1/repay-futures-switch

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Auto-repay-futures-Status

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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/repay-futures-switch"

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

type ApiGetCmAccountDetailRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetCmAccountDetailRequest) RecvWindow(recvWindow int64) ApiGetCmAccountDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCmAccountDetailRequest) Execute() (*common.RestApiResponse[models.GetCmAccountDetailResponse], error) {
	return r.ApiService.GetCmAccountDetailExecute(r)
}

/*
GetCmAccountDetail Get CM Account Detail(USER_DATA)
Get /papi/v1/cm/account

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-CM-Account-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetCmAccountDetailRequest
*/
func (a *AccountAPIService) GetCmAccountDetail(ctx context.Context) ApiGetCmAccountDetailRequest {
	return ApiGetCmAccountDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCmAccountDetailResponse
func (a *AccountAPIService) GetCmAccountDetailExecute(r ApiGetCmAccountDetailRequest) (*common.RestApiResponse[models.GetCmAccountDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCmAccountDetailResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCmCurrentPositionModeRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetCmCurrentPositionModeRequest) RecvWindow(recvWindow int64) ApiGetCmCurrentPositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCmCurrentPositionModeRequest) Execute() (*common.RestApiResponse[models.GetCmCurrentPositionModeResponse], error) {
	return r.ApiService.GetCmCurrentPositionModeExecute(r)
}

/*
GetCmCurrentPositionMode Get CM Current Position Mode(USER_DATA)
Get /papi/v1/cm/positionSide/dual

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-CM-Current-Position-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetCmCurrentPositionModeRequest
*/
func (a *AccountAPIService) GetCmCurrentPositionMode(ctx context.Context) ApiGetCmCurrentPositionModeRequest {
	return ApiGetCmCurrentPositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCmCurrentPositionModeResponse
func (a *AccountAPIService) GetCmCurrentPositionModeExecute(r ApiGetCmCurrentPositionModeRequest) (*common.RestApiResponse[models.GetCmCurrentPositionModeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCmCurrentPositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCmIncomeHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetCmIncomeHistoryRequest) Symbol(symbol string) ApiGetCmIncomeHistoryRequest {
	r.symbol = &symbol
	return r
}

// TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE
func (r ApiGetCmIncomeHistoryRequest) IncomeType(incomeType string) ApiGetCmIncomeHistoryRequest {
	r.incomeType = &incomeType
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiGetCmIncomeHistoryRequest) StartTime(startTime int64) ApiGetCmIncomeHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiGetCmIncomeHistoryRequest) EndTime(endTime int64) ApiGetCmIncomeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetCmIncomeHistoryRequest) Page(page int64) ApiGetCmIncomeHistoryRequest {
	r.page = &page
	return r
}

// Default 100; max 1000
func (r ApiGetCmIncomeHistoryRequest) Limit(limit int64) ApiGetCmIncomeHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCmIncomeHistoryRequest) RecvWindow(recvWindow int64) ApiGetCmIncomeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetCmIncomeHistoryRequest) Execute() (*common.RestApiResponse[models.GetCmIncomeHistoryResponse], error) {
	return r.ApiService.GetCmIncomeHistoryExecute(r)
}

/*
GetCmIncomeHistory Get CM Income History(USER_DATA)
Get /papi/v1/cm/income

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-CM-Income-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param incomeType -  TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param page -
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiGetCmIncomeHistoryRequest
*/
func (a *AccountAPIService) GetCmIncomeHistory(ctx context.Context) ApiGetCmIncomeHistoryRequest {
	return ApiGetCmIncomeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCmIncomeHistoryResponse
func (a *AccountAPIService) GetCmIncomeHistoryExecute(r ApiGetCmIncomeHistoryRequest) (*common.RestApiResponse[models.GetCmIncomeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/income"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.incomeType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "incomeType", r.incomeType, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetCmIncomeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForUmFuturesOrderHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiGetDownloadIdForUmFuturesOrderHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForUmFuturesOrderHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesOrderHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForUmFuturesOrderHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesOrderHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForUmFuturesOrderHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForUmFuturesOrderHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForUmFuturesOrderHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForUmFuturesOrderHistoryExecute(r)
}

/*
GetDownloadIdForUmFuturesOrderHistory Get Download Id For UM Futures Order History (USER_DATA)
Get /papi/v1/um/order/asyn

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Download-Id-For-UM-Futures-Order-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param recvWindow -
@return ApiGetDownloadIdForUmFuturesOrderHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForUmFuturesOrderHistory(ctx context.Context) ApiGetDownloadIdForUmFuturesOrderHistoryRequest {
	return ApiGetDownloadIdForUmFuturesOrderHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForUmFuturesOrderHistoryResponse
func (a *AccountAPIService) GetDownloadIdForUmFuturesOrderHistoryExecute(r ApiGetDownloadIdForUmFuturesOrderHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForUmFuturesOrderHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForUmFuturesOrderHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForUmFuturesTradeHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiGetDownloadIdForUmFuturesTradeHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForUmFuturesTradeHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesTradeHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForUmFuturesTradeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesTradeHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForUmFuturesTradeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForUmFuturesTradeHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForUmFuturesTradeHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForUmFuturesTradeHistoryExecute(r)
}

/*
GetDownloadIdForUmFuturesTradeHistory Get Download Id For UM Futures Trade History (USER_DATA)
Get /papi/v1/um/trade/asyn

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Download-Id-For-UM-Futures-Trade-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param recvWindow -
@return ApiGetDownloadIdForUmFuturesTradeHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForUmFuturesTradeHistory(ctx context.Context) ApiGetDownloadIdForUmFuturesTradeHistoryRequest {
	return ApiGetDownloadIdForUmFuturesTradeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForUmFuturesTradeHistoryResponse
func (a *AccountAPIService) GetDownloadIdForUmFuturesTradeHistoryExecute(r ApiGetDownloadIdForUmFuturesTradeHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForUmFuturesTradeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/trade/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForUmFuturesTradeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetDownloadIdForUmFuturesTransactionHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiGetDownloadIdForUmFuturesTransactionHistoryRequest) StartTime(startTime int64) ApiGetDownloadIdForUmFuturesTransactionHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesTransactionHistoryRequest) EndTime(endTime int64) ApiGetDownloadIdForUmFuturesTransactionHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetDownloadIdForUmFuturesTransactionHistoryRequest) RecvWindow(recvWindow int64) ApiGetDownloadIdForUmFuturesTransactionHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetDownloadIdForUmFuturesTransactionHistoryRequest) Execute() (*common.RestApiResponse[models.GetDownloadIdForUmFuturesTransactionHistoryResponse], error) {
	return r.ApiService.GetDownloadIdForUmFuturesTransactionHistoryExecute(r)
}

/*
GetDownloadIdForUmFuturesTransactionHistory Get Download Id For UM Futures Transaction History (USER_DATA)
Get /papi/v1/um/income/asyn

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Download-Id-For-UM-Futures-Transaction-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param recvWindow -
@return ApiGetDownloadIdForUmFuturesTransactionHistoryRequest
*/
func (a *AccountAPIService) GetDownloadIdForUmFuturesTransactionHistory(ctx context.Context) ApiGetDownloadIdForUmFuturesTransactionHistoryRequest {
	return ApiGetDownloadIdForUmFuturesTransactionHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetDownloadIdForUmFuturesTransactionHistoryResponse
func (a *AccountAPIService) GetDownloadIdForUmFuturesTransactionHistoryExecute(r ApiGetDownloadIdForUmFuturesTransactionHistoryRequest) (*common.RestApiResponse[models.GetDownloadIdForUmFuturesTransactionHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/income/asyn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetDownloadIdForUmFuturesTransactionHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetMarginBorrowLoanInterestHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	archived   *string
	recvWindow *int64
}

func (r ApiGetMarginBorrowLoanInterestHistoryRequest) Asset(asset string) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.asset = &asset
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiGetMarginBorrowLoanInterestHistoryRequest) StartTime(startTime int64) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiGetMarginBorrowLoanInterestHistoryRequest) EndTime(endTime int64) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiGetMarginBorrowLoanInterestHistoryRequest) Current(current int64) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiGetMarginBorrowLoanInterestHistoryRequest) Size(size int64) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.size = &size
	return r
}

// Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago
func (r ApiGetMarginBorrowLoanInterestHistoryRequest) Archived(archived string) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.archived = &archived
	return r
}

func (r ApiGetMarginBorrowLoanInterestHistoryRequest) RecvWindow(recvWindow int64) ApiGetMarginBorrowLoanInterestHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetMarginBorrowLoanInterestHistoryRequest) Execute() (*common.RestApiResponse[models.GetMarginBorrowLoanInterestHistoryResponse], error) {
	return r.ApiService.GetMarginBorrowLoanInterestHistoryExecute(r)
}

/*
GetMarginBorrowLoanInterestHistory Get Margin Borrow/Loan Interest History(USER_DATA)
Get /papi/v1/margin/marginInterestHistory

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Margin-BorrowLoan-Interest-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param archived -  Default: `false`. Set to `true` for archived data from 6 months ago
@param recvWindow -
@return ApiGetMarginBorrowLoanInterestHistoryRequest
*/
func (a *AccountAPIService) GetMarginBorrowLoanInterestHistory(ctx context.Context) ApiGetMarginBorrowLoanInterestHistoryRequest {
	return ApiGetMarginBorrowLoanInterestHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMarginBorrowLoanInterestHistoryResponse
func (a *AccountAPIService) GetMarginBorrowLoanInterestHistoryExecute(r ApiGetMarginBorrowLoanInterestHistoryRequest) (*common.RestApiResponse[models.GetMarginBorrowLoanInterestHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/marginInterestHistory"

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
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}
	if r.archived != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetMarginBorrowLoanInterestHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmAccountDetailRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetUmAccountDetailRequest) RecvWindow(recvWindow int64) ApiGetUmAccountDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmAccountDetailRequest) Execute() (*common.RestApiResponse[models.GetUmAccountDetailResponse], error) {
	return r.ApiService.GetUmAccountDetailExecute(r)
}

/*
GetUmAccountDetail Get UM Account Detail(USER_DATA)
Get /papi/v1/um/account

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Account-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetUmAccountDetailRequest
*/
func (a *AccountAPIService) GetUmAccountDetail(ctx context.Context) ApiGetUmAccountDetailRequest {
	return ApiGetUmAccountDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmAccountDetailResponse
func (a *AccountAPIService) GetUmAccountDetailExecute(r ApiGetUmAccountDetailRequest) (*common.RestApiResponse[models.GetUmAccountDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmAccountDetailResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmAccountDetailV2Request struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetUmAccountDetailV2Request) RecvWindow(recvWindow int64) ApiGetUmAccountDetailV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmAccountDetailV2Request) Execute() (*common.RestApiResponse[models.GetUmAccountDetailV2Response], error) {
	return r.ApiService.GetUmAccountDetailV2Execute(r)
}

/*
GetUmAccountDetailV2 Get UM Account Detail V2(USER_DATA)
Get /papi/v2/um/account

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Account-Detail-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetUmAccountDetailV2Request
*/
func (a *AccountAPIService) GetUmAccountDetailV2(ctx context.Context) ApiGetUmAccountDetailV2Request {
	return ApiGetUmAccountDetailV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmAccountDetailV2Response
func (a *AccountAPIService) GetUmAccountDetailV2Execute(r ApiGetUmAccountDetailV2Request) (*common.RestApiResponse[models.GetUmAccountDetailV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v2/um/account"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmAccountDetailV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmCurrentPositionModeRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiGetUmCurrentPositionModeRequest) RecvWindow(recvWindow int64) ApiGetUmCurrentPositionModeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmCurrentPositionModeRequest) Execute() (*common.RestApiResponse[models.GetUmCurrentPositionModeResponse], error) {
	return r.ApiService.GetUmCurrentPositionModeExecute(r)
}

/*
GetUmCurrentPositionMode Get UM Current Position Mode(USER_DATA)
Get /papi/v1/um/positionSide/dual

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Current-Position-Mode

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiGetUmCurrentPositionModeRequest
*/
func (a *AccountAPIService) GetUmCurrentPositionMode(ctx context.Context) ApiGetUmCurrentPositionModeRequest {
	return ApiGetUmCurrentPositionModeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmCurrentPositionModeResponse
func (a *AccountAPIService) GetUmCurrentPositionModeExecute(r ApiGetUmCurrentPositionModeRequest) (*common.RestApiResponse[models.GetUmCurrentPositionModeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/positionSide/dual"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmCurrentPositionModeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmFuturesOrderDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetUmFuturesOrderDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetUmFuturesOrderDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetUmFuturesOrderDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetUmFuturesOrderDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmFuturesOrderDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetUmFuturesOrderDownloadLinkByIdResponse], error) {
	return r.ApiService.GetUmFuturesOrderDownloadLinkByIdExecute(r)
}

/*
GetUmFuturesOrderDownloadLinkById Get UM Futures Order Download Link by Id(USER_DATA)
Get /papi/v1/um/order/asyn/id

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Order-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetUmFuturesOrderDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetUmFuturesOrderDownloadLinkById(ctx context.Context) ApiGetUmFuturesOrderDownloadLinkByIdRequest {
	return ApiGetUmFuturesOrderDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmFuturesOrderDownloadLinkByIdResponse
func (a *AccountAPIService) GetUmFuturesOrderDownloadLinkByIdExecute(r ApiGetUmFuturesOrderDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetUmFuturesOrderDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/order/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmFuturesOrderDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmFuturesTradeDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetUmFuturesTradeDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetUmFuturesTradeDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetUmFuturesTradeDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetUmFuturesTradeDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmFuturesTradeDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetUmFuturesTradeDownloadLinkByIdResponse], error) {
	return r.ApiService.GetUmFuturesTradeDownloadLinkByIdExecute(r)
}

/*
GetUmFuturesTradeDownloadLinkById Get UM Futures Trade Download Link by Id(USER_DATA)
Get /papi/v1/um/trade/asyn/id

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Trade-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetUmFuturesTradeDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetUmFuturesTradeDownloadLinkById(ctx context.Context) ApiGetUmFuturesTradeDownloadLinkByIdRequest {
	return ApiGetUmFuturesTradeDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmFuturesTradeDownloadLinkByIdResponse
func (a *AccountAPIService) GetUmFuturesTradeDownloadLinkByIdExecute(r ApiGetUmFuturesTradeDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetUmFuturesTradeDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/trade/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmFuturesTradeDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmFuturesTransactionDownloadLinkByIdRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	downloadId *string
	recvWindow *int64
}

// get by download id api
func (r ApiGetUmFuturesTransactionDownloadLinkByIdRequest) DownloadId(downloadId string) ApiGetUmFuturesTransactionDownloadLinkByIdRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetUmFuturesTransactionDownloadLinkByIdRequest) RecvWindow(recvWindow int64) ApiGetUmFuturesTransactionDownloadLinkByIdRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmFuturesTransactionDownloadLinkByIdRequest) Execute() (*common.RestApiResponse[models.GetUmFuturesTransactionDownloadLinkByIdResponse], error) {
	return r.ApiService.GetUmFuturesTransactionDownloadLinkByIdExecute(r)
}

/*
GetUmFuturesTransactionDownloadLinkById Get UM Futures Transaction Download Link by Id(USER_DATA)
Get /papi/v1/um/income/asyn/id

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Transaction-Download-Link-by-Id

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param downloadId -  get by download id api
@param recvWindow -
@return ApiGetUmFuturesTransactionDownloadLinkByIdRequest
*/
func (a *AccountAPIService) GetUmFuturesTransactionDownloadLinkById(ctx context.Context) ApiGetUmFuturesTransactionDownloadLinkByIdRequest {
	return ApiGetUmFuturesTransactionDownloadLinkByIdRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmFuturesTransactionDownloadLinkByIdResponse
func (a *AccountAPIService) GetUmFuturesTransactionDownloadLinkByIdExecute(r ApiGetUmFuturesTransactionDownloadLinkByIdRequest) (*common.RestApiResponse[models.GetUmFuturesTransactionDownloadLinkByIdResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/income/asyn/id"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.downloadId == nil {
		return nil, common.ReportError("downloadId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmFuturesTransactionDownloadLinkByIdResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUmIncomeHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	incomeType *string
	startTime  *int64
	endTime    *int64
	page       *int64
	limit      *int64
	recvWindow *int64
}

func (r ApiGetUmIncomeHistoryRequest) Symbol(symbol string) ApiGetUmIncomeHistoryRequest {
	r.symbol = &symbol
	return r
}

// TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE
func (r ApiGetUmIncomeHistoryRequest) IncomeType(incomeType string) ApiGetUmIncomeHistoryRequest {
	r.incomeType = &incomeType
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiGetUmIncomeHistoryRequest) StartTime(startTime int64) ApiGetUmIncomeHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiGetUmIncomeHistoryRequest) EndTime(endTime int64) ApiGetUmIncomeHistoryRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetUmIncomeHistoryRequest) Page(page int64) ApiGetUmIncomeHistoryRequest {
	r.page = &page
	return r
}

// Default 100; max 1000
func (r ApiGetUmIncomeHistoryRequest) Limit(limit int64) ApiGetUmIncomeHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetUmIncomeHistoryRequest) RecvWindow(recvWindow int64) ApiGetUmIncomeHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUmIncomeHistoryRequest) Execute() (*common.RestApiResponse[models.GetUmIncomeHistoryResponse], error) {
	return r.ApiService.GetUmIncomeHistoryExecute(r)
}

/*
GetUmIncomeHistory Get UM Income History(USER_DATA)
Get /papi/v1/um/income

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Income-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param incomeType -  TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param page -
@param limit -  Default 100; max 1000
@param recvWindow -
@return ApiGetUmIncomeHistoryRequest
*/
func (a *AccountAPIService) GetUmIncomeHistory(ctx context.Context) ApiGetUmIncomeHistoryRequest {
	return ApiGetUmIncomeHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUmIncomeHistoryResponse
func (a *AccountAPIService) GetUmIncomeHistoryExecute(r ApiGetUmIncomeHistoryRequest) (*common.RestApiResponse[models.GetUmIncomeHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/income"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.incomeType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "incomeType", r.incomeType, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUmIncomeHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUserCommissionRateForCmRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiGetUserCommissionRateForCmRequest) Symbol(symbol string) ApiGetUserCommissionRateForCmRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetUserCommissionRateForCmRequest) RecvWindow(recvWindow int64) ApiGetUserCommissionRateForCmRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUserCommissionRateForCmRequest) Execute() (*common.RestApiResponse[models.GetUserCommissionRateForCmResponse], error) {
	return r.ApiService.GetUserCommissionRateForCmExecute(r)
}

/*
GetUserCommissionRateForCm Get User Commission Rate for CM(USER_DATA)
Get /papi/v1/cm/commissionRate

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-User-Commission-Rate-for-CM

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiGetUserCommissionRateForCmRequest
*/
func (a *AccountAPIService) GetUserCommissionRateForCm(ctx context.Context) ApiGetUserCommissionRateForCmRequest {
	return ApiGetUserCommissionRateForCmRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUserCommissionRateForCmResponse
func (a *AccountAPIService) GetUserCommissionRateForCmExecute(r ApiGetUserCommissionRateForCmRequest) (*common.RestApiResponse[models.GetUserCommissionRateForCmResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/commissionRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUserCommissionRateForCmResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetUserCommissionRateForUmRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiGetUserCommissionRateForUmRequest) Symbol(symbol string) ApiGetUserCommissionRateForUmRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetUserCommissionRateForUmRequest) RecvWindow(recvWindow int64) ApiGetUserCommissionRateForUmRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetUserCommissionRateForUmRequest) Execute() (*common.RestApiResponse[models.GetUserCommissionRateForUmResponse], error) {
	return r.ApiService.GetUserCommissionRateForUmExecute(r)
}

/*
GetUserCommissionRateForUm Get User Commission Rate for UM(USER_DATA)
Get /papi/v1/um/commissionRate

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-User-Commission-Rate-for-UM

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiGetUserCommissionRateForUmRequest
*/
func (a *AccountAPIService) GetUserCommissionRateForUm(ctx context.Context) ApiGetUserCommissionRateForUmRequest {
	return ApiGetUserCommissionRateForUmRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetUserCommissionRateForUmResponse
func (a *AccountAPIService) GetUserCommissionRateForUmExecute(r ApiGetUserCommissionRateForUmRequest) (*common.RestApiResponse[models.GetUserCommissionRateForUmResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/commissionRate"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetUserCommissionRateForUmResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMarginMaxBorrowRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	recvWindow *int64
}

func (r ApiMarginMaxBorrowRequest) Asset(asset string) ApiMarginMaxBorrowRequest {
	r.asset = &asset
	return r
}

func (r ApiMarginMaxBorrowRequest) RecvWindow(recvWindow int64) ApiMarginMaxBorrowRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMarginMaxBorrowRequest) Execute() (*common.RestApiResponse[models.MarginMaxBorrowResponse], error) {
	return r.ApiService.MarginMaxBorrowExecute(r)
}

/*
MarginMaxBorrow Margin Max Borrow(USER_DATA)
Get /papi/v1/margin/maxBorrowable

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Margin-Max-Borrow

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param recvWindow -
@return ApiMarginMaxBorrowRequest
*/
func (a *AccountAPIService) MarginMaxBorrow(ctx context.Context) ApiMarginMaxBorrowRequest {
	return ApiMarginMaxBorrowRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MarginMaxBorrowResponse
func (a *AccountAPIService) MarginMaxBorrowExecute(r ApiMarginMaxBorrowRequest) (*common.RestApiResponse[models.MarginMaxBorrowResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/maxBorrowable"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MarginMaxBorrowResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest) Symbol(symbol string) ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest) RecvWindow(recvWindow int64) ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest) Execute() (*common.RestApiResponse[models.PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse], error) {
	return r.ApiService.PortfolioMarginUmTradingQuantitativeRulesIndicatorsExecute(r)
}

/*
PortfolioMarginUmTradingQuantitativeRulesIndicators Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)
Get /papi/v1/um/apiTradingStatus

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Portfolio-Margin-UM-Trading-Quantitative-Rules-Indicators

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest
*/
func (a *AccountAPIService) PortfolioMarginUmTradingQuantitativeRulesIndicators(ctx context.Context) ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest {
	return ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse
func (a *AccountAPIService) PortfolioMarginUmTradingQuantitativeRulesIndicatorsExecute(r ApiPortfolioMarginUmTradingQuantitativeRulesIndicatorsRequest) (*common.RestApiResponse[models.PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/apiTradingStatus"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryCmPositionInformationRequest struct {
	ctx         context.Context
	ApiService  *AccountAPIService
	marginAsset *string
	pair        *string
	recvWindow  *int64
}

func (r ApiQueryCmPositionInformationRequest) MarginAsset(marginAsset string) ApiQueryCmPositionInformationRequest {
	r.marginAsset = &marginAsset
	return r
}

func (r ApiQueryCmPositionInformationRequest) Pair(pair string) ApiQueryCmPositionInformationRequest {
	r.pair = &pair
	return r
}

func (r ApiQueryCmPositionInformationRequest) RecvWindow(recvWindow int64) ApiQueryCmPositionInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryCmPositionInformationRequest) Execute() (*common.RestApiResponse[models.QueryCmPositionInformationResponse], error) {
	return r.ApiService.QueryCmPositionInformationExecute(r)
}

/*
QueryCmPositionInformation Query CM Position Information(USER_DATA)
Get /papi/v1/cm/positionRisk

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-CM-Position-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param marginAsset -
@param pair -
@param recvWindow -
@return ApiQueryCmPositionInformationRequest
*/
func (a *AccountAPIService) QueryCmPositionInformation(ctx context.Context) ApiQueryCmPositionInformationRequest {
	return ApiQueryCmPositionInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryCmPositionInformationResponse
func (a *AccountAPIService) QueryCmPositionInformationExecute(r ApiQueryCmPositionInformationRequest) (*common.RestApiResponse[models.QueryCmPositionInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/cm/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.marginAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "marginAsset", r.marginAsset, "form", "")
	}
	if r.pair != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pair", r.pair, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryCmPositionInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginLoanRecordRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	txId       *int64
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	archived   *string
	recvWindow *int64
}

func (r ApiQueryMarginLoanRecordRequest) Asset(asset string) ApiQueryMarginLoanRecordRequest {
	r.asset = &asset
	return r
}

// the &#x60;tranId&#x60; in &#x60;POST/papi/v1/marginLoan&#x60;
func (r ApiQueryMarginLoanRecordRequest) TxId(txId int64) ApiQueryMarginLoanRecordRequest {
	r.txId = &txId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryMarginLoanRecordRequest) StartTime(startTime int64) ApiQueryMarginLoanRecordRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryMarginLoanRecordRequest) EndTime(endTime int64) ApiQueryMarginLoanRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiQueryMarginLoanRecordRequest) Current(current int64) ApiQueryMarginLoanRecordRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiQueryMarginLoanRecordRequest) Size(size int64) ApiQueryMarginLoanRecordRequest {
	r.size = &size
	return r
}

// Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago
func (r ApiQueryMarginLoanRecordRequest) Archived(archived string) ApiQueryMarginLoanRecordRequest {
	r.archived = &archived
	return r
}

func (r ApiQueryMarginLoanRecordRequest) RecvWindow(recvWindow int64) ApiQueryMarginLoanRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginLoanRecordRequest) Execute() (*common.RestApiResponse[models.QueryMarginLoanRecordResponse], error) {
	return r.ApiService.QueryMarginLoanRecordExecute(r)
}

/*
QueryMarginLoanRecord Query Margin Loan Record(USER_DATA)
Get /papi/v1/margin/marginLoan

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-Loan-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param txId -  the `tranId` in `POST/papi/v1/marginLoan`
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param archived -  Default: `false`. Set to `true` for archived data from 6 months ago
@param recvWindow -
@return ApiQueryMarginLoanRecordRequest
*/
func (a *AccountAPIService) QueryMarginLoanRecord(ctx context.Context) ApiQueryMarginLoanRecordRequest {
	return ApiQueryMarginLoanRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginLoanRecordResponse
func (a *AccountAPIService) QueryMarginLoanRecordExecute(r ApiQueryMarginLoanRecordRequest) (*common.RestApiResponse[models.QueryMarginLoanRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/marginLoan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
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
	if r.archived != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginLoanRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginMaxWithdrawRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	recvWindow *int64
}

func (r ApiQueryMarginMaxWithdrawRequest) Asset(asset string) ApiQueryMarginMaxWithdrawRequest {
	r.asset = &asset
	return r
}

func (r ApiQueryMarginMaxWithdrawRequest) RecvWindow(recvWindow int64) ApiQueryMarginMaxWithdrawRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginMaxWithdrawRequest) Execute() (*common.RestApiResponse[models.QueryMarginMaxWithdrawResponse], error) {
	return r.ApiService.QueryMarginMaxWithdrawExecute(r)
}

/*
QueryMarginMaxWithdraw Query Margin Max Withdraw(USER_DATA)
Get /papi/v1/margin/maxWithdraw

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-Max-Withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param recvWindow -
@return ApiQueryMarginMaxWithdrawRequest
*/
func (a *AccountAPIService) QueryMarginMaxWithdraw(ctx context.Context) ApiQueryMarginMaxWithdrawRequest {
	return ApiQueryMarginMaxWithdrawRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginMaxWithdrawResponse
func (a *AccountAPIService) QueryMarginMaxWithdrawExecute(r ApiQueryMarginMaxWithdrawRequest) (*common.RestApiResponse[models.QueryMarginMaxWithdrawResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/maxWithdraw"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginMaxWithdrawResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryMarginRepayRecordRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	txId       *int64
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	archived   *string
	recvWindow *int64
}

func (r ApiQueryMarginRepayRecordRequest) Asset(asset string) ApiQueryMarginRepayRecordRequest {
	r.asset = &asset
	return r
}

// the &#x60;tranId&#x60; in &#x60;POST/papi/v1/marginLoan&#x60;
func (r ApiQueryMarginRepayRecordRequest) TxId(txId int64) ApiQueryMarginRepayRecordRequest {
	r.txId = &txId
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryMarginRepayRecordRequest) StartTime(startTime int64) ApiQueryMarginRepayRecordRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryMarginRepayRecordRequest) EndTime(endTime int64) ApiQueryMarginRepayRecordRequest {
	r.endTime = &endTime
	return r
}

// Currently querying page. Start from 1. Default:1
func (r ApiQueryMarginRepayRecordRequest) Current(current int64) ApiQueryMarginRepayRecordRequest {
	r.current = &current
	return r
}

// Default:10 Max:100
func (r ApiQueryMarginRepayRecordRequest) Size(size int64) ApiQueryMarginRepayRecordRequest {
	r.size = &size
	return r
}

// Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago
func (r ApiQueryMarginRepayRecordRequest) Archived(archived string) ApiQueryMarginRepayRecordRequest {
	r.archived = &archived
	return r
}

func (r ApiQueryMarginRepayRecordRequest) RecvWindow(recvWindow int64) ApiQueryMarginRepayRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryMarginRepayRecordRequest) Execute() (*common.RestApiResponse[models.QueryMarginRepayRecordResponse], error) {
	return r.ApiService.QueryMarginRepayRecordExecute(r)
}

/*
QueryMarginRepayRecord Query Margin repay Record(USER_DATA)
Get /papi/v1/margin/repayLoan

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-repay-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param txId -  the `tranId` in `POST/papi/v1/marginLoan`
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param current -  Currently querying page. Start from 1. Default:1
@param size -  Default:10 Max:100
@param archived -  Default: `false`. Set to `true` for archived data from 6 months ago
@param recvWindow -
@return ApiQueryMarginRepayRecordRequest
*/
func (a *AccountAPIService) QueryMarginRepayRecord(ctx context.Context) ApiQueryMarginRepayRecordRequest {
	return ApiQueryMarginRepayRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryMarginRepayRecordResponse
func (a *AccountAPIService) QueryMarginRepayRecordExecute(r ApiQueryMarginRepayRecordRequest) (*common.RestApiResponse[models.QueryMarginRepayRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/margin/repayLoan"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
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
	if r.archived != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryMarginRepayRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) Asset(asset string) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	r.asset = &asset
	return r
}

// Timestamp in ms to get funding from INCLUSIVE.
func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) StartTime(startTime int64) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get funding until INCLUSIVE.
func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) EndTime(endTime int64) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	r.endTime = &endTime
	return r
}

// Default:10 Max:100
func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) Size(size int64) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) RecvWindow(recvWindow int64) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) Execute() (*common.RestApiResponse[models.QueryPortfolioMarginNegativeBalanceInterestHistoryResponse], error) {
	return r.ApiService.QueryPortfolioMarginNegativeBalanceInterestHistoryExecute(r)
}

/*
QueryPortfolioMarginNegativeBalanceInterestHistory Query Portfolio Margin Negative Balance Interest History(USER_DATA)
Get /papi/v1/portfolio/interest-history

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Portfolio-Margin-Negative-Balance-Interest-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param startTime -  Timestamp in ms to get funding from INCLUSIVE.
@param endTime -  Timestamp in ms to get funding until INCLUSIVE.
@param size -  Default:10 Max:100
@param recvWindow -
@return ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest
*/
func (a *AccountAPIService) QueryPortfolioMarginNegativeBalanceInterestHistory(ctx context.Context) ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest {
	return ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryPortfolioMarginNegativeBalanceInterestHistoryResponse
func (a *AccountAPIService) QueryPortfolioMarginNegativeBalanceInterestHistoryExecute(r ApiQueryPortfolioMarginNegativeBalanceInterestHistoryRequest) (*common.RestApiResponse[models.QueryPortfolioMarginNegativeBalanceInterestHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/portfolio/interest-history"

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

	resp, err := SendRequest[models.QueryPortfolioMarginNegativeBalanceInterestHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUmPositionInformationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiQueryUmPositionInformationRequest) Symbol(symbol string) ApiQueryUmPositionInformationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiQueryUmPositionInformationRequest) RecvWindow(recvWindow int64) ApiQueryUmPositionInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUmPositionInformationRequest) Execute() (*common.RestApiResponse[models.QueryUmPositionInformationResponse], error) {
	return r.ApiService.QueryUmPositionInformationExecute(r)
}

/*
QueryUmPositionInformation Query UM Position Information(USER_DATA)
Get /papi/v1/um/positionRisk

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-UM-Position-Information

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiQueryUmPositionInformationRequest
*/
func (a *AccountAPIService) QueryUmPositionInformation(ctx context.Context) ApiQueryUmPositionInformationRequest {
	return ApiQueryUmPositionInformationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUmPositionInformationResponse
func (a *AccountAPIService) QueryUmPositionInformationExecute(r ApiQueryUmPositionInformationRequest) (*common.RestApiResponse[models.QueryUmPositionInformationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/positionRisk"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUmPositionInformationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserNegativeBalanceAutoExchangeRecordRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiQueryUserNegativeBalanceAutoExchangeRecordRequest) StartTime(startTime int64) ApiQueryUserNegativeBalanceAutoExchangeRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryUserNegativeBalanceAutoExchangeRecordRequest) EndTime(endTime int64) ApiQueryUserNegativeBalanceAutoExchangeRecordRequest {
	r.endTime = &endTime
	return r
}

func (r ApiQueryUserNegativeBalanceAutoExchangeRecordRequest) RecvWindow(recvWindow int64) ApiQueryUserNegativeBalanceAutoExchangeRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserNegativeBalanceAutoExchangeRecordRequest) Execute() (*common.RestApiResponse[models.QueryUserNegativeBalanceAutoExchangeRecordResponse], error) {
	return r.ApiService.QueryUserNegativeBalanceAutoExchangeRecordExecute(r)
}

/*
QueryUserNegativeBalanceAutoExchangeRecord Query User Negative Balance Auto Exchange Record (USER_DATA)
Get /papi/v1/portfolio/negative-balance-exchange-record

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-User-Negative-Balance-Auto-Exchange-Record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param recvWindow -
@return ApiQueryUserNegativeBalanceAutoExchangeRecordRequest
*/
func (a *AccountAPIService) QueryUserNegativeBalanceAutoExchangeRecord(ctx context.Context) ApiQueryUserNegativeBalanceAutoExchangeRecordRequest {
	return ApiQueryUserNegativeBalanceAutoExchangeRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserNegativeBalanceAutoExchangeRecordResponse
func (a *AccountAPIService) QueryUserNegativeBalanceAutoExchangeRecordExecute(r ApiQueryUserNegativeBalanceAutoExchangeRecordRequest) (*common.RestApiResponse[models.QueryUserNegativeBalanceAutoExchangeRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/portfolio/negative-balance-exchange-record"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUserNegativeBalanceAutoExchangeRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserRateLimitRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiQueryUserRateLimitRequest) RecvWindow(recvWindow int64) ApiQueryUserRateLimitRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserRateLimitRequest) Execute() (*common.RestApiResponse[models.QueryUserRateLimitResponse], error) {
	return r.ApiService.QueryUserRateLimitExecute(r)
}

/*
QueryUserRateLimit Query User Rate Limit (USER_DATA)
Get /papi/v1/rateLimit/order

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-User-Rate-Limit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiQueryUserRateLimitRequest
*/
func (a *AccountAPIService) QueryUserRateLimit(ctx context.Context) ApiQueryUserRateLimitRequest {
	return ApiQueryUserRateLimitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserRateLimitResponse
func (a *AccountAPIService) QueryUserRateLimitExecute(r ApiQueryUserRateLimitRequest) (*common.RestApiResponse[models.QueryUserRateLimitResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/rateLimit/order"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUserRateLimitResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRepayFuturesNegativeBalanceRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
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
Post /papi/v1/repay-futures-negative-balance

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Repay-futures-Negative-Balance

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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
	localVarPath := a.client.cfg.BasePath + "/papi/v1/repay-futures-negative-balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RepayFuturesNegativeBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUmFuturesAccountConfigurationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	recvWindow *int64
}

func (r ApiUmFuturesAccountConfigurationRequest) RecvWindow(recvWindow int64) ApiUmFuturesAccountConfigurationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUmFuturesAccountConfigurationRequest) Execute() (*common.RestApiResponse[models.UmFuturesAccountConfigurationResponse], error) {
	return r.ApiService.UmFuturesAccountConfigurationExecute(r)
}

/*
UmFuturesAccountConfiguration UM Futures Account Configuration(USER_DATA)
Get /papi/v1/um/accountConfig

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Account-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiUmFuturesAccountConfigurationRequest
*/
func (a *AccountAPIService) UmFuturesAccountConfiguration(ctx context.Context) ApiUmFuturesAccountConfigurationRequest {
	return ApiUmFuturesAccountConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UmFuturesAccountConfigurationResponse
func (a *AccountAPIService) UmFuturesAccountConfigurationExecute(r ApiUmFuturesAccountConfigurationRequest) (*common.RestApiResponse[models.UmFuturesAccountConfigurationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/accountConfig"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UmFuturesAccountConfigurationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUmFuturesSymbolConfigurationRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiUmFuturesSymbolConfigurationRequest) Symbol(symbol string) ApiUmFuturesSymbolConfigurationRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUmFuturesSymbolConfigurationRequest) RecvWindow(recvWindow int64) ApiUmFuturesSymbolConfigurationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUmFuturesSymbolConfigurationRequest) Execute() (*common.RestApiResponse[models.UmFuturesSymbolConfigurationResponse], error) {
	return r.ApiService.UmFuturesSymbolConfigurationExecute(r)
}

/*
UmFuturesSymbolConfiguration UM Futures Symbol Configuration(USER_DATA)
Get /papi/v1/um/symbolConfig

https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Symbol-Config

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiUmFuturesSymbolConfigurationRequest
*/
func (a *AccountAPIService) UmFuturesSymbolConfiguration(ctx context.Context) ApiUmFuturesSymbolConfigurationRequest {
	return ApiUmFuturesSymbolConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UmFuturesSymbolConfigurationResponse
func (a *AccountAPIService) UmFuturesSymbolConfigurationExecute(r ApiUmFuturesSymbolConfigurationRequest) (*common.RestApiResponse[models.UmFuturesSymbolConfigurationResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/symbolConfig"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UmFuturesSymbolConfigurationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUmNotionalAndLeverageBracketsRequest struct {
	ctx        context.Context
	ApiService *AccountAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiUmNotionalAndLeverageBracketsRequest) Symbol(symbol string) ApiUmNotionalAndLeverageBracketsRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUmNotionalAndLeverageBracketsRequest) RecvWindow(recvWindow int64) ApiUmNotionalAndLeverageBracketsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUmNotionalAndLeverageBracketsRequest) Execute() (*common.RestApiResponse[models.UmNotionalAndLeverageBracketsResponse], error) {
	return r.ApiService.UmNotionalAndLeverageBracketsExecute(r)
}

/*
UmNotionalAndLeverageBrackets UM Notional and Leverage Brackets (USER_DATA)
Get /papi/v1/um/leverageBracket

https://developers.binance.com/docs/derivatives/portfolio-margin/account/UM-Notional-and-Leverage-Brackets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiUmNotionalAndLeverageBracketsRequest
*/
func (a *AccountAPIService) UmNotionalAndLeverageBrackets(ctx context.Context) ApiUmNotionalAndLeverageBracketsRequest {
	return ApiUmNotionalAndLeverageBracketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UmNotionalAndLeverageBracketsResponse
func (a *AccountAPIService) UmNotionalAndLeverageBracketsExecute(r ApiUmNotionalAndLeverageBracketsRequest) (*common.RestApiResponse[models.UmNotionalAndLeverageBracketsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/papi/v1/um/leverageBracket"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UmNotionalAndLeverageBracketsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
