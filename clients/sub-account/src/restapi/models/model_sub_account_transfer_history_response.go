/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SubAccountTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubAccountTransferHistoryResponse{}

// SubAccountTransferHistoryResponse struct for SubAccountTransferHistoryResponse
type SubAccountTransferHistoryResponse struct {
	Items []SubAccountTransferHistoryResponseInner
}

// NewSubAccountTransferHistoryResponse instantiates a new SubAccountTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountTransferHistoryResponse() *SubAccountTransferHistoryResponse {
	this := SubAccountTransferHistoryResponse{}
	return &this
}

// NewSubAccountTransferHistoryResponseWithDefaults instantiates a new SubAccountTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountTransferHistoryResponseWithDefaults() *SubAccountTransferHistoryResponse {
	this := SubAccountTransferHistoryResponse{}
	return &this
}

func (o SubAccountTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccountTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SubAccountTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSubAccountTransferHistoryResponse struct {
	value SubAccountTransferHistoryResponse
	isSet bool
}

func (v NullableSubAccountTransferHistoryResponse) Get() SubAccountTransferHistoryResponse {
	return v.value
}

func (v *NullableSubAccountTransferHistoryResponse) Set(val SubAccountTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountTransferHistoryResponse) Unset() {
	v.value = SubAccountTransferHistoryResponse{}
	v.isSet = false
}

func NewNullableSubAccountTransferHistoryResponse(val SubAccountTransferHistoryResponse) *NullableSubAccountTransferHistoryResponse {
	return &NullableSubAccountTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableSubAccountTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
