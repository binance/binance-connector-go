### Changelog

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
