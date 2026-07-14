/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package binancespotwebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiAggTradeRequest struct {
	ApiService *DefaultAPIService
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
AggTrade Aggregate Trade Streams
/<symbol>@aggTrade

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#agg-trade

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiAggTradeRequest
*/
func (a *DefaultAPIService) AggTrade() ApiAggTradeRequest {
	return ApiAggTradeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggTradeResponse
func (a *DefaultAPIService) AggTradeExecute(r ApiAggTradeRequest) (*common.StreamHandler[models.AggTradeResponse], error) {
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
	ApiService *DefaultAPIService
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
AllMarketRollingWindowTicker All Market Rolling Window Statistics Streams
/!ticker_<windowSize>@arr

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#all-market-rolling-window-ticker

@param windowSize	@param id Unique WebSocket request ID.
@return ApiAllMarketRollingWindowTickerRequest
*/
func (a *DefaultAPIService) AllMarketRollingWindowTicker() ApiAllMarketRollingWindowTickerRequest {
	return ApiAllMarketRollingWindowTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketRollingWindowTickerResponse
func (a *DefaultAPIService) AllMarketRollingWindowTickerExecute(r ApiAllMarketRollingWindowTickerRequest) (*common.StreamHandler[models.AllMarketRollingWindowTickerResponse], error) {
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
	ApiService *DefaultAPIService
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
AllMiniTicker All Market Mini Tickers Stream
/!miniTicker@arr

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#all-mini-ticker

@param id Unique WebSocket request ID.
@return ApiAllMiniTickerRequest
*/
func (a *DefaultAPIService) AllMiniTicker() ApiAllMiniTickerRequest {
	return ApiAllMiniTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMiniTickerResponse
func (a *DefaultAPIService) AllMiniTickerExecute(r ApiAllMiniTickerRequest) (*common.StreamHandler[models.AllMiniTickerResponse], error) {

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
	ApiService *DefaultAPIService
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
AvgPrice Average Price
/<symbol>@avgPrice

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#avg-price

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiAvgPriceRequest
*/
func (a *DefaultAPIService) AvgPrice() ApiAvgPriceRequest {
	return ApiAvgPriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AvgPriceResponse
func (a *DefaultAPIService) AvgPriceExecute(r ApiAvgPriceRequest) (*common.StreamHandler[models.AvgPriceResponse], error) {
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

type ApiBlockTradeRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiBlockTradeRequest) Symbol(symbol string) ApiBlockTradeRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiBlockTradeRequest) Id(id string) ApiBlockTradeRequest {
	r.id = &id
	return r
}

func (r ApiBlockTradeRequest) Execute() (*common.StreamHandler[models.BlockTradeResponse], error) {
	return r.ApiService.BlockTradeExecute(r)
}

/*
BlockTrade Block Trade Streams
/<symbol>@blockTrade

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#block-trade

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiBlockTradeRequest
*/
func (a *DefaultAPIService) BlockTrade() ApiBlockTradeRequest {
	return ApiBlockTradeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BlockTradeResponse
func (a *DefaultAPIService) BlockTradeExecute(r ApiBlockTradeRequest) (*common.StreamHandler[models.BlockTradeResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@blockTrade"[1:],
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
	resp, err := common.CreateStreamHandler[models.BlockTradeResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiBookTickerRequest struct {
	ApiService *DefaultAPIService
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
BookTicker Individual Symbol Book Ticker Streams
/<symbol>@bookTicker

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#book-ticker

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiBookTickerRequest
*/
func (a *DefaultAPIService) BookTicker() ApiBookTickerRequest {
	return ApiBookTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BookTickerResponse
func (a *DefaultAPIService) BookTickerExecute(r ApiBookTickerRequest) (*common.StreamHandler[models.BookTickerResponse], error) {
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
	ApiService  *DefaultAPIService
	symbol      *string
	id          *string
	updateSpeed *models.DiffBookDepthUpdateSpeedParameter
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

// Optional stream update speed suffix
func (r ApiDiffBookDepthRequest) UpdateSpeed(updateSpeed models.DiffBookDepthUpdateSpeedParameter) ApiDiffBookDepthRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiDiffBookDepthRequest) Execute() (*common.StreamHandler[models.DiffBookDepthResponse], error) {
	return r.ApiService.DiffBookDepthExecute(r)
}

/*
DiffBookDepth Diff. Depth Stream
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#diff-book-depth

@param symbol Symbol to query	@param id Unique WebSocket request ID.	@param updateSpeed Optional stream update speed suffix
@return ApiDiffBookDepthRequest
*/
func (a *DefaultAPIService) DiffBookDepth() ApiDiffBookDepthRequest {
	return ApiDiffBookDepthRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DiffBookDepthResponse
func (a *DefaultAPIService) DiffBookDepthExecute(r ApiDiffBookDepthRequest) (*common.StreamHandler[models.DiffBookDepthResponse], error) {
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
				return string(*r.updateSpeed)
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
	ApiService *DefaultAPIService
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
Kline Kline/Candlestick Streams for UTC
/<symbol>@kline_<interval>

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#kline

@param symbol Symbol to query	@param interval	@param id Unique WebSocket request ID.
@return ApiKlineRequest
*/
func (a *DefaultAPIService) Kline() ApiKlineRequest {
	return ApiKlineRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineResponse
func (a *DefaultAPIService) KlineExecute(r ApiKlineRequest) (*common.StreamHandler[models.KlineResponse], error) {
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
	ApiService *DefaultAPIService
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
KlineOffset Kline/Candlestick Streams with timezone offset
/<symbol>@kline_<interval>@+08:00

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#kline-offset

@param symbol Symbol to query	@param interval	@param id Unique WebSocket request ID.
@return ApiKlineOffsetRequest
*/
func (a *DefaultAPIService) KlineOffset() ApiKlineOffsetRequest {
	return ApiKlineOffsetRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineOffsetResponse
func (a *DefaultAPIService) KlineOffsetExecute(r ApiKlineOffsetRequest) (*common.StreamHandler[models.KlineOffsetResponse], error) {
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
	ApiService *DefaultAPIService
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
MiniTicker Individual Symbol Mini Ticker Stream
/<symbol>@miniTicker

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#mini-ticker

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiMiniTickerRequest
*/
func (a *DefaultAPIService) MiniTicker() ApiMiniTickerRequest {
	return ApiMiniTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MiniTickerResponse
func (a *DefaultAPIService) MiniTickerExecute(r ApiMiniTickerRequest) (*common.StreamHandler[models.MiniTickerResponse], error) {
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
	ApiService  *DefaultAPIService
	symbol      *string
	levels      *models.PartialBookDepthLevelsParameter
	id          *string
	updateSpeed *models.DiffBookDepthUpdateSpeedParameter
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

// Optional stream update speed suffix
func (r ApiPartialBookDepthRequest) UpdateSpeed(updateSpeed models.DiffBookDepthUpdateSpeedParameter) ApiPartialBookDepthRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiPartialBookDepthRequest) Execute() (*common.StreamHandler[models.PartialBookDepthResponse], error) {
	return r.ApiService.PartialBookDepthExecute(r)
}

/*
PartialBookDepth WebSocket Partial Book Depth Streams
/<symbol>@depth<levels>@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#partial-book-depth

@param symbol Symbol to query	@param levels	@param id Unique WebSocket request ID.	@param updateSpeed Optional stream update speed suffix
@return ApiPartialBookDepthRequest
*/
func (a *DefaultAPIService) PartialBookDepth() ApiPartialBookDepthRequest {
	return ApiPartialBookDepthRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthResponse
func (a *DefaultAPIService) PartialBookDepthExecute(r ApiPartialBookDepthRequest) (*common.StreamHandler[models.PartialBookDepthResponse], error) {
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
				return string(*r.updateSpeed)
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

type ApiReferencePriceRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *string
}

// Symbol to query
func (r ApiReferencePriceRequest) Symbol(symbol string) ApiReferencePriceRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiReferencePriceRequest) Id(id string) ApiReferencePriceRequest {
	r.id = &id
	return r
}

func (r ApiReferencePriceRequest) Execute() (*common.StreamHandler[models.ReferencePriceResponse], error) {
	return r.ApiService.ReferencePriceExecute(r)
}

/*
ReferencePrice Reference Price Streams
/<symbol>@referencePrice

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#reference-price

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiReferencePriceRequest
*/
func (a *DefaultAPIService) ReferencePrice() ApiReferencePriceRequest {
	return ApiReferencePriceRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ReferencePriceResponse
func (a *DefaultAPIService) ReferencePriceExecute(r ApiReferencePriceRequest) (*common.StreamHandler[models.ReferencePriceResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@referencePrice"[1:],
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
	resp, err := common.CreateStreamHandler[models.ReferencePriceResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRollingWindowTickerRequest struct {
	ApiService *DefaultAPIService
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
RollingWindowTicker Individual Symbol Rolling Window Statistics Streams
/<symbol>@ticker_<windowSize>

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#rolling-window-ticker

@param symbol Symbol to query	@param windowSize	@param id Unique WebSocket request ID.
@return ApiRollingWindowTickerRequest
*/
func (a *DefaultAPIService) RollingWindowTicker() ApiRollingWindowTickerRequest {
	return ApiRollingWindowTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RollingWindowTickerResponse
func (a *DefaultAPIService) RollingWindowTickerExecute(r ApiRollingWindowTickerRequest) (*common.StreamHandler[models.RollingWindowTickerResponse], error) {
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
	ApiService *DefaultAPIService
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
Ticker Individual Symbol Ticker Streams
/<symbol>@ticker

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#ticker

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiTickerRequest
*/
func (a *DefaultAPIService) Ticker() ApiTickerRequest {
	return ApiTickerRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerResponse
func (a *DefaultAPIService) TickerExecute(r ApiTickerRequest) (*common.StreamHandler[models.TickerResponse], error) {
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
	ApiService *DefaultAPIService
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
Trade Trade Streams
/<symbol>@trade

https://developers.binance.com/en/docs/catalog/core-trading-spot-trading/api/ws-streams/~#trade

@param symbol Symbol to query	@param id Unique WebSocket request ID.
@return ApiTradeRequest
*/
func (a *DefaultAPIService) Trade() ApiTradeRequest {
	return ApiTradeRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradeResponse
func (a *DefaultAPIService) TradeExecute(r ApiTradeRequest) (*common.StreamHandler[models.TradeResponse], error) {
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
