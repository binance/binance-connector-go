# CancelEquityOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Echoes the requested order id. | [optional] 
**Status** | Pointer to **string** | Acknowledgement code: &#x60;S&#x60; &#x3D; cancel accepted by upstream, &#x60;F&#x60; &#x3D; cancel failed. Not an order lifecycle status — use &#x60;/order/detail&#x60; for lifecycle poll. | [optional] 

## Methods

### NewCancelEquityOrderResponse

`func NewCancelEquityOrderResponse() *CancelEquityOrderResponse`

NewCancelEquityOrderResponse instantiates a new CancelEquityOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelEquityOrderResponseWithDefaults

`func NewCancelEquityOrderResponseWithDefaults() *CancelEquityOrderResponse`

NewCancelEquityOrderResponseWithDefaults instantiates a new CancelEquityOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CancelEquityOrderResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelEquityOrderResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelEquityOrderResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CancelEquityOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *CancelEquityOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelEquityOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelEquityOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelEquityOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


