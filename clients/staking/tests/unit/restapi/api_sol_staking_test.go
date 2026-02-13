/*
Binance Staking REST API TEST

Testing SolStakingAPIService

*/

package binancestakingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/staking"
	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancestakingrestapi_SolStakingAPIService(t *testing.T) {

	t.Run("Test SolStakingAPIService ClaimBoostRewards Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/claim", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.ClaimBoostRewardsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.ClaimBoostRewards(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.ClaimBoostRewardsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.ClaimBoostRewardsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService ClaimBoostRewards Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.ClaimBoostRewards(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetBnsolRateHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"annualPercentageRate":"0.00006408","exchangeRate":"1.001212343432","boostRewards":[{"boostAPR":"0.12000000","rewardsAsset":"SOL"},{"boostAPR":"0.00200000","rewardsAsset":"BNB"}],"time":1577233578000}],"total":"1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/rateHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetBnsolRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRateHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetBnsolRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetBnsolRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetBnsolRateHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetBnsolRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"estRewardsInSOL":"1.23230920","rows":[{"time":1575018510000,"amountInSOL":"0.23223","holding":"2.3223","holdingInSOL":"2.4231","annualPercentageRate":"0.5"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/bnsolRewardsHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetBnsolRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRewardsHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetBnsolRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetBnsolRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetBnsolRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBnsolRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetBoostRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1729520680,"token":"SOL","amount":"1.20291028","bnsolHolding":"2.0928798","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/boostRewardsHistory", r.URL.Path)
			require.Equal(t, "CLAIM", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetBoostRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBoostRewardsHistory(context.Background()).Type("CLAIM").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetBoostRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetBoostRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetBoostRewardsHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBoostRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetBoostRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetBoostRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetSolRedemptionHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1575018510000,"arrivalTime":1575018510000,"asset":"BNSOL","amount":"21312.23223","distributeAsset":"SOL","distributeAmount":"21338.0699","exchangeRate":"1.00121234","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/redemptionHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSolRedemptionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolRedemptionHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSolRedemptionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSolRedemptionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetSolRedemptionHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolRedemptionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetSolStakingHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"time":1575018510000,"asset":"SOL","amount":"21312.23223","distributeAsset":"BNSOL","distributeAmount":"21286.42584","exchangeRate":"1.00121234","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/stakingHistory", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSolStakingHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSolStakingHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSolStakingHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetSolStakingHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetSolStakingQuotaDetails Success", func(t *testing.T) {

		mockedJSON := `{"leftStakingPersonalQuota":"1000","leftRedemptionPersonalQuota":"1000","minStakeAmount":"0.01000000","minRedeemAmount":"0.00000001","redeemPeriod":4,"stakeable":true,"redeemable":true,"soldOut":false,"commissionFee":"0.25000000","nextEpochTime":725993969475,"calculating":false}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/quota", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSolStakingQuotaDetailsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingQuotaDetails(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSolStakingQuotaDetailsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSolStakingQuotaDetailsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetSolStakingQuotaDetails Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetSolStakingQuotaDetails(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService GetUnclaimedRewards Success", func(t *testing.T) {

		mockedJSON := `[{"amount":"1.00000011","rewardsAsset":"SOL"},{"amount":"2.00202321","rewardsAsset":"BNB"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/history/unclaimedRewards", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetUnclaimedRewardsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetUnclaimedRewards(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetUnclaimedRewardsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetUnclaimedRewardsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService GetUnclaimedRewards Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.GetUnclaimedRewards(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService RedeemSol Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"solAmount":"0.23092091","exchangeRate":"1.00121234","arrivalTime":1575018510000}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/redeem", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RedeemSolResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.RedeemSol(context.Background()).Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RedeemSolResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RedeemSolResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService RedeemSol Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.RedeemSol(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService RedeemSol Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.RedeemSol(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService SolStakingAccount Success", func(t *testing.T) {

		mockedJSON := `{"bnsolAmount":"1.10928781","holdingInSOL":"1.22330928","thirtyDaysProfitInSOL":"0.22330928"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SolStakingAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.SolStakingAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SolStakingAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SolStakingAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService SolStakingAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.SolStakingAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService SubscribeSolStaking Success", func(t *testing.T) {

		mockedJSON := `{"success":true,"bnsolAmount":"0.23092091","exchangeRate":"1.001212342342"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/sol-staking/sol/stake", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubscribeSolStakingResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.SubscribeSolStaking(context.Background()).Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubscribeSolStakingResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubscribeSolStakingResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test SolStakingAPIService SubscribeSolStaking Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.SubscribeSolStaking(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test SolStakingAPIService SubscribeSolStaking Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.SolStakingAPI.SubscribeSolStaking(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
