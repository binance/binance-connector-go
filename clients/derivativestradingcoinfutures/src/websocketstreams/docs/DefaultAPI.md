# \DefaultAPI

All URIs are relative to *http://localhost*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**AggregateTradeStreams**](DefaultAPI.md#AggregateTradeStreams) | /&lt;symbol&gt;@aggTrade | Aggregate Trade Streams
[**AllBookTickersStream**](DefaultAPI.md#AllBookTickersStream) | /!bookTicker | All Book Tickers Stream
[**AllMarketLiquidationOrderStreams**](DefaultAPI.md#AllMarketLiquidationOrderStreams) | /!forceOrder@arr | All Market Liquidation Order Streams
[**AllMarketMiniTickersStream**](DefaultAPI.md#AllMarketMiniTickersStream) | /!miniTicker@arr | All Market Mini Tickers Stream
[**AllMarketTickersStreams**](DefaultAPI.md#AllMarketTickersStreams) | /!ticker@arr | All Market Tickers Streams
[**ContinuousContractKlineCandlestickStreams**](DefaultAPI.md#ContinuousContractKlineCandlestickStreams) | /&lt;pair&gt;_&lt;contractType&gt;@continuousKline_&lt;interval&gt; | Continuous Contract Kline/Candlestick Streams
[**ContractInfoStream**](DefaultAPI.md#ContractInfoStream) | /!contractInfo | Contract Info Stream
[**DiffBookDepthStreams**](DefaultAPI.md#DiffBookDepthStreams) | /&lt;symbol&gt;@depth@&lt;updateSpeed&gt; | Diff. Book Depth Streams
[**IndexKlineCandlestickStreams**](DefaultAPI.md#IndexKlineCandlestickStreams) | /&lt;pair&gt;@indexPriceKline_&lt;interval&gt; | Index Kline/Candlestick Streams
[**IndexPriceStream**](DefaultAPI.md#IndexPriceStream) | /&lt;pair&gt;@indexPrice@&lt;updateSpeed&gt; | Index Price Stream
[**IndividualSymbolBookTickerStreams**](DefaultAPI.md#IndividualSymbolBookTickerStreams) | /&lt;symbol&gt;@bookTicker | Individual Symbol Book Ticker Streams
[**IndividualSymbolMiniTickerStream**](DefaultAPI.md#IndividualSymbolMiniTickerStream) | /&lt;symbol&gt;@miniTicker | Individual Symbol Mini Ticker Stream
[**IndividualSymbolTickerStreams**](DefaultAPI.md#IndividualSymbolTickerStreams) | /&lt;symbol&gt;@ticker | Individual Symbol Ticker Streams
[**KlineCandlestickStreams**](DefaultAPI.md#KlineCandlestickStreams) | /&lt;symbol&gt;@kline_&lt;interval&gt; | Kline/Candlestick Streams
[**MarkPriceKlineCandlestickStreams**](DefaultAPI.md#MarkPriceKlineCandlestickStreams) | /&lt;symbol&gt;@markPriceKline_&lt;interval&gt; | Mark Price Kline/Candlestick Streams
[**MarkPriceOfAllSymbolsOfAPair**](DefaultAPI.md#MarkPriceOfAllSymbolsOfAPair) | /&lt;pair&gt;@markPrice@&lt;updateSpeed&gt; | Mark Price of All Symbols of a Pair
[**MarkPriceStream**](DefaultAPI.md#MarkPriceStream) | /&lt;symbol&gt;@markPrice@&lt;updateSpeed&gt; | Mark Price Stream
[**MarketLiquidationOrderStreams**](DefaultAPI.md#MarketLiquidationOrderStreams) | /&lt;symbol&gt;@forceOrder | Market Liquidation Order Streams
[**PartialBookDepthStreams**](DefaultAPI.md#PartialBookDepthStreams) | /&lt;symbol&gt;@depth&lt;levels&gt;@&lt;updateSpeed&gt; | Partial Book Depth Streams


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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AggregateTradeStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AggregateTradeStreams``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllBookTickersStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllBookTickersStream``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllMarketLiquidationOrderStreams().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllMarketLiquidationOrderStreams``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllMarketMiniTickersStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllMarketMiniTickersStream``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.AllMarketTickersStreams().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.AllMarketTickersStreams``: %v\n", err)
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
	contractType := models.ContinuousContractKlineCandlestickStreamsContractTypeParameterPerpetual // ContinuousContractKlineCandlestickStreamsContractTypeParameter | The contractType parameter
	interval := models.ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1m // ContinuousContractKlineCandlestickStreamsIntervalParameter | The interval parameter
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ContinuousContractKlineCandlestickStreams().Pair(pair).ContractType(contractType).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.ContinuousContractKlineCandlestickStreams``: %v\n", err)
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
 **contractType** | [**ContinuousContractKlineCandlestickStreamsContractTypeParameter**](ContinuousContractKlineCandlestickStreamsContractTypeParameter.md) | The contractType parameter | 
 **interval** | [**ContinuousContractKlineCandlestickStreamsIntervalParameter**](ContinuousContractKlineCandlestickStreamsIntervalParameter.md) | The interval parameter | 
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.ContractInfoStream().Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.ContractInfoStream``: %v\n", err)
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
	updateSpeed := models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthStreamsUpdateSpeedParameter | WebSocket stream update speed (optional)

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.DiffBookDepthStreams().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.DiffBookDepthStreams``: %v\n", err)
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
 **updateSpeed** | [**DiffBookDepthStreamsUpdateSpeedParameter**](DiffBookDepthStreamsUpdateSpeedParameter.md) | WebSocket stream update speed | 

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
	interval := models.ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1m // ContinuousContractKlineCandlestickStreamsIntervalParameter | The interval parameter
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.IndexKlineCandlestickStreams().Pair(pair).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.IndexKlineCandlestickStreams``: %v\n", err)
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
 **interval** | [**ContinuousContractKlineCandlestickStreamsIntervalParameter**](ContinuousContractKlineCandlestickStreamsIntervalParameter.md) | The interval parameter | 
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
	updateSpeed := models.IndexPriceStreamUpdateSpeedParameterUpdateSpeed1s // IndexPriceStreamUpdateSpeedParameter | WebSocket stream update speed (optional)

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.IndexPriceStream().Pair(pair).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.IndexPriceStream``: %v\n", err)
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
 **updateSpeed** | [**IndexPriceStreamUpdateSpeedParameter**](IndexPriceStreamUpdateSpeedParameter.md) | WebSocket stream update speed | 

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.IndividualSymbolBookTickerStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.IndividualSymbolBookTickerStreams``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.IndividualSymbolMiniTickerStream().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.IndividualSymbolMiniTickerStream``: %v\n", err)
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.IndividualSymbolTickerStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.IndividualSymbolTickerStreams``: %v\n", err)
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
	interval := models.ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1m // ContinuousContractKlineCandlestickStreamsIntervalParameter | The interval parameter
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.KlineCandlestickStreams().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.KlineCandlestickStreams``: %v\n", err)
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
 **interval** | [**ContinuousContractKlineCandlestickStreamsIntervalParameter**](ContinuousContractKlineCandlestickStreamsIntervalParameter.md) | The interval parameter | 
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
	interval := models.ContinuousContractKlineCandlestickStreamsIntervalParameterInterval1m // ContinuousContractKlineCandlestickStreamsIntervalParameter | The interval parameter
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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.MarkPriceKlineCandlestickStreams().Symbol(symbol).Interval(interval).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MarkPriceKlineCandlestickStreams``: %v\n", err)
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
 **interval** | [**ContinuousContractKlineCandlestickStreamsIntervalParameter**](ContinuousContractKlineCandlestickStreamsIntervalParameter.md) | The interval parameter | 
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
	updateSpeed := models.IndexPriceStreamUpdateSpeedParameterUpdateSpeed1s // IndexPriceStreamUpdateSpeedParameter | WebSocket stream update speed (optional)

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.MarkPriceOfAllSymbolsOfAPair().Pair(pair).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MarkPriceOfAllSymbolsOfAPair``: %v\n", err)
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
 **updateSpeed** | [**IndexPriceStreamUpdateSpeedParameter**](IndexPriceStreamUpdateSpeedParameter.md) | WebSocket stream update speed | 

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
	updateSpeed := models.IndexPriceStreamUpdateSpeedParameterUpdateSpeed1s // IndexPriceStreamUpdateSpeedParameter | WebSocket stream update speed (optional)

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.MarkPriceStream().Symbol(symbol).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MarkPriceStream``: %v\n", err)
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
 **updateSpeed** | [**IndexPriceStreamUpdateSpeedParameter**](IndexPriceStreamUpdateSpeedParameter.md) | WebSocket stream update speed | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)


## MarketLiquidationOrderStreams

Market Liquidation Order Streams


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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.MarketLiquidationOrderStreams().Symbol(symbol).Id(id).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.MarketLiquidationOrderStreams``: %v\n", err)
		return
	}

	handler.On("message", func(message responseModels.MarketLiquidationOrderStreamsResponse) {
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
	levels := models.PartialBookDepthStreamsLevelsParameterLevels5 // PartialBookDepthStreamsLevelsParameter | The levels parameter
	id := "e9d6b4349871b40611412680b3445fac" // string | Unique WebSocket request ID. (optional)
	updateSpeed := models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms // DiffBookDepthStreamsUpdateSpeedParameter | WebSocket stream update speed (optional)

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

	handler, err := wsClient.WebsocketStreams.DefaultAPI.PartialBookDepthStreams().Symbol(symbol).Levels(levels).Id(id).UpdateSpeed(updateSpeed).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `DefaultAPI.PartialBookDepthStreams``: %v\n", err)
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
 **levels** | [**PartialBookDepthStreamsLevelsParameter**](PartialBookDepthStreamsLevelsParameter.md) | The levels parameter | 
 **id** | **string** | Unique WebSocket request ID. | 
 **updateSpeed** | [**DiffBookDepthStreamsUpdateSpeedParameter**](DiffBookDepthStreamsUpdateSpeedParameter.md) | WebSocket stream update speed | 

### Authorization

No authorization required

[[Back to README]](../../../README.md)

