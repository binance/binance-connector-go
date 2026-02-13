/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCurrentMarginOrderCountUsageResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentMarginOrderCountUsageResponse{}

// QueryCurrentMarginOrderCountUsageResponse struct for QueryCurrentMarginOrderCountUsageResponse
type QueryCurrentMarginOrderCountUsageResponse struct {
	Items []QueryCurrentMarginOrderCountUsageResponseInner
}

// NewQueryCurrentMarginOrderCountUsageResponse instantiates a new QueryCurrentMarginOrderCountUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentMarginOrderCountUsageResponse() *QueryCurrentMarginOrderCountUsageResponse {
	this := QueryCurrentMarginOrderCountUsageResponse{}
	return &this
}

// NewQueryCurrentMarginOrderCountUsageResponseWithDefaults instantiates a new QueryCurrentMarginOrderCountUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentMarginOrderCountUsageResponseWithDefaults() *QueryCurrentMarginOrderCountUsageResponse {
	this := QueryCurrentMarginOrderCountUsageResponse{}
	return &this
}

func (o QueryCurrentMarginOrderCountUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentMarginOrderCountUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCurrentMarginOrderCountUsageResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCurrentMarginOrderCountUsageResponse struct {
	value QueryCurrentMarginOrderCountUsageResponse
	isSet bool
}

func (v NullableQueryCurrentMarginOrderCountUsageResponse) Get() QueryCurrentMarginOrderCountUsageResponse {
	return v.value
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponse) Set(val QueryCurrentMarginOrderCountUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentMarginOrderCountUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponse) Unset() {
	v.value = QueryCurrentMarginOrderCountUsageResponse{}
	v.isSet = false
}

func NewNullableQueryCurrentMarginOrderCountUsageResponse(val QueryCurrentMarginOrderCountUsageResponse) *NullableQueryCurrentMarginOrderCountUsageResponse {
	return &NullableQueryCurrentMarginOrderCountUsageResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentMarginOrderCountUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentMarginOrderCountUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
