### Changelog

## 1.7.0 - 2026-05-20

### Changed (3)

- Modified response for `getLockedRedemptionRecord()` (`GET /sapi/v1/simple-earn/locked/history/redemptionRecord`):
  - `rows`.items.`deliverDate`: type `string` → `integer`
  - `rows`.items.`deliverDate`: type `string` → `integer`

- Modified response for `getLockedProductPosition()` (`GET /sapi/v1/simple-earn/locked/position`):
  - `rows`.items.`deliverDate`: type `string` → `integer`
  - `rows`.items.`nextPayDate`: type `string` → `integer`
  - `rows`.items.`partialAmtDeliverDate`: type `string` → `integer`
  - `rows`.items.`purchaseTime`: type `string` → `integer`
  - `rows`.items.`rewardsEndDate`: type `string` → `integer`
  - `rows`.items.`deliverDate`: type `string` → `integer`
  - `rows`.items.`nextPayDate`: type `string` → `integer`
  - `rows`.items.`partialAmtDeliverDate`: type `string` → `integer`
  - `rows`.items.`purchaseTime`: type `string` → `integer`
  - `rows`.items.`rewardsEndDate`: type `string` → `integer`

- Modified response for `getLockedSubscriptionPreview()` (`GET /sapi/v1/simple-earn/locked/subscriptionPreview`):
  - items.`deliverDate`: type `string` → `integer`
  - items.`nextPayDate`: type `string` → `integer`
  - items.`nextSubscriptionDate`: type `string` → `integer`
  - items.`rewardsEndDate`: type `string` → `integer`
  - items.`valueDate`: type `string` → `integer`
  - items.`deliverDate`: type `string` → `integer`
  - items.`nextPayDate`: type `string` → `integer`
  - items.`nextSubscriptionDate`: type `string` → `integer`
  - items.`rewardsEndDate`: type `string` → `integer`
  - items.`valueDate`: type `string` → `integer`

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.
- Modified response for `getBfusdQuotaDetails()` (`GET /sapi/v1/bfusd/quota`):
  - property `subscriptionQuota` added

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

### Changed (1)

- Modified response for `getBfusdQuotaDetails()` (`GET /sapi/v1/bfusd/quota`):
  - property `subscribeEnable` deleted
  - property `redeemEnable` deleted

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

## 1.0.0 - 2025-12-17

- Initial release 
