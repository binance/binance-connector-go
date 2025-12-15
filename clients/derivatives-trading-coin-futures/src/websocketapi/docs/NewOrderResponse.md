# NewOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**NewOrderResponseResult**](NewOrderResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]CancelOrderResponseRateLimitsInner**](CancelOrderResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewNewOrderResponse

`func NewNewOrderResponse() *NewOrderResponse`

NewNewOrderResponse instantiates a new NewOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOrderResponseWithDefaults

`func NewNewOrderResponseWithDefaults() *NewOrderResponse`

NewNewOrderResponseWithDefaults instantiates a new NewOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *NewOrderResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewOrderResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewOrderResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *NewOrderResponse) GetResult() NewOrderResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *NewOrderResponse) GetResultOk() (*NewOrderResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *NewOrderResponse) SetResult(v NewOrderResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *NewOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *NewOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *NewOrderResponse) GetRateLimitsOk() (*[]CancelOrderResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *NewOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *NewOrderResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


