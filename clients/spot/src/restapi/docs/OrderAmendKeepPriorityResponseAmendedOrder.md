# OrderAmendKeepPriorityResponseAmendedOrder

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**PreventedQty** | Pointer to **string** |  | [optional] 
**QuoteOrderQty** | Pointer to **string** |  | [optional] 
**CumulativeQuoteQty** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**WorkingTime** | Pointer to **int64** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**IcebergQty** | Pointer to **string** | Appears only if the parameter icebergQty was sent in the request. | [optional] 
**PreventedMatchId** | Pointer to **int64** | Appears only if the order expired due to STP. | [optional] 
**PreventedQuantity** | Pointer to **string** | Order quantity that expired due to STP. | [optional] 
**StopPrice** | Pointer to **string** | Appears for STOP_LOSS, TAKE_PROFIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT orders. | [optional] 
**StrategyId** | Pointer to **int64** | Appears only if the strategyId parameter was provided upon order placement. | [optional] 
**StrategyType** | Pointer to **int64** | Appears only if the strategyType parameter was provided upon order placement. | [optional] 
**TrailingDelta** | Pointer to **int64** | Delta price change required before order activation. | [optional] 
**TrailingTime** | Pointer to **int64** | Time when the trailing order is now active and tracking price changes. | [optional] 
**UsedSor** | Pointer to **bool** | Field that determines whether order used SOR. | [optional] 
**WorkingFloor** | Pointer to **string** | Determines whether the order is being filled by the SOR or by the order book. | [optional] 
**PegPriceType** | Pointer to **string** | Price peg type. Only for pegged orders. | [optional] 
**PegOffsetType** | Pointer to **string** | Price peg offset type. Only for pegged orders, if requested. | [optional] 
**PegOffsetValue** | Pointer to **int64** | Price peg offset value. Only for pegged orders, if requested. | [optional] 
**PeggedPrice** | Pointer to **string** | Current price order is pegged at. Only for pegged orders, once determined. | [optional] 
**ExpiryReason** | Pointer to **string** | Cause of the order&#39;s expiration. Appears when an order has expired. | [optional] 

## Methods

### NewOrderAmendKeepPriorityResponseAmendedOrder

`func NewOrderAmendKeepPriorityResponseAmendedOrder() *OrderAmendKeepPriorityResponseAmendedOrder`

NewOrderAmendKeepPriorityResponseAmendedOrder instantiates a new OrderAmendKeepPriorityResponseAmendedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseAmendedOrderWithDefaults

`func NewOrderAmendKeepPriorityResponseAmendedOrderWithDefaults() *OrderAmendKeepPriorityResponseAmendedOrder`

NewOrderAmendKeepPriorityResponseAmendedOrderWithDefaults instantiates a new OrderAmendKeepPriorityResponseAmendedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetPreventedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQty() string`

GetPreventedQty returns the PreventedQty field if non-nil, zero value otherwise.

### GetPreventedQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQtyOk() (*string, bool)`

GetPreventedQtyOk returns a tuple with the PreventedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPreventedQty(v string)`

SetPreventedQty sets PreventedQty field to given value.

### HasPreventedQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPreventedQty() bool`

HasPreventedQty returns a boolean if a field has been set.

### GetQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQuoteOrderQty() string`

GetQuoteOrderQty returns the QuoteOrderQty field if non-nil, zero value otherwise.

### GetQuoteOrderQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetQuoteOrderQtyOk() (*string, bool)`

GetQuoteOrderQtyOk returns a tuple with the QuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetQuoteOrderQty(v string)`

SetQuoteOrderQty sets QuoteOrderQty field to given value.

### HasQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasQuoteOrderQty() bool`

HasQuoteOrderQty returns a boolean if a field has been set.

### GetCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetCumulativeQuoteQty() string`

GetCumulativeQuoteQty returns the CumulativeQuoteQty field if non-nil, zero value otherwise.

### GetCumulativeQuoteQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetCumulativeQuoteQtyOk() (*string, bool)`

GetCumulativeQuoteQtyOk returns a tuple with the CumulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetCumulativeQuoteQty(v string)`

SetCumulativeQuoteQty sets CumulativeQuoteQty field to given value.

### HasCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasCumulativeQuoteQty() bool`

HasCumulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetWorkingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetUsedSor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OrderAmendKeepPriorityResponseAmendedOrder) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


