/*
Futures (USDⓈ-M) WebSocket Market Streams

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package binancederivativestradingusdsfutureswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// MarketAPIService MarketAPI Service
type MarketAPIService Service

type ApiAggregateTradeStreamsRequest struct {
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#aggregate-trade-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiAggregateTradeStreamsRequest
*/
func (a *MarketAPIService) AggregateTradeStreams() ApiAggregateTradeStreamsRequest {
	return ApiAggregateTradeStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AggregateTradeStreamsResponse
func (a *MarketAPIService) AggregateTradeStreamsExecute(r ApiAggregateTradeStreamsRequest) (*common.StreamHandler[models.AggregateTradeStreamsResponse], error) {
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
	ws := a.client.WsMarket

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

type ApiAllMarketLiquidationOrderStreamsRequest struct {
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#all-market-liquidation-order-streams

@param id Unique WebSocket request ID.
@return ApiAllMarketLiquidationOrderStreamsRequest
*/
func (a *MarketAPIService) AllMarketLiquidationOrderStreams() ApiAllMarketLiquidationOrderStreamsRequest {
	return ApiAllMarketLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketLiquidationOrderStreamsResponse
func (a *MarketAPIService) AllMarketLiquidationOrderStreamsExecute(r ApiAllMarketLiquidationOrderStreamsRequest) (*common.StreamHandler[models.AllMarketLiquidationOrderStreamsResponse], error) {

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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#all-market-mini-tickers-stream

@param id Unique WebSocket request ID.
@return ApiAllMarketMiniTickersStreamRequest
*/
func (a *MarketAPIService) AllMarketMiniTickersStream() ApiAllMarketMiniTickersStreamRequest {
	return ApiAllMarketMiniTickersStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketMiniTickersStreamResponse
func (a *MarketAPIService) AllMarketMiniTickersStreamExecute(r ApiAllMarketMiniTickersStreamRequest) (*common.StreamHandler[models.AllMarketMiniTickersStreamResponse], error) {

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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#all-market-tickers-streams

@param id Unique WebSocket request ID.
@return ApiAllMarketTickersStreamsRequest
*/
func (a *MarketAPIService) AllMarketTickersStreams() ApiAllMarketTickersStreamsRequest {
	return ApiAllMarketTickersStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AllMarketTickersStreamsResponse
func (a *MarketAPIService) AllMarketTickersStreamsExecute(r ApiAllMarketTickersStreamsRequest) (*common.StreamHandler[models.AllMarketTickersStreamsResponse], error) {

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
	ws := a.client.WsMarket

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

type ApiAssetIndexRequest struct {
	ApiService *MarketAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiAssetIndexRequest) Id(id string) ApiAssetIndexRequest {
	r.id = &id
	return r
}

func (r ApiAssetIndexRequest) Execute() (*common.StreamHandler[models.AssetIndexResponse], error) {
	return r.ApiService.AssetIndexExecute(r)
}

/*
AssetIndex Multi-Assets Mode Asset Index
/!assetIndex@arr

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#asset-index

@param id Unique WebSocket request ID.
@return ApiAssetIndexRequest
*/
func (a *MarketAPIService) AssetIndex() ApiAssetIndexRequest {
	return ApiAssetIndexRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return AssetIndexResponse
func (a *MarketAPIService) AssetIndexExecute(r ApiAssetIndexRequest) (*common.StreamHandler[models.AssetIndexResponse], error) {

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
	ws := a.client.WsMarket

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.AssetIndexResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ApiCompositeIndexSymbolInformationStreamsRequest struct {
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#composite-index-symbol-information-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiCompositeIndexSymbolInformationStreamsRequest
*/
func (a *MarketAPIService) CompositeIndexSymbolInformationStreams() ApiCompositeIndexSymbolInformationStreamsRequest {
	return ApiCompositeIndexSymbolInformationStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return CompositeIndexSymbolInformationStreamsResponse
func (a *MarketAPIService) CompositeIndexSymbolInformationStreamsExecute(r ApiCompositeIndexSymbolInformationStreamsRequest) (*common.StreamHandler[models.CompositeIndexSymbolInformationStreamsResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService   *MarketAPIService
	pair         *string
	contractType *models.ContinuousContractKlineCandlestickStreamsContractTypeParameter
	interval     *models.ContinuousContractKlineCandlestickStreamsIntervalParameter
	id           *string
}

func (r ApiContinuousContractKlineCandlestickStreamsRequest) Pair(pair string) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.pair = &pair
	return r
}

func (r ApiContinuousContractKlineCandlestickStreamsRequest) ContractType(contractType models.ContinuousContractKlineCandlestickStreamsContractTypeParameter) ApiContinuousContractKlineCandlestickStreamsRequest {
	r.contractType = &contractType
	return r
}

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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#continuous-contract-kline-candlestick-streams

@param pair	@param contractType	@param interval	@param id Unique WebSocket request ID.
@return ApiContinuousContractKlineCandlestickStreamsRequest
*/
func (a *MarketAPIService) ContinuousContractKlineCandlestickStreams() ApiContinuousContractKlineCandlestickStreamsRequest {
	return ApiContinuousContractKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContinuousContractKlineCandlestickStreamsResponse
func (a *MarketAPIService) ContinuousContractKlineCandlestickStreamsExecute(r ApiContinuousContractKlineCandlestickStreamsRequest) (*common.StreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#contract-info-stream

@param id Unique WebSocket request ID.
@return ApiContractInfoStreamRequest
*/
func (a *MarketAPIService) ContractInfoStream() ApiContractInfoStreamRequest {
	return ApiContractInfoStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ContractInfoStreamResponse
func (a *MarketAPIService) ContractInfoStreamExecute(r ApiContractInfoStreamRequest) (*common.StreamHandler[models.ContractInfoStreamResponse], error) {

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
	ws := a.client.WsMarket

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

type ApiIndividualSymbolMiniTickerStreamRequest struct {
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#individual-symbol-mini-ticker-stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolMiniTickerStreamRequest
*/
func (a *MarketAPIService) IndividualSymbolMiniTickerStream() ApiIndividualSymbolMiniTickerStreamRequest {
	return ApiIndividualSymbolMiniTickerStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolMiniTickerStreamResponse
func (a *MarketAPIService) IndividualSymbolMiniTickerStreamExecute(r ApiIndividualSymbolMiniTickerStreamRequest) (*common.StreamHandler[models.IndividualSymbolMiniTickerStreamResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#individual-symbol-ticker-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiIndividualSymbolTickerStreamsRequest
*/
func (a *MarketAPIService) IndividualSymbolTickerStreams() ApiIndividualSymbolTickerStreamsRequest {
	return ApiIndividualSymbolTickerStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return IndividualSymbolTickerStreamsResponse
func (a *MarketAPIService) IndividualSymbolTickerStreamsExecute(r ApiIndividualSymbolTickerStreamsRequest) (*common.StreamHandler[models.IndividualSymbolTickerStreamsResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
	symbol     *string
	interval   *models.KlineCandlestickStreamsIntervalParameter
	id         *string
}

// The symbol parameter
func (r ApiKlineCandlestickStreamsRequest) Symbol(symbol string) ApiKlineCandlestickStreamsRequest {
	r.symbol = &symbol
	return r
}

// The interval parameter
func (r ApiKlineCandlestickStreamsRequest) Interval(interval models.KlineCandlestickStreamsIntervalParameter) ApiKlineCandlestickStreamsRequest {
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#kline-candlestick-streams

@param symbol The symbol parameter	@param interval The interval parameter	@param id Unique WebSocket request ID.
@return ApiKlineCandlestickStreamsRequest
*/
func (a *MarketAPIService) KlineCandlestickStreams() ApiKlineCandlestickStreamsRequest {
	return ApiKlineCandlestickStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return KlineCandlestickStreamsResponse
func (a *MarketAPIService) KlineCandlestickStreamsExecute(r ApiKlineCandlestickStreamsRequest) (*common.StreamHandler[models.KlineCandlestickStreamsResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#liquidation-order-streams

@param symbol The symbol parameter	@param id Unique WebSocket request ID.
@return ApiLiquidationOrderStreamsRequest
*/
func (a *MarketAPIService) LiquidationOrderStreams() ApiLiquidationOrderStreamsRequest {
	return ApiLiquidationOrderStreamsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return LiquidationOrderStreamsResponse
func (a *MarketAPIService) LiquidationOrderStreamsExecute(r ApiLiquidationOrderStreamsRequest) (*common.StreamHandler[models.LiquidationOrderStreamsResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService  *MarketAPIService
	symbol      *string
	id          *string
	updateSpeed *models.MarkPriceStreamUpdateSpeedParameter
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
func (r ApiMarkPriceStreamRequest) UpdateSpeed(updateSpeed models.MarkPriceStreamUpdateSpeedParameter) ApiMarkPriceStreamRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceStreamRequest) Execute() (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
	return r.ApiService.MarkPriceStreamExecute(r)
}

/*
MarkPriceStream Mark Price Stream
/<symbol>@markPrice@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#mark-price-stream

@param symbol The symbol parameter	@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceStreamRequest
*/
func (a *MarketAPIService) MarkPriceStream() ApiMarkPriceStreamRequest {
	return ApiMarkPriceStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceStreamResponse
func (a *MarketAPIService) MarkPriceStreamExecute(r ApiMarkPriceStreamRequest) (*common.StreamHandler[models.MarkPriceStreamResponse], error) {
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
	ws := a.client.WsMarket

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
	ApiService  *MarketAPIService
	id          *string
	updateSpeed *models.MarkPriceStreamUpdateSpeedParameter
}

// Unique WebSocket request ID.
func (r ApiMarkPriceStreamForAllMarketRequest) Id(id string) ApiMarkPriceStreamForAllMarketRequest {
	r.id = &id
	return r
}

// WebSocket stream update speed
func (r ApiMarkPriceStreamForAllMarketRequest) UpdateSpeed(updateSpeed models.MarkPriceStreamUpdateSpeedParameter) ApiMarkPriceStreamForAllMarketRequest {
	r.updateSpeed = &updateSpeed
	return r
}

func (r ApiMarkPriceStreamForAllMarketRequest) Execute() (*common.StreamHandler[models.MarkPriceStreamForAllMarketResponse], error) {
	return r.ApiService.MarkPriceStreamForAllMarketExecute(r)
}

/*
MarkPriceStreamForAllMarket Mark Price Stream for All market
/!markPrice@arr@<updateSpeed>

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#mark-price-stream-for-all-market

@param id Unique WebSocket request ID.	@param updateSpeed WebSocket stream update speed
@return ApiMarkPriceStreamForAllMarketRequest
*/
func (a *MarketAPIService) MarkPriceStreamForAllMarket() ApiMarkPriceStreamForAllMarketRequest {
	return ApiMarkPriceStreamForAllMarketRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return MarkPriceStreamForAllMarketResponse
func (a *MarketAPIService) MarkPriceStreamForAllMarketExecute(r ApiMarkPriceStreamForAllMarketRequest) (*common.StreamHandler[models.MarkPriceStreamForAllMarketResponse], error) {

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
				return string(*r.updateSpeed)
			}(),
		},
	)
	ws := a.client.WsMarket

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

type ApiTradingSessionStreamRequest struct {
	ApiService *MarketAPIService
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

https://developers.binance.com/en/docs/catalog/core-trading-derivatives-trading-usd-s-m-futures/api/ws-streams/market#trading-session-stream

@param id Unique WebSocket request ID.
@return ApiTradingSessionStreamRequest
*/
func (a *MarketAPIService) TradingSessionStream() ApiTradingSessionStreamRequest {
	return ApiTradingSessionStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return TradingSessionStreamResponse
func (a *MarketAPIService) TradingSessionStreamExecute(r ApiTradingSessionStreamRequest) (*common.StreamHandler[models.TradingSessionStreamResponse], error) {

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
	ws := a.client.WsMarket

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
