/*
Wallet REST API TEST

Testing TravelRuleAPIService

*/

package binancewalletrestapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/clients/wallet/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	"github.com/stretchr/testify/require"
)

func Test_binancewalletrestapi_TravelRuleAPIService(t *testing.T) {

	t.Run("Test TravelRuleAPIService BrokerWithdraw Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"trId":123456,"accepted":true,"info":"Withdraw request accepted"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/broker/withdraw/apply", r.URL.Path)
			require.Equal(t, "address_example", r.URL.Query().Get("address"))
			require.Equal(t, "BTC", r.URL.Query().Get("coin"))
			require.Equal(t, fmt.Sprintf("%v", float32(1.0)), r.URL.Query().Get("amount"))
			require.Equal(t, "1", r.URL.Query().Get("withdrawOrderId"))
			require.Equal(t, "questionnaire_example", r.URL.Query().Get("questionnaire"))
			require.Equal(t, "originatorPii_example", r.URL.Query().Get("originatorPii"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.BrokerWithdrawResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.BrokerWithdraw(context.Background()).Address("address_example").Coin("BTC").Amount(float32(1.0)).WithdrawOrderId("1").Questionnaire("questionnaire_example").OriginatorPii("originatorPii_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.BrokerWithdrawResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.BrokerWithdrawResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService BrokerWithdraw Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.BrokerWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService BrokerWithdraw Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.BrokerWithdraw(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService CheckQuestionnaireRequirements Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"questionnaireCountryCode":"AE"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/questionnaire-requirements", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.CheckQuestionnaireRequirementsResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.CheckQuestionnaireRequirements(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.CheckQuestionnaireRequirementsResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.CheckQuestionnaireRequirementsResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService CheckQuestionnaireRequirements Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.CheckQuestionnaireRequirements(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService DepositHistoryTravelRule Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"trId":123451123,"tranId":17644346245865,"amount":"0.001","coin":"BNB","network":"BNB","depositStatus":0,"travelRuleStatus":1,"travelRuleStatusV2":"PENDING","address":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","addressTag":"101764890","txId":"98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC","insertTime":1661493146000,"completeTime":1661493206000,"transferType":0,"confirmTimes":"1/1","requireQuestionnaire":false,"questionnaire":"{\"question1\":\"answer1\",\"question2\":\"answer2\"}"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/deposit/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DepositHistoryTravelRuleResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryTravelRule(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DepositHistoryTravelRuleResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DepositHistoryTravelRuleResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService DepositHistoryTravelRule Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryTravelRule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService DepositHistoryV2 Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"depositId":"4615328107052018945","amount":"0.01","network":"AVAXC","coin":"AVAX","depositStatus":1,"travelRuleReqStatus":0,"address":"0x0010627ab66d69232f4080d54e0f838b4dc3894a","addressTag":"","txId":"0xdde578983015741eed764e7ca10defb5a2caafdca3db5f92872d24a96beb1879","transferType":0,"confirmTimes":"12/12","requireQuestionnaire":false,"questionnaire":{"vaspName":"BINANCE","depositOriginator":0},"insertTime":1753053392000}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/localentity/deposit/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.DepositHistoryV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryV2(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.DepositHistoryV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.DepositHistoryV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService DepositHistoryV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService FetchAddressVerificationList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"status":"PENDING","token":"AVAX","network":"AVAXC","walletAddress":"0xc03a6aa728a8dde7464c33828424ede7553a0021","addressQuestionnaire":{"sendTo":1,"satoshiToken":"AVAX","isAddressOwner":1,"verifyMethod":1}}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/addressVerify/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.FetchAddressVerificationListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.FetchAddressVerificationList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.FetchAddressVerificationListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.FetchAddressVerificationListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService FetchAddressVerificationList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.FetchAddressVerificationList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService GetCountryList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"countries":[{"countryCode":"au","countryName":"Australia","blockType":"supported","depositAllowed":true,"withdrawalAllowed":true,"hasRegionRestrictions":false}],"lastUpdated":1716300000000}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/country/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetCountryListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.GetCountryList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetCountryListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetCountryListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService GetCountryList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.GetCountryList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService GetRegionList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"countryCode":"au","regions":[{"regionName":"New South Wales","blockType":"supported","depositAllowed":true,"withdrawalAllowed":true}],"lastUpdated":1716300000000}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/region/list", r.URL.Path)
			require.Equal(t, "au", r.URL.Query().Get("countryCode"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetRegionListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.GetRegionList(context.Background()).CountryCode("au").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetRegionListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetRegionListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService GetRegionList Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.GetRegionList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService GetRegionList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.GetRegionList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaire Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"trId":765127651,"accepted":true,"info":"Deposit questionnaire accepted."}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/broker/deposit/provide-info", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("subAccountId"))
			require.Equal(t, "1", r.URL.Query().Get("depositId"))
			require.Equal(t, "questionnaire_example", r.URL.Query().Get("questionnaire"))
			require.Equal(t, "beneficiaryPii_example", r.URL.Query().Get("beneficiaryPii"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubmitDepositQuestionnaireResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaire(context.Background()).SubAccountId("1").DepositId(int64(1)).Questionnaire("questionnaire_example").BeneficiaryPii("beneficiaryPii_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubmitDepositQuestionnaireResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubmitDepositQuestionnaireResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaire Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaire(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaire Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaire(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireTravelRule Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"trId":765127651,"accepted":true,"info":"Deposit questionnaire accepted."}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/deposit/provide-info", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("tranId"))
			require.Equal(t, "questionnaire_example", r.URL.Query().Get("questionnaire"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubmitDepositQuestionnaireTravelRuleResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireTravelRule(context.Background()).TranId(int64(1)).Questionnaire("questionnaire_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubmitDepositQuestionnaireTravelRuleResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubmitDepositQuestionnaireTravelRuleResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireTravelRule Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireTravelRule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireTravelRule Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireTravelRule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireV2 Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"trId":765127651,"accepted":true,"info":"Deposit questionnaire accepted."}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/localentity/deposit/provide-info", r.URL.Path)
			require.Equal(t, "1", r.URL.Query().Get("depositId"))
			require.Equal(t, "questionnaire_example", r.URL.Query().Get("questionnaire"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.SubmitDepositQuestionnaireV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireV2(context.Background()).DepositId(int64(1)).Questionnaire("questionnaire_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.SubmitDepositQuestionnaireV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.SubmitDepositQuestionnaireV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireV2 Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService SubmitDepositQuestionnaireV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService VaspList Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"vaspName":"Binance","vaspCode":"BINANCE","identifier":"xxx"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/vasp", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.VaspListResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.VaspList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.VaspListResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.VaspListResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService VaspList Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.VaspList(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService WithdrawHistoryV1 Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"id":"b6ae22b3aa844210a7041aee7589627c","trId":1234456,"amount":"8.91000000","transactionFee":"0.004","coin":"USDT","withdrawalStatus":6,"travelRuleStatus":0,"address":"0x94df8b352de7f46f64b01d3666bf6e936e44ce60","txId":"0xb5ef8c13b968a406cc62a93a8bd80f9e9a906ef1b3fcf20a2e48573c17659268","applyTime":"2019-10-12 11:12:02","network":"ETH","transferType":0,"withdrawOrderId":"WITHDRAWtest123","info":"The address is not valid. Please confirm with the recipient","confirmNo":3,"walletType":1,"txKey":"","questionnaire":"{\"question1\":\"answer1\",\"question2\":\"answer2\"}","completeTime":"2023-03-23 16:52:41"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/withdraw/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.WithdrawHistoryV1Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV1(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.WithdrawHistoryV1Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.WithdrawHistoryV1Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService WithdrawHistoryV1 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV1(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService WithdrawHistoryV2 Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `[{"id":"b6ae22b3aa844210a7041aee7589627c","trId":1234456,"amount":"8.91000000","transactionFee":"0.004","coin":"USDT","withdrawalStatus":6,"travelRuleStatus":0,"address":"0x94df8b352de7f46f64b01d3666bf6e936e44ce60","txId":"0xb5ef8c13b968a406cc62a93a8bd80f9e9a906ef1b3fcf20a2e48573c17659268","applyTime":"2019-10-12 11:12:02","network":"ETH","transferType":0,"withdrawOrderId":"WITHDRAWtest123","info":"The address is not valid. Please confirm with the recipient","confirmNo":3,"walletType":1,"txKey":"","questionnaire":"{\"question1\":\"answer1\",\"question2\":\"answer2\"}","completeTime":"2023-03-23 16:52:41"}]`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v2/localentity/withdraw/history", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.WithdrawHistoryV2Response
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV2(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.WithdrawHistoryV2Response]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.WithdrawHistoryV2Response{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService WithdrawHistoryV2 Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV2(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService WithdrawTravelRule Success", func(t *testing.T) {

		var mockedJSON string
		mockedJSON = `{"trId":123456,"accepted":true,"info":"Withdraw request accepted"}`
		if mockedJSON == "" {
			mockedJSON = `{}`
		}
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/sapi/v1/localentity/withdraw/apply", r.URL.Path)
			require.Equal(t, "BTC", r.URL.Query().Get("coin"))
			require.Equal(t, "address_example", r.URL.Query().Get("address"))
			require.Equal(t, fmt.Sprintf("%v", float32(1.0)), r.URL.Query().Get("amount"))
			require.Equal(t, "questionnaire_example", r.URL.Query().Get("questionnaire"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.WithdrawTravelRuleResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawTravelRule(context.Background()).Coin("BTC").Address("address_example").Amount(float32(1.0)).Questionnaire("questionnaire_example").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.WithdrawTravelRuleResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.WithdrawTravelRuleResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test TravelRuleAPIService WithdrawTravelRule Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawTravelRule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test TravelRuleAPIService WithdrawTravelRule Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceWalletClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawTravelRule(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
