/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OldTradesLookupResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OldTradesLookupResponseInner{}

// OldTradesLookupResponseInner struct for OldTradesLookupResponseInner
type OldTradesLookupResponseInner struct {
	Id                   *string `json:"id,omitempty"`
	TradeId              *string `json:"tradeId,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Side                 *int64  `json:"side,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OldTradesLookupResponseInner OldTradesLookupResponseInner

// NewOldTradesLookupResponseInner instantiates a new OldTradesLookupResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOldTradesLookupResponseInner() *OldTradesLookupResponseInner {
	this := OldTradesLookupResponseInner{}
	return &this
}

// NewOldTradesLookupResponseInnerWithDefaults instantiates a new OldTradesLookupResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOldTradesLookupResponseInnerWithDefaults() *OldTradesLookupResponseInner {
	this := OldTradesLookupResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OldTradesLookupResponseInner) SetId(v string) {
	o.Id = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetTradeId() string {
	if o == nil || common.IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetTradeIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *OldTradesLookupResponseInner) SetTradeId(v string) {
	o.TradeId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OldTradesLookupResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OldTradesLookupResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *OldTradesLookupResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetSide() int64 {
	if o == nil || common.IsNil(o.Side) {
		var ret int64
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetSideOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given int64 and assigns it to the Side field.
func (o *OldTradesLookupResponseInner) SetSide(v int64) {
	o.Side = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OldTradesLookupResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OldTradesLookupResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OldTradesLookupResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OldTradesLookupResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o OldTradesLookupResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OldTradesLookupResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
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

func (o *OldTradesLookupResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOldTradesLookupResponseInner := _OldTradesLookupResponseInner{}

	err = json.Unmarshal(data, &varOldTradesLookupResponseInner)

	if err != nil {
		return err
	}

	*o = OldTradesLookupResponseInner(varOldTradesLookupResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tradeId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "side")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOldTradesLookupResponseInner struct {
	value *OldTradesLookupResponseInner
	isSet bool
}

func (v NullableOldTradesLookupResponseInner) Get() *OldTradesLookupResponseInner {
	return v.value
}

func (v *NullableOldTradesLookupResponseInner) Set(val *OldTradesLookupResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOldTradesLookupResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOldTradesLookupResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOldTradesLookupResponseInner(val *OldTradesLookupResponseInner) *NullableOldTradesLookupResponseInner {
	return &NullableOldTradesLookupResponseInner{value: val, isSet: true}
}

func (v NullableOldTradesLookupResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOldTradesLookupResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
