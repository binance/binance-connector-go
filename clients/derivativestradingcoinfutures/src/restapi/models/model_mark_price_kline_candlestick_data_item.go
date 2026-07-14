/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarkPriceKlineCandlestickDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceKlineCandlestickDataItem{}

// MarkPriceKlineCandlestickDataItem struct for MarkPriceKlineCandlestickDataItem
type MarkPriceKlineCandlestickDataItem struct {
	Items []IndexPriceKlineCandlestickDataItemInner
}

// NewMarkPriceKlineCandlestickDataItem instantiates a new MarkPriceKlineCandlestickDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceKlineCandlestickDataItem() *MarkPriceKlineCandlestickDataItem {
	this := MarkPriceKlineCandlestickDataItem{}
	return &this
}

// NewMarkPriceKlineCandlestickDataItemWithDefaults instantiates a new MarkPriceKlineCandlestickDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceKlineCandlestickDataItemWithDefaults() *MarkPriceKlineCandlestickDataItem {
	this := MarkPriceKlineCandlestickDataItem{}
	return &this
}

func (o MarkPriceKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceKlineCandlestickDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceKlineCandlestickDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceKlineCandlestickDataItem struct {
	value MarkPriceKlineCandlestickDataItem
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickDataItem) Get() MarkPriceKlineCandlestickDataItem {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickDataItem) Set(val MarkPriceKlineCandlestickDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickDataItem) Unset() {
	v.value = MarkPriceKlineCandlestickDataItem{}
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickDataItem(val MarkPriceKlineCandlestickDataItem) *NullableMarkPriceKlineCandlestickDataItem {
	return &NullableMarkPriceKlineCandlestickDataItem{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
