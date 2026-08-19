# BalanceUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**A** | Pointer to **string** | Asset | [optional] 
**D** | Pointer to **string** | Balance Delta | [optional] 
**U** | Pointer to **int64** | Event updateId | [optional] 
**T** | Pointer to **int64** | Clear Time | [optional] 

## Methods

### NewBalanceUpdate

`func NewBalanceUpdate() *BalanceUpdate`

NewBalanceUpdate instantiates a new BalanceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceUpdateWithDefaults

`func NewBalanceUpdateWithDefaults() *BalanceUpdate`

NewBalanceUpdateWithDefaults instantiates a new BalanceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *BalanceUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *BalanceUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *BalanceUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *BalanceUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetA

`func (o *BalanceUpdate) GetA() string`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *BalanceUpdate) GetAOk() (*string, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *BalanceUpdate) SetA(v string)`

SetA sets A field to given value.

### HasA

`func (o *BalanceUpdate) HasA() bool`

HasA returns a boolean if a field has been set.

### GetD

`func (o *BalanceUpdate) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *BalanceUpdate) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *BalanceUpdate) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *BalanceUpdate) HasD() bool`

HasD returns a boolean if a field has been set.

### GetU

`func (o *BalanceUpdate) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *BalanceUpdate) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *BalanceUpdate) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *BalanceUpdate) HasU() bool`

HasU returns a boolean if a field has been set.

### GetT

`func (o *BalanceUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *BalanceUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *BalanceUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *BalanceUpdate) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to README]](../README.md)


