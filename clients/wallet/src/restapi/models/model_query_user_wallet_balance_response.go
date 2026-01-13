/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserWalletBalanceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserWalletBalanceResponse{}

// QueryUserWalletBalanceResponse struct for QueryUserWalletBalanceResponse
type QueryUserWalletBalanceResponse struct {
	Items []QueryUserWalletBalanceResponseInner
}

// NewQueryUserWalletBalanceResponse instantiates a new QueryUserWalletBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserWalletBalanceResponse() *QueryUserWalletBalanceResponse {
	this := QueryUserWalletBalanceResponse{}
	return &this
}

// NewQueryUserWalletBalanceResponseWithDefaults instantiates a new QueryUserWalletBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserWalletBalanceResponseWithDefaults() *QueryUserWalletBalanceResponse {
	this := QueryUserWalletBalanceResponse{}
	return &this
}

func (o QueryUserWalletBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserWalletBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryUserWalletBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryUserWalletBalanceResponse struct {
	value QueryUserWalletBalanceResponse
	isSet bool
}

func (v NullableQueryUserWalletBalanceResponse) Get() QueryUserWalletBalanceResponse {
	return v.value
}

func (v *NullableQueryUserWalletBalanceResponse) Set(val QueryUserWalletBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserWalletBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserWalletBalanceResponse) Unset() {
	v.value = QueryUserWalletBalanceResponse{}
	v.isSet = false
}

func NewNullableQueryUserWalletBalanceResponse(val QueryUserWalletBalanceResponse) *NullableQueryUserWalletBalanceResponse {
	return &NullableQueryUserWalletBalanceResponse{value: val, isSet: true}
}

func (v NullableQueryUserWalletBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserWalletBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
