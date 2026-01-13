/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponseResultAsksItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponseResultAsksItem{}

// OrderBookResponseResultAsksItem struct for OrderBookResponseResultAsksItem
type OrderBookResponseResultAsksItem struct {
	Items []string
}

// NewOrderBookResponseResultAsksItem instantiates a new OrderBookResponseResultAsksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponseResultAsksItem() *OrderBookResponseResultAsksItem {
	this := OrderBookResponseResultAsksItem{}
	return &this
}

// NewOrderBookResponseResultAsksItemWithDefaults instantiates a new OrderBookResponseResultAsksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseResultAsksItemWithDefaults() *OrderBookResponseResultAsksItem {
	this := OrderBookResponseResultAsksItem{}
	return &this
}

func (o OrderBookResponseResultAsksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponseResultAsksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OrderBookResponseResultAsksItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOrderBookResponseResultAsksItem struct {
	value OrderBookResponseResultAsksItem
	isSet bool
}

func (v NullableOrderBookResponseResultAsksItem) Get() OrderBookResponseResultAsksItem {
	return v.value
}

func (v *NullableOrderBookResponseResultAsksItem) Set(val OrderBookResponseResultAsksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponseResultAsksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponseResultAsksItem) Unset() {
	v.value = OrderBookResponseResultAsksItem{}
	v.isSet = false
}

func NewNullableOrderBookResponseResultAsksItem(val OrderBookResponseResultAsksItem) *NullableOrderBookResponseResultAsksItem {
	return &NullableOrderBookResponseResultAsksItem{value: val, isSet: true}
}

func (v NullableOrderBookResponseResultAsksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponseResultAsksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
