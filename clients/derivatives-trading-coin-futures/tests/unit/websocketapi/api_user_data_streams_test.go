package binancederivativestradingcoinfutureswebsocketapi

import (
	"encoding/json"
	"errors"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
	tests "github.com/binance/binance-connector-go/common/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingcoinfutureswebsocketapi_UserDataStreamsAPIService(t *testing.T) {

	t.Run("Test UserDataStreamsAPIService CloseUserDataStream AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.UserDataStreamsAPI.CloseUserDataStream().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":2}]}`
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
			require.Equal(t, "/userDataStream.stop"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.CloseUserDataStreamResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test UserDataStreamsAPIService CloseUserDataStream Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.CloseUserDataStreamResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.UserDataStreamsAPI.CloseUserDataStream().Execute()
			resultChan <- common.ResultWebsocket[models.CloseUserDataStreamResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":2}]}`
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
		require.Equal(t, "/userDataStream.stop"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.CloseUserDataStreamResponse{}, typedResp)
	})

	t.Run("Test UserDataStreamsAPIService CloseUserDataStream Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.UserDataStreamsAPI.CloseUserDataStream().ExecuteAsync()
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
	t.Run("Test UserDataStreamsAPIService KeepaliveUserDataStream AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.UserDataStreamsAPI.KeepaliveUserDataStream().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"listenKey":"3HBntNTepshgEdjIwSUIBgB9keLyOCg5qv3n6bYAtktG8ejcaW5HXz9Vx1JgIieg"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":2}]}`
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
			require.Equal(t, "/userDataStream.ping"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.KeepaliveUserDataStreamResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test UserDataStreamsAPIService KeepaliveUserDataStream Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.KeepaliveUserDataStreamResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.UserDataStreamsAPI.KeepaliveUserDataStream().Execute()
			resultChan <- common.ResultWebsocket[models.KeepaliveUserDataStreamResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"listenKey":"3HBntNTepshgEdjIwSUIBgB9keLyOCg5qv3n6bYAtktG8ejcaW5HXz9Vx1JgIieg"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":2}]}`
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
		require.Equal(t, "/userDataStream.ping"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.KeepaliveUserDataStreamResponse{}, typedResp)
	})

	t.Run("Test UserDataStreamsAPIService KeepaliveUserDataStream Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.UserDataStreamsAPI.KeepaliveUserDataStream().ExecuteAsync()
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
	t.Run("Test UserDataStreamsAPIService StartUserDataStream AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.UserDataStreamsAPI.StartUserDataStream().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"listenKey":"xs0mRXdAKlIPDRFrlPcw0qI41Eh3ixNntmymGyhrhgqo7L6FuLaWArTD7RLP"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":8}]}`
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
			require.Equal(t, "/userDataStream.start"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.StartUserDataStreamResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test UserDataStreamsAPIService StartUserDataStream Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.StartUserDataStreamResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.UserDataStreamsAPI.StartUserDataStream().Execute()
			resultChan <- common.ResultWebsocket[models.StartUserDataStreamResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"listenKey":"xs0mRXdAKlIPDRFrlPcw0qI41Eh3ixNntmymGyhrhgqo7L6FuLaWArTD7RLP"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":8}]}`
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
		require.Equal(t, "/userDataStream.start"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.StartUserDataStreamResponse{}, typedResp)
	})

	t.Run("Test UserDataStreamsAPIService StartUserDataStream Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.UserDataStreamsAPI.StartUserDataStream().ExecuteAsync()
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
