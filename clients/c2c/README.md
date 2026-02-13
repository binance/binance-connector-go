# Binance Go C2C SDK

[![Build Status](https://img.shields.io/github/actions/workflow/status/binance/binance-connector-go/ci.yaml)](https://github.com/binance/binance-connector-go/actions)
[![Open Issues](https://img.shields.io/github/issues/binance/binance-connector-go)](https://github.com/binance/binance-connector-go/issues)
[![Go Version](https://img.shields.io/github/go-mod/go-version/binance/binance-connector-go)](https://github.com/binance/binance-connector-go)
[![Known Vulnerabilities](https://img.shields.io/badge/security-scanned-brightgreen)](https://github.com/binance/binance-connector-go/security)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a client library for the Binance C2C API, enabling developers to interact programmatically with Binance's C2C trading platform. The library provides tools to query Fiat transaction history through the REST API:

- [REST API](./src/restapi/rest_api.go)

## Table of Contents

- [Supported Features](#supported-features)
- [Installation](#installation)
- [Documentation](#documentation)
- [REST APIs](#rest-apis)
- [Testing](#testing)
- [Migration Guide](#migration-guide)
- [Contributing](#contributing)
- [License](#license)

## Supported Features

- REST API Endpoints:
  - `/sapi/v1/c2c/*`
- Inclusion of test cases and examples for quick onboarding.

## Installation

To use this library, ensure you have Go installed (version **1.24** or higher is recommended). You can install the library using the following command:

```bash
go get github.com/binance/binance-connector-go/clients/c2c
```

## Documentation

For detailed information, refer to the [Binance API Documentation](https://developers.binance.com/docs/c2c).

### REST APIs

All REST API endpoints are available through the [`restapi`](./src/restapi/rest_api.go) module. Note that some endpoints require authentication using your Binance API credentials.

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/c2c"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetC2CTradeHistory()
}

func GetC2CTradeHistory() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.C2CRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceC2CClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.C2CAPI.GetC2CTradeHistory(context.Background()).Execute()
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

## Testing

To run the test cases, use the following command:

```bash
go test ./tests/...
```

The tests cover: 
- REST API endpoints
- Error handling and edge cases

## Migration Guide

For information on migrating from previous versions of the Binance Go C2C SDK, please refer to the [Migration Guide](./docs/migration-guide.md).

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