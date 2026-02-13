/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUmPositionInformationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmPositionInformationResponse{}

// QueryUmPositionInformationResponse struct for QueryUmPositionInformationResponse
type QueryUmPositionInformationResponse struct {
	Items []QueryUmPositionInformationResponseInner
}

// NewQueryUmPositionInformationResponse instantiates a new QueryUmPositionInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmPositionInformationResponse() *QueryUmPositionInformationResponse {
	this := QueryUmPositionInformationResponse{}
	return &this
}

// NewQueryUmPositionInformationResponseWithDefaults instantiates a new QueryUmPositionInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmPositionInformationResponseWithDefaults() *QueryUmPositionInformationResponse {
	this := QueryUmPositionInformationResponse{}
	return &this
}

func (o QueryUmPositionInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmPositionInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUmPositionInformationResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUmPositionInformationResponse struct {
	value QueryUmPositionInformationResponse
	isSet bool
}

func (v NullableQueryUmPositionInformationResponse) Get() QueryUmPositionInformationResponse {
	return v.value
}

func (v *NullableQueryUmPositionInformationResponse) Set(val QueryUmPositionInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmPositionInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmPositionInformationResponse) Unset() {
	v.value = QueryUmPositionInformationResponse{}
	v.isSet = false
}

func NewNullableQueryUmPositionInformationResponse(val QueryUmPositionInformationResponse) *NullableQueryUmPositionInformationResponse {
	return &NullableQueryUmPositionInformationResponse{value: val, isSet: true}
}

func (v NullableQueryUmPositionInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmPositionInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
