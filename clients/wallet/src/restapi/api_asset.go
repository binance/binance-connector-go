/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package binancewalletrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AssetAPIService AssetAPI Service
type AssetAPIService Service

type ApiAssetDetailRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	recvWindow *int64
}

func (r ApiAssetDetailRequest) RecvWindow(recvWindow int64) ApiAssetDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAssetDetailRequest) Execute() (*common.RestApiResponse[models.AssetDetailResponse], error) {
	return r.ApiService.AssetDetailExecute(r)
}

/*
AssetDetail Asset Detail (USER_DATA)
Get /sapi/v1/asset/assetDetail

https://developers.binance.com/docs/wallet/asset/Asset-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiAssetDetailRequest
*/
func (a *AssetAPIService) AssetDetail(ctx context.Context) ApiAssetDetailRequest {
	return ApiAssetDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AssetDetailResponse
func (a *AssetAPIService) AssetDetailExecute(r ApiAssetDetailRequest) (*common.RestApiResponse[models.AssetDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/assetDetail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AssetDetailResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAssetDividendRecordRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	asset      *string
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

// If asset is blank, then query all positive assets user have.
func (r ApiAssetDividendRecordRequest) Asset(asset string) ApiAssetDividendRecordRequest {
	r.asset = &asset
	return r
}

func (r ApiAssetDividendRecordRequest) StartTime(startTime int64) ApiAssetDividendRecordRequest {
	r.startTime = &startTime
	return r
}

func (r ApiAssetDividendRecordRequest) EndTime(endTime int64) ApiAssetDividendRecordRequest {
	r.endTime = &endTime
	return r
}

// min 7, max 30, default 7
func (r ApiAssetDividendRecordRequest) Limit(limit int64) ApiAssetDividendRecordRequest {
	r.limit = &limit
	return r
}

func (r ApiAssetDividendRecordRequest) RecvWindow(recvWindow int64) ApiAssetDividendRecordRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAssetDividendRecordRequest) Execute() (*common.RestApiResponse[models.AssetDividendRecordResponse], error) {
	return r.ApiService.AssetDividendRecordExecute(r)
}

/*
AssetDividendRecord Asset Dividend Record (USER_DATA)
Get /sapi/v1/asset/assetDividend

https://developers.binance.com/docs/wallet/asset/assets-divided-record

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  If asset is blank, then query all positive assets user have.
@param startTime -
@param endTime -
@param limit -  min 7, max 30, default 7
@param recvWindow -
@return ApiAssetDividendRecordRequest
*/
func (a *AssetAPIService) AssetDividendRecord(ctx context.Context) ApiAssetDividendRecordRequest {
	return ApiAssetDividendRecordRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AssetDividendRecordResponse
func (a *AssetAPIService) AssetDividendRecordExecute(r ApiAssetDividendRecordRequest) (*common.RestApiResponse[models.AssetDividendRecordResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/assetDividend"

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
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AssetDividendRecordResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDustConvertRequest struct {
	ctx                              context.Context
	ApiService                       *AssetAPIService
	asset                            *string
	clientId                         *string
	targetAsset                      *string
	thirdPartyClientId               *string
	dustQuotaAssetToTargetAssetPrice *float32
}

func (r ApiDustConvertRequest) Asset(asset string) ApiDustConvertRequest {
	r.asset = &asset
	return r
}

// A unique id for the request
func (r ApiDustConvertRequest) ClientId(clientId string) ApiDustConvertRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDustConvertRequest) TargetAsset(targetAsset string) ApiDustConvertRequest {
	r.targetAsset = &targetAsset
	return r
}

func (r ApiDustConvertRequest) ThirdPartyClientId(thirdPartyClientId string) ApiDustConvertRequest {
	r.thirdPartyClientId = &thirdPartyClientId
	return r
}

func (r ApiDustConvertRequest) DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice float32) ApiDustConvertRequest {
	r.dustQuotaAssetToTargetAssetPrice = &dustQuotaAssetToTargetAssetPrice
	return r
}

func (r ApiDustConvertRequest) Execute() (*common.RestApiResponse[models.DustConvertResponse], error) {
	return r.ApiService.DustConvertExecute(r)
}

/*
DustConvert Dust Convert (USER_DATA)
Post /sapi/v1/asset/dust-convert/convert

https://developers.binance.com/docs/wallet/asset/Dust-Convert

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param clientId -  A unique id for the request
@param targetAsset -
@param thirdPartyClientId -
@param dustQuotaAssetToTargetAssetPrice -
@return ApiDustConvertRequest
*/
func (a *AssetAPIService) DustConvert(ctx context.Context) ApiDustConvertRequest {
	return ApiDustConvertRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DustConvertResponse
func (a *AssetAPIService) DustConvertExecute(r ApiDustConvertRequest) (*common.RestApiResponse[models.DustConvertResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/dust-convert/convert"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.clientId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "form", "")
	}
	if r.targetAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "targetAsset", r.targetAsset, "form", "")
	}
	if r.thirdPartyClientId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "thirdPartyClientId", r.thirdPartyClientId, "form", "")
	}
	if r.dustQuotaAssetToTargetAssetPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dustQuotaAssetToTargetAssetPrice", r.dustQuotaAssetToTargetAssetPrice, "form", "")
	}

	resp, err := SendRequest[models.DustConvertResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDustConvertibleAssetsRequest struct {
	ctx                              context.Context
	ApiService                       *AssetAPIService
	targetAsset                      *string
	dustQuotaAssetToTargetAssetPrice *float32
}

func (r ApiDustConvertibleAssetsRequest) TargetAsset(targetAsset string) ApiDustConvertibleAssetsRequest {
	r.targetAsset = &targetAsset
	return r
}

func (r ApiDustConvertibleAssetsRequest) DustQuotaAssetToTargetAssetPrice(dustQuotaAssetToTargetAssetPrice float32) ApiDustConvertibleAssetsRequest {
	r.dustQuotaAssetToTargetAssetPrice = &dustQuotaAssetToTargetAssetPrice
	return r
}

func (r ApiDustConvertibleAssetsRequest) Execute() (*common.RestApiResponse[models.DustConvertibleAssetsResponse], error) {
	return r.ApiService.DustConvertibleAssetsExecute(r)
}

/*
DustConvertibleAssets Dust Convertible Assets (USER_DATA)
Post /sapi/v1/asset/dust-convert/query-convertible-assets

https://developers.binance.com/docs/wallet/asset/Dust-Convertible-Assets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param targetAsset -
@param dustQuotaAssetToTargetAssetPrice -
@return ApiDustConvertibleAssetsRequest
*/
func (a *AssetAPIService) DustConvertibleAssets(ctx context.Context) ApiDustConvertibleAssetsRequest {
	return ApiDustConvertibleAssetsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DustConvertibleAssetsResponse
func (a *AssetAPIService) DustConvertibleAssetsExecute(r ApiDustConvertibleAssetsRequest) (*common.RestApiResponse[models.DustConvertibleAssetsResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/dust-convert/query-convertible-assets"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.targetAsset == nil {
		return nil, common.ReportError("targetAsset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "targetAsset", r.targetAsset, "form", "")
	if r.dustQuotaAssetToTargetAssetPrice != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "dustQuotaAssetToTargetAssetPrice", r.dustQuotaAssetToTargetAssetPrice, "form", "")
	}

	resp, err := SendRequest[models.DustConvertibleAssetsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDustTransferRequest struct {
	ctx         context.Context
	ApiService  *AssetAPIService
	asset       *string
	accountType *string
	recvWindow  *int64
}

func (r ApiDustTransferRequest) Asset(asset string) ApiDustTransferRequest {
	r.asset = &asset
	return r
}

// &#x60;SPOT&#x60; or &#x60;MARGIN&#x60;,default &#x60;SPOT&#x60;
func (r ApiDustTransferRequest) AccountType(accountType string) ApiDustTransferRequest {
	r.accountType = &accountType
	return r
}

func (r ApiDustTransferRequest) RecvWindow(recvWindow int64) ApiDustTransferRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDustTransferRequest) Execute() (*common.RestApiResponse[models.DustTransferResponse], error) {
	return r.ApiService.DustTransferExecute(r)
}

/*
DustTransfer Dust Transfer (USER_DATA)
Post /sapi/v1/asset/dust

https://developers.binance.com/docs/wallet/asset/Dust-Transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -
@param accountType -  `SPOT` or `MARGIN`,default `SPOT`
@param recvWindow -
@return ApiDustTransferRequest
*/
func (a *AssetAPIService) DustTransfer(ctx context.Context) ApiDustTransferRequest {
	return ApiDustTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DustTransferResponse
func (a *AssetAPIService) DustTransferExecute(r ApiDustTransferRequest) (*common.RestApiResponse[models.DustTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/dust"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.accountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DustTransferResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDustlogRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	startTime  *int64
	endTime    *int64
	recvWindow *int64
}

func (r ApiDustlogRequest) StartTime(startTime int64) ApiDustlogRequest {
	r.startTime = &startTime
	return r
}

func (r ApiDustlogRequest) EndTime(endTime int64) ApiDustlogRequest {
	r.endTime = &endTime
	return r
}

func (r ApiDustlogRequest) RecvWindow(recvWindow int64) ApiDustlogRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiDustlogRequest) Execute() (*common.RestApiResponse[models.DustlogResponse], error) {
	return r.ApiService.DustlogExecute(r)
}

/*
Dustlog DustLog(USER_DATA)
Get /sapi/v1/asset/dribblet

https://developers.binance.com/docs/wallet/asset/dust-log

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param recvWindow -
@return ApiDustlogRequest
*/
func (a *AssetAPIService) Dustlog(ctx context.Context) ApiDustlogRequest {
	return ApiDustlogRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DustlogResponse
func (a *AssetAPIService) DustlogExecute(r ApiDustlogRequest) (*common.RestApiResponse[models.DustlogResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/dribblet"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.DustlogResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFundingWalletRequest struct {
	ctx              context.Context
	ApiService       *AssetAPIService
	asset            *string
	needBtcValuation *string
	recvWindow       *int64
}

// If asset is blank, then query all positive assets user have.
func (r ApiFundingWalletRequest) Asset(asset string) ApiFundingWalletRequest {
	r.asset = &asset
	return r
}

// true or false
func (r ApiFundingWalletRequest) NeedBtcValuation(needBtcValuation string) ApiFundingWalletRequest {
	r.needBtcValuation = &needBtcValuation
	return r
}

func (r ApiFundingWalletRequest) RecvWindow(recvWindow int64) ApiFundingWalletRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFundingWalletRequest) Execute() (*common.RestApiResponse[models.FundingWalletResponse], error) {
	return r.ApiService.FundingWalletExecute(r)
}

/*
FundingWallet Funding Wallet (USER_DATA)
Post /sapi/v1/asset/get-funding-asset

https://developers.binance.com/docs/wallet/asset/Funding-Wallet

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  If asset is blank, then query all positive assets user have.
@param needBtcValuation -  true or false
@param recvWindow -
@return ApiFundingWalletRequest
*/
func (a *AssetAPIService) FundingWallet(ctx context.Context) ApiFundingWalletRequest {
	return ApiFundingWalletRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FundingWalletResponse
func (a *AssetAPIService) FundingWalletExecute(r ApiFundingWalletRequest) (*common.RestApiResponse[models.FundingWalletResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/get-funding-asset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.needBtcValuation != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "needBtcValuation", r.needBtcValuation, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FundingWalletResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetAssetsThatCanBeConvertedIntoBnbRequest struct {
	ctx         context.Context
	ApiService  *AssetAPIService
	accountType *string
	recvWindow  *int64
}

// &#x60;SPOT&#x60; or &#x60;MARGIN&#x60;,default &#x60;SPOT&#x60;
func (r ApiGetAssetsThatCanBeConvertedIntoBnbRequest) AccountType(accountType string) ApiGetAssetsThatCanBeConvertedIntoBnbRequest {
	r.accountType = &accountType
	return r
}

func (r ApiGetAssetsThatCanBeConvertedIntoBnbRequest) RecvWindow(recvWindow int64) ApiGetAssetsThatCanBeConvertedIntoBnbRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetAssetsThatCanBeConvertedIntoBnbRequest) Execute() (*common.RestApiResponse[models.GetAssetsThatCanBeConvertedIntoBnbResponse], error) {
	return r.ApiService.GetAssetsThatCanBeConvertedIntoBnbExecute(r)
}

/*
GetAssetsThatCanBeConvertedIntoBnb Get Assets That Can Be Converted Into BNB (USER_DATA)
Post /sapi/v1/asset/dust-btc

https://developers.binance.com/docs/wallet/asset/assets-can-convert-bnb

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param accountType -  `SPOT` or `MARGIN`,default `SPOT`
@param recvWindow -
@return ApiGetAssetsThatCanBeConvertedIntoBnbRequest
*/
func (a *AssetAPIService) GetAssetsThatCanBeConvertedIntoBnb(ctx context.Context) ApiGetAssetsThatCanBeConvertedIntoBnbRequest {
	return ApiGetAssetsThatCanBeConvertedIntoBnbRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAssetsThatCanBeConvertedIntoBnbResponse
func (a *AssetAPIService) GetAssetsThatCanBeConvertedIntoBnbExecute(r ApiGetAssetsThatCanBeConvertedIntoBnbRequest) (*common.RestApiResponse[models.GetAssetsThatCanBeConvertedIntoBnbResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/dust-btc"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.accountType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "accountType", r.accountType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.GetAssetsThatCanBeConvertedIntoBnbResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetCloudMiningPaymentAndRefundHistoryRequest struct {
	ctx          context.Context
	ApiService   *AssetAPIService
	startTime    *int64
	endTime      *int64
	tranId       *int64
	clientTranId *string
	asset        *string
	current      *int64
	size         *int64
}

func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) StartTime(startTime int64) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) EndTime(endTime int64) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.endTime = &endTime
	return r
}

// The transaction id
func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) TranId(tranId int64) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.tranId = &tranId
	return r
}

// The unique flag
func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) ClientTranId(clientTranId string) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.clientTranId = &clientTranId
	return r
}

// If asset is blank, then query all positive assets user have.
func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) Asset(asset string) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.asset = &asset
	return r
}

