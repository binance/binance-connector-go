### Changelog

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