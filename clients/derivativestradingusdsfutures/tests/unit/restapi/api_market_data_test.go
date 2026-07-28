/*
Futures (USDⓈ-M) REST API TEST

Testing MarketDataAPIService

*/

package binancederivativestradingusdsfuturesrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures"
	"github.com/binance/binance-connector-go/clients/derivativestradingusdsfutures/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancederivativestradingusdsfuturesrestapi_MarketDataAPIService(t *testing.T) {

	t.Run("Test MarketDataAPIService AdlRisk Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","adlRisk":"low","updateTime":1597370495002}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/symbolAdlRisk", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AdlRiskResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AdlRisk(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AdlRiskResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AdlRiskResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService AdlRisk Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AdlRisk(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService AssetIndex Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"ADAUSD","time":1635740268004,"index":"1.92957370","bidBuffer":"0.10000000","askBuffer":"0.10000000","bidRate":"1.73661633","askRate":"2.12253107","autoExchangeBidBuffer":"0.05000000","autoExchangeAskBuffer":"0.05000000","autoExchangeBidRate":"1.83309501","autoExchangeAskRate":"2.02605238"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/assetIndex", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AssetIndexResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AssetIndex(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AssetIndexResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AssetIndexResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService AssetIndex Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.AssetIndex(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService Basis Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"indexPrice":"34400.15945055","contractType":"PERPETUAL","basisRate":"0.0004","futuresPrice":"34414.10","annualizedBasisRate":"","basis":"13.94054945","pair":"BTCUSDT","timestamp":1698742800000}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/basis", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("pair"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Pair("BTCUSDT").ContractType(models.BasisContractTypeParameterPerpetual).Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Basis(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CheckServerTime Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"serverTime":1499827319559}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/time", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckServerTimeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CheckServerTime(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CompositeIndexSymbolInformation Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"DEFIUSDT","time":1589437530011,"component":"baseAsset","baseAssetList":[]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/indexInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CompositeIndexSymbolInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompositeIndexSymbolInformation(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CompositeIndexSymbolInformationResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CompositeIndexSymbolInformationResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService CompositeIndexSymbolInformation Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompositeIndexSymbolInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService CompressedAggregateTradesList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"a":26129,"p":"0.01633102","q":"4.70443515","nq":"100","f":27781,"l":27781,"T":1498793709153,"m":true}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/aggTrades", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CompressedAggregateTradesListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.CompressedAggregateTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ContinuousContractKlineCandlestickData Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[[1607444700000]]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/continuousKlines", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("pair"))
			require.Equal(t, string(models.ContinuousContractKlineCandlestickDataContractTypeParameterPerpetual), r.URL.Query().Get("contractType"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Pair("BTCUSDT").ContractType(models.ContinuousContractKlineCandlestickDataContractTypeParameterPerpetual).Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ContinuousContractKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService ExchangeInformation Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"exchangeFilters":[""],"rateLimits":[{"interval":"MINUTE","intervalNum":1,"limit":2400,"rateLimitType":"REQUEST_WEIGHT"}],"serverTime":1565613908500,"assets":[{"asset":"BTC","marginAvailable":true,"autoAssetExchange":"-0.10"}],"symbols":[{"symbol":"BLZUSDT","pair":"BLZUSDT","contractType":"PERPETUAL","deliveryDate":4133404800000,"onboardDate":1598252400000,"status":"TRADING","maintMarginPercent":"2.5000","requiredMarginPercent":"5.0000","baseAsset":"BLZ","quoteAsset":"USDT","marginAsset":"USDT","pricePrecision":5,"quantityPrecision":0,"baseAssetPrecision":8,"quotePrecision":8,"underlyingType":"COIN","underlyingSubType":["STORAGE"],"settlePlan":0,"triggerProtect":"0.15","filters":[{"filterType":"PRICE_FILTER","maxPrice":"300","minPrice":"0.0001","tickSize":"0.0001","maxQty":"10000000","minQty":"1","stepSize":"1","limit":200,"notional":"5.0","multiplierUp":"1.1500","multiplierDown":"0.8500","multiplierDecimal":"4"}],"orderTypes":["LIMIT"],"timeInForce":["GTC"],"liquidationFee":"0.010000","marketTakeBound":"0.30"}],"timezone":"UTC"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/exchangeInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ExchangeInformationResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.ExchangeInformation(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetFundingRateHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","fundingRate":"-0.03750000","fundingTime":1570608000000,"markPrice":"34287.54619963","rateType":"Regular"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/fundingRate", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFundingRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFundingRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFundingRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService GetFundingRateHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService GetFundingRateInfo Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BLZUSDT","adjustedFundingRateCap":"0.02500000","adjustedFundingRateFloor":"-0.02500000","fundingIntervalHours":8,"disclaimer":false}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/fundingInfo", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFundingRateInfoResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.GetFundingRateInfo(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService IndexPriceKlineCandlestickData Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[[1591256400000]]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/indexPriceKlines", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("pair"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Pair("BTCUSDT").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.IndexPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService KlineCandlestickData Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[[1499040000000]]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/klines", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Symbol("BTCUSDT").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.KlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService LongShortRatio Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","longShortRatio":"0.1960","longAccount":"0.6622","shortAccount":"0.3378","timestamp":1591261042378}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/globalLongShortAccountRatio", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Symbol("BTCUSDT").Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.LongShortRatio(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarkPrice Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","markPrice":"11793.63104562","indexPrice":"11781.80495970","estimatedSettlePrice":"11781.16138815","lastFundingRate":"0.00038246","interestRate":"0.00010000","nextFundingTime":1597392000000,"time":1597370495002}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/premiumIndex", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MarkPriceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPrice(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MarkPriceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MarkPriceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService MarkPrice Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService MarkPriceKlineCandlestickData Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[[1591256460000]]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/markPriceKlines", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Symbol("BTCUSDT").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.MarkPriceKlineCandlestickData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OldTradesLookup Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"id":28457,"price":"4.00000100","qty":"12.00000000","quoteQty":"8000.00","time":1499865549590,"isBuyerMaker":true,"isRPITrade":true}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/historicalTrades", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OldTradesLookupResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OldTradesLookup(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterest Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"openInterest":"10659.509","symbol":"BTCUSDT","time":1589437530011}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/openInterest", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OpenInterestResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterest(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OpenInterestStatistics Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","sumOpenInterest":"20403.12345678","sumOpenInterestValue":"176196512.12345678","CMCCirculatingSupply":"165880.538","timestamp":1591261042378}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/openInterestHist", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Symbol("BTCUSDT").Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OpenInterestStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService OrderBook Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"lastUpdateId":1027024,"E":1589436922972,"T":1589436922959,"bids":[["4.00000000","431.00000000"]],"asks":[["4.00000200","12.00000000"]]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/depth", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderBookResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.OrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService PremiumIndexKlineData Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[[1691603820000]]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/premiumIndexKlines", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Symbol("BTCUSDT").Interval(models.ContinuousContractKlineCandlestickDataIntervalParameterInterval1m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.PremiumIndexKlineData(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QuarterlyContractSettlementPrice Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"deliveryTime":1695945600000,"deliveryPrice":27103}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/delivery-price", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("pair"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuarterlyContractSettlementPriceResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QuarterlyContractSettlementPrice(context.Background()).Pair("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuarterlyContractSettlementPriceResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuarterlyContractSettlementPriceResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService QuarterlyContractSettlementPrice Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QuarterlyContractSettlementPrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QuarterlyContractSettlementPrice Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QuarterlyContractSettlementPrice(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryIndexPriceConstituents Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","time":1745401553408,"constituents":[]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/constituents", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryIndexPriceConstituentsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryIndexPriceConstituents(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService QueryInsuranceFundBalanceSnapshot Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbols":[],"assets":[]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/insuranceBalance", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryInsuranceFundBalanceSnapshotResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryInsuranceFundBalanceSnapshot(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryInsuranceFundBalanceSnapshotResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryInsuranceFundBalanceSnapshotResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService QueryInsuranceFundBalanceSnapshot Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.QueryInsuranceFundBalanceSnapshot(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService RecentTradesList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"id":28457,"price":"4.00000100","qty":"12.00000000","quoteQty":"48.00","time":1499865549590,"isBuyerMaker":true,"isRPITrade":true}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/trades", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RecentTradesListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Symbol("BTCUSDT").Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RecentTradesList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService RpiOrderBook Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"lastUpdateId":1027024,"E":1589436922972,"T":1589436922959,"bids":[["4.00000000","431.00000000"]],"asks":[["4.00000200","12.00000000"]]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/rpiDepth", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RpiOrderBookResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RpiOrderBook(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RpiOrderBookResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RpiOrderBookResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService RpiOrderBook Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RpiOrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService RpiOrderBook Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.RpiOrderBook(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService SymbolOrderBookTicker Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","bidPrice":"4.00000000","bidQty":"431.00000000","askPrice":"4.00000200","askQty":"9.00000000","time":1589437530011}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/ticker/bookTicker", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolOrderBookTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolOrderBookTicker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService SymbolPriceTicker Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","price":"6000.01","time":1589437530011}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/ticker/price", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolPriceTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTicker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService SymbolPriceTickerV2 Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","price":"6000.01","time":1589437530011}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v2/ticker/price", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SymbolPriceTickerV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTickerV2(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SymbolPriceTickerV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SymbolPriceTickerV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService SymbolPriceTickerV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.SymbolPriceTickerV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TakerBuySellVolume Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"buySellRatio":"1.5586","buyVol":"387.3300","sellVol":"248.5030","timestamp":1591261042378}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/takerlongshortRatio", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Symbol("BTCUSDT").Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TakerBuySellVolume(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TestConnectivity Success", func(t *testing.T) {

		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/ping", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		_, err := apiClient.RestApi.MarketDataAPI.TestConnectivity(context.Background()).Execute()

		require.Error(t, err)
	})

	t.Run("Test MarketDataAPIService Ticker24hrPriceChangeStatistics Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","priceChange":"-94.99999800","priceChangePercent":"-95.960","weightedAvgPrice":"0.29628482","lastPrice":"4.00000200","lastQty":"200.00000000","openPrice":"99.00000000","highPrice":"100.00000000","lowPrice":"0.10000000","volume":"8913.30000000","quoteVolume":"15.30000000","openTime":1499783499040,"closeTime":1499869899040,"firstId":28385,"lastId":28460,"count":76}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/ticker/24hr", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.Ticker24hrPriceChangeStatisticsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.Ticker24hrPriceChangeStatistics(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioAccounts Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","longShortRatio":"1.8105","longAccount":"0.6442","shortAccount":"0.3558","timestamp":1591261042378}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/topLongShortAccountRatio", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Symbol("BTCUSDT").Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioAccounts(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TopTraderLongShortRatioPositions Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","longShortRatio":"1.4342","longAccount":"0.5344","shortAccount":"0.4238","timestamp":1591261042378}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/futures/data/topLongShortPositionRatio", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Symbol("BTCUSDT").Period(models.BasisPeriodParameterPeriod5m).Execute()
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
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

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TopTraderLongShortRatioPositions(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test MarketDataAPIService TradingSchedule Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"updateTime":1761286643918,"marketSchedules":{}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/fapi/v1/tradingSchedule", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.TradingScheduleResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TradingSchedule(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.TradingScheduleResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.TradingScheduleResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test MarketDataAPIService TradingSchedule Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceDerivativesTradingUsdsFuturesClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.MarketDataAPI.TradingSchedule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
