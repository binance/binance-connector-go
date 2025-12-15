/*
Binance Margin Trading REST API TEST

Testing TransferAPIService

*/

package binancemargintradingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/margintrading/src"
	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancemargintradingrestapi_TransferAPIService(t *testing.T) {

	t.Run("Test TransferAPIService GetCrossMarginTransferHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"amount":"0.10000000","asset":"BNB","status":"CONFIRMED","timestamp":1566898617,"txId":5240372201,"type":"ROLL_IN","transFrom":"SPOT","transTo":"ISOLATED_MARGIN"},{"amount":"5.00000000","asset":"USDT","status":"CONFIRMED","timestamp":1566888436,"txId":5239810406,"type":"ROLL_OUT","transFrom":"ISOLATED_MARGIN","transTo":"ISOLATED_MARGIN","fromSymbol":"BNBUSDT","toSymbol":"BTCUSDT"},{"amount":"1.00000000","asset":"EOS","status":"CONFIRMED","timestamp":1566888403,"txId":5239808703,"type":"ROLL_IN"}],"total":3}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/transfer", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCrossMarginTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.GetCrossMarginTransferHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCrossMarginTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCrossMarginTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService GetCrossMarginTransferHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.TransferAPI.GetCrossMarginTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryMaxTransferOutAmount Success", func(t *testing.T) {

		mockedJSON := `{"amount":"3.59498107"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/maxTransferable", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMaxTransferOutAmountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryMaxTransferOutAmount(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMaxTransferOutAmountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMaxTransferOutAmountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService QueryMaxTransferOutAmount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryMaxTransferOutAmount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryMaxTransferOutAmount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.TransferAPI.QueryMaxTransferOutAmount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
