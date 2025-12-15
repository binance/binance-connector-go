# DepositHistoryV2ResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**DepositId** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**DepositStatus** | Pointer to **int64** |  | [optional] 
**TravelRuleReqStatus** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressTag** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**TransferType** | Pointer to **int64** |  | [optional] 
**ConfirmTimes** | Pointer to **string** |  | [optional] 
**RequireQuestionnaire** | Pointer to **bool** |  | [optional] 
**Questionnaire** | Pointer to [**DepositHistoryV2ResponseInnerQuestionnaire**](DepositHistoryV2ResponseInnerQuestionnaire.md) |  | [optional] 
**InsertTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewDepositHistoryV2ResponseInner

`func NewDepositHistoryV2ResponseInner() *DepositHistoryV2ResponseInner`

NewDepositHistoryV2ResponseInner instantiates a new DepositHistoryV2ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositHistoryV2ResponseInnerWithDefaults

`func NewDepositHistoryV2ResponseInnerWithDefaults() *DepositHistoryV2ResponseInner`

NewDepositHistoryV2ResponseInnerWithDefaults instantiates a new DepositHistoryV2ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositId

`func (o *DepositHistoryV2ResponseInner) GetDepositId() string`

GetDepositId returns the DepositId field if non-nil, zero value otherwise.

### GetDepositIdOk

`func (o *DepositHistoryV2ResponseInner) GetDepositIdOk() (*string, bool)`

GetDepositIdOk returns a tuple with the DepositId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositId

`func (o *DepositHistoryV2ResponseInner) SetDepositId(v string)`

SetDepositId sets DepositId field to given value.

### HasDepositId

`func (o *DepositHistoryV2ResponseInner) HasDepositId() bool`

HasDepositId returns a boolean if a field has been set.

### GetAmount

`func (o *DepositHistoryV2ResponseInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositHistoryV2ResponseInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositHistoryV2ResponseInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DepositHistoryV2ResponseInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetNetwork

`func (o *DepositHistoryV2ResponseInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DepositHistoryV2ResponseInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DepositHistoryV2ResponseInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DepositHistoryV2ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCoin

`func (o *DepositHistoryV2ResponseInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *DepositHistoryV2ResponseInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *DepositHistoryV2ResponseInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *DepositHistoryV2ResponseInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetDepositStatus

`func (o *DepositHistoryV2ResponseInner) GetDepositStatus() int64`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *DepositHistoryV2ResponseInner) GetDepositStatusOk() (*int64, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *DepositHistoryV2ResponseInner) SetDepositStatus(v int64)`

SetDepositStatus sets DepositStatus field to given value.

### HasDepositStatus

`func (o *DepositHistoryV2ResponseInner) HasDepositStatus() bool`

HasDepositStatus returns a boolean if a field has been set.

### GetTravelRuleReqStatus

`func (o *DepositHistoryV2ResponseInner) GetTravelRuleReqStatus() int64`

GetTravelRuleReqStatus returns the TravelRuleReqStatus field if non-nil, zero value otherwise.

### GetTravelRuleReqStatusOk

`func (o *DepositHistoryV2ResponseInner) GetTravelRuleReqStatusOk() (*int64, bool)`

GetTravelRuleReqStatusOk returns a tuple with the TravelRuleReqStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelRuleReqStatus

`func (o *DepositHistoryV2ResponseInner) SetTravelRuleReqStatus(v int64)`

SetTravelRuleReqStatus sets TravelRuleReqStatus field to given value.

### HasTravelRuleReqStatus

`func (o *DepositHistoryV2ResponseInner) HasTravelRuleReqStatus() bool`

HasTravelRuleReqStatus returns a boolean if a field has been set.

### GetAddress

`func (o *DepositHistoryV2ResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DepositHistoryV2ResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DepositHistoryV2ResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DepositHistoryV2ResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressTag

`func (o *DepositHistoryV2ResponseInner) GetAddressTag() string`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *DepositHistoryV2ResponseInner) GetAddressTagOk() (*string, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *DepositHistoryV2ResponseInner) SetAddressTag(v string)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *DepositHistoryV2ResponseInner) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetTxId

`func (o *DepositHistoryV2ResponseInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *DepositHistoryV2ResponseInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *DepositHistoryV2ResponseInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *DepositHistoryV2ResponseInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetTransferType

`func (o *DepositHistoryV2ResponseInner) GetTransferType() int64`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *DepositHistoryV2ResponseInner) GetTransferTypeOk() (*int64, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *DepositHistoryV2ResponseInner) SetTransferType(v int64)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *DepositHistoryV2ResponseInner) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetConfirmTimes

