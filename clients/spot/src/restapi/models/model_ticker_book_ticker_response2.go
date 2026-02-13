/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerBookTickerResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerBookTickerResponse2{}

// TickerBookTickerResponse2 struct for TickerBookTickerResponse2
type TickerBookTickerResponse2 struct {
	Items []TickerBookTickerResponse2Inner
}

// NewTickerBookTickerResponse2 instantiates a new TickerBookTickerResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerBookTickerResponse2() *TickerBookTickerResponse2 {
	this := TickerBookTickerResponse2{}
	return &this
}

// NewTickerBookTickerResponse2WithDefaults instantiates a new TickerBookTickerResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerBookTickerResponse2WithDefaults() *TickerBookTickerResponse2 {
	this := TickerBookTickerResponse2{}
	return &this
}

func (o TickerBookTickerResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerBookTickerResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TickerBookTickerResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTickerBookTickerResponse2 struct {
	value TickerBookTickerResponse2
	isSet bool
}

func (v NullableTickerBookTickerResponse2) Get() TickerBookTickerResponse2 {
	return v.value
}

func (v *NullableTickerBookTickerResponse2) Set(val TickerBookTickerResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerBookTickerResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerBookTickerResponse2) Unset() {
	v.value = TickerBookTickerResponse2{}
	v.isSet = false
}

func NewNullableTickerBookTickerResponse2(val TickerBookTickerResponse2) *NullableTickerBookTickerResponse2 {
	return &NullableTickerBookTickerResponse2{value: val, isSet: true}
}

func (v NullableTickerBookTickerResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerBookTickerResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
