### Changelog

## 1.8.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.7.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.6.0 - 2026-07-14

### Changed (4)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Modified parameter `direction`:
  - enum added: `ADDITIONAL`, `REDUCED`
  - affected methods:
    - `flexibleLoanAdjustLtv()` (`POST /sapi/v2/loan/flexible/adjust/ltv`)
- Modified parameter `repaymentType`:
  - enum added: `1`, `2`
  - affected methods:
    - `flexibleLoanRepay()` (`POST /sapi/v2/loan/flexible/repay`)
- Modified parameter `type`:
  - enum added: `borrowIn`, `collateralSpent`, `repayAmount`, `collateralReturn`, `addCollateral`, `removeCollateral`, `collateralReturnAfterLiquidation`
  - affected methods:
    - `getCryptoLoansIncomeHistory()` (`GET /sapi/v1/loan/income`)

### Removed (1)

- `checkCollateralRepayRateStableRate()` (`GET /sapi/v1/loan/repay/collateral/rate`)

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

### Added (1)

- `GetFlexibleLoanInterestRateHistory()` (`GET /sapi/v2/loan/interestRateHistory`)

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

## 1.0.0 - 2025-12-17

- Initial release 
