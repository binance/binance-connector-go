/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceKlineCandlestickDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceKlineCandlestickDataItem{}

// IndexPriceKlineCandlestickDataItem struct for IndexPriceKlineCandlestickDataItem
type IndexPriceKlineCandlestickDataItem struct {
	Items []IndexPriceKlineCandlestickDataItemInner
}

// NewIndexPriceKlineCandlestickDataItem instantiates a new IndexPriceKlineCandlestickDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceKlineCandlestickDataItem() *IndexPriceKlineCandlestickDataItem {
	this := IndexPriceKlineCandlestickDataItem{}
	return &this
}

// NewIndexPriceKlineCandlestickDataItemWithDefaults instantiates a new IndexPriceKlineCandlestickDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceKlineCandlestickDataItemWithDefaults() *IndexPriceKlineCandlestickDataItem {
	this := IndexPriceKlineCandlestickDataItem{}
	return &this
}

func (o IndexPriceKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceKlineCandlestickDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *IndexPriceKlineCandlestickDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableIndexPriceKlineCandlestickDataItem struct {
	value IndexPriceKlineCandlestickDataItem
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataItem) Get() IndexPriceKlineCandlestickDataItem {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataItem) Set(val IndexPriceKlineCandlestickDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataItem) Unset() {
	v.value = IndexPriceKlineCandlestickDataItem{}
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataItem(val IndexPriceKlineCandlestickDataItem) *NullableIndexPriceKlineCandlestickDataItem {
	return &NullableIndexPriceKlineCandlestickDataItem{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
