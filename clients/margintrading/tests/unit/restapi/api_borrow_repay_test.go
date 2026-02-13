/*
Binance Margin Trading REST API TEST

Testing BorrowRepayAPIService

*/

package binancemargintradingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/margintrading"
	"github.com/binance/binance-connector-go/clients/margintrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancemargintradingrestapi_BorrowRepayAPIService(t *testing.T) {

	t.Run("Test BorrowRepayAPIService GetFutureHourlyInterestRate Success", func(t *testing.T) {

		mockedJSON := `[{"asset":"BTC","nextHourlyInterestRate":"0.00000571"},{"asset":"ETH","nextHourlyInterestRate":"0.00000578"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/next-hourly-interest-rate", r.URL.Path)
			require.Equal(t, "assets_example", r.URL.Query().Get("assets"))
			require.Equal(t, "false", r.URL.Query().Get("isIsolated"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFutureHourlyInterestRateResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.GetFutureHourlyInterestRate(context.Background()).Assets("assets_example").IsIsolated(false).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFutureHourlyInterestRateResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFutureHourlyInterestRateResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService GetFutureHourlyInterestRate Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.GetFutureHourlyInterestRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService GetFutureHourlyInterestRate Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.GetFutureHourlyInterestRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService GetInterestHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"txId":1352286576452864800,"interestAccuredTime":1672160400000,"asset":"USDT","rawAsset":"USDT","principal":"45.3313","interest":"0.00024995","interestRate":"0.00013233","type":"ON_BORROW","isolatedSymbol":"BNBUSDT"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/interestHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetInterestHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.GetInterestHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetInterestHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetInterestHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService GetInterestHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.GetInterestHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService MarginAccountBorrowRepay Success", func(t *testing.T) {

		mockedJSON := `{"tranId":100000001}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/borrow-repay", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "FALSE", r.URL.Query().Get("isIsolated"))
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, "amount_example", r.URL.Query().Get("amount"))
			require.Equal(t, "type__example", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginAccountBorrowRepayResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.MarginAccountBorrowRepay(context.Background()).Asset("asset_example").IsIsolated("FALSE").Symbol("symbol_example").Amount("amount_example").Type("type__example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginAccountBorrowRepayResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginAccountBorrowRepayResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService MarginAccountBorrowRepay Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.MarginAccountBorrowRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService MarginAccountBorrowRepay Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.MarginAccountBorrowRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryBorrowRepayRecordsInMarginAccount Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"type":"AUTO","isolatedSymbol":"BNBUSDT","amount":"14.00000000","asset":"BNB","interest":"0.01866667","principal":"13.98133333","status":"CONFIRMED","timestamp":1563438204000,"txId":2970933056}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/borrow-repay", r.URL.Path)
			require.Equal(t, "type__example", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryBorrowRepayRecordsInMarginAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryBorrowRepayRecordsInMarginAccount(context.Background()).Type("type__example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryBorrowRepayRecordsInMarginAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryBorrowRepayRecordsInMarginAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService QueryBorrowRepayRecordsInMarginAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryBorrowRepayRecordsInMarginAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryBorrowRepayRecordsInMarginAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryBorrowRepayRecordsInMarginAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryMarginInterestRateHistory Success", func(t *testing.T) {

		mockedJSON := `[{"asset":"BTC","dailyInterestRate":"0.00025000","timestamp":1611544731000,"vipLevel":1},{"asset":"BTC","dailyInterestRate":"0.00035000","timestamp":1610248118000,"vipLevel":1}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/interestRateHistory", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMarginInterestRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMarginInterestRateHistory(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMarginInterestRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMarginInterestRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService QueryMarginInterestRateHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMarginInterestRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryMarginInterestRateHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMarginInterestRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryMaxBorrow Success", func(t *testing.T) {

		mockedJSON := `{"amount":"1.69248805","borrowLimit":"60"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/margin/maxBorrowable", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryMaxBorrowResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMaxBorrow(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryMaxBorrowResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryMaxBorrowResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test BorrowRepayAPIService QueryMaxBorrow Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceMarginTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMaxBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test BorrowRepayAPIService QueryMaxBorrow Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.BorrowRepayAPI.QueryMaxBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
