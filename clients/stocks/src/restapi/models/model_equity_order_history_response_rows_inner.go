/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityOrderHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityOrderHistoryResponseRowsInner{}

// EquityOrderHistoryResponseRowsInner struct for EquityOrderHistoryResponseRowsInner
type EquityOrderHistoryResponseRowsInner struct {
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
	// Average fill price (USD). `null` until the first fill. For `MARKET` orders this is the only meaningful price field.
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

type _EquityOrderHistoryResponseRowsInner EquityOrderHistoryResponseRowsInner

// NewEquityOrderHistoryResponseRowsInner instantiates a new EquityOrderHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityOrderHistoryResponseRowsInner() *EquityOrderHistoryResponseRowsInner {
	this := EquityOrderHistoryResponseRowsInner{}
	return &this
}

// NewEquityOrderHistoryResponseRowsInnerWithDefaults instantiates a new EquityOrderHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityOrderHistoryResponseRowsInnerWithDefaults() *EquityOrderHistoryResponseRowsInner {
	this := EquityOrderHistoryResponseRowsInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EquityOrderHistoryResponseRowsInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EquityOrderHistoryResponseRowsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *EquityOrderHistoryResponseRowsInner) SetQuote(v string) {
	o.Quote = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *EquityOrderHistoryResponseRowsInner) SetSide(v string) {
	o.Side = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *EquityOrderHistoryResponseRowsInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetLimitPrice() string {
	if o == nil || common.IsNil(o.LimitPrice) {
		var ret string
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetLimitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LimitPrice) {
		return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasLimitPrice() bool {
	if o != nil && !common.IsNil(o.LimitPrice) {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given NullableString and assigns it to the LimitPrice field.
func (o *EquityOrderHistoryResponseRowsInner) SetLimitPrice(v string) {
	o.LimitPrice = &v
}

// GetAvgFilledPrice returns the AvgFilledPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetAvgFilledPrice() string {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		var ret string
		return ret
	}
	return *o.AvgFilledPrice
}

// GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetAvgFilledPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		return nil, false
	}
	return o.AvgFilledPrice, true
}

// HasAvgFilledPrice returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasAvgFilledPrice() bool {
	if o != nil && !common.IsNil(o.AvgFilledPrice) {
		return true
	}

	return false
}

// SetAvgFilledPrice gets a reference to the given NullableString and assigns it to the AvgFilledPrice field.
func (o *EquityOrderHistoryResponseRowsInner) SetAvgFilledPrice(v string) {
	o.AvgFilledPrice = &v
}

// GetQty returns the Qty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given NullableString and assigns it to the Qty field.
func (o *EquityOrderHistoryResponseRowsInner) SetQty(v string) {
	o.Qty = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given NullableString and assigns it to the Notional field.
func (o *EquityOrderHistoryResponseRowsInner) SetNotional(v string) {
	o.Notional = &v
}

// GetFilledQty returns the FilledQty field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetFilledQty() string {
	if o == nil || common.IsNil(o.FilledQty) {
		var ret string
		return ret
	}
	return *o.FilledQty
}

// GetFilledQtyOk returns a tuple with the FilledQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetFilledQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledQty) {
		return nil, false
	}
	return o.FilledQty, true
}

// HasFilledQty returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasFilledQty() bool {
	if o != nil && !common.IsNil(o.FilledQty) {
		return true
	}

	return false
}

// SetFilledQty gets a reference to the given string and assigns it to the FilledQty field.
func (o *EquityOrderHistoryResponseRowsInner) SetFilledQty(v string) {
	o.FilledQty = &v
}

// GetFilledTotal returns the FilledTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetFilledTotal() string {
	if o == nil || common.IsNil(o.FilledTotal) {
		var ret string
		return ret
	}
	return *o.FilledTotal
}

// GetFilledTotalOk returns a tuple with the FilledTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetFilledTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledTotal) {
		return nil, false
	}
	return o.FilledTotal, true
}

// HasFilledTotal returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasFilledTotal() bool {
	if o != nil && !common.IsNil(o.FilledTotal) {
		return true
	}

	return false
}

// SetFilledTotal gets a reference to the given NullableString and assigns it to the FilledTotal field.
func (o *EquityOrderHistoryResponseRowsInner) SetFilledTotal(v string) {
	o.FilledTotal = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *EquityOrderHistoryResponseRowsInner) SetFee(v string) {
	o.Fee = &v
}

// GetSession returns the Session field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderHistoryResponseRowsInner) GetSession() string {
	if o == nil || common.IsNil(o.Session) {
		var ret string
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderHistoryResponseRowsInner) GetSessionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasSession() bool {
	if o != nil && !common.IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given NullableString and assigns it to the Session field.
func (o *EquityOrderHistoryResponseRowsInner) SetSession(v string) {
	o.Session = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EquityOrderHistoryResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetCreatedAt() int64 {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetCreatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *EquityOrderHistoryResponseRowsInner) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EquityOrderHistoryResponseRowsInner) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EquityOrderHistoryResponseRowsInner) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *EquityOrderHistoryResponseRowsInner) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o EquityOrderHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityOrderHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
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

func (o *EquityOrderHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varEquityOrderHistoryResponseRowsInner := _EquityOrderHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varEquityOrderHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = EquityOrderHistoryResponseRowsInner(varEquityOrderHistoryResponseRowsInner)

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

type NullableEquityOrderHistoryResponseRowsInner struct {
	value *EquityOrderHistoryResponseRowsInner
	isSet bool
}

func (v NullableEquityOrderHistoryResponseRowsInner) Get() *EquityOrderHistoryResponseRowsInner {
	return v.value
}

func (v *NullableEquityOrderHistoryResponseRowsInner) Set(val *EquityOrderHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityOrderHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityOrderHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityOrderHistoryResponseRowsInner(val *EquityOrderHistoryResponseRowsInner) *NullableEquityOrderHistoryResponseRowsInner {
	return &NullableEquityOrderHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableEquityOrderHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityOrderHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
