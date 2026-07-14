/*
Sub Account REST API TEST

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

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesubaccountrestapi_ApiManagementAPIService(t *testing.T) {

	t.Run("Test ApiManagementAPIService AddIpRestrictionForSubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"status":"2","ipList":["69.210.67.14"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/sub-account/subAccountApi/ipRestriction", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf", r.URL.Query().Get("subAccountApiKey"))
			require.Equal(t, "1", r.URL.Query().Get("status"))
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

		resp, err := apiClient.RestApi.ApiManagementAPI.AddIpRestrictionForSubAccountApiKey(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").Status(int64(1)).Execute()
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

	t.Run("Test ApiManagementAPIService CreateSubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"apiName":"myKey","apiKey":"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A","secretKey":"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j","canTrade":true,"canMarginLoanRepay":false,"canFuturesTrade":false,"canUniversalTransfer":false,"canVanillaOptions":false,"status":2,"ipList":["69.210.67.14"]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "myKey", r.URL.Query().Get("apiName"))
			require.Equal(t, "2", r.URL.Query().Get("status"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateSubAccountApiKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.CreateSubAccountApiKey(context.Background()).Email("123@test.com").ApiName("myKey").Status(int64(2)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateSubAccountApiKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateSubAccountApiKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService CreateSubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.CreateSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService CreateSubAccountApiKey Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.ApiManagementAPI.CreateSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService DeleteIpListForASubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"ipRestrict":"true","ipList":["69.210.67.14"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf", r.URL.Query().Get("subAccountApiKey"))
			require.Equal(t, "69.210.67.14", r.URL.Query().Get("ipAddress"))
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

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteIpListForASubAccountApiKey(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").IpAddress("69.210.67.14").Execute()
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

	t.Run("Test ApiManagementAPIService DeleteSubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf", r.URL.Query().Get("subAccountApiKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected map[string]interface{}
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteSubAccountApiKey(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[map[string]interface{}]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(map[string]interface{}{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService DeleteSubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService DeleteSubAccountApiKey Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.ApiManagementAPI.DeleteSubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService GetIpRestrictionForASubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"ipRestrict":"true","ipList":["69.210.67.14"],"updateTime":1636371437000,"apiKey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi/ipRestriction", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf", r.URL.Query().Get("subAccountApiKey"))
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

		resp, err := apiClient.RestApi.ApiManagementAPI.GetIpRestrictionForASubAccountApiKey(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").Execute()
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

	t.Run("Test ApiManagementAPIService ModifySubAccountApiKeyPermission Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"apiName":"myKey","apikey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf","canTrade":true,"canMarginLoanRepay":false,"canFuturesTrade":true,"canUniversalTransfer":false,"canVanillaOptions":false,"timestamp":1640000000000}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApiPermission", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			require.Equal(t, "k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf", r.URL.Query().Get("subAccountApiKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ModifySubAccountApiKeyPermissionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.ModifySubAccountApiKeyPermission(context.Background()).Email("123@test.com").SubAccountApiKey("k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ModifySubAccountApiKeyPermissionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ModifySubAccountApiKeyPermissionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService ModifySubAccountApiKeyPermission Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.ModifySubAccountApiKeyPermission(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService ModifySubAccountApiKeyPermission Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.ApiManagementAPI.ModifySubAccountApiKeyPermission(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService QuerySubAccountApiKey Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"total":1,"list":[{"email":"123@test.com","apiName":"myKey","apikey":"k5V49ldtn4tszj6W3hystegdfvmGbqDzjmkCtpTvC0G74WhK7yd4rfCTo4lShf","canTrade":true,"canMarginLoanRepay":false,"canFuturesTrade":false,"canUniversalTransfer":false,"canVanillaOptions":false,"timestamp":1640000000000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/subAccountApi", r.URL.Path)
			require.Equal(t, "123@test.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountApiKeyResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.QuerySubAccountApiKey(context.Background()).Email("123@test.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountApiKeyResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountApiKeyResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ApiManagementAPIService QuerySubAccountApiKey Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ApiManagementAPI.QuerySubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ApiManagementAPIService QuerySubAccountApiKey Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.ApiManagementAPI.QuerySubAccountApiKey(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
