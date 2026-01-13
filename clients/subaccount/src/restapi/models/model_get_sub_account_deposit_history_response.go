/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSubAccountDepositHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSubAccountDepositHistoryResponse{}

// GetSubAccountDepositHistoryResponse struct for GetSubAccountDepositHistoryResponse
type GetSubAccountDepositHistoryResponse struct {
	Items []GetSubAccountDepositHistoryResponseInner
}

// NewGetSubAccountDepositHistoryResponse instantiates a new GetSubAccountDepositHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountDepositHistoryResponse() *GetSubAccountDepositHistoryResponse {
	this := GetSubAccountDepositHistoryResponse{}
	return &this
}

// NewGetSubAccountDepositHistoryResponseWithDefaults instantiates a new GetSubAccountDepositHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountDepositHistoryResponseWithDefaults() *GetSubAccountDepositHistoryResponse {
	this := GetSubAccountDepositHistoryResponse{}
	return &this
}

func (o GetSubAccountDepositHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountDepositHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetSubAccountDepositHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetSubAccountDepositHistoryResponse struct {
	value GetSubAccountDepositHistoryResponse
	isSet bool
}

func (v NullableGetSubAccountDepositHistoryResponse) Get() GetSubAccountDepositHistoryResponse {
	return v.value
}

func (v *NullableGetSubAccountDepositHistoryResponse) Set(val GetSubAccountDepositHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountDepositHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountDepositHistoryResponse) Unset() {
	v.value = GetSubAccountDepositHistoryResponse{}
	v.isSet = false
}

func NewNullableGetSubAccountDepositHistoryResponse(val GetSubAccountDepositHistoryResponse) *NullableGetSubAccountDepositHistoryResponse {
	return &NullableGetSubAccountDepositHistoryResponse{value: val, isSet: true}
}

func (v NullableGetSubAccountDepositHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountDepositHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
