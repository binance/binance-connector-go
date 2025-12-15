/*
Binance Staking REST API TEST

Testing OnChainYieldsAPIService

*/

package binancestakingrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/staking/src"
	"github.com/binance/binance-connector-go/clients/staking/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancestakingrestapi_OnChainYieldsAPIService(t *testing.T) {

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedPersonalLeftQuota Success", func(t *testing.T) {

		mockedJSON := `{"leftPersonalQuota":"1000"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/personalLeftQuota", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedPersonalLeftQuotaResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedPersonalLeftQuota(context.Background()).ProjectId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedPersonalLeftQuotaResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedPersonalLeftQuotaResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedPersonalLeftQuota Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedPersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedPersonalLeftQuota Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedPersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedProductList Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"projectId":"Solv-60d","detail":{"asset":"BTC","rewardAsset":"SOLV","duration":60,"renewable":true,"isSoldOut":true,"apr":"0.039","status":"PREHEATING","subscriptionStartTime":1646182276000,"canRedeemToFlex":true},"quota":{"totalPersonalQuota":"2","minimum":"0.001"}}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedProductListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedProductListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedProductListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedProductList Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedProductPosition Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":"123123","projectId":"Solv-60d","asset":"BTC","amount":"122.09202928","purchaseTime":"1646182276000","duration":"60","accrualDays":"4","rewardAsset":"SOLV","APY":"0.039","rewardAmt":"5.17181528","nextPay":"1.29295383","nextPayDate":"1646697600000","payPeriod":"1","rewardsPayDate":"1646697600000","rewardsEndDate":"1651449600000","deliverDate":"1651536000000","nextSubscriptionDate":"1651536000000","redeemingAmt":"232.2323","redeemTo":"FLEXIBLE","canRedeemEarly":true,"autoSubscribe":true,"type":"AUTO","status":"HOLDING"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/position", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedProductPositionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductPosition(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedProductPositionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedProductPositionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedProductPosition Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedProductPosition(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedRedemptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":"123123","redeemId":40607,"time":1575018510000,"asset":"BTC","lockPeriod":"30","amount":"21312.23223","originalAmount":"21312.23223","type":"NORMAL","deliverDate":"1575018510000","lossAmount":"0.00001232","isComplete":true,"rewardAsset":"SOLV","rewardAmt":"5.17181528","status":"PAID"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/history/redemptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedRedemptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRedemptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedRedemptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedRedemptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedRedemptionRecord Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRedemptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":"123123","time":1575018510000,"asset":"BNB","lockPeriod":"30","amount":"21312.23223"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/history/rewardsRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRewardsHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedRewardsHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedSubscriptionPreview Success", func(t *testing.T) {

		mockedJSON := `{"rewardAsset":"SOLV","totalRewardAmt":"5.17181528","nextPay":"1.29295383","nextPayDate":"1646697600000","rewardsPayDate":"1646697600000","valueDate":"1646697600000","rewardsEndDate":"1651449600000","deliverDate":"1651536000000","nextSubscriptionDate":"1651536000000"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/subscriptionPreview", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedSubscriptionPreviewResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionPreview(context.Background()).ProjectId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionPreviewResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedSubscriptionPreviewResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedSubscriptionPreview Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedSubscriptionPreview Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedSubscriptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":"123123","purchaseId":"26055","projectId":"Solv-60d","clientId":"ABC","time":1575018510000,"asset":"BTC","amount":"21312.23223","lockPeriod":"30","type":"AUTO","sourceAccount":"SPOT","amtFromSpot":"30","amtFromFunding":"70","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/history/subscriptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetOnChainYieldsLockedSubscriptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetOnChainYieldsLockedSubscriptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetOnChainYieldsLockedSubscriptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService GetOnChainYieldsLockedSubscriptionRecord Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.GetOnChainYieldsLockedSubscriptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService OnChainYieldsAccount Success", func(t *testing.T) {

		mockedJSON := `{"totalAmountInBTC":"0.01067982","totalAmountInUSDT":"77.13289230","totalFlexibleAmountInBTC":"0.00000000","totalFlexibleAmountInUSDT":"0.00000000","totalLockedInBTC":"0.01067982","totalLockedInUSDT":"77.13289230"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.OnChainYieldsAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.OnChainYieldsAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.OnChainYieldsAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.OnChainYieldsAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService OnChainYieldsAccount Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.OnChainYieldsAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService RedeemOnChainYieldsLockedProduct Success", func(t *testing.T) {

		mockedJSON := `{"redeemId":40607,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/redeem", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RedeemOnChainYieldsLockedProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.RedeemOnChainYieldsLockedProduct(context.Background()).PositionId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RedeemOnChainYieldsLockedProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RedeemOnChainYieldsLockedProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService RedeemOnChainYieldsLockedProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.RedeemOnChainYieldsLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService RedeemOnChainYieldsLockedProduct Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.RedeemOnChainYieldsLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedAutoSubscribe Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/setAutoSubscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			require.Equal(t, "true", r.URL.Query().Get("autoSubscribe"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetOnChainYieldsLockedAutoSubscribeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedAutoSubscribe(context.Background()).PositionId("1").AutoSubscribe(true).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetOnChainYieldsLockedAutoSubscribeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetOnChainYieldsLockedAutoSubscribeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedAutoSubscribe Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedAutoSubscribe Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedProductRedeemOption Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/setRedeemOption", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			require.Equal(t, "redeemTo_example", r.URL.Query().Get("redeemTo"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetOnChainYieldsLockedProductRedeemOptionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedProductRedeemOption(context.Background()).PositionId("1").RedeemTo("redeemTo_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetOnChainYieldsLockedProductRedeemOptionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetOnChainYieldsLockedProductRedeemOptionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedProductRedeemOption Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedProductRedeemOption(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SetOnChainYieldsLockedProductRedeemOption Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SetOnChainYieldsLockedProductRedeemOption(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SubscribeOnChainYieldsLockedProduct Success", func(t *testing.T) {

		mockedJSON := `{"purchaseId":40607,"positionId":"12345","amount":"75.46000000","success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/onchain-yields/locked/subscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubscribeOnChainYieldsLockedProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SubscribeOnChainYieldsLockedProduct(context.Background()).ProjectId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubscribeOnChainYieldsLockedProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubscribeOnChainYieldsLockedProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test OnChainYieldsAPIService SubscribeOnChainYieldsLockedProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceStakingClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SubscribeOnChainYieldsLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test OnChainYieldsAPIService SubscribeOnChainYieldsLockedProduct Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.OnChainYieldsAPI.SubscribeOnChainYieldsLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
