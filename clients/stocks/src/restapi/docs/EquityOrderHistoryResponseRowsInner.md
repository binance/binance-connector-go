# EquityOrderHistoryResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Equity order id. | [optional] 
**Symbol** | Pointer to **string** | US-equity ticker. | [optional] 
**Quote** | Pointer to **string** | Quote asset (e.g. &#x60;USDC&#x60;). | [optional] 
**Side** | Pointer to **string** | &#x60;BUY&#x60; / &#x60;SELL&#x60;. | [optional] 
**OrderType** | Pointer to **string** | &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | [optional] 
**LimitPrice** | Pointer to **NullableString** | Limit price (USD). Non-null for &#x60;LIMIT&#x60; orders, &#x60;null&#x60; for &#x60;MARKET&#x60;. | [optional] 
**AvgFilledPrice** | Pointer to **NullableString** | Average fill price (USD). &#x60;null&#x60; until the first fill. For &#x60;MARKET&#x60; orders this is the only meaningful price field. | [optional] 
**Qty** | Pointer to **NullableString** | Requested quantity. &#x60;null&#x60; for &#x60;BUY MARKET&#x60; (use &#x60;notional&#x60; instead). | [optional] 
**Notional** | Pointer to **NullableString** | Requested notional. Non-null for &#x60;BUY MARKET&#x60;; &#x60;null&#x60; otherwise. | [optional] 
**FilledQty** | Pointer to **string** | Cumulative filled quantity. | [optional] 
**FilledTotal** | Pointer to **NullableString** | Cumulative filled notional. Populated only for &#x60;BUY MARKET&#x60;. | [optional] 
**Fee** | Pointer to **string** | Total commission fee (USD). | [optional] 
**Session** | Pointer to **NullableString** | Trading session the order was placed under: &#x60;RTH&#x60; / &#x60;EXTENDED&#x60; / &#x60;24H&#x60;. &#x60;null&#x60; for &#x60;MARKET&#x60; orders. | [optional] 
**Status** | Pointer to **string** | Order lifecycle status — one of &#x60;NEW&#x60; / &#x60;ACCEPTED&#x60; / &#x60;PARTIALLY_FILLED&#x60; / &#x60;FILLED&#x60; / &#x60;CANCELED&#x60; / &#x60;EXPIRED&#x60; / &#x60;REJECTED&#x60;. | [optional] 
**CreatedAt** | Pointer to **int64** | Order creation time (ms epoch). | [optional] 
**UpdatedAt** | Pointer to **int64** | Last update time (ms epoch). | [optional] 

## Methods

### NewEquityOrderHistoryResponseRowsInner

`func NewEquityOrderHistoryResponseRowsInner() *EquityOrderHistoryResponseRowsInner`

NewEquityOrderHistoryResponseRowsInner instantiates a new EquityOrderHistoryResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityOrderHistoryResponseRowsInnerWithDefaults

`func NewEquityOrderHistoryResponseRowsInnerWithDefaults() *EquityOrderHistoryResponseRowsInner`

NewEquityOrderHistoryResponseRowsInnerWithDefaults instantiates a new EquityOrderHistoryResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *EquityOrderHistoryResponseRowsInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *EquityOrderHistoryResponseRowsInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *EquityOrderHistoryResponseRowsInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *EquityOrderHistoryResponseRowsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *EquityOrderHistoryResponseRowsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EquityOrderHistoryResponseRowsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EquityOrderHistoryResponseRowsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EquityOrderHistoryResponseRowsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetQuote

`func (o *EquityOrderHistoryResponseRowsInner) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *EquityOrderHistoryResponseRowsInner) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *EquityOrderHistoryResponseRowsInner) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *EquityOrderHistoryResponseRowsInner) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetSide

`func (o *EquityOrderHistoryResponseRowsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquityOrderHistoryResponseRowsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquityOrderHistoryResponseRowsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquityOrderHistoryResponseRowsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *EquityOrderHistoryResponseRowsInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *EquityOrderHistoryResponseRowsInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *EquityOrderHistoryResponseRowsInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *EquityOrderHistoryResponseRowsInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetLimitPrice

`func (o *EquityOrderHistoryResponseRowsInner) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *EquityOrderHistoryResponseRowsInner) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *EquityOrderHistoryResponseRowsInner) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *EquityOrderHistoryResponseRowsInner) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### SetLimitPriceNil

`func (o *EquityOrderHistoryResponseRowsInner) SetLimitPriceNil(b bool)`

 SetLimitPriceNil sets the value for LimitPrice to be an explicit nil

### UnsetLimitPrice
`func (o *EquityOrderHistoryResponseRowsInner) UnsetLimitPrice()`

UnsetLimitPrice ensures that no value is present for LimitPrice, not even an explicit nil
### GetAvgFilledPrice

`func (o *EquityOrderHistoryResponseRowsInner) GetAvgFilledPrice() string`

GetAvgFilledPrice returns the AvgFilledPrice field if non-nil, zero value otherwise.

### GetAvgFilledPriceOk

`func (o *EquityOrderHistoryResponseRowsInner) GetAvgFilledPriceOk() (*string, bool)`

GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgFilledPrice

`func (o *EquityOrderHistoryResponseRowsInner) SetAvgFilledPrice(v string)`

SetAvgFilledPrice sets AvgFilledPrice field to given value.

### HasAvgFilledPrice

`func (o *EquityOrderHistoryResponseRowsInner) HasAvgFilledPrice() bool`

HasAvgFilledPrice returns a boolean if a field has been set.

### SetAvgFilledPriceNil

`func (o *EquityOrderHistoryResponseRowsInner) SetAvgFilledPriceNil(b bool)`

 SetAvgFilledPriceNil sets the value for AvgFilledPrice to be an explicit nil

### UnsetAvgFilledPrice
`func (o *EquityOrderHistoryResponseRowsInner) UnsetAvgFilledPrice()`

UnsetAvgFilledPrice ensures that no value is present for AvgFilledPrice, not even an explicit nil
### GetQty

`func (o *EquityOrderHistoryResponseRowsInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *EquityOrderHistoryResponseRowsInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *EquityOrderHistoryResponseRowsInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *EquityOrderHistoryResponseRowsInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *EquityOrderHistoryResponseRowsInner) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *EquityOrderHistoryResponseRowsInner) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetNotional

`func (o *EquityOrderHistoryResponseRowsInner) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *EquityOrderHistoryResponseRowsInner) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *EquityOrderHistoryResponseRowsInner) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *EquityOrderHistoryResponseRowsInner) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### SetNotionalNil

`func (o *EquityOrderHistoryResponseRowsInner) SetNotionalNil(b bool)`

 SetNotionalNil sets the value for Notional to be an explicit nil

### UnsetNotional
`func (o *EquityOrderHistoryResponseRowsInner) UnsetNotional()`

UnsetNotional ensures that no value is present for Notional, not even an explicit nil
### GetFilledQty

`func (o *EquityOrderHistoryResponseRowsInner) GetFilledQty() string`

GetFilledQty returns the FilledQty field if non-nil, zero value otherwise.

### GetFilledQtyOk

`func (o *EquityOrderHistoryResponseRowsInner) GetFilledQtyOk() (*string, bool)`

GetFilledQtyOk returns a tuple with the FilledQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQty

`func (o *EquityOrderHistoryResponseRowsInner) SetFilledQty(v string)`

SetFilledQty sets FilledQty field to given value.

### HasFilledQty

`func (o *EquityOrderHistoryResponseRowsInner) HasFilledQty() bool`

HasFilledQty returns a boolean if a field has been set.

### GetFilledTotal

`func (o *EquityOrderHistoryResponseRowsInner) GetFilledTotal() string`

GetFilledTotal returns the FilledTotal field if non-nil, zero value otherwise.

### GetFilledTotalOk

`func (o *EquityOrderHistoryResponseRowsInner) GetFilledTotalOk() (*string, bool)`

GetFilledTotalOk returns a tuple with the FilledTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTotal

`func (o *EquityOrderHistoryResponseRowsInner) SetFilledTotal(v string)`

SetFilledTotal sets FilledTotal field to given value.

### HasFilledTotal

`func (o *EquityOrderHistoryResponseRowsInner) HasFilledTotal() bool`

HasFilledTotal returns a boolean if a field has been set.

### SetFilledTotalNil

`func (o *EquityOrderHistoryResponseRowsInner) SetFilledTotalNil(b bool)`

 SetFilledTotalNil sets the value for FilledTotal to be an explicit nil

### UnsetFilledTotal
`func (o *EquityOrderHistoryResponseRowsInner) UnsetFilledTotal()`

UnsetFilledTotal ensures that no value is present for FilledTotal, not even an explicit nil
### GetFee

`func (o *EquityOrderHistoryResponseRowsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EquityOrderHistoryResponseRowsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EquityOrderHistoryResponseRowsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EquityOrderHistoryResponseRowsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSession

`func (o *EquityOrderHistoryResponseRowsInner) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *EquityOrderHistoryResponseRowsInner) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *EquityOrderHistoryResponseRowsInner) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *EquityOrderHistoryResponseRowsInner) HasSession() bool`

HasSession returns a boolean if a field has been set.

### SetSessionNil

`func (o *EquityOrderHistoryResponseRowsInner) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *EquityOrderHistoryResponseRowsInner) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetStatus

`func (o *EquityOrderHistoryResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquityOrderHistoryResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquityOrderHistoryResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquityOrderHistoryResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EquityOrderHistoryResponseRowsInner) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EquityOrderHistoryResponseRowsInner) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EquityOrderHistoryResponseRowsInner) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EquityOrderHistoryResponseRowsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EquityOrderHistoryResponseRowsInner) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EquityOrderHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EquityOrderHistoryResponseRowsInner) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EquityOrderHistoryResponseRowsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to README]](../README.md)


