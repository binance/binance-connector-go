### Changelog

## 1.7.0 - 2026-09-02

### Changed (4)

#### REST API

- Modified response for `getOtcBlocktradeDetail()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/detail`):
  - `orderData`.`price`: type `float32` → `float64`
- Modified response for `getQuote()` (`POST /sapi/v1/w3w/wallet/prediction/trade/get-quote`):
  - `averagePrice`: type `float32` → `float64`
  - `lastPrice`: type `float32` → `float64`
  - `priceImpact`: type `float32` → `float64`
- Modified response for `listOtcBlocktrades()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/list`):
  - `blocktrades`.items.`price`: type `float32` → `float64`
- Modified response for `previewOtcBlocktrade()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/preview`):
  - `orderData`.`price`: type `float32` → `float64`

## 1.6.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.5.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.4.0 - 2026-08-18

### Changed (1)

- Updated dependencies

## 1.3.0 - 2026-08-18

### Added (10)

- `applyMmDeposit()` (`POST /sapi/v1/w3w/wallet/prediction/deposit/apply`)
- `applyMmWithdraw()` (`POST /sapi/v1/w3w/wallet/prediction/withdraw/apply`)
- `createOtcBlocktrade()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/create`)
- `fulfilOtcBlocktrade()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/fulfil`)
- `getOtcBlocktradeDetail()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/detail`)
- `getOtcBlocktradeEvents()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/events`)
- `getOtcReservedBalances()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/reserved-balances`)
- `listOtcBlocktrades()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/list`)
- `previewOtcBlocktrade()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/preview`)
- `removeOtcBlocktrades()` (`POST /sapi/v1/w3w/wallet/prediction/otc/blocktrade/remove`)

### Changed (10)

- Added response schema `getOtcBlocktradeEventsResponse`
- Added response schema `previewOtcBlocktradeResponse`
- Added response schema `createOtcBlocktradeResponse`
- Added response schema `fulfilOtcBlocktradeResponse`
- Added response schema `applyMmDepositResponse`
- Added response schema `removeOtcBlocktradesResponse`
- Added response schema `getOtcReservedBalancesResponse`
- Added response schema `listOtcBlocktradesResponse`
- Added response schema `applyMmWithdrawResponse`
- Added response schema `getOtcBlocktradeDetailResponse`

## 1.2.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.1.0 - 2026-07-14

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

## 1.0.0 - 2026-06-29

- Initial release 
