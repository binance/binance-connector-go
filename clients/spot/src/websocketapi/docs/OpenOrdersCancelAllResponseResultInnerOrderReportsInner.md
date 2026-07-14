# OpenOrdersCancelAllResponseResultInnerOrderReportsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
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

### NewOpenOrdersCancelAllResponseResultInnerOrderReportsInner

`func NewOpenOrdersCancelAllResponseResultInnerOrderReportsInner() *OpenOrdersCancelAllResponseResultInnerOrderReportsInner`

NewOpenOrdersCancelAllResponseResultInnerOrderReportsInner instantiates a new OpenOrdersCancelAllResponseResultInnerOrderReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenOrdersCancelAllResponseResultInnerOrderReportsInnerWithDefaults

`func NewOpenOrdersCancelAllResponseResultInnerOrderReportsInnerWithDefaults() *OpenOrdersCancelAllResponseResultInnerOrderReportsInner`

NewOpenOrdersCancelAllResponseResultInnerOrderReportsInnerWithDefaults instantiates a new OpenOrdersCancelAllResponseResultInnerOrderReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTransactTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrigQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetIcebergQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetPreventedQuantity

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPreventedQuantity() string`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPreventedQuantityOk() (*string, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPreventedQuantity(v string)`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### GetStopPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStrategyType() int64`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetStrategyTypeOk() (*int64, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetStrategyType(v int64)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetTrailingDelta

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTrailingDelta() int64`

GetTrailingDelta returns the TrailingDelta field if non-nil, zero value otherwise.

### GetTrailingDeltaOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTrailingDeltaOk() (*int64, bool)`

GetTrailingDeltaOk returns a tuple with the TrailingDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingDelta

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetTrailingDelta(v int64)`

SetTrailingDelta sets TrailingDelta field to given value.

### HasTrailingDelta

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasTrailingDelta() bool`

HasTrailingDelta returns a boolean if a field has been set.

### GetTrailingTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTrailingTime() int64`

GetTrailingTime returns the TrailingTime field if non-nil, zero value otherwise.

### GetTrailingTimeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetTrailingTimeOk() (*int64, bool)`

GetTrailingTimeOk returns a tuple with the TrailingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetTrailingTime(v int64)`

SetTrailingTime sets TrailingTime field to given value.

### HasTrailingTime

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasTrailingTime() bool`

HasTrailingTime returns a boolean if a field has been set.

### GetUsedSor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetPegPriceType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegPriceType() string`

GetPegPriceType returns the PegPriceType field if non-nil, zero value otherwise.

### GetPegPriceTypeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegPriceTypeOk() (*string, bool)`

GetPegPriceTypeOk returns a tuple with the PegPriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegPriceType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPegPriceType(v string)`

SetPegPriceType sets PegPriceType field to given value.

### HasPegPriceType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPegPriceType() bool`

HasPegPriceType returns a boolean if a field has been set.

### GetPegOffsetType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegOffsetType() string`

GetPegOffsetType returns the PegOffsetType field if non-nil, zero value otherwise.

### GetPegOffsetTypeOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegOffsetTypeOk() (*string, bool)`

GetPegOffsetTypeOk returns a tuple with the PegOffsetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPegOffsetType(v string)`

SetPegOffsetType sets PegOffsetType field to given value.

### HasPegOffsetType

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPegOffsetType() bool`

HasPegOffsetType returns a boolean if a field has been set.

### GetPegOffsetValue

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegOffsetValue() int64`

GetPegOffsetValue returns the PegOffsetValue field if non-nil, zero value otherwise.

### GetPegOffsetValueOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPegOffsetValueOk() (*int64, bool)`

GetPegOffsetValueOk returns a tuple with the PegOffsetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegOffsetValue

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPegOffsetValue(v int64)`

SetPegOffsetValue sets PegOffsetValue field to given value.

### HasPegOffsetValue

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPegOffsetValue() bool`

HasPegOffsetValue returns a boolean if a field has been set.

### GetPeggedPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPeggedPrice() string`

GetPeggedPrice returns the PeggedPrice field if non-nil, zero value otherwise.

### GetPeggedPriceOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetPeggedPriceOk() (*string, bool)`

GetPeggedPriceOk returns a tuple with the PeggedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeggedPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetPeggedPrice(v string)`

SetPeggedPrice sets PeggedPrice field to given value.

### HasPeggedPrice

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasPeggedPrice() bool`

HasPeggedPrice returns a boolean if a field has been set.

### GetExpiryReason

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetExpiryReason() string`

GetExpiryReason returns the ExpiryReason field if non-nil, zero value otherwise.

### GetExpiryReasonOk

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) GetExpiryReasonOk() (*string, bool)`

GetExpiryReasonOk returns a tuple with the ExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryReason

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) SetExpiryReason(v string)`

SetExpiryReason sets ExpiryReason field to given value.

### HasExpiryReason

`func (o *OpenOrdersCancelAllResponseResultInnerOrderReportsInner) HasExpiryReason() bool`

HasExpiryReason returns a boolean if a field has been set.


[[Back to README]](../README.md)


