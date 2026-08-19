/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the EquityTradeHistoryResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EquityTradeHistoryResponseRowsInner{}

// EquityTradeHistoryResponseRowsInner struct for EquityTradeHistoryResponseRowsInner
type EquityTradeHistoryResponseRowsInner struct {
	// Execution (per-fill) id.
	ExecutionId *string `json:"executionId,omitempty"`
	// The owning order's id.
	OrderId *string `json:"orderId,omitempty"`
	// US-equity ticker.
	Symbol *string `json:"symbol,omitempty"`
	// Quote asset.
	Quote *string `json:"quote,omitempty"`
	// `BUY` / `SELL`.
	Side *string `json:"side,omitempty"`
	// `MARKET` / `LIMIT`.
	OrderType *string `json:"orderType,omitempty"`
	// Execution price (USD).
	Price *string `json:"price,omitempty"`
	// Executed quantity.
	Qty *string `json:"qty,omitempty"`
	// Notional of this execution (`qty × price`).
	Total *string `json:"total,omitempty"`
	// Execution time (ms epoch).
	ExecutionAt *int64 `json:"executionAt,omitempty"`
	// Last update time (ms epoch).
	UpdatedAt            *int64 `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquityTradeHistoryResponseRowsInner EquityTradeHistoryResponseRowsInner

// NewEquityTradeHistoryResponseRowsInner instantiates a new EquityTradeHistoryResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityTradeHistoryResponseRowsInner() *EquityTradeHistoryResponseRowsInner {
	this := EquityTradeHistoryResponseRowsInner{}
	return &this
}

// NewEquityTradeHistoryResponseRowsInnerWithDefaults instantiates a new EquityTradeHistoryResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityTradeHistoryResponseRowsInnerWithDefaults() *EquityTradeHistoryResponseRowsInner {
	this := EquityTradeHistoryResponseRowsInner{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetExecutionId() string {
	if o == nil || common.IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetExecutionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasExecutionId() bool {
	if o != nil && !common.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *EquityTradeHistoryResponseRowsInner) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetOrderId() string {
	if o == nil || common.IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *EquityTradeHistoryResponseRowsInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EquityTradeHistoryResponseRowsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetQuote() string {
	if o == nil || common.IsNil(o.Quote) {
		var ret string
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetQuoteOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasQuote() bool {
	if o != nil && !common.IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given string and assigns it to the Quote field.
func (o *EquityTradeHistoryResponseRowsInner) SetQuote(v string) {
	o.Quote = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *EquityTradeHistoryResponseRowsInner) SetSide(v string) {
	o.Side = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *EquityTradeHistoryResponseRowsInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *EquityTradeHistoryResponseRowsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *EquityTradeHistoryResponseRowsInner) SetQty(v string) {
	o.Qty = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetTotal() string {
	if o == nil || common.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetTotalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *EquityTradeHistoryResponseRowsInner) SetTotal(v string) {
	o.Total = &v
}

// GetExecutionAt returns the ExecutionAt field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetExecutionAt() int64 {
	if o == nil || common.IsNil(o.ExecutionAt) {
		var ret int64
		return ret
	}
	return *o.ExecutionAt
}

// GetExecutionAtOk returns a tuple with the ExecutionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetExecutionAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExecutionAt) {
		return nil, false
	}
	return o.ExecutionAt, true
}

// HasExecutionAt returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasExecutionAt() bool {
	if o != nil && !common.IsNil(o.ExecutionAt) {
		return true
	}

	return false
}

// SetExecutionAt gets a reference to the given int64 and assigns it to the ExecutionAt field.
func (o *EquityTradeHistoryResponseRowsInner) SetExecutionAt(v int64) {
	o.ExecutionAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EquityTradeHistoryResponseRowsInner) GetUpdatedAt() int64 {
	if o == nil || common.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityTradeHistoryResponseRowsInner) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EquityTradeHistoryResponseRowsInner) HasUpdatedAt() bool {
	if o != nil && !common.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *EquityTradeHistoryResponseRowsInner) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o EquityTradeHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityTradeHistoryResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
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
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.ExecutionAt) {
		toSerialize["executionAt"] = o.ExecutionAt
	}
	if !common.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquityTradeHistoryResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varEquityTradeHistoryResponseRowsInner := _EquityTradeHistoryResponseRowsInner{}

	err = json.Unmarshal(data, &varEquityTradeHistoryResponseRowsInner)

	if err != nil {
		return err
	}

	*o = EquityTradeHistoryResponseRowsInner(varEquityTradeHistoryResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "quote")
		delete(additionalProperties, "side")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "total")
		delete(additionalProperties, "executionAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquityTradeHistoryResponseRowsInner struct {
	value *EquityTradeHistoryResponseRowsInner
	isSet bool
}

func (v NullableEquityTradeHistoryResponseRowsInner) Get() *EquityTradeHistoryResponseRowsInner {
	return v.value
}

func (v *NullableEquityTradeHistoryResponseRowsInner) Set(val *EquityTradeHistoryResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityTradeHistoryResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityTradeHistoryResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityTradeHistoryResponseRowsInner(val *EquityTradeHistoryResponseRowsInner) *NullableEquityTradeHistoryResponseRowsInner {
	return &NullableEquityTradeHistoryResponseRowsInner{value: val, isSet: true}
}

func (v NullableEquityTradeHistoryResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityTradeHistoryResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
