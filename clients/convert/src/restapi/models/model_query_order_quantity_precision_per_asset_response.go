/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryOrderQuantityPrecisionPerAssetResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderQuantityPrecisionPerAssetResponse{}

// QueryOrderQuantityPrecisionPerAssetResponse struct for QueryOrderQuantityPrecisionPerAssetResponse
type QueryOrderQuantityPrecisionPerAssetResponse struct {
	Items []QueryOrderQuantityPrecisionPerAssetResponseInner
}

// NewQueryOrderQuantityPrecisionPerAssetResponse instantiates a new QueryOrderQuantityPrecisionPerAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderQuantityPrecisionPerAssetResponse() *QueryOrderQuantityPrecisionPerAssetResponse {
	this := QueryOrderQuantityPrecisionPerAssetResponse{}
	return &this
}

// NewQueryOrderQuantityPrecisionPerAssetResponseWithDefaults instantiates a new QueryOrderQuantityPrecisionPerAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderQuantityPrecisionPerAssetResponseWithDefaults() *QueryOrderQuantityPrecisionPerAssetResponse {
	this := QueryOrderQuantityPrecisionPerAssetResponse{}
	return &this
}

func (o QueryOrderQuantityPrecisionPerAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderQuantityPrecisionPerAssetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryOrderQuantityPrecisionPerAssetResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryOrderQuantityPrecisionPerAssetResponse struct {
	value QueryOrderQuantityPrecisionPerAssetResponse
	isSet bool
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponse) Get() QueryOrderQuantityPrecisionPerAssetResponse {
	return v.value
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponse) Set(val QueryOrderQuantityPrecisionPerAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponse) Unset() {
	v.value = QueryOrderQuantityPrecisionPerAssetResponse{}
	v.isSet = false
}

func NewNullableQueryOrderQuantityPrecisionPerAssetResponse(val QueryOrderQuantityPrecisionPerAssetResponse) *NullableQueryOrderQuantityPrecisionPerAssetResponse {
	return &NullableQueryOrderQuantityPrecisionPerAssetResponse{value: val, isSet: true}
}

func (v NullableQueryOrderQuantityPrecisionPerAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderQuantityPrecisionPerAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
