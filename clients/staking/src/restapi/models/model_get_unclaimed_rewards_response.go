/*
Binance Staking REST API

OpenAPI Specification for the Binance Staking REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetUnclaimedRewardsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetUnclaimedRewardsResponse{}

// GetUnclaimedRewardsResponse struct for GetUnclaimedRewardsResponse
type GetUnclaimedRewardsResponse struct {
	Items []GetUnclaimedRewardsResponseInner
}

// NewGetUnclaimedRewardsResponse instantiates a new GetUnclaimedRewardsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUnclaimedRewardsResponse() *GetUnclaimedRewardsResponse {
	this := GetUnclaimedRewardsResponse{}
	return &this
}

// NewGetUnclaimedRewardsResponseWithDefaults instantiates a new GetUnclaimedRewardsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUnclaimedRewardsResponseWithDefaults() *GetUnclaimedRewardsResponse {
	this := GetUnclaimedRewardsResponse{}
	return &this
}

func (o GetUnclaimedRewardsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUnclaimedRewardsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetUnclaimedRewardsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetUnclaimedRewardsResponse struct {
	value GetUnclaimedRewardsResponse
	isSet bool
}

func (v NullableGetUnclaimedRewardsResponse) Get() GetUnclaimedRewardsResponse {
	return v.value
}

func (v *NullableGetUnclaimedRewardsResponse) Set(val GetUnclaimedRewardsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUnclaimedRewardsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUnclaimedRewardsResponse) Unset() {
	v.value = GetUnclaimedRewardsResponse{}
	v.isSet = false
}

func NewNullableGetUnclaimedRewardsResponse(val GetUnclaimedRewardsResponse) *NullableGetUnclaimedRewardsResponse {
	return &NullableGetUnclaimedRewardsResponse{value: val, isSet: true}
}

func (v NullableGetUnclaimedRewardsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUnclaimedRewardsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
