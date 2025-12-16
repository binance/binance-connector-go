/*
Binance Derivatives Trading Options REST API TEST

Testing MarketMakerEndpointsAPIService

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

func Test_binancederivativestradingoptionsrestapi_MarketMakerEndpointsAPIService(t *testing.T) {

	t.Run("Test MarketMakerEndpointsAPIService AutoCancelAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `{"underlyings":["BTCUSDT","ETHUSDT"]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/countdownCancelAllHeartBeat", r.URL.Path)
			require.Equal(t, "underlyings_example", r.URL.Query().Get("underlyings"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AutoCancelAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.AutoCancelAllOpenOrders(context.Background()).Underlyings("underlyings_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AutoCancelAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AutoCancelAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService AutoCancelAllOpenOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.AutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService AutoCancelAllOpenOrders Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.AutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService GetAutoCancelAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `{"underlying":"ETHUSDT","countdownTime":100000}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/countdownCancelAll", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetAutoCancelAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetAutoCancelAllOpenOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetAutoCancelAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetAutoCancelAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService GetAutoCancelAllOpenOrders Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetAutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService GetMarketMakerProtectionConfig Success", func(t *testing.T) {

		mockedJSON := `{"underlyingId":2,"underlying":"BTCUSDT","windowTimeInMilliseconds":3000,"frozenTimeInMilliseconds":300000,"qtyLimit":"2","deltaLimit":"2.3","lastTriggerTime":0}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/mmp", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetMarketMakerProtectionConfigResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetMarketMakerProtectionConfig(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetMarketMakerProtectionConfigResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetMarketMakerProtectionConfigResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService GetMarketMakerProtectionConfig Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.GetMarketMakerProtectionConfig(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService ResetMarketMakerProtectionConfig Success", func(t *testing.T) {

		mockedJSON := `{"underlyingId":2,"underlying":"BTCUSDT","windowTimeInMilliseconds":3000,"frozenTimeInMilliseconds":300000,"qtyLimit":"2","deltaLimit":"2.3","lastTriggerTime":0}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/mmpReset", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ResetMarketMakerProtectionConfigResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.ResetMarketMakerProtectionConfig(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ResetMarketMakerProtectionConfigResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ResetMarketMakerProtectionConfigResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService ResetMarketMakerProtectionConfig Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.ResetMarketMakerProtectionConfig(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService SetAutoCancelAllOpenOrders Success", func(t *testing.T) {

		mockedJSON := `{"underlying":"ETHUSDT","countdownTime":30000}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/countdownCancelAll", r.URL.Path)
			require.Equal(t, "underlying_example", r.URL.Query().Get("underlying"))
			require.Equal(t, "789", r.URL.Query().Get("countdownTime"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetAutoCancelAllOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetAutoCancelAllOpenOrders(context.Background()).Underlying("underlying_example").CountdownTime(int64(789)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetAutoCancelAllOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetAutoCancelAllOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService SetAutoCancelAllOpenOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetAutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService SetAutoCancelAllOpenOrders Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetAutoCancelAllOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketMakerEndpointsAPIService SetMarketMakerProtectionConfig Success", func(t *testing.T) {

		mockedJSON := `{"underlyingId":2,"underlying":"BTCUSDT","windowTimeInMilliseconds":3000,"frozenTimeInMilliseconds":300000,"qtyLimit":"2","deltaLimit":"2.3","lastTriggerTime":0}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/mmpSet", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetMarketMakerProtectionConfigResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetMarketMakerProtectionConfig(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetMarketMakerProtectionConfigResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetMarketMakerProtectionConfigResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketMakerEndpointsAPIService SetMarketMakerProtectionConfig Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketMakerEndpointsAPI.SetMarketMakerProtectionConfig(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
