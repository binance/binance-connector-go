/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentOpenOptionOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentOpenOptionOrdersResponse{}

// QueryCurrentOpenOptionOrdersResponse struct for QueryCurrentOpenOptionOrdersResponse
type QueryCurrentOpenOptionOrdersResponse struct {
	Items []QueryCurrentOpenOptionOrdersResponseInner
}

// NewQueryCurrentOpenOptionOrdersResponse instantiates a new QueryCurrentOpenOptionOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentOpenOptionOrdersResponse() *QueryCurrentOpenOptionOrdersResponse {
	this := QueryCurrentOpenOptionOrdersResponse{}
	return &this
}

// NewQueryCurrentOpenOptionOrdersResponseWithDefaults instantiates a new QueryCurrentOpenOptionOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentOpenOptionOrdersResponseWithDefaults() *QueryCurrentOpenOptionOrdersResponse {
	this := QueryCurrentOpenOptionOrdersResponse{}
	return &this
}

func (o QueryCurrentOpenOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentOpenOptionOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCurrentOpenOptionOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCurrentOpenOptionOrdersResponse struct {
	value QueryCurrentOpenOptionOrdersResponse
	isSet bool
}

func (v NullableQueryCurrentOpenOptionOrdersResponse) Get() QueryCurrentOpenOptionOrdersResponse {
	return v.value
}

func (v *NullableQueryCurrentOpenOptionOrdersResponse) Set(val QueryCurrentOpenOptionOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentOpenOptionOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentOpenOptionOrdersResponse) Unset() {
	v.value = QueryCurrentOpenOptionOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryCurrentOpenOptionOrdersResponse(val QueryCurrentOpenOptionOrdersResponse) *NullableQueryCurrentOpenOptionOrdersResponse {
	return &NullableQueryCurrentOpenOptionOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentOpenOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentOpenOptionOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
