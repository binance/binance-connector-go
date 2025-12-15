/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package binancespotwebsocketapi

import (
	"strconv"

	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// UserDataStreamAPIService UserDataStreamAPI Service
type UserDataStreamAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiSessionSubscriptionsRequest struct {
	ApiService *UserDataStreamAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiSessionSubscriptionsRequest) Id(id string) ApiSessionSubscriptionsRequest {
	r.id = &id
	return r
}

func (r ApiSessionSubscriptionsRequest) Execute() (*common.ResponseOrRaw[models.SessionSubscriptionsResponse], error) {
	respChan, errChan, err := r.ApiService.SessionSubscriptionsExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiSessionSubscriptionsRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionSubscriptionsResponse], chan error, error) {
	return r.ApiService.SessionSubscriptionsExecute(r)
}

/*
SessionSubscriptions WebSocket Listing all subscriptions
/session.subscriptions

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/user-Data-Stream-requests#listing-all-subscriptions

@param id Unique WebSocket request ID.
@return ApiSessionSubscriptionsRequest
*/
func (a *UserDataStreamAPIService) SessionSubscriptions() ApiSessionSubscriptionsRequest {
	return ApiSessionSubscriptionsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionSubscriptionsResponse
func (a *UserDataStreamAPIService) SessionSubscriptionsExecute(r ApiSessionSubscriptionsRequest) (chan *common.ResponseOrRaw[models.SessionSubscriptionsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/session.subscriptions"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SessionSubscriptionsResponse](a.Ws, localPayload, sendParams)
}

type ApiUserDataStreamSubscribeRequest struct {
	ApiService *UserDataStreamAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiUserDataStreamSubscribeRequest) Id(id string) ApiUserDataStreamSubscribeRequest {
	r.id = &id
	return r
}

func (r ApiUserDataStreamSubscribeRequest) Execute() (*common.ResponseOrRaw[models.UserDataStreamSubscribeResponse], *common.StreamHandler[models.UserDataStreamEventsResponse], chan error, error) {
	return r.ApiService.UserDataStreamSubscribeExecute(r)
}

/*
UserDataStreamSubscribe WebSocket Subscribe to User Data Stream
/userDataStream.subscribe

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/user-Data-Stream-requests#subscribe-to-user-data-stream-user_stream

@param id Unique WebSocket request ID.
@return ApiUserDataStreamSubscribeRequest
*/
func (a *UserDataStreamAPIService) UserDataStreamSubscribe() ApiUserDataStreamSubscribeRequest {
	return ApiUserDataStreamSubscribeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return UserDataStreamSubscribeResponse
func (a *UserDataStreamAPIService) UserDataStreamSubscribeExecute(r ApiUserDataStreamSubscribeRequest) (*common.ResponseOrRaw[models.UserDataStreamSubscribeResponse], *common.StreamHandler[models.UserDataStreamEventsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/userDataStream.subscribe"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	responseChan, errorChan, err := SendMessage[models.UserDataStreamSubscribeResponse](a.Ws, localPayload, sendParams)

	if err != nil {
		return nil, nil, nil, err
	}

	select {
	case resp := <-responseChan:
		sid := resp.Typed.GetResult().SubscriptionId
		streamId := ""
		if sid != nil {
			streamId = strconv.FormatInt(*sid, 10)
		}
		ws := a.Ws
		stream, err := common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketAPI: ws,
		}, streamId, []string{common.GenerateUUID()})

		if err != nil {
			return nil, nil, nil, err
		}

		return resp, stream, nil, nil
	case err := <-errorChan:
		errChan := make(chan error, 1)
		errChan <- err
		close(errChan)
		return nil, nil, errChan, err
	}

}

type ApiUserDataStreamSubscribeSignatureRequest struct {
	ApiService *UserDataStreamAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiUserDataStreamSubscribeSignatureRequest) Id(id string) ApiUserDataStreamSubscribeSignatureRequest {
	r.id = &id
	return r
}

func (r ApiUserDataStreamSubscribeSignatureRequest) Execute() (*common.ResponseOrRaw[models.UserDataStreamSubscribeSignatureResponse], *common.StreamHandler[models.UserDataStreamEventsResponse], chan error, error) {
	return r.ApiService.UserDataStreamSubscribeSignatureExecute(r)
}

/*
UserDataStreamSubscribeSignature WebSocket Subscribe to User Data Stream through signature subscription
/userDataStream.subscribe.signature

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/user-Data-Stream-requests#subscribe-to-user-data-stream-through-signature-subscription-user_stream

@param id Unique WebSocket request ID.
@return ApiUserDataStreamSubscribeSignatureRequest
*/
func (a *UserDataStreamAPIService) UserDataStreamSubscribeSignature() ApiUserDataStreamSubscribeSignatureRequest {
	return ApiUserDataStreamSubscribeSignatureRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return UserDataStreamSubscribeSignatureResponse
func (a *UserDataStreamAPIService) UserDataStreamSubscribeSignatureExecute(r ApiUserDataStreamSubscribeSignatureRequest) (*common.ResponseOrRaw[models.UserDataStreamSubscribeSignatureResponse], *common.StreamHandler[models.UserDataStreamEventsResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/userDataStream.subscribe.signature"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	responseChan, errorChan, err := SendMessage[models.UserDataStreamSubscribeSignatureResponse](a.Ws, localPayload, sendParams)

	if err != nil {
		return nil, nil, nil, err
	}

	select {
	case resp := <-responseChan:
		sid := resp.Typed.GetResult().SubscriptionId
		streamId := ""
		if sid != nil {
			streamId = strconv.FormatInt(*sid, 10)
		}
		ws := a.Ws
		stream, err := common.CreateStreamHandler[models.UserDataStreamEventsResponse](&common.StreamHandlerWrapper{
			WebsocketAPI: ws,
		}, streamId, []string{common.GenerateUUID()})

		if err != nil {
			return nil, nil, nil, err
		}

		return resp, stream, nil, nil
	case err := <-errorChan:
		errChan := make(chan error, 1)
		errChan <- err
		close(errChan)
		return nil, nil, errChan, err
	}

}

type ApiUserDataStreamUnsubscribeRequest struct {
	ApiService     *UserDataStreamAPIService
	id             *string
	subscriptionId *int32
}

// Unique WebSocket request ID.
func (r ApiUserDataStreamUnsubscribeRequest) Id(id string) ApiUserDataStreamUnsubscribeRequest {
	r.id = &id
	return r
}

// When called with no parameter, this will close all subscriptions. &lt;br&gt;When called with the &#x60;subscriptionId&#x60; parameter, this will attempt to close the subscription with that subscription id, if it exists.
func (r ApiUserDataStreamUnsubscribeRequest) SubscriptionId(subscriptionId int32) ApiUserDataStreamUnsubscribeRequest {
	r.subscriptionId = &subscriptionId
	return r
}

func (r ApiUserDataStreamUnsubscribeRequest) Execute() (*common.ResponseOrRaw[models.UserDataStreamUnsubscribeResponse], error) {
	respChan, errChan, err := r.ApiService.UserDataStreamUnsubscribeExecute(r)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	}
}

func (r ApiUserDataStreamUnsubscribeRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.UserDataStreamUnsubscribeResponse], chan error, error) {
	return r.ApiService.UserDataStreamUnsubscribeExecute(r)
}

/*
UserDataStreamUnsubscribe WebSocket Unsubscribe from User Data Stream
/userDataStream.unsubscribe

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/user-Data-Stream-requests#unsubscribe-from-user-data-stream

@param id Unique WebSocket request ID.	@param subscriptionId When called with no parameter, this will close all subscriptions. <br>When called with the `subscriptionId` parameter, this will attempt to close the subscription with that subscription id, if it exists.
@return ApiUserDataStreamUnsubscribeRequest
*/
func (a *UserDataStreamAPIService) UserDataStreamUnsubscribe() ApiUserDataStreamUnsubscribeRequest {
	return ApiUserDataStreamUnsubscribeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return UserDataStreamUnsubscribeResponse
func (a *UserDataStreamAPIService) UserDataStreamUnsubscribeExecute(r ApiUserDataStreamUnsubscribeRequest) (chan *common.ResponseOrRaw[models.UserDataStreamUnsubscribeResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.subscriptionId != nil {
		localVarQueryParams["subscriptionId"] = *r.subscriptionId
	}

	localPayload := map[string]any{
		"method": "/userDataStream.unsubscribe"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.UserDataStreamUnsubscribeResponse](a.Ws, localPayload, sendParams)
}
