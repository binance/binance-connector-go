/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the WithdrawHistoryV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawHistoryV2Response{}

// WithdrawHistoryV2Response struct for WithdrawHistoryV2Response
type WithdrawHistoryV2Response struct {
	Items []WithdrawHistoryV2ResponseInner
}

// NewWithdrawHistoryV2Response instantiates a new WithdrawHistoryV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawHistoryV2Response() *WithdrawHistoryV2Response {
	this := WithdrawHistoryV2Response{}
	return &this
}

// NewWithdrawHistoryV2ResponseWithDefaults instantiates a new WithdrawHistoryV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawHistoryV2ResponseWithDefaults() *WithdrawHistoryV2Response {
	this := WithdrawHistoryV2Response{}
	return &this
}

func (o WithdrawHistoryV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawHistoryV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *WithdrawHistoryV2Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableWithdrawHistoryV2Response struct {
	value WithdrawHistoryV2Response
	isSet bool
}

func (v NullableWithdrawHistoryV2Response) Get() WithdrawHistoryV2Response {
	return v.value
}

func (v *NullableWithdrawHistoryV2Response) Set(val WithdrawHistoryV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawHistoryV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawHistoryV2Response) Unset() {
	v.value = WithdrawHistoryV2Response{}
	v.isSet = false
}

func NewNullableWithdrawHistoryV2Response(val WithdrawHistoryV2Response) *NullableWithdrawHistoryV2Response {
	return &NullableWithdrawHistoryV2Response{value: val, isSet: true}
}

func (v NullableWithdrawHistoryV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawHistoryV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
