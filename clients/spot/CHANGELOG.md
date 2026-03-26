### Changelog

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
