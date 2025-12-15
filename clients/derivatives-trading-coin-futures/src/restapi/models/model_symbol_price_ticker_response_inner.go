/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponseInner{}

// SymbolPriceTickerResponseInner struct for SymbolPriceTickerResponseInner
type SymbolPriceTickerResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Ps                   *string `json:"ps,omitempty"`
	Price                *string `json:"price,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerResponseInner SymbolPriceTickerResponseInner

// NewSymbolPriceTickerResponseInner instantiates a new SymbolPriceTickerResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponseInner() *SymbolPriceTickerResponseInner {
	this := SymbolPriceTickerResponseInner{}
	return &this
}

// NewSymbolPriceTickerResponseInnerWithDefaults instantiates a new SymbolPriceTickerResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponseInnerWithDefaults() *SymbolPriceTickerResponseInner {
	this := SymbolPriceTickerResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponseInner) GetPs() string {
	if o == nil || common.IsNil(o.Ps) {
		var ret string
		return ret
	}
	return *o.Ps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponseInner) GetPsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ps) {
		return nil, false
	}
	return o.Ps, true
}

// HasPs returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponseInner) HasPs() bool {
	if o != nil && !common.IsNil(o.Ps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *SymbolPriceTickerResponseInner) SetPs(v string) {
	o.Ps = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerResponseInner) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Ps) {
		toSerialize["ps"] = o.Ps
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

func (o *SymbolPriceTickerResponseInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerResponseInner := _SymbolPriceTickerResponseInner{}

	err = json.Unmarshal(data, &varSymbolPriceTickerResponseInner)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerResponseInner(varSymbolPriceTickerResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerResponseInner struct {
	value *SymbolPriceTickerResponseInner
	isSet bool
}

func (v NullableSymbolPriceTickerResponseInner) Get() *SymbolPriceTickerResponseInner {
	return v.value
}

func (v *NullableSymbolPriceTickerResponseInner) Set(val *SymbolPriceTickerResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponseInner(val *SymbolPriceTickerResponseInner) *NullableSymbolPriceTickerResponseInner {
	return &NullableSymbolPriceTickerResponseInner{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
