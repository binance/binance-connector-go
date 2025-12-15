/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFundingRateHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateHistoryResponse{}

// GetFundingRateHistoryResponse struct for GetFundingRateHistoryResponse
type GetFundingRateHistoryResponse struct {
	Items []GetFundingRateHistoryResponseInner
}

// NewGetFundingRateHistoryResponse instantiates a new GetFundingRateHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateHistoryResponse() *GetFundingRateHistoryResponse {
	this := GetFundingRateHistoryResponse{}
	return &this
}

// NewGetFundingRateHistoryResponseWithDefaults instantiates a new GetFundingRateHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateHistoryResponseWithDefaults() *GetFundingRateHistoryResponse {
	this := GetFundingRateHistoryResponse{}
	return &this
}

func (o GetFundingRateHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetFundingRateHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetFundingRateHistoryResponse struct {
	value GetFundingRateHistoryResponse
	isSet bool
}

func (v NullableGetFundingRateHistoryResponse) Get() GetFundingRateHistoryResponse {
	return v.value
}

func (v *NullableGetFundingRateHistoryResponse) Set(val GetFundingRateHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateHistoryResponse) Unset() {
	v.value = GetFundingRateHistoryResponse{}
	v.isSet = false
}

func NewNullableGetFundingRateHistoryResponse(val GetFundingRateHistoryResponse) *NullableGetFundingRateHistoryResponse {
	return &NullableGetFundingRateHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFundingRateHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
