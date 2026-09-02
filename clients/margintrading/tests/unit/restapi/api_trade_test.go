/*
Margin REST API TEST

Testing TradeAPIService

*/

package binancemargintradingrestapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancemargintradingrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService CreateSpecialKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"apiKey":"npOzOAeLVgr2TuxWfNo43AaPWpBbJEoKezh1o8mSQb6ryE2odE11A4AoVlJbQoGx","secretKey":"87ssWB7azoy6ACRfyp6OVOL5U3rtZptX31QWw2kWjl1jHEYRbyM1pd6qykRBQw8p","type":"HMAC_SHA256"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/apiKey", r.URL.Path)
			require.Equal(t, "apiName", r.URL.Query().Get("apiName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateSpecialKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CreateSpecialKey(context.Background()).ApiName("apiName").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateSpecialKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateSpecialKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService CreateSpecialKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CreateSpecialKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService CreateSpecialKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.CreateSpecialKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService DeleteSpecialKey Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/apiKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.DeleteSpecialKey(context.Background()).Execute()
		require.NoError(t, err)
	})

	t.Run("Test TradeAPIService DeleteSpecialKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.DeleteSpecialKey(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test TradeAPIService EditIpForSpecialKey Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/apiKey/ip", r.URL.Path)
			require.Equal(t, "24.156.99.202", r.URL.Query().Get("ip"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Ip("24.156.99.202").Execute()
		require.NoError(t, err)
	})

	t.Run("Test TradeAPIService EditIpForSpecialKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test TradeAPIService EditIpForSpecialKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.EditIpForSpecialKey(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test TradeAPIService ExitSpecialKeyMode Success", func(t *testing.T) {

		var mockedJSON string
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/exit-special-key-mode", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected map[string]interface{}
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ExitSpecialKeyMode(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[map[string]interface{}]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(map[string]interface{}{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService ExitSpecialKeyMode Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.ExitSpecialKeyMode(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetForceLiquidationRecord Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"rows":[{"avgPrice":"0.00388359","executedQty":"31.39000000","orderId":180015097,"price":"0.00388110","qty":"31.39000000","side":"SELL","symbol":"BNBBTC","timeInForce":"GTC","isIsolated":true,"updatedTime":1558941374745}],"total":1}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/forceLiquidationRec", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetForceLiquidationRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetForceLiquidationRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetForceLiquidationRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetForceLiquidationRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetForceLiquidationRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetForceLiquidationRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetSmallLiabilityExchangeCoinList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"asset":"ETH","interest":"0.00083334","principal":"0.001","liabilityAsset":"USDT","liabilityQty":0.3552}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/exchange-small-liability", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSmallLiabilityExchangeCoinListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeCoinList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSmallLiabilityExchangeCoinListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSmallLiabilityExchangeCoinListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetSmallLiabilityExchangeCoinList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeCoinList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetSmallLiabilityExchangeHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"total":1,"rows":[{"asset":"ETH","amount":"0.00083434","targetAsset":"BUSD","targetAmount":"1.37576819","bizType":"EXCHANGE_SMALL_LIABILITY","timestamp":1672801339253}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/exchange-small-liability-history", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("current"))
			require.Equal(t, "10", r.URL.Query().Get("size"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSmallLiabilityExchangeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeHistory(context.Background()).Current(int64(1)).Size(int64(10)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSmallLiabilityExchangeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSmallLiabilityExchangeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetSmallLiabilityExchangeHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetSmallLiabilityExchangeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetSmallLiabilityExchangeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService LiquidationLoanRepay Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"repayId":12345678,"asset":"USDT","amount":"300.00000000","status":"SUCCESS","createTime":1714492800000}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/liquidation-loan/repay", r.URL.Path)
			require.Equal(t, "USDT", r.URL.Query().Get("asset"))
			require.Equal(t, fmt.Sprintf("%v", float64(300.00)), r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.LiquidationLoanRepayResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.LiquidationLoanRepay(context.Background()).Asset("USDT").Amount(float64(300.00)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.LiquidationLoanRepayResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.LiquidationLoanRepayResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService LiquidationLoanRepay Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.LiquidationLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService LiquidationLoanRepay Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.LiquidationLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelAllOpenOrdersOnASymbol Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","isIsolated":true,"origClientOrderId":"E6APeyTJvkMvLMYMqu1KQ4","orderId":11,"orderListId":-1,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","price":"0.089853","origQty":"0.178622","executedQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY","selfTradePreventionMode":"NONE","contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"2inzWQdDvZLHbbAmAozX2N","transactionTime":1585230948299,"orders":[{"symbol":"BTCUSDT","orderId":20,"clientOrderId":"CwOOIPHSmYywx6jZX77TdL"}],"orderReports":[{"symbol":"BTCUSDT","origClientOrderId":"CwOOIPHSmYywx6jZX77TdL","orderId":20,"orderListId":1929,"clientOrderId":"pXLV6Hz6mprAcVYpVMTGgx","price":"0.668611","origQty":"0.690354","executedQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"BUY","stopPrice":"0.378131","icebergQty":"0.017083"}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/openOrders", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountCancelAllOpenOrdersOnASymbolResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountCancelAllOpenOrdersOnASymbolResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountCancelAllOpenOrdersOnASymbolResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountCancelAllOpenOrdersOnASymbol Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelAllOpenOrdersOnASymbol Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelAllOpenOrdersOnASymbol(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":0,"contingencyType":"OCO","listStatusType":"ALL_DONE","listOrderStatus":"ALL_DONE","listClientOrderId":"C3wyj4WVEktd7u9aVBRXcN","transactionTime":1574040868128,"symbol":"LTCBTC","isIsolated":false,"orders":[{"symbol":"LTCBTC","orderId":2,"clientOrderId":"pO9ufTiFGg3nw2fOdgeOXa"}],"orderReports":[{"symbol":"LTCBTC","origClientOrderId":"pO9ufTiFGg3nw2fOdgeOXa","orderId":2,"orderListId":0,"clientOrderId":"unfWT8ig8i0uj6lPuYLez6","price":"1.00000000","origQty":"10.00000000","executedQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"STOP_LOSS_LIMIT","side":"SELL","stopPrice":"1.00000000","selfTradePreventionMode":"NONE"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/orderList", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountCancelOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOco(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountCancelOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountCancelOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"LTCBTC","orderId":"28","origClientOrderId":"myOrder1","clientOrderId":"cancelMyOrder1","price":"1.00000000","origQty":"10.00000000","executedQty":"8.00000000","cummulativeQuoteQty":"8.00000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"SELL","isIsolated":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order", r.URL.Path)
			require.Equal(t, "LTCBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountCancelOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOrder(context.Background()).Symbol("LTCBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountCancelOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountCancelOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountCancelOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountCancelOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":0,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"JYVpp3F0f5CAG15DhtrqLp","transactionTime":1563417480525,"symbol":"LTCBTC","marginBuyBorrowAmount":"5","marginBuyBorrowAsset":"BTC","isIsolated":false,"orders":[{"symbol":"LTCBTC","orderId":2,"clientOrderId":"Kk7sqHb9J6mJWTMDVW7Vos"}],"orderReports":[{"symbol":"LTCBTC","orderId":2,"orderListId":0,"clientOrderId":"Kk7sqHb9J6mJWTMDVW7Vos","transactTime":1563417480525,"price":"0.000000","origQty":"0.624363","executedQty":"0.000000","cummulativeQuoteQty":"0.000000","status":"NEW","timeInForce":"GTC","type":"STOP_LOSS","side":"BUY","stopPrice":"0.960664","selfTradePreventionMode":"NONE"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order/oco", r.URL.Path)
			require.Equal(t, "LTCBTC", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.MarginAccountNewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("quantity"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("price"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("stopPrice"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountNewOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Symbol("LTCBTC").Side(models.MarginAccountNewOrderSideParameterBuy).Quantity(float64(1.0)).Price(float64(1.0)).StopPrice(float64(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountNewOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountNewOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountNewOco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","orderId":26769564559,"clientOrderId":"E156O3KP4gOif65bjuUK5V","isIsolated":false,"transactTime":1713873075893,"price":"0","origQty":"0.001","executedQty":"0.001","cummulativeQuoteQty":"65.98253","status":"FILLED","timeInForce":"GTC","type":"MARKET","side":"SELL","selfTradePreventionMode":"EXPIRE_MAKER","marginBuyBorrowAmount":5,"marginBuyBorrowAsset":"BTC","fills":[{"price":"65982.53","qty":"0.001","commission":"0.06598253","commissionAsset":"USDT","tradeId":3570680726}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.MarginAccountNewOrderSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, string(models.MarginAccountNewOrderTypeParameterLimit), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountNewOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOrder(context.Background()).Symbol("BTCUSDT").Side(models.MarginAccountNewOrderSideParameterBuy).Type(models.MarginAccountNewOrderTypeParameterLimit).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountNewOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountNewOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountNewOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOto Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":13551,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"JDuOrsu0Ge8GTyvx8J7VTD","transactionTime":1725521998054,"symbol":"BTCUSDT","isIsolated":false,"orders":[{"symbol":"BTCUSDT","orderId":29896699,"clientOrderId":"y8RB6tQEMuHUXybqbtzTxk"}],"orderReports":[{"symbol":"BTCUSDT","orderId":29896699,"orderListId":13551,"clientOrderId":"y8RB6tQEMuHUXybqbtzTxk","transactTime":1725521998054,"price":"80000.00000000","origQty":"0.02000000","executedQty":"0","cummulativeQuoteQty":"0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"SELL","selfTradePreventionMode":"NONE"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order/oto", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.MarginAccountNewOtoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.MarginAccountNewOtoWorkingSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("workingPrice"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("workingQuantity"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("workingIcebergQty"))
			require.Equal(t, string(models.MarginAccountNewOrderTypeParameterLimit), r.URL.Query().Get("pendingType"))
			require.Equal(t, string(models.MarginAccountNewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("pendingQuantity"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountNewOtoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOto(context.Background()).Symbol("BTCUSDT").WorkingType(models.MarginAccountNewOtoWorkingTypeParameterLimit).WorkingSide(models.MarginAccountNewOtoWorkingSideParameterBuy).WorkingPrice(float64(1.0)).WorkingQuantity(float64(1.0)).WorkingIcebergQty(float64(1.0)).PendingType(models.MarginAccountNewOrderTypeParameterLimit).PendingSide(models.MarginAccountNewOrderSideParameterBuy).PendingQuantity(float64(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountNewOtoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountNewOtoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountNewOto Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOto(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOto Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOto(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOtoco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":13509,"contingencyType":"OTO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"u2AUo48LLef5qVenRtwJZy","transactionTime":1725521881300,"symbol":"BNBUSDT","isIsolated":false,"orders":[{"symbol":"BNBUSDT","orderId":28282534,"clientOrderId":"IfYDxvrZI4kiyqYpRH13iI"}],"orderReports":[{"symbol":"BNBUSDT","orderId":28282534,"orderListId":13509,"clientOrderId":"IfYDxvrZI4kiyqYpRH13iI","transactTime":1725521881300,"price":"300.00000000","origQty":"1.00000000","executedQty":"0","cummulativeQuoteQty":"0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","selfTradePreventionMode":"NONE","stopPrice":"299.00000000"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order/otoco", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.MarginAccountNewOtoWorkingTypeParameterLimit), r.URL.Query().Get("workingType"))
			require.Equal(t, string(models.MarginAccountNewOtoWorkingSideParameterBuy), r.URL.Query().Get("workingSide"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("workingPrice"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("workingQuantity"))
			require.Equal(t, string(models.MarginAccountNewOrderSideParameterBuy), r.URL.Query().Get("pendingSide"))
			require.Equal(t, fmt.Sprintf("%v", float64(1.0)), r.URL.Query().Get("pendingQuantity"))
			require.Equal(t, string(models.MarginAccountNewOtocoPendingAboveTypeParameterLimitMaker), r.URL.Query().Get("pendingAboveType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountNewOtocoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Symbol("BTCUSDT").WorkingType(models.MarginAccountNewOtoWorkingTypeParameterLimit).WorkingSide(models.MarginAccountNewOtoWorkingSideParameterBuy).WorkingPrice(float64(1.0)).WorkingQuantity(float64(1.0)).PendingSide(models.MarginAccountNewOrderSideParameterBuy).PendingQuantity(float64(1.0)).PendingAboveType(models.MarginAccountNewOtocoPendingAboveTypeParameterLimitMaker).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountNewOtocoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountNewOtocoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginAccountNewOtoco Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginAccountNewOtoco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginAccountNewOtoco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginManualLiquidation Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"asset":"ETH","interest":"0.00083334","principal":"0.001","liabilityAsset":"USDT","liabilityQty":0.3552}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/manual-liquidation", r.URL.Path)
			require.Equal(t, string(models.QueryMarginAvailableInventoryTypeParameterMargin), r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginManualLiquidationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginManualLiquidation(context.Background()).Type(models.QueryMarginAvailableInventoryTypeParameterMargin).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginManualLiquidationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginManualLiquidationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService MarginManualLiquidation Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginManualLiquidation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService MarginManualLiquidation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.MarginManualLiquidation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryCurrentMarginOrderCountUsage Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":10000,"count":0}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/rateLimit/order", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryCurrentMarginOrderCountUsageResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryCurrentMarginOrderCountUsage(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryCurrentMarginOrderCountUsageResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryCurrentMarginOrderCountUsageResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryCurrentMarginOrderCountUsage Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryCurrentMarginOrderCountUsage(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryLiquidationLoan Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"asset":"USDC","amount":"1000.00000000","repaidAmount":"300.00000000","remainingAmount":"700.00000000"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/liquidation-loan", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryLiquidationLoanResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoan(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryLiquidationLoanResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryLiquidationLoanResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryLiquidationLoan Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoan(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryLiquidationLoanRepayHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"total":2,"rows":[{"repayId":12345678,"asset":"USDC","amount":"300.00000000","status":"SUCCESS","createTime":1714492800000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/liquidation-loan/repay-history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryLiquidationLoanRepayHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoanRepayHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryLiquidationLoanRepayHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryLiquidationLoanRepayHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryLiquidationLoanRepayHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryLiquidationLoanRepayHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsAllOco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"orderListId":29,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"amEEAXryFzFwYF1FeRpUoZ","transactionTime":1565245913483,"symbol":"LTCBTC","isIsolated":true,"orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"oD7aesZqjEGlZrbtRpy5zB"}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/allOrderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsAllOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOco(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsAllOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsAllOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsAllOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsAllOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"clientOrderId":"D2KDy4DIeS56PvkM13f8cP","cummulativeQuoteQty":"0.00000000","executedQty":"0.00000000","icebergQty":"0.00000000","isWorking":false,"orderId":41295,"origQty":"5.31000000","price":"0.22500000","side":"SELL","status":"CANCELED","stopPrice":"0.18000000","symbol":"BNBBTC","isIsolated":false,"time":1565769338806,"timeInForce":"GTC","type":"TAKE_PROFIT_LIMIT","selfTradePreventionMode":"NONE","updateTime":1565769342148}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/allOrders", r.URL.Path)
			require.Equal(t, "BNBBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsAllOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOrders(context.Background()).Symbol("BNBBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsAllOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsAllOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsAllOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsAllOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsAllOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":27,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"h2USkA5YQpaXHPIrkd96xE","transactionTime":1565245656253,"symbol":"LTCBTC","isIsolated":true,"orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"qD1gy3kc3Gx0rihm9Y3xwS"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/orderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOco(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOpenOco Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"orderListId":31,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"wuB13fmulKj3YjdqWEcsnp","transactionTime":1565246080644,"symbol":"LTCBTC","isIsolated":true,"orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"r3EH2N76dHfLoSZWIUw1bT"}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/openOrderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsOpenOcoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOco(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsOpenOcoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsOpenOcoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOpenOco Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOco(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOpenOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"clientOrderId":"qhcZw71gAkCCTv0t0k8LUK","cummulativeQuoteQty":"0.00000000","executedQty":"0.00000000","icebergQty":"0.00000000","isWorking":true,"orderId":211842552,"origQty":"0.30000000","price":"0.00475010","side":"SELL","status":"NEW","stopPrice":"0.00000000","symbol":"BNBBTC","isIsolated":true,"time":1562040170089,"timeInForce":"GTC","type":"LIMIT","selfTradePreventionMode":"NONE","updateTime":1562040170089}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/openOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"clientOrderId":"ZwfQzuDIGpceVhKW5DvCmO","cummulativeQuoteQty":"0.00000000","executedQty":"0.00000000","icebergQty":"0.00000000","isWorking":true,"orderId":213205622,"origQty":"0.30000000","price":"0.00493630","side":"SELL","status":"NEW","stopPrice":"0.00000000","symbol":"BNBBTC","isIsolated":true,"time":1562133008725,"timeInForce":"GTC","type":"LIMIT","selfTradePreventionMode":"NONE","updateTime":1562133008725}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/order", r.URL.Path)
			require.Equal(t, "BNBBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOrder(context.Background()).Symbol("BNBBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsTradeList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"commission":"0.00006000","commissionAsset":"BTC","id":34,"isBestMatch":true,"isBuyer":false,"isMaker":false,"orderId":39324,"price":"0.02000000","qty":"3.00000000","symbol":"BNBBTC","isIsolated":false,"time":1561973357171}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/myTrades", r.URL.Path)
			require.Equal(t, "BNBBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginAccountsTradeListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsTradeList(context.Background()).Symbol("BNBBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginAccountsTradeListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginAccountsTradeListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsTradeList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsTradeList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryMarginAccountsTradeList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryMarginAccountsTradeList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryPreventedMatches Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","price":"1.100000","makerPreventedQuantity":"1.300000","transactTime":1669101687094}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/myPreventedMatches", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryPreventedMatchesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryPreventedMatches(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryPreventedMatchesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryPreventedMatchesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryPreventedMatches Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryPreventedMatches(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryPreventedMatches Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryPreventedMatches(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QuerySpecialKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"apiKey":"npOzOAeLVgr2TuxWfNo43AaPWpBbJEoKezh1o8mSQb6ryE2odE11A4AoVlJbQoGx","ip":"0.0.0.0,192.168.0.1,192.168.0.2","apiName":"testName","type":"RSA","permissionMode":"TRADE"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/apiKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySpecialKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKey(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySpecialKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySpecialKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QuerySpecialKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QuerySpecialKeyList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"apiName":"testName1","apiKey":"znpOzOAeLVgr2TuxWfNo43AaPWpBbJEoKezh1o8mSQb6ryE2odE11A4AoVlJbQoG","ip":"192.168.0.1,192.168.0.2","type":"RSA","permissionMode":"TRADE"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/api-key-list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySpecialKeyListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKeyList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySpecialKeyListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySpecialKeyListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QuerySpecialKeyList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QuerySpecialKeyList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService SmallLiabilityExchange Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/exchange-small-liability", r.URL.Path)
			require.Equal(t, "BTC,ETH", r.URL.Query().Get("assetNames"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).AssetNames("BTC,ETH").Execute()
		require.NoError(t, err)
	})

	t.Run("Test TradeAPIService SmallLiabilityExchange Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test TradeAPIService SmallLiabilityExchange Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.TradeAPI.SmallLiabilityExchange(context.Background()).Execute()

		require.Error(t, err)
	})

}
