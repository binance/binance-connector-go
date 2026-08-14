### Changelog

## 1.12.0 - 2026-08-14

### Added (1)

- `getSpotAssetTags()` (`GET /sapi/v1/spot/asset/tags`)

## 1.11.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.10.0 - 2026-07-14

### Changed (14)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Added parameter `recvWindow`
  - affected methods:
    - `getCountryList()` (`GET /sapi/v1/localentity/country/list`)
    - `getRegionList()` (`GET /sapi/v1/localentity/region/list`)
- Deleted parameter `signature`
  - affected methods:
    - `submitDepositQuestionnaire()` (`PUT /sapi/v1/localentity/broker/deposit/provide-info`)
    - `brokerWithdraw()` (`POST /sapi/v1/localentity/broker/withdraw/apply`)
- Modified parameter `accountType`:
  - enum added: `SPOT`, `MARGIN`
  - affected methods:
    - `dustlog()` (`GET /sapi/v1/asset/dribblet`)
    - `dustTransfer()` (`POST /sapi/v1/asset/dust`)
    - `getAssetsThatCanBeConvertedIntoBnb()` (`POST /sapi/v1/asset/dust-btc`)
- Modified parameter `depositId`:
  - type `string` → `integer`
  - affected methods:
    - `depositHistoryV2()` (`GET /sapi/v2/localentity/deposit/history`)
- Modified parameter `fromSymbol`:
  - enum added: `ISOLATEDMARGIN_MARGIN`, `ISOLATEDMARGIN_ISOLATEDMARGIN`
  - affected methods:
    - `queryUserUniversalTransferHistory()` (`GET /sapi/v1/asset/transfer`)
    - `userUniversalTransfer()` (`POST /sapi/v1/asset/transfer`)
- Modified parameter `needBtcValuation`:
  - type `string` → `boolean`
  - affected methods:
    - `fundingWallet()` (`POST /sapi/v1/asset/get-funding-asset`)
- Modified parameter `status`:
  - enum added: `0`, `1`, `2`, `6`, `7`, `8`
  - affected methods:
    - `depositHistory()` (`GET /sapi/v1/capital/deposit/hisrec`)
- Modified parameter `subAccountId`:
  - type `integer` → `string`
  - affected methods:
    - `oneClickArrivalDepositApply()` (`POST /sapi/v1/capital/deposit/credit-apply`)
- Modified parameter `toSymbol`:
  - enum added: `MARGIN_ISOLATEDMARGIN`, `ISOLATEDMARGIN_ISOLATEDMARGIN`
  - affected methods:
    - `queryUserUniversalTransferHistory()` (`GET /sapi/v1/asset/transfer`)
    - `userUniversalTransfer()` (`POST /sapi/v1/asset/transfer`)
- Modified parameter `type`:
  - enum added: `SPOT`, `MARGIN`, `FUTURES`
  - affected methods:
    - `dailyAccountSnapshot()` (`GET /sapi/v1/accountSnapshot`)
- Modified parameter `type`:
  - enum added: `DELEGATE`, `UNDELEGATE`
  - affected methods:
    - `queryUserDelegationHistory()` (`GET /sapi/v1/asset/custody/transfer-history`)
- Modified parameter `type`:
  - enum added: `MAIN_UMFUTURE`, `MAIN_CMFUTURE`, `MAIN_MARGIN`, `UMFUTURE_MAIN`, `UMFUTURE_MARGIN`, `CMFUTURE_MAIN`, `CMFUTURE_MARGIN`, `MARGIN_MAIN`, `MARGIN_UMFUTURE`, `MARGIN_CMFUTURE`, `ISOLATEDMARGIN_MARGIN`, `MARGIN_ISOLATEDMARGIN`, `ISOLATEDMARGIN_ISOLATEDMARGIN`, `MAIN_FUNDING`, `FUNDING_MAIN`, `FUNDING_UMFUTURE`, `UMFUTURE_FUNDING`, `MARGIN_FUNDING`, `FUNDING_MARGIN`, `FUNDING_CMFUTURE`, `CMFUTURE_FUNDING`, `MAIN_OPTION`, `OPTION_MAIN`, `UMFUTURE_OPTION`, `OPTION_UMFUTURE`, `MARGIN_OPTION`, `OPTION_MARGIN`, `FUNDING_OPTION`, `OPTION_FUNDING`, `MAIN_PORTFOLIO_MARGIN`, `PORTFOLIO_MARGIN_MAIN`
  - affected methods:
    - `userUniversalTransfer()` (`POST /sapi/v1/asset/transfer`)
- Modified response for `depositHistoryTravelRule()` (`GET /sapi/v1/localentity/deposit/history`):
  - items: property `completeTime` added
  - items: property `travelRuleStatusV2` added
  - items: property `unlockConfirm` deleted
  - items: property `walletType` deleted
  - items: item property `completeTime` added
  - items: item property `travelRuleStatusV2` added
  - items: item property `unlockConfirm` deleted
  - items: item property `walletType` deleted

## 1.9.0 - 2026-06-30

### Changed (2)

- Modified response for `brokerWithdraw()` (`POST /sapi/v1/localentity/broker/withdraw/apply`):
  - property `accepted` added
  - property `accpted` deleted

- Modified response for `withdrawTravelRule()` (`POST /sapi/v1/localentity/withdraw/apply`):
  - property `accepted` added
  - property `accpted` deleted

## 1.8.0 - 2026-06-16

### Added (2)

- `getCountryList()` (`GET /sapi/v1/localentity/country/list`)
- `getRegionList()` (`GET /sapi/v1/localentity/region/list`)

### Changed (1)

- Added parameter `accountType`
  - affected methods:
    - `dustConvert()` (`POST /sapi/v1/asset/dust-convert/convert`)
    - `dustConvertibleAssets()` (`POST /sapi/v1/asset/dust-convert/query-convertible-assets`)

## 1.7.0 - 2026-03-26

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

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
