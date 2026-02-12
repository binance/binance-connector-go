### Changelog

## 2.0.0 - 2026-02-12

### Changed (2)

- Updated `WebSocketCommon` connect method to accept streams parameter for subscribing upon connection.
- Updated `retry` logic in `utils.go` to use `SleepContext` for better handling of context cancellation during retries.

## 1.2.0 - 2026-01-23

### Added (1)

- Added `Alpha` base url

### Changed (2)

- Updated `PrepareRequest` and `SendRequest` methods to include `signed` boolean parameter to indicate if the request requires signing.
- Fixed Lock issue in `WebsocketStreams` `Unsubscribe` method by ensuring proper unlocking of mutex after stream removal.

## 1.1.0 - 2026-01-13

- Updated `WebsocketStreams` `Subscribe` method to support `int32` format for `id` parameter
- Support Derivatives Trading Options different WS Streams URL paths.

## 1.0.1 - 2025-12-17

- Fixed `common` tags issue

## 1.0.0 - 2025-12-16

- Initial release 