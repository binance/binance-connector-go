### Changelog

## 1.11.0 - 2026-09-02

### Changed (5)

#### REST API

- Modified response for `earningsList()` (`GET /sapi/v1/mining/payment/list`):
  - `data`.`accountProfits`.items.`profitAmount`: type `float32` → `float64`
  - `data`.`accountProfits`.items.`transferAmount`: type `float32` → `float64`
- Modified response for `extraBonusList()` (`GET /sapi/v1/mining/payment/other`):
  - `data`.`otherProfits`.items.`profitAmount`: type `float32` → `float64`
- Modified response for `hashrateResaleDetail()` (`GET /sapi/v1/mining/hash-transfer/profit/details`):
  - `data`.`profitTransferDetails`.items.`amount`: type `float32` → `float64`
- Modified response for `miningAccountEarning()` (`GET /sapi/v1/mining/payment/uid`):
  - `data`.`accountProfits`.items.`amount`: type `float32` → `float64`
- Modified response for `requestForMinerList()` (`GET /sapi/v1/mining/worker/list`):
  - `data`.`workerDatas`.items.`dayHashRate`: type `float32` → `float64`

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

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Modified response for `statisticList()` (`GET /sapi/v1/mining/statistics/user/status`):
  - `data`.`profitToday`: property `BSV` deleted
  - `data`.`profitToday`: property `BTC` deleted
  - `data`.`profitToday`: property `BCH` deleted
  - `data`.`profitYesterday`: property `BSV` deleted
  - `data`.`profitYesterday`: property `BTC` deleted
  - `data`.`profitYesterday`: property `BCH` deleted

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `2.1.0`.

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (2)

- Deleted parameter `userName`
  - affected methods:
    - `hashrateResaleDetail()` (`GET /sapi/v1/mining/hash-transfer/profit/details`)
- Modified response for `hashrateResaleList()` (`GET /sapi/v1/mining/hash-transfer/config/details/list`):
  - `data`.`configDetails`.items: property `type` added
  - `data`.`configDetails`.items: item property `type` added

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
