# Binance Spot Go Connector

This is a lightweight library that works as a connector to [Binance public API](https://github.com/binance/binance-spot-api-docs)

## Supported API Endpoints:
- Account/Trade: `account.go`
- Wallet: `wallet.go`
- Margin Account/Trade: `margin.go`
- Market Data: `market.go`
- Sub-Accounts: `subaccount.go`
- Staking: `staking.go`
- Websocket Market/User Data Stream: `websocket.go`
- Websocket User Data Stream: `user_stream.go`

## Installation
```shell
go get github.com/binance/binance-connector-go
```

To reference the package in your code, use the following import statement:
```golang
import (
    "github.com/binance/binance-connector-go"
)
```
## Authentication
```go
// The Client can be initiated with apiKey, secretKey and baseURL.
// The baseURL is optional. If not specified, it will default to "https://api.binance.com".
client := binance_connector.NewClient("yourApiKey", "yourSecretKey")
```

## Extra Options
```go
client := binance_connector.NewClient("yourApiKey", "yourSecretKey", "https://api.binance.com")

// Debug Mode
client.Debug = true

// TimeOffset (in milliseconds) - used to adjust the request timestamp by subtracting/adding the current time with it:
client.TimeOffset = -1000 // implies adding: request timestamp = current time - (-1000)
```

## REST API

Create an order example

```go
package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	apiKey := "yourApiKey"
	secretKey := "yourSecretKey"
	baseURL := "https://testnet.binance.vision"

	// Initialise the client
	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// Create new order
	newOrder, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side("BUY").Type("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(newOrder))
}
```

Please find more examples for each supported endpoint in the `examples` folder.

## Websocket Stream
Initialising Websocket Client
- Websocket Client can be initialized with 2 parameters, `NewWebsocketStreamClient(isCombined, baseURL)`:
- `isCombined` is a MANDATORY boolean value that specifies whether you are calling a combined stream or not.
  - If `isCombined` is set to `true`, `"/stream?streams="` will be appended to the `baseURL` to allow for Combining streams.
  - Otherwise, if set to `false`, `"/ws/"` will be appended to the `baseURL`.
- `baseURL` is an OPTIONAL string value that determines the base URL to use for the websocket connection.
  - If `baseURL` is not set, it will default to the Live Exchange URL: `"wss://stream.binance.com:9443"`.

```go
// Initialise Websocket Client with Production baseURL and false for "isCombined" parameter

websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, "wss://testnet.binance.vision")

// Initialise Websocket Client with Production baseURL and true for "isCombined" parameter

websocketStreamClient := binance_connector.NewWebsocketStreamClient(true)
```

Diff. Depth Stream Example

```go
package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	// Initialise Websocket Client with Testnet BaseURL and false for "isCombined" parameter
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, "wss://testnet.binance.vision")

	wsDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}

	errHandler := func(err error) {
		fmt.Println(err)
	}

	// Depth stream subscription
	doneCh, stopCh, err := websocketStreamClient.WsDepthServe("BNBUSDT", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(30 * time.Second)
		stopCh <- struct{}{} // use stopCh to stop streaming
	}()

	<-doneCh
}
```

Combined Diff. Depth Stream Example

```go
package main

import (
	"fmt"
	"time"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	// Set isCombined parameter to true as we are using Combined Depth Stream
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(true)

	wsCombinedDepthHandler := func(event *binance_connector.WsDepthEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	// Use WsCombinedDepthServe to subscribe to multiple streams
	doneCh, stopCh, err := websocketStreamClient.WsCombinedDepthServe([]string{"LTCBTC", "BTCUSDT", "MATICUSDT"}, wsCombinedDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	<-doneCh
}
```

## Websocket API

```go
func OCOHistoryExample() {
	// Initialise Websocket API Client
	client := binance_connector.NewWebsocketAPIClient("api_key", "secret_key")
	// Connect to Websocket API
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	// Send request to Websocket API
	response, err := client.NewAccountOCOHistoryService().Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Print the response
	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
```

## Base URL
- Binance provides alternative Production URLs in case of performance issues:
  - https://api1.binance.com
  - https://api2.binance.com
  - https://api3.binance.com

## Testnet Support
- In order to use the Testnet, simply set the `baseURL` to "https://testnet.binance.vision"
- You can find step-by-step instructions on how to use the get a Testnet API and Secret Key [here](https://dev.binance.vision/t/binance-testnet-environments/99)

## Pretty Print vs PrintLn
- The `fmt.Println(<response>)` function will print the struct in a single line, which is not very readable.
- The `fmt.Println(binance_connector.PrettyPrint(<response>))` function will print the struct, including both the key and value, in a multi-line format which is more easily readable.

### Regular PrintLn Example Output
```bash
&{depthUpdate 1680092520368 LTCBTC 1989614201 1989614210 [{0.00322300 70.96700000} {0.00322200 52.57100000} {0.00322000 248.64000000} {0.00321900 34.98300000}] [{0.00322600 71.52600000} {0.00323400 53.88900000} {0.00323500 27.37000000}]}
&{depthUpdate 1680092521368 LTCBTC 1989614211 1989614212 [{0.00320700 197.10100000} {0.00320100 15.76800000}] []}
&{depthUpdate 1680092522368 LTCBTC 1989614213 1989614224 [{0.00322300 86.15400000} {0.00322200 37.38400000} {0.00322100 252.53900000} {0.00322000 60.01300000}] [{0.00322800 75.48400000} {0.00322900 254.84500000} {0.00323000 8.74700000} {0.00323100 37.42800000}]}
&{depthUpdate 1680092523369 LTCBTC 1989614225 1989614226 [{0.00322300 103.57400000}] [{0.00399500 11.75400000}]}
&{depthUpdate 1680092524369 LTCBTC 1989614227 1989614276 [{0.00322500 0.00000000} {0.00322400 101.32700000} {0.00322300 138.82600000} {0.00322200 58.49100000} {0.00322100 249.65400000} {0.00321900 47.34800000} {0.00317800 16.08500000} {0.00317500 38.36500000}] [{0.00322500 75.14300000} {0.00322600 48.19100000} {0.00322700 44.97900000} {0.00322800 242.74300000} {0.00322900 20.73400000} {0.00532700 0.18900000} {0.00779700 0.05600000}]}
```

### binance_connector.PrettyPrint Example Output
```bash
{
	"e": "depthUpdate",
	"E": 1680092041346,
	"s": "LTCBTC",
    "U": 1989606566,
	"u": 1989606596,
	"b": [
		{
			"Price": "0.00322800",
			"Quantity": "83.05100000"
		},
		{
			"Price": "0.00322700",
			"Quantity": "12.50200000"
		},
		{
			"Price": "0.00322500",
			"Quantity": "48.53700000"
		},
		{
			"Price": "0.00322400",
			"Quantity": "244.13500000"
		}
	],
	"a": [
		{
			"Price": "0.00322900",
			"Quantity": "79.52900000"
		},
		{
			"Price": "0.00323000",
			"Quantity": "42.68400000"
		},
		{
			"Price": "0.00323100",
			"Quantity": "68.75500000"
		}
	]
}
{
	"e": "depthUpdate",
	"E": 1680092042346,
	"s": "LTCBTC",
	"U": 1989606597,
	"u": 1989606611,
	"b": [
		{
			"Price": "0.00321400",
			"Quantity": "0.24700000"
		},
		{
			"Price": "0.00318000",
			"Quantity": "1.91600000"
		}
	],
	"a": [
		{
			"Price": "0.00322900",
			"Quantity": "79.27900000"
		}
	]
}
```

## Limitations
Futures and European Options APIs are not supported:
- /fapi/*
- /dapi/*
- /vapi/*
- Associated Websocket Market and User Data Streams

## Contributing
Contributions are welcome.<br/>
If you've found a bug within this project, please open an issue to discuss what you would like to change.<br/>
If it's an issue with the API, please open a topic at [Binance Developer Community](https://dev.binance.vision)
