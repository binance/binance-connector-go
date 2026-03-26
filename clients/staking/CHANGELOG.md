### Changelog

## 1.5.0 - 2026-03-26

### Changed (7)

- Added parameter `purchaseId`
  - affected methods:
    - `getEthStakingHistory()` (`GET /sapi/v1/eth-staking/eth/history/stakingHistory`)
    - `getSolStakingHistory()` (`GET /sapi/v1/sol-staking/sol/history/stakingHistory`)
- Added parameter `redeemId`
  - affected methods:
    - `getEthRedemptionHistory()` (`GET /sapi/v1/eth-staking/eth/history/redemptionHistory`)
    - `getSolRedemptionHistory()` (`GET /sapi/v1/sol-staking/sol/history/redemptionHistory`)
- Modified response for `redeemEth()` (`POST /sapi/v1/eth-staking/eth/redeem`):
  - property `redeemId` added

- Modified response for `redeemSol()` (`POST /sapi/v1/sol-staking/sol/redeem`):
  - property `redeemId` added

- Modified response for `subscribeSolStaking()` (`POST /sapi/v1/sol-staking/sol/stake`):
  - property `purchaseId` added

- Modified response for `subscribeEthStaking()` (`POST /sapi/v2/eth-staking/eth/stake`):
  - property `purchaseId` added
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

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