`func (o *DepositHistoryV2ResponseInner) GetConfirmTimes() string`

GetConfirmTimes returns the ConfirmTimes field if non-nil, zero value otherwise.

### GetConfirmTimesOk

`func (o *DepositHistoryV2ResponseInner) GetConfirmTimesOk() (*string, bool)`

GetConfirmTimesOk returns a tuple with the ConfirmTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTimes

`func (o *DepositHistoryV2ResponseInner) SetConfirmTimes(v string)`

SetConfirmTimes sets ConfirmTimes field to given value.

### HasConfirmTimes

`func (o *DepositHistoryV2ResponseInner) HasConfirmTimes() bool`

HasConfirmTimes returns a boolean if a field has been set.

### GetRequireQuestionnaire

`func (o *DepositHistoryV2ResponseInner) GetRequireQuestionnaire() bool`

GetRequireQuestionnaire returns the RequireQuestionnaire field if non-nil, zero value otherwise.

### GetRequireQuestionnaireOk

`func (o *DepositHistoryV2ResponseInner) GetRequireQuestionnaireOk() (*bool, bool)`

GetRequireQuestionnaireOk returns a tuple with the RequireQuestionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireQuestionnaire

`func (o *DepositHistoryV2ResponseInner) SetRequireQuestionnaire(v bool)`

SetRequireQuestionnaire sets RequireQuestionnaire field to given value.

### HasRequireQuestionnaire

`func (o *DepositHistoryV2ResponseInner) HasRequireQuestionnaire() bool`

HasRequireQuestionnaire returns a boolean if a field has been set.

### GetQuestionnaire

`func (o *DepositHistoryV2ResponseInner) GetQuestionnaire() DepositHistoryV2ResponseInnerQuestionnaire`

GetQuestionnaire returns the Questionnaire field if non-nil, zero value otherwise.

### GetQuestionnaireOk

`func (o *DepositHistoryV2ResponseInner) GetQuestionnaireOk() (*DepositHistoryV2ResponseInnerQuestionnaire, bool)`

GetQuestionnaireOk returns a tuple with the Questionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaire

`func (o *DepositHistoryV2ResponseInner) SetQuestionnaire(v DepositHistoryV2ResponseInnerQuestionnaire)`

SetQuestionnaire sets Questionnaire field to given value.

### HasQuestionnaire

`func (o *DepositHistoryV2ResponseInner) HasQuestionnaire() bool`

HasQuestionnaire returns a boolean if a field has been set.

### GetInsertTime

`func (o *DepositHistoryV2ResponseInner) GetInsertTime() int64`

GetInsertTime returns the InsertTime field if non-nil, zero value otherwise.

### GetInsertTimeOk

`func (o *DepositHistoryV2ResponseInner) GetInsertTimeOk() (*int64, bool)`

GetInsertTimeOk returns a tuple with the InsertTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTime

`func (o *DepositHistoryV2ResponseInner) SetInsertTime(v int64)`

SetInsertTime sets InsertTime field to given value.

### HasInsertTime

`func (o *DepositHistoryV2ResponseInner) HasInsertTime() bool`

HasInsertTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


