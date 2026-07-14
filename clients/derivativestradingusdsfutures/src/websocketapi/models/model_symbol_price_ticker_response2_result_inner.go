/*
Futures (USDⓈ-M) WebSocket API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolPriceTickerResponse2ResultInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse2ResultInner{}

// SymbolPriceTickerResponse2ResultInner struct for SymbolPriceTickerResponse2ResultInner
type SymbolPriceTickerResponse2ResultInner struct {
	Symbol *string `json:"symbol,omitempty"`
	Price  *string `json:"price,omitempty"`
	// Transaction time
	Time                 *int64 `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SymbolPriceTickerResponse2ResultInner SymbolPriceTickerResponse2ResultInner

// NewSymbolPriceTickerResponse2ResultInner instantiates a new SymbolPriceTickerResponse2ResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse2ResultInner() *SymbolPriceTickerResponse2ResultInner {
	this := SymbolPriceTickerResponse2ResultInner{}
	return &this
}

// NewSymbolPriceTickerResponse2ResultInnerWithDefaults instantiates a new SymbolPriceTickerResponse2ResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponse2ResultInnerWithDefaults() *SymbolPriceTickerResponse2ResultInner {
	this := SymbolPriceTickerResponse2ResultInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse2ResultInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse2ResultInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse2ResultInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SymbolPriceTickerResponse2ResultInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse2ResultInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse2ResultInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse2ResultInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SymbolPriceTickerResponse2ResultInner) SetPrice(v string) {
	o.Price = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SymbolPriceTickerResponse2ResultInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SymbolPriceTickerResponse2ResultInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SymbolPriceTickerResponse2ResultInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SymbolPriceTickerResponse2ResultInner) SetTime(v int64) {
	o.Time = &v
}

func (o SymbolPriceTickerResponse2ResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse2ResultInner) ToMap() (map[string]interface{}, error) {
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

func (o *SymbolPriceTickerResponse2ResultInner) UnmarshalJSON(data []byte) (err error) {
	varSymbolPriceTickerResponse2ResultInner := _SymbolPriceTickerResponse2ResultInner{}

	err = json.Unmarshal(data, &varSymbolPriceTickerResponse2ResultInner)

	if err != nil {
		return err
	}

	*o = SymbolPriceTickerResponse2ResultInner(varSymbolPriceTickerResponse2ResultInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSymbolPriceTickerResponse2ResultInner struct {
	value *SymbolPriceTickerResponse2ResultInner
	isSet bool
}

func (v NullableSymbolPriceTickerResponse2ResultInner) Get() *SymbolPriceTickerResponse2ResultInner {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse2ResultInner) Set(val *SymbolPriceTickerResponse2ResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse2ResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse2ResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse2ResultInner(val *SymbolPriceTickerResponse2ResultInner) *NullableSymbolPriceTickerResponse2ResultInner {
	return &NullableSymbolPriceTickerResponse2ResultInner{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse2ResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse2ResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
