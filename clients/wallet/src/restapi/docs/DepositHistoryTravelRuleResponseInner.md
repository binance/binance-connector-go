# DepositHistoryTravelRuleResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TrId** | Pointer to **int64** |  | [optional] 
**TranId** | Pointer to **int64** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**DepositStatus** | Pointer to **int64** |  | [optional] 
**TravelRuleStatus** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressTag** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**InsertTime** | Pointer to **int64** |  | [optional] 
**TransferType** | Pointer to **int64** |  | [optional] 
**ConfirmTimes** | Pointer to **string** |  | [optional] 
**UnlockConfirm** | Pointer to **int64** |  | [optional] 
**WalletType** | Pointer to **int64** |  | [optional] 
**RequireQuestionnaire** | Pointer to **bool** |  | [optional] 
**Questionnaire** | Pointer to **string** |  | [optional] 

## Methods

### NewDepositHistoryTravelRuleResponseInner

`func NewDepositHistoryTravelRuleResponseInner() *DepositHistoryTravelRuleResponseInner`

NewDepositHistoryTravelRuleResponseInner instantiates a new DepositHistoryTravelRuleResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositHistoryTravelRuleResponseInnerWithDefaults

`func NewDepositHistoryTravelRuleResponseInnerWithDefaults() *DepositHistoryTravelRuleResponseInner`

NewDepositHistoryTravelRuleResponseInnerWithDefaults instantiates a new DepositHistoryTravelRuleResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrId

`func (o *DepositHistoryTravelRuleResponseInner) GetTrId() int64`

GetTrId returns the TrId field if non-nil, zero value otherwise.

### GetTrIdOk

`func (o *DepositHistoryTravelRuleResponseInner) GetTrIdOk() (*int64, bool)`

GetTrIdOk returns a tuple with the TrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrId

`func (o *DepositHistoryTravelRuleResponseInner) SetTrId(v int64)`

SetTrId sets TrId field to given value.

### HasTrId

`func (o *DepositHistoryTravelRuleResponseInner) HasTrId() bool`

HasTrId returns a boolean if a field has been set.

### GetTranId

`func (o *DepositHistoryTravelRuleResponseInner) GetTranId() int64`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *DepositHistoryTravelRuleResponseInner) GetTranIdOk() (*int64, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *DepositHistoryTravelRuleResponseInner) SetTranId(v int64)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *DepositHistoryTravelRuleResponseInner) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetAmount

`func (o *DepositHistoryTravelRuleResponseInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositHistoryTravelRuleResponseInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositHistoryTravelRuleResponseInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DepositHistoryTravelRuleResponseInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCoin

`func (o *DepositHistoryTravelRuleResponseInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *DepositHistoryTravelRuleResponseInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *DepositHistoryTravelRuleResponseInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *DepositHistoryTravelRuleResponseInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetNetwork

`func (o *DepositHistoryTravelRuleResponseInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DepositHistoryTravelRuleResponseInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DepositHistoryTravelRuleResponseInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DepositHistoryTravelRuleResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetDepositStatus

`func (o *DepositHistoryTravelRuleResponseInner) GetDepositStatus() int64`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *DepositHistoryTravelRuleResponseInner) GetDepositStatusOk() (*int64, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *DepositHistoryTravelRuleResponseInner) SetDepositStatus(v int64)`

SetDepositStatus sets DepositStatus field to given value.

### HasDepositStatus

`func (o *DepositHistoryTravelRuleResponseInner) HasDepositStatus() bool`

HasDepositStatus returns a boolean if a field has been set.

### GetTravelRuleStatus

`func (o *DepositHistoryTravelRuleResponseInner) GetTravelRuleStatus() int64`

GetTravelRuleStatus returns the TravelRuleStatus field if non-nil, zero value otherwise.

### GetTravelRuleStatusOk

`func (o *DepositHistoryTravelRuleResponseInner) GetTravelRuleStatusOk() (*int64, bool)`

GetTravelRuleStatusOk returns a tuple with the TravelRuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelRuleStatus

`func (o *DepositHistoryTravelRuleResponseInner) SetTravelRuleStatus(v int64)`

SetTravelRuleStatus sets TravelRuleStatus field to given value.

### HasTravelRuleStatus

`func (o *DepositHistoryTravelRuleResponseInner) HasTravelRuleStatus() bool`

HasTravelRuleStatus returns a boolean if a field has been set.

### GetAddress

`func (o *DepositHistoryTravelRuleResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DepositHistoryTravelRuleResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DepositHistoryTravelRuleResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DepositHistoryTravelRuleResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressTag

`func (o *DepositHistoryTravelRuleResponseInner) GetAddressTag() string`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *DepositHistoryTravelRuleResponseInner) GetAddressTagOk() (*string, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *DepositHistoryTravelRuleResponseInner) SetAddressTag(v string)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *DepositHistoryTravelRuleResponseInner) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetTxId

`func (o *DepositHistoryTravelRuleResponseInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *DepositHistoryTravelRuleResponseInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *DepositHistoryTravelRuleResponseInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *DepositHistoryTravelRuleResponseInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetInsertTime

`func (o *DepositHistoryTravelRuleResponseInner) GetInsertTime() int64`

GetInsertTime returns the InsertTime field if non-nil, zero value otherwise.

### GetInsertTimeOk

`func (o *DepositHistoryTravelRuleResponseInner) GetInsertTimeOk() (*int64, bool)`

GetInsertTimeOk returns a tuple with the InsertTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTime

`func (o *DepositHistoryTravelRuleResponseInner) SetInsertTime(v int64)`

SetInsertTime sets InsertTime field to given value.

### HasInsertTime

`func (o *DepositHistoryTravelRuleResponseInner) HasInsertTime() bool`

HasInsertTime returns a boolean if a field has been set.

### GetTransferType

`func (o *DepositHistoryTravelRuleResponseInner) GetTransferType() int64`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *DepositHistoryTravelRuleResponseInner) GetTransferTypeOk() (*int64, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *DepositHistoryTravelRuleResponseInner) SetTransferType(v int64)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *DepositHistoryTravelRuleResponseInner) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetConfirmTimes

