# AccountUpdateA

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**M** | Pointer to **string** |  | [optional] 
**B** | Pointer to [**[]AccountUpdateABInner**](AccountUpdateABInner.md) |  | [optional] 
**P** | Pointer to [**[]AccountUpdateAPInner**](AccountUpdateAPInner.md) |  | [optional] 

## Methods

### NewAccountUpdateA

`func NewAccountUpdateA() *AccountUpdateA`

NewAccountUpdateA instantiates a new AccountUpdateA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateAWithDefaults

`func NewAccountUpdateAWithDefaults() *AccountUpdateA`

NewAccountUpdateAWithDefaults instantiates a new AccountUpdateA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetM

`func (o *AccountUpdateA) GetM() string`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *AccountUpdateA) GetMOk() (*string, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *AccountUpdateA) SetM(v string)`

SetM sets M field to given value.

### HasM

`func (o *AccountUpdateA) HasM() bool`

HasM returns a boolean if a field has been set.

### GetB

`func (o *AccountUpdateA) GetB() []AccountUpdateABInner`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *AccountUpdateA) GetBOk() (*[]AccountUpdateABInner, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *AccountUpdateA) SetB(v []AccountUpdateABInner)`

SetB sets B field to given value.

### HasB

`func (o *AccountUpdateA) HasB() bool`

HasB returns a boolean if a field has been set.

### GetP

`func (o *AccountUpdateA) GetP() []AccountUpdateAPInner`

GetP returns the P field if non-nil, zero value otherwise.

### GetPOk

`func (o *AccountUpdateA) GetPOk() (*[]AccountUpdateAPInner, bool)`

GetPOk returns a tuple with the P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP

`func (o *AccountUpdateA) SetP(v []AccountUpdateAPInner)`

SetP sets P field to given value.

### HasP

`func (o *AccountUpdateA) HasP() bool`

HasP returns a boolean if a field has been set.


[[Back to README]](../README.md)


