# GetAutoCancelAllOpenOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Underlying** | Pointer to **string** |  | [optional] 
**CountdownTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAutoCancelAllOpenOrdersResponse

`func NewGetAutoCancelAllOpenOrdersResponse() *GetAutoCancelAllOpenOrdersResponse`

NewGetAutoCancelAllOpenOrdersResponse instantiates a new GetAutoCancelAllOpenOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutoCancelAllOpenOrdersResponseWithDefaults

`func NewGetAutoCancelAllOpenOrdersResponseWithDefaults() *GetAutoCancelAllOpenOrdersResponse`

NewGetAutoCancelAllOpenOrdersResponseWithDefaults instantiates a new GetAutoCancelAllOpenOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnderlying

`func (o *GetAutoCancelAllOpenOrdersResponse) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *GetAutoCancelAllOpenOrdersResponse) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *GetAutoCancelAllOpenOrdersResponse) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *GetAutoCancelAllOpenOrdersResponse) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetCountdownTime

`func (o *GetAutoCancelAllOpenOrdersResponse) GetCountdownTime() int64`

GetCountdownTime returns the CountdownTime field if non-nil, zero value otherwise.

### GetCountdownTimeOk

`func (o *GetAutoCancelAllOpenOrdersResponse) GetCountdownTimeOk() (*int64, bool)`

GetCountdownTimeOk returns a tuple with the CountdownTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountdownTime

`func (o *GetAutoCancelAllOpenOrdersResponse) SetCountdownTime(v int64)`

SetCountdownTime sets CountdownTime field to given value.

### HasCountdownTime

`func (o *GetAutoCancelAllOpenOrdersResponse) HasCountdownTime() bool`

HasCountdownTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


