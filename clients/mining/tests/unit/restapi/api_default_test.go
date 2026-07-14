/*
Mining REST API TEST

Testing DefaultAPIService

*/

package binanceminingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/mining"
	"github.com/binance/binance-connector-go/clients/mining/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binanceminingrestapi_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPIService AccountList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":[{"type":"H_hashrate","userName":"test","list":[{"time":1585267200000,"hashrate":"0.00000000","reject":"0.00000000"}]}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/statistics/user/list", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AccountList(context.Background()).Algo("sha256").UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService AccountList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AccountList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AcquiringAlgorithm Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":[{"algoName":"SHA256-BTC","algoId":1,"poolIndex":100,"unit":"h/s"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/pub/algoList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AcquiringAlgorithmResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AcquiringAlgorithm(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AcquiringAlgorithmResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AcquiringAlgorithmResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService AcquiringAlgorithm Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AcquiringAlgorithm(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService AcquiringCoinname Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":[{"coinName":"BTC","coinId":1,"poolIndex":0,"algoId":1,"algoName":"SHA256-BTC"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/pub/coinList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AcquiringCoinnameResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AcquiringCoinname(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AcquiringCoinnameResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AcquiringCoinnameResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService AcquiringCoinname Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.AcquiringCoinname(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService CancelHashrateResaleConfiguration Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/config/cancel", r.URL.Path)
			require.Equal(t, "168", r.URL.Query().Get("configId"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CancelHashrateResaleConfigurationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.CancelHashrateResaleConfiguration(context.Background()).ConfigId(int64(168)).UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CancelHashrateResaleConfigurationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CancelHashrateResaleConfigurationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService CancelHashrateResaleConfiguration Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.CancelHashrateResaleConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService CancelHashrateResaleConfiguration Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.CancelHashrateResaleConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService EarningsList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"accountProfits":[{"time":1586188800000,"type":31,"hashTransfer":200000000000,"transferAmount":0.02180958,"dayHashRate":129129903378244,"profitAmount":8.6083060304,"coinName":"BTC","status":2}],"totalNum":3,"pageSize":20}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/list", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.EarningsListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.EarningsList(context.Background()).Algo("sha256").UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.EarningsListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.EarningsListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService EarningsList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.EarningsList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService EarningsList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.EarningsList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ExtraBonusList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"otherProfits":[{"time":1607443200000,"coinName":"BTC","type":4,"profitAmount":0.0011859,"status":2}],"totalNum":3,"pageSize":20}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/other", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ExtraBonusListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.ExtraBonusList(context.Background()).Algo("sha256").UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ExtraBonusListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ExtraBonusListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService ExtraBonusList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.ExtraBonusList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService ExtraBonusList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.ExtraBonusList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService HashrateResaleDetail Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"profitTransferDetails":[{"poolUsername":"test4001","toPoolUsername":"pop","algoName":"sha256","hashRate":200000000000,"day":20201213,"amount":0.2256872,"coinName":"BTC"}],"totalNum":8,"pageSize":200}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/profit/details", r.URL.Path)
			require.Equal(t, "168", r.URL.Query().Get("configId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.HashrateResaleDetailResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleDetail(context.Background()).ConfigId(int64(168)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.HashrateResaleDetailResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.HashrateResaleDetailResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService HashrateResaleDetail Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService HashrateResaleDetail Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService HashrateResaleList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"configDetails":[{"configId":168,"poolUsername":"123","toPoolUsername":"user1","algoName":"Ethash","hashRate":5000000,"startDay":20201210,"endDay":20210405,"status":1,"type":0}],"totalNum":21,"pageSize":200}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/config/details/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.HashrateResaleListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.HashrateResaleListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.HashrateResaleListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService HashrateResaleList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService HashrateResaleRequest Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":171}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/config", r.URL.Path)
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "1770736694138", r.URL.Query().Get("endDate"))
			require.Equal(t, "1770736694138", r.URL.Query().Get("startDate"))
			require.Equal(t, "S19pro", r.URL.Query().Get("toPoolUser"))
			require.Equal(t, "100000000", r.URL.Query().Get("hashRate"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.HashrateResaleRequestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleRequest(context.Background()).UserName("test").Algo("sha256").EndDate(int64(1770736694138)).StartDate(int64(1770736694138)).ToPoolUser("S19pro").HashRate(int64(100000000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.HashrateResaleRequestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.HashrateResaleRequestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService HashrateResaleRequest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService HashrateResaleRequest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.HashrateResaleRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService MiningAccountEarning Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"accountProfits":[{"time":1607443200000,"coinName":"BTC","type":2,"puid":59985472,"subName":"vdvaghani","amount":0.09186957}],"totalNum":3,"pageSize":20}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/uid", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MiningAccountEarningResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.MiningAccountEarning(context.Background()).Algo("sha256").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MiningAccountEarningResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MiningAccountEarningResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService MiningAccountEarning Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.MiningAccountEarning(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService MiningAccountEarning Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.MiningAccountEarning(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RequestForDetailMinerList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":[{"workerName":"bhdc1.16A10404B","type":"H_hashrate","hashrateDatas":[{"time":1587902400000,"hashrate":"0","reject":0}]}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/worker/detail", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			require.Equal(t, "bhdc1.16A10404B", r.URL.Query().Get("workerName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RequestForDetailMinerListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForDetailMinerList(context.Background()).Algo("sha256").UserName("test").WorkerName("bhdc1.16A10404B").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RequestForDetailMinerListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RequestForDetailMinerListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService RequestForDetailMinerList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForDetailMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RequestForDetailMinerList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForDetailMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RequestForMinerList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"workerDatas":[{"workerId":"1420554439452400131","workerName":"2X73","status":3,"hashRate":0,"dayHashRate":1.269778129801366E13,"rejectRate":0,"lastShareTime":1587712919000}],"totalNum":18530,"pageSize":20}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/worker/list", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RequestForMinerListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForMinerList(context.Background()).Algo("sha256").UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RequestForMinerListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RequestForMinerListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService RequestForMinerList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService RequestForMinerList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.RequestForMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService StatisticList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":0,"msg":"","data":{"fifteenMinHashRate":"457835490067496409.00000000","dayHashRate":"214289268068874127.65000000","validNum":0,"invalidNum":17562,"profitToday":{"BTC":"0.00314332","BSV":"56.17055953","BCH":"106.61586001"},"profitYesterday":{"BTC":"0.00314332","BSV":"56.17055953","BCH":"106.61586001"},"userName":"test","unit":"h/s","algo":"sha256"}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/statistics/user/status", r.URL.Path)
			require.Equal(t, "sha256", r.URL.Query().Get("algo"))
			require.Equal(t, "test", r.URL.Query().Get("userName"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.StatisticListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.StatisticList(context.Background()).Algo("sha256").UserName("test").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.StatisticListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.StatisticListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test DefaultAPIService StatisticList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.StatisticList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService StatisticList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.StatisticList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
