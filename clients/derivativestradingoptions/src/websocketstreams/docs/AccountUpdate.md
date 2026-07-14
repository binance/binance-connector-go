# AccountUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** | Event Time | [optional] 
**T** | Pointer to **int64** | Transaction Time | [optional] 
**Eq** | Pointer to **string** | Account equity in USDT | [optional] 
**Aeq** | Pointer to **string** | Account adjusted equity in USDT | [optional] 
**B** | Pointer to **string** | Account wallet balance in USDT | [optional] 
**M** | Pointer to **string** | Position value | [optional] 
**U** | Pointer to **string** | Unrealized PnL | [optional] 
**I** | Pointer to **string** | Initial margin in USDT | [optional] 
**M** | Pointer to **string** | Maintenance margin in USDT | [optional] 

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

### GetEq

`func (o *AccountUpdate) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *AccountUpdate) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *AccountUpdate) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *AccountUpdate) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetAeq

`func (o *AccountUpdate) GetAeq() string`

GetAeq returns the Aeq field if non-nil, zero value otherwise.

### GetAeqOk

`func (o *AccountUpdate) GetAeqOk() (*string, bool)`

GetAeqOk returns a tuple with the Aeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAeq

`func (o *AccountUpdate) SetAeq(v string)`

SetAeq sets Aeq field to given value.

### HasAeq

`func (o *AccountUpdate) HasAeq() bool`

HasAeq returns a boolean if a field has been set.

### GetB

`func (o *AccountUpdate) GetB() string`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *AccountUpdate) GetBOk() (*string, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *AccountUpdate) SetB(v string)`

SetB sets B field to given value.

### HasB

`func (o *AccountUpdate) HasB() bool`

HasB returns a boolean if a field has been set.

### GetM

`func (o *AccountUpdate) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AccountUpdate) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AccountUpdate) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *AccountUpdate) HasM() bool`

HasM returns a boolean if a field has been set.

### GetU

`func (o *AccountUpdate) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *AccountUpdate) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *AccountUpdate) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *AccountUpdate) HasU() bool`

HasU returns a boolean if a field has been set.

### GetI

`func (o *AccountUpdate) GetI() string`

GetI returns the I field if non-nil, zero value otherwise.

### GetIOk

`func (o *AccountUpdate) GetIOk() (*string, bool)`

GetIOk returns a tuple with the I field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI

`func (o *AccountUpdate) SetI(v string)`

SetI sets I field to given value.

### HasI

`func (o *AccountUpdate) HasI() bool`

HasI returns a boolean if a field has been set.

### GetM

`func (o *AccountUpdate) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AccountUpdate) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AccountUpdate) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *AccountUpdate) HasM() bool`

HasM returns a boolean if a field has been set.


[[Back to README]](../README.md)


