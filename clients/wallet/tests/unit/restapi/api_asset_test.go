/*
Binance Wallet REST API TEST

Testing AssetAPIService

*/

package binancewalletrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancewalletrestapi_AssetAPIService(t *testing.T) {

	t.Run("Test AssetAPIService AssetDetail Success", func(t *testing.T) {

		mockedJSON := `{"CTR":{"minWithdrawAmount":"70.00000000","depositStatus":false,"withdrawFee":35,"withdrawStatus":true,"depositTip":"Delisted, Deposit Suspended"},"SKY":{"minWithdrawAmount":"0.02000000","depositStatus":true,"withdrawFee":0.01,"withdrawStatus":true}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/assetDetail", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AssetDetailResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.AssetDetail(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AssetDetailResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AssetDetailResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService AssetDetail Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.AssetDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService AssetDividendRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"id":1637366104,"amount":"10.00000000","asset":"BHFT","divTime":1563189166000,"enInfo":"BHFT distribution","tranId":2968885920},{"id":1631750237,"amount":"10.00000000","asset":"BHFT","divTime":1563189165000,"enInfo":"BHFT distribution","tranId":2968885920}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/assetDividend", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AssetDividendRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.AssetDividendRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AssetDividendRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AssetDividendRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService AssetDividendRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.AssetDividendRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustConvert Success", func(t *testing.T) {

		mockedJSON := `{"totalTransfered":"3.5971223","totalServiceCharge":"0.0794964","transferResult":[{"tranId":2987331510,"fromAsset":"USDT","amount":"1","transferedAmount":"3.5971223","serviceChargeAmount":"0.0794964","operateTime":1765212029749}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/dust-convert/convert", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DustConvertResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvert(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DustConvertResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DustConvertResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService DustConvert Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvert(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustConvert Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvert(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustConvertibleAssets Success", func(t *testing.T) {

		mockedJSON := `{"dribbletPercentage":"0.02","totalTransferQuotaAssetAmount":"0.7899968","totalTransferTargetAssetAmount":"0.7899968","dribbletBase":"10","details":[{"asset":"AR","assetFullName":"AR","amountFree":"0.00856","exchange":"0.00073616","toQuotaAssetAmount":"0.036808","toTargetAssetAmount":"0.036808","toTargetAssetOffExchange":"0.03607184"},{"asset":"BNB","assetFullName":"BNB","amountFree":"0.00082768","exchange":"0.01506378","toQuotaAssetAmount":"0.7531888","toTargetAssetAmount":"0.7531888","toTargetAssetOffExchange":"0.73812502"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/dust-convert/query-convertible-assets", r.URL.Path)
			require.Equal(t, "targetAsset_example", r.URL.Query().Get("targetAsset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DustConvertibleAssetsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvertibleAssets(context.Background()).TargetAsset("targetAsset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DustConvertibleAssetsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DustConvertibleAssetsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService DustConvertibleAssets Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvertibleAssets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustConvertibleAssets Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustConvertibleAssets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustTransfer Success", func(t *testing.T) {

		mockedJSON := `{"totalServiceCharge":"0.02102542","totalTransfered":"1.05127099","transferResult":[{"amount":"0.03000000","fromAsset":"ETH","operateTime":1563368549307,"serviceChargeAmount":"0.00500000","tranId":2970932918,"transferedAmount":"0.25000000"},{"amount":"0.09000000","fromAsset":"LTC","operateTime":1563368549404,"serviceChargeAmount":"0.01548000","tranId":2970932918,"transferedAmount":"0.77400000"},{"amount":"248.61878453","fromAsset":"TRX","operateTime":1563368549489,"serviceChargeAmount":"0.00054542","tranId":2970932918,"transferedAmount":"0.02727099"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/dust", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DustTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustTransfer(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DustTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DustTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService DustTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService DustTransfer Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.DustTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService Dustlog Success", func(t *testing.T) {

		mockedJSON := `{"total":8,"userAssetDribblets":[{"operateTime":1615985535000,"totalTransferedAmount":"0.00132256","totalServiceChargeAmount":"0.00002699","transId":45178372831,"userAssetDribbletDetails":[{"transId":4359321,"serviceChargeAmount":"0.000009","amount":"0.0009","operateTime":1615985535000,"transferedAmount":"0.000441","fromAsset":"USDT"},{"transId":4359321,"serviceChargeAmount":"0.00001799","amount":"0.0009","operateTime":1615985535000,"transferedAmount":"0.00088156","fromAsset":"ETH"}]},{"operateTime":1616203180000,"totalTransferedAmount":"0.00058795","totalServiceChargeAmount":"0.000012","transId":4357015,"userAssetDribbletDetails":[{"transId":4357015,"serviceChargeAmount":"0.00001","amount":"0.001","operateTime":1616203180000,"transferedAmount":"0.00049","fromAsset":"USDT"},{"transId":4357015,"serviceChargeAmount":"0.000002","amount":"0.0001","operateTime":1616203180000,"transferedAmount":"0.00009795","fromAsset":"ETH"}]}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/dribblet", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DustlogResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.Dustlog(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DustlogResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DustlogResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService Dustlog Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.Dustlog(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService FundingWallet Success", func(t *testing.T) {

		mockedJSON := `[{"asset":"USDT","free":"1","locked":"0","freeze":"0","withdrawing":"0","btcValuation":"0.00000091"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/get-funding-asset", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FundingWalletResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.FundingWallet(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FundingWalletResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FundingWalletResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService FundingWallet Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.FundingWallet(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService GetAssetsThatCanBeConvertedIntoBnb Success", func(t *testing.T) {

		mockedJSON := `{"details":[{"asset":"ADA","assetFullName":"ADA","amountFree":"6.21","toBTC":"0.00016848","toBNB":"0.01777302","toBNBOffExchange":"0.01741756","exchange":"0.00035546"}],"totalTransferBtc":"0.00016848","totalTransferBNB":"0.01777302","dribbletPercentage":"0.02"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/dust-btc", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetAssetsThatCanBeConvertedIntoBnbResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetAssetsThatCanBeConvertedIntoBnb(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetAssetsThatCanBeConvertedIntoBnbResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetAssetsThatCanBeConvertedIntoBnbResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService GetAssetsThatCanBeConvertedIntoBnb Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetAssetsThatCanBeConvertedIntoBnb(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService GetCloudMiningPaymentAndRefundHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":5,"rows":[{"createTime":1667880112000,"tranId":121230610120,"type":248,"asset":"USDT","amount":"25.0068","status":"S"},{"createTime":1666776366000,"tranId":119991507468,"type":249,"asset":"USDT","amount":"0.027","status":"S"},{"createTime":1666764505000,"tranId":119977966327,"type":248,"asset":"USDT","amount":"0.027","status":"S"},{"createTime":1666758189000,"tranId":119973601721,"type":248,"asset":"USDT","amount":"0.018","status":"S"},{"createTime":1666757278000,"tranId":119973028551,"type":248,"asset":"USDT","amount":"0.018","status":"S"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage", r.URL.Path)
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCloudMiningPaymentAndRefundHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetCloudMiningPaymentAndRefundHistory(context.Background()).StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCloudMiningPaymentAndRefundHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCloudMiningPaymentAndRefundHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService GetCloudMiningPaymentAndRefundHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetCloudMiningPaymentAndRefundHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService GetCloudMiningPaymentAndRefundHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetCloudMiningPaymentAndRefundHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService GetOpenSymbolList Success", func(t *testing.T) {

		mockedJSON := `[{"openTime":1686161202000,"symbols":["BNBBTC","BNBETH"]},{"openTime":1686222232000,"symbols":["BTCUSDT"]}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/spot/open-symbol-list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOpenSymbolListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetOpenSymbolList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOpenSymbolListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOpenSymbolListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService GetOpenSymbolList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.GetOpenSymbolList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService QueryUserDelegationHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":3316,"rows":[{"clientTranId":"293915932290879488","transferType":"Undelegate","asset":"ETH","amount":"1","time":1695205406000},{"clientTranId":"293915892281413632","transferType":"Delegate","asset":"ETH","amount":"1","time":1695205396000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/custody/transfer-history", r.URL.Path)
			require.Equal(t, "email_example", r.URL.Query().Get("email"))
			require.Equal(t, "1623319461670", r.URL.Query().Get("startTime"))
			require.Equal(t, "1641782889000", r.URL.Query().Get("endTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryUserDelegationHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserDelegationHistory(context.Background()).Email("email_example").StartTime(int64(1623319461670)).EndTime(int64(1641782889000)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryUserDelegationHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryUserDelegationHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService QueryUserDelegationHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserDelegationHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService QueryUserDelegationHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserDelegationHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService QueryUserUniversalTransferHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":2,"rows":[{"asset":"USDT","amount":"1","type":"MAIN_UMFUTURE","status":"CONFIRMED","tranId":11415955596,"timestamp":1544433328000},{"asset":"USDT","amount":"2","type":"MAIN_UMFUTURE","status":"CONFIRMED","tranId":11366865406,"timestamp":1544433328000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/transfer", r.URL.Path)
			require.Equal(t, "type__example", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryUserUniversalTransferHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserUniversalTransferHistory(context.Background()).Type("type__example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryUserUniversalTransferHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryUserUniversalTransferHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService QueryUserUniversalTransferHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserUniversalTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService QueryUserUniversalTransferHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserUniversalTransferHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService QueryUserWalletBalance Success", func(t *testing.T) {

		mockedJSON := `[{"activate":true,"balance":"0","walletName":"Spot"},{"activate":true,"balance":"0","walletName":"Funding"},{"activate":true,"balance":"0","walletName":"Cross Margin"},{"activate":true,"balance":"0","walletName":"Isolated Margin"},{"activate":true,"balance":"0.71842752","walletName":"USDⓈ-M Futures"},{"activate":true,"balance":"0","walletName":"COIN-M Futures"},{"activate":true,"balance":"0","walletName":"Earn"},{"activate":false,"balance":"0","walletName":"Options"},{"activate":true,"balance":"0","walletName":"Trading Bots"},{"activate":true,"balance":"0","walletName":"Copy Trading"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/wallet/balance", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryUserWalletBalanceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserWalletBalance(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryUserWalletBalanceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryUserWalletBalanceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService QueryUserWalletBalance Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.QueryUserWalletBalance(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService ToggleBnbBurnOnSpotTradeAndMarginInterest Success", func(t *testing.T) {

		mockedJSON := `{"spotBNBBurn":true,"interestBNBBurn":false}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/bnbBurn", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.ToggleBnbBurnOnSpotTradeAndMarginInterest(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ToggleBnbBurnOnSpotTradeAndMarginInterestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService ToggleBnbBurnOnSpotTradeAndMarginInterest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.ToggleBnbBurnOnSpotTradeAndMarginInterest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService TradeFee Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"ADABNB","makerCommission":"0.001","takerCommission":"0.001"},{"symbol":"BNBBTC","makerCommission":"0.001","takerCommission":"0.001"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/tradeFee", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TradeFeeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.TradeFee(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TradeFeeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TradeFeeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService TradeFee Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.TradeFee(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService UserAsset Success", func(t *testing.T) {

		mockedJSON := `[{"asset":"AVAX","free":"1","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"},{"asset":"BCH","free":"0.9","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"},{"asset":"BNB","free":"887.47061626","locked":"0","freeze":"10.52","withdrawing":"0.1","ipoable":"0","btcValuation":"0"},{"asset":"BUSD","free":"9999.7","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"},{"asset":"SHIB","free":"532.32","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"},{"asset":"USDT","free":"50300000001.44911105","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"},{"asset":"WRZ","free":"1","locked":"0","freeze":"0","withdrawing":"0","ipoable":"0","btcValuation":"0"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v3/asset/getUserAsset", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.UserAssetResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.UserAsset(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.UserAssetResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.UserAssetResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService UserAsset Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.UserAsset(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService UserUniversalTransfer Success", func(t *testing.T) {

		mockedJSON := `{"tranId":13526853623}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/asset/transfer", r.URL.Path)
			require.Equal(t, "type__example", r.URL.Query().Get("type"))
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.UserUniversalTransferResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.UserUniversalTransfer(context.Background()).Type("type__example").Asset("asset_example").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.UserUniversalTransferResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.UserUniversalTransferResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AssetAPIService UserUniversalTransfer Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.UserUniversalTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AssetAPIService UserUniversalTransfer Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AssetAPI.UserUniversalTransfer(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
