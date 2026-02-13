/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package binancegiftcardrestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/giftcard/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService Service

type ApiCreateADualTokenGiftCardRequest struct {
	ctx             context.Context
	ApiService      *MarketDataAPIService
	baseToken       *string
	faceToken       *string
	baseTokenAmount *float32
	recvWindow      *int64
}

// The token you want to pay, example: BUSD
func (r ApiCreateADualTokenGiftCardRequest) BaseToken(baseToken string) ApiCreateADualTokenGiftCardRequest {
	r.baseToken = &baseToken
	return r
}

// The token you want to buy, example: BNB. If faceToken &#x3D; baseToken, it&#39;s the same as createCode endpoint.
func (r ApiCreateADualTokenGiftCardRequest) FaceToken(faceToken string) ApiCreateADualTokenGiftCardRequest {
	r.faceToken = &faceToken
	return r
}

// The base token asset quantity, example : 1.002
func (r ApiCreateADualTokenGiftCardRequest) BaseTokenAmount(baseTokenAmount float32) ApiCreateADualTokenGiftCardRequest {
	r.baseTokenAmount = &baseTokenAmount
	return r
}

func (r ApiCreateADualTokenGiftCardRequest) RecvWindow(recvWindow int64) ApiCreateADualTokenGiftCardRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateADualTokenGiftCardRequest) Execute() (*common.RestApiResponse[models.CreateADualTokenGiftCardResponse], error) {
	return r.ApiService.CreateADualTokenGiftCardExecute(r)
}

