/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PremiumIndexKlineDataResponseItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PremiumIndexKlineDataResponseItem{}

// PremiumIndexKlineDataResponseItem struct for PremiumIndexKlineDataResponseItem
type PremiumIndexKlineDataResponseItem struct {
	Items []PremiumIndexKlineDataResponseItemInner
}

// NewPremiumIndexKlineDataResponseItem instantiates a new PremiumIndexKlineDataResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPremiumIndexKlineDataResponseItem() *PremiumIndexKlineDataResponseItem {
	this := PremiumIndexKlineDataResponseItem{}
	return &this
}

// NewPremiumIndexKlineDataResponseItemWithDefaults instantiates a new PremiumIndexKlineDataResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPremiumIndexKlineDataResponseItemWithDefaults() *PremiumIndexKlineDataResponseItem {
	this := PremiumIndexKlineDataResponseItem{}
	return &this
}

func (o PremiumIndexKlineDataResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PremiumIndexKlineDataResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PremiumIndexKlineDataResponseItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePremiumIndexKlineDataResponseItem struct {
	value PremiumIndexKlineDataResponseItem
	isSet bool
}

func (v NullablePremiumIndexKlineDataResponseItem) Get() PremiumIndexKlineDataResponseItem {
	return v.value
}

func (v *NullablePremiumIndexKlineDataResponseItem) Set(val PremiumIndexKlineDataResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePremiumIndexKlineDataResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePremiumIndexKlineDataResponseItem) Unset() {
	v.value = PremiumIndexKlineDataResponseItem{}
	v.isSet = false
}

func NewNullablePremiumIndexKlineDataResponseItem(val PremiumIndexKlineDataResponseItem) *NullablePremiumIndexKlineDataResponseItem {
	return &NullablePremiumIndexKlineDataResponseItem{value: val, isSet: true}
}

func (v NullablePremiumIndexKlineDataResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePremiumIndexKlineDataResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
