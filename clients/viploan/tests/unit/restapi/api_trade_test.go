/*
VIP Loan REST API TEST

Testing TradeAPIService

*/

package binanceviploanrestapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/viploan"
	"github.com/binance/binance-connector-go/clients/viploan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binanceviploanrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService VipLoanBorrow Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"loanAccountId":"12345678","requestId":"12345678","loanCoin":"BTC","isFlexibleRate":"Yes","loanAmount":"100.55","collateralAccountId":"12345678,12345678,12345678","collateralCoin":"BUSD,USDT,ETH","loanTerm":"30"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/borrow", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("loanAccountId"))
			require.Equal(t, "BTC", r.URL.Query().Get("loanCoin"))
			require.Equal(t, fmt.Sprintf("%v", float32(1.0)), r.URL.Query().Get("loanAmount"))
			require.Equal(t, "12345678,12345678,12345678", r.URL.Query().Get("collateralAccountId"))
			require.Equal(t, "BUSD,USDT,ETH", r.URL.Query().Get("collateralCoin"))
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

		resp, err := apiClient.RestApi.TradeAPI.VipLoanBorrow(context.Background()).LoanAccountId(int64(1)).LoanCoin("BTC").LoanAmount(float32(1.0)).CollateralAccountId("12345678,12345678,12345678").CollateralCoin("BUSD,USDT,ETH").IsFlexibleRate(true).Execute()
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

	t.Run("Test TradeAPIService VipLoanFixedRateBorrow Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"borrowCoin":"BUSD","borrowAmount":"100.5","actualReceivedAmount":"98.75","collateralCoin":"BNB,ETH,BTC","collateralAccountId":"12345,67890,13579","borrowInterestRate":"0.01501231","duration":"30Days","autoRepay":true,"orderId":123456789,"status":"Succeeds"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/fixed/borrow", r.URL.Path)
			require.Equal(t, "1212:0.12:100;3434:0.13:50", r.URL.Query().Get("supplyRequest"))
			require.Equal(t, "BUSD", r.URL.Query().Get("borrowCoin"))
			require.Equal(t, "30", r.URL.Query().Get("loanTerm"))
			require.Equal(t, "12345678", r.URL.Query().Get("borrowUid"))
			require.Equal(t, "BNB,ETH,BTC", r.URL.Query().Get("collateralCoin"))
			require.Equal(t, "12345,67890,13579", r.URL.Query().Get("collateralAccountId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VipLoanFixedRateBorrowResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanFixedRateBorrow(context.Background()).SupplyRequest("1212:0.12:100;3434:0.13:50").BorrowCoin("BUSD").LoanTerm(int64(30)).BorrowUid(int64(12345678)).CollateralCoin("BNB,ETH,BTC").CollateralAccountId("12345,67890,13579").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VipLoanFixedRateBorrowResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VipLoanFixedRateBorrowResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService VipLoanFixedRateBorrow Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.VipLoanFixedRateBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanFixedRateBorrow Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.TradeAPI.VipLoanFixedRateBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService VipLoanRenew Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"loanAccountId":"12345678","loanCoin":"BTC","loanAmount":"100.55","collateralAccountId":"12345677,12345678,12345679","collateralCoin":"BUSD,USDT,ETH","loanTerm":"30"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/renew", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("orderId"))
			require.Equal(t, "30", r.URL.Query().Get("loanTerm"))
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

		resp, err := apiClient.RestApi.TradeAPI.VipLoanRenew(context.Background()).OrderId(int64(1)).LoanTerm(int64(30)).Execute()
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

		var mockedJSON string
		mockedJSON = `{"loanCoin":"BUSD","repayAmount":"200.5","remainingPrincipal":"100.5","remainingInterest":"0","collateralCoin":"BNB,BTC,ETH","currentLTV":"0.25","repayStatus":"Repaid"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/repay", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("orderId"))
			require.Equal(t, fmt.Sprintf("%v", float32(1.0)), r.URL.Query().Get("amount"))
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
