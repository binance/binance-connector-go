# \PublicAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**DiffBookDepthStreams**](PublicAPI.md#DiffBookDepthStreams) | /&lt;symbol&gt;@depth@&lt;updateSpeed&gt; | Diff Book Depth Streams
[**Hour24Ticker**](PublicAPI.md#Hour24Ticker) | /&lt;symbol&gt;@optionTicker&lt;expirationDate&gt; | 24-hour TICKER
[**IndividualSymbolBookTickerStreams**](PublicAPI.md#IndividualSymbolBookTickerStreams) | /&lt;symbol&gt;@bookTicker | Individual Symbol Book Ticker Streams
[**PartialBookDepthStreams**](PublicAPI.md#PartialBookDepthStreams) | /&lt;symbol&gt;@depth&lt;level&gt;@&lt;updateSpeed&gt; | Partial Book Depth Streams
[**TradeStreams**](PublicAPI.md#TradeStreams) | /&lt;symbol&gt;@optionTrade | Trade Streams


## DiffBookDepthStreams

Diff Book Depth Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	updateSpeed := models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthStreamsUpdateSpeedParameter | WebSocket stream update speed
	id := int32(532601580) // int32 | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.PublicAPI.DiffBookDepthStreams().Symbol(symbol).UpdateSpeed(updateSpeed).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PublicAPI.DiffBookDepthStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.DiffBookDepthStreamsResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The symbol parameter | 
 **updateSpeed** | [**DiffBookDepthStreamsUpdateSpeedParameter**](DiffBookDepthStreamsUpdateSpeedParameter.md) | WebSocket stream update speed | 
 **id** | **int32** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Hour24Ticker

24-hour TICKER


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := int32(532601580) // int32 | Unique WebSocket request ID. (optional)
	expirationDate := "251230" // string | The expiration date parameter (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.PublicAPI.Hour24Ticker().Symbol(symbol).Id(id).ExpirationDate(expirationDate).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PublicAPI.Hour24Ticker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.Hour24TickerResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The symbol parameter | 
 **id** | **int32** | Unique WebSocket request ID. | 
 **expirationDate** | **string** | The expiration date parameter | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## IndividualSymbolBookTickerStreams

Individual Symbol Book Ticker Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := int32(532601580) // int32 | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.PublicAPI.IndividualSymbolBookTickerStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PublicAPI.IndividualSymbolBookTickerStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.IndividualSymbolBookTickerStreamsResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The symbol parameter | 
 **id** | **int32** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PartialBookDepthStreams

Partial Book Depth Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	level := models.PartialBookDepthStreamsLevelParameterLevel5 // PartialBookDepthStreamsLevelParameter | The level parameter
	updateSpeed := models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthStreamsUpdateSpeedParameter | WebSocket stream update speed
	id := int32(532601580) // int32 | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Symbol(symbol).Level(level).UpdateSpeed(updateSpeed).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PublicAPI.PartialBookDepthStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.PartialBookDepthStreamsResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The symbol parameter | 
 **level** | [**PartialBookDepthStreamsLevelParameter**](PartialBookDepthStreamsLevelParameter.md) | The level parameter | 
 **updateSpeed** | [**DiffBookDepthStreamsUpdateSpeedParameter**](DiffBookDepthStreamsUpdateSpeedParameter.md) | WebSocket stream update speed | 
 **id** | **int32** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradeStreams

Trade Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := int32(532601580) // int32 | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingOptionsClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.PublicAPI.TradeStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `PublicAPI.TradeStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.TradeStreamsResponse) {
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

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The symbol parameter | 
 **id** | **int32** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

