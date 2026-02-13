### Changelog

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `2.0.0`.

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