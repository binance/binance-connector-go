### Changelog

## 1.14.0 - 2026-09-02

### Changed (4)

#### REST API

- Modified parameter `amount`:
  - type `float32` → `float64`
  - affected methods:
    - `depositAssetsIntoTheManagedSubAccount()` (`POST /sapi/v1/managed-subaccount/deposit`)
    - `futuresTransferForSubAccount()` (`POST /sapi/v1/sub-account/futures/transfer`)
    - `getManagedSubAccountDepositAddress()` (`GET /sapi/v1/managed-subaccount/deposit/address`)
    - `getSubAccountDepositAddress()` (`GET /sapi/v1/capital/deposit/subAddress`)
    - `marginTransferForSubAccount()` (`POST /sapi/v1/sub-account/margin/transfer`)
    - `subAccountFuturesAssetTransfer()` (`POST /sapi/v1/sub-account/futures/internalTransfer`)
    - `transferToMaster()` (`POST /sapi/v1/sub-account/transfer/subToMaster`)
    - `transferToSubAccountOfSameMaster()` (`POST /sapi/v1/sub-account/transfer/subToSub`)
    - `universalTransfer()` (`POST /sapi/v1/sub-account/universalTransfer`)
    - `withdrawlAssetsFromTheManagedSubAccount()` (`POST /sapi/v1/managed-subaccount/withdraw`)
- Modified parameter `orderArgs`:
  - items.`quantity`: type `float32` → `float64`
  - affected methods:
    - `movePositionForSubAccount()` (`POST /sapi/v1/sub-account/futures/move-position`)
- Modified response for `queryManagedSubAccountFuturesAssetDetails()` (`GET /sapi/v1/managed-subaccount/fetch-future-asset`):
  - `snapshotVos`.items.`data`.`position`.items.`positionAmt`: type `float32` → `float64`
- Modified response for `querySubAccountAssets()` (`GET /sapi/v3/sub-account/assets`):
  - `balances`.items.`free`: type `float32` → `float64`

## 1.13.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.12.0 - 2026-08-24

### Changed (4)

- Added parameter `productType`
  - affected methods:
    - `getMovePositionHistoryForSubAccount()` (`GET /sapi/v1/sub-account/futures/move-position`)
- Modified parameter `productType`:
  - enum added: `OPTION`
  - affected methods:
    - `movePositionForSubAccount()` (`POST /sapi/v1/sub-account/futures/move-position`)
- Modified response for `movePositionForSubAccount()` (`POST /sapi/v1/sub-account/futures/move-position`):
  - `movePositionOrders`.items.`positionSide`: nullable `false` → `true`
  - `movePositionOrders`.items.`positionSide`: nullable `false` → `true`

- Modified response field `movePositionOrders`:
  - items.`positionSide`: nullable `false` → `true`
  - items.`positionSide`: nullable `false` → `true`
  - affected events:
    - `movePositionForSubAccountResponse`

## 1.11.0 - 2026-08-19

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.
- Modified response for `querySubAccountApiKey()` (`GET /sapi/v1/sub-account/subAccountApi`):
  - property `rows` added
  - property `list` deleted

## 1.10.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.9.0 - 2026-07-14

### Added (4)

- `createSubAccountApiKey()` (`POST /sapi/v1/sub-account/subAccountApi`)
- `deleteSubAccountApiKey()` (`DELETE /sapi/v1/sub-account/subAccountApi`)
- `modifySubAccountApiKeyPermission()` (`POST /sapi/v1/sub-account/subAccountApiPermission`)
- `querySubAccountApiKey()` (`GET /sapi/v1/sub-account/subAccountApi`)

### Changed (7)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Modified parameter `fromAccountType`:
  - enum added: `SPOT`, `USDT_FUTURE`, `COIN_FUTURE`, `MARGIN`, `ISOLATED_MARGIN`
  - affected methods:
    - `universalTransfer()` (`POST /sapi/v1/sub-account/universalTransfer`)
- Modified parameter `productType`:
  - enum added: `UM`
  - affected methods:
    - `movePositionForSubAccount()` (`POST /sapi/v1/sub-account/futures/move-position`)
- Modified parameter `status`:
  - type `string` → `integer`
  - affected methods:
    - `addIpRestrictionForSubAccountApiKey()` (`POST /sapi/v2/sub-account/subAccountApi/ipRestriction`)
- Modified parameter `toAccountType`:
  - enum added: `SPOT`, `USDT_FUTURE`, `COIN_FUTURE`, `MARGIN`, `ISOLATED_MARGIN`
  - affected methods:
    - `universalTransfer()` (`POST /sapi/v1/sub-account/universalTransfer`)
- Modified parameter `transferFunctionAccountType`:
  - enum added: `SPOT`, `MARGIN`, `ISOLATED_MARGIN`, `USDT_FUTURE`, `COIN_FUTURE`
  - affected methods:
    - `queryManagedSubAccountTransferLogSubAccountTrading()` (`GET /sapi/v1/managed-subaccount/query-trans-log`)
    - `queryManagedSubAccountTransferLogMasterAccountInvestor()` (`GET /sapi/v1/managed-subaccount/queryTransLogForInvestor`)
    - `queryManagedSubAccountTransferLogMasterAccountTrading()` (`GET /sapi/v1/managed-subaccount/queryTransLogForTradeParent`)
- Modified parameter `type`:
  - enum added: `SPOT`, `MARGIN`, `FUTURES`
  - affected methods:
    - `queryManagedSubAccountSnapshot()` (`GET /sapi/v1/managed-subaccount/accountSnapshot`)

## 1.8.0 - 2026-06-16

### Changed (1)

- Added parameter `includeSource`
  - affected methods:
    - `getSubAccountDepositHistory()` (`GET /sapi/v1/capital/deposit/subHisrec`)

## 1.7.0 - 2026-05-20

### Changed (2)

- Added parameter `rows`
  - affected methods:
    - `getMovePositionHistoryForSubAccount()` (`GET /sapi/v1/sub-account/futures/move-position`)
- Deleted parameter `row`
  - affected methods:
    - `getMovePositionHistoryForSubAccount()` (`GET /sapi/v1/sub-account/futures/move-position`)

## 1.6.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.5.0 - 2026-03-16

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

## 1.4.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.3.0 - 2026-02-02

### Changed (2)

- Added parameter `limit`
  - affected methods:
    - `getSummaryOfSubAccountsFuturesAccount()` (`GET /sapi/v1/sub-account/futures/accountSummary`)
- Added parameter `page`
  - affected methods:
    - `getSummaryOfSubAccountsFuturesAccount()` (`GET /sapi/v1/sub-account/futures/accountSummary`)

## 1.2.0 - 2026-01-23

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

## 1.0.0 - 2025-12-17

- Initial release 
