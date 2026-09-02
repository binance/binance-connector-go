### Changelog

## 1.11.0 - 2026-09-02

### Changed (1)

- Updated the connector version for this release; no API changes.

## 1.10.0 - 2026-08-28

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.8.0`.
- Updated dependencies.

## 1.9.0 - 2026-08-19

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.7.0`.

## 1.8.0 - 2026-07-15

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.6.0`.

## 1.7.0 - 2026-07-14

### Changed (4)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.
- Deleted parameter `recvWindow`
  - affected methods:
    - `getC2CTradeHistory()` (`GET /sapi/v1/c2c/orderMatch/listUserOrderHistory`)
- Modified parameter `tradeType`:
  - enum added: `BUY`, `SELL`
  - affected methods:
    - `getC2CTradeHistory()` (`GET /sapi/v1/c2c/orderMatch/listUserOrderHistory`)
- Modified response for `getC2CTradeHistory()` (`GET /sapi/v1/c2c/orderMatch/listUserOrderHistory`):
  - property `code` added
  - property `data` added
  - property `message` added
  - property `success` added
  - property `total` added
  - property `orderStatus` deleted
  - property `tradeType` deleted
  - property `additionalKycVerify` deleted
  - property `asset` deleted
  - property `commission` deleted
  - property `payMethodName` deleted
  - property `orderNumber` deleted
  - property `takerCommissionRate` deleted
  - property `takerCommission` deleted
  - property `counterPartNickName` deleted
  - property `fiat` deleted
  - property `createTime` deleted
  - property `amount` deleted
  - property `unitPrice` deleted
  - property `advNo` deleted
  - property `totalPrice` deleted
  - property `fiatSymbol` deleted
  - property `takerAmount` deleted

## 1.6.0 - 2026-04-20

### Changed (1)

- Modified response for `getC2CTradeHistory()` (`GET /sapi/v1/c2c/orderMatch/listUserOrderHistory`):
  - property `asset` added
  - property `createTime` added
  - property `amount` added
  - property `takerCommission` added
  - property `orderStatus` added
  - property `counterPartNickName` added
  - property `orderNumber` added
  - property `takerAmount` added
  - property `commission` added
  - property `tradeType` added
  - property `payMethodName` added
  - property `additionalKycVerify` added
  - property `fiatSymbol` added
  - property `advNo` added
  - property `totalPrice` added
  - property `fiat` added
  - property `unitPrice` added
  - property `takerCommissionRate` added
  - property `data` deleted
  - property `message` deleted
  - property `success` deleted
  - property `total` deleted
  - property `code` deleted

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

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.

## 1.0.0 - 2025-12-17

- Initial release 
