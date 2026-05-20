### Changelog

## 1.8.0 - 2026-05-20

### Added (3)

#### REST API

- `historicalBlockTrades()` (`GET /api/v3/historicalBlockTrades`)

#### WebSocket API

- `blockTradesHistorical()` (`blockTrades.historical` method)

#### WebSocket Streams

- `blockTrade()` (`<symbol>@blockTrade` stream)

### Changed (3)

#### REST API

- Modified parameter `selfTradePreventionMode`:
  - enum added: `TRANSFER`
  - affected methods:
    - `newOrder()` (`POST /api/v3/order`)
    - `orderCancelReplace()` (`POST /api/v3/order/cancelReplace`)
    - `orderOco()` (`POST /api/v3/order/oco`)
    - `orderTest()` (`POST /api/v3/order/test`)
    - `orderListOco()` (`POST /api/v3/orderList/oco`)
    - `orderListOpo()` (`POST /api/v3/orderList/opo`)
    - `orderListOpoco()` (`POST /api/v3/orderList/opoco`)
    - `orderListOto()` (`POST /api/v3/orderList/oto`)
    - `orderListOtoco()` (`POST /api/v3/orderList/otoco`)
    - `sorOrder()` (`POST /api/v3/sor/order`)
    - `sorOrderTest()` (`POST /api/v3/sor/order/test`)
- Marked `orderOco()` (`POST /api/v3/order/oco`) as deprecated.

#### WebSocket API

- Modified parameter `selfTradePreventionMode`:
  - enum added: `TRANSFER`
  - affected methods:
    - `orderCancelReplace()` (`order.cancelReplace` method)
    - `orderPlace()` (`order.place` method)
    - `orderTest()` (`order.test` method)
    - `orderListPlace()` (`orderList.place` method)
    - `orderListPlaceOco()` (`orderList.place.oco` method)
    - `orderListPlaceOpo()` (`orderList.place.opo` method)
    - `orderListPlaceOpoco()` (`orderList.place.opoco` method)
    - `orderListPlaceOto()` (`orderList.place.oto` method)
    - `orderListPlaceOtoco()` (`orderList.place.otoco` method)
    - `sorOrderPlace()` (`sor.order.place` method)
    - `sorOrderTest()` (`sor.order.test` method)

## 1.7.0 - 2026-04-20

### Changed (1)

#### WebSocket API

- Modified response for `referencePrice()` (`referencePrice` method):
  - `result`: property `code` added
  - `result`: property `msg` added

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Added (7)

#### REST API

- `executionRules()` (`GET /api/v3/executionRules`)
- `referencePrice()` (`GET /api/v3/referencePrice`)
- `referencePriceCalculation()` (`GET /api/v3/referencePrice/calculation`)

#### WebSocket API

- `executionRules()` (`executionRules` method)
- `referencePrice()` (`referencePrice` method)
- `referencePriceCalculation()` (`referencePrice.calculation` method)

#### WebSocket Streams

- `referencePrice()` (`<symbol>@referencePrice` stream)

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (1)

#### WebSocket API

- Added parameter `recvWindow`
  - affected methods:
    - `userDataStreamSubscribeSignature()` (`userDataStream.subscribe.signature` method)

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
