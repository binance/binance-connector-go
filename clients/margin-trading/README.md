# Binance Go Margin Trading SDK

[![Build Status](https://img.shields.io/github/actions/workflow/status/binance/binance-connector-go/ci.yaml)](https://github.com/binance/binance-connector-go/actions)
[![Open Issues](https://img.shields.io/github/issues/binance/binance-connector-go)](https://github.com/binance/binance-connector-go/issues)
[![Go Version](https://img.shields.io/github/go-mod/go-version/binance/binance-connector-go)](https://github.com/binance/binance-connector-go)
[![Known Vulnerabilities](https://img.shields.io/badge/security-scanned-brightgreen)](https://github.com/binance/binance-connector-go/security)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a client library for the Binance Margin Trading SDK API, enabling developers to interact programmatically with Binance's Margin Trading trading platform. The library provides tools for retrieving market data, executing trades, and managing orders through two distinct endpoints:

- [REST API](./src/restapi/rest_api.go)
- [Websocket Stream](./src/websocketstreams/websocket_streams.go)

## Table of Contents

- [Supported Features](#supported-features)
- [Installation](#installation)
- [Documentation](#documentation)
- [REST APIs](#rest-apis)
- [Websocket Streams](#websocket-streams)
- [Testing](#testing)
- [Migration Guide](#migration-guide)
- [Contributing](#contributing)
- [License](#license)

## Supported Features

- REST API Endpoints:
  - `/sapi/v1/margin/*`
  - `/sapi/v1/bnbBurn/*`
  - `/sapi/v1/userDataStream/*`
- Inclusion of test cases and examples for quick onboarding.

## Installation

To use this library, ensure you have Go installed (version **1.24** or higher is recommended). You can install the library using the following command:

```bash
go get github.com/binance/binance-connector-go/clients/margintrading
```

## Documentation

For detailed information, refer to the [Binance API Documentation](https://developers.binance.com/docs/margin_trading).

### REST APIs

All REST API endpoints are available through the [`restapi`](./src/restapi/rest_api.go) module. Note that some endpoints require authentication using your Binance API credentials.

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetSummaryOfMarginAccount()
}

func GetSummaryOfMarginAccount() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.MarginTradingRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceMarginTradingClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.GetSummaryOfMarginAccount(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

More examples can be found in the [`examples/restapi`](./examples/restapi/) folder.

#### Configuration Options

The REST API supports the following advanced configuration options:

- `Timeout`: Timeout for requests in milliseconds (default: 1000 ms).
- `Proxy`: Proxy configuration:
  - `Host`: Proxy server hostname.
  - `Port`: Proxy server port.
  - `Protocol`: Proxy protocol (http or https).
  - `Auth`: Proxy authentication credentials:
    - `Username`: Proxy username.
    - `Password`: Proxy password.
- `KeepAlive`: Enable HTTP keep-alive (default: true).
- `Compression`: Enable response compression (default: true).
- `Retries`: Number of retry attempts for failed requests (default: 3).
- `Backoff`: Delay in milliseconds between retries (default: 1000 ms).
- `HTTPSAgent`: Custom HTTPS agent for advanced TLS configuration.
- `PrivateKey`: RSA or ED25519 private key for authentication.
- `PrivateKeyPassphrase`: Passphrase for the private key, if encrypted.

##### Timeout

You can configure a timeout for requests in milliseconds. If the request exceeds the specified timeout, it will be aborted. See the [Timeout example](./docs/restapi/timeout.md) for detailed usage.

##### Proxy

The REST API supports HTTP/HTTPS proxy configurations. See the [Proxy example](./docs/restapi/proxy.md) for detailed usage.

##### Keep-Alive

Enable HTTP keep-alive for persistent connections. See the [Keep-Alive example](./docs/restapi/keep-alive.md) for detailed usage.

##### Compression

Enable or disable response compression. See the [Compression example](./docs/restapi/compression.md) for detailed usage.

##### Retries

Configure the number of retry attempts and delay in milliseconds between retries for failed requests. See the [Retries example](./docs/restapi/retries.md) for detailed usage.

##### HTTPS Agent

Customize the HTTPS agent for advanced TLS configurations. See the [HTTPS Agent example](./docs/restapi/https-agent.md) for detailed usage.

##### Key Pair Based Authentication

The REST API supports key pair-based authentication for secure communication. You can use `RSA` or `ED25519` keys for signing requests. See the [Key Pair Based Authentication example](./docs/restapi/key-pair-authentication.md) for detailed usage.

##### Certificate Pinning

To enhance security, you can use certificate pinning with the `HTTPSAgent` option in the configuration. This ensures the client only communicates with servers using specific certificates. See the [Certificate Pinning example](./docs/restapi/certificate-pinning.md) for detailed usage.

#### Error Handling

The REST API provides detailed error types to help you handle issues effectively:

- `ConnectorClientError`: General client error.
- `RequiredError`: Thrown when a required parameter is missing.
- `UnauthorizedError`: Indicates missing or invalid authentication credentials.
- `ForbiddenError`: Access to the requested resource is forbidden.
- `TooManyRequestsError`: Rate limit exceeded.
- `RateLimitBanError`: IP address banned for exceeding rate limits.
- `ServerError`: Internal server error.
- `NetworkError`: Issues with network connectivity.
- `NotFoundError`: Resource not found.
- `BadRequestError`: Invalid request.

See the [Error Handling example](./docs/restapi/error-handling.md) for detailed usage.


### Websocket Streams

WebSocket Streams in Margin Trading is used for subscribing to risk and trade data streams. Use the [websocket-streams](./src/websocketstreams/websocket_streams.go) module to interact with it.

#### Configuration Options

The WebSocket Streams API supports the following advanced configuration options:

- `ReconnectDelay`: Delay (ms) between reconnections.
- `Compression`: Enable response compression.
- `Mode`: Choose between `single` and `pool` connection modes.
  - `Single`: A single WebSocket connection.
  - `Pool`: A pool of WebSocket connections.
- `PoolSize`: Define the number of WebSocket connections in pool mode.
- `Agent`: Customize the WebSocket Agent for advanced configurations.

#### Subscribe to Risk and Trade Data Streams

You can consume the risk and trade data stream, which sends account-level events such as account and order updates. First create a listen-key via REST API; then:

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	riskData()
}

func riskData() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath("wss://fstream.binance.com/pm/stream"),
	)

	wsClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.RiskData("listenKey", "") // You can replace `RiskData` with `TradeData` to subscribe to trade data stream
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.RiskDataStreamEventsResponse) {
		b, _ := json.MarshalIndent(message, "", "  ")
		log.Printf("Received message: %s\n", string(b))
	})

	for {
		time.Sleep(1 * time.Second)
	}
}
```

#### Unsubscribing from Streams

You can unsubscribe from specific WebSocket streams using the `unsubscribe` method. This is useful for managing active subscriptions without closing the connection.

```go
package main

import (
	"encoding/json"
	"log"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	riskData()
}

func riskData() {
	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsStreamsBasePath("wss://fstream.binance.com/pm/stream"),
	)

	wsClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
		client.WithWebsocketStreams(configuration),
	)

	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Fatalf("Error connecting to WebSocket: %v", err)
	}

	handler, err := wsClient.WebsocketStreams.RiskData("listenKey", "")
	if err != nil {
		log.Fatalf("Error subscribing to stream: %v", err)
	}

	handler.On("message", func(message models.RiskDataStreamEventsResponse) {
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
- REST API endpoints
- Error handling and edge cases

## Migration Guide

For information on migrating from previous versions of the Binance Go Margin Trading SDK, please refer to the [Migration Guide](./docs/migration-guide.md).

## Contributing

Contributions are welcome!

Since this repository contains auto-generated code, we encourage you to start by opening a GitHub issue to discuss your ideas or suggest improvements. This helps ensure that changes align with the project's goals and auto-generation processes.

To contribute:

1. Open a GitHub issue describing your suggestion or the bug you've identified.
2. If it's determined that changes are necessary, the maintainers will merge the changes into the main branch.

Please ensure that all tests pass if you're making a direct contribution. Submit a pull request only after discussing and confirming the change.

Thank you for your contributions!

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.