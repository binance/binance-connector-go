/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestResponse{}

// OpenInterestResponse struct for OpenInterestResponse
type OpenInterestResponse struct {
	Items []OpenInterestResponseInner
}

// NewOpenInterestResponse instantiates a new OpenInterestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestResponse() *OpenInterestResponse {
	this := OpenInterestResponse{}
	return &this
}

// NewOpenInterestResponseWithDefaults instantiates a new OpenInterestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestResponseWithDefaults() *OpenInterestResponse {
	this := OpenInterestResponse{}
	return &this
}

func (o OpenInterestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OpenInterestResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOpenInterestResponse struct {
	value OpenInterestResponse
	isSet bool
}

func (v NullableOpenInterestResponse) Get() OpenInterestResponse {
	return v.value
}

func (v *NullableOpenInterestResponse) Set(val OpenInterestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestResponse) Unset() {
	v.value = OpenInterestResponse{}
	v.isSet = false
}

func NewNullableOpenInterestResponse(val OpenInterestResponse) *NullableOpenInterestResponse {
	return &NullableOpenInterestResponse{value: val, isSet: true}
}

func (v NullableOpenInterestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
