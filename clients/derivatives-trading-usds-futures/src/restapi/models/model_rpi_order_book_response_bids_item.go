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

// checks if the RpiOrderBookResponseBidsItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiOrderBookResponseBidsItem{}

// RpiOrderBookResponseBidsItem struct for RpiOrderBookResponseBidsItem
type RpiOrderBookResponseBidsItem struct {
	Items []string
}

// NewRpiOrderBookResponseBidsItem instantiates a new RpiOrderBookResponseBidsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiOrderBookResponseBidsItem() *RpiOrderBookResponseBidsItem {
	this := RpiOrderBookResponseBidsItem{}
	return &this
}

// NewRpiOrderBookResponseBidsItemWithDefaults instantiates a new RpiOrderBookResponseBidsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiOrderBookResponseBidsItemWithDefaults() *RpiOrderBookResponseBidsItem {
	this := RpiOrderBookResponseBidsItem{}
	return &this
}

func (o RpiOrderBookResponseBidsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiOrderBookResponseBidsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RpiOrderBookResponseBidsItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRpiOrderBookResponseBidsItem struct {
	value RpiOrderBookResponseBidsItem
	isSet bool
}

func (v NullableRpiOrderBookResponseBidsItem) Get() RpiOrderBookResponseBidsItem {
	return v.value
}

func (v *NullableRpiOrderBookResponseBidsItem) Set(val RpiOrderBookResponseBidsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiOrderBookResponseBidsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiOrderBookResponseBidsItem) Unset() {
	v.value = RpiOrderBookResponseBidsItem{}
	v.isSet = false
}

func NewNullableRpiOrderBookResponseBidsItem(val RpiOrderBookResponseBidsItem) *NullableRpiOrderBookResponseBidsItem {
	return &NullableRpiOrderBookResponseBidsItem{value: val, isSet: true}
}

func (v NullableRpiOrderBookResponseBidsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiOrderBookResponseBidsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
