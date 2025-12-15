/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UiKlinesItem type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UiKlinesItem{}

// UiKlinesItem struct for UiKlinesItem
type UiKlinesItem struct {
	Items []KlinesItemInner
}

// NewUiKlinesItem instantiates a new UiKlinesItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiKlinesItem() *UiKlinesItem {
	this := UiKlinesItem{}
	return &this
}

// NewUiKlinesItemWithDefaults instantiates a new UiKlinesItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiKlinesItemWithDefaults() *UiKlinesItem {
	this := UiKlinesItem{}
	return &this
}

func (o UiKlinesItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiKlinesItem) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UiKlinesItem) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUiKlinesItem struct {
	value UiKlinesItem
	isSet bool
}

func (v NullableUiKlinesItem) Get() UiKlinesItem {
	return v.value
}

func (v *NullableUiKlinesItem) Set(val UiKlinesItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUiKlinesItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUiKlinesItem) Unset() {
	v.value = UiKlinesItem{}
	v.isSet = false
}

func NewNullableUiKlinesItem(val UiKlinesItem) *NullableUiKlinesItem {
	return &NullableUiKlinesItem{value: val, isSet: true}
}

func (v NullableUiKlinesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiKlinesItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
