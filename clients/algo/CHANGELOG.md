### Changelog

## 1.15.0 - 2026-10-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.11.0`.

## 1.14.0 - 2026-10-15

### Changed (3)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.10.0`.
- Updated dependencies.
- Raised the minimum Go version to `1.26`.

## 1.13.0 - 2026-10-12

### Added (1)

- `volumeParticipation()` (`POST /sapi/v1/algo/spot/newOrderVp`)

### Changed (4)

- Modified response for `queryHistoricalAlgoOrdersSpotAlgo()` (`GET /sapi/v1/algo/spot/historicalOrders`):
  - `orders`.items: property `duration` added
  - `orders`.items: item property `duration` added

- Modified response for `queryCurrentAlgoOpenOrdersSpotAlgo()` (`GET /sapi/v1/algo/spot/openOrders`):
  - `orders`.items: property `duration` added
  - `orders`.items: item property `duration` added

- Modified response field `orders`:
  - items: property `duration` added
  - items: item property `duration` added
  - affected events:
    - `queryCurrentAlgoOpenOrdersSpotAlgoResponse`
    - `queryHistoricalAlgoOrdersSpotAlgoResponse`
- Added response schema `volumeParticipationResponse`

## 1.12.0 - 2026-09-28

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.9.0`.

## 1.11.1 - 2026-09-09

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.1`.

## 1.11.0 - 2026-09-02

### Changed (2)

#### REST API

- Modified parameter `limitPrice`:
  - type `float32` → `float64`
  - affected methods:
    - `timeWeightedAveragePriceFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderTwap`)
    - `timeWeightedAveragePriceSpotAlgo()` (`POST /sapi/v1/algo/spot/newOrderTwap`)
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)
- Modified parameter `quantity`:
  - type `float32` → `float64`
  - affected methods:
    - `timeWeightedAveragePriceFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderTwap`)
    - `timeWeightedAveragePriceSpotAlgo()` (`POST /sapi/v1/algo/spot/newOrderTwap`)
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)

## 1.10.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.9.0 - 2026-08-25

### Changed (2)

- Added parameter `clientAlgoId`
  - affected methods:
    - `cancelAlgoOrderFutureAlgo()` (`DELETE /sapi/v1/algo/futures/order`)
    - `cancelAlgoOrderSpotAlgo()` (`DELETE /sapi/v1/algo/spot/order`)
- Modified parameter `algoId`:
  - required: `true` → `false`
  - affected methods:
    - `cancelAlgoOrderFutureAlgo()` (`DELETE /sapi/v1/algo/futures/order`)
    - `cancelAlgoOrderSpotAlgo()` (`DELETE /sapi/v1/algo/spot/order`)

## 1.8.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.7.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.6.0 - 2026-07-14

### Changed (5)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Modified parameter `clientAlgoId`:
  - minLength `0` → `32`
  - maxLength `null` → `32`
  - affected methods:
    - `timeWeightedAveragePriceFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderTwap`)
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)
    - `timeWeightedAveragePriceSpotAlgo()` (`POST /sapi/v1/algo/spot/newOrderTwap`)
- Modified parameter `positionSide`:
  - enum added: `BOTH`, `LONG`, `SHORT`
  - affected methods:
    - `timeWeightedAveragePriceFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderTwap`)
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)
- Modified parameter `side`:
  - enum added: `BUY`, `SELL`
  - affected methods:
    - `queryHistoricalAlgoOrdersFutureAlgo()` (`GET /sapi/v1/algo/futures/historicalOrders`)
    - `timeWeightedAveragePriceFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderTwap`)
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)
    - `queryHistoricalAlgoOrdersSpotAlgo()` (`GET /sapi/v1/algo/spot/historicalOrders`)
    - `timeWeightedAveragePriceSpotAlgo()` (`POST /sapi/v1/algo/spot/newOrderTwap`)
- Modified parameter `urgency`:
  - enum added: `LOW`, `MEDIUM`, `HIGH`
  - affected methods:
    - `volumeParticipationFutureAlgo()` (`POST /sapi/v1/algo/futures/newOrderVp`)

## 1.5.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.4.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.3.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.2.0 - 2026-01-29

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
