/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarkPriceKlineCandlestickDataResponseItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceKlineCandlestickDataResponseItem{}

// MarkPriceKlineCandlestickDataResponseItem struct for MarkPriceKlineCandlestickDataResponseItem
type MarkPriceKlineCandlestickDataResponseItem struct {
	Items []MarkPriceKlineCandlestickDataResponseItemInner
}

// NewMarkPriceKlineCandlestickDataResponseItem instantiates a new MarkPriceKlineCandlestickDataResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceKlineCandlestickDataResponseItem() *MarkPriceKlineCandlestickDataResponseItem {
	this := MarkPriceKlineCandlestickDataResponseItem{}
	return &this
}

// NewMarkPriceKlineCandlestickDataResponseItemWithDefaults instantiates a new MarkPriceKlineCandlestickDataResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceKlineCandlestickDataResponseItemWithDefaults() *MarkPriceKlineCandlestickDataResponseItem {
	this := MarkPriceKlineCandlestickDataResponseItem{}
	return &this
}

func (o MarkPriceKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceKlineCandlestickDataResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceKlineCandlestickDataResponseItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceKlineCandlestickDataResponseItem struct {
	value MarkPriceKlineCandlestickDataResponseItem
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickDataResponseItem) Get() MarkPriceKlineCandlestickDataResponseItem {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItem) Set(val MarkPriceKlineCandlestickDataResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickDataResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItem) Unset() {
	v.value = MarkPriceKlineCandlestickDataResponseItem{}
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickDataResponseItem(val MarkPriceKlineCandlestickDataResponseItem) *NullableMarkPriceKlineCandlestickDataResponseItem {
	return &NullableMarkPriceKlineCandlestickDataResponseItem{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickDataResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickDataResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
