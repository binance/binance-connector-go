/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package binanceminingrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/mining/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// MiningAPIService MiningAPI Service
type MiningAPIService Service

type ApiAccountListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	userName   *string
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiAccountListRequest) Algo(algo string) ApiAccountListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiAccountListRequest) UserName(userName string) ApiAccountListRequest {
	r.userName = &userName
	return r
}

func (r ApiAccountListRequest) RecvWindow(recvWindow int64) ApiAccountListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountListRequest) Execute() (*common.RestApiResponse[models.AccountListResponse], error) {
	return r.ApiService.AccountListExecute(r)
}

/*
AccountList Account List(USER_DATA)
Get /sapi/v1/mining/statistics/user/list

https://developers.binance.com/docs/mining/rest-api/Account-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param recvWindow -
@return ApiAccountListRequest
*/
func (a *MiningAPIService) AccountList(ctx context.Context) ApiAccountListRequest {
	return ApiAccountListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountListResponse
func (a *MiningAPIService) AccountListExecute(r ApiAccountListRequest) (*common.RestApiResponse[models.AccountListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/statistics/user/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.AccountListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAcquiringAlgorithmRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
}

func (r ApiAcquiringAlgorithmRequest) Execute() (*common.RestApiResponse[models.AcquiringAlgorithmResponse], error) {
	return r.ApiService.AcquiringAlgorithmExecute(r)
}

/*
AcquiringAlgorithm Acquiring Algorithm(MARKET_DATA)
Get /sapi/v1/mining/pub/algoList

https://developers.binance.com/docs/mining/rest-api/Acquiring-Algorithm

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAcquiringAlgorithmRequest
*/
func (a *MiningAPIService) AcquiringAlgorithm(ctx context.Context) ApiAcquiringAlgorithmRequest {
	return ApiAcquiringAlgorithmRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AcquiringAlgorithmResponse
func (a *MiningAPIService) AcquiringAlgorithmExecute(r ApiAcquiringAlgorithmRequest) (*common.RestApiResponse[models.AcquiringAlgorithmResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/pub/algoList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.AcquiringAlgorithmResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiAcquiringCoinnameRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
}

func (r ApiAcquiringCoinnameRequest) Execute() (*common.RestApiResponse[models.AcquiringCoinnameResponse], error) {
	return r.ApiService.AcquiringCoinnameExecute(r)
}

/*
AcquiringCoinname Acquiring CoinName(MARKET_DATA)
Get /sapi/v1/mining/pub/coinList

https://developers.binance.com/docs/mining/rest-api/Acquiring-CoinName

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return ApiAcquiringCoinnameRequest
*/
func (a *MiningAPIService) AcquiringCoinname(ctx context.Context) ApiAcquiringCoinnameRequest {
	return ApiAcquiringCoinnameRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AcquiringCoinnameResponse
func (a *MiningAPIService) AcquiringCoinnameExecute(r ApiAcquiringCoinnameRequest) (*common.RestApiResponse[models.AcquiringCoinnameResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/pub/coinList"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	resp, err := SendRequest[models.AcquiringCoinnameResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, false)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCancelHashrateResaleConfigurationRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	configId   *int64
	userName   *string
	recvWindow *int64
}

// Mining ID 168
func (r ApiCancelHashrateResaleConfigurationRequest) ConfigId(configId int64) ApiCancelHashrateResaleConfigurationRequest {
	r.configId = &configId
	return r
}

// Mining account test
func (r ApiCancelHashrateResaleConfigurationRequest) UserName(userName string) ApiCancelHashrateResaleConfigurationRequest {
	r.userName = &userName
	return r
}

func (r ApiCancelHashrateResaleConfigurationRequest) RecvWindow(recvWindow int64) ApiCancelHashrateResaleConfigurationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCancelHashrateResaleConfigurationRequest) Execute() (*common.RestApiResponse[models.CancelHashrateResaleConfigurationResponse], error) {
	return r.ApiService.CancelHashrateResaleConfigurationExecute(r)
}

/*
CancelHashrateResaleConfiguration Cancel hashrate resale configuration(USER_DATA)
Post /sapi/v1/mining/hash-transfer/config/cancel

https://developers.binance.com/docs/mining/rest-api/Cancel-hashrate-resale-configuration

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param configId -  Mining ID 168
@param userName -  Mining account test
@param recvWindow -
@return ApiCancelHashrateResaleConfigurationRequest
*/
func (a *MiningAPIService) CancelHashrateResaleConfiguration(ctx context.Context) ApiCancelHashrateResaleConfigurationRequest {
	return ApiCancelHashrateResaleConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CancelHashrateResaleConfigurationResponse
func (a *MiningAPIService) CancelHashrateResaleConfigurationExecute(r ApiCancelHashrateResaleConfigurationRequest) (*common.RestApiResponse[models.CancelHashrateResaleConfigurationResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/hash-transfer/config/cancel"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.configId == nil {
		return nil, common.ReportError("configId is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "configId", r.configId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CancelHashrateResaleConfigurationResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiEarningsListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	userName   *string
	coin       *string
	startDate  *int64
	endDate    *int64
	pageIndex  *int64
	pageSize   *int64
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiEarningsListRequest) Algo(algo string) ApiEarningsListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiEarningsListRequest) UserName(userName string) ApiEarningsListRequest {
	r.userName = &userName
	return r
}

// Coin Name
func (r ApiEarningsListRequest) Coin(coin string) ApiEarningsListRequest {
	r.coin = &coin
	return r
}

// Millisecond timestamp
func (r ApiEarningsListRequest) StartDate(startDate int64) ApiEarningsListRequest {
	r.startDate = &startDate
	return r
}

// Millisecond timestamp
func (r ApiEarningsListRequest) EndDate(endDate int64) ApiEarningsListRequest {
	r.endDate = &endDate
	return r
}

// Page number, empty default first page, starting from 1
func (r ApiEarningsListRequest) PageIndex(pageIndex int64) ApiEarningsListRequest {
	r.pageIndex = &pageIndex
	return r
}

// Min 10,Max 200
func (r ApiEarningsListRequest) PageSize(pageSize int64) ApiEarningsListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiEarningsListRequest) RecvWindow(recvWindow int64) ApiEarningsListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiEarningsListRequest) Execute() (*common.RestApiResponse[models.EarningsListResponse], error) {
	return r.ApiService.EarningsListExecute(r)
}

/*
EarningsList Earnings List(USER_DATA)
Get /sapi/v1/mining/payment/list

https://developers.binance.com/docs/mining/rest-api/Earnings-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param coin -  Coin Name
@param startDate -  Millisecond timestamp
@param endDate -  Millisecond timestamp
@param pageIndex -  Page number, empty default first page, starting from 1
@param pageSize -  Min 10,Max 200
@param recvWindow -
@return ApiEarningsListRequest
*/
func (a *MiningAPIService) EarningsList(ctx context.Context) ApiEarningsListRequest {
	return ApiEarningsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EarningsListResponse
func (a *MiningAPIService) EarningsListExecute(r ApiEarningsListRequest) (*common.RestApiResponse[models.EarningsListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/payment/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.startDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.EarningsListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiExtraBonusListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	userName   *string
	coin       *string
	startDate  *int64
	endDate    *int64
	pageIndex  *int64
	pageSize   *int64
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiExtraBonusListRequest) Algo(algo string) ApiExtraBonusListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiExtraBonusListRequest) UserName(userName string) ApiExtraBonusListRequest {
	r.userName = &userName
	return r
}

// Coin Name
func (r ApiExtraBonusListRequest) Coin(coin string) ApiExtraBonusListRequest {
	r.coin = &coin
	return r
}

// Millisecond timestamp
func (r ApiExtraBonusListRequest) StartDate(startDate int64) ApiExtraBonusListRequest {
	r.startDate = &startDate
	return r
}

// Millisecond timestamp
func (r ApiExtraBonusListRequest) EndDate(endDate int64) ApiExtraBonusListRequest {
	r.endDate = &endDate
	return r
}

// Page number, empty default first page, starting from 1
func (r ApiExtraBonusListRequest) PageIndex(pageIndex int64) ApiExtraBonusListRequest {
	r.pageIndex = &pageIndex
	return r
}

// Min 10,Max 200
func (r ApiExtraBonusListRequest) PageSize(pageSize int64) ApiExtraBonusListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiExtraBonusListRequest) RecvWindow(recvWindow int64) ApiExtraBonusListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiExtraBonusListRequest) Execute() (*common.RestApiResponse[models.ExtraBonusListResponse], error) {
	return r.ApiService.ExtraBonusListExecute(r)
}

/*
ExtraBonusList Extra Bonus List(USER_DATA)
Get /sapi/v1/mining/payment/other

https://developers.binance.com/docs/mining/rest-api/Extra-Bonus-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param coin -  Coin Name
@param startDate -  Millisecond timestamp
@param endDate -  Millisecond timestamp
@param pageIndex -  Page number, empty default first page, starting from 1
@param pageSize -  Min 10,Max 200
@param recvWindow -
@return ApiExtraBonusListRequest
*/
func (a *MiningAPIService) ExtraBonusList(ctx context.Context) ApiExtraBonusListRequest {
	return ApiExtraBonusListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExtraBonusListResponse
func (a *MiningAPIService) ExtraBonusListExecute(r ApiExtraBonusListRequest) (*common.RestApiResponse[models.ExtraBonusListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/payment/other"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.startDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.ExtraBonusListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiHashrateResaleDetailRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	configId   *int64
	pageIndex  *int64
	pageSize   *int64
	recvWindow *int64
}

// Mining ID 168
func (r ApiHashrateResaleDetailRequest) ConfigId(configId int64) ApiHashrateResaleDetailRequest {
	r.configId = &configId
	return r
}

// Page number, empty default first page, starting from 1
func (r ApiHashrateResaleDetailRequest) PageIndex(pageIndex int64) ApiHashrateResaleDetailRequest {
	r.pageIndex = &pageIndex
	return r
}

// Min 10,Max 200
func (r ApiHashrateResaleDetailRequest) PageSize(pageSize int64) ApiHashrateResaleDetailRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiHashrateResaleDetailRequest) RecvWindow(recvWindow int64) ApiHashrateResaleDetailRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiHashrateResaleDetailRequest) Execute() (*common.RestApiResponse[models.HashrateResaleDetailResponse], error) {
	return r.ApiService.HashrateResaleDetailExecute(r)
}

/*
HashrateResaleDetail Hashrate Resale Detail(USER_DATA)
Get /sapi/v1/mining/hash-transfer/profit/details

https://developers.binance.com/docs/mining/rest-api/Hashrate-Resale-Detail

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param configId -  Mining ID 168
@param pageIndex -  Page number, empty default first page, starting from 1
@param pageSize -  Min 10,Max 200
@param recvWindow -
@return ApiHashrateResaleDetailRequest
*/
func (a *MiningAPIService) HashrateResaleDetail(ctx context.Context) ApiHashrateResaleDetailRequest {
	return ApiHashrateResaleDetailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HashrateResaleDetailResponse
func (a *MiningAPIService) HashrateResaleDetailExecute(r ApiHashrateResaleDetailRequest) (*common.RestApiResponse[models.HashrateResaleDetailResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/hash-transfer/profit/details"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.configId == nil {
		return nil, common.ReportError("configId is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "configId", r.configId, "form", "")
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.HashrateResaleDetailResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiHashrateResaleListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	pageIndex  *int64
	pageSize   *int64
	recvWindow *int64
}

// Page number, empty default first page, starting from 1
func (r ApiHashrateResaleListRequest) PageIndex(pageIndex int64) ApiHashrateResaleListRequest {
	r.pageIndex = &pageIndex
	return r
}

// Min 10,Max 200
func (r ApiHashrateResaleListRequest) PageSize(pageSize int64) ApiHashrateResaleListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiHashrateResaleListRequest) RecvWindow(recvWindow int64) ApiHashrateResaleListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiHashrateResaleListRequest) Execute() (*common.RestApiResponse[models.HashrateResaleListResponse], error) {
	return r.ApiService.HashrateResaleListExecute(r)
}

/*
HashrateResaleList Hashrate Resale List
Get /sapi/v1/mining/hash-transfer/config/details/list

https://developers.binance.com/docs/mining/rest-api/Hashrate-Resale-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param pageIndex -  Page number, empty default first page, starting from 1
@param pageSize -  Min 10,Max 200
@param recvWindow -
@return ApiHashrateResaleListRequest
*/
func (a *MiningAPIService) HashrateResaleList(ctx context.Context) ApiHashrateResaleListRequest {
	return ApiHashrateResaleListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HashrateResaleListResponse
func (a *MiningAPIService) HashrateResaleListExecute(r ApiHashrateResaleListRequest) (*common.RestApiResponse[models.HashrateResaleListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/hash-transfer/config/details/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.HashrateResaleListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiHashrateResaleRequestRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	userName   *string
	algo       *string
	endDate    *int64
	startDate  *int64
	toPoolUser *string
	hashRate   *int64
	recvWindow *int64
}

// Mining account test
func (r ApiHashrateResaleRequestRequest) UserName(userName string) ApiHashrateResaleRequestRequest {
	r.userName = &userName
	return r
}

// Algorithm(sha256) sha256
func (r ApiHashrateResaleRequestRequest) Algo(algo string) ApiHashrateResaleRequestRequest {
	r.algo = &algo
	return r
}

// Resale End Time (Millisecond timestamp) 1617659086000
func (r ApiHashrateResaleRequestRequest) EndDate(endDate int64) ApiHashrateResaleRequestRequest {
	r.endDate = &endDate
	return r
}

// Resale Start Time(Millisecond timestamp) 1607659086000
func (r ApiHashrateResaleRequestRequest) StartDate(startDate int64) ApiHashrateResaleRequestRequest {
	r.startDate = &startDate
	return r
}

// Mining Account S19pro
func (r ApiHashrateResaleRequestRequest) ToPoolUser(toPoolUser string) ApiHashrateResaleRequestRequest {
	r.toPoolUser = &toPoolUser
	return r
}

// Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) 100000000
func (r ApiHashrateResaleRequestRequest) HashRate(hashRate int64) ApiHashrateResaleRequestRequest {
	r.hashRate = &hashRate
	return r
}

func (r ApiHashrateResaleRequestRequest) RecvWindow(recvWindow int64) ApiHashrateResaleRequestRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiHashrateResaleRequestRequest) Execute() (*common.RestApiResponse[models.HashrateResaleRequestResponse], error) {
	return r.ApiService.HashrateResaleRequestExecute(r)
}

/*
HashrateResaleRequest Hashrate Resale Request(USER_DATA)
Post /sapi/v1/mining/hash-transfer/config

https://developers.binance.com/docs/mining/rest-api/Hashrate-Resale-Request

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param userName -  Mining account test
@param algo -  Algorithm(sha256) sha256
@param endDate -  Resale End Time (Millisecond timestamp) 1617659086000
@param startDate -  Resale Start Time(Millisecond timestamp) 1607659086000
@param toPoolUser -  Mining Account S19pro
@param hashRate -  Resale hashrate h/s must be transferred (BTC is greater than 500000000000 ETH is greater than 500000) 100000000
@param recvWindow -
@return ApiHashrateResaleRequestRequest
*/
func (a *MiningAPIService) HashrateResaleRequest(ctx context.Context) ApiHashrateResaleRequestRequest {
	return ApiHashrateResaleRequestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HashrateResaleRequestResponse
func (a *MiningAPIService) HashrateResaleRequestExecute(r ApiHashrateResaleRequestRequest) (*common.RestApiResponse[models.HashrateResaleRequestResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/hash-transfer/config"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}
	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.endDate == nil {
		return nil, common.ReportError("endDate is required and must be specified")
	}
	if r.startDate == nil {
		return nil, common.ReportError("startDate is required and must be specified")
	}
	if r.toPoolUser == nil {
		return nil, common.ReportError("toPoolUser is required and must be specified")
	}
	if r.hashRate == nil {
		return nil, common.ReportError("hashRate is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "toPoolUser", r.toPoolUser, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "hashRate", r.hashRate, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.HashrateResaleRequestResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiMiningAccountEarningRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	startDate  *int64
	endDate    *int64
	pageIndex  *int64
	pageSize   *int64
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiMiningAccountEarningRequest) Algo(algo string) ApiMiningAccountEarningRequest {
	r.algo = &algo
	return r
}

// Millisecond timestamp
func (r ApiMiningAccountEarningRequest) StartDate(startDate int64) ApiMiningAccountEarningRequest {
	r.startDate = &startDate
	return r
}

// Millisecond timestamp
func (r ApiMiningAccountEarningRequest) EndDate(endDate int64) ApiMiningAccountEarningRequest {
	r.endDate = &endDate
	return r
}

// Page number, empty default first page, starting from 1
func (r ApiMiningAccountEarningRequest) PageIndex(pageIndex int64) ApiMiningAccountEarningRequest {
	r.pageIndex = &pageIndex
	return r
}

// Min 10,Max 200
func (r ApiMiningAccountEarningRequest) PageSize(pageSize int64) ApiMiningAccountEarningRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiMiningAccountEarningRequest) RecvWindow(recvWindow int64) ApiMiningAccountEarningRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiMiningAccountEarningRequest) Execute() (*common.RestApiResponse[models.MiningAccountEarningResponse], error) {
	return r.ApiService.MiningAccountEarningExecute(r)
}

/*
MiningAccountEarning Mining Account Earning(USER_DATA)
Get /sapi/v1/mining/payment/uid

https://developers.binance.com/docs/mining/rest-api/Mining-Account-Earning

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param startDate -  Millisecond timestamp
@param endDate -  Millisecond timestamp
@param pageIndex -  Page number, empty default first page, starting from 1
@param pageSize -  Min 10,Max 200
@param recvWindow -
@return ApiMiningAccountEarningRequest
*/
func (a *MiningAPIService) MiningAccountEarning(ctx context.Context) ApiMiningAccountEarningRequest {
	return ApiMiningAccountEarningRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MiningAccountEarningResponse
func (a *MiningAPIService) MiningAccountEarningExecute(r ApiMiningAccountEarningRequest) (*common.RestApiResponse[models.MiningAccountEarningResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/payment/uid"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	if r.startDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.pageSize != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.MiningAccountEarningResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRequestForDetailMinerListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	userName   *string
	workerName *string
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiRequestForDetailMinerListRequest) Algo(algo string) ApiRequestForDetailMinerListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiRequestForDetailMinerListRequest) UserName(userName string) ApiRequestForDetailMinerListRequest {
	r.userName = &userName
	return r
}

// Miner’s name(required) bhdc1.16A10404B
func (r ApiRequestForDetailMinerListRequest) WorkerName(workerName string) ApiRequestForDetailMinerListRequest {
	r.workerName = &workerName
	return r
}

func (r ApiRequestForDetailMinerListRequest) RecvWindow(recvWindow int64) ApiRequestForDetailMinerListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRequestForDetailMinerListRequest) Execute() (*common.RestApiResponse[models.RequestForDetailMinerListResponse], error) {
	return r.ApiService.RequestForDetailMinerListExecute(r)
}

/*
RequestForDetailMinerList Request for Detail Miner List(USER_DATA)
Get /sapi/v1/mining/worker/detail

https://developers.binance.com/docs/mining/rest-api/Request-for-Detail-Miner-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param workerName -  Miner’s name(required) bhdc1.16A10404B
@param recvWindow -
@return ApiRequestForDetailMinerListRequest
*/
func (a *MiningAPIService) RequestForDetailMinerList(ctx context.Context) ApiRequestForDetailMinerListRequest {
	return ApiRequestForDetailMinerListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestForDetailMinerListResponse
func (a *MiningAPIService) RequestForDetailMinerListExecute(r ApiRequestForDetailMinerListRequest) (*common.RestApiResponse[models.RequestForDetailMinerListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/worker/detail"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}
	if r.workerName == nil {
		return nil, common.ReportError("workerName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workerName", r.workerName, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RequestForDetailMinerListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRequestForMinerListRequest struct {
	ctx          context.Context
	ApiService   *MiningAPIService
	algo         *string
	userName     *string
	pageIndex    *int64
	sort         *int64
	sortColumn   *int64
	workerStatus *int64
	recvWindow   *int64
}

// Algorithm(sha256) sha256
func (r ApiRequestForMinerListRequest) Algo(algo string) ApiRequestForMinerListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiRequestForMinerListRequest) UserName(userName string) ApiRequestForMinerListRequest {
	r.userName = &userName
	return r
}

// Page number, empty default first page, starting from 1
func (r ApiRequestForMinerListRequest) PageIndex(pageIndex int64) ApiRequestForMinerListRequest {
	r.pageIndex = &pageIndex
	return r
}

// sort sequence(default&#x3D;0)0 positive sequence，1 negative sequence
func (r ApiRequestForMinerListRequest) Sort(sort int64) ApiRequestForMinerListRequest {
	r.sort = &sort
	return r
}

// Sort by( default 1): &lt;br&gt;&lt;/br&gt;1: miner name, &lt;br&gt;&lt;/br&gt;2: real-time computing power, &lt;br&gt;&lt;/br&gt;3: daily average computing power, &lt;br&gt;&lt;/br&gt;4: real-time rejection rate, &lt;br&gt;&lt;/br&gt;5: last submission time
func (r ApiRequestForMinerListRequest) SortColumn(sortColumn int64) ApiRequestForMinerListRequest {
	r.sortColumn = &sortColumn
	return r
}

// miners status(default&#x3D;0),0 all，1 valid，2 invalid，3 failure
func (r ApiRequestForMinerListRequest) WorkerStatus(workerStatus int64) ApiRequestForMinerListRequest {
	r.workerStatus = &workerStatus
	return r
}

func (r ApiRequestForMinerListRequest) RecvWindow(recvWindow int64) ApiRequestForMinerListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRequestForMinerListRequest) Execute() (*common.RestApiResponse[models.RequestForMinerListResponse], error) {
	return r.ApiService.RequestForMinerListExecute(r)
}

/*
RequestForMinerList Request for Miner List(USER_DATA)
Get /sapi/v1/mining/worker/list

https://developers.binance.com/docs/mining/rest-api/Request-for-Miner-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param pageIndex -  Page number, empty default first page, starting from 1
@param sort -  sort sequence(default=0)0 positive sequence，1 negative sequence
@param sortColumn -  Sort by( default 1): <br></br>1: miner name, <br></br>2: real-time computing power, <br></br>3: daily average computing power, <br></br>4: real-time rejection rate, <br></br>5: last submission time
@param workerStatus -  miners status(default=0),0 all，1 valid，2 invalid，3 failure
@param recvWindow -
@return ApiRequestForMinerListRequest
*/
func (a *MiningAPIService) RequestForMinerList(ctx context.Context) ApiRequestForMinerListRequest {
	return ApiRequestForMinerListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RequestForMinerListResponse
func (a *MiningAPIService) RequestForMinerListExecute(r ApiRequestForMinerListRequest) (*common.RestApiResponse[models.RequestForMinerListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/worker/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.pageIndex != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pageIndex", r.pageIndex, "form", "")
	}
	if r.sort != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.sortColumn != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "sortColumn", r.sortColumn, "form", "")
	}
	if r.workerStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "workerStatus", r.workerStatus, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RequestForMinerListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiStatisticListRequest struct {
	ctx        context.Context
	ApiService *MiningAPIService
	algo       *string
	userName   *string
	recvWindow *int64
}

// Algorithm(sha256) sha256
func (r ApiStatisticListRequest) Algo(algo string) ApiStatisticListRequest {
	r.algo = &algo
	return r
}

// Mining account test
func (r ApiStatisticListRequest) UserName(userName string) ApiStatisticListRequest {
	r.userName = &userName
	return r
}

func (r ApiStatisticListRequest) RecvWindow(recvWindow int64) ApiStatisticListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiStatisticListRequest) Execute() (*common.RestApiResponse[models.StatisticListResponse], error) {
	return r.ApiService.StatisticListExecute(r)
}

/*
StatisticList Statistic List(USER_DATA)
Get /sapi/v1/mining/statistics/user/status

https://developers.binance.com/docs/mining/rest-api/Statistic-List

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param algo -  Algorithm(sha256) sha256
@param userName -  Mining account test
@param recvWindow -
@return ApiStatisticListRequest
*/
func (a *MiningAPIService) StatisticList(ctx context.Context) ApiStatisticListRequest {
	return ApiStatisticListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StatisticListResponse
func (a *MiningAPIService) StatisticListExecute(r ApiStatisticListRequest) (*common.RestApiResponse[models.StatisticListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/mining/statistics/user/status"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.algo == nil {
		return nil, common.ReportError("algo is required and must be specified")
	}
	if r.userName == nil {
		return nil, common.ReportError("userName is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "algo", r.algo, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "userName", r.userName, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.StatisticListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
