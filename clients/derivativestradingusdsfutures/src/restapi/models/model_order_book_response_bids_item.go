/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponseBidsItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponseBidsItem{}

// OrderBookResponseBidsItem struct for OrderBookResponseBidsItem
type OrderBookResponseBidsItem struct {
	Items []string
}

// NewOrderBookResponseBidsItem instantiates a new OrderBookResponseBidsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponseBidsItem() *OrderBookResponseBidsItem {
	this := OrderBookResponseBidsItem{}
	return &this
}

// NewOrderBookResponseBidsItemWithDefaults instantiates a new OrderBookResponseBidsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseBidsItemWithDefaults() *OrderBookResponseBidsItem {
	this := OrderBookResponseBidsItem{}
	return &this
}

func (o OrderBookResponseBidsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponseBidsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OrderBookResponseBidsItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOrderBookResponseBidsItem struct {
	value OrderBookResponseBidsItem
	isSet bool
}

func (v NullableOrderBookResponseBidsItem) Get() OrderBookResponseBidsItem {
	return v.value
}

func (v *NullableOrderBookResponseBidsItem) Set(val OrderBookResponseBidsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponseBidsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponseBidsItem) Unset() {
	v.value = OrderBookResponseBidsItem{}
	v.isSet = false
}

func NewNullableOrderBookResponseBidsItem(val OrderBookResponseBidsItem) *NullableOrderBookResponseBidsItem {
	return &NullableOrderBookResponseBidsItem{value: val, isSet: true}
}

func (v NullableOrderBookResponseBidsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponseBidsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
