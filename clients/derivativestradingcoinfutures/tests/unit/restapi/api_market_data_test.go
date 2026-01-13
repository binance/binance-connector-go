/*
Binance Derivatives Trading COIN Futures REST API TEST

Testing MarketDataAPIService

*/

package binancederivativestradingcoinfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingcoinfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingcoinfuturesrestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService Basis Success", func(t *testing.T) {

		mockedJSON := `[{"indexPrice":"29269.93972727","contractType":"CURRENT_QUARTER","basisRate":"0.0024","futuresPrice":"29341.3","annualizedBasisRate":"0.0283","basis":"71.36027273","pair":"BTCUSD","timestamp":1653381600000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/basis", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisContractTypeParameterPerpetual), r.URL.Query().Get("contractType"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.BasisResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Pair("pair_example").ContractType(models.BasisContractTypeParameterPerpetual).Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.BasisResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.BasisResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService Basis Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Basis Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CheckServerTime Success", func(t *testing.T) {

		mockedJSON := `{"serverTime":1499827319559}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/time", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckServerTimeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CheckServerTime(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CheckServerTimeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CheckServerTimeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService CheckServerTime Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CheckServerTime(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CompressedAggregateTradesList Success", func(t *testing.T) {

		mockedJSON := `[{"a":416690,"p":"9642.4","q":"3","f":595259,"l":595259,"T":1591250548649,"m":false}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/aggTrades", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CompressedAggregateTradesListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CompressedAggregateTradesListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CompressedAggregateTradesListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService CompressedAggregateTradesList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CompressedAggregateTradesList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ContinuousContractKlineCandlestickData Success", func(t *testing.T) {

		mockedJSON := `[[1591258320000,"9640.7","9642.4","9640.6","9642.0","206",1591258379999,"2.13660389",48,"119","1.23424865","0"]]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/continuousKlines", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisContractTypeParameterPerpetual), r.URL.Query().Get("contractType"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ContinuousContractKlineCandlestickDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Pair("pair_example").ContractType(models.BasisContractTypeParameterPerpetual).Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ContinuousContractKlineCandlestickDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ContinuousContractKlineCandlestickDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService ContinuousContractKlineCandlestickData Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ContinuousContractKlineCandlestickData Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ExchangeInformation Success", func(t *testing.T) {

		mockedJSON := `{"exchangeFilters":[],"rateLimits":[{"interval":"MINUTE","intervalNum":1,"limit":6000,"rateLimitType":"REQUEST_WEIGHT"},{"interval":"MINUTE","intervalNum":1,"limit":6000,"rateLimitType":"ORDERS"}],"serverTime":1565613908500,"symbols":[{"filters":[{"filterType":"PRICE_FILTER","maxPrice":"100000","minPrice":"0.1","tickSize":"0.1"},{"filterType":"LOT_SIZE","maxQty":"100000","minQty":"1","stepSize":"1"},{"filterType":"MARKET_LOT_SIZE","maxQty":"100000","minQty":"1","stepSize":"1"},{"filterType":"MAX_NUM_ORDERS","limit":200},{"filterType":"PERCENT_PRICE","multiplierUp":"1.0500","multiplierDown":"0.9500","multiplierDecimal":"4"}],"OrderType":["LIMIT","MARKET","STOP","TAKE_PROFIT","TRAILING_STOP_MARKET"],"timeInForce":["GTC","IOC","FOK","GTX"],"liquidationFee":"0.010000","marketTakeBound":"0.30","symbol":"BTCUSD_200925","pair":"BTCUSD","contractType":"CURRENT_QUARTER","deliveryDate":1601020800000,"onboardDate":1590739200000,"contractStatus":"TRADING","contractSize":100,"quoteAsset":"USD","baseAsset":"BTC","marginAsset":"BTC","pricePrecision":1,"quantityPrecision":0,"baseAssetPrecision":8,"quotePrecision":8,"equalQtyPrecision":4,"triggerProtect":"0.0500","maintMarginPercent":"2.5000","requiredMarginPercent":"5.0000","underlyingType":"COIN","underlyingSubType":[]}],"timezone":"UTC"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/exchangeInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ExchangeInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInformation(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ExchangeInformationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ExchangeInformationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService ExchangeInformation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetFundingRateHistoryOfPerpetualFutures Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_PERP","fundingTime":1596038400000,"fundingRate":"-0.00300000"},{"symbol":"BTCUSD_PERP","fundingTime":1596067200000,"fundingRate":"-0.00300000"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/fundingRate", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFundingRateHistoryOfPerpetualFuturesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistoryOfPerpetualFutures(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFundingRateHistoryOfPerpetualFuturesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFundingRateHistoryOfPerpetualFuturesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetFundingRateHistoryOfPerpetualFutures Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistoryOfPerpetualFutures(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetFundingRateHistoryOfPerpetualFutures Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistoryOfPerpetualFutures(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetFundingRateInfo Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_PERP","adjustedFundingRateCap":"0.02500000","adjustedFundingRateFloor":"-0.02500000","fundingIntervalHours":8,"disclaimer":false}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/fundingInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFundingRateInfoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateInfo(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFundingRateInfoResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFundingRateInfoResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetFundingRateInfo Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateInfo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService IndexPriceAndMarkPrice Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_PERP","pair":"BTCUSD","markPrice":"11029.69574559","indexPrice":"10979.14437500","estimatedSettlePrice":"10981.74168236","lastFundingRate":"0.00071003","interestRate":"0.00010000","nextFundingTime":1596096000000,"time":1596094042000},{"symbol":"BTCUSD_200925","pair":"BTCUSD","markPrice":"12077.01343750","indexPrice":"10979.10312500","estimatedSettlePrice":"10981.74168236","lastFundingRate":"","interestRate":"","nextFundingTime":0,"time":1596094042000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/premiumIndex", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.IndexPriceAndMarkPriceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceAndMarkPrice(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.IndexPriceAndMarkPriceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.IndexPriceAndMarkPriceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService IndexPriceAndMarkPrice Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceAndMarkPrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService IndexPriceKlineCandlestickData Success", func(t *testing.T) {

		mockedJSON := `[[1591256400000,"9653.69440000","9653.69640000","9651.38600000","9651.55200000","0",1591256459999,"0",60,"0","0","0"]]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/indexPriceKlines", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.IndexPriceKlineCandlestickDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Pair("pair_example").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.IndexPriceKlineCandlestickDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.IndexPriceKlineCandlestickDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService IndexPriceKlineCandlestickData Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService IndexPriceKlineCandlestickData Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService KlineCandlestickData Success", func(t *testing.T) {

		mockedJSON := `[[1591258320000,"9640.7","9642.4","9640.6","9642.0","206",1591258379999,"2.13660389",48,"119","1.23424865","0"]]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/klines", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.KlineCandlestickDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Symbol("symbol_example").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.KlineCandlestickDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.KlineCandlestickDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService KlineCandlestickData Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService KlineCandlestickData Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService LongShortRatio Success", func(t *testing.T) {

		mockedJSON := `[{"pair":"BTCUSD","longShortRatio":"0.1960","longAccount":"0.6622","shortAccount":"0.3378","timestamp":1583139600000},{"pair":"BTCUSD","longShortRatio":"1.9559","longAccount":"0.6617","shortAccount":"0.3382","timestamp":1583139900000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/globalLongShortAccountRatio", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.LongShortRatioResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Pair("pair_example").Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.LongShortRatioResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.LongShortRatioResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService LongShortRatio Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService LongShortRatio Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarkPriceKlineCandlestickData Success", func(t *testing.T) {

		mockedJSON := `[[1591256460000,"9653.29201333","9654.56401333","9653.07367333","9653.07367333","0",1591256519999,"0",60,"0","0","0"]]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/markPriceKlines", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarkPriceKlineCandlestickDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Symbol("symbol_example").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarkPriceKlineCandlestickDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarkPriceKlineCandlestickDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService MarkPriceKlineCandlestickData Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarkPriceKlineCandlestickData Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OldTradesLookup Success", func(t *testing.T) {

		mockedJSON := `[{"id":595103,"price":"9642.2","qty":"1","baseQty":"0.01037108","time":1499865549590,"isBuyerMaker":true}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/historicalTrades", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OldTradesLookupResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OldTradesLookupResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OldTradesLookupResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService OldTradesLookup Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OldTradesLookup Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterest Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSD_200626","pair":"BTCUSD","openInterest":"15004","contractType":"CURRENT_QUARTER","time":1591261042378}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/openInterest", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OpenInterestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OpenInterestResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OpenInterestResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService OpenInterest Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterest Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterestStatistics Success", func(t *testing.T) {

		mockedJSON := `[{"pair":"BTCUSD","contractType":"CURRENT_QUARTER","sumOpenInterest":"20403","sumOpenInterestValue":"176196512.23400000","timestamp":1591261042378},{"pair":"BTCUSD","contractType":"CURRENT_QUARTER","sumOpenInterest":"20401","sumOpenInterestValue":"176178704.98700000","timestamp":1583128200000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/openInterestHist", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisContractTypeParameterPerpetual), r.URL.Query().Get("contractType"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OpenInterestStatisticsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Pair("pair_example").ContractType(models.BasisContractTypeParameterPerpetual).Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OpenInterestStatisticsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OpenInterestStatisticsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService OpenInterestStatistics Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterestStatistics Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OrderBook Success", func(t *testing.T) {

		mockedJSON := `{"lastUpdateId":16769853,"symbol":"BTCUSD_PERP","pair":"BTCUSD","E":1591250106370,"T":1591250106368,"bids":[["9638.0","431"]],"asks":[["9638.2","12"]]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/depth", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderBookResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderBookResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderBookResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService OrderBook Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OrderBook Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService PremiumIndexKlineData Success", func(t *testing.T) {

		mockedJSON := `[[1691603820000,"-0.00042931","-0.00023641","-0.00059406","-0.00043659","0",1691603879999,"0",12,"0","0","0"]]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/premiumIndexKlines", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m), r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.PremiumIndexKlineDataResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Symbol("symbol_example").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.PremiumIndexKlineDataResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.PremiumIndexKlineDataResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService PremiumIndexKlineData Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService PremiumIndexKlineData Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryIndexPriceConstituents Success", func(t *testing.T) {

		mockedJSON := `{"symbol":"BTCUSD","time":1697422647853,"constituents":[{"exchange":"bitstamp","symbol":"btcusd"},{"exchange":"coinbase","symbol":"BTC-USD"},{"exchange":"kraken","symbol":"XBT/USD"},{"exchange":"binance_cross","symbol":"BTCUSDC*index(USDCUSD)"}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/constituents", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryIndexPriceConstituentsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryIndexPriceConstituentsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryIndexPriceConstituentsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService QueryIndexPriceConstituents Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryIndexPriceConstituents Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService RecentTradesList Success", func(t *testing.T) {

		mockedJSON := `[{"id":28457,"price":"9635.0","qty":"1","baseQty":"0.01037883","time":1591250192508,"isBuyerMaker":true}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/trades", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RecentTradesListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Symbol("symbol_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RecentTradesListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RecentTradesListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService RecentTradesList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService RecentTradesList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService SymbolOrderBookTicker Success", func(t *testing.T) {

		mockedJSON := `[{"lastUpdateId":1027024,"symbol":"BTCUSD_200626","pair":"BTCUSD","bidPrice":"9650.1","bidQty":"16","askPrice":"9650.3","askQty":"7","time":1591257300345}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/ticker/bookTicker", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolOrderBookTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolOrderBookTicker(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SymbolOrderBookTickerResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SymbolOrderBookTickerResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService SymbolOrderBookTicker Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolOrderBookTicker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService SymbolPriceTicker Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_200626","ps":"9647.8","price":"9647.8","time":1591257246176}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/ticker/price", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolPriceTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTicker(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SymbolPriceTickerResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SymbolPriceTickerResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService SymbolPriceTicker Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTicker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TakerBuySellVolume Success", func(t *testing.T) {

		mockedJSON := `[{"pair":"BTCUSD","contractType":"CURRENT_QUARTER","takerBuyVol":"387","takerSellVol":"248","takerBuyVolValue":"2342.1220","takerSellVolValue":"4213.9800","timestamp":1591261042378},{"pair":"BTCUSD","contractType":"CURRENT_QUARTER","takerBuyVol":"234","takerSellVol":"121","takerBuyVolValue":"4563.1320","takerSellVolValue":"3313.3940","timestamp":1585615200000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/takerBuySellVol", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisContractTypeParameterPerpetual), r.URL.Query().Get("contractType"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TakerBuySellVolumeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Pair("pair_example").ContractType(models.BasisContractTypeParameterPerpetual).Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TakerBuySellVolumeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TakerBuySellVolumeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService TakerBuySellVolume Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TakerBuySellVolume Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TestConnectivity Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/ping", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()
		require.NoError(t, err)
	})

	t.Run("Test MarketDataAPIService TestConnectivity Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test MarketDataAPIService Ticker24hrPriceChangeStatistics Success", func(t *testing.T) {

		mockedJSON := `[{"symbol":"BTCUSD_200925","pair":"BTCUSD","priceChange":"136.6","priceChangePercent":"1.436","weightedAvgPrice":"9547.3","lastPrice":"9651.6","lastQty":"1","openPrice":"9515.0","highPrice":"9687.0","lowPrice":"9499.5","volume":"494109","baseVolume":"5192.94797687","openTime":1591170300000,"closeTime":1591256718418,"firstId":600507,"lastId":697803,"count":97297}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/dapi/v1/ticker/24hr", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.Ticker24hrPriceChangeStatisticsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker24hrPriceChangeStatistics(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.Ticker24hrPriceChangeStatisticsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.Ticker24hrPriceChangeStatisticsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService Ticker24hrPriceChangeStatistics Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker24hrPriceChangeStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioAccounts Success", func(t *testing.T) {

		mockedJSON := `[{"pair":"BTCUSD","longShortRatio":"1.8105","longAccount":"0.6442","shortAccount":"0.3558","timestamp":1591261042378},{"pair":"BTCUSD","longShortRatio":"1.1110","longAccount":"0.5263","shortAccount":"0.4737","timestamp":1592870400000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/topLongShortAccountRatio", r.URL.Path)
			require.Equal(t, "symbol_example", r.URL.Query().Get("symbol"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TopTraderLongShortRatioAccountsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Symbol("symbol_example").Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TopTraderLongShortRatioAccountsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TopTraderLongShortRatioAccountsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioAccounts Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioAccounts Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioPositions Success", func(t *testing.T) {

		mockedJSON := `[{"pair":"BTCUSD","longShortRatio":"0.7869","longPosition":"0.6442","shortPosition":"0.4404","timestamp":1592870400000},{"pair":"BTCUSD","longShortRatio":"1.1231","longPosition":"0.2363","shortPosition":"0.4537","timestamp":1592956800000}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/topLongShortPositionRatio", r.URL.Path)
			require.Equal(t, "pair_example", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.BasisPeriodParameterPeriod5m), r.URL.Query().Get("period"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TopTraderLongShortRatioPositionsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Pair("pair_example").Period(models.BasisPeriodParameterPeriod5m).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TopTraderLongShortRatioPositionsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TopTraderLongShortRatioPositionsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioPositions Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioPositions Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingCoinFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
