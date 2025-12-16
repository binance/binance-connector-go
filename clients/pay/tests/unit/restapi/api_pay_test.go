/*
Binance Pay REST API TEST

Testing PayAPIService

*/

package binancepayrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/pay"
	"github.com/binance/binance-connector-go/clients/pay/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancepayrestapi_PayAPIService(t *testing.T) {

	t.Run("Test PayAPIService GetPayTradeHistory Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":[{"orderType":"C2C","transactionId":"M_P_71505104267788288","transactionTime":1610090460133,"amount":"23.72469206","currency":"BNB","walletType":1,"walletTypes":[1,2],"fundsDetail":[{"currency":"USDT","amount":"1.2","walletAssetCost":[{"1":"0.6"},{"2":"0.6"}]},{"currency":"ETH","amount":"0.0001","walletAssetCost":[{"1":"0.00005"},{"2":"0.00005"}]}],"payerInfo":{"name":"Jack","type":"USER","binanceId":"12345678"},"receiverInfo":{"name":"Alan","type":"MERCHANT","email":"alan@binance.com","binanceId":"34355667","accountId":"21326891","countryCode":"1","phoneNumber":"8057651210","mobileCode":"US","extend":{"institutionName":"","cardNumber":"","digitalWalletId":""}}}],"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/pay/transactions", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetPayTradeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinancePayClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetPayTradeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetPayTradeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PayAPIService GetPayTradeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinancePayClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PayAPI.GetPayTradeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
