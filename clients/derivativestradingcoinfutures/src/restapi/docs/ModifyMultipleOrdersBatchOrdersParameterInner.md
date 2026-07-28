# ModifyMultipleOrdersBatchOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Sub-order ID, required when clientOrderId is not sent. | [optional] 
**OrigClientOrderId** | Pointer to **string** | Client order ID, required when orderId is not sent. Only support letters, numbers and underscores, and must be unique within 24 hours. | [optional] 
**Symbol** | **string** | Trading symbol | 
**Side** | [**ModifyMultipleOrdersBatchOrdersParameterInnerSide**](ModifyMultipleOrdersBatchOrdersParameterInnerSide.md) |  | 
**Quantity** | Pointer to **float32** | Order quantity, cannot be sent with closePosition&#x3D;true. **After CM migration, this parameter becomes mandatory** (each batch element must send both &#x60;price&#x60; and &#x60;quantity&#x60;). | [optional] 
**Price** | Pointer to **float32** | Latest token price. **After CM migration, this parameter becomes mandatory** (each batch element must send both &#x60;price&#x60; and &#x60;quantity&#x60;). | [optional] 
**ModifyId** | Pointer to **int64** | User-defined modification identifier, returned as-is in the response. Optional; not validated for uniqueness. | [optional] 
**RecvWindow** | Pointer to **int64** |  | [optional] 
**Timestamp** | **int64** | Unix timestamp in milliseconds used to sign the request. The value must reflect the current client time and is validated by the server for signed endpoints. | 

## Methods

### NewModifyMultipleOrdersBatchOrdersParameterInner

`func NewModifyMultipleOrdersBatchOrdersParameterInner(symbol string, side ModifyMultipleOrdersBatchOrdersParameterInnerSide, timestamp int64, ) *ModifyMultipleOrdersBatchOrdersParameterInner`

NewModifyMultipleOrdersBatchOrdersParameterInner instantiates a new ModifyMultipleOrdersBatchOrdersParameterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMultipleOrdersBatchOrdersParameterInnerWithDefaults

`func NewModifyMultipleOrdersBatchOrdersParameterInnerWithDefaults() *ModifyMultipleOrdersBatchOrdersParameterInner`

NewModifyMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new ModifyMultipleOrdersBatchOrdersParameterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSide

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSide() ModifyMultipleOrdersBatchOrdersParameterInnerSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*ModifyMultipleOrdersBatchOrdersParameterInnerSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetSide(v ModifyMultipleOrdersBatchOrdersParameterInnerSide)`

SetSide sets Side field to given value.


### GetQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetModifyId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetModifyId() int64`

GetModifyId returns the ModifyId field if non-nil, zero value otherwise.

### GetModifyIdOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetModifyIdOk() (*int64, bool)`

GetModifyIdOk returns a tuple with the ModifyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetModifyId(v int64)`

SetModifyId sets ModifyId field to given value.

### HasModifyId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasModifyId() bool`

HasModifyId returns a boolean if a field has been set.

### GetRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindow() int64`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindowOk() (*int64, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetRecvWindow(v int64)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.

### GetTimestamp

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to README]](../README.md)


