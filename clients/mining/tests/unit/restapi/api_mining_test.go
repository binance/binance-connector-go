/*
Binance Mining REST API TEST

Testing MiningAPIService

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
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binanceminingrestapi_MiningAPIService(t *testing.T) {

	t.Run("Test MiningAPIService AccountList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":[{"type":"H_hashrate","userName":"test","list":[{"time":1585267200000,"hashrate":"0.00000000","reject":"0.00000000"},{"time":1585353600000,"hashrate":"0.00000000","reject":"0.00000000"}]},{"type":"D_hashrate","userName":"test","list":[{"time":1587906000000,"hashrate":"0.00000000","reject":"0.00000000"},{"time":1587909600000,"hashrate":"0.00000000","reject":"0.00000000"}]}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/statistics/user/list", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.AccountList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService AccountList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.AccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService AccountList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.AccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService AcquiringAlgorithm Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":[{"algoName":"sha256","algoId":1,"poolIndex":0,"unit":"h/s"}]}`
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

		resp, err := apiClient.RestApi.MiningAPI.AcquiringAlgorithm(context.Background()).Execute()
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

	t.Run("Test MiningAPIService AcquiringAlgorithm Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.AcquiringAlgorithm(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService AcquiringCoinname Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":[{"coinName":"BTC","coinId":1,"poolIndex":0,"algoId":1,"algoName":"sha256"}]}`
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

		resp, err := apiClient.RestApi.MiningAPI.AcquiringCoinname(context.Background()).Execute()
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

	t.Run("Test MiningAPIService AcquiringCoinname Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.AcquiringCoinname(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService CancelHashrateResaleConfiguration Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/config/cancel", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("configId"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.CancelHashrateResaleConfiguration(context.Background()).ConfigId(int64(1)).UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService CancelHashrateResaleConfiguration Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.CancelHashrateResaleConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService CancelHashrateResaleConfiguration Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.CancelHashrateResaleConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService EarningsList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"accountProfits":[{"time":1586188800000,"type":31,"hashTransfer":null,"transferAmount":null,"dayHashRate":129129903378244,"profitAmount":8.6083060304,"coinName":"BTC","status":2},{"time":1607529600000,"coinName":"BTC","type":0,"dayHashRate":9942053925926,"profitAmount":0.85426469,"hashTransfer":200000000000,"transferAmount":0.02180958,"status":2},{"time":1607443200000,"coinName":"BTC","type":31,"dayHashRate":200000000000,"profitAmount":0.02905916,"hashTransfer":null,"transferAmount":null,"status":2}],"totalNum":3,"pageSize":20}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/list", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.EarningsList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService EarningsList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.EarningsList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService EarningsList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.EarningsList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService ExtraBonusList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"otherProfits":[{"time":1607443200000,"coinName":"BTC","type":4,"profitAmount":0.0011859,"status":2}],"totalNum":3,"pageSize":20}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/other", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.ExtraBonusList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService ExtraBonusList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.ExtraBonusList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService ExtraBonusList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.ExtraBonusList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService HashrateResaleDetail Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"profitTransferDetails":[{"poolUsername":"test4001","toPoolUsername":"pop","algoName":"sha256","hashRate":200000000000,"day":20201213,"amount":0.2256872,"coinName":"BTC"},{"poolUsername":"test4001","toPoolUsername":"pop","algoName":"sha256","hashRate":200000000000,"day":20201213,"amount":0.2256872,"coinName":"BTC"}],"totalNum":8,"pageSize":200}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/profit/details", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("configId"))
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleDetail(context.Background()).ConfigId(int64(1)).Execute()
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

	t.Run("Test MiningAPIService HashrateResaleDetail Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService HashrateResaleDetail Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService HashrateResaleList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"configDetails":[{"configId":168,"poolUsername":"123","toPoolUsername":"user1","algoName":"Ethash","hashRate":5000000,"startDay":20201210,"endDay":20210405,"status":1,"type":0},{"configId":166,"poolUsername":"pop","toPoolUsername":"111111","algoName":"Ethash","hashRate":3320000,"startDay":20201226,"endDay":20201227,"status":0,"type":0}],"totalNum":21,"pageSize":200}}`
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleList(context.Background()).Execute()
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

	t.Run("Test MiningAPIService HashrateResaleList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService HashrateResaleRequest Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":171}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/hash-transfer/config", r.URL.Path)
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "789", r.URL.Query().Get("endDate"))
			require.Equal(t, "789", r.URL.Query().Get("startDate"))
			require.Equal(t, "toPoolUser_example", r.URL.Query().Get("toPoolUser"))
			require.Equal(t, "789", r.URL.Query().Get("hashRate"))
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleRequest(context.Background()).UserName("userName_example").Algo("algo_example").EndDate(int64(789)).StartDate(int64(789)).ToPoolUser("toPoolUser_example").HashRate(int64(789)).Execute()
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

	t.Run("Test MiningAPIService HashrateResaleRequest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService HashrateResaleRequest Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.HashrateResaleRequest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService MiningAccountEarning Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"accountProfits":[{"time":1607443200000,"coinName":"BTC","type":2,"puid":59985472,"subName":"vdvaghani","amount":0.09186957}],"totalNum":3,"pageSize":20}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/payment/uid", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
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

		resp, err := apiClient.RestApi.MiningAPI.MiningAccountEarning(context.Background()).Algo("algo_example").Execute()
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

	t.Run("Test MiningAPIService MiningAccountEarning Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.MiningAccountEarning(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService MiningAccountEarning Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.MiningAccountEarning(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService RequestForDetailMinerList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":[{"workerName":"bhdc1.16A10404B","type":"H_hashrate","hashrateDatas":[{"time":1587902400000,"hashrate":"0","reject":0},{"time":1587906000000,"hashrate":"0","reject":0}]},{"workerName":"bhdc1.16A10404B","type":"D_hashrate","hashrateDatas":[{"time":1587902400000,"hashrate":"0","reject":0},{"time":1587906000000,"hashrate":"0","reject":0}]}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/worker/detail", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
			require.Equal(t, "workerName_example", r.URL.Query().Get("workerName"))
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

		resp, err := apiClient.RestApi.MiningAPI.RequestForDetailMinerList(context.Background()).Algo("algo_example").UserName("userName_example").WorkerName("workerName_example").Execute()
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

	t.Run("Test MiningAPIService RequestForDetailMinerList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.RequestForDetailMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService RequestForDetailMinerList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.RequestForDetailMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService RequestForMinerList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"workerDatas":[{"workerId":"1420554439452400131","workerName":"2X73","status":3,"hashRate":0,"dayHashRate":0,"rejectRate":0,"lastShareTime":1587712919000},{"workerId":"7893926126382807951","workerName":"AZDC1.1A10101","status":2,"hashRate":29711247541680,"dayHashRate":1.269778129801366E13,"rejectRate":0,"lastShareTime":1587969727000}],"totalNum":18530,"pageSize":20}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/worker/list", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.RequestForMinerList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService RequestForMinerList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.RequestForMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService RequestForMinerList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.RequestForMinerList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService StatisticList Success", func(t *testing.T) {

		mockedJSON := `{"code":0,"msg":"","data":{"fifteenMinHashRate":"457835490067496409.00000000","dayHashRate":"214289268068874127.65000000","validNum":0,"invalidNum":17562,"profitToday":{"BTC":"0.00314332","BSV":"56.17055953","BCH":"106.61586001"},"profitYesterday":{"BTC":"0.00314332","BSV":"56.17055953","BCH":"106.61586001"},"userName":"test","unit":"h/s","algo":"sha256"}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/mining/statistics/user/status", r.URL.Path)
			require.Equal(t, "algo_example", r.URL.Query().Get("algo"))
			require.Equal(t, "userName_example", r.URL.Query().Get("userName"))
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

		resp, err := apiClient.RestApi.MiningAPI.StatisticList(context.Background()).Algo("algo_example").UserName("userName_example").Execute()
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

	t.Run("Test MiningAPIService StatisticList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMiningClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MiningAPI.StatisticList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MiningAPIService StatisticList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MiningAPI.StatisticList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
