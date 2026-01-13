# OrderBookResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**OrderBookResponseResult**](OrderBookResponseResult.md) |  | [optional] 
**RateLimits** | Pointer to [**[]OrderBookResponseRateLimitsInner**](OrderBookResponseRateLimitsInner.md) |  | [optional] 

## Methods

### NewOrderBookResponse

`func NewOrderBookResponse() *OrderBookResponse`

NewOrderBookResponse instantiates a new OrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookResponseWithDefaults

`func NewOrderBookResponseWithDefaults() *OrderBookResponse`

NewOrderBookResponseWithDefaults instantiates a new OrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderBookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderBookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderBookResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderBookResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderBookResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderBookResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderBookResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderBookResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OrderBookResponse) GetResult() OrderBookResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderBookResponse) GetResultOk() (*OrderBookResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderBookResponse) SetResult(v OrderBookResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderBookResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRateLimits

`func (o *OrderBookResponse) GetRateLimits() []OrderBookResponseRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OrderBookResponse) GetRateLimitsOk() (*[]OrderBookResponseRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OrderBookResponse) SetRateLimits(v []OrderBookResponseRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OrderBookResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to README]](../README.md)


