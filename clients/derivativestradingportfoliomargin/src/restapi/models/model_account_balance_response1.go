/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountBalanceResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBalanceResponse1{}

// AccountBalanceResponse1 struct for AccountBalanceResponse1
type AccountBalanceResponse1 struct {
	Items []AccountBalanceResponse1Inner
}

// NewAccountBalanceResponse1 instantiates a new AccountBalanceResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBalanceResponse1() *AccountBalanceResponse1 {
	this := AccountBalanceResponse1{}
	return &this
}

// NewAccountBalanceResponse1WithDefaults instantiates a new AccountBalanceResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBalanceResponse1WithDefaults() *AccountBalanceResponse1 {
	this := AccountBalanceResponse1{}
	return &this
}

func (o AccountBalanceResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBalanceResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AccountBalanceResponse1) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAccountBalanceResponse1 struct {
	value AccountBalanceResponse1
	isSet bool
}

func (v NullableAccountBalanceResponse1) Get() AccountBalanceResponse1 {
	return v.value
}

func (v *NullableAccountBalanceResponse1) Set(val AccountBalanceResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBalanceResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBalanceResponse1) Unset() {
	v.value = AccountBalanceResponse1{}
	v.isSet = false
}

func NewNullableAccountBalanceResponse1(val AccountBalanceResponse1) *NullableAccountBalanceResponse1 {
	return &NullableAccountBalanceResponse1{value: val, isSet: true}
}

func (v NullableAccountBalanceResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBalanceResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
