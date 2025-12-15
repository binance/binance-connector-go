package binancespotwebsocketapi

import (
	"encoding/json"
	"errors"
	"testing"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/binance-connector-go/clients/spot/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/common"
	tests "github.com/binance/binance-connector-go/common/tests"
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

		mockedJSON := `{"id":"123","status":200,"result":{"mins":5,"price":"9.35751834","closeTime":1694061154503},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"mins":5,"price":"9.35751834","closeTime":1694061154503},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/avgPrice"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":{"lastUpdateId":2731179239,"bids":[["0.01379900","3.43200000"],["0.01379800","3.24300000"],["0.01379700","10.45500000"],["0.01379600","3.82100000"],["0.01379500","10.26200000"]],"asks":[["0.01380000","5.91700000"],["0.01380100","6.01400000"],["0.01380200","0.26800000"],["0.01380300","0.33800000"],["0.01380400","0.26800000"]]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"lastUpdateId":2731179239,"bids":[["0.01379900","3.43200000"],["0.01379800","3.24300000"],["0.01379700","10.45500000"],["0.01379600","3.82100000"],["0.01379500","10.26200000"]],"asks":[["0.01380000","5.91700000"],["0.01380100","6.01400000"],["0.01380200","0.26800000"],["0.01380300","0.33800000"],["0.01380400","0.26800000"]]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/depth"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[[1655971200000,"0.01086000","0.01086600","0.01083600","0.01083800","2290.53800000",1655974799999,"24.85074442",2283,"1171.64000000","12.71225884","0"]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[[1655971200000,"0.01086000","0.01086600","0.01083600","0.01083800","2290.53800000",1655974799999,"24.85074442",2283,"1171.64000000","12.71225884","0"]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/klines"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","priceChange":"0.00061500","priceChangePercent":"4.735","weightedAvgPrice":"0.01368242","openPrice":"0.01298900","highPrice":"0.01418800","lowPrice":"0.01296000","lastPrice":"0.01360400","volume":"587179.23900000","quoteVolume":"8034.03382165","openTime":1659580020000,"closeTime":1660184865291,"firstId":192977765,"lastId":195365758,"count":2387994},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","priceChange":"0.00061500","priceChangePercent":"4.735","weightedAvgPrice":"0.01368242","openPrice":"0.01298900","highPrice":"0.01418800","lowPrice":"0.01296000","lastPrice":"0.01360400","volume":"587179.23900000","quoteVolume":"8034.03382165","openTime":1659580020000,"closeTime":1660184865291,"firstId":192977765,"lastId":195365758,"count":2387994},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
		require.Equal(t, "/ticker"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","priceChange":"0.00013900","priceChangePercent":"1.020","weightedAvgPrice":"0.01382453","prevClosePrice":"0.01362800","lastPrice":"0.01376700","lastQty":"1.78800000","bidPrice":"0.01376700","bidQty":"4.64600000","askPrice":"0.01376800","askQty":"14.31400000","openPrice":"0.01362800","highPrice":"0.01414900","lowPrice":"0.01346600","volume":"69412.40500000","quoteVolume":"959.59411487","openTime":1660014164909,"closeTime":1660100564909,"firstId":194696115,"lastId":194968287,"count":272173},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","priceChange":"0.00013900","priceChangePercent":"1.020","weightedAvgPrice":"0.01382453","prevClosePrice":"0.01362800","lastPrice":"0.01376700","lastQty":"1.78800000","bidPrice":"0.01376700","bidQty":"4.64600000","askPrice":"0.01376800","askQty":"14.31400000","openPrice":"0.01362800","highPrice":"0.01414900","lowPrice":"0.01346600","volume":"69412.40500000","quoteVolume":"959.59411487","openTime":1660014164909,"closeTime":1660100564909,"firstId":194696115,"lastId":194968287,"count":272173},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/ticker.24hr"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","bidPrice":"0.01358000","bidQty":"12.53400000","askPrice":"0.01358100","askQty":"17.83700000"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","bidPrice":"0.01358000","bidQty":"12.53400000","askPrice":"0.01358100","askQty":"17.83700000"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/ticker.book"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","price":"0.01361900"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BNBBTC","price":"0.01361900"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/ticker.price"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BNBUSDT","priceChange":"2.60000000","priceChangePercent":"1.238","weightedAvgPrice":"211.92276958","openPrice":"210.00000000","highPrice":"213.70000000","lowPrice":"209.70000000","lastPrice":"212.60000000","volume":"280709.58900000","quoteVolume":"59488753.54750000","openTime":1695686400000,"closeTime":1695772799999,"firstId":672397461,"lastId":672496158,"count":98698},{"symbol":"BTCUSDT","priceChange":"-83.13000000","priceChangePercent":"-0.317","weightedAvgPrice":"26234.58803036","openPrice":"26304.80000000","highPrice":"26397.46000000","lowPrice":"26088.34000000","lastPrice":"26221.67000000","volume":"18495.35066000","quoteVolume":"485217905.04210480","openTime":1695686400000,"closeTime":1695772799999,"firstId":3220151555,"lastId":3220849281,"count":697727}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":8}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BNBUSDT","priceChange":"2.60000000","priceChangePercent":"1.238","weightedAvgPrice":"211.92276958","openPrice":"210.00000000","highPrice":"213.70000000","lowPrice":"209.70000000","lastPrice":"212.60000000","volume":"280709.58900000","quoteVolume":"59488753.54750000","openTime":1695686400000,"closeTime":1695772799999,"firstId":672397461,"lastId":672496158,"count":98698},{"symbol":"BTCUSDT","priceChange":"-83.13000000","priceChangePercent":"-0.317","weightedAvgPrice":"26234.58803036","openPrice":"26304.80000000","highPrice":"26397.46000000","lowPrice":"26088.34000000","lastPrice":"26221.67000000","volume":"18495.35066000","quoteVolume":"485217905.04210480","openTime":1695686400000,"closeTime":1695772799999,"firstId":3220151555,"lastId":3220849281,"count":697727}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":8}]}`
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
		require.Equal(t, "/ticker.tradingDay"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[{"a":50000000,"p":"0.00274100","q":"57.19000000","f":59120167,"l":59120170,"T":1565877971222,"m":true,"M":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[{"a":50000000,"p":"0.00274100","q":"57.19000000","f":59120167,"l":59120170,"T":1565877971222,"m":true,"M":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/trades.aggregate"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[{"id":0,"price":"0.00005000","qty":"40.00000000","quoteQty":"0.00200000","time":1500004800376,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":10}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[{"id":0,"price":"0.00005000","qty":"40.00000000","quoteQty":"0.00200000","time":1500004800376,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":10}]}`
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
		require.Equal(t, "/trades.historical"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[{"id":194686783,"price":"0.01361000","qty":"0.01400000","quoteQty":"0.00019054","time":1660009530807,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[{"id":194686783,"price":"0.01361000","qty":"0.01400000","quoteQty":"0.00019054","time":1660009530807,"isBuyerMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/trades.recent"[1:], sent["method"])

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

		mockedJSON := `{"id":"123","status":200,"result":[[1655971200000,"0.01086000","0.01086600","0.01083600","0.01083800","2290.53800000",1655974799999,"24.85074442",2283,"1171.64000000","12.71225884","0"]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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

		mockedJSON := `{"id":"123","status":200,"result":[[1655971200000,"0.01086000","0.01086600","0.01083600","0.01083800","2290.53800000",1655974799999,"24.85074442",2283,"1171.64000000","12.71225884","0"]],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":2}]}`
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
		require.Equal(t, "/uiKlines"[1:], sent["method"])

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
