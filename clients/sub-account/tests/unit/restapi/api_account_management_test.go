/*
Binance Sub Account REST API TEST

Testing AccountManagementAPIService

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

func Test_binancesubaccountrestapi_AccountManagementAPIService(t *testing.T) {

	t.Run("Test AccountManagementAPIService CreateAVirtualSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"email":"addsdd_virtual@aasaixwqnoemail.com"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/virtualSubAccount", r.URL.Path)
			require.Equal(t, "subAccountString_example", r.URL.Query().Get("subAccountString"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateAVirtualSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.CreateAVirtualSubAccount(context.Background()).SubAccountString("subAccountString_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateAVirtualSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateAVirtualSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService CreateAVirtualSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.CreateAVirtualSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService CreateAVirtualSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.CreateAVirtualSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService EnableFuturesForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"email":"123@test.com","isFuturesEnabled":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/enable", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.EnableFuturesForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableFuturesForSubAccount(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.EnableFuturesForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.EnableFuturesForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService EnableFuturesForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableFuturesForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService EnableFuturesForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableFuturesForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService EnableOptionsForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"email":"123@test.com","isEOptionsEnabled":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/eoptions/enable", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.EnableOptionsForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableOptionsForSubAccount(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.EnableOptionsForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.EnableOptionsForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService EnableOptionsForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableOptionsForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService EnableOptionsForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.EnableOptionsForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccount Success", func(t *testing.T) {

		mockedJSON := `[{"entryPrice":"9975.12000","leverage":"50","maxNotional":"1000000","liquidationPrice":"7963.54","markPrice":"9973.50770517","positionAmount":"0.010","symbol":"BTCUSDT","unrealizedProfit":"-0.01612295"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/positionRisk", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesPositionRiskOfSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccount(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesPositionRiskOfSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccountV2 Success", func(t *testing.T) {

		mockedJSON := `{"futurePositionRiskVos":[{"entryPrice":"9975.12000","leverage":"50","maxNotional":"1000000","liquidationPrice":"7963.54","markPrice":"9973.50770517","positionAmount":"0.010","symbol":"BTCUSDT","unrealizedProfit":"-0.01612295"}],"deliveryPositionRiskVos":[{"entryPrice":"9975.12000","markPrice":"9973.50770517","leverage":"20","isolated":"false","isolatedWallet":"9973.50770517","isolatedMargin":"0.00000000","isAutoAddMargin":"false","positionSide":"BOTH","positionAmount":"1.230","symbol":"BTCUSD_201225","unrealizedProfit":"-0.01612295"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/sub-account/futures/positionRisk", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "789", r.URL.Query().Get("futuresType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesPositionRiskOfSubAccountV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccountV2(context.Background()).Email("sub-account-email@email.com").FuturesType(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesPositionRiskOfSubAccountV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesPositionRiskOfSubAccountV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccountV2 Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService GetFuturesPositionRiskOfSubAccountV2 Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.GetFuturesPositionRiskOfSubAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService GetSubAccountsStatusOnMarginOrFutures Success", func(t *testing.T) {

		mockedJSON := `[{"email":"123@test.com","isSubUserEnabled":true,"isUserActive":true,"insertTime":1570791523523,"isMarginEnabled":true,"isFutureEnabled":true,"mobile":1570791523523}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/status", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSubAccountsStatusOnMarginOrFuturesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.GetSubAccountsStatusOnMarginOrFutures(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSubAccountsStatusOnMarginOrFuturesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSubAccountsStatusOnMarginOrFuturesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService GetSubAccountsStatusOnMarginOrFutures Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.GetSubAccountsStatusOnMarginOrFutures(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService QuerySubAccountList Success", func(t *testing.T) {

		mockedJSON := `{"subAccounts":[{"subUserId":123456,"email":"testsub@gmail.com","remark":"remark","isFreeze":false,"createTime":1544433328000,"isManagedSubAccount":false,"isAssetManagementSubAccount":false},{"subUserId":1234567,"email":"virtual@oxebmvfonoemail.com","remark":"remarks","isFreeze":false,"createTime":1544433328000,"isManagedSubAccount":false,"isAssetManagementSubAccount":false}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService QuerySubAccountList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountManagementAPIService QuerySubAccountTransactionStatistics Success", func(t *testing.T) {

		mockedJSON := `{"recent30BtcTotal":"0","recent30BtcFuturesTotal":"0","recent30BtcMarginTotal":"0","recent30BusdTotal":"0","recent30BusdFuturesTotal":"0","recent30BusdMarginTotal":"0","tradeInfoVos":[{"userId":1000138138384,"btc":0,"btcFutures":0,"btcMargin":0,"busd":0,"busdFutures":0,"busdMargin":0,"date":1676851200000},{"userId":1000138138384,"btc":0,"btcFutures":0,"btcMargin":0,"busd":0,"busdFutures":0,"busdMargin":0,"date":1677110400000},{"userId":1000138138384,"btc":0,"btcFutures":0,"btcMargin":0,"busd":0,"busdFutures":0,"busdMargin":0,"date":1677369600000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/transaction-statistics", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountTransactionStatisticsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountTransactionStatistics(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountTransactionStatisticsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountTransactionStatisticsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountManagementAPIService QuerySubAccountTransactionStatistics Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountManagementAPI.QuerySubAccountTransactionStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
