# StrategyUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**Su** | Pointer to [**StrategyUpdateSu**](StrategyUpdateSu.md) |  | [optional] 

## Methods

### NewStrategyUpdate

`func NewStrategyUpdate() *StrategyUpdate`

NewStrategyUpdate instantiates a new StrategyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyUpdateWithDefaults

`func NewStrategyUpdateWithDefaults() *StrategyUpdate`

NewStrategyUpdateWithDefaults instantiates a new StrategyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *StrategyUpdate) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *StrategyUpdate) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *StrategyUpdate) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *StrategyUpdate) HasT() bool`

HasT returns a boolean if a field has been set.

### GetE

`func (o *StrategyUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *StrategyUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *StrategyUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *StrategyUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetSu

`func (o *StrategyUpdate) GetSu() StrategyUpdateSu`

GetSu returns the Su field if non-nil, zero value otherwise.

### GetSuOk

`func (o *StrategyUpdate) GetSuOk() (*StrategyUpdateSu, bool)`

GetSuOk returns a tuple with the Su field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSu

`func (o *StrategyUpdate) SetSu(v StrategyUpdateSu)`

SetSu sets Su field to given value.

### HasSu

`func (o *StrategyUpdate) HasSu() bool`

HasSu returns a boolean if a field has been set.


[[Back to README]](../README.md)


