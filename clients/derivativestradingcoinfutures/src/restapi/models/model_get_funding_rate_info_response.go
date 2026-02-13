/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFundingRateInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFundingRateInfoResponse{}

// GetFundingRateInfoResponse struct for GetFundingRateInfoResponse
type GetFundingRateInfoResponse struct {
	Items []GetFundingRateInfoResponseInner
}

// NewGetFundingRateInfoResponse instantiates a new GetFundingRateInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundingRateInfoResponse() *GetFundingRateInfoResponse {
	this := GetFundingRateInfoResponse{}
	return &this
}

// NewGetFundingRateInfoResponseWithDefaults instantiates a new GetFundingRateInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundingRateInfoResponseWithDefaults() *GetFundingRateInfoResponse {
	this := GetFundingRateInfoResponse{}
	return &this
}

func (o GetFundingRateInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFundingRateInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetFundingRateInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetFundingRateInfoResponse struct {
	value GetFundingRateInfoResponse
	isSet bool
}

func (v NullableGetFundingRateInfoResponse) Get() GetFundingRateInfoResponse {
	return v.value
}

func (v *NullableGetFundingRateInfoResponse) Set(val GetFundingRateInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundingRateInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundingRateInfoResponse) Unset() {
	v.value = GetFundingRateInfoResponse{}
	v.isSet = false
}

func NewNullableGetFundingRateInfoResponse(val GetFundingRateInfoResponse) *NullableGetFundingRateInfoResponse {
	return &NullableGetFundingRateInfoResponse{value: val, isSet: true}
}

func (v NullableGetFundingRateInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundingRateInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
