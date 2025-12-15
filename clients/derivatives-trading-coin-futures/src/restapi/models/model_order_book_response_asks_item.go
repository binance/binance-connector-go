/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponseAsksItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponseAsksItem{}

// OrderBookResponseAsksItem struct for OrderBookResponseAsksItem
type OrderBookResponseAsksItem struct {
	Items []string
}

// NewOrderBookResponseAsksItem instantiates a new OrderBookResponseAsksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponseAsksItem() *OrderBookResponseAsksItem {
	this := OrderBookResponseAsksItem{}
	return &this
}

// NewOrderBookResponseAsksItemWithDefaults instantiates a new OrderBookResponseAsksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseAsksItemWithDefaults() *OrderBookResponseAsksItem {
	this := OrderBookResponseAsksItem{}
	return &this
}

func (o OrderBookResponseAsksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponseAsksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OrderBookResponseAsksItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOrderBookResponseAsksItem struct {
	value OrderBookResponseAsksItem
	isSet bool
}

func (v NullableOrderBookResponseAsksItem) Get() OrderBookResponseAsksItem {
	return v.value
}

func (v *NullableOrderBookResponseAsksItem) Set(val OrderBookResponseAsksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponseAsksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponseAsksItem) Unset() {
	v.value = OrderBookResponseAsksItem{}
	v.isSet = false
}

func NewNullableOrderBookResponseAsksItem(val OrderBookResponseAsksItem) *NullableOrderBookResponseAsksItem {
	return &NullableOrderBookResponseAsksItem{value: val, isSet: true}
}

func (v NullableOrderBookResponseAsksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponseAsksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
