package binancederivativestradingcoinfutureswebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/common"
	tests "github.com/binance/binance-connector-go/common/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingcoinfutureswebsocketstreams_WebsocketMarketStreamsAPIService(t *testing.T) {

	t.Run("Test WebsocketMarketStreamsAPI AggregateTradeStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"aggTrade","E":1591261134288,"a":424951,"s":"BTCUSD_200626","p":"9643.5","q":"2","f":606073,"l":606073,"T":1591261134199,"m":false}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AggregateTradeStreams().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AggregateTradeStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AggregateTradeStreamsResponse
		mockCallback := func(msg models.AggregateTradeStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService AggregateTradeStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AggregateTradeStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService AggregateTradeStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AggregateTradeStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI AllBookTickersStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"bookTicker","u":17242169,"s":"BTCUSD_200626","ps":"BTCUSD","b":"9548.1","B":"52","a":"9548.5","A":"11","T":1591268628155,"E":1591268628166}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllBookTickersStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllBookTickersStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllBookTickersStreamResponse
		mockCallback := func(msg models.AllBookTickersStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService AllBookTickersStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllBookTickersStream().Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI AllMarketLiquidationOrderStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"forceOrder","E":1591154240950,"o":{"s":"BTCUSD_200925","ps":"BTCUSD","S":"SELL","o":"LIMIT","f":"IOC","q":"1","p":"9425.5","ap":"9496.5","X":"FILLED","l":"1","z":"1","T":1591154240949}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketLiquidationOrderStreams().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMarketLiquidationOrderStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMarketLiquidationOrderStreamsResponse
		mockCallback := func(msg models.AllMarketLiquidationOrderStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService AllMarketLiquidationOrderStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketLiquidationOrderStreams().Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI AllMarketMiniTickersStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"24hrMiniTicker","E":1591267704450,"s":"BTCUSD_200626","ps":"BTCUSD","c":"9561.7","o":"9580.9","h":"10000.0","l":"7000.0","v":"487476","q":"33264343847.22378500"}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketMiniTickersStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMarketMiniTickersStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMarketMiniTickersStreamResponse
		mockCallback := func(msg models.AllMarketMiniTickersStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService AllMarketMiniTickersStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketMiniTickersStream().Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI AllMarketTickersStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"24hrTicker","E":1591268262453,"s":"BTCUSD_200626","ps":"BTCUSD","p":"-43.4","P":"-0.452","w":"0.00147974","c":"9548.5","Q":"2","o":"9591.9","h":"10000.0","l":"7000.0","v":"487850","q":"32968676323.46222700","O":1591181820000,"C":1591268262442,"F":512014,"L":615289,"n":103272}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketTickersStreams().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.AllMarketTickersStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.AllMarketTickersStreamsResponse
		mockCallback := func(msg models.AllMarketTickersStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService AllMarketTickersStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.AllMarketTickersStreams().Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI ContinuousContractKlineCandlestickStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"continuous_kline","E":1591261542539,"ps":"BTCUSD","ct":"NEXT_QUARTER","k":{"t":1591261500000,"T":1591261559999,"i":"1m","f":606400,"L":606430,"o":"9638.9","c":"9639.8","h":"9639.8","l":"9638.6","v":"156","n":31,"x":false,"q":"1.61836886","V":"73","Q":"0.75731156","B":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Pair("btcusdt").ContractType("next_quarter").Interval("1m").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.ContinuousContractKlineCandlestickStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.ContinuousContractKlineCandlestickStreamsResponse
		mockCallback := func(msg models.ContinuousContractKlineCandlestickStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService ContinuousContractKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService ContinuousContractKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService ContinuousContractKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService ContinuousContractKlineCandlestickStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContinuousContractKlineCandlestickStreams().Pair("btcusdt").ContractType("next_quarter").Interval("1m").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI ContractInfoStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"contractInfo","E":1669647330375,"s":"APTUSD_PERP","ps":"APTUSD","ct":"PERPETUAL","dt":4133404800000,"ot":1666594800000,"cs":"TRADING","bks":[{"bs":1,"bnf":0,"bnc":5000,"mmr":0.01,"cf":0,"mi":21,"ma":50},{"bs":2,"bnf":5000,"bnc":25000,"mmr":0.025,"cf":75,"mi":11,"ma":20}]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContractInfoStream().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.ContractInfoStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.ContractInfoStreamResponse
		mockCallback := func(msg models.ContractInfoStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService ContractInfoStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.ContractInfoStream().Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI DiffBookDepthStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1591270260907,"T":1591270260891,"s":"BTCUSD_200626","ps":"BTCUSD","U":17285681,"u":17285702,"pu":17285675,"b":[["9517.6","10"]],"a":[["9518.5","45"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.DiffBookDepthStreams().Symbol("btcusdt").Execute()
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

	t.Run("Test WebsocketMarketStreamsAPIService DiffBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.DiffBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService DiffBookDepthStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.DiffBookDepthStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI IndexKlineCandlestickStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"indexPrice_kline","E":1591267070033,"ps":"BTCUSD","k":{"t":1591267020000,"T":1591267079999,"s":"0","i":"1m","f":1591267020000,"L":1591267070000,"o":"9542.21900000","c":"9542.50440000","h":"9542.71640000","l":"9542.21040000","v":"0","n":51,"x":false,"q":"0","V":"0","Q":"0","B":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams().Pair("btcusdt").Interval("1m").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndexKlineCandlestickStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndexKlineCandlestickStreamsResponse
		mockCallback := func(msg models.IndexKlineCandlestickStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService IndexKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndexKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndexKlineCandlestickStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexKlineCandlestickStreams().Pair("btcusdt").Interval("1m").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI IndexPriceStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"indexPriceUpdate","E":1591261236000,"i":"BTCUSD","p":"9636.57860000"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexPriceStream().Pair("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndexPriceStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndexPriceStreamResponse
		mockCallback := func(msg models.IndexPriceStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService IndexPriceStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexPriceStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndexPriceStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndexPriceStream().Pair("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI IndividualSymbolBookTickerStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"bookTicker","u":17242169,"s":"BTCUSD_200626","ps":"BTCUSD","b":"9548.1","B":"52","a":"9548.5","A":"11","T":1591268628155,"E":1591268628166}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams().Symbol("btcusdt").Execute()
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

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolBookTickerStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolBookTickerStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolBookTickerStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI IndividualSymbolMiniTickerStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrMiniTicker","E":1591267704450,"s":"BTCUSD_200626","ps":"BTCUSD","c":"9561.7","o":"9580.9","h":"10000.0","l":"7000.0","v":"487476","q":"33264343847.22378500"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolMiniTickerStream().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndividualSymbolMiniTickerStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndividualSymbolMiniTickerStreamResponse
		mockCallback := func(msg models.IndividualSymbolMiniTickerStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolMiniTickerStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolMiniTickerStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolMiniTickerStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolMiniTickerStream().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI IndividualSymbolTickerStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"24hrTicker","E":1591268262453,"s":"BTCUSD_200626","ps":"BTCUSD","p":"-43.4","P":"-0.452","w":"0.00147974","c":"9548.5","Q":"2","o":"9591.9","h":"10000.0","l":"7000.0","v":"487850","q":"32968676323.46222700","O":1591181820000,"C":1591268262442,"F":512014,"L":615289,"n":103272}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolTickerStreams().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.IndividualSymbolTickerStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.IndividualSymbolTickerStreamsResponse
		mockCallback := func(msg models.IndividualSymbolTickerStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolTickerStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolTickerStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService IndividualSymbolTickerStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.IndividualSymbolTickerStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI KlineCandlestickStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"kline","E":1591261542539,"s":"BTCUSD_200626","k":{"t":1591261500000,"T":1591261559999,"s":"BTCUSD_200626","i":"1m","f":606400,"L":606430,"o":"9638.9","c":"9639.8","h":"9639.8","l":"9638.6","v":"156","n":31,"x":false,"q":"1.61836886","V":"73","Q":"0.75731156","B":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.KlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
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

	t.Run("Test WebsocketMarketStreamsAPIService KlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.KlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService KlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.KlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService KlineCandlestickStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.KlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI LiquidationOrderStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"forceOrder","E":1591154240950,"o":{"s":"BTCUSD_200925","ps":"BTCUSD","S":"SELL","o":"LIMIT","f":"IOC","q":"1","p":"9425.5","ap":"9496.5","X":"FILLED","l":"1","z":"1","T":1591154240949}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.LiquidationOrderStreams().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.LiquidationOrderStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.LiquidationOrderStreamsResponse
		mockCallback := func(msg models.LiquidationOrderStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService LiquidationOrderStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.LiquidationOrderStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService LiquidationOrderStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.LiquidationOrderStreams().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI MarkPriceKlineCandlestickStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"markPrice_kline","E":1591267398004,"ps":"BTCUSD","k":{"t":1591267380000,"T":1591267439999,"s":"BTCUSD_200626","i":"1m","f":1591267380000,"L":1591267398000,"o":"9539.67161333","c":"9540.82761333","h":"9540.82761333","l":"9539.66961333","v":"0","n":19,"x":false,"q":"0","V":"0","Q":"0","B":"0"}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarkPriceKlineCandlestickStreamsResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarkPriceKlineCandlestickStreamsResponse
		mockCallback := func(msg models.MarkPriceKlineCandlestickStreamsResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceKlineCandlestickStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceKlineCandlestickStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceKlineCandlestickStreams().Symbol("btcusdt").Interval("1m").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI MarkPriceOfAllSymbolsOfAPair Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `[{"e":"markPriceUpdate","E":1596095725000,"s":"BTCUSD_201225","p":"10934.62615417","P":"10962.17178236","i":"10933.62615417","r":"","T":0},{"e":"markPriceUpdate","E":1596095725000,"s":"BTCUSD_PERP","p":"11012.31359011","P":"10962.17178236","i":"10933.62615417","r":"0.00000000","T":1596096000000}]`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceOfAllSymbolsOfAPair().Pair("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarkPriceOfAllSymbolsOfAPairResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarkPriceOfAllSymbolsOfAPairResponse
		mockCallback := func(msg models.MarkPriceOfAllSymbolsOfAPairResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceOfAllSymbolsOfAPair Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceOfAllSymbolsOfAPair().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceOfAllSymbolsOfAPair Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceOfAllSymbolsOfAPair().Pair("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI MarkPriceStream Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"markPriceUpdate","E":1596095725000,"s":"BTCUSD_201225","p":"10934.62615417","P":"10962.17178236","i":"10933.62615417","r":"","T":0}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceStream().Symbol("btcusdt").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarkPriceStreamResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarkPriceStreamResponse
		mockCallback := func(msg models.MarkPriceStreamResponse) {
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

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceStream Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceStream().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService MarkPriceStream Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.MarkPriceStream().Symbol("btcusdt").Execute()
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
	t.Run("Test WebsocketMarketStreamsAPI PartialBookDepthStreams Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"e":"depthUpdate","E":1591269996801,"T":1591269996646,"s":"BTCUSD_200626","ps":"BTCUSD","U":17276694,"u":17276701,"pu":17276678,"b":[["9523.0","5"],["9522.8","8"],["9522.6","2"],["9522.4","1"],["9522.0","5"]],"a":[["9524.6","2"],["9524.7","3"],["9524.9","16"],["9525.1","10"],["9525.3","6"]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.PartialBookDepthStreams().Symbol("btcusdt").Levels(int64(10)).Execute()
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

	t.Run("Test WebsocketMarketStreamsAPIService PartialBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.PartialBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService PartialBookDepthStreams Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.PartialBookDepthStreams().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WebsocketMarketStreamsAPIService PartialBookDepthStreams Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.WebsocketMarketStreamsAPI.PartialBookDepthStreams().Symbol("btcusdt").Levels(int64(10)).Execute()
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
