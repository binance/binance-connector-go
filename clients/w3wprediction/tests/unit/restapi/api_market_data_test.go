/*
Prediction Trading REST API TEST

Testing MarketDataAPIService

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

func Test_binancew3wpredictionrestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService GetMarketDetail Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"marketTopicId":4229564,"vendor":"PREDICT_FUN","chainId":"56","slug":"btc-price-1h-up-or-down","title":"BTC Price 1h Up or Down?","question":"Will BTC price go UP in the next 1 hour?","description":"Resolves YES if BTC spot price is higher than the starting price at resolution time.","imageUrl":"https://cdn.example.com/prediction/btc-1h.png","topicType":"FLAT","chartType":"CRYPTO_UP_DOWN","symbol":"BTCUSDT","variantData":{"type":"CRYPTO_UP_DOWN","startPrice":"67890.12","endPrice":"endPrice","priceFeedId":"0xpricefeedid123","priceFeedProvider":"BINANCE","priceFeedSymbol":"BTCUSDT"},"participantCount":3420,"collateral":"USDT","feeRateBps":200,"slippageBps":1200,"isYieldBearing":false,"tradeVolume":"158234.56","liquidity":"45000.00","publishedAt":1748100000000,"startDate":1748131200000,"endDate":1748134800000,"status":"REGISTERED","timeline":[{"marketTopicId":4229560,"startDate":1748127600000,"endDate":1748131200000}],"markets":[{"marketId":5567895,"externalId":"ext_001","title":"UP","question":"Will BTC go UP?","description":"Resolves YES if BTC price increases.","conditionId":"0xabc123","status":"REGISTERED","tradingStatus":"OPEN","tradeVolume":"90000.00","liquidity":"25000.00","decimalPrecision":2,"outcomes":[{"name":"YES","price":"0.52","chance":"0.52","index":0,"tokenId":"112233"}]}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/market/detail", r.URL.Path)
			require.Equal(t, "4229564", r.URL.Query().Get("marketTopicId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetMarketDetailResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetMarketDetail(context.Background()).MarketTopicId(int64(4229564)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetMarketDetailResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetMarketDetailResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetMarketDetail Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetMarketDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetMarketDetail Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.GetMarketDetail(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ListPredictionCategories Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"categories":[{"id":"crypto","name":"Crypto","icon":"crypto_icon","order":1,"subcategories":[{"id":"up-down","name":"Up/Down","icon":"updown_icon","order":1}]}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/category/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ListPredictionCategoriesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionCategories(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ListPredictionCategoriesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ListPredictionCategoriesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService ListPredictionCategories Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionCategories(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ListPredictionMarkets Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"marketTopics":[{"marketTopicId":4229564,"vendor":"PREDICT_FUN","chainId":"56","slug":"btc-price-1h-up-or-down","title":"BTC Price 1h Up or Down?","question":"Will BTC price go UP in the next 1 hour?","description":"Resolves YES if BTC spot price is higher than the starting price at resolution time.","imageUrl":"https://cdn.example.com/prediction/btc-1h.png","topicType":"FLAT","chartType":"CRYPTO_UP_DOWN","symbol":"BTCUSDT","participantCount":3420,"collateral":"USDT","feeRateBps":200,"slippageBps":1200,"isYieldBearing":false,"tradeVolume":"158234.56","liquidity":"45000.00","publishedAt":1748100000000,"startDate":1748131200000,"endDate":1748134800000,"status":"REGISTERED","markets":[{}]}],"total":128,"offset":0,"limit":20,"hasMore":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/market/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ListPredictionMarketsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionMarkets(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ListPredictionMarketsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ListPredictionMarketsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService ListPredictionMarkets Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.ListPredictionMarkets(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarketSearch Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"marketTopicId":4229564,"vendor":"PREDICT_FUN","chainId":"56","slug":"btc-price-1h-up-or-down","title":"BTC Price 1h Up or Down?","question":"Will BTC price go UP in the next 1 hour?","description":"Resolves YES if BTC spot price is higher than the starting price.","topicType":"FLAT","chartType":"CRYPTO_UP_DOWN","symbol":"BTCUSDT","participantCount":3420,"collateral":"USDT","tradeVolume":"158234.56","liquidity":"45000.00","startDate":1748131200000,"endDate":1748134800000,"status":"REGISTERED","markets":[{}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/market/search", r.URL.Path)
			require.Equal(t, "BTC price", r.URL.Query().Get("query"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarketSearchResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarketSearch(context.Background()).Query("BTC price").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarketSearchResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarketSearchResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService MarketSearch Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarketSearch(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarketSearch Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.MarketSearch(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryLastTradePrice Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"marketId":5567895,"lastTradePrice":"0.52"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/order-book/last-trade-price", r.URL.Path)
			require.Equal(t, "5567895", r.URL.Query().Get("marketId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryLastTradePriceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryLastTradePrice(context.Background()).MarketId(int64(5567895)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryLastTradePriceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryLastTradePriceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService QueryLastTradePrice Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryLastTradePrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryLastTradePrice Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.QueryLastTradePrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryOrderBook Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"outcome":"YES","tokenId":"112233","timestamp":1748131800000,"bids":[{"price":"0.51","size":"5000.00"}],"asks":[{"price":"0.52","size":"3000.00"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/order-book", r.URL.Path)
			require.Equal(t, "predict_fun", r.URL.Query().Get("vendor"))
			require.Equal(t, "5567895", r.URL.Query().Get("marketId"))
			require.Equal(t, "112233", r.URL.Query().Get("tokenId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryOrderBookResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryOrderBook(context.Background()).Vendor("predict_fun").MarketId(int64(5567895)).TokenId("112233").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryOrderBookResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryOrderBookResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService QueryOrderBook Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryOrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryOrderBook Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.MarketDataAPI.QueryOrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
