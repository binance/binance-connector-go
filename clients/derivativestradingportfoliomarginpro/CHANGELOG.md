### Changelog

## 1.9.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.8.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.7.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.6.0 - 2026-07-14

### Changed (7)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

#### REST API

- Modified parameter `asset`:
  - enum added: `LDUSDT`, `RWUSD`
  - affected methods:
    - `transferLdusdtRwusdForPortfolioMargin()` (`POST /sapi/v1/portfolio/earn-asset-transfer`)
- Modified parameter `autoRepay`:
  - enum added: `true`, `false`
  - affected methods:
    - `changeAutoRepayFuturesStatus()` (`POST /sapi/v1/portfolio/repay-futures-switch`)
- Modified parameter `deltaEnabled`:
  - enum added: `true`, `false`
  - affected methods:
    - `switchDeltaMode()` (`POST /sapi/v1/portfolio/delta-mode`)
- Modified parameter `from`:
  - enum added: `SPOT`, `MARGIN`
  - affected methods:
    - `portfolioMarginProBankruptcyLoanRepay()` (`POST /sapi/v1/portfolio/repay`)
    - `repayFuturesNegativeBalance()` (`POST /sapi/v1/portfolio/repay-futures-negative-balance`)
- Modified parameter `transferSide`:
  - enum added: `TO_UM`, `FROM_UM`
  - affected methods:
    - `bnbTransfer()` (`POST /sapi/v1/portfolio/bnb-transfer`)
- Modified parameter `transferType`:
  - enum added: `EARN_TO_FUTURE`, `FUTURE_TO_EARN`
  - affected methods:
    - `getTransferableEarnAssetBalanceForPortfolioMargin()` (`GET /sapi/v1/portfolio/earn-asset-balance`)
    - `transferLdusdtRwusdForPortfolioMargin()` (`POST /sapi/v1/portfolio/earn-asset-transfer`)

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

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

## 1.0.0 - 2025-12-17

- Initial release 
