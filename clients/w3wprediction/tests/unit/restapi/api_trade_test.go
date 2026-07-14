/*
Prediction Trading REST API TEST

Testing TradeAPIService

*/

package binancew3wpredictionrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/w3wprediction"
	"github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancew3wpredictionrestapi_TradeAPIService(t *testing.T) {

	t.Run("Test TradeAPIService BatchCancelOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"canceled":["54124"],"failed":[{"orderId":"54126","reason":"ORDER_NOT_FOUND"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/trade/batch-cancel", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "5b5c1ec3be4e4416a5872b21c1ca5d20", r.URL.Query().Get("walletId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.BatchCancelOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.BatchCancelOrders(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.BatchCancelOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.BatchCancelOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService BatchCancelOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.BatchCancelOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService BatchCancelOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.BatchCancelOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetQuote Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"quoteId":"q_20260525_abc123xyz","tokenId":"112233","chance":"0.52","vendor":"PREDICT_FUN","marketTitle":"UP","marketExtId":"ext_001","side":"BUY","amountIn":"1000000000000000000","amountOut":"1923070000000000000","isMinAmountOut":false,"feeAmount":"20000000000000000","feeDiscountBps":"0","averagePrice":0.52,"lastPrice":0.52,"priceImpact":0.001,"timestamp":1748131500000,"chainId":"56","userId":100103755893,"walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","orderType":"MARKET","slippageBps":1200,"feeRateBps":200,"minReceive":"1900000000000000000","expireAt":1748131800000,"priceLimit":"priceLimit"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/trade/get-quote", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "112233", r.URL.Query().Get("tokenId"))
			require.Equal(t, string(models.GetQuoteSideParameterBuy), r.URL.Query().Get("side"))
			require.Equal(t, "1000000000000000000", r.URL.Query().Get("amountIn"))
			require.Equal(t, string(models.GetQuoteOrderTypeParameterMarket), r.URL.Query().Get("orderType"))
			require.Equal(t, "1200", r.URL.Query().Get("slippageBps"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetQuoteResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetQuote(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").TokenId("112233").Side(models.GetQuoteSideParameterBuy).AmountIn("1000000000000000000").OrderType(models.GetQuoteOrderTypeParameterMarket).SlippageBps(int32(1200)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetQuoteResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetQuoteResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService GetQuote Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetQuote(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService GetQuote Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.GetQuote(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PlaceOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderId":"54124"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/trade/place-order-bundle", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "5b5c1ec3be4e4416a5872b21c1ca5d20", r.URL.Query().Get("walletId"))
			require.Equal(t, "q_20260525_abc123xyz", r.URL.Query().Get("quoteId"))
			require.Equal(t, "FOK", r.URL.Query().Get("timeInForce"))
			require.Equal(t, string(models.PlaceOrderAccountTypeParameterSpot), r.URL.Query().Get("accountType"))
			require.Equal(t, string(models.GetQuoteOrderTypeParameterMarket), r.URL.Query().Get("orderType"))
			require.Equal(t, "1200", r.URL.Query().Get("slippageBps"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.PlaceOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceOrder(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").WalletId("5b5c1ec3be4e4416a5872b21c1ca5d20").QuoteId("q_20260525_abc123xyz").TimeInForce("FOK").AccountType(models.PlaceOrderAccountTypeParameterSpot).OrderType(models.GetQuoteOrderTypeParameterMarket).SlippageBps(int32(1200)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.PlaceOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.PlaceOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService PlaceOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService PlaceOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.PlaceOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryActiveOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"total":2,"offset":0,"limit":20,"orders":[{"orderId":"54124","vendorOrderId":"0x1234abcd...","vendor":"PREDICT_FUN","marketTopicId":4229564,"slug":"btc-price-1h-up-or-down","marketTopicTitle":"BTC Price 1h Up or Down?","marketId":5567895,"marketTitle":"UP","outcome":"YES","outcomeIndex":0,"status":"OPENING","side":"BUY","orderType":"LIMIT","createTime":1748131500000,"modifyTime":1748131500000,"makerUsdtAmount":"1.00","makerShareQty":"2000.00","filledUsdtAmount":"0.00","filledShareQty":"0.00","fillPercentage":"0.00","price":"0.50","marketProviderFee":"0.02","networkFee":"0.000001"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/order/list", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryActiveOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryActiveOrders(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryActiveOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryActiveOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryActiveOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryActiveOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryActiveOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryActiveOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryOrderHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"total":15,"offset":0,"limit":20,"orders":[{"orderId":"54100","vendorOrderId":"0xabcd5678...","vendor":"PREDICT_FUN","marketTopicId":4229500,"slug":"btc-price-1h-up-or-down-prev","marketTopicTitle":"BTC Price 1h Up or Down?","marketId":5567800,"marketTitle":"UP","outcome":"YES","outcomeIndex":0,"status":"CLOSED","side":"BUY","orderType":"MARKET","createTime":1748045100000,"modifyTime":1748045101000,"terminalTime":1748045101000,"makerUsdtAmount":"1.00","makerShareQty":"1923.07","filledUsdtAmount":"1.00","filledShareQty":"1923.07","fillPercentage":"1.00","price":"0.52","marketProviderFee":"0.02","networkFee":"0.000001"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/order/history", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryOrderHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrderHistory(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryOrderHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryOrderHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TradeAPIService QueryOrderHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrderHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TradeAPIService QueryOrderHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TradeAPI.QueryOrderHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
