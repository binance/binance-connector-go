### Changelog

## 1.10.0 - 2026-09-02

### Changed (1)

#### REST API

- Modified parameter `depositAmount`:
  - type `float32` → `float64`
  - affected methods:
    - `subscribeDualInvestmentProducts()` (`POST /sapi/v1/dci/product/subscribe`)

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
- Added parameter `autoCompoundPlan`
  - affected methods:
    - `changeAutoCompoundStatus()` (`POST /sapi/v1/dci/product/auto_compound/edit-status`)
- Deleted parameter `AutoCompoundPlan`
  - affected methods:
    - `changeAutoCompoundStatus()` (`POST /sapi/v1/dci/product/auto_compound/edit-status`)
- Modified parameter `autoCompoundPlan`:
  - enum added: `NONE`, `STANDARD`, `ADVANCED`
  - affected methods:
    - `subscribeDualInvestmentProducts()` (`POST /sapi/v1/dci/product/subscribe`)
- Modified parameter `optionType`:
  - enum added: `CALL`, `PUT`
  - affected methods:
    - `getDualInvestmentProductList()` (`GET /sapi/v1/dci/product/list`)
- Modified parameter `status`:
  - enum added: `PENDING`, `PURCHASE_SUCCESS`, `SETTLED`, `PURCHASE_FAIL`, `REFUNDING`, `REFUND_SUCCESS`, `SETTLING`
  - affected methods:
    - `getDualInvestmentPositions()` (`GET /sapi/v1/dci/product/positions`)
- Modified response for `getDualInvestmentPositions()` (`GET /sapi/v1/dci/product/positions`):
  - `list`.items: property `subscriptionTime` added
  - `list`.items: item property `subscriptionTime` added

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
