/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RecentBlockTradesListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecentBlockTradesListResponseInner{}

// RecentBlockTradesListResponseInner struct for RecentBlockTradesListResponseInner
type RecentBlockTradesListResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	TradeId              *int64  `json:"tradeId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Side                 *int64  `json:"side,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentBlockTradesListResponseInner RecentBlockTradesListResponseInner

// NewRecentBlockTradesListResponseInner instantiates a new RecentBlockTradesListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentBlockTradesListResponseInner() *RecentBlockTradesListResponseInner {
	this := RecentBlockTradesListResponseInner{}
	return &this
}

// NewRecentBlockTradesListResponseInnerWithDefaults instantiates a new RecentBlockTradesListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentBlockTradesListResponseInnerWithDefaults() *RecentBlockTradesListResponseInner {
	this := RecentBlockTradesListResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RecentBlockTradesListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *RecentBlockTradesListResponseInner) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *RecentBlockTradesListResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *RecentBlockTradesListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *RecentBlockTradesListResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *RecentBlockTradesListResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetSide() int64 {
	if o == nil || common.IsNil(o.Side) {
		var ret int64
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetSideOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given int64 and assigns it to the Side field.
func (o *RecentBlockTradesListResponseInner) SetSide(v int64) {
	o.Side = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RecentBlockTradesListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecentBlockTradesListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RecentBlockTradesListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *RecentBlockTradesListResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o RecentBlockTradesListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentBlockTradesListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentBlockTradesListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varRecentBlockTradesListResponseInner := _RecentBlockTradesListResponseInner{}

	err = json.Unmarshal(data, &varRecentBlockTradesListResponseInner)

	if err != nil {
		return err
	}

	*o = RecentBlockTradesListResponseInner(varRecentBlockTradesListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tradeId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "side")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentBlockTradesListResponseInner struct {
	value *RecentBlockTradesListResponseInner
	isSet bool
}

func (v NullableRecentBlockTradesListResponseInner) Get() *RecentBlockTradesListResponseInner {
	return v.value
}

func (v *NullableRecentBlockTradesListResponseInner) Set(val *RecentBlockTradesListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentBlockTradesListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentBlockTradesListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentBlockTradesListResponseInner(val *RecentBlockTradesListResponseInner) *NullableRecentBlockTradesListResponseInner {
	return &NullableRecentBlockTradesListResponseInner{value: val, isSet: true}
}

func (v NullableRecentBlockTradesListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentBlockTradesListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
