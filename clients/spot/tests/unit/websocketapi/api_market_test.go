package binancespotwebsocketapi

import (
	"encoding/json"
	"errors"
	"testing"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancespotwebsocketapi_MarketAPIService(t *testing.T) {

	t.Run("Test MarketAPIService AvgPrice AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.AvgPrice().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"mins":5,"closeTime":1694061154503},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/avgPrice"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AvgPriceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService AvgPrice Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AvgPriceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.AvgPrice().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.AvgPriceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"mins":5,"closeTime":1694061154503},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/avgPrice"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AvgPriceResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService AvgPrice Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.AvgPrice().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService AvgPrice Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.AvgPrice().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test MarketAPIService BlockTradesHistorical AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().Symbol("BNBBTC").FromId(int64(582)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":582,"price":"0.052","qty":"5838","quoteQty":"303.576","time":1772506983321,"isBuyerMaker":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":10}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/blockTrades.historical"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.BlockTradesHistoricalResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService BlockTradesHistorical Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.BlockTradesHistoricalResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().Symbol("BNBBTC").FromId(int64(582)).Execute()
			resultChan <- common.ResultWebsocket[models.BlockTradesHistoricalResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":582,"price":"0.052","qty":"5838","quoteQty":"303.576","time":1772506983321,"isBuyerMaker":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":10}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/blockTrades.historical"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.BlockTradesHistoricalResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService BlockTradesHistorical Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService BlockTradesHistorical Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService BlockTradesHistorical Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.BlockTradesHistorical().Symbol("BNBBTC").FromId(int64(582)).ExecuteAsync()
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
	t.Run("Test MarketAPIService Depth AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.Depth().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"lastUpdateId":2731179239,"bids":[["0.01379900","3.43200000"]],"asks":[["0.01380000","5.91700000"]]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/depth"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.DepthResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService Depth Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.DepthResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.Depth().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.DepthResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"lastUpdateId":2731179239,"bids":[["0.01379900","3.43200000"]],"asks":[["0.01380000","5.91700000"]]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/depth"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.DepthResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService Depth Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.Depth().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService Depth Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.Depth().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test MarketAPIService Klines AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.Klines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[[1499040000000]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/klines"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.KlinesResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService Klines Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.KlinesResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.Klines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).Execute()
			resultChan <- common.ResultWebsocket[models.KlinesResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[[1499040000000]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/klines"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.KlinesResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService Klines Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.Klines().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService Klines Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.Klines().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService Klines Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.Klines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).ExecuteAsync()
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
	t.Run("Test MarketAPIService ReferencePrice AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.ReferencePrice().Symbol("BAZUSD").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BAZUSD","referencePrice":"0.00501900","timestamp":1770946889251,"code":-2043,"msg":"This symbol doesn't have a reference price."},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/referencePrice"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.ReferencePriceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService ReferencePrice Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.ReferencePriceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.ReferencePrice().Symbol("BAZUSD").Execute()
			resultChan <- common.ResultWebsocket[models.ReferencePriceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BAZUSD","referencePrice":"0.00501900","timestamp":1770946889251,"code":-2043,"msg":"This symbol doesn't have a reference price."},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/referencePrice"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.ReferencePriceResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService ReferencePrice Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.ReferencePrice().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService ReferencePrice Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.ReferencePrice().Symbol("BAZUSD").ExecuteAsync()
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
	t.Run("Test MarketAPIService ReferencePriceCalculation AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.ReferencePriceCalculation().Symbol("BAZUSD").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BAZUSD","calculationType":"ARITHMETIC_MEAN","bucketCount":10,"bucketWidthMs":1000,"externalCalculationId":42},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/referencePrice.calculation"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.ReferencePriceCalculationResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService ReferencePriceCalculation Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.ReferencePriceCalculationResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.ReferencePriceCalculation().Symbol("BAZUSD").Execute()
			resultChan <- common.ResultWebsocket[models.ReferencePriceCalculationResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BAZUSD","calculationType":"ARITHMETIC_MEAN","bucketCount":10,"bucketWidthMs":1000,"externalCalculationId":42},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/referencePrice.calculation"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.ReferencePriceCalculationResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService ReferencePriceCalculation Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.ReferencePriceCalculation().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService ReferencePriceCalculation Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.ReferencePriceCalculation().Symbol("BAZUSD").ExecuteAsync()
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
	t.Run("Test MarketAPIService Ticker AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.Ticker().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","openTime":1659580020000,"closeTime":1660184865291,"firstId":192977765,"lastId":195365758,"count":2387994},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/ticker"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TickerResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService Ticker Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TickerResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.Ticker().Execute()
			resultChan <- common.ResultWebsocket[models.TickerResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","openTime":1659580020000,"closeTime":1660184865291,"firstId":192977765,"lastId":195365758,"count":2387994},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/ticker"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TickerResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService Ticker Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.Ticker().ExecuteAsync()
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
	t.Run("Test MarketAPIService Ticker24hr AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.Ticker24hr().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","openTime":1660014164909,"closeTime":1660100564909,"firstId":194696115,"lastId":194968287,"count":272173},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/ticker.24hr"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.Ticker24hrResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService Ticker24hr Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.Ticker24hrResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.Ticker24hr().Execute()
			resultChan <- common.ResultWebsocket[models.Ticker24hrResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","openTime":1660014164909,"closeTime":1660100564909,"firstId":194696115,"lastId":194968287,"count":272173},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/ticker.24hr"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.Ticker24hrResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService Ticker24hr Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.Ticker24hr().ExecuteAsync()
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
	t.Run("Test MarketAPIService TickerBook AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TickerBook().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/ticker.book"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TickerBookResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TickerBook Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TickerBookResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TickerBook().Execute()
			resultChan <- common.ResultWebsocket[models.TickerBookResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/ticker.book"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TickerBookResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TickerBook Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TickerBook().ExecuteAsync()
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
	t.Run("Test MarketAPIService TickerPrice AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TickerPrice().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/ticker.price"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TickerPriceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TickerPrice Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TickerPriceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TickerPrice().Execute()
			resultChan <- common.ResultWebsocket[models.TickerPriceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/ticker.price"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TickerPriceResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TickerPrice Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TickerPrice().ExecuteAsync()
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
	t.Run("Test MarketAPIService TickerTradingDay AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TickerTradingDay().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","openTime":1695686400000,"closeTime":1695772799999,"firstId":3220151555,"lastId":3220849281,"count":697727}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/ticker.tradingDay"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TickerTradingDayResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TickerTradingDay Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TickerTradingDayResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TickerTradingDay().Execute()
			resultChan <- common.ResultWebsocket[models.TickerTradingDayResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","openTime":1695686400000,"closeTime":1695772799999,"firstId":3220151555,"lastId":3220849281,"count":697727}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/ticker.tradingDay"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TickerTradingDayResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TickerTradingDay Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TickerTradingDay().ExecuteAsync()
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
	t.Run("Test MarketAPIService TradesAggregate AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TradesAggregate().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"a":50000000,"f":59120167,"l":59120170,"T":1565877971222,"m":true,"M":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/trades.aggregate"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TradesAggregateResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TradesAggregate Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TradesAggregateResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TradesAggregate().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.TradesAggregateResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"a":50000000,"f":59120167,"l":59120170,"T":1565877971222,"m":true,"M":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/trades.aggregate"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TradesAggregateResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TradesAggregate Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.TradesAggregate().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService TradesAggregate Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TradesAggregate().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test MarketAPIService TradesHistorical AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TradesHistorical().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":0,"time":1500004800376,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/trades.historical"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TradesHistoricalResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TradesHistorical Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TradesHistoricalResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TradesHistorical().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.TradesHistoricalResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":0,"time":1500004800376,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/trades.historical"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TradesHistoricalResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TradesHistorical Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.TradesHistorical().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService TradesHistorical Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TradesHistorical().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test MarketAPIService TradesRecent AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.TradesRecent().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":194686783,"time":1660009530807,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/trades.recent"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.TradesRecentResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService TradesRecent Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.TradesRecentResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.TradesRecent().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.TradesRecentResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"id":194686783,"time":1660009530807,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/trades.recent"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.TradesRecentResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService TradesRecent Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.TradesRecent().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService TradesRecent Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.TradesRecent().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test MarketAPIService UiKlines AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.MarketAPI.UiKlines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[[1499040000000]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		select {
		case resp := <-responseChan:
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotEmpty(t, mockWS.MessagesWritten)

			require.Len(t, mockWS.MessagesWritten, 1)
			var sent map[string]any
			err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
			require.NoError(t, err)
			require.Equal(t, "/uiKlines"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.UiKlinesResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test MarketAPIService UiKlines Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.UiKlinesResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.MarketAPI.UiKlines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).Execute()
			resultChan <- common.ResultWebsocket[models.UiKlinesResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[[1499040000000]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

		var mocked map[string]interface{}
		err = json.Unmarshal([]byte(mockedJSON), &mocked)
		require.NoError(t, err)

		mocked["id"] = sent["id"]

		finalJSON, err := json.Marshal(mocked)
		require.NoError(t, err)

		mockWS.QueueMessage(finalJSON)

		res := <-resultChan
		resp := res.Value
		err = res.Err

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.Len(t, mockWS.MessagesWritten, 1)
		var sentCheck map[string]any
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sentCheck)
		require.NoError(t, err)
		require.Equal(t, "/uiKlines"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.UiKlinesResponse{}, typedResp)
	})

	t.Run("Test MarketAPIService UiKlines Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.UiKlines().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService UiKlines Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.MarketAPI.UiKlines().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test MarketAPIService UiKlines Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.MarketAPI.UiKlines().Symbol("BNBUSDT").Interval(models.KlinesIntervalParameterInterval1s).ExecuteAsync()
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
