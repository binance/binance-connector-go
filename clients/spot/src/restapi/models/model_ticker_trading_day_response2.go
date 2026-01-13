/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TickerTradingDayResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerTradingDayResponse2{}

// TickerTradingDayResponse2 struct for TickerTradingDayResponse2
type TickerTradingDayResponse2 struct {
	Items []TickerTradingDayResponse2Inner
}

// NewTickerTradingDayResponse2 instantiates a new TickerTradingDayResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerTradingDayResponse2() *TickerTradingDayResponse2 {
	this := TickerTradingDayResponse2{}
	return &this
}

// NewTickerTradingDayResponse2WithDefaults instantiates a new TickerTradingDayResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerTradingDayResponse2WithDefaults() *TickerTradingDayResponse2 {
	this := TickerTradingDayResponse2{}
	return &this
}

func (o TickerTradingDayResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerTradingDayResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TickerTradingDayResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTickerTradingDayResponse2 struct {
	value TickerTradingDayResponse2
	isSet bool
}

func (v NullableTickerTradingDayResponse2) Get() TickerTradingDayResponse2 {
	return v.value
}

func (v *NullableTickerTradingDayResponse2) Set(val TickerTradingDayResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerTradingDayResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerTradingDayResponse2) Unset() {
	v.value = TickerTradingDayResponse2{}
	v.isSet = false
}

func NewNullableTickerTradingDayResponse2(val TickerTradingDayResponse2) *NullableTickerTradingDayResponse2 {
	return &NullableTickerTradingDayResponse2{value: val, isSet: true}
}

func (v NullableTickerTradingDayResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerTradingDayResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
