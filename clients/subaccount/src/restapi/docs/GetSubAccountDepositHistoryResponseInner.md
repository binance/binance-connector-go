# GetSubAccountDepositHistoryResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressTag** | Pointer to **string** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**InsertTime** | Pointer to **int64** |  | [optional] 
**TransferType** | Pointer to **int64** |  | [optional] 
**ConfirmTimes** | Pointer to **string** |  | [optional] 
**UnlockConfirm** | Pointer to **int64** |  | [optional] 
**WalletType** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetSubAccountDepositHistoryResponseInner

`func NewGetSubAccountDepositHistoryResponseInner() *GetSubAccountDepositHistoryResponseInner`

NewGetSubAccountDepositHistoryResponseInner instantiates a new GetSubAccountDepositHistoryResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountDepositHistoryResponseInnerWithDefaults

`func NewGetSubAccountDepositHistoryResponseInnerWithDefaults() *GetSubAccountDepositHistoryResponseInner`

NewGetSubAccountDepositHistoryResponseInnerWithDefaults instantiates a new GetSubAccountDepositHistoryResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSubAccountDepositHistoryResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSubAccountDepositHistoryResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSubAccountDepositHistoryResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *GetSubAccountDepositHistoryResponseInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSubAccountDepositHistoryResponseInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetSubAccountDepositHistoryResponseInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCoin

`func (o *GetSubAccountDepositHistoryResponseInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *GetSubAccountDepositHistoryResponseInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *GetSubAccountDepositHistoryResponseInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetNetwork

`func (o *GetSubAccountDepositHistoryResponseInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetSubAccountDepositHistoryResponseInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetSubAccountDepositHistoryResponseInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *GetSubAccountDepositHistoryResponseInner) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSubAccountDepositHistoryResponseInner) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSubAccountDepositHistoryResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddress

`func (o *GetSubAccountDepositHistoryResponseInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetSubAccountDepositHistoryResponseInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetSubAccountDepositHistoryResponseInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressTag

`func (o *GetSubAccountDepositHistoryResponseInner) GetAddressTag() string`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetAddressTagOk() (*string, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *GetSubAccountDepositHistoryResponseInner) SetAddressTag(v string)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *GetSubAccountDepositHistoryResponseInner) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetTxId

`func (o *GetSubAccountDepositHistoryResponseInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetSubAccountDepositHistoryResponseInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *GetSubAccountDepositHistoryResponseInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetInsertTime

`func (o *GetSubAccountDepositHistoryResponseInner) GetInsertTime() int64`

GetInsertTime returns the InsertTime field if non-nil, zero value otherwise.

### GetInsertTimeOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetInsertTimeOk() (*int64, bool)`

GetInsertTimeOk returns a tuple with the InsertTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTime

`func (o *GetSubAccountDepositHistoryResponseInner) SetInsertTime(v int64)`

SetInsertTime sets InsertTime field to given value.

### HasInsertTime

`func (o *GetSubAccountDepositHistoryResponseInner) HasInsertTime() bool`

HasInsertTime returns a boolean if a field has been set.

### GetTransferType

`func (o *GetSubAccountDepositHistoryResponseInner) GetTransferType() int64`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetTransferTypeOk() (*int64, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *GetSubAccountDepositHistoryResponseInner) SetTransferType(v int64)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *GetSubAccountDepositHistoryResponseInner) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetConfirmTimes

`func (o *GetSubAccountDepositHistoryResponseInner) GetConfirmTimes() string`

GetConfirmTimes returns the ConfirmTimes field if non-nil, zero value otherwise.

### GetConfirmTimesOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetConfirmTimesOk() (*string, bool)`

GetConfirmTimesOk returns a tuple with the ConfirmTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTimes

`func (o *GetSubAccountDepositHistoryResponseInner) SetConfirmTimes(v string)`

SetConfirmTimes sets ConfirmTimes field to given value.

### HasConfirmTimes

`func (o *GetSubAccountDepositHistoryResponseInner) HasConfirmTimes() bool`

HasConfirmTimes returns a boolean if a field has been set.

### GetUnlockConfirm

`func (o *GetSubAccountDepositHistoryResponseInner) GetUnlockConfirm() int64`

GetUnlockConfirm returns the UnlockConfirm field if non-nil, zero value otherwise.

### GetUnlockConfirmOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetUnlockConfirmOk() (*int64, bool)`

GetUnlockConfirmOk returns a tuple with the UnlockConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockConfirm

`func (o *GetSubAccountDepositHistoryResponseInner) SetUnlockConfirm(v int64)`

SetUnlockConfirm sets UnlockConfirm field to given value.

### HasUnlockConfirm

`func (o *GetSubAccountDepositHistoryResponseInner) HasUnlockConfirm() bool`

HasUnlockConfirm returns a boolean if a field has been set.

### GetWalletType

`func (o *GetSubAccountDepositHistoryResponseInner) GetWalletType() int64`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *GetSubAccountDepositHistoryResponseInner) GetWalletTypeOk() (*int64, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *GetSubAccountDepositHistoryResponseInner) SetWalletType(v int64)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *GetSubAccountDepositHistoryResponseInner) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to README]](../README.md)


