# \DefaultAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggTrade**](DefaultAPI.md#AggTrade) | /&lt;symbol&gt;@aggTrade | Aggregate Trade Streams
[**AllMarketRollingWindowTicker**](DefaultAPI.md#AllMarketRollingWindowTicker) | /!ticker_&lt;windowSize&gt;@arr | All Market Rolling Window Statistics Streams
[**AllMiniTicker**](DefaultAPI.md#AllMiniTicker) | /!miniTicker@arr | All Market Mini Tickers Stream
[**AvgPrice**](DefaultAPI.md#AvgPrice) | /&lt;symbol&gt;@avgPrice | Average Price
[**BlockTrade**](DefaultAPI.md#BlockTrade) | /&lt;symbol&gt;@blockTrade | Block Trade Streams
[**BookTicker**](DefaultAPI.md#BookTicker) | /&lt;symbol&gt;@bookTicker | Individual Symbol Book Ticker Streams
[**DiffBookDepth**](DefaultAPI.md#DiffBookDepth) | /&lt;symbol&gt;@depth@&lt;updateSpeed&gt; | Diff. Depth Stream
[**Kline**](DefaultAPI.md#Kline) | /&lt;symbol&gt;@kline_&lt;interval&gt; | Kline/Candlestick Streams for UTC
[**KlineOffset**](DefaultAPI.md#KlineOffset) | /&lt;symbol&gt;@kline_&lt;interval&gt;@+08:00 | Kline/Candlestick Streams with timezone offset
[**MiniTicker**](DefaultAPI.md#MiniTicker) | /&lt;symbol&gt;@miniTicker | Individual Symbol Mini Ticker Stream
[**PartialBookDepth**](DefaultAPI.md#PartialBookDepth) | /&lt;symbol&gt;@depth&lt;levels&gt;@&lt;updateSpeed&gt; | WebSocket Partial Book Depth Streams
[**ReferencePrice**](DefaultAPI.md#ReferencePrice) | /&lt;symbol&gt;@referencePrice | Reference Price Streams
[**RollingWindowTicker**](DefaultAPI.md#RollingWindowTicker) | /&lt;symbol&gt;@ticker_&lt;windowSize&gt; | Individual Symbol Rolling Window Statistics Streams
[**Ticker**](DefaultAPI.md#Ticker) | /&lt;symbol&gt;@ticker | Individual Symbol Ticker Streams
[**Trade**](DefaultAPI.md#Trade) | /&lt;symbol&gt;@trade | Trade Streams


## AggTrade

Aggregate Trade Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AggTrade().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AggTrade``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AggTradeResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllMarketRollingWindowTicker

All Market Rolling Window Statistics Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	windowSize := models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h // AllMarketRollingWindowTickerWindowSizeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllMarketRollingWindowTicker().WindowSize(windowSize).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllMarketRollingWindowTicker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllMarketRollingWindowTickerResponse) {
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
 **windowSize** | [**AllMarketRollingWindowTickerWindowSizeParameter**](AllMarketRollingWindowTickerWindowSizeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllMiniTicker

All Market Mini Tickers Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllMiniTicker().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllMiniTicker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllMiniTickerResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AvgPrice

Average Price


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AvgPrice().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AvgPrice``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AvgPriceResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## BlockTrade

Block Trade Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.BlockTrade().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.BlockTrade``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.BlockTradeResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## BookTicker

Individual Symbol Book Ticker Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.BookTicker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.BookTicker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.BookTickerResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## DiffBookDepth

Diff. Depth Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := models.DiffBookDepthUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthUpdateSpeedParameter | Optional stream update speed suffix (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.DiffBookDepth().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.DiffBookDepth``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.DiffBookDepthResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | [**DiffBookDepthUpdateSpeedParameter**](DiffBookDepthUpdateSpeedParameter.md) | Optional stream update speed suffix | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Kline

Kline/Candlestick Streams for UTC


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	interval := models.KlineIntervalParameterInterval1s // KlineIntervalParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.Kline().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.Kline``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.KlineResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **interval** | [**KlineIntervalParameter**](KlineIntervalParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## KlineOffset

Kline/Candlestick Streams with timezone offset


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	interval := models.KlineIntervalParameterInterval1s // KlineIntervalParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.KlineOffset().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.KlineOffset``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.KlineOffsetResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **interval** | [**KlineIntervalParameter**](KlineIntervalParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MiniTicker

Individual Symbol Mini Ticker Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.MiniTicker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MiniTicker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.MiniTickerResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PartialBookDepth

WebSocket Partial Book Depth Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	levels := models.PartialBookDepthLevelsParameterLevels5 // PartialBookDepthLevelsParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := models.DiffBookDepthUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthUpdateSpeedParameter | Optional stream update speed suffix (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.PartialBookDepth().Symbol(symbol).Levels(levels).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.PartialBookDepth``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.PartialBookDepthResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **levels** | [**PartialBookDepthLevelsParameter**](PartialBookDepthLevelsParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | [**DiffBookDepthUpdateSpeedParameter**](DiffBookDepthUpdateSpeedParameter.md) | Optional stream update speed suffix | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ReferencePrice

Reference Price Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bazusd" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ReferencePrice().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.ReferencePrice``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.ReferencePriceResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## RollingWindowTicker

Individual Symbol Rolling Window Statistics Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	windowSize := models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h // AllMarketRollingWindowTickerWindowSizeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.RollingWindowTicker().Symbol(symbol).WindowSize(windowSize).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.RollingWindowTicker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.RollingWindowTickerResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **windowSize** | [**AllMarketRollingWindowTickerWindowSizeParameter**](AllMarketRollingWindowTickerWindowSizeParameter.md) |  | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Ticker

Individual Symbol Ticker Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.Ticker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.Ticker``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.TickerResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Trade

Trade Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/spot"
	responseModels "github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.DefaultAPI.Trade().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.Trade``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.TradeResponse) {
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
 **symbol** | **string** | Symbol to query | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

