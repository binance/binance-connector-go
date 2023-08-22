# Change log

## v0.5.2 - 2023-08-22

### Fixed
- Fixed formatting issues causing JSON Unmarshal error in some `wallet` methods.

## v0.5.1 - 2023-08-21

### Fixed
- Implemented a fix for the automatic conversion to scientific notation when handling large `quantity` values.

## v0.5.0 - 2023-06-30

### Added
- Added Websocket API Endpoints
- Added User-Agent header to Websocket and Websocket API

### Fixed
- bookTickerService Endpoint fixed to correctly support individual `symbol` parameter.

## v0.4.0 - 2023-05-08

### Added
- Staking Endpoints Implemented

## v0.3.0 - 2023-05-02

### Added
- New Endpoint, `NewGetManagedSubAccountDepositAddressService()`: `GET /sapi/v1/managed-subaccount/deposit/address` - Get Managed Sub-account Deposit Address (For Investor Master Account) (USER_DATA)

### Fixed
- Added separate USDT-M and COIN-M Response Types for relevant Subaccount Endpoints
- Added support for Conditional Fields in relevant responses on Account Endpoints
- Responses for `POST /api/v3/order/cancelReplace` support all 4 Response Types
- Support for specifying `recvWindow` via `WithRecvWindow()` function
- `strategyId` and `strategyType` parameters added to `TestNewOrder` and `CreateOrderService` endpoints
- `UiKlines`: corrected `limit` parameter to not be sent as `interval`

## v0.2.0 - 2023-04-21

### Added
- WebsocketStreamClient
  - Websocket Client can be initialized with 2 parameters, `NewWebsocketStreamClient(isCombined, baseURL)`:
  - `isCombined` is a MANDATORY boolean value that specifies whether you are calling a combined stream or not.
    - If `isCombined` is set to `true`, `"/stream?streams="` will be appended to the `baseURL` to allow for Combining streams.
    - Otherwise, if set to `false`, `"/ws/"` will be appended to the `baseURL`.
  - `baseURL` is an OPTIONAL string value that determines the base URL to use for the websocket connection.
    - If `baseURL` is not set, it will default to the Live Exchange URL: `"wss://stream.binance.com:9443"`.

### Fixed
- Order Response Types for `CreateOrderService` - `POST /api/v3/order`
  - Added support for all 3 Order Response Types - `ACK`, `RESULT` and `FULL`
- Order Response Types for `MarginAccountNewOrderService` - `POST /sapi/v1/margin/order`
  - Added support for all 3 Order Response Types - `ACK`, `RESULT` and `FULL`
- Different Response Types for `MarginIsolatedAccountInfoService` - `GET /sapi/v1/margin/isolated/account`
  - Added support for both Response Types, depending on whether `symbols` is set or not.

### Changed
- Renamed `WsAllMiniMarketsStatServe` to `WsAllMarketMiniTickersStatServe`.
- Renamed `WsMarketsStatServe` to `WsMarketTickersStatServe`.
- Renamed `WsAllMarketsStatServe` to `WsAllMarketTickersStatServe`.
- Updated Github Action `UnitTest`.

### Removed
- Removed unused `setFormParam`, `setFormParams`, `WithRecvWindow`, `WithHeader` and `WithHeaders` functions from `request.go`.

## v0.1.0 - 2023-03-31

- First release
