### Changelog

## 1.18.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.17.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.16.0 - 2026-08-18

### Changed (1)

- Updated dependencies

## 1.15.0 - 2026-08-18

### Changed (4)

#### REST API

- Added parameter `contractType`
  - affected methods:
    - `longShortRatio()` (`GET /futures/data/globalLongShortAccountRatio`)
    - `topTraderLongShortRatioAccounts()` (`GET /futures/data/topLongShortAccountRatio`)
    - `topTraderLongShortRatioPositions()` (`GET /futures/data/topLongShortPositionRatio`)
- Added parameter `pair`
  - affected methods:
    - `topTraderLongShortRatioAccounts()` (`GET /futures/data/topLongShortAccountRatio`)
- Deleted parameter `symbol`
  - affected methods:
    - `topTraderLongShortRatioAccounts()` (`GET /futures/data/topLongShortAccountRatio`)
- Modified parameter `contractType`:
  - required: `true` → `false`
  - enum removed: `ALL`
  - affected methods:
    - `openInterestStatistics()` (`GET /futures/data/openInterestHist`)

## 1.14.0 - 2026-08-14

### Changed (1)

#### WebSocket Streams

- Modified response field `a`:
  - property `S` added
  - affected events:
    - `UserDataStreamEventsResponse`
    - `accountUpdate`

## 1.13.0 - 2026-08-06

### Changed (3)

#### REST API

- Modified response for `allOrders()` (`GET /dapi/v1/allOrders`):
  - items: property `cumQuote` added
  - items: property `goodTillDate` added
  - items: item property `cumQuote` added
  - items: item property `goodTillDate` added

- Modified response for `usersForceOrders()` (`GET /dapi/v1/forceOrders`):
  - items: property `cumQuote` added
  - items: property `goodTillDate` added
  - items: item property `cumQuote` added
  - items: item property `goodTillDate` added

- Modified response for `accountTradeList()` (`GET /dapi/v1/userTrades`):
  - items: property `quoteQty` added
  - items: item property `quoteQty` added

## 1.12.0 - 2026-07-28

### Changed (7)

#### REST API

- Added parameter `modifyId`
  - affected methods:
    - `modifyOrder()` (`PUT /dapi/v1/order`)
- Modified parameter `batchOrders`:
  - items: property `modifyId` added
  - items: item property `modifyId` added
  - affected methods:
    - `modifyMultipleOrders()` (`PUT /dapi/v1/batchOrders`)
- Modified response for `modifyMultipleOrders()` (`PUT /dapi/v1/batchOrders`):
  - items: property `modifyId` added
  - items: item property `modifyId` added

- Modified response for `modifyOrder()` (`PUT /dapi/v1/order`):
  - property `modifyId` added

- Modified response for `getOrderModifyHistory()` (`GET /dapi/v1/orderAmendment`):
  - items.`amendment`: property `modifyId` added
  - items.`amendment`: property `modifyId` added

#### WebSocket API

- Added parameter `modifyId`
  - affected methods:
    - `modifyOrder()` (`order.modify` method)
- Modified response for `modifyOrder()` (`order.modify` method):
  - `result`: property `modifyId` added

## 1.11.0 - 2026-07-15

### Changed (10)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

#### REST API

- Modified response for `cancelMultipleOrders()` (`DELETE /dapi/v1/batchOrders`):
  - items: property `avgPrice` deleted
  - items: property `cumBase` deleted
  - items: item property `avgPrice` deleted
  - items: item property `cumBase` deleted

- Modified response for `placeMultipleOrders()` (`POST /dapi/v1/batchOrders`):
  - items: property `cumBase` deleted
  - items: property `avgPrice` deleted
  - items: item property `cumBase` deleted
  - items: item property `avgPrice` deleted

- Modified response for `modifyMultipleOrders()` (`PUT /dapi/v1/batchOrders`):
  - items: property `avgPrice` deleted
  - items: property `cumBase` deleted
  - items: item property `avgPrice` deleted
  - items: item property `cumBase` deleted

- Modified response for `cancelOrder()` (`DELETE /dapi/v1/order`):
  - property `avgPrice` deleted
  - property `cumBase` deleted

- Modified response for `newOrder()` (`POST /dapi/v1/order`):
  - property `cumBase` deleted
  - property `avgPrice` deleted

- Modified response for `modifyOrder()` (`PUT /dapi/v1/order`):
  - property `avgPrice` deleted
  - property `cumBase` deleted

#### WebSocket API

- Modified response for `cancelOrder()` (`order.cancel` method):
  - `result`: property `cumBase` deleted
  - `result`: property `avgPrice` deleted

- Modified response for `modifyOrder()` (`order.modify` method):
  - `result`: property `avgPrice` deleted
  - `result`: property `cumBase` deleted

