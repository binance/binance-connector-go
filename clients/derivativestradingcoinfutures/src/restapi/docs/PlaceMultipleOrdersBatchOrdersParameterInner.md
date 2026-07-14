# PlaceMultipleOrdersBatchOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol | 
**Side** | [**PlaceMultipleOrdersBatchOrdersParameterInnerSide**](PlaceMultipleOrdersBatchOrdersParameterInnerSide.md) |  | 
**PositionSide** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide**](PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide.md) |  | [optional] 
**Type** | [**PlaceMultipleOrdersBatchOrdersParameterInnerType**](PlaceMultipleOrdersBatchOrdersParameterInnerType.md) |  | 
**TimeInForce** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce**](PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce.md) |  | [optional] 
**Quantity** | **float32** | quantity measured by contract number | 
**ReduceOnly** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly**](PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly.md) |  | [optional] [default to PLACEMULTIPLEORDERSBATCHORDERSPARAMETERINNERREDUCEONLY_FALSE]
**Price** | Pointer to **float32** | Order price | [optional] 
**NewClientOrderId** | Pointer to **string** | A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: &#x60;^[\\.A-Z\\:/a-z0-9_-]{1,36}$&#x60; | [optional] 
**StopPrice** | Pointer to **float32** | Used with &#x60;STOP/STOP_MARKET&#x60; or &#x60;TAKE_PROFIT/TAKE_PROFIT_MARKET&#x60; orders. | [optional] 
**ActivationPrice** | Pointer to **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, default as the latest price(supporting different &#x60;workingType&#x60;) | [optional] 
**CallbackRate** | Pointer to **float32** | Used with &#x60;TRAILING_STOP_MARKET&#x60; orders, min 0.1, max 4 where 1 for 1% | [optional] 
**WorkingType** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType**](PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType.md) |  | [optional] [default to PLACEMULTIPLEORDERSBATCHORDERSPARAMETERINNERWORKINGTYPE_CONTRACT_PRICE]
**PriceProtect** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect**](PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect.md) |  | [optional] [default to PLACEMULTIPLEORDERSBATCHORDERSPARAMETERINNERPRICEPROTECT_FALSE]
**NewOrderRespType** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType**](PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType.md) |  | [optional] [default to PLACEMULTIPLEORDERSBATCHORDERSPARAMETERINNERNEWORDERRESPTYPE_ACK]
**PriceMatch** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch**](PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch.md) |  | [optional] 
**SelfTradePreventionMode** | Pointer to [**PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode**](PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode.md) |  | [optional] [default to PLACEMULTIPLEORDERSBATCHORDERSPARAMETERINNERSELFTRADEPREVENTIONMODE_EXPIRE_MAKER]

## Methods

### NewPlaceMultipleOrdersBatchOrdersParameterInner

`func NewPlaceMultipleOrdersBatchOrdersParameterInner(symbol string, side PlaceMultipleOrdersBatchOrdersParameterInnerSide, type_ PlaceMultipleOrdersBatchOrdersParameterInnerType, quantity float32, ) *PlaceMultipleOrdersBatchOrdersParameterInner`

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


### GetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() PlaceMultipleOrdersBatchOrdersParameterInnerSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v PlaceMultipleOrdersBatchOrdersParameterInnerSide)`

SetSide sets Side field to given value.


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

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnly() PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnlyOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetReduceOnly(v PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPrice(v float32)`

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

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPrice() float32`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*float32, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v float32)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPrice() float32`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPriceOk() (*float32, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetActivationPrice(v float32)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRate() float32`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRateOk() (*float32, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetCallbackRate(v float32)`

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

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtect() PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtectOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceProtect(v PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect)`

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


