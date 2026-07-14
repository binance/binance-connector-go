/*
Options REST API TEST

Testing AccountAPIService

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
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingoptionsrestapi_AccountAPIService(t *testing.T) {

	t.Run("Test AccountAPIService AccountFundingFlow Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"id":1125899906842624000,"asset":"USDT","amount":"-0.552","type":"FEE","createDate":1592449456000}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/bill", r.URL.Path)
			require.Equal(t, string(models.AccountFundingFlowCurrencyParameterUsdt), r.URL.Query().Get("currency"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountFundingFlowResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountFundingFlow(context.Background()).Currency(models.AccountFundingFlowCurrencyParameterUsdt).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountFundingFlowResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountFundingFlowResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AccountFundingFlow Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountFundingFlow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AccountFundingFlow Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountAPI.AccountFundingFlow(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService OptionMarginAccountInformation Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"asset":[{"asset":"USDT","marginBalance":"10099.448","equity":"10094.44662","available":"8725.92524","initialMargin":"1084.52138","maintMargin":"151.00138","unrealizedPNL":"-5.00138","adjustedEquity":"34.13282285"}],"greek":[{"underlying":"BTCUSDT","delta":"-0.05","gamma":"-0.002","theta":"-0.05","vega":"-0.002"}],"time":1762843368098,"canTrade":true,"canDeposit":true,"canWithdraw":true,"reduceOnly":false,"tradeGroupId":-1}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/eapi/v1/marginAccount", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OptionMarginAccountInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingOptionsClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OptionMarginAccountInformation(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OptionMarginAccountInformationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OptionMarginAccountInformationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService OptionMarginAccountInformation Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.AccountAPI.OptionMarginAccountInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
