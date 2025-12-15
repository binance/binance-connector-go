# \WebSocketStreamsAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggTrade**](WebSocketStreamsAPI.md#AggTrade) | /&lt;symbol&gt;@aggTrade | WebSocket Aggregate Trade Streams
[**AllMarketRollingWindowTicker**](WebSocketStreamsAPI.md#AllMarketRollingWindowTicker) | /!ticker_&lt;windowSize&gt;@arr | WebSocket All Market Rolling Window Statistics Streams
[**AllMiniTicker**](WebSocketStreamsAPI.md#AllMiniTicker) | /!miniTicker@arr | WebSocket All Market Mini Tickers Stream
[**AvgPrice**](WebSocketStreamsAPI.md#AvgPrice) | /&lt;symbol&gt;@avgPrice | WebSocket Average Price
[**BookTicker**](WebSocketStreamsAPI.md#BookTicker) | /&lt;symbol&gt;@bookTicker | WebSocket Individual Symbol Book Ticker Streams
[**DiffBookDepth**](WebSocketStreamsAPI.md#DiffBookDepth) | /&lt;symbol&gt;@depth@&lt;updateSpeed&gt; | WebSocket Diff. Depth Stream
[**Kline**](WebSocketStreamsAPI.md#Kline) | /&lt;symbol&gt;@kline_&lt;interval&gt; | WebSocket Kline/Candlestick Streams for UTC
[**KlineOffset**](WebSocketStreamsAPI.md#KlineOffset) | /&lt;symbol&gt;@kline_&lt;interval&gt;@+08:00 | WebSocket Kline/Candlestick Streams with timezone offset
[**MiniTicker**](WebSocketStreamsAPI.md#MiniTicker) | /&lt;symbol&gt;@miniTicker | WebSocket Individual Symbol Mini Ticker Stream
[**PartialBookDepth**](WebSocketStreamsAPI.md#PartialBookDepth) | /&lt;symbol&gt;@depth&lt;levels&gt;@&lt;updateSpeed&gt; | WebSocket Partial Book Depth Streams
[**RollingWindowTicker**](WebSocketStreamsAPI.md#RollingWindowTicker) | /&lt;symbol&gt;@ticker_&lt;windowSize&gt; | WebSocket Individual Symbol Rolling Window Statistics Streams
[**Ticker**](WebSocketStreamsAPI.md#Ticker) | /&lt;symbol&gt;@ticker | WebSocket Individual Symbol Ticker Streams
[**Trade**](WebSocketStreamsAPI.md#Trade) | /&lt;symbol&gt;@trade | WebSocket Trade Streams


## AggTrade

WebSocket Aggregate Trade Streams


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.AggTrade().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.AggTrade``: %v\n", err)
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

WebSocket All Market Rolling Window Statistics Streams


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	windowSize := models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h // AllMarketRollingWindowTickerWindowSizeParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.AllMarketRollingWindowTicker().WindowSize(windowSize).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.AllMarketRollingWindowTicker``: %v\n", err)
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

WebSocket All Market Mini Tickers Stream


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.AllMiniTicker().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.AllMiniTicker``: %v\n", err)
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

WebSocket Average Price


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.AvgPrice().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.AvgPrice``: %v\n", err)
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


## BookTicker

WebSocket Individual Symbol Book Ticker Streams


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.BookTicker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.BookTicker``: %v\n", err)
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

WebSocket Diff. Depth Stream


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "100ms" // string | 1000ms or 100ms (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.DiffBookDepth().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.DiffBookDepth``: %v\n", err)
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
 **updateSpeed** | **string** | 1000ms or 100ms | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## Kline

WebSocket Kline/Candlestick Streams for UTC


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
	"github.com/binance/binance-connector-go/common/common"
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
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.Kline().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.Kline``: %v\n", err)
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

WebSocket Kline/Candlestick Streams with timezone offset


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
	"github.com/binance/binance-connector-go/common/common"
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
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.KlineOffset().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.KlineOffset``: %v\n", err)
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

WebSocket Individual Symbol Mini Ticker Stream


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.MiniTicker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.MiniTicker``: %v\n", err)
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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	levels := models.PartialBookDepthLevelsParameterLevels5 // PartialBookDepthLevelsParameter | 
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "100ms" // string | 1000ms or 100ms (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.PartialBookDepth().Symbol(symbol).Levels(levels).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.PartialBookDepth``: %v\n", err)
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
 **updateSpeed** | **string** | 1000ms or 100ms | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## RollingWindowTicker

WebSocket Individual Symbol Rolling Window Statistics Streams


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
	"github.com/binance/binance-connector-go/common/common"
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
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.RollingWindowTicker().Symbol(symbol).WindowSize(windowSize).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.RollingWindowTicker``: %v\n", err)
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

WebSocket Individual Symbol Ticker Streams


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.Ticker().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.Ticker``: %v\n", err)
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

WebSocket Trade Streams


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
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	symbol := "bnbusdt" // string | Symbol to query
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceSpotClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebSocketStreamsAPI.Trade().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebSocketStreamsAPI.Trade``: %v\n", err)
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

