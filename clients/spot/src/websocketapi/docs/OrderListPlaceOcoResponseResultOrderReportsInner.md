# OrderListPlaceOcoResponseResultOrderReportsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**OrigQuoteOrderQty** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
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

### NewOrderListPlaceOcoResponseResultOrderReportsInner

`func NewOrderListPlaceOcoResponseResultOrderReportsInner() *OrderListPlaceOcoResponseResultOrderReportsInner`

NewOrderListPlaceOcoResponseResultOrderReportsInner instantiates a new OrderListPlaceOcoResponseResultOrderReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOcoResponseResultOrderReportsInnerWithDefaults

`func NewOrderListPlaceOcoResponseResultOrderReportsInnerWithDefaults() *OrderListPlaceOcoResponseResultOrderReportsInner`

NewOrderListPlaceOcoResponseResultOrderReportsInnerWithDefaults instantiates a new OrderListPlaceOcoResponseResultOrderReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetWorkingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetUsedSor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OrderListPlaceOcoResponseResultOrderReportsInner) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


