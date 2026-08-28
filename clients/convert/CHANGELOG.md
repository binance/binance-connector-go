### Changelog

## 1.10.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.9.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.8.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.7.0 - 2026-07-14

### Changed (5)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Modified parameter `expiredType`:
  - enum added: `1_D`, `3_D`, `7_D`, `30_D`
  - affected methods:
    - `placeLimitOrder()` (`POST /sapi/v1/convert/limit/placeOrder`)
- Modified parameter `side`:
  - enum added: `BUY`, `SELL`
  - affected methods:
    - `placeLimitOrder()` (`POST /sapi/v1/convert/limit/placeOrder`)
- Modified parameter `validTime`:
  - enum added: `10s`, `30s`, `1m`
  - affected methods:
    - `sendQuoteRequest()` (`POST /sapi/v1/convert/getQuote`)
- Modified parameter `walletType`:
  - enum added: `SPOT`, `FUNDING`, `EARN`, `SPOT_FUNDING`, `FUNDING_EARN`, `SPOT_FUNDING_EARN`, `SPOT_EARN`
  - affected methods:
    - `sendQuoteRequest()` (`POST /sapi/v1/convert/getQuote`)
    - `placeLimitOrder()` (`POST /sapi/v1/convert/limit/placeOrder`)

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (1)

- Modified response for `PlaceLimitOrder()` (`POST /sapi/v1/convert/limit/placeOrder`):
  - property `Status` added
  - property `OrderId` added
  - property `FromAmount` deleted
  - property `InverseRatio` deleted
  - property `QuoteId` deleted
  - property `Ratio` deleted
  - property `ToAmount` deleted
  - property `ValidTimestamp` deleted

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
