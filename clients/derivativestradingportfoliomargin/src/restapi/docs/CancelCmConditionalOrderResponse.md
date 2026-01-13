# CancelCmConditionalOrderResponse

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

## Methods

### NewCancelCmConditionalOrderResponse

`func NewCancelCmConditionalOrderResponse() *CancelCmConditionalOrderResponse`

NewCancelCmConditionalOrderResponse instantiates a new CancelCmConditionalOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelCmConditionalOrderResponseWithDefaults

`func NewCancelCmConditionalOrderResponseWithDefaults() *CancelCmConditionalOrderResponse`

NewCancelCmConditionalOrderResponseWithDefaults instantiates a new CancelCmConditionalOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewClientStrategyId

`func (o *CancelCmConditionalOrderResponse) GetNewClientStrategyId() string`

GetNewClientStrategyId returns the NewClientStrategyId field if non-nil, zero value otherwise.

### GetNewClientStrategyIdOk

`func (o *CancelCmConditionalOrderResponse) GetNewClientStrategyIdOk() (*string, bool)`

GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientStrategyId

`func (o *CancelCmConditionalOrderResponse) SetNewClientStrategyId(v string)`

SetNewClientStrategyId sets NewClientStrategyId field to given value.

### HasNewClientStrategyId

`func (o *CancelCmConditionalOrderResponse) HasNewClientStrategyId() bool`

HasNewClientStrategyId returns a boolean if a field has been set.

### GetStrategyId

`func (o *CancelCmConditionalOrderResponse) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *CancelCmConditionalOrderResponse) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *CancelCmConditionalOrderResponse) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *CancelCmConditionalOrderResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyStatus

`func (o *CancelCmConditionalOrderResponse) GetStrategyStatus() string`

GetStrategyStatus returns the StrategyStatus field if non-nil, zero value otherwise.

### GetStrategyStatusOk

`func (o *CancelCmConditionalOrderResponse) GetStrategyStatusOk() (*string, bool)`

GetStrategyStatusOk returns a tuple with the StrategyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyStatus

`func (o *CancelCmConditionalOrderResponse) SetStrategyStatus(v string)`

SetStrategyStatus sets StrategyStatus field to given value.

### HasStrategyStatus

`func (o *CancelCmConditionalOrderResponse) HasStrategyStatus() bool`

HasStrategyStatus returns a boolean if a field has been set.

### GetStrategyType

`func (o *CancelCmConditionalOrderResponse) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *CancelCmConditionalOrderResponse) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *CancelCmConditionalOrderResponse) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *CancelCmConditionalOrderResponse) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetOrigQty

`func (o *CancelCmConditionalOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *CancelCmConditionalOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *CancelCmConditionalOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *CancelCmConditionalOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *CancelCmConditionalOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CancelCmConditionalOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CancelCmConditionalOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CancelCmConditionalOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CancelCmConditionalOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CancelCmConditionalOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CancelCmConditionalOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CancelCmConditionalOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CancelCmConditionalOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CancelCmConditionalOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CancelCmConditionalOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CancelCmConditionalOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *CancelCmConditionalOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CancelCmConditionalOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CancelCmConditionalOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CancelCmConditionalOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *CancelCmConditionalOrderResponse) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CancelCmConditionalOrderResponse) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CancelCmConditionalOrderResponse) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CancelCmConditionalOrderResponse) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *CancelCmConditionalOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelCmConditionalOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelCmConditionalOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelCmConditionalOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CancelCmConditionalOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CancelCmConditionalOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CancelCmConditionalOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CancelCmConditionalOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetActivatePrice

`func (o *CancelCmConditionalOrderResponse) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *CancelCmConditionalOrderResponse) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *CancelCmConditionalOrderResponse) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *CancelCmConditionalOrderResponse) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetPriceRate

`func (o *CancelCmConditionalOrderResponse) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *CancelCmConditionalOrderResponse) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *CancelCmConditionalOrderResponse) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *CancelCmConditionalOrderResponse) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetBookTime

`func (o *CancelCmConditionalOrderResponse) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *CancelCmConditionalOrderResponse) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *CancelCmConditionalOrderResponse) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *CancelCmConditionalOrderResponse) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CancelCmConditionalOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CancelCmConditionalOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CancelCmConditionalOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CancelCmConditionalOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *CancelCmConditionalOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *CancelCmConditionalOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *CancelCmConditionalOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *CancelCmConditionalOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *CancelCmConditionalOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *CancelCmConditionalOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *CancelCmConditionalOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *CancelCmConditionalOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.


[[Back to README]](../README.md)


