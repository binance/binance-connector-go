/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerResponse2{}

// TickerResponse2 struct for TickerResponse2
type TickerResponse2 struct {
	Items []TickerResponse2Inner
}

// NewTickerResponse2 instantiates a new TickerResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerResponse2() *TickerResponse2 {
	this := TickerResponse2{}
	return &this
}

// NewTickerResponse2WithDefaults instantiates a new TickerResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerResponse2WithDefaults() *TickerResponse2 {
	this := TickerResponse2{}
	return &this
}

func (o TickerResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TickerResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTickerResponse2 struct {
	value TickerResponse2
	isSet bool
}

func (v NullableTickerResponse2) Get() TickerResponse2 {
	return v.value
}

func (v *NullableTickerResponse2) Set(val TickerResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerResponse2) Unset() {
	v.value = TickerResponse2{}
	v.isSet = false
}

func NewNullableTickerResponse2(val TickerResponse2) *NullableTickerResponse2 {
	return &NullableTickerResponse2{value: val, isSet: true}
}

func (v NullableTickerResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
