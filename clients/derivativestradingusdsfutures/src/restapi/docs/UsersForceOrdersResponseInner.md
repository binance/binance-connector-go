# UsersForceOrdersResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewUsersForceOrdersResponseInner

`func NewUsersForceOrdersResponseInner() *UsersForceOrdersResponseInner`

NewUsersForceOrdersResponseInner instantiates a new UsersForceOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersForceOrdersResponseInnerWithDefaults

`func NewUsersForceOrdersResponseInnerWithDefaults() *UsersForceOrdersResponseInner`

NewUsersForceOrdersResponseInnerWithDefaults instantiates a new UsersForceOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *UsersForceOrdersResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UsersForceOrdersResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UsersForceOrdersResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UsersForceOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *UsersForceOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UsersForceOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UsersForceOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UsersForceOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStatus

`func (o *UsersForceOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersForceOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersForceOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersForceOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UsersForceOrdersResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UsersForceOrdersResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UsersForceOrdersResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UsersForceOrdersResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *UsersForceOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UsersForceOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UsersForceOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UsersForceOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAvgPrice

`func (o *UsersForceOrdersResponseInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *UsersForceOrdersResponseInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *UsersForceOrdersResponseInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *UsersForceOrdersResponseInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *UsersForceOrdersResponseInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *UsersForceOrdersResponseInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *UsersForceOrdersResponseInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *UsersForceOrdersResponseInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *UsersForceOrdersResponseInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *UsersForceOrdersResponseInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *UsersForceOrdersResponseInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *UsersForceOrdersResponseInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetCumQuote

`func (o *UsersForceOrdersResponseInner) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *UsersForceOrdersResponseInner) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *UsersForceOrdersResponseInner) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *UsersForceOrdersResponseInner) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UsersForceOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UsersForceOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UsersForceOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UsersForceOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UsersForceOrdersResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsersForceOrdersResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsersForceOrdersResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UsersForceOrdersResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReduceOnly

`func (o *UsersForceOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UsersForceOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UsersForceOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UsersForceOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetClosePosition

`func (o *UsersForceOrdersResponseInner) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *UsersForceOrdersResponseInner) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *UsersForceOrdersResponseInner) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *UsersForceOrdersResponseInner) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetSide

`func (o *UsersForceOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UsersForceOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UsersForceOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UsersForceOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *UsersForceOrdersResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UsersForceOrdersResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UsersForceOrdersResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UsersForceOrdersResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *UsersForceOrdersResponseInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UsersForceOrdersResponseInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UsersForceOrdersResponseInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UsersForceOrdersResponseInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetWorkingType

`func (o *UsersForceOrdersResponseInner) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UsersForceOrdersResponseInner) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UsersForceOrdersResponseInner) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UsersForceOrdersResponseInner) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetOrigType

`func (o *UsersForceOrdersResponseInner) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *UsersForceOrdersResponseInner) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *UsersForceOrdersResponseInner) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *UsersForceOrdersResponseInner) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetTime

`func (o *UsersForceOrdersResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UsersForceOrdersResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UsersForceOrdersResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UsersForceOrdersResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UsersForceOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UsersForceOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UsersForceOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UsersForceOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


