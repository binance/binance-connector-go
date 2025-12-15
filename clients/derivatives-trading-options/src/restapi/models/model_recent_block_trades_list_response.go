/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RecentBlockTradesListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecentBlockTradesListResponse{}

// RecentBlockTradesListResponse struct for RecentBlockTradesListResponse
type RecentBlockTradesListResponse struct {
	Items []RecentBlockTradesListResponseInner
}

// NewRecentBlockTradesListResponse instantiates a new RecentBlockTradesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentBlockTradesListResponse() *RecentBlockTradesListResponse {
	this := RecentBlockTradesListResponse{}
	return &this
}

// NewRecentBlockTradesListResponseWithDefaults instantiates a new RecentBlockTradesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentBlockTradesListResponseWithDefaults() *RecentBlockTradesListResponse {
	this := RecentBlockTradesListResponse{}
	return &this
}

func (o RecentBlockTradesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentBlockTradesListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RecentBlockTradesListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRecentBlockTradesListResponse struct {
	value RecentBlockTradesListResponse
	isSet bool
}

func (v NullableRecentBlockTradesListResponse) Get() RecentBlockTradesListResponse {
	return v.value
}

func (v *NullableRecentBlockTradesListResponse) Set(val RecentBlockTradesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentBlockTradesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentBlockTradesListResponse) Unset() {
	v.value = RecentBlockTradesListResponse{}
	v.isSet = false
}

func NewNullableRecentBlockTradesListResponse(val RecentBlockTradesListResponse) *NullableRecentBlockTradesListResponse {
	return &NullableRecentBlockTradesListResponse{value: val, isSet: true}
}

func (v NullableRecentBlockTradesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentBlockTradesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
