/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DepositHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryResponse{}

// DepositHistoryResponse struct for DepositHistoryResponse
type DepositHistoryResponse struct {
	Items []DepositHistoryResponseInner
}

// NewDepositHistoryResponse instantiates a new DepositHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryResponse() *DepositHistoryResponse {
	this := DepositHistoryResponse{}
	return &this
}

// NewDepositHistoryResponseWithDefaults instantiates a new DepositHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryResponseWithDefaults() *DepositHistoryResponse {
	this := DepositHistoryResponse{}
	return &this
}

func (o DepositHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DepositHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDepositHistoryResponse struct {
	value DepositHistoryResponse
	isSet bool
}

func (v NullableDepositHistoryResponse) Get() DepositHistoryResponse {
	return v.value
}

func (v *NullableDepositHistoryResponse) Set(val DepositHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryResponse) Unset() {
	v.value = DepositHistoryResponse{}
	v.isSet = false
}

func NewNullableDepositHistoryResponse(val DepositHistoryResponse) *NullableDepositHistoryResponse {
	return &NullableDepositHistoryResponse{value: val, isSet: true}
}

func (v NullableDepositHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
