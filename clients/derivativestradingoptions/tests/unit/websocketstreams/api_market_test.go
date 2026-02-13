package binancederivativestradingoptionswebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingoptionswebsocketstreams_MarketAPIService(t *testing.T) {

	t.Run("Test MarketAPI IndexPriceStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"indexPrice","E":1763092572229,"s":"ETHUSDT","p":"3224.51976744"},{"e":"indexPrice","E":1763092572229,"s":"BTCUSDT","p":"99102.32326087"}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.MarketAPI.IndexPriceStreams().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndexPriceStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndexPriceStreamsResponse
		mockCallback := func(msg models.IndexPriceStreamsResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test MarketAPIService IndexPriceStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.IndexPriceStreams().Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
	t.Run("Test MarketAPI KlineCandlestickStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"kline","E":1638747660000,"s":"BTC-200630-9000-P","k":{"t":1638747660000,"T":1638747719999,"s":"BTC-200630-9000-P","i":"1m","f":0,"L":0,"o":"1000","c":"1000","h":"1000","l":"1000","v":"0","n":0,"x":false,"q":"0","V":"0","Q":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.MarketAPI.KlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.KlineCandlestickStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.KlineCandlestickStreamsResponse
		mockCallback := func(msg models.KlineCandlestickStreamsResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test MarketAPIService KlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.KlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketAPIService KlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.KlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketAPIService KlineCandlestickStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.KlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
	t.Run("Test MarketAPI MarkPrice Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"s":"BTC-251120-126000-C","mp":"770.543","E":1762867543321,"e":"markPrice","i":"104334.60217391","P":"0.000","bo":"0.000","ao":"900.000","bq":"0.0000","aq":"0.2000","b":"-1.0","a":"0.98161161","hl":"924.652","ll":"616.435","vo":"0.9408058","rf":"0.0","d":"0.11111964","t":"-164.26702615","g":"0.00001245","v":"30.63855919"},{"s":"BTC-251123-126000-C","mp":"1249.61","E":1762867543321,"e":"markPrice","i":"104334.60217391","P":"0.000","bo":"1200.000","ao":"1300.000","bq":"0.3000","aq":"0.6000","b":"0.92159033","a":"0.94461441","hl":"1499.533","ll":"999.688","vo":"0.93310237","rf":"0.0","d":"0.14869196","t":"-172.12148811","g":"0.00001326","v":"43.43627792"}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.MarketAPI.MarkPrice().Underlying("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarkPriceResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarkPriceResponse
		mockCallback := func(msg models.MarkPriceResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test MarketAPIService MarkPrice Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.MarkPrice().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketAPIService MarkPrice Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.MarkPrice().Underlying("btcusdt").Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
	t.Run("Test MarketAPI NewSymbolInfo Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"optionSymbol","E":1669356423908,"s":"BTC-250926-140000-C","ps":"BTCUSDT","qa":"USDT","d":"CALL","sp":"21000","dt":4133404800000,"u":1,"ot":1569398400000,"cs":"TRADING"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.MarketAPI.NewSymbolInfo().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.NewSymbolInfoResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.NewSymbolInfoResponse
		mockCallback := func(msg models.NewSymbolInfoResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test MarketAPIService NewSymbolInfo Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.NewSymbolInfo().Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
	t.Run("Test MarketAPI OpenInterest Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"openInterest","E":1668759300045,"s":"ETH-221125-2700-C","o":"1580.87","h":"1912992.178168204"}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.MarketAPI.OpenInterest().ExpirationDate("220930").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.OpenInterestResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.OpenInterestResponse
		mockCallback := func(msg models.OpenInterestResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test MarketAPIService OpenInterest Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.OpenInterest().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketAPIService OpenInterest Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsMarket.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.MarketAPI.OpenInterest().ExpirationDate("220930").Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
}
