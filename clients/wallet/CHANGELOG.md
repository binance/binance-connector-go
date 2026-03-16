### Changelog

## 1.6.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.5.0 - 2026-03-09

### Changed (1)

- Modified response for `vasp_list()` (`GET /sapi/v1/localentity/vasp`):
  - items: property `identifier` added
  - items: item property `identifier` added

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (1)

- Modified response for `assetDividendRecord()` (`GET /sapi/v1/asset/assetDividend`):
  - `rows`.items: property `direction` added
  - `rows`.items: item property `direction` added

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Added (1)

- `SubmitDepositQuestionnaireV2()` (`PUT /sapi/v2/localentity/deposit/provide-info`)

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

- Modified parameter `depositId`:
  - type `string` → `integer`
  - affected methods:
    - `SubmitDepositQuestionnaire()` (`PUT /sapi/v1/localentity/broker/deposit/provide-info`)

## 1.0.0 - 2025-12-17

- Initial release 
