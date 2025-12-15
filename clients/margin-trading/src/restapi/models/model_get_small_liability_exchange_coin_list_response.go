/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSmallLiabilityExchangeCoinListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSmallLiabilityExchangeCoinListResponse{}

// GetSmallLiabilityExchangeCoinListResponse struct for GetSmallLiabilityExchangeCoinListResponse
type GetSmallLiabilityExchangeCoinListResponse struct {
	Items []GetSmallLiabilityExchangeCoinListResponseInner
}

// NewGetSmallLiabilityExchangeCoinListResponse instantiates a new GetSmallLiabilityExchangeCoinListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSmallLiabilityExchangeCoinListResponse() *GetSmallLiabilityExchangeCoinListResponse {
	this := GetSmallLiabilityExchangeCoinListResponse{}
	return &this
}

// NewGetSmallLiabilityExchangeCoinListResponseWithDefaults instantiates a new GetSmallLiabilityExchangeCoinListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSmallLiabilityExchangeCoinListResponseWithDefaults() *GetSmallLiabilityExchangeCoinListResponse {
	this := GetSmallLiabilityExchangeCoinListResponse{}
	return &this
}

func (o GetSmallLiabilityExchangeCoinListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSmallLiabilityExchangeCoinListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetSmallLiabilityExchangeCoinListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetSmallLiabilityExchangeCoinListResponse struct {
	value GetSmallLiabilityExchangeCoinListResponse
	isSet bool
}

func (v NullableGetSmallLiabilityExchangeCoinListResponse) Get() GetSmallLiabilityExchangeCoinListResponse {
	return v.value
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponse) Set(val GetSmallLiabilityExchangeCoinListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmallLiabilityExchangeCoinListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponse) Unset() {
	v.value = GetSmallLiabilityExchangeCoinListResponse{}
	v.isSet = false
}

func NewNullableGetSmallLiabilityExchangeCoinListResponse(val GetSmallLiabilityExchangeCoinListResponse) *NullableGetSmallLiabilityExchangeCoinListResponse {
	return &NullableGetSmallLiabilityExchangeCoinListResponse{value: val, isSet: true}
}

func (v NullableGetSmallLiabilityExchangeCoinListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmallLiabilityExchangeCoinListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
