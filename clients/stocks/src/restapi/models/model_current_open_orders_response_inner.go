/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CurrentOpenOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentOpenOrdersResponseInner{}

// CurrentOpenOrdersResponseInner struct for CurrentOpenOrdersResponseInner
type CurrentOpenOrdersResponseInner struct {
	// Equity order id.
	OrderId *string `json:"orderId,omitempty"`
	// US-equity ticker.
	Symbol *string `json:"symbol,omitempty"`
	// Quote asset (e.g. `USDC`).
	Quote *string `json:"quote,omitempty"`
	// `BUY` / `SELL`.
	Side *string `json:"side,omitempty"`
	// `MARKET` / `LIMIT`.
	OrderType *string `json:"orderType,omitempty"`
	// Limit price (USD). Non-null for `LIMIT` orders, `null` for `MARKET`.
	LimitPrice *string `json:"limitPrice,omitempty"`
	// Average fill price (USD). `null` until the first fill.
	AvgFilledPrice *string `json:"avgFilledPrice,omitempty"`
	// Requested quantity. `null` for `BUY MARKET` (use `notional` instead).
	Qty *string `json:"qty,omitempty"`
	// Requested notional. Non-null for `BUY MARKET`; `null` otherwise.
	Notional *string `json:"notional,omitempty"`
	// Cumulative filled quantity.
	FilledQty *string `json:"filledQty,omitempty"`
	// Cumulative filled notional. Populated only for `BUY MARKET`.
	FilledTotal *string `json:"filledTotal,omitempty"`
	// Total commission fee (USD).
	Fee *string `json:"fee,omitempty"`
	// Trading session the order was placed under: `RTH` / `EXTENDED` / `24H`. `null` for `MARKET` orders.
	Session *string `json:"session,omitempty"`
	// Order lifecycle status — one of `NEW` / `ACCEPTED` / `PARTIALLY_FILLED` / `FILLED` / `CANCELED` / `EXPIRED` / `REJECTED`.
	Status *string `json:"status,omitempty"`
	// Order creation time (ms epoch).
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Last update time (ms epoch).
	UpdatedAt            *int64 `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CurrentOpenOrdersResponseInner CurrentOpenOrdersResponseInner

// NewCurrentOpenOrdersResponseInner instantiates a new CurrentOpenOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentOpenOrdersResponseInner() *CurrentOpenOrdersResponseInner {
	this := CurrentOpenOrdersResponseInner{}
	return &this
}

// NewCurrentOpenOrdersResponseInnerWithDefaults instantiates a new CurrentOpenOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentOpenOrdersResponseInnerWithDefaults() *CurrentOpenOrdersResponseInner {
	this := CurrentOpenOrdersResponseInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CurrentOpenOrdersResponseInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CurrentOpenOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *CurrentOpenOrdersResponseInner) SetQuote(v string) {
	o.Quote = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CurrentOpenOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *CurrentOpenOrdersResponseInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetLimitPrice() string {
	if o == nil || common.IsNil(o.LimitPrice) {
		var ret string
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetLimitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LimitPrice) {
		return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasLimitPrice() bool {
	if o != nil && !common.IsNil(o.LimitPrice) {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given NullableString and assigns it to the LimitPrice field.
func (o *CurrentOpenOrdersResponseInner) SetLimitPrice(v string) {
	o.LimitPrice = &v
}

// GetAvgFilledPrice returns the AvgFilledPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetAvgFilledPrice() string {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		var ret string
		return ret
	}
	return *o.AvgFilledPrice
}

// GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetAvgFilledPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		return nil, false
	}
	return o.AvgFilledPrice, true
}

// HasAvgFilledPrice returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasAvgFilledPrice() bool {
	if o != nil && !common.IsNil(o.AvgFilledPrice) {
		return true
	}

	return false
}

// SetAvgFilledPrice gets a reference to the given NullableString and assigns it to the AvgFilledPrice field.
func (o *CurrentOpenOrdersResponseInner) SetAvgFilledPrice(v string) {
	o.AvgFilledPrice = &v
}

// GetQty returns the Qty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given NullableString and assigns it to the Qty field.
func (o *CurrentOpenOrdersResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given NullableString and assigns it to the Notional field.
func (o *CurrentOpenOrdersResponseInner) SetNotional(v string) {
	o.Notional = &v
}

// GetFilledQty returns the FilledQty field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetFilledQty() string {
	if o == nil || common.IsNil(o.FilledQty) {
		var ret string
		return ret
	}
	return *o.FilledQty
}

// GetFilledQtyOk returns a tuple with the FilledQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetFilledQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledQty) {
		return nil, false
	}
	return o.FilledQty, true
}

// HasFilledQty returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasFilledQty() bool {
	if o != nil && !common.IsNil(o.FilledQty) {
		return true
	}

	return false
}

// SetFilledQty gets a reference to the given string and assigns it to the FilledQty field.
func (o *CurrentOpenOrdersResponseInner) SetFilledQty(v string) {
	o.FilledQty = &v
}

// GetFilledTotal returns the FilledTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetFilledTotal() string {
	if o == nil || common.IsNil(o.FilledTotal) {
		var ret string
		return ret
	}
	return *o.FilledTotal
}

// GetFilledTotalOk returns a tuple with the FilledTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetFilledTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledTotal) {
		return nil, false
	}
	return o.FilledTotal, true
}

// HasFilledTotal returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasFilledTotal() bool {
	if o != nil && !common.IsNil(o.FilledTotal) {
		return true
	}

	return false
}

// SetFilledTotal gets a reference to the given NullableString and assigns it to the FilledTotal field.
func (o *CurrentOpenOrdersResponseInner) SetFilledTotal(v string) {
	o.FilledTotal = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *CurrentOpenOrdersResponseInner) SetFee(v string) {
	o.Fee = &v
}

// GetSession returns the Session field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentOpenOrdersResponseInner) GetSession() string {
	if o == nil || common.IsNil(o.Session) {
		var ret string
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentOpenOrdersResponseInner) GetSessionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasSession() bool {
	if o != nil && !common.IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given NullableString and assigns it to the Session field.
func (o *CurrentOpenOrdersResponseInner) SetSession(v string) {
	o.Session = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CurrentOpenOrdersResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetCreatedAt() int64 {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetCreatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CurrentOpenOrdersResponseInner) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CurrentOpenOrdersResponseInner) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentOpenOrdersResponseInner) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CurrentOpenOrdersResponseInner) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CurrentOpenOrdersResponseInner) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o CurrentOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentOpenOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.LimitPrice) {
		toSerialize["limitPrice"] = o.LimitPrice
	}
	if !common.IsNil(o.AvgFilledPrice) {
		toSerialize["avgFilledPrice"] = o.AvgFilledPrice
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !common.IsNil(o.FilledQty) {
		toSerialize["filledQty"] = o.FilledQty
	}
	if !common.IsNil(o.FilledTotal) {
		toSerialize["filledTotal"] = o.FilledTotal
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !common.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CurrentOpenOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCurrentOpenOrdersResponseInner := _CurrentOpenOrdersResponseInner{}

	err = json.Unmarshal(data, &varCurrentOpenOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = CurrentOpenOrdersResponseInner(varCurrentOpenOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "quote")
		delete(additionalProperties, "side")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "limitPrice")
		delete(additionalProperties, "avgFilledPrice")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "notional")
		delete(additionalProperties, "filledQty")
		delete(additionalProperties, "filledTotal")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "session")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCurrentOpenOrdersResponseInner struct {
	value *CurrentOpenOrdersResponseInner
	isSet bool
}

func (v NullableCurrentOpenOrdersResponseInner) Get() *CurrentOpenOrdersResponseInner {
	return v.value
}

func (v *NullableCurrentOpenOrdersResponseInner) Set(val *CurrentOpenOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentOpenOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentOpenOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentOpenOrdersResponseInner(val *CurrentOpenOrdersResponseInner) *NullableCurrentOpenOrdersResponseInner {
	return &NullableCurrentOpenOrdersResponseInner{value: val, isSet: true}
}

func (v NullableCurrentOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentOpenOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
