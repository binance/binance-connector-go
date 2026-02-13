/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelMultipleOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelMultipleOrdersResponse{}

// CancelMultipleOrdersResponse struct for CancelMultipleOrdersResponse
type CancelMultipleOrdersResponse struct {
	Items []CancelMultipleOrdersResponseInner
}

// NewCancelMultipleOrdersResponse instantiates a new CancelMultipleOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelMultipleOrdersResponse() *CancelMultipleOrdersResponse {
	this := CancelMultipleOrdersResponse{}
	return &this
}

// NewCancelMultipleOrdersResponseWithDefaults instantiates a new CancelMultipleOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelMultipleOrdersResponseWithDefaults() *CancelMultipleOrdersResponse {
	this := CancelMultipleOrdersResponse{}
	return &this
}

func (o CancelMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelMultipleOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CancelMultipleOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCancelMultipleOrdersResponse struct {
	value CancelMultipleOrdersResponse
	isSet bool
}

func (v NullableCancelMultipleOrdersResponse) Get() CancelMultipleOrdersResponse {
	return v.value
}

func (v *NullableCancelMultipleOrdersResponse) Set(val CancelMultipleOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelMultipleOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelMultipleOrdersResponse) Unset() {
	v.value = CancelMultipleOrdersResponse{}
	v.isSet = false
}

func NewNullableCancelMultipleOrdersResponse(val CancelMultipleOrdersResponse) *NullableCancelMultipleOrdersResponse {
	return &NullableCancelMultipleOrdersResponse{value: val, isSet: true}
}

func (v NullableCancelMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelMultipleOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
