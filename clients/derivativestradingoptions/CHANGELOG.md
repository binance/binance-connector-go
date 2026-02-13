### Changelog

## 1.3.0 - 2026-02-13

### Changed (1)

- Updated `github.com/binance/binance-connector-go/common` library to version `2.0.0`.

## 1.3.0 - 2026-01-29

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