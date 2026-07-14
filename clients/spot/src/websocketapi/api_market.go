/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotwebsocketapi

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketAPIService MarketAPI Service
type MarketAPIService struct {
	Ws *common.WebsocketAPI
}

type ApiAvgPriceRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	id         *string
}

func (r ApiAvgPriceRequest) Symbol(symbol string) ApiAvgPriceRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiAvgPriceRequest) Id(id string) ApiAvgPriceRequest {
	r.id = &id
	return r
}

func (r ApiAvgPriceRequest) Execute() (*common.ResponseOrRaw[models.AvgPriceResponse], error) {
	respChan, errChan, err := r.ApiService.AvgPriceExecute(r)
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

func (r ApiAvgPriceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.AvgPriceResponse], chan error, error) {
	return r.ApiService.AvgPriceExecute(r)
}

/*
AvgPrice Current average price
/avgPrice

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#avg-price

@param symbol	@param id Client-generated request identifier.
@return ApiAvgPriceRequest
*/
func (a *MarketAPIService) AvgPrice() ApiAvgPriceRequest {
	return ApiAvgPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AvgPriceResponse
func (a *MarketAPIService) AvgPriceExecute(r ApiAvgPriceRequest) (chan *common.ResponseOrRaw[models.AvgPriceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/avgPrice"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.AvgPriceResponse](a.Ws, localPayload, sendParams)
}

type ApiBlockTradesHistoricalRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	fromId     *int64
	id         *string
	limit      *int64
}

func (r ApiBlockTradesHistoricalRequest) Symbol(symbol string) ApiBlockTradesHistoricalRequest {
	r.symbol = &symbol
	return r
}

// Block trade ID to fetch from
func (r ApiBlockTradesHistoricalRequest) FromId(fromId int64) ApiBlockTradesHistoricalRequest {
	r.fromId = &fromId
	return r
}

// Client-generated request identifier.
func (r ApiBlockTradesHistoricalRequest) Id(id string) ApiBlockTradesHistoricalRequest {
	r.id = &id
	return r
}

// Default: 500; Maximum: 1000
func (r ApiBlockTradesHistoricalRequest) Limit(limit int64) ApiBlockTradesHistoricalRequest {
	r.limit = &limit
	return r
}

func (r ApiBlockTradesHistoricalRequest) Execute() (*common.ResponseOrRaw[models.BlockTradesHistoricalResponse], error) {
	respChan, errChan, err := r.ApiService.BlockTradesHistoricalExecute(r)
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

func (r ApiBlockTradesHistoricalRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.BlockTradesHistoricalResponse], chan error, error) {
	return r.ApiService.BlockTradesHistoricalExecute(r)
}

/*
BlockTradesHistorical Historical Block Trades
/blockTrades.historical

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#block-trades-historical

@param symbol	@param fromId Block trade ID to fetch from	@param id Client-generated request identifier.	@param limit Default: 500; Maximum: 1000
@return ApiBlockTradesHistoricalRequest
*/
func (a *MarketAPIService) BlockTradesHistorical() ApiBlockTradesHistoricalRequest {
	return ApiBlockTradesHistoricalRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BlockTradesHistoricalResponse
func (a *MarketAPIService) BlockTradesHistoricalExecute(r ApiBlockTradesHistoricalRequest) (chan *common.ResponseOrRaw[models.BlockTradesHistoricalResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.fromId == nil {
		return nil, nil, common.ReportError("fromId is required and must be specified")
	}
	localVarQueryParams["fromId"] = *r.fromId

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/blockTrades.historical"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.BlockTradesHistoricalResponse](a.Ws, localPayload, sendParams)
}

type ApiDepthRequest struct {
	ApiService   *MarketAPIService
	symbol       *string
	id           *string
	limit        *int32
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

func (r ApiDepthRequest) Symbol(symbol string) ApiDepthRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiDepthRequest) Id(id string) ApiDepthRequest {
	r.id = &id
	return r
}

func (r ApiDepthRequest) Limit(limit int32) ApiDepthRequest {
	r.limit = &limit
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. A status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;.
func (r ApiDepthRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiDepthRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiDepthRequest) Execute() (*common.ResponseOrRaw[models.DepthResponse], error) {
	respChan, errChan, err := r.ApiService.DepthExecute(r)
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

func (r ApiDepthRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.DepthResponse], chan error, error) {
	return r.ApiService.DepthExecute(r)
}

/*
Depth Order book
/depth

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#depth

@param symbol	@param id Client-generated request identifier.	@param limit	@param symbolStatus Filters for symbols that have this `tradingStatus`. A status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`.
@return ApiDepthRequest
*/
func (a *MarketAPIService) Depth() ApiDepthRequest {
	return ApiDepthRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DepthResponse
func (a *MarketAPIService) DepthExecute(r ApiDepthRequest) (chan *common.ResponseOrRaw[models.DepthResponse], chan error, error) {
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
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
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

	return SendMessage[models.DepthResponse](a.Ws, localPayload, sendParams)
}

type ApiKlinesRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	interval   *models.KlinesIntervalParameter
	id         *string
	startTime  *int64
	endTime    *int64
	timeZone   *string
	limit      *int32
}

func (r ApiKlinesRequest) Symbol(symbol string) ApiKlinesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlinesRequest) Interval(interval models.KlinesIntervalParameter) ApiKlinesRequest {
	r.interval = &interval
	return r
}

// Client-generated request identifier.
func (r ApiKlinesRequest) Id(id string) ApiKlinesRequest {
	r.id = &id
	return r
}

func (r ApiKlinesRequest) StartTime(startTime int64) ApiKlinesRequest {
	r.startTime = &startTime
	return r
}

func (r ApiKlinesRequest) EndTime(endTime int64) ApiKlinesRequest {
	r.endTime = &endTime
	return r
}

// Default: 0 (UTC)
func (r ApiKlinesRequest) TimeZone(timeZone string) ApiKlinesRequest {
	r.timeZone = &timeZone
	return r
}

func (r ApiKlinesRequest) Limit(limit int32) ApiKlinesRequest {
	r.limit = &limit
	return r
}

func (r ApiKlinesRequest) Execute() (*common.ResponseOrRaw[models.KlinesResponse], error) {
	respChan, errChan, err := r.ApiService.KlinesExecute(r)
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

func (r ApiKlinesRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.KlinesResponse], chan error, error) {
	return r.ApiService.KlinesExecute(r)
}

/*
Klines Klines
/klines

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#klines

@param symbol	@param interval	@param id Client-generated request identifier.	@param startTime	@param endTime	@param timeZone Default: 0 (UTC)	@param limit
@return ApiKlinesRequest
*/
func (a *MarketAPIService) Klines() ApiKlinesRequest {
	return ApiKlinesRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlinesResponse
func (a *MarketAPIService) KlinesExecute(r ApiKlinesRequest) (chan *common.ResponseOrRaw[models.KlinesResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.interval == nil {
		return nil, nil, common.ReportError("interval is required and must be specified")
	}
	localVarQueryParams["interval"] = *r.interval

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.timeZone != nil {
		localVarQueryParams["timeZone"] = *r.timeZone
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/klines"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.KlinesResponse](a.Ws, localPayload, sendParams)
}

type ApiReferencePriceRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	id         *string
}

func (r ApiReferencePriceRequest) Symbol(symbol string) ApiReferencePriceRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiReferencePriceRequest) Id(id string) ApiReferencePriceRequest {
	r.id = &id
	return r
}

func (r ApiReferencePriceRequest) Execute() (*common.ResponseOrRaw[models.ReferencePriceResponse], error) {
	respChan, errChan, err := r.ApiService.ReferencePriceExecute(r)
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

func (r ApiReferencePriceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.ReferencePriceResponse], chan error, error) {
	return r.ApiService.ReferencePriceExecute(r)
}

/*
ReferencePrice Query Reference Price
/referencePrice

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#reference-price

@param symbol	@param id Client-generated request identifier.
@return ApiReferencePriceRequest
*/
func (a *MarketAPIService) ReferencePrice() ApiReferencePriceRequest {
	return ApiReferencePriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ReferencePriceResponse
func (a *MarketAPIService) ReferencePriceExecute(r ApiReferencePriceRequest) (chan *common.ResponseOrRaw[models.ReferencePriceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}

	localPayload := map[string]any{
		"method": "/referencePrice"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.ReferencePriceResponse](a.Ws, localPayload, sendParams)
}

type ApiReferencePriceCalculationRequest struct {
	ApiService   *MarketAPIService
	symbol       *string
	id           *string
	symbolStatus *models.ReferencePriceCalculationSymbolStatusParameter
}

func (r ApiReferencePriceCalculationRequest) Symbol(symbol string) ApiReferencePriceCalculationRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiReferencePriceCalculationRequest) Id(id string) ApiReferencePriceCalculationRequest {
	r.id = &id
	return r
}

func (r ApiReferencePriceCalculationRequest) SymbolStatus(symbolStatus models.ReferencePriceCalculationSymbolStatusParameter) ApiReferencePriceCalculationRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiReferencePriceCalculationRequest) Execute() (*common.ResponseOrRaw[models.ReferencePriceCalculationResponse], error) {
	respChan, errChan, err := r.ApiService.ReferencePriceCalculationExecute(r)
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

func (r ApiReferencePriceCalculationRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.ReferencePriceCalculationResponse], chan error, error) {
	return r.ApiService.ReferencePriceCalculationExecute(r)
}

/*
ReferencePriceCalculation Query Reference Price Calculation
/referencePrice.calculation

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#reference-price-calculation

@param symbol	@param id Client-generated request identifier.	@param symbolStatus
@return ApiReferencePriceCalculationRequest
*/
func (a *MarketAPIService) ReferencePriceCalculation() ApiReferencePriceCalculationRequest {
	return ApiReferencePriceCalculationRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ReferencePriceCalculationResponse
func (a *MarketAPIService) ReferencePriceCalculationExecute(r ApiReferencePriceCalculationRequest) (chan *common.ResponseOrRaw[models.ReferencePriceCalculationResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
	}

	localPayload := map[string]any{
		"method": "/referencePrice.calculation"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.ReferencePriceCalculationResponse](a.Ws, localPayload, sendParams)
}

type ApiTickerRequest struct {
	ApiService   *MarketAPIService
	id           *string
	symbol       *string
	symbols      *[]string
	type_        *models.TickerTypeParameter
	windowSize   *models.TickerWindowSizeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Client-generated request identifier.
func (r ApiTickerRequest) Id(id string) ApiTickerRequest {
	r.id = &id
	return r
}

// Query ticker of a single symbol
func (r ApiTickerRequest) Symbol(symbol string) ApiTickerRequest {
	r.symbol = &symbol
	return r
}

// Query ticker for multiple symbols
func (r ApiTickerRequest) Symbols(symbols []string) ApiTickerRequest {
	r.symbols = &symbols
	return r
}

// Ticker type. Supported values: FULL (default) or MINI
func (r ApiTickerRequest) Type(type_ models.TickerTypeParameter) ApiTickerRequest {
	r.type_ = &type_
	return r
}

// Defaults to 1d if no parameter provided.
func (r ApiTickerRequest) WindowSize(windowSize models.TickerWindowSizeParameter) ApiTickerRequest {
	r.windowSize = &windowSize
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;
func (r ApiTickerRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerRequest) Execute() (*common.ResponseOrRaw[models.TickerResponse], error) {
	respChan, errChan, err := r.ApiService.TickerExecute(r)
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

func (r ApiTickerRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TickerResponse], chan error, error) {
	return r.ApiService.TickerExecute(r)
}

/*
Ticker Rolling window price change statistics
/ticker

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ticker

@param id Client-generated request identifier.	@param symbol Query ticker of a single symbol	@param symbols Query ticker for multiple symbols	@param type_ Ticker type. Supported values: FULL (default) or MINI	@param windowSize Defaults to 1d if no parameter provided.	@param symbolStatus Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK`
@return ApiTickerRequest
*/
func (a *MarketAPIService) Ticker() ApiTickerRequest {
	return ApiTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerResponse
func (a *MarketAPIService) TickerExecute(r ApiTickerRequest) (chan *common.ResponseOrRaw[models.TickerResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.type_ != nil {
		localVarQueryParams["type_"] = *r.type_
	}
	if r.windowSize != nil {
		localVarQueryParams["windowSize"] = *r.windowSize
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
	}

	localPayload := map[string]any{
		"method": "/ticker"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TickerResponse](a.Ws, localPayload, sendParams)
}

type ApiTicker24hrRequest struct {
	ApiService   *MarketAPIService
	id           *string
	symbol       *string
	symbols      *[]string
	type_        *models.TickerTypeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Client-generated request identifier.
func (r ApiTicker24hrRequest) Id(id string) ApiTicker24hrRequest {
	r.id = &id
	return r
}

func (r ApiTicker24hrRequest) Symbol(symbol string) ApiTicker24hrRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTicker24hrRequest) Symbols(symbols []string) ApiTicker24hrRequest {
	r.symbols = &symbols
	return r
}

// Ticker type. Supported values: FULL (default) or MINI
func (r ApiTicker24hrRequest) Type(type_ models.TickerTypeParameter) ApiTicker24hrRequest {
	r.type_ = &type_
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;
func (r ApiTicker24hrRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTicker24hrRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTicker24hrRequest) Execute() (*common.ResponseOrRaw[models.Ticker24hrResponse], error) {
	respChan, errChan, err := r.ApiService.Ticker24hrExecute(r)
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

func (r ApiTicker24hrRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.Ticker24hrResponse], chan error, error) {
	return r.ApiService.Ticker24hrExecute(r)
}

/*
Ticker24hr 24hr ticker price change statistics
/ticker.24hr

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ticker24hr

@param id Client-generated request identifier.	@param symbol	@param symbols	@param type_ Ticker type. Supported values: FULL (default) or MINI	@param symbolStatus Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK`
@return ApiTicker24hrRequest
*/
func (a *MarketAPIService) Ticker24hr() ApiTicker24hrRequest {
	return ApiTicker24hrRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return Ticker24hrResponse
func (a *MarketAPIService) Ticker24hrExecute(r ApiTicker24hrRequest) (chan *common.ResponseOrRaw[models.Ticker24hrResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.type_ != nil {
		localVarQueryParams["type_"] = *r.type_
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
	}

	localPayload := map[string]any{
		"method": "/ticker.24hr"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.Ticker24hrResponse](a.Ws, localPayload, sendParams)
}

type ApiTickerBookRequest struct {
	ApiService   *MarketAPIService
	id           *string
	symbol       *string
	symbols      *[]string
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Client-generated request identifier.
func (r ApiTickerBookRequest) Id(id string) ApiTickerBookRequest {
	r.id = &id
	return r
}

// Query ticker for a single symbol
func (r ApiTickerBookRequest) Symbol(symbol string) ApiTickerBookRequest {
	r.symbol = &symbol
	return r
}

// Query ticker for multiple symbols
func (r ApiTickerBookRequest) Symbols(symbols []string) ApiTickerBookRequest {
	r.symbols = &symbols
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;
func (r ApiTickerBookRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerBookRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerBookRequest) Execute() (*common.ResponseOrRaw[models.TickerBookResponse], error) {
	respChan, errChan, err := r.ApiService.TickerBookExecute(r)
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

func (r ApiTickerBookRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TickerBookResponse], chan error, error) {
	return r.ApiService.TickerBookExecute(r)
}

/*
TickerBook Symbol order book ticker
/ticker.book

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ticker-book

@param id Client-generated request identifier.	@param symbol Query ticker for a single symbol	@param symbols Query ticker for multiple symbols	@param symbolStatus Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK`
@return ApiTickerBookRequest
*/
func (a *MarketAPIService) TickerBook() ApiTickerBookRequest {
	return ApiTickerBookRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerBookResponse
func (a *MarketAPIService) TickerBookExecute(r ApiTickerBookRequest) (chan *common.ResponseOrRaw[models.TickerBookResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
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

	return SendMessage[models.TickerBookResponse](a.Ws, localPayload, sendParams)
}

type ApiTickerPriceRequest struct {
	ApiService   *MarketAPIService
	id           *string
	symbol       *string
	symbols      *[]string
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Client-generated request identifier.
func (r ApiTickerPriceRequest) Id(id string) ApiTickerPriceRequest {
	r.id = &id
	return r
}

// Query price for a single symbol
func (r ApiTickerPriceRequest) Symbol(symbol string) ApiTickerPriceRequest {
	r.symbol = &symbol
	return r
}

// Query price for multiple symbols
func (r ApiTickerPriceRequest) Symbols(symbols []string) ApiTickerPriceRequest {
	r.symbols = &symbols
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;
func (r ApiTickerPriceRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerPriceRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerPriceRequest) Execute() (*common.ResponseOrRaw[models.TickerPriceResponse], error) {
	respChan, errChan, err := r.ApiService.TickerPriceExecute(r)
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

func (r ApiTickerPriceRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TickerPriceResponse], chan error, error) {
	return r.ApiService.TickerPriceExecute(r)
}

/*
TickerPrice Symbol price ticker
/ticker.price

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ticker-price

@param id Client-generated request identifier.	@param symbol Query price for a single symbol	@param symbols Query price for multiple symbols	@param symbolStatus Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK`
@return ApiTickerPriceRequest
*/
func (a *MarketAPIService) TickerPrice() ApiTickerPriceRequest {
	return ApiTickerPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerPriceResponse
func (a *MarketAPIService) TickerPriceExecute(r ApiTickerPriceRequest) (chan *common.ResponseOrRaw[models.TickerPriceResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
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

	return SendMessage[models.TickerPriceResponse](a.Ws, localPayload, sendParams)
}

type ApiTickerTradingDayRequest struct {
	ApiService   *MarketAPIService
	id           *string
	symbol       *string
	symbols      *[]string
	timeZone     *string
	type_        *models.TickerTypeParameter
	symbolStatus *models.ExchangeInfoSymbolStatusParameter
}

// Client-generated request identifier.
func (r ApiTickerTradingDayRequest) Id(id string) ApiTickerTradingDayRequest {
	r.id = &id
	return r
}

func (r ApiTickerTradingDayRequest) Symbol(symbol string) ApiTickerTradingDayRequest {
	r.symbol = &symbol
	return r
}

func (r ApiTickerTradingDayRequest) Symbols(symbols []string) ApiTickerTradingDayRequest {
	r.symbols = &symbols
	return r
}

// Default: 0 (UTC)
func (r ApiTickerTradingDayRequest) TimeZone(timeZone string) ApiTickerTradingDayRequest {
	r.timeZone = &timeZone
	return r
}

// Ticker type. Supported values: FULL (default) or MINI
func (r ApiTickerTradingDayRequest) Type(type_ models.TickerTypeParameter) ApiTickerTradingDayRequest {
	r.type_ = &type_
	return r
}

// Filters for symbols that have this &#x60;tradingStatus&#x60;. For a single symbol, a status mismatch returns error &#x60;-1220 SYMBOL_DOES_NOT_MATCH_STATUS&#x60;. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60;
func (r ApiTickerTradingDayRequest) SymbolStatus(symbolStatus models.ExchangeInfoSymbolStatusParameter) ApiTickerTradingDayRequest {
	r.symbolStatus = &symbolStatus
	return r
}

func (r ApiTickerTradingDayRequest) Execute() (*common.ResponseOrRaw[models.TickerTradingDayResponse], error) {
	respChan, errChan, err := r.ApiService.TickerTradingDayExecute(r)
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

func (r ApiTickerTradingDayRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TickerTradingDayResponse], chan error, error) {
	return r.ApiService.TickerTradingDayExecute(r)
}

/*
TickerTradingDay Trading Day Ticker
/ticker.tradingDay

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ticker-trading-day

@param id Client-generated request identifier.	@param symbol	@param symbols	@param timeZone Default: 0 (UTC)	@param type_ Ticker type. Supported values: FULL (default) or MINI	@param symbolStatus Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple or all symbols, non-matching ones are simply excluded from the response. Valid values: `TRADING`, `HALT`, `BREAK`
@return ApiTickerTradingDayRequest
*/
func (a *MarketAPIService) TickerTradingDay() ApiTickerTradingDayRequest {
	return ApiTickerTradingDayRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerTradingDayResponse
func (a *MarketAPIService) TickerTradingDayExecute(r ApiTickerTradingDayRequest) (chan *common.ResponseOrRaw[models.TickerTradingDayResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.symbol != nil {
		localVarQueryParams["symbol"] = *r.symbol
	}
	if r.symbols != nil {
		localVarQueryParams["symbols"] = *r.symbols
	}
	if r.timeZone != nil {
		localVarQueryParams["timeZone"] = *r.timeZone
	}
	if r.type_ != nil {
		localVarQueryParams["type_"] = *r.type_
	}
	if r.symbolStatus != nil {
		localVarQueryParams["symbolStatus"] = *r.symbolStatus
	}

	localPayload := map[string]any{
		"method": "/ticker.tradingDay"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TickerTradingDayResponse](a.Ws, localPayload, sendParams)
}

type ApiTradesAggregateRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	id         *string
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int32
}

func (r ApiTradesAggregateRequest) Symbol(symbol string) ApiTradesAggregateRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiTradesAggregateRequest) Id(id string) ApiTradesAggregateRequest {
	r.id = &id
	return r
}

// Aggregate trade ID to begin at
func (r ApiTradesAggregateRequest) FromId(fromId int64) ApiTradesAggregateRequest {
	r.fromId = &fromId
	return r
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (r ApiTradesAggregateRequest) StartTime(startTime int64) ApiTradesAggregateRequest {
	r.startTime = &startTime
	return r
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (r ApiTradesAggregateRequest) EndTime(endTime int64) ApiTradesAggregateRequest {
	r.endTime = &endTime
	return r
}

func (r ApiTradesAggregateRequest) Limit(limit int32) ApiTradesAggregateRequest {
	r.limit = &limit
	return r
}

func (r ApiTradesAggregateRequest) Execute() (*common.ResponseOrRaw[models.TradesAggregateResponse], error) {
	respChan, errChan, err := r.ApiService.TradesAggregateExecute(r)
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

func (r ApiTradesAggregateRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TradesAggregateResponse], chan error, error) {
	return r.ApiService.TradesAggregateExecute(r)
}

/*
TradesAggregate Aggregate trades
/trades.aggregate

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#trades-aggregate

@param symbol	@param id Client-generated request identifier.	@param fromId Aggregate trade ID to begin at	@param startTime Timestamp in ms to get aggregate trades from INCLUSIVE.	@param endTime Timestamp in ms to get aggregate trades until INCLUSIVE.	@param limit
@return ApiTradesAggregateRequest
*/
func (a *MarketAPIService) TradesAggregate() ApiTradesAggregateRequest {
	return ApiTradesAggregateRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradesAggregateResponse
func (a *MarketAPIService) TradesAggregateExecute(r ApiTradesAggregateRequest) (chan *common.ResponseOrRaw[models.TradesAggregateResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.fromId != nil {
		localVarQueryParams["fromId"] = *r.fromId
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/trades.aggregate"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TradesAggregateResponse](a.Ws, localPayload, sendParams)
}

type ApiTradesHistoricalRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	id         *string
	fromId     *int64
	limit      *int32
}

func (r ApiTradesHistoricalRequest) Symbol(symbol string) ApiTradesHistoricalRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiTradesHistoricalRequest) Id(id string) ApiTradesHistoricalRequest {
	r.id = &id
	return r
}

// Trade ID to begin at
func (r ApiTradesHistoricalRequest) FromId(fromId int64) ApiTradesHistoricalRequest {
	r.fromId = &fromId
	return r
}

func (r ApiTradesHistoricalRequest) Limit(limit int32) ApiTradesHistoricalRequest {
	r.limit = &limit
	return r
}

func (r ApiTradesHistoricalRequest) Execute() (*common.ResponseOrRaw[models.TradesHistoricalResponse], error) {
	respChan, errChan, err := r.ApiService.TradesHistoricalExecute(r)
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

func (r ApiTradesHistoricalRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TradesHistoricalResponse], chan error, error) {
	return r.ApiService.TradesHistoricalExecute(r)
}

/*
TradesHistorical Historical trades
/trades.historical

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#trades-historical

@param symbol	@param id Client-generated request identifier.	@param fromId Trade ID to begin at	@param limit
@return ApiTradesHistoricalRequest
*/
func (a *MarketAPIService) TradesHistorical() ApiTradesHistoricalRequest {
	return ApiTradesHistoricalRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradesHistoricalResponse
func (a *MarketAPIService) TradesHistoricalExecute(r ApiTradesHistoricalRequest) (chan *common.ResponseOrRaw[models.TradesHistoricalResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.fromId != nil {
		localVarQueryParams["fromId"] = *r.fromId
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/trades.historical"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TradesHistoricalResponse](a.Ws, localPayload, sendParams)
}

type ApiTradesRecentRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	id         *string
	limit      *int32
}

func (r ApiTradesRecentRequest) Symbol(symbol string) ApiTradesRecentRequest {
	r.symbol = &symbol
	return r
}

// Client-generated request identifier.
func (r ApiTradesRecentRequest) Id(id string) ApiTradesRecentRequest {
	r.id = &id
	return r
}

func (r ApiTradesRecentRequest) Limit(limit int32) ApiTradesRecentRequest {
	r.limit = &limit
	return r
}

func (r ApiTradesRecentRequest) Execute() (*common.ResponseOrRaw[models.TradesRecentResponse], error) {
	respChan, errChan, err := r.ApiService.TradesRecentExecute(r)
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

func (r ApiTradesRecentRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.TradesRecentResponse], chan error, error) {
	return r.ApiService.TradesRecentExecute(r)
}

/*
TradesRecent Recent trades
/trades.recent

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#trades-recent

@param symbol	@param id Client-generated request identifier.	@param limit
@return ApiTradesRecentRequest
*/
func (a *MarketAPIService) TradesRecent() ApiTradesRecentRequest {
	return ApiTradesRecentRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradesRecentResponse
func (a *MarketAPIService) TradesRecentExecute(r ApiTradesRecentRequest) (chan *common.ResponseOrRaw[models.TradesRecentResponse], chan error, error) {
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
		"method": "/trades.recent"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.TradesRecentResponse](a.Ws, localPayload, sendParams)
}

type ApiUiKlinesRequest struct {
	ApiService *MarketAPIService
	symbol     *string
	interval   *models.KlinesIntervalParameter
	id         *string
	startTime  *int64
	endTime    *int64
	timeZone   *string
	limit      *int32
}

func (r ApiUiKlinesRequest) Symbol(symbol string) ApiUiKlinesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiUiKlinesRequest) Interval(interval models.KlinesIntervalParameter) ApiUiKlinesRequest {
	r.interval = &interval
	return r
}

// Client-generated request identifier.
func (r ApiUiKlinesRequest) Id(id string) ApiUiKlinesRequest {
	r.id = &id
	return r
}

func (r ApiUiKlinesRequest) StartTime(startTime int64) ApiUiKlinesRequest {
	r.startTime = &startTime
	return r
}

func (r ApiUiKlinesRequest) EndTime(endTime int64) ApiUiKlinesRequest {
	r.endTime = &endTime
	return r
}

// Default: 0 (UTC)
func (r ApiUiKlinesRequest) TimeZone(timeZone string) ApiUiKlinesRequest {
	r.timeZone = &timeZone
	return r
}

func (r ApiUiKlinesRequest) Limit(limit int32) ApiUiKlinesRequest {
	r.limit = &limit
	return r
}

func (r ApiUiKlinesRequest) Execute() (*common.ResponseOrRaw[models.UiKlinesResponse], error) {
	respChan, errChan, err := r.ApiService.UiKlinesExecute(r)
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

func (r ApiUiKlinesRequest) ExecuteAsync() (chan *common.ResponseOrRaw[models.UiKlinesResponse], chan error, error) {
	return r.ApiService.UiKlinesExecute(r)
}

/*
UiKlines UI Klines
/uiKlines

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-api/market#ui-klines

@param symbol	@param interval	@param id Client-generated request identifier.	@param startTime	@param endTime	@param timeZone Default: 0 (UTC)	@param limit
@return ApiUiKlinesRequest
*/
func (a *MarketAPIService) UiKlines() ApiUiKlinesRequest {
	return ApiUiKlinesRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return UiKlinesResponse
func (a *MarketAPIService) UiKlinesExecute(r ApiUiKlinesRequest) (chan *common.ResponseOrRaw[models.UiKlinesResponse], chan error, error) {
	localVarQueryParams := map[string]any{}

	if r.symbol == nil {
		return nil, nil, common.ReportError("symbol is required and must be specified")
	}
	localVarQueryParams["symbol"] = *r.symbol

	if r.interval == nil {
		return nil, nil, common.ReportError("interval is required and must be specified")
	}
	localVarQueryParams["interval"] = *r.interval

	if r.id != nil {
		localVarQueryParams["id"] = *r.id
	}
	if r.startTime != nil {
		localVarQueryParams["startTime"] = *r.startTime
	}
	if r.endTime != nil {
		localVarQueryParams["endTime"] = *r.endTime
	}
	if r.timeZone != nil {
		localVarQueryParams["timeZone"] = *r.timeZone
	}
	if r.limit != nil {
		localVarQueryParams["limit"] = *r.limit
	}

	localPayload := map[string]any{
		"method": "/uiKlines"[1:],
		"params": localVarQueryParams,
	}

	sendParams := common.SendParams{
		Signed:           false,
		WithAPIKey:       false,
		WithSessionLogon: false,
	}

	return SendMessage[models.UiKlinesResponse](a.Ws, localPayload, sendParams)
}
