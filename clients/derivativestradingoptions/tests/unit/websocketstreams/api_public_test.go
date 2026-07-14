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

func Test_binancederivativestradingoptionswebsocketstreams_PublicAPIService(t *testing.T) {

	t.Run("Test PublicAPI DiffBookDepthStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1762866729459,"T":1762866729358,"s":"BTC-251123-126000-C","U":465,"u":465,"pu":464,"b":[["1100.000"]],"a":[["1300.000"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.PublicAPI.DiffBookDepthStreams().Symbol("btcusdt").UpdateSpeed(models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.DiffBookDepthStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.DiffBookDepthStreamsResponse
		mockCallback := func(msg models.DiffBookDepthStreamsResponse) {
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

	t.Run("Test PublicAPIService DiffBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.DiffBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService DiffBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.DiffBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService DiffBookDepthStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.DiffBookDepthStreams().Symbol("btcusdt").UpdateSpeed(models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms).Execute()
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
	t.Run("Test PublicAPI Hour24Ticker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrTicker","E":1764080707933,"s":"ETH-251226-3000-C","p":"0.0000","P":"0.00","w":"200.0000","c":"200.0000","Q":"1.0000","o":"200.0000","h":"200.0000","l":"200.0000","v":"9.0000","q":"1800.0000","O":1764051060000,"C":1764080707933,"F":1,"L":22,"n":9}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.PublicAPI.Hour24Ticker().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.Hour24TickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.Hour24TickerResponse
		mockCallback := func(msg models.Hour24TickerResponse) {
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

	t.Run("Test PublicAPIService Hour24Ticker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.Hour24Ticker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService Hour24Ticker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.Hour24Ticker().Symbol("btcusdt").Execute()
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
	t.Run("Test PublicAPI IndividualSymbolBookTickerStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"bookTicker","u":2472,"s":"BTC-251226-110000-C","b":"5000.000","B":"0.2000","a":"5100.000","A":"0.1000","T":1763041762942,"E":1763041762942}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.PublicAPI.IndividualSymbolBookTickerStreams().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndividualSymbolBookTickerStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndividualSymbolBookTickerStreamsResponse
		mockCallback := func(msg models.IndividualSymbolBookTickerStreamsResponse) {
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

	t.Run("Test PublicAPIService IndividualSymbolBookTickerStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.IndividualSymbolBookTickerStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService IndividualSymbolBookTickerStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.IndividualSymbolBookTickerStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test PublicAPI PartialBookDepthStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1762866729459,"T":1762866729358,"s":"BTC-251123-126000-C","U":465,"u":465,"pu":464,"b":[["1100.000"]],"a":[["1300.000"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Symbol("btcusdt").Level(models.PartialBookDepthStreamsLevelParameterLevel5).UpdateSpeed(models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.PartialBookDepthStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.PartialBookDepthStreamsResponse
		mockCallback := func(msg models.PartialBookDepthStreamsResponse) {
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

	t.Run("Test PublicAPIService PartialBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService PartialBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService PartialBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService PartialBookDepthStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.PartialBookDepthStreams().Symbol("btcusdt").Level(models.PartialBookDepthStreamsLevelParameterLevel5).UpdateSpeed(models.DiffBookDepthStreamsUpdateSpeedParameterUpdateSpeed100ms).Execute()
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
	t.Run("Test PublicAPI TradeStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"trade","E":1762856064204,"T":1762856064203,"s":"BTC-251123-126000-C","t":4,"p":"1300.000","q":"0.1000","X":"MARKET","S":"BUY","m":false}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.PublicAPI.TradeStreams().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.TradeStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.TradeStreamsResponse
		mockCallback := func(msg models.TradeStreamsResponse) {
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

	t.Run("Test PublicAPIService TradeStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.TradeStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PublicAPIService TradeStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.WsPublic.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.PublicAPI.TradeStreams().Symbol("btcusdt").Execute()
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
