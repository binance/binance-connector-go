/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package binancederivativestradingusdsfutureswebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketDataAPIService MarketDataAPI Service
type MarketDataAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiOrderBookRequest struct {
	ApiService *MarketDataAPIService
	symbol     *string
	id         *string
	limit      *int64
}

func (r ApiOrderBookRequest) Symbol(symbol string) ApiOrderBookRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiOrderBookRequest) Id(id string) ApiOrderBookRequest {
	r.id = &id
	return r
}

// Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000]
func (r ApiOrderBookRequest) Limit(limit int64) ApiOrderBookRequest {
	r.limit = &limit
	return r
}

func (r ApiOrderBookRequest) Execute() (*common.ResponseOrRaw[models.OrderBookResponse], error) {
	respChan, errChan, err := r.ApiService.OrderBookExecute(r)
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

func (r ApiOrderBookRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.OrderBookResponse], chan error, error) {
	return r.ApiService.OrderBookExecute(r)
}

/*
OrderBook Order Book
/depth

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/websocket-api/Order-Book

@param symbol	@param id Unique WebSocket request ID.	@param limit Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000]
@return ApiOrderBookRequest
*/
func (a *MarketDataAPIService) OrderBook() ApiOrderBookRequest {
	return ApiOrderBookRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderBookResponse
func (a *MarketDataAPIService) OrderBookExecute(r ApiOrderBookRequest) (chan *common.ResponseOrRaw[models.OrderBookResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/depth"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.OrderBookResponse](a.Ws, localPayload, sendParams)
}

type ApiSymbolOrderBookTickerRequest struct {
	ApiService *MarketDataAPIService
	id         *string
	symbol     *string
}

// Unique WebSocket request ID.
func (r ApiSymbolOrderBookTickerRequest) Id(id string) ApiSymbolOrderBookTickerRequest {
	r.id = &id
	return r
}

func (r ApiSymbolOrderBookTickerRequest) Symbol(symbol string) ApiSymbolOrderBookTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolOrderBookTickerRequest) Execute() (*common.ResponseOrRaw[models.SymbolOrderBookTickerResponse], error) {
	respChan, errChan, err := r.ApiService.SymbolOrderBookTickerExecute(r)
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

func (r ApiSymbolOrderBookTickerRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SymbolOrderBookTickerResponse], chan error, error) {
	return r.ApiService.SymbolOrderBookTickerExecute(r)
}

/*
SymbolOrderBookTicker Symbol Order Book Ticker
/ticker.book

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/websocket-api/Symbol-Order-Book-Ticker

@param id Unique WebSocket request ID.	@param symbol
@return ApiSymbolOrderBookTickerRequest
*/
func (a *MarketDataAPIService) SymbolOrderBookTicker() ApiSymbolOrderBookTickerRequest {
	return ApiSymbolOrderBookTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SymbolOrderBookTickerResponse
func (a *MarketDataAPIService) SymbolOrderBookTickerExecute(r ApiSymbolOrderBookTickerRequest) (chan *common.ResponseOrRaw[models.SymbolOrderBookTickerResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}

	localPayload := map[string]any{
		"method": "/ticker.book"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SymbolOrderBookTickerResponse](a.Ws, localPayload, sendParams)
}

type ApiSymbolPriceTickerRequest struct {
	ApiService *MarketDataAPIService
	id         *string
	symbol     *string
}

// Unique WebSocket request ID.
func (r ApiSymbolPriceTickerRequest) Id(id string) ApiSymbolPriceTickerRequest {
	r.id = &id
	return r
}

func (r ApiSymbolPriceTickerRequest) Symbol(symbol string) ApiSymbolPriceTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiSymbolPriceTickerRequest) Execute() (*common.ResponseOrRaw[models.SymbolPriceTickerResponse], error) {
	respChan, errChan, err := r.ApiService.SymbolPriceTickerExecute(r)
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

func (r ApiSymbolPriceTickerRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.SymbolPriceTickerResponse], chan error, error) {
	return r.ApiService.SymbolPriceTickerExecute(r)
}

/*
SymbolPriceTicker Symbol Price Ticker
/ticker.price

https://developers.binance.com/docs/derivatives/usds-margined-futures/market-data/websocket-api/Symbol-Price-Ticker

@param id Unique WebSocket request ID.	@param symbol
@return ApiSymbolPriceTickerRequest
*/
func (a *MarketDataAPIService) SymbolPriceTicker() ApiSymbolPriceTickerRequest {
	return ApiSymbolPriceTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return SymbolPriceTickerResponse
func (a *MarketDataAPIService) SymbolPriceTickerExecute(r ApiSymbolPriceTickerRequest) (chan *common.ResponseOrRaw[models.SymbolPriceTickerResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}

	localPayload := map[string]any{
		"method": "/ticker.price"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.SymbolPriceTickerResponse](a.Ws, localPayload, sendParams)
}
