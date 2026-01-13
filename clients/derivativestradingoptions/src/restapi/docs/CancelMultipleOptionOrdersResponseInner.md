# CancelMultipleOptionOrdersResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **int64** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewCancelMultipleOptionOrdersResponseInner

`func NewCancelMultipleOptionOrdersResponseInner() *CancelMultipleOptionOrdersResponseInner`

NewCancelMultipleOptionOrdersResponseInner instantiates a new CancelMultipleOptionOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelMultipleOptionOrdersResponseInnerWithDefaults

`func NewCancelMultipleOptionOrdersResponseInnerWithDefaults() *CancelMultipleOptionOrdersResponseInner`

NewCancelMultipleOptionOrdersResponseInnerWithDefaults instantiates a new CancelMultipleOptionOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *CancelMultipleOptionOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelMultipleOptionOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelMultipleOptionOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *CancelMultipleOptionOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CancelMultipleOptionOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CancelMultipleOptionOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CancelMultipleOptionOrdersResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CancelMultipleOptionOrdersResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CancelMultipleOptionOrdersResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetExecutedQty

`func (o *CancelMultipleOptionOrdersResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *CancelMultipleOptionOrdersResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *CancelMultipleOptionOrdersResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFee

`func (o *CancelMultipleOptionOrdersResponseInner) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CancelMultipleOptionOrdersResponseInner) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CancelMultipleOptionOrdersResponseInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSide

`func (o *CancelMultipleOptionOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CancelMultipleOptionOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CancelMultipleOptionOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *CancelMultipleOptionOrdersResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancelMultipleOptionOrdersResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CancelMultipleOptionOrdersResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CancelMultipleOptionOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CancelMultipleOptionOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CancelMultipleOptionOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetCreateTime

`func (o *CancelMultipleOptionOrdersResponseInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CancelMultipleOptionOrdersResponseInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CancelMultipleOptionOrdersResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetStatus

`func (o *CancelMultipleOptionOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelMultipleOptionOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelMultipleOptionOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvgPrice

`func (o *CancelMultipleOptionOrdersResponseInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CancelMultipleOptionOrdersResponseInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *CancelMultipleOptionOrdersResponseInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CancelMultipleOptionOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CancelMultipleOptionOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CancelMultipleOptionOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetClientOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *CancelMultipleOptionOrdersResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CancelMultipleOptionOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CancelMultipleOptionOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CancelMultipleOptionOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CancelMultipleOptionOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


