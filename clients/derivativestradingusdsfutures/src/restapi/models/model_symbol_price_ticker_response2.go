/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse2{}

// SymbolPriceTickerResponse2 struct for SymbolPriceTickerResponse2
type SymbolPriceTickerResponse2 struct {
	Items []SymbolPriceTickerV2Response2Inner
}

// NewSymbolPriceTickerResponse2 instantiates a new SymbolPriceTickerResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse2() *SymbolPriceTickerResponse2 {
	this := SymbolPriceTickerResponse2{}
	return &this
}

// NewSymbolPriceTickerResponse2WithDefaults instantiates a new SymbolPriceTickerResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponse2WithDefaults() *SymbolPriceTickerResponse2 {
	this := SymbolPriceTickerResponse2{}
	return &this
}

func (o SymbolPriceTickerResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SymbolPriceTickerResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSymbolPriceTickerResponse2 struct {
	value SymbolPriceTickerResponse2
	isSet bool
}

func (v NullableSymbolPriceTickerResponse2) Get() SymbolPriceTickerResponse2 {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse2) Set(val SymbolPriceTickerResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse2) Unset() {
	v.value = SymbolPriceTickerResponse2{}
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse2(val SymbolPriceTickerResponse2) *NullableSymbolPriceTickerResponse2 {
	return &NullableSymbolPriceTickerResponse2{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
