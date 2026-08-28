### Changelog

## 1.19.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.18.1 - 2026-08-25

### Changed (1)

#### REST API

- Modified parameter `symbol`:
  - required: `true` → `false`
  - affected methods:
    - `allOrders()` (`GET /fapi/v1/allOrders`)

## 1.18.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.17.0 - 2026-08-18

### Changed (1)

- Updated dependencies

## 1.16.0 - 2026-08-18

### Changed (2)

#### REST API

- Modified response for `tradingSchedule()` (`GET /fapi/v1/tradingSchedule`):
  - `marketSchedules`: property `CN_EQUITY` added

- Modified response field `marketSchedules`:
  - property `CN_EQUITY` added
  - affected events:
    - `tradingScheduleResponse`

## 1.15.0 - 2026-08-14

### Changed (2)

#### WebSocket Streams

- Modified response field `a`:
  - property `S` added
  - affected events:
    - `UserDataStreamEventsResponse`
    - `accountUpdate`
- Modified response field `o`:
  - property `ia` added
  - affected events:
    - `UserDataStreamEventsResponse`
    - `algoUpdate`

## 1.14.0 - 2026-08-06

### Changed (3)

#### REST API

- Modified response for `allOrders()` (`GET /fapi/v1/allOrders`):
  - items: property `cumBase` added
  - items: property `pair` added
  - items: item property `cumBase` added
  - items: item property `pair` added

- Modified response for `usersForceOrders()` (`GET /fapi/v1/forceOrders`):
  - items: property `pair` added
  - items: property `cumBase` added
  - items: item property `pair` added
  - items: item property `cumBase` added

- Modified response for `accountTradeList()` (`GET /fapi/v1/userTrades`):
  - items: property `marginAsset` added
  - items: property `baseQty` added
  - items: property `pair` added
  - items: item property `marginAsset` added
  - items: item property `baseQty` added
  - items: item property `pair` added

## 1.13.0 - 2026-07-28

### Changed (9)

#### REST API

- Added parameter `modifyId`
  - affected methods:
    - `modifyOrder()` (`PUT /fapi/v1/order`)
- Modified parameter `batchOrders`:
  - items: property `modifyId` added
  - items: item property `modifyId` added
  - affected methods:
    - `modifyMultipleOrders()` (`PUT /fapi/v1/batchOrders`)
- Modified response for `modifyMultipleOrders()` (`PUT /fapi/v1/batchOrders`):
  - items: property `modifyId` added
  - items: item property `modifyId` added

- Modified response for `modifyOrder()` (`PUT /fapi/v1/order`):
  - property `modifyId` added

- Modified response for `getOrderModifyHistory()` (`GET /fapi/v1/orderAmendment`):
  - items.`amendment`: property `modifyId` added
  - items.`amendment`: property `modifyId` added

- Modified response for `tradingSchedule()` (`GET /fapi/v1/tradingSchedule`):
  - `marketSchedules`: property `HK_EQUITY` added

- Modified response for `getFundingRateHistory()` (`GET /fapi/v1/fundingRate`):
  - items: property `rateType` added
  - items: item property `rateType` added

#### WebSocket API

- Added parameter `modifyId`
  - affected methods:
    - `modifyOrder()` (`order.modify` method)
- Modified response for `modifyOrder()` (`order.modify` method):
  - `result`: property `modifyId` added

## 1.12.0 - 2026-07-15

### Changed (10)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

#### REST API

- Modified response for `cancelMultipleOrders()` (`DELETE /fapi/v1/batchOrders`):
  - items: property `cumQuote` deleted
  - items: item property `cumQuote` deleted

- Modified response for `placeMultipleOrders()` (`POST /fapi/v1/batchOrders`):
  - items: property `cumQuote` deleted
  - items: property `avgPrice` deleted
  - items: item property `cumQuote` deleted
  - items: item property `avgPrice` deleted

- Modified response for `modifyMultipleOrders()` (`PUT /fapi/v1/batchOrders`):
  - items: property `avgPrice` deleted
  - items: property `cumBase` deleted
  - items: item property `avgPrice` deleted
  - items: item property `cumBase` deleted

