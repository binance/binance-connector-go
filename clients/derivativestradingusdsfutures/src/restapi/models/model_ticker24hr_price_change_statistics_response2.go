/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24hrPriceChangeStatisticsResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrPriceChangeStatisticsResponse2{}

// Ticker24hrPriceChangeStatisticsResponse2 struct for Ticker24hrPriceChangeStatisticsResponse2
type Ticker24hrPriceChangeStatisticsResponse2 struct {
	Items []Ticker24hrPriceChangeStatisticsResponse2Inner
}

// NewTicker24hrPriceChangeStatisticsResponse2 instantiates a new Ticker24hrPriceChangeStatisticsResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrPriceChangeStatisticsResponse2() *Ticker24hrPriceChangeStatisticsResponse2 {
	this := Ticker24hrPriceChangeStatisticsResponse2{}
	return &this
}

// NewTicker24hrPriceChangeStatisticsResponse2WithDefaults instantiates a new Ticker24hrPriceChangeStatisticsResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrPriceChangeStatisticsResponse2WithDefaults() *Ticker24hrPriceChangeStatisticsResponse2 {
	this := Ticker24hrPriceChangeStatisticsResponse2{}
	return &this
}

func (o Ticker24hrPriceChangeStatisticsResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrPriceChangeStatisticsResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *Ticker24hrPriceChangeStatisticsResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTicker24hrPriceChangeStatisticsResponse2 struct {
	value Ticker24hrPriceChangeStatisticsResponse2
	isSet bool
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2) Get() Ticker24hrPriceChangeStatisticsResponse2 {
	return v.value
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2) Set(val Ticker24hrPriceChangeStatisticsResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2) Unset() {
	v.value = Ticker24hrPriceChangeStatisticsResponse2{}
	v.isSet = false
}

func NewNullableTicker24hrPriceChangeStatisticsResponse2(val Ticker24hrPriceChangeStatisticsResponse2) *NullableTicker24hrPriceChangeStatisticsResponse2 {
	return &NullableTicker24hrPriceChangeStatisticsResponse2{value: val, isSet: true}
}

func (v NullableTicker24hrPriceChangeStatisticsResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
