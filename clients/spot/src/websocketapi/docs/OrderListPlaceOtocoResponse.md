# OrderListPlaceOtocoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**OrderListPlaceOtocoResponseResult**](OrderListPlaceOtocoResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]RateLimits**](RateLimits.md) |  | [optional] 

## Methods

### NewOrderListPlaceOtocoResponse

`func NewOrderListPlaceOtocoResponse() *OrderListPlaceOtocoResponse`

NewOrderListPlaceOtocoResponse instantiates a new OrderListPlaceOtocoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOtocoResponseWithDefaults

`func NewOrderListPlaceOtocoResponseWithDefaults() *OrderListPlaceOtocoResponse`

NewOrderListPlaceOtocoResponseWithDefaults instantiates a new OrderListPlaceOtocoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderListPlaceOtocoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderListPlaceOtocoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderListPlaceOtocoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderListPlaceOtocoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderListPlaceOtocoResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListPlaceOtocoResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListPlaceOtocoResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderListPlaceOtocoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OrderListPlaceOtocoResponse) GetResult() OrderListPlaceOtocoResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderListPlaceOtocoResponse) GetResultOk() (*OrderListPlaceOtocoResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderListPlaceOtocoResponse) SetResult(v OrderListPlaceOtocoResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderListPlaceOtocoResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OrderListPlaceOtocoResponse) GetRateLimits() []RateLimits`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OrderListPlaceOtocoResponse) GetRateLimitsOk() (*[]RateLimits, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OrderListPlaceOtocoResponse) SetRateLimits(v []RateLimits)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OrderListPlaceOtocoResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


