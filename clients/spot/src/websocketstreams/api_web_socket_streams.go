/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package binancespotwebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

// WebSocketStreamsAPIService WebSocketStreamsAPI Service
type WebSocketStreamsAPIService Service

type ApiAggTradeRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiAggTradeRequest) Symbol(symbol string) ApiAggTradeRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiAggTradeRequest) Id(id string) ApiAggTradeRequest {
	r.id = &id
	return r
}

func (r ApiAggTradeRequest) Execute() (*common.StreamHandler[models.AggTradeResponse], error) {
	return r.ApiService.AggTradeExecute(r)
}

/*
AggTrade WebSocket Aggregate Trade Streams
/<symbol>@aggTrade

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#aggregate-trade-streams

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiAggTradeRequest
*/
func (a *WebSocketStreamsAPIService) AggTrade() ApiAggTradeRequest {
	return ApiAggTradeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggTradeResponse
func (a *WebSocketStreamsAPIService) AggTradeExecute(r ApiAggTradeRequest) (*common.StreamHandler[models.AggTradeResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@aggTrade"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.AggTradeResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMarketRollingWindowTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	windowSize *models.AllMarketRollingWindowTickerWindowSizeParameter
	id         *string
}

func (r ApiAllMarketRollingWindowTickerRequest) WindowSize(windowSize models.AllMarketRollingWindowTickerWindowSizeParameter) ApiAllMarketRollingWindowTickerRequest {
	r.windowSize = &windowSize
	return r
}

// Unique WebSocket request ID.
func (r ApiAllMarketRollingWindowTickerRequest) Id(id string) ApiAllMarketRollingWindowTickerRequest {
	r.id = &id
	return r
}

func (r ApiAllMarketRollingWindowTickerRequest) Execute() (*common.StreamHandler[models.AllMarketRollingWindowTickerResponse], error) {
	return r.ApiService.AllMarketRollingWindowTickerExecute(r)
}

/*
AllMarketRollingWindowTicker WebSocket All Market Rolling Window Statistics Streams
/!ticker_<windowSize>@arr

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#all-market-rolling-window-statistics-streams

@param windowSize	@param id Unique WebSocket request ID.
@return ApiAllMarketRollingWindowTickerRequest
*/
func (a *WebSocketStreamsAPIService) AllMarketRollingWindowTicker() ApiAllMarketRollingWindowTickerRequest {
	return ApiAllMarketRollingWindowTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketRollingWindowTickerResponse
func (a *WebSocketStreamsAPIService) AllMarketRollingWindowTickerExecute(r ApiAllMarketRollingWindowTickerRequest) (*common.StreamHandler[models.AllMarketRollingWindowTickerResponse], error) {
	if r.windowSize == nil {
		return nil, common.ReportError("windowSize is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/!ticker_<windowSize>@arr"[1:],
		map[string]string{
			"windowSize": func() string {
				if r.windowSize == nil {
					return ""
				}
				return string(*r.windowSize)
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
	resp, err := common.CreateStreamHandler[models.AllMarketRollingWindowTickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMiniTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllMiniTickerRequest) Id(id string) ApiAllMiniTickerRequest {
	r.id = &id
	return r
}

func (r ApiAllMiniTickerRequest) Execute() (*common.StreamHandler[models.AllMiniTickerResponse], error) {
	return r.ApiService.AllMiniTickerExecute(r)
}

/*
AllMiniTicker WebSocket All Market Mini Tickers Stream
/!miniTicker@arr

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#all-market-mini-tickers-stream

@param id Unique WebSocket request ID.
@return ApiAllMiniTickerRequest
*/
func (a *WebSocketStreamsAPIService) AllMiniTicker() ApiAllMiniTickerRequest {
	return ApiAllMiniTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMiniTickerResponse
func (a *WebSocketStreamsAPIService) AllMiniTickerExecute(r ApiAllMiniTickerRequest) (*common.StreamHandler[models.AllMiniTickerResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!miniTicker@arr"[1:],
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
	resp, err := common.CreateStreamHandler[models.AllMiniTickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAvgPriceRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiAvgPriceRequest) Symbol(symbol string) ApiAvgPriceRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiAvgPriceRequest) Id(id string) ApiAvgPriceRequest {
	r.id = &id
	return r
}

func (r ApiAvgPriceRequest) Execute() (*common.StreamHandler[models.AvgPriceResponse], error) {
	return r.ApiService.AvgPriceExecute(r)
}

/*
AvgPrice WebSocket Average Price
/<symbol>@avgPrice

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#average-price

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiAvgPriceRequest
*/
func (a *WebSocketStreamsAPIService) AvgPrice() ApiAvgPriceRequest {
	return ApiAvgPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AvgPriceResponse
func (a *WebSocketStreamsAPIService) AvgPriceExecute(r ApiAvgPriceRequest) (*common.StreamHandler[models.AvgPriceResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@avgPrice"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.AvgPriceResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiBookTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiBookTickerRequest) Symbol(symbol string) ApiBookTickerRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiBookTickerRequest) Id(id string) ApiBookTickerRequest {
	r.id = &id
	return r
}

func (r ApiBookTickerRequest) Execute() (*common.StreamHandler[models.BookTickerResponse], error) {
	return r.ApiService.BookTickerExecute(r)
}

/*
BookTicker WebSocket Individual Symbol Book Ticker Streams
/<symbol>@bookTicker

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#individual-symbol-book-ticker-streams

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiBookTickerRequest
*/
func (a *WebSocketStreamsAPIService) BookTicker() ApiBookTickerRequest {
	return ApiBookTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BookTickerResponse
func (a *WebSocketStreamsAPIService) BookTickerExecute(r ApiBookTickerRequest) (*common.StreamHandler[models.BookTickerResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@bookTicker"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.BookTickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiDiffBookDepthRequest struct {
	ApiService  *WebSocketStreamsAPIService
	symbol      *string
	id          *string
	updateSpeed *string
}

// Symbol to query
func (r ApiDiffBookDepthRequest) Symbol(symbol string) ApiDiffBookDepthRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiDiffBookDepthRequest) Id(id string) ApiDiffBookDepthRequest {
	r.id = &id
	return r
}

// 1000ms or 100ms
func (r ApiDiffBookDepthRequest) UpdateSpeed(updateSpeed string) ApiDiffBookDepthRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiDiffBookDepthRequest) Execute() (*common.StreamHandler[models.DiffBookDepthResponse], error) {
	return r.ApiService.DiffBookDepthExecute(r)
}

/*
DiffBookDepth WebSocket Diff. Depth Stream
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#diff-depth-stream

@param symbol Symbol to query	@param id Unique WebSocket request ID.	@param updateSpeed 1000ms or 100ms
@return ApiDiffBookDepthRequest
*/
func (a *WebSocketStreamsAPIService) DiffBookDepth() ApiDiffBookDepthRequest {
	return ApiDiffBookDepthRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DiffBookDepthResponse
func (a *WebSocketStreamsAPIService) DiffBookDepthExecute(r ApiDiffBookDepthRequest) (*common.StreamHandler[models.DiffBookDepthResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth@<updateSpeed>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return *r.updateSpeed
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.DiffBookDepthResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	interval   *models.KlineIntervalParameter
	id         *string
}

// Symbol to query
func (r ApiKlineRequest) Symbol(symbol string) ApiKlineRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlineRequest) Interval(interval models.KlineIntervalParameter) ApiKlineRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiKlineRequest) Id(id string) ApiKlineRequest {
	r.id = &id
	return r
}

func (r ApiKlineRequest) Execute() (*common.StreamHandler[models.KlineResponse], error) {
	return r.ApiService.KlineExecute(r)
}

/*
Kline WebSocket Kline/Candlestick Streams for UTC
/<symbol>@kline_<interval>

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#klinecandlestick-streams-for-utc

@param symbol Symbol to query	@param interval	@param id Unique WebSocket request ID.
@return ApiKlineRequest
*/
func (a *WebSocketStreamsAPIService) Kline() ApiKlineRequest {
	return ApiKlineRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineResponse
func (a *WebSocketStreamsAPIService) KlineExecute(r ApiKlineRequest) (*common.StreamHandler[models.KlineResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@kline_<interval>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"interval": func() string {
				if r.interval == nil {
					return ""
				}
				return string(*r.interval)
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
	resp, err := common.CreateStreamHandler[models.KlineResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineOffsetRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	interval   *models.KlineIntervalParameter
	id         *string
}

// Symbol to query
func (r ApiKlineOffsetRequest) Symbol(symbol string) ApiKlineOffsetRequest {
	r.symbol = &symbol
	return r
}

func (r ApiKlineOffsetRequest) Interval(interval models.KlineIntervalParameter) ApiKlineOffsetRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiKlineOffsetRequest) Id(id string) ApiKlineOffsetRequest {
	r.id = &id
	return r
}

func (r ApiKlineOffsetRequest) Execute() (*common.StreamHandler[models.KlineOffsetResponse], error) {
	return r.ApiService.KlineOffsetExecute(r)
}

/*
KlineOffset WebSocket Kline/Candlestick Streams with timezone offset
/<symbol>@kline_<interval>@+08:00

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#klinecandlestick-streams-with-timezone-offset

@param symbol Symbol to query	@param interval	@param id Unique WebSocket request ID.
@return ApiKlineOffsetRequest
*/
func (a *WebSocketStreamsAPIService) KlineOffset() ApiKlineOffsetRequest {
	return ApiKlineOffsetRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineOffsetResponse
func (a *WebSocketStreamsAPIService) KlineOffsetExecute(r ApiKlineOffsetRequest) (*common.StreamHandler[models.KlineOffsetResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@kline_<interval>@+08:00"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"interval": func() string {
				if r.interval == nil {
					return ""
				}
				return string(*r.interval)
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
	resp, err := common.CreateStreamHandler[models.KlineOffsetResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMiniTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiMiniTickerRequest) Symbol(symbol string) ApiMiniTickerRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiMiniTickerRequest) Id(id string) ApiMiniTickerRequest {
	r.id = &id
	return r
}

func (r ApiMiniTickerRequest) Execute() (*common.StreamHandler[models.MiniTickerResponse], error) {
	return r.ApiService.MiniTickerExecute(r)
}

/*
MiniTicker WebSocket Individual Symbol Mini Ticker Stream
/<symbol>@miniTicker

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#individual-symbol-mini-ticker-stream

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiMiniTickerRequest
*/
func (a *WebSocketStreamsAPIService) MiniTicker() ApiMiniTickerRequest {
	return ApiMiniTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MiniTickerResponse
func (a *WebSocketStreamsAPIService) MiniTickerExecute(r ApiMiniTickerRequest) (*common.StreamHandler[models.MiniTickerResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@miniTicker"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.MiniTickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthRequest struct {
	ApiService  *WebSocketStreamsAPIService
	symbol      *string
	levels      *models.PartialBookDepthLevelsParameter
	id          *string
	updateSpeed *string
}

// Symbol to query
func (r ApiPartialBookDepthRequest) Symbol(symbol string) ApiPartialBookDepthRequest {
	r.symbol = &symbol
	return r
}

func (r ApiPartialBookDepthRequest) Levels(levels models.PartialBookDepthLevelsParameter) ApiPartialBookDepthRequest {
	r.levels = &levels
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthRequest) Id(id string) ApiPartialBookDepthRequest {
	r.id = &id
	return r
}

// 1000ms or 100ms
func (r ApiPartialBookDepthRequest) UpdateSpeed(updateSpeed string) ApiPartialBookDepthRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiPartialBookDepthRequest) Execute() (*common.StreamHandler[models.PartialBookDepthResponse], error) {
	return r.ApiService.PartialBookDepthExecute(r)
}

/*
PartialBookDepth WebSocket Partial Book Depth Streams
/<symbol>@depth<levels>@<updateSpeed>

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#partial-book-depth-streams

@param symbol Symbol to query	@param levels	@param id Unique WebSocket request ID.	@param updateSpeed 1000ms or 100ms
@return ApiPartialBookDepthRequest
*/
func (a *WebSocketStreamsAPIService) PartialBookDepth() ApiPartialBookDepthRequest {
	return ApiPartialBookDepthRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthResponse
func (a *WebSocketStreamsAPIService) PartialBookDepthExecute(r ApiPartialBookDepthRequest) (*common.StreamHandler[models.PartialBookDepthResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.levels == nil {
		return nil, common.ReportError("levels is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth<levels>@<updateSpeed>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"levels": func() string {
				if r.levels == nil {
					return ""
				}
				return string(*r.levels)
			}(),
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
			"updateSpeed": func() string {
				if r.updateSpeed == nil {
					return ""
				}
				return *r.updateSpeed
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PartialBookDepthResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRollingWindowTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	windowSize *models.AllMarketRollingWindowTickerWindowSizeParameter
	id         *string
}

// Symbol to query
func (r ApiRollingWindowTickerRequest) Symbol(symbol string) ApiRollingWindowTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiRollingWindowTickerRequest) WindowSize(windowSize models.AllMarketRollingWindowTickerWindowSizeParameter) ApiRollingWindowTickerRequest {
	r.windowSize = &windowSize
	return r
}

// Unique WebSocket request ID.
func (r ApiRollingWindowTickerRequest) Id(id string) ApiRollingWindowTickerRequest {
	r.id = &id
	return r
}

func (r ApiRollingWindowTickerRequest) Execute() (*common.StreamHandler[models.RollingWindowTickerResponse], error) {
	return r.ApiService.RollingWindowTickerExecute(r)
}

/*
RollingWindowTicker WebSocket Individual Symbol Rolling Window Statistics Streams
/<symbol>@ticker_<windowSize>

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#individual-symbol-rolling-window-statistics-streams

@param symbol Symbol to query	@param windowSize	@param id Unique WebSocket request ID.
@return ApiRollingWindowTickerRequest
*/
func (a *WebSocketStreamsAPIService) RollingWindowTicker() ApiRollingWindowTickerRequest {
	return ApiRollingWindowTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RollingWindowTickerResponse
func (a *WebSocketStreamsAPIService) RollingWindowTickerExecute(r ApiRollingWindowTickerRequest) (*common.StreamHandler[models.RollingWindowTickerResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.windowSize == nil {
		return nil, common.ReportError("windowSize is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@ticker_<windowSize>"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
			}(),
			"windowSize": func() string {
				if r.windowSize == nil {
					return ""
				}
				return string(*r.windowSize)
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
	resp, err := common.CreateStreamHandler[models.RollingWindowTickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTickerRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiTickerRequest) Symbol(symbol string) ApiTickerRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTickerRequest) Id(id string) ApiTickerRequest {
	r.id = &id
	return r
}

func (r ApiTickerRequest) Execute() (*common.StreamHandler[models.TickerResponse], error) {
	return r.ApiService.TickerExecute(r)
}

/*
Ticker WebSocket Individual Symbol Ticker Streams
/<symbol>@ticker

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#individual-symbol-ticker-streams

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiTickerRequest
*/
func (a *WebSocketStreamsAPIService) Ticker() ApiTickerRequest {
	return ApiTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerResponse
func (a *WebSocketStreamsAPIService) TickerExecute(r ApiTickerRequest) (*common.StreamHandler[models.TickerResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@ticker"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.TickerResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradeRequest struct {
	ApiService *WebSocketStreamsAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiTradeRequest) Symbol(symbol string) ApiTradeRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTradeRequest) Id(id string) ApiTradeRequest {
	r.id = &id
	return r
}

func (r ApiTradeRequest) Execute() (*common.StreamHandler[models.TradeResponse], error) {
	return r.ApiService.TradeExecute(r)
}

/*
Trade WebSocket Trade Streams
/<symbol>@trade

https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#trade-streams

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiTradeRequest
*/
func (a *WebSocketStreamsAPIService) Trade() ApiTradeRequest {
	return ApiTradeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradeResponse
func (a *WebSocketStreamsAPIService) TradeExecute(r ApiTradeRequest) (*common.StreamHandler[models.TradeResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@trade"[1:],
		map[string]string{
			"symbol": func() string {
				if r.symbol == nil {
					return ""
				}
				return *r.symbol
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
	resp, err := common.CreateStreamHandler[models.TradeResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
