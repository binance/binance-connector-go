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

// checks if the SymbolPriceTickerResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse1{}

// SymbolPriceTickerResponse1 struct for SymbolPriceTickerResponse1
type SymbolPriceTickerResponse1 struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerResponse1 SymbolPriceTickerResponse1

// NewSymbolPriceTickerResponse1 instantiates a new SymbolPriceTickerResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse1() *SymbolPriceTickerResponse1 {
	this := SymbolPriceTickerResponse1{}
	return &this
}

// NewSymbolPriceTickerResponse1WithDefaults instantiates a new SymbolPriceTickerResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponse1WithDefaults() *SymbolPriceTickerResponse1 {
	this := SymbolPriceTickerResponse1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerResponse1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerResponse1) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerResponse1) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse1) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolPriceTickerResponse1) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerResponse1 := _SymbolPriceTickerResponse1{}

	err = json.Unmarshal(data, &varSymbolPriceTickerResponse1)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerResponse1(varSymbolPriceTickerResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerResponse1 struct {
	value *SymbolPriceTickerResponse1
	isSet bool
}

func (v NullableSymbolPriceTickerResponse1) Get() *SymbolPriceTickerResponse1 {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse1) Set(val *SymbolPriceTickerResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse1(val *SymbolPriceTickerResponse1) *NullableSymbolPriceTickerResponse1 {
	return &NullableSymbolPriceTickerResponse1{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
