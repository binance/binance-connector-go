/*
Binance Derivatives Trading Portfolio Margin REST API TEST

Testing MarketDataAPIService

*/

package binancederivativestradingportfoliomarginrestapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingportfoliomargin/src"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingportfoliomarginrestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService TestConnectivity Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/papi/v1/ping", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
		require.NoError(t, err)
	})

	t.Run("Test MarketDataAPIService TestConnectivity Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingPortfolioMarginClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()

		require.Error(t, err)
	})

}
