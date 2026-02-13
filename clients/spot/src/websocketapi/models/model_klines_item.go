/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlinesItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlinesItem{}

// KlinesItem struct for KlinesItem
type KlinesItem struct {
	Items []KlinesItemInner
}

// NewKlinesItem instantiates a new KlinesItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlinesItem() *KlinesItem {
	this := KlinesItem{}
	return &this
}

// NewKlinesItemWithDefaults instantiates a new KlinesItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlinesItemWithDefaults() *KlinesItem {
	this := KlinesItem{}
	return &this
}

func (o KlinesItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlinesItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *KlinesItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableKlinesItem struct {
	value KlinesItem
	isSet bool
}

func (v NullableKlinesItem) Get() KlinesItem {
	return v.value
}

func (v *NullableKlinesItem) Set(val KlinesItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesItem) Unset() {
	v.value = KlinesItem{}
	v.isSet = false
}

func NewNullableKlinesItem(val KlinesItem) *NullableKlinesItem {
	return &NullableKlinesItem{value: val, isSet: true}
}

func (v NullableKlinesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
