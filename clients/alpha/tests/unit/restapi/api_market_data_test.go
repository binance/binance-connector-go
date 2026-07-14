/*
Alpha Trading REST API TEST

Testing MarketDataAPIService

*/

package binancealpharestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/clients/alpha/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancealpharestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService AggregatedTrades Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"","messageDetail":"","data":[{"a":58470,"p":"1.00","q":"1.00","f":58470,"l":58665,"T":1752568680000,"m":true}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/agg-trades", r.URL.Path)
			require.Equal(t, "ALPHA_118USDC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AggregatedTradesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AggregatedTrades(context.Background()).Symbol("ALPHA_118USDC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AggregatedTradesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AggregatedTradesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService AggregatedTrades Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AggregatedTrades(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService AggregatedTrades Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AggregatedTrades(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService FullDepth Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"message","messageDetail":"messageDetail","success":true,"data":{"lastUpdateId":47534656223,"symbol":"ALPHA_175USDT","bids":[["0.00040161","365980.85000000"]],"asks":[["0.00046996","61994.70000000"]],"E":1775027836086,"T":1775027836072}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/fullDepth", r.URL.Path)
			require.Equal(t, "ALPHA_175USDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FullDepthResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.FullDepth(context.Background()).Symbol("ALPHA_175USDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FullDepthResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FullDepthResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService FullDepth Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.FullDepth(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService FullDepth Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.FullDepth(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetExchangeInfo Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"","messageDetail":"","success":true,"data":{"timezone":"UTC","assets":[{"asset":"USDT"}],"symbols":[{"symbol":"ALPHA_105USDT","status":"TRADING","baseAsset":"ALPHA_105","quoteAsset":"USDT","pricePrecision":8,"quantityPrecision":8,"baseAssetPrecision":8,"quotePrecision":8,"filters":[{"filterType":"PRICE_FILTER","minPrice":"0.00000001","maxPrice":"1000","tickSize":"0.00000001","stepSize":"0.01000000","maxQty":"277778","minQty":"0.01000000","limit":200,"minNotional":"0.1","maxNotional":"1000000","multiplierDown":"0.20000","multiplierUp":"5","bidMultiplierUp":"5","askMultiplierUp":"5","bidMultiplierDown":"0.20000","askMultiplierDown":"0.20000"}],"orderTypes":["LIMIT"]}],"orderTypes":""}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/get-exchange-info", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetExchangeInfoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetExchangeInfo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetExchangeInfoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetExchangeInfoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetExchangeInfo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetExchangeInfo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Klines Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"","messageDetail":"","success":true,"data":[["1752642000000"]]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/klines", r.URL.Path)
			require.Equal(t, "ALPHA_175USDT", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.KlinesIntervalParameterInterval1s), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.KlinesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Klines(context.Background()).Symbol("ALPHA_175USDT").Interval(models.KlinesIntervalParameterInterval1s).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.KlinesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.KlinesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService Klines Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Klines(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Klines Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Klines(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Ticker Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"","messageDetail":"","data":{"symbol":"ALPHA_175USDT","priceChange":"-0.00007172","priceChangePercent":"-8.841","weightedAvgPrice":"0.00079608","lastPrice":"0.00073954","lastQty":"12600.43000000","openPrice":"0.00081126","highPrice":"0.00081126","lowPrice":"0.00073954","volume":"1204754.30000000","quoteVolume":"959.07729927","openTime":1768808100000,"closeTime":1768893244772,"firstId":93742,"lastId":93768,"count":38},"success":true}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/ticker", r.URL.Path)
			require.Equal(t, "ALPHA_175USDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker(context.Background()).Symbol("ALPHA_175USDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TickerResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TickerResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService Ticker Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Ticker Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TokenList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"code":"000000","message":"","messageDetail":"","success":true,"data":[{"tokenId":"3F350C8B3621770A673159B7A19BC034","chainId":"56","chainIconUrl":"https://bin.bnbstatic.com/image/admin_mgs_image_upload/20250228/d0216ce4-a3e9-4bda-8937-4a6aa943ccf2.png","chainName":"BSC","contractAddress":"0xcf640fdf9b3d9e45cbd69fda91d7e22579c14444","name":"gorilla","symbol":"gorilla","iconUrl":"https://bin.bnbstatic.com/images/web3-data/public/token/logos/248b5406f88a4ee28913a29107875339.png","price":"0.00080595003978242023","percentChange24h":"-7.42","volume24h":"14847.711222633409558029675","marketCap":"805950.03978242","fdv":"805950.03978242","liquidity":"263192.72813121626034","totalSupply":"1000000000","circulatingSupply":"1000000000","holders":"7120","decimals":18,"listingCex":false,"hotTag":false,"cexCoinName":"","canTransfer":false,"denomination":1,"offline":false,"tradeDecimal":8,"alphaId":"ALPHA_175","offsell":false,"priceHigh24h":"0.00088873864475645879","priceLow24h":"0.0008020805165487083","count24h":"166","onlineTge":false,"onlineAirdrop":false,"score":1,"cexOffDisplay":false,"stockState":false,"listingTime":1746686700000,"mulPoint":1,"bnExclusiveState":false}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TokenListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TokenList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TokenListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TokenListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService TokenList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TokenList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
