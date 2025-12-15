/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AuthAPIService AuthAPI Service
type AuthAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiSessionLogonRequest struct {
	ApiService *AuthAPIService
	id         *string
	recvWindow *float32
}

// Unique WebSocket request ID.
func (r ApiSessionLogonRequest) Id(id string) ApiSessionLogonRequest {
	r.id = &id
	return r
}

// The value cannot be greater than &#x60;60000&#x60;. &lt;br&gt; Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
func (r ApiSessionLogonRequest) RecvWindow(recvWindow float32) ApiSessionLogonRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSessionLogonRequest) Execute() (*common.ResponseOrRaw[models.SessionLogonResponse], error) {
	respChan, errChan, err := r.ApiService.SessionLogonExecute(r)
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

func (r ApiSessionLogonRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionLogonResponse], chan error, error) {
	return r.ApiService.SessionLogonExecute(r)
}

/*
SessionLogon WebSocket Log in with API key
/session.logon

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/authentication-requests#log-in-with-api-key-signed

@param id Unique WebSocket request ID.	@param recvWindow The value cannot be greater than `60000`. <br> Supports up to three decimal places of precision (e.g., 6000.346) so that microseconds may be specified.
@return ApiSessionLogonRequest
*/
func (a *AuthAPIService) SessionLogon() ApiSessionLogonRequest {
	return ApiSessionLogonRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionLogonResponse
func (a *AuthAPIService) SessionLogonExecute(r ApiSessionLogonRequest) (chan *common.ResponseOrRaw[models.SessionLogonResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/session.logon"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: true,
	}

	return SendMessage[models.SessionLogonResponse](a.Ws, localPayload, sendParams)
}

type ApiSessionLogoutRequest struct {
	ApiService *AuthAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiSessionLogoutRequest) Id(id string) ApiSessionLogoutRequest {
	r.id = &id
	return r
}

func (r ApiSessionLogoutRequest) Execute() (*common.ResponseOrRaw[models.SessionLogoutResponse], error) {
	respChan, errChan, err := r.ApiService.SessionLogoutExecute(r)
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

func (r ApiSessionLogoutRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionLogoutResponse], chan error, error) {
	return r.ApiService.SessionLogoutExecute(r)
}

/*
SessionLogout WebSocket Log out of the session
/session.logout

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/authentication-requests#log-out-of-the-session

@param id Unique WebSocket request ID.
@return ApiSessionLogoutRequest
*/
func (a *AuthAPIService) SessionLogout() ApiSessionLogoutRequest {
	return ApiSessionLogoutRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionLogoutResponse
func (a *AuthAPIService) SessionLogoutExecute(r ApiSessionLogoutRequest) (chan *common.ResponseOrRaw[models.SessionLogoutResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/session.logout"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SessionLogoutResponse](a.Ws, localPayload, sendParams)
}

type ApiSessionStatusRequest struct {
	ApiService *AuthAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiSessionStatusRequest) Id(id string) ApiSessionStatusRequest {
	r.id = &id
	return r
}

func (r ApiSessionStatusRequest) Execute() (*common.ResponseOrRaw[models.SessionStatusResponse], error) {
	respChan, errChan, err := r.ApiService.SessionStatusExecute(r)
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

func (r ApiSessionStatusRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SessionStatusResponse], chan error, error) {
	return r.ApiService.SessionStatusExecute(r)
}

/*
SessionStatus WebSocket Query session status
/session.status

https://developers.binance.com/docs/binance-spot-api-docs/websocket-api/authentication-requests#query-session-status

@param id Unique WebSocket request ID.
@return ApiSessionStatusRequest
*/
func (a *AuthAPIService) SessionStatus() ApiSessionStatusRequest {
	return ApiSessionStatusRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SessionStatusResponse
func (a *AuthAPIService) SessionStatusExecute(r ApiSessionStatusRequest) (chan *common.ResponseOrRaw[models.SessionStatusResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/session.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SessionStatusResponse](a.Ws, localPayload, sendParams)
}
