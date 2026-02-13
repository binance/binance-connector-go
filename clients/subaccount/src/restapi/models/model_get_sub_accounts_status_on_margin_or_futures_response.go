/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSubAccountsStatusOnMarginOrFuturesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSubAccountsStatusOnMarginOrFuturesResponse{}

// GetSubAccountsStatusOnMarginOrFuturesResponse struct for GetSubAccountsStatusOnMarginOrFuturesResponse
type GetSubAccountsStatusOnMarginOrFuturesResponse struct {
	Items []GetSubAccountsStatusOnMarginOrFuturesResponseInner
}

// NewGetSubAccountsStatusOnMarginOrFuturesResponse instantiates a new GetSubAccountsStatusOnMarginOrFuturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountsStatusOnMarginOrFuturesResponse() *GetSubAccountsStatusOnMarginOrFuturesResponse {
	this := GetSubAccountsStatusOnMarginOrFuturesResponse{}
	return &this
}

// NewGetSubAccountsStatusOnMarginOrFuturesResponseWithDefaults instantiates a new GetSubAccountsStatusOnMarginOrFuturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountsStatusOnMarginOrFuturesResponseWithDefaults() *GetSubAccountsStatusOnMarginOrFuturesResponse {
	this := GetSubAccountsStatusOnMarginOrFuturesResponse{}
	return &this
}

func (o GetSubAccountsStatusOnMarginOrFuturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountsStatusOnMarginOrFuturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetSubAccountsStatusOnMarginOrFuturesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetSubAccountsStatusOnMarginOrFuturesResponse struct {
	value GetSubAccountsStatusOnMarginOrFuturesResponse
	isSet bool
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponse) Get() GetSubAccountsStatusOnMarginOrFuturesResponse {
	return v.value
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponse) Set(val GetSubAccountsStatusOnMarginOrFuturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponse) Unset() {
	v.value = GetSubAccountsStatusOnMarginOrFuturesResponse{}
	v.isSet = false
}

func NewNullableGetSubAccountsStatusOnMarginOrFuturesResponse(val GetSubAccountsStatusOnMarginOrFuturesResponse) *NullableGetSubAccountsStatusOnMarginOrFuturesResponse {
	return &NullableGetSubAccountsStatusOnMarginOrFuturesResponse{value: val, isSet: true}
}

func (v NullableGetSubAccountsStatusOnMarginOrFuturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountsStatusOnMarginOrFuturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
