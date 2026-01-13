# SubAccountTransferHistoryResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CounterParty** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int64** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**FromAccountType** | Pointer to **string** |  | [optional] 
**ToAccountType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TranId** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubAccountTransferHistoryResponseInner

`func NewSubAccountTransferHistoryResponseInner() *SubAccountTransferHistoryResponseInner`

NewSubAccountTransferHistoryResponseInner instantiates a new SubAccountTransferHistoryResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubAccountTransferHistoryResponseInnerWithDefaults

`func NewSubAccountTransferHistoryResponseInnerWithDefaults() *SubAccountTransferHistoryResponseInner`

NewSubAccountTransferHistoryResponseInnerWithDefaults instantiates a new SubAccountTransferHistoryResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterParty

`func (o *SubAccountTransferHistoryResponseInner) GetCounterParty() string`

GetCounterParty returns the CounterParty field if non-nil, zero value otherwise.

### GetCounterPartyOk

`func (o *SubAccountTransferHistoryResponseInner) GetCounterPartyOk() (*string, bool)`

GetCounterPartyOk returns a tuple with the CounterParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterParty

`func (o *SubAccountTransferHistoryResponseInner) SetCounterParty(v string)`

SetCounterParty sets CounterParty field to given value.

### HasCounterParty

`func (o *SubAccountTransferHistoryResponseInner) HasCounterParty() bool`

HasCounterParty returns a boolean if a field has been set.

### GetEmail

`func (o *SubAccountTransferHistoryResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubAccountTransferHistoryResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubAccountTransferHistoryResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubAccountTransferHistoryResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *SubAccountTransferHistoryResponseInner) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubAccountTransferHistoryResponseInner) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubAccountTransferHistoryResponseInner) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *SubAccountTransferHistoryResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAsset

`func (o *SubAccountTransferHistoryResponseInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SubAccountTransferHistoryResponseInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SubAccountTransferHistoryResponseInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *SubAccountTransferHistoryResponseInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetQty

`func (o *SubAccountTransferHistoryResponseInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *SubAccountTransferHistoryResponseInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *SubAccountTransferHistoryResponseInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *SubAccountTransferHistoryResponseInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetFromAccountType

`func (o *SubAccountTransferHistoryResponseInner) GetFromAccountType() string`

GetFromAccountType returns the FromAccountType field if non-nil, zero value otherwise.

### GetFromAccountTypeOk

`func (o *SubAccountTransferHistoryResponseInner) GetFromAccountTypeOk() (*string, bool)`

GetFromAccountTypeOk returns a tuple with the FromAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccountType

`func (o *SubAccountTransferHistoryResponseInner) SetFromAccountType(v string)`

SetFromAccountType sets FromAccountType field to given value.

### HasFromAccountType

`func (o *SubAccountTransferHistoryResponseInner) HasFromAccountType() bool`

HasFromAccountType returns a boolean if a field has been set.

### GetToAccountType

`func (o *SubAccountTransferHistoryResponseInner) GetToAccountType() string`

GetToAccountType returns the ToAccountType field if non-nil, zero value otherwise.

### GetToAccountTypeOk

`func (o *SubAccountTransferHistoryResponseInner) GetToAccountTypeOk() (*string, bool)`

GetToAccountTypeOk returns a tuple with the ToAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountType

`func (o *SubAccountTransferHistoryResponseInner) SetToAccountType(v string)`

SetToAccountType sets ToAccountType field to given value.

### HasToAccountType

`func (o *SubAccountTransferHistoryResponseInner) HasToAccountType() bool`

HasToAccountType returns a boolean if a field has been set.

### GetStatus

`func (o *SubAccountTransferHistoryResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubAccountTransferHistoryResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubAccountTransferHistoryResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubAccountTransferHistoryResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTranId

`func (o *SubAccountTransferHistoryResponseInner) GetTranId() int64`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *SubAccountTransferHistoryResponseInner) GetTranIdOk() (*int64, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *SubAccountTransferHistoryResponseInner) SetTranId(v int64)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *SubAccountTransferHistoryResponseInner) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTime

`func (o *SubAccountTransferHistoryResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SubAccountTransferHistoryResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SubAccountTransferHistoryResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SubAccountTransferHistoryResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


