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


[[Back to README]](../README.md)


