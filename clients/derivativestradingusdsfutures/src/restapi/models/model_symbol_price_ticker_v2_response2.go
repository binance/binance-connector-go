/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolPriceTickerV2Response2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerV2Response2{}

// SymbolPriceTickerV2Response2 struct for SymbolPriceTickerV2Response2
type SymbolPriceTickerV2Response2 struct {
	Items []SymbolPriceTickerV2Response2Inner
}

// NewSymbolPriceTickerV2Response2 instantiates a new SymbolPriceTickerV2Response2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerV2Response2() *SymbolPriceTickerV2Response2 {
	this := SymbolPriceTickerV2Response2{}
	return &this
}

// NewSymbolPriceTickerV2Response2WithDefaults instantiates a new SymbolPriceTickerV2Response2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerV2Response2WithDefaults() *SymbolPriceTickerV2Response2 {
	this := SymbolPriceTickerV2Response2{}
	return &this
}

func (o SymbolPriceTickerV2Response2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerV2Response2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SymbolPriceTickerV2Response2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSymbolPriceTickerV2Response2 struct {
	value SymbolPriceTickerV2Response2
	isSet bool
}

func (v NullableSymbolPriceTickerV2Response2) Get() SymbolPriceTickerV2Response2 {
	return v.value
}

func (v *NullableSymbolPriceTickerV2Response2) Set(val SymbolPriceTickerV2Response2) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerV2Response2) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerV2Response2) Unset() {
	v.value = SymbolPriceTickerV2Response2{}
	v.isSet = false
}

func NewNullableSymbolPriceTickerV2Response2(val SymbolPriceTickerV2Response2) *NullableSymbolPriceTickerV2Response2 {
	return &NullableSymbolPriceTickerV2Response2{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerV2Response2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerV2Response2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
