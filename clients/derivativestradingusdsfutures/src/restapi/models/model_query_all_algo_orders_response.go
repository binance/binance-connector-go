/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryAllAlgoOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllAlgoOrdersResponse{}

// QueryAllAlgoOrdersResponse struct for QueryAllAlgoOrdersResponse
type QueryAllAlgoOrdersResponse struct {
	Items []QueryAllAlgoOrdersResponseInner
}

// NewQueryAllAlgoOrdersResponse instantiates a new QueryAllAlgoOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllAlgoOrdersResponse() *QueryAllAlgoOrdersResponse {
	this := QueryAllAlgoOrdersResponse{}
	return &this
}

// NewQueryAllAlgoOrdersResponseWithDefaults instantiates a new QueryAllAlgoOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllAlgoOrdersResponseWithDefaults() *QueryAllAlgoOrdersResponse {
	this := QueryAllAlgoOrdersResponse{}
	return &this
}

func (o QueryAllAlgoOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllAlgoOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryAllAlgoOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryAllAlgoOrdersResponse struct {
	value QueryAllAlgoOrdersResponse
	isSet bool
}

func (v NullableQueryAllAlgoOrdersResponse) Get() QueryAllAlgoOrdersResponse {
	return v.value
}

func (v *NullableQueryAllAlgoOrdersResponse) Set(val QueryAllAlgoOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllAlgoOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllAlgoOrdersResponse) Unset() {
	v.value = QueryAllAlgoOrdersResponse{}
	v.isSet = false
}

func NewNullableQueryAllAlgoOrdersResponse(val QueryAllAlgoOrdersResponse) *NullableQueryAllAlgoOrdersResponse {
	return &NullableQueryAllAlgoOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryAllAlgoOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllAlgoOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
