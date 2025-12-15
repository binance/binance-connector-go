# Binance Go Connectors

[![Build Status](https://img.shields.io/github/actions/workflow/status/binance/binance-connector-go/ci.yaml)](https://github.com/binance/binance-connector-go/actions)
[![Open Issues](https://img.shields.io/github/issues/binance/binance-connector-go)](https://github.com/binance/binance-connector-go/issues)
[![Go Version](https://img.shields.io/github/go-mod/go-version/binance/binance-connector-go)](https://github.com/binance/binance-connector-go)
[![Known Vulnerabilities](https://img.shields.io/badge/security-scanned-brightgreen)](https://github.com/binance/binance-connector-go/security)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Collection of auto-generated Go SDK for Binance APIs.

## Prerequisites

Before using the SDK, ensure you have:
- **Go** (version 1.24 or later)
- **Go Modules** (for dependency management)

## Available SDK

- [binance-connector-algo](./clients/algo) - Algo Trading connector
- [binance-connector-c2c](./clients/c2c/) - C2
- [binance-sdk-convert](./clients/convert/) - Convert connector
- [binance-sdk-copy-trading](./clients/copy-trading/) - Copy Trading connector
- [binance-sdk-crypto-loan](./clients/crypto-loan/) - Crypto Loan connector
- [binance-sdk-derivatives-trading-coin-futures](./clients/derivatives-trading-coin-futures/) - Coin Futures Trading connector
- [binance-sdk-derivatives-trading-options](./clients/derivatives-trading-options/) - Options Trading connector
- [binance-sdk-derivatives-trading-portfolio-margin](./clients/derivatives-trading-portfolio-margin/) - Portfolio Margin Futures Trading connector
- [binance-sdk-derivatives-trading-portfolio-margin-pro](./clients/derivatives-trading-portfolio-margin-pro/) - Portfolio Margin Pro Trading connector
- [binance-sdk-derivatives-trading-usds-futures](./clients/derivatives-trading-usds-futures/) - USDs Futures Trading connector
- [binance-sdk-dual-investment](./clients/dual-investment/) - Dual Investment connector
- [binance-sdk-fiat](./clients/fiat/) - Fiat connector
- [binance-sdk-gift-card](./clients/gift-card/) - Gift Card connector
- [binance-sdk-margin-trading](./clients/margin-trading/) - Margin Trading connector
- [binance-sdk-mining](./clients/mining/) - Mining connector
- [binance-sdk-nft](./clients/nft/) - NFT connector
- [binance-sdk-pay](./clients/pay/) - Pay connector
- [binance-sdk-rebate](./clients/rebate/) - Rebate connector
- [binance-sdk-simple-earn](./clients/simple-earn/) - Simple Earn connector
- [binance-sdk-spot](./clients/spot/) - Spot Trading connector
- [binance-sdk-staking](./clients/staking/) - Staking connector
- [binance-sdk-sub-account](./clients/sub-account/) - Sub Account connector
- [binance-sdk-vip-loan](./clients/vip-loan/) - VIP Loan connector
- [binance-sdk-wallet](./clients/wallet/) - Wallet connector

## Documentation

For detailed information, refer to the [Binance API Documentation](https://developers.binance.com).

## Installation

Each connector is published as a separate Python package. You can install them via `go get`. For example:

```bash
go get github.com/binance/binance-connector-go/clients/spot
```

Or to install multiple connectors:

```bash
go get github.com/binance/binance-connector-go/clients/spot github.com/binance/binance-connector-go/clients/margintrading github.com/binance/binance-connector-go/clients/wallet
```

## Contributing

Since this repository contains auto-generated code using OpenAPI Generator, we encourage you to:

1. Open a GitHub issue to discuss your ideas or report bugs
2. Allow maintainers to implement necessary changes through the code generation process

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE.md) file for details.
