### Changelog

## 1.10.0 - 2026-09-02

### Changed (6)

#### REST API

- Modified parameter `deltaLimit`:
  - type `float32` → `float64`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `price`:
  - type `float32` → `float64`
  - affected methods:
    - `newOrder()` (`POST /eapi/v1/order`)
- Modified parameter `qtyLimit`:
  - type `float32` → `float64`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `quantity`:
  - type `float32` → `float64`
  - affected methods:
    - `newOrder()` (`POST /eapi/v1/order`)
- Modified parameter `orders`:
  - items.`price`: type `float32` → `float64`
  - items.`quantity`: type `float32` → `float64`
  - affected methods:
    - `placeMultipleOrders()` (`POST /eapi/v1/batchOrders`)
- Modified response for `accountBlockTradeList()` (`GET /eapi/v1/block/user-trades`):
  - items.`legs`.items.`commission`: type `float32` → `float64`
  - items.`legs`.items.`executedAmount`: type `float32` → `float64`
  - items.`legs`.items.`executedQty`: type `float32` → `float64`
  - items.`legs`.items.`fee`: type `float32` → `float64`
  - items.`legs`.items.`orderPrice`: type `float32` → `float64`
  - items.`legs`.items.`orderQuantity`: type `float32` → `float64`
  - items.`legs`.items.`tradePrice`: type `float32` → `float64`
  - items.`legs`.items.`tradeQty`: type `float32` → `float64`

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

### Added (3)

#### REST API

- `tradfiOptionsContract()` (`POST /eapi/v1/stock/contract`)

#### WebSocket Streams

- `hour24Ticker()` (`<symbol>@optionTicker<expirationDate>` stream)
- `openInterest()` (`<underlying>@openInterest@<expirationDate>` stream)

### Changed (22)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.4.0`.

#### REST API

- Modified parameter `currency`:
  - enum added: `USDT`
  - affected methods:
    - `accountFundingFlow()` (`GET /eapi/v1/bill`)
- Modified parameter `deltaLimit`:
  - required: `false` → `true`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `frozenTimeInMilliseconds`:
  - required: `false` → `true`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `interval`:
  - enum added: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `8h`, `12h`, `1d`, `3d`, `1w`, `1M`
  - affected methods:
    - `klineCandlestickData()` (`GET /eapi/v1/klines`)
- Modified parameter `legs`:
  - items: required added: `side`, `type`, `quantity`, `symbol`
  - items: property `type` added
  - items: property `price` added
  - items: property `quantity` added
  - items: property `side` added
  - items: property `symbol` added
  - items: item property `type` added
  - items: item property `price` added
  - items: item property `quantity` added
  - items: item property `side` added
  - items: item property `symbol` added
  - affected methods:
    - `newBlockTradeOrder()` (`POST /eapi/v1/block/order/create`)
- Modified parameter `liquidity`:
  - enum added: `MAKER`, `TAKER`
  - affected methods:
    - `newBlockTradeOrder()` (`POST /eapi/v1/block/order/create`)
- Modified parameter `orders`:
  - items: required added: `quantity`, `symbol`, `side`, `type`
  - items.`isMmp`: type `string` → `boolean`
  - items.`postOnly`: type `string` → `boolean`
  - items.`price`: type `string` → `number`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: type `string` → `boolean`
  - items.`isMmp`: type `string` → `boolean`
  - items.`postOnly`: type `string` → `boolean`
  - items.`price`: type `string` → `number`
  - items.`quantity`: type `string` → `number`
  - items.`reduceOnly`: type `string` → `boolean`
  - affected methods:
    - `placeMultipleOrders()` (`POST /eapi/v1/batchOrders`)
- Modified parameter `qtyLimit`:
  - required: `false` → `true`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `selfTradePreventionMode`:
  - enum added: `NONE`
  - affected methods:
    - `newOrder()` (`POST /eapi/v1/order`)
- Modified parameter `symbol`:
  - required: `false` → `true`
  - affected methods:
    - `accountTradeList()` (`GET /eapi/v1/userTrades`)
- Modified parameter `underlying`:
  - required: `false` → `true`
  - affected methods:
    - `getMarketMakerProtectionConfig()` (`GET /eapi/v1/mmp`)
    - `resetMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpReset`)
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified parameter `windowTimeInMilliseconds`:
  - required: `false` → `true`
  - affected methods:
    - `setMarketMakerProtectionConfig()` (`POST /eapi/v1/mmpSet`)
- Modified response for `cancelMultipleOptionOrders()` (`DELETE /eapi/v1/batchOrders`):
  - items: property `fee` added
  - items: item property `fee` added

- Modified response for `placeMultipleOrders()` (`POST /eapi/v1/batchOrders`):
  - items: property `postOnly` added
  - items: property `fee` added
  - items: item property `postOnly` added
  - items: item property `fee` added

- Modified response for `orderBook()` (`GET /eapi/v1/depth`):
  - `asks`.items: minItems `0` → `2`
  - `asks`.items: maxItems `null` → `2`
  - `bids`.items: minItems `0` → `2`
  - `bids`.items: maxItems `null` → `2`

- Modified response for `exchangeInformation()` (`GET /eapi/v1/exchangeInfo`):
  - `optionSymbols`.items: property `underlyingType` added
  - `optionSymbols`.items: property `contractType` added
  - `optionSymbols`.items: property `nakedSell` added
  - `optionSymbols`.items: item property `underlyingType` added
  - `optionSymbols`.items: item property `contractType` added
  - `optionSymbols`.items: item property `nakedSell` added

