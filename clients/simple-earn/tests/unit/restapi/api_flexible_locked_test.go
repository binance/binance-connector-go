/*
Binance Simple Earn REST API TEST

Testing FlexibleLockedAPIService

*/

package binancesimpleearnrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/simpleearn/src"
	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesimpleearnrestapi_FlexibleLockedAPIService(t *testing.T) {

	t.Run("Test FlexibleLockedAPIService GetCollateralRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"amount":"100.00000000","productId":"BUSD001","asset":"USDT","createTime":1575018510000,"type":"REPAY","productName":"USDT","orderId":26055}],"total":"1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/history/collateralRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCollateralRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetCollateralRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCollateralRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCollateralRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetCollateralRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetCollateralRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexiblePersonalLeftQuota Success", func(t *testing.T) {

		mockedJSON := `{"leftPersonalQuota":"1000"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/personalLeftQuota", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexiblePersonalLeftQuotaResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexiblePersonalLeftQuota(context.Background()).ProductId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexiblePersonalLeftQuotaResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexiblePersonalLeftQuotaResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexiblePersonalLeftQuota Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexiblePersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexiblePersonalLeftQuota Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexiblePersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleProductPosition Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"totalAmount":"75.46000000","tierAnnualPercentageRate":{"0-5BTC":0.05,"5-10BTC":0.03},"latestAnnualPercentageRate":"0.02599895","yesterdayAirdropPercentageRate":"0.02599895","asset":"USDT","airDropAsset":"BETH","canRedeem":true,"collateralAmount":"232.23123213","productId":"USDT001","yesterdayRealTimeRewards":"0.10293829","cumulativeBonusRewards":"0.22759183","cumulativeRealTimeRewards":"0.22759183","cumulativeTotalRewards":"0.45459183","autoSubscribe":true}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/position", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleProductPositionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleProductPosition(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleProductPositionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleProductPositionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleProductPosition Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleProductPosition(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleRedemptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"amount":"10.54000000","asset":"USDT","time":1577257222000,"projectId":"USDT001","redeemId":40607,"destAccount":"SPOT","status":"PAID"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/history/redemptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleRedemptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRedemptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleRedemptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleRedemptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleRedemptionRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRedemptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"asset":"BUSD","rewards":"0.00006408","projectId":"USDT001","type":"BONUS","time":1577233578000},{"asset":"USDT","rewards":"0.00687654","projectId":"USDT001","type":"REALTIME","time":1577233562000}],"total":2}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/history/rewardsRecord", r.URL.Path)
			require.Equal(t, "s", r.URL.Query().Get("type"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRewardsHistory(context.Background()).Type("s").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleRewardsHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleSubscriptionPreview Success", func(t *testing.T) {

		mockedJSON := `{"totalAmount":"1232.32230982","rewardAsset":"BUSD","airDropAsset":"BETH","estDailyBonusRewards":"0.22759183","estDailyRealTimeRewards":"0.22759183","estDailyAirdropRewards":"0.22759183"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/subscriptionPreview", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleSubscriptionPreviewResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionPreview(context.Background()).ProductId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleSubscriptionPreviewResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleSubscriptionPreviewResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleSubscriptionPreview Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleSubscriptionPreview Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleSubscriptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"amount":"100.00000000","asset":"USDT","time":1575018510000,"purchaseId":26055,"productId":"USDT001","type":"AUTO","sourceAccount":"SPOT","amtFromSpot":"30","amtFromFunding":"70","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/history/subscriptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetFlexibleSubscriptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetFlexibleSubscriptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetFlexibleSubscriptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetFlexibleSubscriptionRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetFlexibleSubscriptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedPersonalLeftQuota Success", func(t *testing.T) {

		mockedJSON := `{"leftPersonalQuota":"1000"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/personalLeftQuota", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedPersonalLeftQuotaResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedPersonalLeftQuota(context.Background()).ProjectId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedPersonalLeftQuotaResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedPersonalLeftQuotaResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedPersonalLeftQuota Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedPersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedPersonalLeftQuota Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedPersonalLeftQuota(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedProductPosition Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":123123,"parentPositionId":123122,"projectId":"Axs*90","asset":"AXS","amount":"122.09202928","purchaseTime":"1646182276000","duration":"60","accrualDays":"4","rewardAsset":"AXS","APY":"0.2032","rewardAmt":"5.17181528","extraRewardAsset":"BNB","extraRewardAPR":"0.0203","estExtraRewardAmt":"5.17181528","boostRewardAsset":"AXS","boostApr":"0.0121","totalBoostRewardAmt":"3.98201829","nextPay":"1.29295383","nextPayDate":"1646697600000","payPeriod":"1","redeemAmountEarly":"2802.24068892","rewardsEndDate":"1651449600000","deliverDate":"1651536000000","redeemPeriod":"1","redeemingAmt":"232.2323","redeemTo":"FLEXIBLE","partialAmtDeliverDate":"1651536000000","canRedeemEarly":true,"canFastRedemption":true,"autoSubscribe":true,"type":"AUTO","status":"HOLDING","canReStake":true}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/position", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedProductPositionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedProductPosition(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedProductPositionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedProductPositionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedProductPosition Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedProductPosition(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedRedemptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":123123,"redeemId":40607,"time":1575018510000,"asset":"BNB","lockPeriod":"30","amount":"21312.23223","originalAmount":"21312.23223","type":"MATURE","deliverDate":"1575018510000","lossAmount":"0.00001232","isComplete":true,"rewardAsset":"AXS","rewardAmt":"5.17181528","extraRewardAsset":"BNB","estExtraRewardAmt":"5.17181528","status":"PAID"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/history/redemptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedRedemptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRedemptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedRedemptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedRedemptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedRedemptionRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRedemptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedRewardsHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":123123,"time":1575018510000,"asset":"BNB","lockPeriod":"30","amount":"21312.23223","type":"Locked Rewards"},{"positionId":123123,"time":1575018510000,"asset":"BNB","amount":"1.23223","type":"Boost Rewards"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/history/rewardsRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedRewardsHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRewardsHistory(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedRewardsHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedRewardsHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedRewardsHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedRewardsHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedSubscriptionPreview Success", func(t *testing.T) {

		mockedJSON := `[{"rewardAsset":"AXS","totalRewardAmt":"5.17181528","extraRewardAsset":"BNB","estTotalExtraRewardAmt":"5.17181528","boostRewardAsset":"AXS","estDailyRewardAmt":"1.20928901","nextPay":"1.29295383","nextPayDate":"1646697600000","valueDate":"1646697600000","rewardsEndDate":"1651449600000","deliverDate":"1651536000000","nextSubscriptionDate":"1651536000000"}]`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/subscriptionPreview", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedSubscriptionPreviewResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionPreview(context.Background()).ProjectId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedSubscriptionPreviewResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedSubscriptionPreviewResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedSubscriptionPreview Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedSubscriptionPreview Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionPreview(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedSubscriptionRecord Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"positionId":123123,"purchaseId":"26055","projectId":"Axs*90","time":1575018510000,"asset":"BNB","amount":"21312.23223","lockPeriod":"30","type":"AUTO","sourceAccount":"SPOT","amtFromSpot":"30","amtFromFunding":"70","status":"SUCCESS"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/history/subscriptionRecord", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetLockedSubscriptionRecordResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionRecord(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetLockedSubscriptionRecordResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetLockedSubscriptionRecordResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetLockedSubscriptionRecord Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetLockedSubscriptionRecord(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetRateHistory Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"productId":"BUSD001","asset":"BUSD","annualPercentageRate":"0.00006408","time":1577233578000}],"total":"1"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/history/rateHistory", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRateHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetRateHistory(context.Background()).ProductId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRateHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRateHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetRateHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetRateHistory Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetRateHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetSimpleEarnFlexibleProductList Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"asset":"BTC","latestAnnualPercentageRate":"0.05000000","tierAnnualPercentageRate":{"0-5BTC":0.05,"5-10BTC":0.03},"airDropPercentageRate":"0.05000000","canPurchase":true,"canRedeem":true,"isSoldOut":true,"hot":true,"minPurchaseAmount":"0.01000000","productId":"BTC001","subscriptionStartTime":1646182276000,"status":"PURCHASING"}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSimpleEarnFlexibleProductListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnFlexibleProductList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSimpleEarnFlexibleProductListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSimpleEarnFlexibleProductListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetSimpleEarnFlexibleProductList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnFlexibleProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService GetSimpleEarnLockedProductList Success", func(t *testing.T) {

		mockedJSON := `{"rows":[{"projectId":"Axs*90","detail":{"asset":"AXS","rewardAsset":"AXS","duration":90,"renewable":true,"isSoldOut":true,"apr":"1.2069","status":"CREATED","subscriptionStartTime":1646182276000,"extraRewardAsset":"BNB","extraRewardAPR":"0.23","boostRewardAsset":"AXS","boostApr":"0.0121","boostEndTime":1646182276000},"quota":{"totalPersonalQuota":"2","minimum":"0.001"}}],"total":1}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetSimpleEarnLockedProductListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnLockedProductList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetSimpleEarnLockedProductListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetSimpleEarnLockedProductListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService GetSimpleEarnLockedProductList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.GetSimpleEarnLockedProductList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService RedeemFlexibleProduct Success", func(t *testing.T) {

		mockedJSON := `{"redeemId":40607,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/redeem", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RedeemFlexibleProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemFlexibleProduct(context.Background()).ProductId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RedeemFlexibleProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RedeemFlexibleProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService RedeemFlexibleProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemFlexibleProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService RedeemFlexibleProduct Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemFlexibleProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService RedeemLockedProduct Success", func(t *testing.T) {

		mockedJSON := `{"redeemId":40607,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/redeem", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.RedeemLockedProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemLockedProduct(context.Background()).PositionId("1").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.RedeemLockedProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.RedeemLockedProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService RedeemLockedProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService RedeemLockedProduct Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.RedeemLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetFlexibleAutoSubscribe Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/setAutoSubscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			require.Equal(t, "true", r.URL.Query().Get("autoSubscribe"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetFlexibleAutoSubscribeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetFlexibleAutoSubscribe(context.Background()).ProductId("1").AutoSubscribe(true).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetFlexibleAutoSubscribeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetFlexibleAutoSubscribeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SetFlexibleAutoSubscribe Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetFlexibleAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetFlexibleAutoSubscribe Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetFlexibleAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedAutoSubscribe Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/setAutoSubscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			require.Equal(t, "true", r.URL.Query().Get("autoSubscribe"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetLockedAutoSubscribeResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedAutoSubscribe(context.Background()).PositionId("1").AutoSubscribe(true).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetLockedAutoSubscribeResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetLockedAutoSubscribeResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedAutoSubscribe Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedAutoSubscribe Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedAutoSubscribe(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedProductRedeemOption Success", func(t *testing.T) {

		mockedJSON := `{"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/setRedeemOption", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("positionId"))
			require.Equal(t, "redeemTo_example", r.URL.Query().Get("redeemTo"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SetLockedProductRedeemOptionResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedProductRedeemOption(context.Background()).PositionId("1").RedeemTo("redeemTo_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SetLockedProductRedeemOptionResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SetLockedProductRedeemOptionResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedProductRedeemOption Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedProductRedeemOption(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SetLockedProductRedeemOption Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SetLockedProductRedeemOption(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SimpleAccount Success", func(t *testing.T) {

		mockedJSON := `{"totalAmountInBTC":"0.01067982","totalAmountInUSDT":"77.13289230","totalFlexibleAmountInBTC":"0.00000000","totalFlexibleAmountInUSDT":"0.00000000","totalLockedInBTC":"0.01067982","totalLockedInUSDT":"77.13289230"}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/account", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SimpleAccountResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SimpleAccount(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SimpleAccountResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SimpleAccountResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SimpleAccount Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SimpleAccount(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeFlexibleProduct Success", func(t *testing.T) {

		mockedJSON := `{"purchaseId":40607,"success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/flexible/subscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("productId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubscribeFlexibleProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeFlexibleProduct(context.Background()).ProductId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubscribeFlexibleProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubscribeFlexibleProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeFlexibleProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeFlexibleProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeFlexibleProduct Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeFlexibleProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeLockedProduct Success", func(t *testing.T) {

		mockedJSON := `{"purchaseId":40607,"positionId":"12345","success":true}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/simple-earn/locked/subscribe", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("projectId"))
			require.Equal(t, "1", r.URL.Query().Get("amount"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubscribeLockedProductResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeLockedProduct(context.Background()).ProjectId("1").Amount(float32(1.0)).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubscribeLockedProductResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubscribeLockedProductResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeLockedProduct Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test FlexibleLockedAPIService SubscribeLockedProduct Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.FlexibleLockedAPI.SubscribeLockedProduct(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
