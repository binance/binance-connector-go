# NewUmAlgoOrderResponse

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

### NewNewUmAlgoOrderResponse

`func NewNewUmAlgoOrderResponse() *NewUmAlgoOrderResponse`

NewNewUmAlgoOrderResponse instantiates a new NewUmAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUmAlgoOrderResponseWithDefaults

`func NewNewUmAlgoOrderResponseWithDefaults() *NewUmAlgoOrderResponse`

NewNewUmAlgoOrderResponseWithDefaults instantiates a new NewUmAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *NewUmAlgoOrderResponse) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *NewUmAlgoOrderResponse) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *NewUmAlgoOrderResponse) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *NewUmAlgoOrderResponse) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *NewUmAlgoOrderResponse) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *NewUmAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *NewUmAlgoOrderResponse) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *NewUmAlgoOrderResponse) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetAlgoType

`func (o *NewUmAlgoOrderResponse) GetAlgoType() string`

GetAlgoType returns the AlgoType field if non-nil, zero value otherwise.

### GetAlgoTypeOk

`func (o *NewUmAlgoOrderResponse) GetAlgoTypeOk() (*string, bool)`

GetAlgoTypeOk returns a tuple with the AlgoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoType

`func (o *NewUmAlgoOrderResponse) SetAlgoType(v string)`

SetAlgoType sets AlgoType field to given value.

### HasAlgoType

`func (o *NewUmAlgoOrderResponse) HasAlgoType() bool`

HasAlgoType returns a boolean if a field has been set.

### GetOrderType

`func (o *NewUmAlgoOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *NewUmAlgoOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *NewUmAlgoOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *NewUmAlgoOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSymbol

`func (o *NewUmAlgoOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *NewUmAlgoOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *NewUmAlgoOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *NewUmAlgoOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *NewUmAlgoOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *NewUmAlgoOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *NewUmAlgoOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *NewUmAlgoOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *NewUmAlgoOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *NewUmAlgoOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *NewUmAlgoOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *NewUmAlgoOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *NewUmAlgoOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *NewUmAlgoOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *NewUmAlgoOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *NewUmAlgoOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *NewUmAlgoOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *NewUmAlgoOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *NewUmAlgoOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *NewUmAlgoOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAlgoStatus

`func (o *NewUmAlgoOrderResponse) GetAlgoStatus() string`

GetAlgoStatus returns the AlgoStatus field if non-nil, zero value otherwise.

### GetAlgoStatusOk

`func (o *NewUmAlgoOrderResponse) GetAlgoStatusOk() (*string, bool)`

GetAlgoStatusOk returns a tuple with the AlgoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoStatus

`func (o *NewUmAlgoOrderResponse) SetAlgoStatus(v string)`

SetAlgoStatus sets AlgoStatus field to given value.

### HasAlgoStatus

`func (o *NewUmAlgoOrderResponse) HasAlgoStatus() bool`

HasAlgoStatus returns a boolean if a field has been set.

### GetTriggerPrice

`func (o *NewUmAlgoOrderResponse) GetTriggerPrice() string`

GetTriggerPrice returns the TriggerPrice field if non-nil, zero value otherwise.

### GetTriggerPriceOk

`func (o *NewUmAlgoOrderResponse) GetTriggerPriceOk() (*string, bool)`

GetTriggerPriceOk returns a tuple with the TriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPrice

`func (o *NewUmAlgoOrderResponse) SetTriggerPrice(v string)`

SetTriggerPrice sets TriggerPrice field to given value.

### HasTriggerPrice

`func (o *NewUmAlgoOrderResponse) HasTriggerPrice() bool`

HasTriggerPrice returns a boolean if a field has been set.

### GetPrice

`func (o *NewUmAlgoOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NewUmAlgoOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NewUmAlgoOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NewUmAlgoOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIcebergQuantity

`func (o *NewUmAlgoOrderResponse) GetIcebergQuantity() string`

GetIcebergQuantity returns the IcebergQuantity field if non-nil, zero value otherwise.

### GetIcebergQuantityOk

`func (o *NewUmAlgoOrderResponse) GetIcebergQuantityOk() (*string, bool)`

GetIcebergQuantityOk returns a tuple with the IcebergQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQuantity

`func (o *NewUmAlgoOrderResponse) SetIcebergQuantity(v string)`

SetIcebergQuantity sets IcebergQuantity field to given value.

### HasIcebergQuantity

`func (o *NewUmAlgoOrderResponse) HasIcebergQuantity() bool`

HasIcebergQuantity returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *NewUmAlgoOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *NewUmAlgoOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *NewUmAlgoOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *NewUmAlgoOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetWorkingType

`func (o *NewUmAlgoOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *NewUmAlgoOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *NewUmAlgoOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *NewUmAlgoOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *NewUmAlgoOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *NewUmAlgoOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *NewUmAlgoOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *NewUmAlgoOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetClosePosition

`func (o *NewUmAlgoOrderResponse) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *NewUmAlgoOrderResponse) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *NewUmAlgoOrderResponse) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *NewUmAlgoOrderResponse) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetPriceProtect

`func (o *NewUmAlgoOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *NewUmAlgoOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *NewUmAlgoOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *NewUmAlgoOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *NewUmAlgoOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *NewUmAlgoOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *NewUmAlgoOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *NewUmAlgoOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetActivatePrice

`func (o *NewUmAlgoOrderResponse) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *NewUmAlgoOrderResponse) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *NewUmAlgoOrderResponse) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *NewUmAlgoOrderResponse) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *NewUmAlgoOrderResponse) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *NewUmAlgoOrderResponse) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *NewUmAlgoOrderResponse) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *NewUmAlgoOrderResponse) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetCreateTime

`func (o *NewUmAlgoOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NewUmAlgoOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NewUmAlgoOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NewUmAlgoOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *NewUmAlgoOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *NewUmAlgoOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *NewUmAlgoOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *NewUmAlgoOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *NewUmAlgoOrderResponse) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *NewUmAlgoOrderResponse) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *NewUmAlgoOrderResponse) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *NewUmAlgoOrderResponse) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *NewUmAlgoOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *NewUmAlgoOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *NewUmAlgoOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *NewUmAlgoOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


