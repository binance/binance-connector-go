### Changelog

## 1.7.0 - 2026-04-20

### Added (7)

#### REST API

- `cancelAllUmAlgoOpenOrders()` (`DELETE /papi/v1/um/algo/allOpenOrders`)
- `cancelUmAlgoOrder()` (`DELETE /papi/v1/um/algo/order`)
- `newUmAlgoOrder()` (`POST /papi/v1/um/algo/order`)
- `queryAllCurrentUmOpenAlgoOrders()` (`GET /papi/v1/um/algo/openAlgoOrders`)
- `queryCurrentUmOpenAlgoOrder()` (`GET /papi/v1/um/algo/algoOrder`)
- `queryUmAlgoOrderHistory()` (`GET /papi/v1/um/algo/allAlgoOrders`)
- `futuresTradfiPerpsContract()` (`POST /papi/v1/um/stock/contract`)

### Changed (8)

#### REST API

- Marked `cancelAllUmOpenConditionalOrders()` (`DELETE /papi/v1/um/conditional/allOpenOrders`) as deprecated.
- Marked `cancelUmConditionalOrder()` (`DELETE /papi/v1/um/conditional/order`) as deprecated.
- Marked `newUmConditionalOrder()` (`POST /papi/v1/um/conditional/order`) as deprecated.
- Marked `queryAllCurrentUmOpenConditionalOrders()` (`GET /papi/v1/um/conditional/openOrders`) as deprecated.
- Marked `queryAllUmConditionalOrders()` (`GET /papi/v1/um/conditional/allOrders`) as deprecated.
- Marked `queryCurrentUmOpenConditionalOrder()` (`GET /papi/v1/um/conditional/openOrder`) as deprecated.
- Marked `queryUmConditionalOrderHistory()` (`GET /papi/v1/um/conditional/orderHistory`) as deprecated.
- Modified response for `getUmIncomeHistory()` (`GET /papi/v1/um/income`):
  - items.`tranId`: type `string` → `integer`
  - items.`tranId`: type `string` → `integer`

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

#### REST API

- Modified parameter `StrategyType`:
  - enum added: `LIMIT_MAKER`
  - affected methods:
    - `NewCmConditionalOrder()` (`POST /papi/v1/cm/conditional/order`)
    - `NewUmConditionalOrder()` (`POST /papi/v1/um/conditional/order`)

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

## 1.0.0 - 2025-12-17

- Initial release 
