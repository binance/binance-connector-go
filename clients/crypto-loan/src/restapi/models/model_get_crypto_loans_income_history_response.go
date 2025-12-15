/*
Binance Crypto Loan REST API

OpenAPI Specification for the Binance Crypto Loan REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCryptoLoansIncomeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCryptoLoansIncomeHistoryResponse{}

// GetCryptoLoansIncomeHistoryResponse struct for GetCryptoLoansIncomeHistoryResponse
type GetCryptoLoansIncomeHistoryResponse struct {
	Items []GetCryptoLoansIncomeHistoryResponseInner
}

// NewGetCryptoLoansIncomeHistoryResponse instantiates a new GetCryptoLoansIncomeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCryptoLoansIncomeHistoryResponse() *GetCryptoLoansIncomeHistoryResponse {
	this := GetCryptoLoansIncomeHistoryResponse{}
	return &this
}

// NewGetCryptoLoansIncomeHistoryResponseWithDefaults instantiates a new GetCryptoLoansIncomeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCryptoLoansIncomeHistoryResponseWithDefaults() *GetCryptoLoansIncomeHistoryResponse {
	this := GetCryptoLoansIncomeHistoryResponse{}
	return &this
}

func (o GetCryptoLoansIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCryptoLoansIncomeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetCryptoLoansIncomeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetCryptoLoansIncomeHistoryResponse struct {
	value GetCryptoLoansIncomeHistoryResponse
	isSet bool
}

func (v NullableGetCryptoLoansIncomeHistoryResponse) Get() GetCryptoLoansIncomeHistoryResponse {
	return v.value
}

func (v *NullableGetCryptoLoansIncomeHistoryResponse) Set(val GetCryptoLoansIncomeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCryptoLoansIncomeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCryptoLoansIncomeHistoryResponse) Unset() {
	v.value = GetCryptoLoansIncomeHistoryResponse{}
	v.isSet = false
}

func NewNullableGetCryptoLoansIncomeHistoryResponse(val GetCryptoLoansIncomeHistoryResponse) *NullableGetCryptoLoansIncomeHistoryResponse {
	return &NullableGetCryptoLoansIncomeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetCryptoLoansIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCryptoLoansIncomeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
