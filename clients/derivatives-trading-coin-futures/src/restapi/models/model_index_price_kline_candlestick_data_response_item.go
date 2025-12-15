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

// checks if the IndexPriceKlineCandlestickDataResponseItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceKlineCandlestickDataResponseItem{}

// IndexPriceKlineCandlestickDataResponseItem struct for IndexPriceKlineCandlestickDataResponseItem
type IndexPriceKlineCandlestickDataResponseItem struct {
	Items []IndexPriceKlineCandlestickDataResponseItemInner
}

// NewIndexPriceKlineCandlestickDataResponseItem instantiates a new IndexPriceKlineCandlestickDataResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceKlineCandlestickDataResponseItem() *IndexPriceKlineCandlestickDataResponseItem {
	this := IndexPriceKlineCandlestickDataResponseItem{}
	return &this
}

// NewIndexPriceKlineCandlestickDataResponseItemWithDefaults instantiates a new IndexPriceKlineCandlestickDataResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceKlineCandlestickDataResponseItemWithDefaults() *IndexPriceKlineCandlestickDataResponseItem {
	this := IndexPriceKlineCandlestickDataResponseItem{}
	return &this
}

func (o IndexPriceKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceKlineCandlestickDataResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *IndexPriceKlineCandlestickDataResponseItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableIndexPriceKlineCandlestickDataResponseItem struct {
	value IndexPriceKlineCandlestickDataResponseItem
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataResponseItem) Get() IndexPriceKlineCandlestickDataResponseItem {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItem) Set(val IndexPriceKlineCandlestickDataResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItem) Unset() {
	v.value = IndexPriceKlineCandlestickDataResponseItem{}
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataResponseItem(val IndexPriceKlineCandlestickDataResponseItem) *NullableIndexPriceKlineCandlestickDataResponseItem {
	return &NullableIndexPriceKlineCandlestickDataResponseItem{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
