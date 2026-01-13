# PlaceMultipleOrdersBatchOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**ModifyMultipleOrdersBatchOrdersParameterInnerSide**](ModifyMultipleOrdersBatchOrdersParameterInnerSide.md) |  | [optional] 
**PositionSide** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide**](PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide.md) |  | [optional] 
**Type** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) |  | [optional] 
**TimeInForce** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce**](PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce.md) |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**NewClientOrderId** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**ActivationPrice** | Pointer to **string** |  | [optional] 
**CallbackRate** | Pointer to **string** |  | [optional] 
**WorkingType** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType**](PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType.md) |  | [optional] 
**PriceProtect** | Pointer to **string** |  | [optional] 
**NewOrderRespType** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType**](PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType.md) |  | [optional] 
**PriceMatch** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch**](PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch.md) |  | [optional] 
**SelfTradePreventionMode** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode**](PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode.md) |  | [optional] 

## Methods

### NewPlaceMultipleOrdersBatchOrdersParameterInner

`func NewPlaceMultipleOrdersBatchOrdersParameterInner() *PlaceMultipleOrdersBatchOrdersParameterInner`

NewPlaceMultipleOrdersBatchOrdersParameterInner instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults

`func NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersBatchOrdersParameterInner`

NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() ModifyMultipleOrdersBatchOrdersParameterInnerSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*ModifyMultipleOrdersBatchOrdersParameterInnerSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v ModifyMultipleOrdersBatchOrdersParameterInnerSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSide() PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSideOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPositionSide(v PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetType() PlaceMultipleOrdersBatchOrdersParameterInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetType(v PlaceMultipleOrdersBatchOrdersParameterInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForce() PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForceOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetTimeInForce(v PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetStopPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPrice() string`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPriceOk() (*string, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetActivationPrice(v string)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetWorkingType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetWorkingType() PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetWorkingTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetWorkingType(v PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtect() string`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtectOk() (*string, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceProtect(v string)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespType() PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewOrderRespType(v PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatch() PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatchOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceMatch(v PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionMode() PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionModeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSelfTradePreventionMode(v PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.


[[Back to README]](../README.md)


