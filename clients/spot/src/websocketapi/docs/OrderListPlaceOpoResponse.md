# OrderListPlaceOpoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**Result** | Pointer to [**OrderListPlaceOpoResponseResult**](OrderListPlaceOpoResponseResult.md) |  | [optional] 

## Methods

### NewOrderListPlaceOpoResponse

`func NewOrderListPlaceOpoResponse() *OrderListPlaceOpoResponse`

NewOrderListPlaceOpoResponse instantiates a new OrderListPlaceOpoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOpoResponseWithDefaults

`func NewOrderListPlaceOpoResponseWithDefaults() *OrderListPlaceOpoResponse`

NewOrderListPlaceOpoResponseWithDefaults instantiates a new OrderListPlaceOpoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderListPlaceOpoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderListPlaceOpoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderListPlaceOpoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderListPlaceOpoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderListPlaceOpoResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListPlaceOpoResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListPlaceOpoResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderListPlaceOpoResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *OrderListPlaceOpoResponse) GetResult() OrderListPlaceOpoResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OrderListPlaceOpoResponse) GetResultOk() (*OrderListPlaceOpoResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OrderListPlaceOpoResponse) SetResult(v OrderListPlaceOpoResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *OrderListPlaceOpoResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