- Modified response for `cancelOrder()` (`DELETE /fapi/v1/order`):
  - property `cumQuote` deleted
  - property `avgPrice` deleted

- Modified response for `newOrder()` (`POST /fapi/v1/order`):
  - property `avgPrice` deleted
  - property `cumQuote` deleted

- Modified response for `modifyOrder()` (`PUT /fapi/v1/order`):
  - property `cumBase` deleted
  - property `avgPrice` deleted

#### WebSocket API

- Modified response for `cancelOrder()` (`order.cancel` method):
  - `result`: property `cumQuote` deleted

- Modified response for `modifyOrder()` (`order.modify` method):
  - `result`: property `cumQuote` deleted
  - `result`: property `avgPrice` deleted

- Modified response for `newOrder()` (`order.place` method):
  - `result`: property `cumQuote` deleted
  - `result`: property `avgPrice` deleted

## 1.11.0 - 2026-07-14

### Changed (57)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

#### REST API

- Modified parameter `algoType`:
  - enum added: `CONDITIONAL`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
- Modified parameter `batchOrders`:
  - items.`goodTillDate`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: enum added: `true`, `false`
  - items.`selfTradePreventionMode`: enum added: `NONE`
  - items.`type`: enum added: `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`
  - items.`goodTillDate`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: enum added: `true`, `false`
  - items.`selfTradePreventionMode`: enum added: `NONE`
  - items.`type`: enum added: `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`
  - affected methods:
    - `placeMultipleOrders()` (`POST /fapi/v1/batchOrders`)
- Modified parameter `batchOrders`:
  - items: property `timestamp` added
  - items.`orderId`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`quantity`: type `string` → `number`
  - items.`recvWindow`: type `string` → `integer`
  - items.`stopPrice`: type `string` → `number`
  - items: item property `timestamp` added
  - items.`orderId`: type `string` → `integer`
  - items.`price`: type `string` → `number`
  - items.`priceMatch`: enum removed: `NONE`
  - items.`quantity`: type `string` → `number`
  - items.`recvWindow`: type `string` → `integer`
  - items.`stopPrice`: type `string` → `number`
  - affected methods:
    - `modifyMultipleOrders()` (`PUT /fapi/v1/batchOrders`)
- Modified parameter `closePosition`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `contractType`:
  - enum removed: `CURRENT_MONTH`, `NEXT_MONTH`, `PERPETUAL_DELIVERING`
  - enum added: `TRADIFI_PERPETUAL`
  - affected methods:
    - `continuousContractKlineCandlestickData()` (`GET /fapi/v1/continuousKlines`)
- Modified parameter `contractType`:
  - enum removed: `CURRENT_MONTH`, `NEXT_MONTH`, `PERPETUAL_DELIVERING`
  - affected methods:
    - `basis()` (`GET /futures/data/basis`)
- Modified parameter `incomeType`:
  - enum added: `TRANSFER`, `WELCOME_BONUS`, `REALIZED_PNL`, `FUNDING_FEE`, `COMMISSION`, `INSURANCE_CLEAR`, `REFERRAL_KICKBACK`, `COMMISSION_REBATE`, `API_REBATE`, `CONTEST_REWARD`, `CROSS_COLLATERAL_TRANSFER`, `OPTIONS_PREMIUM_FEE`, `OPTIONS_SETTLE_PROFIT`, `INTERNAL_TRANSFER`, `AUTO_EXCHANGE`, `DELIVERED_SETTELMENT`, `COIN_SWAP_DEPOSIT`, `COIN_SWAP_WITHDRAW`, `POSITION_LIMIT_INCREASE_FEE`, `STRATEGY_UMFUTURES_TRANSFER`, `FEE_RETURN`, `BFUSD_REWARD`
  - affected methods:
    - `getIncomeHistory()` (`GET /fapi/v1/income`)
