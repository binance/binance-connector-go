/*
Binance C2C REST API TEST

Testing C2CAPIService

*/

package binancec2crestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/c2c"
	"github.com/binance/binance-connector-go/clients/c2c/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancec2crestapi_C2CAPIService(t *testing.T) {

	t.Run("Test C2CAPIService GetC2CTradeHistory Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":"success","data":[{"orderNumber":"20219644646554779648","advNo":"11218246497340923904","tradeType":"SELL","asset":"BUSD","fiat":"CNY","fiatSymbol":"￥","amount":"5000.00000000","totalPrice":"33400.00000000","unitPrice":"6.68","orderStatus":"COMPLETED","createTime":1619361369000,"commission":"0","counterPartNickName":"ab***","advertisementRole":"TAKER"}],"total":1,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/c2c/orderMatch/listUserOrderHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetC2CTradeHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceC2CClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.C2CAPI.GetC2CTradeHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetC2CTradeHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetC2CTradeHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test C2CAPIService GetC2CTradeHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceC2CClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.C2CAPI.GetC2CTradeHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
