/*
Binance Algo REST API TEST

Testing SpotAlgoAPIService

*/

package binancealgorestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/clients/algo/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancealgorestapi_SpotAlgoAPIService(t *testing.T) {

	t.Run("Test SpotAlgoAPIService CancelAlgoOrderSpotAlgo Success", func(t *testing.T) {

		mockedJSON := `{"algoId":14511,"success":true,"code":0,"msg":"OK"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/spot/order", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("algoId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelAlgoOrderSpotAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.CancelAlgoOrderSpotAlgo(context.Background()).AlgoId(int64(1)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelAlgoOrderSpotAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelAlgoOrderSpotAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SpotAlgoAPIService CancelAlgoOrderSpotAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.CancelAlgoOrderSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService CancelAlgoOrderSpotAlgo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.CancelAlgoOrderSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService QueryCurrentAlgoOpenOrdersSpotAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"orders":[{"algoId":14517,"symbol":"ETHUSDT","side":"SELL","totalQty":"5.000","executedQty":"0.000","executedAmt":"0.00000000","avgPrice":"0.00","clientAlgoId":"d7096549481642f8a0bb69e9e2e31f2e","bookTime":1649756817004,"endTime":0,"algoStatus":"WORKING","algoType":"TWAP","urgency":"LOW"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/spot/openOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QueryCurrentAlgoOpenOrdersSpotAlgo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryCurrentAlgoOpenOrdersSpotAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SpotAlgoAPIService QueryCurrentAlgoOpenOrdersSpotAlgo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QueryCurrentAlgoOpenOrdersSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService QueryHistoricalAlgoOrdersSpotAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"orders":[{"algoId":14518,"symbol":"BNBUSDT","side":"BUY","totalQty":"100.00","executedQty":"0.00","executedAmt":"0.00000000","avgPrice":"0.000","clientAlgoId":"acacab56b3c44bef9f6a8f8ebd2a8408","bookTime":1649757019503,"endTime":1649757088101,"algoStatus":"CANCELLED","algoType":"VP","urgency":"LOW"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/spot/historicalOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryHistoricalAlgoOrdersSpotAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryHistoricalAlgoOrdersSpotAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryHistoricalAlgoOrdersSpotAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SpotAlgoAPIService QueryHistoricalAlgoOrdersSpotAlgo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService QuerySubOrdersSpotAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"executedQty":"1.000","executedAmt":"3229.44000000","subOrders":[{"algoId":13723,"orderId":8389765519993909000,"orderStatus":"FILLED","executedQty":"1.000","executedAmt":"3229.44000000","feeAmt":"-1.61471999","feeAsset":"USDT","bookTime":1649319001964,"avgPrice":"3229.44","side":"SELL","symbol":"ETHUSDT","subId":1,"timeInForce":"IMMEDIATE_OR_CANCEL","origQty":"1.000"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/spot/subOrders", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("algoId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubOrdersSpotAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QuerySubOrdersSpotAlgo(context.Background()).AlgoId(int64(1)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubOrdersSpotAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubOrdersSpotAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SpotAlgoAPIService QuerySubOrdersSpotAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QuerySubOrdersSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService QuerySubOrdersSpotAlgo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.QuerySubOrdersSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService TimeWeightedAveragePriceSpotAlgo Success", func(t *testing.T) {

		mockedJSON := `{"clientAlgoId":"65ce1630101a480b85915d7e11fd5078","success":true,"code":0,"msg":"OK"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/spot/newOrderTwap", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "BUY", r.URL.Query().Get("side"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			require.Equal(t, "5000", r.URL.Query().Get("duration"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TimeWeightedAveragePriceSpotAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.TimeWeightedAveragePriceSpotAlgo(context.Background()).Symbol("BTCUSDT").Side("BUY").Quantity(float32(1.0)).Duration(int64(5000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TimeWeightedAveragePriceSpotAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TimeWeightedAveragePriceSpotAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SpotAlgoAPIService TimeWeightedAveragePriceSpotAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.TimeWeightedAveragePriceSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SpotAlgoAPIService TimeWeightedAveragePriceSpotAlgo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SpotAlgoAPI.TimeWeightedAveragePriceSpotAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
