/*
Binance Crypto Loan REST API TEST

Testing FlexibleRateAPIService

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

func Test_binancecryptoloanrestapi_FlexibleRateAPIService(t *testing.T) {

	t.Run("Test FlexibleRateAPIService CheckCollateralRepayRate Success", func(t *testing.T) {

		mockedJSON := `{"loanCoin":"BUSD","collateralCoin":"BNB","rate":"300.36781234"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/repay/rate", r.URL.Path)
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckCollateralRepayRateResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.CheckCollateralRepayRate(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CheckCollateralRepayRateResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CheckCollateralRepayRateResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService CheckCollateralRepayRate Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.CheckCollateralRepayRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService CheckCollateralRepayRate Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.CheckCollateralRepayRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanAdjustLtv Success", func(t *testing.T) {

		mockedJSON := `{"loanCoin":"BUSD","collateralCoin":"BNB","direction":"ADDITIONAL","adjustmentAmount":"5.235","currentLTV":"0.52","status":"Succeeds"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/adjust/ltv", r.URL.Path)
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			require.Equal(t, "1", r.URL.Query().Get("adjustmentAmount"))
			require.Equal(t, "direction_example", r.URL.Query().Get("direction"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FlexibleLoanAdjustLtvResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanAdjustLtv(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").AdjustmentAmount(float32(1.0)).Direction("direction_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FlexibleLoanAdjustLtvResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FlexibleLoanAdjustLtvResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanAdjustLtv Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanAdjustLtv(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanAdjustLtv Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanAdjustLtv(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanBorrow Success", func(t *testing.T) {

		mockedJSON := `{"loanCoin":"BUSD","loanAmount":"100.5","collateralCoin":"BNB","collateralAmount":"50.5","status":"Succeeds"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/borrow", r.URL.Path)
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FlexibleLoanBorrowResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanBorrow(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FlexibleLoanBorrowResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FlexibleLoanBorrowResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanBorrow Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanBorrow Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanBorrow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanRepay Success", func(t *testing.T) {

		mockedJSON := `{"loanCoin":"BUSD","collateralCoin":"BNB","remainingDebt":"100.5","remainingCollateral":"5.253","fullRepayment":false,"currentLTV":"0.25","repayStatus":"REPAID"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/repay", r.URL.Path)
			require.Equal(t, "loanCoin_example", r.URL.Query().Get("loanCoin"))
			require.Equal(t, "collateralCoin_example", r.URL.Query().Get("collateralCoin"))
			require.Equal(t, "1", r.URL.Query().Get("repayAmount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FlexibleLoanRepayResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanRepay(context.Background()).LoanCoin("loanCoin_example").CollateralCoin("collateralCoin_example").RepayAmount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FlexibleLoanRepayResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FlexibleLoanRepayResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanRepay Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService FlexibleLoanRepay Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.FlexibleLoanRepay(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanAssetsData Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","flexibleInterestRate":"0.00000491","flexibleMinLimit":"100","flexibleMaxLimit":"1000000"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/loanable/data", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanAssetsDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanAssetsData(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanAssetsDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanAssetsDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanAssetsData Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanAssetsData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanBorrowHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","initialLoanAmount":"10000","collateralCoin":"BNB","initialCollateralAmount":"49.27565492","borrowTime":1575018510000,"status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/borrow/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanBorrowHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanBorrowHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanBorrowHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanBorrowHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanBorrowHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanBorrowHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanCollateralAssetsData Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"collateralCoin":"BNB","initialLTV":"0.65","marginCallLTV":"0.75","liquidationLTV":"0.83","maxLimit":"1000000"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/collateral/data", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanCollateralAssetsDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanCollateralAssetsData(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanCollateralAssetsDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanCollateralAssetsDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanCollateralAssetsData Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanCollateralAssetsData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanInterestRateHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"coin":"USDT","annualizedInterestRate":"0.0647","time":1575018510000},{"coin":"USDT","annualizedInterestRate":"0.0647","time":1575018510000}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/interestRateHistory", r.URL.Path)
			require.Equal(t, "coin_example", r.URL.Query().Get("coin"))
			require.Equal(t, "5000", r.URL.Query().Get("recvWindow"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanInterestRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanInterestRateHistory(context.Background()).Coin("coin_example").RecvWindow(int64(5000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanInterestRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanInterestRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanInterestRateHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanInterestRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanInterestRateHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanInterestRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanLiquidationHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","liquidationDebt":"10000","collateralCoin":"BNB","liquidationCollateralAmount":"123","returnCollateralAmount":"0.2","liquidationFee":"1.2","liquidationStartingPrice":"49.27565492","liquidationStartingTime":1575018510000,"status":"Liquidated"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/liquidation/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanLiquidationHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLiquidationHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanLiquidationHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanLiquidationHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanLiquidationHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLiquidationHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanLtvAdjustmentHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","collateralCoin":"BNB","direction":"ADDITIONAL","collateralAmount":"5.235","preLTV":"0.78","afterLTV":"0.56","adjustTime":1575018510000}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/ltv/adjustment/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanLtvAdjustmentHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLtvAdjustmentHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanLtvAdjustmentHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanLtvAdjustmentHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanLtvAdjustmentHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanLtvAdjustmentHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanOngoingOrders Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","totalDebt":"10000","collateralCoin":"BNB","collateralAmount":"49.27565492","currentLTV":"0.57"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/ongoing/orders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanOngoingOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanOngoingOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanOngoingOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanOngoingOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanOngoingOrders Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanOngoingOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanRepaymentHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"loanCoin":"BUSD","repayAmount":"10000","collateralCoin":"BNB","collateralReturn":"49.27565492","repayStatus":"REPAID","repayTime":1575018510000}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/loan/flexible/repay/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleLoanRepaymentHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCryptoLoanClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanRepaymentHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleLoanRepaymentHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleLoanRepaymentHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleRateAPIService GetFlexibleLoanRepaymentHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.FlexibleRateAPI.GetFlexibleLoanRepaymentHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
