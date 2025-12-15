/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerV2Response1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerV2Response1{}

// SymbolPriceTickerV2Response1 struct for SymbolPriceTickerV2Response1
type SymbolPriceTickerV2Response1 struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerV2Response1 SymbolPriceTickerV2Response1

// NewSymbolPriceTickerV2Response1 instantiates a new SymbolPriceTickerV2Response1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerV2Response1() *SymbolPriceTickerV2Response1 {
	this := SymbolPriceTickerV2Response1{}
	return &this
}

// NewSymbolPriceTickerV2Response1WithDefaults instantiates a new SymbolPriceTickerV2Response1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerV2Response1WithDefaults() *SymbolPriceTickerV2Response1 {
	this := SymbolPriceTickerV2Response1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerV2Response1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response1) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response1) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response1) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerV2Response1) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response1) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response1) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response1) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerV2Response1) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerV2Response1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerV2Response1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SymbolPriceTickerV2Response1) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerV2Response1 := _SymbolPriceTickerV2Response1{}

	err = json.Unmarshal(data, &varSymbolPriceTickerV2Response1)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerV2Response1(varSymbolPriceTickerV2Response1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerV2Response1 struct {
	value *SymbolPriceTickerV2Response1
	isSet bool
}

func (v NullableSymbolPriceTickerV2Response1) Get() *SymbolPriceTickerV2Response1 {
	return v.value
}

func (v *NullableSymbolPriceTickerV2Response1) Set(val *SymbolPriceTickerV2Response1) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerV2Response1) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerV2Response1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerV2Response1(val *SymbolPriceTickerV2Response1) *NullableSymbolPriceTickerV2Response1 {
	return &NullableSymbolPriceTickerV2Response1{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerV2Response1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerV2Response1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
