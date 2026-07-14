# PmProAccountUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**U** | Pointer to **string** | uniMMR level | [optional] 
**Eq** | Pointer to **string** | Account equity in USD | [optional] 
**Ae** | Pointer to **string** | Actual equity without collateral rate in USD | [optional] 
**Im** | Pointer to **string** | Total initial margin in USD | [optional] 
**Mm** | Pointer to **string** | Total maintenance margin in USD | [optional] 
**Avb** | Pointer to **string** | Total available balance in USD | [optional] 
**Vmw** | Pointer to **string** | Virtual maxWithdraw amount in USD | [optional] 

## Methods

### NewPmProAccountUpdate

`func NewPmProAccountUpdate() *PmProAccountUpdate`

NewPmProAccountUpdate instantiates a new PmProAccountUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmProAccountUpdateWithDefaults

`func NewPmProAccountUpdateWithDefaults() *PmProAccountUpdate`

NewPmProAccountUpdateWithDefaults instantiates a new PmProAccountUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *PmProAccountUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *PmProAccountUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *PmProAccountUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *PmProAccountUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetU

`func (o *PmProAccountUpdate) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *PmProAccountUpdate) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *PmProAccountUpdate) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *PmProAccountUpdate) HasU() bool`

HasU returns a boolean if a field has been set.

### GetEq

`func (o *PmProAccountUpdate) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *PmProAccountUpdate) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *PmProAccountUpdate) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *PmProAccountUpdate) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAe

`func (o *PmProAccountUpdate) GetAe() string`

GetAe returns the Ae field if non-nil, zero value otherwise.

### GetAeOk

`func (o *PmProAccountUpdate) GetAeOk() (*string, bool)`

GetAeOk returns a tuple with the Ae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAe

`func (o *PmProAccountUpdate) SetAe(v string)`

SetAe sets Ae field to given value.

### HasAe

`func (o *PmProAccountUpdate) HasAe() bool`

HasAe returns a boolean if a field has been set.

### GetIm

`func (o *PmProAccountUpdate) GetIm() string`

GetIm returns the Im field if non-nil, zero value otherwise.

### GetImOk

`func (o *PmProAccountUpdate) GetImOk() (*string, bool)`

GetImOk returns a tuple with the Im field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIm

`func (o *PmProAccountUpdate) SetIm(v string)`

SetIm sets Im field to given value.

### HasIm

`func (o *PmProAccountUpdate) HasIm() bool`

HasIm returns a boolean if a field has been set.

### GetMm

`func (o *PmProAccountUpdate) GetMm() string`

GetMm returns the Mm field if non-nil, zero value otherwise.

### GetMmOk

`func (o *PmProAccountUpdate) GetMmOk() (*string, bool)`

GetMmOk returns a tuple with the Mm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMm

`func (o *PmProAccountUpdate) SetMm(v string)`

SetMm sets Mm field to given value.

### HasMm

`func (o *PmProAccountUpdate) HasMm() bool`

HasMm returns a boolean if a field has been set.

### GetAvb

`func (o *PmProAccountUpdate) GetAvb() string`

GetAvb returns the Avb field if non-nil, zero value otherwise.

### GetAvbOk

`func (o *PmProAccountUpdate) GetAvbOk() (*string, bool)`

GetAvbOk returns a tuple with the Avb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvb

`func (o *PmProAccountUpdate) SetAvb(v string)`

SetAvb sets Avb field to given value.

### HasAvb

`func (o *PmProAccountUpdate) HasAvb() bool`

HasAvb returns a boolean if a field has been set.

### GetVmw

`func (o *PmProAccountUpdate) GetVmw() string`

GetVmw returns the Vmw field if non-nil, zero value otherwise.

### GetVmwOk

`func (o *PmProAccountUpdate) GetVmwOk() (*string, bool)`

GetVmwOk returns a tuple with the Vmw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmw

`func (o *PmProAccountUpdate) SetVmw(v string)`

SetVmw sets Vmw field to given value.

### HasVmw

`func (o *PmProAccountUpdate) HasVmw() bool`

HasVmw returns a boolean if a field has been set.


[[Back to README]](../README.md)


