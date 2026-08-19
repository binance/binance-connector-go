# EquityTradeHistoryResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** | Execution (per-fill) id. | [optional] 
**OrderId** | Pointer to **string** | The owning order&#39;s id. | [optional] 
**Symbol** | Pointer to **string** | US-equity ticker. | [optional] 
**Quote** | Pointer to **string** | Quote asset. | [optional] 
**Side** | Pointer to **string** | &#x60;BUY&#x60; / &#x60;SELL&#x60;. | [optional] 
**OrderType** | Pointer to **string** | &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | [optional] 
**Price** | Pointer to **string** | Execution price (USD). | [optional] 
**Qty** | Pointer to **string** | Executed quantity. | [optional] 
**Total** | Pointer to **string** | Notional of this execution (&#x60;qty × price&#x60;). | [optional] 
**ExecutionAt** | Pointer to **int64** | Execution time (ms epoch). | [optional] 
**UpdatedAt** | Pointer to **int64** | Last update time (ms epoch). | [optional] 

## Methods

### NewEquityTradeHistoryResponseRowsInner

`func NewEquityTradeHistoryResponseRowsInner() *EquityTradeHistoryResponseRowsInner`

NewEquityTradeHistoryResponseRowsInner instantiates a new EquityTradeHistoryResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityTradeHistoryResponseRowsInnerWithDefaults

`func NewEquityTradeHistoryResponseRowsInnerWithDefaults() *EquityTradeHistoryResponseRowsInner`

NewEquityTradeHistoryResponseRowsInnerWithDefaults instantiates a new EquityTradeHistoryResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *EquityTradeHistoryResponseRowsInner) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EquityTradeHistoryResponseRowsInner) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EquityTradeHistoryResponseRowsInner) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EquityTradeHistoryResponseRowsInner) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetOrderId

`func (o *EquityTradeHistoryResponseRowsInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *EquityTradeHistoryResponseRowsInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *EquityTradeHistoryResponseRowsInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *EquityTradeHistoryResponseRowsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *EquityTradeHistoryResponseRowsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EquityTradeHistoryResponseRowsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EquityTradeHistoryResponseRowsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EquityTradeHistoryResponseRowsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetQuote

`func (o *EquityTradeHistoryResponseRowsInner) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *EquityTradeHistoryResponseRowsInner) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *EquityTradeHistoryResponseRowsInner) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *EquityTradeHistoryResponseRowsInner) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetSide

`func (o *EquityTradeHistoryResponseRowsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquityTradeHistoryResponseRowsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquityTradeHistoryResponseRowsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquityTradeHistoryResponseRowsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *EquityTradeHistoryResponseRowsInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *EquityTradeHistoryResponseRowsInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *EquityTradeHistoryResponseRowsInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *EquityTradeHistoryResponseRowsInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPrice

`func (o *EquityTradeHistoryResponseRowsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *EquityTradeHistoryResponseRowsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *EquityTradeHistoryResponseRowsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *EquityTradeHistoryResponseRowsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *EquityTradeHistoryResponseRowsInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *EquityTradeHistoryResponseRowsInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *EquityTradeHistoryResponseRowsInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *EquityTradeHistoryResponseRowsInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetTotal

`func (o *EquityTradeHistoryResponseRowsInner) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EquityTradeHistoryResponseRowsInner) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EquityTradeHistoryResponseRowsInner) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EquityTradeHistoryResponseRowsInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetExecutionAt

`func (o *EquityTradeHistoryResponseRowsInner) GetExecutionAt() int64`

GetExecutionAt returns the ExecutionAt field if non-nil, zero value otherwise.

### GetExecutionAtOk

`func (o *EquityTradeHistoryResponseRowsInner) GetExecutionAtOk() (*int64, bool)`

GetExecutionAtOk returns a tuple with the ExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionAt

`func (o *EquityTradeHistoryResponseRowsInner) SetExecutionAt(v int64)`

SetExecutionAt sets ExecutionAt field to given value.

### HasExecutionAt

`func (o *EquityTradeHistoryResponseRowsInner) HasExecutionAt() bool`

HasExecutionAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EquityTradeHistoryResponseRowsInner) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EquityTradeHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EquityTradeHistoryResponseRowsInner) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EquityTradeHistoryResponseRowsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to README]](../README.md)


