/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetUmIncomeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUmIncomeHistoryResponse{}

// GetUmIncomeHistoryResponse struct for GetUmIncomeHistoryResponse
type GetUmIncomeHistoryResponse struct {
	Items []GetUmIncomeHistoryResponseInner
}

// NewGetUmIncomeHistoryResponse instantiates a new GetUmIncomeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUmIncomeHistoryResponse() *GetUmIncomeHistoryResponse {
	this := GetUmIncomeHistoryResponse{}
	return &this
}

// NewGetUmIncomeHistoryResponseWithDefaults instantiates a new GetUmIncomeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUmIncomeHistoryResponseWithDefaults() *GetUmIncomeHistoryResponse {
	this := GetUmIncomeHistoryResponse{}
	return &this
}

func (o GetUmIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUmIncomeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetUmIncomeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetUmIncomeHistoryResponse struct {
	value GetUmIncomeHistoryResponse
	isSet bool
}

func (v NullableGetUmIncomeHistoryResponse) Get() GetUmIncomeHistoryResponse {
	return v.value
}

func (v *NullableGetUmIncomeHistoryResponse) Set(val GetUmIncomeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUmIncomeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUmIncomeHistoryResponse) Unset() {
	v.value = GetUmIncomeHistoryResponse{}
	v.isSet = false
}

func NewNullableGetUmIncomeHistoryResponse(val GetUmIncomeHistoryResponse) *NullableGetUmIncomeHistoryResponse {
	return &NullableGetUmIncomeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetUmIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUmIncomeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