// current page, default 1, the min value is 1
func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) Current(current int64) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.current = &current
	return r
}

// page size, default 10, the max value is 100
func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) Size(size int64) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	r.size = &size
	return r
}

func (r ApiGetCloudMiningPaymentAndRefundHistoryRequest) Execute() (*common.RestApiResponse[models.GetCloudMiningPaymentAndRefundHistoryResponse], error) {
	return r.ApiService.GetCloudMiningPaymentAndRefundHistoryExecute(r)
}

/*
GetCloudMiningPaymentAndRefundHistory Get Cloud-Mining payment and refund history (USER_DATA)
Get /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage

https://developers.binance.com/docs/wallet/asset/cloud-mining-payment-and-refund-history

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param startTime -
@param endTime -
@param tranId -  The transaction id
@param clientTranId -  The unique flag
@param asset -  If asset is blank, then query all positive assets user have.
@param current -  current page, default 1, the min value is 1
@param size -  page size, default 10, the max value is 100
@return ApiGetCloudMiningPaymentAndRefundHistoryRequest
*/
func (a *AssetAPIService) GetCloudMiningPaymentAndRefundHistory(ctx context.Context) ApiGetCloudMiningPaymentAndRefundHistoryRequest {
	return ApiGetCloudMiningPaymentAndRefundHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetCloudMiningPaymentAndRefundHistoryResponse
func (a *AssetAPIService) GetCloudMiningPaymentAndRefundHistoryExecute(r ApiGetCloudMiningPaymentAndRefundHistoryRequest) (*common.RestApiResponse[models.GetCloudMiningPaymentAndRefundHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	if r.tranId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tranId", r.tranId, "form", "")
	}
	if r.clientTranId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "clientTranId", r.clientTranId, "form", "")
	}
	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.current != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "current", r.current, "form", "")
	}
	if r.size != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "form", "")
	}

	resp, err := SendRequest[models.GetCloudMiningPaymentAndRefundHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiGetOpenSymbolListRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
}

