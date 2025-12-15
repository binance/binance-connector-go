/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RpiOrderBookResponseAsksItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiOrderBookResponseAsksItem{}

// RpiOrderBookResponseAsksItem struct for RpiOrderBookResponseAsksItem
type RpiOrderBookResponseAsksItem struct {
	Items []string
}

// NewRpiOrderBookResponseAsksItem instantiates a new RpiOrderBookResponseAsksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiOrderBookResponseAsksItem() *RpiOrderBookResponseAsksItem {
	this := RpiOrderBookResponseAsksItem{}
	return &this
}

// NewRpiOrderBookResponseAsksItemWithDefaults instantiates a new RpiOrderBookResponseAsksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiOrderBookResponseAsksItemWithDefaults() *RpiOrderBookResponseAsksItem {
	this := RpiOrderBookResponseAsksItem{}
	return &this
}

func (o RpiOrderBookResponseAsksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiOrderBookResponseAsksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RpiOrderBookResponseAsksItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRpiOrderBookResponseAsksItem struct {
	value RpiOrderBookResponseAsksItem
	isSet bool
}

func (v NullableRpiOrderBookResponseAsksItem) Get() RpiOrderBookResponseAsksItem {
	return v.value
}

func (v *NullableRpiOrderBookResponseAsksItem) Set(val RpiOrderBookResponseAsksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiOrderBookResponseAsksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiOrderBookResponseAsksItem) Unset() {
	v.value = RpiOrderBookResponseAsksItem{}
	v.isSet = false
}

func NewNullableRpiOrderBookResponseAsksItem(val RpiOrderBookResponseAsksItem) *NullableRpiOrderBookResponseAsksItem {
	return &NullableRpiOrderBookResponseAsksItem{value: val, isSet: true}
}

func (v NullableRpiOrderBookResponseAsksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiOrderBookResponseAsksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
