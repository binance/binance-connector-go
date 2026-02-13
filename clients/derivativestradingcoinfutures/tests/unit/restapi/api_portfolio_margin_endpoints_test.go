/*
Binance Derivatives Trading COIN Futures REST API TEST

Testing PortfolioMarginEndpointsAPIService

*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingcoinfuturesrestapi_PortfolioMarginEndpointsAPIService(t *testing.T) {

	t.Run("Test PortfolioMarginEndpointsAPIService ClassicPortfolioMarginAccountInformation Success", func(t *testing.T) {

		mockedJSON := `{"maxWithdrawAmountUSD":"25347.92083245","asset":"BTC","maxWithdrawAmount":"1.33663654"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/pmAccountInfo", r.URL.Path)
			require.Equal(t, "asset_example", r.URL.Query().Get("asset"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ClassicPortfolioMarginAccountInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PortfolioMarginEndpointsAPI.ClassicPortfolioMarginAccountInformation(context.Background()).Asset("asset_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ClassicPortfolioMarginAccountInformationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ClassicPortfolioMarginAccountInformationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PortfolioMarginEndpointsAPIService ClassicPortfolioMarginAccountInformation Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PortfolioMarginEndpointsAPI.ClassicPortfolioMarginAccountInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PortfolioMarginEndpointsAPIService ClassicPortfolioMarginAccountInformation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PortfolioMarginEndpointsAPI.ClassicPortfolioMarginAccountInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
