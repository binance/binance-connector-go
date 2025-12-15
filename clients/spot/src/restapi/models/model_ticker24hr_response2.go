/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24hrResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrResponse2{}

// Ticker24hrResponse2 struct for Ticker24hrResponse2
type Ticker24hrResponse2 struct {
	Items []Ticker24hrResponse2Inner
}

// NewTicker24hrResponse2 instantiates a new Ticker24hrResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrResponse2() *Ticker24hrResponse2 {
	this := Ticker24hrResponse2{}
	return &this
}

// NewTicker24hrResponse2WithDefaults instantiates a new Ticker24hrResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrResponse2WithDefaults() *Ticker24hrResponse2 {
	this := Ticker24hrResponse2{}
	return &this
}

func (o Ticker24hrResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *Ticker24hrResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTicker24hrResponse2 struct {
	value Ticker24hrResponse2
	isSet bool
}

func (v NullableTicker24hrResponse2) Get() Ticker24hrResponse2 {
	return v.value
}

func (v *NullableTicker24hrResponse2) Set(val Ticker24hrResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrResponse2) Unset() {
	v.value = Ticker24hrResponse2{}
	v.isSet = false
}

func NewNullableTicker24hrResponse2(val Ticker24hrResponse2) *NullableTicker24hrResponse2 {
	return &NullableTicker24hrResponse2{value: val, isSet: true}
}

func (v NullableTicker24hrResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