- Modified response for `querySingleOrder()` (`GET /eapi/v1/order`):
  - property `postOnly` added

- Modified response for `newOrder()` (`POST /eapi/v1/order`):
  - property `fee` added
  - property `postOnly` added

#### WebSocket Streams

- Modified parameter `interval`:
  - enum added: `1m`, `3m`, `5m`, `15m`, `30m`, `1h`, `2h`, `4h`, `6h`, `12h`, `1d`, `3d`, `1w`
  - affected methods:
    - `klineCandlestickStreams()` (`<symbol>@kline_<interval>` stream)
- Modified parameter `level`:
  - enum added: `5`, `10`, `20`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<level>@<updateSpeed>` stream)
- Modified parameter `updateSpeed`:
  - required: `false` → `true`
  - enum added: `100ms`, `500ms`
  - affected methods:
    - `partialBookDepthStreams()` (`<symbol>@depth<level>@<updateSpeed>` stream)
    - `diffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream)

### Removed (2)

#### WebSocket Streams

- `/<symbol>@optionTicker()` (`<symbol>@optionTicker` stream)
- `/underlying@optionOpenInterest@<expirationDate>()` (`underlying@optionOpenInterest@<expirationDate>` stream)

## 1.5.0 - 2026-03-26

### Changed (10)

#### REST API

- Added parameter `selfTradePreventionMode`
  - affected methods:
    - `newOrder()` (`POST /eapi/v1/order`)
- Modified parameter `orders`:
  - items: property `selfTradePreventionMode` added
  - items: item property `selfTradePreventionMode` added
  - affected methods:
    - `placeMultipleOrders()` (`POST /eapi/v1/batchOrders`)
- Modified response for `cancelMultipleOptionOrders()` (`DELETE /eapi/v1/batchOrders`):
  - items: property `selfTradePreventionMode` added
  - items: item property `selfTradePreventionMode` added

- Modified response for `placeMultipleOrders()` (`POST /eapi/v1/batchOrders`):
  - items: property `selfTradePreventionMode` added
  - items: item property `selfTradePreventionMode` added

- Modified response for `optionMarginAccountInformation()` (`GET /eapi/v1/marginAccount`):
  - property `tradeGroupId` added

- Modified response for `queryCurrentOpenOptionOrders()` (`GET /eapi/v1/openOrders`):
  - items: property `selfTradePreventionMode` added
  - items: item property `selfTradePreventionMode` added

- Modified response for `cancelOptionOrder()` (`DELETE /eapi/v1/order`):
  - property `selfTradePreventionMode` added

- Modified response for `querySingleOrder()` (`GET /eapi/v1/order`):
  - property `selfTradePreventionMode` added

- Modified response for `newOrder()` (`POST /eapi/v1/order`):
  - property `selfTradePreventionMode` added
- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.2.0`.

## 1.4.0 - 2026-03-16

### Changed (2)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.1.0`.

#### REST API

- Modified response for `cancelAllOptionOrdersOnSpecificSymbol()` (`DELETE /eapi/v1/allOpenOrders`):
  - `code`: type `integer` → `string`

## 1.3.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common/v2` library to version `2.0.0`.

## 1.2.0 - 2026-01-29

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.2.0`.

## 1.1.0 - 2026-01-13

### Added (3)

#### WebSocket Streams

`Market WebSocket Streams`:
- `NewSymbolInfo()` (`!optionSymbol` stream)

`Public WebSocket Streams`:
- `DiffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream)
- `IndividualSymbolBookTickerStreams()` (`<symbol>@bookTicker` stream)

### Changed (4)

- Updated `github.com/binance/binance-connector-go/common` library to version `1.1.0`.
- Updated directory name.

#### WebSocket Streams

- Modified parameter `id`:
  - type `string` → `int32`
  - affected methods:
    - `PartialBookDepthStreams()` (`<symbol>@depth<levels>@<updateSpeed>` stream)
    - `IndexPriceStreams()` (`!index@arr` stream)
    - `KlineCandlestickStreams()` (`<symbol>@kline_<interval>` stream)
    - `Ticker24Hour()` (`<symbol>@ticker` stream)
    - `TradeStreams()` (`<symbol>@trade` stream)
    - `MarkPrice()` (`<underlyingAsset>@markPrice` stream)
    - `OpenInterest()` (`<underlyingAsset>@openInterest@<expirationDate>` stream)
    - `NewSymbolInfo()` (`!optionSymbol` stream)
    - `DiffBookDepthStreams()` (`<symbol>@depth@<updateSpeed>` stream)
    - `IndividualSymbolBookTickerStreams()` (`<symbol>@bookTicker` stream)

- Handled separated base urls:
  - `Public WebSocket Streams`:
    - `Base URL`: `wss://fstream.binance.com/public/stream`
  - `Market WebSocket Streams`:
    - `Base URL`: `wss://fstream.binance.com/market/stream`
  - `User Data WebSocket Streams`:
    - `Base URL`: `wss://fstream.binance.com/private/ws`

### Removed (1)

#### WebSocket Streams

- `Ticker24HourByUnderlyingAssetAndExpirationData()` (`<underlyingAsset>@ticker@<expirationDate>` stream)

## 1.0.0 - 2025-12-17

- Initial release 
