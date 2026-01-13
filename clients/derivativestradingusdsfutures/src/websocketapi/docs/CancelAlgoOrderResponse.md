# CancelAlgoOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**CancelAlgoOrderResponseResult**](CancelAlgoOrderResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CancelAlgoOrderResponseRateLimitsInner**](CancelAlgoOrderResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewCancelAlgoOrderResponse

`func NewCancelAlgoOrderResponse() *CancelAlgoOrderResponse`

NewCancelAlgoOrderResponse instantiates a new CancelAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelAlgoOrderResponseWithDefaults

`func NewCancelAlgoOrderResponseWithDefaults() *CancelAlgoOrderResponse`

NewCancelAlgoOrderResponseWithDefaults instantiates a new CancelAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelAlgoOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelAlgoOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelAlgoOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CancelAlgoOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CancelAlgoOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelAlgoOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelAlgoOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelAlgoOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *CancelAlgoOrderResponse) GetResult() CancelAlgoOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CancelAlgoOrderResponse) GetResultOk() (*CancelAlgoOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CancelAlgoOrderResponse) SetResult(v CancelAlgoOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CancelAlgoOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *CancelAlgoOrderResponse) GetRateLimits() []CancelAlgoOrderResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *CancelAlgoOrderResponse) GetRateLimitsOk() (*[]CancelAlgoOrderResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *CancelAlgoOrderResponse) SetRateLimits(v []CancelAlgoOrderResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *CancelAlgoOrderResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


