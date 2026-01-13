# OrderAmendKeepPriorityResponseResultAmendedOrder

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

## Methods

### NewOrderAmendKeepPriorityResponseResultAmendedOrder

`func NewOrderAmendKeepPriorityResponseResultAmendedOrder() *OrderAmendKeepPriorityResponseResultAmendedOrder`

NewOrderAmendKeepPriorityResponseResultAmendedOrder instantiates a new OrderAmendKeepPriorityResponseResultAmendedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseResultAmendedOrderWithDefaults

`func NewOrderAmendKeepPriorityResponseResultAmendedOrderWithDefaults() *OrderAmendKeepPriorityResponseResultAmendedOrder`

NewOrderAmendKeepPriorityResponseResultAmendedOrderWithDefaults instantiates a new OrderAmendKeepPriorityResponseResultAmendedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetPreventedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPreventedQty() string`

GetPreventedQty returns the PreventedQty field if non-nil, zero value otherwise.

### GetPreventedQtyOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetPreventedQtyOk() (*string, bool)`

GetPreventedQtyOk returns a tuple with the PreventedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetPreventedQty(v string)`

SetPreventedQty sets PreventedQty field to given value.

### HasPreventedQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasPreventedQty() bool`

HasPreventedQty returns a boolean if a field has been set.

### GetQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQuoteOrderQty() string`

GetQuoteOrderQty returns the QuoteOrderQty field if non-nil, zero value otherwise.

### GetQuoteOrderQtyOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetQuoteOrderQtyOk() (*string, bool)`

GetQuoteOrderQtyOk returns a tuple with the QuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetQuoteOrderQty(v string)`

SetQuoteOrderQty sets QuoteOrderQty field to given value.

### HasQuoteOrderQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasQuoteOrderQty() bool`

HasQuoteOrderQty returns a boolean if a field has been set.

### GetCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetCumulativeQuoteQty() string`

GetCumulativeQuoteQty returns the CumulativeQuoteQty field if non-nil, zero value otherwise.

### GetCumulativeQuoteQtyOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetCumulativeQuoteQtyOk() (*string, bool)`

GetCumulativeQuoteQtyOk returns a tuple with the CumulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetCumulativeQuoteQty(v string)`

SetCumulativeQuoteQty sets CumulativeQuoteQty field to given value.

### HasCumulativeQuoteQty

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasCumulativeQuoteQty() bool`

HasCumulativeQuoteQty returns a boolean if a field has been set.

### GetStatus

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetWorkingTime

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *OrderAmendKeepPriorityResponseResultAmendedOrder) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.


[[Back to README]](../README.md)