func (r ApiGetOpenSymbolListRequest) Execute() (*common.RestApiResponse[models.GetOpenSymbolListResponse], error) {
	return r.ApiService.GetOpenSymbolListExecute(r)
}

/*
GetOpenSymbolList Get Open Symbol List (MARKET_DATA)
Get /sapi/v1/spot/open-symbol-list

https://developers.binance.com/docs/wallet/asset/open-symbol-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiGetOpenSymbolListRequest
*/
func (a *AssetAPIService) GetOpenSymbolList(ctx context.Context) ApiGetOpenSymbolListRequest {
	return ApiGetOpenSymbolListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetOpenSymbolListResponse
func (a *AssetAPIService) GetOpenSymbolListExecute(r ApiGetOpenSymbolListRequest) (*common.RestApiResponse[models.GetOpenSymbolListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/spot/open-symbol-list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.GetOpenSymbolListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserDelegationHistoryRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	email      *string
	startTime  *int64
	endTime    *int64
	type_      *string
	asset      *string
	current    *int64
	size       *int64
	recvWindow *int64
}

func (r ApiQueryUserDelegationHistoryRequest) Email(email string) ApiQueryUserDelegationHistoryRequest {
	r.email = &email
	return r
}

func (r ApiQueryUserDelegationHistoryRequest) StartTime(startTime int64) ApiQueryUserDelegationHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryUserDelegationHistoryRequest) EndTime(endTime int64) ApiQueryUserDelegationHistoryRequest {
	r.endTime = &endTime
	return r
}

// Delegate/Undelegate
func (r ApiQueryUserDelegationHistoryRequest) Type(type_ string) ApiQueryUserDelegationHistoryRequest {
	r.type_ = &type_
	return r
}

// If asset is blank, then query all positive assets user have.
func (r ApiQueryUserDelegationHistoryRequest) Asset(asset string) ApiQueryUserDelegationHistoryRequest {
	r.asset = &asset
	return r
}

// current page, default 1, the min value is 1
func (r ApiQueryUserDelegationHistoryRequest) Current(current int64) ApiQueryUserDelegationHistoryRequest {
	r.current = &current
	return r
}

// page size, default 10, the max value is 100
func (r ApiQueryUserDelegationHistoryRequest) Size(size int64) ApiQueryUserDelegationHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryUserDelegationHistoryRequest) RecvWindow(recvWindow int64) ApiQueryUserDelegationHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserDelegationHistoryRequest) Execute() (*common.RestApiResponse[models.QueryUserDelegationHistoryResponse], error) {
	return r.ApiService.QueryUserDelegationHistoryExecute(r)
}

