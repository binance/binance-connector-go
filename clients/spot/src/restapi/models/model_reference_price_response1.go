/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ReferencePriceResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReferencePriceResponse1{}

// ReferencePriceResponse1 If a reference price is set:
type ReferencePriceResponse1 struct {
	Symbol *string `json:"symbol,omitempty"`
	// Reference price. Can be `null` if no reference price is set.
	ReferencePrice *string `json:"referencePrice,omitempty"`
	// Timestamp when reference price was valid.
	Timestamp            *int64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReferencePriceResponse1 ReferencePriceResponse1

// NewReferencePriceResponse1 instantiates a new ReferencePriceResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencePriceResponse1() *ReferencePriceResponse1 {
	this := ReferencePriceResponse1{}
	return &this
}

// NewReferencePriceResponse1WithDefaults instantiates a new ReferencePriceResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencePriceResponse1WithDefaults() *ReferencePriceResponse1 {
	this := ReferencePriceResponse1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ReferencePriceResponse1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ReferencePriceResponse1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ReferencePriceResponse1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetReferencePrice returns the ReferencePrice field value if set, zero value otherwise.
func (o *ReferencePriceResponse1) GetReferencePrice() string {
	if o == nil || common.IsNil(o.ReferencePrice) {
		var ret string
		return ret
	}
	return *o.ReferencePrice
}

// GetReferencePriceOk returns a tuple with the ReferencePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1) GetReferencePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ReferencePrice) {
		return nil, false
	}
	return o.ReferencePrice, true
}

// HasReferencePrice returns a boolean if a field has been set.
func (o *ReferencePriceResponse1) HasReferencePrice() bool {
	if o != nil && !common.IsNil(o.ReferencePrice) {
		return true
	}

	return false
}

// SetReferencePrice gets a reference to the given string and assigns it to the ReferencePrice field.
func (o *ReferencePriceResponse1) SetReferencePrice(v string) {
	o.ReferencePrice = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReferencePriceResponse1) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencePriceResponse1) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReferencePriceResponse1) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ReferencePriceResponse1) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ReferencePriceResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferencePriceResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.ReferencePrice) {
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

func (o *ReferencePriceResponse1) UnmarshalJSON(data []byte) (err error) {
	varReferencePriceResponse1 := _ReferencePriceResponse1{}

	err = json.Unmarshal(data, &varReferencePriceResponse1)

	if err != nil {
		return err
	}

	*o = ReferencePriceResponse1(varReferencePriceResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "referencePrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReferencePriceResponse1 struct {
	value *ReferencePriceResponse1
	isSet bool
}

func (v NullableReferencePriceResponse1) Get() *ReferencePriceResponse1 {
	return v.value
}

func (v *NullableReferencePriceResponse1) Set(val *ReferencePriceResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceResponse1(val *ReferencePriceResponse1) *NullableReferencePriceResponse1 {
	return &NullableReferencePriceResponse1{value: val, isSet: true}
}

func (v NullableReferencePriceResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
