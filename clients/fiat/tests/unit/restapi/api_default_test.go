/*
Fiat REST API TEST

Testing DefaultAPIService

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

func Test_binancefiatrestapi_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPIService Deposit Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"success","data":{"orderId":"04595xxxxxxxxx37"}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/deposit", r.URL.Path)
			require.Equal(t, "currency_example", r.URL.Query().Get("currency"))
			require.Equal(t, string(models.DepositRequestApiPaymentMethodPix), r.URL.Query().Get("apiPaymentMethod"))
			require.Equal(t, "amount_example", r.URL.Query().Get("amount"))
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

		resp, err := apiClient.RestApi.DefaultAPI.Deposit(context.Background()).Currency("currency_example").ApiPaymentMethod(models.DepositRequestApiPaymentMethodPix).Amount("amount_example").Execute()
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

	t.Run("Test DefaultAPIService Deposit Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.Deposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService Deposit Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.DefaultAPI.Deposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService FiatWithdraw Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"success","data":{"orderId":"04595xxxxxxxxx37"}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/fiat/withdraw", r.URL.Path)
			require.Equal(t, "currency_example", r.URL.Query().Get("currency"))
			require.Equal(t, string(models.FiatWithdrawRequestApiPaymentMethodBankTransfer), r.URL.Query().Get("apiPaymentMethod"))
			require.Equal(t, "789", r.URL.Query().Get("amount"))
			var got models.FiatWithdrawRequestAccountInfo
			err := json.Unmarshal([]byte(r.URL.Query().Get("accountInfo")), &got)
			require.NoError(t, err)
			if got.AdditionalProperties != nil && len(got.AdditionalProperties) == 0 {
				got.AdditionalProperties = nil
			}
			require.Equal(t, *models.NewFiatWithdrawRequestAccountInfo("1056894222"), got)
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

		resp, err := apiClient.RestApi.DefaultAPI.FiatWithdraw(context.Background()).Currency("currency_example").ApiPaymentMethod(models.FiatWithdrawRequestApiPaymentMethodBankTransfer).Amount(int64(789)).AccountInfo(*models.NewFiatWithdrawRequestAccountInfo("1056894222")).Execute()
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

	t.Run("Test DefaultAPIService FiatWithdraw Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.FiatWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService FiatWithdraw Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.DefaultAPI.FiatWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetFiatDepositWithdrawHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"success","data":[{"orderNo":"7d76d611-0568-4f43-afb6-24cac7767365","fiatCurrency":"BRL","indicatedAmount":"10.00","amount":"10.00","totalFee":"0.00","method":"BankAccount","status":"Expired","createTime":1626144956000,"updateTime":1626400907000}],"total":1,"success":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/orders", r.URL.Path)
			require.Equal(t, "0", r.URL.Query().Get("transactionType"))
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

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatDepositWithdrawHistory(context.Background()).TransactionType("0").Execute()
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

	t.Run("Test DefaultAPIService GetFiatDepositWithdrawHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatDepositWithdrawHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetFiatDepositWithdrawHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatDepositWithdrawHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetFiatPaymentsHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"success","data":[{"orderNo":"353fca443f06466db0c4dc89f94f027a","sourceAmount":"20.0","fiatCurrency":"EUR","obtainAmount":"4.462","cryptoCurrency":"LUNA","totalFee":"0.2","price":"4.437472","status":"Failed","paymentMethod":"Credit Card","createTime":1624529919000,"updateTime":1624529919000}],"total":1,"success":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/payments", r.URL.Path)
			require.Equal(t, "0", r.URL.Query().Get("transactionType"))
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

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatPaymentsHistory(context.Background()).TransactionType("0").Execute()
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

	t.Run("Test DefaultAPIService GetFiatPaymentsHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatPaymentsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetFiatPaymentsHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.DefaultAPI.GetFiatPaymentsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetOrderDetail Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"success","data":{"orderId":"036752*678","orderStatus":"ORDER_INITIAL","amount":"4.33","fee":"0.43","fiatCurrency":"***","errorCode":"","errorMessage":"","ext":{}}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/fiat/get-order-detail", r.URL.Path)
			require.Equal(t, "036752*678", r.URL.Query().Get("orderNo"))
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

		resp, err := apiClient.RestApi.DefaultAPI.GetOrderDetail(context.Background()).OrderNo("036752*678").Execute()
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

	t.Run("Test DefaultAPIService GetOrderDetail Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceFiatClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.DefaultAPI.GetOrderDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test DefaultAPIService GetOrderDetail Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.DefaultAPI.GetOrderDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
