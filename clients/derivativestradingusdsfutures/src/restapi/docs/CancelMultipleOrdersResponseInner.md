# CancelMultipleOrdersResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CumQty** | Pointer to **string** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
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
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ActivatePrice** | Pointer to **string** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewCancelMultipleOrdersResponseInner

`func NewCancelMultipleOrdersResponseInner() *CancelMultipleOrdersResponseInner`

NewCancelMultipleOrdersResponseInner instantiates a new CancelMultipleOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelMultipleOrdersResponseInnerWithDefaults

`func NewCancelMultipleOrdersResponseInnerWithDefaults() *CancelMultipleOrdersResponseInner`

NewCancelMultipleOrdersResponseInnerWithDefaults instantiates a new CancelMultipleOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *CancelMultipleOrdersResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *CancelMultipleOrdersResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *CancelMultipleOrdersResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *CancelMultipleOrdersResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumQty

`func (o *CancelMultipleOrdersResponseInner) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *CancelMultipleOrdersResponseInner) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *CancelMultipleOrdersResponseInner) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *CancelMultipleOrdersResponseInner) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetCumQuote

`func (o *CancelMultipleOrdersResponseInner) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *CancelMultipleOrdersResponseInner) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *CancelMultipleOrdersResponseInner) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *CancelMultipleOrdersResponseInner) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *CancelMultipleOrdersResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *CancelMultipleOrdersResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *CancelMultipleOrdersResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *CancelMultipleOrdersResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrderId

`func (o *CancelMultipleOrdersResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelMultipleOrdersResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelMultipleOrdersResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CancelMultipleOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *CancelMultipleOrdersResponseInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *CancelMultipleOrdersResponseInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *CancelMultipleOrdersResponseInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *CancelMultipleOrdersResponseInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *CancelMultipleOrdersResponseInner) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *CancelMultipleOrdersResponseInner) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *CancelMultipleOrdersResponseInner) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *CancelMultipleOrdersResponseInner) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPrice

`func (o *CancelMultipleOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CancelMultipleOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CancelMultipleOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CancelMultipleOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CancelMultipleOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CancelMultipleOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CancelMultipleOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CancelMultipleOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CancelMultipleOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CancelMultipleOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CancelMultipleOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CancelMultipleOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *CancelMultipleOrdersResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CancelMultipleOrdersResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CancelMultipleOrdersResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CancelMultipleOrdersResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStatus

`func (o *CancelMultipleOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelMultipleOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelMultipleOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelMultipleOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *CancelMultipleOrdersResponseInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CancelMultipleOrdersResponseInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CancelMultipleOrdersResponseInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CancelMultipleOrdersResponseInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetClosePosition

`func (o *CancelMultipleOrdersResponseInner) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *CancelMultipleOrdersResponseInner) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *CancelMultipleOrdersResponseInner) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *CancelMultipleOrdersResponseInner) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetSymbol

`func (o *CancelMultipleOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CancelMultipleOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CancelMultipleOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CancelMultipleOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CancelMultipleOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CancelMultipleOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CancelMultipleOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CancelMultipleOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CancelMultipleOrdersResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CancelMultipleOrdersResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CancelMultipleOrdersResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CancelMultipleOrdersResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActivatePrice

`func (o *CancelMultipleOrdersResponseInner) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *CancelMultipleOrdersResponseInner) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *CancelMultipleOrdersResponseInner) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *CancelMultipleOrdersResponseInner) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetPriceRate

`func (o *CancelMultipleOrdersResponseInner) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *CancelMultipleOrdersResponseInner) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *CancelMultipleOrdersResponseInner) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *CancelMultipleOrdersResponseInner) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CancelMultipleOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CancelMultipleOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CancelMultipleOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CancelMultipleOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *CancelMultipleOrdersResponseInner) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *CancelMultipleOrdersResponseInner) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *CancelMultipleOrdersResponseInner) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *CancelMultipleOrdersResponseInner) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceProtect

`func (o *CancelMultipleOrdersResponseInner) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *CancelMultipleOrdersResponseInner) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *CancelMultipleOrdersResponseInner) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *CancelMultipleOrdersResponseInner) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceMatch

`func (o *CancelMultipleOrdersResponseInner) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *CancelMultipleOrdersResponseInner) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *CancelMultipleOrdersResponseInner) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *CancelMultipleOrdersResponseInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *CancelMultipleOrdersResponseInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *CancelMultipleOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *CancelMultipleOrdersResponseInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *CancelMultipleOrdersResponseInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *CancelMultipleOrdersResponseInner) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *CancelMultipleOrdersResponseInner) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *CancelMultipleOrdersResponseInner) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *CancelMultipleOrdersResponseInner) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetCode

`func (o *CancelMultipleOrdersResponseInner) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CancelMultipleOrdersResponseInner) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CancelMultipleOrdersResponseInner) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *CancelMultipleOrdersResponseInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *CancelMultipleOrdersResponseInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CancelMultipleOrdersResponseInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CancelMultipleOrdersResponseInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CancelMultipleOrdersResponseInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to README]](../README.md)


