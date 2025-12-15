/*
Binance Algo REST API TEST

Testing FutureAlgoAPIService

*/

package binancealgorestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/algo/src"
	"github.com/binance/binance-connector-go/clients/algo/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancealgorestapi_FutureAlgoAPIService(t *testing.T) {

	t.Run("Test FutureAlgoAPIService CancelAlgoOrderFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"algoId":14511,"success":true,"code":0,"msg":"OK"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/order", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("algoId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelAlgoOrderFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.CancelAlgoOrderFutureAlgo(context.Background()).AlgoId(int64(1)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelAlgoOrderFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelAlgoOrderFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService CancelAlgoOrderFutureAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.CancelAlgoOrderFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService CancelAlgoOrderFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.CancelAlgoOrderFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService QueryCurrentAlgoOpenOrdersFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"orders":[{"algoId":14517,"symbol":"ETHUSDT","side":"SELL","positionSide":"SHORT","totalQty":"5.000","executedQty":"0.000","executedAmt":"0.00000000","avgPrice":"0.00","clientAlgoId":"d7096549481642f8a0bb69e9e2e31f2e","bookTime":1649756817004,"endTime":0,"algoStatus":"WORKING","algoType":"VP","urgency":"LOW"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/openOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.QueryCurrentAlgoOpenOrdersFutureAlgo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryCurrentAlgoOpenOrdersFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService QueryCurrentAlgoOpenOrdersFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.QueryCurrentAlgoOpenOrdersFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService QueryHistoricalAlgoOrdersFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"orders":[{"algoId":14518,"symbol":"BNBUSDT","side":"BUY","positionSide":"BOTH","totalQty":"100.00","executedQty":"0.00","executedAmt":"0.00000000","avgPrice":"0.000","clientAlgoId":"acacab56b3c44bef9f6a8f8ebd2a8408","bookTime":1649757019503,"endTime":1649757088101,"algoStatus":"CANCELLED","algoType":"VP","urgency":"LOW"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/historicalOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryHistoricalAlgoOrdersFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.QueryHistoricalAlgoOrdersFutureAlgo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryHistoricalAlgoOrdersFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryHistoricalAlgoOrdersFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService QueryHistoricalAlgoOrdersFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.QueryHistoricalAlgoOrdersFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService QuerySubOrdersFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"executedQty":"1.000","executedAmt":"3229.44000000","subOrders":[{"algoId":13723,"orderId":8389765519993909000,"orderStatus":"FILLED","executedQty":"1.000","executedAmt":"3229.44000000","feeAmt":"-1.61471999","feeAsset":"USDT","bookTime":1649319001964,"avgPrice":"3229.44","side":"SELL","symbol":"ETHUSDT","subId":1,"timeInForce":"IMMEDIATE_OR_CANCEL","origQty":"1.000"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/subOrders", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("algoId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubOrdersFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.QuerySubOrdersFutureAlgo(context.Background()).AlgoId(int64(1)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubOrdersFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubOrdersFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService QuerySubOrdersFutureAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.QuerySubOrdersFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService QuerySubOrdersFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.QuerySubOrdersFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService TimeWeightedAveragePriceFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"clientAlgoId":"65ce1630101a480b85915d7e11fd5078","success":true,"code":0,"msg":"OK"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/newOrderTwap", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "BUY", r.URL.Query().Get("side"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			require.Equal(t, "5000", r.URL.Query().Get("duration"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TimeWeightedAveragePriceFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.TimeWeightedAveragePriceFutureAlgo(context.Background()).Symbol("BTCUSDT").Side("BUY").Quantity(float32(1.0)).Duration(int64(5000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TimeWeightedAveragePriceFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TimeWeightedAveragePriceFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService TimeWeightedAveragePriceFutureAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.TimeWeightedAveragePriceFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService TimeWeightedAveragePriceFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.TimeWeightedAveragePriceFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService VolumeParticipationFutureAlgo Success", func(t *testing.T) {

		mockedJSON := `{"clientAlgoId":"00358ce6a268403398bd34eaa36dffe7","success":true,"code":0,"msg":"OK"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/algo/futures/newOrderVp", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "BUY", r.URL.Query().Get("side"))
			require.Equal(t, "1", r.URL.Query().Get("quantity"))
			require.Equal(t, "LOW", r.URL.Query().Get("urgency"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VolumeParticipationFutureAlgoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.VolumeParticipationFutureAlgo(context.Background()).Symbol("BTCUSDT").Side("BUY").Quantity(float32(1.0)).Urgency("LOW").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VolumeParticipationFutureAlgoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VolumeParticipationFutureAlgoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureAlgoAPIService VolumeParticipationFutureAlgo Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlgoClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureAlgoAPI.VolumeParticipationFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureAlgoAPIService VolumeParticipationFutureAlgo Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FutureAlgoAPI.VolumeParticipationFutureAlgo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
