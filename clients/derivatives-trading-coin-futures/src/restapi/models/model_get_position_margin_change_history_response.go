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

// checks if the GetPositionMarginChangeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPositionMarginChangeHistoryResponse{}

// GetPositionMarginChangeHistoryResponse struct for GetPositionMarginChangeHistoryResponse
type GetPositionMarginChangeHistoryResponse struct {
	Items []GetPositionMarginChangeHistoryResponseInner
}

// NewGetPositionMarginChangeHistoryResponse instantiates a new GetPositionMarginChangeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPositionMarginChangeHistoryResponse() *GetPositionMarginChangeHistoryResponse {
	this := GetPositionMarginChangeHistoryResponse{}
	return &this
}

// NewGetPositionMarginChangeHistoryResponseWithDefaults instantiates a new GetPositionMarginChangeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPositionMarginChangeHistoryResponseWithDefaults() *GetPositionMarginChangeHistoryResponse {
	this := GetPositionMarginChangeHistoryResponse{}
	return &this
}

func (o GetPositionMarginChangeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPositionMarginChangeHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetPositionMarginChangeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetPositionMarginChangeHistoryResponse struct {
	value GetPositionMarginChangeHistoryResponse
	isSet bool
}

func (v NullableGetPositionMarginChangeHistoryResponse) Get() GetPositionMarginChangeHistoryResponse {
	return v.value
}

func (v *NullableGetPositionMarginChangeHistoryResponse) Set(val GetPositionMarginChangeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPositionMarginChangeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPositionMarginChangeHistoryResponse) Unset() {
	v.value = GetPositionMarginChangeHistoryResponse{}
	v.isSet = false
}

func NewNullableGetPositionMarginChangeHistoryResponse(val GetPositionMarginChangeHistoryResponse) *NullableGetPositionMarginChangeHistoryResponse {
	return &NullableGetPositionMarginChangeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetPositionMarginChangeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPositionMarginChangeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
