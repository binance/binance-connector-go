/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawHistoryResponse{}

// WithdrawHistoryResponse struct for WithdrawHistoryResponse
type WithdrawHistoryResponse struct {
	Items []WithdrawHistoryResponseInner
}

// NewWithdrawHistoryResponse instantiates a new WithdrawHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawHistoryResponse() *WithdrawHistoryResponse {
	this := WithdrawHistoryResponse{}
	return &this
}

// NewWithdrawHistoryResponseWithDefaults instantiates a new WithdrawHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawHistoryResponseWithDefaults() *WithdrawHistoryResponse {
	this := WithdrawHistoryResponse{}
	return &this
}

func (o WithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *WithdrawHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableWithdrawHistoryResponse struct {
	value WithdrawHistoryResponse
	isSet bool
}

func (v NullableWithdrawHistoryResponse) Get() WithdrawHistoryResponse {
	return v.value
}

func (v *NullableWithdrawHistoryResponse) Set(val WithdrawHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawHistoryResponse) Unset() {
	v.value = WithdrawHistoryResponse{}
	v.isSet = false
}

func NewNullableWithdrawHistoryResponse(val WithdrawHistoryResponse) *NullableWithdrawHistoryResponse {
	return &NullableWithdrawHistoryResponse{value: val, isSet: true}
}

func (v NullableWithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
