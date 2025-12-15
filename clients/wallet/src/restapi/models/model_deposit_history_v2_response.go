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

// checks if the DepositHistoryV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryV2Response{}

// DepositHistoryV2Response struct for DepositHistoryV2Response
type DepositHistoryV2Response struct {
	Items []DepositHistoryV2ResponseInner
}

// NewDepositHistoryV2Response instantiates a new DepositHistoryV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryV2Response() *DepositHistoryV2Response {
	this := DepositHistoryV2Response{}
	return &this
}

// NewDepositHistoryV2ResponseWithDefaults instantiates a new DepositHistoryV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryV2ResponseWithDefaults() *DepositHistoryV2Response {
	this := DepositHistoryV2Response{}
	return &this
}

func (o DepositHistoryV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DepositHistoryV2Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDepositHistoryV2Response struct {
	value DepositHistoryV2Response
	isSet bool
}

func (v NullableDepositHistoryV2Response) Get() DepositHistoryV2Response {
	return v.value
}

func (v *NullableDepositHistoryV2Response) Set(val DepositHistoryV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryV2Response) Unset() {
	v.value = DepositHistoryV2Response{}
	v.isSet = false
}

func NewNullableDepositHistoryV2Response(val DepositHistoryV2Response) *NullableDepositHistoryV2Response {
	return &NullableDepositHistoryV2Response{value: val, isSet: true}
}

func (v NullableDepositHistoryV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