- Modified response for `newOrder()` (`order.place` method):
  - `result`: property `avgPrice` deleted
  - `result`: property `cumBase` deleted

## 1.10.0 - 2026-07-14

### Added (1)

#### WebSocket Streams

- `indexPriceStream()` (`<pair>@indexPrice@<updateSpeed>` stream)

### Changed (25)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

#### REST API

- Modified parameter `batchOrders`:
  - items: required added: `side`, `type`, `quantity`, `symbol`
  - items.`activationPrice`: type `string` → `number`
  - items.`callbackRate`: type `string` → `number`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`priceProtect`: enum added: `true`, `false`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: enum added: `true`, `false`
  - items.`selfTradePreventionMode`: enum removed: `NONE`
  - items.`stopPrice`: type `string` → `number`
  - items.`activationPrice`: type `string` → `number`
  - items.`callbackRate`: type `string` → `number`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`priceProtect`: enum added: `true`, `false`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: enum added: `true`, `false`
  - items.`selfTradePreventionMode`: enum removed: `NONE`
  - items.`stopPrice`: type `string` → `number`
  - affected methods:
    - `placeMultipleOrders()` (`POST /dapi/v1/batchOrders`)
- Modified parameter `batchOrders`:
  - items: required added: `symbol`, `side`, `timestamp`
  - items: property `timestamp` added
  - items.`orderId`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`quantity`: type `string` → `number`
  - items.`recvWindow`: type `string` → `integer`
  - items: item property `timestamp` added
  - items.`orderId`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`quantity`: type `string` → `number`
  - items.`recvWindow`: type `string` → `integer`
  - affected methods:
    - `modifyMultipleOrders()` (`PUT /dapi/v1/batchOrders`)
- Modified parameter `contractType`:
  - enum removed: `CURRENT_QUARTER_DELIVERING`, `NEXT_QUARTER_DELIVERING`, `PERPETUAL_DELIVERING`
  - affected methods:
    - `continuousContractKlineCandlestickData()` (`GET /dapi/v1/continuousKlines`)
    - `basis()` (`GET /futures/data/basis`)
- Modified parameter `contractType`:
  - enum removed: `CURRENT_QUARTER_DELIVERING`, `NEXT_QUARTER_DELIVERING`, `PERPETUAL_DELIVERING`
  - enum added: `ALL`
  - affected methods:
    - `openInterestStatistics()` (`GET /futures/data/openInterestHist`)
    - `takerBuySellVolume()` (`GET /futures/data/takerBuySellVol`)
- Modified parameter `incomeType`:
  - enum added: `TRANSFER`, `WELCOME_BONUS`, `FUNDING_FEE`, `REALIZED_PNL`, `COMMISSION`, `INSURANCE_CLEAR`, `DELIVERED_SETTELMENT`
  - affected methods:
    - `getIncomeHistory()` (`GET /dapi/v1/income`)
- Modified parameter `orderIdList`:
  - maxItems `null` → `10`
  - affected methods:
    - `cancelMultipleOrders()` (`DELETE /dapi/v1/batchOrders`)
- Modified parameter `origClientOrderIdList`:
  - maxItems `null` → `10`
  - affected methods:
    - `cancelMultipleOrders()` (`DELETE /dapi/v1/batchOrders`)
- Modified parameter `priceMatch`:
  - enum removed: `NONE`
  - affected methods:
    - `newOrder()` (`POST /dapi/v1/order`)
    - `modifyOrder()` (`PUT /dapi/v1/order`)
- Modified parameter `priceProtect`:
  - enum added: `true`, `false`
  - affected methods:
    - `newOrder()` (`POST /dapi/v1/order`)
- Modified parameter `reduceOnly`:
  - enum added: `true`, `false`
  - affected methods:
    - `newOrder()` (`POST /dapi/v1/order`)
- Modified parameter `type`:
  - type `string` → `integer`
  - enum removed: `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`
  - affected methods:
    - `modifyIsolatedPositionMargin()` (`POST /dapi/v1/positionMargin`)
- Modified response for `placeMultipleOrders()` (`POST /dapi/v1/batchOrders`):
  - items: property `closePosition` added
  - items: item property `closePosition` added

- Modified response for `orderBook()` (`GET /dapi/v1/depth`):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `markPriceKlineCandlestickData()` (`GET /dapi/v1/markPriceKlines`):
  - items.items: oneOf added 2 schema(s)
  - items.items: oneOf removed 2 schema(s)

#### WebSocket API

- Modified parameter `closePosition`:
  - enum added: `true`, `false`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Modified parameter `priceMatch`:
  - enum removed: `NONE`
  - affected methods:
    - `modifyOrder()` (`order.modify` method)
    - `newOrder()` (`order.place` method)
