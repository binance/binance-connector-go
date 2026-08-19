# EquityOrderDetailResponseTradesInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** | Execution (fill) id. | [optional] 
**ExecutionAt** | Pointer to **int64** | Execution time (ms epoch). | [optional] 
**Price** | Pointer to **string** | Fill price for this execution (USD). | [optional] 
**Qty** | Pointer to **string** | Fill quantity. | [optional] 

## Methods

### NewEquityOrderDetailResponseTradesInner

`func NewEquityOrderDetailResponseTradesInner() *EquityOrderDetailResponseTradesInner`

NewEquityOrderDetailResponseTradesInner instantiates a new EquityOrderDetailResponseTradesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityOrderDetailResponseTradesInnerWithDefaults

`func NewEquityOrderDetailResponseTradesInnerWithDefaults() *EquityOrderDetailResponseTradesInner`

NewEquityOrderDetailResponseTradesInnerWithDefaults instantiates a new EquityOrderDetailResponseTradesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *EquityOrderDetailResponseTradesInner) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EquityOrderDetailResponseTradesInner) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EquityOrderDetailResponseTradesInner) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EquityOrderDetailResponseTradesInner) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetExecutionAt

`func (o *EquityOrderDetailResponseTradesInner) GetExecutionAt() int64`

GetExecutionAt returns the ExecutionAt field if non-nil, zero value otherwise.

### GetExecutionAtOk

`func (o *EquityOrderDetailResponseTradesInner) GetExecutionAtOk() (*int64, bool)`

GetExecutionAtOk returns a tuple with the ExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionAt

`func (o *EquityOrderDetailResponseTradesInner) SetExecutionAt(v int64)`

SetExecutionAt sets ExecutionAt field to given value.

### HasExecutionAt

`func (o *EquityOrderDetailResponseTradesInner) HasExecutionAt() bool`

HasExecutionAt returns a boolean if a field has been set.

### GetPrice

`func (o *EquityOrderDetailResponseTradesInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EquityOrderDetailResponseTradesInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EquityOrderDetailResponseTradesInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EquityOrderDetailResponseTradesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *EquityOrderDetailResponseTradesInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *EquityOrderDetailResponseTradesInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *EquityOrderDetailResponseTradesInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *EquityOrderDetailResponseTradesInner) HasQty() bool`

HasQty returns a boolean if a field has been set.


[[Back to README]](../README.md)


