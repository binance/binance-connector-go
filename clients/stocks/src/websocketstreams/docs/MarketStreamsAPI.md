# \MarketStreamsAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**CalendarStream**](MarketStreamsAPI.md#CalendarStream) | /calendar | Calendar Stream
[**KlineStream**](MarketStreamsAPI.md#KlineStream) | /&lt;symbol&gt;@kline_&lt;interval&gt; | Kline Stream
[**PriceStream**](MarketStreamsAPI.md#PriceStream) | /price | Price Stream
[**QuoteStream**](MarketStreamsAPI.md#QuoteStream) | /&lt;symbol&gt;@quote | Quote Stream
[**TradabilityStream**](MarketStreamsAPI.md#TradabilityStream) | /&lt;symbol&gt;@tradability | Tradability Stream
[**TradingStatusStream**](MarketStreamsAPI.md#TradingStatusStream) | /&lt;symbol&gt;@tradingStatus | Trading Status Stream


## CalendarStream

Calendar Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.CalendarStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.CalendarStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.CalendarStreamResponse) {
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


## KlineStream

Kline Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US-equity ticker (UPPERCASE), e.g. `AAPL`.
	interval := "5m" // string | Kline interval — `5m` / `1h` / `1d` / `1w` / `1M`.
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.KlineStream().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.KlineStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.KlineStreamResponse) {
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
 **symbol** | **string** | US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;. | 
 **interval** | **string** | Kline interval — &#x60;5m&#x60; / &#x60;1h&#x60; / &#x60;1d&#x60; / &#x60;1w&#x60; / &#x60;1M&#x60;. | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## PriceStream

Price Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.PriceStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.PriceStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.PriceStreamResponse) {
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


## QuoteStream

Quote Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US-equity ticker (UPPERCASE), e.g. `AAPL`.
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.QuoteStream().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.QuoteStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.QuoteStreamResponse) {
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
 **symbol** | **string** | US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;. | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradabilityStream

Tradability Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US-equity ticker (UPPERCASE), e.g. `AAPL`.
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.TradabilityStream().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.TradabilityStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.TradabilityStreamResponse) {
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
 **symbol** | **string** | US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;. | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## TradingStatusStream

Trading Status Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/stocks"
	responseModels "github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "AAPL" // string | US-equity ticker (UPPERCASE), e.g. `AAPL`.
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceStocksClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect([]string{})
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.MarketStreamsAPI.TradingStatusStream().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `MarketStreamsAPI.TradingStatusStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.TradingStatusStreamResponse) {
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
 **symbol** | **string** | US-equity ticker (UPPERCASE), e.g. &#x60;AAPL&#x60;. | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

