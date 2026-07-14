package binancespotwebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancespotwebsocketstreams_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPI AggTrade Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"aggTrade","E":1672515782136,"s":"BNBBTC","a":12345,"f":100,"l":105,"T":1672515782136,"m":true,"M":true}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggTrade().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AggTradeResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AggTradeResponse
		mockCallback := func(msg models.AggTradeResponse) {
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

	t.Run("Test DefaultAPIService AggTrade Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggTrade().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AggTrade Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggTrade().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI AllMarketRollingWindowTicker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"1hTicker","E":1672515782136,"s":"BNBBTC","O":0,"C":1675216573749,"F":0,"L":18150,"n":18151}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMarketRollingWindowTicker().WindowSize(models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMarketRollingWindowTickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMarketRollingWindowTickerResponse
		mockCallback := func(msg models.AllMarketRollingWindowTickerResponse) {
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

	t.Run("Test DefaultAPIService AllMarketRollingWindowTicker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMarketRollingWindowTicker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AllMarketRollingWindowTicker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMarketRollingWindowTicker().WindowSize(models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h).Execute()
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
	t.Run("Test DefaultAPI AllMiniTicker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"24hrMiniTicker","E":1672515782136,"s":"BNBBTC"}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMiniTicker().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMiniTickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMiniTickerResponse
		mockCallback := func(msg models.AllMiniTickerResponse) {
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

	t.Run("Test DefaultAPIService AllMiniTicker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMiniTicker().Execute()
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
	t.Run("Test DefaultAPI AvgPrice Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"avgPrice","E":1693907033000,"s":"BTCUSDT","i":"5m","T":1693907032213}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AvgPrice().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AvgPriceResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AvgPriceResponse
		mockCallback := func(msg models.AvgPriceResponse) {
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

	t.Run("Test DefaultAPIService AvgPrice Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AvgPrice().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AvgPrice Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AvgPrice().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI BlockTrade Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"blockTrade","E":1772506983582,"s":"BNBBTC","t":582,"p":"0.052","q":"5838","T":1772506983321,"m":true}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BlockTrade().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.BlockTradeResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.BlockTradeResponse
		mockCallback := func(msg models.BlockTradeResponse) {
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

	t.Run("Test DefaultAPIService BlockTrade Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BlockTrade().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService BlockTrade Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BlockTrade().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI BookTicker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"u":400900217,"s":"BNBUSDT"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTicker().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.BookTickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.BookTickerResponse
		mockCallback := func(msg models.BookTickerResponse) {
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

	t.Run("Test DefaultAPIService BookTicker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTicker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService BookTicker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTicker().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI DiffBookDepth Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1672515782136,"s":"BNBBTC","U":157,"u":160,"b":[["0.0024","10"]],"a":[["0.0026","100"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.DiffBookDepth().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.DiffBookDepthResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.DiffBookDepthResponse
		mockCallback := func(msg models.DiffBookDepthResponse) {
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

	t.Run("Test DefaultAPIService DiffBookDepth Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.DiffBookDepth().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService DiffBookDepth Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.DiffBookDepth().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI Kline Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"kline","E":1672515782136,"s":"BNBBTC","k":{"t":1672515780000,"T":1672515839999,"s":"BNBBTC","i":"1m","f":100,"L":200,"n":100,"x":false}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Kline().Symbol("bnbusdt").Interval(models.KlineIntervalParameterInterval1s).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.KlineResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.KlineResponse
		mockCallback := func(msg models.KlineResponse) {
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

	t.Run("Test DefaultAPIService Kline Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Kline().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService Kline Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Kline().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService Kline Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Kline().Symbol("bnbusdt").Interval(models.KlineIntervalParameterInterval1s).Execute()
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
	t.Run("Test DefaultAPI KlineOffset Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"kline","E":1672515782136,"s":"BNBBTC","k":{"t":1672515780000,"T":1672515839999,"s":"BNBBTC","i":"1m","f":100,"L":200,"n":100,"x":false}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineOffset().Symbol("bnbusdt").Interval(models.KlineIntervalParameterInterval1s).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.KlineOffsetResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.KlineOffsetResponse
		mockCallback := func(msg models.KlineOffsetResponse) {
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

	t.Run("Test DefaultAPIService KlineOffset Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineOffset().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService KlineOffset Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineOffset().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService KlineOffset Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineOffset().Symbol("bnbusdt").Interval(models.KlineIntervalParameterInterval1s).Execute()
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
	t.Run("Test DefaultAPI MiniTicker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrMiniTicker","E":1672515782136,"s":"BNBBTC"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTicker().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MiniTickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MiniTickerResponse
		mockCallback := func(msg models.MiniTickerResponse) {
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

	t.Run("Test DefaultAPIService MiniTicker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTicker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService MiniTicker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTicker().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI PartialBookDepth Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"lastUpdateId":160,"bids":[["0.0024","10"]],"asks":[["0.0026","100"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialBookDepth().Symbol("bnbusdt").Levels(models.PartialBookDepthLevelsParameterLevels5).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.PartialBookDepthResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.PartialBookDepthResponse
		mockCallback := func(msg models.PartialBookDepthResponse) {
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

	t.Run("Test DefaultAPIService PartialBookDepth Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialBookDepth().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PartialBookDepth Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialBookDepth().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PartialBookDepth Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialBookDepth().Symbol("bnbusdt").Levels(models.PartialBookDepthLevelsParameterLevels5).Execute()
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
	t.Run("Test DefaultAPI ReferencePrice Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"referencePrice","s":"BAZUSD","r":"1.00","t":1770313263917}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ReferencePrice().Symbol("bazusd").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.ReferencePriceResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.ReferencePriceResponse
		mockCallback := func(msg models.ReferencePriceResponse) {
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

	t.Run("Test DefaultAPIService ReferencePrice Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ReferencePrice().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ReferencePrice Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ReferencePrice().Symbol("bazusd").Execute()
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
	t.Run("Test DefaultAPI RollingWindowTicker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"1hTicker","E":1672515782136,"s":"BNBBTC","O":0,"C":1675216573749,"F":0,"L":18150,"n":18151}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RollingWindowTicker().Symbol("bnbusdt").WindowSize(models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.RollingWindowTickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.RollingWindowTickerResponse
		mockCallback := func(msg models.RollingWindowTickerResponse) {
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

	t.Run("Test DefaultAPIService RollingWindowTicker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RollingWindowTicker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RollingWindowTicker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RollingWindowTicker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RollingWindowTicker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RollingWindowTicker().Symbol("bnbusdt").WindowSize(models.AllMarketRollingWindowTickerWindowSizeParameterWindowSize1h).Execute()
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
	t.Run("Test DefaultAPI Ticker Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrTicker","E":1672515782136,"s":"BNBBTC","O":0,"C":1675216573749,"F":0,"L":18150,"n":18151}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Ticker().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.TickerResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.TickerResponse
		mockCallback := func(msg models.TickerResponse) {
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

	t.Run("Test DefaultAPIService Ticker Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Ticker().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService Ticker Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Ticker().Symbol("bnbusdt").Execute()
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
	t.Run("Test DefaultAPI Trade Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"trade","E":1672515782136,"s":"BNBBTC","t":12345,"T":1672515782136,"m":true,"M":true}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Trade().Symbol("bnbusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.TradeResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.TradeResponse
		mockCallback := func(msg models.TradeResponse) {
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

	t.Run("Test DefaultAPIService Trade Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Trade().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService Trade Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.Trade().Symbol("bnbusdt").Execute()
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
