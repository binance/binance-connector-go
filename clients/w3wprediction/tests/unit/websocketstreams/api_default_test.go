package binancew3wpredictionwebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancew3wpredictionwebsocketstreams_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPI BatchClaimPartiallySucceeded Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_CLAIM_BATCH_20260528003_pm_claim_partial_success_e1f2a3b4"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BatchClaimPartiallySucceeded().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.BatchClaimPartiallySucceededResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.BatchClaimPartiallySucceededResponse
		mockCallback := func(msg models.BatchClaimPartiallySucceededResponse) {
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

	t.Run("Test DefaultAPIService BatchClaimPartiallySucceeded Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.BatchClaimPartiallySucceeded().Execute()
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
	t.Run("Test DefaultAPI FundTransferFailed Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_TRANSFER_20260528002_pm_transfer_fail_a3b4c5d6","amount":"100.00"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FundTransferFailed().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.FundTransferFailedResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.FundTransferFailedResponse
		mockCallback := func(msg models.FundTransferFailedResponse) {
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

	t.Run("Test DefaultAPIService FundTransferFailed Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FundTransferFailed().Execute()
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
	t.Run("Test DefaultAPI FundTransferSucceeded Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_TRANSFER_20260528001_pm_transfer_success_f2a3b4c5","amount":"100.00"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FundTransferSucceeded().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.FundTransferSucceededResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.FundTransferSucceededResponse
		mockCallback := func(msg models.FundTransferSucceededResponse) {
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

	t.Run("Test DefaultAPIService FundTransferSucceeded Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.FundTransferSucceeded().Execute()
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
	t.Run("Test DefaultAPI LimitOrderFullyFilled Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812347_pm_limit_order_filled_a7b8c9d0","amount":"100.00","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderFullyFilled().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.LimitOrderFullyFilledResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.LimitOrderFullyFilledResponse
		mockCallback := func(msg models.LimitOrderFullyFilledResponse) {
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

	t.Run("Test DefaultAPIService LimitOrderFullyFilled Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderFullyFilled().Execute()
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
	t.Run("Test DefaultAPI LimitOrderPartiallyFilled Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812347_pm_limit_order_partial_fill_b8c9d0e1","amount":"30.00","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderPartiallyFilled().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.LimitOrderPartiallyFilledResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.LimitOrderPartiallyFilledResponse
		mockCallback := func(msg models.LimitOrderPartiallyFilledResponse) {
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

	t.Run("Test DefaultAPIService LimitOrderPartiallyFilled Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderPartiallyFilled().Execute()
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
	t.Run("Test DefaultAPI LimitOrderSubmissionFailed Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812347_pm_limit_submit_fail_f6a7b8c9","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderSubmissionFailed().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.LimitOrderSubmissionFailedResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.LimitOrderSubmissionFailedResponse
		mockCallback := func(msg models.LimitOrderSubmissionFailedResponse) {
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

	t.Run("Test DefaultAPIService LimitOrderSubmissionFailed Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderSubmissionFailed().Execute()
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
	t.Run("Test DefaultAPI LimitOrderSubmittedSuccessfully Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812347_pm_limit_submit_success_e5f6a7b8","amount":"100.00","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderSubmittedSuccessfully().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.LimitOrderSubmittedSuccessfullyResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.LimitOrderSubmittedSuccessfullyResponse
		mockCallback := func(msg models.LimitOrderSubmittedSuccessfullyResponse) {
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

	t.Run("Test DefaultAPIService LimitOrderSubmittedSuccessfully Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.LimitOrderSubmittedSuccessfully().Execute()
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
	t.Run("Test DefaultAPI MarketBuyOrderFailed Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812345_pm_market_buy_fail_b2c3d4e5","amount":"12.34","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketBuyOrderFailed().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarketBuyOrderFailedResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarketBuyOrderFailedResponse
		mockCallback := func(msg models.MarketBuyOrderFailedResponse) {
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

	t.Run("Test DefaultAPIService MarketBuyOrderFailed Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketBuyOrderFailed().Execute()
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
	t.Run("Test DefaultAPI MarketBuyOrderFullyFilled Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812345_pm_market_buy_success_a1b2c3d4","amount":"12.34","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketBuyOrderFullyFilled().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarketBuyOrderFullyFilledResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarketBuyOrderFullyFilledResponse
		mockCallback := func(msg models.MarketBuyOrderFullyFilledResponse) {
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

	t.Run("Test DefaultAPIService MarketBuyOrderFullyFilled Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketBuyOrderFullyFilled().Execute()
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
	t.Run("Test DefaultAPI MarketResolutionNotification Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_8859231_pm_market_close_b4c5d6e7","topic":"BTC will hi.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketResolutionNotification().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarketResolutionNotificationResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarketResolutionNotificationResponse
		mockCallback := func(msg models.MarketResolutionNotificationResponse) {
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

	t.Run("Test DefaultAPIService MarketResolutionNotification Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketResolutionNotification().Execute()
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
	t.Run("Test DefaultAPI MarketSellOrderFailedExpired Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812346_pm_market_sell_fail_d4e5f6a7","amount":"10.50","topic":"ETH price.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketSellOrderFailedExpired().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarketSellOrderFailedExpiredResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarketSellOrderFailedExpiredResponse
		mockCallback := func(msg models.MarketSellOrderFailedExpiredResponse) {
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

	t.Run("Test DefaultAPIService MarketSellOrderFailedExpired Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketSellOrderFailedExpired().Execute()
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
	t.Run("Test DefaultAPI MarketSellOrderFullyFilled Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_PM2026052812346_pm_market_sell_success_c3d4e5f6","amount":"10.50","topic":"ETH price.."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketSellOrderFullyFilled().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.MarketSellOrderFullyFilledResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.MarketSellOrderFullyFilledResponse
		mockCallback := func(msg models.MarketSellOrderFullyFilledResponse) {
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

	t.Run("Test DefaultAPIService MarketSellOrderFullyFilled Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.MarketSellOrderFullyFilled().Execute()
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
	t.Run("Test DefaultAPI PredictionOrderBookByMarketId Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"msgType":"orderbook","marketId":8859231,"updateTimestampMs":1717420800123,"asks":[[0.32,500],[0.33,1200.5]],"bids":[[0.31,800],[0.3,2000]]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PredictionOrderBookByMarketId().MarketId("8859231").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.PredictionOrderBookByMarketIdResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.PredictionOrderBookByMarketIdResponse
		mockCallback := func(msg models.PredictionOrderBookByMarketIdResponse) {
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

	t.Run("Test DefaultAPIService PredictionOrderBookByMarketId Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PredictionOrderBookByMarketId().Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService PredictionOrderBookByMarketId Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.PredictionOrderBookByMarketId().MarketId("8859231").Execute()
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
	t.Run("Test DefaultAPI RewardClaimFailed Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_CLAIM_BATCH_20260528002_pm_claim_fail_d0e1f2a3","amount":"10.00","outcome":"Yes"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RewardClaimFailed().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.RewardClaimFailedResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.RewardClaimFailedResponse
		mockCallback := func(msg models.RewardClaimFailedResponse) {
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

	t.Run("Test DefaultAPIService RewardClaimFailed Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RewardClaimFailed().Execute()
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
	t.Run("Test DefaultAPI RewardClaimSucceeded Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"pushId":"pm_CLAIM_BATCH_20260528001_pm_claim_success_c9d0e1f2","amount":"50.00"}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RewardClaimSucceeded().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.RewardClaimSucceededResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.RewardClaimSucceededResponse
		mockCallback := func(msg models.RewardClaimSucceededResponse) {
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

	t.Run("Test DefaultAPIService RewardClaimSucceeded Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceW3wPredictionClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.RewardClaimSucceeded().Execute()
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
