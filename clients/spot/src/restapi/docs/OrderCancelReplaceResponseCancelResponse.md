# OrderCancelReplaceResponseCancelResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** | Unless it&#39;s part of an order list, value will be -1 | [optional] 
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

### NewOrderCancelReplaceResponseCancelResponse

`func NewOrderCancelReplaceResponseCancelResponse() *OrderCancelReplaceResponseCancelResponse`

NewOrderCancelReplaceResponseCancelResponse instantiates a new OrderCancelReplaceResponseCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCancelReplaceResponseCancelResponseWithDefaults

`func NewOrderCancelReplaceResponseCancelResponseWithDefaults() *OrderCancelReplaceResponseCancelResponse`

NewOrderCancelReplaceResponseCancelResponseWithDefaults instantiates a new OrderCancelReplaceResponseCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderCancelReplaceResponseCancelResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderCancelReplaceResponseCancelResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderCancelReplaceResponseCancelResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderCancelReplaceResponseCancelResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderCancelReplaceResponseCancelResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderCancelReplaceResponseCancelResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *OrderCancelReplaceResponseCancelResponse) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderCancelReplaceResponseCancelResponse) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderCancelReplaceResponseCancelResponse) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *OrderCancelReplaceResponseCancelResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderCancelReplaceResponseCancelResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderCancelReplaceResponseCancelResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderCancelReplaceResponseCancelResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderCancelReplaceResponseCancelResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderCancelReplaceResponseCancelResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderCancelReplaceResponseCancelResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderCancelReplaceResponseCancelResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OrderCancelReplaceResponseCancelResponse) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OrderCancelReplaceResponseCancelResponse) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OrderCancelReplaceResponseCancelResponse) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OrderCancelReplaceResponseCancelResponse) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OrderCancelReplaceResponseCancelResponse) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderCancelReplaceResponseCancelResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderCancelReplaceResponseCancelResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderCancelReplaceResponseCancelResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderCancelReplaceResponseCancelResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderCancelReplaceResponseCancelResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderCancelReplaceResponseCancelResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderCancelReplaceResponseCancelResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCancelReplaceResponseCancelResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderCancelReplaceResponseCancelResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderCancelReplaceResponseCancelResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderCancelReplaceResponseCancelResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderCancelReplaceResponseCancelResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderCancelReplaceResponseCancelResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderCancelReplaceResponseCancelResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderCancelReplaceResponseCancelResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OrderCancelReplaceResponseCancelResponse) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OrderCancelReplaceResponseCancelResponse) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OrderCancelReplaceResponseCancelResponse) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OrderCancelReplaceResponseCancelResponse) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OrderCancelReplaceResponseCancelResponse) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OrderCancelReplaceResponseCancelResponse) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OrderCancelReplaceResponseCancelResponse) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OrderCancelReplaceResponseCancelResponse) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OrderCancelReplaceResponseCancelResponse) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderCancelReplaceResponseCancelResponse) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderCancelReplaceResponseCancelResponse) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderCancelReplaceResponseCancelResponse) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *OrderCancelReplaceResponseCancelResponse) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OrderCancelReplaceResponseCancelResponse) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OrderCancelReplaceResponseCancelResponse) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OrderCancelReplaceResponseCancelResponse) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OrderCancelReplaceResponseCancelResponse) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OrderCancelReplaceResponseCancelResponse) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OrderCancelReplaceResponseCancelResponse) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OrderCancelReplaceResponseCancelResponse) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OrderCancelReplaceResponseCancelResponse) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OrderCancelReplaceResponseCancelResponse) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OrderCancelReplaceResponseCancelResponse) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OrderCancelReplaceResponseCancelResponse) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetUsedSor

`func (o *OrderCancelReplaceResponseCancelResponse) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OrderCancelReplaceResponseCancelResponse) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OrderCancelReplaceResponseCancelResponse) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OrderCancelReplaceResponseCancelResponse) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OrderCancelReplaceResponseCancelResponse) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OrderCancelReplaceResponseCancelResponse) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OrderCancelReplaceResponseCancelResponse) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OrderCancelReplaceResponseCancelResponse) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OrderCancelReplaceResponseCancelResponse) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OrderCancelReplaceResponseCancelResponse) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OrderCancelReplaceResponseCancelResponse) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OrderCancelReplaceResponseCancelResponse) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OrderCancelReplaceResponseCancelResponse) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OrderCancelReplaceResponseCancelResponse) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OrderCancelReplaceResponseCancelResponse) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OrderCancelReplaceResponseCancelResponse) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OrderCancelReplaceResponseCancelResponse) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OrderCancelReplaceResponseCancelResponse) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OrderCancelReplaceResponseCancelResponse) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


