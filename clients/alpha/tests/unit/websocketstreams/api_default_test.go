package binancealphawebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/clients/alpha/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancealphawebsocketstreams_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPI AggregateTradeStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"aggTrade","E":1771828569861,"T":1771828569702,"a":2684294,"f":2684294,"l":2684294,"m":false,"p":"0.08530000","q":"879.24000000","s":"ALPHA_474USDT"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggregateTradeStream().Symbol("alpha_116usdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AggregateTradeStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AggregateTradeStreamResponse
		mockCallback := func(msg models.AggregateTradeStreamResponse) {
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

	t.Run("Test DefaultAPIService AggregateTradeStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggregateTradeStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AggregateTradeStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AggregateTradeStream().Symbol("alpha_116usdt").Execute()
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
	t.Run("Test DefaultAPI AllBookTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"bookTicker","E":1773108379067,"T":1773108379054,"u":6663207693,"s":"ALPHA_116USDT","b":"8.30000000","B":"0.91076900","a":"9.30000000","A":"2.00353200"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllBookTickerStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllBookTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllBookTickerStreamResponse
		mockCallback := func(msg models.AllBookTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService AllBookTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllBookTickerStream().Execute()
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
	t.Run("Test DefaultAPI AllMiniTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrMiniTicker","E":1773109449908,"s":"ALPHA_116USDT","c":"8.40000000","o":"8.40000000","h":"8.50000000","l":"8.30000000","v":"64739.03978900","q":"543810.36005090"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMiniTickerStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMiniTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMiniTickerStreamResponse
		mockCallback := func(msg models.AllMiniTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService AllMiniTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllMiniTickerStream().Execute()
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
	t.Run("Test DefaultAPI AllTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrTicker","E":1773109631569,"s":"ALPHA_116USDT","p":"0.00000000","P":"0.00","w":"8.40003127","c":"8.40000000","Q":"0.49293200","o":"8.40000000","h":"8.50000000","l":"8.30000000","v":"64750.63418500","q":"543907.35169250","O":1773023220000,"C":1773109631555,"F":19847634,"L":19911287,"n":217505}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllTickerStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllTickerStreamResponse
		mockCallback := func(msg models.AllTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService AllTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllTickerStream().Execute()
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
	t.Run("Test DefaultAPI AllTokens24hTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"tickerList","d":[{"ca":"0x8fce...@56","cnt24":10426,"fdv":"72834145.18711072","hc":"6822","liq":"1442996.98151177662835","mc":"13542034.24008014","p":"0.072849536538593663","pc24":"-2.32","s":"1","t":1771825733000,"vol24":"12930712.213702157742688594285"}]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllTokens24hTickerStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllTokens24hTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllTokens24hTickerStreamResponse
		mockCallback := func(msg models.AllTokens24hTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService AllTokens24hTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.AllTokens24hTickerStream().Execute()
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
	t.Run("Test DefaultAPI BookTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"bookTicker","E":1773108379067,"T":1773108379054,"u":6663207693,"s":"ALPHA_116USDT","b":"8.30000000","B":"0.91076900","a":"9.30000000","A":"2.00353200"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTickerStream().Symbol("alpha_116usdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.BookTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.BookTickerStreamResponse
		mockCallback := func(msg models.BookTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService BookTickerStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTickerStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService BookTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BookTickerStream().Symbol("alpha_116usdt").Execute()
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
	t.Run("Test DefaultAPI ContractKlineStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"ca":"G7vQW...@CT_501","e":"kline","k":{"o":"0.15584372258086979622","c":"0.15584957424118055058","h":"0.15584957424118055058","l":"0.15584372258086979622","v":"2.8584092999332775480","ot":1771828655000,"ct":1771828656000,"i":"1s"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ContractKlineStream().ContractAddress("G7vQWurMkMMm2dU3iZpXYFTHT9Biio4F4gZCrwFpKNwG").ChainId("CT_501").Interval(models.ContractKlineStreamIntervalParameterInterval1s).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.ContractKlineStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.ContractKlineStreamResponse
		mockCallback := func(msg models.ContractKlineStreamResponse) {
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

	t.Run("Test DefaultAPIService ContractKlineStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ContractKlineStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ContractKlineStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ContractKlineStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ContractKlineStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ContractKlineStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ContractKlineStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ContractKlineStream().ContractAddress("G7vQWurMkMMm2dU3iZpXYFTHT9Biio4F4gZCrwFpKNwG").ChainId("CT_501").Interval(models.ContractKlineStreamIntervalParameterInterval1s).Execute()
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
	t.Run("Test DefaultAPI FullDepthStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1771828614381,"T":1771828614258,"U":42648947917,"u":42648948229,"pu":42648947321,"s":"ALPHA_474USDT","b":[["0.16479999","0"],["0.16480000","18.68000000"]],"a":[["0.16490000","500.00000000"],["0.16500000","0"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FullDepthStream().Symbol("alpha_116usdt").Interval(models.PartialDepthStreamIntervalParameterInterval0ms).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.FullDepthStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.FullDepthStreamResponse
		mockCallback := func(msg models.FullDepthStreamResponse) {
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

	t.Run("Test DefaultAPIService FullDepthStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FullDepthStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService FullDepthStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FullDepthStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService FullDepthStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FullDepthStream().Symbol("alpha_116usdt").Interval(models.PartialDepthStreamIntervalParameterInterval0ms).Execute()
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
	t.Run("Test DefaultAPI KlineStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"kline","E":1773111055144,"s":"ALPHA_116USDT","k":{"t":1773111000000,"T":1773111059999,"s":"ALPHA_116USDT","i":"1m","f":19912696,"L":19912765,"o":"8.40000000","c":"8.40000000","h":"8.40000000","l":"8.40000000","v":"63.64876300","n":218,"x":false,"q":"534.64960920","V":"63.64876300","Q":"534.64960920","B":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineStream().Symbol("alpha_116usdt").Interval(models.KlineStreamIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.KlineStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.KlineStreamResponse
		mockCallback := func(msg models.KlineStreamResponse) {
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

	t.Run("Test DefaultAPIService KlineStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService KlineStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService KlineStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.KlineStream().Symbol("alpha_116usdt").Interval(models.KlineStreamIntervalParameterInterval1m).Execute()
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
	t.Run("Test DefaultAPI MiniTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrMiniTicker","E":1773109449908,"s":"ALPHA_116USDT","c":"8.40000000","o":"8.40000000","h":"8.50000000","l":"8.30000000","v":"64739.03978900","q":"543810.36005090"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTickerStream().Symbol("alpha_116usdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MiniTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MiniTickerStreamResponse
		mockCallback := func(msg models.MiniTickerStreamResponse) {
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

	t.Run("Test DefaultAPIService MiniTickerStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTickerStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService MiniTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MiniTickerStream().Symbol("alpha_116usdt").Execute()
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
	t.Run("Test DefaultAPI PartialDepthStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1773110222878,"T":1773110222767,"U":6663367135,"u":6663367147,"pu":6663366971,"s":"ALPHA_116USDT","b":[["0.16479999","0"],["0.16480000","18.68000000"]],"a":[["0.16490000","500.00000000"],["0.16500000","0"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialDepthStream().Symbol("alpha_116usdt").Levels(models.PartialDepthStreamLevelsParameterLevels5).Interval(models.PartialDepthStreamIntervalParameterInterval0ms).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.PartialDepthStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.PartialDepthStreamResponse
		mockCallback := func(msg models.PartialDepthStreamResponse) {
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

	t.Run("Test DefaultAPIService PartialDepthStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialDepthStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PartialDepthStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialDepthStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PartialDepthStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialDepthStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PartialDepthStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PartialDepthStream().Symbol("alpha_116usdt").Levels(models.PartialDepthStreamLevelsParameterLevels5).Interval(models.PartialDepthStreamIntervalParameterInterval0ms).Execute()
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
	t.Run("Test DefaultAPI TickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrTicker","E":1773109631569,"s":"ALPHA_116USDT","p":"0.00000000","P":"0.00","w":"8.40003127","c":"8.40000000","Q":"0.49293200","o":"8.40000000","h":"8.50000000","l":"8.30000000","v":"64750.63418500","q":"543907.35169250","O":1773023220000,"C":1773109631555,"F":19847634,"L":19911287,"n":217505}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TickerStream().Symbol("alpha_116usdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.TickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.TickerStreamResponse
		mockCallback := func(msg models.TickerStreamResponse) {
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

	t.Run("Test DefaultAPIService TickerStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TickerStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService TickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TickerStream().Symbol("alpha_116usdt").Execute()
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
	t.Run("Test DefaultAPI TradeStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"trade","E":1773110023891,"T":1773110023877,"s":"ALPHA_116USDT","t":19911650,"p":"8.40000000","q":"0.97915700","m":false}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TradeStream().Symbol("alpha_116usdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.TradeStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.TradeStreamResponse
		mockCallback := func(msg models.TradeStreamResponse) {
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

	t.Run("Test DefaultAPIService TradeStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TradeStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService TradeStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceAlphaClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.TradeStream().Symbol("alpha_116usdt").Execute()
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
