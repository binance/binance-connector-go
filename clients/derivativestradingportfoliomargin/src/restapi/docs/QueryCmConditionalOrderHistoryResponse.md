# QueryCmConditionalOrderHistoryResponse

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
**OrderId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BookTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**TriggerTime** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ActivatePrice** | Pointer to **string** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryCmConditionalOrderHistoryResponse

`func NewQueryCmConditionalOrderHistoryResponse() *QueryCmConditionalOrderHistoryResponse`

NewQueryCmConditionalOrderHistoryResponse instantiates a new QueryCmConditionalOrderHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCmConditionalOrderHistoryResponseWithDefaults

`func NewQueryCmConditionalOrderHistoryResponseWithDefaults() *QueryCmConditionalOrderHistoryResponse`

NewQueryCmConditionalOrderHistoryResponseWithDefaults instantiates a new QueryCmConditionalOrderHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewClientStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) GetNewClientStrategyId() string`

GetNewClientStrategyId returns the NewClientStrategyId field if non-nil, zero value otherwise.

### GetNewClientStrategyIdOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetNewClientStrategyIdOk() (*string, bool)`

GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) SetNewClientStrategyId(v string)`

SetNewClientStrategyId sets NewClientStrategyId field to given value.

### HasNewClientStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) HasNewClientStrategyId() bool`

HasNewClientStrategyId returns a boolean if a field has been set.

### GetStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *QueryCmConditionalOrderHistoryResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyStatus

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyStatus() string`

GetStrategyStatus returns the StrategyStatus field if non-nil, zero value otherwise.

### GetStrategyStatusOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyStatusOk() (*string, bool)`

GetStrategyStatusOk returns a tuple with the StrategyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyStatus

`func (o *QueryCmConditionalOrderHistoryResponse) SetStrategyStatus(v string)`

SetStrategyStatus sets StrategyStatus field to given value.

### HasStrategyStatus

`func (o *QueryCmConditionalOrderHistoryResponse) HasStrategyStatus() bool`

HasStrategyStatus returns a boolean if a field has been set.

### GetStrategyType

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *QueryCmConditionalOrderHistoryResponse) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *QueryCmConditionalOrderHistoryResponse) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetOrigQty

`func (o *QueryCmConditionalOrderHistoryResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *QueryCmConditionalOrderHistoryResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *QueryCmConditionalOrderHistoryResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *QueryCmConditionalOrderHistoryResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryCmConditionalOrderHistoryResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryCmConditionalOrderHistoryResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryCmConditionalOrderHistoryResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryCmConditionalOrderHistoryResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryCmConditionalOrderHistoryResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *QueryCmConditionalOrderHistoryResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryCmConditionalOrderHistoryResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryCmConditionalOrderHistoryResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryCmConditionalOrderHistoryResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryCmConditionalOrderHistoryResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryCmConditionalOrderHistoryResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *QueryCmConditionalOrderHistoryResponse) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *QueryCmConditionalOrderHistoryResponse) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *QueryCmConditionalOrderHistoryResponse) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryCmConditionalOrderHistoryResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryCmConditionalOrderHistoryResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryCmConditionalOrderHistoryResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryCmConditionalOrderHistoryResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryCmConditionalOrderHistoryResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryCmConditionalOrderHistoryResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *QueryCmConditionalOrderHistoryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryCmConditionalOrderHistoryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryCmConditionalOrderHistoryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBookTime

`func (o *QueryCmConditionalOrderHistoryResponse) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *QueryCmConditionalOrderHistoryResponse) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *QueryCmConditionalOrderHistoryResponse) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryCmConditionalOrderHistoryResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryCmConditionalOrderHistoryResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryCmConditionalOrderHistoryResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *QueryCmConditionalOrderHistoryResponse) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *QueryCmConditionalOrderHistoryResponse) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *QueryCmConditionalOrderHistoryResponse) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryCmConditionalOrderHistoryResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryCmConditionalOrderHistoryResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryCmConditionalOrderHistoryResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *QueryCmConditionalOrderHistoryResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryCmConditionalOrderHistoryResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryCmConditionalOrderHistoryResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActivatePrice

`func (o *QueryCmConditionalOrderHistoryResponse) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *QueryCmConditionalOrderHistoryResponse) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *QueryCmConditionalOrderHistoryResponse) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetPriceRate

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *QueryCmConditionalOrderHistoryResponse) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *QueryCmConditionalOrderHistoryResponse) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetWorkingType

`func (o *QueryCmConditionalOrderHistoryResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *QueryCmConditionalOrderHistoryResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *QueryCmConditionalOrderHistoryResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *QueryCmConditionalOrderHistoryResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *QueryCmConditionalOrderHistoryResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryCmConditionalOrderHistoryResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryCmConditionalOrderHistoryResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryCmConditionalOrderHistoryResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.


[[Back to README]](../README.md)