/*
QueryUserDelegationHistory Query User Delegation History(For Master Account)(USER_DATA)
Get /sapi/v1/asset/custody/transfer-history

https://developers.binance.com/docs/wallet/asset/query-user-delegation

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param email -
@param startTime -
@param endTime -
@param type_ -  Delegate/Undelegate
@param asset -  If asset is blank, then query all positive assets user have.
@param current -  current page, default 1, the min value is 1
@param size -  page size, default 10, the max value is 100
@param recvWindow -
@return ApiQueryUserDelegationHistoryRequest
*/
func (a *AssetAPIService) QueryUserDelegationHistory(ctx context.Context) ApiQueryUserDelegationHistoryRequest {
	return ApiQueryUserDelegationHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserDelegationHistoryResponse
func (a *AssetAPIService) QueryUserDelegationHistoryExecute(r ApiQueryUserDelegationHistoryRequest) (*common.RestApiResponse[models.QueryUserDelegationHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/custody/transfer-history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.email == nil {
		return nil, common.ReportError("email is required and must be specified")
	}
	if r.startTime == nil {
		return nil, common.ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return nil, common.ReportError("endTime is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	if r.type_ != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
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

	resp, err := SendRequest[models.QueryUserDelegationHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserUniversalTransferHistoryRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	type_      *string
	startTime  *int64
	endTime    *int64
	current    *int64
	size       *int64
	fromSymbol *string
	toSymbol   *string
	recvWindow *int64
}

func (r ApiQueryUserUniversalTransferHistoryRequest) Type(type_ string) ApiQueryUserUniversalTransferHistoryRequest {
	r.type_ = &type_
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) StartTime(startTime int64) ApiQueryUserUniversalTransferHistoryRequest {
	r.startTime = &startTime
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) EndTime(endTime int64) ApiQueryUserUniversalTransferHistoryRequest {
	r.endTime = &endTime
	return r
}

// current page, default 1, the min value is 1
func (r ApiQueryUserUniversalTransferHistoryRequest) Current(current int64) ApiQueryUserUniversalTransferHistoryRequest {
	r.current = &current
	return r
}

// page size, default 10, the max value is 100
func (r ApiQueryUserUniversalTransferHistoryRequest) Size(size int64) ApiQueryUserUniversalTransferHistoryRequest {
	r.size = &size
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) FromSymbol(fromSymbol string) ApiQueryUserUniversalTransferHistoryRequest {
	r.fromSymbol = &fromSymbol
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) ToSymbol(toSymbol string) ApiQueryUserUniversalTransferHistoryRequest {
	r.toSymbol = &toSymbol
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) RecvWindow(recvWindow int64) ApiQueryUserUniversalTransferHistoryRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserUniversalTransferHistoryRequest) Execute() (*common.RestApiResponse[models.QueryUserUniversalTransferHistoryResponse], error) {
	return r.ApiService.QueryUserUniversalTransferHistoryExecute(r)
}

/*
QueryUserUniversalTransferHistory Query User Universal Transfer History(USER_DATA)
Get /sapi/v1/asset/transfer

https://developers.binance.com/docs/wallet/asset/query-user-universal-transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -
@param startTime -
@param endTime -
@param current -  current page, default 1, the min value is 1
@param size -  page size, default 10, the max value is 100
@param fromSymbol -
@param toSymbol -
@param recvWindow -
@return ApiQueryUserUniversalTransferHistoryRequest
*/
func (a *AssetAPIService) QueryUserUniversalTransferHistory(ctx context.Context) ApiQueryUserUniversalTransferHistoryRequest {
	return ApiQueryUserUniversalTransferHistoryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserUniversalTransferHistoryResponse
func (a *AssetAPIService) QueryUserUniversalTransferHistoryExecute(r ApiQueryUserUniversalTransferHistoryRequest) (*common.RestApiResponse[models.QueryUserUniversalTransferHistoryResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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
	if r.fromSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromSymbol", r.fromSymbol, "form", "")
	}
	if r.toSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toSymbol", r.toSymbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUserUniversalTransferHistoryResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiQueryUserWalletBalanceRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	quoteAsset *string
	recvWindow *int64
}

// &#x60;USDT&#x60;, &#x60;ETH&#x60;, &#x60;USDC&#x60;, &#x60;BNB&#x60;, etc. default &#x60;BTC&#x60;
func (r ApiQueryUserWalletBalanceRequest) QuoteAsset(quoteAsset string) ApiQueryUserWalletBalanceRequest {
	r.quoteAsset = &quoteAsset
	return r
}

func (r ApiQueryUserWalletBalanceRequest) RecvWindow(recvWindow int64) ApiQueryUserWalletBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiQueryUserWalletBalanceRequest) Execute() (*common.RestApiResponse[models.QueryUserWalletBalanceResponse], error) {
	return r.ApiService.QueryUserWalletBalanceExecute(r)
}

/*
QueryUserWalletBalance Query User Wallet Balance (USER_DATA)
Get /sapi/v1/asset/wallet/balance

https://developers.binance.com/docs/wallet/asset/Query-User-Wallet-Balance

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param quoteAsset -  `USDT`, `ETH`, `USDC`, `BNB`, etc. default `BTC`
@param recvWindow -
@return ApiQueryUserWalletBalanceRequest
*/
func (a *AssetAPIService) QueryUserWalletBalance(ctx context.Context) ApiQueryUserWalletBalanceRequest {
	return ApiQueryUserWalletBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QueryUserWalletBalanceResponse
func (a *AssetAPIService) QueryUserWalletBalanceExecute(r ApiQueryUserWalletBalanceRequest) (*common.RestApiResponse[models.QueryUserWalletBalanceResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/wallet/balance"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.quoteAsset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "quoteAsset", r.quoteAsset, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.QueryUserWalletBalanceResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest struct {
	ctx             context.Context
	ApiService      *AssetAPIService
	spotBNBBurn     *string
	interestBNBBurn *string
	recvWindow      *int64
}

// \&quot;true\&quot; or \&quot;false\&quot;; Determines whether to use BNB to pay for trading fees on SPOT
func (r ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest) SpotBNBBurn(spotBNBBurn string) ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest {
	r.spotBNBBurn = &spotBNBBurn
	return r
}

// \&quot;true\&quot; or \&quot;false\&quot;; Determines whether to use BNB to pay for margin loan&#39;s interest
func (r ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest) InterestBNBBurn(interestBNBBurn string) ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest {
	r.interestBNBBurn = &interestBNBBurn
	return r
}

func (r ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest) RecvWindow(recvWindow int64) ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest) Execute() (*common.RestApiResponse[models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse], error) {
	return r.ApiService.ToggleBnbBurnOnSpotTradeAndMarginInterestExecute(r)
}

/*
ToggleBnbBurnOnSpotTradeAndMarginInterest Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
Post /sapi/v1/bnbBurn

https://developers.binance.com/docs/wallet/asset/Toggle-BNB-Burn-On-Spot-Trade-And-Margin-Interest

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param spotBNBBurn -  \"true\" or \"false\"; Determines whether to use BNB to pay for trading fees on SPOT
@param interestBNBBurn -  \"true\" or \"false\"; Determines whether to use BNB to pay for margin loan's interest
@param recvWindow -
@return ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest
*/
func (a *AssetAPIService) ToggleBnbBurnOnSpotTradeAndMarginInterest(ctx context.Context) ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest {
	return ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToggleBnbBurnOnSpotTradeAndMarginInterestResponse
func (a *AssetAPIService) ToggleBnbBurnOnSpotTradeAndMarginInterestExecute(r ApiToggleBnbBurnOnSpotTradeAndMarginInterestRequest) (*common.RestApiResponse[models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/bnbBurn"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.spotBNBBurn != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "spotBNBBurn", r.spotBNBBurn, "form", "")
	}
	if r.interestBNBBurn != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "interestBNBBurn", r.interestBNBBurn, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiTradeFeeRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	symbol     *string
	recvWindow *int64
}

func (r ApiTradeFeeRequest) Symbol(symbol string) ApiTradeFeeRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTradeFeeRequest) RecvWindow(recvWindow int64) ApiTradeFeeRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiTradeFeeRequest) Execute() (*common.RestApiResponse[models.TradeFeeResponse], error) {
	return r.ApiService.TradeFeeExecute(r)
}

/*
TradeFee Trade Fee (USER_DATA)
Get /sapi/v1/asset/tradeFee

https://developers.binance.com/docs/wallet/asset/Trade-Fee

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol -
@param recvWindow -
@return ApiTradeFeeRequest
*/
func (a *AssetAPIService) TradeFee(ctx context.Context) ApiTradeFeeRequest {
	return ApiTradeFeeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TradeFeeResponse
func (a *AssetAPIService) TradeFeeExecute(r ApiTradeFeeRequest) (*common.RestApiResponse[models.TradeFeeResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/tradeFee"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.TradeFeeResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUserAssetRequest struct {
	ctx              context.Context
	ApiService       *AssetAPIService
	asset            *string
	needBtcValuation *bool
	recvWindow       *int64
}

// If asset is blank, then query all positive assets user have.
func (r ApiUserAssetRequest) Asset(asset string) ApiUserAssetRequest {
	r.asset = &asset
	return r
}

// Whether need btc valuation or not.
func (r ApiUserAssetRequest) NeedBtcValuation(needBtcValuation bool) ApiUserAssetRequest {
	r.needBtcValuation = &needBtcValuation
	return r
}

func (r ApiUserAssetRequest) RecvWindow(recvWindow int64) ApiUserAssetRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUserAssetRequest) Execute() (*common.RestApiResponse[models.UserAssetResponse], error) {
	return r.ApiService.UserAssetExecute(r)
}

/*
UserAsset User Asset (USER_DATA)
Post /sapi/v3/asset/getUserAsset

https://developers.binance.com/docs/wallet/asset/user-assets

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param asset -  If asset is blank, then query all positive assets user have.
@param needBtcValuation -  Whether need btc valuation or not.
@param recvWindow -
@return ApiUserAssetRequest
*/
func (a *AssetAPIService) UserAsset(ctx context.Context) ApiUserAssetRequest {
	return ApiUserAssetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserAssetResponse
func (a *AssetAPIService) UserAssetExecute(r ApiUserAssetRequest) (*common.RestApiResponse[models.UserAssetResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v3/asset/getUserAsset"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.asset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	}
	if r.needBtcValuation != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "needBtcValuation", r.needBtcValuation, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UserAssetResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiUserUniversalTransferRequest struct {
	ctx        context.Context
	ApiService *AssetAPIService
	type_      *string
	asset      *string
	amount     *float32
	fromSymbol *string
	toSymbol   *string
	recvWindow *int64
}

func (r ApiUserUniversalTransferRequest) Type(type_ string) ApiUserUniversalTransferRequest {
	r.type_ = &type_
	return r
}

func (r ApiUserUniversalTransferRequest) Asset(asset string) ApiUserUniversalTransferRequest {
	r.asset = &asset
	return r
}

func (r ApiUserUniversalTransferRequest) Amount(amount float32) ApiUserUniversalTransferRequest {
	r.amount = &amount
	return r
}

func (r ApiUserUniversalTransferRequest) FromSymbol(fromSymbol string) ApiUserUniversalTransferRequest {
	r.fromSymbol = &fromSymbol
	return r
}

func (r ApiUserUniversalTransferRequest) ToSymbol(toSymbol string) ApiUserUniversalTransferRequest {
	r.toSymbol = &toSymbol
	return r
}

func (r ApiUserUniversalTransferRequest) RecvWindow(recvWindow int64) ApiUserUniversalTransferRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiUserUniversalTransferRequest) Execute() (*common.RestApiResponse[models.UserUniversalTransferResponse], error) {
	return r.ApiService.UserUniversalTransferExecute(r)
}

/*
UserUniversalTransfer User Universal Transfer (USER_DATA)
Post /sapi/v1/asset/transfer

https://developers.binance.com/docs/wallet/asset/User-Universal-Transfer

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param type_ -
@param asset -
@param amount -
@param fromSymbol -
@param toSymbol -
@param recvWindow -
@return ApiUserUniversalTransferRequest
*/
func (a *AssetAPIService) UserUniversalTransfer(ctx context.Context) ApiUserUniversalTransferRequest {
	return ApiUserUniversalTransferRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserUniversalTransferResponse
func (a *AssetAPIService) UserUniversalTransferExecute(r ApiUserUniversalTransferRequest) (*common.RestApiResponse[models.UserUniversalTransferResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/asset/transfer"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.type_ == nil {
		return nil, common.ReportError("type_ is required and must be specified")
	}
	if r.asset == nil {
		return nil, common.ReportError("asset is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.fromSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "fromSymbol", r.fromSymbol, "form", "")
	}
	if r.toSymbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toSymbol", r.toSymbol, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.UserUniversalTransferResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
