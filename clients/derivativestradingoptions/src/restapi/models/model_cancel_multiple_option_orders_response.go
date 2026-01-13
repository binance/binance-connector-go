/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelMultipleOptionOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMultipleOptionOrdersResponse{}

// CancelMultipleOptionOrdersResponse struct for CancelMultipleOptionOrdersResponse
type CancelMultipleOptionOrdersResponse struct {
	Items []CancelMultipleOptionOrdersResponseInner
}

// NewCancelMultipleOptionOrdersResponse instantiates a new CancelMultipleOptionOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMultipleOptionOrdersResponse() *CancelMultipleOptionOrdersResponse {
	this := CancelMultipleOptionOrdersResponse{}
	return &this
}

// NewCancelMultipleOptionOrdersResponseWithDefaults instantiates a new CancelMultipleOptionOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMultipleOptionOrdersResponseWithDefaults() *CancelMultipleOptionOrdersResponse {
	this := CancelMultipleOptionOrdersResponse{}
	return &this
}

func (o CancelMultipleOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMultipleOptionOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CancelMultipleOptionOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCancelMultipleOptionOrdersResponse struct {
	value CancelMultipleOptionOrdersResponse
	isSet bool
}

func (v NullableCancelMultipleOptionOrdersResponse) Get() CancelMultipleOptionOrdersResponse {
	return v.value
}

func (v *NullableCancelMultipleOptionOrdersResponse) Set(val CancelMultipleOptionOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMultipleOptionOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMultipleOptionOrdersResponse) Unset() {
	v.value = CancelMultipleOptionOrdersResponse{}
	v.isSet = false
}

func NewNullableCancelMultipleOptionOrdersResponse(val CancelMultipleOptionOrdersResponse) *NullableCancelMultipleOptionOrdersResponse {
	return &NullableCancelMultipleOptionOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelMultipleOptionOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMultipleOptionOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
