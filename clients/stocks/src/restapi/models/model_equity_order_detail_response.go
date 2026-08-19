/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityOrderDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityOrderDetailResponse{}

// EquityOrderDetailResponse Top-level fields mirror an order row from `/order/history` (orderId, symbol, quote, side, orderType, limitPrice, avgFilledPrice, qty, notional, filledQty, filledTotal, fee, session, status, createdAt, updatedAt), plus a `trades` array. `clientOrderId` is present only in Order Detail, not in Order History.
type EquityOrderDetailResponse struct {
	// Equity order id.
	OrderId *string `json:"orderId,omitempty"`
	// Client-supplied order id. Present only in Order Detail, not in Order History.
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	// US-equity ticker.
	Symbol *string `json:"symbol,omitempty"`
	// Quote asset.
	Quote *string `json:"quote,omitempty"`
	// `BUY` / `SELL`.
	Side *string `json:"side,omitempty"`
	// `MARKET` / `LIMIT`.
	OrderType *string `json:"orderType,omitempty"`
	// Limit price (USD). Non-null for `LIMIT`, `null` for `MARKET`.
	LimitPrice *string `json:"limitPrice,omitempty"`
	// Average fill price (USD). Only present when the order has at least one fill.
	AvgFilledPrice *string `json:"avgFilledPrice,omitempty"`
	// Requested quantity.
	Qty *string `json:"qty,omitempty"`
	// Requested notional.
	Notional *string `json:"notional,omitempty"`
	// Cumulative filled quantity.
	FilledQty *string `json:"filledQty,omitempty"`
	// Cumulative filled notional.
	FilledTotal *string `json:"filledTotal,omitempty"`
	// Total commission fee (USD).
	Fee *string `json:"fee,omitempty"`
	// Trading session.
	Session *string `json:"session,omitempty"`
	// Order lifecycle status.
	Status *string `json:"status,omitempty"`
	// Order creation time (ms epoch).
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// Last update time (ms epoch).
	UpdatedAt *int64 `json:"updatedAt,omitempty"`
	// Trade executions for this order, most recent first. Empty array when no fills.
	Trades               []EquityOrderDetailResponseTradesInner `json:"trades,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquityOrderDetailResponse EquityOrderDetailResponse

// NewEquityOrderDetailResponse instantiates a new EquityOrderDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityOrderDetailResponse() *EquityOrderDetailResponse {
	this := EquityOrderDetailResponse{}
	return &this
}

// NewEquityOrderDetailResponseWithDefaults instantiates a new EquityOrderDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityOrderDetailResponseWithDefaults() *EquityOrderDetailResponse {
	this := EquityOrderDetailResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EquityOrderDetailResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *EquityOrderDetailResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EquityOrderDetailResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *EquityOrderDetailResponse) SetQuote(v string) {
	o.Quote = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *EquityOrderDetailResponse) SetSide(v string) {
	o.Side = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *EquityOrderDetailResponse) SetOrderType(v string) {
	o.OrderType = &v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetLimitPrice() string {
	if o == nil || common.IsNil(o.LimitPrice) {
		var ret string
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetLimitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.LimitPrice) {
		return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasLimitPrice() bool {
	if o != nil && !common.IsNil(o.LimitPrice) {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given NullableString and assigns it to the LimitPrice field.
func (o *EquityOrderDetailResponse) SetLimitPrice(v string) {
	o.LimitPrice = &v
}

// GetAvgFilledPrice returns the AvgFilledPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetAvgFilledPrice() string {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		var ret string
		return ret
	}
	return *o.AvgFilledPrice
}

// GetAvgFilledPriceOk returns a tuple with the AvgFilledPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetAvgFilledPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgFilledPrice) {
		return nil, false
	}
	return o.AvgFilledPrice, true
}

// HasAvgFilledPrice returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasAvgFilledPrice() bool {
	if o != nil && !common.IsNil(o.AvgFilledPrice) {
		return true
	}

	return false
}

// SetAvgFilledPrice gets a reference to the given NullableString and assigns it to the AvgFilledPrice field.
func (o *EquityOrderDetailResponse) SetAvgFilledPrice(v string) {
	o.AvgFilledPrice = &v
}

// GetQty returns the Qty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given NullableString and assigns it to the Qty field.
func (o *EquityOrderDetailResponse) SetQty(v string) {
	o.Qty = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetNotional() string {
	if o == nil || common.IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetNotionalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasNotional() bool {
	if o != nil && !common.IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given NullableString and assigns it to the Notional field.
func (o *EquityOrderDetailResponse) SetNotional(v string) {
	o.Notional = &v
}

// GetFilledQty returns the FilledQty field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetFilledQty() string {
	if o == nil || common.IsNil(o.FilledQty) {
		var ret string
		return ret
	}
	return *o.FilledQty
}

// GetFilledQtyOk returns a tuple with the FilledQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetFilledQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledQty) {
		return nil, false
	}
	return o.FilledQty, true
}

// HasFilledQty returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasFilledQty() bool {
	if o != nil && !common.IsNil(o.FilledQty) {
		return true
	}

	return false
}

// SetFilledQty gets a reference to the given string and assigns it to the FilledQty field.
func (o *EquityOrderDetailResponse) SetFilledQty(v string) {
	o.FilledQty = &v
}

// GetFilledTotal returns the FilledTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetFilledTotal() string {
	if o == nil || common.IsNil(o.FilledTotal) {
		var ret string
		return ret
	}
	return *o.FilledTotal
}

// GetFilledTotalOk returns a tuple with the FilledTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetFilledTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.FilledTotal) {
		return nil, false
	}
	return o.FilledTotal, true
}

// HasFilledTotal returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasFilledTotal() bool {
	if o != nil && !common.IsNil(o.FilledTotal) {
		return true
	}

	return false
}

// SetFilledTotal gets a reference to the given NullableString and assigns it to the FilledTotal field.
func (o *EquityOrderDetailResponse) SetFilledTotal(v string) {
	o.FilledTotal = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *EquityOrderDetailResponse) SetFee(v string) {
	o.Fee = &v
}

// GetSession returns the Session field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityOrderDetailResponse) GetSession() string {
	if o == nil || common.IsNil(o.Session) {
		var ret string
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityOrderDetailResponse) GetSessionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasSession() bool {
	if o != nil && !common.IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given NullableString and assigns it to the Session field.
func (o *EquityOrderDetailResponse) SetSession(v string) {
	o.Session = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EquityOrderDetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetCreatedAt() int64 {
	if o == nil || common.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasCreatedAt() bool {
	if o != nil && !common.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *EquityOrderDetailResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *EquityOrderDetailResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetTrades returns the Trades field value if set, zero value otherwise.
func (o *EquityOrderDetailResponse) GetTrades() []EquityOrderDetailResponseTradesInner {
	if o == nil || common.IsNil(o.Trades) {
		var ret []EquityOrderDetailResponseTradesInner
		return ret
	}
	return o.Trades
}

// GetTradesOk returns a tuple with the Trades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityOrderDetailResponse) GetTradesOk() ([]EquityOrderDetailResponseTradesInner, bool) {
	if o == nil || common.IsNil(o.Trades) {
		return nil, false
	}
	return o.Trades, true
}

// HasTrades returns a boolean if a field has been set.
func (o *EquityOrderDetailResponse) HasTrades() bool {
	if o != nil && !common.IsNil(o.Trades) {
		return true
	}

	return false
}

// SetTrades gets a reference to the given []EquityOrderDetailResponseTradesInner and assigns it to the Trades field.
func (o *EquityOrderDetailResponse) SetTrades(v []EquityOrderDetailResponseTradesInner) {
	o.Trades = v
}

func (o EquityOrderDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityOrderDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
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
	if !common.IsNil(o.Trades) {
		toSerialize["trades"] = o.Trades
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquityOrderDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varEquityOrderDetailResponse := _EquityOrderDetailResponse{}

	err = json.Unmarshal(data, &varEquityOrderDetailResponse)

	if err != nil {
		return err
	}

	*o = EquityOrderDetailResponse(varEquityOrderDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "clientOrderId")
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
		delete(additionalProperties, "trades")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquityOrderDetailResponse struct {
	value *EquityOrderDetailResponse
	isSet bool
}

func (v NullableEquityOrderDetailResponse) Get() *EquityOrderDetailResponse {
	return v.value
}

func (v *NullableEquityOrderDetailResponse) Set(val *EquityOrderDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityOrderDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityOrderDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityOrderDetailResponse(val *EquityOrderDetailResponse) *NullableEquityOrderDetailResponse {
	return &NullableEquityOrderDetailResponse{value: val, isSet: true}
}

func (v NullableEquityOrderDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityOrderDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
