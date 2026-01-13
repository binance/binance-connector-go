# CurrentAllOpenOrdersResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CumBase** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ActivatePrice** | Pointer to **string** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrentAllOpenOrdersResponseInner

`func NewCurrentAllOpenOrdersResponseInner() *CurrentAllOpenOrdersResponseInner`

NewCurrentAllOpenOrdersResponseInner instantiates a new CurrentAllOpenOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentAllOpenOrdersResponseInnerWithDefaults

`func NewCurrentAllOpenOrdersResponseInnerWithDefaults() *CurrentAllOpenOrdersResponseInner`

NewCurrentAllOpenOrdersResponseInnerWithDefaults instantiates a new CurrentAllOpenOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *CurrentAllOpenOrdersResponseInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *CurrentAllOpenOrdersResponseInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *CurrentAllOpenOrdersResponseInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *CurrentAllOpenOrdersResponseInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *CurrentAllOpenOrdersResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *CurrentAllOpenOrdersResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *CurrentAllOpenOrdersResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *CurrentAllOpenOrdersResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumBase

`func (o *CurrentAllOpenOrdersResponseInner) GetCumBase() string`

GetCumBase returns the CumBase field if non-nil, zero value otherwise.

### GetCumBaseOk

`func (o *CurrentAllOpenOrdersResponseInner) GetCumBaseOk() (*string, bool)`

GetCumBaseOk returns a tuple with the CumBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumBase

`func (o *CurrentAllOpenOrdersResponseInner) SetCumBase(v string)`

SetCumBase sets CumBase field to given value.

### HasCumBase

`func (o *CurrentAllOpenOrdersResponseInner) HasCumBase() bool`

HasCumBase returns a boolean if a field has been set.

### GetExecutedQty

`func (o *CurrentAllOpenOrdersResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *CurrentAllOpenOrdersResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *CurrentAllOpenOrdersResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *CurrentAllOpenOrdersResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrderId

`func (o *CurrentAllOpenOrdersResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CurrentAllOpenOrdersResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CurrentAllOpenOrdersResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CurrentAllOpenOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *CurrentAllOpenOrdersResponseInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *CurrentAllOpenOrdersResponseInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *CurrentAllOpenOrdersResponseInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *CurrentAllOpenOrdersResponseInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *CurrentAllOpenOrdersResponseInner) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *CurrentAllOpenOrdersResponseInner) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *CurrentAllOpenOrdersResponseInner) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPrice

`func (o *CurrentAllOpenOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CurrentAllOpenOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CurrentAllOpenOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CurrentAllOpenOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CurrentAllOpenOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CurrentAllOpenOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CurrentAllOpenOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CurrentAllOpenOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CurrentAllOpenOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CurrentAllOpenOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CurrentAllOpenOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *CurrentAllOpenOrdersResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CurrentAllOpenOrdersResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CurrentAllOpenOrdersResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CurrentAllOpenOrdersResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStatus

`func (o *CurrentAllOpenOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrentAllOpenOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrentAllOpenOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CurrentAllOpenOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *CurrentAllOpenOrdersResponseInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CurrentAllOpenOrdersResponseInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CurrentAllOpenOrdersResponseInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CurrentAllOpenOrdersResponseInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetClosePosition

`func (o *CurrentAllOpenOrdersResponseInner) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *CurrentAllOpenOrdersResponseInner) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *CurrentAllOpenOrdersResponseInner) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *CurrentAllOpenOrdersResponseInner) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetSymbol

`func (o *CurrentAllOpenOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrentAllOpenOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrentAllOpenOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrentAllOpenOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *CurrentAllOpenOrdersResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CurrentAllOpenOrdersResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *CurrentAllOpenOrdersResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CurrentAllOpenOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CurrentAllOpenOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CurrentAllOpenOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CurrentAllOpenOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CurrentAllOpenOrdersResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrentAllOpenOrdersResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CurrentAllOpenOrdersResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActivatePrice

`func (o *CurrentAllOpenOrdersResponseInner) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *CurrentAllOpenOrdersResponseInner) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *CurrentAllOpenOrdersResponseInner) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *CurrentAllOpenOrdersResponseInner) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetPriceRate

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *CurrentAllOpenOrdersResponseInner) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *CurrentAllOpenOrdersResponseInner) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CurrentAllOpenOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CurrentAllOpenOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CurrentAllOpenOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *CurrentAllOpenOrdersResponseInner) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *CurrentAllOpenOrdersResponseInner) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *CurrentAllOpenOrdersResponseInner) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *CurrentAllOpenOrdersResponseInner) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *CurrentAllOpenOrdersResponseInner) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceMatch

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *CurrentAllOpenOrdersResponseInner) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *CurrentAllOpenOrdersResponseInner) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *CurrentAllOpenOrdersResponseInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *CurrentAllOpenOrdersResponseInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *CurrentAllOpenOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *CurrentAllOpenOrdersResponseInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *CurrentAllOpenOrdersResponseInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.


[[Back to README]](../README.md)


