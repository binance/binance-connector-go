/*
Spot REST API TEST

Testing AccountAPIService

*/

package binancespotrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/spot"
	"github.com/binance/binance-connector-go/clients/spot/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancespotrestapi_AccountAPIService(t *testing.T) {

	t.Run("Test AccountAPIService AccountCommission Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"BTCUSDT","standardCommission":{"maker":"0.00000010","taker":"0.00000020","buyer":"0.00000030","seller":"0.00000040"},"specialCommission":{"maker":"0.01000000","taker":"0.02000000","buyer":"0.03000000","seller":"0.04000000"},"taxCommission":{"maker":"0.00000112","taker":"0.00000114","buyer":"0.00000118","seller":"0.00000116"},"discount":{"enabledForAccount":true,"enabledForSymbol":true,"discountAsset":"BNB","discount":"0.75000000"}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/account/commission", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AccountCommissionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountCommission(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AccountCommissionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AccountCommissionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AccountCommission Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountCommission(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AccountCommission Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AccountCommission(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AllOrderList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"orderListId":29,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"amEEAXryFzFwYF1FeRpUoZ","transactionTime":1565245913483,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"oD7aesZqjEGlZrbtRpy5zB"}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/allOrderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AllOrderListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AllOrderList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AllOrderListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AllOrderListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AllOrderList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AllOrderList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AllOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"LTCBTC","orderId":1,"orderListId":-1,"clientOrderId":"myOrder1","price":"0.1","origQty":"1.0","executedQty":"0.0","cummulativeQuoteQty":"0.0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1499827319559,"updateTime":1499827319559,"isWorking":true,"origQuoteOrderQty":"0.000000","workingTime":1499827319559,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/allOrders", r.URL.Path)
			require.Equal(t, "LTCBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.AllOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AllOrders(context.Background()).Symbol("LTCBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.AllOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.AllOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService AllOrders Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AllOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService AllOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.AllOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetAccount Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"makerCommission":15,"takerCommission":15,"buyerCommission":0,"sellerCommission":0,"commissionRates":{"maker":"0.00150000","taker":"0.00150000","buyer":"0.00000000","seller":"0.00000000"},"canTrade":true,"canWithdraw":true,"canDeposit":true,"brokered":false,"requireSelfTradePrevention":false,"preventSor":false,"updateTime":123456789,"accountType":"SPOT","balances":[{"asset":"BTC","free":"4723846.89208129","locked":"0.00000000"}],"permissions":["SPOT"],"uid":354937868}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetOpenOrders Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"LTCBTC","orderId":1,"orderListId":-1,"clientOrderId":"myOrder1","price":"0.1","origQty":"1.0","executedQty":"0.0","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1499827319559,"updateTime":1499827319559,"isWorking":true,"workingTime":1499827319559,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/openOrders", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOpenOrdersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOpenOrders(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOpenOrdersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOpenOrdersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetOpenOrders Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOpenOrders(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"symbol":"LTCBTC","orderId":1,"orderListId":-1,"clientOrderId":"myOrder1","price":"0.1","origQty":"1.0","executedQty":"0.0","origQuoteOrderQty":"0.000000","cummulativeQuoteQty":"0.0","status":"NEW","timeInForce":"GTC","type":"LIMIT","side":"BUY","stopPrice":"0.00000000","icebergQty":"0.00000000","time":1499827319559,"updateTime":1499827319559,"isWorking":true,"workingTime":1499827319559,"selfTradePreventionMode":"NONE","preventedMatchId":0,"preventedQuantity":"1.200000","strategyId":1,"strategyType":1000000,"trailingDelta":10,"trailingTime":-1,"usedSor":true,"workingFloor":"SOR","pegPriceType":"PRIMARY_PEG","pegOffsetType":"PRICE_LEVEL","pegOffsetValue":5,"peggedPrice":"87523.83710000","expiryReason":"INSUFFICIENT_LIQUIDITY"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order", r.URL.Path)
			require.Equal(t, "LTCBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOrder(context.Background()).Symbol("LTCBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetOrder Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService GetOrderList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"orderListId":27,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"h2USkA5YQpaXHPIrkd96xE","transactionTime":1565245656253,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"qD1gy3kc3Gx0rihm9Y3xwS"}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/orderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOrderListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOrderList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOrderListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOrderListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService GetOrderList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.GetOrderList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyAllocations Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","allocationId":0,"allocationType":"SOR","orderId":1,"orderListId":-1,"price":"1.00000000","qty":"5.00000000","quoteQty":"5.00000000","commission":"0.00000000","commissionAsset":"BTC","time":1687506878118,"isBuyer":true,"isMaker":false,"isAllocator":false}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/myAllocations", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MyAllocationsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyAllocations(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MyAllocationsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MyAllocationsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService MyAllocations Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyAllocations(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyAllocations Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyAllocations(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyFilters Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"exchangeFilters":[{"filterType":"EXCHANGE_MAX_NUM_ORDERS","maxNumOrders":1000}],"symbolFilters":[{"filterType":"PRICE_FILTER","priceExponent":8,"minPrice":"0.00000100","maxPrice":"100000.00000000","tickSize":"0.00000100"}],"assetFilters":[{"filterType":"MAX_ASSET","qtyExponent":8,"limit":"100.00000000","asset":"BNB"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":321}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/myFilters", r.URL.Path)
			require.Equal(t, "BNBUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MyFiltersResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyFilters(context.Background()).Symbol("BNBUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MyFiltersResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MyFiltersResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService MyFilters Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyFilters(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyFilters Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyFilters(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyPreventedMatches Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","preventedMatchId":1,"takerOrderId":5,"makerSymbol":"BTCUSDT","makerOrderId":3,"tradeGroupId":1,"selfTradePreventionMode":"EXPIRE_MAKER","price":"1.100000","makerPreventedQuantity":"1.300000","transactTime":1669101687094}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/myPreventedMatches", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MyPreventedMatchesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyPreventedMatches(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MyPreventedMatchesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MyPreventedMatchesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService MyPreventedMatches Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyPreventedMatches(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyPreventedMatches Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyPreventedMatches(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyTrades Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BNBBTC","id":28457,"orderId":100234,"orderListId":-1,"price":"4.00000100","qty":"12.00000000","quoteQty":"48.000012","commission":"10.10000000","commissionAsset":"BNB","time":1499865549590,"isBuyer":true,"isMaker":false,"isBestMatch":true}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/myTrades", r.URL.Path)
			require.Equal(t, "BNBBTC", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.MyTradesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyTrades(context.Background()).Symbol("BNBBTC").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.MyTradesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.MyTradesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService MyTrades Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyTrades(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService MyTrades Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.MyTrades(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService OpenOrderList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"orderListId":31,"contingencyType":"OCO","listStatusType":"EXEC_STARTED","listOrderStatus":"EXECUTING","listClientOrderId":"wuB13fmulKj3YjdqWEcsnp","transactionTime":1565246080644,"symbol":"LTCBTC","orders":[{"symbol":"LTCBTC","orderId":4,"clientOrderId":"r3EH2N76dHfLoSZWIUw1bT"}]}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/openOrderList", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OpenOrderListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OpenOrderList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OpenOrderListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OpenOrderListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService OpenOrderList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OpenOrderList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService OrderAmendments Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"symbol":"BTCUSDT","orderId":9,"executionId":22,"origClientOrderId":"W0fJ9fiLKHOJutovPK3oJp","newClientOrderId":"UQ1Np3bmQ71jJzsSDW9Vpi","origQty":"5.00000000","newQty":"4.00000000","time":1741669661670}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/order/amendments", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "9", r.URL.Query().Get("orderId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OrderAmendmentsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OrderAmendments(context.Background()).Symbol("BTCUSDT").OrderId(int64(9)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OrderAmendmentsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OrderAmendmentsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService OrderAmendments Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OrderAmendments(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService OrderAmendments Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.OrderAmendments(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AccountAPIService RateLimitOrder Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/api/v3/rateLimit/order", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RateLimitOrderResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.RateLimitOrder(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RateLimitOrderResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RateLimitOrderResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AccountAPIService RateLimitOrder Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSpotClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AccountAPI.RateLimitOrder(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
