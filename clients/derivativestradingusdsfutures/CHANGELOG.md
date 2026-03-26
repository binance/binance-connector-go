### Changelog

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
