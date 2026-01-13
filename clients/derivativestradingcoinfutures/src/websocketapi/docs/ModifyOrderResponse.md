# ModifyOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**ModifyOrderResponseResult**](ModifyOrderResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CancelOrderResponseRateLimitsInner**](CancelOrderResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewModifyOrderResponse

`func NewModifyOrderResponse() *ModifyOrderResponse`

NewModifyOrderResponse instantiates a new ModifyOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyOrderResponseWithDefaults

`func NewModifyOrderResponseWithDefaults() *ModifyOrderResponse`

NewModifyOrderResponseWithDefaults instantiates a new ModifyOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModifyOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModifyOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ModifyOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModifyOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModifyOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModifyOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *ModifyOrderResponse) GetResult() ModifyOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ModifyOrderResponse) GetResultOk() (*ModifyOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ModifyOrderResponse) SetResult(v ModifyOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ModifyOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *ModifyOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ModifyOrderResponse) GetRateLimitsOk() (*[]CancelOrderResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ModifyOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ModifyOrderResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


