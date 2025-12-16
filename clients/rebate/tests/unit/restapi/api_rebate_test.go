/*
Binance Rebate REST API TEST

Testing RebateAPIService

*/

package binancerebaterestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/rebate"
	"github.com/binance/binance-connector-go/clients/rebate/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancerebaterestapi_RebateAPIService(t *testing.T) {

	t.Run("Test RebateAPIService GetSpotRebateHistoryRecords Success", func(t *testing.T) {

		mockedJSON := `{"status":"OK","type":"GENERAL","code":"000000000","data":{"page":1,"totalRecords":2,"totalPageNum":1,"data":[{"asset":"USDT","type":1,"amount":"0.0001126","updateTime":1637651320000},{"asset":"ETH","type":1,"amount":"0.00000056","updateTime":1637928379000}]}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/rebate/taxQuery", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSpotRebateHistoryRecordsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceRebateClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RebateAPI.GetSpotRebateHistoryRecords(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSpotRebateHistoryRecordsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSpotRebateHistoryRecordsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test RebateAPIService GetSpotRebateHistoryRecords Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceRebateClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.RebateAPI.GetSpotRebateHistoryRecords(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
