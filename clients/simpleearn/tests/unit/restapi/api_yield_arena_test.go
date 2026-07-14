/*
Simple Earn REST API TEST

Testing YieldArenaAPIService

*/

package binancesimpleearnrestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/simpleearn"
	"github.com/binance/binance-connector-go/clients/simpleearn/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancesimpleearnrestapi_YieldArenaAPIService(t *testing.T) {

	t.Run("Test YieldArenaAPIService GetYieldArenaActivities Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"activities":[{"activityId":10001,"activityType":"AIRDROP","title":"Hold FDUSD & Earn BNB Airdrop","description":"Subscribe to FDUSD Flexible and earn bonus BNB airdrop.","rewardPoolInUsd":"50000","rewardToken":["BNB"],"redirectUrl":"https://www.binance.com/en/earn/arena/airdrop-123","startTime":1713052800000,"endTime":1713657600000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/earn/arena/activities", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetYieldArenaActivitiesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceSimpleEarnClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.YieldArenaAPI.GetYieldArenaActivities(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetYieldArenaActivitiesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetYieldArenaActivitiesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test YieldArenaAPIService GetYieldArenaActivities Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.YieldArenaAPI.GetYieldArenaActivities(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
