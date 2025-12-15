/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package binancewalletrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// TravelRuleAPIService TravelRuleAPI Service
type TravelRuleAPIService Service

type ApiBrokerWithdrawRequest struct {
	ctx                context.Context
	ApiService         *TravelRuleAPIService
	address            *string
	coin               *string
	amount             *float32
	withdrawOrderId    *string
	questionnaire      *string
	originatorPii      *string
	signature          *string
	addressTag         *string
	network            *string
	addressName        *string
	transactionFeeFlag *bool
	walletType         *int64
}

func (r ApiBrokerWithdrawRequest) Address(address string) ApiBrokerWithdrawRequest {
	r.address = &address
	return r
}

func (r ApiBrokerWithdrawRequest) Coin(coin string) ApiBrokerWithdrawRequest {
	r.coin = &coin
	return r
}

func (r ApiBrokerWithdrawRequest) Amount(amount float32) ApiBrokerWithdrawRequest {
	r.amount = &amount
	return r
}

// withdrawID defined by the client (i.e. client&#39;s internal withdrawID)
func (r ApiBrokerWithdrawRequest) WithdrawOrderId(withdrawOrderId string) ApiBrokerWithdrawRequest {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

// JSON format questionnaire answers.
func (r ApiBrokerWithdrawRequest) Questionnaire(questionnaire string) ApiBrokerWithdrawRequest {
	r.questionnaire = &questionnaire
	return r
}

// JSON format originator Pii, see StandardPii section below
func (r ApiBrokerWithdrawRequest) OriginatorPii(originatorPii string) ApiBrokerWithdrawRequest {
	r.originatorPii = &originatorPii
	return r
}

// Must be the last parameter.
func (r ApiBrokerWithdrawRequest) Signature(signature string) ApiBrokerWithdrawRequest {
	r.signature = &signature
	return r
}

// Secondary address identifier for coins like XRP,XMR etc.
func (r ApiBrokerWithdrawRequest) AddressTag(addressTag string) ApiBrokerWithdrawRequest {
	r.addressTag = &addressTag
	return r
}

func (r ApiBrokerWithdrawRequest) Network(network string) ApiBrokerWithdrawRequest {
	r.network = &network
	return r
}

// Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60;
func (r ApiBrokerWithdrawRequest) AddressName(addressName string) ApiBrokerWithdrawRequest {
	r.addressName = &addressName
	return r
}

// When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;.
func (r ApiBrokerWithdrawRequest) TransactionFeeFlag(transactionFeeFlag bool) ApiBrokerWithdrawRequest {
	r.transactionFeeFlag = &transactionFeeFlag
	return r
}

// The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit
func (r ApiBrokerWithdrawRequest) WalletType(walletType int64) ApiBrokerWithdrawRequest {
	r.walletType = &walletType
	return r
}

func (r ApiBrokerWithdrawRequest) Execute() (*common.RestApiResponse[models.BrokerWithdrawResponse], error) {
	return r.ApiService.BrokerWithdrawExecute(r)
}

/*
BrokerWithdraw Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)
Post /sapi/v1/localentity/broker/withdraw/apply

https://developers.binance.com/docs/wallet/travel-rule/Broker-Withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param address -
@param coin -
@param amount -
@param withdrawOrderId -  withdrawID defined by the client (i.e. client's internal withdrawID)
@param questionnaire -  JSON format questionnaire answers.
@param originatorPii -  JSON format originator Pii, see StandardPii section below
@param signature -  Must be the last parameter.
@param addressTag -  Secondary address identifier for coins like XRP,XMR etc.
@param network -
@param addressName -  Description of the address. Address book cap is 200, space in name should be encoded into `%20`
@param transactionFeeFlag -  When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`.
@param walletType -  The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit
@return ApiBrokerWithdrawRequest
*/
func (a *TravelRuleAPIService) BrokerWithdraw(ctx context.Context) ApiBrokerWithdrawRequest {
	return ApiBrokerWithdrawRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrokerWithdrawResponse
func (a *TravelRuleAPIService) BrokerWithdrawExecute(r ApiBrokerWithdrawRequest) (*common.RestApiResponse[models.BrokerWithdrawResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/broker/withdraw/apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.address == nil {
		return nil, common.ReportError("address is required and must be specified")
	}
	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.withdrawOrderId == nil {
		return nil, common.ReportError("withdrawOrderId is required and must be specified")
	}
	if r.questionnaire == nil {
		return nil, common.ReportError("questionnaire is required and must be specified")
	}
	if r.originatorPii == nil {
		return nil, common.ReportError("originatorPii is required and must be specified")
	}
	if r.signature == nil {
		return nil, common.ReportError("signature is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "form", "")
	if r.addressTag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "addressTag", r.addressTag, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.addressName != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "addressName", r.addressName, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	if r.transactionFeeFlag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transactionFeeFlag", r.transactionFeeFlag, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "questionnaire", r.questionnaire, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "originatorPii", r.originatorPii, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "signature", r.signature, "form", "")

	resp, err := SendRequest[models.BrokerWithdrawResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCheckQuestionnaireRequirementsRequest struct {
	ctx        context.Context
	ApiService *TravelRuleAPIService
	recvWindow *int64
}

func (r ApiCheckQuestionnaireRequirementsRequest) RecvWindow(recvWindow int64) ApiCheckQuestionnaireRequirementsRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCheckQuestionnaireRequirementsRequest) Execute() (*common.RestApiResponse[models.CheckQuestionnaireRequirementsResponse], error) {
	return r.ApiService.CheckQuestionnaireRequirementsExecute(r)
}

/*
CheckQuestionnaireRequirements Check Questionnaire Requirements (for local entities that require travel rule) (supporting network) (USER_DATA)
Get /sapi/v1/localentity/questionnaire-requirements

https://developers.binance.com/docs/wallet/travel-rule/questionnaire-requirements

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiCheckQuestionnaireRequirementsRequest
*/
func (a *TravelRuleAPIService) CheckQuestionnaireRequirements(ctx context.Context) ApiCheckQuestionnaireRequirementsRequest {
	return ApiCheckQuestionnaireRequirementsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CheckQuestionnaireRequirementsResponse
func (a *TravelRuleAPIService) CheckQuestionnaireRequirementsExecute(r ApiCheckQuestionnaireRequirementsRequest) (*common.RestApiResponse[models.CheckQuestionnaireRequirementsResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/questionnaire-requirements"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CheckQuestionnaireRequirementsResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDepositHistoryTravelRuleRequest struct {
	ctx                  context.Context
	ApiService           *TravelRuleAPIService
	trId                 *string
	txId                 *string
	tranId               *string
	network              *string
	coin                 *string
	travelRuleStatus     *int64
	pendingQuestionnaire *bool
	startTime            *int64
	endTime              *int64
	offset               *int64
	limit                *int64
}

// Comma(,) separated list of travel rule record Ids.
func (r ApiDepositHistoryTravelRuleRequest) TrId(trId string) ApiDepositHistoryTravelRuleRequest {
	r.trId = &trId
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) TxId(txId string) ApiDepositHistoryTravelRuleRequest {
	r.txId = &txId
	return r
}

// Comma(,) separated list of wallet tran Ids.
func (r ApiDepositHistoryTravelRuleRequest) TranId(tranId string) ApiDepositHistoryTravelRuleRequest {
	r.tranId = &tranId
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) Network(network string) ApiDepositHistoryTravelRuleRequest {
	r.network = &network
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) Coin(coin string) ApiDepositHistoryTravelRuleRequest {
	r.coin = &coin
	return r
}

// 0:Completed,1:Pending,2:Failed
func (r ApiDepositHistoryTravelRuleRequest) TravelRuleStatus(travelRuleStatus int64) ApiDepositHistoryTravelRuleRequest {
	r.travelRuleStatus = &travelRuleStatus
	return r
}

// true: Only return records that pending deposit questionnaire. false/not provided: return all records.
func (r ApiDepositHistoryTravelRuleRequest) PendingQuestionnaire(pendingQuestionnaire bool) ApiDepositHistoryTravelRuleRequest {
	r.pendingQuestionnaire = &pendingQuestionnaire
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) StartTime(startTime int64) ApiDepositHistoryTravelRuleRequest {
	r.startTime = &startTime
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) EndTime(endTime int64) ApiDepositHistoryTravelRuleRequest {
	r.endTime = &endTime
	return r
}

// Default: 0
func (r ApiDepositHistoryTravelRuleRequest) Offset(offset int64) ApiDepositHistoryTravelRuleRequest {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiDepositHistoryTravelRuleRequest) Limit(limit int64) ApiDepositHistoryTravelRuleRequest {
	r.limit = &limit
	return r
}

func (r ApiDepositHistoryTravelRuleRequest) Execute() (*common.RestApiResponse[models.DepositHistoryTravelRuleResponse], error) {
	return r.ApiService.DepositHistoryTravelRuleExecute(r)
}

/*
DepositHistoryTravelRule Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)
Get /sapi/v1/localentity/deposit/history

https://developers.binance.com/docs/wallet/travel-rule/Deposit-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param trId -  Comma(,) separated list of travel rule record Ids.
@param txId -
@param tranId -  Comma(,) separated list of wallet tran Ids.
@param network -
@param coin -
@param travelRuleStatus -  0:Completed,1:Pending,2:Failed
@param pendingQuestionnaire -  true: Only return records that pending deposit questionnaire. false/not provided: return all records.
@param startTime -
@param endTime -
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@return ApiDepositHistoryTravelRuleRequest
*/
func (a *TravelRuleAPIService) DepositHistoryTravelRule(ctx context.Context) ApiDepositHistoryTravelRuleRequest {
	return ApiDepositHistoryTravelRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositHistoryTravelRuleResponse
func (a *TravelRuleAPIService) DepositHistoryTravelRuleExecute(r ApiDepositHistoryTravelRuleRequest) (*common.RestApiResponse[models.DepositHistoryTravelRuleResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/deposit/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.trId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trId", r.trId, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.tranId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tranId", r.tranId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.travelRuleStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "travelRuleStatus", r.travelRuleStatus, "form", "")
	}
	if r.pendingQuestionnaire != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "pendingQuestionnaire", r.pendingQuestionnaire, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.DepositHistoryTravelRuleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiDepositHistoryV2Request struct {
	ctx                   context.Context
	ApiService            *TravelRuleAPIService
	depositId             *string
	txId                  *string
	network               *string
	coin                  *string
	retrieveQuestionnaire *bool
	startTime             *int64
	endTime               *int64
	offset                *int64
	limit                 *int64
}

// Comma(,) separated list of wallet tran Ids.
func (r ApiDepositHistoryV2Request) DepositId(depositId string) ApiDepositHistoryV2Request {
	r.depositId = &depositId
	return r
}

func (r ApiDepositHistoryV2Request) TxId(txId string) ApiDepositHistoryV2Request {
	r.txId = &txId
	return r
}

func (r ApiDepositHistoryV2Request) Network(network string) ApiDepositHistoryV2Request {
	r.network = &network
	return r
}

func (r ApiDepositHistoryV2Request) Coin(coin string) ApiDepositHistoryV2Request {
	r.coin = &coin
	return r
}

// true: return &#x60;questionnaire&#x60; within response.
func (r ApiDepositHistoryV2Request) RetrieveQuestionnaire(retrieveQuestionnaire bool) ApiDepositHistoryV2Request {
	r.retrieveQuestionnaire = &retrieveQuestionnaire
	return r
}

func (r ApiDepositHistoryV2Request) StartTime(startTime int64) ApiDepositHistoryV2Request {
	r.startTime = &startTime
	return r
}

func (r ApiDepositHistoryV2Request) EndTime(endTime int64) ApiDepositHistoryV2Request {
	r.endTime = &endTime
	return r
}

// Default: 0
func (r ApiDepositHistoryV2Request) Offset(offset int64) ApiDepositHistoryV2Request {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiDepositHistoryV2Request) Limit(limit int64) ApiDepositHistoryV2Request {
	r.limit = &limit
	return r
}

func (r ApiDepositHistoryV2Request) Execute() (*common.RestApiResponse[models.DepositHistoryV2Response], error) {
	return r.ApiService.DepositHistoryV2Execute(r)
}

/*
DepositHistoryV2 Deposit History V2 (for local entities that required travel rule) (supporting network) (USER_DATA)
Get /sapi/v2/localentity/deposit/history

https://developers.binance.com/docs/wallet/travel-rule/Deposit-History-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param depositId -  Comma(,) separated list of wallet tran Ids.
@param txId -
@param network -
@param coin -
@param retrieveQuestionnaire -  true: return `questionnaire` within response.
@param startTime -
@param endTime -
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@return ApiDepositHistoryV2Request
*/
func (a *TravelRuleAPIService) DepositHistoryV2(ctx context.Context) ApiDepositHistoryV2Request {
	return ApiDepositHistoryV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DepositHistoryV2Response
func (a *TravelRuleAPIService) DepositHistoryV2Execute(r ApiDepositHistoryV2Request) (*common.RestApiResponse[models.DepositHistoryV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/localentity/deposit/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.depositId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "depositId", r.depositId, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.retrieveQuestionnaire != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "retrieveQuestionnaire", r.retrieveQuestionnaire, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.DepositHistoryV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchAddressVerificationListRequest struct {
	ctx        context.Context
	ApiService *TravelRuleAPIService
	recvWindow *int64
}

func (r ApiFetchAddressVerificationListRequest) RecvWindow(recvWindow int64) ApiFetchAddressVerificationListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFetchAddressVerificationListRequest) Execute() (*common.RestApiResponse[models.FetchAddressVerificationListResponse], error) {
	return r.ApiService.FetchAddressVerificationListExecute(r)
}

/*
FetchAddressVerificationList Fetch address verification list (USER_DATA)
Get /sapi/v1/addressVerify/list

https://developers.binance.com/docs/wallet/travel-rule/address-verification-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFetchAddressVerificationListRequest
*/
func (a *TravelRuleAPIService) FetchAddressVerificationList(ctx context.Context) ApiFetchAddressVerificationListRequest {
	return ApiFetchAddressVerificationListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchAddressVerificationListResponse
func (a *TravelRuleAPIService) FetchAddressVerificationListExecute(r ApiFetchAddressVerificationListRequest) (*common.RestApiResponse[models.FetchAddressVerificationListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/addressVerify/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FetchAddressVerificationListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubmitDepositQuestionnaireRequest struct {
	ctx            context.Context
	ApiService     *TravelRuleAPIService
	subAccountId   *string
	depositId      *string
	questionnaire  *string
	beneficiaryPii *string
	signature      *string
	network        *string
	coin           *string
	amount         *float32
	address        *string
	addressTag     *string
}

// External user ID.
func (r ApiSubmitDepositQuestionnaireRequest) SubAccountId(subAccountId string) ApiSubmitDepositQuestionnaireRequest {
	r.subAccountId = &subAccountId
	return r
}

// Wallet deposit ID.
func (r ApiSubmitDepositQuestionnaireRequest) DepositId(depositId string) ApiSubmitDepositQuestionnaireRequest {
	r.depositId = &depositId
	return r
}

// JSON format questionnaire answers.
func (r ApiSubmitDepositQuestionnaireRequest) Questionnaire(questionnaire string) ApiSubmitDepositQuestionnaireRequest {
	r.questionnaire = &questionnaire
	return r
}

// JSON format beneficiary Pii.
func (r ApiSubmitDepositQuestionnaireRequest) BeneficiaryPii(beneficiaryPii string) ApiSubmitDepositQuestionnaireRequest {
	r.beneficiaryPii = &beneficiaryPii
	return r
}

// Must be the last parameter.
func (r ApiSubmitDepositQuestionnaireRequest) Signature(signature string) ApiSubmitDepositQuestionnaireRequest {
	r.signature = &signature
	return r
}

func (r ApiSubmitDepositQuestionnaireRequest) Network(network string) ApiSubmitDepositQuestionnaireRequest {
	r.network = &network
	return r
}

func (r ApiSubmitDepositQuestionnaireRequest) Coin(coin string) ApiSubmitDepositQuestionnaireRequest {
	r.coin = &coin
	return r
}

func (r ApiSubmitDepositQuestionnaireRequest) Amount(amount float32) ApiSubmitDepositQuestionnaireRequest {
	r.amount = &amount
	return r
}

func (r ApiSubmitDepositQuestionnaireRequest) Address(address string) ApiSubmitDepositQuestionnaireRequest {
	r.address = &address
	return r
}

// Secondary address identifier for coins like XRP,XMR etc.
func (r ApiSubmitDepositQuestionnaireRequest) AddressTag(addressTag string) ApiSubmitDepositQuestionnaireRequest {
	r.addressTag = &addressTag
	return r
}

func (r ApiSubmitDepositQuestionnaireRequest) Execute() (*common.RestApiResponse[models.SubmitDepositQuestionnaireResponse], error) {
	return r.ApiService.SubmitDepositQuestionnaireExecute(r)
}

/*
SubmitDepositQuestionnaire Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
Put /sapi/v1/localentity/broker/deposit/provide-info

https://developers.binance.com/docs/wallet/travel-rule/deposit-provide-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param subAccountId -  External user ID.
@param depositId -  Wallet deposit ID.
@param questionnaire -  JSON format questionnaire answers.
@param beneficiaryPii -  JSON format beneficiary Pii.
@param signature -  Must be the last parameter.
@param network -
@param coin -
@param amount -
@param address -
@param addressTag -  Secondary address identifier for coins like XRP,XMR etc.
@return ApiSubmitDepositQuestionnaireRequest
*/
func (a *TravelRuleAPIService) SubmitDepositQuestionnaire(ctx context.Context) ApiSubmitDepositQuestionnaireRequest {
	return ApiSubmitDepositQuestionnaireRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubmitDepositQuestionnaireResponse
func (a *TravelRuleAPIService) SubmitDepositQuestionnaireExecute(r ApiSubmitDepositQuestionnaireRequest) (*common.RestApiResponse[models.SubmitDepositQuestionnaireResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/broker/deposit/provide-info"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.subAccountId == nil {
		return nil, common.ReportError("subAccountId is required and must be specified")
	}
	if r.depositId == nil {
		return nil, common.ReportError("depositId is required and must be specified")
	}
	if r.questionnaire == nil {
		return nil, common.ReportError("questionnaire is required and must be specified")
	}
	if r.beneficiaryPii == nil {
		return nil, common.ReportError("beneficiaryPii is required and must be specified")
	}
	if r.signature == nil {
		return nil, common.ReportError("signature is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "subAccountId", r.subAccountId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "depositId", r.depositId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "questionnaire", r.questionnaire, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "beneficiaryPii", r.beneficiaryPii, "form", "")
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.amount != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	}
	if r.address != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "form", "")
	}
	if r.addressTag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "addressTag", r.addressTag, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "signature", r.signature, "form", "")

	resp, err := SendRequest[models.SubmitDepositQuestionnaireResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiSubmitDepositQuestionnaireTravelRuleRequest struct {
	ctx           context.Context
	ApiService    *TravelRuleAPIService
	tranId        *int64
	questionnaire *string
}

// Wallet tran ID
func (r ApiSubmitDepositQuestionnaireTravelRuleRequest) TranId(tranId int64) ApiSubmitDepositQuestionnaireTravelRuleRequest {
	r.tranId = &tranId
	return r
}

// JSON format questionnaire answers.
func (r ApiSubmitDepositQuestionnaireTravelRuleRequest) Questionnaire(questionnaire string) ApiSubmitDepositQuestionnaireTravelRuleRequest {
	r.questionnaire = &questionnaire
	return r
}

func (r ApiSubmitDepositQuestionnaireTravelRuleRequest) Execute() (*common.RestApiResponse[models.SubmitDepositQuestionnaireTravelRuleResponse], error) {
	return r.ApiService.SubmitDepositQuestionnaireTravelRuleExecute(r)
}

/*
SubmitDepositQuestionnaireTravelRule Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
Put /sapi/v1/localentity/deposit/provide-info

https://developers.binance.com/docs/wallet/travel-rule/deposit-provide-info

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param tranId -  Wallet tran ID
@param questionnaire -  JSON format questionnaire answers.
@return ApiSubmitDepositQuestionnaireTravelRuleRequest
*/
func (a *TravelRuleAPIService) SubmitDepositQuestionnaireTravelRule(ctx context.Context) ApiSubmitDepositQuestionnaireTravelRuleRequest {
	return ApiSubmitDepositQuestionnaireTravelRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubmitDepositQuestionnaireTravelRuleResponse
func (a *TravelRuleAPIService) SubmitDepositQuestionnaireTravelRuleExecute(r ApiSubmitDepositQuestionnaireTravelRuleRequest) (*common.RestApiResponse[models.SubmitDepositQuestionnaireTravelRuleResponse], error) {
	localVarHTTPMethod := http.MethodPut
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/deposit/provide-info"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.tranId == nil {
		return nil, common.ReportError("tranId is required and must be specified")
	}
	if r.questionnaire == nil {
		return nil, common.ReportError("questionnaire is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "tranId", r.tranId, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "questionnaire", r.questionnaire, "form", "")

	resp, err := SendRequest[models.SubmitDepositQuestionnaireTravelRuleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiVaspListRequest struct {
	ctx        context.Context
	ApiService *TravelRuleAPIService
	recvWindow *int64
}

func (r ApiVaspListRequest) RecvWindow(recvWindow int64) ApiVaspListRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVaspListRequest) Execute() (*common.RestApiResponse[models.VaspListResponse], error) {
	return r.ApiService.VaspListExecute(r)
}

/*
VaspList VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)
Get /sapi/v1/localentity/vasp

https://developers.binance.com/docs/wallet/travel-rule/onboarded-vasp-list

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiVaspListRequest
*/
func (a *TravelRuleAPIService) VaspList(ctx context.Context) ApiVaspListRequest {
	return ApiVaspListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VaspListResponse
func (a *TravelRuleAPIService) VaspListExecute(r ApiVaspListRequest) (*common.RestApiResponse[models.VaspListResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/vasp"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VaspListResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawHistoryV1Request struct {
	ctx              context.Context
	ApiService       *TravelRuleAPIService
	trId             *string
	txId             *string
	withdrawOrderId  *string
	network          *string
	coin             *string
	travelRuleStatus *int64
	offset           *int64
	limit            *int64
	startTime        *int64
	endTime          *int64
	recvWindow       *int64
}

// Comma(,) separated list of travel rule record Ids.
func (r ApiWithdrawHistoryV1Request) TrId(trId string) ApiWithdrawHistoryV1Request {
	r.trId = &trId
	return r
}

func (r ApiWithdrawHistoryV1Request) TxId(txId string) ApiWithdrawHistoryV1Request {
	r.txId = &txId
	return r
}

// client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query.
func (r ApiWithdrawHistoryV1Request) WithdrawOrderId(withdrawOrderId string) ApiWithdrawHistoryV1Request {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

func (r ApiWithdrawHistoryV1Request) Network(network string) ApiWithdrawHistoryV1Request {
	r.network = &network
	return r
}

func (r ApiWithdrawHistoryV1Request) Coin(coin string) ApiWithdrawHistoryV1Request {
	r.coin = &coin
	return r
}

// 0:Completed,1:Pending,2:Failed
func (r ApiWithdrawHistoryV1Request) TravelRuleStatus(travelRuleStatus int64) ApiWithdrawHistoryV1Request {
	r.travelRuleStatus = &travelRuleStatus
	return r
}

// Default: 0
func (r ApiWithdrawHistoryV1Request) Offset(offset int64) ApiWithdrawHistoryV1Request {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiWithdrawHistoryV1Request) Limit(limit int64) ApiWithdrawHistoryV1Request {
	r.limit = &limit
	return r
}

func (r ApiWithdrawHistoryV1Request) StartTime(startTime int64) ApiWithdrawHistoryV1Request {
	r.startTime = &startTime
	return r
}

func (r ApiWithdrawHistoryV1Request) EndTime(endTime int64) ApiWithdrawHistoryV1Request {
	r.endTime = &endTime
	return r
}

func (r ApiWithdrawHistoryV1Request) RecvWindow(recvWindow int64) ApiWithdrawHistoryV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawHistoryV1Request) Execute() (*common.RestApiResponse[models.WithdrawHistoryV1Response], error) {
	return r.ApiService.WithdrawHistoryV1Execute(r)
}

/*
WithdrawHistoryV1 Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)
Get /sapi/v1/localentity/withdraw/history

https://developers.binance.com/docs/wallet/travel-rule/Withdraw-History

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param trId -  Comma(,) separated list of travel rule record Ids.
@param txId -
@param withdrawOrderId -  client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query.
@param network -
@param coin -
@param travelRuleStatus -  0:Completed,1:Pending,2:Failed
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@param startTime -
@param endTime -
@param recvWindow -
@return ApiWithdrawHistoryV1Request
*/
func (a *TravelRuleAPIService) WithdrawHistoryV1(ctx context.Context) ApiWithdrawHistoryV1Request {
	return ApiWithdrawHistoryV1Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawHistoryV1Response
func (a *TravelRuleAPIService) WithdrawHistoryV1Execute(r ApiWithdrawHistoryV1Request) (*common.RestApiResponse[models.WithdrawHistoryV1Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/withdraw/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.trId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trId", r.trId, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.travelRuleStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "travelRuleStatus", r.travelRuleStatus, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WithdrawHistoryV1Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawHistoryV2Request struct {
	ctx              context.Context
	ApiService       *TravelRuleAPIService
	trId             *string
	txId             *string
	withdrawOrderId  *string
	network          *string
	coin             *string
	travelRuleStatus *int64
	offset           *int64
	limit            *int64
	startTime        *int64
	endTime          *int64
	recvWindow       *int64
}

// Comma(,) separated list of travel rule record Ids.
func (r ApiWithdrawHistoryV2Request) TrId(trId string) ApiWithdrawHistoryV2Request {
	r.trId = &trId
	return r
}

func (r ApiWithdrawHistoryV2Request) TxId(txId string) ApiWithdrawHistoryV2Request {
	r.txId = &txId
	return r
}

// client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query.
func (r ApiWithdrawHistoryV2Request) WithdrawOrderId(withdrawOrderId string) ApiWithdrawHistoryV2Request {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

func (r ApiWithdrawHistoryV2Request) Network(network string) ApiWithdrawHistoryV2Request {
	r.network = &network
	return r
}

func (r ApiWithdrawHistoryV2Request) Coin(coin string) ApiWithdrawHistoryV2Request {
	r.coin = &coin
	return r
}

// 0:Completed,1:Pending,2:Failed
func (r ApiWithdrawHistoryV2Request) TravelRuleStatus(travelRuleStatus int64) ApiWithdrawHistoryV2Request {
	r.travelRuleStatus = &travelRuleStatus
	return r
}

// Default: 0
func (r ApiWithdrawHistoryV2Request) Offset(offset int64) ApiWithdrawHistoryV2Request {
	r.offset = &offset
	return r
}

// min 7, max 30, default 7
func (r ApiWithdrawHistoryV2Request) Limit(limit int64) ApiWithdrawHistoryV2Request {
	r.limit = &limit
	return r
}

func (r ApiWithdrawHistoryV2Request) StartTime(startTime int64) ApiWithdrawHistoryV2Request {
	r.startTime = &startTime
	return r
}

func (r ApiWithdrawHistoryV2Request) EndTime(endTime int64) ApiWithdrawHistoryV2Request {
	r.endTime = &endTime
	return r
}

func (r ApiWithdrawHistoryV2Request) RecvWindow(recvWindow int64) ApiWithdrawHistoryV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawHistoryV2Request) Execute() (*common.RestApiResponse[models.WithdrawHistoryV2Response], error) {
	return r.ApiService.WithdrawHistoryV2Execute(r)
}

/*
WithdrawHistoryV2 Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)
Get /sapi/v2/localentity/withdraw/history

https://developers.binance.com/docs/wallet/travel-rule/Withdraw-History-V2

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param trId -  Comma(,) separated list of travel rule record Ids.
@param txId -
@param withdrawOrderId -  client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query.
@param network -
@param coin -
@param travelRuleStatus -  0:Completed,1:Pending,2:Failed
@param offset -  Default: 0
@param limit -  min 7, max 30, default 7
@param startTime -
@param endTime -
@param recvWindow -
@return ApiWithdrawHistoryV2Request
*/
func (a *TravelRuleAPIService) WithdrawHistoryV2(ctx context.Context) ApiWithdrawHistoryV2Request {
	return ApiWithdrawHistoryV2Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawHistoryV2Response
func (a *TravelRuleAPIService) WithdrawHistoryV2Execute(r ApiWithdrawHistoryV2Request) (*common.RestApiResponse[models.WithdrawHistoryV2Response], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v2/localentity/withdraw/history"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.trId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "trId", r.trId, "form", "")
	}
	if r.txId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "txId", r.txId, "form", "")
	}
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	if r.coin != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	}
	if r.travelRuleStatus != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "travelRuleStatus", r.travelRuleStatus, "form", "")
	}
	if r.offset != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.WithdrawHistoryV2Response](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiWithdrawTravelRuleRequest struct {
	ctx                context.Context
	ApiService         *TravelRuleAPIService
	coin               *string
	address            *string
	amount             *float32
	questionnaire      *string
	withdrawOrderId    *string
	network            *string
	addressTag         *string
	transactionFeeFlag *bool
	name               *string
	walletType         *int64
	recvWindow         *int64
}

func (r ApiWithdrawTravelRuleRequest) Coin(coin string) ApiWithdrawTravelRuleRequest {
	r.coin = &coin
	return r
}

func (r ApiWithdrawTravelRuleRequest) Address(address string) ApiWithdrawTravelRuleRequest {
	r.address = &address
	return r
}

func (r ApiWithdrawTravelRuleRequest) Amount(amount float32) ApiWithdrawTravelRuleRequest {
	r.amount = &amount
	return r
}

// JSON format questionnaire answers.
func (r ApiWithdrawTravelRuleRequest) Questionnaire(questionnaire string) ApiWithdrawTravelRuleRequest {
	r.questionnaire = &questionnaire
	return r
}

// client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query.
func (r ApiWithdrawTravelRuleRequest) WithdrawOrderId(withdrawOrderId string) ApiWithdrawTravelRuleRequest {
	r.withdrawOrderId = &withdrawOrderId
	return r
}

func (r ApiWithdrawTravelRuleRequest) Network(network string) ApiWithdrawTravelRuleRequest {
	r.network = &network
	return r
}

// Secondary address identifier for coins like XRP,XMR etc.
func (r ApiWithdrawTravelRuleRequest) AddressTag(addressTag string) ApiWithdrawTravelRuleRequest {
	r.addressTag = &addressTag
	return r
}

// When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;.
func (r ApiWithdrawTravelRuleRequest) TransactionFeeFlag(transactionFeeFlag bool) ApiWithdrawTravelRuleRequest {
	r.transactionFeeFlag = &transactionFeeFlag
	return r
}

// Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60;
func (r ApiWithdrawTravelRuleRequest) Name(name string) ApiWithdrawTravelRuleRequest {
	r.name = &name
	return r
}

// The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit
func (r ApiWithdrawTravelRuleRequest) WalletType(walletType int64) ApiWithdrawTravelRuleRequest {
	r.walletType = &walletType
	return r
}

func (r ApiWithdrawTravelRuleRequest) RecvWindow(recvWindow int64) ApiWithdrawTravelRuleRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiWithdrawTravelRuleRequest) Execute() (*common.RestApiResponse[models.WithdrawTravelRuleResponse], error) {
	return r.ApiService.WithdrawTravelRuleExecute(r)
}

/*
WithdrawTravelRule Withdraw (for local entities that require travel rule) (USER_DATA)
Post /sapi/v1/localentity/withdraw/apply

https://developers.binance.com/docs/wallet/travel-rule/Withdraw

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param coin -
@param address -
@param amount -
@param questionnaire -  JSON format questionnaire answers.
@param withdrawOrderId -  client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query.
@param network -
@param addressTag -  Secondary address identifier for coins like XRP,XMR etc.
@param transactionFeeFlag -  When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`.
@param name -  Description of the address. Address book cap is 200, space in name should be encoded into `%20`
@param walletType -  The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit
@param recvWindow -
@return ApiWithdrawTravelRuleRequest
*/
func (a *TravelRuleAPIService) WithdrawTravelRule(ctx context.Context) ApiWithdrawTravelRuleRequest {
	return ApiWithdrawTravelRuleRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WithdrawTravelRuleResponse
func (a *TravelRuleAPIService) WithdrawTravelRuleExecute(r ApiWithdrawTravelRuleRequest) (*common.RestApiResponse[models.WithdrawTravelRuleResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/localentity/withdraw/apply"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.coin == nil {
		return nil, common.ReportError("coin is required and must be specified")
	}
	if r.address == nil {
		return nil, common.ReportError("address is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}
	if r.questionnaire == nil {
		return nil, common.ReportError("questionnaire is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "coin", r.coin, "form", "")
	if r.withdrawOrderId != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "withdrawOrderId", r.withdrawOrderId, "form", "")
	}
	if r.network != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "network", r.network, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "form", "")
	if r.addressTag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "addressTag", r.addressTag, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.transactionFeeFlag != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "transactionFeeFlag", r.transactionFeeFlag, "form", "")
	}
	if r.name != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.walletType != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "walletType", r.walletType, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "questionnaire", r.questionnaire, "form", "")

	resp, err := SendRequest[models.WithdrawTravelRuleResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
