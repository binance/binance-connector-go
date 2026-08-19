/*
Stocks Trading REST API TEST

Testing UserDataStreamsAPIService

*/

package binancestocksrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancestocksrestapi_UserDataStreamsAPIService(t *testing.T) {

	t.Run("Test UserDataStreamsAPIService CreateRenewListenKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"listenKey":"pqia91ma19a5s61cv6a81va65sdf19v8a65s1cv1zuz3ee1c5xz2ef6ad7"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/equity/listenKey", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateRenewListenKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserDataStreamsAPI.CreateRenewListenKey(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateRenewListenKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateRenewListenKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserDataStreamsAPIService CreateRenewListenKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserDataStreamsAPI.CreateRenewListenKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
