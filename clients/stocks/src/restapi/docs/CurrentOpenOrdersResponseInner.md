# CurrentOpenOrdersResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Equity order id. | [optional] 
**Symbol** | Pointer to **string** | US-equity ticker. | [optional] 
**Quote** | Pointer to **string** | Quote asset (e.g. &#x60;USDC&#x60;). | [optional] 
**Side** | Pointer to **string** | &#x60;BUY&#x60; / &#x60;SELL&#x60;. | [optional] 
**OrderType** | Pointer to **string** | &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | [optional] 
**LimitPrice** | Pointer to **NullableString** | Limit price (USD). Non-null for &#x60;LIMIT&#x60; orders, &#x60;null&#x60; for &#x60;MARKET&#x60;. | [optional] 
**AvgFilledPrice** | Pointer to **NullableString** | Average fill price (USD). &#x60;null&#x60; until the first fill. | [optional] 
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

### NewCurrentOpenOrdersResponseInner

`func NewCurrentOpenOrdersResponseInner() *CurrentOpenOrdersResponseInner`

NewCurrentOpenOrdersResponseInner instantiates a new CurrentOpenOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentOpenOrdersResponseInnerWithDefaults

`func NewCurrentOpenOrdersResponseInnerWithDefaults() *CurrentOpenOrdersResponseInner`

NewCurrentOpenOrdersResponseInnerWithDefaults instantiates a new CurrentOpenOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CurrentOpenOrdersResponseInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CurrentOpenOrdersResponseInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CurrentOpenOrdersResponseInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CurrentOpenOrdersResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *CurrentOpenOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrentOpenOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrentOpenOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrentOpenOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetQuote

`func (o *CurrentOpenOrdersResponseInner) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *CurrentOpenOrdersResponseInner) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *CurrentOpenOrdersResponseInner) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *CurrentOpenOrdersResponseInner) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetSide

`func (o *CurrentOpenOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CurrentOpenOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CurrentOpenOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CurrentOpenOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *CurrentOpenOrdersResponseInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *CurrentOpenOrdersResponseInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *CurrentOpenOrdersResponseInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *CurrentOpenOrdersResponseInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetLimitPrice

`func (o *CurrentOpenOrdersResponseInner) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *CurrentOpenOrdersResponseInner) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *CurrentOpenOrdersResponseInner) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *CurrentOpenOrdersResponseInner) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### SetLimitPriceNil

`func (o *CurrentOpenOrdersResponseInner) SetLimitPriceNil(b bool)`

 SetLimitPriceNil sets the value for LimitPrice to be an explicit nil

### UnsetLimitPrice
`func (o *CurrentOpenOrdersResponseInner) UnsetLimitPrice()`

UnsetLimitPrice ensures that no value is present for LimitPrice, not even an explicit nil
### GetAvgFilledPrice

`func (o *CurrentOpenOrdersResponseInner) GetAvgFilledPrice() string`

GetAvgFilledPrice returns the AvgFilledPrice field if non-nil, zero value otherwise.

### GetAvgFilledPriceOk

`func (o *CurrentOpenOrdersResponseInner) GetAvgFilledPriceOk() (*string, bool)`

GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgFilledPrice

`func (o *CurrentOpenOrdersResponseInner) SetAvgFilledPrice(v string)`

SetAvgFilledPrice sets AvgFilledPrice field to given value.

### HasAvgFilledPrice

`func (o *CurrentOpenOrdersResponseInner) HasAvgFilledPrice() bool`

HasAvgFilledPrice returns a boolean if a field has been set.

### SetAvgFilledPriceNil

`func (o *CurrentOpenOrdersResponseInner) SetAvgFilledPriceNil(b bool)`

 SetAvgFilledPriceNil sets the value for AvgFilledPrice to be an explicit nil

### UnsetAvgFilledPrice
`func (o *CurrentOpenOrdersResponseInner) UnsetAvgFilledPrice()`

UnsetAvgFilledPrice ensures that no value is present for AvgFilledPrice, not even an explicit nil
### GetQty

`func (o *CurrentOpenOrdersResponseInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *CurrentOpenOrdersResponseInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *CurrentOpenOrdersResponseInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *CurrentOpenOrdersResponseInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *CurrentOpenOrdersResponseInner) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *CurrentOpenOrdersResponseInner) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetNotional

`func (o *CurrentOpenOrdersResponseInner) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *CurrentOpenOrdersResponseInner) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *CurrentOpenOrdersResponseInner) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *CurrentOpenOrdersResponseInner) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### SetNotionalNil

`func (o *CurrentOpenOrdersResponseInner) SetNotionalNil(b bool)`

 SetNotionalNil sets the value for Notional to be an explicit nil

### UnsetNotional
`func (o *CurrentOpenOrdersResponseInner) UnsetNotional()`

UnsetNotional ensures that no value is present for Notional, not even an explicit nil
### GetFilledQty

`func (o *CurrentOpenOrdersResponseInner) GetFilledQty() string`

GetFilledQty returns the FilledQty field if non-nil, zero value otherwise.

### GetFilledQtyOk

`func (o *CurrentOpenOrdersResponseInner) GetFilledQtyOk() (*string, bool)`

GetFilledQtyOk returns a tuple with the FilledQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQty

`func (o *CurrentOpenOrdersResponseInner) SetFilledQty(v string)`

SetFilledQty sets FilledQty field to given value.

### HasFilledQty

`func (o *CurrentOpenOrdersResponseInner) HasFilledQty() bool`

HasFilledQty returns a boolean if a field has been set.

### GetFilledTotal

`func (o *CurrentOpenOrdersResponseInner) GetFilledTotal() string`

GetFilledTotal returns the FilledTotal field if non-nil, zero value otherwise.

### GetFilledTotalOk

`func (o *CurrentOpenOrdersResponseInner) GetFilledTotalOk() (*string, bool)`

GetFilledTotalOk returns a tuple with the FilledTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTotal

`func (o *CurrentOpenOrdersResponseInner) SetFilledTotal(v string)`

SetFilledTotal sets FilledTotal field to given value.

### HasFilledTotal

`func (o *CurrentOpenOrdersResponseInner) HasFilledTotal() bool`

HasFilledTotal returns a boolean if a field has been set.

### SetFilledTotalNil

`func (o *CurrentOpenOrdersResponseInner) SetFilledTotalNil(b bool)`

 SetFilledTotalNil sets the value for FilledTotal to be an explicit nil

### UnsetFilledTotal
`func (o *CurrentOpenOrdersResponseInner) UnsetFilledTotal()`

UnsetFilledTotal ensures that no value is present for FilledTotal, not even an explicit nil
### GetFee

`func (o *CurrentOpenOrdersResponseInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CurrentOpenOrdersResponseInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CurrentOpenOrdersResponseInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CurrentOpenOrdersResponseInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSession

`func (o *CurrentOpenOrdersResponseInner) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CurrentOpenOrdersResponseInner) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CurrentOpenOrdersResponseInner) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *CurrentOpenOrdersResponseInner) HasSession() bool`

HasSession returns a boolean if a field has been set.

### SetSessionNil

`func (o *CurrentOpenOrdersResponseInner) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *CurrentOpenOrdersResponseInner) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetStatus

`func (o *CurrentOpenOrdersResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrentOpenOrdersResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrentOpenOrdersResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CurrentOpenOrdersResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CurrentOpenOrdersResponseInner) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentOpenOrdersResponseInner) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentOpenOrdersResponseInner) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CurrentOpenOrdersResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CurrentOpenOrdersResponseInner) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrentOpenOrdersResponseInner) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrentOpenOrdersResponseInner) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CurrentOpenOrdersResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to README]](../README.md)


