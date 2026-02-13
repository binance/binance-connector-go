/*
Binance Derivatives Trading USDS Futures REST API TEST

Testing AccountAPIService

*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingusdsfuturesrestapi_AccountAPIService(t *testing.T) {

	t.Run("Test AccountAPIService AccountInformationV2 Success", func(t *testing.T) {

		mockedJSON := `{"feeTier":0,"feeBurn":true,"canDeposit":true,"canWithdraw":true,"updateTime":0,"multiAssetsMargin":true,"tradeGroupId":-1,"totalInitialMargin":"0.00000000","totalMaintMargin":"0.00000000","totalWalletBalance":"126.72469206","totalUnrealizedProfit":"0.00000000","totalMarginBalance":"126.72469206","totalPositionInitialMargin":"0.00000000","totalOpenOrderInitialMargin":"0.00000000","totalCrossWalletBalance":"126.72469206","totalCrossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"126.72469206","assets":[{"asset":"USDT","walletBalance":"23.72469206","unrealizedProfit":"0.00000000","marginBalance":"23.72469206","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"23.72469206","maxWithdrawAmount":"23.72469206","marginAvailable":true,"updateTime":1625474304765},{"asset":"BUSD","walletBalance":"103.12345678","unrealizedProfit":"0.00000000","marginBalance":"103.12345678","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"103.12345678","crossUnPnl":"0.00000000","availableBalance":"103.12345678","maxWithdrawAmount":"103.12345678","marginAvailable":true,"updateTime":1625474304765},{"asset":"USDT","walletBalance":"23.72469206","unrealizedProfit":"0.00000000","marginBalance":"23.72469206","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"23.72469206","marginAvailable":true,"updateTime":1625474304765},{"asset":"BUSD","walletBalance":"103.12345678","unrealizedProfit":"0.00000000","marginBalance":"103.12345678","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"103.12345678","crossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"103.12345678","marginAvailable":true,"updateTime":1625474304765}],"positions":[{"symbol":"BTCUSDT","initialMargin":"0","maintMargin":"0","unrealizedProfit":"0.00000000","positionInitialMargin":"0","openOrderInitialMargin":"0","leverage":"100","isolated":true,"entryPrice":"0.00000","maxNotional":"250000","bidNotional":"0","askNotional":"0","positionSide":"BOTH","positionAmt":"0","updateTime":0},{"symbol":"BTCUSDT","initialMargin":"0","maintMargin":"0","unrealizedProfit":"0.00000000","positionInitialMargin":"0","openOrderInitialMargin":"0","leverage":"100","isolated":true,"entryPrice":"0.00000","maxNotional":"250000","bidNotional":"0","askNotional":"0","positionSide":"BOTH","positionAmt":"0","updateTime":0}],"canTrade":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v2/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountInformationV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountInformationV2(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountInformationV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountInformationV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AccountInformationV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountInformationV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AccountInformationV3 Success", func(t *testing.T) {

		mockedJSON := `{"totalInitialMargin":"0.00000000","totalMaintMargin":"0.00000000","totalWalletBalance":"126.72469206","totalUnrealizedProfit":"0.00000000","totalMarginBalance":"126.72469206","totalPositionInitialMargin":"0.00000000","totalOpenOrderInitialMargin":"0.00000000","totalCrossWalletBalance":"126.72469206","totalCrossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"126.72469206","assets":[{"asset":"USDT","walletBalance":"23.72469206","unrealizedProfit":"0.00000000","marginBalance":"23.72469206","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"23.72469206","maxWithdrawAmount":"23.72469206","updateTime":1625474304765},{"asset":"USDC","walletBalance":"103.12345678","unrealizedProfit":"0.00000000","marginBalance":"103.12345678","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"103.12345678","crossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"103.12345678","updateTime":1625474304765},{"asset":"USDT","walletBalance":"23.72469206","unrealizedProfit":"0.00000000","marginBalance":"23.72469206","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"23.72469206","updateTime":1625474304765},{"asset":"BUSD","walletBalance":"103.12345678","unrealizedProfit":"0.00000000","marginBalance":"103.12345678","maintMargin":"0.00000000","initialMargin":"0.00000000","positionInitialMargin":"0.00000000","openOrderInitialMargin":"0.00000000","crossWalletBalance":"103.12345678","crossUnPnl":"0.00000000","availableBalance":"126.72469206","maxWithdrawAmount":"103.12345678","updateTime":1625474304765}],"positions":[{"symbol":"BTCUSDT","positionSide":"BOTH","positionAmt":"1.000","unrealizedProfit":"0.00000000","isolatedMargin":"0.00000000","notional":"0","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","updateTime":0},{"symbol":"BTCUSDT","positionSide":"BOTH","positionAmt":"1.000","unrealizedProfit":"0.00000000","isolatedMargin":"0.00000000","notional":"0","isolatedWallet":"0","initialMargin":"0","maintMargin":"0","updateTime":0}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v3/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountInformationV3Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountInformationV3(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountInformationV3Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountInformationV3Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AccountInformationV3 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountInformationV3(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService FuturesAccountBalanceV2 Success", func(t *testing.T) {

		mockedJSON := `[{"accountAlias":"SgsR","asset":"USDT","balance":"122607.35137903","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"23.72469206","maxWithdrawAmount":"23.72469206","marginAvailable":true,"updateTime":1617939110373}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v2/balance", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FuturesAccountBalanceV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV2(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FuturesAccountBalanceV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FuturesAccountBalanceV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService FuturesAccountBalanceV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService FuturesAccountBalanceV3 Success", func(t *testing.T) {

		mockedJSON := `[{"accountAlias":"SgsR","asset":"USDT","balance":"122607.35137903","crossWalletBalance":"23.72469206","crossUnPnl":"0.00000000","availableBalance":"23.72469206","maxWithdrawAmount":"23.72469206","marginAvailable":true,"updateTime":1617939110373}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v3/balance", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FuturesAccountBalanceV3Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV3(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FuturesAccountBalanceV3Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FuturesAccountBalanceV3Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService FuturesAccountBalanceV3 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountBalanceV3(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService FuturesAccountConfiguration Success", func(t *testing.T) {

		mockedJSON := `{"feeTier":0,"canTrade":true,"canDeposit":true,"canWithdraw":true,"dualSidePosition":true,"updateTime":0,"multiAssetsMargin":false,"tradeGroupId":-1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/accountConfig", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FuturesAccountConfigurationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountConfiguration(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FuturesAccountConfigurationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FuturesAccountConfigurationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService FuturesAccountConfiguration Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesAccountConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService FuturesTradingQuantitativeRulesIndicators Success", func(t *testing.T) {

		mockedJSON := `{"indicators":{"BTCUSDT":[{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"UFR","value":0.05,"triggerValue":0.995},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"IFER","value":0.99,"triggerValue":0.99},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"GCR","value":0.99,"triggerValue":0.99},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"DR","value":0.99,"triggerValue":0.99}],"ETHUSDT":[{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"UFR","value":0.05,"triggerValue":0.995},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"IFER","value":0.99,"triggerValue":0.99},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"GCR","value":0.99,"triggerValue":0.99},{"isLocked":true,"plannedRecoverTime":1545741270000,"indicator":"DR","value":0.99,"triggerValue":0.99}],"ACCOUNT":[{"indicator":"TMV","value":10,"triggerValue":1,"plannedRecoverTime":1644919865000,"isLocked":true}]},"updateTime":1644913304748}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/apiTradingStatus", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FuturesTradingQuantitativeRulesIndicatorsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesTradingQuantitativeRulesIndicators(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FuturesTradingQuantitativeRulesIndicatorsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FuturesTradingQuantitativeRulesIndicatorsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService FuturesTradingQuantitativeRulesIndicators Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.FuturesTradingQuantitativeRulesIndicators(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetBnbBurnStatus Success", func(t *testing.T) {

		mockedJSON := `{"feeBurn":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/feeBurn", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetBnbBurnStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetBnbBurnStatus(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetBnbBurnStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetBnbBurnStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetBnbBurnStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetBnbBurnStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetCurrentMultiAssetsMode Success", func(t *testing.T) {

		mockedJSON := `{"multiAssetsMargin":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/multiAssetsMargin", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCurrentMultiAssetsModeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetCurrentMultiAssetsMode(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCurrentMultiAssetsModeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCurrentMultiAssetsModeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetCurrentMultiAssetsMode Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetCurrentMultiAssetsMode(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetCurrentPositionMode Success", func(t *testing.T) {

		mockedJSON := `{"dualSidePosition":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/positionSide/dual", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCurrentPositionModeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetCurrentPositionMode(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCurrentPositionModeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCurrentPositionModeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetCurrentPositionMode Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetCurrentPositionMode(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesOrderHistory Success", func(t *testing.T) {

		mockedJSON := `{"avgCostTimestampOfLast30d":7241837,"downloadId":"546975389218332672"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/order/asyn", r.URL.Path)
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDownloadIdForFuturesOrderHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesOrderHistory(context.Background()).StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDownloadIdForFuturesOrderHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDownloadIdForFuturesOrderHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesOrderHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesOrderHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesOrderHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesOrderHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTradeHistory Success", func(t *testing.T) {

		mockedJSON := `{"avgCostTimestampOfLast30d":7241837,"downloadId":"546975389218332672"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/trade/asyn", r.URL.Path)
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDownloadIdForFuturesTradeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTradeHistory(context.Background()).StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDownloadIdForFuturesTradeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDownloadIdForFuturesTradeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTradeHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTradeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTradeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTradeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTransactionHistory Success", func(t *testing.T) {

		mockedJSON := `{"avgCostTimestampOfLast30d":7241837,"downloadId":"546975389218332672"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/income/asyn", r.URL.Path)
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDownloadIdForFuturesTransactionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTransactionHistory(context.Background()).StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDownloadIdForFuturesTransactionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDownloadIdForFuturesTransactionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTransactionHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTransactionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetDownloadIdForFuturesTransactionHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetDownloadIdForFuturesTransactionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesOrderHistoryDownloadLinkById Success", func(t *testing.T) {

		mockedJSON := `{"downloadId":"545923594199212032","status":"processing","url":"","notified":false,"expirationTimestamp":-1,"isExpired":null}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/order/asyn/id", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("downloadId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesOrderHistoryDownloadLinkByIdResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesOrderHistoryDownloadLinkById(context.Background()).DownloadId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesOrderHistoryDownloadLinkByIdResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesOrderHistoryDownloadLinkByIdResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetFuturesOrderHistoryDownloadLinkById Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesOrderHistoryDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesOrderHistoryDownloadLinkById Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesOrderHistoryDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesTradeDownloadLinkById Success", func(t *testing.T) {

		mockedJSON := `{"downloadId":"545923594199212032","status":"processing","url":"","notified":false,"expirationTimestamp":-1,"isExpired":null}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/trade/asyn/id", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("downloadId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesTradeDownloadLinkByIdResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTradeDownloadLinkById(context.Background()).DownloadId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesTradeDownloadLinkByIdResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesTradeDownloadLinkByIdResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetFuturesTradeDownloadLinkById Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTradeDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesTradeDownloadLinkById Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTradeDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesTransactionHistoryDownloadLinkById Success", func(t *testing.T) {

		mockedJSON := `{"downloadId":"545923594199212032","status":"processing","url":"","notified":false,"expirationTimestamp":-1,"isExpired":null}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/income/asyn/id", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("downloadId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesTransactionHistoryDownloadLinkByIdResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTransactionHistoryDownloadLinkById(context.Background()).DownloadId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesTransactionHistoryDownloadLinkByIdResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesTransactionHistoryDownloadLinkByIdResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetFuturesTransactionHistoryDownloadLinkById Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTransactionHistoryDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetFuturesTransactionHistoryDownloadLinkById Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetFuturesTransactionHistoryDownloadLinkById(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetIncomeHistory Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"","incomeType":"TRANSFER","income":"-0.37500000","asset":"USDT","info":"TRANSFER","time":1570608000000,"tranId":9689322392,"tradeId":""},{"symbol":"BTCUSDT","incomeType":"COMMISSION","income":"-0.01000000","asset":"USDT","info":"COMMISSION","time":1570636800000,"tranId":9689322392,"tradeId":"2059192"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/income", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetIncomeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetIncomeHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetIncomeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetIncomeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetIncomeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetIncomeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService NotionalAndLeverageBrackets Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"ETHUSDT","notionalCoef":1.5,"brackets":[{"bracket":1,"initialLeverage":75,"notionalCap":10000,"notionalFloor":0,"maintMarginRatio":0.0065,"cum":0}]}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/leverageBracket", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.NotionalAndLeverageBracketsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.NotionalAndLeverageBrackets(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.NotionalAndLeverageBracketsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.NotionalAndLeverageBracketsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService NotionalAndLeverageBrackets Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.NotionalAndLeverageBrackets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService QueryUserRateLimit Success", func(t *testing.T) {

		mockedJSON := `[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":10000},{"rateLimitType":"ORDERS","interval":"MINUTE","intervalNum":1,"limit":20000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/rateLimit/order", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryUserRateLimitResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.QueryUserRateLimit(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryUserRateLimitResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryUserRateLimitResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService QueryUserRateLimit Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.QueryUserRateLimit(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService SymbolConfiguration Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSDT","marginType":"CROSSED","isAutoAddMargin":false,"leverage":21,"maxNotionalValue":"1000000"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/symbolConfig", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolConfigurationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.SymbolConfiguration(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SymbolConfigurationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SymbolConfigurationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService SymbolConfiguration Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.SymbolConfiguration(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService ToggleBnbBurnOnFuturesTrade Success", func(t *testing.T) {

		mockedJSON := `{"code":200,"msg":"success"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/feeBurn", r.URL.Path)
			require.Equal(t, "feeBurn_example", r.URL.Query().Get("feeBurn"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ToggleBnbBurnOnFuturesTradeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.ToggleBnbBurnOnFuturesTrade(context.Background()).FeeBurn("feeBurn_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ToggleBnbBurnOnFuturesTradeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ToggleBnbBurnOnFuturesTradeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService ToggleBnbBurnOnFuturesTrade Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.ToggleBnbBurnOnFuturesTrade(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService ToggleBnbBurnOnFuturesTrade Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.ToggleBnbBurnOnFuturesTrade(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService UserCommissionRate Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSDT","makerCommissionRate":"0.0002","takerCommissionRate":"0.0004","rpiCommissionRate":"0.00005"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/commissionRate", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.UserCommissionRateResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.UserCommissionRate(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.UserCommissionRateResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.UserCommissionRateResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService UserCommissionRate Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.UserCommissionRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService UserCommissionRate Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.UserCommissionRate(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
