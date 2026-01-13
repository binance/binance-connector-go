/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TickerPriceResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerPriceResponse2{}

// TickerPriceResponse2 struct for TickerPriceResponse2
type TickerPriceResponse2 struct {
	Items []TickerPriceResponse2Inner
}

// NewTickerPriceResponse2 instantiates a new TickerPriceResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerPriceResponse2() *TickerPriceResponse2 {
	this := TickerPriceResponse2{}
	return &this
}

// NewTickerPriceResponse2WithDefaults instantiates a new TickerPriceResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerPriceResponse2WithDefaults() *TickerPriceResponse2 {
	this := TickerPriceResponse2{}
	return &this
}

func (o TickerPriceResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerPriceResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TickerPriceResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTickerPriceResponse2 struct {
	value TickerPriceResponse2
	isSet bool
}

func (v NullableTickerPriceResponse2) Get() TickerPriceResponse2 {
	return v.value
}

func (v *NullableTickerPriceResponse2) Set(val TickerPriceResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerPriceResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerPriceResponse2) Unset() {
	v.value = TickerPriceResponse2{}
	v.isSet = false
}

func NewNullableTickerPriceResponse2(val TickerPriceResponse2) *NullableTickerPriceResponse2 {
	return &NullableTickerPriceResponse2{value: val, isSet: true}
}

func (v NullableTickerPriceResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerPriceResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
