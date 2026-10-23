### Changelog

## 1.5.0 - 2026-10-23

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.11.0`.
- Fixed WebSocket Streams names being converted to lowercase: the case of stream names is now kept.

## 1.4.0 - 2026-10-15

### Changed (3)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.10.0`.
- Updated dependencies.
- Raised the minimum Go version to `1.26`.

## 1.3.0 - 2026-09-28

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.9.0`.

## 1.2.1 - 2026-09-09

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.1`.

## 1.2.0 - 2026-09-02

### Added (1)

#### WebSocket Streams

- Added parameter `id`
  - affected methods:
    - `calendarStream()` (`calendar` stream)
    - `klineStream()` (`<symbol>@kline_<interval>` stream)
    - `orderReportStream()` (`<listenKey>@orderReport` stream)
    - `priceStream()` (`price` stream)
    - `quoteStream()` (`<symbol>@quote` stream)
    - `tradabilityStream()` (`<symbol>@tradability` stream)
    - `tradingStatusStream()` (`<symbol>@tradingStatus` stream)

### Changed (1)

#### WebSocket Streams

- Modified response for `orderReportStream()` (`<listenKey>@orderReport` stream):
  - `FN`: type `float32` → `float64`
  - `N`: type `float32` → `float64`
  - `Q`: type `float32` → `float64`
  - `Z`: type `float32` → `float64`
  - `fq`: type `float32` → `float64`
  - `p`: type `float32` → `float64`
  - `tc`: type `float32` → `float64`

## 1.1.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.0.0 - 2026-08-19

- Initial release 
