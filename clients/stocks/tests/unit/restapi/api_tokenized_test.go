/*
Stocks Trading REST API TEST

Testing TokenizedAPIService

*/

package binancestocksrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancestocksrestapi_TokenizedAPIService(t *testing.T) {

	t.Run("Test TokenizedAPIService TokenizedConvertHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"rows":[{"underlyingAsset":"AAPL","underlyingAssetAmount":"1.000000000","tokenizedAsset":"AAPLB","tokenizedAssetAmount":"1.000000000","issuerRequestId":"mint-20260505-8f3b9e1a2d3c4b5a","convertType":"MINT","status":"S","createdAt":1735900000000,"updatedAt":1735900060000}],"hasMore":true,"nextLastId":10020}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/equity/tokenized/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TokenizedConvertHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TokenizedConvertHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TokenizedConvertHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TokenizedAPIService TokenizedConvertHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedConvertStatus Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"underlyingAsset":"AAPL","underlyingAssetAmount":"1.000000000","tokenizedAsset":"AAPLB","tokenizedAssetAmount":"1.000000000","issuerRequestId":"mint-20260505-8f3b9e1a2d3c4b5a","convertType":"MINT","status":"S","createdAt":1735900000000,"updatedAt":1735900060000}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/equity/tokenized/convert-status", r.URL.Path)
			require.Equal(t, "mint-20260505-8f3b9e1a2d3c4b5a", r.URL.Query().Get("issuerRequestId"))
			require.Equal(t, string(models.TokenizedConvertStatusConvertTypeParameterMint), r.URL.Query().Get("convertType"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TokenizedConvertStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertStatus(context.Background()).IssuerRequestId("mint-20260505-8f3b9e1a2d3c4b5a").ConvertType(models.TokenizedConvertStatusConvertTypeParameterMint).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TokenizedConvertStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TokenizedConvertStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TokenizedAPIService TokenizedConvertStatus Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedConvertStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedMint Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"issuerRequestId":"mint-20260505-8f3b9e1a2d3c4b5a","status":"P"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/equity/tokenized/mint", r.URL.Path)
			require.Equal(t, "AAPL", r.URL.Query().Get("underlyingAsset"))
			require.Equal(t, "1", r.URL.Query().Get("underlyingAssetAmount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TokenizedMintResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedMint(context.Background()).UnderlyingAsset("AAPL").UnderlyingAssetAmount("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TokenizedMintResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TokenizedMintResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TokenizedAPIService TokenizedMint Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedMint(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedMint Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedMint(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedRedeem Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"issuerRequestId":"redeem-20260505-a1b2c3d4e5f6","status":"P"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/equity/tokenized/redeem", r.URL.Path)
			require.Equal(t, "AAPLB", r.URL.Query().Get("tokenizedAsset"))
			require.Equal(t, "1", r.URL.Query().Get("tokenizedAssetAmount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TokenizedRedeemResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedRedeem(context.Background()).TokenizedAsset("AAPLB").TokenizedAssetAmount("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TokenizedRedeemResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TokenizedRedeemResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TokenizedAPIService TokenizedRedeem Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedRedeem(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TokenizedAPIService TokenizedRedeem Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStocksClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TokenizedAPI.TokenizedRedeem(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
