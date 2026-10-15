# Binance Go CMS SDK

[![Build Status](https://img.shields.io/github/actions/workflow/status/binance/binance-connector-go/ci.yaml)](https://github.com/binance/binance-connector-go/actions)
[![Open Issues](https://img.shields.io/github/issues/binance/binance-connector-go)](https://github.com/binance/binance-connector-go/issues)
[![Known Vulnerabilities](https://img.shields.io/badge/security-scanned-brightgreen)](https://github.com/binance/binance-connector-go/security)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a client library for the Binance CMS API, enabling developers to receive Binance announcements programmatically. The library provides real-time push notifications for new Binance announcements through the WebSocket Streams:

- [Websocket Stream](./src/websocketstreams/websocket_streams.go)

## Table of Contents

- [Supported Features](#supported-features)
- [Installation](#installation)
- [Documentation](#documentation)
- [Websocket Streams](#websocket-streams)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Supported Features

- WebSocket Streams:
  - Real-time push notifications for Binance announcements (`com_announcement_en`).
- Inclusion of test cases and examples for quick onboarding.

## Installation

To use this library, ensure you have Go installed (version **1.26** or higher is required). You can install the library using the following command:

```bash
go get github.com/binance/binance-connector-go/clients/cms
```

## Documentation

For detailed information, refer to the [Binance API Documentation](https://developers.binance.com/docs/cms).

### Websocket Streams

WebSocket Streams push a notification each time Binance publishes an announcement. Use the [websocket-streams](./src/websocketstreams/websocket_streams.go) module to subscribe to these streams.

The streams are signed: provide your Binance API credentials in the configuration. The connection URL is signed again on every connect and reconnect.

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/cms"
	"github.com/binance/binance-connector-go/clients/cms/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ComAnnouncementEn()
}

func ComAnnouncementEn() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.CmsWebsocketStreamsProdUrl),
		common.WithWsStreamsApiKey("Your API Key"),
		common.WithWsStreamsApiSecret("Your API Secret"),
	)

	wsClient := client.NewBinanceCmsClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ComAnnouncementEn().Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.ComAnnouncementEnResponse) {
		b, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received message: %s\n", string(b))
	})

	for {
		time.Sleep(1 * time.Second)
	}
}
```

More examples are available in the [`examples/websocketstreams/DefaultAPI`](./examples/websocketstreams/DefaultAPI/) folder.

#### Configuration Options

The WebSocket Streams API supports the following advanced configuration options:

- `ApiKey`: Binance API key, sent with every connection.
- `ApiSecret`: Binance API secret, used to sign the connection URL.
- `PrivateKey`: RSA or ED25519 private key, used to sign the connection URL instead of `ApiSecret`.
- `PrivateKeyPassphrase`: Passphrase for the private key, if encrypted.
- `ReconnectDelay`: Delay between reconnection attempts.
- `MaxReconnectAttempts`: Number of reconnection attempts before giving up (1-10, default: 3).
- `Compression`: Enable response compression.
- `Proxy`: Proxy configuration:
  - `Host`: Proxy server hostname.
  - `Port`: Proxy server port.
  - `Protocol`: Proxy protocol (http or https).
  - `Auth`: Proxy authentication credentials:
    - `Username`: Proxy username.
    - `Password`: Proxy password.
- `Mode`: Choose between `single` and `pool` connection modes.
  - `Single`: A single WebSocket connection.
  - `Pool`: A pool of WebSocket connections.
- `PoolSize`: Define the number of WebSocket connections in pool mode.
- `Agent`: Customize the WebSocket Agent for advanced configurations.

##### Key Pair Based Authentication

The WebSocket Streams support key pair-based authentication. You can use `RSA` or `ED25519` keys to sign the connection URL. See the [Key Pair Based Authentication example](./docs/websocketstreams/key-pair-authentication.md) for detailed usage.

##### Reconnect Delay

Specify the delay in milliseconds between WebSocket reconnection attempts for streams. See the [Reconnect Delay example](./docs/websocketstreams/reconnect-delay.md) for detailed usage.

##### Max Reconnect Attempts

Specify how many times a reconnection is attempted before giving up and reporting `ErrReconnectAttemptsExhausted`, from 1 to 10 and defaulting to 3. See the [Max Reconnect Attempts example](./docs/websocketstreams/max-reconnect-attempts.md) for detailed usage.

##### Compression

Enable or disable compression for WebSocket Streams messages. See the [Compression example](./docs/websocketstreams/compression.md) for detailed usage.

##### Proxy

The WebSocket Streams supports HTTP/HTTPS proxy configurations. See the [Proxy example](./docs/websocketstreams/proxy.md) for detailed usage.

##### Connection Mode

Choose between `single` and `pool` connection modes for WebSocket Streams. The `single` mode uses a single WebSocket connection, while the `pool` mode uses a pool of WebSocket connections. See the [Connection Mode example](./docs/websocketstreams/connection-mode.md) for detailed usage.

##### WebSocket Agent

Customize the agent for advanced configurations. See the [WebSocket Agent example](./docs/websocketstreams/agent.md) for detailed usage.

#### Unsubscribing from Streams

You can unsubscribe from specific WebSocket streams using the `Unsubscribe` method. This is useful for managing active subscriptions without closing the connection.

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/cms"
	"github.com/binance/binance-connector-go/clients/cms/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	ComAnnouncementEn()
}

func ComAnnouncementEn() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath(common.CmsWebsocketStreamsProdUrl),
		common.WithWsStreamsApiKey("Your API Key"),
		common.WithWsStreamsApiSecret("Your API Secret"),
	)

	wsClient := client.NewBinanceCmsClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ComAnnouncementEn().Execute()
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.ComAnnouncementEnResponse) {
		b, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received message: %s\n", string(b))
	})

	log.Println("Subscribed. Waiting 10 seconds...")
	time.Sleep(10 * time.Second)

	log.Println("Unsubscribing from stream...")
	handler.Unsubscribe()

	log.Println("Closing WebSocket connection...")
	err = wsClient.WebsocketStreams.CloseWebSocketStreamConnection()
	if err != nil {
		log.Fatalf("Error closing WebSocket connection: %v", err)
	}
}
```

## Testing

To run the test cases, use the following command:

```bash
go test ./tests/...
```

The tests cover: 
- WebSocket Streams endpoints
- Error handling and edge cases

## Contributing

Contributions are welcome!

Since this repository contains auto-generated code, we encourage you to start by opening a GitHub issue to discuss your ideas or suggest improvements. This helps ensure that changes align with the project's goals and auto-generation processes.

To contribute:

1. Open a GitHub issue describing your suggestion or the bug you've identified.
2. If it's determined that changes are necessary, the maintainers will merge the changes into the main branch.

Please ensure that all tests pass if you're making a direct contribution. Submit a pull request only after discussing and confirming the change.

Thank you for your contributions!

## Disclaimer

This SDK is provided by Binance on an "as is" and "as available" basis for use at your own risk. Binance makes no representations or warranties of any kind, whether express or implied, as to the operation of the SDK, its accuracy, reliability, completeness, or fitness for any particular purpose.

To the fullest extent permitted by law, Binance shall not be liable for any losses, damages, or expenses of any kind arising from or in connection with your use of, or inability to use, this SDK, including but not limited to any financial losses resulting from errors, bugs, interruptions, or inaccuracies in the SDK.

Your use of this SDK to access the Binance Platform is subject to the Binance API Key Terms and the Binance Terms of Use, which shall prevail in the event of any conflict with this disclaimer. You are solely responsible for any orders or transactions executed through the Binance Platform using this SDK.

This SDK is not intended to constitute investment advice or a recommendation to buy, sell, or hold any digital asset. You should independently evaluate and verify all information before acting.

- [Binance Terms of Use](https://www.binance.com/en/terms)
- [Binance API Key Terms](https://www.binance.com/en/about-legal/terms-binance-api)

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.