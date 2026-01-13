/*
Binance Derivatives Trading Options REST API TEST

Testing MarketMakerBlockTradeAPIService

*/

package binancederivativestradingoptionsrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingoptions"
	"github.com/binance/binance-connector-go/clients/derivativestradingoptions/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingoptionsrestapi_MarketMakerBlockTradeAPIService(t *testing.T) {

	t.Run("Test MarketMakerBlockTradeAPIService AcceptBlockTradeOrder Success", func(t *testing.T) {

		mockedJSON := `{"blockTradeSettlementKey":"7d046e6e-a429-4335-ab9d-6a681febcde5","expireTime":1730172115801,"liquidity":"MAKER","status":"ACCEPTED","createTime":1730170315803,"legs":[{"symbol":"BNB-241101-700-C","side":"SELL","quantity":"1.2","price":"2.8"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/execute", r.URL.Path)
			require.Equal(t, "blockOrderMatchingKey_example", r.URL.Query().Get("blockOrderMatchingKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AcceptBlockTradeOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AcceptBlockTradeOrder(context.Background()).BlockOrderMatchingKey("blockOrderMatchingKey_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AcceptBlockTradeOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AcceptBlockTradeOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService AcceptBlockTradeOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AcceptBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService AcceptBlockTradeOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AcceptBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService AccountBlockTradeList Success", func(t *testing.T) {

		mockedJSON := `[{"parentOrderId":"4675011431944499201","crossType":"USER_BLOCK","legs":[{"createTime":1730170445600,"updateTime":1730170445600,"symbol":"BNB-241101-700-C","orderId":"4675011431944499203","orderPrice":2.8,"orderQuantity":1.2,"orderStatus":"FILLED","executedQty":1.2,"executedAmount":3.36,"fee":0.336,"orderType":"PREV_QUOTED","orderSide":"BUY","id":"1125899906900937837","tradeId":1,"tradePrice":2.8,"tradeQty":1.2,"tradeTime":1730170445600,"liquidity":"TAKER","commission":0.336}],"blockTradeSettlementKey":"7d085e6e-a229-2335-ab9d-6a581febcd25"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/user-trades", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountBlockTradeListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AccountBlockTradeList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountBlockTradeListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountBlockTradeListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService AccountBlockTradeList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.AccountBlockTradeList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService CancelBlockTradeOrder Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/create", r.URL.Path)
			require.Equal(t, "blockOrderMatchingKey_example", r.URL.Query().Get("blockOrderMatchingKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).BlockOrderMatchingKey("blockOrderMatchingKey_example").Execute()
		require.NoError(t, err)
	})

	t.Run("Test MarketMakerBlockTradeAPIService CancelBlockTradeOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test MarketMakerBlockTradeAPIService CancelBlockTradeOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketMakerBlockTradeAPI.CancelBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test MarketMakerBlockTradeAPIService ExtendBlockTradeOrder Success", func(t *testing.T) {

		mockedJSON := `{"blockTradeSettlementKey":"3668822b8-1baa-6a2f-adb8-d3de6289b361","expireTime":1730172007000,"liquidity":"TAKER","status":"RECEIVED","createTime":1730170088111,"legs":[{"symbol":"BNB-241101-700-C","side":"BUY","quantity":"1.2","price":"2.8"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/create", r.URL.Path)
			require.Equal(t, "blockOrderMatchingKey_example", r.URL.Query().Get("blockOrderMatchingKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ExtendBlockTradeOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.ExtendBlockTradeOrder(context.Background()).BlockOrderMatchingKey("blockOrderMatchingKey_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ExtendBlockTradeOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ExtendBlockTradeOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService ExtendBlockTradeOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.ExtendBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService ExtendBlockTradeOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.ExtendBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService NewBlockTradeOrder Success", func(t *testing.T) {

		mockedJSON := `{"blockTradeSettlementKey":"3668822b8-1baa-6a2f-adb8-d3de6289b361","expireTime":1730171888109,"liquidity":"TAKER","status":"RECEIVED","legs":[{"symbol":"BNB-241101-700-C","side":"BUY","quantity":"1.2","price":"2.8"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/create", r.URL.Path)
			require.Equal(t, "liquidity_example", r.URL.Query().Get("liquidity"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.NewBlockTradeOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.NewBlockTradeOrder(context.Background()).Liquidity("liquidity_example").Legs([]map[string]interface{}{{}}).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.NewBlockTradeOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.NewBlockTradeOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService NewBlockTradeOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.NewBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService NewBlockTradeOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.NewBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService QueryBlockTradeDetails Success", func(t *testing.T) {

		mockedJSON := `{"blockTradeSettlementKey":"12b96c28-ba05-8906-c89t-703215cfb2e6","expireTime":1730171860460,"liquidity":"MAKER","status":"RECEIVED","createTime":1730170060462,"legs":[{"symbol":"BNB-241101-700-C","side":"SELL","quantity":"1.66","price":"20"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/execute", r.URL.Path)
			require.Equal(t, "blockOrderMatchingKey_example", r.URL.Query().Get("blockOrderMatchingKey"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryBlockTradeDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeDetails(context.Background()).BlockOrderMatchingKey("blockOrderMatchingKey_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryBlockTradeDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryBlockTradeDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService QueryBlockTradeDetails Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService QueryBlockTradeDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerBlockTradeAPIService QueryBlockTradeOrder Success", func(t *testing.T) {

		mockedJSON := `[{"blockTradeSettlementKey":"7d046e6e-a429-4335-ab9d-6a681febcde5","expireTime":1730172115801,"liquidity":"TAKER","status":"RECEIVED","createTime":1730170315803,"legs":[{"symbol":"BNB-241101-700-C","side":"BUY","quantity":"1.2","price":"2.8"}]},{"blockTradeSettlementKey":"28b96c28-ba05-6906-a47c-703215cfbfe6","expireTime":1730171860460,"liquidity":"TAKER","status":"RECEIVED","createTime":1730170060462,"legs":[{"symbol":"BNB-241101-700-C","side":"BUY","quantity":"1.66","price":"20"}]}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/block/order/orders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryBlockTradeOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeOrder(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryBlockTradeOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryBlockTradeOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerBlockTradeAPIService QueryBlockTradeOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerBlockTradeAPI.QueryBlockTradeOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