`func (o *DepositHistoryTravelRuleResponseInner) GetConfirmTimes() string`

GetConfirmTimes returns the ConfirmTimes field if non-nil, zero value otherwise.

### GetConfirmTimesOk

`func (o *DepositHistoryTravelRuleResponseInner) GetConfirmTimesOk() (*string, bool)`

GetConfirmTimesOk returns a tuple with the ConfirmTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTimes

`func (o *DepositHistoryTravelRuleResponseInner) SetConfirmTimes(v string)`

SetConfirmTimes sets ConfirmTimes field to given value.

### HasConfirmTimes

`func (o *DepositHistoryTravelRuleResponseInner) HasConfirmTimes() bool`

HasConfirmTimes returns a boolean if a field has been set.

### GetUnlockConfirm

`func (o *DepositHistoryTravelRuleResponseInner) GetUnlockConfirm() int64`

GetUnlockConfirm returns the UnlockConfirm field if non-nil, zero value otherwise.

### GetUnlockConfirmOk

`func (o *DepositHistoryTravelRuleResponseInner) GetUnlockConfirmOk() (*int64, bool)`

GetUnlockConfirmOk returns a tuple with the UnlockConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockConfirm

`func (o *DepositHistoryTravelRuleResponseInner) SetUnlockConfirm(v int64)`

SetUnlockConfirm sets UnlockConfirm field to given value.

### HasUnlockConfirm

`func (o *DepositHistoryTravelRuleResponseInner) HasUnlockConfirm() bool`

HasUnlockConfirm returns a boolean if a field has been set.

### GetWalletType

`func (o *DepositHistoryTravelRuleResponseInner) GetWalletType() int64`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *DepositHistoryTravelRuleResponseInner) GetWalletTypeOk() (*int64, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *DepositHistoryTravelRuleResponseInner) SetWalletType(v int64)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *DepositHistoryTravelRuleResponseInner) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetRequireQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) GetRequireQuestionnaire() bool`

GetRequireQuestionnaire returns the RequireQuestionnaire field if non-nil, zero value otherwise.

### GetRequireQuestionnaireOk

`func (o *DepositHistoryTravelRuleResponseInner) GetRequireQuestionnaireOk() (*bool, bool)`

GetRequireQuestionnaireOk returns a tuple with the RequireQuestionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) SetRequireQuestionnaire(v bool)`

SetRequireQuestionnaire sets RequireQuestionnaire field to given value.

### HasRequireQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) HasRequireQuestionnaire() bool`

HasRequireQuestionnaire returns a boolean if a field has been set.

### GetQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) GetQuestionnaire() string`

GetQuestionnaire returns the Questionnaire field if non-nil, zero value otherwise.

### GetQuestionnaireOk

`func (o *DepositHistoryTravelRuleResponseInner) GetQuestionnaireOk() (*string, bool)`

GetQuestionnaireOk returns a tuple with the Questionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) SetQuestionnaire(v string)`

SetQuestionnaire sets Questionnaire field to given value.

### HasQuestionnaire

`func (o *DepositHistoryTravelRuleResponseInner) HasQuestionnaire() bool`

HasQuestionnaire returns a boolean if a field has been set.


[[Back to README]](../README.md)


