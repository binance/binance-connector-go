package binancespotwebsocketapi

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancespotwebsocketapi_UserDataStreamAPIService(t *testing.T) {

	t.Run("Test UserDataStreamAPIService SessionSubscriptions AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.UserDataStreamAPI.SessionSubscriptions().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"subscriptionId":1},{"subscriptionId":0}]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/session.subscriptions"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.SessionSubscriptionsResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test UserDataStreamAPIService SessionSubscriptions Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.SessionSubscriptionsResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.UserDataStreamAPI.SessionSubscriptions().Execute()
			resultChan <- common.ResultWebsocket[models.SessionSubscriptionsResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"subscriptionId":1},{"subscriptionId":0}]}`
		mockWS.QueueMessage([]byte(mockedJSON))

		res := <-resultChan
		resp := res.Value
		err := res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sent map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)
		require.Equal(t, "/session.subscriptions"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.SessionSubscriptionsResponse{}, typedResp)
	})

	t.Run("Test UserDataStreamAPIService SessionSubscriptions Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.SessionSubscriptions().ExecuteAsync()
			if err != nil {
				var wsErr *common.WebSocketError
				if errors.As(err, &wsErr) {
					require.Contains(t, wsErr.Error(), "[-1001] Internal server error")
					require.Equal(t, "123", wsErr.ConnID)
					require.Equal(t, "error_response", wsErr.Op)
				} else {
					t.Errorf("unexpected error type: %T", err)
				}
				_, ok := <-respChan
				require.False(t, ok, "response channel should be closed")
			}
			close(done)
		}()

		<-mockWS.HasSentChan

		mockWS.QueueMessage([]byte(`{
			"id":"123",
			"status":500,
			"error":{
				"code":-1001,
				"msg":"Internal server error"
			}
		}`))

		<-done
	})
	t.Run("Test UserDataStreamAPIService UserDataStreamSubscribe Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultUserData[models.UserDataStreamSubscribeResponse, models.UserDataStreamEventsResponse], 1)
		go func() {
			resp, stream, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribe().Execute()
			resultChan <- common.ResultUserData[models.UserDataStreamSubscribeResponse, models.UserDataStreamEventsResponse]{Value: resp, Stream: stream, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"subscriptionId":0}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		res := <-resultChan
		resp := res.Value
		stream := res.Stream
		err := res.Err
		require.NoError(t, err)

		require.NotNil(t, resp)
		require.NotNil(t, stream)
		require.IsType(t, &common.StreamHandler[models.UserDataStreamEventsResponse]{}, stream)
		require.NotNil(t, stream.On)
		require.NotNil(t, stream.Unsubscribe)
	})

	t.Run("Test UserDataStreamAPIService UserDataStreamSubscribe Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resultChan := make(chan common.ResultUserData[models.UserDataStreamSubscribeResponse, models.UserDataStreamEventsResponse], 1)
		go func() {
			resp, stream, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribe().Execute()
			resultChan <- common.ResultUserData[models.UserDataStreamSubscribeResponse, models.UserDataStreamEventsResponse]{Value: resp, Stream: stream, Err: err}
		}()

		<-mockWS.HasSentChan

		mockWS.QueueMessage([]byte(`{
			"id":"123",
			"status":500,
			"error":{
				"code":-1001,
				"msg":"Internal server error"
			}
		}`))

		res := <-resultChan
		stream := res.Stream
		require.Nil(t, res.Value)
		require.Error(t, res.Err)

		require.Nil(t, stream)
		var wsErr *common.WebSocketError
		switch e := res.Err.(type) {
		case *common.WebSocketError:
			wsErr = e
		default:
			errStr := e.Error()
			if strings.Contains(errStr, "WebSocket error_response error") {
				connID := "123"
				if parts := strings.Split(errStr, "conn "); len(parts) > 1 {
					if idParts := strings.Split(parts[1], ")"); len(idParts) > 0 {
						connID = idParts[0]
					}
				}
				require.Contains(t, errStr, "[-1001] Internal server error")
				require.Contains(t, errStr, "conn "+connID)
				require.Contains(t, errStr, "error_response")
				return
			}
			t.Fatalf("Unexpected error type: %T", e)
		}

		require.Contains(t, wsErr.Error(), "[-1001] Internal server error")
		require.Equal(t, "123", wsErr.ConnID)
		require.Equal(t, "error_response", wsErr.Op)
	})
	t.Run("Test UserDataStreamAPIService UserDataStreamSubscribeSignature Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultUserData[models.UserDataStreamSubscribeSignatureResponse, models.UserDataStreamEventsResponse], 1)
		go func() {
			resp, stream, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribeSignature().Execute()
			resultChan <- common.ResultUserData[models.UserDataStreamSubscribeSignatureResponse, models.UserDataStreamEventsResponse]{Value: resp, Stream: stream, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"subscriptionId":0}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		res := <-resultChan
		resp := res.Value
		stream := res.Stream
		err := res.Err
		require.NoError(t, err)

		require.NotNil(t, resp)
		require.NotNil(t, stream)
		require.IsType(t, &common.StreamHandler[models.UserDataStreamEventsResponse]{}, stream)
		require.NotNil(t, stream.On)
		require.NotNil(t, stream.Unsubscribe)
	})

	t.Run("Test UserDataStreamAPIService UserDataStreamSubscribeSignature Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resultChan := make(chan common.ResultUserData[models.UserDataStreamSubscribeSignatureResponse, models.UserDataStreamEventsResponse], 1)
		go func() {
			resp, stream, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamSubscribeSignature().Execute()
			resultChan <- common.ResultUserData[models.UserDataStreamSubscribeSignatureResponse, models.UserDataStreamEventsResponse]{Value: resp, Stream: stream, Err: err}
		}()

		<-mockWS.HasSentChan

		mockWS.QueueMessage([]byte(`{
			"id":"123",
			"status":500,
			"error":{
				"code":-1001,
				"msg":"Internal server error"
			}
		}`))

		res := <-resultChan
		stream := res.Stream
		require.Nil(t, res.Value)
		require.Error(t, res.Err)

		require.Nil(t, stream)
		var wsErr *common.WebSocketError
		switch e := res.Err.(type) {
		case *common.WebSocketError:
			wsErr = e
		default:
			errStr := e.Error()
			if strings.Contains(errStr, "WebSocket error_response error") {
				connID := "123"
				if parts := strings.Split(errStr, "conn "); len(parts) > 1 {
					if idParts := strings.Split(parts[1], ")"); len(idParts) > 0 {
						connID = idParts[0]
					}
				}
				require.Contains(t, errStr, "[-1001] Internal server error")
				require.Contains(t, errStr, "conn "+connID)
				require.Contains(t, errStr, "error_response")
				return
			}
			t.Fatalf("Unexpected error type: %T", e)
		}

		require.Contains(t, wsErr.Error(), "[-1001] Internal server error")
		require.Equal(t, "123", wsErr.ConnID)
		require.Equal(t, "error_response", wsErr.Op)
	})
	t.Run("Test UserDataStreamAPIService UserDataStreamUnsubscribe AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamUnsubscribe().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/userDataStream.unsubscribe"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.UserDataStreamUnsubscribeResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test UserDataStreamAPIService UserDataStreamUnsubscribe Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.UserDataStreamUnsubscribeResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamUnsubscribe().Execute()
			resultChan <- common.ResultWebsocket[models.UserDataStreamUnsubscribeResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{}}`
		mockWS.QueueMessage([]byte(mockedJSON))

		res := <-resultChan
		resp := res.Value
		err := res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sent map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)
		require.Equal(t, "/userDataStream.unsubscribe"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.UserDataStreamUnsubscribeResponse{}, typedResp)
	})

	t.Run("Test UserDataStreamAPIService UserDataStreamUnsubscribe Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.UserDataStreamAPI.UserDataStreamUnsubscribe().ExecuteAsync()
			if err != nil {
				var wsErr *common.WebSocketError
				if errors.As(err, &wsErr) {
					require.Contains(t, wsErr.Error(), "[-1001] Internal server error")
					require.Equal(t, "123", wsErr.ConnID)
					require.Equal(t, "error_response", wsErr.Op)
				} else {
					t.Errorf("unexpected error type: %T", err)
				}
				_, ok := <-respChan
				require.False(t, ok, "response channel should be closed")
			}
			close(done)
		}()

		<-mockWS.HasSentChan

		mockWS.QueueMessage([]byte(`{
			"id":"123",
			"status":500,
			"error":{
				"code":-1001,
				"msg":"Internal server error"
			}
		}`))

		<-done
	})
}
