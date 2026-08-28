### Changelog

## 2.8.0 - 2026-08-28

### Added (2)

- Added bounded reconnection through `MaxReconnectAttempts` and the `WithWsMaxReconnectAttempts` and `WithWsStreamsMaxReconnectAttempts` options: a reconnect is retried that many times with a backoff before giving up and reporting `ErrReconnectAttemptsExhausted` on the error channel and to `OnError`, from 1 to 10 attempts and defaulting to 3.
- Added `Unwrap` to `WebSocketError`, for `errors.Is` and `errors.As`.

### Changed (2)

- Fixed the websocket not being renewed after a network interruption, including the keep-alive refresh leaving the connection closed and stopping for good after a failed refresh.
- Fixed the connection handling under concurrency: in-flight requests waiting for ever when the connection is closed, the panic on a closed channel when it is closed twice, the read loop being started more than once, the `concurrent map writes` crash on the stream connection map, the data race on the pool round-robin index, and the pool connect waiting on connections it never dialled.

## 2.7.0 - 2026-08-18

### Added (1)

- Added `Stocks` base url

## 2.6.0 - 2026-07-15

### Changed (1)

- Updated `Alpha Websocket Streams` base url

### Removed (1)

- Removed `Nft` base url

## 2.5.0 - 2026-07-14

### Changed (1)

- Updated dependencies

## 2.4.0 - 2026-07-14

### Added (1)

- Added `Alpha Websocket Streams` base url

## 2.3.0 - 2026-06-26

### Added (1)

- Added `W3W Prediction` base url

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