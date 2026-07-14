/*
Prediction Trading REST API TEST

Testing WalletAPIService

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

func Test_binancew3wpredictionrestapi_WalletAPIService(t *testing.T) {

	t.Run("Test WalletAPIService GetPortfolio Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"chainId":"56","walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","activePositionsCount":3,"totalRealizedPnl":"120.50","totalUnrealizedPnl":"15.30","totalPnl":"135.80","totalCostBasis":"450.00","totalCurrentValue":"465.30","positions":[{"id":10001,"walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","marketTopicId":4229564,"marketId":5567895,"tokenId":"112233","vendor":"PREDICT_FUN","currentShares":"1923.07","avgPrice":"0.52","currentPrice":"0.55","realizedPnl":"0.00","unrealizedPnl":"0.06","totalPnl":"0.06","pnlPercentage":"6.00","isResolved":false}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/pnl/portfolio", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetPortfolioResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.WalletAPI.GetPortfolio(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetPortfolioResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetPortfolioResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test WalletAPIService GetPortfolio Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.WalletAPI.GetPortfolio(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WalletAPIService GetPortfolio Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.WalletAPI.GetPortfolio(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WalletAPIService GetQuotaStatus Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"dailyLimit":"10000.00","remainingDailyLimit":"8500.00"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/quota/limit/status", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetQuotaStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.WalletAPI.GetQuotaStatus(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetQuotaStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetQuotaStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test WalletAPIService GetQuotaStatus Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.WalletAPI.GetQuotaStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WalletAPIService ListPredictionWallets Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"wallets":[{"walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","walletId":"5b5c1ec3be4e4416a5872b21c1ca5d20","registeredTime":1748000000000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/wallet/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ListPredictionWalletsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.WalletAPI.ListPredictionWallets(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ListPredictionWalletsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ListPredictionWalletsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test WalletAPIService ListPredictionWallets Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.WalletAPI.ListPredictionWallets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test WalletAPIService QueryPaymentOptionBalances Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"items":[{"accountType":"SPOT","availableBalanceDisplay":"1000.00","enabled":true}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/balance/payment-options", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryPaymentOptionBalancesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.WalletAPI.QueryPaymentOptionBalances(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryPaymentOptionBalancesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryPaymentOptionBalancesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test WalletAPIService QueryPaymentOptionBalances Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.WalletAPI.QueryPaymentOptionBalances(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
