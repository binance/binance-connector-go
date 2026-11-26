### Changelog

## 1.7.0 - 2026-11-26

### Added (2)

#### REST API

- `bstocksCorporateActionRecords()` (`GET /sapi/v1/bstocks/corp-action`)
- `corporateActionRecords()` (`GET /sapi/v1/equity/corp-action`)

### Changed (2)

#### REST API

- Added response schema `bstocksCorporateActionRecordsResponse`
- Added response schema `corporateActionRecordsResponse`

## 1.6.0 - 2026-11-24

### Changed (7)

#### REST API

- Added parameter `tokenizedAssetSettleWalletType`
  - affected methods:
    - `placeEquityOrder()` (`POST /sapi/v1/equity/order/place`)
    - `tokenizedMint()` (`POST /sapi/v1/equity/tokenized/mint`)
    - `tokenizedRedeem()` (`POST /sapi/v1/equity/tokenized/redeem`)
- Modified response for `tokenizedConvertStatus()` (`GET /sapi/v1/equity/tokenized/convert-status`):
  - property `tokenizedAssetSettleWalletType` added

- Modified response for `tokenizedConvertHistory()` (`GET /sapi/v1/equity/tokenized/history`):
  - `rows`.items: property `tokenizedAssetSettleWalletType` added
  - `rows`.items: item property `tokenizedAssetSettleWalletType` added

- Modified response for `tokenizedMint()` (`POST /sapi/v1/equity/tokenized/mint`):
  - property `tokenizedAssetSettleWalletType` added

- Modified response for `tokenizedRedeem()` (`POST /sapi/v1/equity/tokenized/redeem`):
  - property `tokenizedAssetSettleWalletType` added

- Added response field `tokenizedAssetSettleWalletType`
  - affected events:
    - `tokenizedConvertStatusResponse`
    - `tokenizedMintResponse`
    - `tokenizedRedeemResponse`
- Modified response field `rows`:
  - items: property `tokenizedAssetSettleWalletType` added
  - items: item property `tokenizedAssetSettleWalletType` added
  - affected events:
    - `tokenizedConvertHistoryResponse`

## 1.5.1 - 2026-11-12

### Changed (2)

#### WebSocket Streams

- Modified response for `priceStream()` (`price` stream):
  - `rates`.items: property `prc` added
  - `rates`.items: property `lrc` added
  - `rates`.items: item property `prc` added
  - `rates`.items: item property `lrc` added

- Modified response field `rates`:
  - items: property `prc` added
  - items: property `lrc` added
  - items: item property `prc` added
  - items: item property `lrc` added
  - affected events:
    - `priceStreamResponse`

## 1.5.0 - 2026-10-23

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.11.0`.
- Fixed WebSocket Streams names being converted to lowercase: the case of stream names is now kept.

## 1.4.0 - 2026-10-15

### Changed (3)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.10.0`.
- Updated dependencies.
- Raised the minimum Go version to `1.26`.

## 1.3.0 - 2026-09-28

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.9.0`.

## 1.2.1 - 2026-09-09

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.1`.

## 1.2.0 - 2026-09-02

### Added (1)

#### WebSocket Streams

- Added parameter `id`
  - affected methods:
    - `calendarStream()` (`calendar` stream)
    - `klineStream()` (`<symbol>@kline_<interval>` stream)
    - `orderReportStream()` (`<listenKey>@orderReport` stream)
    - `priceStream()` (`price` stream)
    - `quoteStream()` (`<symbol>@quote` stream)
    - `tradabilityStream()` (`<symbol>@tradability` stream)
    - `tradingStatusStream()` (`<symbol>@tradingStatus` stream)

### Changed (1)

#### WebSocket Streams

- Modified response for `orderReportStream()` (`<listenKey>@orderReport` stream):
  - `FN`: type `float32` → `float64`
  - `N`: type `float32` → `float64`
  - `Q`: type `float32` → `float64`
  - `Z`: type `float32` → `float64`
  - `fq`: type `float32` → `float64`
  - `p`: type `float32` → `float64`
  - `tc`: type `float32` → `float64`

## 1.1.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.0.0 - 2026-08-19

- Initial release 
