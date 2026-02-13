/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawHistoryV1Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawHistoryV1Response{}

// WithdrawHistoryV1Response struct for WithdrawHistoryV1Response
type WithdrawHistoryV1Response struct {
	Items []WithdrawHistoryV2ResponseInner
}

// NewWithdrawHistoryV1Response instantiates a new WithdrawHistoryV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawHistoryV1Response() *WithdrawHistoryV1Response {
	this := WithdrawHistoryV1Response{}
	return &this
}

// NewWithdrawHistoryV1ResponseWithDefaults instantiates a new WithdrawHistoryV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawHistoryV1ResponseWithDefaults() *WithdrawHistoryV1Response {
	this := WithdrawHistoryV1Response{}
	return &this
}

func (o WithdrawHistoryV1Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawHistoryV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *WithdrawHistoryV1Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableWithdrawHistoryV1Response struct {
	value WithdrawHistoryV1Response
	isSet bool
}

func (v NullableWithdrawHistoryV1Response) Get() WithdrawHistoryV1Response {
	return v.value
}

func (v *NullableWithdrawHistoryV1Response) Set(val WithdrawHistoryV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawHistoryV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawHistoryV1Response) Unset() {
	v.value = WithdrawHistoryV1Response{}
	v.isSet = false
}

func NewNullableWithdrawHistoryV1Response(val WithdrawHistoryV1Response) *NullableWithdrawHistoryV1Response {
	return &NullableWithdrawHistoryV1Response{value: val, isSet: true}
}

func (v NullableWithdrawHistoryV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawHistoryV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
