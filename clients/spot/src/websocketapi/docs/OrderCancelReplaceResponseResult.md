# OrderCancelReplaceResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CancelResult** | Pointer to **string** |  | [optional] 
**NewOrderResult** | Pointer to **string** |  | [optional] 
**CancelResponse** | Pointer to [**OrderCancelReplaceResponseResultCancelResponse**](OrderCancelReplaceResponseResultCancelResponse.md) |  | [optional] 
**NewOrderResponse** | Pointer to [**OrderCancelReplaceResponseResultNewOrderResponse**](OrderCancelReplaceResponseResultNewOrderResponse.md) |  | [optional] 

## Methods

### NewOrderCancelReplaceResponseResult

`func NewOrderCancelReplaceResponseResult() *OrderCancelReplaceResponseResult`

NewOrderCancelReplaceResponseResult instantiates a new OrderCancelReplaceResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCancelReplaceResponseResultWithDefaults

`func NewOrderCancelReplaceResponseResultWithDefaults() *OrderCancelReplaceResponseResult`

NewOrderCancelReplaceResponseResultWithDefaults instantiates a new OrderCancelReplaceResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelResult

`func (o *OrderCancelReplaceResponseResult) GetCancelResult() string`

GetCancelResult returns the CancelResult field if non-nil, zero value otherwise.

### GetCancelResultOk

`func (o *OrderCancelReplaceResponseResult) GetCancelResultOk() (*string, bool)`

GetCancelResultOk returns a tuple with the CancelResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResult

`func (o *OrderCancelReplaceResponseResult) SetCancelResult(v string)`

SetCancelResult sets CancelResult field to given value.

### HasCancelResult

`func (o *OrderCancelReplaceResponseResult) HasCancelResult() bool`

HasCancelResult returns a boolean if a field has been set.

### GetNewOrderResult

`func (o *OrderCancelReplaceResponseResult) GetNewOrderResult() string`

GetNewOrderResult returns the NewOrderResult field if non-nil, zero value otherwise.

### GetNewOrderResultOk

`func (o *OrderCancelReplaceResponseResult) GetNewOrderResultOk() (*string, bool)`

GetNewOrderResultOk returns a tuple with the NewOrderResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResult

`func (o *OrderCancelReplaceResponseResult) SetNewOrderResult(v string)`

SetNewOrderResult sets NewOrderResult field to given value.

### HasNewOrderResult

`func (o *OrderCancelReplaceResponseResult) HasNewOrderResult() bool`

HasNewOrderResult returns a boolean if a field has been set.

### GetCancelResponse

`func (o *OrderCancelReplaceResponseResult) GetCancelResponse() OrderCancelReplaceResponseResultCancelResponse`

GetCancelResponse returns the CancelResponse field if non-nil, zero value otherwise.

### GetCancelResponseOk

`func (o *OrderCancelReplaceResponseResult) GetCancelResponseOk() (*OrderCancelReplaceResponseResultCancelResponse, bool)`

GetCancelResponseOk returns a tuple with the CancelResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResponse

`func (o *OrderCancelReplaceResponseResult) SetCancelResponse(v OrderCancelReplaceResponseResultCancelResponse)`

SetCancelResponse sets CancelResponse field to given value.

### HasCancelResponse

`func (o *OrderCancelReplaceResponseResult) HasCancelResponse() bool`

HasCancelResponse returns a boolean if a field has been set.

### GetNewOrderResponse

`func (o *OrderCancelReplaceResponseResult) GetNewOrderResponse() OrderCancelReplaceResponseResultNewOrderResponse`

GetNewOrderResponse returns the NewOrderResponse field if non-nil, zero value otherwise.

### GetNewOrderResponseOk

`func (o *OrderCancelReplaceResponseResult) GetNewOrderResponseOk() (*OrderCancelReplaceResponseResultNewOrderResponse, bool)`

GetNewOrderResponseOk returns a tuple with the NewOrderResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResponse

`func (o *OrderCancelReplaceResponseResult) SetNewOrderResponse(v OrderCancelReplaceResponseResultNewOrderResponse)`

SetNewOrderResponse sets NewOrderResponse field to given value.

### HasNewOrderResponse

`func (o *OrderCancelReplaceResponseResult) HasNewOrderResponse() bool`

HasNewOrderResponse returns a boolean if a field has been set.


[[Back to README]](../README.md)


