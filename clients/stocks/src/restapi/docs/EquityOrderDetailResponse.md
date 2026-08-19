# EquityOrderDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Equity order id. | [optional] 
**ClientOrderId** | Pointer to **string** | Client-supplied order id. Present only in Order Detail, not in Order History. | [optional] 
**Symbol** | Pointer to **string** | US-equity ticker. | [optional] 
**Quote** | Pointer to **string** | Quote asset. | [optional] 
**Side** | Pointer to **string** | &#x60;BUY&#x60; / &#x60;SELL&#x60;. | [optional] 
**OrderType** | Pointer to **string** | &#x60;MARKET&#x60; / &#x60;LIMIT&#x60;. | [optional] 
**LimitPrice** | Pointer to **NullableString** | Limit price (USD). Non-null for &#x60;LIMIT&#x60;, &#x60;null&#x60; for &#x60;MARKET&#x60;. | [optional] 
**AvgFilledPrice** | Pointer to **NullableString** | Average fill price (USD). Only present when the order has at least one fill. | [optional] 
**Qty** | Pointer to **NullableString** | Requested quantity. | [optional] 
**Notional** | Pointer to **NullableString** | Requested notional. | [optional] 
**FilledQty** | Pointer to **string** | Cumulative filled quantity. | [optional] 
**FilledTotal** | Pointer to **NullableString** | Cumulative filled notional. | [optional] 
**Fee** | Pointer to **string** | Total commission fee (USD). | [optional] 
**Session** | Pointer to **NullableString** | Trading session. | [optional] 
**Status** | Pointer to **string** | Order lifecycle status. | [optional] 
**CreatedAt** | Pointer to **int64** | Order creation time (ms epoch). | [optional] 
**UpdatedAt** | Pointer to **int64** | Last update time (ms epoch). | [optional] 
**Trades** | Pointer to [**[]EquityOrderDetailResponseTradesInner**](EquityOrderDetailResponseTradesInner.md) | Trade executions for this order, most recent first. Empty array when no fills. | [optional] 

## Methods

### NewEquityOrderDetailResponse

`func NewEquityOrderDetailResponse() *EquityOrderDetailResponse`

NewEquityOrderDetailResponse instantiates a new EquityOrderDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityOrderDetailResponseWithDefaults

`func NewEquityOrderDetailResponseWithDefaults() *EquityOrderDetailResponse`

NewEquityOrderDetailResponseWithDefaults instantiates a new EquityOrderDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *EquityOrderDetailResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *EquityOrderDetailResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *EquityOrderDetailResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *EquityOrderDetailResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *EquityOrderDetailResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *EquityOrderDetailResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *EquityOrderDetailResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *EquityOrderDetailResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *EquityOrderDetailResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EquityOrderDetailResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EquityOrderDetailResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EquityOrderDetailResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetQuote

`func (o *EquityOrderDetailResponse) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *EquityOrderDetailResponse) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *EquityOrderDetailResponse) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *EquityOrderDetailResponse) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetSide

`func (o *EquityOrderDetailResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *EquityOrderDetailResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *EquityOrderDetailResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *EquityOrderDetailResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderType

`func (o *EquityOrderDetailResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *EquityOrderDetailResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *EquityOrderDetailResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *EquityOrderDetailResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetLimitPrice

`func (o *EquityOrderDetailResponse) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *EquityOrderDetailResponse) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *EquityOrderDetailResponse) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *EquityOrderDetailResponse) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### SetLimitPriceNil

`func (o *EquityOrderDetailResponse) SetLimitPriceNil(b bool)`

 SetLimitPriceNil sets the value for LimitPrice to be an explicit nil

### UnsetLimitPrice
`func (o *EquityOrderDetailResponse) UnsetLimitPrice()`

UnsetLimitPrice ensures that no value is present for LimitPrice, not even an explicit nil
### GetAvgFilledPrice

`func (o *EquityOrderDetailResponse) GetAvgFilledPrice() string`

GetAvgFilledPrice returns the AvgFilledPrice field if non-nil, zero value otherwise.

### GetAvgFilledPriceOk

`func (o *EquityOrderDetailResponse) GetAvgFilledPriceOk() (*string, bool)`

GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgFilledPrice

`func (o *EquityOrderDetailResponse) SetAvgFilledPrice(v string)`

SetAvgFilledPrice sets AvgFilledPrice field to given value.

### HasAvgFilledPrice

`func (o *EquityOrderDetailResponse) HasAvgFilledPrice() bool`

HasAvgFilledPrice returns a boolean if a field has been set.

### SetAvgFilledPriceNil