/*
CreateADualTokenGiftCard Create a dual-token gift card(fixed value, discount feature)(TRADE)
Post /sapi/v1/giftcard/buyCode

https://developers.binance.com/docs/gift_card/market-data/Create-a-dual-token-gift-card

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param baseToken -  The token you want to pay, example: BUSD
@param faceToken -  The token you want to buy, example: BNB. If faceToken = baseToken, it's the same as createCode endpoint.
@param baseTokenAmount -  The base token asset quantity, example : 1.002
@param recvWindow -
@return ApiCreateADualTokenGiftCardRequest
*/
func (a *MarketDataAPIService) CreateADualTokenGiftCard(ctx context.Context) ApiCreateADualTokenGiftCardRequest {
	return ApiCreateADualTokenGiftCardRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateADualTokenGiftCardResponse
func (a *MarketDataAPIService) CreateADualTokenGiftCardExecute(r ApiCreateADualTokenGiftCardRequest) (*common.RestApiResponse[models.CreateADualTokenGiftCardResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/buyCode"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.baseToken == nil {
		return nil, common.ReportError("baseToken is required and must be specified")
	}
	if r.faceToken == nil {
		return nil, common.ReportError("faceToken is required and must be specified")
	}
	if r.baseTokenAmount == nil {
		return nil, common.ReportError("baseTokenAmount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "baseToken", r.baseToken, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "faceToken", r.faceToken, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "baseTokenAmount", r.baseTokenAmount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateADualTokenGiftCardResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiCreateASingleTokenGiftCardRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	token      *string
	amount     *float32
	recvWindow *int64
}

// The token type contained in the Binance Gift Card
func (r ApiCreateASingleTokenGiftCardRequest) Token(token string) ApiCreateASingleTokenGiftCardRequest {
	r.token = &token
	return r
}

// The amount of the token contained in the Binance Gift Card
func (r ApiCreateASingleTokenGiftCardRequest) Amount(amount float32) ApiCreateASingleTokenGiftCardRequest {
	r.amount = &amount
	return r
}

func (r ApiCreateASingleTokenGiftCardRequest) RecvWindow(recvWindow int64) ApiCreateASingleTokenGiftCardRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCreateASingleTokenGiftCardRequest) Execute() (*common.RestApiResponse[models.CreateASingleTokenGiftCardResponse], error) {
	return r.ApiService.CreateASingleTokenGiftCardExecute(r)
}

/*
CreateASingleTokenGiftCard Create a single-token gift card (USER_DATA)
Post /sapi/v1/giftcard/createCode

https://developers.binance.com/docs/gift_card/market-data/Create-a-single-token-gift-card

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param token -  The token type contained in the Binance Gift Card
@param amount -  The amount of the token contained in the Binance Gift Card
@param recvWindow -
@return ApiCreateASingleTokenGiftCardRequest
*/
func (a *MarketDataAPIService) CreateASingleTokenGiftCard(ctx context.Context) ApiCreateASingleTokenGiftCardRequest {
	return ApiCreateASingleTokenGiftCardRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateASingleTokenGiftCardResponse
func (a *MarketDataAPIService) CreateASingleTokenGiftCardExecute(r ApiCreateASingleTokenGiftCardRequest) (*common.RestApiResponse[models.CreateASingleTokenGiftCardResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/createCode"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.token == nil {
		return nil, common.ReportError("token is required and must be specified")
	}
	if r.amount == nil {
		return nil, common.ReportError("amount is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.CreateASingleTokenGiftCardResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchRsaPublicKeyRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	recvWindow *int64
}

func (r ApiFetchRsaPublicKeyRequest) RecvWindow(recvWindow int64) ApiFetchRsaPublicKeyRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFetchRsaPublicKeyRequest) Execute() (*common.RestApiResponse[models.FetchRsaPublicKeyResponse], error) {
	return r.ApiService.FetchRsaPublicKeyExecute(r)
}

/*
FetchRsaPublicKey Fetch RSA Public Key(USER_DATA)
Get /sapi/v1/giftcard/cryptography/rsa-public-key

https://developers.binance.com/docs/gift_card/market-data/Fetch-RSA-Public-Key

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param recvWindow -
@return ApiFetchRsaPublicKeyRequest
*/
func (a *MarketDataAPIService) FetchRsaPublicKey(ctx context.Context) ApiFetchRsaPublicKeyRequest {
	return ApiFetchRsaPublicKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchRsaPublicKeyResponse
func (a *MarketDataAPIService) FetchRsaPublicKeyExecute(r ApiFetchRsaPublicKeyRequest) (*common.RestApiResponse[models.FetchRsaPublicKeyResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/cryptography/rsa-public-key"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FetchRsaPublicKeyResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiFetchTokenLimitRequest struct {
	ctx        context.Context
	ApiService *MarketDataAPIService
	baseToken  *string
	recvWindow *int64
}

// The token you want to pay, example: BUSD
func (r ApiFetchTokenLimitRequest) BaseToken(baseToken string) ApiFetchTokenLimitRequest {
	r.baseToken = &baseToken
	return r
}

func (r ApiFetchTokenLimitRequest) RecvWindow(recvWindow int64) ApiFetchTokenLimitRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFetchTokenLimitRequest) Execute() (*common.RestApiResponse[models.FetchTokenLimitResponse], error) {
	return r.ApiService.FetchTokenLimitExecute(r)
}

/*
FetchTokenLimit Fetch Token Limit(USER_DATA)
Get /sapi/v1/giftcard/buyCode/token-limit

https://developers.binance.com/docs/gift_card/market-data/Fetch-Token-Limit

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param baseToken -  The token you want to pay, example: BUSD
@param recvWindow -
@return ApiFetchTokenLimitRequest
*/
func (a *MarketDataAPIService) FetchTokenLimit(ctx context.Context) ApiFetchTokenLimitRequest {
	return ApiFetchTokenLimitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FetchTokenLimitResponse
func (a *MarketDataAPIService) FetchTokenLimitExecute(r ApiFetchTokenLimitRequest) (*common.RestApiResponse[models.FetchTokenLimitResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/buyCode/token-limit"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.baseToken == nil {
		return nil, common.ReportError("baseToken is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "baseToken", r.baseToken, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.FetchTokenLimitResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiRedeemABinanceGiftCardRequest struct {
	ctx         context.Context
	ApiService  *MarketDataAPIService
	code        *string
	externalUid *string
	recvWindow  *int64
}

// Redemption code of Binance Gift Card to be redeemed, supports both Plaintext &amp; Encrypted code.
func (r ApiRedeemABinanceGiftCardRequest) Code(code string) ApiRedeemABinanceGiftCardRequest {
	r.code = &code
	return r
}

// Each external unique ID represents a unique user on the partner platform. The function helps you to identify the redemption behavior of different users, such as redemption frequency and amount. It also helps risk and limit control of a single account, such as daily limit on redemption volume, frequency, and incorrect number of entries. This will also prevent a single user account reach the partner&#39;s daily redemption limits. We strongly recommend you to use this feature and transfer us the User ID of your users if you have different users redeeming Binance Gift Cards on your platform. To protect user data privacy, you may choose to transfer the user id in any desired format (max. 400 characters).
func (r ApiRedeemABinanceGiftCardRequest) ExternalUid(externalUid string) ApiRedeemABinanceGiftCardRequest {
	r.externalUid = &externalUid
	return r
}

func (r ApiRedeemABinanceGiftCardRequest) RecvWindow(recvWindow int64) ApiRedeemABinanceGiftCardRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiRedeemABinanceGiftCardRequest) Execute() (*common.RestApiResponse[models.RedeemABinanceGiftCardResponse], error) {
	return r.ApiService.RedeemABinanceGiftCardExecute(r)
}

/*
RedeemABinanceGiftCard Redeem a Binance Gift Card(USER_DATA)
Post /sapi/v1/giftcard/redeemCode

https://developers.binance.com/docs/gift_card/market-data/Redeem-a-Binance-Gift-Card

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param code -  Redemption code of Binance Gift Card to be redeemed, supports both Plaintext & Encrypted code.
@param externalUid -  Each external unique ID represents a unique user on the partner platform. The function helps you to identify the redemption behavior of different users, such as redemption frequency and amount. It also helps risk and limit control of a single account, such as daily limit on redemption volume, frequency, and incorrect number of entries. This will also prevent a single user account reach the partner's daily redemption limits. We strongly recommend you to use this feature and transfer us the User ID of your users if you have different users redeeming Binance Gift Cards on your platform. To protect user data privacy, you may choose to transfer the user id in any desired format (max. 400 characters).
@param recvWindow -
@return ApiRedeemABinanceGiftCardRequest
*/
func (a *MarketDataAPIService) RedeemABinanceGiftCard(ctx context.Context) ApiRedeemABinanceGiftCardRequest {
	return ApiRedeemABinanceGiftCardRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedeemABinanceGiftCardResponse
func (a *MarketDataAPIService) RedeemABinanceGiftCardExecute(r ApiRedeemABinanceGiftCardRequest) (*common.RestApiResponse[models.RedeemABinanceGiftCardResponse], error) {
	localVarHTTPMethod := http.MethodPost
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/redeemCode"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.code == nil {
		return nil, common.ReportError("code is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
	if r.externalUid != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "externalUid", r.externalUid, "form", "")
	}
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.RedeemABinanceGiftCardResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

type ApiVerifyBinanceGiftCardByGiftCardNumberRequest struct {
	ctx         context.Context
	ApiService  *MarketDataAPIService
	referenceNo *string
	recvWindow  *int64
}

// Enter the Gift Card Number
func (r ApiVerifyBinanceGiftCardByGiftCardNumberRequest) ReferenceNo(referenceNo string) ApiVerifyBinanceGiftCardByGiftCardNumberRequest {
	r.referenceNo = &referenceNo
	return r
}

func (r ApiVerifyBinanceGiftCardByGiftCardNumberRequest) RecvWindow(recvWindow int64) ApiVerifyBinanceGiftCardByGiftCardNumberRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiVerifyBinanceGiftCardByGiftCardNumberRequest) Execute() (*common.RestApiResponse[models.VerifyBinanceGiftCardByGiftCardNumberResponse], error) {
	return r.ApiService.VerifyBinanceGiftCardByGiftCardNumberExecute(r)
}

/*
VerifyBinanceGiftCardByGiftCardNumber Verify Binance Gift Card by Gift Card Number(USER_DATA)
Get /sapi/v1/giftcard/verify

https://developers.binance.com/docs/gift_card/market-data/Verify-Binance-Gift-Card-by-Gift-Card-Number

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param referenceNo -  Enter the Gift Card Number
@param recvWindow -
@return ApiVerifyBinanceGiftCardByGiftCardNumberRequest
*/
func (a *MarketDataAPIService) VerifyBinanceGiftCardByGiftCardNumber(ctx context.Context) ApiVerifyBinanceGiftCardByGiftCardNumberRequest {
	return ApiVerifyBinanceGiftCardByGiftCardNumberRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VerifyBinanceGiftCardByGiftCardNumberResponse
func (a *MarketDataAPIService) VerifyBinanceGiftCardByGiftCardNumberExecute(r ApiVerifyBinanceGiftCardByGiftCardNumberRequest) (*common.RestApiResponse[models.VerifyBinanceGiftCardByGiftCardNumberResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/sapi/v1/giftcard/verify"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.referenceNo == nil {
		return nil, common.ReportError("referenceNo is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "referenceNo", r.referenceNo, "form", "")
	if r.recvWindow != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}

	resp, err := SendRequest[models.VerifyBinanceGiftCardByGiftCardNumberResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg, true)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
