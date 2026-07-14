/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package binancederivativestradingcoinfutureswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiAggregateTradeStreamsRequest struct {
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#aggregate-trade-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiAggregateTradeStreamsRequest
*/
func (a *DefaultAPIService) AggregateTradeStreams() ApiAggregateTradeStreamsRequest {
	return ApiAggregateTradeStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggregateTradeStreamsResponse
func (a *DefaultAPIService) AggregateTradeStreamsExecute(r ApiAggregateTradeStreamsRequest) (*common.StreamHandler[models.AggregateTradeStreamsResponse], error) {
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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#all-book-tickers-stream

@param id Unique WebSocket request ID.
@return ApiAllBookTickersStreamRequest
*/
func (a *DefaultAPIService) AllBookTickersStream() ApiAllBookTickersStreamRequest {
	return ApiAllBookTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllBookTickersStreamResponse
func (a *DefaultAPIService) AllBookTickersStreamExecute(r ApiAllBookTickersStreamRequest) (*common.StreamHandler[models.AllBookTickersStreamResponse], error) {

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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#all-market-liquidation-order-streams

@param id Unique WebSocket request ID.
@return ApiAllMarketLiquidationOrderStreamsRequest
*/
func (a *DefaultAPIService) AllMarketLiquidationOrderStreams() ApiAllMarketLiquidationOrderStreamsRequest {
	return ApiAllMarketLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketLiquidationOrderStreamsResponse
func (a *DefaultAPIService) AllMarketLiquidationOrderStreamsExecute(r ApiAllMarketLiquidationOrderStreamsRequest) (*common.StreamHandler[models.AllMarketLiquidationOrderStreamsResponse], error) {

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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#all-market-mini-tickers-stream

@param id Unique WebSocket request ID.
@return ApiAllMarketMiniTickersStreamRequest
*/
func (a *DefaultAPIService) AllMarketMiniTickersStream() ApiAllMarketMiniTickersStreamRequest {
	return ApiAllMarketMiniTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketMiniTickersStreamResponse
func (a *DefaultAPIService) AllMarketMiniTickersStreamExecute(r ApiAllMarketMiniTickersStreamRequest) (*common.StreamHandler[models.AllMarketMiniTickersStreamResponse], error) {

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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#all-market-tickers-streams

@param id Unique WebSocket request ID.
@return ApiAllMarketTickersStreamsRequest
*/
func (a *DefaultAPIService) AllMarketTickersStreams() ApiAllMarketTickersStreamsRequest {
	return ApiAllMarketTickersStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketTickersStreamsResponse
func (a *DefaultAPIService) AllMarketTickersStreamsExecute(r ApiAllMarketTickersStreamsRequest) (*common.StreamHandler[models.AllMarketTickersStreamsResponse], error) {

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

type ApiContinuousContractKlineCandlestickStreamsRequest struct {
	ApiService   *DefaultAPIService
	pair         *string
	contractType *models.ContinuousContractKlineCandlestickStreamsContractTypeParameter
	interval     *models.ContinuousContractKlineCandlestickStreamsIntervalParameter
	id           *string
}

// The pair parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) Pair(pair string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.pair = &pair
	return r
}

// The contractType parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) ContractType(contractType models.ContinuousContractKlineCandlestickStreamsContractTypeParameter) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.contractType = &contractType
	return r
}

// The interval parameter
func (r ApiContinuousContractKlineCandlestickStreamsRequest) Interval(interval models.ContinuousContractKlineCandlestickStreamsIntervalParameter) ApiContinuousContractKlineCandlestickStreamsRequest {
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#continuous-contract-kline-candlestick-streams

@param pair The pair parameter	@param contractType The contractType parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiContinuousContractKlineCandlestickStreamsRequest
*/
func (a *DefaultAPIService) ContinuousContractKlineCandlestickStreams() ApiContinuousContractKlineCandlestickStreamsRequest {
	return ApiContinuousContractKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContinuousContractKlineCandlestickStreamsResponse
func (a *DefaultAPIService) ContinuousContractKlineCandlestickStreamsExecute(r ApiContinuousContractKlineCandlestickStreamsRequest) (*common.StreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse], error) {
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
				return string(*r.contractType)
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
	resp, err := common.CreateStreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiContractInfoStreamRequest struct {
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#contract-info-stream

@param id Unique WebSocket request ID.
@return ApiContractInfoStreamRequest
*/
func (a *DefaultAPIService) ContractInfoStream() ApiContractInfoStreamRequest {
	return ApiContractInfoStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContractInfoStreamResponse
func (a *DefaultAPIService) ContractInfoStreamExecute(r ApiContractInfoStreamRequest) (*common.StreamHandler[models.ContractInfoStreamResponse], error) {

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
	ApiService  *DefaultAPIService
	symbol      *string
	id          *string
	updateSpeed *models.DiffBookDepthStreamsUpdateSpeedParameter
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
func (r ApiDiffBookDepthStreamsRequest) UpdateSpeed(updateSpeed models.DiffBookDepthStreamsUpdateSpeedParameter) ApiDiffBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiDiffBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
	return r.ApiService.DiffBookDepthStreamsExecute(r)
}

/*
DiffBookDepthStreams Diff. Book Depth Streams
/<symbol>@depth@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#diff-book-depth-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiDiffBookDepthStreamsRequest
*/
func (a *DefaultAPIService) DiffBookDepthStreams() ApiDiffBookDepthStreamsRequest {
	return ApiDiffBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return DiffBookDepthStreamsResponse
func (a *DefaultAPIService) DiffBookDepthStreamsExecute(r ApiDiffBookDepthStreamsRequest) (*common.StreamHandler[models.DiffBookDepthStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.DiffBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndexKlineCandlestickStreamsRequest struct {
	ApiService *DefaultAPIService
	pair       *string
	interval   *models.ContinuousContractKlineCandlestickStreamsIntervalParameter
	id         *string
}

// The pair parameter
func (r ApiIndexKlineCandlestickStreamsRequest) Pair(pair string) ApiIndexKlineCandlestickStreamsRequest {
	r.pair = &pair
	return r
}

// The interval parameter
func (r ApiIndexKlineCandlestickStreamsRequest) Interval(interval models.ContinuousContractKlineCandlestickStreamsIntervalParameter) ApiIndexKlineCandlestickStreamsRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiIndexKlineCandlestickStreamsRequest) Id(id string) ApiIndexKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiIndexKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.IndexKlineCandlestickStreamsResponse], error) {
	return r.ApiService.IndexKlineCandlestickStreamsExecute(r)
}

/*
IndexKlineCandlestickStreams Index Kline/Candlestick Streams
/<pair>@indexPriceKline_<interval>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#index-kline-candlestick-streams

@param pair The pair parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiIndexKlineCandlestickStreamsRequest
*/
func (a *DefaultAPIService) IndexKlineCandlestickStreams() ApiIndexKlineCandlestickStreamsRequest {
	return ApiIndexKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndexKlineCandlestickStreamsResponse
func (a *DefaultAPIService) IndexKlineCandlestickStreamsExecute(r ApiIndexKlineCandlestickStreamsRequest) (*common.StreamHandler[models.IndexKlineCandlestickStreamsResponse], error) {
	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<pair>@indexPriceKline_<interval>"[1:],
		map[string]string{
			"pair": func() string {
				if r.pair == nil {
					return ""
				}
				return *r.pair
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
	resp, err := common.CreateStreamHandler[models.IndexKlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndexPriceStreamRequest struct {
	ApiService  *DefaultAPIService
	pair        *string
	id          *string
	updateSpeed *models.IndexPriceStreamUpdateSpeedParameter
}

// The pair parameter
func (r ApiIndexPriceStreamRequest) Pair(pair string) ApiIndexPriceStreamRequest {
	r.pair = &pair
	return r
}

// Unique WebSocket request ID.
func (r ApiIndexPriceStreamRequest) Id(id string) ApiIndexPriceStreamRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiIndexPriceStreamRequest) UpdateSpeed(updateSpeed models.IndexPriceStreamUpdateSpeedParameter) ApiIndexPriceStreamRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiIndexPriceStreamRequest) Execute() (*common.StreamHandler[models.IndexPriceStreamResponse], error) {
	return r.ApiService.IndexPriceStreamExecute(r)
}

/*
IndexPriceStream Index Price Stream
/<pair>@indexPrice@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#index-price-stream

@param pair The pair parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiIndexPriceStreamRequest
*/
func (a *DefaultAPIService) IndexPriceStream() ApiIndexPriceStreamRequest {
	return ApiIndexPriceStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndexPriceStreamResponse
func (a *DefaultAPIService) IndexPriceStreamExecute(r ApiIndexPriceStreamRequest) (*common.StreamHandler[models.IndexPriceStreamResponse], error) {
	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<pair>@indexPrice@<updateSpeed>"[1:],
		map[string]string{
			"pair": func() string {
				if r.pair == nil {
					return ""
				}
				return *r.pair
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
	resp, err := common.CreateStreamHandler[models.IndexPriceStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiIndividualSymbolBookTickerStreamsRequest struct {
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#individual-symbol-book-ticker-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolBookTickerStreamsRequest
*/
func (a *DefaultAPIService) IndividualSymbolBookTickerStreams() ApiIndividualSymbolBookTickerStreamsRequest {
	return ApiIndividualSymbolBookTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolBookTickerStreamsResponse
func (a *DefaultAPIService) IndividualSymbolBookTickerStreamsExecute(r ApiIndividualSymbolBookTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse], error) {
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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#individual-symbol-mini-ticker-stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolMiniTickerStreamRequest
*/
func (a *DefaultAPIService) IndividualSymbolMiniTickerStream() ApiIndividualSymbolMiniTickerStreamRequest {
	return ApiIndividualSymbolMiniTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolMiniTickerStreamResponse
func (a *DefaultAPIService) IndividualSymbolMiniTickerStreamExecute(r ApiIndividualSymbolMiniTickerStreamRequest) (*common.StreamHandler[models.IndividualSymbolMiniTickerStreamResponse], error) {
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
	ApiService *DefaultAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#individual-symbol-ticker-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolTickerStreamsRequest
*/
func (a *DefaultAPIService) IndividualSymbolTickerStreams() ApiIndividualSymbolTickerStreamsRequest {
	return ApiIndividualSymbolTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolTickerStreamsResponse
func (a *DefaultAPIService) IndividualSymbolTickerStreamsExecute(r ApiIndividualSymbolTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolTickerStreamsResponse], error) {
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
	ApiService *DefaultAPIService
	symbol     *string
	interval   *models.ContinuousContractKlineCandlestickStreamsIntervalParameter
	id         *string
}

// The symbol parameter
func (r ApiKlineCandlestickStreamsRequest) Symbol(symbol string) ApiKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiKlineCandlestickStreamsRequest) Interval(interval models.ContinuousContractKlineCandlestickStreamsIntervalParameter) ApiKlineCandlestickStreamsRequest {
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#kline-candlestick-streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiKlineCandlestickStreamsRequest
*/
func (a *DefaultAPIService) KlineCandlestickStreams() ApiKlineCandlestickStreamsRequest {
	return ApiKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickStreamsResponse
func (a *DefaultAPIService) KlineCandlestickStreamsExecute(r ApiKlineCandlestickStreamsRequest) (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.KlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceKlineCandlestickStreamsRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	interval   *models.ContinuousContractKlineCandlestickStreamsIntervalParameter
	id         *string
}

// The symbol parameter
func (r ApiMarkPriceKlineCandlestickStreamsRequest) Symbol(symbol string) ApiMarkPriceKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiMarkPriceKlineCandlestickStreamsRequest) Interval(interval models.ContinuousContractKlineCandlestickStreamsIntervalParameter) ApiMarkPriceKlineCandlestickStreamsRequest {
	r.interval = &interval
	return r
}

// Unique WebSocket request ID.
func (r ApiMarkPriceKlineCandlestickStreamsRequest) Id(id string) ApiMarkPriceKlineCandlestickStreamsRequest {
	r.id = &id
	return r
}

func (r ApiMarkPriceKlineCandlestickStreamsRequest) Execute() (*common.StreamHandler[models.MarkPriceKlineCandlestickStreamsResponse], error) {
	return r.ApiService.MarkPriceKlineCandlestickStreamsExecute(r)
}

/*
MarkPriceKlineCandlestickStreams Mark Price Kline/Candlestick Streams
/<symbol>@markPriceKline_<interval>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#mark-price-kline-candlestick-streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiMarkPriceKlineCandlestickStreamsRequest
*/
func (a *DefaultAPIService) MarkPriceKlineCandlestickStreams() ApiMarkPriceKlineCandlestickStreamsRequest {
	return ApiMarkPriceKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceKlineCandlestickStreamsResponse
func (a *DefaultAPIService) MarkPriceKlineCandlestickStreamsExecute(r ApiMarkPriceKlineCandlestickStreamsRequest) (*common.StreamHandler[models.MarkPriceKlineCandlestickStreamsResponse], error) {
	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<symbol>@markPriceKline_<interval>"[1:],
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
	resp, err := common.CreateStreamHandler[models.MarkPriceKlineCandlestickStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceOfAllSymbolsOfAPairRequest struct {
	ApiService  *DefaultAPIService
	pair        *string
	id          *string
	updateSpeed *models.IndexPriceStreamUpdateSpeedParameter
}

// The pair parameter
func (r ApiMarkPriceOfAllSymbolsOfAPairRequest) Pair(pair string) ApiMarkPriceOfAllSymbolsOfAPairRequest {
	r.pair = &pair
	return r
}

// Unique WebSocket request ID.
func (r ApiMarkPriceOfAllSymbolsOfAPairRequest) Id(id string) ApiMarkPriceOfAllSymbolsOfAPairRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiMarkPriceOfAllSymbolsOfAPairRequest) UpdateSpeed(updateSpeed models.IndexPriceStreamUpdateSpeedParameter) ApiMarkPriceOfAllSymbolsOfAPairRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceOfAllSymbolsOfAPairRequest) Execute() (*common.StreamHandler[models.MarkPriceOfAllSymbolsOfAPairResponse], error) {
	return r.ApiService.MarkPriceOfAllSymbolsOfAPairExecute(r)
}

/*
MarkPriceOfAllSymbolsOfAPair Mark Price of All Symbols of a Pair
/<pair>@markPrice@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#mark-price-of-all-symbols-of-apair

@param pair The pair parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceOfAllSymbolsOfAPairRequest
*/
func (a *DefaultAPIService) MarkPriceOfAllSymbolsOfAPair() ApiMarkPriceOfAllSymbolsOfAPairRequest {
	return ApiMarkPriceOfAllSymbolsOfAPairRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceOfAllSymbolsOfAPairResponse
func (a *DefaultAPIService) MarkPriceOfAllSymbolsOfAPairExecute(r ApiMarkPriceOfAllSymbolsOfAPairRequest) (*common.StreamHandler[models.MarkPriceOfAllSymbolsOfAPairResponse], error) {
	if r.pair == nil {
		return nil, common.ReportError("pair is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<pair>@markPrice@<updateSpeed>"[1:],
		map[string]string{
			"pair": func() string {
				if r.pair == nil {
					return ""
				}
				return *r.pair
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
	resp, err := common.CreateStreamHandler[models.MarkPriceOfAllSymbolsOfAPairResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiMarkPriceStreamRequest struct {
	ApiService  *DefaultAPIService
	symbol      *string
	id          *string
	updateSpeed *models.IndexPriceStreamUpdateSpeedParameter
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
func (r ApiMarkPriceStreamRequest) UpdateSpeed(updateSpeed models.IndexPriceStreamUpdateSpeedParameter) ApiMarkPriceStreamRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceStreamRequest) Execute() (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
	return r.ApiService.MarkPriceStreamExecute(r)
}

/*
MarkPriceStream Mark Price Stream
/<symbol>@markPrice@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#mark-price-stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceStreamRequest
*/
func (a *DefaultAPIService) MarkPriceStream() ApiMarkPriceStreamRequest {
	return ApiMarkPriceStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceStreamResponse
func (a *DefaultAPIService) MarkPriceStreamExecute(r ApiMarkPriceStreamRequest) (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
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
				return string(*r.updateSpeed)
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

type ApiMarketLiquidationOrderStreamsRequest struct {
	ApiService *DefaultAPIService
	symbol     *string
	id         *string
}

// The symbol parameter
func (r ApiMarketLiquidationOrderStreamsRequest) Symbol(symbol string) ApiMarketLiquidationOrderStreamsRequest {
	r.symbol = &symbol
	return r
}

// Unique WebSocket request ID.
func (r ApiMarketLiquidationOrderStreamsRequest) Id(id string) ApiMarketLiquidationOrderStreamsRequest {
	r.id = &id
	return r
}

func (r ApiMarketLiquidationOrderStreamsRequest) Execute() (*common.StreamHandler[models.MarketLiquidationOrderStreamsResponse], error) {
	return r.ApiService.MarketLiquidationOrderStreamsExecute(r)
}

/*
MarketLiquidationOrderStreams Market Liquidation Order Streams
/<symbol>@forceOrder

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#market-liquidation-order-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiMarketLiquidationOrderStreamsRequest
*/
func (a *DefaultAPIService) MarketLiquidationOrderStreams() ApiMarketLiquidationOrderStreamsRequest {
	return ApiMarketLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarketLiquidationOrderStreamsResponse
func (a *DefaultAPIService) MarketLiquidationOrderStreamsExecute(r ApiMarketLiquidationOrderStreamsRequest) (*common.StreamHandler[models.MarketLiquidationOrderStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.MarketLiquidationOrderStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiPartialBookDepthStreamsRequest struct {
	ApiService  *DefaultAPIService
	symbol      *string
	levels      *models.PartialBookDepthStreamsLevelsParameter
	id          *string
	updateSpeed *models.DiffBookDepthStreamsUpdateSpeedParameter
}

// The symbol parameter
func (r ApiPartialBookDepthStreamsRequest) Symbol(symbol string) ApiPartialBookDepthStreamsRequest {
	r.symbol = &symbol
	return r
}

// The levels parameter
func (r ApiPartialBookDepthStreamsRequest) Levels(levels models.PartialBookDepthStreamsLevelsParameter) ApiPartialBookDepthStreamsRequest {
	r.levels = &levels
	return r
}

// Unique WebSocket request ID.
func (r ApiPartialBookDepthStreamsRequest) Id(id string) ApiPartialBookDepthStreamsRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiPartialBookDepthStreamsRequest) UpdateSpeed(updateSpeed models.DiffBookDepthStreamsUpdateSpeedParameter) ApiPartialBookDepthStreamsRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiPartialBookDepthStreamsRequest) Execute() (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
	return r.ApiService.PartialBookDepthStreamsExecute(r)
}

/*
PartialBookDepthStreams Partial Book Depth Streams
/<symbol>@depth<levels>@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-coin-m-futures/api/ws-streams/~#partial-book-depth-streams

@param symbol The symbol parameter	@param levels The levels parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiPartialBookDepthStreamsRequest
*/
func (a *DefaultAPIService) PartialBookDepthStreams() ApiPartialBookDepthStreamsRequest {
	return ApiPartialBookDepthStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return PartialBookDepthStreamsResponse
func (a *DefaultAPIService) PartialBookDepthStreamsExecute(r ApiPartialBookDepthStreamsRequest) (*common.StreamHandler[models.PartialBookDepthStreamsResponse], error) {
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
	resp, err := common.CreateStreamHandler[models.PartialBookDepthStreamsResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
