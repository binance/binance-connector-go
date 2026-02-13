/*
Binance Staking REST API TEST

Testing SoftStakingAPIService

*/

package binancestakingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancestakingrestapi_SoftStakingAPIService(t *testing.T) {

	t.Run("Test SoftStakingAPIService GetSoftStakingProductList Success", func(t *testing.T) {

		mockedJSON := `{"status":true,"totalRewardsUsdt":"3.09827182","rows":[{"asset":"BNB","minAmount":"0.5","maxCap":"1000","apr":"0.0015","stakedAmount":"2.14","totalProfit":"0.00171234"},{"asset":"SUI","minAmount":"100","maxCap":"50000","apr":"0.01","stakedAmount":"100","totalProfit":"0.1"}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/soft-staking/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSoftStakingProductListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingProductList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSoftStakingProductListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSoftStakingProductListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SoftStakingAPIService GetSoftStakingProductList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SoftStakingAPIService GetSoftStakingRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"asset":"BNB","rewards":"0.00000557","rewardAsset":"BNB","avgAmount":"2.14","time":1754007978000},{"asset":"SUI","rewards":"0.00274257","rewardAsset":"SUI","avgAmount":"100","time":1754007978000}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/soft-staking/history/rewardsRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSoftStakingRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingRewardsHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSoftStakingRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSoftStakingRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SoftStakingAPIService GetSoftStakingRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.GetSoftStakingRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SoftStakingAPIService SetSoftStaking Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/soft-staking/set", r.URL.Path)
			require.Equal(t, "true", r.URL.Query().Get("softStaking"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetSoftStakingResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.SetSoftStaking(context.Background()).SoftStaking(true).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetSoftStakingResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetSoftStakingResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SoftStakingAPIService SetSoftStaking Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.SetSoftStaking(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SoftStakingAPIService SetSoftStaking Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SoftStakingAPI.SetSoftStaking(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
