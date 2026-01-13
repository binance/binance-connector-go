# NewAlgoOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**NewAlgoOrderResponseResult**](NewAlgoOrderResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CancelOrderResponseRateLimitsInner**](CancelOrderResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewNewAlgoOrderResponse

`func NewNewAlgoOrderResponse() *NewAlgoOrderResponse`

NewNewAlgoOrderResponse instantiates a new NewAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAlgoOrderResponseWithDefaults

`func NewNewAlgoOrderResponseWithDefaults() *NewAlgoOrderResponse`

NewNewAlgoOrderResponseWithDefaults instantiates a new NewAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewAlgoOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewAlgoOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewAlgoOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewAlgoOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *NewAlgoOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewAlgoOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewAlgoOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewAlgoOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *NewAlgoOrderResponse) GetResult() NewAlgoOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *NewAlgoOrderResponse) GetResultOk() (*NewAlgoOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *NewAlgoOrderResponse) SetResult(v NewAlgoOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *NewAlgoOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *NewAlgoOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *NewAlgoOrderResponse) GetRateLimitsOk() (*[]CancelOrderResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *NewAlgoOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *NewAlgoOrderResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