- Modified parameter `priceProtect`:
  - enum added: `true`, `false`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Modified parameter `reduceOnly`:
  - enum added: `true`, `false`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Modified response for `queryOrder()` (`order.status` method):
  - `result`: property `cumQty` added

#### WebSocket Streams

- Modified parameter `contractType`:
  - enum added: `perpetual`, `current_quarter`, `next_quarter`
  - affected methods:
    - `continuousContractKlineCandlestickStreams()` (`<pair>_<contractType>@continuousKline_<interval>` stream)
- Modified parameter `interval`:
  - enum added: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1M`
  - affected methods:
    - `indexKlineCandlestickStreams()` (`<pair>@indexPriceKline_<interval>` stream)
    - `continuousContractKlineCandlestickStreams()` (`<pair>_<contractType>@continuousKline_<interval>` stream)
    - `klineCandlestickStreams()` (`<symbol>@kline_<interval>` stream)
    - `markPriceKlineCandlestickStreams()` (`<symbol>@markPriceKline_<interval>` stream)
- Modified parameter `levels`:
  - type `integer` → `string`
  - enum added: `5`, `10`, `20`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
- Modified parameter `updateSpeed`:
  - enum added: `1s`
  - affected methods:
    - `markPriceOfAllSymbolsOfAPair()` (`<pair>@markPrice@<updateSpeed>` stream)
    - `markPriceStream()` (`<symbol>@markPrice@<updateSpeed>` stream)
- Modified parameter `updateSpeed`:
  - enum added: `100ms`, `500ms`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
    - `diffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream)

### Removed (1)

#### WebSocket Streams

- `/<pair>@indexPrice()` (`<pair>@indexPrice` stream)

## 1.9.0 - 2026-06-30

### Changed (1)

#### REST API

- Modified response for `queryIndexPriceConstituents()` (`GET /dapi/v1/constituents`):
  - `constituents`.items: property `weight` added
  - `constituents`.items: property `price` added
  - `constituents`.items: item property `weight` added
  - `constituents`.items: item property `price` added

## 1.8.0 - 2026-06-16

### Changed (13)

- Modified response for `markPriceOfAllSymbolsOfAPair()` (`<pair>@markPrice@<updateSpeed>` stream):
  - items: property `st` added
  - items: item property `st` added

#### WebSocket Streams

- Modified response for `allBookTickersStream()` (`!bookTicker` stream):
  - property `st` added

- Modified response for `contractInfoStream()` (`!contractInfo` stream):
  - property `st` added

- Modified response for `allMarketLiquidationOrderStreams()` (`!forceOrder@arr` stream):
  - property `st` added

- Modified response for `allMarketMiniTickersStream()` (`!miniTicker@arr` stream):
  - items: property `st` added
  - items: item property `st` added

- Modified response for `allMarketTickersStreams()` (`!ticker@arr` stream):
  - items: property `st` added
  - items: item property `st` added

- Modified response for `aggregateTradeStreams()` (`<symbol>@aggTrade` stream):
  - property `st` added

- Modified response for `individualSymbolBookTickerStreams()` (`<symbol>@bookTicker` stream):
  - property `st` added

- Modified response for `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream):
  - property `st` added

- Modified response for `diffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream):
  - property `st` added

- Modified response for `markPriceStream()` (`<symbol>@markPrice@<updateSpeed>` stream):
  - property `st` added

- Modified response for `individualSymbolMiniTickerStream()` (`<symbol>@miniTicker` stream):
  - property `st` added

- Modified response for `individualSymbolTickerStreams()` (`<symbol>@ticker` stream):
  - property `st` added

## 1.7.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.6.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.5.0 - 2026-03-09

### Changed (1)

#### REST API

- Modified response for `exchange_information()` (`GET /dapi/v1/exchangeInfo`):
  - `symbols`.items: property `orderTypes` added
  - `symbols`.items: property `OrderType` deleted
  - `symbols`.items: item property `orderTypes` added
  - `symbols`.items: item property `OrderType` deleted

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (2)

#### REST API

- Modified response for `CancelMultipleOrders()` (`DELETE /dapi/v1/batchOrders`):
  - items: property `Pair` added
  - items: item property `Pair` added

- Modified response for `CurrentAllOpenOrders()` (`GET /dapi/v1/openOrders`):
  - items: property `Pair` added
  - items: item property `Pair` added

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Added (1)

#### REST API

- `PlaceMultipleOrders()` (`POST /dapi/v1/batchOrders`)

### Changed (3)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

#### REST API

- Modified response for `AutoCancelAllOpenOrders()` (`POST /dapi/v1/countdownCancelAll`) and `KeepaliveUserDataStreamExecute` (`PUT /dapi/v1/listenKey`) endpoints.

## 1.0.0 - 2025-12-17

- Initial release 
