/*
Prediction Trading REST API TEST

Testing PositionAPIService

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

func Test_binancew3wpredictionrestapi_PositionAPIService(t *testing.T) {

	t.Run("Test PositionAPIService GetPositionByToken Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"position":{"positionId":1001,"vendor":"PREDICT_FUN","chainId":"56","tokenId":"112233","collateralSymbol":"USDT","topicType":"FLAT","marketTopicId":4229564,"marketId":5567895,"marketTopicTitle":"BTC Price 1h Up or Down?","marketTitle":"UP","outcomeName":"YES","outcomeIndex":0,"shares":"1923.07","avgPrice":"0.52","totalCost":"1.00","value":"1.06","currentPrice":"0.55","positionStatus":"OPEN","endDate":1748134800000,"unrealizedPnl":"0.06","realizedPnl":"0.00","pnl":"0.06","createdTime":1748131500000,"updatedTime":1748132000000}}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/position/token", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			require.Equal(t, "112233", r.URL.Query().Get("tokenId"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetPositionByTokenResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.GetPositionByToken(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").TokenId("112233").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetPositionByTokenResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetPositionByTokenResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PositionAPIService GetPositionByToken Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.GetPositionByToken(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService GetPositionByToken Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.PositionAPI.GetPositionByToken(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QueryPnL Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"chainId":"56","walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","pnl":{},"pnlList":[{"id":10001,"walletAddress":"0x12e32db8817e292508c34111cbc4b23340df542c","marketTopicId":4229564,"marketId":5567895,"tokenId":"112233","vendor":"PREDICT_FUN","currentShares":"1923.07","avgPrice":"0.52","currentPrice":"0.55","realizedPnl":"0.00","unrealizedPnl":"0.06","totalPnl":"0.06","pnlPercentage":"6.00","isResolved":false}],"totalCount":1,"totalRealizedPnl":"0.00","totalUnrealizedPnl":"0.06","totalPnl":"0.06"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/pnl/query", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryPnLResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QueryPnL(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryPnLResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryPnLResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PositionAPIService QueryPnL Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QueryPnL(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QueryPnL Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.PositionAPI.QueryPnL(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QueryPositions Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"summary":{"totalValue":"1523.45","positionValue":"523.45","walletBalance":"1000.00","totalClaimableAmount":"50.00","todayRealizedPnl":"15.30","todayRealizedPnlPercent":"3.10","todayTotalCost":"493.55"},"counts":{"ongoingCount":3,"endedCount":12,"pendingClaimCount":1},"positions":[{"positionId":1001,"vendor":"PREDICT_FUN","chainId":"56","tokenId":"112233","collateralSymbol":"USDT","topicType":"FLAT","marketTopicId":4229564,"marketId":5567895,"marketTopicTitle":"BTC Price 1h Up or Down?","marketTitle":"UP","outcomeName":"YES","outcomeIndex":0,"shares":"1923.07","avgPrice":"0.52","totalCost":"1.00","value":"1.06","currentPrice":"0.55","toWin":"1923.07","positionStatus":"OPEN","canClaim":false,"endDate":1748134800000,"unrealizedPnl":"0.06","unrealizedPnlPercent":"6.00","realizedPnl":"0.00","pnl":"0.06","createdTime":1748131500000,"updatedTime":1748132000000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/position/list", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryPositionsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QueryPositions(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryPositionsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryPositionsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PositionAPIService QueryPositions Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QueryPositions(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QueryPositions Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.PositionAPI.QueryPositions(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QueryPositionsByFilter Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"positions":[{"positionId":1001,"vendor":"PREDICT_FUN","chainId":"56","tokenId":"112233","collateralSymbol":"USDT","topicType":"FLAT","marketTopicId":4229564,"marketId":5567895,"marketTopicTitle":"BTC Price 1h Up or Down?","marketTitle":"UP","outcomeName":"YES","outcomeIndex":0,"shares":"1923.07","avgPrice":"0.52","totalCost":"1.00","value":"1.06","currentPrice":"0.55","positionStatus":"OPEN","endDate":1748134800000,"unrealizedPnl":"0.06","realizedPnl":"0.00","pnl":"0.06","createdTime":1748131500000,"updatedTime":1748132000000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/position/filter", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QueryPositionsByFilterResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QueryPositionsByFilter(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QueryPositionsByFilterResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QueryPositionsByFilterResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PositionAPIService QueryPositionsByFilter Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.PositionAPI.QueryPositionsByFilter(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QuerySettledPositionHistory Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"totalCount":5,"positions":[{"positionId":998,"vendor":"PREDICT_FUN","chainId":"56","tokenId":"112200","collateralSymbol":"USDT","topicType":"FLAT","marketTopicId":4229500,"marketId":5567800,"marketTopicTitle":"BTC Price 1h Up or Down?","marketTitle":"UP","outcomeName":"YES","outcomeIndex":0,"shares":"5000.00","avgPrice":"0.30","totalCost":"1.50","value":"5000.00","currentPrice":"1.00","toWin":"5000.00","positionStatus":"CLAIMED","isWinner":true,"redeemStatus":"CONFIRMED","endDate":1748040000000,"finalOutcome":"YES","realizedPnl":"4998.50","pnl":"4998.50","claimAmount":"5000.00","settledDate":1748040000000,"createdTime":1748020000000,"updatedTime":1748040500000}]}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/w3w/wallet/prediction/position/settled-history", r.URL.Path)
			require.Equal(t, "0x12e32db8817e292508c34111cbc4b23340df542c", r.URL.Query().Get("walletAddress"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.QuerySettledPositionHistoryResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QuerySettledPositionHistory(context.Background()).WalletAddress("0x12e32db8817e292508c34111cbc4b23340df542c").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.QuerySettledPositionHistoryResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.QuerySettledPositionHistoryResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test PositionAPIService QuerySettledPositionHistory Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceW3wPredictionClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.PositionAPI.QuerySettledPositionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test PositionAPIService QuerySettledPositionHistory Server Error", func(t *testing.T) {
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

		resp, err := apiClient.RestApi.PositionAPI.QuerySettledPositionHistory(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
