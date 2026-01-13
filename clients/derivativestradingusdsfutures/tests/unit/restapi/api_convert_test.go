/*
Binance Derivatives Trading USDS Futures REST API TEST

Testing ConvertAPIService

*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingusdsfuturesrestapi_ConvertAPIService(t *testing.T) {

	t.Run("Test ConvertAPIService AcceptTheOfferedQuote Success", func(t *testing.T) {

		mockedJSON := `{"orderId":"933256278426274426","createTime":1623381330472,"orderStatus":"PROCESS"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/convert/acceptQuote", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("quoteId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AcceptTheOfferedQuoteResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.AcceptTheOfferedQuote(context.Background()).QuoteId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AcceptTheOfferedQuoteResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AcceptTheOfferedQuoteResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ConvertAPIService AcceptTheOfferedQuote Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.AcceptTheOfferedQuote(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ConvertAPIService AcceptTheOfferedQuote Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.AcceptTheOfferedQuote(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ConvertAPIService ListAllConvertPairs Success", func(t *testing.T) {

		mockedJSON := `[{"fromAsset":"BTC","toAsset":"USDT","fromAssetMinAmount":"0.0004","fromAssetMaxAmount":"50","toAssetMinAmount":"20","toAssetMaxAmount":"2500000"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/convert/exchangeInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ListAllConvertPairsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.ListAllConvertPairs(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ListAllConvertPairsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ListAllConvertPairsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ConvertAPIService ListAllConvertPairs Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.ListAllConvertPairs(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ConvertAPIService OrderStatus Success", func(t *testing.T) {

		mockedJSON := `{"orderId":933256278426274400,"orderStatus":"SUCCESS","fromAsset":"BTC","fromAmount":"0.00054414","toAsset":"USDT","toAmount":"20","ratio":"36755","inverseRatio":"0.00002721","createTime":1623381330472}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/convert/orderStatus", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.OrderStatus(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ConvertAPIService OrderStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.OrderStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ConvertAPIService SendQuoteRequest Success", func(t *testing.T) {

		mockedJSON := `{"quoteId":"12415572564","ratio":"38163.7","inverseRatio":"0.0000262","validTimestamp":1623319461670,"toAmount":"3816.37","fromAmount":"0.1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/convert/getQuote", r.URL.Path)
			require.Equal(t, "fromAsset_example", r.URL.Query().Get("fromAsset"))
			require.Equal(t, "toAsset_example", r.URL.Query().Get("toAsset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SendQuoteRequestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.SendQuoteRequest(context.Background()).FromAsset("fromAsset_example").ToAsset("toAsset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SendQuoteRequestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SendQuoteRequestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ConvertAPIService SendQuoteRequest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.SendQuoteRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ConvertAPIService SendQuoteRequest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ConvertAPI.SendQuoteRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
