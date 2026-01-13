/*
Binance VIP Loan REST API TEST

Testing TradeAPIService

*/

package binanceviploanrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binanceviploanrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService VipLoanBorrow Success", func(t *testing.T) {

		mockedJSON := `{"loanAccountId":"12345678","requestId":"12345678","loanCoin":"BTC","isFlexibleRate":"Yes","loanAmount":"100.55","collateralAccountId":"12345678,12345678,12345678","collateralCoin":"BUSD,USDT,ETH","loanTerm":"30"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/borrow", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("loanAccountId"))
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "1", r.URL.Query().Get("loanAmount"))
			require.Equal(t, "1", r.URL.Query().Get("collateralAccountId"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			require.Equal(t, "true", r.URL.Query().Get("isFlexibleRate"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VipLoanBorrowResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).LoanAccountId(int64(1)).LoanCoin("loanCoin_example").LoanAmount(float32(1.0)).CollateralAccountId("1").CollateralCoin("collateralCoin_example").IsFlexibleRate(true).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VipLoanBorrowResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VipLoanBorrowResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService VipLoanBorrow Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanBorrow Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanRenew Success", func(t *testing.T) {

		mockedJSON := `{"loanAccountId":"12345678","loanCoin":"BTC","loanAmount":"100.55","collateralAccountId":"12345677,12345678,12345679","collateralCoin":"BUSD,USDT,ETH","loanTerm":"30"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/renew", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("orderId"))
			require.Equal(t, "789", r.URL.Query().Get("loanTerm"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VipLoanRenewResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRenew(context.Background()).OrderId(int64(1)).LoanTerm(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VipLoanRenewResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VipLoanRenewResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService VipLoanRenew Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRenew(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanRenew Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRenew(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanRepay Success", func(t *testing.T) {

		mockedJSON := `{"loanCoin":"BUSD","repayAmount":"200.5","remainingPrincipal":"100.5","remainingInterest":"0","collateralCoin":"BNB,BTC,ETH","currentLTV":"0.25","repayStatus":"Repaid"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/repay", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("orderId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VipLoanRepayResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRepay(context.Background()).OrderId(int64(1)).Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VipLoanRepayResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VipLoanRepayResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService VipLoanRepay Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanRepay Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
