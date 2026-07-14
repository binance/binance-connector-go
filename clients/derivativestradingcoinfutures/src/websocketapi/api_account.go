/*
Futures (COIN-M) WebSocket API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package binancederivativestradingcoinfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// AccountAPIService AccountAPI Service
type AccountAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiAccountInformationRequest struct {
	ApiService *AccountAPIService
	id         *string
	recvWindow *int64
}

func (r ApiAccountInformationRequest) Id(id string) ApiAccountInformationRequest {
	r.id = &id
	return r
}

func (r ApiAccountInformationRequest) RecvWindow(recvWindow int64) ApiAccountInformationRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiAccountInformationRequest) Execute() (*common.ResponseOrRaw[models.AccountInformationResponse], error) {
	respChan, errChan, err := r.ApiService.AccountInformationExecute(r)
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

func (r ApiAccountInformationRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AccountInformationResponse], chan error, error) {
	return r.ApiService.AccountInformationExecute(r)
}

/*
AccountInformation Account Information (USER_DATA)
/account.status

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/account#account-information

@param id	@param recvWindow
@return ApiAccountInformationRequest
*/
func (a *AccountAPIService) AccountInformation() ApiAccountInformationRequest {
	return ApiAccountInformationRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AccountInformationResponse
func (a *AccountAPIService) AccountInformationExecute(r ApiAccountInformationRequest) (chan *common.ResponseOrRaw[models.AccountInformationResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.status"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AccountInformationResponse](a.Ws, localPayload, sendParams)
}

type ApiFuturesAccountBalanceRequest struct {
	ApiService *AccountAPIService
	id         *string
	recvWindow *int64
}

func (r ApiFuturesAccountBalanceRequest) Id(id string) ApiFuturesAccountBalanceRequest {
	r.id = &id
	return r
}

func (r ApiFuturesAccountBalanceRequest) RecvWindow(recvWindow int64) ApiFuturesAccountBalanceRequest {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiFuturesAccountBalanceRequest) Execute() (*common.ResponseOrRaw[models.FuturesAccountBalanceResponse], error) {
	respChan, errChan, err := r.ApiService.FuturesAccountBalanceExecute(r)
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

func (r ApiFuturesAccountBalanceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.FuturesAccountBalanceResponse], chan error, error) {
	return r.ApiService.FuturesAccountBalanceExecute(r)
}

/*
FuturesAccountBalance Futures Account Balance (USER_DATA)
/account.balance

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-api/account#futures-account-balance

@param id	@param recvWindow
@return ApiFuturesAccountBalanceRequest
*/
func (a *AccountAPIService) FuturesAccountBalance() ApiFuturesAccountBalanceRequest {
	return ApiFuturesAccountBalanceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return FuturesAccountBalanceResponse
func (a *AccountAPIService) FuturesAccountBalanceExecute(r ApiFuturesAccountBalanceRequest) (chan *common.ResponseOrRaw[models.FuturesAccountBalanceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.recvWindow != nil {
		localVarQueryParams["recvWindow"] = *r.recvWindow
	}

	localPayload := map[string]any{
		"method": "/account.balance"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           true,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.FuturesAccountBalanceResponse](a.Ws, localPayload, sendParams)
}
