/*
Binance Copy Trading REST API TEST

Testing FutureCopyTradingAPIService

*/

package binancecopytradingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/copytrading"
	"github.com/binance/binance-connector-go/clients/copytrading/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancecopytradingrestapi_FutureCopyTradingAPIService(t *testing.T) {

	t.Run("Test FutureCopyTradingAPIService GetFuturesLeadTraderStatus Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":{"isLeadTrader":true,"time":1717382310843},"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/copyTrading/futures/userStatus", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesLeadTraderStatusResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCopyTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTraderStatus(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesLeadTraderStatusResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesLeadTraderStatusResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureCopyTradingAPIService GetFuturesLeadTraderStatus Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCopyTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTraderStatus(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FutureCopyTradingAPIService GetFuturesLeadTradingSymbolWhitelist Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":[{"symbol":"BTCUSDT","baseAsset":"BTC","quoteAsset":"USDT"},{"symbol":"ETHUSDT","baseAsset":"ETH","quoteAsset":"USDT"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/copyTrading/futures/leadSymbol", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFuturesLeadTradingSymbolWhitelistResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceCopyTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTradingSymbolWhitelist(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFuturesLeadTradingSymbolWhitelistResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFuturesLeadTradingSymbolWhitelistResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FutureCopyTradingAPIService GetFuturesLeadTradingSymbolWhitelist Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceCopyTradingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FutureCopyTradingAPI.GetFuturesLeadTradingSymbolWhitelist(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
