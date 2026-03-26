### Changelog

## 2.2.0 - 2026-03-26

### Changed (4)

- Updated `apiKey`, `apiSecret`, `privateKey`, and `privateKeyPassphrase` fields in `ConfigurationRestAPI` and `ConfigurationWebsocketApi` to be unexported and added corresponding getter methods for better encapsulation.
- Updated `panic` calls in `utils.go` to return errors instead, allowing the caller to handle them gracefully instead of crashing the application.
- Updated `Decoder` method to generate temp files with a unique prefix to avoid conflicts when multiple decoders are used concurrently.
- Fixed error returned when signing requests with a private key by ensuring the error is returned to the caller instead of being ignored.

## 2.1.0 - 2026-03-16

### Added (1)

- Added `serverShutdown` event handler.

## 2.0.1 - 2026-02-13

### Changed (1)

- Fixed imported `common` package for websocket mocks.

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