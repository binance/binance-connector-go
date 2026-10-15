/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiBatchClaimPartiallySucceededRequest) Id(id string) ApiBatchClaimPartiallySucceededRequest {
	r.id = &id
	return r
}

func (r ApiBatchClaimPartiallySucceededRequest) Execute() (*common.StreamHandler[models.BatchClaimPartiallySucceededResponse], error) {
	return r.ApiService.BatchClaimPartiallySucceededExecute(r)
}

/*
BatchClaimPartiallySucceeded Batch claim partially succeeded
/web3_prediction_pm_claim_partial_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#batch-claim-partially-succeeded

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiFundTransferFailedRequest) Id(id string) ApiFundTransferFailedRequest {
	r.id = &id
	return r
}

func (r ApiFundTransferFailedRequest) Execute() (*common.StreamHandler[models.FundTransferFailedResponse], error) {
	return r.ApiService.FundTransferFailedExecute(r)
}

/*
FundTransferFailed Fund transfer failed
/web3_prediction_pm_transfer_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#fund-transfer-failed

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiFundTransferSucceededRequest) Id(id string) ApiFundTransferSucceededRequest {
	r.id = &id
	return r
}

func (r ApiFundTransferSucceededRequest) Execute() (*common.StreamHandler[models.FundTransferSucceededResponse], error) {
	return r.ApiService.FundTransferSucceededExecute(r)
}

/*
FundTransferSucceeded Fund transfer succeeded
/web3_prediction_pm_transfer_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#fund-transfer-succeeded

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiLimitOrderFullyFilledRequest) Id(id string) ApiLimitOrderFullyFilledRequest {
	r.id = &id
	return r
}

func (r ApiLimitOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.LimitOrderFullyFilledResponse], error) {
	return r.ApiService.LimitOrderFullyFilledExecute(r)
}

/*
LimitOrderFullyFilled Limit order fully filled
/web3_prediction_pm_limit_order_filled

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-fully-filled

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiLimitOrderPartiallyFilledRequest) Id(id string) ApiLimitOrderPartiallyFilledRequest {
	r.id = &id
	return r
}

func (r ApiLimitOrderPartiallyFilledRequest) Execute() (*common.StreamHandler[models.LimitOrderPartiallyFilledResponse], error) {
	return r.ApiService.LimitOrderPartiallyFilledExecute(r)
}

/*
LimitOrderPartiallyFilled Limit order partially filled
/web3_prediction_pm_limit_order_partial_fill

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-partially-filled

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiLimitOrderSubmissionFailedRequest) Id(id string) ApiLimitOrderSubmissionFailedRequest {
	r.id = &id
	return r
}

func (r ApiLimitOrderSubmissionFailedRequest) Execute() (*common.StreamHandler[models.LimitOrderSubmissionFailedResponse], error) {
	return r.ApiService.LimitOrderSubmissionFailedExecute(r)
}

/*
LimitOrderSubmissionFailed Limit order submission failed
/web3_prediction_pm_limit_submit_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-submission-failed

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiLimitOrderSubmittedSuccessfullyRequest) Id(id string) ApiLimitOrderSubmittedSuccessfullyRequest {
	r.id = &id
	return r
}

func (r ApiLimitOrderSubmittedSuccessfullyRequest) Execute() (*common.StreamHandler[models.LimitOrderSubmittedSuccessfullyResponse], error) {
	return r.ApiService.LimitOrderSubmittedSuccessfullyExecute(r)
}

/*
LimitOrderSubmittedSuccessfully Limit order submitted successfully
/web3_prediction_pm_limit_submit_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#limit-order-submitted-successfully

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMarketBuyOrderFailedRequest) Id(id string) ApiMarketBuyOrderFailedRequest {
	r.id = &id
	return r
}

func (r ApiMarketBuyOrderFailedRequest) Execute() (*common.StreamHandler[models.MarketBuyOrderFailedResponse], error) {
	return r.ApiService.MarketBuyOrderFailedExecute(r)
}

/*
MarketBuyOrderFailed Market buy order failed / expired
/web3_prediction_pm_market_buy_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-buy-order-failed

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMarketBuyOrderFullyFilledRequest) Id(id string) ApiMarketBuyOrderFullyFilledRequest {
	r.id = &id
	return r
}

func (r ApiMarketBuyOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.MarketBuyOrderFullyFilledResponse], error) {
	return r.ApiService.MarketBuyOrderFullyFilledExecute(r)
}

/*
MarketBuyOrderFullyFilled Market buy order fully filled
/web3_prediction_pm_market_buy_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-buy-order-fully-filled

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMarketResolutionNotificationRequest) Id(id string) ApiMarketResolutionNotificationRequest {
	r.id = &id
	return r
}

func (r ApiMarketResolutionNotificationRequest) Execute() (*common.StreamHandler[models.MarketResolutionNotificationResponse], error) {
	return r.ApiService.MarketResolutionNotificationExecute(r)
}

/*
MarketResolutionNotification Market resolution notification
/web3_prediction_pm_market_close

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-resolution-notification

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMarketSellOrderFailedExpiredRequest) Id(id string) ApiMarketSellOrderFailedExpiredRequest {
	r.id = &id
	return r
}

func (r ApiMarketSellOrderFailedExpiredRequest) Execute() (*common.StreamHandler[models.MarketSellOrderFailedExpiredResponse], error) {
	return r.ApiService.MarketSellOrderFailedExpiredExecute(r)
}

/*
MarketSellOrderFailedExpired Market sell order failed / expired
/web3_prediction_pm_market_sell_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-sell-order-failed-expired

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMarketSellOrderFullyFilledRequest) Id(id string) ApiMarketSellOrderFullyFilledRequest {
	r.id = &id
	return r
}

func (r ApiMarketSellOrderFullyFilledRequest) Execute() (*common.StreamHandler[models.MarketSellOrderFullyFilledResponse], error) {
	return r.ApiService.MarketSellOrderFullyFilledExecute(r)
}

/*
MarketSellOrderFullyFilled Market sell order fully filled
/web3_prediction_pm_market_sell_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#market-sell-order-fully-filled

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.MarketSellOrderFullyFilledResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPredictionOrderBookAllRequest struct {
	ApiService *DefaultAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiPredictionOrderBookAllRequest) Id(id string) ApiPredictionOrderBookAllRequest {
	r.id = &id
	return r
}

func (r ApiPredictionOrderBookAllRequest) Execute() (*common.StreamHandler[models.PredictionOrderBookAllResponse], error) {
	return r.ApiService.PredictionOrderBookAllExecute(r)
}

/*
PredictionOrderBookAll Prediction Order Book All
/web3_prediction_orderbook_data

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#prediction-order-book-all

@param id Unique WebSocket request ID.
@return ApiPredictionOrderBookAllRequest
*/
func (a *DefaultAPIService) PredictionOrderBookAll() ApiPredictionOrderBookAllRequest {
	return ApiPredictionOrderBookAllRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PredictionOrderBookAllResponse
func (a *DefaultAPIService) PredictionOrderBookAllExecute(r ApiPredictionOrderBookAllRequest) (*common.StreamHandler[models.PredictionOrderBookAllResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/web3_prediction_orderbook_data"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PredictionOrderBookAllResponse](&common.StreamHandlerWrapper{
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
	id         *string
}

// Binance internal market ID to subscribe to.
func (r ApiPredictionOrderBookByMarketIdRequest) MarketId(marketId string) ApiPredictionOrderBookByMarketIdRequest {
	r.marketId = &marketId
	return r
}

// Unique WebSocket request ID.
func (r ApiPredictionOrderBookByMarketIdRequest) Id(id string) ApiPredictionOrderBookByMarketIdRequest {
	r.id = &id
	return r
}

func (r ApiPredictionOrderBookByMarketIdRequest) Execute() (*common.StreamHandler[models.PredictionOrderBookByMarketIdResponse], error) {
	return r.ApiService.PredictionOrderBookByMarketIdExecute(r)
}

/*
PredictionOrderBookByMarketId Prediction Order Book by marketId
/web3_prediction_orderbook_<marketId>

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#prediction-order-book-by-market-id

@param marketId Binance internal market ID to subscribe to.	@param id Unique WebSocket request ID.
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
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiRewardClaimFailedRequest) Id(id string) ApiRewardClaimFailedRequest {
	r.id = &id
	return r
}

func (r ApiRewardClaimFailedRequest) Execute() (*common.StreamHandler[models.RewardClaimFailedResponse], error) {
	return r.ApiService.RewardClaimFailedExecute(r)
}

/*
RewardClaimFailed Reward claim failed
/web3_prediction_pm_claim_fail

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#reward-claim-failed

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
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
	id         *string
}

// Unique WebSocket request ID.
func (r ApiRewardClaimSucceededRequest) Id(id string) ApiRewardClaimSucceededRequest {
	r.id = &id
	return r
}

func (r ApiRewardClaimSucceededRequest) Execute() (*common.StreamHandler[models.RewardClaimSucceededResponse], error) {
	return r.ApiService.RewardClaimSucceededExecute(r)
}

/*
RewardClaimSucceeded Reward claim succeeded
/web3_prediction_pm_claim_success

https://developers.binance.com/en/docs/catalog/web3-wallet-prediction-trading/api/ws-streams/~#reward-claim-succeeded

@param id Unique WebSocket request ID.
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
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.RewardClaimSucceededResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
