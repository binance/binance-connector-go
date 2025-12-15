# OrderTradeUpdate

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int64** |  | [optional] 
**O** | Pointer to [**[]OrderTradeUpdateOInner**](OrderTradeUpdateOInner.md) |  | [optional] 

## Methods

### NewOrderTradeUpdate

`func NewOrderTradeUpdate() *OrderTradeUpdate`

NewOrderTradeUpdate instantiates a new OrderTradeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTradeUpdateWithDefaults

`func NewOrderTradeUpdateWithDefaults() *OrderTradeUpdate`

NewOrderTradeUpdateWithDefaults instantiates a new OrderTradeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OrderTradeUpdate) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OrderTradeUpdate) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OrderTradeUpdate) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OrderTradeUpdate) HasE() bool`

HasE returns a boolean if a field has been set.

### GetO

`func (o *OrderTradeUpdate) GetO() []OrderTradeUpdateOInner`

GetO returns the O field if non-nil, zero value otherwise.

### GetOOk

`func (o *OrderTradeUpdate) GetOOk() (*[]OrderTradeUpdateOInner, bool)`

GetOOk returns a tuple with the O field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetO

`func (o *OrderTradeUpdate) SetO(v []OrderTradeUpdateOInner)`

SetO sets O field to given value.

### HasO

`func (o *OrderTradeUpdate) HasO() bool`

HasO returns a boolean if a field has been set.


[[Back to README]](../README.md)


