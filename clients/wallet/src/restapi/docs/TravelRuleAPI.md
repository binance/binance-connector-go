# \TravelRuleAPI

All URIs are relative to *https://api.binance.com*

Method        | HTTP request  | Description
------------- | ------------- | -------------
[**BrokerWithdraw**](TravelRuleAPI.md#BrokerWithdraw) | **Post** /sapi/v1/localentity/broker/withdraw/apply | Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)
[**CheckQuestionnaireRequirements**](TravelRuleAPI.md#CheckQuestionnaireRequirements) | **Get** /sapi/v1/localentity/questionnaire-requirements | Check Questionnaire Requirements (for local entities that require travel rule) (supporting network) (USER_DATA)
[**DepositHistoryTravelRule**](TravelRuleAPI.md#DepositHistoryTravelRule) | **Get** /sapi/v1/localentity/deposit/history | Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)
[**DepositHistoryV2**](TravelRuleAPI.md#DepositHistoryV2) | **Get** /sapi/v2/localentity/deposit/history | Deposit History V2 (for local entities that required travel rule) (supporting network) (USER_DATA)
[**FetchAddressVerificationList**](TravelRuleAPI.md#FetchAddressVerificationList) | **Get** /sapi/v1/addressVerify/list | Fetch address verification list (USER_DATA)
[**SubmitDepositQuestionnaire**](TravelRuleAPI.md#SubmitDepositQuestionnaire) | **Put** /sapi/v1/localentity/broker/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
[**SubmitDepositQuestionnaireTravelRule**](TravelRuleAPI.md#SubmitDepositQuestionnaireTravelRule) | **Put** /sapi/v1/localentity/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
[**VaspList**](TravelRuleAPI.md#VaspList) | **Get** /sapi/v1/localentity/vasp | VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WithdrawHistoryV1**](TravelRuleAPI.md#WithdrawHistoryV1) | **Get** /sapi/v1/localentity/withdraw/history | Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WithdrawHistoryV2**](TravelRuleAPI.md#WithdrawHistoryV2) | **Get** /sapi/v2/localentity/withdraw/history | Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WithdrawTravelRule**](TravelRuleAPI.md#WithdrawTravelRule) | **Post** /sapi/v1/localentity/withdraw/apply | Withdraw (for local entities that require travel rule) (USER_DATA)


## BrokerWithdraw

> BrokerWithdrawResponse BrokerWithdraw(ctx).Address(address).Coin(coin).Amount(amount).WithdrawOrderId(withdrawOrderId).Questionnaire(questionnaire).OriginatorPii(originatorPii).Signature(signature).AddressTag(addressTag).Network(network).AddressName(addressName).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()

Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	address := "address_example" // string | 
	coin := "coin_example" // string | 
	amount := float32(1.0) // float32 | 
	withdrawOrderId := "1" // string | withdrawID defined by the client (i.e. client's internal withdrawID)
	questionnaire := "questionnaire_example" // string | JSON format questionnaire answers.
	originatorPii := "originatorPii_example" // string | JSON format originator Pii, see StandardPii section below
	signature := "signature_example" // string | Must be the last parameter.
	addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)
	network := "network_example" // string |  (optional)
	addressName := "addressName_example" // string | Description of the address. Address book cap is 200, space in name should be encoded into `%20` (optional)
	transactionFeeFlag := false // bool | When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`. (optional)
	walletType := int64(0) // int64 | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.BrokerWithdraw(context.Background()).Address(address).Coin(coin).Amount(amount).WithdrawOrderId(withdrawOrderId).Questionnaire(questionnaire).OriginatorPii(originatorPii).Signature(signature).AddressTag(addressTag).Network(network).AddressName(addressName).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.BrokerWithdraw``: %v\n", err)
		return
	}

	// response from `BrokerWithdraw`: BrokerWithdrawResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | 
 **coin** | **string** |  | 
 **amount** | **float32** |  | 
 **withdrawOrderId** | **string** | withdrawID defined by the client (i.e. client&#39;s internal withdrawID) | 
 **questionnaire** | **string** | JSON format questionnaire answers. | 
 **originatorPii** | **string** | JSON format originator Pii, see StandardPii section below | 
 **signature** | **string** | Must be the last parameter. | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 
 **network** | **string** |  | 
 **addressName** | **string** | Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60; | 
 **transactionFeeFlag** | **bool** | When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;. | 
 **walletType** | **int64** | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit | 

### Return type

[**BrokerWithdrawResponse**](BrokerWithdrawResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## CheckQuestionnaireRequirements

> CheckQuestionnaireRequirementsResponse CheckQuestionnaireRequirements(ctx).RecvWindow(recvWindow).Execute()

Check Questionnaire Requirements (for local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.CheckQuestionnaireRequirements(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.CheckQuestionnaireRequirements``: %v\n", err)
		return
	}

	// response from `CheckQuestionnaireRequirements`: CheckQuestionnaireRequirementsResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**CheckQuestionnaireRequirementsResponse**](CheckQuestionnaireRequirementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DepositHistoryTravelRule

> DepositHistoryTravelRuleResponse DepositHistoryTravelRule(ctx).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()

Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	trId := "1" // string | Comma(,) separated list of travel rule record Ids. (optional)
	txId := "1" // string |  (optional)
	tranId := "1" // string | Comma(,) separated list of wallet tran Ids. (optional)
	network := "network_example" // string |  (optional)
	coin := "coin_example" // string |  (optional)
	travelRuleStatus := int64(789) // int64 | 0:Completed,1:Pending,2:Failed (optional)
	pendingQuestionnaire := true // bool | true: Only return records that pending deposit questionnaire. false/not provided: return all records. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryTravelRule(context.Background()).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.DepositHistoryTravelRule``: %v\n", err)
		return
	}

	// response from `DepositHistoryTravelRule`: DepositHistoryTravelRuleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | 
 **txId** | **string** |  | 
 **tranId** | **string** | Comma(,) separated list of wallet tran Ids. | 
 **network** | **string** |  | 
 **coin** | **string** |  | 
 **travelRuleStatus** | **int64** | 0:Completed,1:Pending,2:Failed | 
 **pendingQuestionnaire** | **bool** | true: Only return records that pending deposit questionnaire. false/not provided: return all records. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 

### Return type

[**DepositHistoryTravelRuleResponse**](DepositHistoryTravelRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## DepositHistoryV2

> DepositHistoryV2Response DepositHistoryV2(ctx).DepositId(depositId).TxId(txId).Network(network).Coin(coin).RetrieveQuestionnaire(retrieveQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()

Deposit History V2 (for local entities that required travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	depositId := "1" // string | Comma(,) separated list of wallet tran Ids. (optional)
	txId := "1" // string |  (optional)
	network := "network_example" // string |  (optional)
	coin := "coin_example" // string |  (optional)
	retrieveQuestionnaire := true // bool | true: return `questionnaire` within response. (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.DepositHistoryV2(context.Background()).DepositId(depositId).TxId(txId).Network(network).Coin(coin).RetrieveQuestionnaire(retrieveQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.DepositHistoryV2``: %v\n", err)
		return
	}

	// response from `DepositHistoryV2`: DepositHistoryV2Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **depositId** | **string** | Comma(,) separated list of wallet tran Ids. | 
 **txId** | **string** |  | 
 **network** | **string** |  | 
 **coin** | **string** |  | 
 **retrieveQuestionnaire** | **bool** | true: return &#x60;questionnaire&#x60; within response. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 

### Return type

[**DepositHistoryV2Response**](DepositHistoryV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## FetchAddressVerificationList

> FetchAddressVerificationListResponse FetchAddressVerificationList(ctx).RecvWindow(recvWindow).Execute()

Fetch address verification list (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.FetchAddressVerificationList(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.FetchAddressVerificationList``: %v\n", err)
		return
	}

	// response from `FetchAddressVerificationList`: FetchAddressVerificationListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**FetchAddressVerificationListResponse**](FetchAddressVerificationListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubmitDepositQuestionnaire

> SubmitDepositQuestionnaireResponse SubmitDepositQuestionnaire(ctx).SubAccountId(subAccountId).DepositId(depositId).Questionnaire(questionnaire).BeneficiaryPii(beneficiaryPii).Signature(signature).Network(network).Coin(coin).Amount(amount).Address(address).AddressTag(addressTag).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	subAccountId := "1" // string | External user ID.
	depositId := "1" // string | Wallet deposit ID.
	questionnaire := "questionnaire_example" // string | JSON format questionnaire answers.
	beneficiaryPii := "beneficiaryPii_example" // string | JSON format beneficiary Pii.
	signature := "signature_example" // string | Must be the last parameter.
	network := "network_example" // string |  (optional)
	coin := "coin_example" // string |  (optional)
	amount := float32(1.0) // float32 |  (optional)
	address := "address_example" // string |  (optional)
	addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaire(context.Background()).SubAccountId(subAccountId).DepositId(depositId).Questionnaire(questionnaire).BeneficiaryPii(beneficiaryPii).Signature(signature).Network(network).Coin(coin).Amount(amount).Address(address).AddressTag(addressTag).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.SubmitDepositQuestionnaire``: %v\n", err)
		return
	}

	// response from `SubmitDepositQuestionnaire`: SubmitDepositQuestionnaireResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **subAccountId** | **string** | External user ID. | 
 **depositId** | **string** | Wallet deposit ID. | 
 **questionnaire** | **string** | JSON format questionnaire answers. | 
 **beneficiaryPii** | **string** | JSON format beneficiary Pii. | 
 **signature** | **string** | Must be the last parameter. | 
 **network** | **string** |  | 
 **coin** | **string** |  | 
 **amount** | **float32** |  | 
 **address** | **string** |  | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 

### Return type

[**SubmitDepositQuestionnaireResponse**](SubmitDepositQuestionnaireResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## SubmitDepositQuestionnaireTravelRule

> SubmitDepositQuestionnaireTravelRuleResponse SubmitDepositQuestionnaireTravelRule(ctx).TranId(tranId).Questionnaire(questionnaire).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	tranId := int64(1) // int64 | Wallet tran ID
	questionnaire := "questionnaire_example" // string | JSON format questionnaire answers.

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.SubmitDepositQuestionnaireTravelRule(context.Background()).TranId(tranId).Questionnaire(questionnaire).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.SubmitDepositQuestionnaireTravelRule``: %v\n", err)
		return
	}

	// response from `SubmitDepositQuestionnaireTravelRule`: SubmitDepositQuestionnaireTravelRuleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **tranId** | **int64** | Wallet tran ID | 
 **questionnaire** | **string** | JSON format questionnaire answers. | 

### Return type

[**SubmitDepositQuestionnaireTravelRuleResponse**](SubmitDepositQuestionnaireTravelRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## VaspList

> VaspListResponse VaspList(ctx).RecvWindow(recvWindow).Execute()

VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.VaspList(context.Background()).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.VaspList``: %v\n", err)
		return
	}

	// response from `VaspList`: VaspListResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 

### Return type

[**VaspListResponse**](VaspListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WithdrawHistoryV1

> WithdrawHistoryV1Response WithdrawHistoryV1(ctx).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	trId := "1" // string | Comma(,) separated list of travel rule record Ids. (optional)
	txId := "1" // string |  (optional)
	withdrawOrderId := "1" // string | client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query. (optional)
	network := "network_example" // string |  (optional)
	coin := "coin_example" // string |  (optional)
	travelRuleStatus := int64(789) // int64 | 0:Completed,1:Pending,2:Failed (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV1(context.Background()).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.WithdrawHistoryV1``: %v\n", err)
		return
	}

	// response from `WithdrawHistoryV1`: WithdrawHistoryV1Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | 
 **txId** | **string** |  | 
 **withdrawOrderId** | **string** | client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query. | 
 **network** | **string** |  | 
 **coin** | **string** |  | 
 **travelRuleStatus** | **int64** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawHistoryV1Response**](WithdrawHistoryV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WithdrawHistoryV2

> WithdrawHistoryV2Response WithdrawHistoryV2(ctx).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	trId := "1" // string | Comma(,) separated list of travel rule record Ids. (optional)
	txId := "1" // string |  (optional)
	withdrawOrderId := "1" // string | client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query. (optional)
	network := "network_example" // string |  (optional)
	coin := "coin_example" // string |  (optional)
	travelRuleStatus := int64(789) // int64 | 0:Completed,1:Pending,2:Failed (optional)
	offset := int64(0) // int64 | Default: 0 (optional)
	limit := int64(7) // int64 | min 7, max 30, default 7 (optional)
	startTime := int64(1623319461670) // int64 |  (optional)
	endTime := int64(1641782889000) // int64 |  (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawHistoryV2(context.Background()).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.WithdrawHistoryV2``: %v\n", err)
		return
	}

	// response from `WithdrawHistoryV2`: WithdrawHistoryV2Response
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | 
 **txId** | **string** |  | 
 **withdrawOrderId** | **string** | client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query. | 
 **network** | **string** |  | 
 **coin** | **string** |  | 
 **travelRuleStatus** | **int64** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int64** | Default: 0 | 
 **limit** | **int64** | min 7, max 30, default 7 | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawHistoryV2Response**](WithdrawHistoryV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)


## WithdrawTravelRule

> WithdrawTravelRuleResponse WithdrawTravelRule(ctx).Coin(coin).Address(address).Amount(amount).Questionnaire(questionnaire).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()

Withdraw (for local entities that require travel rule) (USER_DATA)


### Example

```go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	models "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	coin := "coin_example" // string | 
	address := "address_example" // string | 
	amount := float32(1.0) // float32 | 
	questionnaire := "questionnaire_example" // string | JSON format questionnaire answers.
	withdrawOrderId := "1" // string | client side id for withdrawal, if provided in POST `/sapi/v1/capital/withdraw/apply`, can be used here for query. (optional)
	network := "network_example" // string |  (optional)
	addressTag := "addressTag_example" // string | Secondary address identifier for coins like XRP,XMR etc. (optional)
	transactionFeeFlag := false // bool | When making internal transfer, `true` for returning the fee to the destination account; `false` for returning the fee back to the departure account. Default `false`. (optional)
	name := "name_example" // string | Description of the address. Address book cap is 200, space in name should be encoded into `%20` (optional)
	walletType := int64(0) // int64 | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \"selected wallet\" under wallet->Fiat and Spot/Funding->Deposit (optional)
	recvWindow := int64(5000) // int64 |  (optional)

	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := models.NewBinanceWalletClient(models.WithRestAPI(configuration))

	resp, err := apiClient.RestApi.TravelRuleAPI.WithdrawTravelRule(context.Background()).Coin(coin).Address(address).Amount(amount).Questionnaire(questionnaire).WithdrawOrderId(withdrawOrderId).Network(network).AddressTag(addressTag).TransactionFeeFlag(transactionFeeFlag).Name(name).WalletType(walletType).RecvWindow(recvWindow).Execute()
	if err != nil {
		log.Println(os.Stderr, "Error when calling `TravelRuleAPI.WithdrawTravelRule``: %v\n", err)
		return
	}

	// response from `WithdrawTravelRule`: WithdrawTravelRuleResponse
	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### Path Parameters

Name          | Type          | Description   | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | 
 **address** | **string** |  | 
 **amount** | **float32** |  | 
 **questionnaire** | **string** | JSON format questionnaire answers. | 
 **withdrawOrderId** | **string** | client side id for withdrawal, if provided in POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, can be used here for query. | 
 **network** | **string** |  | 
 **addressTag** | **string** | Secondary address identifier for coins like XRP,XMR etc. | 
 **transactionFeeFlag** | **bool** | When making internal transfer, &#x60;true&#x60; for returning the fee to the destination account; &#x60;false&#x60; for returning the fee back to the departure account. Default &#x60;false&#x60;. | 
 **name** | **string** | Description of the address. Address book cap is 200, space in name should be encoded into &#x60;%20&#x60; | 
 **walletType** | **int64** | The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current \&quot;selected wallet\&quot; under wallet-&gt;Fiat and Spot/Funding-&gt;Deposit | 
 **recvWindow** | **int64** |  | 

### Return type

[**WithdrawTravelRuleResponse**](WithdrawTravelRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Accept**: application/json

[[Back to README]](../../../README.md)

