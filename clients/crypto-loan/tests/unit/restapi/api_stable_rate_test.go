/*
Binance Crypto Loan REST API TEST

Testing StableRateAPIService

*/

package binancecryptoloanrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/cryptoloan"
	"github.com/binance/binance-connector-go/clients/cryptoloan/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancecryptoloanrestapi_StableRateAPIService(t *testing.T) {

	t.Run("Test StableRateAPIService CheckCollateralRepayRateStableRate Success", func(t *testing.T) {

		mockedJSON := `{"loanlCoin":"BUSD","collateralCoin":"BNB","repayAmount":"1000","rate":"300.36781234"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/repay/collateral/rate", r.URL.Path)
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			require.Equal(t, "1", r.URL.Query().Get("repayAmount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckCollateralRepayRateStableRateResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.CheckCollateralRepayRateStableRate(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").RepayAmount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CheckCollateralRepayRateStableRateResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CheckCollateralRepayRateStableRateResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test StableRateAPIService CheckCollateralRepayRateStableRate Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.CheckCollateralRepayRateStableRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test StableRateAPIService CheckCollateralRepayRateStableRate Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.CheckCollateralRepayRateStableRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test StableRateAPIService GetCryptoLoansIncomeHistory Success", func(t *testing.T) {

		mockedJSON := `[{"asset":"BUSD","type":"borrowIn","amount":"100","timestamp":1633771139847,"tranId":"80423589583"},{"asset":"BUSD","type":"borrowIn","amount":"100","timestamp":1634638371496,"tranId":"81685123491"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/income", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCryptoLoansIncomeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetCryptoLoansIncomeHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCryptoLoansIncomeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCryptoLoansIncomeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test StableRateAPIService GetCryptoLoansIncomeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetCryptoLoansIncomeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test StableRateAPIService GetLoanBorrowHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"orderId":100000001,"loanCoin":"BUSD","initialLoanAmount":"10000","hourlyInterestRate":"0.000057","loanTerm":"7","collateralCoin":"BNB","initialCollateralAmount":"49.27565492","borrowTime":1575018510000,"status":"Repaid"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/borrow/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLoanBorrowHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanBorrowHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLoanBorrowHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLoanBorrowHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test StableRateAPIService GetLoanBorrowHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanBorrowHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test StableRateAPIService GetLoanLtvAdjustmentHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","collateralCoin":"BNB","direction":"ADDITIONAL","amount":"5.235","preLTV":"0.78","afterLTV":"0.56","adjustTime":1575018510000,"orderId":756783308056935400}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/ltv/adjustment/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLoanLtvAdjustmentHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanLtvAdjustmentHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLoanLtvAdjustmentHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLoanLtvAdjustmentHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test StableRateAPIService GetLoanLtvAdjustmentHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanLtvAdjustmentHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test StableRateAPIService GetLoanRepaymentHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","repayAmount":"10000","collateralCoin":"BNB","collateralUsed":"0","collateralReturn":"49.27565492","repayType":"1","repayStatus":"Repaid","repayTime":1575018510000,"orderId":756783308056935400}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/loan/repay/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLoanRepaymentHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanRepaymentHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLoanRepaymentHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLoanRepaymentHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test StableRateAPIService GetLoanRepaymentHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.StableRateAPI.GetLoanRepaymentHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
