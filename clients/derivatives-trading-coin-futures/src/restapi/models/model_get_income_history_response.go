/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetIncomeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetIncomeHistoryResponse{}

// GetIncomeHistoryResponse struct for GetIncomeHistoryResponse
type GetIncomeHistoryResponse struct {
	Items []GetIncomeHistoryResponseInner
}

// NewGetIncomeHistoryResponse instantiates a new GetIncomeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIncomeHistoryResponse() *GetIncomeHistoryResponse {
	this := GetIncomeHistoryResponse{}
	return &this
}

// NewGetIncomeHistoryResponseWithDefaults instantiates a new GetIncomeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIncomeHistoryResponseWithDefaults() *GetIncomeHistoryResponse {
	this := GetIncomeHistoryResponse{}
	return &this
}

func (o GetIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIncomeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetIncomeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetIncomeHistoryResponse struct {
	value GetIncomeHistoryResponse
	isSet bool
}

func (v NullableGetIncomeHistoryResponse) Get() GetIncomeHistoryResponse {
	return v.value
}

func (v *NullableGetIncomeHistoryResponse) Set(val GetIncomeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeHistoryResponse) Unset() {
	v.value = GetIncomeHistoryResponse{}
	v.isSet = false
}

func NewNullableGetIncomeHistoryResponse(val GetIncomeHistoryResponse) *NullableGetIncomeHistoryResponse {
	return &NullableGetIncomeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetIncomeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
