# OrderCancelReplaceResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CancelResult** | Pointer to **string** |  | [optional] 
**NewOrderResult** | Pointer to **string** |  | [optional] 
**CancelResponse** | Pointer to [**OrderCancelReplaceResponseDataCancelResponse**](OrderCancelReplaceResponseDataCancelResponse.md) |  | [optional] 
**NewOrderResponse** | Pointer to [**OrderCancelReplaceResponseDataNewOrderResponse**](OrderCancelReplaceResponseDataNewOrderResponse.md) |  | [optional] 

## Methods

### NewOrderCancelReplaceResponseData

`func NewOrderCancelReplaceResponseData() *OrderCancelReplaceResponseData`

NewOrderCancelReplaceResponseData instantiates a new OrderCancelReplaceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCancelReplaceResponseDataWithDefaults

`func NewOrderCancelReplaceResponseDataWithDefaults() *OrderCancelReplaceResponseData`

NewOrderCancelReplaceResponseDataWithDefaults instantiates a new OrderCancelReplaceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelResult

`func (o *OrderCancelReplaceResponseData) GetCancelResult() string`

GetCancelResult returns the CancelResult field if non-nil, zero value otherwise.

### GetCancelResultOk

`func (o *OrderCancelReplaceResponseData) GetCancelResultOk() (*string, bool)`

GetCancelResultOk returns a tuple with the CancelResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResult

`func (o *OrderCancelReplaceResponseData) SetCancelResult(v string)`

SetCancelResult sets CancelResult field to given value.

### HasCancelResult

`func (o *OrderCancelReplaceResponseData) HasCancelResult() bool`

HasCancelResult returns a boolean if a field has been set.

### GetNewOrderResult

`func (o *OrderCancelReplaceResponseData) GetNewOrderResult() string`

GetNewOrderResult returns the NewOrderResult field if non-nil, zero value otherwise.

### GetNewOrderResultOk

`func (o *OrderCancelReplaceResponseData) GetNewOrderResultOk() (*string, bool)`

GetNewOrderResultOk returns a tuple with the NewOrderResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResult

`func (o *OrderCancelReplaceResponseData) SetNewOrderResult(v string)`

SetNewOrderResult sets NewOrderResult field to given value.

### HasNewOrderResult

`func (o *OrderCancelReplaceResponseData) HasNewOrderResult() bool`

HasNewOrderResult returns a boolean if a field has been set.

### GetCancelResponse

`func (o *OrderCancelReplaceResponseData) GetCancelResponse() OrderCancelReplaceResponseDataCancelResponse`

GetCancelResponse returns the CancelResponse field if non-nil, zero value otherwise.

### GetCancelResponseOk

`func (o *OrderCancelReplaceResponseData) GetCancelResponseOk() (*OrderCancelReplaceResponseDataCancelResponse, bool)`

GetCancelResponseOk returns a tuple with the CancelResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResponse

`func (o *OrderCancelReplaceResponseData) SetCancelResponse(v OrderCancelReplaceResponseDataCancelResponse)`

SetCancelResponse sets CancelResponse field to given value.

### HasCancelResponse

`func (o *OrderCancelReplaceResponseData) HasCancelResponse() bool`

HasCancelResponse returns a boolean if a field has been set.

### GetNewOrderResponse

`func (o *OrderCancelReplaceResponseData) GetNewOrderResponse() OrderCancelReplaceResponseDataNewOrderResponse`

GetNewOrderResponse returns the NewOrderResponse field if non-nil, zero value otherwise.

### GetNewOrderResponseOk

`func (o *OrderCancelReplaceResponseData) GetNewOrderResponseOk() (*OrderCancelReplaceResponseDataNewOrderResponse, bool)`

GetNewOrderResponseOk returns a tuple with the NewOrderResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResponse

`func (o *OrderCancelReplaceResponseData) SetNewOrderResponse(v OrderCancelReplaceResponseDataNewOrderResponse)`

SetNewOrderResponse sets NewOrderResponse field to given value.

### HasNewOrderResponse

`func (o *OrderCancelReplaceResponseData) HasNewOrderResponse() bool`

HasNewOrderResponse returns a boolean if a field has been set.


[[Back to README]](../README.md)


