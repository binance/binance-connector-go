/*
Binance Simple Earn REST API TEST

Testing RwusdAPIService

*/

package binancesimpleearnrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/simpleearn/src"
	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesimpleearnrestapi_RwusdAPIService(t *testing.T) {

	t.Run("Test RwusdAPIService GetRwusdAccount Success", func(t *testing.T) {

		mockedJSON := `{"rwusdAmount":"100","totalProfit":"12.81"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService GetRwusdQuotaDetails Success", func(t *testing.T) {

		mockedJSON := `{"subscriptionQuota":{"assets":["USDT","USDC"],"leftQuota":"1000","minimum":"0.10000000"},"fastRedemptionQuota":{"leftQuota":"2","minimum":"0.1","fee":"0.0005","freeQuota":"100"},"standardRedemptionQuota":{"leftQuota":"2","minimum":"0.1","fee":"0.001","redeemPeriod":3},"subscribeEnable":true,"redeemEnable":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/quota", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdQuotaDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdQuotaDetails(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdQuotaDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdQuotaDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdQuotaDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdQuotaDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService GetRwusdRateHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"annualPercentageRate":"0.0418","time":1577233578000}],"total":"1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/history/rateHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRateHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdRateHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService GetRwusdRedemptionHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1575018510000,"asset":"RWUSD","amount":"51","receiveAsset":"USDC","receiveAmount":"50","fee":"1","arrivalTime":1575018510000,"status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/history/redemptionHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdRedemptionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRedemptionHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdRedemptionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdRedemptionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdRedemptionHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRedemptionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService GetRwusdRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1575018510000,"rewardsAmount":"1","rwusdPosition":"100","annualPercentageRate":"0.0418"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/history/rewardsHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRewardsHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService GetRwusdSubscriptionHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1575018510000,"asset":"USDC","amount":"100","receiveAsset":"RWUSD","receiveAmount":"100","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/history/subscriptionHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRwusdSubscriptionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdSubscriptionHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRwusdSubscriptionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRwusdSubscriptionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService GetRwusdSubscriptionHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.GetRwusdSubscriptionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService RedeemRwusd Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"receiveAmount":"0.23092091","fee":"0.00000012","arrivalTime":1575018510000}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/redeem", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			require.Equal(t, "s", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RedeemRwusdResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.RedeemRwusd(context.Background()).Amount(float32(1.0)).Type("s").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RedeemRwusdResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RedeemRwusdResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService RedeemRwusd Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.RedeemRwusd(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService RedeemRwusd Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.RedeemRwusd(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService SubscribeRwusd Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"rwusdAmount":"0.22091092"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rwusd/subscribe", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubscribeRwusdResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.SubscribeRwusd(context.Background()).Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubscribeRwusdResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubscribeRwusdResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RwusdAPIService SubscribeRwusd Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.SubscribeRwusd(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test RwusdAPIService SubscribeRwusd Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RwusdAPI.SubscribeRwusd(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
