### Changelog

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