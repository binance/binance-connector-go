# UserCommissionResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Commissions** | Pointer to [**[]UserCommissionResponseCommissionsInner**](UserCommissionResponseCommissionsInner.md) |  | [optional] 

## Methods

### NewUserCommissionResponse

`func NewUserCommissionResponse() *UserCommissionResponse`

NewUserCommissionResponse instantiates a new UserCommissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCommissionResponseWithDefaults

`func NewUserCommissionResponseWithDefaults() *UserCommissionResponse`

NewUserCommissionResponseWithDefaults instantiates a new UserCommissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissions

`func (o *UserCommissionResponse) GetCommissions() []UserCommissionResponseCommissionsInner`

GetCommissions returns the Commissions field if non-nil, zero value otherwise.

### GetCommissionsOk

`func (o *UserCommissionResponse) GetCommissionsOk() (*[]UserCommissionResponseCommissionsInner, bool)`

GetCommissionsOk returns a tuple with the Commissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissions

`func (o *UserCommissionResponse) SetCommissions(v []UserCommissionResponseCommissionsInner)`

SetCommissions sets Commissions field to given value.

### HasCommissions

`func (o *UserCommissionResponse) HasCommissions() bool`

HasCommissions returns a boolean if a field has been set.


[[Back to README]](../README.md)


