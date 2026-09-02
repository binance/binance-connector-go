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

func Test_binancespotwebsocketapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService OpenOrdersCancelAll AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OpenOrdersCancelAll().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":-1,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"iuVNVJYYrByz6C4yGOPPK0","transactionTime":1660803702431,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569099453,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"bX5wROblo6YeDwa9iTLeyY","orderId":12569099453,"orderListId":19431,"clientOrderId":"OFFXQtxVFZ6Nbcg4PgE2DA","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"origClientOrderId":"4d96324ff9d44481926157","orderId":12569099453,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"trailingTime":-1,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/openOrders.cancelAll"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OpenOrdersCancelAllResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OpenOrdersCancelAll Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OpenOrdersCancelAllResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OpenOrdersCancelAll().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.OpenOrdersCancelAllResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":-1,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"iuVNVJYYrByz6C4yGOPPK0","transactionTime":1660803702431,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569099453,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"bX5wROblo6YeDwa9iTLeyY","orderId":12569099453,"orderListId":19431,"clientOrderId":"OFFXQtxVFZ6Nbcg4PgE2DA","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"origClientOrderId":"4d96324ff9d44481926157","orderId":12569099453,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"trailingTime":-1,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/openOrders.cancelAll"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OpenOrdersCancelAllResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OpenOrdersCancelAll Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OpenOrdersCancelAll().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OpenOrdersCancelAll Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OpenOrdersCancelAll().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderAmendKeepPriority AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().Symbol("BNBUSDT").NewQty(float64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"transactTime":1741924229819,"executionId":60,"amendedOrder":{"symbol":"BTUCSDT","orderId":23,"orderListId":4,"origClientOrderId":"my_pending_order","clientOrderId":"xbxXh5SSwaHS7oUEOCI88B","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1741924204920,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"listStatus":{"orderListId":4,"contingencyType":"OTO","listOrderStatus":"EXECUTING","listClientOrderId":"8nOGLLawudj1QoOiwbroRH","symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":22,"clientOrderId":"g04EWsjaackzedjC9wRkWD"}]}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/order.amend.keepPriority"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderAmendKeepPriorityResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderAmendKeepPriorityResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().Symbol("BNBUSDT").NewQty(float64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.OrderAmendKeepPriorityResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"transactTime":1741924229819,"executionId":60,"amendedOrder":{"symbol":"BTUCSDT","orderId":23,"orderListId":4,"origClientOrderId":"my_pending_order","clientOrderId":"xbxXh5SSwaHS7oUEOCI88B","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1741924204920,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"listStatus":{"orderListId":4,"contingencyType":"OTO","listOrderStatus":"EXECUTING","listClientOrderId":"8nOGLLawudj1QoOiwbroRH","symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":22,"clientOrderId":"g04EWsjaackzedjC9wRkWD"}]}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.amend.keepPriority"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderAmendKeepPriorityResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderAmendKeepPriority().Symbol("BNBUSDT").NewQty(float64(1)).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderCancel AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancel().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","origClientOrderId":"4d96324ff9d44481926157","orderId":12569099453,"orderListId":19431,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY","contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"iuVNVJYYrByz6C4yGOPPK0","transactionTime":1660803702431,"orders":[{"symbol":"BTCUSDT","orderId":12569099453,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"bX5wROblo6YeDwa9iTLeyY","orderId":12569099453,"orderListId":19431,"clientOrderId":"OFFXQtxVFZ6Nbcg4PgE2DA","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/order.cancel"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderCancelResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderCancel Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderCancelResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderCancel().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.OrderCancelResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","origClientOrderId":"4d96324ff9d44481926157","orderId":12569099453,"orderListId":19431,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY","contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"iuVNVJYYrByz6C4yGOPPK0","transactionTime":1660803702431,"orders":[{"symbol":"BTCUSDT","orderId":12569099453,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"bX5wROblo6YeDwa9iTLeyY","orderId":12569099453,"orderListId":19431,"clientOrderId":"OFFXQtxVFZ6Nbcg4PgE2DA","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.cancel"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderCancelResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderCancel Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancel().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderCancel Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderCancel().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderCancelReplace AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().Symbol("BNBUSDT").CancelReplaceMode(models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure).Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"cancelResult":"SUCCESS","newOrderResult":"SUCCESS","cancelResponse":{"symbol":"BTCUSDT","origClientOrderId":"4d96324ff9d44481926157","orderId":125690984230,"orderListId":-1,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"newOrderResponse":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY","transactTime":1660813156959,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/order.cancelReplace"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderCancelReplaceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderCancelReplace Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderCancelReplaceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().Symbol("BNBUSDT").CancelReplaceMode(models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure).Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).Execute()
			resultChan <- common.ResultWebsocket[models.OrderCancelReplaceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"cancelResult":"SUCCESS","newOrderResult":"SUCCESS","cancelResponse":{"symbol":"BTCUSDT","origClientOrderId":"4d96324ff9d44481926157","orderId":125690984230,"orderListId":-1,"clientOrderId":"91fe37ce9e69c90d6358c0","transactTime":1684804350068,"status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"newOrderResponse":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"bX5wROblo6YeDwa9iTLeyY","transactTime":1660813156959,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.cancelReplace"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderCancelReplaceResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderCancelReplace().Symbol("BNBUSDT").CancelReplaceMode(models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure).Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListCancel AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListCancel().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"6023531d7edaad348f5aff","transactionTime":1660801720215,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}],"orderReports":[{"symbol":"BTCUSDT","orderId":12569138901,"orderListId":1274512,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU","transactTime":1660801720215,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/orderList.cancel"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListCancelResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListCancel Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListCancelResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListCancel().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.OrderListCancelResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"6023531d7edaad348f5aff","transactionTime":1660801720215,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}],"orderReports":[{"symbol":"BTCUSDT","orderId":12569138901,"orderListId":1274512,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU","transactTime":1660801720215,"status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.cancel"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListCancelResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListCancel Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListCancel().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListCancel Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListCancel().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlace AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Price(float64(1)).Quantity(float64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}],"orderReports":[{"symbol":"BTCUSDT","orderId":12569138901,"orderListId":1274512,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU","transactTime":1660801713793,"status":"NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","workingTime":-1,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/orderList.place"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlace Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Price(float64(1)).Quantity(float64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}],"orderReports":[{"symbol":"BTCUSDT","orderId":12569138901,"orderListId":1274512,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU","transactTime":1660801713793,"status":"NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","workingTime":-1,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.place"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlace Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Price(float64(1)).Quantity(float64(1)).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlaceOco AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Quantity(float64(1)).AboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).BelowType(models.OrderListPlaceOcoBelowTypeParameterStopLoss).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":2,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"cKPMnDCbcLQILtDYM4f4fX","transactionTime":1711062760648,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":2,"clientOrderId":"0m6I4wfxvTUrOBSMUl0OPU"}],"orderReports":[{"symbol":"LTCBNB","orderId":2,"orderListId":2,"clientOrderId":"0m6I4wfxvTUrOBSMUl0OPU","transactTime":1711062760648,"status":"NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","workingTime":-1,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/orderList.place.oco"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceOcoResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceOcoResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Quantity(float64(1)).AboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).BelowType(models.OrderListPlaceOcoBelowTypeParameterStopLoss).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceOcoResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":2,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"cKPMnDCbcLQILtDYM4f4fX","transactionTime":1711062760648,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":2,"clientOrderId":"0m6I4wfxvTUrOBSMUl0OPU"}],"orderReports":[{"symbol":"LTCBNB","orderId":2,"orderListId":2,"clientOrderId":"0m6I4wfxvTUrOBSMUl0OPU","transactTime":1711062760648,"status":"NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","workingTime":-1,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.place.oco"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceOcoResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOco Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOco().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Quantity(float64(1)).AboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).BelowType(models.OrderListPlaceOcoBelowTypeParameterStopLoss).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlaceOpo AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":2,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"OiOgqvRagBefpzdM5gjYX3","transactionTime":1762941318142,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":2,"clientOrderId":"pUzhKBbc0ZVdMScIRAqitH"}],"orderReports":[{"symbol":"BTCUSDT","orderId":2,"orderListId":2,"clientOrderId":"pUzhKBbc0ZVdMScIRAqitH","transactTime":1762941318142,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1762941318142,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]}}`

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
			require.Equal(t, "/orderList.place.opo"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceOpoResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceOpoResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceOpoResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":2,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"OiOgqvRagBefpzdM5gjYX3","transactionTime":1762941318142,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":2,"clientOrderId":"pUzhKBbc0ZVdMScIRAqitH"}],"orderReports":[{"symbol":"BTCUSDT","orderId":2,"orderListId":2,"clientOrderId":"pUzhKBbc0ZVdMScIRAqitH","transactTime":1762941318142,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1762941318142,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]}}`

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
		require.Equal(t, "/orderList.place.opo"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceOpoResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpo Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpo().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlaceOpoco AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":1,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"TVbG6ymkYMXTj7tczbOsBf","transactionTime":1763000139104,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":6,"clientOrderId":"3czuJSeyjPwV9Xo28j1Dv3"}],"orderReports":[{"symbol":"BTCUSDT","orderId":6,"orderListId":1,"clientOrderId":"3czuJSeyjPwV9Xo28j1Dv3","transactTime":1763000139104,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1763000139104,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]}}`

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
			require.Equal(t, "/orderList.place.opoco"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceOpocoResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceOpocoResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceOpocoResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":1,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"TVbG6ymkYMXTj7tczbOsBf","transactionTime":1763000139104,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":6,"clientOrderId":"3czuJSeyjPwV9Xo28j1Dv3"}],"orderReports":[{"symbol":"BTCUSDT","orderId":6,"orderListId":1,"clientOrderId":"3czuJSeyjPwV9Xo28j1Dv3","transactTime":1763000139104,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1763000139104,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]}}`

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
		require.Equal(t, "/orderList.place.opoco"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceOpocoResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOpoco Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOpoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlaceOto AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":626,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"KA4EBjGnzvSwSCQsDdTrlf","transactionTime":1712544395981,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":13,"clientOrderId":"YiAUtM9yJjl1a2jXHSp9Ny"}],"orderReports":[{"symbol":"LTCBNB","orderId":13,"orderListId":626,"clientOrderId":"YiAUtM9yJjl1a2jXHSp9Ny","transactTime":1712544395981,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1712544395981,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/orderList.place.oto"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceOtoResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceOtoResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceOtoResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":626,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"KA4EBjGnzvSwSCQsDdTrlf","transactionTime":1712544395981,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":13,"clientOrderId":"YiAUtM9yJjl1a2jXHSp9Ny"}],"orderReports":[{"symbol":"LTCBNB","orderId":13,"orderListId":626,"clientOrderId":"YiAUtM9yJjl1a2jXHSp9Ny","transactTime":1712544395981,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1712544395981,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.place.oto"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceOtoResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOto Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOto().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingType(models.OrderListPlaceOpoPendingTypeParameterLimit).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderListPlaceOtoco AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":629,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"GaeJHjZPasPItFj4x7Mqm6","transactionTime":1712544408537,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":23,"clientOrderId":"OVQOpKwfmPCfaBTD0n7e7H"}],"orderReports":[{"symbol":"LTCBNB","orderId":23,"orderListId":629,"clientOrderId":"OVQOpKwfmPCfaBTD0n7e7H","transactTime":1712544408537,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1712544408537,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/orderList.place.otoco"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListPlaceOtocoResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListPlaceOtocoResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).Execute()
			resultChan <- common.ResultWebsocket[models.OrderListPlaceOtocoResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"orderListId":629,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"GaeJHjZPasPItFj4x7Mqm6","transactionTime":1712544408537,"symbol":"LTCBNB","orders":[{"symbol":"LTCBNB","orderId":23,"clientOrderId":"OVQOpKwfmPCfaBTD0n7e7H"}],"orderReports":[{"symbol":"LTCBNB","orderId":23,"orderListId":629,"clientOrderId":"OVQOpKwfmPCfaBTD0n7e7H","transactTime":1712544408537,"status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1712544408537,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.place.otoco"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListPlaceOtocoResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderListPlaceOtoco Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderListPlaceOtoco().Symbol("BNBUSDT").WorkingType(models.OrderListPlaceOpoWorkingTypeParameterLimit).WorkingSide(models.OrderCancelReplaceSideParameterBuy).WorkingPrice(float64(1)).WorkingQuantity(float64(1)).PendingSide(models.OrderCancelReplaceSideParameterBuy).PendingQuantity(float64(1)).PendingAboveType(models.OrderListPlaceOcoAboveTypeParameterStopLossLimit).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderPlace AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157ec08158a40","transactTime":1660801715793,"status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1660801715639,"selfTradePreventionMode":"NONE","stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY","fills":[{"commissionAsset":"BNB","tradeId":1650422481}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/order.place"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderPlaceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderPlace Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderPlaceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).Execute()
			resultChan <- common.ResultWebsocket[models.OrderPlaceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157ec08158a40","transactTime":1660801715793,"status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1660801715639,"selfTradePreventionMode":"NONE","stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY","fills":[{"commissionAsset":"BNB","tradeId":1650422481}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.place"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderPlaceResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderPlace Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
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
	t.Run("Test TradeAPIService OrderTest AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/order.test"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderTestResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService OrderTest Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderTestResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).Execute()
			resultChan <- common.ResultWebsocket[models.OrderTestResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.test"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderTestResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService OrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService OrderTest Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.OrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.OrderCancelReplaceTypeParameterMarket).ExecuteAsync()
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
	t.Run("Test TradeAPIService SorOrderPlace AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":2,"orderListId":-1,"clientOrderId":"sBI1KM6nNtOfj5tccZSKly","transactTime":1689149087774,"status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1689149087774,"fills":[{"matchType":"ONE_PARTY_TRADE_REPORT","commissionAsset":"BTC","tradeId":-1,"allocId":0}],"workingFloor":"SOR","selfTradePreventionMode":"NONE","usedSor":true,"stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/sor.order.place"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.SorOrderPlaceResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService SorOrderPlace Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.SorOrderPlaceResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.SorOrderPlaceResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":2,"orderListId":-1,"clientOrderId":"sBI1KM6nNtOfj5tccZSKly","transactTime":1689149087774,"status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1689149087774,"fills":[{"matchType":"ONE_PARTY_TRADE_REPORT","commissionAsset":"BTC","tradeId":-1,"allocId":0}],"workingFloor":"SOR","selfTradePreventionMode":"NONE","usedSor":true,"stopPrice":"0.00000000","trailingDelta":10,"icebergQty":"0.00000000","strategyId":1,"strategyType":1000000,"preventedMatchId":0,"preventedQuantity":"1.200000","trailingTime":-1,"pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/sor.order.place"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.SorOrderPlaceResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService SorOrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderPlace Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderPlace Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.SorOrderPlace().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).ExecuteAsync()
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
	t.Run("Test TradeAPIService SorOrderTest AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
			require.Equal(t, "/sor.order.test"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.SorOrderTestResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService SorOrderTest Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.SorOrderTestResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.SorOrderTestResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/sor.order.test"[1:], sentCheck["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.SorOrderTestResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService SorOrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderTest Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService SorOrderTest Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.SorOrderTest().Symbol("BNBUSDT").Side(models.OrderCancelReplaceSideParameterBuy).Type(models.SorOrderPlaceTypeParameterMarket).Quantity(float64(1)).ExecuteAsync()
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
