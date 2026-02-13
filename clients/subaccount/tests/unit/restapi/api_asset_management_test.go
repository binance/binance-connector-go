/*
Binance Sub Account REST API TEST

Testing AssetManagementAPIService

*/

package binancesubaccountrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/subaccount"
	"github.com/binance/binance-connector-go/clients/subaccount/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesubaccountrestapi_AssetManagementAPIService(t *testing.T) {

	t.Run("Test AssetManagementAPIService FuturesTransferForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"txnId":"2966662589"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/transfer", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			require.Equal(t, "789", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FuturesTransferForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.FuturesTransferForSubAccount(context.Background()).Email("sub-account-email@email.com").Asset("asset_example").Amount(float32(1.0)).Type(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FuturesTransferForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FuturesTransferForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService FuturesTransferForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.FuturesTransferForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService FuturesTransferForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.FuturesTransferForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccount Success", func(t *testing.T) {

		mockedJSON := `{"email":"abc@test.com","asset":"USDT","assets":[{"asset":"USDT","initialMargin":"0.00000000","maintenanceMargin":"0.00000000","marginBalance":"0.88308000","maxWithdrawAmount":"0.88308000","openOrderInitialMargin":"0.00000000","positionInitialMargin":"0.00000000","unrealizedProfit":"0.00000000","walletBalance":"0.88308000"}],"canDeposit":true,"canTrade":true,"canWithdraw":true,"feeTier":2,"maxWithdrawAmount":"0.88308000","totalInitialMargin":"0.00000000","totalMaintenanceMargin":"0.00000000","totalMarginBalance":"0.88308000","totalOpenOrderInitialMargin":"0.00000000","totalPositionInitialMargin":"0.00000000","totalUnrealizedProfit":"0.00000000","totalWalletBalance":"0.88308000","updateTime":1576756674610}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/account", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDetailOnSubAccountsFuturesAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccount(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDetailOnSubAccountsFuturesAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccountV2 Success", func(t *testing.T) {

		mockedJSON := `{"futureAccountResp":{"email":"abc@test.com","assets":[{"asset":"USDT","initialMargin":"0.00000000","maintenanceMargin":"0.00000000","marginBalance":"0.88308000","maxWithdrawAmount":"0.88308000","openOrderInitialMargin":"0.00000000","positionInitialMargin":"0.00000000","unrealizedProfit":"0.00000000","walletBalance":"0.88308000"}],"canDeposit":true,"canTrade":true,"canWithdraw":true,"feeTier":2,"maxWithdrawAmount":"0.88308000","totalInitialMargin":"0.00000000","totalMaintenanceMargin":"0.00000000","totalMarginBalance":"0.88308000","totalOpenOrderInitialMargin":"0.00000000","totalPositionInitialMargin":"0.00000000","totalUnrealizedProfit":"0.00000000","totalWalletBalance":"0.88308000","updateTime":1576756674610},"deliveryAccountResp":{"email":"abc@test.com","assets":[{"asset":"BTC","initialMargin":"0.00000000","maintenanceMargin":"0.00000000","marginBalance":"0.88308000","maxWithdrawAmount":"0.88308000","openOrderInitialMargin":"0.00000000","positionInitialMargin":"0.00000000","unrealizedProfit":"0.00000000","walletBalance":"0.88308000"}],"canDeposit":true,"canTrade":true,"canWithdraw":true,"feeTier":2,"updateTime":1598959682001}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/sub-account/futures/account", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "789", r.URL.Query().Get("futuresType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDetailOnSubAccountsFuturesAccountV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccountV2(context.Background()).Email("sub-account-email@email.com").FuturesType(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDetailOnSubAccountsFuturesAccountV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDetailOnSubAccountsFuturesAccountV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccountV2 Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsFuturesAccountV2 Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsFuturesAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsMarginAccount Success", func(t *testing.T) {

		mockedJSON := `{"email":"123@test.com","marginLevel":"11.64405625","totalAssetOfBtc":"6.82728457","totalLiabilityOfBtc":"0.58633215","totalNetAssetOfBtc":"6.24095242","marginTradeCoeffVo":{"forceLiquidationBar":"1.10000000","marginCallBar":"1.50000000","normalBar":"2.00000000"},"marginUserAssetVoList":[{"asset":"BTC","borrowed":"0.00000000","free":"0.00499500","interest":"0.00000000","locked":"0.00000000","netAsset":"0.00499500"},{"asset":"BNB","borrowed":"201.66666672","free":"2346.50000000","interest":"0.00000000","locked":"0.00000000","netAsset":"2144.83333328"},{"asset":"ETH","borrowed":"0.00000000","free":"0.00000000","interest":"0.00000000","locked":"0.00000000","netAsset":"0.00000000"},{"asset":"USDT","borrowed":"0.00000000","free":"0.00000000","interest":"0.00000000","locked":"0.00000000","netAsset":"0.00000000"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/margin/account", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDetailOnSubAccountsMarginAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsMarginAccount(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDetailOnSubAccountsMarginAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDetailOnSubAccountsMarginAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsMarginAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsMarginAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetDetailOnSubAccountsMarginAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetDetailOnSubAccountsMarginAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetMovePositionHistoryForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"total":3,"futureMovePositionOrderVoList":[{"fromUserEmail":"testFrom@google.com","toUserEmail":"testTo@google.com","productType":"UM","symbol":"BTCUSDT","price":"105025.50981609","quantity":"0.00100000","positionSide":"BOTH","side":"SELL","timeStamp":1737544712000},{"fromUserEmail":"testFrom1@google.com","toUserEmail":"testTo1@google.com","productType":"UM","symbol":"BTCUSDT","price":"97100.00000000","quantity":"0.00100000","positionSide":"BOTH","side":"SELL","timeStamp":1740041627000},{"fromUserEmail":"testFrom2@google.com","toUserEmail":"testTo2@google.com","productType":"UM","symbol":"BTCUSDT","price":"97108.62068889","quantity":"0.00100000","positionSide":"BOTH","side":"SELL","timeStamp":1740041959000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/move-position", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, "789", r.URL.Query().Get("page"))
			require.Equal(t, "789", r.URL.Query().Get("row"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetMovePositionHistoryForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetMovePositionHistoryForSubAccount(context.Background()).Symbol("symbol_example").Page(int64(789)).Row(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetMovePositionHistoryForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetMovePositionHistoryForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetMovePositionHistoryForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetMovePositionHistoryForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetMovePositionHistoryForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetMovePositionHistoryForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositAddress Success", func(t *testing.T) {

		mockedJSON := `{"address":"TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV","coin":"USDT","tag":"","url":"https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/capital/deposit/subAddress", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "coin_example", r.URL.Query().Get("coin"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSubAccountDepositAddressResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositAddress(context.Background()).Email("sub-account-email@email.com").Coin("coin_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSubAccountDepositAddressResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSubAccountDepositAddressResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositAddress Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositAddress(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositAddress Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositAddress(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositHistory Success", func(t *testing.T) {

		mockedJSON := `[{"id":"769800519366885376","amount":"0.001","coin":"BNB","network":"BNB","status":0,"address":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","addressTag":"101764890","txId":"98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC","insertTime":1661493146000,"transferType":0,"confirmTimes":"1/1","unlockConfirm":0,"walletType":0},{"id":"769754833590042625","amount":"0.50000000","coin":"IOTA","network":"IOTA","status":1,"address":"SIZ9VLMHWATXKV99LH99CIGFJFUMLEHGWVZVNNZXRJJVWBPHYWPPBOSDORZ9EQSHCZAMPVAPGFYQAUUV9DROOXJLNW","addressTag":"","txId":"ESBFVQUTPIWQNJSPXFNHNYHSQNTGKRVKPRABQWTAXCDWOAKDKYWPTVG9BGXNVNKTLEJGESAVXIKIZ9999","insertTime":1599620082000,"transferType":0,"confirmTimes":"1/1","unlockConfirm":0,"walletType":0}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/capital/deposit/subHisrec", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSubAccountDepositHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositHistory(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSubAccountDepositHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSubAccountDepositHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSubAccountDepositHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSubAccountDepositHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccount Success", func(t *testing.T) {

		mockedJSON := `{"totalInitialMargin":"9.83137400","totalMaintenanceMargin":"0.41568700","totalMarginBalance":"23.03235621","totalOpenOrderInitialMargin":"9.00000000","totalPositionInitialMargin":"0.83137400","totalUnrealizedProfit":"0.03219710","totalWalletBalance":"22.15879444","asset":"USD","subAccountList":[{"email":"123@test.com","totalInitialMargin":"9.00000000","totalMaintenanceMargin":"0.00000000","totalMarginBalance":"22.12659734","totalOpenOrderInitialMargin":"9.00000000","totalPositionInitialMargin":"0.00000000","totalUnrealizedProfit":"0.00000000","totalWalletBalance":"22.12659734","asset":"USD"},{"email":"345@test.com","totalInitialMargin":"0.83137400","totalMaintenanceMargin":"0.41568700","totalMarginBalance":"0.90575887","totalOpenOrderInitialMargin":"0.00000000","totalPositionInitialMargin":"0.83137400","totalUnrealizedProfit":"0.03219710","totalWalletBalance":"0.87356177","asset":"USD"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/accountSummary", r.URL.Path)
			require.Equal(t, "789", r.URL.Query().Get("page"))
			require.Equal(t, "789", r.URL.Query().Get("limit"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSummaryOfSubAccountsFuturesAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccount(context.Background()).Page(int64(789)).Limit(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSummaryOfSubAccountsFuturesAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccountV2 Success", func(t *testing.T) {

		mockedJSON := `{"futureAccountSummaryResp":{"totalInitialMargin":"9.83137400","totalMaintenanceMargin":"0.41568700","totalMarginBalance":"23.03235621","totalOpenOrderInitialMargin":"9.00000000","totalPositionInitialMargin":"0.83137400","totalUnrealizedProfit":"0.03219710","totalWalletBalance":"22.15879444","asset":"USD","subAccountList":[{"email":"123@test.com","totalInitialMargin":"9.00000000","totalMaintenanceMargin":"0.00000000","totalMarginBalance":"22.12659734","totalOpenOrderInitialMargin":"9.00000000","totalPositionInitialMargin":"0.00000000","totalUnrealizedProfit":"0.00000000","totalWalletBalance":"22.12659734","asset":"USD"},{"email":"345@test.com","totalInitialMargin":"0.83137400","totalMaintenanceMargin":"0.41568700","totalMarginBalance":"0.90575887","totalOpenOrderInitialMargin":"0.00000000","totalPositionInitialMargin":"0.83137400","totalUnrealizedProfit":"0.03219710","totalWalletBalance":"0.87356177","asset":"USD"}]},"deliveryAccountSummaryResp":{"totalMarginBalanceOfBTC":"25.03221121","totalUnrealizedProfitOfBTC":"0.12233410","totalWalletBalanceOfBTC":"22.15879444","asset":"BTC","subAccountList":[{"email":"123@test.com","totalMarginBalance":"22.12659734","totalUnrealizedProfit":"0.00000000","totalWalletBalance":"22.12659734","asset":"BTC"},{"email":"345@test.com","totalMarginBalance":"0.90575887","totalUnrealizedProfit":"0.03219710","totalWalletBalance":"0.87356177","asset":"BTC"}]}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/sub-account/futures/accountSummary", r.URL.Path)
			require.Equal(t, "789", r.URL.Query().Get("futuresType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSummaryOfSubAccountsFuturesAccountV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccountV2(context.Background()).FuturesType(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSummaryOfSubAccountsFuturesAccountV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSummaryOfSubAccountsFuturesAccountV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccountV2 Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsFuturesAccountV2 Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsFuturesAccountV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsMarginAccount Success", func(t *testing.T) {

		mockedJSON := `{"totalAssetOfBtc":"4.33333333","totalLiabilityOfBtc":"2.11111112","totalNetAssetOfBtc":"2.22222221","subAccountList":[{"email":"123@test.com","totalAssetOfBtc":"2.11111111","totalLiabilityOfBtc":"1.11111111","totalNetAssetOfBtc":"1.00000000"},{"email":"345@test.com","totalAssetOfBtc":"2.22222222","totalLiabilityOfBtc":"1.00000001","totalNetAssetOfBtc":"1.22222221"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/margin/accountSummary", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSummaryOfSubAccountsMarginAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsMarginAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSummaryOfSubAccountsMarginAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSummaryOfSubAccountsMarginAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService GetSummaryOfSubAccountsMarginAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.GetSummaryOfSubAccountsMarginAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService MarginTransferForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"txnId":"2966662589"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/margin/transfer", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			require.Equal(t, "789", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarginTransferForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.MarginTransferForSubAccount(context.Background()).Email("sub-account-email@email.com").Asset("asset_example").Amount(float32(1.0)).Type(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarginTransferForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarginTransferForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService MarginTransferForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.MarginTransferForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService MarginTransferForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.MarginTransferForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService MovePositionForSubAccount Success", func(t *testing.T) {

		mockedJSON := `{"movePositionOrders":[{"fromUserEmail":"testFrom@google.com","toUserEmail":"testTo@google.com","productType":"UM","symbol":"BTCUSDT","priceType":"MARK_PRICE","price":"97139.00000000","quantity":"0.001","positionSide":"BOTH","side":"BUY","success":true},{"fromUserEmail":"testFrom1@google.com","toUserEmail":"1testTo@google.com","productType":"UM","symbol":"BTCUSDT","priceType":"MARK_PRICE","price":"97139.00000000","quantity":"0.0011","positionSide":"BOTH","side":"BUY","success":true}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/move-position", r.URL.Path)
			require.Equal(t, "fromUserEmail_example", r.URL.Query().Get("fromUserEmail"))
			require.Equal(t, "toUserEmail_example", r.URL.Query().Get("toUserEmail"))
			require.Equal(t, "productType_example", r.URL.Query().Get("productType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MovePositionForSubAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.MovePositionForSubAccount(context.Background()).FromUserEmail("fromUserEmail_example").ToUserEmail("toUserEmail_example").ProductType("productType_example").OrderArgs([]models.MovePositionForSubAccountOrderArgsParameterInner{*models.NewMovePositionForSubAccountOrderArgsParameterInner()}).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MovePositionForSubAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MovePositionForSubAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService MovePositionForSubAccount Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.MovePositionForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService MovePositionForSubAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.MovePositionForSubAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssets Success", func(t *testing.T) {

		mockedJSON := `{"balances":[{"freeze":0,"withdrawing":0,"asset":"ADA","free":10000,"locked":0},{"freeze":0,"withdrawing":0,"asset":"BNB","free":10003,"locked":0},{"freeze":0,"withdrawing":0,"asset":"BTC","free":11467.6399,"locked":0},{"freeze":0,"withdrawing":0,"asset":"ETH","free":10004.995,"locked":0},{"freeze":0,"withdrawing":0,"asset":"USDT","free":11652.14213,"locked":0}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v3/sub-account/assets", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountAssetsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssets(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountAssetsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountAssetsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssets Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssets Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssetsAssetManagement Success", func(t *testing.T) {

		mockedJSON := `{"balances":[{"freeze":"0","withdrawing":"0","asset":"ADA","free":"10000","locked":"0"},{"freeze":"0","withdrawing":"0","asset":"BNB","free":"10003","locked":"0"},{"freeze":"0","withdrawing":"0","asset":"BTC","free":"11467.6399","locked":"0"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v4/sub-account/assets", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountAssetsAssetManagementResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssetsAssetManagement(context.Background()).Email("sub-account-email@email.com").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountAssetsAssetManagementResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountAssetsAssetManagementResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssetsAssetManagement Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssetsAssetManagement(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountAssetsAssetManagement Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountAssetsAssetManagement(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountFuturesAssetTransferHistory Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"futuresType":2,"transfers":[{"from":"aaa@test.com","to":"bbb@test.com","asset":"BTC","qty":"1","tranId":11897001102,"time":1544433328000},{"from":"bbb@test.com","to":"ccc@test.com","asset":"ETH","qty":"2","tranId":11631474902,"time":1544433328000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/internalTransfer", r.URL.Path)
			require.Equal(t, "sub-account-email@email.com", r.URL.Query().Get("email"))
			require.Equal(t, "789", r.URL.Query().Get("futuresType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountFuturesAssetTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountFuturesAssetTransferHistory(context.Background()).Email("sub-account-email@email.com").FuturesType(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountFuturesAssetTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountFuturesAssetTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountFuturesAssetTransferHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountFuturesAssetTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountFuturesAssetTransferHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountFuturesAssetTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountSpotAssetTransferHistory Success", func(t *testing.T) {

		mockedJSON := `[{"from":"aaa@test.com","to":"bbb@test.com","asset":"BTC","qty":"10","status":"SUCCESS","tranId":6489943656,"time":1544433328000},{"from":"bbb@test.com","to":"ccc@test.com","asset":"ETH","qty":"2","status":"SUCCESS","tranId":6489938713,"time":1544433328000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/sub/transfer/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountSpotAssetTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetTransferHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountSpotAssetTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountSpotAssetTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountSpotAssetTransferHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountSpotAssetsSummary Success", func(t *testing.T) {

		mockedJSON := `{"totalCount":2,"masterAccountTotalAsset":"0.23231201","spotSubUserAssetBtcVoList":[{"email":"sub123@test.com","totalAsset":"9999.00000000"},{"email":"test456@test.com","totalAsset":"0.00000000"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/spotSummary", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySubAccountSpotAssetsSummaryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetsSummary(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySubAccountSpotAssetsSummaryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySubAccountSpotAssetsSummaryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QuerySubAccountSpotAssetsSummary Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QuerySubAccountSpotAssetsSummary(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService QueryUniversalTransferHistory Success", func(t *testing.T) {

		mockedJSON := `{"result":[{"tranId":92275823339,"fromEmail":"abctest@gmail.com","toEmail":"deftest@gmail.com","asset":"BNB","amount":"0.01","createTimeStamp":1640317374000,"fromAccountType":"USDT_FUTURE","toAccountType":"SPOT","status":"SUCCESS","clientTranId":"test"}],"totalCount":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/universalTransfer", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryUniversalTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.QueryUniversalTransferHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryUniversalTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryUniversalTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService QueryUniversalTransferHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.QueryUniversalTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService SubAccountFuturesAssetTransfer Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"txnId":"2934662589"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/futures/internalTransfer", r.URL.Path)
			require.Equal(t, "fromEmail_example", r.URL.Query().Get("fromEmail"))
			require.Equal(t, "toEmail_example", r.URL.Query().Get("toEmail"))
			require.Equal(t, "789", r.URL.Query().Get("futuresType"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubAccountFuturesAssetTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountFuturesAssetTransfer(context.Background()).FromEmail("fromEmail_example").ToEmail("toEmail_example").FuturesType(int64(789)).Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubAccountFuturesAssetTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubAccountFuturesAssetTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService SubAccountFuturesAssetTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountFuturesAssetTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService SubAccountFuturesAssetTransfer Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountFuturesAssetTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService SubAccountTransferHistory Success", func(t *testing.T) {

		mockedJSON := `[{"counterParty":"master","email":"master@test.com","type":1,"asset":"BTC","qty":"1","fromAccountType":"SPOT","toAccountType":"SPOT","status":"SUCCESS","tranId":11798835829,"time":1544433325000},{"counterParty":"subAccount","email":"sub2@test.com","type":1,"asset":"ETH","qty":"2","fromAccountType":"SPOT","toAccountType":"COIN_FUTURE","status":"SUCCESS","tranId":11798829519,"time":1544433326000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/transfer/subUserHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubAccountTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountTransferHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubAccountTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubAccountTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService SubAccountTransferHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.SubAccountTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService TransferToMaster Success", func(t *testing.T) {

		mockedJSON := `{"txnId":"2966662589"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/transfer/subToMaster", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TransferToMasterResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToMaster(context.Background()).Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TransferToMasterResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TransferToMasterResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService TransferToMaster Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToMaster(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService TransferToMaster Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToMaster(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService TransferToSubAccountOfSameMaster Success", func(t *testing.T) {

		mockedJSON := `{"txnId":"2966662589"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/transfer/subToSub", r.URL.Path)
			require.Equal(t, "toEmail_example", r.URL.Query().Get("toEmail"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TransferToSubAccountOfSameMasterResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToSubAccountOfSameMaster(context.Background()).ToEmail("toEmail_example").Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TransferToSubAccountOfSameMasterResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TransferToSubAccountOfSameMasterResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService TransferToSubAccountOfSameMaster Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToSubAccountOfSameMaster(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService TransferToSubAccountOfSameMaster Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.TransferToSubAccountOfSameMaster(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService UniversalTransfer Success", func(t *testing.T) {

		mockedJSON := `{"tranId":11945860693,"clientTranId":"test"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sub-account/universalTransfer", r.URL.Path)
			require.Equal(t, "fromAccountType_example", r.URL.Query().Get("fromAccountType"))
			require.Equal(t, "toAccountType_example", r.URL.Query().Get("toAccountType"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.UniversalTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.UniversalTransfer(context.Background()).FromAccountType("fromAccountType_example").ToAccountType("toAccountType_example").Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.UniversalTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.UniversalTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetManagementAPIService UniversalTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSubAccountClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetManagementAPI.UniversalTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetManagementAPIService UniversalTransfer Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AssetManagementAPI.UniversalTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
