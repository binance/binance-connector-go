/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AssetIndexResponse2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetIndexResponse2{}

// AssetIndexResponse2 struct for AssetIndexResponse2
type AssetIndexResponse2 struct {
	Items []AssetIndexResponse2Inner
}

// NewAssetIndexResponse2 instantiates a new AssetIndexResponse2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetIndexResponse2() *AssetIndexResponse2 {
	this := AssetIndexResponse2{}
	return &this
}

// NewAssetIndexResponse2WithDefaults instantiates a new AssetIndexResponse2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetIndexResponse2WithDefaults() *AssetIndexResponse2 {
	this := AssetIndexResponse2{}
	return &this
}

func (o AssetIndexResponse2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetIndexResponse2) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AssetIndexResponse2) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAssetIndexResponse2 struct {
	value AssetIndexResponse2
	isSet bool
}

func (v NullableAssetIndexResponse2) Get() AssetIndexResponse2 {
	return v.value
}

func (v *NullableAssetIndexResponse2) Set(val AssetIndexResponse2) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetIndexResponse2) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetIndexResponse2) Unset() {
	v.value = AssetIndexResponse2{}
	v.isSet = false
}

func NewNullableAssetIndexResponse2(val AssetIndexResponse2) *NullableAssetIndexResponse2 {
	return &NullableAssetIndexResponse2{value: val, isSet: true}
}

func (v NullableAssetIndexResponse2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetIndexResponse2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
