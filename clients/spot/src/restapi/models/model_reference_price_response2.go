/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse2{}

// ReferencePriceResponse2 If no reference price is set:
type ReferencePriceResponse2 struct {
	Symbol *string `json:"symbol,omitempty"`
	// Reference price. Can be `null` if no reference price is set.
	ReferencePrice interface{} `json:"referencePrice,omitempty"`
	// Timestamp when reference price was valid
	Timestamp            *int64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse2 ReferencePriceResponse2

// NewReferencePriceResponse2 instantiates a new ReferencePriceResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse2() *ReferencePriceResponse2 {
	this := ReferencePriceResponse2{}
	return &this
}

// NewReferencePriceResponse2WithDefaults instantiates a new ReferencePriceResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponse2WithDefaults() *ReferencePriceResponse2 {
	this := ReferencePriceResponse2{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceResponse2) SetSymbol(v string) {
	o.Symbol = &v
}

// GetReferencePrice returns the ReferencePrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReferencePriceResponse2) GetReferencePrice() interface{} {
	if o == nil || common.IsNil(o.ReferencePrice) {
		var ret interface{}
		return ret
	}
	return o.ReferencePrice
}

// GetReferencePriceOk returns a tuple with the ReferencePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReferencePriceResponse2) GetReferencePriceOk() (*interface{}, bool) {
	if o == nil || common.IsNil(o.ReferencePrice) {
		return nil, false
	}
	return &o.ReferencePrice, true
}

// HasReferencePrice returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasReferencePrice() bool {
	if o != nil && !common.IsNil(o.ReferencePrice) {
		return true
	}

	return false
}

// SetReferencePrice gets a reference to the given interface{} and assigns it to the ReferencePrice field.
func (o *ReferencePriceResponse2) SetReferencePrice(v interface{}) {
	o.ReferencePrice = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReferencePriceResponse2) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse2) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReferencePriceResponse2) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ReferencePriceResponse2) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ReferencePriceResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if o.ReferencePrice != nil {
		toSerialize["referencePrice"] = o.ReferencePrice
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReferencePriceResponse2) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse2 := _ReferencePriceResponse2{}

	err = json.Unmarshal(data, &varReferencePriceResponse2)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse2(varReferencePriceResponse2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "referencePrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponse2 struct {
	value *ReferencePriceResponse2
	isSet bool
}

func (v NullableReferencePriceResponse2) Get() *ReferencePriceResponse2 {
	return v.value
}

func (v *NullableReferencePriceResponse2) Set(val *ReferencePriceResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse2(val *ReferencePriceResponse2) *NullableReferencePriceResponse2 {
	return &NullableReferencePriceResponse2{value: val, isSet: true}
}

func (v NullableReferencePriceResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
