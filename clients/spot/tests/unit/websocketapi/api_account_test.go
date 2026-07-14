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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","standardCommission":{"maker":"0.00000010","taker":"0.00000020","buyer":"0.00000030","seller":"0.00000040"},"specialCommission":{"maker":"0.01000000","taker":"0.02000000","buyer":"0.03000000","seller":"0.04000000"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.75000000"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","standardCommission":{"maker":"0.00000010","taker":"0.00000020","buyer":"0.00000030","seller":"0.00000040"},"specialCommission":{"maker":"0.01000000","taker":"0.02000000","buyer":"0.03000000","seller":"0.04000000"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.75000000"}},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/account.commission"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/account.rateLimits.orders"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"makerCommission":15,"takerCommission":15,"buyerCommission":0,"sellerCommission":0,"canTrade":true,"canWithdraw":true,"canDeposit":true,"brokered":false,"requireSelfTradePrevention":false,"preventSor":false,"updateTime":1660801833000,"accountType":"SPOT","balances":[{"asset":"BNB"}],"permissions":["SPOT"],"uid":354937868},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"makerCommission":15,"takerCommission":15,"buyerCommission":0,"sellerCommission":0,"canTrade":true,"canWithdraw":true,"canDeposit":true,"brokered":false,"requireSelfTradePrevention":false,"preventSor":false,"updateTime":1660801833000,"accountType":"SPOT","balances":[{"asset":"BNB"}],"permissions":["SPOT"],"uid":354937868},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/account.status"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/allOrderLists"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE","preventedMatchId":0,"icebergQty":"0.00000000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE","preventedMatchId":0,"icebergQty":"0.00000000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/allOrders"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","allocationId":0,"allocationType":"SOR","orderId":500,"orderListId":-1,"commissionAsset":"BTC","time":1687319487614,"isBuyer":false,"isMaker":false,"isAllocator":false}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","allocationId":0,"allocationType":"SOR","orderId":500,"orderListId":-1,"commissionAsset":"BTC","time":1687319487614,"isBuyer":false,"isMaker":false,"isAllocator":false}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/myAllocations"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"exchangeFilters":[{"filterType":"EXCHANGE_MAX_NUM_ORDERS","maxNumOrders":1000}],"symbolFilters":[{"filterType":"PRICE_FILTER","priceExponent":8,"minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"}],"assetFilters":[{"filterType":"MAX_ASSET","qtyExponent":8,"limit":"1000000.00000000","asset":"JPY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"status":200,"result":{"exchangeFilters":[{"filterType":"EXCHANGE_MAX_NUM_ORDERS","maxNumOrders":1000}],"symbolFilters":[{"filterType":"PRICE_FILTER","priceExponent":8,"minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"}],"assetFilters":[{"filterType":"MAX_ASSET","qtyExponent":8,"limit":"1000000.00000000","asset":"JPY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}}`

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
		require.Equal(t, "/myFilters"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","transactTime":1669101687094}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","transactTime":1669101687094}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/myPreventedMatches"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","id":1650422481,"orderId":12569099453,"orderListId":-1,"commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","id":1650422481,"orderId":12569099453,"orderListId":-1,"commissionAsset":"BNB","time":1660801715793,"isBuyer":false,"isMaker":true,"isBestMatch":true}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/myTrades"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":4,"clientOrderId":"CUhLgTXnX5n2c0gWiLpV4d"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":4,"clientOrderId":"CUhLgTXnX5n2c0gWiLpV4d"}]}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/openOrderLists.status"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"PARTIALLY_FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"PARTIALLY_FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"selfTradePreventionMode":"NONE","icebergQty":"0.00000000","preventedMatchId":0,"preventedQuantity":"1.200000","stopPrice":"0.00000000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/openOrders.status"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":23,"executionId":60,"origClientOrderId":"my_pending_order","newClientOrderId":"xbxXh5SSwaHS7oUEOCI88B","time":1741924229819}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":[{"symbol":"BTCUSDT","orderId":23,"executionId":60,"origClientOrderId":"my_pending_order","newClientOrderId":"xbxXh5SSwaHS7oUEOCI88B","time":1741924229819}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.amendments"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"orderListId":1274512,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"08985fedd9ea2cf6b28996","transactionTime":1660801713793,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":12569138901,"clientOrderId":"BqtFCj5odMoWtSqGk2X9tU"}]},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/orderList.status"[1:], sentCheck["method"])

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

		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","trailingDelta":10,"trailingTime":-1,"time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"strategyId":37463720,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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

		var err error
		var sent map[string]interface{}
		err = json.Unmarshal(mockWS.MessagesWritten[0], &sent)
		require.NoError(t, err)

		mockedJSON := `{"id":"123","status":200,"result":{"symbol":"BTCUSDT","orderId":12569099453,"orderListId":-1,"clientOrderId":"4d96324ff9d44481926157","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"SELL","trailingDelta":10,"trailingTime":-1,"time":1660801715639,"updateTime":1660801717945,"isWorking":true,"workingTime":1660801715639,"strategyId":37463720,"strategyType":1000000,"selfTradePreventionMode":"NONE","preventedMatchId":0,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`

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
		require.Equal(t, "/order.status"[1:], sentCheck["method"])

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
