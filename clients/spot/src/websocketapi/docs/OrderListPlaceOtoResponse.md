# OrderListPlaceOtoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**OrderListPlaceOtoResponseResult**](OrderListPlaceOtoResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOrderListPlaceOtoResponse

`func NewOrderListPlaceOtoResponse() *OrderListPlaceOtoResponse`

NewOrderListPlaceOtoResponse instantiates a new OrderListPlaceOtoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOtoResponseWithDefaults

`func NewOrderListPlaceOtoResponseWithDefaults() *OrderListPlaceOtoResponse`

NewOrderListPlaceOtoResponseWithDefaults instantiates a new OrderListPlaceOtoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderListPlaceOtoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderListPlaceOtoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderListPlaceOtoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderListPlaceOtoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderListPlaceOtoResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListPlaceOtoResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListPlaceOtoResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderListPlaceOtoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OrderListPlaceOtoResponse) GetResult() OrderListPlaceOtoResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderListPlaceOtoResponse) GetResultOk() (*OrderListPlaceOtoResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderListPlaceOtoResponse) SetResult(v OrderListPlaceOtoResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderListPlaceOtoResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OrderListPlaceOtoResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OrderListPlaceOtoResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OrderListPlaceOtoResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OrderListPlaceOtoResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


