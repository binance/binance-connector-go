# AccountUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**B** | Pointer to [**[]AccountUpdateBInner**](AccountUpdateBInner.md) |  | [optional] 
**G** | Pointer to [**[]AccountUpdateGInner**](AccountUpdateGInner.md) |  | [optional] 
**P** | Pointer to [**[]AccountUpdatePInner**](AccountUpdatePInner.md) |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountUpdate

`func NewAccountUpdate() *AccountUpdate`

NewAccountUpdate instantiates a new AccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateWithDefaults

`func NewAccountUpdateWithDefaults() *AccountUpdate`

NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AccountUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AccountUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AccountUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *AccountUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetB

`func (o *AccountUpdate) GetB() []AccountUpdateBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *AccountUpdate) GetBOk() (*[]AccountUpdateBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *AccountUpdate) SetB(v []AccountUpdateBInner)`

SetB sets B field to given value.

### HasB

`func (o *AccountUpdate) HasB() bool`

HasB returns a boolean if a field has been set.

### GetG

`func (o *AccountUpdate) GetG() []AccountUpdateGInner`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *AccountUpdate) GetGOk() (*[]AccountUpdateGInner, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *AccountUpdate) SetG(v []AccountUpdateGInner)`

SetG sets G field to given value.

### HasG

`func (o *AccountUpdate) HasG() bool`

HasG returns a boolean if a field has been set.

### GetP

`func (o *AccountUpdate) GetP() []AccountUpdatePInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AccountUpdate) GetPOk() (*[]AccountUpdatePInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AccountUpdate) SetP(v []AccountUpdatePInner)`

SetP sets P field to given value.

### HasP

`func (o *AccountUpdate) HasP() bool`

HasP returns a boolean if a field has been set.

### GetUid

`func (o *AccountUpdate) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AccountUpdate) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AccountUpdate) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AccountUpdate) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to README]](../README.md)


