/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerResponse1Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse1Result{}

// SymbolPriceTickerResponse1Result struct for SymbolPriceTickerResponse1Result
type SymbolPriceTickerResponse1Result struct {
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerResponse1Result SymbolPriceTickerResponse1Result

// NewSymbolPriceTickerResponse1Result instantiates a new SymbolPriceTickerResponse1Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse1Result() *SymbolPriceTickerResponse1Result {
	this := SymbolPriceTickerResponse1Result{}
	return &this
}

// NewSymbolPriceTickerResponse1ResultWithDefaults instantiates a new SymbolPriceTickerResponse1Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponse1ResultWithDefaults() *SymbolPriceTickerResponse1Result {
	this := SymbolPriceTickerResponse1Result{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1Result) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1Result) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1Result) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerResponse1Result) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1Result) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1Result) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1Result) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerResponse1Result) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse1Result) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse1Result) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse1Result) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerResponse1Result) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerResponse1Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse1Result) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolPriceTickerResponse1Result) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerResponse1Result := _SymbolPriceTickerResponse1Result{}

	err = json.Unmarshal(data, &varSymbolPriceTickerResponse1Result)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerResponse1Result(varSymbolPriceTickerResponse1Result)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerResponse1Result struct {
	value *SymbolPriceTickerResponse1Result
	isSet bool
}

func (v NullableSymbolPriceTickerResponse1Result) Get() *SymbolPriceTickerResponse1Result {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse1Result) Set(val *SymbolPriceTickerResponse1Result) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse1Result) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse1Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse1Result(val *SymbolPriceTickerResponse1Result) *NullableSymbolPriceTickerResponse1Result {
	return &NullableSymbolPriceTickerResponse1Result{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse1Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse1Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
