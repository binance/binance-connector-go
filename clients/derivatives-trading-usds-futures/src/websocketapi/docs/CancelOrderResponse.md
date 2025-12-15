# CancelOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**CancelOrderResponseResult**](CancelOrderResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CancelOrderResponseRateLimitsInner**](CancelOrderResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewCancelOrderResponse

`func NewCancelOrderResponse() *CancelOrderResponse`

NewCancelOrderResponse instantiates a new CancelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderResponseWithDefaults

`func NewCancelOrderResponseWithDefaults() *CancelOrderResponse`

NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CancelOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CancelOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *CancelOrderResponse) GetResult() CancelOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CancelOrderResponse) GetResultOk() (*CancelOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CancelOrderResponse) SetResult(v CancelOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CancelOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *CancelOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *CancelOrderResponse) GetRateLimitsOk() (*[]CancelOrderResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *CancelOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *CancelOrderResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


