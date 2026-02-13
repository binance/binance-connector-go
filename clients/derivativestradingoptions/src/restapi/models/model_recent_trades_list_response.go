/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RecentTradesListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RecentTradesListResponse{}

// RecentTradesListResponse struct for RecentTradesListResponse
type RecentTradesListResponse struct {
	Items []RecentTradesListResponseInner
}

// NewRecentTradesListResponse instantiates a new RecentTradesListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentTradesListResponse() *RecentTradesListResponse {
	this := RecentTradesListResponse{}
	return &this
}

// NewRecentTradesListResponseWithDefaults instantiates a new RecentTradesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentTradesListResponseWithDefaults() *RecentTradesListResponse {
	this := RecentTradesListResponse{}
	return &this
}

func (o RecentTradesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentTradesListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RecentTradesListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRecentTradesListResponse struct {
	value RecentTradesListResponse
	isSet bool
}

func (v NullableRecentTradesListResponse) Get() RecentTradesListResponse {
	return v.value
}

func (v *NullableRecentTradesListResponse) Set(val RecentTradesListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentTradesListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentTradesListResponse) Unset() {
	v.value = RecentTradesListResponse{}
	v.isSet = false
}

func NewNullableRecentTradesListResponse(val RecentTradesListResponse) *NullableRecentTradesListResponse {
	return &NullableRecentTradesListResponse{value: val, isSet: true}
}

func (v NullableRecentTradesListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentTradesListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
