/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package binancederivativestradingusdsfutureswebsocketstreams

import (
	"strconv"

	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// WebsocketMarketStreamsAPIService WebsocketMarketStreamsAPI Service
type WebsocketMarketStreamsAPIService Service

type ApiAggregateTradeStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiAggregateTradeStreamsRequest) Symbol(symbol string) ApiAggregateTradeStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiAggregateTradeStreamsRequest) Id(id string) ApiAggregateTradeStreamsRequest {
	r.id = &id
	return r
}

func (r ApiAggregateTradeStreamsRequest) Execute() (*common.StreamHandler[models.AggregateTradeStreamsResponse], error) {
	return r.ApiService.AggregateTradeStreamsExecute(r)
}

/*
AggregateTradeStreams Aggregate Trade Streams
/<symbol>@aggTrade

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Aggregate-Trade-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiAggregateTradeStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) AggregateTradeStreams() ApiAggregateTradeStreamsRequest {
	return ApiAggregateTradeStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggregateTradeStreamsResponse
func (a *WebsocketMarketStreamsAPIService) AggregateTradeStreamsExecute(r ApiAggregateTradeStreamsRequest) (*common.StreamHandler[models.AggregateTradeStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.AggregateTradeStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllBookTickersStreamRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllBookTickersStreamRequest) Id(id string) ApiAllBookTickersStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllBookTickersStreamRequest) Execute() (*common.StreamHandler[models.AllBookTickersStreamResponse], error) {
	return r.ApiService.AllBookTickersStreamExecute(r)
}

/*
AllBookTickersStream All Book Tickers Stream
/!bookTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/All-Book-Tickers-Stream

@param id Unique WebSocket request ID.
@return ApiAllBookTickersStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) AllBookTickersStream() ApiAllBookTickersStreamRequest {
	return ApiAllBookTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllBookTickersStreamResponse
func (a *WebsocketMarketStreamsAPIService) AllBookTickersStreamExecute(r ApiAllBookTickersStreamRequest) (*common.StreamHandler[models.AllBookTickersStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!bookTicker"[1:],
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
	resp, err := common.CreateStreamHandler[models.AllBookTickersStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMarketLiquidationOrderStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllMarketLiquidationOrderStreamsRequest) Id(id string) ApiAllMarketLiquidationOrderStreamsRequest {
	r.id = &id
	return r
}

func (r ApiAllMarketLiquidationOrderStreamsRequest) Execute() (*common.StreamHandler[models.AllMarketLiquidationOrderStreamsResponse], error) {
	return r.ApiService.AllMarketLiquidationOrderStreamsExecute(r)
}

/*
AllMarketLiquidationOrderStreams All Market Liquidation Order Streams
/!forceOrder@arr

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/All-Market-Liquidation-Order-Streams

@param id Unique WebSocket request ID.
@return ApiAllMarketLiquidationOrderStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) AllMarketLiquidationOrderStreams() ApiAllMarketLiquidationOrderStreamsRequest {
	return ApiAllMarketLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketLiquidationOrderStreamsResponse
func (a *WebsocketMarketStreamsAPIService) AllMarketLiquidationOrderStreamsExecute(r ApiAllMarketLiquidationOrderStreamsRequest) (*common.StreamHandler[models.AllMarketLiquidationOrderStreamsResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!forceOrder@arr"[1:],
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
	resp, err := common.CreateStreamHandler[models.AllMarketLiquidationOrderStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMarketMiniTickersStreamRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllMarketMiniTickersStreamRequest) Id(id string) ApiAllMarketMiniTickersStreamRequest {
	r.id = &id
	return r
}

func (r ApiAllMarketMiniTickersStreamRequest) Execute() (*common.StreamHandler[models.AllMarketMiniTickersStreamResponse], error) {
	return r.ApiService.AllMarketMiniTickersStreamExecute(r)
}

/*
AllMarketMiniTickersStream All Market Mini Tickers Stream
/!miniTicker@arr

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/All-Market-Mini-Tickers-Stream

@param id Unique WebSocket request ID.
@return ApiAllMarketMiniTickersStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) AllMarketMiniTickersStream() ApiAllMarketMiniTickersStreamRequest {
	return ApiAllMarketMiniTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketMiniTickersStreamResponse
func (a *WebsocketMarketStreamsAPIService) AllMarketMiniTickersStreamExecute(r ApiAllMarketMiniTickersStreamRequest) (*common.StreamHandler[models.AllMarketMiniTickersStreamResponse], error) {

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
	resp, err := common.CreateStreamHandler[models.AllMarketMiniTickersStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiAllMarketTickersStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAllMarketTickersStreamsRequest) Id(id string) ApiAllMarketTickersStreamsRequest {
	r.id = &id
	return r
}

func (r ApiAllMarketTickersStreamsRequest) Execute() (*common.StreamHandler[models.AllMarketTickersStreamsResponse], error) {
	return r.ApiService.AllMarketTickersStreamsExecute(r)
}

/*
AllMarketTickersStreams All Market Tickers Streams
/!ticker@arr

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/All-Market-Tickers-Streams

@param id Unique WebSocket request ID.
@return ApiAllMarketTickersStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) AllMarketTickersStreams() ApiAllMarketTickersStreamsRequest {
	return ApiAllMarketTickersStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketTickersStreamsResponse
func (a *WebsocketMarketStreamsAPIService) AllMarketTickersStreamsExecute(r ApiAllMarketTickersStreamsRequest) (*common.StreamHandler[models.AllMarketTickersStreamsResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!ticker@arr"[1:],
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
	resp, err := common.CreateStreamHandler[models.AllMarketTickersStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiCompositeIndexSymbolInformationStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiCompositeIndexSymbolInformationStreamsRequest) Symbol(symbol string) ApiCompositeIndexSymbolInformationStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiCompositeIndexSymbolInformationStreamsRequest) Id(id string) ApiCompositeIndexSymbolInformationStreamsRequest {
	r.id = &id
	return r
}

func (r ApiCompositeIndexSymbolInformationStreamsRequest) Execute() (*common.StreamHandler[models.CompositeIndexSymbolInformationStreamsResponse], error) {
	return r.ApiService.CompositeIndexSymbolInformationStreamsExecute(r)
}

/*
CompositeIndexSymbolInformationStreams Composite Index Symbol Information Streams
/<symbol>@compositeIndex

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Composite-Index-Symbol-Information-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiCompositeIndexSymbolInformationStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) CompositeIndexSymbolInformationStreams() ApiCompositeIndexSymbolInformationStreamsRequest {
	return ApiCompositeIndexSymbolInformationStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CompositeIndexSymbolInformationStreamsResponse
func (a *WebsocketMarketStreamsAPIService) CompositeIndexSymbolInformationStreamsExecute(r ApiCompositeIndexSymbolInformationStreamsRequest) (*common.StreamHandler[models.CompositeIndexSymbolInformationStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@compositeIndex"[1:],
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
	resp, err := common.CreateStreamHandler[models.CompositeIndexSymbolInformationStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiContinuousContractKlineCandlestickStreamsRequest struct {
	ApiService   *WebsocketMarketStreamsAPIService
	pair         *string
	contractType *string
	interval     *string
	id           *string
}

// The pair parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) Pair(pair string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.pair = &pair
	return r
}

// The contractType parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) ContractType(contractType string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.contractType = &contractType
	return r
}

// The interval parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) Interval(interval string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiContinuousContractKlineCandlestickStreamsRequest) Id(id string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiContinuousContractKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse], error) {
	return r.ApiService.ContinuousContractKlineCandlestickStreamsExecute(r)
}

/*
ContinuousContractKlineCandlestickStreams Continuous Contract Kline/Candlestick Streams
/<pair>_<contractType>@continuousKline_<interval>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Continuous-Contract-Kline-Candlestick-Streams

@param pair The pair parameter	@param contractType The contractType parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiContinuousContractKlineCandlestickStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) ContinuousContractKlineCandlestickStreams() ApiContinuousContractKlineCandlestickStreamsRequest {
	return ApiContinuousContractKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContinuousContractKlineCandlestickStreamsResponse
func (a *WebsocketMarketStreamsAPIService) ContinuousContractKlineCandlestickStreamsExecute(r ApiContinuousContractKlineCandlestickStreamsRequest) (*common.StreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse], error) {
	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.contractType == nil {
		return nil, common.ReportError("contractType is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<pair>_<contractType>@continuousKline_<interval>"[1:],
		map[string]string{
			"pair": func() string {
				if r.pair == nil {
					return ""
				}
				return *r.pair
			}(),
			"contractType": func() string {
				if r.contractType == nil {
					return ""
				}
				return *r.contractType
			}(),
			"interval": func() string {
				if r.interval == nil {
					return ""
				}
				return *r.interval
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
	resp, err := common.CreateStreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiContractInfoStreamRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiContractInfoStreamRequest) Id(id string) ApiContractInfoStreamRequest {
	r.id = &id
	return r
}

func (r ApiContractInfoStreamRequest) Execute() (*common.StreamHandler[models.ContractInfoStreamResponse], error) {
	return r.ApiService.ContractInfoStreamExecute(r)
}

/*
ContractInfoStream Contract Info Stream
/!contractInfo

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Contract-Info-Stream

@param id Unique WebSocket request ID.
@return ApiContractInfoStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) ContractInfoStream() ApiContractInfoStreamRequest {
	return ApiContractInfoStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContractInfoStreamResponse
func (a *WebsocketMarketStreamsAPIService) ContractInfoStreamExecute(r ApiContractInfoStreamRequest) (*common.StreamHandler[models.ContractInfoStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!contractInfo"[1:],
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
	resp, err := common.CreateStreamHandler[models.ContractInfoStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiDiffBookDepthStreamsRequest struct {
	ApiService  *WebsocketMarketStreamsAPIService
	symbol      *string
	id          *string
	updateSpeed *string
}

// The symbol parameter
func (r ApiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiDiffBookDepthStreamsRequest) Id(id string) ApiDiffBookDepthStreamsRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiDiffBookDepthStreamsRequest) UpdateSpeed(updateSpeed string) ApiDiffBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
	return r.ApiService.DiffBookDepthStreamsExecute(r)
}

/*
DiffBookDepthStreams Diff. Book Depth Streams
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Diff-Book-Depth-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiDiffBookDepthStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) DiffBookDepthStreams() ApiDiffBookDepthStreamsRequest {
	return ApiDiffBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DiffBookDepthStreamsResponse
func (a *WebsocketMarketStreamsAPIService) DiffBookDepthStreamsExecute(r ApiDiffBookDepthStreamsRequest) (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.DiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolBookTickerStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiIndividualSymbolBookTickerStreamsRequest) Symbol(symbol string) ApiIndividualSymbolBookTickerStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndividualSymbolBookTickerStreamsRequest) Id(id string) ApiIndividualSymbolBookTickerStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndividualSymbolBookTickerStreamsRequest) Execute() (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
	return r.ApiService.IndividualSymbolBookTickerStreamsExecute(r)
}

/*
IndividualSymbolBookTickerStreams Individual Symbol Book Ticker Streams
/<symbol>@bookTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Individual-Symbol-Book-Ticker-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolBookTickerStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolBookTickerStreams() ApiIndividualSymbolBookTickerStreamsRequest {
	return ApiIndividualSymbolBookTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolBookTickerStreamsResponse
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolBookTickerStreamsExecute(r ApiIndividualSymbolBookTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.IndividualSymbolBookTickerStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolMiniTickerStreamRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiIndividualSymbolMiniTickerStreamRequest) Symbol(symbol string) ApiIndividualSymbolMiniTickerStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndividualSymbolMiniTickerStreamRequest) Id(id string) ApiIndividualSymbolMiniTickerStreamRequest {
	r.id = &id
	return r
}

func (r ApiIndividualSymbolMiniTickerStreamRequest) Execute() (*common.StreamHandler[models.IndividualSymbolMiniTickerStreamResponse], error) {
	return r.ApiService.IndividualSymbolMiniTickerStreamExecute(r)
}

/*
IndividualSymbolMiniTickerStream Individual Symbol Mini Ticker Stream
/<symbol>@miniTicker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Individual-Symbol-Mini-Ticker-Stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolMiniTickerStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolMiniTickerStream() ApiIndividualSymbolMiniTickerStreamRequest {
	return ApiIndividualSymbolMiniTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolMiniTickerStreamResponse
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolMiniTickerStreamExecute(r ApiIndividualSymbolMiniTickerStreamRequest) (*common.StreamHandler[models.IndividualSymbolMiniTickerStreamResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.IndividualSymbolMiniTickerStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolTickerStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiIndividualSymbolTickerStreamsRequest) Symbol(symbol string) ApiIndividualSymbolTickerStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiIndividualSymbolTickerStreamsRequest) Id(id string) ApiIndividualSymbolTickerStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndividualSymbolTickerStreamsRequest) Execute() (*common.StreamHandler[models.IndividualSymbolTickerStreamsResponse], error) {
	return r.ApiService.IndividualSymbolTickerStreamsExecute(r)
}

/*
IndividualSymbolTickerStreams Individual Symbol Ticker Streams
/<symbol>@ticker

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Individual-Symbol-Ticker-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolTickerStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolTickerStreams() ApiIndividualSymbolTickerStreamsRequest {
	return ApiIndividualSymbolTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolTickerStreamsResponse
func (a *WebsocketMarketStreamsAPIService) IndividualSymbolTickerStreamsExecute(r ApiIndividualSymbolTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolTickerStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.IndividualSymbolTickerStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiKlineCandlestickStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	interval   *string
	id         *string
}

// The symbol parameter
func (r ApiKlineCandlestickStreamsRequest) Symbol(symbol string) ApiKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiKlineCandlestickStreamsRequest) Interval(interval string) ApiKlineCandlestickStreamsRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiKlineCandlestickStreamsRequest) Id(id string) ApiKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
	return r.ApiService.KlineCandlestickStreamsExecute(r)
}

/*
KlineCandlestickStreams Kline/Candlestick Streams
/<symbol>@kline_<interval>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Kline-Candlestick-Streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiKlineCandlestickStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) KlineCandlestickStreams() ApiKlineCandlestickStreamsRequest {
	return ApiKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickStreamsResponse
func (a *WebsocketMarketStreamsAPIService) KlineCandlestickStreamsExecute(r ApiKlineCandlestickStreamsRequest) (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
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
				return *r.interval
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
	resp, err := common.CreateStreamHandler[models.KlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiLiquidationOrderStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiLiquidationOrderStreamsRequest) Symbol(symbol string) ApiLiquidationOrderStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiLiquidationOrderStreamsRequest) Id(id string) ApiLiquidationOrderStreamsRequest {
	r.id = &id
	return r
}

func (r ApiLiquidationOrderStreamsRequest) Execute() (*common.StreamHandler[models.LiquidationOrderStreamsResponse], error) {
	return r.ApiService.LiquidationOrderStreamsExecute(r)
}

/*
LiquidationOrderStreams Liquidation Order Streams
/<symbol>@forceOrder

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Liquidation-Order-Streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiLiquidationOrderStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) LiquidationOrderStreams() ApiLiquidationOrderStreamsRequest {
	return ApiLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LiquidationOrderStreamsResponse
func (a *WebsocketMarketStreamsAPIService) LiquidationOrderStreamsExecute(r ApiLiquidationOrderStreamsRequest) (*common.StreamHandler[models.LiquidationOrderStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@forceOrder"[1:],
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
	resp, err := common.CreateStreamHandler[models.LiquidationOrderStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceStreamRequest struct {
	ApiService  *WebsocketMarketStreamsAPIService
	symbol      *string
	id          *string
	updateSpeed *string
}

// The symbol parameter
func (r ApiMarkPriceStreamRequest) Symbol(symbol string) ApiMarkPriceStreamRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiMarkPriceStreamRequest) Id(id string) ApiMarkPriceStreamRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiMarkPriceStreamRequest) UpdateSpeed(updateSpeed string) ApiMarkPriceStreamRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceStreamRequest) Execute() (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
	return r.ApiService.MarkPriceStreamExecute(r)
}

/*
MarkPriceStream Mark Price Stream
/<symbol>@markPrice@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Mark-Price-Stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) MarkPriceStream() ApiMarkPriceStreamRequest {
	return ApiMarkPriceStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceStreamResponse
func (a *WebsocketMarketStreamsAPIService) MarkPriceStreamExecute(r ApiMarkPriceStreamRequest) (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@markPrice@<updateSpeed>"[1:],
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
	resp, err := common.CreateStreamHandler[models.MarkPriceStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceStreamForAllMarketRequest struct {
	ApiService  *WebsocketMarketStreamsAPIService
	id          *string
	updateSpeed *string
}

// Unique WebSocket request ID.
func (r ApiMarkPriceStreamForAllMarketRequest) Id(id string) ApiMarkPriceStreamForAllMarketRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiMarkPriceStreamForAllMarketRequest) UpdateSpeed(updateSpeed string) ApiMarkPriceStreamForAllMarketRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceStreamForAllMarketRequest) Execute() (*common.StreamHandler[models.MarkPriceStreamForAllMarketResponse], error) {
	return r.ApiService.MarkPriceStreamForAllMarketExecute(r)
}

/*
MarkPriceStreamForAllMarket Mark Price Stream for All market
/!markPrice@arr@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Mark-Price-Stream-for-All-market

@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceStreamForAllMarketRequest
*/
func (a *WebsocketMarketStreamsAPIService) MarkPriceStreamForAllMarket() ApiMarkPriceStreamForAllMarketRequest {
	return ApiMarkPriceStreamForAllMarketRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceStreamForAllMarketResponse
func (a *WebsocketMarketStreamsAPIService) MarkPriceStreamForAllMarketExecute(r ApiMarkPriceStreamForAllMarketRequest) (*common.StreamHandler[models.MarkPriceStreamForAllMarketResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!markPrice@arr@<updateSpeed>"[1:],
		map[string]string{
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
	resp, err := common.CreateStreamHandler[models.MarkPriceStreamForAllMarketResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMultiAssetsModeAssetIndexRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiMultiAssetsModeAssetIndexRequest) Id(id string) ApiMultiAssetsModeAssetIndexRequest {
	r.id = &id
	return r
}

func (r ApiMultiAssetsModeAssetIndexRequest) Execute() (*common.StreamHandler[models.MultiAssetsModeAssetIndexResponse], error) {
	return r.ApiService.MultiAssetsModeAssetIndexExecute(r)
}

/*
MultiAssetsModeAssetIndex Multi-Assets Mode Asset Index
/!assetIndex@arr

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Multi-Assets-Mode-Asset-Index

@param id Unique WebSocket request ID.
@return ApiMultiAssetsModeAssetIndexRequest
*/
func (a *WebsocketMarketStreamsAPIService) MultiAssetsModeAssetIndex() ApiMultiAssetsModeAssetIndexRequest {
	return ApiMultiAssetsModeAssetIndexRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MultiAssetsModeAssetIndexResponse
func (a *WebsocketMarketStreamsAPIService) MultiAssetsModeAssetIndexExecute(r ApiMultiAssetsModeAssetIndexRequest) (*common.StreamHandler[models.MultiAssetsModeAssetIndexResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/!assetIndex@arr"[1:],
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
	resp, err := common.CreateStreamHandler[models.MultiAssetsModeAssetIndexResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthStreamsRequest struct {
	ApiService  *WebsocketMarketStreamsAPIService
	symbol      *string
	levels      *int64
	id          *string
	updateSpeed *string
}

// The symbol parameter
func (r ApiPartialBookDepthStreamsRequest) Symbol(symbol string) ApiPartialBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// The levels parameter
func (r ApiPartialBookDepthStreamsRequest) Levels(levels int64) ApiPartialBookDepthStreamsRequest {
	r.levels = &levels
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthStreamsRequest) Id(id string) ApiPartialBookDepthStreamsRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiPartialBookDepthStreamsRequest) UpdateSpeed(updateSpeed string) ApiPartialBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiPartialBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
	return r.ApiService.PartialBookDepthStreamsExecute(r)
}

/*
PartialBookDepthStreams Partial Book Depth Streams
/<symbol>@depth<levels>@<updateSpeed>

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Partial-Book-Depth-Streams

@param symbol The symbol parameter	@param levels The levels parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiPartialBookDepthStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) PartialBookDepthStreams() ApiPartialBookDepthStreamsRequest {
	return ApiPartialBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthStreamsResponse
func (a *WebsocketMarketStreamsAPIService) PartialBookDepthStreamsExecute(r ApiPartialBookDepthStreamsRequest) (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
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
				return strconv.FormatInt(*r.levels, 10)
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
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiRpiDiffBookDepthStreamsRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiRpiDiffBookDepthStreamsRequest) Symbol(symbol string) ApiRpiDiffBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiRpiDiffBookDepthStreamsRequest) Id(id string) ApiRpiDiffBookDepthStreamsRequest {
	r.id = &id
	return r
}

func (r ApiRpiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.RpiDiffBookDepthStreamsResponse], error) {
	return r.ApiService.RpiDiffBookDepthStreamsExecute(r)
}

/*
RpiDiffBookDepthStreams RPI Diff. Book Depth Streams
/<symbol>@rpiDepth@500ms

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Diff-Book-Depth-Streams-RPI

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiRpiDiffBookDepthStreamsRequest
*/
func (a *WebsocketMarketStreamsAPIService) RpiDiffBookDepthStreams() ApiRpiDiffBookDepthStreamsRequest {
	return ApiRpiDiffBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return RpiDiffBookDepthStreamsResponse
func (a *WebsocketMarketStreamsAPIService) RpiDiffBookDepthStreamsExecute(r ApiRpiDiffBookDepthStreamsRequest) (*common.StreamHandler[models.RpiDiffBookDepthStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@rpiDepth@500ms"[1:],
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
	resp, err := common.CreateStreamHandler[models.RpiDiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiTradingSessionStreamRequest struct {
	ApiService *WebsocketMarketStreamsAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiTradingSessionStreamRequest) Id(id string) ApiTradingSessionStreamRequest {
	r.id = &id
	return r
}

func (r ApiTradingSessionStreamRequest) Execute() (*common.StreamHandler[models.TradingSessionStreamResponse], error) {
	return r.ApiService.TradingSessionStreamExecute(r)
}

/*
TradingSessionStream Trading Session Stream
/tradingSession

https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Trading-Session-Stream

@param id Unique WebSocket request ID.
@return ApiTradingSessionStreamRequest
*/
func (a *WebsocketMarketStreamsAPIService) TradingSessionStream() ApiTradingSessionStreamRequest {
	return ApiTradingSessionStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradingSessionStreamResponse
func (a *WebsocketMarketStreamsAPIService) TradingSessionStreamExecute(r ApiTradingSessionStreamRequest) (*common.StreamHandler[models.TradingSessionStreamResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/tradingSession"[1:],
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
	resp, err := common.CreateStreamHandler[models.TradingSessionStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
