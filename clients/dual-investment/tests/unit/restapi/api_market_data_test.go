/*
Binance Dual Investment REST API TEST

Testing MarketDataAPIService

*/

package binancedualinvestmentrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/dualinvestment/src"
	"github.com/binance/binance-connector-go/clients/dualinvestment/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancedualinvestmentrestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService GetDualInvestmentProductList Success", func(t *testing.T) {

		mockedJSON := `{"total":1,"list":[{"id":"741590","investCoin":"USDT","exercisedCoin":"BNB","strikePrice":"380","duration":4,"settleDate":1709020800000,"purchaseDecimal":8,"purchaseEndTime":1708934400000,"canPurchase":true,"apr":"0.6076","orderId":8257205859,"minAmount":"0.1","maxAmount":"25265.7","createTimestamp":1708560084000,"optionType":"PUT","isAutoCompoundEnable":true,"autoCompoundPlanList":["STANDARD","ADVANCE"]}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/dci/product/list", r.URL.Path)
			require.Equal(t, "optionType_example", r.URL.Query().Get("optionType"))
			require.Equal(t, "exercisedCoin_example", r.URL.Query().Get("exercisedCoin"))
			require.Equal(t, "investCoin_example", r.URL.Query().Get("investCoin"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetDualInvestmentProductListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDualInvestmentClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetDualInvestmentProductList(context.Background()).OptionType("optionType_example").ExercisedCoin("exercisedCoin_example").InvestCoin("investCoin_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetDualInvestmentProductListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetDualInvestmentProductListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetDualInvestmentProductList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDualInvestmentClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetDualInvestmentProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetDualInvestmentProductList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDualInvestmentClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetDualInvestmentProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
