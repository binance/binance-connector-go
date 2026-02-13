package binancederivativestradingusdsfutureswebsocketapi

import (
	"encoding/json"
	"errors"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/websocketapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingusdsfutureswebsocketapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService CancelAlgoOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.CancelAlgoOrder().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"algoId":2000000002162519,"clientAlgoId":"rDMG8WSde6LkyMNtk6s825","code":"200","msg":"success"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":6}]}`
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
			require.Equal(t, "/algoOrder.cancel"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.CancelAlgoOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService CancelAlgoOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.CancelAlgoOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.CancelAlgoOrder().Execute()
			resultChan <- common.ResultWebsocket[models.CancelAlgoOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"algoId":2000000002162519,"clientAlgoId":"rDMG8WSde6LkyMNtk6s825","code":"200","msg":"success"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":6}]}`
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
		require.Equal(t, "/algoOrder.cancel"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.CancelAlgoOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService CancelAlgoOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.CancelAlgoOrder().ExecuteAsync()
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
	t.Run("Test TradeAPIService CancelOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.CancelOrder().Symbol("symbol_example").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"clientOrderId":"myOrder1","cumQty":"0","cumQuote":"0","executedQty":"0","orderId":283194212,"origQty":"11","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"CANCELED","stopPrice":"9300","closePosition":false,"symbol":"BTCUSDT","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1571110484038,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
			require.Equal(t, "/order.cancel"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.CancelOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService CancelOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.CancelOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.CancelOrder().Symbol("symbol_example").Execute()
			resultChan <- common.ResultWebsocket[models.CancelOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"clientOrderId":"myOrder1","cumQty":"0","cumQuote":"0","executedQty":"0","orderId":283194212,"origQty":"11","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"CANCELED","stopPrice":"9300","closePosition":false,"symbol":"BTCUSDT","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1571110484038,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
		require.Equal(t, "/order.cancel"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.CancelOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService CancelOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.CancelOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService CancelOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.CancelOrder().Symbol("symbol_example").ExecuteAsync()
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
	t.Run("Test TradeAPIService ModifyOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Quantity(float32(1.0)).Price(float32(1.0)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderId":328971409,"symbol":"BTCUSDT","status":"NEW","clientOrderId":"xGHfltUMExx0TbQstQQfRX","price":"43769.10","avgPrice":"0.00","origQty":"0.110","executedQty":"0.000","cumQty":"0.000","cumQuote":"0.00000","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"SELL","positionSide":"SHORT","stopPrice":"0.00","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0,"updateTime":1703426756190},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":300,"count":1},{"rateLimitType":"ORDERS","interval":"MINUTE","intervalNum":1,"limit":1200,"count":1},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
			require.Equal(t, "/order.modify"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.ModifyOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService ModifyOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.ModifyOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Quantity(float32(1.0)).Price(float32(1.0)).Execute()
			resultChan <- common.ResultWebsocket[models.ModifyOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderId":328971409,"symbol":"BTCUSDT","status":"NEW","clientOrderId":"xGHfltUMExx0TbQstQQfRX","price":"43769.10","avgPrice":"0.00","origQty":"0.110","executedQty":"0.000","cumQty":"0.000","cumQuote":"0.00000","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"SELL","positionSide":"SHORT","stopPrice":"0.00","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0,"updateTime":1703426756190},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":300,"count":1},{"rateLimitType":"ORDERS","interval":"MINUTE","intervalNum":1,"limit":1200,"count":1},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
		require.Equal(t, "/order.modify"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.ModifyOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService ModifyOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService ModifyOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService ModifyOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService ModifyOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService ModifyOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.ModifyOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Quantity(float32(1.0)).Price(float32(1.0)).ExecuteAsync()
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
	t.Run("Test TradeAPIService NewAlgoOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType("algoType_example").Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"algoId":3000000000003505,"clientAlgoId":"0Xkl1p621E4EryvufmYre1","algoType":"CONDITIONAL","orderType":"TAKE_PROFIT","symbol":"BTCUSDT","side":"SELL","positionSide":"SHORT","timeInForce":"GTC","quantity":"1.000","algoStatus":"NEW","triggerPrice":"120000.00","price":"160000.00","icebergQuantity":null,"selfTradePreventionMode":"EXPIRE_MAKER","workingType":"CONTRACT_PRICE","priceMatch":"NONE","closePosition":false,"priceProtect":false,"reduceOnly":false,"createTime":1762507264142,"updateTime":1762507264143,"triggerTime":0,"goodTillDate":0},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
			require.Equal(t, "/algoOrder.place"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.NewAlgoOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService NewAlgoOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.NewAlgoOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType("algoType_example").Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").Execute()
			resultChan <- common.ResultWebsocket[models.NewAlgoOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"algoId":3000000000003505,"clientAlgoId":"0Xkl1p621E4EryvufmYre1","algoType":"CONDITIONAL","orderType":"TAKE_PROFIT","symbol":"BTCUSDT","side":"SELL","positionSide":"SHORT","timeInForce":"GTC","quantity":"1.000","algoStatus":"NEW","triggerPrice":"120000.00","price":"160000.00","icebergQuantity":null,"selfTradePreventionMode":"EXPIRE_MAKER","workingType":"CONTRACT_PRICE","priceMatch":"NONE","closePosition":false,"priceProtect":false,"reduceOnly":false,"createTime":1762507264142,"updateTime":1762507264143,"triggerTime":0,"goodTillDate":0},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
		require.Equal(t, "/algoOrder.place"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.NewAlgoOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService NewAlgoOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewAlgoOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewAlgoOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewAlgoOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewAlgoOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.NewAlgoOrder().AlgoType("algoType_example").Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").ExecuteAsync()
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
	t.Run("Test TradeAPIService NewOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderId":325078477,"symbol":"BTCUSDT","status":"NEW","clientOrderId":"iCXL1BywlBaf2sesNUrVl3","price":"43187.00","avgPrice":"0.00","origQty":"0.100","executedQty":"0.000","cumQty":"0.000","cumQuote":"0.00000","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"BUY","positionSide":"BOTH","stopPrice":"0.00","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0,"updateTime":1702555534435},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":300,"count":1},{"rateLimitType":"ORDERS","interval":"MINUTE","intervalNum":1,"limit":1200,"count":1},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
			require.Equal(t, "/order.place"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.NewOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService NewOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.NewOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").Execute()
			resultChan <- common.ResultWebsocket[models.NewOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderId":325078477,"symbol":"BTCUSDT","status":"NEW","clientOrderId":"iCXL1BywlBaf2sesNUrVl3","price":"43187.00","avgPrice":"0.00","origQty":"0.100","executedQty":"0.000","cumQty":"0.000","cumQuote":"0.00000","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"BUY","positionSide":"BOTH","stopPrice":"0.00","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","goodTillDate":0,"updateTime":1702555534435},"rateLimits":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":300,"count":1},{"rateLimitType":"ORDERS","interval":"MINUTE","intervalNum":1,"limit":1200,"count":1},{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":1}]}`
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
		require.Equal(t, "/order.place"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.NewOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService NewOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService NewOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.NewOrder().Symbol("symbol_example").Side(models.ModifyOrderSideParameterBuy).Type("type__example").ExecuteAsync()
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
	t.Run("Test TradeAPIService PositionInformation AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.PositionInformation().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"entryPrice":"0.00000","breakEvenPrice":"0.0","marginType":"isolated","isAutoAddMargin":"false","isolatedMargin":"0.00000000","leverage":"10","liquidationPrice":"0","markPrice":"6679.50671178","maxNotionalValue":"20000000","positionAmt":"0.000","notional":"0","isolatedWallet":"0","symbol":"BTCUSDT","unRealizedProfit":"0.00000000","positionSide":"BOTH","updateTime":0},{"symbol":"BTCUSDT","positionAmt":"0.001","entryPrice":"22185.2","breakEvenPrice":"0.0","markPrice":"21123.05052574","unRealizedProfit":"-1.06214947","liquidationPrice":"19731.45529116","leverage":"4","maxNotionalValue":"100000000","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"LONG","notional":"21.12305052","isolatedWallet":"0","updateTime":1655217461579},{"symbol":"BTCUSDT","positionAmt":"0.000","entryPrice":"0.0","breakEvenPrice":"0.0","markPrice":"21123.05052574","unRealizedProfit":"0.00000000","liquidationPrice":"0","leverage":"4","maxNotionalValue":"100000000","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"SHORT","notional":"0","isolatedWallet":"0","updateTime":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":20}]}`
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
			require.Equal(t, "/account.position"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.PositionInformationResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService PositionInformation Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.PositionInformationResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.PositionInformation().Execute()
			resultChan <- common.ResultWebsocket[models.PositionInformationResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"entryPrice":"0.00000","breakEvenPrice":"0.0","marginType":"isolated","isAutoAddMargin":"false","isolatedMargin":"0.00000000","leverage":"10","liquidationPrice":"0","markPrice":"6679.50671178","maxNotionalValue":"20000000","positionAmt":"0.000","notional":"0","isolatedWallet":"0","symbol":"BTCUSDT","unRealizedProfit":"0.00000000","positionSide":"BOTH","updateTime":0},{"symbol":"BTCUSDT","positionAmt":"0.001","entryPrice":"22185.2","breakEvenPrice":"0.0","markPrice":"21123.05052574","unRealizedProfit":"-1.06214947","liquidationPrice":"19731.45529116","leverage":"4","maxNotionalValue":"100000000","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"LONG","notional":"21.12305052","isolatedWallet":"0","updateTime":1655217461579},{"symbol":"BTCUSDT","positionAmt":"0.000","entryPrice":"0.0","breakEvenPrice":"0.0","markPrice":"21123.05052574","unRealizedProfit":"0.00000000","liquidationPrice":"0","leverage":"4","maxNotionalValue":"100000000","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"SHORT","notional":"0","isolatedWallet":"0","updateTime":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":20}]}`
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
		require.Equal(t, "/account.position"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.PositionInformationResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService PositionInformation Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.PositionInformation().ExecuteAsync()
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
	t.Run("Test TradeAPIService PositionInformationV2 AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.PositionInformationV2().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","positionSide":"BOTH","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0},{"symbol":"BTCUSDT","positionSide":"LONG","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0},{"symbol":"BTCUSDT","positionSide":"SHORT","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":20}]}`
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
			require.Equal(t, "/v2/account.position"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.PositionInformationV2Response{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService PositionInformationV2 Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.PositionInformationV2Response], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.PositionInformationV2().Execute()
			resultChan <- common.ResultWebsocket[models.PositionInformationV2Response]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","positionSide":"BOTH","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0},{"symbol":"BTCUSDT","positionSide":"LONG","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0},{"symbol":"BTCUSDT","positionSide":"SHORT","positionAmt":"1.000","entryPrice":"0.00000","breakEvenPrice":"0.0","markPrice":"6679.50671178","unRealizedProfit":"0.00000000","liquidationPrice":"0","isolatedMargin":"0.00000000","notional":"0","marginAsset":"USDT","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","positionInitialMargin":"0","openOrderInitialMargin":"0","adl":0,"bidNotional":"0","askNotional":"0","updateTime":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":2400,"count":20}]}`
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
		require.Equal(t, "/v2/account.position"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.PositionInformationV2Response{}, typedResp)
	})

	t.Run("Test TradeAPIService PositionInformationV2 Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.PositionInformationV2().ExecuteAsync()
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
	t.Run("Test TradeAPIService QueryOrder AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.TradeAPI.QueryOrder().Symbol("symbol_example").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"avgPrice":"0.00000","clientOrderId":"abc","cumQuote":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSDT","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false}}`
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
			require.Equal(t, "/order.status"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.QueryOrderResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test TradeAPIService QueryOrder Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.QueryOrderResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.TradeAPI.QueryOrder().Symbol("symbol_example").Execute()
			resultChan <- common.ResultWebsocket[models.QueryOrderResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"avgPrice":"0.00000","clientOrderId":"abc","cumQuote":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSDT","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false}}`
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
		require.Equal(t, "/order.status"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.QueryOrderResponse{}, typedResp)
	})

	t.Run("Test TradeAPIService QueryOrder Missing Required Params", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		respChan, errChan, err := mockClient.WebsocketAPI.TradeAPI.QueryOrder().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test TradeAPIService QueryOrder Server Error", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer func() {
			close(conn.Done)
			cleanup()
		}()
		conn.Id = "123"

		conn.Listen()
		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}
		done := make(chan struct{})

		go func() {
			respChan, _, err := mockClient.WebsocketAPI.TradeAPI.QueryOrder().Symbol("symbol_example").ExecuteAsync()
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
