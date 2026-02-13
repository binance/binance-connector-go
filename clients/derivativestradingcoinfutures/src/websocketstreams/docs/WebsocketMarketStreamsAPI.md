# \WebsocketMarketStreamsAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggregateTradeStreams**](WebsocketMarketStreamsAPI.md#AggregateTradeStreams) | /&lt;symbol&gt;@aggTrade | Aggregate Trade Streams
[**AllBookTickersStream**](WebsocketMarketStreamsAPI.md#AllBookTickersStream) | /!bookTicker | All Book Tickers Stream
[**AllMarketLiquidationOrderStreams**](WebsocketMarketStreamsAPI.md#AllMarketLiquidationOrderStreams) | /!forceOrder@arr | All Market Liquidation Order Streams
[**AllMarketMiniTickersStream**](WebsocketMarketStreamsAPI.md#AllMarketMiniTickersStream) | /!miniTicker@arr | All Market Mini Tickers Stream
[**AllMarketTickersStreams**](WebsocketMarketStreamsAPI.md#AllMarketTickersStreams) | /!ticker@arr | All Market Tickers Streams
[**ContinuousContractKlineCandlestickStreams**](WebsocketMarketStreamsAPI.md#ContinuousContractKlineCandlestickStreams) | /&lt;pair&gt;_&lt;contractType&gt;@continuousKline_&lt;interval&gt; | Continuous Contract Kline/Candlestick Streams
[**ContractInfoStream**](WebsocketMarketStreamsAPI.md#ContractInfoStream) | /!contractInfo | Contract Info Stream
[**DiffBookDepthStreams**](WebsocketMarketStreamsAPI.md#DiffBookDepthStreams) | /&lt;symbol&gt;@depth@&lt;updateSpeed&gt; | Diff. Book Depth Streams
[**IndexKlineCandlestickStreams**](WebsocketMarketStreamsAPI.md#IndexKlineCandlestickStreams) | /&lt;pair&gt;@indexPriceKline_&lt;interval&gt; | Index Kline/Candlestick Streams
[**IndexPriceStream**](WebsocketMarketStreamsAPI.md#IndexPriceStream) | /&lt;pair&gt;@indexPrice@&lt;updateSpeed&gt; | Index Price Stream
[**IndividualSymbolBookTickerStreams**](WebsocketMarketStreamsAPI.md#IndividualSymbolBookTickerStreams) | /&lt;symbol&gt;@bookTicker | Individual Symbol Book Ticker Streams
[**IndividualSymbolMiniTickerStream**](WebsocketMarketStreamsAPI.md#IndividualSymbolMiniTickerStream) | /&lt;symbol&gt;@miniTicker | Individual Symbol Mini Ticker Stream
[**IndividualSymbolTickerStreams**](WebsocketMarketStreamsAPI.md#IndividualSymbolTickerStreams) | /&lt;symbol&gt;@ticker | Individual Symbol Ticker Streams
[**KlineCandlestickStreams**](WebsocketMarketStreamsAPI.md#KlineCandlestickStreams) | /&lt;symbol&gt;@kline_&lt;interval&gt; | Kline/Candlestick Streams
[**LiquidationOrderStreams**](WebsocketMarketStreamsAPI.md#LiquidationOrderStreams) | /&lt;symbol&gt;@forceOrder | Liquidation Order Streams
[**MarkPriceKlineCandlestickStreams**](WebsocketMarketStreamsAPI.md#MarkPriceKlineCandlestickStreams) | /&lt;symbol&gt;@markPriceKline_&lt;interval&gt; | Mark Price Kline/Candlestick Streams
[**MarkPriceOfAllSymbolsOfAPair**](WebsocketMarketStreamsAPI.md#MarkPriceOfAllSymbolsOfAPair) | /&lt;pair&gt;@markPrice@&lt;updateSpeed&gt; | Mark Price of All Symbols of a Pair
[**MarkPriceStream**](WebsocketMarketStreamsAPI.md#MarkPriceStream) | /&lt;symbol&gt;@markPrice@&lt;updateSpeed&gt; | Mark Price Stream
[**PartialBookDepthStreams**](WebsocketMarketStreamsAPI.md#PartialBookDepthStreams) | /&lt;symbol&gt;@depth&lt;levels&gt;@&lt;updateSpeed&gt; | Partial Book Depth Streams


## AggregateTradeStreams

Aggregate Trade Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AggregateTradeStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.AggregateTradeStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AggregateTradeStreamsResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## AllBookTickersStream

All Book Tickers Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllBookTickersStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.AllBookTickersStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllBookTickersStreamResponse) {
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


## AllMarketLiquidationOrderStreams

All Market Liquidation Order Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketLiquidationOrderStreams().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.AllMarketLiquidationOrderStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllMarketLiquidationOrderStreamsResponse) {
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


## AllMarketMiniTickersStream

All Market Mini Tickers Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketMiniTickersStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.AllMarketMiniTickersStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllMarketMiniTickersStreamResponse) {
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


## AllMarketTickersStreams

All Market Tickers Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketTickersStreams().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.AllMarketTickersStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.AllMarketTickersStreamsResponse) {
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


## ContinuousContractKlineCandlestickStreams

Continuous Contract Kline/Candlestick Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "btcusdt" // string | The pair parameter
	contractType := "next_quarter" // string | The contractType parameter
	interval := "1m" // string | The interval parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Pair(pair).ContractType(contractType).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.ContinuousContractKlineCandlestickStreamsResponse) {
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
 **pair** | **string** | The pair parameter | 
 **contractType** | **string** | The contractType parameter | 
 **interval** | **string** | The interval parameter | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## ContractInfoStream

Contract Info Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContractInfoStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.ContractInfoStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.ContractInfoStreamResponse) {
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


## DiffBookDepthStreams

Diff. Book Depth Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "updateSpeed_example" // string | WebSocket stream update speed (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.DiffBookDepthStreams().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.DiffBookDepthStreams``: %v\n", err)
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
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | **string** | WebSocket stream update speed | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## IndexKlineCandlestickStreams

Index Kline/Candlestick Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "btcusdt" // string | The pair parameter
	interval := "1m" // string | The interval parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams().Pair(pair).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.IndexKlineCandlestickStreamsResponse) {
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
 **pair** | **string** | The pair parameter | 
 **interval** | **string** | The interval parameter | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## IndexPriceStream

Index Price Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "btcusdt" // string | The pair parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "updateSpeed_example" // string | WebSocket stream update speed (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexPriceStream().Pair(pair).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.IndexPriceStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.IndexPriceStreamResponse) {
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
 **pair** | **string** | The pair parameter | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | **string** | WebSocket stream update speed | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams``: %v\n", err)
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## IndividualSymbolMiniTickerStream

Individual Symbol Mini Ticker Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolMiniTickerStream().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.IndividualSymbolMiniTickerStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.IndividualSymbolMiniTickerStreamResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## IndividualSymbolTickerStreams

Individual Symbol Ticker Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolTickerStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.IndividualSymbolTickerStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.IndividualSymbolTickerStreamsResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## KlineCandlestickStreams

Kline/Candlestick Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	interval := "1m" // string | The interval parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.KlineCandlestickStreams().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.KlineCandlestickStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.KlineCandlestickStreamsResponse) {
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
 **interval** | **string** | The interval parameter | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## LiquidationOrderStreams

Liquidation Order Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.LiquidationOrderStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.LiquidationOrderStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.LiquidationOrderStreamsResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MarkPriceKlineCandlestickStreams

Mark Price Kline/Candlestick Streams


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	interval := "1m" // string | The interval parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.MarkPriceKlineCandlestickStreamsResponse) {
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
 **interval** | **string** | The interval parameter | 
 **id** | **string** | Unique WebSocket request ID. | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MarkPriceOfAllSymbolsOfAPair

Mark Price of All Symbols of a Pair


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	pair := "btcusdt" // string | The pair parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "updateSpeed_example" // string | WebSocket stream update speed (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceOfAllSymbolsOfAPair().Pair(pair).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.MarkPriceOfAllSymbolsOfAPair``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.MarkPriceOfAllSymbolsOfAPairResponse) {
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
 **pair** | **string** | The pair parameter | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | **string** | WebSocket stream update speed | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MarkPriceStream

Mark Price Stream


### Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "updateSpeed_example" // string | WebSocket stream update speed (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceStream().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.MarkPriceStream``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.MarkPriceStreamResponse) {
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
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | **string** | WebSocket stream update speed | 

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

	models "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	responseModels "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	symbol := "btcusdt" // string | The symbol parameter
	levels := int64(10) // int64 | The levels parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := "updateSpeed_example" // string | WebSocket stream update speed (optional)

	configuration := common.NewConfigurationWebsocketStreams(
		common.WithWsBasePath(common.SpotWebsocketStreamsProdUrl),
	)
	wsClient := models.NewBinanceDerivativesTradingCoinFuturesClient(models.WithWebsocketStreams(configuration))

	// Connect to WebSocket
	err := wsClient.WebsocketStreams.Connect()
	if err != nil {
		log.Printf("Error connecting to WebSocket: %v\n", err)
		return
	}

	handler, err := wsClient.WebsocketStreams.WebsocketMarketStreamsAPI.PartialBookDepthStreams().Symbol(symbol).Levels(levels).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `WebsocketMarketStreamsAPI.PartialBookDepthStreams``: %v\n", err)
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
 **levels** | **int64** | The levels parameter | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | **string** | WebSocket stream update speed | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

