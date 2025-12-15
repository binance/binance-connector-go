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

// checks if the GetOrderModifyHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderModifyHistoryResponse{}

// GetOrderModifyHistoryResponse struct for GetOrderModifyHistoryResponse
type GetOrderModifyHistoryResponse struct {
	Items []GetOrderModifyHistoryResponseInner
}

// NewGetOrderModifyHistoryResponse instantiates a new GetOrderModifyHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderModifyHistoryResponse() *GetOrderModifyHistoryResponse {
	this := GetOrderModifyHistoryResponse{}
	return &this
}

// NewGetOrderModifyHistoryResponseWithDefaults instantiates a new GetOrderModifyHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderModifyHistoryResponseWithDefaults() *GetOrderModifyHistoryResponse {
	this := GetOrderModifyHistoryResponse{}
	return &this
}

func (o GetOrderModifyHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderModifyHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetOrderModifyHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetOrderModifyHistoryResponse struct {
	value GetOrderModifyHistoryResponse
	isSet bool
}

func (v NullableGetOrderModifyHistoryResponse) Get() GetOrderModifyHistoryResponse {
	return v.value
}

func (v *NullableGetOrderModifyHistoryResponse) Set(val GetOrderModifyHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderModifyHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderModifyHistoryResponse) Unset() {
	v.value = GetOrderModifyHistoryResponse{}
	v.isSet = false
}

func NewNullableGetOrderModifyHistoryResponse(val GetOrderModifyHistoryResponse) *NullableGetOrderModifyHistoryResponse {
	return &NullableGetOrderModifyHistoryResponse{value: val, isSet: true}
}

func (v NullableGetOrderModifyHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderModifyHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