`func (o *EquityOrderDetailResponse) SetAvgFilledPriceNil(b bool)`

 SetAvgFilledPriceNil sets the value for AvgFilledPrice to be an explicit nil

### UnsetAvgFilledPrice
`func (o *EquityOrderDetailResponse) UnsetAvgFilledPrice()`

UnsetAvgFilledPrice ensures that no value is present for AvgFilledPrice, not even an explicit nil
### GetQty

`func (o *EquityOrderDetailResponse) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *EquityOrderDetailResponse) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *EquityOrderDetailResponse) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *EquityOrderDetailResponse) HasQty() bool`

HasQty returns a boolean if a field has been set.

### SetQtyNil

`func (o *EquityOrderDetailResponse) SetQtyNil(b bool)`

 SetQtyNil sets the value for Qty to be an explicit nil

### UnsetQty
`func (o *EquityOrderDetailResponse) UnsetQty()`

UnsetQty ensures that no value is present for Qty, not even an explicit nil
### GetNotional

`func (o *EquityOrderDetailResponse) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *EquityOrderDetailResponse) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *EquityOrderDetailResponse) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *EquityOrderDetailResponse) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### SetNotionalNil

`func (o *EquityOrderDetailResponse) SetNotionalNil(b bool)`

 SetNotionalNil sets the value for Notional to be an explicit nil

### UnsetNotional
`func (o *EquityOrderDetailResponse) UnsetNotional()`

UnsetNotional ensures that no value is present for Notional, not even an explicit nil
### GetFilledQty

`func (o *EquityOrderDetailResponse) GetFilledQty() string`

GetFilledQty returns the FilledQty field if non-nil, zero value otherwise.

### GetFilledQtyOk

`func (o *EquityOrderDetailResponse) GetFilledQtyOk() (*string, bool)`

GetFilledQtyOk returns a tuple with the FilledQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQty

`func (o *EquityOrderDetailResponse) SetFilledQty(v string)`

SetFilledQty sets FilledQty field to given value.

### HasFilledQty

`func (o *EquityOrderDetailResponse) HasFilledQty() bool`

HasFilledQty returns a boolean if a field has been set.

### GetFilledTotal

`func (o *EquityOrderDetailResponse) GetFilledTotal() string`

GetFilledTotal returns the FilledTotal field if non-nil, zero value otherwise.

### GetFilledTotalOk

`func (o *EquityOrderDetailResponse) GetFilledTotalOk() (*string, bool)`

GetFilledTotalOk returns a tuple with the FilledTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTotal

`func (o *EquityOrderDetailResponse) SetFilledTotal(v string)`

SetFilledTotal sets FilledTotal field to given value.

### HasFilledTotal

`func (o *EquityOrderDetailResponse) HasFilledTotal() bool`

HasFilledTotal returns a boolean if a field has been set.

### SetFilledTotalNil

`func (o *EquityOrderDetailResponse) SetFilledTotalNil(b bool)`

 SetFilledTotalNil sets the value for FilledTotal to be an explicit nil

### UnsetFilledTotal
`func (o *EquityOrderDetailResponse) UnsetFilledTotal()`

UnsetFilledTotal ensures that no value is present for FilledTotal, not even an explicit nil
### GetFee

`func (o *EquityOrderDetailResponse) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *EquityOrderDetailResponse) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *EquityOrderDetailResponse) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *EquityOrderDetailResponse) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetSession

`func (o *EquityOrderDetailResponse) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *EquityOrderDetailResponse) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *EquityOrderDetailResponse) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *EquityOrderDetailResponse) HasSession() bool`

HasSession returns a boolean if a field has been set.

### SetSessionNil

`func (o *EquityOrderDetailResponse) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *EquityOrderDetailResponse) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetStatus

`func (o *EquityOrderDetailResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EquityOrderDetailResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EquityOrderDetailResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EquityOrderDetailResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EquityOrderDetailResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EquityOrderDetailResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EquityOrderDetailResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EquityOrderDetailResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EquityOrderDetailResponse) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EquityOrderDetailResponse) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EquityOrderDetailResponse) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EquityOrderDetailResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTrades

`func (o *EquityOrderDetailResponse) GetTrades() []EquityOrderDetailResponseTradesInner`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *EquityOrderDetailResponse) GetTradesOk() (*[]EquityOrderDetailResponseTradesInner, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *EquityOrderDetailResponse) SetTrades(v []EquityOrderDetailResponseTradesInner)`

SetTrades sets Trades field to given value.

### HasTrades

`func (o *EquityOrderDetailResponse) HasTrades() bool`

HasTrades returns a boolean if a field has been set.


[[Back to README]](../README.md)


