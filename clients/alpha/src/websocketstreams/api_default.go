/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package binancealphawebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/alpha/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiAggregateTradeStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiAggregateTradeStreamRequest) Symbol(symbol string) ApiAggregateTradeStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiAggregateTradeStreamRequest) Id(id int32) ApiAggregateTradeStreamRequest {
	r.id = &id
	return r
}

func (r ApiAggregateTradeStreamRequest) Execute() (*common.StreamHandler[models.AggregateTradeStreamResponse], error) {
	return r.ApiService.AggregateTradeStreamExecute(r)
}

/*
AggregateTradeStream Aggregate Trade Stream
/<symbol>@aggTrade

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#aggregate-trade-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param id Unique WebSocket request ID.
@return ApiAggregateTradeStreamRequest
*/
func (a *DefaultAPIService) AggregateTradeStream() ApiAggregateTradeStreamRequest {
	return ApiAggregateTradeStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggregateTradeStreamResponse
func (a *DefaultAPIService) AggregateTradeStreamExecute(r ApiAggregateTradeStreamRequest) (*common.StreamHandler[models.AggregateTradeStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AggregateTradeStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllBookTickerStreamRequest struct {
	ApiService *DefaultAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiAllBookTickerStreamRequest) Id(id int32) ApiAllBookTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllBookTickerStreamRequest) Execute() (*common.StreamHandler[models.AllBookTickerStreamResponse], error) {
	return r.ApiService.AllBookTickerStreamExecute(r)
}

/*
AllBookTickerStream All Book Ticker Stream
/!bookTicker

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#all-book-ticker-stream

@param id Unique WebSocket request ID.
@return ApiAllBookTickerStreamRequest
*/
func (a *DefaultAPIService) AllBookTickerStream() ApiAllBookTickerStreamRequest {
	return ApiAllBookTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllBookTickerStreamResponse
func (a *DefaultAPIService) AllBookTickerStreamExecute(r ApiAllBookTickerStreamRequest) (*common.StreamHandler[models.AllBookTickerStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!bookTicker"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AllBookTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMiniTickerStreamRequest struct {
	ApiService *DefaultAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiAllMiniTickerStreamRequest) Id(id int32) ApiAllMiniTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllMiniTickerStreamRequest) Execute() (*common.StreamHandler[models.AllMiniTickerStreamResponse], error) {
	return r.ApiService.AllMiniTickerStreamExecute(r)
}

/*
AllMiniTickerStream All Mini Ticker Stream
/!miniTicker@arr

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#all-mini-ticker-stream

@param id Unique WebSocket request ID.
@return ApiAllMiniTickerStreamRequest
*/
func (a *DefaultAPIService) AllMiniTickerStream() ApiAllMiniTickerStreamRequest {
	return ApiAllMiniTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMiniTickerStreamResponse
func (a *DefaultAPIService) AllMiniTickerStreamExecute(r ApiAllMiniTickerStreamRequest) (*common.StreamHandler[models.AllMiniTickerStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!miniTicker@arr"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AllMiniTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllTickerStreamRequest struct {
	ApiService *DefaultAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiAllTickerStreamRequest) Id(id int32) ApiAllTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllTickerStreamRequest) Execute() (*common.StreamHandler[models.AllTickerStreamResponse], error) {
	return r.ApiService.AllTickerStreamExecute(r)
}

/*
AllTickerStream All Ticker Stream
/!ticker@arr

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#all-ticker-stream

@param id Unique WebSocket request ID.
@return ApiAllTickerStreamRequest
*/
func (a *DefaultAPIService) AllTickerStream() ApiAllTickerStreamRequest {
	return ApiAllTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllTickerStreamResponse
func (a *DefaultAPIService) AllTickerStreamExecute(r ApiAllTickerStreamRequest) (*common.StreamHandler[models.AllTickerStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!ticker@arr"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AllTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllTokens24hTickerStreamRequest struct {
	ApiService *DefaultAPIService
	id         *int32
}

// Unique WebSocket request ID.
func (r ApiAllTokens24hTickerStreamRequest) Id(id int32) ApiAllTokens24hTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllTokens24hTickerStreamRequest) Execute() (*common.StreamHandler[models.AllTokens24hTickerStreamResponse], error) {
	return r.ApiService.AllTokens24hTickerStreamExecute(r)
}

/*
AllTokens24hTickerStream All Tokens 24h Ticker Stream
/came@allTokens@ticker24

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#all-tokens24h-ticker-stream

@param id Unique WebSocket request ID.
@return ApiAllTokens24hTickerStreamRequest
*/
func (a *DefaultAPIService) AllTokens24hTickerStream() ApiAllTokens24hTickerStreamRequest {
	return ApiAllTokens24hTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllTokens24hTickerStreamResponse
func (a *DefaultAPIService) AllTokens24hTickerStreamExecute(r ApiAllTokens24hTickerStreamRequest) (*common.StreamHandler[models.AllTokens24hTickerStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/came@allTokens@ticker24"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AllTokens24hTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiBookTickerStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiBookTickerStreamRequest) Symbol(symbol string) ApiBookTickerStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiBookTickerStreamRequest) Id(id int32) ApiBookTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiBookTickerStreamRequest) Execute() (*common.StreamHandler[models.BookTickerStreamResponse], error) {
	return r.ApiService.BookTickerStreamExecute(r)
}

/*
BookTickerStream Book Ticker Stream
/<symbol>@bookTicker

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#book-ticker-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param id Unique WebSocket request ID.
@return ApiBookTickerStreamRequest
*/
func (a *DefaultAPIService) BookTickerStream() ApiBookTickerStreamRequest {
	return ApiBookTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return BookTickerStreamResponse
func (a *DefaultAPIService) BookTickerStreamExecute(r ApiBookTickerStreamRequest) (*common.StreamHandler[models.BookTickerStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.BookTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiContractKlineStreamRequest struct {
	ApiService      *DefaultAPIService
	contractAddress *string
	chainId         *string
	interval        *models.ContractKlineStreamIntervalParameter
	id              *int32
}

// Contract address.
func (r ApiContractKlineStreamRequest) ContractAddress(contractAddress string) ApiContractKlineStreamRequest {
	r.contractAddress = &contractAddress
	return r
}

// Chain ID.
func (r ApiContractKlineStreamRequest) ChainId(chainId string) ApiContractKlineStreamRequest {
	r.chainId = &chainId
	return r
}

// Kline interval.
func (r ApiContractKlineStreamRequest) Interval(interval models.ContractKlineStreamIntervalParameter) ApiContractKlineStreamRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiContractKlineStreamRequest) Id(id int32) ApiContractKlineStreamRequest {
	r.id = &id
	return r
}

func (r ApiContractKlineStreamRequest) Execute() (*common.StreamHandler[models.ContractKlineStreamResponse], error) {
	return r.ApiService.ContractKlineStreamExecute(r)
}

/*
ContractKlineStream Contract Kline Stream
/came@<contractAddress>@<chainId>@kline_<interval>

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#contract-kline-stream

@param contractAddress Contract address.	@param chainId Chain ID.	@param interval Kline interval.	@param id Unique WebSocket request ID.
@return ApiContractKlineStreamRequest
*/
func (a *DefaultAPIService) ContractKlineStream() ApiContractKlineStreamRequest {
	return ApiContractKlineStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContractKlineStreamResponse
func (a *DefaultAPIService) ContractKlineStreamExecute(r ApiContractKlineStreamRequest) (*common.StreamHandler[models.ContractKlineStreamResponse], error) {
	if r.contractAddress == nil {
		return nil, common.ReportError("contractAddress is required and must be specified")
	}
	if r.chainId == nil {
		return nil, common.ReportError("chainId is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/came@<contractAddress>@<chainId>@kline_<interval>"[1:],
		map[string]string{
			"contractAddress": func() string {
				if r.contractAddress == nil {
					return ""
				}
				return *r.contractAddress
			}(),
			"chainId": func() string {
				if r.chainId == nil {
					return ""
				}
				return *r.chainId
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.ContractKlineStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiFullDepthStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	interval   *models.PartialDepthStreamIntervalParameter
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiFullDepthStreamRequest) Symbol(symbol string) ApiFullDepthStreamRequest {
	r.symbol = &symbol
	return r
}

// Update interval.
func (r ApiFullDepthStreamRequest) Interval(interval models.PartialDepthStreamIntervalParameter) ApiFullDepthStreamRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiFullDepthStreamRequest) Id(id int32) ApiFullDepthStreamRequest {
	r.id = &id
	return r
}

func (r ApiFullDepthStreamRequest) Execute() (*common.StreamHandler[models.FullDepthStreamResponse], error) {
	return r.ApiService.FullDepthStreamExecute(r)
}

/*
FullDepthStream Full Depth Stream
/<symbol>@fulldepth@<interval>

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#full-depth-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param interval Update interval.	@param id Unique WebSocket request ID.
@return ApiFullDepthStreamRequest
*/
func (a *DefaultAPIService) FullDepthStream() ApiFullDepthStreamRequest {
	return ApiFullDepthStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return FullDepthStreamResponse
func (a *DefaultAPIService) FullDepthStreamExecute(r ApiFullDepthStreamRequest) (*common.StreamHandler[models.FullDepthStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@fulldepth@<interval>"[1:],
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.FullDepthStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	interval   *models.KlineStreamIntervalParameter
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiKlineStreamRequest) Symbol(symbol string) ApiKlineStreamRequest {
	r.symbol = &symbol
	return r
}

// Kline interval.
func (r ApiKlineStreamRequest) Interval(interval models.KlineStreamIntervalParameter) ApiKlineStreamRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiKlineStreamRequest) Id(id int32) ApiKlineStreamRequest {
	r.id = &id
	return r
}

func (r ApiKlineStreamRequest) Execute() (*common.StreamHandler[models.KlineStreamResponse], error) {
	return r.ApiService.KlineStreamExecute(r)
}

/*
KlineStream Kline Stream
/<symbol>@kline_<interval>

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#kline-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param interval Kline interval.	@param id Unique WebSocket request ID.
@return ApiKlineStreamRequest
*/
func (a *DefaultAPIService) KlineStream() ApiKlineStreamRequest {
	return ApiKlineStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineStreamResponse
func (a *DefaultAPIService) KlineStreamExecute(r ApiKlineStreamRequest) (*common.StreamHandler[models.KlineStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.KlineStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMiniTickerStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiMiniTickerStreamRequest) Symbol(symbol string) ApiMiniTickerStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiMiniTickerStreamRequest) Id(id int32) ApiMiniTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiMiniTickerStreamRequest) Execute() (*common.StreamHandler[models.MiniTickerStreamResponse], error) {
	return r.ApiService.MiniTickerStreamExecute(r)
}

/*
MiniTickerStream Mini Ticker Stream
/<symbol>@miniTicker

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#mini-ticker-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param id Unique WebSocket request ID.
@return ApiMiniTickerStreamRequest
*/
func (a *DefaultAPIService) MiniTickerStream() ApiMiniTickerStreamRequest {
	return ApiMiniTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MiniTickerStreamResponse
func (a *DefaultAPIService) MiniTickerStreamExecute(r ApiMiniTickerStreamRequest) (*common.StreamHandler[models.MiniTickerStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.MiniTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialDepthStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	levels     *models.PartialDepthStreamLevelsParameter
	interval   *models.PartialDepthStreamIntervalParameter
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiPartialDepthStreamRequest) Symbol(symbol string) ApiPartialDepthStreamRequest {
	r.symbol = &symbol
	return r
}

// Depth levels.
func (r ApiPartialDepthStreamRequest) Levels(levels models.PartialDepthStreamLevelsParameter) ApiPartialDepthStreamRequest {
	r.levels = &levels
	return r
}

// Update interval.
func (r ApiPartialDepthStreamRequest) Interval(interval models.PartialDepthStreamIntervalParameter) ApiPartialDepthStreamRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialDepthStreamRequest) Id(id int32) ApiPartialDepthStreamRequest {
	r.id = &id
	return r
}

func (r ApiPartialDepthStreamRequest) Execute() (*common.StreamHandler[models.PartialDepthStreamResponse], error) {
	return r.ApiService.PartialDepthStreamExecute(r)
}

/*
PartialDepthStream Partial Depth Stream
/<symbol>@depth<levels>@<interval>

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#partial-depth-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param levels Depth levels.	@param interval Update interval.	@param id Unique WebSocket request ID.
@return ApiPartialDepthStreamRequest
*/
func (a *DefaultAPIService) PartialDepthStream() ApiPartialDepthStreamRequest {
	return ApiPartialDepthStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialDepthStreamResponse
func (a *DefaultAPIService) PartialDepthStreamExecute(r ApiPartialDepthStreamRequest) (*common.StreamHandler[models.PartialDepthStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.levels == nil {
		return nil, common.ReportError("levels is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@depth<levels>@<interval>"[1:],
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.PartialDepthStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTickerStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiTickerStreamRequest) Symbol(symbol string) ApiTickerStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTickerStreamRequest) Id(id int32) ApiTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiTickerStreamRequest) Execute() (*common.StreamHandler[models.TickerStreamResponse], error) {
	return r.ApiService.TickerStreamExecute(r)
}

/*
TickerStream Ticker Stream
/<symbol>@ticker

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#ticker-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param id Unique WebSocket request ID.
@return ApiTickerStreamRequest
*/
func (a *DefaultAPIService) TickerStream() ApiTickerStreamRequest {
	return ApiTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TickerStreamResponse
func (a *DefaultAPIService) TickerStreamExecute(r ApiTickerStreamRequest) (*common.StreamHandler[models.TickerStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.TickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradeStreamRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *int32
}

// Symbol to subscribe, in lowercase stream format.
func (r ApiTradeStreamRequest) Symbol(symbol string) ApiTradeStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiTradeStreamRequest) Id(id int32) ApiTradeStreamRequest {
	r.id = &id
	return r
}

func (r ApiTradeStreamRequest) Execute() (*common.StreamHandler[models.TradeStreamResponse], error) {
	return r.ApiService.TradeStreamExecute(r)
}

/*
TradeStream Trade Stream
/<symbol>@trade

https://developers.binance.com/en/docs/catalog/advanced-trading-alpha-trading/api/ws-streams/~#trade-stream

@param symbol Symbol to subscribe, in lowercase stream format.	@param id Unique WebSocket request ID.
@return ApiTradeStreamRequest
*/
func (a *DefaultAPIService) TradeStream() ApiTradeStreamRequest {
	return ApiTradeStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradeStreamResponse
func (a *DefaultAPIService) TradeStreamExecute(r ApiTradeStreamRequest) (*common.StreamHandler[models.TradeStreamResponse], error) {
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
				return string(*r.id)
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateIntUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.TradeStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, true)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
