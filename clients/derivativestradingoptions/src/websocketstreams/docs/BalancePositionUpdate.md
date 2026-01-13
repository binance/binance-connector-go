# BalancePositionUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**M** | Pointer to **string** |  | [optional] 
**B** | Pointer to [**[]BalancePositionUpdateBInner**](BalancePositionUpdateBInner.md) |  | [optional] 
**P** | Pointer to [**[]BalancePositionUpdatePInner**](BalancePositionUpdatePInner.md) |  | [optional] 

## Methods

### NewBalancePositionUpdate

`func NewBalancePositionUpdate() *BalancePositionUpdate`

NewBalancePositionUpdate instantiates a new BalancePositionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancePositionUpdateWithDefaults

`func NewBalancePositionUpdateWithDefaults() *BalancePositionUpdate`

NewBalancePositionUpdateWithDefaults instantiates a new BalancePositionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *BalancePositionUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *BalancePositionUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *BalancePositionUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *BalancePositionUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *BalancePositionUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *BalancePositionUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *BalancePositionUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *BalancePositionUpdate) HasT() bool`

HasT returns a boolean if a field has been set.

### GetM

`func (o *BalancePositionUpdate) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *BalancePositionUpdate) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *BalancePositionUpdate) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *BalancePositionUpdate) HasM() bool`

HasM returns a boolean if a field has been set.

### GetB

`func (o *BalancePositionUpdate) GetB() []BalancePositionUpdateBInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *BalancePositionUpdate) GetBOk() (*[]BalancePositionUpdateBInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *BalancePositionUpdate) SetB(v []BalancePositionUpdateBInner)`

SetB sets B field to given value.

### HasB

`func (o *BalancePositionUpdate) HasB() bool`

HasB returns a boolean if a field has been set.

### GetP

`func (o *BalancePositionUpdate) GetP() []BalancePositionUpdatePInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *BalancePositionUpdate) GetPOk() (*[]BalancePositionUpdatePInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *BalancePositionUpdate) SetP(v []BalancePositionUpdatePInner)`

SetP sets P field to given value.

### HasP

`func (o *BalancePositionUpdate) HasP() bool`

HasP returns a boolean if a field has been set.


[[Back to README]](../README.md)


