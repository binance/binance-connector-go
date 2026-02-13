/*
Binance Derivatives Trading Portfolio Margin REST API TEST

Testing UserDataStreamsAPIService
*/

package binancederivativestradingportfoliomarginrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin"
	"github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingportfoliomarginrestapi_UserDataStreamsAPIService(t *testing.T) {

	t.Run("Test UserDataStreamsAPIService CloseUserDataStream Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/papi/v1/listenKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.UserDataStreamsAPI.CloseUserDataStream(context.Background()).Execute()
		require.NoError(t, err)
	})

	t.Run("Test UserDataStreamsAPIService CloseUserDataStream Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.UserDataStreamsAPI.CloseUserDataStream(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test UserDataStreamsAPIService KeepaliveUserDataStream Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/papi/v1/listenKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.UserDataStreamsAPI.KeepaliveUserDataStream(context.Background()).Execute()
		require.NoError(t, err)
	})

	t.Run("Test UserDataStreamsAPIService KeepaliveUserDataStream Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.UserDataStreamsAPI.KeepaliveUserDataStream(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test UserDataStreamsAPIService StartUserDataStream Success", func(t *testing.T) {

		mockedJSON := `{"listenKey":"pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/papi/v1/listenKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.StartUserDataStreamResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserDataStreamsAPI.StartUserDataStream(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.StartUserDataStreamResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.StartUserDataStreamResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserDataStreamsAPIService StartUserDataStream Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserDataStreamsAPI.StartUserDataStream(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
