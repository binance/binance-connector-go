# WithdrawHistoryV2ResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TrId** | Pointer to **int64** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**TransactionFee** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**WithdrawalStatus** | Pointer to **int64** |  | [optional] 
**TravelRuleStatus** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**ApplyTime** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**TransferType** | Pointer to **int64** |  | [optional] 
**WithdrawOrderId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**ConfirmNo** | Pointer to **int64** |  | [optional] 
**WalletType** | Pointer to **int64** |  | [optional] 
**TxKey** | Pointer to **string** |  | [optional] 
**Questionnaire** | Pointer to **string** |  | [optional] 
**CompleteTime** | Pointer to **string** |  | [optional] 

## Methods

### NewWithdrawHistoryV2ResponseInner

`func NewWithdrawHistoryV2ResponseInner() *WithdrawHistoryV2ResponseInner`

NewWithdrawHistoryV2ResponseInner instantiates a new WithdrawHistoryV2ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawHistoryV2ResponseInnerWithDefaults

`func NewWithdrawHistoryV2ResponseInnerWithDefaults() *WithdrawHistoryV2ResponseInner`

NewWithdrawHistoryV2ResponseInnerWithDefaults instantiates a new WithdrawHistoryV2ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WithdrawHistoryV2ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WithdrawHistoryV2ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WithdrawHistoryV2ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WithdrawHistoryV2ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrId

`func (o *WithdrawHistoryV2ResponseInner) GetTrId() int64`

GetTrId returns the TrId field if non-nil, zero value otherwise.

### GetTrIdOk

`func (o *WithdrawHistoryV2ResponseInner) GetTrIdOk() (*int64, bool)`

GetTrIdOk returns a tuple with the TrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrId

`func (o *WithdrawHistoryV2ResponseInner) SetTrId(v int64)`

SetTrId sets TrId field to given value.

### HasTrId

`func (o *WithdrawHistoryV2ResponseInner) HasTrId() bool`

HasTrId returns a boolean if a field has been set.

### GetAmount

`func (o *WithdrawHistoryV2ResponseInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawHistoryV2ResponseInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawHistoryV2ResponseInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WithdrawHistoryV2ResponseInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionFee

`func (o *WithdrawHistoryV2ResponseInner) GetTransactionFee() string`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *WithdrawHistoryV2ResponseInner) GetTransactionFeeOk() (*string, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *WithdrawHistoryV2ResponseInner) SetTransactionFee(v string)`

SetTransactionFee sets TransactionFee field to given value.

### HasTransactionFee

`func (o *WithdrawHistoryV2ResponseInner) HasTransactionFee() bool`

HasTransactionFee returns a boolean if a field has been set.

### GetCoin

`func (o *WithdrawHistoryV2ResponseInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *WithdrawHistoryV2ResponseInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *WithdrawHistoryV2ResponseInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *WithdrawHistoryV2ResponseInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetWithdrawalStatus

`func (o *WithdrawHistoryV2ResponseInner) GetWithdrawalStatus() int64`

GetWithdrawalStatus returns the WithdrawalStatus field if non-nil, zero value otherwise.

### GetWithdrawalStatusOk

`func (o *WithdrawHistoryV2ResponseInner) GetWithdrawalStatusOk() (*int64, bool)`

GetWithdrawalStatusOk returns a tuple with the WithdrawalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalStatus

`func (o *WithdrawHistoryV2ResponseInner) SetWithdrawalStatus(v int64)`

SetWithdrawalStatus sets WithdrawalStatus field to given value.

### HasWithdrawalStatus

`func (o *WithdrawHistoryV2ResponseInner) HasWithdrawalStatus() bool`

HasWithdrawalStatus returns a boolean if a field has been set.

### GetTravelRuleStatus

`func (o *WithdrawHistoryV2ResponseInner) GetTravelRuleStatus() int64`

GetTravelRuleStatus returns the TravelRuleStatus field if non-nil, zero value otherwise.

### GetTravelRuleStatusOk

`func (o *WithdrawHistoryV2ResponseInner) GetTravelRuleStatusOk() (*int64, bool)`

GetTravelRuleStatusOk returns a tuple with the TravelRuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelRuleStatus

`func (o *WithdrawHistoryV2ResponseInner) SetTravelRuleStatus(v int64)`

SetTravelRuleStatus sets TravelRuleStatus field to given value.

### HasTravelRuleStatus

`func (o *WithdrawHistoryV2ResponseInner) HasTravelRuleStatus() bool`

HasTravelRuleStatus returns a boolean if a field has been set.

### GetAddress

`func (o *WithdrawHistoryV2ResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WithdrawHistoryV2ResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WithdrawHistoryV2ResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WithdrawHistoryV2ResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTxId

`func (o *WithdrawHistoryV2ResponseInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *WithdrawHistoryV2ResponseInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *WithdrawHistoryV2ResponseInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *WithdrawHistoryV2ResponseInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetApplyTime

`func (o *WithdrawHistoryV2ResponseInner) GetApplyTime() string`

GetApplyTime returns the ApplyTime field if non-nil, zero value otherwise.

### GetApplyTimeOk

`func (o *WithdrawHistoryV2ResponseInner) GetApplyTimeOk() (*string, bool)`

GetApplyTimeOk returns a tuple with the ApplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTime

`func (o *WithdrawHistoryV2ResponseInner) SetApplyTime(v string)`

SetApplyTime sets ApplyTime field to given value.

### HasApplyTime

`func (o *WithdrawHistoryV2ResponseInner) HasApplyTime() bool`

HasApplyTime returns a boolean if a field has been set.

### GetNetwork

`func (o *WithdrawHistoryV2ResponseInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WithdrawHistoryV2ResponseInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WithdrawHistoryV2ResponseInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WithdrawHistoryV2ResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTransferType

`func (o *WithdrawHistoryV2ResponseInner) GetTransferType() int64`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *WithdrawHistoryV2ResponseInner) GetTransferTypeOk() (*int64, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *WithdrawHistoryV2ResponseInner) SetTransferType(v int64)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *WithdrawHistoryV2ResponseInner) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetWithdrawOrderId

