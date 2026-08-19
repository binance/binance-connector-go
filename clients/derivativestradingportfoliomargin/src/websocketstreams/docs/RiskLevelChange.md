# RiskLevelChange

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**U** | Pointer to **string** | uniMMR level | [optional] 
**S** | Pointer to **string** | Risk level: MARGIN_CALL, REDUCE_ONLY, FORCE_LIQUIDATION | [optional] 
**Eq** | Pointer to **string** | Account equity in USD value | [optional] 
**Ae** | Pointer to **string** | Actual equity without collateral rate in USD value | [optional] 
**M** | Pointer to **string** | Total maintenance margin in USD value | [optional] 

## Methods

### NewRiskLevelChange

`func NewRiskLevelChange() *RiskLevelChange`

NewRiskLevelChange instantiates a new RiskLevelChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLevelChangeWithDefaults

`func NewRiskLevelChangeWithDefaults() *RiskLevelChange`

NewRiskLevelChangeWithDefaults instantiates a new RiskLevelChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *RiskLevelChange) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *RiskLevelChange) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *RiskLevelChange) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *RiskLevelChange) HasE() bool`

HasE returns a boolean if a field has been set.

### GetU

`func (o *RiskLevelChange) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *RiskLevelChange) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *RiskLevelChange) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *RiskLevelChange) HasU() bool`

HasU returns a boolean if a field has been set.

### GetS

`func (o *RiskLevelChange) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *RiskLevelChange) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *RiskLevelChange) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *RiskLevelChange) HasS() bool`

HasS returns a boolean if a field has been set.

### GetEq

`func (o *RiskLevelChange) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *RiskLevelChange) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *RiskLevelChange) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *RiskLevelChange) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAe

`func (o *RiskLevelChange) GetAe() string`

GetAe returns the Ae field if non-nil, zero value otherwise.

### GetAeOk

`func (o *RiskLevelChange) GetAeOk() (*string, bool)`

GetAeOk returns a tuple with the Ae field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAe

`func (o *RiskLevelChange) SetAe(v string)`

SetAe sets Ae field to given value.

### HasAe

`func (o *RiskLevelChange) HasAe() bool`

HasAe returns a boolean if a field has been set.

### GetM

`func (o *RiskLevelChange) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *RiskLevelChange) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *RiskLevelChange) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *RiskLevelChange) HasM() bool`

HasM returns a boolean if a field has been set.


[[Back to README]](../README.md)


