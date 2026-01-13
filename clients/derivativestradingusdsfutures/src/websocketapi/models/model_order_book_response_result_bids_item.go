/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponseResultBidsItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponseResultBidsItem{}

// OrderBookResponseResultBidsItem struct for OrderBookResponseResultBidsItem
type OrderBookResponseResultBidsItem struct {
	Items []string
}

// NewOrderBookResponseResultBidsItem instantiates a new OrderBookResponseResultBidsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponseResultBidsItem() *OrderBookResponseResultBidsItem {
	this := OrderBookResponseResultBidsItem{}
	return &this
}

// NewOrderBookResponseResultBidsItemWithDefaults instantiates a new OrderBookResponseResultBidsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseResultBidsItemWithDefaults() *OrderBookResponseResultBidsItem {
	this := OrderBookResponseResultBidsItem{}
	return &this
}

func (o OrderBookResponseResultBidsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponseResultBidsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OrderBookResponseResultBidsItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOrderBookResponseResultBidsItem struct {
	value OrderBookResponseResultBidsItem
	isSet bool
}

func (v NullableOrderBookResponseResultBidsItem) Get() OrderBookResponseResultBidsItem {
	return v.value
}

func (v *NullableOrderBookResponseResultBidsItem) Set(val OrderBookResponseResultBidsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponseResultBidsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponseResultBidsItem) Unset() {
	v.value = OrderBookResponseResultBidsItem{}
	v.isSet = false
}

func NewNullableOrderBookResponseResultBidsItem(val OrderBookResponseResultBidsItem) *NullableOrderBookResponseResultBidsItem {
	return &NullableOrderBookResponseResultBidsItem{value: val, isSet: true}
}

func (v NullableOrderBookResponseResultBidsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponseResultBidsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
