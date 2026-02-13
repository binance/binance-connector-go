/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolOrderBookTickerResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse2{}

// SymbolOrderBookTickerResponse2 struct for SymbolOrderBookTickerResponse2
type SymbolOrderBookTickerResponse2 struct {
	Items []SymbolOrderBookTickerResponse2Inner
}

// NewSymbolOrderBookTickerResponse2 instantiates a new SymbolOrderBookTickerResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse2() *SymbolOrderBookTickerResponse2 {
	this := SymbolOrderBookTickerResponse2{}
	return &this
}

// NewSymbolOrderBookTickerResponse2WithDefaults instantiates a new SymbolOrderBookTickerResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponse2WithDefaults() *SymbolOrderBookTickerResponse2 {
	this := SymbolOrderBookTickerResponse2{}
	return &this
}

func (o SymbolOrderBookTickerResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SymbolOrderBookTickerResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSymbolOrderBookTickerResponse2 struct {
	value SymbolOrderBookTickerResponse2
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse2) Get() SymbolOrderBookTickerResponse2 {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse2) Set(val SymbolOrderBookTickerResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse2) Unset() {
	v.value = SymbolOrderBookTickerResponse2{}
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse2(val SymbolOrderBookTickerResponse2) *NullableSymbolOrderBookTickerResponse2 {
	return &NullableSymbolOrderBookTickerResponse2{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
