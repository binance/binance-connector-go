# OrderListPlaceOtocoResponseResultOrderReportsInner

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

### NewOrderListPlaceOtocoResponseResultOrderReportsInner

`func NewOrderListPlaceOtocoResponseResultOrderReportsInner() *OrderListPlaceOtocoResponseResultOrderReportsInner`

NewOrderListPlaceOtocoResponseResultOrderReportsInner instantiates a new OrderListPlaceOtocoResponseResultOrderReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListPlaceOtocoResponseResultOrderReportsInnerWithDefaults

`func NewOrderListPlaceOtocoResponseResultOrderReportsInnerWithDefaults() *OrderListPlaceOtocoResponseResultOrderReportsInner`

NewOrderListPlaceOtocoResponseResultOrderReportsInnerWithDefaults instantiates a new OrderListPlaceOtocoResponseResultOrderReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetWorkingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetUsedSor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OrderListPlaceOtocoResponseResultOrderReportsInner) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


