# NewAlgoOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **int64** |  | [optional] 
**ClientAlgoId** | Pointer to **string** |  | [optional] 
**AlgoType** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**AlgoStatus** | Pointer to **string** |  | [optional] 
**TriggerPrice** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**IcebergQuantity** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**ActivatePrice** | Pointer to **string** |  | [optional] 
**CallbackRate** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**TriggerTime** | Pointer to **int64** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewNewAlgoOrderResponse

`func NewNewAlgoOrderResponse() *NewAlgoOrderResponse`

NewNewAlgoOrderResponse instantiates a new NewAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAlgoOrderResponseWithDefaults

`func NewNewAlgoOrderResponseWithDefaults() *NewAlgoOrderResponse`

NewNewAlgoOrderResponseWithDefaults instantiates a new NewAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *NewAlgoOrderResponse) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *NewAlgoOrderResponse) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *NewAlgoOrderResponse) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *NewAlgoOrderResponse) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *NewAlgoOrderResponse) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *NewAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *NewAlgoOrderResponse) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *NewAlgoOrderResponse) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetAlgoType

`func (o *NewAlgoOrderResponse) GetAlgoType() string`

GetAlgoType returns the AlgoType field if non-nil, zero value otherwise.

### GetAlgoTypeOk

`func (o *NewAlgoOrderResponse) GetAlgoTypeOk() (*string, bool)`

GetAlgoTypeOk returns a tuple with the AlgoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoType

`func (o *NewAlgoOrderResponse) SetAlgoType(v string)`

SetAlgoType sets AlgoType field to given value.

### HasAlgoType

`func (o *NewAlgoOrderResponse) HasAlgoType() bool`

HasAlgoType returns a boolean if a field has been set.

### GetOrderType

`func (o *NewAlgoOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *NewAlgoOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *NewAlgoOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *NewAlgoOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSymbol

`func (o *NewAlgoOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NewAlgoOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NewAlgoOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NewAlgoOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *NewAlgoOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *NewAlgoOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *NewAlgoOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *NewAlgoOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *NewAlgoOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *NewAlgoOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *NewAlgoOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *NewAlgoOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *NewAlgoOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *NewAlgoOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *NewAlgoOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *NewAlgoOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *NewAlgoOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NewAlgoOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NewAlgoOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NewAlgoOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAlgoStatus

`func (o *NewAlgoOrderResponse) GetAlgoStatus() string`

GetAlgoStatus returns the AlgoStatus field if non-nil, zero value otherwise.

### GetAlgoStatusOk

`func (o *NewAlgoOrderResponse) GetAlgoStatusOk() (*string, bool)`

GetAlgoStatusOk returns a tuple with the AlgoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoStatus

`func (o *NewAlgoOrderResponse) SetAlgoStatus(v string)`

SetAlgoStatus sets AlgoStatus field to given value.

### HasAlgoStatus

`func (o *NewAlgoOrderResponse) HasAlgoStatus() bool`

HasAlgoStatus returns a boolean if a field has been set.

### GetTriggerPrice

`func (o *NewAlgoOrderResponse) GetTriggerPrice() string`

GetTriggerPrice returns the TriggerPrice field if non-nil, zero value otherwise.

### GetTriggerPriceOk

`func (o *NewAlgoOrderResponse) GetTriggerPriceOk() (*string, bool)`

GetTriggerPriceOk returns a tuple with the TriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPrice

`func (o *NewAlgoOrderResponse) SetTriggerPrice(v string)`

SetTriggerPrice sets TriggerPrice field to given value.

### HasTriggerPrice

`func (o *NewAlgoOrderResponse) HasTriggerPrice() bool`

HasTriggerPrice returns a boolean if a field has been set.

### GetPrice

`func (o *NewAlgoOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewAlgoOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NewAlgoOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NewAlgoOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIcebergQuantity

`func (o *NewAlgoOrderResponse) GetIcebergQuantity() string`

GetIcebergQuantity returns the IcebergQuantity field if non-nil, zero value otherwise.

### GetIcebergQuantityOk

`func (o *NewAlgoOrderResponse) GetIcebergQuantityOk() (*string, bool)`

GetIcebergQuantityOk returns a tuple with the IcebergQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQuantity

`func (o *NewAlgoOrderResponse) SetIcebergQuantity(v string)`

SetIcebergQuantity sets IcebergQuantity field to given value.

### HasIcebergQuantity

`func (o *NewAlgoOrderResponse) HasIcebergQuantity() bool`

HasIcebergQuantity returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *NewAlgoOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *NewAlgoOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *NewAlgoOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *NewAlgoOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetWorkingType

`func (o *NewAlgoOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *NewAlgoOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *NewAlgoOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *NewAlgoOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *NewAlgoOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *NewAlgoOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *NewAlgoOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *NewAlgoOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetClosePosition

`func (o *NewAlgoOrderResponse) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *NewAlgoOrderResponse) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *NewAlgoOrderResponse) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *NewAlgoOrderResponse) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetPriceProtect

`func (o *NewAlgoOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *NewAlgoOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *NewAlgoOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *NewAlgoOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *NewAlgoOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *NewAlgoOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *NewAlgoOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *NewAlgoOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetActivatePrice

`func (o *NewAlgoOrderResponse) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *NewAlgoOrderResponse) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *NewAlgoOrderResponse) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *NewAlgoOrderResponse) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *NewAlgoOrderResponse) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *NewAlgoOrderResponse) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *NewAlgoOrderResponse) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *NewAlgoOrderResponse) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetCreateTime

`func (o *NewAlgoOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NewAlgoOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NewAlgoOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NewAlgoOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *NewAlgoOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *NewAlgoOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *NewAlgoOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *NewAlgoOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *NewAlgoOrderResponse) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *NewAlgoOrderResponse) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *NewAlgoOrderResponse) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *NewAlgoOrderResponse) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *NewAlgoOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *NewAlgoOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *NewAlgoOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *NewAlgoOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


