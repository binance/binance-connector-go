/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryCmPositionInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCmPositionInformationResponse{}

// QueryCmPositionInformationResponse struct for QueryCmPositionInformationResponse
type QueryCmPositionInformationResponse struct {
	Items []QueryCmPositionInformationResponseInner
}

// NewQueryCmPositionInformationResponse instantiates a new QueryCmPositionInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCmPositionInformationResponse() *QueryCmPositionInformationResponse {
	this := QueryCmPositionInformationResponse{}
	return &this
}

// NewQueryCmPositionInformationResponseWithDefaults instantiates a new QueryCmPositionInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCmPositionInformationResponseWithDefaults() *QueryCmPositionInformationResponse {
	this := QueryCmPositionInformationResponse{}
	return &this
}

func (o QueryCmPositionInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCmPositionInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryCmPositionInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryCmPositionInformationResponse struct {
	value QueryCmPositionInformationResponse
	isSet bool
}

func (v NullableQueryCmPositionInformationResponse) Get() QueryCmPositionInformationResponse {
	return v.value
}

func (v *NullableQueryCmPositionInformationResponse) Set(val QueryCmPositionInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCmPositionInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCmPositionInformationResponse) Unset() {
	v.value = QueryCmPositionInformationResponse{}
	v.isSet = false
}

func NewNullableQueryCmPositionInformationResponse(val QueryCmPositionInformationResponse) *NullableQueryCmPositionInformationResponse {
	return &NullableQueryCmPositionInformationResponse{value: val, isSet: true}
}

func (v NullableQueryCmPositionInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCmPositionInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
