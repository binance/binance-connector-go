/*
Prediction Trading REST API TEST

Testing TransferAPIService

*/

package binancew3wpredictionrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancew3wpredictionrestapi_TransferAPIService(t *testing.T) {

	t.Run("Test TransferAPIService ApplyMmDeposit Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"transferId":"26080400000000454343","status":"PROCESSING"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/deposit/apply", r.URL.Path)
			require.Equal(t, "USDT", r.URL.Query().Get("fromToken"))
			require.Equal(t, "1000000000000000000", r.URL.Query().Get("fromTokenAmount"))
			require.Equal(t, "USDT", r.URL.Query().Get("toToken"))
			require.Equal(t, string(models.PlaceOrderAccountTypeParameterSpot), r.URL.Query().Get("accountType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ApplyMmDepositResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmDeposit(context.Background()).FromToken("USDT").FromTokenAmount("1000000000000000000").ToToken("USDT").AccountType(models.PlaceOrderAccountTypeParameterSpot).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ApplyMmDepositResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ApplyMmDepositResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService ApplyMmDeposit Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmDeposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService ApplyMmDeposit Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmDeposit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService ApplyMmWithdraw Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"id":"123456789","walletId":"0ecd3be8e3674b54b1258e24f9c3706b","walletAddress":"0x2ac3a1fb164da5c4e76f67ae6dbe6be821c000c8","transferId":"26080400000000454344"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/withdraw/apply", r.URL.Path)
			require.Equal(t, "USDT", r.URL.Query().Get("coin"))
			require.Equal(t, "BEP20", r.URL.Query().Get("network"))
			require.Equal(t, "100.00", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ApplyMmWithdrawResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmWithdraw(context.Background()).Coin("USDT").Network("BEP20").Amount("100.00").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ApplyMmWithdrawResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ApplyMmWithdrawResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService ApplyMmWithdraw Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService ApplyMmWithdraw Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.ApplyMmWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService CreateInboundTransfer Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"transferId":"tf_20260525_in_001","status":"PROCESSING"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/transfer/inbound", r.URL.Path)
			require.Equal(t, "5b5c1ec3be4e4416a5872b21c1ca5d20", r.URL.Query().Get("walletId"))
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "1000000000000000000", r.URL.Query().Get("fromTokenAmount"))
			require.Equal(t, string(models.PlaceOrderAccountTypeParameterSpot), r.URL.Query().Get("accountType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateInboundTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateInboundTransfer(context.Background()).WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").FromTokenAmount("1000000000000000000").AccountType(models.PlaceOrderAccountTypeParameterSpot).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateInboundTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateInboundTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService CreateInboundTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateInboundTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService CreateInboundTransfer Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateInboundTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService CreateOutboundTransfer Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"transferId":"tf_20260525_out_001","status":"PROCESSING"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/transfer/outbound", r.URL.Path)
			require.Equal(t, "5b5c1ec3be4e4416a5872b21c1ca5d20", r.URL.Query().Get("walletId"))
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "1000000000000000000", r.URL.Query().Get("fromTokenAmount"))
			require.Equal(t, string(models.PlaceOrderAccountTypeParameterSpot), r.URL.Query().Get("accountType"))
			require.Equal(t, string(models.CreateOutboundTransferSourceBizParameterUserTransfer), r.URL.Query().Get("sourceBiz"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CreateOutboundTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateOutboundTransfer(context.Background()).WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").FromTokenAmount("1000000000000000000").AccountType(models.PlaceOrderAccountTypeParameterSpot).SourceBiz(models.CreateOutboundTransferSourceBizParameterUserTransfer).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CreateOutboundTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CreateOutboundTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService CreateOutboundTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateOutboundTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService CreateOutboundTransfer Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.CreateOutboundTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryTransferList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"transfers":[{"transferId":"tf_20260525_out_001","direction":"OUTBOUND","status":"SUCCESS","walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","fromToken":"USDT","fromTokenAmount":"100.00","toToken":"USDT","toTokenAmount":"100.00","errorCode":"errorCode","errorMessage":"errorMessage","createTime":"2026-05-25T04:00:00.000+00:00","updateTime":"2026-05-25T04:00:05.000+00:00","completeAt":"2026-05-25T04:00:05.000+00:00"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/transfer/list", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "2026-05-01", r.URL.Query().Get("startDate"))
			require.Equal(t, "2026-05-25", r.URL.Query().Get("endDate"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryTransferListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferList(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").StartDate("2026-05-01").EndDate("2026-05-25").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryTransferListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryTransferListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService QueryTransferList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryTransferList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryTransferStatus Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"transferId":"tf_20260525_out_001","direction":"OUTBOUND","status":"COMPLETED","fromToken":"USDT","fromTokenAmount":"100.00","toToken":"USDT","toTokenAmount":"100.00","errorCode":"errorCode","errorMessage":"errorMessage","createTime":"2026-05-25T04:00:00.000+00:00","updateTime":"2026-05-25T04:00:05.000+00:00"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/transfer/status", r.URL.Path)
			require.Equal(t, "tf_20260525_out_001", r.URL.Query().Get("transferId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryTransferStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferStatus(context.Background()).TransferId("tf_20260525_out_001").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryTransferStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryTransferStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TransferAPIService QueryTransferStatus Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TransferAPIService QueryTransferStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TransferAPI.QueryTransferStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
