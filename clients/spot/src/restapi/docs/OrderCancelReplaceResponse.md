# OrderCancelReplaceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**CancelResult** | Pointer to **string** |  | [optional] 
**NewOrderResult** | Pointer to **string** |  | [optional] 
**CancelResponse** | Pointer to [**OrderCancelReplaceResponseCancelResponse**](OrderCancelReplaceResponseCancelResponse.md) |  | [optional] 
**NewOrderResponse** | Pointer to [**OrderCancelReplaceResponseNewOrderResponse**](OrderCancelReplaceResponseNewOrderResponse.md) |  | [optional] 
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**OrderCancelReplaceResponseData**](OrderCancelReplaceResponseData.md) |  | [optional] 

## Methods

### NewOrderCancelReplaceResponse

`func NewOrderCancelReplaceResponse() *OrderCancelReplaceResponse`

NewOrderCancelReplaceResponse instantiates a new OrderCancelReplaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCancelReplaceResponseWithDefaults

`func NewOrderCancelReplaceResponseWithDefaults() *OrderCancelReplaceResponse`

NewOrderCancelReplaceResponseWithDefaults instantiates a new OrderCancelReplaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelResult

`func (o *OrderCancelReplaceResponse) GetCancelResult() string`

GetCancelResult returns the CancelResult field if non-nil, zero value otherwise.

### GetCancelResultOk

`func (o *OrderCancelReplaceResponse) GetCancelResultOk() (*string, bool)`

GetCancelResultOk returns a tuple with the CancelResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResult

`func (o *OrderCancelReplaceResponse) SetCancelResult(v string)`

SetCancelResult sets CancelResult field to given value.

### HasCancelResult

`func (o *OrderCancelReplaceResponse) HasCancelResult() bool`

HasCancelResult returns a boolean if a field has been set.

### GetNewOrderResult

`func (o *OrderCancelReplaceResponse) GetNewOrderResult() string`

GetNewOrderResult returns the NewOrderResult field if non-nil, zero value otherwise.

### GetNewOrderResultOk

`func (o *OrderCancelReplaceResponse) GetNewOrderResultOk() (*string, bool)`

GetNewOrderResultOk returns a tuple with the NewOrderResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResult

`func (o *OrderCancelReplaceResponse) SetNewOrderResult(v string)`

SetNewOrderResult sets NewOrderResult field to given value.

### HasNewOrderResult

`func (o *OrderCancelReplaceResponse) HasNewOrderResult() bool`

HasNewOrderResult returns a boolean if a field has been set.

### GetCancelResponse

`func (o *OrderCancelReplaceResponse) GetCancelResponse() OrderCancelReplaceResponseCancelResponse`

GetCancelResponse returns the CancelResponse field if non-nil, zero value otherwise.

### GetCancelResponseOk

`func (o *OrderCancelReplaceResponse) GetCancelResponseOk() (*OrderCancelReplaceResponseCancelResponse, bool)`

GetCancelResponseOk returns a tuple with the CancelResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResponse

`func (o *OrderCancelReplaceResponse) SetCancelResponse(v OrderCancelReplaceResponseCancelResponse)`

SetCancelResponse sets CancelResponse field to given value.

### HasCancelResponse

`func (o *OrderCancelReplaceResponse) HasCancelResponse() bool`

HasCancelResponse returns a boolean if a field has been set.

### GetNewOrderResponse

`func (o *OrderCancelReplaceResponse) GetNewOrderResponse() OrderCancelReplaceResponseNewOrderResponse`

GetNewOrderResponse returns the NewOrderResponse field if non-nil, zero value otherwise.

### GetNewOrderResponseOk

`func (o *OrderCancelReplaceResponse) GetNewOrderResponseOk() (*OrderCancelReplaceResponseNewOrderResponse, bool)`

GetNewOrderResponseOk returns a tuple with the NewOrderResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResponse

`func (o *OrderCancelReplaceResponse) SetNewOrderResponse(v OrderCancelReplaceResponseNewOrderResponse)`

SetNewOrderResponse sets NewOrderResponse field to given value.

### HasNewOrderResponse

`func (o *OrderCancelReplaceResponse) HasNewOrderResponse() bool`

HasNewOrderResponse returns a boolean if a field has been set.

### GetCode

`func (o *OrderCancelReplaceResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OrderCancelReplaceResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OrderCancelReplaceResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *OrderCancelReplaceResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *OrderCancelReplaceResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OrderCancelReplaceResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OrderCancelReplaceResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OrderCancelReplaceResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *OrderCancelReplaceResponse) GetData() OrderCancelReplaceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderCancelReplaceResponse) GetDataOk() (*OrderCancelReplaceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderCancelReplaceResponse) SetData(v OrderCancelReplaceResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *OrderCancelReplaceResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


