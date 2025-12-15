/*
Binance Sub Account REST API TEST

Testing ManagedSubAccountAPIService

*/

package binancesubaccountrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/subaccount/src"
	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesubaccountrestapi_ManagedSubAccountAPIService(t *testing.T) {

	t.Run("Test ManagedSubAccountAPIService DepositAssetsIntoTheManagedSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"tranId":66157362489}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/deposit", r.URL.Path)
			require.Equal(t, "toEmail_example", r.URL.Query().Get("toEmail"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DepositAssetsIntoTheManagedSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.DepositAssetsIntoTheManagedSubAccount(context.Background()).ToEmail("toEmail_example").Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DepositAssetsIntoTheManagedSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DepositAssetsIntoTheManagedSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService DepositAssetsIntoTheManagedSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.DepositAssetsIntoTheManagedSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService DepositAssetsIntoTheManagedSubAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.DepositAssetsIntoTheManagedSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService GetManagedSubAccountDepositAddress Success", func(t *testing.T) {

		mockedJSON := `{"coin":"USDT","address":"0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d","tag":"","url":"https://etherscan.io/address/0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/deposit/address", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "coin_example", r.URL.Query().Get("coin"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetManagedSubAccountDepositAddressResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.GetManagedSubAccountDepositAddress(context.Background()).Email("sub-account-email@email.com").Coin("coin_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetManagedSubAccountDepositAddressResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetManagedSubAccountDepositAddressResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService GetManagedSubAccountDepositAddress Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.GetManagedSubAccountDepositAddress(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService GetManagedSubAccountDepositAddress Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.GetManagedSubAccountDepositAddress(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountAssetDetails Success", func(t *testing.T) {

		mockedJSON := `[{"coin":"INJ","name":"Injective Protocol","totalBalance":"0","availableBalance":"0","inOrder":"0","btcValue":"0"},{"coin":"FILDOWN","name":"FILDOWN","totalBalance":"0","availableBalance":"0","inOrder":"0","btcValue":"0"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/asset", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountAssetDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountAssetDetails(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountAssetDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountAssetDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountAssetDetails Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountAssetDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountFuturesAssetDetails Success", func(t *testing.T) {

		mockedJSON := `{"code":"200","message":"OK","snapshotVos":[{"type":"FUTURES","updateTime":1672893855394,"data":{"assets":[{"asset":"USDT","marginBalance":100,"walletBalance":120}],"position":[{"symbol":"BTCUSDT","entryPrice":17000,"markPrice":17000,"positionAmt":1.0E-4}]}}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/fetch-future-asset", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountFuturesAssetDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountFuturesAssetDetails(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountFuturesAssetDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountFuturesAssetDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountFuturesAssetDetails Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountFuturesAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountFuturesAssetDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountFuturesAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountList Success", func(t *testing.T) {

		mockedJSON := `{"total":3,"managerSubUserInfoVoList":[{"rootUserId":1000138475670,"managersubUserId":1000137842513,"bindParentUserId":1000138475669,"email":"test_0_virtual@kq3kno9imanagedsub.com","insertTimeStamp":1678435149000,"bindParentEmail":"wdyw8xsh8pey@test.com","isSubUserEnabled":true,"isUserActive":true,"isMarginEnabled":false,"isFutureEnabled":false,"isSignedLVTRiskAgreement":false},{"rootUserId":1000138475670,"managersubUserId":1000137842514,"bindParentUserId":1000138475669,"email":"test_1_virtual@4qd2u7zxmanagedsub.com","insertTimeStamp":1678435152000,"bindParentEmail":"wdyw8xsh8pey@test.com","isSubUserEnabled":true,"isUserActive":true,"isMarginEnabled":false,"isFutureEnabled":false,"isSignedLVTRiskAgreement":false},{"rootUserId":1000138475670,"managersubUserId":1000137842515,"bindParentUserId":1000138475669,"email":"test_2_virtual@akc05o8hmanagedsub.com","insertTimeStamp":1678435153000,"bindParentEmail":"wdyw8xsh8pey@test.com","isSubUserEnabled":true,"isUserActive":true,"isMarginEnabled":false,"isFutureEnabled":false,"isSignedLVTRiskAgreement":false}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/info", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountMarginAssetDetails Success", func(t *testing.T) {

		mockedJSON := `{"marginLevel":"999","totalAssetOfBtc":"0","totalLiabilityOfBtc":"0","totalNetAssetOfBtc":"0","userAssets":[{"asset":"MATIC","borrowed":"0","free":"0","interest":"0","locked":"0","netAsset":"0"},{"asset":"VET","borrowed":"0","free":"0","interest":"0","locked":"0","netAsset":"0"},{"asset":"BAKE","borrowed":"0","free":"0","interest":"0","locked":"0","netAsset":"0"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/marginAsset", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountMarginAssetDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountMarginAssetDetails(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountMarginAssetDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountMarginAssetDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountMarginAssetDetails Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountMarginAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountMarginAssetDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountMarginAssetDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountSnapshot Success", func(t *testing.T) {

		mockedJSON := `{"code":200,"msg":"","snapshotVos":[{"data":{"balances":[{"asset":"BTC","free":"0.09905021","locked":"0.00000000"},{"asset":"USDT","free":"1.89109409","locked":"0.00000000"}],"totalAssetOfBtc":"0.09942700"},"type":"spot","updateTime":1576281599000},{"data":{"marginLevel":"2748.02909813","totalAssetOfBtc":"0.00274803","totalLiabilityOfBtc":"0.00000100","totalNetAssetOfBtc":"0.00274750","userAssets":[{"asset":"XRP","borrowed":"0.00000000","free":"1.00000000","interest":"0.00000000","locked":"0.00000000","netAsset":"1.00000000"}]},"type":"margin","updateTime":1576281599000},{"data":{"assets":[{"asset":"USDT","marginBalance":"118.99782335","walletBalance":"120.23811389"}],"position":[{"entryPrice":"7130.41000000","markPrice":"7257.66239673","positionAmt":"0.01000000","symbol":"BTCUSDT","unRealizedProfit":"1.24029054"}]},"type":"futures","updateTime":1576281599000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/accountSnapshot", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "type__example", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountSnapshotResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountSnapshot(context.Background()).Email("sub-account-email@email.com").Type("type__example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountSnapshotResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountSnapshotResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountSnapshot Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountSnapshot(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountSnapshot Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountSnapshot(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountInvestor Success", func(t *testing.T) {

		mockedJSON := `{"managerSubTransferHistoryVos":[{"fromEmail":"test_0_virtual@kq3kno9imanagedsub.com","fromAccountType":"SPOT","toEmail":"wdywl0lddakh@test.com","toAccountType":"SPOT","asset":"BNB","amount":"0.01","scheduledData":1679416673000,"createTime":1679416673000,"status":"SUCCESS","tranId":91077779},{"fromEmail":"wdywl0lddakh@test.com","fromAccountType":"SPOT","toEmail":"test_0_virtual@kq3kno9imanagedsub.com","toAccountType":"SPOT","asset":"BNB","amount":"1","scheduledData":1679416616000,"createTime":1679416616000,"status":"SUCCESS","tranId":91077676}],"count":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/queryTransLogForInvestor", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			require.Equal(t, "789", r.URL.Query().Get("page"))
			require.Equal(t, "789", r.URL.Query().Get("limit"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor(context.Background()).Email("sub-account-email@email.com").StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Page(int64(789)).Limit(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountTransferLogMasterAccountInvestorResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountInvestor Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountInvestor Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountInvestor(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountTrading Success", func(t *testing.T) {

		mockedJSON := `{"managerSubTransferHistoryVos":[{"fromEmail":"test_0_virtual@kq3kno9imanagedsub.com","fromAccountType":"SPOT","toEmail":"wdywl0lddakh@test.com","toAccountType":"SPOT","asset":"BNB","amount":"0.01","scheduledData":1679416673000,"createTime":1679416673000,"status":"SUCCESS","tranId":91077779},{"fromEmail":"wdywl0lddakh@test.com","fromAccountType":"SPOT","toEmail":"test_0_virtual@kq3kno9imanagedsub.com","toAccountType":"SPOT","asset":"BNB","amount":"1","scheduledData":1679416616000,"createTime":1679416616000,"status":"SUCCESS","tranId":91077676}],"count":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/queryTransLogForTradeParent", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			require.Equal(t, "789", r.URL.Query().Get("page"))
			require.Equal(t, "789", r.URL.Query().Get("limit"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountTrading(context.Background()).Email("sub-account-email@email.com").StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Page(int64(789)).Limit(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountTransferLogMasterAccountTradingResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountTrading Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountTrading(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogMasterAccountTrading Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogMasterAccountTrading(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogSubAccountTrading Success", func(t *testing.T) {

		mockedJSON := `{"managerSubTransferHistoryVos":[{"fromEmail":"test_0_virtual@kq3kno9imanagedsub.com","fromAccountType":"SPOT","toEmail":"wdywl0lddakh@test.com","toAccountType":"SPOT","asset":"BNB","amount":"0.01","scheduledData":1679416673000,"createTime":1679416673000,"status":"SUCCESS","tranId":91077779},{"fromEmail":"wdywl0lddakh@test.com","fromAccountType":"SPOT","toEmail":"test_0_virtual@kq3kno9imanagedsub.com","toAccountType":"SPOT","asset":"BNB","amount":"1","scheduledData":1679416616000,"createTime":1679416616000,"status":"SUCCESS","tranId":91077676}],"count":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/query-trans-log", r.URL.Path)
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			require.Equal(t, "789", r.URL.Query().Get("page"))
			require.Equal(t, "789", r.URL.Query().Get("limit"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryManagedSubAccountTransferLogSubAccountTradingResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogSubAccountTrading(context.Background()).StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Page(int64(789)).Limit(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryManagedSubAccountTransferLogSubAccountTradingResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryManagedSubAccountTransferLogSubAccountTradingResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogSubAccountTrading Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogSubAccountTrading(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService QueryManagedSubAccountTransferLogSubAccountTrading Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.QueryManagedSubAccountTransferLogSubAccountTrading(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService WithdrawlAssetsFromTheManagedSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"tranId":66157362489}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/managed-subaccount/withdraw", r.URL.Path)
			require.Equal(t, "fromEmail_example", r.URL.Query().Get("fromEmail"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.WithdrawlAssetsFromTheManagedSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.WithdrawlAssetsFromTheManagedSubAccount(context.Background()).FromEmail("fromEmail_example").Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.WithdrawlAssetsFromTheManagedSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.WithdrawlAssetsFromTheManagedSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test ManagedSubAccountAPIService WithdrawlAssetsFromTheManagedSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.WithdrawlAssetsFromTheManagedSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test ManagedSubAccountAPIService WithdrawlAssetsFromTheManagedSubAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.ManagedSubAccountAPI.WithdrawlAssetsFromTheManagedSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
