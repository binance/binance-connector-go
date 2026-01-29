# PlaceLimitOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaceLimitOrderResponse

`func NewPlaceLimitOrderResponse() *PlaceLimitOrderResponse`

NewPlaceLimitOrderResponse instantiates a new PlaceLimitOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceLimitOrderResponseWithDefaults

`func NewPlaceLimitOrderResponseWithDefaults() *PlaceLimitOrderResponse`

NewPlaceLimitOrderResponseWithDefaults instantiates a new PlaceLimitOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *PlaceLimitOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PlaceLimitOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PlaceLimitOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PlaceLimitOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *PlaceLimitOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceLimitOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaceLimitOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlaceLimitOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


