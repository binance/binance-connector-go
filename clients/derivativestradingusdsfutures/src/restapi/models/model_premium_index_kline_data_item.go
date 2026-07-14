/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PremiumIndexKlineDataItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PremiumIndexKlineDataItem{}

// PremiumIndexKlineDataItem struct for PremiumIndexKlineDataItem
type PremiumIndexKlineDataItem struct {
	Items []PremiumIndexKlineDataItemInner
}

// NewPremiumIndexKlineDataItem instantiates a new PremiumIndexKlineDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPremiumIndexKlineDataItem() *PremiumIndexKlineDataItem {
	this := PremiumIndexKlineDataItem{}
	return &this
}

// NewPremiumIndexKlineDataItemWithDefaults instantiates a new PremiumIndexKlineDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPremiumIndexKlineDataItemWithDefaults() *PremiumIndexKlineDataItem {
	this := PremiumIndexKlineDataItem{}
	return &this
}

func (o PremiumIndexKlineDataItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PremiumIndexKlineDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PremiumIndexKlineDataItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePremiumIndexKlineDataItem struct {
	value PremiumIndexKlineDataItem
	isSet bool
}

func (v NullablePremiumIndexKlineDataItem) Get() PremiumIndexKlineDataItem {
	return v.value
}

func (v *NullablePremiumIndexKlineDataItem) Set(val PremiumIndexKlineDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePremiumIndexKlineDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePremiumIndexKlineDataItem) Unset() {
	v.value = PremiumIndexKlineDataItem{}
	v.isSet = false
}

func NewNullablePremiumIndexKlineDataItem(val PremiumIndexKlineDataItem) *NullablePremiumIndexKlineDataItem {
	return &NullablePremiumIndexKlineDataItem{value: val, isSet: true}
}

func (v NullablePremiumIndexKlineDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePremiumIndexKlineDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
