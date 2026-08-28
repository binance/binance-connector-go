### Changelog

## 1.12.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.11.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.10.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.9.0 - 2026-07-14

### Changed (12)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Added parameter `lang`
  - affected methods:
    - `getYieldArenaActivities()` (`GET /sapi/v1/earn/arena/activities`)
- Modified parameter `aprPeriod`:
  - enum added: `DAY`, `YEAR`
  - affected methods:
    - `getRateHistory()` (`GET /sapi/v1/simple-earn/flexible/history/rateHistory`)
- Modified parameter `asset`:
  - enum added: `USDC`, `USDT`
  - affected methods:
    - `getBfusdSubscriptionHistory()` (`GET /sapi/v1/bfusd/history/subscriptionHistory`)
    - `getRwusdSubscriptionHistory()` (`GET /sapi/v1/rwusd/history/subscriptionHistory`)
- Modified parameter `asset`:
  - enum added: `USDT`, `USDC`
  - affected methods:
    - `subscribeRwusd()` (`POST /sapi/v1/rwusd/subscribe`)
- Modified parameter `destAccount`:
  - enum added: `SPOT`, `FUND`
  - affected methods:
    - `redeemFlexibleProduct()` (`POST /sapi/v1/simple-earn/flexible/redeem`)
- Modified parameter `positionId`:
  - type `integer` → `string`
  - affected methods:
    - `getLockedRedemptionRecord()` (`GET /sapi/v1/simple-earn/locked/history/redemptionRecord`)
    - `getLockedRewardsHistory()` (`GET /sapi/v1/simple-earn/locked/history/rewardsRecord`)
    - `getLockedProductPosition()` (`GET /sapi/v1/simple-earn/locked/position`)
- Modified parameter `redeemTo`:
  - enum added: `SPOT`, `FLEXIBLE`
  - affected methods:
    - `setLockedProductRedeemOption()` (`POST /sapi/v1/simple-earn/locked/setRedeemOption`)
- Modified parameter `redeemTo`:
  - enum added: `SPOT`, `FLEXIBLE`
  - affected methods:
    - `subscribeLockedProduct()` (`POST /sapi/v1/simple-earn/locked/subscribe`)
- Modified parameter `sourceAccount`:
  - enum added: `SPOT`, `FUND`, `ALL`
  - affected methods:
    - `subscribeFlexibleProduct()` (`POST /sapi/v1/simple-earn/flexible/subscribe`)
    - `subscribeLockedProduct()` (`POST /sapi/v1/simple-earn/locked/subscribe`)
- Modified parameter `type`:
  - enum added: `FAST`, `STANDARD`
  - affected methods:
    - `redeemBfusd()` (`POST /sapi/v1/bfusd/redeem`)
    - `redeemRwusd()` (`POST /sapi/v1/rwusd/redeem`)
- Modified parameter `type`:
  - required: `true` → `false`
  - enum added: `BONUS`, `REALTIME`, `REWARDS`, `ALL`
  - affected methods:
    - `getFlexibleRewardsHistory()` (`GET /sapi/v1/simple-earn/flexible/history/rewardsRecord`)

## 1.8.0 - 2026-06-16

### Added (1)

- `getYieldArenaActivities()` (`GET /sapi/v1/earn/arena/activities`)

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
