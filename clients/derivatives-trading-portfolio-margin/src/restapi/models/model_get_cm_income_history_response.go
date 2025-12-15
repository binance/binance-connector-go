/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetCmIncomeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetCmIncomeHistoryResponse{}

// GetCmIncomeHistoryResponse struct for GetCmIncomeHistoryResponse
type GetCmIncomeHistoryResponse struct {
	Items []GetCmIncomeHistoryResponseInner
}

// NewGetCmIncomeHistoryResponse instantiates a new GetCmIncomeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCmIncomeHistoryResponse() *GetCmIncomeHistoryResponse {
	this := GetCmIncomeHistoryResponse{}
	return &this
}

// NewGetCmIncomeHistoryResponseWithDefaults instantiates a new GetCmIncomeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCmIncomeHistoryResponseWithDefaults() *GetCmIncomeHistoryResponse {
	this := GetCmIncomeHistoryResponse{}
	return &this
}

func (o GetCmIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCmIncomeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetCmIncomeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetCmIncomeHistoryResponse struct {
	value GetCmIncomeHistoryResponse
	isSet bool
}

func (v NullableGetCmIncomeHistoryResponse) Get() GetCmIncomeHistoryResponse {
	return v.value
}

func (v *NullableGetCmIncomeHistoryResponse) Set(val GetCmIncomeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCmIncomeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCmIncomeHistoryResponse) Unset() {
	v.value = GetCmIncomeHistoryResponse{}
	v.isSet = false
}

func NewNullableGetCmIncomeHistoryResponse(val GetCmIncomeHistoryResponse) *NullableGetCmIncomeHistoryResponse {
	return &NullableGetCmIncomeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetCmIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCmIncomeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
