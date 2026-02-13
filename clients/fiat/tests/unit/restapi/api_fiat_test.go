/*
Binance Fiat REST API TEST

Testing FiatAPIService

*/

package binancefiatrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/fiat"
	"github.com/binance/binance-connector-go/clients/fiat/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancefiatrestapi_FiatAPIService(t *testing.T) {

	t.Run("Test FiatAPIService Deposit Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":{"orderId":"04595xxxxxxxxx37"}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/deposit", r.URL.Path)
			require.Equal(t, "currency_example", r.URL.Query().Get("currency"))
			require.Equal(t, "apiPaymentMethod_example", r.URL.Query().Get("apiPaymentMethod"))
			require.Equal(t, "789", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DepositResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.Deposit(context.Background()).Currency("currency_example").ApiPaymentMethod("apiPaymentMethod_example").Amount(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DepositResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DepositResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FiatAPIService Deposit Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.Deposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService Deposit Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.Deposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService FiatWithdraw Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":{"orderId":"04595xxxxxxxxx37"}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/fiat/withdraw", r.URL.Path)
			require.Equal(t, "currency_example", r.URL.Query().Get("currency"))
			require.Equal(t, "apiPaymentMethod_example", r.URL.Query().Get("apiPaymentMethod"))
			require.Equal(t, "789", r.URL.Query().Get("amount"))
			require.Equal(t, string("{}"), r.URL.Query().Get("accountInfo"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FiatWithdrawResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.FiatWithdraw(context.Background()).Currency("currency_example").ApiPaymentMethod("apiPaymentMethod_example").Amount(int64(789)).AccountInfo(*models.NewAccountInfo()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FiatWithdrawResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FiatWithdrawResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FiatAPIService FiatWithdraw Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.FiatWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService FiatWithdraw Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.FiatWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetFiatDepositWithdrawHistory Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":[{"orderNo":"7d76d611-0568-4f43-afb6-24cac7767365","fiatCurrency":"BRL","indicatedAmount":"10.00","amount":"10.00","totalFee":"0.00","method":"BankAccount","status":"Expired","createTime":1626144956000,"updateTime":1626400907000}],"total":1,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/orders", r.URL.Path)
			require.Equal(t, "transactionType_example", r.URL.Query().Get("transactionType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFiatDepositWithdrawHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType("transactionType_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFiatDepositWithdrawHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFiatDepositWithdrawHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FiatAPIService GetFiatDepositWithdrawHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetFiatDepositWithdrawHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatDepositWithdrawHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetFiatPaymentsHistory Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":[{"orderNo":"353fca443f06466db0c4dc89f94f027a","sourceAmount":"20.0","fiatCurrency":"EUR","obtainAmount":"4.462","cryptoCurrency":"LUNA","totalFee":"0.2","price":"4.437472","status":"Failed","paymentMethod":"Credit Card","createTime":1624529919000,"updateTime":1624529919000}],"total":1,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/payments", r.URL.Path)
			require.Equal(t, "transactionType_example", r.URL.Query().Get("transactionType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFiatPaymentsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatPaymentsHistory(context.Background()).TransactionType("transactionType_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFiatPaymentsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFiatPaymentsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FiatAPIService GetFiatPaymentsHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatPaymentsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetFiatPaymentsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetFiatPaymentsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetOrderDetail Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":{"orderId":"036752*678","orderStatus":"ORDER_INITIAL","amount":"4.33","fee":"0.43","fiatCurrency":"***","errorCode":"","errorMessage":"","ext":{}}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/get-order-detail", r.URL.Path)
			require.Equal(t, "orderNo_example", r.URL.Query().Get("orderNo"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOrderDetailResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetOrderDetail(context.Background()).OrderNo("orderNo_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOrderDetailResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOrderDetailResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FiatAPIService GetOrderDetail Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetOrderDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FiatAPIService GetOrderDetail Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FiatAPI.GetOrderDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
