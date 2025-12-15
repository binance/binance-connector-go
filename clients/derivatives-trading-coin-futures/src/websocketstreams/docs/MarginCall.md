# MarginCall

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**I** | Pointer to **string** |  | [optional] 
**Cw** | Pointer to **string** |  | [optional] 
**P** | Pointer to [**[]MarginCallPInner**](MarginCallPInner.md) |  | [optional] 

## Methods

### NewMarginCall

`func NewMarginCall() *MarginCall`

NewMarginCall instantiates a new MarginCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginCallWithDefaults

`func NewMarginCallWithDefaults() *MarginCall`

NewMarginCallWithDefaults instantiates a new MarginCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *MarginCall) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *MarginCall) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *MarginCall) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *MarginCall) HasE() bool`

HasE returns a boolean if a field has been set.

### GetI

`func (o *MarginCall) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *MarginCall) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *MarginCall) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *MarginCall) HasI() bool`

HasI returns a boolean if a field has been set.

### GetCw

`func (o *MarginCall) GetCw() string`

GetCw returns the Cw field if non-nil, zero value otherwise.

### GetCwOk

`func (o *MarginCall) GetCwOk() (*string, bool)`

GetCwOk returns a tuple with the Cw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCw

`func (o *MarginCall) SetCw(v string)`

SetCw sets Cw field to given value.

### HasCw

`func (o *MarginCall) HasCw() bool`

HasCw returns a boolean if a field has been set.

### GetP

`func (o *MarginCall) GetP() []MarginCallPInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *MarginCall) GetPOk() (*[]MarginCallPInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *MarginCall) SetP(v []MarginCallPInner)`

SetP sets P field to given value.

### HasP

`func (o *MarginCall) HasP() bool`

HasP returns a boolean if a field has been set.


[[Back to README]](../README.md)


