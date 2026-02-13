/*
Binance VIP Loan REST API TEST

Testing UserInformationAPIService

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
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binanceviploanrestapi_UserInformationAPIService(t *testing.T) {

	t.Run("Test UserInformationAPIService CheckVIPLoanCollateralAccount Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"collateralAccountId":"12345678","collateralCoin":"BNB,BTC,ETH"},{"collateralAccountId":"23456789","collateralCoin":"BNB,BTC,ETH"}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/collateral/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckVIPLoanCollateralAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserInformationAPI.CheckVIPLoanCollateralAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CheckVIPLoanCollateralAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CheckVIPLoanCollateralAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserInformationAPIService CheckVIPLoanCollateralAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.UserInformationAPI.CheckVIPLoanCollateralAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test UserInformationAPIService GetVIPLoanAccruedInterest Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"USDT","principalAmount":"10000","interestAmount":"1.2","annualInterestRate":"0.001273","accrualTime":1575018510000,"orderId":756783308056935400}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/accruedInterest", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetVIPLoanAccruedInterestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserInformationAPI.GetVIPLoanAccruedInterest(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetVIPLoanAccruedInterestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetVIPLoanAccruedInterestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserInformationAPIService GetVIPLoanAccruedInterest Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.UserInformationAPI.GetVIPLoanAccruedInterest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test UserInformationAPIService GetVIPLoanOngoingOrders Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"orderId":100000001,"loanCoin":"BUSD","totalDebt":"10000","loanRate":"0.0123","residualInterest":"10.27687923","collateralAccountId":"12345678,23456789","collateralCoin":"BNB,BTC,ETH","totalCollateralValueAfterHaircut":"25000.27565492","lockedCollateralValue":"25000.27565492","currentLTV":"0.57","expirationTime":1575018510000,"loanDate":"1676851200000","loanTerm":"30days"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/ongoing/orders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetVIPLoanOngoingOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserInformationAPI.GetVIPLoanOngoingOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetVIPLoanOngoingOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetVIPLoanOngoingOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserInformationAPIService GetVIPLoanOngoingOrders Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.UserInformationAPI.GetVIPLoanOngoingOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test UserInformationAPIService QueryApplicationStatus Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanAccountId":"12345678","orderId":"12345678","requestId":"12345678","loanCoin":"BTC","loanAmount":"100.55","collateralAccountId":"12345678,12345678,12345678","collateralCoin":"BUSD,USDT,ETH","loanTerm":"30","status":"Repaid","loanDate":"1676851200000"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/vip/request/data", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryApplicationStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceVipLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.UserInformationAPI.QueryApplicationStatus(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryApplicationStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryApplicationStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test UserInformationAPIService QueryApplicationStatus Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.UserInformationAPI.QueryApplicationStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
