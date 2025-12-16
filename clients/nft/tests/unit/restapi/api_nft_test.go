/*
Binance NFT REST API TEST

Testing NFTAPIService

*/

package binancenftrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/clients/nft/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancenftrestapi_NFTAPIService(t *testing.T) {

	t.Run("Test NFTAPIService GetNFTAsset Success", func(t *testing.T) {

		mockedJSON := `{"total":347,"list":[{"network":"BSC","contractAddress":"REGULAR11234567891779","tokenId":"100900000017"},{"network":"BSC","contractAddress":"SSMDQ8W59","tokenId":"200500000011"},{"network":"BSC","contractAddress":"SSMDQ8W59","tokenId":"200500000019"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/nft/user/getAsset", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetNFTAssetResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTAsset(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetNFTAssetResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetNFTAssetResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test NFTAPIService GetNFTAsset Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTAsset(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test NFTAPIService GetNFTDepositHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":2,"list":[{"network":"ETH","txID":null,"contractAdrress":"0xe507c961ee127d4439977a61af39c34eafee0dc6","tokenId":"10014","timestamp":1629986047000},{"network":"BSC","txID":null,"contractAdrress":"0x058451b463bab04f52c0799d55c4094f507acfa9","tokenId":"10016","timestamp":1630083581000}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/nft/history/deposit", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetNFTDepositHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTDepositHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetNFTDepositHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetNFTDepositHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test NFTAPIService GetNFTDepositHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTDepositHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test NFTAPIService GetNFTTransactionHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":2,"list":[{"orderNo":"1_470502070600699904","tokens":[{"network":"BSC","tokenId":"216000000496","contractAddress":"MYSTERY_BOX0000087"}],"tradeTime":1626941236000,"tradeAmount":"19.60000000","tradeCurrency":"BNB"},{"orderNo":"1_488306442479116288","tokens":[{"network":"BSC","tokenId":"132900000007","contractAddress":"0xAf12111a592e408DAbC740849fcd5e68629D9fb6"}],"tradeTime":1631186130000,"tradeAmount":"192.00000000","tradeCurrency":"BNB"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/nft/history/transactions", r.URL.Path)
			require.Equal(t, "789", r.URL.Query().Get("orderType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetNFTTransactionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTTransactionHistory(context.Background()).OrderType(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetNFTTransactionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetNFTTransactionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test NFTAPIService GetNFTTransactionHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTTransactionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test NFTAPIService GetNFTTransactionHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTTransactionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test NFTAPIService GetNFTWithdrawHistory Success", func(t *testing.T) {

		mockedJSON := `{"total":178,"list":[{"network":"ETH","txID":"0x2be5eed31d787fdb4880bc631c8e76bdfb6150e137f5cf1732e0416ea206f57f","contractAdrress":"0xe507c961ee127d4439977a61af39c34eafee0dc6","tokenId":"1000001247","timestamp":1633674433000,"fee":0.1,"feeAsset":"ETH"},{"network":"ETH","txID":"0x3b3aea5c0a4faccd6f306641e6deb9713ab229ac233be3be227f580311e4362a","contractAdrress":"0xe507c961ee127d4439977a61af39c34eafee0dc6","tokenId":"40000030","timestamp":1633677022000,"fee":0.1,"feeAsset":"ETH"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/nft/history/withdraw", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetNFTWithdrawHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTWithdrawHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetNFTWithdrawHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetNFTWithdrawHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test NFTAPIService GetNFTWithdrawHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceNFTClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.NFTAPI.GetNFTWithdrawHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
