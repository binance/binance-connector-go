# SetAutoCancelAllOpenOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Underlying** | Pointer to **string** |  | [optional] 
**CountdownTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSetAutoCancelAllOpenOrdersResponse

`func NewSetAutoCancelAllOpenOrdersResponse() *SetAutoCancelAllOpenOrdersResponse`

NewSetAutoCancelAllOpenOrdersResponse instantiates a new SetAutoCancelAllOpenOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAutoCancelAllOpenOrdersResponseWithDefaults

`func NewSetAutoCancelAllOpenOrdersResponseWithDefaults() *SetAutoCancelAllOpenOrdersResponse`

NewSetAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new SetAutoCancelAllOpenOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderlying

`func (o *SetAutoCancelAllOpenOrdersResponse) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *SetAutoCancelAllOpenOrdersResponse) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *SetAutoCancelAllOpenOrdersResponse) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *SetAutoCancelAllOpenOrdersResponse) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetCountdownTime

`func (o *SetAutoCancelAllOpenOrdersResponse) GetCountdownTime() int64`

GetCountdownTime returns the CountdownTime field if non-nil, zero value otherwise.

### GetCountdownTimeOk

`func (o *SetAutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*int64, bool)`

GetCountdownTimeOk returns a tuple with the CountdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdownTime

`func (o *SetAutoCancelAllOpenOrdersResponse) SetCountdownTime(v int64)`

SetCountdownTime sets CountdownTime field to given value.

### HasCountdownTime

`func (o *SetAutoCancelAllOpenOrdersResponse) HasCountdownTime() bool`

HasCountdownTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