`func (o *WithdrawHistoryV2ResponseInner) GetWithdrawOrderId() string`

GetWithdrawOrderId returns the WithdrawOrderId field if non-nil, zero value otherwise.

### GetWithdrawOrderIdOk

`func (o *WithdrawHistoryV2ResponseInner) GetWithdrawOrderIdOk() (*string, bool)`

GetWithdrawOrderIdOk returns a tuple with the WithdrawOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawOrderId

`func (o *WithdrawHistoryV2ResponseInner) SetWithdrawOrderId(v string)`

SetWithdrawOrderId sets WithdrawOrderId field to given value.

### HasWithdrawOrderId

`func (o *WithdrawHistoryV2ResponseInner) HasWithdrawOrderId() bool`

HasWithdrawOrderId returns a boolean if a field has been set.

### GetInfo

`func (o *WithdrawHistoryV2ResponseInner) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WithdrawHistoryV2ResponseInner) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WithdrawHistoryV2ResponseInner) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WithdrawHistoryV2ResponseInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetConfirmNo

`func (o *WithdrawHistoryV2ResponseInner) GetConfirmNo() int64`

GetConfirmNo returns the ConfirmNo field if non-nil, zero value otherwise.

### GetConfirmNoOk

`func (o *WithdrawHistoryV2ResponseInner) GetConfirmNoOk() (*int64, bool)`

GetConfirmNoOk returns a tuple with the ConfirmNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmNo

`func (o *WithdrawHistoryV2ResponseInner) SetConfirmNo(v int64)`

SetConfirmNo sets ConfirmNo field to given value.

### HasConfirmNo

`func (o *WithdrawHistoryV2ResponseInner) HasConfirmNo() bool`

HasConfirmNo returns a boolean if a field has been set.

### GetWalletType

`func (o *WithdrawHistoryV2ResponseInner) GetWalletType() int64`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *WithdrawHistoryV2ResponseInner) GetWalletTypeOk() (*int64, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *WithdrawHistoryV2ResponseInner) SetWalletType(v int64)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *WithdrawHistoryV2ResponseInner) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetTxKey

`func (o *WithdrawHistoryV2ResponseInner) GetTxKey() string`

GetTxKey returns the TxKey field if non-nil, zero value otherwise.

### GetTxKeyOk

`func (o *WithdrawHistoryV2ResponseInner) GetTxKeyOk() (*string, bool)`

GetTxKeyOk returns a tuple with the TxKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxKey

`func (o *WithdrawHistoryV2ResponseInner) SetTxKey(v string)`

SetTxKey sets TxKey field to given value.

### HasTxKey

`func (o *WithdrawHistoryV2ResponseInner) HasTxKey() bool`

HasTxKey returns a boolean if a field has been set.

### GetQuestionnaire

`func (o *WithdrawHistoryV2ResponseInner) GetQuestionnaire() string`

GetQuestionnaire returns the Questionnaire field if non-nil, zero value otherwise.

### GetQuestionnaireOk

`func (o *WithdrawHistoryV2ResponseInner) GetQuestionnaireOk() (*string, bool)`

GetQuestionnaireOk returns a tuple with the Questionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaire

`func (o *WithdrawHistoryV2ResponseInner) SetQuestionnaire(v string)`

SetQuestionnaire sets Questionnaire field to given value.

### HasQuestionnaire

`func (o *WithdrawHistoryV2ResponseInner) HasQuestionnaire() bool`

HasQuestionnaire returns a boolean if a field has been set.

### GetCompleteTime

`func (o *WithdrawHistoryV2ResponseInner) GetCompleteTime() string`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *WithdrawHistoryV2ResponseInner) GetCompleteTimeOk() (*string, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *WithdrawHistoryV2ResponseInner) SetCompleteTime(v string)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *WithdrawHistoryV2ResponseInner) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


