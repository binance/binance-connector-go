# AcceptQuoteResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewAcceptQuoteResponse

`func NewAcceptQuoteResponse() *AcceptQuoteResponse`

NewAcceptQuoteResponse instantiates a new AcceptQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptQuoteResponseWithDefaults

`func NewAcceptQuoteResponseWithDefaults() *AcceptQuoteResponse`

NewAcceptQuoteResponseWithDefaults instantiates a new AcceptQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *AcceptQuoteResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AcceptQuoteResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AcceptQuoteResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AcceptQuoteResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcceptQuoteResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcceptQuoteResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcceptQuoteResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcceptQuoteResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetOrderStatus

`func (o *AcceptQuoteResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *AcceptQuoteResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *AcceptQuoteResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *AcceptQuoteResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


