# AccountUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**A** | Pointer to [**AccountUpdateA**](AccountUpdateA.md) |  | [optional] 

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

### GetT

`func (o *AccountUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AccountUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AccountUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *AccountUpdate) HasT() bool`

HasT returns a boolean if a field has been set.

### GetA

`func (o *AccountUpdate) GetA() AccountUpdateA`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *AccountUpdate) GetAOk() (*AccountUpdateA, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *AccountUpdate) SetA(v AccountUpdateA)`

SetA sets A field to given value.

### HasA

`func (o *AccountUpdate) HasA() bool`

HasA returns a boolean if a field has been set.


[[Back to README]](../README.md)


