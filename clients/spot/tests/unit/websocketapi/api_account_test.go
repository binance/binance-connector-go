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

func Test_binancespotwebsocketapi_AccountAPIService(t *testing.T) {

	t.Run("Test AccountAPIService AccountCommission AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.AccountCommission().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","standardCommission":{"maker":"0.00000010","taker":"0.00000020","buyer":"0.00000030","seller":"0.00000040"},"specialCommission":{"maker":"0.01000000","taker":"0.02000000","buyer":"0.03000000","seller":"0.04000000"},"taxCommission":{"maker":"0.00000112","taker":"0.00000114","buyer":"0.00000118","seller":"0.00000116"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.75000000"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/account.commission"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AccountCommissionResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService AccountCommission Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AccountCommissionResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.AccountCommission().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.AccountCommissionResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","standardCommission":{"maker":"0.00000010","taker":"0.00000020","buyer":"0.00000030","seller":"0.00000040"},"specialCommission":{"maker":"0.01000000","taker":"0.02000000","buyer":"0.03000000","seller":"0.04000000"},"taxCommission":{"maker":"0.00000112","taker":"0.00000114","buyer":"0.00000118","seller":"0.00000116"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.75000000"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/account.commission"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AccountCommissionResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService AccountCommission Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.AccountCommission().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService AccountCommission Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.AccountCommission().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService AccountRateLimitsOrders AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.AccountRateLimitsOrders().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":0},{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":40}]}`
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
			require.Equal(t, "/account.rateLimits.orders"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AccountRateLimitsOrdersResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService AccountRateLimitsOrders Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AccountRateLimitsOrdersResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.AccountRateLimitsOrders().Execute()
			resultChan <- common.ResultWebsocket[models.AccountRateLimitsOrdersResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":0},{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":40}]}`
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
		require.Equal(t, "/account.rateLimits.orders"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AccountRateLimitsOrdersResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService AccountRateLimitsOrders Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.AccountRateLimitsOrders().ExecuteAsync()
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
	t.Run("Test AccountAPIService AccountStatus AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.AccountStatus().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"makerCommission":15,"takerCommission":15,"buyerCommission":0,"sellerCommission":0,"canTrade":true,"canWithdraw":true,"canDeposit":true,"commissionRates":{"maker":"0.00150000","taker":"0.00150000","buyer":"0.00000000","seller":"0.00000000"},"brokered":false,"requireSelfTradePrevention":false,"preventSor":false,"updateTime":1660801833000,"accountType":"SPOT","balances":[{"asset":"USDT","free":"1021.21000000","locked":"0.00000000"},{"asset":"BTC","free":"1.3447112","locked":"0.08600000"},{"asset":"BNB","free":"0.00000000","locked":"0.00000000"}],"permissions":["SPOT"],"uid":354937868},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/account.status"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AccountStatusResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService AccountStatus Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AccountStatusResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.AccountStatus().Execute()
			resultChan <- common.ResultWebsocket[models.AccountStatusResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"makerCommission":15,"takerCommission":15,"buyerCommission":0,"sellerCommission":0,"canTrade":true,"canWithdraw":true,"canDeposit":true,"commissionRates":{"maker":"0.00150000","taker":"0.00150000","buyer":"0.00000000","seller":"0.00000000"},"brokered":false,"requireSelfTradePrevention":false,"preventSor":false,"updateTime":1660801833000,"accountType":"SPOT","balances":[{"asset":"USDT","free":"1021.21000000","locked":"0.00000000"},{"asset":"BTC","free":"1.3447112","locked":"0.08600000"},{"asset":"BNB","free":"0.00000000","locked":"0.00000000"}],"permissions":["SPOT"],"uid":354937868},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/account.status"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AccountStatusResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService AccountStatus Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.AccountStatus().ExecuteAsync()
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
	t.Run("Test AccountAPIService AllOrderLists AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.AllOrderLists().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138902,"clientOrderId":"jLnZpj5enfMXTuhKB1d0us"},{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/allOrderLists"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AllOrderListsResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService AllOrderLists Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AllOrderListsResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.AllOrderLists().Execute()
			resultChan <- common.ResultWebsocket[models.AllOrderListsResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138902,"clientOrderId":"jLnZpj5enfMXTuhKB1d0us"},{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/allOrderLists"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AllOrderListsResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService AllOrderLists Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.AllOrderLists().ExecuteAsync()
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
	t.Run("Test AccountAPIService AllOrders AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.AllOrders().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00847000","cummulativeQuoteQty":"198.33521500","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"origQuoteOrderQty":"0.00000000","selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/allOrders"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.AllOrdersResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService AllOrders Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.AllOrdersResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.AllOrders().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.AllOrdersResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00847000","cummulativeQuoteQty":"198.33521500","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"origQuoteOrderQty":"0.00000000","selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/allOrders"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.AllOrdersResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService AllOrders Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.AllOrders().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService AllOrders Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.AllOrders().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService MyAllocations AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.MyAllocations().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","allocationId":0,"allocationType":"SOR","orderId":500,"orderListId":-1,"price":"1.00000000","qty":"0.10000000","quoteQty":"0.10000000","commission":"0.00000000","commissionAsset":"BTC","time":1687319487614,"isBuyer":false,"isMaker":false,"isAllocator":false}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/myAllocations"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.MyAllocationsResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService MyAllocations Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.MyAllocationsResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.MyAllocations().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.MyAllocationsResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","allocationId":0,"allocationType":"SOR","orderId":500,"orderListId":-1,"price":"1.00000000","qty":"0.10000000","quoteQty":"0.10000000","commission":"0.00000000","commissionAsset":"BTC","time":1687319487614,"isBuyer":false,"isMaker":false,"isAllocator":false}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/myAllocations"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.MyAllocationsResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService MyAllocations Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.MyAllocations().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService MyAllocations Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.MyAllocations().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService MyFilters AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.MyFilters().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"exchangeFilters":[{"filterType":"EXCHANGE_MAX_NUM_ORDERS","maxNumOrders":1000}],"symbolFilters":[{"filterType":"MAX_NUM_ORDER_LISTS","maxNumOrderLists":20}],"assetFilters":[{"filterType":"MAX_ASSET","asset":"JPY","limit":"1000000.00000000"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000},{"rateLimitType":"RAW_REQUESTS","interval":"MINUTE","intervalNum":5,"limit":61000}]}`
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
			require.Equal(t, "/myFilters"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.MyFiltersResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService MyFilters Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.MyFiltersResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.MyFilters().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.MyFiltersResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"exchangeFilters":[{"filterType":"EXCHANGE_MAX_NUM_ORDERS","maxNumOrders":1000}],"symbolFilters":[{"filterType":"MAX_NUM_ORDER_LISTS","maxNumOrderLists":20}],"assetFilters":[{"filterType":"MAX_ASSET","asset":"JPY","limit":"1000000.00000000"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000},{"rateLimitType":"RAW_REQUESTS","interval":"MINUTE","intervalNum":5,"limit":61000}]}`
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
		require.Equal(t, "/myFilters"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.MyFiltersResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService MyFilters Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.MyFilters().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService MyFilters Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.MyFilters().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService MyPreventedMatches AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.MyPreventedMatches().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","price":"1.100000","makerPreventedQuantity":"1.300000","transactTime":1669101687094}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/myPreventedMatches"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.MyPreventedMatchesResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService MyPreventedMatches Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.MyPreventedMatchesResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.MyPreventedMatches().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.MyPreventedMatchesResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","price":"1.100000","makerPreventedQuantity":"1.300000","transactTime":1669101687094}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/myPreventedMatches"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.MyPreventedMatchesResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService MyPreventedMatches Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.MyPreventedMatches().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService MyPreventedMatches Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.MyPreventedMatches().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService MyTrades AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.MyTrades().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","id":1650422482,"orderId":12569099453,"orderListId":-1,"price":"23416.50000000","qty":"0.00212000","quoteQty":"49.64298000","commission":"0.00000000","commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true},{"symbol":"BTCUSDT","id":1650422481,"orderId":12569099453,"orderListId":-1,"price":"23416.10000000","qty":"0.00635000","quoteQty":"148.69223500","commission":"0.00000000","commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
			require.Equal(t, "/myTrades"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.MyTradesResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService MyTrades Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.MyTradesResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.MyTrades().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.MyTradesResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","id":1650422482,"orderId":12569099453,"orderListId":-1,"price":"23416.50000000","qty":"0.00212000","quoteQty":"49.64298000","commission":"0.00000000","commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true},{"symbol":"BTCUSDT","id":1650422481,"orderId":12569099453,"orderListId":-1,"price":"23416.10000000","qty":"0.00635000","quoteQty":"148.69223500","commission":"0.00000000","commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":20}]}`
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
		require.Equal(t, "/myTrades"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.MyTradesResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService MyTrades Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.MyTrades().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService MyTrades Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.MyTrades().Symbol("BNBUSDT").ExecuteAsync()
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
	t.Run("Test AccountAPIService OpenOrderListsStatus AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.OpenOrderListsStatus().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":5,"clientOrderId":"1ZqG7bBuYwaF4SU8CwnwHm"},{"symbol":"BTCUSDT","orderId":4,"clientOrderId":"CUhLgTXnX5n2c0gWiLpV4d"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":6}]}`
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
			require.Equal(t, "/openOrderLists.status"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OpenOrderListsStatusResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService OpenOrderListsStatus Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OpenOrderListsStatusResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.OpenOrderListsStatus().Execute()
			resultChan <- common.ResultWebsocket[models.OpenOrderListsStatusResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":5,"clientOrderId":"1ZqG7bBuYwaF4SU8CwnwHm"},{"symbol":"BTCUSDT","orderId":4,"clientOrderId":"CUhLgTXnX5n2c0gWiLpV4d"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":6}]}`
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
		require.Equal(t, "/openOrderLists.status"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OpenOrderListsStatusResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService OpenOrderListsStatus Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.OpenOrderListsStatus().ExecuteAsync()
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
	t.Run("Test AccountAPIService OpenOrdersStatus AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.OpenOrdersStatus().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00720000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"172.43931000","status":"PARTIALLY_FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":6}]}`
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
			require.Equal(t, "/openOrders.status"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OpenOrdersStatusResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService OpenOrdersStatus Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OpenOrdersStatusResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.OpenOrdersStatus().Execute()
			resultChan <- common.ResultWebsocket[models.OpenOrdersStatusResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00720000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"172.43931000","status":"PARTIALLY_FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":6}]}`
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
		require.Equal(t, "/openOrders.status"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OpenOrdersStatusResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService OpenOrdersStatus Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.OpenOrdersStatus().ExecuteAsync()
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
	t.Run("Test AccountAPIService OrderAmendments AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.OrderAmendments().Symbol("BNBUSDT").OrderId(int64(1)).ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":23,"executionId":60,"origClientOrderId":"my_pending_order","newClientOrderId":"xbxXh5SSwaHS7oUEOCI88B","origQty":"7.00000000","newQty":"5.00000000","time":1741924229819}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
			require.Equal(t, "/order.amendments"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderAmendmentsResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService OrderAmendments Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderAmendmentsResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.OrderAmendments().Symbol("BNBUSDT").OrderId(int64(1)).Execute()
			resultChan <- common.ResultWebsocket[models.OrderAmendmentsResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":23,"executionId":60,"origClientOrderId":"my_pending_order","newClientOrderId":"xbxXh5SSwaHS7oUEOCI88B","origQty":"7.00000000","newQty":"5.00000000","time":1741924229819}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
		require.Equal(t, "/order.amendments"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderAmendmentsResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService OrderAmendments Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.OrderAmendments().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService OrderAmendments Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.OrderAmendments().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService OrderAmendments Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.OrderAmendments().Symbol("BNBUSDT").OrderId(int64(1)).ExecuteAsync()
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
	t.Run("Test AccountAPIService OrderListStatus AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.OrderListStatus().ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138902,"clientOrderId":"jLnZpj5enfMXTuhKB1d0us"},{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
			require.Equal(t, "/orderList.status"[1:], sent["method"])

			typedResp := resp.Typed
			require.IsType(t, &models.OrderListStatusResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService OrderListStatus Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderListStatusResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.OrderListStatus().Execute()
			resultChan <- common.ResultWebsocket[models.OrderListStatusResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138902,"clientOrderId":"jLnZpj5enfMXTuhKB1d0us"},{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
		require.Equal(t, "/orderList.status"[1:], sent["method"])

		typedResp := resp.Typed
		require.IsType(t, &models.OrderListStatusResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService OrderListStatus Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.OrderListStatus().ExecuteAsync()
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
	t.Run("Test AccountAPIService OrderStatus AsyncExecute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		responseChan, errorChan, err := mockClient.WebsocketAPI.AccountAPI.OrderStatus().Symbol("BNBUSDT").ExecuteAsync()
		require.NoError(t, err)

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00847000","cummulativeQuoteQty":"198.33521500","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"trailingTime":-1,"icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"origQuoteOrderQty":"0.00000000","strategyId":37463720,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
			require.IsType(t, &models.OrderStatusResponse{}, typedResp)
		case err := <-errorChan:
			t.Fatalf("Unexpected error: %v", err)
		}
	})

	t.Run("Test AccountAPIService OrderStatus Execute Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		cfg := common.NewConfigurationWebsocketApi()
		mockClient := client.NewBinanceSpotClient(
			client.WithWebsocketAPI(cfg),
		)
		mockClient.WebsocketAPI.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		conn.Listen()

		resultChan := make(chan common.ResultWebsocket[models.OrderStatusResponse], 1)
		go func() {
			resp, err := mockClient.WebsocketAPI.AccountAPI.OrderStatus().Symbol("BNBUSDT").Execute()
			resultChan <- common.ResultWebsocket[models.OrderStatusResponse]{Value: resp, Err: err}
		}()

		<-mockWS.HasSentChan

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","price":"23416.10000000","origQty":"0.00847000","executedQty":"0.00847000","cummulativeQuoteQty":"198.33521500","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","stopPrice":"0.00000000","trailingDelta":10,"trailingTime":-1,"icebergQty":"0.00000000","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"origQuoteOrderQty":"0.00000000","strategyId":37463720,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`
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
		require.IsType(t, &models.OrderStatusResponse{}, typedResp)
	})

	t.Run("Test AccountAPIService OrderStatus Missing Required Params", func(t *testing.T) {
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

		respChan, errChan, err := mockClient.WebsocketAPI.AccountAPI.OrderStatus().ExecuteAsync()
		require.Error(t, err)
		require.Nil(t, respChan)
		require.Nil(t, errChan)
	})

	t.Run("Test AccountAPIService OrderStatus Server Error", func(t *testing.T) {
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
			respChan, _, err := mockClient.WebsocketAPI.AccountAPI.OrderStatus().Symbol("BNBUSDT").ExecuteAsync()
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
