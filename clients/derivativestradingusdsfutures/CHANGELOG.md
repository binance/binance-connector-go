### Changelog

## 1.4.0 - 2026-02-13

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `2.0.0`.

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
