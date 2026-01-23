/*
Binance Derivatives Trading COIN Futures REST API TEST

Testing TradeAPIService

*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingcoinfuturesrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService AccountTradeList Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_200626","id":6,"orderId":28,"pair":"BTCUSD","side":"SELL","price":"8800","qty":"1","realizedPnl":"0","marginAsset":"BTC","baseQty":"0.01136364","commission":"0.00000454","commissionAsset":"BTC","time":1590743483586,"positionSide":"BOTH","buyer":false,"maker":false}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/userTrades", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountTradeListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AccountTradeList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountTradeListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountTradeListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService AccountTradeList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AccountTradeList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService AllOrders Success", func(t *testing.T) {

		mockedJSON := `[{"avgPrice":"0.0","clientOrderId":"abc","cumBase":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","pair":"BTCUSD","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/allOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AllOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AllOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AllOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AllOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService AllOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AllOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService AutoCancelAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSD_200925","countdownTime":"100000"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/countdownCancelAll", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, "789", r.URL.Query().Get("countdownTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AutoCancelAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AutoCancelAllOpenOrders(context.Background()).Symbol("symbol_example").CountdownTime(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AutoCancelAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService AutoCancelAllOpenOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService AutoCancelAllOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.AutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `{"code":200,"msg":"The operation of cancel all open order is done."}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/allOpenOrders", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelAllOpenOrders(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService CancelAllOpenOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelAllOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelMultipleOrders Success", func(t *testing.T) {

		mockedJSON := `[{"avgPrice":"0.0","clientOrderId":"myOrder1","cumQty":"0","cumBase":"0","executedQty":"0","orderId":283194212,"origQty":"11","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"CANCELED","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE","updateTime":1571110484038},{"code":-2011,"msg":"Unknown order sent."}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/batchOrders", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelMultipleOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelMultipleOrders(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelMultipleOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelMultipleOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService CancelMultipleOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelMultipleOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelOrder Success", func(t *testing.T) {

		mockedJSON := `{"avgPrice":"0.0","clientOrderId":"myOrder1","cumQty":"0","cumBase":"0","executedQty":"0","orderId":283194212,"origQty":"11","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"CANCELED","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","pair":"BTCUSD","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1571110484038,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/order", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelOrder(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService CancelOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CancelOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CancelOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangeInitialLeverage Success", func(t *testing.T) {

		mockedJSON := `{"leverage":21,"maxQty":"1000","symbol":"BTCUSD_200925"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/leverage", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, "789", r.URL.Query().Get("leverage"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ChangeInitialLeverageResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeInitialLeverage(context.Background()).Symbol("symbol_example").Leverage(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ChangeInitialLeverageResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ChangeInitialLeverageResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ChangeInitialLeverage Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeInitialLeverage(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangeInitialLeverage Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeInitialLeverage(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangeMarginType Success", func(t *testing.T) {

		mockedJSON := `{"code":200,"msg":"success"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/marginType", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ChangeMarginTypeMarginTypeParameterIsolated), r.URL.Query().Get("marginType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ChangeMarginTypeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeMarginType(context.Background()).Symbol("symbol_example").MarginType(models.ChangeMarginTypeMarginTypeParameterIsolated).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ChangeMarginTypeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ChangeMarginTypeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ChangeMarginType Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeMarginType(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangeMarginType Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangeMarginType(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangePositionMode Success", func(t *testing.T) {

		mockedJSON := `{"code":200,"msg":"success"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/positionSide/dual", r.URL.Path)
			require.Equal(t, "dualSidePosition_example", r.URL.Query().Get("dualSidePosition"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ChangePositionModeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangePositionMode(context.Background()).DualSidePosition("dualSidePosition_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ChangePositionModeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ChangePositionModeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ChangePositionMode Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangePositionMode(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ChangePositionMode Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ChangePositionMode(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CurrentAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `[{"avgPrice":"0.0","clientOrderId":"abc","cumBase":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/openOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CurrentAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CurrentAllOpenOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CurrentAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CurrentAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService CurrentAllOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CurrentAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetOrderModifyHistory Success", func(t *testing.T) {

		mockedJSON := `[{"amendmentId":5363,"symbol":"BTCUSD_PERP","pair":"BTCUSD","orderId":20072994037,"clientOrderId":"LJ9R4QZDihCaS8UAOOLpgW","time":1629184560899,"amendment":{"price":{"before":"30004","after":"30003.2"},"origQty":{"before":"1","after":"1"},"count":3}},{"amendmentId":5361,"symbol":"BTCUSD_PERP","pair":"BTCUSD","orderId":20072994037,"clientOrderId":"LJ9R4QZDihCaS8UAOOLpgW","time":1629184533946,"amendment":{"price":{"before":"30005","after":"30004"},"origQty":{"before":"1","after":"1"},"count":2}},{"amendmentId":5325,"symbol":"BTCUSD_PERP","pair":"BTCUSD","orderId":20072994037,"clientOrderId":"LJ9R4QZDihCaS8UAOOLpgW","time":1629182711787,"amendment":{"price":{"before":"30002","after":"30005"},"origQty":{"before":"1","after":"1"},"count":1}}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/orderAmendment", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOrderModifyHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetOrderModifyHistory(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOrderModifyHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOrderModifyHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetOrderModifyHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetOrderModifyHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetOrderModifyHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetOrderModifyHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetPositionMarginChangeHistory Success", func(t *testing.T) {

		mockedJSON := `[{"amount":"23.36332311","asset":"BTC","symbol":"BTCUSD_200925","time":1578047897183,"type":1,"positionSide":"BOTH"},{"amount":"100","asset":"BTC","symbol":"BTCUSD_200925","time":1578047900425,"type":1,"positionSide":"LONG"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/positionMargin/history", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetPositionMarginChangeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetPositionMarginChangeHistory(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetPositionMarginChangeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetPositionMarginChangeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetPositionMarginChangeHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetPositionMarginChangeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetPositionMarginChangeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetPositionMarginChangeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyIsolatedPositionMargin Success", func(t *testing.T) {

		mockedJSON := `{"amount":100,"code":200,"msg":"Successfully modify position margin.","type":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/positionMargin", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			require.Equal(t, string(models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ModifyIsolatedPositionMarginResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyIsolatedPositionMargin(context.Background()).Symbol("symbol_example").Amount(float32(1.0)).Type(models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ModifyIsolatedPositionMarginResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ModifyIsolatedPositionMarginResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ModifyIsolatedPositionMargin Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyIsolatedPositionMargin(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyIsolatedPositionMargin Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyIsolatedPositionMargin(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyMultipleOrders Success", func(t *testing.T) {

		mockedJSON := `[{"orderId":20072994037,"symbol":"BTCUSD_PERP","pair":"BTCUSD","status":"NEW","clientOrderId":"LJ9R4QZDihCaS8UAOOLpgW","price":"30005","avgPrice":"0.0","origQty":"1","executedQty":"0","cumQty":"0","cumBase":"0","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"BUY","positionSide":"LONG","stopPrice":"0","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","updateTime":1629182711600},{"code":-2022,"msg":"ReduceOnly Order is rejected."}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/batchOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ModifyMultipleOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyMultipleOrders(context.Background()).BatchOrders([]models.ModifyMultipleOrdersBatchOrdersParameterInner{*models.NewModifyMultipleOrdersBatchOrdersParameterInner()}).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ModifyMultipleOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ModifyMultipleOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ModifyMultipleOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyMultipleOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyOrder Success", func(t *testing.T) {

		mockedJSON := `{"orderId":20072994037,"symbol":"BTCUSD_PERP","pair":"BTCUSD","status":"NEW","clientOrderId":"LJ9R4QZDihCaS8UAOOLpgW","price":"30005","avgPrice":"0.0","origQty":"1","executedQty":"0","cumQty":"0","cumBase":"0","timeInForce":"GTC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"BUY","positionSide":"LONG","stopPrice":"0","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","priceMatch":"NONE","selfTradePreventionMode":"NONE","updateTime":1629182711600}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/order", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy), r.URL.Query().Get("side"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ModifyOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyOrder(context.Background()).Symbol("symbol_example").Side(models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ModifyOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ModifyOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ModifyOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService ModifyOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ModifyOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService NewOrder Success", func(t *testing.T) {

		mockedJSON := `{"clientOrderId":"testOrder","cumQty":"0","cumBase":"0","executedQty":"0","orderId":22542179,"avgPrice":"0.0","origQty":"10","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","pair":"BTCUSD","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","origType":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1566818724722,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/order", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.NewOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Symbol("symbol_example").Side(models.ModifyMultipleOrdersBatchOrdersParameterInnerSideBuy).Type(models.PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.NewOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PlaceMultipleOrders Success", func(t *testing.T) {

		mockedJSON := `[{"clientOrderId":"testOrder","cumQty":"0","cumBase":"0","executedQty":"0","orderId":22542179,"avgPrice":"0.0","origQty":"10","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","symbol":"BTCUSD_200925","pair":"BTCUSD","timeInForce":"GTC","type":"TRAILING_STOP_MARKET","origType":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1566818724722,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"},{"code":-2022,"msg":"ReduceOnly Order is rejected."}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/batchOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.PlaceMultipleOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceMultipleOrders(context.Background()).BatchOrders([]models.PlaceMultipleOrdersBatchOrdersParameterInner{*models.NewPlaceMultipleOrdersBatchOrdersParameterInner()}).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.PlaceMultipleOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.PlaceMultipleOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService PlaceMultipleOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PlaceMultipleOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceMultipleOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PositionAdlQuantileEstimation Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_200925","adlQuantile":{"LONG":3,"SHORT":3,"HEDGE":0}},{"symbol":"BTCUSD_201225","adlQuantile":{"LONG":1,"SHORT":2,"BOTH":0}}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/adlQuantile", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.PositionAdlQuantileEstimationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PositionAdlQuantileEstimation(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.PositionAdlQuantileEstimationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.PositionAdlQuantileEstimationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService PositionAdlQuantileEstimation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PositionAdlQuantileEstimation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PositionInformation Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_201225","positionAmt":"0","entryPrice":"0.0","breakEvenPrice":"0.0","markPrice":"0.00000000","unRealizedProfit":"0.00000000","liquidationPrice":"0","leverage":"125","maxQty":"50","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"BOTH","updateTime":0},{"symbol":"BTCUSD_201225","positionAmt":"1","entryPrice":"11707.70000003","breakEvenPrice":"11707.80000005","markPrice":"11788.66626667","unRealizedProfit":"0.00005866","liquidationPrice":"11667.63509587","leverage":"125","maxQty":"50","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"LONG","updateTime":1627026881327},{"symbol":"BTCUSD_201225","positionAmt":"0","entryPrice":"0.0","breakEvenPrice":"0.0","markPrice":"0.00000000","unRealizedProfit":"0.00000000","liquidationPrice":"0","leverage":"125","maxQty":"50","marginType":"cross","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"SHORT","updateTime":1627026881327}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/positionRisk", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.PositionInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PositionInformation(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.PositionInformationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.PositionInformationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService PositionInformation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PositionInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryCurrentOpenOrder Success", func(t *testing.T) {

		mockedJSON := `{"avgPrice":"0.0","clientOrderId":"abc","cumBase":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","pair":"BTCUSD","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/openOrder", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryCurrentOpenOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryCurrentOpenOrder(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryCurrentOpenOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryCurrentOpenOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryCurrentOpenOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryCurrentOpenOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryCurrentOpenOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryCurrentOpenOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryOrder Success", func(t *testing.T) {

		mockedJSON := `{"avgPrice":"0.0","clientOrderId":"abc","cumBase":"0","executedQty":"0","orderId":1917641,"origQty":"0.40","origType":"TRAILING_STOP_MARKET","price":"0","reduceOnly":false,"side":"BUY","positionSide":"SHORT","status":"NEW","stopPrice":"9300","closePosition":false,"symbol":"BTCUSD_200925","pair":"BTCUSD","time":1579276756075,"timeInForce":"GTC","type":"TRAILING_STOP_MARKET","activatePrice":"9020","priceRate":"0.3","updateTime":1579276756075,"workingType":"CONTRACT_PRICE","priceProtect":false,"priceMatch":"NONE","selfTradePreventionMode":"NONE"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/order", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrder(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService UsersForceOrders Success", func(t *testing.T) {

		mockedJSON := `[{"orderId":165123080,"symbol":"BTCUSD_200925","pair":"BTCUSD","status":"FILLED","clientOrderId":"autoclose-1596542005017000006","price":"11326.9","avgPrice":"11326.9","origQty":"1","executedQty":"1","cumBase":"0.00882854","timeInForce":"IOC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"SELL","positionSide":"BOTH","stopPrice":"0","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","time":1596542005019,"updateTime":1596542005050},{"orderId":207251986,"symbol":"BTCUSD_200925","pair":"BTCUSD","status":"FILLED","clientOrderId":"autoclose-1597307316020000006","price":"11619.4","avgPrice":"11661.2","origQty":"1","executedQty":"1","cumBase":"0.00857544","timeInForce":"IOC","type":"LIMIT","reduceOnly":false,"closePosition":false,"side":"SELL","positionSide":"LONG","stopPrice":"0","workingType":"CONTRACT_PRICE","priceProtect":false,"origType":"LIMIT","time":1597307316022,"updateTime":1597307316035}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/forceOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.UsersForceOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.UsersForceOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.UsersForceOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.UsersForceOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService UsersForceOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.UsersForceOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
