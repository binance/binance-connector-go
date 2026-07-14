# OrderStatusResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** | Present only for orders that belong to an order list. | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**OrigQuoteOrderQty** | Pointer to **string** | Always present. Zero if the order type does not use &#x60;quoteOrderQty&#x60;. | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** | Always present. Zero if the order type does not use &#x60;stopPrice&#x60;. | [optional] 
**TrailingDelta** | Pointer to **int64** | Present only if &#x60;trailingDelta&#x60; was set on the order. | [optional] 
**TrailingTime** | Pointer to **int64** | Present only if &#x60;trailingDelta&#x60; was set on the order. | [optional] 
**IcebergQty** | Pointer to **string** | Always present. Zero for non-iceberg orders. | [optional] 
**Time** | Pointer to **int64** | Order placement time. | [optional] 
**UpdateTime** | Pointer to **int64** | Time of the last update to the order. | [optional] 
**IsWorking** | Pointer to **bool** |  | [optional] 
**WorkingTime** | Pointer to **int64** |  | [optional] 
**StrategyId** | Pointer to **int64** | Present only if &#x60;strategyId&#x60; was set on the order. | [optional] 
**StrategyType** | Pointer to **int64** | Present only if &#x60;strategyType&#x60; was set on the order. | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**PreventedMatchId** | Pointer to **int64** | Present only if the order expired due to STP. | [optional] 
**PreventedQuantity** | Pointer to **string** | Present only if the order expired due to STP. | [optional] 
**UsedSor** | Pointer to **bool** | Field that determines whether order used SOR. | [optional] 
**WorkingFloor** | Pointer to **string** | Determines whether the order is being filled by the SOR or by the order book. | [optional] 
**PegPriceType** | Pointer to **string** | Price peg type. Only for pegged orders. | [optional] 
**PegOffsetType** | Pointer to **string** | Price peg offset type. Only for pegged orders, if requested. | [optional] 
**PegOffsetValue** | Pointer to **int64** | Price peg offset value. Only for pegged orders, if requested. | [optional] 
**PeggedPrice** | Pointer to **string** | Current price order is pegged at. Only for pegged orders, once determined. | [optional] 
**ExpiryReason** | Pointer to **string** | Cause of the order&#39;s expiration. Appears when an order has expired. | [optional] 

## Methods

### NewOrderStatusResponseResult

`func NewOrderStatusResponseResult() *OrderStatusResponseResult`

NewOrderStatusResponseResult instantiates a new OrderStatusResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusResponseResultWithDefaults

`func NewOrderStatusResponseResultWithDefaults() *OrderStatusResponseResult`

NewOrderStatusResponseResultWithDefaults instantiates a new OrderStatusResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderStatusResponseResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderStatusResponseResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderStatusResponseResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderStatusResponseResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderStatusResponseResult) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderStatusResponseResult) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderStatusResponseResult) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderStatusResponseResult) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderStatusResponseResult) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderStatusResponseResult) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderStatusResponseResult) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderStatusResponseResult) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderStatusResponseResult) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderStatusResponseResult) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderStatusResponseResult) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderStatusResponseResult) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *OrderStatusResponseResult) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderStatusResponseResult) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderStatusResponseResult) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderStatusResponseResult) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderStatusResponseResult) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderStatusResponseResult) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderStatusResponseResult) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderStatusResponseResult) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderStatusResponseResult) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderStatusResponseResult) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderStatusResponseResult) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderStatusResponseResult) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OrderStatusResponseResult) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OrderStatusResponseResult) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OrderStatusResponseResult) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OrderStatusResponseResult) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderStatusResponseResult) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderStatusResponseResult) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderStatusResponseResult) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderStatusResponseResult) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderStatusResponseResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderStatusResponseResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderStatusResponseResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderStatusResponseResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderStatusResponseResult) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderStatusResponseResult) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderStatusResponseResult) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderStatusResponseResult) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderStatusResponseResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderStatusResponseResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderStatusResponseResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderStatusResponseResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderStatusResponseResult) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderStatusResponseResult) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderStatusResponseResult) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderStatusResponseResult) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderStatusResponseResult) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderStatusResponseResult) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderStatusResponseResult) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderStatusResponseResult) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OrderStatusResponseResult) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OrderStatusResponseResult) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OrderStatusResponseResult) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OrderStatusResponseResult) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OrderStatusResponseResult) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OrderStatusResponseResult) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OrderStatusResponseResult) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OrderStatusResponseResult) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderStatusResponseResult) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderStatusResponseResult) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderStatusResponseResult) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderStatusResponseResult) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetTime

`func (o *OrderStatusResponseResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OrderStatusResponseResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OrderStatusResponseResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OrderStatusResponseResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OrderStatusResponseResult) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OrderStatusResponseResult) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OrderStatusResponseResult) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OrderStatusResponseResult) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetIsWorking

`func (o *OrderStatusResponseResult) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *OrderStatusResponseResult) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *OrderStatusResponseResult) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *OrderStatusResponseResult) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetWorkingTime

`func (o *OrderStatusResponseResult) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *OrderStatusResponseResult) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *OrderStatusResponseResult) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *OrderStatusResponseResult) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.

### GetStrategyId

`func (o *OrderStatusResponseResult) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OrderStatusResponseResult) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OrderStatusResponseResult) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OrderStatusResponseResult) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OrderStatusResponseResult) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OrderStatusResponseResult) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OrderStatusResponseResult) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OrderStatusResponseResult) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderStatusResponseResult) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderStatusResponseResult) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderStatusResponseResult) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderStatusResponseResult) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OrderStatusResponseResult) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OrderStatusResponseResult) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OrderStatusResponseResult) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OrderStatusResponseResult) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OrderStatusResponseResult) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OrderStatusResponseResult) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OrderStatusResponseResult) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OrderStatusResponseResult) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetUsedSor

`func (o *OrderStatusResponseResult) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OrderStatusResponseResult) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OrderStatusResponseResult) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OrderStatusResponseResult) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OrderStatusResponseResult) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OrderStatusResponseResult) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OrderStatusResponseResult) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OrderStatusResponseResult) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OrderStatusResponseResult) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OrderStatusResponseResult) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OrderStatusResponseResult) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OrderStatusResponseResult) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OrderStatusResponseResult) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OrderStatusResponseResult) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OrderStatusResponseResult) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OrderStatusResponseResult) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OrderStatusResponseResult) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OrderStatusResponseResult) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OrderStatusResponseResult) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OrderStatusResponseResult) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OrderStatusResponseResult) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OrderStatusResponseResult) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OrderStatusResponseResult) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OrderStatusResponseResult) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OrderStatusResponseResult) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OrderStatusResponseResult) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OrderStatusResponseResult) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OrderStatusResponseResult) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


