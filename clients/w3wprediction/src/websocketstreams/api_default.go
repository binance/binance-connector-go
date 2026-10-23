/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 kinds of real-time events over the Binance SApi WebSocket channel, covering market and limit trading, reward claims, fund transfers and market resolution, plus order book (bid/ask) quotes per market. Trading events go only to the user they belong to; order book quotes go to everyone subscribed. A message may occasionally be missed, order is not guaranteed, and nothing is stored while you are offline — after reconnecting, fetch the latest state from the REST API and drop duplicates using `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package binancew3wpredictionwebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiBatchClaimPartiallySucceededRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiBatchClaimPartiallySucceededRequest) Execute() (*common.StreamHandler[models.BatchClaimPartiallySucceededResponse], error) {
	return r.ApiService.BatchClaimPartiallySucceededExecute(r)
}

/*
BatchClaimPartiallySucceeded Batch claim partially succeeded
/web3_prediction_pm_claim_partial_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#batch-claim-partially-succeeded

@return ApiBatchClaimPartiallySucceededRequest
*/
func (a *DefaultAPIService) BatchClaimPartiallySucceeded() ApiBatchClaimPartiallySucceededRequest {
	return ApiBatchClaimPartiallySucceededRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BatchClaimPartiallySucceededResponse
func (a *DefaultAPIService) BatchClaimPartiallySucceededExecute(r ApiBatchClaimPartiallySucceededRequest) (*common.StreamHandler[models.BatchClaimPartiallySucceededResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_claim_partial_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.BatchClaimPartiallySucceededResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiFundTransferFailedRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiFundTransferFailedRequest) Execute() (*common.StreamHandler[models.FundTransferFailedResponse], error) {
	return r.ApiService.FundTransferFailedExecute(r)
}

/*
FundTransferFailed Fund transfer failed
/web3_prediction_pm_transfer_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#fund-transfer-failed

@return ApiFundTransferFailedRequest
*/
func (a *DefaultAPIService) FundTransferFailed() ApiFundTransferFailedRequest {
	return ApiFundTransferFailedRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return FundTransferFailedResponse
func (a *DefaultAPIService) FundTransferFailedExecute(r ApiFundTransferFailedRequest) (*common.StreamHandler[models.FundTransferFailedResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_transfer_fail"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.FundTransferFailedResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiFundTransferSucceededRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiFundTransferSucceededRequest) Execute() (*common.StreamHandler[models.FundTransferSucceededResponse], error) {
	return r.ApiService.FundTransferSucceededExecute(r)
}

/*
FundTransferSucceeded Fund transfer succeeded
/web3_prediction_pm_transfer_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#fund-transfer-succeeded

@return ApiFundTransferSucceededRequest
*/
func (a *DefaultAPIService) FundTransferSucceeded() ApiFundTransferSucceededRequest {
	return ApiFundTransferSucceededRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return FundTransferSucceededResponse
func (a *DefaultAPIService) FundTransferSucceededExecute(r ApiFundTransferSucceededRequest) (*common.StreamHandler[models.FundTransferSucceededResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_transfer_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.FundTransferSucceededResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiLimitOrderFullyFilledRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiLimitOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.LimitOrderFullyFilledResponse], error) {
	return r.ApiService.LimitOrderFullyFilledExecute(r)
}

/*
LimitOrderFullyFilled Limit order fully filled
/web3_prediction_pm_limit_order_filled

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-fully-filled

@return ApiLimitOrderFullyFilledRequest
*/
func (a *DefaultAPIService) LimitOrderFullyFilled() ApiLimitOrderFullyFilledRequest {
	return ApiLimitOrderFullyFilledRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LimitOrderFullyFilledResponse
func (a *DefaultAPIService) LimitOrderFullyFilledExecute(r ApiLimitOrderFullyFilledRequest) (*common.StreamHandler[models.LimitOrderFullyFilledResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_limit_order_filled"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.LimitOrderFullyFilledResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiLimitOrderPartiallyFilledRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiLimitOrderPartiallyFilledRequest) Execute() (*common.StreamHandler[models.LimitOrderPartiallyFilledResponse], error) {
	return r.ApiService.LimitOrderPartiallyFilledExecute(r)
}

/*
LimitOrderPartiallyFilled Limit order partially filled
/web3_prediction_pm_limit_order_partial_fill

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-partially-filled

@return ApiLimitOrderPartiallyFilledRequest
*/
func (a *DefaultAPIService) LimitOrderPartiallyFilled() ApiLimitOrderPartiallyFilledRequest {
	return ApiLimitOrderPartiallyFilledRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LimitOrderPartiallyFilledResponse
func (a *DefaultAPIService) LimitOrderPartiallyFilledExecute(r ApiLimitOrderPartiallyFilledRequest) (*common.StreamHandler[models.LimitOrderPartiallyFilledResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_limit_order_partial_fill"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.LimitOrderPartiallyFilledResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiLimitOrderSubmissionFailedRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiLimitOrderSubmissionFailedRequest) Execute() (*common.StreamHandler[models.LimitOrderSubmissionFailedResponse], error) {
	return r.ApiService.LimitOrderSubmissionFailedExecute(r)
}

/*
LimitOrderSubmissionFailed Limit order submission failed
/web3_prediction_pm_limit_submit_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-submission-failed

@return ApiLimitOrderSubmissionFailedRequest
*/
func (a *DefaultAPIService) LimitOrderSubmissionFailed() ApiLimitOrderSubmissionFailedRequest {
	return ApiLimitOrderSubmissionFailedRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LimitOrderSubmissionFailedResponse
func (a *DefaultAPIService) LimitOrderSubmissionFailedExecute(r ApiLimitOrderSubmissionFailedRequest) (*common.StreamHandler[models.LimitOrderSubmissionFailedResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_limit_submit_fail"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.LimitOrderSubmissionFailedResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiLimitOrderSubmittedSuccessfullyRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiLimitOrderSubmittedSuccessfullyRequest) Execute() (*common.StreamHandler[models.LimitOrderSubmittedSuccessfullyResponse], error) {
	return r.ApiService.LimitOrderSubmittedSuccessfullyExecute(r)
}

/*
LimitOrderSubmittedSuccessfully Limit order submitted successfully
/web3_prediction_pm_limit_submit_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-submitted-successfully

@return ApiLimitOrderSubmittedSuccessfullyRequest
*/
func (a *DefaultAPIService) LimitOrderSubmittedSuccessfully() ApiLimitOrderSubmittedSuccessfullyRequest {
	return ApiLimitOrderSubmittedSuccessfullyRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LimitOrderSubmittedSuccessfullyResponse
func (a *DefaultAPIService) LimitOrderSubmittedSuccessfullyExecute(r ApiLimitOrderSubmittedSuccessfullyRequest) (*common.StreamHandler[models.LimitOrderSubmittedSuccessfullyResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_limit_submit_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.LimitOrderSubmittedSuccessfullyResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarketBuyOrderFailedRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiMarketBuyOrderFailedRequest) Execute() (*common.StreamHandler[models.MarketBuyOrderFailedResponse], error) {
	return r.ApiService.MarketBuyOrderFailedExecute(r)
}

/*
MarketBuyOrderFailed Market buy order failed / expired
/web3_prediction_pm_market_buy_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-buy-order-failed

@return ApiMarketBuyOrderFailedRequest
*/
func (a *DefaultAPIService) MarketBuyOrderFailed() ApiMarketBuyOrderFailedRequest {
	return ApiMarketBuyOrderFailedRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketBuyOrderFailedResponse
func (a *DefaultAPIService) MarketBuyOrderFailedExecute(r ApiMarketBuyOrderFailedRequest) (*common.StreamHandler[models.MarketBuyOrderFailedResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_market_buy_fail"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.MarketBuyOrderFailedResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarketBuyOrderFullyFilledRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiMarketBuyOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.MarketBuyOrderFullyFilledResponse], error) {
	return r.ApiService.MarketBuyOrderFullyFilledExecute(r)
}

/*
MarketBuyOrderFullyFilled Market buy order fully filled
/web3_prediction_pm_market_buy_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-buy-order-fully-filled

@return ApiMarketBuyOrderFullyFilledRequest
*/
func (a *DefaultAPIService) MarketBuyOrderFullyFilled() ApiMarketBuyOrderFullyFilledRequest {
	return ApiMarketBuyOrderFullyFilledRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketBuyOrderFullyFilledResponse
func (a *DefaultAPIService) MarketBuyOrderFullyFilledExecute(r ApiMarketBuyOrderFullyFilledRequest) (*common.StreamHandler[models.MarketBuyOrderFullyFilledResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_market_buy_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.MarketBuyOrderFullyFilledResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarketResolutionNotificationRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiMarketResolutionNotificationRequest) Execute() (*common.StreamHandler[models.MarketResolutionNotificationResponse], error) {
	return r.ApiService.MarketResolutionNotificationExecute(r)
}

/*
MarketResolutionNotification Market resolution notification
/web3_prediction_pm_market_close

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-resolution-notification

@return ApiMarketResolutionNotificationRequest
*/
func (a *DefaultAPIService) MarketResolutionNotification() ApiMarketResolutionNotificationRequest {
	return ApiMarketResolutionNotificationRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketResolutionNotificationResponse
func (a *DefaultAPIService) MarketResolutionNotificationExecute(r ApiMarketResolutionNotificationRequest) (*common.StreamHandler[models.MarketResolutionNotificationResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_market_close"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.MarketResolutionNotificationResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarketSellOrderFailedExpiredRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiMarketSellOrderFailedExpiredRequest) Execute() (*common.StreamHandler[models.MarketSellOrderFailedExpiredResponse], error) {
	return r.ApiService.MarketSellOrderFailedExpiredExecute(r)
}

/*
MarketSellOrderFailedExpired Market sell order failed / expired
/web3_prediction_pm_market_sell_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-sell-order-failed-expired

@return ApiMarketSellOrderFailedExpiredRequest
*/
func (a *DefaultAPIService) MarketSellOrderFailedExpired() ApiMarketSellOrderFailedExpiredRequest {
	return ApiMarketSellOrderFailedExpiredRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketSellOrderFailedExpiredResponse
func (a *DefaultAPIService) MarketSellOrderFailedExpiredExecute(r ApiMarketSellOrderFailedExpiredRequest) (*common.StreamHandler[models.MarketSellOrderFailedExpiredResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_market_sell_fail"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.MarketSellOrderFailedExpiredResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarketSellOrderFullyFilledRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiMarketSellOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.MarketSellOrderFullyFilledResponse], error) {
	return r.ApiService.MarketSellOrderFullyFilledExecute(r)
}

/*
MarketSellOrderFullyFilled Market sell order fully filled
/web3_prediction_pm_market_sell_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-sell-order-fully-filled

@return ApiMarketSellOrderFullyFilledRequest
*/
func (a *DefaultAPIService) MarketSellOrderFullyFilled() ApiMarketSellOrderFullyFilledRequest {
	return ApiMarketSellOrderFullyFilledRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketSellOrderFullyFilledResponse
func (a *DefaultAPIService) MarketSellOrderFullyFilledExecute(r ApiMarketSellOrderFullyFilledRequest) (*common.StreamHandler[models.MarketSellOrderFullyFilledResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_market_sell_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.MarketSellOrderFullyFilledResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPredictionOrderBookByMarketIdRequest struct {
	ApiService *DefaultAPIService
	marketId   *string
}

// Market ID to subscribe to — the &#x60;marketId&#x60; inside &#x60;markets[]&#x60; returned by &#x60;market/list&#x60; or &#x60;market/detail&#x60;. It is **not** &#x60;marketTopicId&#x60; and **not** &#x60;externalId&#x60;. A wrong id is never rejected: subscribing with &#x60;marketTopicId&#x60;, &#x60;externalId&#x60;, or with a closed or resolved market returns the same successful &#x60;COMMAND&#x60; reply as a correct id and then pushes nothing at all, so a stream that stays silent is far more likely a wrong id than a broken connection.
func (r ApiPredictionOrderBookByMarketIdRequest) MarketId(marketId string) ApiPredictionOrderBookByMarketIdRequest {
	r.marketId = &marketId
	return r
}

func (r ApiPredictionOrderBookByMarketIdRequest) Execute() (*common.StreamHandler[models.PredictionOrderBookByMarketIdResponse], error) {
	return r.ApiService.PredictionOrderBookByMarketIdExecute(r)
}

/*
PredictionOrderBookByMarketId Prediction Order Book by marketId
/web3_prediction_orderbook_<marketId>

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#prediction-order-book-by-market-id

@param marketId Market ID to subscribe to — the `marketId` inside `markets[]` returned by `market/list` or `market/detail`. It is **not** `marketTopicId` and **not** `externalId`. A wrong id is never rejected: subscribing with `marketTopicId`, `externalId`, or with a closed or resolved market returns the same successful `COMMAND` reply as a correct id and then pushes nothing at all, so a stream that stays silent is far more likely a wrong id than a broken connection.
@return ApiPredictionOrderBookByMarketIdRequest
*/
func (a *DefaultAPIService) PredictionOrderBookByMarketId() ApiPredictionOrderBookByMarketIdRequest {
	return ApiPredictionOrderBookByMarketIdRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PredictionOrderBookByMarketIdResponse
func (a *DefaultAPIService) PredictionOrderBookByMarketIdExecute(r ApiPredictionOrderBookByMarketIdRequest) (*common.StreamHandler[models.PredictionOrderBookByMarketIdResponse], error) {
	if r.marketId == nil {
		return nil, common.ReportError("marketId is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_orderbook_<marketId>"[1:],
		map[string]string{
			"marketId": func() string {
				if r.marketId == nil {
					return ""
				}
				return *r.marketId
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.PredictionOrderBookByMarketIdResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRewardClaimFailedRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiRewardClaimFailedRequest) Execute() (*common.StreamHandler[models.RewardClaimFailedResponse], error) {
	return r.ApiService.RewardClaimFailedExecute(r)
}

/*
RewardClaimFailed Reward claim failed
/web3_prediction_pm_claim_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#reward-claim-failed

@return ApiRewardClaimFailedRequest
*/
func (a *DefaultAPIService) RewardClaimFailed() ApiRewardClaimFailedRequest {
	return ApiRewardClaimFailedRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RewardClaimFailedResponse
func (a *DefaultAPIService) RewardClaimFailedExecute(r ApiRewardClaimFailedRequest) (*common.StreamHandler[models.RewardClaimFailedResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_claim_fail"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.RewardClaimFailedResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRewardClaimSucceededRequest struct {
	ApiService *DefaultAPIService
}

func (r ApiRewardClaimSucceededRequest) Execute() (*common.StreamHandler[models.RewardClaimSucceededResponse], error) {
	return r.ApiService.RewardClaimSucceededExecute(r)
}

/*
RewardClaimSucceeded Reward claim succeeded
/web3_prediction_pm_claim_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#reward-claim-succeeded

@return ApiRewardClaimSucceededRequest
*/
func (a *DefaultAPIService) RewardClaimSucceeded() ApiRewardClaimSucceededRequest {
	return ApiRewardClaimSucceededRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RewardClaimSucceededResponse
func (a *DefaultAPIService) RewardClaimSucceededExecute(r ApiRewardClaimSucceededRequest) (*common.StreamHandler[models.RewardClaimSucceededResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_pm_claim_success"[1:],
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.RewardClaimSucceededResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
