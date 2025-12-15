/*
Binance Sub Account REST API TEST

Testing ApiManagementAPIService

*/

package binancesubaccountrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/subaccount/src"
	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesubaccountrestapi_ApiManagementAPIService(t *testing.T) {

	t.Run("Test ApiManagementAPIService AddIpRestrictionForSubAccountApiKey Success", func(t *testing.T) {

		mockedJSON := `{"status":"2","ipList":["69.210.67.14","8.34.21.10"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/sub-account/subAccountApi/ipRestriction", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "subAccountApiKey_example", r.URL.Query().Get("subAccountApiKey"))
			require.Equal(t, "status_example", r.URL.Query().Get("status"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AddIpRestrictionForSubAccountApiKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Email("sub-account-email@email.com").SubAccountApiKey("subAccountApiKey_example").Status("status_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AddIpRestrictionForSubAccountApiKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AddIpRestrictionForSubAccountApiKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService AddIpRestrictionForSubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService AddIpRestrictionForSubAccountApiKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService DeleteIpListForASubAccountApiKey Success", func(t *testing.T) {

		mockedJSON := `{"ipRestrict":"true","ipList":["69.210.67.14","8.34.21.10"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "subAccountApiKey_example", r.URL.Query().Get("subAccountApiKey"))
			require.Equal(t, "ipAddress_example", r.URL.Query().Get("ipAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DeleteIpListForASubAccountApiKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteIpListForASubAccountApiKey(context.Background()).Email("sub-account-email@email.com").SubAccountApiKey("subAccountApiKey_example").IpAddress("ipAddress_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DeleteIpListForASubAccountApiKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DeleteIpListForASubAccountApiKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService DeleteIpListForASubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteIpListForASubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService DeleteIpListForASubAccountApiKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteIpListForASubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService GetIpRestrictionForASubAccountApiKey Success", func(t *testing.T) {

		mockedJSON := `{"ipRestrict":"true","ipList":["69.210.67.14","8.34.21.10"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi/ipRestriction", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "subAccountApiKey_example", r.URL.Query().Get("subAccountApiKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetIpRestrictionForASubAccountApiKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.GetIpRestrictionForASubAccountApiKey(context.Background()).Email("sub-account-email@email.com").SubAccountApiKey("subAccountApiKey_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetIpRestrictionForASubAccountApiKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetIpRestrictionForASubAccountApiKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService GetIpRestrictionForASubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.GetIpRestrictionForASubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService GetIpRestrictionForASubAccountApiKey Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.GetIpRestrictionForASubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
