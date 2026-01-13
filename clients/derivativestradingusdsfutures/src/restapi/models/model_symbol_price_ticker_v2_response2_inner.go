/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerV2Response2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerV2Response2Inner{}

// SymbolPriceTickerV2Response2Inner struct for SymbolPriceTickerV2Response2Inner
type SymbolPriceTickerV2Response2Inner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerV2Response2Inner SymbolPriceTickerV2Response2Inner

// NewSymbolPriceTickerV2Response2Inner instantiates a new SymbolPriceTickerV2Response2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerV2Response2Inner() *SymbolPriceTickerV2Response2Inner {
	this := SymbolPriceTickerV2Response2Inner{}
	return &this
}

// NewSymbolPriceTickerV2Response2InnerWithDefaults instantiates a new SymbolPriceTickerV2Response2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerV2Response2InnerWithDefaults() *SymbolPriceTickerV2Response2Inner {
	this := SymbolPriceTickerV2Response2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerV2Response2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response2Inner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response2Inner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response2Inner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerV2Response2Inner) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerV2Response2Inner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerV2Response2Inner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerV2Response2Inner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerV2Response2Inner) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerV2Response2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerV2Response2Inner) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolPriceTickerV2Response2Inner) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerV2Response2Inner := _SymbolPriceTickerV2Response2Inner{}

	err = json.Unmarshal(data, &varSymbolPriceTickerV2Response2Inner)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerV2Response2Inner(varSymbolPriceTickerV2Response2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerV2Response2Inner struct {
	value *SymbolPriceTickerV2Response2Inner
	isSet bool
}

func (v NullableSymbolPriceTickerV2Response2Inner) Get() *SymbolPriceTickerV2Response2Inner {
	return v.value
}

func (v *NullableSymbolPriceTickerV2Response2Inner) Set(val *SymbolPriceTickerV2Response2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerV2Response2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerV2Response2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerV2Response2Inner(val *SymbolPriceTickerV2Response2Inner) *NullableSymbolPriceTickerV2Response2Inner {
	return &NullableSymbolPriceTickerV2Response2Inner{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerV2Response2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerV2Response2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
