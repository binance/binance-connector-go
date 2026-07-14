/*
Futures (USDⓈ-M) WebSocket Market Streams

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AssetIndexResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AssetIndexResponse{}

// AssetIndexResponse struct for AssetIndexResponse
type AssetIndexResponse struct {
	Items []AssetIndexResponseInner
}

// NewAssetIndexResponse instantiates a new AssetIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetIndexResponse() *AssetIndexResponse {
	this := AssetIndexResponse{}
	return &this
}

// NewAssetIndexResponseWithDefaults instantiates a new AssetIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetIndexResponseWithDefaults() *AssetIndexResponse {
	this := AssetIndexResponse{}
	return &this
}

func (o AssetIndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetIndexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AssetIndexResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAssetIndexResponse struct {
	value AssetIndexResponse
	isSet bool
}

func (v NullableAssetIndexResponse) Get() AssetIndexResponse {
	return v.value
}

func (v *NullableAssetIndexResponse) Set(val AssetIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetIndexResponse) Unset() {
	v.value = AssetIndexResponse{}
	v.isSet = false
}

func NewNullableAssetIndexResponse(val AssetIndexResponse) *NullableAssetIndexResponse {
	return &NullableAssetIndexResponse{value: val, isSet: true}
}

func (v NullableAssetIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
