# UserDataStreamEventsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**U** | Pointer to **string** | uniMMR level | [optional] 
**Eq** | Pointer to **string** | Account equity in USD value | [optional] 
**Ae** | Pointer to **string** | Actual equity without collateral rate in USD value | [optional] 
**Im** | Pointer to **string** | Total initial margin in USD | [optional] 
**Mm** | Pointer to **string** | Total maintenance margin in USD | [optional] 
**Avb** | Pointer to **string** | Total available balance in USD | [optional] 
**Vmw** | Pointer to **string** | Virtual maxWithdraw amount in USD | [optional] 
**S** | Pointer to **string** | Risk level: MARGIN_CALL, REDUCE_ONLY, FORCE_LIQUIDATION | [optional] 
**M** | Pointer to **string** | Total maintenance margin in USD value | [optional] 

## Methods

### NewUserDataStreamEventsResponse

`func NewUserDataStreamEventsResponse() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponse instantiates a new UserDataStreamEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataStreamEventsResponseWithDefaults

`func NewUserDataStreamEventsResponseWithDefaults() *UserDataStreamEventsResponse`

NewUserDataStreamEventsResponseWithDefaults instantiates a new UserDataStreamEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *UserDataStreamEventsResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *UserDataStreamEventsResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *UserDataStreamEventsResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *UserDataStreamEventsResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetU

`func (o *UserDataStreamEventsResponse) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *UserDataStreamEventsResponse) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *UserDataStreamEventsResponse) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *UserDataStreamEventsResponse) HasU() bool`

HasU returns a boolean if a field has been set.

### GetEq

`func (o *UserDataStreamEventsResponse) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *UserDataStreamEventsResponse) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *UserDataStreamEventsResponse) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *UserDataStreamEventsResponse) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAe

`func (o *UserDataStreamEventsResponse) GetAe() string`

GetAe returns the Ae field if non-nil, zero value otherwise.

### GetAeOk

`func (o *UserDataStreamEventsResponse) GetAeOk() (*string, bool)`

GetAeOk returns a tuple with the Ae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAe

`func (o *UserDataStreamEventsResponse) SetAe(v string)`

SetAe sets Ae field to given value.

### HasAe

`func (o *UserDataStreamEventsResponse) HasAe() bool`

HasAe returns a boolean if a field has been set.

### GetIm

`func (o *UserDataStreamEventsResponse) GetIm() string`

GetIm returns the Im field if non-nil, zero value otherwise.

### GetImOk

`func (o *UserDataStreamEventsResponse) GetImOk() (*string, bool)`

GetImOk returns a tuple with the Im field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIm

`func (o *UserDataStreamEventsResponse) SetIm(v string)`

SetIm sets Im field to given value.

### HasIm

`func (o *UserDataStreamEventsResponse) HasIm() bool`

HasIm returns a boolean if a field has been set.

### GetMm

`func (o *UserDataStreamEventsResponse) GetMm() string`

GetMm returns the Mm field if non-nil, zero value otherwise.

### GetMmOk

`func (o *UserDataStreamEventsResponse) GetMmOk() (*string, bool)`

GetMmOk returns a tuple with the Mm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMm

`func (o *UserDataStreamEventsResponse) SetMm(v string)`

SetMm sets Mm field to given value.

### HasMm

`func (o *UserDataStreamEventsResponse) HasMm() bool`

HasMm returns a boolean if a field has been set.

### GetAvb

`func (o *UserDataStreamEventsResponse) GetAvb() string`

GetAvb returns the Avb field if non-nil, zero value otherwise.

### GetAvbOk

`func (o *UserDataStreamEventsResponse) GetAvbOk() (*string, bool)`

GetAvbOk returns a tuple with the Avb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvb

`func (o *UserDataStreamEventsResponse) SetAvb(v string)`

SetAvb sets Avb field to given value.

### HasAvb

`func (o *UserDataStreamEventsResponse) HasAvb() bool`

HasAvb returns a boolean if a field has been set.

### GetVmw

`func (o *UserDataStreamEventsResponse) GetVmw() string`

GetVmw returns the Vmw field if non-nil, zero value otherwise.

### GetVmwOk

`func (o *UserDataStreamEventsResponse) GetVmwOk() (*string, bool)`

GetVmwOk returns a tuple with the Vmw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmw

`func (o *UserDataStreamEventsResponse) SetVmw(v string)`

SetVmw sets Vmw field to given value.

### HasVmw

`func (o *UserDataStreamEventsResponse) HasVmw() bool`

HasVmw returns a boolean if a field has been set.

### GetS

`func (o *UserDataStreamEventsResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *UserDataStreamEventsResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *UserDataStreamEventsResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *UserDataStreamEventsResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetM

`func (o *UserDataStreamEventsResponse) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *UserDataStreamEventsResponse) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *UserDataStreamEventsResponse) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *UserDataStreamEventsResponse) HasM() bool`

HasM returns a boolean if a field has been set.


[[Back to README]](../README.md)


