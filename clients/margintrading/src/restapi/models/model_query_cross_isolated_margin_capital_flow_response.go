/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCrossIsolatedMarginCapitalFlowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCrossIsolatedMarginCapitalFlowResponse{}

// QueryCrossIsolatedMarginCapitalFlowResponse struct for QueryCrossIsolatedMarginCapitalFlowResponse
type QueryCrossIsolatedMarginCapitalFlowResponse struct {
	Items []QueryCrossIsolatedMarginCapitalFlowResponseInner
}

// NewQueryCrossIsolatedMarginCapitalFlowResponse instantiates a new QueryCrossIsolatedMarginCapitalFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCrossIsolatedMarginCapitalFlowResponse() *QueryCrossIsolatedMarginCapitalFlowResponse {
	this := QueryCrossIsolatedMarginCapitalFlowResponse{}
	return &this
}

// NewQueryCrossIsolatedMarginCapitalFlowResponseWithDefaults instantiates a new QueryCrossIsolatedMarginCapitalFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCrossIsolatedMarginCapitalFlowResponseWithDefaults() *QueryCrossIsolatedMarginCapitalFlowResponse {
	this := QueryCrossIsolatedMarginCapitalFlowResponse{}
	return &this
}

func (o QueryCrossIsolatedMarginCapitalFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCrossIsolatedMarginCapitalFlowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCrossIsolatedMarginCapitalFlowResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCrossIsolatedMarginCapitalFlowResponse struct {
	value QueryCrossIsolatedMarginCapitalFlowResponse
	isSet bool
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponse) Get() QueryCrossIsolatedMarginCapitalFlowResponse {
	return v.value
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponse) Set(val QueryCrossIsolatedMarginCapitalFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponse) Unset() {
	v.value = QueryCrossIsolatedMarginCapitalFlowResponse{}
	v.isSet = false
}

func NewNullableQueryCrossIsolatedMarginCapitalFlowResponse(val QueryCrossIsolatedMarginCapitalFlowResponse) *NullableQueryCrossIsolatedMarginCapitalFlowResponse {
	return &NullableQueryCrossIsolatedMarginCapitalFlowResponse{value: val, isSet: true}
}

func (v NullableQueryCrossIsolatedMarginCapitalFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCrossIsolatedMarginCapitalFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