- Modified parameter `interval`:
  - enum removed: `1s`
  - affected methods:
    - `continuousContractKlineCandlestickData()` (`GET /fapi/v1/continuousKlines`)
    - `indexPriceKlineCandlestickData()` (`GET /fapi/v1/indexPriceKlines`)
    - `klineCandlestickData()` (`GET /fapi/v1/klines`)
    - `markPriceKlineCandlestickData()` (`GET /fapi/v1/markPriceKlines`)
    - `premiumIndexKlineData()` (`GET /fapi/v1/premiumIndexKlines`)
- Modified parameter `orderIdList`:
  - maxLength `null` → `10`
  - affected methods:
    - `cancelMultipleOrders()` (`DELETE /fapi/v1/batchOrders`)
- Modified parameter `origClientOrderIdList`:
  - maxLength `null` → `10`
  - affected methods:
    - `cancelMultipleOrders()` (`DELETE /fapi/v1/batchOrders`)
- Modified parameter `positionSide`:
  - enum removed: `BOTH`, `LONG`, `SHORT`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `newOrder()` (`POST /fapi/v1/order`)
    - `modifyIsolatedPositionMargin()` (`POST /fapi/v1/positionMargin`)
- Modified parameter `priceMatch`:
  - enum removed: `NONE`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `newOrder()` (`POST /fapi/v1/order`)
    - `modifyOrder()` (`PUT /fapi/v1/order`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `priceProtect`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `reduceOnly`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `newOrder()` (`POST /fapi/v1/order`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `selfTradePreventionMode`:
  - enum added: `NONE`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `type`:
  - enum added: `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
- Modified parameter `type`:
  - enum added: `LIMIT`, `MARKET`, `STOP`, `STOP_MARKET`, `TAKE_PROFIT`, `TAKE_PROFIT_MARKET`, `TRAILING_STOP_MARKET`
  - affected methods:
    - `newOrder()` (`POST /fapi/v1/order`)
    - `testOrder()` (`POST /fapi/v1/order/test`)
- Modified parameter `type`:
  - type `string` → `integer`
  - affected methods:
    - `modifyIsolatedPositionMargin()` (`POST /fapi/v1/positionMargin`)
- Modified parameter `type`:
  - type `integer` → `string`
  - affected methods:
    - `getPositionMarginChangeHistory()` (`GET /fapi/v1/positionMargin/history`)
- Modified response for `orderBook()` (`GET /fapi/v1/depth`):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `queryInsuranceFundBalanceSnapshot()` (`GET /fapi/v1/insuranceBalance`):
  - oneOf modified

- Modified response for `notionalAndLeverageBrackets()` (`GET /fapi/v1/leverageBracket`):
  - oneOf modified

- Modified response for `queryOrder()` (`GET /fapi/v1/order`):
  - property `priceMatch` added
  - property `goodTillDate` added
  - property `selfTradePreventionMode` added

- Modified response for `markPrice()` (`GET /fapi/v1/premiumIndex`):
  - oneOf modified

- Modified response for `rpiOrderBook()` (`GET /fapi/v1/rpiDepth`):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `adlRisk()` (`GET /fapi/v1/symbolAdlRisk`):
  - oneOf modified

- Modified response for `ticker24hrPriceChangeStatistics()` (`GET /fapi/v1/ticker/24hr`):
  - oneOf modified

- Modified response for `symbolOrderBookTicker()` (`GET /fapi/v1/ticker/bookTicker`):
  - oneOf modified

- Modified response for `symbolPriceTicker()` (`GET /fapi/v1/ticker/price`):
  - oneOf modified

- Modified response for `symbolPriceTickerV2()` (`GET /fapi/v2/ticker/price`):
  - oneOf modified

- Modified response for `longShortRatio()` (`GET /futures/data/globalLongShortAccountRatio`):
  - items.`timestamp`: type `string` → `integer`
  - items.`timestamp`: type `string` → `integer`

- Modified response for `openInterestStatistics()` (`GET /futures/data/openInterestHist`):
  - items.`timestamp`: type `string` → `integer`
  - items.`timestamp`: type `string` → `integer`

- Modified response for `takerBuySellVolume()` (`GET /futures/data/takerlongshortRatio`):
  - items.`timestamp`: type `string` → `integer`
  - items.`timestamp`: type `string` → `integer`

- Modified response for `topTraderLongShortRatioAccounts()` (`GET /futures/data/topLongShortAccountRatio`):
  - items.`timestamp`: type `string` → `integer`
  - items.`timestamp`: type `string` → `integer`

- Modified response for `topTraderLongShortRatioPositions()` (`GET /futures/data/topLongShortPositionRatio`):
  - items.`timestamp`: type `string` → `integer`
  - items.`timestamp`: type `string` → `integer`

- Marked `symbolPriceTicker()` (`GET /fapi/v1/ticker/price`) as deprecated.

#### WebSocket API

- Modified parameter `algoType`:
  - enum added: `CONDITIONAL`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `closePosition`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `priceMatch`:
  - enum removed: `NONE`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
    - `modifyOrder()` (`order.modify` method)
    - `newOrder()` (`order.place` method)
- Modified parameter `priceProtect`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `reduceOnly`:
  - enum added: `true`, `false`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
    - `newOrder()` (`order.place` method)
- Modified parameter `selfTradePreventionMode`:
  - enum added: `NONE`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `selfTradePreventionMode`:
  - enum added: `NONE`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Modified parameter `timeInForce`:
  - enum removed: `GTX`, `GTD`, `RPI`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `type`:
  - enum added: `STOP_MARKET`, `TAKE_PROFIT_MARKET`, `STOP`, `TAKE_PROFIT`, `TRAILING_STOP_MARKET`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified parameter `type`:
  - enum added: `LIMIT`, `MARKET`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Modified response for `orderBook()` (`depth` method):
  - property `bids` added
  - property `asks` added
  - `result`: property `asks` deleted
  - `result`: property `bids` deleted

- Modified response for `queryOrder()` (`order.status` method):
  - `result`: property `priceMatch` added
  - `result`: property `goodTillDate` added
  - `result`: property `selfTradePreventionMode` added

- Modified response for `symbolOrderBookTicker()` (`ticker.book` method):
  - oneOf modified

- Modified response for `symbolPriceTicker()` (`ticker.price` method):
  - oneOf modified

#### WebSocket Streams

- Modified parameter `contractType`:
  - enum added: `perpetual`, `current_quarter`, `next_quarter`, `tradifi_perpetual`
  - affected methods:
    - `continuousContractKlineCandlestickStreams()` (`<pair>_<contractType>@continuousKline_<interval>` stream)
- Modified parameter `interval`:
  - enum added: `1s`, `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1M`
  - affected methods:
    - `continuousContractKlineCandlestickStreams()` (`<pair>_<contractType>@continuousKline_<interval>` stream)
- Modified parameter `interval`:
  - enum added: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1M`
  - affected methods:
    - `klineCandlestickStreams()` (`<symbol>@kline_<interval>` stream)
- Modified parameter `levels`:
  - type `integer` → `string`
  - enum added: `5`, `10`, `20`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
- Modified parameter `updateSpeed`:
  - enum added: `1s`
  - affected methods:
    - `markPriceStreamForAllMarket()` (`!markPrice@arr@<updateSpeed>` stream)
    - `markPriceStream()` (`<symbol>@markPrice@<updateSpeed>` stream)
- Modified parameter `updateSpeed`:
  - enum added: `100ms`, `500ms`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
    - `diffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream)

## 1.11.0 - 2026-06-30

### Changed (1)

#### WebSocket Streams

- Modified response for `individualSymbolBookTickerStreams()` (`<symbol>@bookTicker` stream):
  - property `ps` added

## 1.10.0 - 2026-06-16

### Changed (24)

#### REST API

- Modified response for `assetIndex()` (`GET /fapi/v1/assetIndex`):
  - oneOf added 2 schema(s)
  - oneOf removed 2 schema(s)

- Modified response for `compressedAggregateTradesList()` (`GET /fapi/v1/aggTrades`):
  - items: property `nq` added
  - items: item property `nq` added

- Modified response for `queryAlgoOrder()` (`GET /fapi/v1/algoOrder`):
  - property `actualQty` added
  - property `actualType` added
  - property `slPrice` deleted
  - property `slTriggerPrice` deleted
  - property `tpPrice` deleted
  - property `tpTriggerPrice` deleted

- Modified response for `tradingSchedule()` (`GET /fapi/v1/tradingSchedule`):
  - `marketSchedules`: property `KR_EQUITY` added

#### WebSocket API

- Deleted parameter `activationPrice`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Deleted parameter `callbackRate`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Deleted parameter `closePosition`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Deleted parameter `priceProtect`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Deleted parameter `stopPrice`
  - affected methods:
    - `newOrder()` (`order.place` method)
- Deleted parameter `workingType`
  - affected methods:
    - `newOrder()` (`order.place` method)

#### WebSocket Streams

- Modified response for `allBookTickersStream()` (`!bookTicker` stream):
  - property `st` added
  - property `ps` added

- Modified response for `contractInfoStream()` (`!contractInfo` stream):
  - property `st` added
  - property `ps` deleted

- Modified response for `allMarketLiquidationOrderStreams()` (`!forceOrder@arr` stream):
  - property `ps` added
  - property `st` added

- Modified response for `markPriceStreamForAllMarket()` (`!markPrice@arr@<updateSpeed>` stream):
  - items: property `st` added
  - items: item property `st` added

- Modified response for `allMarketMiniTickersStream()` (`!miniTicker@arr` stream):
  - items: property `ps` added
  - items: property `st` added
  - items: item property `ps` added
  - items: item property `st` added

- Modified response for `allMarketTickersStreams()` (`!ticker@arr` stream):
  - items: property `ps` added
  - items: property `st` added
  - items: item property `ps` added
  - items: item property `st` added

- Modified response for `aggregateTradeStreams()` (`<symbol>@aggTrade` stream):
  - property `st` added

- Modified response for `individualSymbolBookTickerStreams()` (`<symbol>@bookTicker` stream):
  - property `st` added

- Modified response for `partialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream):
  - property `ps` added
  - property `st` added

- Modified response for `diffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream):
  - property `ps` added
  - property `st` added

- Modified response for `markPriceStream()` (`<symbol>@markPrice@<updateSpeed>` stream):
  - property `st` added

- Modified response for `individualSymbolMiniTickerStream()` (`<symbol>@miniTicker` stream):
  - property `st` added
  - property `ps` added

- Modified response for `rpiDiffBookDepthStreams()` (`<symbol>@rpiDepth@500ms` stream):
  - property `ps` added
  - property `st` added

- Modified response for `individualSymbolTickerStreams()` (`<symbol>@ticker` stream):
  - property `st` added
  - property `ps` added

## 1.9.0 - 2026-05-20

### Changed (1)

#### WebSocket Streams

- Modified response for `Listenkeyexpired`:
  - `E`: type `string` → `int64`

## 1.8.0 - 2026-04-20

### Changed (3)

#### REST API

- Deleted parameter `page`
  - affected methods:
    - `queryAllAlgoOrders()` (`GET /fapi/v1/allAlgoOrders`)
- Modified parameter `interval`:
  - enum added: `1s`
  - affected methods:
    - `continuousContractKlineCandlestickData()` (`GET /fapi/v1/continuousKlines`)
    - `indexPriceKlineCandlestickData()` (`GET /fapi/v1/indexPriceKlines`)
    - `klineCandlestickData()` (`GET /fapi/v1/klines`)
    - `markPriceKlineCandlestickData()` (`GET /fapi/v1/markPriceKlines`)
    - `premiumIndexKlineData()` (`GET /fapi/v1/premiumIndexKlines`)
- Modified parameter `limit`:
  - required: `true` → `false`
  - affected methods:
    - `basis()` (`GET /futures/data/basis`)

## 1.7.0 - 2026-03-26

### Changed (3)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

#### WebSocket Streams

- Modified response for `markPriceStreamForAllMarket()` (`!markPrice@arr@<updateSpeed>` stream):
  - items: property `ap` added
  - items: item property `ap` added

- Modified response for `markPriceStream()` (`<symbol>@markPrice@<updateSpeed>` stream):
  - property `ap` added

## 1.6.0 - 2026-03-16

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

#### REST API

- Modified response for `queryOrder()` (`GET /fapi/v1/order`):
  - property `workingType` added
  - property `clientOrderId` added
  - property `orderId` added
  - property `time` added
  - property `priceProtect` added
  - property `positionSide` added
  - property `cumQuote` added
  - property `activatePrice` added
  - property `avgPrice` added
  - property `price` added
  - property `reduceOnly` added
  - property `timeInForce` added
  - property `symbol` added
  - property `updateTime` added
  - property `type` added
  - property `origType` added
  - property `origQty` added
  - property `stopPrice` added
  - property `closePosition` added
  - property `priceRate` added
  - property `side` added
  - property `executedQty` added
  - property `result` deleted
  - property `id` deleted
  - `status`: type `integer` → `string`

## 1.5.0 - 2026-03-09

### Changed (2)

#### REST API

- Modified response for `exchange_information()` (`GET /fapi/v1/exchangeInfo`):
  - `symbols`.items: property `orderTypes` added
  - `symbols`.items: property `OrderType` deleted
  - `symbols`.items: item property `orderTypes` added
  - `symbols`.items: item property `OrderType` deleted

- Modified response for `cancel_order()` (`DELETE /fapi/v1/order`):
  - property `avgPrice` added

## 1.4.0 - 2026-02-13

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

#### REST API

- Modified response for `cancelOrder()` (`DELETE /fapi/v1/order`):
  - property `avgPrice` deleted

## 1.3.0 - 2026-01-29

### Changed (6)

#### REST API

- Added parameter `newOrderRespType`
  - affected methods:
    - `newAlgoOrder()` (`POST /fapi/v1/algoOrder`)
- Modified parameter `batchOrders`:
  - items: property `stopPrice` added
  - items: item property `stopPrice` added
  - affected methods:
    - `modifyMultipleOrders()` (`PUT /fapi/v1/batchOrders`)
- Modified response for `placeMultipleOrders()` (`POST /fapi/v1/batchOrders`):
  - items: property `closePosition` added
  - items: item property `closePosition` added

- Modified response for `queryOrder()` (`GET /fapi/v1/order`):
  - property `id` added
  - property `result` added
  - property `cumQuote` deleted
  - property `origQty` deleted
  - property `orderId` deleted
  - property `priceRate` deleted
  - property `type` deleted
  - property `workingType` deleted
  - property `clientOrderId` deleted
  - property `timeInForce` deleted
  - property `symbol` deleted
  - property `updateTime` deleted
  - property `activatePrice` deleted
  - property `closePosition` deleted
  - property `reduceOnly` deleted
  - property `selfTradePreventionMode` deleted
  - property `executedQty` deleted
  - property `origType` deleted
  - property `goodTillDate` deleted
  - property `positionSide` deleted
  - property `stopPrice` deleted
  - property `price` deleted
  - property `avgPrice` deleted
  - property `priceProtect` deleted
  - property `priceMatch` deleted
  - property `time` deleted
  - property `side` deleted
  - `status`: type `string` → `integer`

#### WebSocket API

- Added parameter `newOrderRespType`
  - affected methods:
    - `newAlgoOrder()` (`algoOrder.place` method)
- Modified response for `positionInformationV2()` (`v2/account.position` method):
  - `result`.items: property `unRealizedProfit` added
  - `result`.items: property `unrealizedProfit` deleted
  - `result`.items: item property `unRealizedProfit` added
  - `result`.items: item property `unrealizedProfit` deleted

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (5)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

#### REST API

- Modified parameter `activationPrice` name to `activatePrice`
  - affected methods:
    - `NewAlgoOrder()` (`POST /fapi/v1/algoOrder`)

#### WebSocket API

- Modified parameter `activationPrice` name to `activatePrice`
  - affected methods:
    - `NewAlgoOrder()` (`algoOrder.place`)

#### WebSocket Streams

- Modified response for `AggregateTradeStreams()` method

## 1.0.0 - 2025-12-17

- Initial release 
