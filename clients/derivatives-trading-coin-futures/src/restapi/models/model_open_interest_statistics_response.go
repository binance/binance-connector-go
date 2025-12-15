/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestStatisticsResponse{}

// OpenInterestStatisticsResponse struct for OpenInterestStatisticsResponse
type OpenInterestStatisticsResponse struct {
	Items []OpenInterestStatisticsResponseInner
}

// NewOpenInterestStatisticsResponse instantiates a new OpenInterestStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestStatisticsResponse() *OpenInterestStatisticsResponse {
	this := OpenInterestStatisticsResponse{}
	return &this
}

// NewOpenInterestStatisticsResponseWithDefaults instantiates a new OpenInterestStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestStatisticsResponseWithDefaults() *OpenInterestStatisticsResponse {
	this := OpenInterestStatisticsResponse{}
	return &this
}

func (o OpenInterestStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OpenInterestStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOpenInterestStatisticsResponse struct {
	value OpenInterestStatisticsResponse
	isSet bool
}

func (v NullableOpenInterestStatisticsResponse) Get() OpenInterestStatisticsResponse {
	return v.value
}

func (v *NullableOpenInterestStatisticsResponse) Set(val OpenInterestStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestStatisticsResponse) Unset() {
	v.value = OpenInterestStatisticsResponse{}
	v.isSet = false
}

func NewNullableOpenInterestStatisticsResponse(val OpenInterestStatisticsResponse) *NullableOpenInterestStatisticsResponse {
	return &NullableOpenInterestStatisticsResponse{value: val, isSet: true}
}

func (v NullableOpenInterestStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
