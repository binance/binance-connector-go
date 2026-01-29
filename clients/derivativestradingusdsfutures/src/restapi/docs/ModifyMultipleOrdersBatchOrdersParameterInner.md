# ModifyMultipleOrdersBatchOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**RecvWindow** | Pointer to **string** |  | [optional] 

## Methods

### NewModifyMultipleOrdersBatchOrdersParameterInner

`func NewModifyMultipleOrdersBatchOrdersParameterInner() *ModifyMultipleOrdersBatchOrdersParameterInner`

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

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetOrderId(v string)`

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

### HasSymbol

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSide() NewAlgoOrderSideParameter`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*NewAlgoOrderSideParameter, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetSide(v NewAlgoOrderSideParameter)`

SetSide sets Side field to given value.

### HasSide

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceMatch() NewAlgoOrderPriceMatchParameter`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetPriceMatchOk() (*NewAlgoOrderPriceMatchParameter, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetPriceMatch(v NewAlgoOrderPriceMatchParameter)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetStopPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindow() string`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) GetRecvWindowOk() (*string, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) SetRecvWindow(v string)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *ModifyMultipleOrdersBatchOrdersParameterInner) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.


[[Back to README]](../README.md)


