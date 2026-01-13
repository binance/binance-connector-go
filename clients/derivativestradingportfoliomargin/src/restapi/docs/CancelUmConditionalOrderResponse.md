# CancelUmConditionalOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**NewClientStrategyId** | Pointer to **string** |  | [optional] 
**StrategyId** | Pointer to **int64** |  | [optional] 
**StrategyStatus** | Pointer to **string** |  | [optional] 
**StrategyType** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**ActivatePrice** | Pointer to **string** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**BookTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 

## Methods

### NewCancelUmConditionalOrderResponse

`func NewCancelUmConditionalOrderResponse() *CancelUmConditionalOrderResponse`

NewCancelUmConditionalOrderResponse instantiates a new CancelUmConditionalOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelUmConditionalOrderResponseWithDefaults

`func NewCancelUmConditionalOrderResponseWithDefaults() *CancelUmConditionalOrderResponse`

NewCancelUmConditionalOrderResponseWithDefaults instantiates a new CancelUmConditionalOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewClientStrategyId

`func (o *CancelUmConditionalOrderResponse) GetNewClientStrategyId() string`

GetNewClientStrategyId returns the NewClientStrategyId field if non-nil, zero value otherwise.

### GetNewClientStrategyIdOk

`func (o *CancelUmConditionalOrderResponse) GetNewClientStrategyIdOk() (*string, bool)`

GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientStrategyId

`func (o *CancelUmConditionalOrderResponse) SetNewClientStrategyId(v string)`

SetNewClientStrategyId sets NewClientStrategyId field to given value.

### HasNewClientStrategyId

`func (o *CancelUmConditionalOrderResponse) HasNewClientStrategyId() bool`

HasNewClientStrategyId returns a boolean if a field has been set.

### GetStrategyId

`func (o *CancelUmConditionalOrderResponse) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *CancelUmConditionalOrderResponse) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *CancelUmConditionalOrderResponse) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *CancelUmConditionalOrderResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyStatus

`func (o *CancelUmConditionalOrderResponse) GetStrategyStatus() string`

GetStrategyStatus returns the StrategyStatus field if non-nil, zero value otherwise.

### GetStrategyStatusOk

`func (o *CancelUmConditionalOrderResponse) GetStrategyStatusOk() (*string, bool)`

GetStrategyStatusOk returns a tuple with the StrategyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyStatus

`func (o *CancelUmConditionalOrderResponse) SetStrategyStatus(v string)`

SetStrategyStatus sets StrategyStatus field to given value.

### HasStrategyStatus

`func (o *CancelUmConditionalOrderResponse) HasStrategyStatus() bool`

HasStrategyStatus returns a boolean if a field has been set.

### GetStrategyType

`func (o *CancelUmConditionalOrderResponse) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *CancelUmConditionalOrderResponse) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *CancelUmConditionalOrderResponse) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *CancelUmConditionalOrderResponse) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetOrigQty

`func (o *CancelUmConditionalOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *CancelUmConditionalOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *CancelUmConditionalOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *CancelUmConditionalOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *CancelUmConditionalOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CancelUmConditionalOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CancelUmConditionalOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CancelUmConditionalOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CancelUmConditionalOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CancelUmConditionalOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CancelUmConditionalOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CancelUmConditionalOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CancelUmConditionalOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CancelUmConditionalOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CancelUmConditionalOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CancelUmConditionalOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *CancelUmConditionalOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CancelUmConditionalOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CancelUmConditionalOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CancelUmConditionalOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *CancelUmConditionalOrderResponse) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CancelUmConditionalOrderResponse) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CancelUmConditionalOrderResponse) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CancelUmConditionalOrderResponse) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *CancelUmConditionalOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelUmConditionalOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelUmConditionalOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelUmConditionalOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CancelUmConditionalOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CancelUmConditionalOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CancelUmConditionalOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CancelUmConditionalOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetActivatePrice

`func (o *CancelUmConditionalOrderResponse) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *CancelUmConditionalOrderResponse) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *CancelUmConditionalOrderResponse) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *CancelUmConditionalOrderResponse) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetPriceRate

`func (o *CancelUmConditionalOrderResponse) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *CancelUmConditionalOrderResponse) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *CancelUmConditionalOrderResponse) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *CancelUmConditionalOrderResponse) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetBookTime

`func (o *CancelUmConditionalOrderResponse) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *CancelUmConditionalOrderResponse) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *CancelUmConditionalOrderResponse) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *CancelUmConditionalOrderResponse) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CancelUmConditionalOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CancelUmConditionalOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CancelUmConditionalOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CancelUmConditionalOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *CancelUmConditionalOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *CancelUmConditionalOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *CancelUmConditionalOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *CancelUmConditionalOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *CancelUmConditionalOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *CancelUmConditionalOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *CancelUmConditionalOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *CancelUmConditionalOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *CancelUmConditionalOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *CancelUmConditionalOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *CancelUmConditionalOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *CancelUmConditionalOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *CancelUmConditionalOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *CancelUmConditionalOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *CancelUmConditionalOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *CancelUmConditionalOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetPriceMatch

`func (o *CancelUmConditionalOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *CancelUmConditionalOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *CancelUmConditionalOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *CancelUmConditionalOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.


[[Back to README]](../README.md)


