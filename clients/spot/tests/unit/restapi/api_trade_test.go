/*
Binance Spot REST API TEST

Testing TradeAPIService

*/

package binancespotrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancespotrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService DeleteOpenOrders Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSDT","origClientOrderId":"E6APeyTJvkMvLMYMqu1KQ4","orderId":11,"orderListId":-1,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","transactTime":1684804350068,"price":"0.089853","origQty":"0.178622","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY","selfTradePreventionMode":"NONE"},{"symbol":"BTCUSDT","origClientOrderId":"A3EF2HCwxgZPFMrfwbgrhv","orderId":13,"orderListId":-1,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","transactTime":1684804350069,"price":"0.090430","origQty":"0.178622","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY","selfTradePreventionMode":"NONE"},{"orderListId":1929,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"2inzWQdDvZLHbbAmAozX2N","transactionTime":1585230948299,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":20,"clientOrderId":"CwOOIPHSmYywx6jZX77TdL"},{"symbol":"BTCUSDT","orderId":21,"clientOrderId":"461cPg51vQjV3zIMOXNz39"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"CwOOIPHSmYywx6jZX77TdL","orderId":20,"orderListId":1929,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","transactTime":1688005070874,"price":"0.668611","origQty":"0.690354","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","stopPrice":"0.378131","icebergQty":"0.017083","selfTradePreventionMode":"NONE"},{"symbol":"BTCUSDT","origClientOrderId":"461cPg51vQjV3zIMOXNz39","orderId":21,"orderListId":1929,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","transactTime":1688005070874,"price":"0.008791","origQty":"0.690354","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT_MAKER","side":"BUY","icebergQty":"0.639962","selfTradePreventionMode":"NONE"}]}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/openOrders", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DeleteOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOpenOrders(context.Background()).Symbol("BNBUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DeleteOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DeleteOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService DeleteOpenOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteOrder Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"LTCBTC","origClientOrderId":"myOrder1","orderId":4,"orderListId":-1,"clientOrderId":"cancelMyOrder1","transactTime":1684804350068,"price":"2.00000000","origQty":"1.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY","selfTradePreventionMode":"NONE"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DeleteOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrder(context.Background()).Symbol("BNBUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DeleteOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DeleteOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService DeleteOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteOrderList Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":0,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"C3wyj4WVEktd7u9aVBRXcN","transactionTime":1574040868128,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":3,"clientOrderId":"TXOvglzXuaubXAaENpaRCB"},{"symbol":"LTCBTC","orderId":2,"clientOrderId":"pO9ufTiFGg3nw2fOdgeOXa"}],"orderReports":[{"symbol":"LTCBTC","origClientOrderId":"TXOvglzXuaubXAaENpaRCB","orderId":3,"orderListId":0,"clientOrderId":"unfWT8ig8i0uj6lPuYLez6","transactTime":1688005070874,"price":"3.00000000","origQty":"10.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT_MAKER","side":"SELL","selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","origClientOrderId":"pO9ufTiFGg3nw2fOdgeOXa","orderId":2,"orderListId":0,"clientOrderId":"unfWT8ig8i0uj6lPuYLez6","transactTime":1688005070874,"price":"1.00000000","origQty":"10.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","stopPrice":"1.00000000","selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DeleteOrderListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrderList(context.Background()).Symbol("BNBUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DeleteOrderListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DeleteOrderListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService DeleteOrderList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrderList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteOrderList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.DeleteOrderList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService NewOrder Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSDT","orderId":28,"orderListId":-1,"clientOrderId":"6gCrw2kRUAF9CvJDGP16IP","transactTime":1507725176595,"price":"0.00000000","origQty":"10.00000000","executedQty":"10.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"10.00000000","status":"FILLED","timeInForce":"GTC","type":"MARKET","side":"SELL","workingTime":1507725176595,"selfTradePreventionMode":"NONE","fills":[{"price":"3995.00000000","qty":"1.00000000","commission":"3.99500000","commissionAsset":"USDT","tradeId":60},{"price":"3997.00000000","qty":"1.00000000","commission":"3.99700000","commissionAsset":"USDT","tradeId":59},{"price":"3998.00000000","qty":"2.00000000","commission":"7.99600000","commissionAsset":"USDT","tradeId":58},{"price":"3999.00000000","qty":"5.00000000","commission":"19.99500000","commissionAsset":"USDT","tradeId":57},{"price":"4000.00000000","qty":"1.00000000","commission":"4.00000000","commissionAsset":"USDT","tradeId":56},{"price":"3995.00000000","qty":"1.00000000","commission":"3.99500000","commissionAsset":"USDT","tradeId":60},{"price":"3997.00000000","qty":"1.00000000","commission":"3.99700000","commissionAsset":"USDT","tradeId":59},{"price":"3998.00000000","qty":"2.00000000","commission":"7.99600000","commissionAsset":"USDT","tradeId":58},{"price":"3999.00000000","qty":"5.00000000","commission":"19.99500000","commissionAsset":"USDT","tradeId":57},{"price":"4000.00000000","qty":"1.00000000","commission":"4.00000000","commissionAsset":"USDT","tradeId":56}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.NewOrderTypeParameterMarket), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.NewOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Type(models.NewOrderTypeParameterMarket).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.NewOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.NewOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService NewOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService NewOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Success", func(t *testing.T) {

		mockedJSON := `{"transactTime":1741669661670,"executionId":22,"amendedOrder":{"symbol":"BTCUSDT","orderId":9,"orderListId":1,"origClientOrderId":"W0fJ9fiLKHOJutovPK3oJp","clientOrderId":"UQ1Np3bmQ71jJzsSDW9Vpi","price":"0.00000000","qty":"4.00000000","executedQty":"0.00000000","preventedQty":"0.00000000","quoteOrderQty":"0.00000000","cumulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"MARKET","side":"BUY","workingTime":1741926410242,"selfTradePreventionMode":"NONE"},"listStatus":{"orderListId":1,"contingencyType":"OTO","listOrderStatus":"EXECUTING","listClientOrderId":"AT7FTxZXylVSwRoZs52mt3","symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":9,"clientOrderId":"UQ1Np3bmQ71jJzsSDW9Vpi"},{"symbol":"BTCUSDT","orderId":8,"clientOrderId":"GkwwHZUUbFtZOoH1YsZk9Q"},{"symbol":"BTCUSDT","orderId":9,"clientOrderId":"UQ1Np3bmQ71jJzsSDW9Vpi"},{"symbol":"BTCUSDT","orderId":8,"clientOrderId":"GkwwHZUUbFtZOoH1YsZk9Q"}]}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order/amend/keepPriority", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "1", r.URL.Query().Get("newQty"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderAmendKeepPriorityResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderAmendKeepPriority(context.Background()).Symbol("BNBUSDT").NewQty(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderAmendKeepPriorityResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderAmendKeepPriorityResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderAmendKeepPriority(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderAmendKeepPriority Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderAmendKeepPriority(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Success", func(t *testing.T) {

		mockedJSON := `{"cancelResult":"SUCCESS","newOrderResult":"SUCCESS","cancelResponse":{"symbol":"BTCUSDT","origClientOrderId":"DnLo3vTAQcjha43lAZhZ0y","orderId":9,"orderListId":-1,"clientOrderId":"osxN3JXAtJvKvCqGeMWMVR","transactTime":1684804350068,"price":"0.01000000","origQty":"0.000100","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE"},"newOrderResponse":{"symbol":"BTCUSDT","orderId":10,"orderListId":-1,"clientOrderId":"wOceeeOzNORyLiQfw7jd8S","transactTime":1652928801803,"price":"0.02000000","origQty":"0.040000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1669277163808,"fills":[],"selfTradePreventionMode":"NONE"},"code":-2021,"msg":"Order cancel-replace partially failed.","data":{"cancelResult":"SUCCESS","newOrderResult":"FAILURE","cancelResponse":{"code":-2011,"msg":"Unknown order sent.","symbol":"LTCBNB","origClientOrderId":"GKt5zzfOxRDSQLveDYCTkc","orderId":64,"orderListId":-1,"clientOrderId":"loehOJF3FjoreUBDmv739R","transactTime":1715779007228,"price":"1.00","origQty":"10.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE"},"newOrderResponse":{"code":-1015,"msg":"Too many new orders; current limit is 1 orders per 10 SECOND.","symbol":"BTCUSDT","orderId":11,"orderListId":-1,"clientOrderId":"pfojJMg6IMNDKuJqDxvoxN","transactTime":1648540168818}}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order/cancelReplace", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.NewOrderTypeParameterMarket), r.URL.Query().Get("type"))
			require.Equal(t, string(models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure), r.URL.Query().Get("cancelReplaceMode"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderCancelReplaceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderCancelReplace(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Type(models.NewOrderTypeParameterMarket).CancelReplaceMode(models.OrderCancelReplaceCancelReplaceModeParameterStopOnFailure).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderCancelReplaceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderCancelReplaceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderCancelReplace(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderCancelReplace Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderCancelReplace(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOco Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":1,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"lH1YDkuQKWiXVXHPSKYEIp","transactionTime":1710485608839,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":11,"clientOrderId":"NuMp0nVYnciDiFmVqfpBqK"},{"symbol":"LTCBTC","orderId":10,"clientOrderId":"44nZvqpemY7sVYgPYbvPih"}],"orderReports":[{"symbol":"LTCBTC","orderId":11,"orderListId":1,"clientOrderId":"NuMp0nVYnciDiFmVqfpBqK","transactTime":1710485608839,"price":"3.00000000","origQty":"5.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT_MAKER","side":"SELL","workingTime":1710485608839,"selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","orderId":10,"orderListId":1,"clientOrderId":"44nZvqpemY7sVYgPYbvPih","transactTime":1710485608839,"price":"1.00000000","origQty":"5.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","stopPrice":"1.00000000","workingTime":-1,"icebergQty":"1.00000000","selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList/oco", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			require.Equal(t, string(models.OrderListOcoAboveTypeParameterStopLossLimit), r.URL.Query().Get("aboveType"))
			require.Equal(t, string(models.OrderListOcoBelowTypeParameterStopLoss), r.URL.Query().Get("belowType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderListOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOco(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Quantity(float32(1.0)).AboveType(models.OrderListOcoAboveTypeParameterStopLossLimit).BelowType(models.OrderListOcoBelowTypeParameterStopLoss).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderListOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderListOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderListOco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOpo Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":0,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"H94qCqO27P74OEiO4X8HOG","transactionTime":1762998011671,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":3,"clientOrderId":"2ZJCY0IjOhuYIMLGN8kU8S"},{"symbol":"BTCUSDT","orderId":2,"clientOrderId":"JX6xfdjo0wysiGumfHNmPu"}],"orderReports":[{"symbol":"BTCUSDT","orderId":3,"orderListId":0,"clientOrderId":"2ZJCY0IjOhuYIMLGN8kU8S","transactTime":1762998011671,"price":"0.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"MARKET","side":"SELL","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"BTCUSDT","orderId":2,"orderListId":0,"clientOrderId":"JX6xfdjo0wysiGumfHNmPu","transactTime":1762998011671,"price":"102264.00000000","origQty":"0.00060000","executedQty":"0.00000000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1762998011671,"selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList/opo", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.OrderListOpoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, "1", r.URL.Query().Get("workingPrice"))
			require.Equal(t, "1", r.URL.Query().Get("workingQuantity"))
			require.Equal(t, string(models.OrderListOpoPendingTypeParameterLimit), r.URL.Query().Get("pendingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderListOpoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpo(context.Background()).Symbol("BNBUSDT").WorkingType(models.OrderListOpoWorkingTypeParameterLimit).WorkingSide(models.NewOrderSideParameterBuy).WorkingPrice(float32(1.0)).WorkingQuantity(float32(1.0)).PendingType(models.OrderListOpoPendingTypeParameterLimit).PendingSide(models.NewOrderSideParameterBuy).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderListOpoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderListOpoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderListOpo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOpo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOpoco Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":2,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"bcedxMpQG6nFrZUPQyshoL","transactionTime":1763000506354,"symbol":"BTCUSDT","orders":[{"symbol":"BTCUSDT","orderId":11,"clientOrderId":"yINkaXSJeoi3bU5vWMY8Z8"},{"symbol":"BTCUSDT","orderId":10,"clientOrderId":"mfif39yPTHsB3C0FIXznR2"},{"symbol":"BTCUSDT","orderId":9,"clientOrderId":"OLSBhMWaIlLSzZ9Zm7fnKB"}],"orderReports":[{"symbol":"BTCUSDT","orderId":11,"orderListId":2,"clientOrderId":"yINkaXSJeoi3bU5vWMY8Z8","transactTime":1763000506354,"price":"104261.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"LIMIT_MAKER","side":"SELL","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"BTCUSDT","orderId":10,"orderListId":2,"clientOrderId":"mfif39yPTHsB3C0FIXznR2","transactTime":1763000506354,"price":"101613.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","stopPrice":"10100.00000000","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"BTCUSDT","orderId":9,"orderListId":2,"clientOrderId":"OLSBhMWaIlLSzZ9Zm7fnKB","transactTime":1763000506354,"price":"102496.00000000","origQty":"0.00170000","executedQty":"0.00000000","origQuoteOrderQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1763000506354,"selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList/opoco", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.OrderListOpoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, "1", r.URL.Query().Get("workingPrice"))
			require.Equal(t, "1", r.URL.Query().Get("workingQuantity"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			require.Equal(t, string(models.OrderListOcoAboveTypeParameterStopLossLimit), r.URL.Query().Get("pendingAboveType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderListOpocoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpoco(context.Background()).Symbol("BNBUSDT").WorkingType(models.OrderListOpoWorkingTypeParameterLimit).WorkingSide(models.NewOrderSideParameterBuy).WorkingPrice(float32(1.0)).WorkingQuantity(float32(1.0)).PendingSide(models.NewOrderSideParameterBuy).PendingAboveType(models.OrderListOcoAboveTypeParameterStopLossLimit).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderListOpocoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderListOpocoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderListOpoco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOpoco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOpoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOto Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":0,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"yl2ERtcar1o25zcWtqVBTC","transactionTime":1712289389158,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":5,"clientOrderId":"arLFo0zGJVDE69cvGBaU0d"},{"symbol":"LTCBTC","orderId":4,"clientOrderId":"Bq17mn9fP6vyCn75Jw1xya"}],"orderReports":[{"symbol":"LTCBTC","orderId":5,"orderListId":0,"clientOrderId":"arLFo0zGJVDE69cvGBaU0d","transactTime":1712289389158,"price":"0.00000000","origQty":"5.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"MARKET","side":"BUY","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","orderId":4,"orderListId":0,"clientOrderId":"Bq17mn9fP6vyCn75Jw1xya","transactTime":1712289389158,"price":"1.00000000","origQty":"1.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1712289389158,"selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList/oto", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.OrderListOpoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, "1", r.URL.Query().Get("workingPrice"))
			require.Equal(t, "1", r.URL.Query().Get("workingQuantity"))
			require.Equal(t, string(models.OrderListOpoPendingTypeParameterLimit), r.URL.Query().Get("pendingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			require.Equal(t, "1", r.URL.Query().Get("pendingQuantity"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderListOtoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOto(context.Background()).Symbol("BNBUSDT").WorkingType(models.OrderListOpoWorkingTypeParameterLimit).WorkingSide(models.NewOrderSideParameterBuy).WorkingPrice(float32(1.0)).WorkingQuantity(float32(1.0)).PendingType(models.OrderListOpoPendingTypeParameterLimit).PendingSide(models.NewOrderSideParameterBuy).PendingQuantity(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderListOtoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderListOtoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderListOto Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOto(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOto Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOto(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOtoco Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":1,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"RumwQpBaDctlUu5jyG5rs0","transactionTime":1712291372842,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":8,"clientOrderId":"r4JMv9cwAYYUwwBZfbussx"},{"symbol":"LTCBTC","orderId":7,"clientOrderId":"6pcQbFIzTXGZQ1e2MkGDq4"},{"symbol":"LTCBTC","orderId":6,"clientOrderId":"fM9Y4m23IFJVCQmIrlUmMK"}],"orderReports":[{"symbol":"LTCBTC","orderId":8,"orderListId":1,"clientOrderId":"r4JMv9cwAYYUwwBZfbussx","transactTime":1712291372842,"price":"3.00000000","origQty":"5.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"GTC","type":"LIMIT_MAKER","side":"BUY","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","orderId":7,"orderListId":1,"clientOrderId":"6pcQbFIzTXGZQ1e2MkGDq4","transactTime":1712291372842,"price":"1.00000000","origQty":"5.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"PENDING_NEW","timeInForce":"IOC","type":"STOP_LOSS_LIMIT","side":"BUY","stopPrice":"6.00000000","workingTime":-1,"selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","orderId":6,"orderListId":1,"clientOrderId":"fM9Y4m23IFJVCQmIrlUmMK","transactTime":1712291372842,"price":"1.00000000","origQty":"1.00000000","executedQty":"0.00000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.00000000","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","workingTime":1712291372842,"selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList/otoco", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.OrderListOpoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, "1", r.URL.Query().Get("workingPrice"))
			require.Equal(t, "1", r.URL.Query().Get("workingQuantity"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			require.Equal(t, "1", r.URL.Query().Get("pendingQuantity"))
			require.Equal(t, string(models.OrderListOcoAboveTypeParameterStopLossLimit), r.URL.Query().Get("pendingAboveType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderListOtocoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOtoco(context.Background()).Symbol("BNBUSDT").WorkingType(models.OrderListOpoWorkingTypeParameterLimit).WorkingSide(models.NewOrderSideParameterBuy).WorkingPrice(float32(1.0)).WorkingQuantity(float32(1.0)).PendingSide(models.NewOrderSideParameterBuy).PendingQuantity(float32(1.0)).PendingAboveType(models.OrderListOcoAboveTypeParameterStopLossLimit).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderListOtocoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderListOtocoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderListOtoco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOtoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderListOtoco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderListOtoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderOco Success", func(t *testing.T) {

		mockedJSON := `{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"JYVpp3F0f5CAG15DhtrqLp","transactionTime":1563417480525,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":3,"clientOrderId":"xTXKaGYd4bluPVp78IVRvl"},{"symbol":"LTCBTC","orderId":2,"clientOrderId":"Kk7sqHb9J6mJWTMDVW7Vos"}],"orderReports":[{"symbol":"LTCBTC","orderId":3,"orderListId":0,"clientOrderId":"xTXKaGYd4bluPVp78IVRvl","transactTime":1563417480525,"price":"0.036435","origQty":"0.624363","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"NEW","timeInForce":"GTC","type":"LIMIT_MAKER","side":"BUY","workingTime":1563417480525,"selfTradePreventionMode":"NONE"},{"symbol":"LTCBTC","orderId":2,"orderListId":0,"clientOrderId":"Kk7sqHb9J6mJWTMDVW7Vos","transactTime":1563417480525,"price":"0.000000","origQty":"0.624363","executedQty":"0.000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"NEW","timeInForce":"GTC","type":"STOP_LOSS","side":"BUY","stopPrice":"0.960664","workingTime":-1,"selfTradePreventionMode":"NONE"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order/oco", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			require.Equal(t, "1", r.URL.Query().Get("price"))
			require.Equal(t, "1", r.URL.Query().Get("stopPrice"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderOco(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Quantity(float32(1.0)).Price(float32(1.0)).StopPrice(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderOco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderTest Success", func(t *testing.T) {

		mockedJSON := `{"standardCommissionForOrder":{"maker":"0.00000112","taker":"0.00000114"},"specialCommissionForOrder":{"maker":"0.05000000","taker":"0.06000000"},"taxCommissionForOrder":{"maker":"0.00000112","taker":"0.00000114"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.25000000"}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order/test", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.NewOrderTypeParameterMarket), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderTestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderTest(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Type(models.NewOrderTypeParameterMarket).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderTestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderTestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService OrderTest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderTest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService OrderTest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.OrderTest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService SorOrder Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSDT","orderId":2,"orderListId":-1,"clientOrderId":"sBI1KM6nNtOfj5tccZSKly","transactTime":1689149087774,"price":"31000.00000000","origQty":"0.50000000","executedQty":"0.50000000","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"14000.00000000","status":"FILLED","timeInForce":"GTC","type":"LIMIT","side":"BUY","workingTime":1689149087774,"fills":[{"matchType":"ONE_PARTY_TRADE_REPORT","price":"28000.00000000","qty":"0.50000000","commission":"0.00000000","commissionAsset":"BTC","tradeId":-1,"allocId":0}],"workingFloor":"SOR","selfTradePreventionMode":"NONE","usedSor":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/sor/order", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.NewOrderTypeParameterMarket), r.URL.Query().Get("type"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SorOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrder(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Type(models.NewOrderTypeParameterMarket).Quantity(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SorOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SorOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService SorOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService SorOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService SorOrderTest Success", func(t *testing.T) {

		mockedJSON := `{"standardCommissionForOrder":{"maker":"0.00000112","taker":"0.00000114"},"taxCommissionForOrder":{"maker":"0.00000112","taker":"0.00000114"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.25000000"}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/sor/order/test", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.NewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.NewOrderTypeParameterMarket), r.URL.Query().Get("type"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SorOrderTestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrderTest(context.Background()).Symbol("BNBUSDT").Side(models.NewOrderSideParameterBuy).Type(models.NewOrderTypeParameterMarket).Quantity(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SorOrderTestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SorOrderTestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService SorOrderTest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrderTest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService SorOrderTest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.SorOrderTest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
