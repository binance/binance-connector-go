/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryAllCurrentUmOpenAlgoOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentUmOpenAlgoOrdersResponse{}

// QueryAllCurrentUmOpenAlgoOrdersResponse struct for QueryAllCurrentUmOpenAlgoOrdersResponse
type QueryAllCurrentUmOpenAlgoOrdersResponse struct {
	Items []QueryAllCurrentUmOpenAlgoOrdersResponseInner
}

// NewQueryAllCurrentUmOpenAlgoOrdersResponse instantiates a new QueryAllCurrentUmOpenAlgoOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentUmOpenAlgoOrdersResponse() *QueryAllCurrentUmOpenAlgoOrdersResponse {
	this := QueryAllCurrentUmOpenAlgoOrdersResponse{}
	return &this
}

// NewQueryAllCurrentUmOpenAlgoOrdersResponseWithDefaults instantiates a new QueryAllCurrentUmOpenAlgoOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentUmOpenAlgoOrdersResponseWithDefaults() *QueryAllCurrentUmOpenAlgoOrdersResponse {
	this := QueryAllCurrentUmOpenAlgoOrdersResponse{}
	return &this
}

func (o QueryAllCurrentUmOpenAlgoOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentUmOpenAlgoOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllCurrentUmOpenAlgoOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllCurrentUmOpenAlgoOrdersResponse struct {
	value QueryAllCurrentUmOpenAlgoOrdersResponse
	isSet bool
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponse) Get() QueryAllCurrentUmOpenAlgoOrdersResponse {
	return v.value
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponse) Set(val QueryAllCurrentUmOpenAlgoOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponse) Unset() {
	v.value = QueryAllCurrentUmOpenAlgoOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllCurrentUmOpenAlgoOrdersResponse(val QueryAllCurrentUmOpenAlgoOrdersResponse) *NullableQueryAllCurrentUmOpenAlgoOrdersResponse {
	return &NullableQueryAllCurrentUmOpenAlgoOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
