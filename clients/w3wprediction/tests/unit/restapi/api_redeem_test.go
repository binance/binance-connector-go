/*
Prediction Trading REST API TEST

Testing RedeemAPIService

*/

package binancew3wpredictionrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancew3wpredictionrestapi_RedeemAPIService(t *testing.T) {

	t.Run("Test RedeemAPIService BatchRedeem Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"batchId":"batch_20260525_redeem_001","results":[{"requestId":"req_001","txHash":"0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd","status":"PENDING","error":"error"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/batch-redeem", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "5b5c1ec3be4e4416a5872b21c1ca5d20", r.URL.Query().Get("walletId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.BatchRedeemResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.BatchRedeem(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").TokenIds([]string{"112233"}).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.BatchRedeemResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.BatchRedeemResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RedeemAPIService BatchRedeem Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.BatchRedeem(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RedeemAPIService BatchRedeem Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.BatchRedeem(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RedeemAPIService GetRedeemStatus Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"txHash":"0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd","status":"CONFIRMED"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/redeem/status", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd", r.URL.Query().Get("txHash"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRedeemStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.GetRedeemStatus(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").TxHash("0xabc123def456789abcdef123456789abcdef123456789abcdef123456789abcd").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRedeemStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRedeemStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RedeemAPIService GetRedeemStatus Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.GetRedeemStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RedeemAPIService GetRedeemStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RedeemAPI.GetRedeemStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
